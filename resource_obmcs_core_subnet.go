// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"
)

func SubnetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createSubnet,
		Read:     readSubnet,
		Update:   updateSubnet,
		Delete:   deleteSubnet,
		Schema: map[string]*schema.Schema{
			"availability_domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cidr_block": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dhcp_options_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"route_table_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"security_list_ids": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			// Optional
			"dns_label": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prohibit_public_ip_on_vnic": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},
			// Computed
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"virtual_router_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"virtual_router_mac": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createSubnet(d *schema.ResourceData, m interface{}) (e error) {
	sync := &SubnetResourceCrud{}
	sync.D = d
	sync.Client = m.(*baremetal.Client)
	return crud.CreateResource(d, sync)
}

func readSubnet(d *schema.ResourceData, m interface{}) (e error) {
	sync := &SubnetResourceCrud{}
	sync.D = d
	sync.Client = m.(*baremetal.Client)
	return crud.ReadResource(sync)
}

func updateSubnet(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &SubnetResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.UpdateResource(sync.D, sync)
}

func deleteSubnet(d *schema.ResourceData, m interface{}) (e error) {
	sync := &SubnetResourceCrud{}
	sync.D = d
	sync.Client = m.(*baremetal.Client)
	return crud.DeleteResource(d, sync)
}

type SubnetResourceCrud struct {
	crud.BaseCrud
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

func (s *SubnetResourceCrud) ExtraWaitPostCreateDelete() time.Duration {
	return time.Duration(25 * time.Second)
}

func (s *SubnetResourceCrud) Create() (e error) {
	availabilityDomain := s.D.Get("availability_domain").(string)
	cidrBlock := s.D.Get("cidr_block").(string)
	compartmentID := s.D.Get("compartment_id").(string)
	vcnID := s.D.Get("vcn_id").(string)

	opts := &baremetal.CreateSubnetOptions{}
	if dhcpOptionsID, ok := s.D.GetOk("dhcp_options_id"); ok {
		opts.DHCPOptionsID = dhcpOptionsID.(string)
	}
	if dnsLabel, ok := s.D.GetOk("dns_label"); ok {
		opts.DNSLabel = dnsLabel.(string)
	}
	if displayName, ok := s.D.GetOk("display_name"); ok {
		opts.DisplayName = displayName.(string)
	}

	dnsLabel, ok := s.D.GetOk("dns_label")
	if ok {
		opts.DNSLabel = dnsLabel.(string)
	}

	prohibitPublicIpOnVnic, ok := s.D.GetOk("prohibit_public_ip_on_vnic")
	if ok {
		opts.ProhibitPublicIpOnVnic = prohibitPublicIpOnVnic.(bool)
	}

	if rawSecurityListIDs, ok := s.D.GetOk("security_list_ids"); ok {
		securityListIDs := []string{}
		for _, val := range rawSecurityListIDs.(*schema.Set).List() {
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
	res, e := s.Client.GetSubnet(s.D.Id())
	if e == nil {
		s.Resource = res
	}
	return
}

func (s *SubnetResourceCrud) Update() (e error) {
	opts := &baremetal.IfMatchDisplayNameOptions{}
	
	displayName, ok := s.D.GetOk("display_name")
	if ok {
		opts.DisplayName = displayName.(string)
	}

	s.Resource, e = s.Client.UpdateSubnet(s.D.Id(), opts)
	return
}

func (s *SubnetResourceCrud) SetData() {
	s.D.Set("availability_domain", s.Resource.AvailabilityDomain)
	s.D.Set("compartment_id", s.Resource.CompartmentID)
	s.D.Set("display_name", s.Resource.DisplayName)
	s.D.Set("dns_label", s.Resource.DNSLabel)
	s.D.Set("cidr_block", s.Resource.CIDRBlock)
	s.D.Set("dhcp_options_id", s.Resource.DHCPOptionsID)
	s.D.Set("dns_label", s.Resource.DNSLabel)
	s.D.Set("prohibit_public_ip_on_vnic", s.Resource.ProhibitPublicIpOnVnic)
	s.D.Set("route_table_id", s.Resource.RouteTableID)
	s.D.Set("vcn_id", s.Resource.VcnID)
	s.D.Set("security_list_ids", makeSetFromStrings(s.Resource.SecurityListIDs))
	s.D.Set("state", s.Resource.State)
	s.D.Set("time_created", s.Resource.TimeCreated.String())
	s.D.Set("virtual_router_ip", s.Resource.VirtualRouterIP)
	s.D.Set("virtual_router_mac", s.Resource.VirtualRouterMac)
}

func (s *SubnetResourceCrud) Delete() (e error) {
	return s.Client.DeleteSubnet(s.D.Id(), nil)
}

// makeSetFromStrings encodes an []string into a
// *schema.Set in the appropriate structure for the schema
func makeSetFromStrings(ss []string) *schema.Set {
	st := &schema.Set{F: schema.HashString}
	for _, s := range ss {
		st.Add(s)
	}
	return st
}
