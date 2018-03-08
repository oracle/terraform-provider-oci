// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_core "github.com/oracle/oci-go-sdk/core"
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
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"id": {
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

func createCpe(d *schema.ResourceData, m interface{}) error {
	sync := &CpeResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.CreateResource(d, sync)
}

func readCpe(d *schema.ResourceData, m interface{}) error {
	sync := &CpeResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

func updateCpe(d *schema.ResourceData, m interface{}) error {
	sync := &CpeResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.UpdateResource(d, sync)
}

func deleteCpe(d *schema.ResourceData, m interface{}) error {
	sync := &CpeResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type CpeResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.Cpe
	DisableNotFoundRetries bool
}

func (s *CpeResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CpeResourceCrud) Create() error {
	request := oci_core.CreateCpeRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if ipAddress, ok := s.D.GetOkExists("ip_address"); ok {
		tmp := ipAddress.(string)
		request.IpAddress = &tmp
	}

	response, err := s.Client.CreateCpe(context.Background(), request, getRetryOptions(s.DisableNotFoundRetries, "core")...)
	if err != nil {
		return err
	}

	s.Res = &response.Cpe
	return nil
}

func (s *CpeResourceCrud) Get() error {
	request := oci_core.GetCpeRequest{}

	tmp := s.D.Id()
	request.CpeId = &tmp

	response, err := s.Client.GetCpe(context.Background(), request, getRetryOptions(s.DisableNotFoundRetries, "core")...)
	if err != nil {
		return err
	}

	s.Res = &response.Cpe
	return nil
}

func (s *CpeResourceCrud) Update() error {
	request := oci_core.UpdateCpeRequest{}

	tmp := s.D.Id()
	request.CpeId = &tmp

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	response, err := s.Client.UpdateCpe(context.Background(), request, getRetryOptions(s.DisableNotFoundRetries, "core")...)
	if err != nil {
		return err
	}

	s.Res = &response.Cpe
	return nil
}

func (s *CpeResourceCrud) Delete() error {
	request := oci_core.DeleteCpeRequest{}

	tmp := s.D.Id()
	request.CpeId = &tmp

	_, err := s.Client.DeleteCpe(context.Background(), request, getRetryOptions(s.DisableNotFoundRetries, "core")...)
	return err
}

func (s *CpeResourceCrud) SetData() {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	if s.Res.IpAddress != nil {
		s.D.Set("ip_address", *s.Res.IpAddress)
	}

	s.D.Set("time_created", s.Res.TimeCreated.String())

}
