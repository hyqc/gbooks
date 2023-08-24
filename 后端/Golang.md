# Golang基础总结
## 1. make 和 new 的区别
    make和new的定义：
    ```golang
    func make(t Type,size ...IntegerType) Type
    
    func new(Type) *Type
    ```
    区别：
    - 返回值类型不同，make返回类型值，new返回类型的指针
    - 参数不同，make只能用于slice,map,chanel的初始化，new可以用于任意类型
    - make是分配内存并初始化，new是只分配内存并返回类型的零值指针

## 2. 类型断言使用条件
![示例代码](./imgs/go-interface-type.png)
    类型断言只能作用于interface类型的变量，其他类型使用断言会报语法错误。
    


## 进阶

### 切片底层
切片的结构：
type slice struct {
  p unsafe.Pointer
  len int
  size int
}

### map结构底层

### Golang内存模型
golang的内存模型阐述了一个go程对某个变量的写操作，满足那些条件才能确保被另一个读取该变量的go程检测到。内存模型的本质就是限定写事件和读事件发生的顺序需要满足的条件：
1. 对某变量的写W操作必须发生在该变量的读R操作之前。（保证多go程对同一个变量的读写操作有序）
2. 在该写W操作之后和该读R操作之前，没有其他go程对该变量进行写W操作。（确保该变量的读和写唯一）

		这两个条件限定了，对同一个变量的读和写有序且唯一
		
		
### 如何分析Go项目的cpu、内存、竞态竞争等问题？
使用net包的pprof库，在代码中引入后，开启监听，在浏览器中打开debug路径，会有统计的信息列表：
allocs：查看过去所有内存分配的样本。
block：查看导致阻塞同步的堆栈跟踪。
cmdline： 当前程序的命令行的完整调用路径。
goroutine：查看当前所有运行的 goroutines 堆栈跟踪。
heap：查看活动对象的内存分配情况。
mutex：查看导致互斥锁的竞争持有者的堆栈跟踪。
profile： 默认进行 30s 的 CPU Profiling，得到一个分析用的 profile 文件。
threadcreate：查看创建新 OS 线程的堆栈跟踪。
trace：

runtime.SetBlockProfileRate(1) // 开启对阻塞操作的跟踪，block  
runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪，mutex

### GO pprof分析
pprof是用于可视化和性能分析的工具，pprof以profile.proto读取分析样本集合，并生成报告。使用方式如下：
- runtime/pprof：采集程序（非server）的指定区块的运行数据进行分析
- net/http/pprof：基于http server运行，可以采集运行时数据进行分析
- go test：通过运行测试用例，并指定所需表示来进行采集
支持三方使用方式：
- web界面
- 生成报告
- 交互式命令行
#### 可以做什么
- cpu profiling：cpu分析，按照一定频率采集监听的应用程序cpu使用情况，可以确定应用程序在主动消耗cpu周期时那些代码花费的时机多少
- memory profiling：内存分析，在应用程序进行堆分配时记录堆栈跟踪，用于监视当前和历史内存使用情况，以检查内存泄漏
- block profiling：阻塞分析，记录goroutine阻塞等待同步的位置，默认不开启，需要调用runtime.SetBlockProfileRate进行设置
- mutex profiling：互斥锁分析，报告互斥锁的竞争情况，默认不开启，需要调用runtime.SetMutextProfileFraction进行设置
- goroutine profiling：goroutine分析，可以对当前应用程序正在运行的goroutine进行堆栈跟踪分析。这项功能在实际排查中经常用到，因为很多问题出现时的表现就是goroutine暴增，而这时候我们要做的事情之一就是查看应用程序中goroutine正在做什么事情，因为什么阻塞了，然后在进行下一步分析。

#### xxx:xx/debug/pprof/xxx?debug=1（直接在浏览器访问）或0（下载对应的profile文件）
- allocs：查看过去所有内存分片的样本，访问路径/debug/pprof/allocs
- block：查看导致阻塞同步的堆栈跟踪，访问路径/debug/pprof/block
- cmdline：查看当前程序的命令行完整调用路径
- goroutine：查看当前所有运行的goroutine堆栈跟踪，访问路径：/debug/pprof/goroutine
- heap：查看活动对象的内存分配情况，访问路径：/debug/pprof/heap
- mutex：查看导致互斥锁的竞争持有者的堆栈跟踪，访问路径：/debug/pprof/mutex
- profile：默认进行30s的cpu分析，得到一个分析用的profile文件，访问路径：/debug/pprof/profile
- threadcreate：查看创建新OS线程的堆栈跟踪，访问路径：/debug/pprof/threadcreate

### 可视化查看指定profile文件分析结果
go tool pprof -http=:8889 profile

### 命令行查看
go tool pprof profile
