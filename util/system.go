package util

import (
	"sync"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

// 获取系统信息
var (
	syOnce sync.Once
	Sy     SystemStruct
)

// 获取系统信息结构体
type SystemStruct struct{}

func CreateSystem() SystemStruct {
	syOnce.Do(func() {
		Sy = SystemStruct{}
	})
	return Sy
}

func (s *SystemStruct) GetSystemInfo() map[string]System {
	// 储存 数据
	var mapData map[string]System = make(map[string]System, 7)

	var system System

	system = &SystemInfo{}
	system.GetInfo()
	mapData["system"] = system

	system = &LoadedInfo{}
	system.GetInfo()
	mapData["load"] = system

	system = &CPUInfo{}
	system.GetInfo()
	mapData["cpu"] = system

	system = &MemoryInfo{}
	system.GetInfo()
	mapData["memory"] = system

	system = &SwapInfo{}
	system.GetInfo()
	mapData["swap"] = system

	system = &DiskInfo{}
	system.GetInfo()
	mapData["dis"] = system

	system = &NetInfo{}
	system.GetInfo()
	mapData["net"] = system

	return mapData
}

// system 接口
type System interface {
	GetInfo()
}

// 系统结构体
type SystemInfo struct {
	// 系统名称
	SystemName string
	// 系统 架构
	SA string
}

// 获取系统信息
func (s *SystemInfo) GetInfo() {
	// 判断是否为空，如果为空则获取信息
	if s.SystemName == "" || s.SA == "" {
		// 获取系统名字
		s.SA, _ = host.KernelArch()
		// 获取系统运行的平台
		s.SystemName, _, _, _ = host.PlatformInformation()
	}
}

// 负载状态结构体
type LoadedInfo struct {
	// 负载状态
	Load *load.AvgStat
}

// 获取系统负载状态
func (l *LoadedInfo) GetInfo() {
	// 获取系统负载状态
	l.Load, _ = load.Avg()
}

// cpu信息 结构体
type CPUInfo struct {
	// cpu名字
	ModelName string
	// 处理器核心数
	Core int
	// 处理器核心数 含逻辑处理器
	CoreLogic int
	// cpu占用 率 100% 多个cpu只取第一个
	Used float64
}

// 获取cpu状态
func (c *CPUInfo) GetInfo() {
	if c.ModelName == "" || c.Core == 0 || c.CoreLogic == 0 {
		cpuI, _ := cpu.Info()
		c.ModelName = cpuI[0].ModelName
		c.Core, _ = cpu.Counts(false)
		c.CoreLogic, _ = cpu.Counts(true)
	}

	u, _ := cpu.Percent(0, false)
	c.Used = u[0]
}

// 内存信息
type MemoryInfo struct {
	// 总内存
	Total uint64
	// 使用的内存
	Used uint64
	// 使用的内存 百分比
	UsedPercent float64
}

// 获取内存信息
func (m *MemoryInfo) GetInfo() {
	v, _ := mem.VirtualMemory()
	if m.Total == 0 {
		m.Total = v.Total
	}

	m.Used = v.Used
	m.UsedPercent = v.UsedPercent
}

// swap信息
type SwapInfo struct {
	// 总内存
	Total uint64
	// 使用的内存
	Used uint64
	// 使用的内存 百分比
	UsedPercent float64
}

// 获取swap信息
func (s *SwapInfo) GetInfo() {
	v, _ := mem.SwapMemory()
	if s.Total == 0 {
		s.Total = v.Total
	}

	s.Used = v.Used
	s.UsedPercent = v.UsedPercent
}

// 硬盘信息  只获取系统盘
type DiskInfo struct {
	// 硬盘总容量
	Total uint64
	// 使用的容量
	Used uint64
	// 使用的容量
	UsedPercent float64
}

// 获取硬盘信息
func (d *DiskInfo) GetInfo() {
	ds, _ := disk.Partitions(false)
	dsUsed, _ := disk.Usage(ds[0].Mountpoint)

	if d.Total == 0 {
		d.Total = dsUsed.Total
	}

	d.Used = dsUsed.Used
	d.UsedPercent = dsUsed.UsedPercent
}

// 网络信息
type NetInfo struct {
	// 接收的字节数
	BytesRecv uint64
	// 发送的字节数
	BytesSent uint64
}

func (ni *NetInfo) GetInfo() {
	// 获取所有网卡
	n, _ := net.IOCounters(true)

	// 遍历所有网卡拿到每个网卡的发送和接收的字节数
	for _, v := range n {
		ni.BytesRecv += v.BytesRecv
		ni.BytesSent += v.BytesSent
	}
}
