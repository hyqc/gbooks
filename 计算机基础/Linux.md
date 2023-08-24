# Linux

## 配置ssh免密登录远程服务器

### 本地生成登录秘钥
打开gitbash，执行：
```bash
ssh-keygen
```
进入目录：~/.ssh：
![.ssh目录](./img/ssh-dir.png)
文件说明：
- id_rsa：私钥
- id_rsa.put：公钥

### 上传ssh公钥到远程服务器
```bash
ssh-copy-id -i ~/.ssh/id_rsa.pub root@xxx.xxx.xxx.xxx
```

### 执行免密登录
```bash
ssh root@xxx.xxx.xxx.xxx
```

### 解决长时间不操作ssh连接断开问题
修改参数：/etc/ssh/sshd_config配置：
```bash
#ClientAliveInterval 0 间隔多少秒向客户端发送请求保活
#ClientAliveCountMax 3 服务器向客户端发送请求后没有收到客户端响应的次数最多值后端口连接

改为：
ClientAliveInterval 60
ClientAliveCountMax 3
```
修改完毕后，执行：
```bash
systemctl restart sshd
```

## namespace
Linux namespace提供了一种内核级别隔离系统资源的方法，通过将系统的全局资源放在不同的namespace中，来实现资源隔离的目的。不同的namespace程序，可以享有一份独立的系统资源。目前linux中提供了六类系统资源的隔离机制，分别是：
- mount：隔离文件系统挂载点
- uts：隔离主机名和域名信息
- ipc：隔离进程通信
- pid：隔离进程id
- newwork：隔离网络资源
- user：隔离用户和用户组的ID

## cgroup（control groups）
是linux内核提供的一种可以限制单个或多个进程所使用资源的机制，可以对cpu、内存等资源做精细化控制。cgroups为没中可以控制的资源定义了一个子系统：
- cpu子系统，主要限制进程的cpu使用率
- cpuacct子系统，可以统计cgroups中的进程的cpu使用情况
- cpuset子系统，可以为cgroups中的进程分配单独的cou节点或内存节点
- memory子系统，可以限制进程的内存使用量
- blkio子系统，可以限制进程的块设备io
- devices子系统，可以控制进程能欧访问某些设备
- net_cls子系统，可以标记cgroups中进程的网络数据包，然后使用tc模块对数据包进行控制
- freezer子系统，可以挂起或恢复cgroups中的进程
- ns子系统，可以使不同cgroups下面的进程使用不同的namespace