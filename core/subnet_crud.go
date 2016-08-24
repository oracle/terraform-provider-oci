package core

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type SubnetSync struct {
	D        *schema.ResourceData
	Client   client.BareMetalClient
	Resource *baremetal.Subnet
}

func (s *SubnetSync) ID() string {
	return s.Resource.ID
}

func (s *SubnetSync) CreatedPending() []string {
	return []string{baremetal.ResourceProvisioning}
}

func (s *SubnetSync) CreatedTarget() []string {
	return []string{baremetal.ResourceAvailable}
}

func (s *SubnetSync) DeletedPending() []string {
	return []string{baremetal.ResourceTerminating}
}

func (s *SubnetSync) DeletedTarget() []string {
	return []string{baremetal.ResourceTerminated}
}

func (s *SubnetSync) Create() (e error) {
	opts := baremetal.Options{}
	availabilityDomain := s.D.Get("availability_domain").(string)
	cidrBlock := s.D.Get("cidr_block").(string)
	compartmentID := s.D.Get("compartment_id").(string)
	routeTableID := s.D.Get("route_table_id").(string)
	vcnID := s.D.Get("vcn_id").(string)

	securityListIDs := []string{}
	for _, val := range s.D.Get("security_list_ids").([]interface{}) {
		securityListIDs = append(securityListIDs, val.(string))
	}

	if displayName, ok := s.D.GetOk("display_name"); ok {
		opts.DisplayName = displayName.(string)
	}

	s.Resource, e = s.Client.CreateSubnet(
		availabilityDomain,
		cidrBlock,
		compartmentID,
		routeTableID,
		vcnID,
		securityListIDs,
		opts,
	)

	return
}

func (s *SubnetSync) Get() (e error) {
	s.Resource, e = s.Client.GetSubnet(s.D.Id())
	return
}

func (s *SubnetSync) SetData() {
	s.D.Set("availability_domain", s.Resource.AvailabilityDomain)
	s.D.Set("compartment_id", s.Resource.CompartmentID)
	s.D.Set("display_name", s.Resource.DisplayName)
	s.D.Set("cidr_block", s.Resource.CIDRBlock)
	s.D.Set("route_table_id", s.Resource.RouteTableID)
	s.D.Set("vcn_id", s.Resource.VcnID)
	s.D.Set("security_list_ids", s.Resource.SecurityListIDs)
	s.D.Set("state", s.Resource.State)
	s.D.Set("time_created", s.Resource.TimeCreated.String())
	s.D.Set("virtual_router_id", s.Resource.VirtualRouterID)
	s.D.Set("virtual_router_mac", s.Resource.VirtualRouterMac)
}

func (s *SubnetSync) Delete() (e error) {
	return s.Client.DeleteSubnet(s.D.Id())
}
