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

func CoreComputeImageCapabilitySchemaDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["compute_image_capability_schema_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["is_merge_enabled"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CoreComputeImageCapabilitySchemaResource(), fieldMap, readSingularCoreComputeImageCapabilitySchema)
}

func readSingularCoreComputeImageCapabilitySchema(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeImageCapabilitySchemaDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreComputeImageCapabilitySchemaDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.GetComputeImageCapabilitySchemaResponse
}

func (s *CoreComputeImageCapabilitySchemaDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreComputeImageCapabilitySchemaDataSourceCrud) Get() error {
	request := oci_core.GetComputeImageCapabilitySchemaRequest{}

	if computeImageCapabilitySchemaId, ok := s.D.GetOkExists("compute_image_capability_schema_id"); ok {
		tmp := computeImageCapabilitySchemaId.(string)
		request.ComputeImageCapabilitySchemaId = &tmp
	}

	if isMergeEnabled, ok := s.D.GetOkExists("is_merge_enabled"); ok {
		tmp, err := strconv.ParseBool(isMergeEnabled.(string))
		if err != nil {
			return err
		}
		request.IsMergeEnabled = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetComputeImageCapabilitySchema(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreComputeImageCapabilitySchemaDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComputeGlobalImageCapabilitySchemaId != nil {
		s.D.Set("compute_global_image_capability_schema_id", *s.Res.ComputeGlobalImageCapabilitySchemaId)
	}

	if s.Res.ComputeGlobalImageCapabilitySchemaVersionName != nil {
		s.D.Set("compute_global_image_capability_schema_version_name", *s.Res.ComputeGlobalImageCapabilitySchemaVersionName)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.ImageId != nil {
		s.D.Set("image_id", *s.Res.ImageId)
	}

	s.D.Set("schema_data", schemaDataToMap(s.Res.SchemaData))

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
