// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OpsiOperationsInsightsPrivateEndpointsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOpsiOperationsInsightsPrivateEndpoints,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_used_for_rac_dbs": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"opsi_private_endpoint_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"operations_insights_private_endpoint_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(OpsiOperationsInsightsPrivateEndpointResource()),
						},
					},
				},
			},
		},
	}
}

func readOpsiOperationsInsightsPrivateEndpoints(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiOperationsInsightsPrivateEndpointsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

type OpsiOperationsInsightsPrivateEndpointsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opsi.OperationsInsightsClient
	Res    *oci_opsi.ListOperationsInsightsPrivateEndpointsResponse
}

func (s *OpsiOperationsInsightsPrivateEndpointsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpsiOperationsInsightsPrivateEndpointsDataSourceCrud) Get() error {
	request := oci_opsi.ListOperationsInsightsPrivateEndpointsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if isUsedForRacDbs, ok := s.D.GetOkExists("is_used_for_rac_dbs"); ok {
		tmp := isUsedForRacDbs.(bool)
		request.IsUsedForRacDbs = &tmp
	}

	if opsiPrivateEndpointId, ok := s.D.GetOkExists("opsi_private_endpoint_id"); ok {
		tmp := opsiPrivateEndpointId.(string)
		request.OpsiPrivateEndpointId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		//request.LifecycleState = oci_opsi.ListOperationsInsightsPrivateEndpointsLifecycleStateEnum(state.(string))
		interfaces := state.([]interface{})
		tmp := make([]oci_opsi.OperationsInsightsPrivateEndpointLifecycleStateEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_opsi.OperationsInsightsPrivateEndpointLifecycleStateEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("state") {
			request.LifecycleState = tmp
		}
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "opsi")

	response, err := s.Client.ListOperationsInsightsPrivateEndpoints(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOperationsInsightsPrivateEndpoints(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OpsiOperationsInsightsPrivateEndpointsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OpsiOperationsInsightsPrivateEndpointsDataSource-", OpsiOperationsInsightsPrivateEndpointsDataSource(), s.D))
	resources := []map[string]interface{}{}
	operationsInsightsPrivateEndpoint := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OperationsInsightsPrivateEndpointSummaryToMap(item))
	}
	operationsInsightsPrivateEndpoint["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OpsiOperationsInsightsPrivateEndpointsDataSource().Schema["operations_insights_private_endpoint_collection"].Elem.(*schema.Resource).Schema)
		operationsInsightsPrivateEndpoint["items"] = items
	}

	resources = append(resources, operationsInsightsPrivateEndpoint)
	if err := s.D.Set("operations_insights_private_endpoint_collection", resources); err != nil {
		return err
	}

	return nil
}
