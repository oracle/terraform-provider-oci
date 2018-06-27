// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/identity"

	"github.com/oracle/terraform-provider-oci/crud"
)

func TagsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readTags,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"tag_namespace_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     TagResource(),
			},
		},
	}
}

func readTags(d *schema.ResourceData, m interface{}) error {
	sync := &TagsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.ReadResource(sync)
}

type TagsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListTagsResponse
}

func (s *TagsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *TagsDataSourceCrud) Get() error {
	request := oci_identity.ListTagsRequest{}

	if tagNamespaceId, ok := s.D.GetOkExists("tag_namespace_id"); ok {
		tmp := tagNamespaceId.(string)
		request.TagNamespaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

	response, err := s.Client.ListTags(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListTags(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *TagsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		tag := map[string]interface{}{}

		if r.DefinedTags != nil {
			tag["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			tag["description"] = *r.Description
		}

		tag["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			tag["id"] = *r.Id
		}

		if r.IsRetired != nil {
			tag["is_retired"] = *r.IsRetired
		}

		if r.Name != nil {
			tag["name"] = *r.Name
		}

		if r.TimeCreated != nil {
			tag["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, tag)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, TagsDataSource().Schema["tags"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("tags", resources); err != nil {
		panic(err)
	}

	return
}
