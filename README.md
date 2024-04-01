# miniDocker
miniDocker实现Docker基本功能

Docker本质其实是一个特殊的进程，这个进程被`Linux`中`Namespace`和`Cgroup`技术做了装饰，`NameSpace`将该进程与`Linux`系统进行隔离开来，而`Cgroup`则对该进程做了一系列的资源限制，两者配合模拟出来一个沙盒的环境。



## 创建容器过程

1. 初始化runCommand和initCommand命令
   - runCommand命令中为run命令，包含ti（是否前台显示）、memory（内存设置）、cpu（cpu设置）、cpuset（cpuset 用于限制一组进程只运行在特定的 cpu 节点上和只在特定的 mem 节点上分配内存）等参数。
   - initCommand命令中为init命令。

![img.png](doc/img.png)