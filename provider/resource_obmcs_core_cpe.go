// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/crud"
)

func CpeResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createCpe,
		Read:     readCpe,
		Update:   updateCpe,
		Delete:   deleteCpe,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}

}

func createCpe(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*OracleClients)
	sync := &CpeResourceCrud{}
	sync.D = d
	sync.Client = client.client
	return crud.CreateResource(d, sync)
}

func readCpe(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*OracleClients)
	sync := &CpeResourceCrud{}
	sync.D = d
	sync.Client = client.client
	return crud.ReadResource(sync)
}

func updateCpe(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*OracleClients)
	crd := &CpeResourceCrud{}
	crd.D = d
	crd.Client = client.client
	return crud.UpdateResource(d, crd)
}

func deleteCpe(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*OracleClients)
	sync := &CpeResourceCrud{}
	sync.D = d
	sync.Client = client.clientWithoutNotFoundRetries
	return sync.Delete()
}

type CpeResourceCrud struct {
	crud.BaseCrud
	Resource *baremetal.Cpe
}

func (s *CpeResourceCrud) ID() string {
	return s.Resource.ID
}

func (s *CpeResourceCrud) Create() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	ipAddress := s.D.Get("ip_address").(string)

	opts := &baremetal.CreateOptions{}
	displayName, ok := s.D.GetOk("display_name")
	if ok {
		opts.DisplayName = displayName.(string)
	}

	s.Resource, e = s.Client.CreateCpe(compartmentID, ipAddress, opts)
	return
}

func (s *CpeResourceCrud) Get() (e error) {
	res, e := s.Client.GetCpe(s.D.Id())
	if e == nil {
		s.Resource = res
	}
	return
}

func (s *CpeResourceCrud) Update() (e error) {
	opts := &baremetal.IfMatchDisplayNameOptions{}
	compartmentID := s.D.Get("compartment_id").(string)
	displayName, ok := s.D.GetOk("display_name")
	if ok {
		opts.DisplayName = displayName.(string)
	}

	s.Resource, e = s.Client.UpdateCpe(compartmentID, opts)
	return
}

func (s *CpeResourceCrud) SetData() {
	s.D.Set("compartment_id", s.Resource.CompartmentID)
	s.D.Set("display_name", s.Resource.DisplayName)
	s.D.Set("ip_address", s.Resource.IPAddress)
	s.D.Set("time_created", s.Resource.TimeCreated.String())
}

func (s *CpeResourceCrud) Delete() (e error) {
	return s.Client.DeleteCpe(s.D.Id(), nil)
}
