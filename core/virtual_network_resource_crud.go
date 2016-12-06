package core

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type VirtualNetworkResourceCrud struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.VirtualNetwork
}

func (s *VirtualNetworkResourceCrud) ID() string {
	return s.Res.ID
}

func (s *VirtualNetworkResourceCrud) CreatedPending() []string {
	return []string{baremetal.ResourceProvisioning}
}

func (s *VirtualNetworkResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceAvailable}
}

func (s *VirtualNetworkResourceCrud) DeletedPending() []string {
	return []string{baremetal.ResourceTerminating}
}

func (s *VirtualNetworkResourceCrud) DeletedTarget() []string {
	return []string{baremetal.ResourceTerminated}
}

func (s *VirtualNetworkResourceCrud) State() string {
	return s.Res.State
}

func (s *VirtualNetworkResourceCrud) Create() (e error) {
	cidrBlock := s.D.Get("cidr_block").(string)
	compartmentID := s.D.Get("compartment_id").(string)

	opts := &baremetal.CreateOptions{}
	displayName, ok := s.D.GetOk("display_name")
	if ok {
		opts.DisplayName = displayName.(string)
	}

	s.Res, e = s.Client.CreateVirtualNetwork(cidrBlock, compartmentID, opts)

	return
}

func (s *VirtualNetworkResourceCrud) Get() (e error) {
	s.Res, e = s.Client.GetVirtualNetwork(s.D.Id())
	return
}

func (s *VirtualNetworkResourceCrud) SetData() {
	s.D.Set("cidr_block", s.Res.CidrBlock)
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("default_routing_table_id", s.Res.DefaultRoutingTableID)
	s.D.Set("default_security_list_id", s.Res.DefaultSecurityListID)
	s.D.Set("display_name", s.Res.DisplayName)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
}

func (s *VirtualNetworkResourceCrud) Delete() (e error) {
	return s.Client.DeleteVirtualNetwork(s.D.Id(), nil)
}
