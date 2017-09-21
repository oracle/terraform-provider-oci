// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"
)

func IPSecConnectionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createIPSec,
		Read:     readIPSec,
		Update:   updateIPSec,
		Delete:   deleteIPSec,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cpe_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"drg_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"static_routes": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
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

func createIPSec(d *schema.ResourceData, m interface{}) (e error) {
	sync := &IPSecConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*baremetal.Client)
	return crud.CreateResource(d, sync)
}

func readIPSec(d *schema.ResourceData, m interface{}) (e error) {
	sync := &IPSecConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*baremetal.Client)
	return crud.ReadResource(sync)
}

func updateIPSec(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &IPSecConnectionResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.UpdateResource(sync.D, sync)
}

func deleteIPSec(d *schema.ResourceData, m interface{}) (e error) {
	sync := &IPSecConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*baremetal.Client)
	return crud.DeleteResource(d, sync)
}

type IPSecConnectionResourceCrud struct {
	crud.BaseCrud
	Resource *baremetal.IPSecConnection
}

func (s *IPSecConnectionResourceCrud) ID() string {
	return s.Resource.ID
}

func (s *IPSecConnectionResourceCrud) CreatedPending() []string {
	return []string{baremetal.ResourceProvisioning}
}

func (s *IPSecConnectionResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceAvailable}
}

func (s *IPSecConnectionResourceCrud) DeletedPending() []string {
	return []string{baremetal.ResourceTerminating}
}

func (s *IPSecConnectionResourceCrud) DeletedTarget() []string {
	return []string{baremetal.ResourceTerminated}
}

func (s *IPSecConnectionResourceCrud) Create() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	cpeID := s.D.Get("cpe_id").(string)
	drgID := s.D.Get("drg_id").(string)

	staticRoutes := []string{}
	for _, route := range s.D.Get("static_routes").([]interface{}) {
		staticRoutes = append(staticRoutes, route.(string))
	}

	opts := &baremetal.CreateOptions{}
	if displayName, ok := s.D.GetOk("display_name"); ok {
		opts.DisplayName = displayName.(string)
	}

	s.Resource, e = s.Client.CreateIPSecConnection(
		compartmentID,
		cpeID,
		drgID,
		staticRoutes,
		opts,
	)

	return
}

func (s *IPSecConnectionResourceCrud) Get() (e error) {
	res, e := s.Client.GetIPSecConnection(s.D.Id())
	if e == nil {
		s.Resource = res
	}
	return
}

func (s *IPSecConnectionResourceCrud) Update() (e error) {
	opts := &baremetal.IfMatchDisplayNameOptions{}
	displayName, ok := s.D.GetOk("display_name")
	if ok {
		opts.DisplayName = displayName.(string)
	}

	s.Resource, e = s.Client.UpdateIPSecConnection(s.D.Id(), opts)
	return
}

func (s *IPSecConnectionResourceCrud) SetData() {
	s.D.Set("compartment_id", s.Resource.CompartmentID)
	s.D.Set("cpe_id", s.Resource.CpeID)
	s.D.Set("drg_id", s.Resource.DrgID)
	s.D.Set("static_routes", s.Resource.StaticRoutes)
	s.D.Set("display_name", s.Resource.DisplayName)
	s.D.Set("state", s.Resource.State)
	s.D.Set("time_created", s.Resource.TimeCreated.String())

}

func (s *IPSecConnectionResourceCrud) Delete() (e error) {
	return s.Client.DeleteIPSecConnection(s.D.Id(), nil)
}
