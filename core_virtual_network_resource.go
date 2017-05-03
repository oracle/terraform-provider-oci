// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func VirtualNetworkResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createVirtualNetwork,
		Read:     readVirtualNetwork,
		Update:   updateVirtualNetwork,
		Delete:   deleteVirtualNetwork,
		Schema: map[string]*schema.Schema{
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
			"default_route_table_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_security_list_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_dhcp_options_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dns_label": {
				Type:     schema.TypeString,
				Optional: true,
			},
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
		},
	}
}

func createVirtualNetwork(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &VirtualNetworkResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.CreateResource(d, sync)
}

func readVirtualNetwork(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &VirtualNetworkResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

func updateVirtualNetwork(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &VirtualNetworkResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.UpdateResource(sync.D, sync)
}

func deleteVirtualNetwork(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &VirtualNetworkResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.DeleteResource(d, sync)
}

type VirtualNetworkResourceCrud struct {
	crud.BaseCrud
	Res *baremetal.VirtualNetwork
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

func (s *VirtualNetworkResourceCrud) Create() (e error) {
	cidrBlock := s.D.Get("cidr_block").(string)
	compartmentID := s.D.Get("compartment_id").(string)

	opts := &baremetal.CreateVcnOptions{}
	displayName, ok := s.D.GetOk("display_name")
	if ok {
		opts.DisplayName = displayName.(string)
	}

	dnsLabel, ok := s.D.GetOk("dns_label")
	if ok {
		opts.DnsLabel = dnsLabel.(string)
	}

	s.Res, e = s.Client.CreateVirtualNetwork(cidrBlock, compartmentID, opts)

	return
}

func (s *VirtualNetworkResourceCrud) Get() (e error) {
	s.Res, e = s.Client.GetVirtualNetwork(s.D.Id())
	return
}

func (s *VirtualNetworkResourceCrud) Update() (e error) {
	opts := &baremetal.IfMatchDisplayNameOptions{}
	compartmentID := s.D.Get("compartment_id").(string)
	displayName, ok := s.D.GetOk("display_name")
	if ok {
		opts.DisplayName = displayName.(string)
	}

	s.Res, e = s.Client.UpdateVirtualNetwork(compartmentID, opts)
	return
}

func (s *VirtualNetworkResourceCrud) SetData() {
	s.D.Set("cidr_block", s.Res.CidrBlock)
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("dns_label", s.Res.DnsLabel)
	s.D.Set("default_route_table_id", s.Res.DefaultRouteTableID)
	s.D.Set("default_security_list_id", s.Res.DefaultSecurityListID)
	s.D.Set("default_dhcp_options_id", s.Res.DefaultDHCPOptionsID)
	s.D.Set("display_name", s.Res.DisplayName)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
}

func (s *VirtualNetworkResourceCrud) Delete() (e error) {
	return s.Client.DeleteVirtualNetwork(s.D.Id(), nil)
}
