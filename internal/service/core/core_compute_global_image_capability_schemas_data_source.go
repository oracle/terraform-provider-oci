// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v58/core"
)

func CoreComputeGlobalImageCapabilitySchemasDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreComputeGlobalImageCapabilitySchemas,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compute_global_image_capability_schemas": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"current_version_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"defined_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
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

func readCoreComputeGlobalImageCapabilitySchemas(d *schema.ResourceData, m interface{}) error {
	sync := &CoreComputeGlobalImageCapabilitySchemasDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()

	return tfresource.ReadResource(sync)
}

type CoreComputeGlobalImageCapabilitySchemasDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListComputeGlobalImageCapabilitySchemasResponse
}

func (s *CoreComputeGlobalImageCapabilitySchemasDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreComputeGlobalImageCapabilitySchemasDataSourceCrud) Get() error {
	request := oci_core.ListComputeGlobalImageCapabilitySchemasRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListComputeGlobalImageCapabilitySchemas(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListComputeGlobalImageCapabilitySchemas(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreComputeGlobalImageCapabilitySchemasDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreComputeGlobalImageCapabilitySchemasDataSource-", CoreComputeGlobalImageCapabilitySchemasDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		computeGlobalImageCapabilitySchema := map[string]interface{}{}

		if r.CompartmentId != nil {
			computeGlobalImageCapabilitySchema["compartment_id"] = *r.CompartmentId
		}

		if r.CurrentVersionName != nil {
			computeGlobalImageCapabilitySchema["current_version_name"] = *r.CurrentVersionName
		}

		if r.DefinedTags != nil {
			computeGlobalImageCapabilitySchema["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			computeGlobalImageCapabilitySchema["display_name"] = *r.DisplayName
		}

		computeGlobalImageCapabilitySchema["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			computeGlobalImageCapabilitySchema["id"] = *r.Id
		}

		if r.TimeCreated != nil {
			computeGlobalImageCapabilitySchema["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, computeGlobalImageCapabilitySchema)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreComputeGlobalImageCapabilitySchemasDataSource().Schema["compute_global_image_capability_schemas"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("compute_global_image_capability_schemas", resources); err != nil {
		return err
	}

	return nil
}
