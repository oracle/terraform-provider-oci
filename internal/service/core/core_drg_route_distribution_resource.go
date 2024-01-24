// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"
)

func CoreDrgRouteDistributionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreDrgRouteDistribution,
		Read:     readCoreDrgRouteDistribution,
		Update:   updateCoreDrgRouteDistribution,
		Delete:   deleteCoreDrgRouteDistribution,
		Schema: map[string]*schema.Schema{
			// Required
			"distribution_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"drg_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
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

func createCoreDrgRouteDistribution(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgRouteDistributionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.CreateResource(d, sync)
}

func readCoreDrgRouteDistribution(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgRouteDistributionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

func updateCoreDrgRouteDistribution(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgRouteDistributionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreDrgRouteDistribution(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgRouteDistributionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CoreDrgRouteDistributionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.DrgRouteDistribution
	DisableNotFoundRetries bool
}

func (s *CoreDrgRouteDistributionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreDrgRouteDistributionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.DrgRouteDistributionLifecycleStateProvisioning),
	}
}

func (s *CoreDrgRouteDistributionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.DrgRouteDistributionLifecycleStateAvailable),
	}
}

func (s *CoreDrgRouteDistributionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.DrgRouteDistributionLifecycleStateTerminating),
	}
}

func (s *CoreDrgRouteDistributionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.DrgRouteDistributionLifecycleStateTerminated),
	}
}

func (s *CoreDrgRouteDistributionResourceCrud) Create() error {
	request := oci_core.CreateDrgRouteDistributionRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if distributionType, ok := s.D.GetOkExists("distribution_type"); ok {
		request.DistributionType = oci_core.CreateDrgRouteDistributionDetailsDistributionTypeEnum(distributionType.(string))
	}

	if drgId, ok := s.D.GetOkExists("drg_id"); ok {
		tmp := drgId.(string)
		request.DrgId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateDrgRouteDistribution(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DrgRouteDistribution
	return nil
}

func (s *CoreDrgRouteDistributionResourceCrud) Get() error {
	request := oci_core.GetDrgRouteDistributionRequest{}

	tmp := s.D.Id()
	request.DrgRouteDistributionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetDrgRouteDistribution(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DrgRouteDistribution
	return nil
}

func (s *CoreDrgRouteDistributionResourceCrud) Update() error {
	request := oci_core.UpdateDrgRouteDistributionRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok && s.D.HasChange("defined_tags") {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok && s.D.HasChange("display_name") {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	tmp := s.D.Id()
	request.DrgRouteDistributionId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok && s.D.HasChange("freeform_tags") {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateDrgRouteDistribution(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DrgRouteDistribution
	return nil
}

func (s *CoreDrgRouteDistributionResourceCrud) Delete() error {
	request := oci_core.DeleteDrgRouteDistributionRequest{}

	tmp := s.D.Id()
	request.DrgRouteDistributionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteDrgRouteDistribution(context.Background(), request)
	return err
}

func (s *CoreDrgRouteDistributionResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("distribution_type", s.Res.DistributionType)

	if s.Res.DrgId != nil {
		s.D.Set("drg_id", *s.Res.DrgId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
