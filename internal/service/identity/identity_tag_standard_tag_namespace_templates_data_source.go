// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"

	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v58/identity"
)

func IdentityTagStandardTagNamespaceTemplatesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityTagStandardTagNamespaceTemplates,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"standard_tag_namespace_templates": {
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
						"standard_tag_namespace_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readIdentityTagStandardTagNamespaceTemplates(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityTagStandardTagNamespaceTemplatesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*tf_client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

type IdentityTagStandardTagNamespaceTemplatesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListStandardTagNamespacesResponse
}

func (s *IdentityTagStandardTagNamespaceTemplatesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityTagStandardTagNamespaceTemplatesDataSourceCrud) Get() error {
	request := oci_identity.ListStandardTagNamespacesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity")

	response, err := s.Client.ListStandardTagNamespaces(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListStandardTagNamespaces(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *IdentityTagStandardTagNamespaceTemplatesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentityTagStandardTagNamespaceTemplatesDataSource-", IdentityTagStandardTagNamespaceTemplatesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		tagStandardTagNamespaceTemplate := map[string]interface{}{}

		if r.Description != nil {
			tagStandardTagNamespaceTemplate["description"] = *r.Description
		}

		if r.StandardTagNamespaceName != nil {
			tagStandardTagNamespaceTemplate["standard_tag_namespace_name"] = *r.StandardTagNamespaceName
		}

		if r.Status != nil {
			tagStandardTagNamespaceTemplate["status"] = *r.Status
		}

		resources = append(resources, tagStandardTagNamespaceTemplate)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, IdentityTagStandardTagNamespaceTemplatesDataSource().Schema["standard_tag_namespace_templates"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("standard_tag_namespace_templates", resources); err != nil {
		return err
	}

	return nil
}
