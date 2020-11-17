// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v29/identity"
)

func init() {
	RegisterDatasource("oci_identity_tag_defaults", IdentityTagDefaultsDataSource())
}

func IdentityTagDefaultsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityTagDefaults,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tag_definition_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tag_defaults": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(IdentityTagDefaultResource()),
			},
		},
	}
}

func readIdentityTagDefaults(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityTagDefaultsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient()

	return ReadResource(sync)
}

type IdentityTagDefaultsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListTagDefaultsResponse
}

func (s *IdentityTagDefaultsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityTagDefaultsDataSourceCrud) Get() error {
	request := oci_identity.ListTagDefaultsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_identity.TagDefaultSummaryLifecycleStateEnum(state.(string))
	}

	if tagDefinitionId, ok := s.D.GetOkExists("tag_definition_id"); ok {
		tmp := tagDefinitionId.(string)
		request.TagDefinitionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

	response, err := s.Client.ListTagDefaults(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListTagDefaults(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *IdentityTagDefaultsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("IdentityTagDefaultsDataSource-", IdentityTagDefaultsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		tagDefault := map[string]interface{}{}

		if r.CompartmentId != nil {
			tagDefault["compartment_id"] = *r.CompartmentId
		}

		if r.Id != nil {
			tagDefault["id"] = *r.Id
		}

		if r.IsRequired != nil {
			tagDefault["is_required"] = *r.IsRequired
		}

		tagDefault["state"] = r.LifecycleState

		if r.TagDefinitionId != nil {
			tagDefault["tag_definition_id"] = *r.TagDefinitionId
		}

		if r.TagDefinitionName != nil {
			tagDefault["tag_definition_name"] = *r.TagDefinitionName
		}

		if r.TagNamespaceId != nil {
			tagDefault["tag_namespace_id"] = *r.TagNamespaceId
		}

		if r.TimeCreated != nil {
			tagDefault["time_created"] = r.TimeCreated.String()
		}

		if r.Value != nil {
			tagDefault["value"] = *r.Value
		}

		resources = append(resources, tagDefault)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, IdentityTagDefaultsDataSource().Schema["tag_defaults"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("tag_defaults", resources); err != nil {
		return err
	}

	return nil
}
