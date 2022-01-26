// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v56/core"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CoreImageDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["image_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CoreImageResource(), fieldMap, readSingularCoreImage)
}

func readSingularCoreImage(d *schema.ResourceData, m interface{}) error {
	sync := &CoreImageDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreImageDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.GetImageResponse
}

func (s *CoreImageDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreImageDataSourceCrud) Get() error {
	request := oci_core.GetImageRequest{}

	if imageId, ok := s.D.GetOkExists("image_id"); ok {
		tmp := imageId.(string)
		request.ImageId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetImage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreImageDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AgentFeatures != nil {
		s.D.Set("agent_features", []interface{}{InstanceAgentFeaturesToMap(s.Res.AgentFeatures)})
	} else {
		s.D.Set("agent_features", nil)
	}

	if s.Res.BaseImageId != nil {
		s.D.Set("base_image_id", *s.Res.BaseImageId)
	}

	if s.Res.BillableSizeInGBs != nil {
		s.D.Set("billable_size_in_gbs", strconv.FormatInt(*s.Res.BillableSizeInGBs, 10))
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreateImageAllowed != nil {
		s.D.Set("create_image_allowed", *s.Res.CreateImageAllowed)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("launch_mode", s.Res.LaunchMode)

	if s.Res.LaunchOptions != nil {
		s.D.Set("launch_options", []interface{}{LaunchOptionsToMap(s.Res.LaunchOptions)})
	} else {
		s.D.Set("launch_options", nil)
	}

	s.D.Set("listing_type", s.Res.ListingType)

	if s.Res.OperatingSystem != nil {
		s.D.Set("operating_system", *s.Res.OperatingSystem)
	}

	if s.Res.OperatingSystemVersion != nil {
		s.D.Set("operating_system_version", *s.Res.OperatingSystemVersion)
	}

	if s.Res.SizeInMBs != nil {
		s.D.Set("size_in_mbs", strconv.FormatInt(*s.Res.SizeInMBs, 10))
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
