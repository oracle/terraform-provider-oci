package data_types

import (
	"fmt"
	"net"
)

type SoftLayer_Network_Subnet struct {
	Id                int    `json:"id"`
	NetworkIdentifier string `json:"networkIdentifier"`
	Gateway           string `json:"gateway"`
	BroadcastAddress  string `json:"broadcastAddress"`
	Netmask           string `json:"netmask"`
}

func (s SoftLayer_Network_Subnet) Contains(address string) bool {
	ipNet := net.IPNet{
		IP:   net.ParseIP(s.NetworkIdentifier),
		Mask: net.IPMask(net.ParseIP(s.Netmask)),
	}

	return ipNet.Contains(net.ParseIP(address))
}

type SoftLayer_Network_Subnets []SoftLayer_Network_Subnet

func (s SoftLayer_Network_Subnets) Containing(address string) (SoftLayer_Network_Subnet, error) {
	for _, subnet := range s {
		if subnet.Contains(address) {
			return subnet, nil
		}
	}

	return SoftLayer_Network_Subnet{}, fmt.Errorf("subnet not found for %q", address)
}
