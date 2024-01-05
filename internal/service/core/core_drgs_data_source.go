// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreDrgsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreDrgs,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"drgs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreDrgResource()),
			},
		},
	}
}

func readCoreDrgs(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

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

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreDrgsDataSource-", CoreDrgsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		drg := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefaultDrgRouteTables != nil {
			drg["default_drg_route_tables"] = []interface{}{DefaultDrgRouteTablesToMap(r.DefaultDrgRouteTables)}
		} else {
			drg["default_drg_route_tables"] = nil
		}

		if r.DefaultExportDrgRouteDistributionId != nil {
			drg["default_export_drg_route_distribution_id"] = *r.DefaultExportDrgRouteDistributionId
		}

		if r.DefinedTags != nil {
			drg["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
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
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreDrgsDataSource().Schema["drgs"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("drgs", resources); err != nil {
		return err
	}

	return nil
}
