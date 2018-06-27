// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/identity"

	"github.com/oracle/terraform-provider-oci/crud"
)

func TagNamespacesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readTagNamespaces,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"include_subcompartments": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"tag_namespaces": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     TagNamespaceResource(),
			},
		},
	}
}

func readTagNamespaces(d *schema.ResourceData, m interface{}) error {
	sync := &TagNamespacesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.ReadResource(sync)
}

type TagNamespacesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListTagNamespacesResponse
}

func (s *TagNamespacesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *TagNamespacesDataSourceCrud) Get() error {
	request := oci_identity.ListTagNamespacesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if includeSubcompartments, ok := s.D.GetOkExists("include_subcompartments"); ok {
		tmp := includeSubcompartments.(bool)
		request.IncludeSubcompartments = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

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

func (s *TagNamespacesDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		tagNamespace := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			tagNamespace["defined_tags"] = definedTagsToMap(r.DefinedTags)
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

		if r.TimeCreated != nil {
			tagNamespace["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, tagNamespace)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, TagNamespacesDataSource().Schema["tag_namespaces"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("tag_namespaces", resources); err != nil {
		panic(err)
	}

	return
}
