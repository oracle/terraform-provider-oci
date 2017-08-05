// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"
)

func DrgResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createDrg,
		Read:     readDrg,
		Update:   updateDrg,
		Delete:   deleteDrg,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
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

func createDrg(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &DrgResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.CreateResource(d, sync)
}

func readDrg(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &DrgResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

func updateDrg(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &DrgResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.UpdateResource(sync.D, sync)
}

func deleteDrg(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &DrgResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.DeleteResource(d, sync)
}

type DrgResourceCrud struct {
	crud.BaseCrud
	Res *baremetal.Drg
}

func (s *DrgResourceCrud) ID() string {
	return s.Res.ID
}

func (s *DrgResourceCrud) CreatedPending() []string {
	return []string{baremetal.ResourceProvisioning}
}

func (s *DrgResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceAvailable}
}

func (s *DrgResourceCrud) DeletedPending() []string {
	return []string{baremetal.ResourceTerminating}
}

func (s *DrgResourceCrud) DeletedTarget() []string {
	return []string{baremetal.ResourceTerminated}
}

func (s *DrgResourceCrud) Create() (e error) {
	opts := &baremetal.CreateOptions{}
	compartmentID := s.D.Get("compartment_id").(string)
	displayName, ok := s.D.GetOk("display_name")
	if ok {
		opts.DisplayName = displayName.(string)
	}

	s.Res, e = s.Client.CreateDrg(compartmentID, opts)

	return
}

func (s *DrgResourceCrud) Get() (e error) {
	s.Res, e = s.Client.GetDrg(s.D.Id())
	return
}

func (s *DrgResourceCrud) Update() (e error) {
	opts := &baremetal.IfMatchDisplayNameOptions{}
	compartmentID := s.D.Get("compartment_id").(string)
	displayName, ok := s.D.GetOk("display_name")
	if ok {
		opts.DisplayName = displayName.(string)
	}

	s.Res, e = s.Client.UpdateDrg(compartmentID, opts)
	return
}

func (s *DrgResourceCrud) SetData() {
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("display_name", s.Res.DisplayName)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
}

func (s *DrgResourceCrud) Delete() (e error) {
	return s.Client.DeleteDrg(s.D.Id(), nil)
}
