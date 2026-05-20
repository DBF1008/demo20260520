package controller

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/library/libUtils"
	"os"
	"runtime"
	"strconv"
	"time"
)

var Monitor = sysMonitorController{
	startTime: gtime.Now(),
}

type sysMonitorController struct {
	BaseController
	startTime *gtime.Time
}

func (c *sysMonitorController) List(ctx context.Context, req *system.MonitorSearchReq) (res *system.MonitorSearchRes, err error) {
	cpuNum := runtime.NumCPU()
	var cpuUsed float64 = 0
	var cpuAvg5 float64 = 0
	var cpuAvg15 float64 = 0

	cpuInfo, err := cpu.Percent(time.Duration(time.Second), false)
	if err == nil {
		cpuUsed, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", cpuInfo[0]), 64)
	}

	loadInfo, err := load.Avg()
	if err == nil {
		cpuAvg5, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", loadInfo.Load5), 64)
		cpuAvg15, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", loadInfo.Load5), 64)
	}

	var memTotal uint64 = 0
	var memUsed uint64 = 0
	var memFree uint64 = 0
	var memUsage float64 = 0

	v, err := mem.VirtualMemory()
	if err == nil {
		memTotal = v.Total
		memUsed = v.Used
		memFree = memTotal - memUsed
		memUsage, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", v.UsedPercent), 64)
	}

	var goTotal uint64 = 0
	var goUsed uint64 = 0
	var goFree uint64 = 0
	var goUsage float64 = 0

	p, err := process.NewProcess(int32(os.Getpid()))
	if err == nil {
		memInfo, err := p.MemoryInfo()
		if err == nil {
			goUsed = memInfo.RSS
			goUsage = gconv.Float64(fmt.Sprintf("%.2f", gconv.Float64(goUsed)/gconv.Float64(memTotal)*100))
		}
	}

	sysComputerIp := ""
	ip, err := libUtils.GetLocalIP()
	if err == nil {
		sysComputerIp = ip
	}

	sysComputerName := ""
	sysOsName := ""
	sysOsArch := ""

	sysInfo, err := host.Info()

	if err == nil {
		sysComputerName = sysInfo.Hostname
		sysOsName = sysInfo.OS
		sysOsArch = sysInfo.KernelArch
	}

	goName := "GoLang"
	goVersion := runtime.Version()
	gtime.Date()
	goStartTime := c.startTime

	goRunTime := gtime.Now().Timestamp() - c.startTime.Timestamp()
	goHome := runtime.GOROOT()
	goUserDir := ""

	curDir, err := os.Getwd()

	if err == nil {
		goUserDir = curDir
	}


	diskList := make([]disk.UsageStat, 0)
	diskInfo, err := disk.Partitions(true)
	if err == nil {
		for _, p := range diskInfo {
			diskDetail, err := disk.Usage(p.Mountpoint)
			if err == nil {
				diskDetail.UsedPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", diskDetail.UsedPercent), 64)
				diskList = append(diskList, *diskDetail)
			}
		}
	}
	res = new(system.MonitorSearchRes)
	res = &system.MonitorSearchRes{
		"cpuNum":          cpuNum,
		"cpuUsed":         cpuUsed,
		"cpuAvg5":         gconv.String(cpuAvg5),
		"cpuAvg15":        gconv.String(cpuAvg15),
		"memTotal":        memTotal,
		"goTotal":         goTotal,
		"memUsed":         memUsed,
		"goUsed":          goUsed,
		"memFree":         memFree,
		"goFree":          goFree,
		"memUsage":        memUsage,
		"goUsage":         goUsage,
		"sysComputerName": sysComputerName,
		"sysOsName":       sysOsName,
		"sysComputerIp":   sysComputerIp,
		"sysOsArch":       sysOsArch,
		"goName":          goName,
		"goVersion":       goVersion,
		"goStartTime":     goStartTime,
		"goRunTime":       goRunTime,
		"goHome":          goHome,
		"goUserDir":       goUserDir,
		"diskList":        diskList,
	}
	return
}
