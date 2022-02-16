// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v58/core"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CoreDrgRouteTableRouteRulesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreDrgRouteTableRouteRules,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"drg_route_table_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"route_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"drg_route_rules": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"attributes": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"destination": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"destination_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_blackhole": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_conflict": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"next_hop_drg_attachment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_provenance": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readCoreDrgRouteTableRouteRules(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgRouteTableRouteRulesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreDrgRouteTableRouteRulesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListDrgRouteRulesResponse
}

func (s *CoreDrgRouteTableRouteRulesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreDrgRouteTableRouteRulesDataSourceCrud) Get() error {
	request := oci_core.ListDrgRouteRulesRequest{}

	if drgRouteTableId, ok := s.D.GetOkExists("drg_route_table_id"); ok {
		tmp := drgRouteTableId.(string)
		request.DrgRouteTableId = &tmp
	}

	if routeType, ok := s.D.GetOkExists("route_type"); ok {
		request.RouteType = oci_core.ListDrgRouteRulesRouteTypeEnum(routeType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListDrgRouteRules(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDrgRouteRules(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreDrgRouteTableRouteRulesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreDrgRouteTableRouteRulesDataSource-", CoreDrgRouteTableRouteRulesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		drgRouteTableRouteRule := map[string]interface{}{}

		drgRouteTableRouteRule["attributes"] = r.Attributes

		if r.Destination != nil {
			drgRouteTableRouteRule["destination"] = *r.Destination
		}

		drgRouteTableRouteRule["destination_type"] = r.DestinationType

		if r.Id != nil {
			drgRouteTableRouteRule["id"] = *r.Id
		}

		if r.IsBlackhole != nil {
			drgRouteTableRouteRule["is_blackhole"] = *r.IsBlackhole
		}

		if r.IsConflict != nil {
			drgRouteTableRouteRule["is_conflict"] = *r.IsConflict
		}

		if r.NextHopDrgAttachmentId != nil {
			drgRouteTableRouteRule["next_hop_drg_attachment_id"] = *r.NextHopDrgAttachmentId
		}

		drgRouteTableRouteRule["route_provenance"] = r.RouteProvenance

		drgRouteTableRouteRule["route_type"] = r.RouteType

		resources = append(resources, drgRouteTableRouteRule)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreDrgRouteTableRouteRulesDataSource().Schema["drg_route_rules"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("drg_route_rules", resources); err != nil {
		return err
	}

	return nil
}
