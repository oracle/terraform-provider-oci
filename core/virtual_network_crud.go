package core

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type VirtualNetworkSync struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.VirtualNetwork
}

func (s *VirtualNetworkSync) ID() string {
	return s.Res.ID
}

func (s *VirtualNetworkSync) CreatedPending() []string {
	return []string{baremetal.ResourceProvisioning}
}

func (s *VirtualNetworkSync) CreatedTarget() []string {
	return []string{baremetal.ResourceAvailable}
}

func (s *VirtualNetworkSync) DeletedPending() []string {
	return []string{baremetal.ResourceTerminating}
}

func (s *VirtualNetworkSync) DeletedTarget() []string {
	return []string{baremetal.ResourceTerminated}
}

func (s *VirtualNetworkSync) State() string {
	return s.Res.State
}

func (s *VirtualNetworkSync) Create() (e error) {
	opts := baremetal.Options{}
	cidrBlock := s.D.Get("cidr_block").(string)
	compartmentID := s.D.Get("compartment_id").(string)
	displayName, ok := s.D.GetOk("display_name")
	if ok {
		opts.DisplayName = displayName.(string)
	}

	s.Res, e = s.Client.CreateVirtualNetwork(cidrBlock, compartmentID, opts)

	return
}

func (s *VirtualNetworkSync) Get() (e error) {
	s.Res, e = s.Client.GetVirtualNetwork(s.D.Id())
	return
}

func (s *VirtualNetworkSync) SetData() {
	s.D.Set("cidr_block", s.Res.CidrBlock)
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("default_routing_table_id", s.Res.DefaultRoutingTableID)
	s.D.Set("default_security_list_id", s.Res.DefaultSecurityListID)
	s.D.Set("display_name", s.Res.DisplayName)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
}

func (s *VirtualNetworkSync) Delete() (e error) {
	return s.Client.DeleteVirtualNetwork(s.D.Id())
}
