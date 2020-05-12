// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_apigateway "github.com/oracle/oci-go-sdk/apigateway"
)

func init() {
	RegisterDatasource("oci_apigateway_gateways", ApigatewayGatewaysDataSource())
}

func ApigatewayGatewaysDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readApigatewayGateways,
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

			"gateway_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(ApigatewayGatewayResource()),
			},
		},
	}
}

func readApigatewayGateways(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewayGatewaysDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).gatewayClient()

	return ReadResource(sync)
}

type ApigatewayGatewaysDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apigateway.GatewayClient
	Res    *oci_apigateway.ListGatewaysResponse
}

func (s *ApigatewayGatewaysDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApigatewayGatewaysDataSourceCrud) Get() error {
	request := oci_apigateway.ListGatewaysRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_apigateway.GatewayLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "apigateway")

	response, err := s.Client.ListGateways(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ApigatewayGatewaysDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())

	resources := []map[string]interface{}{}
	for _, item := range s.Res.Items {
		resources = append(resources, GatewaySummaryToMap(item))
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, ApigatewayGatewaysDataSource().Schema["gateway_collection"].Elem.(*schema.Resource).Schema)
	}

	s.D.Set("gateway_collection", resources)

	return nil
}
