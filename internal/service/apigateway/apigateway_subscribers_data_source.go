// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_apigateway "github.com/oracle/oci-go-sdk/v65/apigateway"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApigatewaySubscribersDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readApigatewaySubscribersWithContext,
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
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subscriber_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(ApigatewaySubscriberResource()),
						},
					},
				},
			},
		},
	}
}

func readApigatewaySubscribersWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &ApigatewaySubscribersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SubscribersClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type ApigatewaySubscribersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apigateway.SubscribersClient
	Res    *oci_apigateway.ListSubscribersResponse
}

func (s *ApigatewaySubscribersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApigatewaySubscribersDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_apigateway.ListSubscribersRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_apigateway.SubscriberLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apigateway")

	response, err := s.Client.ListSubscribers(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSubscribers(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ApigatewaySubscribersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ApigatewaySubscribersDataSource-", ApigatewaySubscribersDataSource(), s.D))
	resources := []map[string]interface{}{}
	subscriber := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SubscriberSummaryToMap(item))
	}
	subscriber["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ApigatewaySubscribersDataSource().Schema["subscriber_collection"].Elem.(*schema.Resource).Schema)
		subscriber["items"] = items
	}

	resources = append(resources, subscriber)
	if err := s.D.Set("subscriber_collection", resources); err != nil {
		return err
	}

	return nil
}
