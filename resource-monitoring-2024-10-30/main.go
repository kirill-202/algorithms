package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/cpu"
)

type CPU struct {

}

func main() {

    info, _ := cpu.Info()
    fmt.Printf("CPU: %v, Cores: %v, Mhz: %v\n", info[0].ModelName, len(info), info[0].Mhz)

    usage, _ := cpu.Percent(0, false)
    fmt.Printf("CPU usage: %.2f%%\n", usage)

}
