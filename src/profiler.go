package profiler

import (
	"runtime"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type SystemInfo struct {
	Disks          []DiskInfo
	OS             string
	CPUModel       string
	RAMModel       string
	RAMSize        uint64
	CPUTemperature float64
}

type DiskInfo struct {
	Name      string
	Size      uint64
	FreeSpace uint64
}

func GetSystemInfo() (*SystemInfo, error) {
	sysInfo := &SystemInfo{}

	// Get disks info
	partitions, err := disk.Partitions(true)
	if err != nil {
		return nil, err
	}

	for _, partition := range partitions {
		diskInfo := DiskInfo{Name: partition.Device}

		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			return nil, err
		}

		diskInfo.Size = usage.Total
		diskInfo.FreeSpace = usage.Free

		sysInfo.Disks = append(sysInfo.Disks, diskInfo)
	}

	// Get OS info
	sysInfo.OS = runtime.GOOS

	// Get CPU info
	cpuInfo, err := cpu.Info()
	if err != nil {
		return nil, err
	}
	sysInfo.CPUModel = cpuInfo[0].ModelName

	// Get RAM info
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	sysInfo.RAMModel = memInfo.String()
	sysInfo.RAMSize = memInfo.Total

	// Get CPU temperature info
	tempInfo, err := host.SensorsTemperatures()
	if err != nil {
		return nil, err
	}
	if len(tempInfo) > 0 {
		sysInfo.CPUTemperature = tempInfo[0].Temperature
	}

	return sysInfo, nil
}
