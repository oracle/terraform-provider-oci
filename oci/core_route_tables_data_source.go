// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v46/core"
)

func init() {
	RegisterDatasource("oci_core_route_tables", CoreRouteTablesDataSource())
}

func CoreRouteTablesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreRouteTables,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"route_tables": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(CoreRouteTableResource()),
			},
		},
	}
}

func readCoreRouteTables(d *schema.ResourceData, m interface{}) error {
	sync := &CoreRouteTablesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return ReadResource(sync)
}

type CoreRouteTablesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListRouteTablesResponse
}

func (s *CoreRouteTablesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreRouteTablesDataSourceCrud) Get() error {
	request := oci_core.ListRouteTablesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.RouteTableLifecycleStateEnum(state.(string))
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListRouteTables(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListRouteTables(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreRouteTablesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("CoreRouteTablesDataSource-", CoreRouteTablesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		routeTable := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			routeTable["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			routeTable["display_name"] = *r.DisplayName
		}

		routeTable["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			routeTable["id"] = *r.Id
		}

		routeRules := []interface{}{}
		for _, item := range r.RouteRules {
			routeRules = append(routeRules, RouteRuleToMap(item))
		}
		routeTable["route_rules"] = routeRules

		routeTable["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			routeTable["time_created"] = r.TimeCreated.String()
		}

		if r.VcnId != nil {
			routeTable["vcn_id"] = *r.VcnId
		}

		resources = append(resources, routeTable)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, CoreRouteTablesDataSource().Schema["route_tables"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("route_tables", resources); err != nil {
		return err
	}

	return nil
}
