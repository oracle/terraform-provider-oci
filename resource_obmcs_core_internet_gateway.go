// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func InternetGatewayResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createInternetGateway,
		Read:     readInternetGateway,
		Update:   updateInternetGateway,
		Delete:   deleteInternetGateway,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_modified": {
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

func createInternetGateway(d *schema.ResourceData, m interface{}) (e error) {
	sync := &InternetGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)

	return crud.CreateResource(d, sync)
}

func readInternetGateway(d *schema.ResourceData, m interface{}) (e error) {
	sync := &InternetGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)

	return crud.ReadResource(sync)
}

func updateInternetGateway(d *schema.ResourceData, m interface{}) (e error) {
	sync := &InternetGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)

	return crud.UpdateResource(d, sync)

}

func deleteInternetGateway(d *schema.ResourceData, m interface{}) (e error) {
	sync := &InternetGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.DeleteResource(d, sync)
}

type InternetGatewayResourceCrud struct {
	crud.BaseCrud
	Resource *baremetal.InternetGateway
}

func (s *InternetGatewayResourceCrud) ID() string {
	return s.Resource.ID
}

func (s *InternetGatewayResourceCrud) CreatedPending() []string {
	return []string{
		baremetal.ResourceProvisioning,
	}
}

func (s *InternetGatewayResourceCrud) CreatedTarget() []string {
	return []string{
		baremetal.ResourceAvailable,
	}
}

func (s *InternetGatewayResourceCrud) DeletedPending() []string {
	return []string{
		baremetal.ResourceTerminating,
	}
}

func (s *InternetGatewayResourceCrud) DeletedTarget() []string {
	return []string{
		baremetal.ResourceTerminated,
	}
}

func (s *InternetGatewayResourceCrud) Create() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	vcnID := s.D.Get("vcn_id").(string)
	isEnabled := s.D.Get("enabled").(bool)

	opts := &baremetal.CreateOptions{}
	opts.DisplayName = s.D.Get("display_name").(string)

	s.Resource, e = s.Client.CreateInternetGateway(compartmentID, vcnID, isEnabled, opts)
	return
}

func (s *InternetGatewayResourceCrud) Get() (e error) {
	s.Resource, e = s.Client.GetInternetGateway(s.D.Id())
	return
}

func (s *InternetGatewayResourceCrud) Update() (e error) {
	opts := &baremetal.UpdateGatewayOptions{}
	if isEnabled, ok := s.D.GetOk("enabled"); ok {
		opts.IsEnabled = isEnabled.(bool)
	}

	s.Resource, e = s.Client.UpdateInternetGateway(s.D.Id(), opts)
	return
}

func (s *InternetGatewayResourceCrud) SetData() {
	s.D.Set("compartment_id", s.Resource.CompartmentID)
	s.D.Set("display_name", s.Resource.DisplayName)
	s.D.Set("enabled", s.Resource.IsEnabled)
	s.D.Set("time_modified", s.Resource.ModifiedTime.String())
	s.D.Set("state", s.Resource.State)
	s.D.Set("time_created", s.Resource.TimeCreated.String())
}

func (s *InternetGatewayResourceCrud) Delete() (e error) {
	return s.Client.DeleteInternetGateway(s.D.Id(), nil)
}
