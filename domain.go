/*
 * This file is part of the libvirt-go-xml project
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 *
 * Copyright (C) 2016 Red Hat, Inc.
 *
 */

package libvirtxml

import (
	"encoding/xml"
)

type DomainController struct {
	Type  string `xml:"type,attr"`
	Index string `xml:"index,attr"`
}

type DomainDiskFileSource struct {
	File string `xml:"file,attr"`
}

type DomainDiskDriver struct {
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
}

type DomainDiskTarget struct {
	Dev string `xml:"dev,attr"`
	Bus string `xml:"bus,attr"`
}

type DomainDisk struct {
	Type       string                `xml:"type,attr"`
	Device     string                `xml:"device,attr"`
	Snapshot   string                `xml:"snapshot,attr,omitempty"`
	Driver     *DomainDiskDriver     `xml:"driver"`
	FileSource *DomainDiskFileSource `xml:"source"`
	Target     *DomainDiskTarget     `xml:"target"`
}

type DomainInterfaceMAC struct {
	Address string `xml:"address,attr"`
}

type DomainInterfaceModel struct {
	Type string `xml:type,attr"`
}

type DomainInterface struct {
	Type  string               `xml:"type,attr"`
	MAC   string               `xml:"mac"`
	Model DomainInterfaceModel `xml:"model"`
}

type DomainChardev struct {
	Type string `xml:"type,attr"`
}

type DomainInput struct {
	Type string `xml:"type,attr"`
	Bus  string `xml:"bus,attr"`
}

type DomainGraphicListener struct {
	Type    string `xml:"type,attr"`
	Address string `xml:"address,attr,omitempty"`
	Network string `xml:"network,attr,omitempty"`
	Socket  string `xml:"socket,attr,omitempty"`
}

type DomainGraphic struct {
	Type          string                  `xml:"type,attr"`
	AutoPort      string                  `xml:"autoport,attr,omitempty"`
	Port          int                     `xml:"port,attr,omitempty"`
	TLSPort       int                     `xml:"tlsPort,attr,omitempty"`
	WebSocket     int                     `xml:"websocket,attr,omitempty"`
	Listen        string                  `xml:"listen,attr,omitempty"`
	Socket        string                  `xml:"socket,attr,omitempty"`
	Keymap        string                  `xml:"keymap,attr,omitempty"`
	Passwd        string                  `xml:"passwd,attr,omitempty"`
	PasswdValidTo string                  `xml:"passwdValidTo,attr,omitempty"`
	Connected     string                  `xml:"connected,attr,omitempty"`
	SharePolicy   string                  `xml:"sharePolicy,attr,omitempty"`
	DefaultMode   string                  `xml:"defaultMode,attr,omitempty"`
	Display       string                  `xml:"display,attr,omitempty"`
	XAuth         string                  `xml:"xauth,attr,omitempty"`
	FullScreen    string                  `xml:"fullscreen,attr,omitempty"`
	ReplaceUser   string                  `xml:"replaceUser,attr,omitempty"`
	MultiUser     string                  `xml:"multiUser,attr,omitempty"`
	Listeners     []DomainGraphicListener `xml:"listen"`
}

type DomainVideoModel struct {
	Type string `xml:"type,attr"`
}

type DomainVideo struct {
	Model DomainVideoModel `xml:"model"`
}

type DomainDeviceList struct {
	Controllers []DomainController `xml:"controller"`
	Disks       []DomainDisk       `xml:"disk"`
	Interfaces  []DomainInterface  `xml:"interface"`
	Serials     []DomainChardev    `xml:"serial"`
	Consoles    []DomainChardev    `xml:"console"`
	Inputs      []DomainInput      `xml:"input"`
	Graphics    []DomainGraphic    `xml:"graphics"`
	Videos      []DomainVideo      `xml:"video"`
}

type DomainMemory struct {
	Value int    `xml:",chardata"`
	Unit  string `xml:"unit,attr,omitempty"`
}

type DomainMaxMemory struct {
	Value int    `xml:",chardata"`
	Unit  string `xml:"unit,attr,omitempty"`
	Slots int    `xml:"slots,attr,omitempty"`
}

type DomainOSType struct {
	Arch    string `xml:"arch,attr"`
	Machine string `xml:"machine,attr"`
	Type    string `xml:",chardata"`
}

type DomainSMBios struct {
	Mode string `xml:"mode,attr"`
}

type DomainNVRam struct {
	NVRam    string `xml:",chardata"`
	Template string `xml:"template,attr"`
}

type DomainBootDevice struct {
	Dev string `xml:"dev,attr"`
}

type DomainBootMenu struct {
	Enabled string `xml:"enabled,attr"`
	Timeout string `xml:"timeout,attr"`
}

type DomainSysInfo struct {
	Type      string               `xml:"type,attr"`
	System    []DomainSysInfoEntry `xml:"system>entry"`
	BIOS      []DomainSysInfoEntry `xml:"bios>entry"`
	BaseBoard []DomainSysInfoEntry `xml:"baseBoard>entry"`
}

type DomainSysInfoEntry struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

type DomainBIOS struct {
	UseSerial     string `xml:"useserial,attr"`
	RebootTimeout string `xml:"rebootTimeout,attr"`
}

type DomainLoader struct {
	Path     string `xml:",chardata"`
	Readonly string `xml:"readonly,attr"`
	Secure   string `xml:"secure,attr"`
	Type     string `xml:"type,attr"`
}

type DomainOS struct {
	Type        *DomainOSType      `xml:"type"`
	Loader      *DomainLoader      `xml:"loader"`
	NVRam       *DomainNVRam       `xml:"nvram"`
	Kernel      string             `xml:"kernel,omitempty"`
	Initrd      string             `xml:"initrd,omitempty"`
	KernelArgs  string             `xml:"cmdline,omitempty"`
	BootDevices []DomainBootDevice `xml:"boot"`
	BootMenu    *DomainBootMenu    `xml:"bootmenu"`
	SMBios      *DomainSMBios      `xml:"smbios"`
	BIOS        *DomainBIOS        `xml:"bios"`
	Init        string             `xml:"init,omitempty"`
	InitArgs    []string           `xml:"initarg"`
}

type DomainResource struct {
	Partition string `xml:"partition,omitempty"`
}

type Domain struct {
	XMLName       xml.Name          `xml:"domain"`
	Type          string            `xml:"type,attr,omitempty"`
	Name          string            `xml:"name"`
	UUID          string            `xml:"uuid,omitempty"`
	Memory        *DomainMemory     `xml:"memory"`
	CurrentMemory *DomainMemory     `xml:"currentMemory"`
	MaximumMemory *DomainMaxMemory  `xml:"maxMemory"`
	Resource      *DomainResource   `xml:"resource"`
	Devices       *DomainDeviceList `xml:"devices"`
	OS            *DomainOS         `xml:"os"`
	SysInfo       *DomainSysInfo    `xml:"sysinfo"`
	OnPoweroff    string            `xml:"on_poweroff,omitempty"`
	OnReboot      string            `xml:"on_reboot,omitempty"`
	OnCrash       string            `xml:"on_crash,omitempty"`
}

func (d *Domain) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *Domain) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}
