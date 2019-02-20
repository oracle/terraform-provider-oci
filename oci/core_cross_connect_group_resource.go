// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func CoreCrossConnectGroupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createCoreCrossConnectGroup,
		Read:     readCoreCrossConnectGroup,
		Update:   updateCoreCrossConnectGroup,
		Delete:   deleteCoreCrossConnectGroup,
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

func createCoreCrossConnectGroup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreCrossConnectGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return CreateResource(d, sync)
}

func readCoreCrossConnectGroup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreCrossConnectGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return ReadResource(sync)
}

func updateCoreCrossConnectGroup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreCrossConnectGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return UpdateResource(d, sync)
}

func deleteCoreCrossConnectGroup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreCrossConnectGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type CoreCrossConnectGroupResourceCrud struct {
	BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.CrossConnectGroup
	DisableNotFoundRetries bool
}

func (s *CoreCrossConnectGroupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreCrossConnectGroupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.CrossConnectGroupLifecycleStateProvisioning),
	}
}

func (s *CoreCrossConnectGroupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.CrossConnectGroupLifecycleStateInactive),
		string(oci_core.CrossConnectGroupLifecycleStateProvisioned),
	}
}

func (s *CoreCrossConnectGroupResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_core.CrossConnectGroupLifecycleStateProvisioning),
	}
}

func (s *CoreCrossConnectGroupResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_core.CrossConnectGroupLifecycleStateInactive),
		string(oci_core.CrossConnectGroupLifecycleStateProvisioned),
	}
}

func (s *CoreCrossConnectGroupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.CrossConnectGroupLifecycleStateTerminating),
	}
}

func (s *CoreCrossConnectGroupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.CrossConnectGroupLifecycleStateTerminated),
	}
}

func (s *CoreCrossConnectGroupResourceCrud) Create() error {
	request := oci_core.CreateCrossConnectGroupRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateCrossConnectGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CrossConnectGroup
	return nil
}

func (s *CoreCrossConnectGroupResourceCrud) Get() error {
	request := oci_core.GetCrossConnectGroupRequest{}

	tmp := s.D.Id()
	request.CrossConnectGroupId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetCrossConnectGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CrossConnectGroup
	return nil
}

func (s *CoreCrossConnectGroupResourceCrud) Update() error {
	request := oci_core.UpdateCrossConnectGroupRequest{}

	tmp := s.D.Id()
	request.CrossConnectGroupId = &tmp

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateCrossConnectGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CrossConnectGroup
	return nil
}

func (s *CoreCrossConnectGroupResourceCrud) Delete() error {
	request := oci_core.DeleteCrossConnectGroupRequest{}

	tmp := s.D.Id()
	request.CrossConnectGroupId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteCrossConnectGroup(context.Background(), request)
	return err
}

func (s *CoreCrossConnectGroupResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
