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

func CoreDrgRouteDistributionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreDrgRouteDistributions,
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
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"drg_route_distributions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreDrgRouteDistributionResource()),
			},
		},
	}
}

func readCoreDrgRouteDistributions(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgRouteDistributionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreDrgRouteDistributionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListDrgRouteDistributionsResponse
}

func (s *CoreDrgRouteDistributionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreDrgRouteDistributionsDataSourceCrud) Get() error {
	request := oci_core.ListDrgRouteDistributionsRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if drgId, ok := s.D.GetOkExists("drg_id"); ok {
		tmp := drgId.(string)
		request.DrgId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.DrgRouteDistributionLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListDrgRouteDistributions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDrgRouteDistributions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreDrgRouteDistributionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreDrgRouteDistributionsDataSource-", CoreDrgRouteDistributionsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		drgRouteDistribution := map[string]interface{}{
			"drg_id": *r.DrgId,
		}

		if r.CompartmentId != nil {
			drgRouteDistribution["compartment_id"] = *r.CompartmentId
		}

		if r.DefinedTags != nil {
			drgRouteDistribution["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			drgRouteDistribution["display_name"] = *r.DisplayName
		}

		drgRouteDistribution["distribution_type"] = r.DistributionType

		drgRouteDistribution["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			drgRouteDistribution["id"] = *r.Id
		}

		drgRouteDistribution["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			drgRouteDistribution["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, drgRouteDistribution)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreDrgRouteDistributionsDataSource().Schema["drg_route_distributions"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("drg_route_distributions", resources); err != nil {
		return err
	}

	return nil
}
