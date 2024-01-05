// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_apigateway "github.com/oracle/oci-go-sdk/v65/apigateway"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApigatewayDeploymentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readApigatewayDeployments,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"gateway_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"deployment_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(ApigatewayDeploymentResource()),
			},
		},
	}
}

func readApigatewayDeployments(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewayDeploymentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DeploymentClient()

	return tfresource.ReadResource(sync)
}

type ApigatewayDeploymentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apigateway.DeploymentClient
	Res    *oci_apigateway.ListDeploymentsResponse
}

func (s *ApigatewayDeploymentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApigatewayDeploymentsDataSourceCrud) Get() error {
	request := oci_apigateway.ListDeploymentsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if gatewayId, ok := s.D.GetOkExists("gateway_id"); ok {
		tmp := gatewayId.(string)
		request.GatewayId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_apigateway.DeploymentLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apigateway")

	listResponse, err := s.Client.ListDeployments(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &listResponse
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDeployments(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ApigatewayDeploymentsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ApigatewayDeploymentsDataSource-", ApigatewayDeploymentsDataSource(), s.D))

	resources := []map[string]interface{}{}

	for _, item := range s.Res.Items {
		resources = append(resources, DeploymentSummaryToMap(item))
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, ApigatewayDeploymentsDataSource().Schema["deployment_collection"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("deployment_collection", resources); err != nil {
		return err
	}

	return nil
}
