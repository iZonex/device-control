package status

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

// DeviceRegistered when device was registered
type DeviceRegistered struct {
	Status bool  `json:"status"`
	Date   int64 `json:"date"`
}

// NetworkInterface device information
type NetworkInterface struct {
	Name  string     `json:"name"`
	Addrs []net.Addr `json:"addrs"`
	MAC   string     `json:"mac"`
}

// HardwareInformation Hardware Information structure
type HardwareInformation struct {
	Memory    uint64  `json:"memory"`
	Mhz       float64 `json:"mhz"`
	Cores     int32   `json:"cores"`
	CPU       string  `json:"cpu"`
	Type      string  `json:"type"`
	GPIO      bool    `json:"gpio"`
	Display   bool    `json:"display"`
	Audio     bool    `json:"audio"`
	Camera    bool    `json:"camera"`
	Bluetooth bool    `json:"bluetooth"`
	Ble       bool    `json:"ble"`
	Wifi      bool    `json:"wifi"`
}

// DeviceInfo Device Information structure
type DeviceInfo struct {
	NetworkInterfaces []NetworkInterface  `json:"network_interfaces"`
	SoftwareInfo      SoftWareInformation `json:"sw_info"`
	Registered        DeviceRegistered    `json:"registered"`
	HardwareInfo      HardwareInformation `json:"hw_info"`
	Uptime            uint64              `json:"uptime"`
	HeartBeat         int64               `json:"heartbeat"`
	ID                uuid.UUID           `json:"id"`
	OnLine            bool                `json:"online"`
	HostName          string              `json:"hostname"`
}

// SoftWareInformation info about struct
type SoftWareInformation struct {
	Version  string `json:"version"`
	Name     string `json:"name"`
	Date     int64  `json:"date"`
	Os       string `json:"os"`
	Kernel   string `json:"kernel"`
	Platform string `json:"platform"`
}

func getBoardModel() string {
	path := "/proc/device-tree/model"
	dat, _ := ioutil.ReadFile(path)
	value := string(dat)
	return value
}

func _getUnameInfo() string {
	cmd := exec.Command("uname", "-srm")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("getInfo:", err)
	}
	return out.String()
}

func GetSoftwareInfo() SoftWareInformation {
	out := _getUnameInfo()
	for strings.Index(out, "broken pipe") != -1 {
		out = _getUnameInfo()
		time.Sleep(500 * time.Millisecond)
	}
	osStr := strings.Replace(out, "\n", "", -1)
	osStr = strings.Replace(osStr, "\r\n", "", -1)
	osInfo := strings.Split(osStr, " ")
	gio := SoftWareInformation{
		Kernel:   osInfo[1],
		Platform: osInfo[2],
		Os:       osInfo[0],
		Date:     1530394823,
		Name:     "mesh-iot",
		Version:  "0.1.0",
	}
	return gio
}

func GetNetworkInfo() []NetworkInterface {
	NetworkInterfaceSlice := []NetworkInterface{}
	interfStat, _ := net.Interfaces()

	for _, interf := range interfStat {
		addrs, _ := interf.Addrs()
		name := interf.Name
		mac := interf.HardwareAddr.String()
		if mac != "" && name != "lo" {
			NetworkInterfaceSlice = append(NetworkInterfaceSlice, NetworkInterface{
				Name:  name,
				Addrs: addrs,
				MAC:   mac})
		}

	}
	return NetworkInterfaceSlice
}

// TODO @ added support for read os-release information
// pi@raspberrypi:~ $ cat /etc/os-release
// PRETTY_NAME="Raspbian GNU/Linux 9 (stretch)"
// NAME="Raspbian GNU/Linux"
// VERSION_ID="9"
// VERSION="9 (stretch)"
// ID=raspbian
// ID_LIKE=debian
// HOME_URL="http://www.raspbian.org/"
// SUPPORT_URL="http://www.raspbian.org/RaspbianForums"
// BUG_REPORT_URL="http://www.raspbian.org/RaspbianBugs"

// StatusInfo function return data about Device
func StatusInfo() DeviceInfo {
	hostInfo, _ := host.Info()
	cpuInfo, _ := cpu.Info()
	virtMemInfo, _ := mem.VirtualMemory()
	networkInfo := GetNetworkInfo()
	swInfo := GetSoftwareInfo()
	hostname, _ := os.Hostname()
	uuidNumber := uuid.NewV4()
	boardType := getBoardModel()
	devInfo := DeviceInfo{
		OnLine:            true,
		Uptime:            hostInfo.Uptime,
		HeartBeat:         time.Now().Unix(),
		ID:                uuidNumber,
		HostName:          hostname,
		NetworkInterfaces: networkInfo,
		SoftwareInfo:      swInfo,
		HardwareInfo: HardwareInformation{
			Memory:    virtMemInfo.Total,
			Mhz:       cpuInfo[0].Mhz,
			Cores:     cpuInfo[0].Cores,
			CPU:       cpuInfo[0].ModelName,
			Type:      boardType,
			GPIO:      true,
			Display:   false,
			Audio:     false,
			Camera:    false,
			Bluetooth: false,
			Ble:       false,
			Wifi:      true,
		},
		Registered: DeviceRegistered{
			Status: false,
			Date:   time.Now().Unix(),
		},
	}
	return devInfo
}
