package main

import (
	"MiniDocker/cgroups/subsystem"
	"fmt"
	"github.com/urfave/cli"
)

var runCommand = cli.Command{
	Name:  "run",
	Usage: "Create a container with namespace and cgroup limit",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
		cli.StringFlag{
			Name:  "m",
			Usage: "memory limit",
		},
		cli.StringFlag{
			Name:  "cpushare",
			Usage: "cpushare limit",
		},
		cli.StringFlag{
			Name:  "cpuset",
			Usage: "cpuset limit",
		},
	},
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("missing container args")
		}

		tty := context.Bool("ti")

		res := &subsystem.ResourceConfig{
			MemoryLimit: context.String("m"),
			CpuSet:      context.String("cpuset"),
			CpuShare:    context.String("cpushare"),
		}

		//cmdArray 为容器运行后，执行的第一个命令信息
		//cmdArray[0] 为命令内容，后面的为命令参数
		var cmdArray []string
		for _, arg := range context.Args() {
			cmdArray := append(cmdArray, arg)
		}

		return nil
	},
}
