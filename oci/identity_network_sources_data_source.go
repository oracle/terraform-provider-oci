// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v25/identity"
)

func init() {
	RegisterDatasource("oci_identity_network_sources", IdentityNetworkSourcesDataSource())
}

func IdentityNetworkSourcesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityNetworkSources,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network_sources": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(IdentityNetworkSourceResource()),
			},
		},
	}
}

func readIdentityNetworkSources(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityNetworkSourcesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient()

	return ReadResource(sync)
}

type IdentityNetworkSourcesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListNetworkSourcesResponse
}

func (s *IdentityNetworkSourcesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityNetworkSourcesDataSourceCrud) Get() error {
	request := oci_identity.ListNetworkSourcesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

	response, err := s.Client.ListNetworkSources(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListNetworkSources(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *IdentityNetworkSourcesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		networkSource := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			networkSource["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			networkSource["description"] = *r.Description
		}

		networkSource["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			networkSource["id"] = *r.Id
		}

		if r.Name != nil {
			networkSource["name"] = *r.Name
		}

		networkSource["public_source_list"] = r.PublicSourceList

		networkSource["services"] = r.Services

		if r.TimeCreated != nil {
			networkSource["time_created"] = r.TimeCreated.String()
		}

		virtualSourceList := []interface{}{}
		for _, item := range r.VirtualSourceList {
			virtualSourceList = append(virtualSourceList, networkSourcesVirtualSourceListToMap(item))
		}
		networkSource["virtual_source_list"] = virtualSourceList

		resources = append(resources, networkSource)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, IdentityNetworkSourcesDataSource().Schema["network_sources"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("network_sources", resources); err != nil {
		return err
	}

	return nil
}
