// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v65/identity"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IdentityTagNamespacesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityTagNamespaces,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"include_subcompartments": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tag_namespaces": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(IdentityTagNamespaceResource()),
			},
		},
	}
}

func readIdentityTagNamespaces(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityTagNamespacesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

type IdentityTagNamespacesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListTagNamespacesResponse
}

func (s *IdentityTagNamespacesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityTagNamespacesDataSourceCrud) Get() error {
	request := oci_identity.ListTagNamespacesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if includeSubcompartments, ok := s.D.GetOkExists("include_subcompartments"); ok {
		tmp := includeSubcompartments.(bool)
		request.IncludeSubcompartments = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_identity.TagNamespaceLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity")

	response, err := s.Client.ListTagNamespaces(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListTagNamespaces(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *IdentityTagNamespacesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentityTagNamespacesDataSource-", IdentityTagNamespacesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		tagNamespace := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			tagNamespace["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			tagNamespace["description"] = *r.Description
		}

		tagNamespace["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			tagNamespace["id"] = *r.Id
		}

		if r.IsRetired != nil {
			tagNamespace["is_retired"] = *r.IsRetired
		}

		if r.Name != nil {
			tagNamespace["name"] = *r.Name
		}

		tagNamespace["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			tagNamespace["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, tagNamespace)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, IdentityTagNamespacesDataSource().Schema["tag_namespaces"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("tag_namespaces", resources); err != nil {
		return err
	}

	return nil
}
