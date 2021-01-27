// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v34/core"
)

func init() {
	RegisterDatasource("oci_core_drgs", CoreDrgsDataSource())
}

func CoreDrgsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreDrgs,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"drgs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(CoreDrgResource()),
			},
		},
	}
}

func readCoreDrgs(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return ReadResource(sync)
}

type CoreDrgsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListDrgsResponse
}

func (s *CoreDrgsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreDrgsDataSourceCrud) Get() error {
	request := oci_core.ListDrgsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListDrgs(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDrgs(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreDrgsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("CoreDrgsDataSource-", CoreDrgsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		drg := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			drg["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			drg["display_name"] = *r.DisplayName
		}

		drg["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			drg["id"] = *r.Id
		}

		drg["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			drg["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, drg)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, CoreDrgsDataSource().Schema["drgs"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("drgs", resources); err != nil {
		return err
	}

	return nil
}
