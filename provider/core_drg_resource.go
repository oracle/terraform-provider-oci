// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_core "github.com/oracle/oci-go-sdk/core"
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
			// Required
			"compartment_id": {
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

func createDrg(d *schema.ResourceData, m interface{}) error {
	sync := &DrgResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.CreateResource(d, sync)
}

func readDrg(d *schema.ResourceData, m interface{}) error {
	sync := &DrgResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

func updateDrg(d *schema.ResourceData, m interface{}) error {
	sync := &DrgResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.UpdateResource(d, sync)
}

func deleteDrg(d *schema.ResourceData, m interface{}) error {
	sync := &DrgResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type DrgResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.Drg
	DisableNotFoundRetries bool
}

func (s *DrgResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DrgResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.DrgLifecycleStateProvisioning),
	}
}

func (s *DrgResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.DrgLifecycleStateAvailable),
	}
}

func (s *DrgResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.DrgLifecycleStateTerminating),
	}
}

func (s *DrgResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.DrgLifecycleStateTerminated),
	}
}

func (s *DrgResourceCrud) Create() error {
	request := oci_core.CreateDrgRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateDrg(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Drg
	return nil
}

func (s *DrgResourceCrud) Get() error {
	request := oci_core.GetDrgRequest{}

	tmp := s.D.Id()
	request.DrgId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetDrg(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Drg
	return nil
}

func (s *DrgResourceCrud) Update() error {
	request := oci_core.UpdateDrgRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	tmp := s.D.Id()
	request.DrgId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateDrg(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Drg
	return nil
}

func (s *DrgResourceCrud) Delete() error {
	request := oci_core.DeleteDrgRequest{}

	tmp := s.D.Id()
	request.DrgId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteDrg(context.Background(), request)
	return err
}

func (s *DrgResourceCrud) SetData() {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("time_created", s.Res.TimeCreated.String())

}
