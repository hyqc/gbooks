# Redis

## 什么是Redis？
一个高性能的缓存数据库，支持字符串、列表、哈希、集合等类型，支持持久化到磁盘，支持Lua脚本、发布订阅、自动故障转移等。可用于热点数据缓存等场景。

## Redis与Memcached有什么区别？
回答这个问题前，我们至少要知道Redis和Memcached是什么，各自的使用场景，以及为什么。那么什么事Memcached呢？
### 什么事memcached？
memcached 是一个内存数据库，仅支持key-value类型，不支持数据持久化到磁盘。
那么对比redis和memcache可以得出以下结论：
- 相同点：
    - 数据都在内存中操作
    - 读写性能高
- 不同点：
    - redis支持持久化，memcached不支持
    - redis支持哈希、列表、集合等类型，memcache不支持
    - redis支持发布订阅，lua脚本等，memcached不支持等

## 为什么使用Redis作为Mysql缓存？
- 高性能：内存读写，性能高
- 高并发：性能高带来了高QPS，自然相对Mysql来说并发支持更高

## Redis数据结构
### SDS简单动态字符串
结构：
```
struct {
    len int // 字符串长度
    free int // 可用空间大小
    data []byte // 字节数组，二进制安全
}

```

### Hash
使用链表(链地址发)解决哈希冲突，相同的hash索引加入链表的头部，节点指针指向上一个节点

#### Hash结构
```c
typedef struct dictht {

    // 哈希表数组
    dictEntry **table;

    // 哈希表大小
    unsigned long size;

    // 哈希表大小掩码，用于计算索引值
    // 总是等于 size - 1
    unsigned long sizemask;

    // 该哈希表已有节点的数量
    unsigned long used;

} dictht;

typedef struct dictEntry {

    // 键
    void *key;

    // 值
    union {
        void *val;
        uint64_t u64;
        int64_t s64;
    } v;

    // 指向下个哈希表节点，形成链表
    struct dictEntry *next;

} dictEntry;

typedef struct dict {

    // 类型特定函数
    dictType *type;

    // 私有数据
    void *privdata;

    // 哈希表
    dictht ht[2];

    // rehash 索引
    // 当 rehash 不在进行时，值为 -1
    int rehashidx; /* rehashing not in progress if rehashidx == -1 */

} dict;
```
#### Hash冲突
不同的键，计算哈希值后得到相同的索引，此时该索引上已经存在其他的键，那么就出现了哈希冲突。redis的Hash是通过在索引后面加上一个单向链表解决的

#### rehash重新散列（渐进式rehash）
rehash：重新散列hash表，是为了将哈希的负载因子维持在一个合理的位置而做的处理。触发条件时hash表键太多或太少，此时需要对哈希表的大小做扩展或收缩。
流程：
- 计算新hash需要的大小
- 计算键的新hash索引并迁移到h1哈希表上，多次分配迁移
- 迁移完成后，将h0空间释放，并将h1设置为h0，重新生成空的h1为下一次扩缩准备。  

### 跳跃表skiplist（只有有序集合zset）
在一个节点中维持多个指向不同节点的指针，从而达到快速访问节点的目的。

### 整数集合intset
当一个集合元素不多，且都是整数时，redis机会使用整数集合作为集合的底层实现

## Redis数据类型


### string 字符串类型
数据结构为简单动态字符串SDS，及存储二进制数据、字符串长度熟悉。SDS是安全的，字符串拼接不会造成缓冲区溢出，因为拼接之前SDS会检测空间知否足够，不够会自动扩容。类似结构
```golang
type string struct {
    data unsafe.Pinter // []int
    len int
}
```


### list 列表
底层是快速链表quicklist(双压缩ziplist)


### hash 哈希表



## 场景问题

### Redis的list底层数据结构是什么

### Redis的hash表是如何解决hash冲突的

### 你都支付项目中为什么用list，而不用发布订阅？


### 你的支付项目中为什么用hash表存储订单状态信息？有没有考虑过数据量超出hash的容量范围的问题？

### 你的支付项目中是如何生成订单号的？
雪花算法

### 多节点中你是如何保证队列消费安全的？
使用redis的 redis.call方法调用lua脚本实现的原子加锁和释放锁，创建全局唯一锁，哪个节点拿到这个锁了哪个节点来执行消费，消费完毕后在释放掉这个锁。

#### 那如果并发量上来了，其他节点不就成了摆设了吗？现在需要你提升这个服务的性能，怎么设计？
将队列进行拆分，有几个节点拆分成几个队列，生产者写入时用轮询写入保证每个节点的队列数据数量基本一致，消费时，各个节点只消费自己这个节点的队列数据。考虑极端情况可以在增加一个全局的调度器和全局队列来协调各个队列的消费情况，避免某一个节点挂掉后这个节点的队列数据停止消费问题。