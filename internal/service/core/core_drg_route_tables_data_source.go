// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v56/core"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CoreDrgRouteTablesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreDrgRouteTables,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"drg_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"import_drg_route_distribution_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"drg_route_tables": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreDrgRouteTableResource()),
			},
		},
	}
}

func readCoreDrgRouteTables(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgRouteTablesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreDrgRouteTablesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListDrgRouteTablesResponse
}

func (s *CoreDrgRouteTablesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreDrgRouteTablesDataSourceCrud) Get() error {
	request := oci_core.ListDrgRouteTablesRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if drgId, ok := s.D.GetOkExists("drg_id"); ok {
		tmp := drgId.(string)
		request.DrgId = &tmp
	}

	if importDrgRouteDistributionId, ok := s.D.GetOkExists("import_drg_route_distribution_id"); ok {
		tmp := importDrgRouteDistributionId.(string)
		request.ImportDrgRouteDistributionId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.DrgRouteTableLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListDrgRouteTables(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDrgRouteTables(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreDrgRouteTablesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreDrgRouteTablesDataSource-", CoreDrgRouteTablesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		drgRouteTable := map[string]interface{}{
			"drg_id": *r.DrgId,
		}

		if r.CompartmentId != nil {
			drgRouteTable["compartment_id"] = *r.CompartmentId
		}

		if r.DefinedTags != nil {
			drgRouteTable["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			drgRouteTable["display_name"] = *r.DisplayName
		}

		drgRouteTable["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			drgRouteTable["id"] = *r.Id
		}

		if r.ImportDrgRouteDistributionId != nil {
			drgRouteTable["import_drg_route_distribution_id"] = *r.ImportDrgRouteDistributionId
		}

		if r.IsEcmpEnabled != nil {
			drgRouteTable["is_ecmp_enabled"] = *r.IsEcmpEnabled
		}

		drgRouteTable["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			drgRouteTable["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, drgRouteTable)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreDrgRouteTablesDataSource().Schema["drg_route_tables"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("drg_route_tables", resources); err != nil {
		return err
	}

	return nil
}
