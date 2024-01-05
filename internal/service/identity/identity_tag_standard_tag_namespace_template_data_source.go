// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"

	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v65/identity"
)

func IdentityTagStandardTagNamespaceTemplateDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularIdentityTagStandardTagNamespaceTemplate,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"standard_tag_namespace_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tag_definition_templates": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"enum_mutability": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_cost_tracking": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"possible_values": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"tag_definition_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularIdentityTagStandardTagNamespaceTemplate(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityTagStandardTagNamespaceTemplateDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*tf_client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

type IdentityTagStandardTagNamespaceTemplateDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.GetStandardTagTemplateResponse
}

func (s *IdentityTagStandardTagNamespaceTemplateDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityTagStandardTagNamespaceTemplateDataSourceCrud) Get() error {
	request := oci_identity.GetStandardTagTemplateRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if standardTagNamespaceName, ok := s.D.GetOkExists("standard_tag_namespace_name"); ok {
		tmp := standardTagNamespaceName.(string)
		request.StandardTagNamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity")

	response, err := s.Client.GetStandardTagTemplate(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityTagStandardTagNamespaceTemplateDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentityTagStandardTagNamespaceTemplateDataSource-", IdentityTagStandardTagNamespaceTemplateDataSource(), s.D))

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.Status != nil {
		s.D.Set("status", *s.Res.Status)
	}

	tagDefinitionTemplates := []interface{}{}
	for _, item := range s.Res.TagDefinitionTemplates {
		tagDefinitionTemplates = append(tagDefinitionTemplates, StandardTagDefinitionTemplateToMap(item))
	}
	s.D.Set("tag_definition_templates", tagDefinitionTemplates)

	return nil
}

func StandardTagDefinitionTemplateToMap(obj oci_identity.StandardTagDefinitionTemplate) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["enum_mutability"] = string(obj.EnumMutability)

	if obj.IsCostTracking != nil {
		result["is_cost_tracking"] = bool(*obj.IsCostTracking)
	}

	result["possible_values"] = obj.PossibleValues

	if obj.TagDefinitionName != nil {
		result["tag_definition_name"] = string(*obj.TagDefinitionName)
	}

	result["type"] = string(obj.Type)

	return result
}
