// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package media_services

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_media_services "github.com/oracle/oci-go-sdk/v65/mediaservices"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MediaServicesMediaWorkflowTaskDeclarationDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularMediaServicesMediaWorkflowTaskDeclaration,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_current": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			// Computed
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"parameters_schema": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"parameters_schema_allowing_references": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularMediaServicesMediaWorkflowTaskDeclaration(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesMediaWorkflowTaskDeclarationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.ReadResource(sync)
}

type MediaServicesMediaWorkflowTaskDeclarationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_media_services.MediaServicesClient
	Res    *oci_media_services.ListMediaWorkflowTaskDeclarationsResponse
}

func (s *MediaServicesMediaWorkflowTaskDeclarationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MediaServicesMediaWorkflowTaskDeclarationDataSourceCrud) Get() error {
	request := oci_media_services.ListMediaWorkflowTaskDeclarationsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if isCurrent, ok := s.D.GetOkExists("is_current"); ok {
		tmp := isCurrent.(bool)
		request.IsCurrent = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if version, ok := s.D.GetOkExists("version"); ok {
		tmp := version.(int)
		request.Version = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "media_services")

	response, err := s.Client.ListMediaWorkflowTaskDeclarations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MediaServicesMediaWorkflowTaskDeclarationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MediaServicesMediaWorkflowTaskDeclarationDataSource-", MediaServicesMediaWorkflowTaskDeclarationDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MediaWorkflowTaskDeclarationToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func MediaWorkflowTaskDeclarationToMap(obj oci_media_services.MediaWorkflowTaskDeclaration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ParametersSchema != nil {
		jsonStr, err := json.Marshal(obj.ParametersSchema)
		if err == nil {
			result["parameters_schema"] = string(jsonStr)
		}
	}

	if obj.ParametersSchemaAllowingReferences != nil {
		jsonStr, err := json.Marshal(obj.ParametersSchemaAllowingReferences)
		if err == nil {
			result["parameters_schema_allowing_references"] = string(jsonStr)
		}
	}

	if obj.Version != nil {
		result["version"] = int(*obj.Version)
	}

	return result
}
