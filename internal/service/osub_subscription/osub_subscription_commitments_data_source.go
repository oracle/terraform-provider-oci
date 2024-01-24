// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osub_subscription

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_osub_subscription "github.com/oracle/oci-go-sdk/v65/osubsubscription"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsubSubscriptionCommitmentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsubSubscriptionCommitments,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subscribed_service_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"x_one_gateway_subscription_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"x_one_origin_region": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"commitments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"available_amount": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"funded_allocation_value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"quantity": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_end": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_start": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"used_amount": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readOsubSubscriptionCommitments(d *schema.ResourceData, m interface{}) error {
	sync := &OsubSubscriptionCommitmentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CommitmentClient()

	return tfresource.ReadResource(sync)
}

type OsubSubscriptionCommitmentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_osub_subscription.CommitmentClient
	Res    *oci_osub_subscription.ListCommitmentsResponse
}

func (s *OsubSubscriptionCommitmentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsubSubscriptionCommitmentsDataSourceCrud) Get() error {
	request := oci_osub_subscription.ListCommitmentsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if subscribedServiceId, ok := s.D.GetOkExists("subscribed_service_id"); ok {
		tmp := subscribedServiceId.(string)
		request.SubscribedServiceId = &tmp
	}

	if xOneGatewaySubscriptionId, ok := s.D.GetOkExists("x_one_gateway_subscription_id"); ok {
		tmp := xOneGatewaySubscriptionId.(string)
		request.XOneGatewaySubscriptionId = &tmp
	}

	if xOneOriginRegion, ok := s.D.GetOkExists("x_one_origin_region"); ok {
		tmp := xOneOriginRegion.(string)
		request.XOneOriginRegion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "osub_subscription")

	response, err := s.Client.ListCommitments(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCommitments(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsubSubscriptionCommitmentsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsubSubscriptionCommitmentsDataSource-", OsubSubscriptionCommitmentsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		commitment := map[string]interface{}{}

		if r.AvailableAmount != nil {
			commitment["available_amount"] = *r.AvailableAmount
		}

		if r.FundedAllocationValue != nil {
			commitment["funded_allocation_value"] = *r.FundedAllocationValue
		}

		if r.Id != nil {
			commitment["id"] = *r.Id
		}

		if r.Quantity != nil {
			commitment["quantity"] = *r.Quantity
		}

		if r.TimeEnd != nil {
			commitment["time_end"] = r.TimeEnd.String()
		}

		if r.TimeStart != nil {
			commitment["time_start"] = r.TimeStart.String()
		}

		if r.UsedAmount != nil {
			commitment["used_amount"] = *r.UsedAmount
		}

		resources = append(resources, commitment)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, OsubSubscriptionCommitmentsDataSource().Schema["commitments"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("commitments", resources); err != nil {
		return err
	}

	return nil
}
