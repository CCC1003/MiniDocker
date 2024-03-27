# miniDocker
miniDocker实现Docker基本功能

Docker本质其实是一个特殊的进程，这个进程被`Linux`中`Namespace`和`Cgroup`技术做了装饰，`NameSpace`将该进程与`Linux`系统进行隔离开来，而`Cgroup`则对该进程做了一系列的资源限制，两者配合模拟出来一个沙盒的环境。


![img.png](img.png)