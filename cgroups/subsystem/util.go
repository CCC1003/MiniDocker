package subsystem

// 获取cgroup在文件系统中的绝对路径
func GetCgroupPath(subsystem string, cgroupPath string, autoCreate bool) (string, error) {

}

// 找到了挂载了 subsystem 的hierarchy树 cgroup根节点所在的目录
func findCgroupMountPoint(subsystem string) (string, error) {

}
