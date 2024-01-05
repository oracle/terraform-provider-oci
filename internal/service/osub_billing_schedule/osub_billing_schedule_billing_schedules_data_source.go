// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osub_billing_schedule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_osub_billing_schedule "github.com/oracle/oci-go-sdk/v65/osubbillingschedule"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsubBillingScheduleBillingSchedulesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsubBillingScheduleBillingSchedules,
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
			"x_one_origin_region": {
				Type:     schema.TypeString,
				Optional: true,
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

func readOsubBillingScheduleBillingSchedules(d *schema.ResourceData, m interface{}) error {
	sync := &OsubBillingScheduleBillingSchedulesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BillingScheduleClient()

	return tfresource.ReadResource(sync)
}

type OsubBillingScheduleBillingSchedulesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_osub_billing_schedule.BillingScheduleClient
	Res    *oci_osub_billing_schedule.ListBillingSchedulesResponse
}

func (s *OsubBillingScheduleBillingSchedulesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsubBillingScheduleBillingSchedulesDataSourceCrud) Get() error {
	request := oci_osub_billing_schedule.ListBillingSchedulesRequest{}

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

	if xOneOriginRegion, ok := s.D.GetOkExists("x_one_origin_region"); ok {
		tmp := xOneOriginRegion.(string)
		request.XOneOriginRegion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "osub_billing_schedule")

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

func (s *OsubBillingScheduleBillingSchedulesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsubBillingScheduleBillingSchedulesDataSource-", OsubBillingScheduleBillingSchedulesDataSource(), s.D))
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
			billingSchedule["product"] = []interface{}{ProductToMap(r.Product)}
		} else {
			billingSchedule["product"] = nil
		}

		if r.Quantity != nil {
			billingSchedule["quantity"] = *r.Quantity
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
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, OsubBillingScheduleBillingSchedulesDataSource().Schema["billing_schedules"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("billing_schedules", resources); err != nil {
		return err
	}

	return nil
}

func ProductToMap(obj *oci_osub_billing_schedule.Product) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.PartNumber != nil {
		result["part_number"] = string(*obj.PartNumber)
	}

	return result
}
