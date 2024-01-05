// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osub_organization_subscription

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_osub_organization_subscription "github.com/oracle/oci-go-sdk/v65/osuborganizationsubscription"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsubOrganizationSubscriptionOrganizationSubscriptionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsubOrganizationSubscriptionOrganizationSubscriptions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subscription_ids": {
				Type:     schema.TypeString,
				Required: true,
			},
			"x_one_origin_region": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subscriptions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"currency": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"iso_code": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"std_precision": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"service_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
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
						"total_value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readOsubOrganizationSubscriptionOrganizationSubscriptions(d *schema.ResourceData, m interface{}) error {
	sync := &OsubOrganizationSubscriptionOrganizationSubscriptionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OrganizationSubscriptionClient()

	return tfresource.ReadResource(sync)
}

type OsubOrganizationSubscriptionOrganizationSubscriptionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_osub_organization_subscription.OrganizationSubscriptionClient
	Res    *oci_osub_organization_subscription.ListOrganizationSubscriptionsResponse
}

func (s *OsubOrganizationSubscriptionOrganizationSubscriptionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsubOrganizationSubscriptionOrganizationSubscriptionsDataSourceCrud) Get() error {
	request := oci_osub_organization_subscription.ListOrganizationSubscriptionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if subscriptionIds, ok := s.D.GetOkExists("subscription_ids"); ok {
		tmp := subscriptionIds.(string)
		request.SubscriptionIds = &tmp
	}

	if xOneOriginRegion, ok := s.D.GetOkExists("x_one_origin_region"); ok {
		tmp := xOneOriginRegion.(string)
		request.XOneOriginRegion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "osub_organization_subscription")

	response, err := s.Client.ListOrganizationSubscriptions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOrganizationSubscriptions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsubOrganizationSubscriptionOrganizationSubscriptionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsubOrganizationSubscriptionOrganizationSubscriptionsDataSource-", OsubOrganizationSubscriptionOrganizationSubscriptionsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		organizationSubscription := map[string]interface{}{}

		if r.Currency != nil {
			organizationSubscription["currency"] = []interface{}{CurrencyToMap(r.Currency)}
		} else {
			organizationSubscription["currency"] = nil
		}

		if r.Id != nil {
			organizationSubscription["id"] = *r.Id
		}

		if r.ServiceName != nil {
			organizationSubscription["service_name"] = *r.ServiceName
		}

		if r.Status != nil {
			organizationSubscription["status"] = *r.Status
		}

		if r.TimeEnd != nil {
			organizationSubscription["time_end"] = r.TimeEnd.String()
		}

		if r.TimeStart != nil {
			organizationSubscription["time_start"] = r.TimeStart.String()
		}

		if r.TotalValue != nil {
			organizationSubscription["total_value"] = *r.TotalValue
		}

		if r.Type != nil {
			organizationSubscription["type"] = *r.Type
		}

		resources = append(resources, organizationSubscription)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, OsubOrganizationSubscriptionOrganizationSubscriptionsDataSource().Schema["subscriptions"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("subscriptions", resources); err != nil {
		return err
	}

	return nil
}

func CurrencyToMap(obj *oci_osub_organization_subscription.Currency) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsoCode != nil {
		result["iso_code"] = string(*obj.IsoCode)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.StdPrecision != nil {
		result["std_precision"] = strconv.FormatInt(*obj.StdPrecision, 10)
	}

	return result
}
