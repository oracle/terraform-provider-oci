// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func CrossConnectGroupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createCrossConnectGroup,
		Read:     readCrossConnectGroup,
		Update:   updateCrossConnectGroup,
		Delete:   deleteCrossConnectGroup,
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

func createCrossConnectGroup(d *schema.ResourceData, m interface{}) error {
	sync := &CrossConnectGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return CreateResource(d, sync)
}

func readCrossConnectGroup(d *schema.ResourceData, m interface{}) error {
	sync := &CrossConnectGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return ReadResource(sync)
}

func updateCrossConnectGroup(d *schema.ResourceData, m interface{}) error {
	sync := &CrossConnectGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return UpdateResource(d, sync)
}

func deleteCrossConnectGroup(d *schema.ResourceData, m interface{}) error {
	sync := &CrossConnectGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type CrossConnectGroupResourceCrud struct {
	BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.CrossConnectGroup
	DisableNotFoundRetries bool
}

func (s *CrossConnectGroupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CrossConnectGroupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.CrossConnectGroupLifecycleStateProvisioning),
	}
}

func (s *CrossConnectGroupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.CrossConnectGroupLifecycleStateInactive),
		string(oci_core.CrossConnectGroupLifecycleStateProvisioned),
	}
}

func (s *CrossConnectGroupResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_core.CrossConnectGroupLifecycleStateProvisioning),
	}
}

func (s *CrossConnectGroupResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_core.CrossConnectGroupLifecycleStateInactive),
		string(oci_core.CrossConnectGroupLifecycleStateProvisioned),
	}
}

func (s *CrossConnectGroupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.CrossConnectGroupLifecycleStateTerminating),
	}
}

func (s *CrossConnectGroupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.CrossConnectGroupLifecycleStateTerminated),
	}
}

func (s *CrossConnectGroupResourceCrud) Create() error {
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

func (s *CrossConnectGroupResourceCrud) Get() error {
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

func (s *CrossConnectGroupResourceCrud) Update() error {
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

func (s *CrossConnectGroupResourceCrud) Delete() error {
	request := oci_core.DeleteCrossConnectGroupRequest{}

	tmp := s.D.Id()
	request.CrossConnectGroupId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteCrossConnectGroup(context.Background(), request)
	return err
}

func (s *CrossConnectGroupResourceCrud) SetData() error {
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
