// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package onesubscription

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_onesubscription "github.com/oracle/oci-go-sdk/v65/onesubscription"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OnesubscriptionOrganizationSubscriptionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOnesubscriptionOrganizationSubscriptions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"organization_subscriptions": {
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

func readOnesubscriptionOrganizationSubscriptions(d *schema.ResourceData, m interface{}) error {
	sync := &OnesubscriptionOrganizationSubscriptionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OrganizationSubscriptionRegionalClient()

	return tfresource.ReadResource(sync)
}

type OnesubscriptionOrganizationSubscriptionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_onesubscription.OrganizationSubscriptionClient
	Res    *oci_onesubscription.ListOrganizationSubscriptionsResponse
}

func (s *OnesubscriptionOrganizationSubscriptionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OnesubscriptionOrganizationSubscriptionsDataSourceCrud) Get() error {
	request := oci_onesubscription.ListOrganizationSubscriptionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "onesubscription")

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

func (s *OnesubscriptionOrganizationSubscriptionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OnesubscriptionOrganizationSubscriptionsDataSource-", OnesubscriptionOrganizationSubscriptionsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		organizationSubscription := map[string]interface{}{}

		if r.Currency != nil {
			organizationSubscription["currency"] = []interface{}{OrgnizationSubsCurrencyToMap(r.Currency)}
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
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, OnesubscriptionOrganizationSubscriptionsDataSource().Schema["organization_subscriptions"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("organization_subscriptions", resources); err != nil {
		return err
	}

	return nil
}

func OrgnizationSubsCurrencyToMap(obj *oci_onesubscription.OrgnizationSubsCurrency) map[string]interface{} {
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
