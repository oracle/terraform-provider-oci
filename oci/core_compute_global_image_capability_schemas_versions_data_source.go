// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v46/core"
)

func init() {
	RegisterDatasource("oci_core_compute_global_image_capability_schemas_versions", CoreComputeGlobalImageCapabilitySchemasVersionsDataSource())
}

func CoreComputeGlobalImageCapabilitySchemasVersionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreComputeGlobalImageCapabilitySchemasVersions,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compute_global_image_capability_schema_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compute_global_image_capability_schema_versions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"compute_global_image_capability_schema_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
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
				},
			},
		},
	}
}

func readCoreComputeGlobalImageCapabilitySchemasVersions(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeGlobalImageCapabilitySchemasVersionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient()

	return ReadResource(sync)
}

type CoreComputeGlobalImageCapabilitySchemasVersionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListComputeGlobalImageCapabilitySchemaVersionsResponse
}

func (s *CoreComputeGlobalImageCapabilitySchemasVersionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreComputeGlobalImageCapabilitySchemasVersionsDataSourceCrud) Get() error {
	request := oci_core.ListComputeGlobalImageCapabilitySchemaVersionsRequest{}

	if computeGlobalImageCapabilitySchemaId, ok := s.D.GetOkExists("compute_global_image_capability_schema_id"); ok {
		tmp := computeGlobalImageCapabilitySchemaId.(string)
		request.ComputeGlobalImageCapabilitySchemaId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListComputeGlobalImageCapabilitySchemaVersions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListComputeGlobalImageCapabilitySchemaVersions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreComputeGlobalImageCapabilitySchemasVersionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("CoreComputeGlobalImageCapabilitySchemasVersionsDataSource-", CoreComputeGlobalImageCapabilitySchemasVersionsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		computeGlobalImageCapabilitySchemasVersion := map[string]interface{}{
			"compute_global_image_capability_schema_id": *r.ComputeGlobalImageCapabilitySchemaId,
		}

		if r.DisplayName != nil {
			computeGlobalImageCapabilitySchemasVersion["display_name"] = *r.DisplayName
		}

		if r.Name != nil {
			computeGlobalImageCapabilitySchemasVersion["name"] = *r.Name
		}

		if r.TimeCreated != nil {
			computeGlobalImageCapabilitySchemasVersion["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, computeGlobalImageCapabilitySchemasVersion)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, CoreComputeGlobalImageCapabilitySchemasVersionsDataSource().Schema["compute_global_image_capability_schema_versions"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("compute_global_image_capability_schema_versions", resources); err != nil {
		return err
	}

	return nil
}
