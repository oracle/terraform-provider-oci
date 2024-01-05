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

func CoreDrgRouteDistributionStatementsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreDrgRouteDistributionStatements,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"drg_route_distribution_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"drg_route_distribution_statements": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"action": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"match_criteria": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"attachment_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"drg_attachment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"match_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"priority": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readCoreDrgRouteDistributionStatements(d *schema.ResourceData, m interface{}) error {
	sync := &CoreDrgRouteDistributionStatementsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreDrgRouteDistributionStatementsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListDrgRouteDistributionStatementsResponse
}

func (s *CoreDrgRouteDistributionStatementsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreDrgRouteDistributionStatementsDataSourceCrud) Get() error {
	request := oci_core.ListDrgRouteDistributionStatementsRequest{}

	if drgRouteDistributionId, ok := s.D.GetOkExists("drg_route_distribution_id"); ok {
		tmp := drgRouteDistributionId.(string)
		request.DrgRouteDistributionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListDrgRouteDistributionStatements(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDrgRouteDistributionStatements(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreDrgRouteDistributionStatementsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreDrgRouteDistributionStatementsDataSource-", CoreDrgRouteDistributionStatementsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		drgRouteDistributionStatement := map[string]interface{}{}

		drgRouteDistributionStatement["action"] = r.Action

		if r.Id != nil {
			drgRouteDistributionStatement["id"] = *r.Id
		}

		matchCriteria := []interface{}{}
		for _, item := range r.MatchCriteria {
			matchCriteria = append(matchCriteria, DrgRouteDistributionMatchCriteriaToMap(item))
		}
		drgRouteDistributionStatement["match_criteria"] = matchCriteria

		if r.Priority != nil {
			drgRouteDistributionStatement["priority"] = *r.Priority
		}

		resources = append(resources, drgRouteDistributionStatement)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreDrgRouteDistributionStatementsDataSource().Schema["drg_route_distribution_statements"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("drg_route_distribution_statements", resources); err != nil {
		return err
	}

	return nil
}
