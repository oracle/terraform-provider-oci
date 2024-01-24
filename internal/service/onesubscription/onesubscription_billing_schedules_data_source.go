// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package onesubscription

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_onesubscription "github.com/oracle/oci-go-sdk/v65/onesubscription"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OnesubscriptionBillingSchedulesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOnesubscriptionBillingSchedules,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subscribed_service_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"billing_schedules": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"amount": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ar_customer_transaction_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ar_invoice_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"billing_frequency": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"invoice_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"net_unit_price": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"order_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"product": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"part_number": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"quantity": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subscribed_service_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_end": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_invoicing": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_start": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readOnesubscriptionBillingSchedules(d *schema.ResourceData, m interface{}) error {
	sync := &OnesubscriptionBillingSchedulesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BillingScheduleRegionalClient()

	return tfresource.ReadResource(sync)
}

type OnesubscriptionBillingSchedulesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_onesubscription.BillingScheduleClient
	Res    *oci_onesubscription.ListBillingSchedulesResponse
}

func (s *OnesubscriptionBillingSchedulesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OnesubscriptionBillingSchedulesDataSourceCrud) Get() error {
	request := oci_onesubscription.ListBillingSchedulesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if subscribedServiceId, ok := s.D.GetOkExists("subscribed_service_id"); ok {
		tmp := subscribedServiceId.(string)
		request.SubscribedServiceId = &tmp
	}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "onesubscription")

	response, err := s.Client.ListBillingSchedules(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBillingSchedules(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OnesubscriptionBillingSchedulesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OnesubscriptionBillingSchedulesDataSource-", OnesubscriptionBillingSchedulesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		billingSchedule := map[string]interface{}{}

		if r.Amount != nil {
			billingSchedule["amount"] = *r.Amount
		}

		if r.ArCustomerTransactionId != nil {
			billingSchedule["ar_customer_transaction_id"] = *r.ArCustomerTransactionId
		}

		if r.ArInvoiceNumber != nil {
			billingSchedule["ar_invoice_number"] = *r.ArInvoiceNumber
		}

		if r.BillingFrequency != nil {
			billingSchedule["billing_frequency"] = *r.BillingFrequency
		}

		billingSchedule["invoice_status"] = r.InvoiceStatus

		if r.NetUnitPrice != nil {
			billingSchedule["net_unit_price"] = *r.NetUnitPrice
		}

		if r.OrderNumber != nil {
			billingSchedule["order_number"] = *r.OrderNumber
		}

		if r.Product != nil {
			billingSchedule["product"] = []interface{}{BillingScheduleProductToMap(r.Product)}
		} else {
			billingSchedule["product"] = nil
		}

		if r.Quantity != nil {
			billingSchedule["quantity"] = *r.Quantity
		}

		if r.SubscribedServiceId != nil {
			billingSchedule["subscribed_service_id"] = *r.SubscribedServiceId
		}

		if r.TimeEnd != nil {
			billingSchedule["time_end"] = r.TimeEnd.String()
		}

		if r.TimeInvoicing != nil {
			billingSchedule["time_invoicing"] = r.TimeInvoicing.String()
		}

		if r.TimeStart != nil {
			billingSchedule["time_start"] = r.TimeStart.String()
		}

		resources = append(resources, billingSchedule)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, OnesubscriptionBillingSchedulesDataSource().Schema["billing_schedules"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("billing_schedules", resources); err != nil {
		return err
	}

	return nil
}

func BillingScheduleProductToMap(obj *oci_onesubscription.BillingScheduleProduct) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.PartNumber != nil {
		result["part_number"] = string(*obj.PartNumber)
	}

	return result
}
