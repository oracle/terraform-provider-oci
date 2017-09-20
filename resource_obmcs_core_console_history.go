// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"
)

func ConsoleHistoryResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createConsoleHistory,
		Read:     readConsoleHistory,
		Delete:   deleteConsoleHistory,
		Schema: map[string]*schema.Schema{
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
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
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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

func createConsoleHistory(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	ichCrud := &ConsoleHistoryResourceCrud{}
	ichCrud.D = d
	ichCrud.Client = client
	return crud.CreateResource(d, ichCrud)
}

func readConsoleHistory(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	ichCrud := &ConsoleHistoryResourceCrud{}
	ichCrud.D = d
	ichCrud.Client = client
	return crud.ReadResource(ichCrud)
}

func deleteConsoleHistory(d *schema.ResourceData, m interface{}) (e error) {
	sync := &ConsoleHistoryResourceCrud{}
	sync.D = d
	return crud.DeleteResource(d, sync)
}

type ConsoleHistoryResourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ConsoleHistoryMetadata
}

func (s *ConsoleHistoryResourceCrud) ID() string {
	return s.Res.ID
}

func (s *ConsoleHistoryResourceCrud) CreatedPending() []string {
	return []string{baremetal.ResourceRequested, baremetal.ResourceGettingHistory}
}

func (s *ConsoleHistoryResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceSucceeded}
}

func (s *ConsoleHistoryResourceCrud) State() string {
	return s.Res.State
}

func (s *ConsoleHistoryResourceCrud) Create() (e error) {
	instanceID := s.D.Get("instance_id").(string)

	s.Res, e = s.Client.CaptureConsoleHistory(instanceID, nil)

	return
}

func (s *ConsoleHistoryResourceCrud) Get() (e error) {
	res, e := s.Client.GetConsoleHistory(s.D.Id())
	if e == nil {
		s.Res = res
	}
	return
}

func (s *ConsoleHistoryResourceCrud) Delete() (e error) {
	e = nil
	return
}

func (s *ConsoleHistoryResourceCrud) SetData() {
	s.D.Set("availability_domain", s.Res.AvailabilityDomain)
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("display_name", s.Res.DisplayName)
	s.D.Set("instance_id", s.Res.InstanceID)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
}
