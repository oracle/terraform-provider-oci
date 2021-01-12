// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_core "github.com/oracle/oci-go-sdk/v32/core"
)

func init() {
	RegisterResource("oci_core_drg_attachment", CoreDrgAttachmentResource())
}

func CoreDrgAttachmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createCoreDrgAttachment,
		Read:     readCoreDrgAttachment,
		Update:   updateCoreDrgAttachment,
		Delete:   deleteCoreDrgAttachment,
		Schema: map[string]*schema.Schema{
			// Required
			"drg_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vcn_id": {
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
			"route_table_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"compartment_id": {
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

func createCoreDrgAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return CreateResource(d, sync)
}

func readCoreDrgAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return ReadResource(sync)
}

func updateCoreDrgAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return UpdateResource(d, sync)
}

func deleteCoreDrgAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgAttachmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type CoreDrgAttachmentResourceCrud struct {
	BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.DrgAttachment
	DisableNotFoundRetries bool
}

func (s *CoreDrgAttachmentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreDrgAttachmentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.DrgAttachmentLifecycleStateAttaching),
	}
}

func (s *CoreDrgAttachmentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.DrgAttachmentLifecycleStateAttached),
	}
}

func (s *CoreDrgAttachmentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.DrgAttachmentLifecycleStateDetaching),
	}
}

func (s *CoreDrgAttachmentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.DrgAttachmentLifecycleStateDetached),
	}
}

func (s *CoreDrgAttachmentResourceCrud) Create() error {
	request := oci_core.CreateDrgAttachmentRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if drgId, ok := s.D.GetOkExists("drg_id"); ok {
		tmp := drgId.(string)
		request.DrgId = &tmp
	}

	if routeTableId, ok := s.D.GetOkExists("route_table_id"); ok {
		tmp := routeTableId.(string)
		request.RouteTableId = &tmp
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateDrgAttachment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DrgAttachment
	return nil
}

func (s *CoreDrgAttachmentResourceCrud) Get() error {
	request := oci_core.GetDrgAttachmentRequest{}

	tmp := s.D.Id()
	request.DrgAttachmentId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetDrgAttachment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DrgAttachment
	return nil
}

func (s *CoreDrgAttachmentResourceCrud) Update() error {
	request := oci_core.UpdateDrgAttachmentRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	tmp := s.D.Id()
	request.DrgAttachmentId = &tmp

	if routeTableId, ok := s.D.GetOkExists("route_table_id"); ok {
		tmp := routeTableId.(string)
		request.RouteTableId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateDrgAttachment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DrgAttachment
	return nil
}

func (s *CoreDrgAttachmentResourceCrud) Delete() error {
	request := oci_core.DeleteDrgAttachmentRequest{}

	tmp := s.D.Id()
	request.DrgAttachmentId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteDrgAttachment(context.Background(), request)
	return err
}

func (s *CoreDrgAttachmentResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DrgId != nil {
		s.D.Set("drg_id", *s.Res.DrgId)
	}

	if s.Res.RouteTableId != nil {
		s.D.Set("route_table_id", *s.Res.RouteTableId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	return nil
}
