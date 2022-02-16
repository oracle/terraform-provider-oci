// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v58/identity"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func IdentityTagsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityTags,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tag_namespace_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(IdentityTagResource()),
			},
		},
	}
}

func readIdentityTags(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityTagsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

type IdentityTagsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListTagsResponse
}

func (s *IdentityTagsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityTagsDataSourceCrud) Get() error {
	request := oci_identity.ListTagsRequest{}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_identity.TagLifecycleStateEnum(state.(string))
	}

	if tagNamespaceId, ok := s.D.GetOkExists("tag_namespace_id"); ok {
		tmp := tagNamespaceId.(string)
		request.TagNamespaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity")

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

func (s *IdentityTagsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentityTagsDataSource-", IdentityTagsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		tag := map[string]interface{}{}

		if r.DefinedTags != nil {
			tag["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			tag["description"] = *r.Description
		}

		tag["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			tag["id"] = *r.Id
		}

		if r.IsCostTracking != nil {
			tag["is_cost_tracking"] = *r.IsCostTracking
		}

		if r.IsRetired != nil {
			tag["is_retired"] = *r.IsRetired
		}

		if r.Name != nil {
			tag["name"] = *r.Name
		}

		tag["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			tag["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, tag)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, IdentityTagsDataSource().Schema["tags"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("tags", resources); err != nil {
		return err
	}

	return nil
}
