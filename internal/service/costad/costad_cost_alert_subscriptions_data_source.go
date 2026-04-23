// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package costad

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_costad "github.com/oracle/oci-go-sdk/v65/costad"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CostadCostAlertSubscriptionsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readCostadCostAlertSubscriptionsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cost_alert_subscription_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(CostadCostAlertSubscriptionResource()),
						},
					},
				},
			},
		},
	}
}

func readCostadCostAlertSubscriptionsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &CostadCostAlertSubscriptionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CustomerCostAdClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type CostadCostAlertSubscriptionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_costad.CostAdClient
	Res    *oci_costad.ListCostAlertSubscriptionsResponse
}

func (s *CostadCostAlertSubscriptionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CostadCostAlertSubscriptionsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_costad.ListCostAlertSubscriptionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_costad.CostAlertSubscriptionLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "costad")

	response, err := s.Client.ListCostAlertSubscriptions(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCostAlertSubscriptions(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CostadCostAlertSubscriptionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CostadCostAlertSubscriptionsDataSource-", CostadCostAlertSubscriptionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	costAlertSubscription := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CostAlertSubscriptionSummaryToMap(item))
	}
	costAlertSubscription["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CostadCostAlertSubscriptionsDataSource().Schema["cost_alert_subscription_collection"].Elem.(*schema.Resource).Schema)
		costAlertSubscription["items"] = items
	}

	resources = append(resources, costAlertSubscription)
	if err := s.D.Set("cost_alert_subscription_collection", resources); err != nil {
		return err
	}

	return nil
}
