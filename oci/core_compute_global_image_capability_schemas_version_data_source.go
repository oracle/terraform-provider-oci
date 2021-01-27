// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v34/core"
)

func init() {
	RegisterDatasource("oci_core_compute_global_image_capability_schemas_version", CoreComputeGlobalImageCapabilitySchemasVersionDataSource())
}

func CoreComputeGlobalImageCapabilitySchemasVersionDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCoreComputeGlobalImageCapabilitySchemasVersion,
		Schema: map[string]*schema.Schema{
			"compute_global_image_capability_schema_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compute_global_image_capability_schema_version_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"schema_data": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularCoreComputeGlobalImageCapabilitySchemasVersion(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeGlobalImageCapabilitySchemasVersionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient()

	return ReadResource(sync)
}

type CoreComputeGlobalImageCapabilitySchemasVersionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.GetComputeGlobalImageCapabilitySchemaVersionResponse
}

func (s *CoreComputeGlobalImageCapabilitySchemasVersionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreComputeGlobalImageCapabilitySchemasVersionDataSourceCrud) Get() error {
	request := oci_core.GetComputeGlobalImageCapabilitySchemaVersionRequest{}

	if computeGlobalImageCapabilitySchemaId, ok := s.D.GetOkExists("compute_global_image_capability_schema_id"); ok {
		tmp := computeGlobalImageCapabilitySchemaId.(string)
		request.ComputeGlobalImageCapabilitySchemaId = &tmp
	}

	if computeGlobalImageCapabilitySchemaVersionName, ok := s.D.GetOkExists("compute_global_image_capability_schema_version_name"); ok {
		tmp := computeGlobalImageCapabilitySchemaVersionName.(string)
		request.ComputeGlobalImageCapabilitySchemaVersionName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.GetComputeGlobalImageCapabilitySchemaVersion(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreComputeGlobalImageCapabilitySchemasVersionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("CoreComputeGlobalImageCapabilitySchemasVersionDataSource-", CoreComputeGlobalImageCapabilitySchemasVersionDataSource(), s.D))

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("schema_data", schemaDataToMap(s.Res.SchemaData))

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
