package core

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type SubnetResourceCrud struct {
	D        *schema.ResourceData
	Client   client.BareMetalClient
	Resource *baremetal.Subnet
}

func (s *SubnetResourceCrud) ID() string {
	return s.Resource.ID
}

func (s *SubnetResourceCrud) CreatedPending() []string {
	return []string{baremetal.ResourceProvisioning}
}

func (s *SubnetResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceAvailable}
}

func (s *SubnetResourceCrud) DeletedPending() []string {
	return []string{baremetal.ResourceTerminating}
}

func (s *SubnetResourceCrud) DeletedTarget() []string {
	return []string{baremetal.ResourceTerminated}
}

func (s *SubnetResourceCrud) Create() (e error) {
	availabilityDomain := s.D.Get("availability_domain").(string)
	cidrBlock := s.D.Get("cidr_block").(string)
	compartmentID := s.D.Get("compartment_id").(string)
	vcnID := s.D.Get("vcn_id").(string)

	opts := &baremetal.CreateSubnetOptions{}

	if displayName, ok := s.D.GetOk("display_name"); ok {
		opts.DisplayName = displayName.(string)
	}

	if rawSecurityListIDs, ok := s.D.GetOk("security_list_ids"); ok {
		securityListIDs := []string{}
		for _, val := range rawSecurityListIDs.([]interface{}) {
			securityListIDs = append(securityListIDs, val.(string))
		}
		opts.SecurityListIDs = securityListIDs
	}

	if routeTableID, ok := s.D.GetOk("route_table_id"); ok {
		opts.RouteTableID = routeTableID.(string)
	}

	s.Resource, e = s.Client.CreateSubnet(
		availabilityDomain,
		cidrBlock,
		compartmentID,
		vcnID,
		opts,
	)

	return
}

func (s *SubnetResourceCrud) Get() (e error) {
	s.Resource, e = s.Client.GetSubnet(s.D.Id())
	return
}

func (s *SubnetResourceCrud) SetData() {
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

func (s *SubnetResourceCrud) Delete() (e error) {
	return s.Client.DeleteSubnet(s.D.Id(), nil)
}
