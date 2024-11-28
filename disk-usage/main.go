package main

import (
	"log"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"

	"github.com/StackExchange/wmi"
)

//type diskPerformance struct {
//	BytesRead           int64
//	BytesWritten        int64
//	ReadTime            int64
//	WriteTime           int64
//	IdleTime            int64
//	ReadCount           uint32
//	WriteCount          uint32
//	QueueDepth          uint32
//	SplitCount          uint32
//	QueryTime           int64
//	StorageDeviceNumber uint32
//	StorageManagerName  [8]uint16
//	alignmentPadding    uint32 // necessary for 32bit support, see https://github.com/elastic/beats/pull/16553
//}

type Win32_PerfRawData_PerfDisk_LogicalDisk struct {
	AvgDiskBytesPerRead          uint64
	AvgDiskBytesPerRead_Base     uint32
	AvgDiskBytesPerTransfer      uint64
	AvgDiskBytesPerTransfer_Base uint32
	AvgDiskBytesPerWrite         uint64
	AvgDiskBytesPerWrite_Base    uint32
	AvgDiskQueueLength           uint64
	AvgDiskReadQueueLength       uint64
	AvgDiskSecPerRead            uint32
	AvgDiskSecPerRead_Base       uint32
	AvgDiskSecPerTransfer        uint32
	AvgDiskSecPerTransfer_Base   uint32
	AvgDiskSecPerWrite           uint32
	AvgDiskSecPerWrite_Base      uint32
	AvgDiskWriteQueueLength      uint64
	Caption                      *string
	CurrentDiskQueueLength       uint32
	Description                  *string
	DiskBytesPerSec              uint64
	DiskReadBytesPerSec          uint64
	DiskReadsPerSec              uint32
	DiskTransfersPerSec          uint32
	DiskWriteBytesPerSec         uint64
	DiskWritesPerSec             uint32
	FreeMegabytes                uint32
	Frequency_Object             uint64
	Frequency_PerfTime           uint64
	Frequency_Sys100NS           uint64
	Name                         string
	PercentDiskReadTime          uint64
	PercentDiskReadTime_Base     uint64
	PercentDiskTime              uint64
	PercentDiskTime_Base         uint64
	PercentDiskWriteTime         uint64
	PercentDiskWriteTime_Base    uint64
	PercentFreeSpace             uint32
	PercentFreeSpace_Base        uint32
	PercentIdleTime              uint64
	PercentIdleTime_Base         uint64
	SplitIOPerSec                uint32
	Timestamp_Object             uint64
	Timestamp_PerfTime           uint64
	Timestamp_Sys100NS           uint64
}

type Win32_PerfFormattedData_PerfDisk_PhysicalDisk struct {
	AvgDiskBytesPerRead     uint64
	AvgDiskBytesPerTransfer uint64
	AvgDiskBytesPerWrite    uint64
	AvgDiskQueueLength      uint64
	AvgDiskReadQueueLength  uint64
	AvgDiskSecPerRead       uint32
	AvgDiskSecPerTransfer   uint32
	AvgDiskSecPerWrite      uint32
	AvgDiskWriteQueueLength uint64
	Caption                 string
	CurrentDiskQueueLength  uint32
	Description             string
	DiskBytesPerSec         uint64
	DiskReadBytesPerSec     uint64
	DiskReadsPerSec         uint32
	DiskTransfersPerSec     uint32
	DiskWriteBytesPerSec    uint64
	DiskWritesPerSec        uint32
	Frequency_Object        uint64
	Frequency_PerfTime      uint64
	Frequency_Sys100NS      uint64
	Name                    string
	PercentDiskReadTime     uint64
	PercentDiskTime         uint64
	PercentDiskWriteTime    uint64
	PercentIdleTime         uint64
	SplitIOPerSec           uint32
	Timestamp_Object        uint64
	Timestamp_PerfTime      uint64
	Timestamp_Sys100NS      uint64
}

func main() {

	var dst []Win32_PerfFormattedData_PerfDisk_PhysicalDisk
	q := wmi.CreateQuery(&dst, "")
	for {
		// * CPU
		perc, err := cpu.Percent(1*time.Second, false)
		if err != nil {
			log.Printf("CPU(%%) Error: %v\n", err)
		} else {
			if len(perc) > 0 {
				log.Printf("CPU Usage: %.1f%%\n", perc[0])
			} else {
				log.Printf("CPU Usage: Empty perc\n")
			}
		}

		// * Memory
		v, _ := mem.VirtualMemory()
		log.Printf("Memory Usage: %v%%\n", v.UsedPercent)
		// almost every return value is a struct
		// log.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
		// convert to JSON. String() is also implemented
		// log.Println(v)

		// * Disk
		//diskPart, _ := disk.Partitions(true)
		//for _, dp := range diskPart {
		//	log.Println(dp)
		//	diskUsed, _ := disk.Usage(dp.Mountpoint)
		//	log.Printf("Partition Size: %d MB", diskUsed.Total/1024/1024)
		//	log.Printf("Partition Usage: %.1f%%", diskUsed.UsedPercent)
		//	log.Printf("Partition inode Usage: %.1f%%", diskUsed.InodesUsedPercent)
		//}

		//log.Println("Disk Usage:")
		//stat, err := disk.IOCounters()
		//if err != nil {
		//	log.Printf("Disk IO error: %v\n", err)
		//} else {
		//	for k, v := range stat {
		//		log.Printf("%s --- %+v\n", k, v)
		//	}
		//}

		// * WMI - PerfDisk
		//var dst []Win32_PerfRawData_PerfDisk_LogicalDisk
		err = wmi.Query(q, &dst)
		if err != nil {
			log.Println("ERROR disk", err)
		}
		//log.Printf("WMI Disk: %+v\n", dst)
		for _, pd := range dst {
			if pd.Name == "_Total" {
				activityTime := 100.0
				if pd.PercentIdleTime != 0 {
					activityTime = float64(pd.PercentDiskTime) / float64(pd.PercentDiskTime+pd.PercentIdleTime) * 100.0
				}
				log.Printf("Disk Time: %.1f%%\n", activityTime)
				break
			}
		}

		// *

		//var diskPerformance diskPerformance

		//lpBuffer := make([]uint16, 254)
		//lpBufferLen, err := windows.GetLogicalDriveStrings(uint32(len(lpBuffer)), &lpBuffer[0])
		//if err != nil {
		//	log.Println("xxx GetLogicalDriveStrings error:", err)
		//}
		//for _, v := range lpBuffer[:lpBufferLen] {
		//	if 'A' <= v && v <= 'Z' {
		//		path := string(rune(v)) + ":"
		//		typepath, _ := windows.UTF16PtrFromString(path)
		//		typeret := windows.GetDriveType(typepath)
		//		if typeret == 0 {
		//			log.Println(windows.GetLastError())
		//			return
		//		}
		//		if typeret != windows.DRIVE_FIXED {
		//			continue
		//		}
		//		szDevice := log.Sprintf(`\\.\%s`, path)
		//		const IOCTL_DISK_PERFORMANCE = 0x70020
		//		h, err := windows.CreateFile(syscall.StringToUTF16Ptr(szDevice), 0, windows.FILE_SHARE_READ|windows.FILE_SHARE_WRITE, nil, windows.OPEN_EXISTING, 0, 0)
		//		if err != nil {
		//			if err == windows.ERROR_FILE_NOT_FOUND {
		//				continue
		//			}
		//			log.Println(windows.GetLastError())
		//			return
		//		}
		//		defer windows.CloseHandle(h)

		//		var diskPerformanceSize uint32
		//		err = windows.DeviceIoControl(h, IOCTL_DISK_PERFORMANCE, nil, 0, (*byte)(unsafe.Pointer(&diskPerformance)), uint32(unsafe.Sizeof(diskPerformance)), &diskPerformanceSize, nil)
		//		if err != nil {
		//			log.Println(err)
		//		}

		//		log.Printf("Disk Performance: %+v\n", diskPerformance)

		//		//ReadBytes := uint64(diskPerformance.BytesRead)
		//		//WriteBytes := uint64(diskPerformance.BytesWritten)
		//		//ReadCount := uint64(diskPerformance.ReadCount)
		//		//WriteCount := uint64(diskPerformance.WriteCount)
		//		//ReadTime := uint64(diskPerformance.ReadTime / 10000 / 1000) // convert to ms: https://github.com/giampaolo/psutil/issues/1012
		//		//WriteTime := uint64(diskPerformance.WriteTime / 10000 / 1000)
		//		//IdleTime := uint64(diskPerformance.IdleTime / 10000 / 1000)

		//		//log.Printf("ReadBytes: %d\n", ReadBytes)
		//		//log.Printf("WriteBytes: %d\n", WriteBytes)
		//		//log.Printf("ReadCount: %d\n", ReadCount)
		//		//log.Printf("WriteCount: %d\n", WriteCount)
		//		//log.Printf("ReadTime: %d\n", ReadTime)
		//		//log.Printf("WriteTime: %d\n", WriteTime)
		//		//log.Printf("IdleTime: %d\n", IdleTime)
		//		//log.Printf("ActiveTime: %.1f%%\n", float64(ReadTime+WriteTime)/float64(IdleTime)*100)
		//	}

		//}

		log.Println()
	}

}
