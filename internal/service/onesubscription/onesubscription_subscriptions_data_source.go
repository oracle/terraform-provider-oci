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

func OnesubscriptionSubscriptionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOnesubscriptionSubscriptions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"buyer_email": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_commit_info_required": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"plan_number": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subscription_id": {
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
						"hold_reason": {
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
						"subscribed_services": {
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
									"booking_opty_number": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"commitment_services": {
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
												"line_net_amount": {
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
											},
										},
									},
									"csi": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"data_center_region": {
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
									"is_intent_to_pay": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"net_unit_price": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"operation_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"order_number": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"original_promo_amount": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"partner_transaction_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"pricing_model": {
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
												"provisioning_group": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"unit_of_measure": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"program_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"promo_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"quantity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"term_value": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"term_value_uom": {
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
									"used_amount": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"time_end": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_hold_release_eta": {
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

func readOnesubscriptionSubscriptions(d *schema.ResourceData, m interface{}) error {
	sync := &OnesubscriptionSubscriptionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SubscriptionRegionalClient()

	return tfresource.ReadResource(sync)
}

type OnesubscriptionSubscriptionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_onesubscription.SubscriptionClient
	Res    *oci_onesubscription.ListSubscriptionsResponse
}

func (s *OnesubscriptionSubscriptionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OnesubscriptionSubscriptionsDataSourceCrud) Get() error {
	request := oci_onesubscription.ListSubscriptionsRequest{}

	if buyerEmail, ok := s.D.GetOkExists("buyer_email"); ok {
		tmp := buyerEmail.(string)
		request.BuyerEmail = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if isCommitInfoRequired, ok := s.D.GetOkExists("is_commit_info_required"); ok {
		tmp := isCommitInfoRequired.(bool)
		request.IsCommitInfoRequired = &tmp
	}

	if planNumber, ok := s.D.GetOkExists("plan_number"); ok {
		tmp := planNumber.(string)
		request.PlanNumber = &tmp
	}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "onesubscription")

	response, err := s.Client.ListSubscriptions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSubscriptions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OnesubscriptionSubscriptionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OnesubscriptionSubscriptionsDataSource-", OnesubscriptionSubscriptionsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		subscription := map[string]interface{}{}

		if r.Currency != nil {
			subscription["currency"] = []interface{}{SubscriptionCurrencyToMap(r.Currency)}
		} else {
			subscription["currency"] = nil
		}

		if r.HoldReason != nil {
			subscription["hold_reason"] = *r.HoldReason
		}

		if r.ServiceName != nil {
			subscription["service_name"] = *r.ServiceName
		}

		if r.Status != nil {
			subscription["status"] = *r.Status
		}

		subscribedServices := []interface{}{}
		for _, item := range r.SubscribedServices {
			subscribedServices = append(subscribedServices, SubscriptionSubscribedServiceToMap(item))
		}
		subscription["subscribed_services"] = subscribedServices

		if r.TimeEnd != nil {
			subscription["time_end"] = r.TimeEnd.String()
		}

		if r.TimeHoldReleaseEta != nil {
			subscription["time_hold_release_eta"] = r.TimeHoldReleaseEta.String()
		}

		if r.TimeStart != nil {
			subscription["time_start"] = r.TimeStart.String()
		}

		resources = append(resources, subscription)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, OnesubscriptionSubscriptionsDataSource().Schema["subscriptions"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("subscriptions", resources); err != nil {
		return err
	}

	return nil
}

func CommitmentServiceToMap(obj oci_onesubscription.CommitmentService) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailableAmount != nil {
		result["available_amount"] = string(*obj.AvailableAmount)
	}

	if obj.FundedAllocationValue != nil {
		result["funded_allocation_value"] = string(*obj.FundedAllocationValue)
	}

	if obj.LineNetAmount != nil {
		result["line_net_amount"] = string(*obj.LineNetAmount)
	}

	if obj.Quantity != nil {
		result["quantity"] = string(*obj.Quantity)
	}

	if obj.TimeEnd != nil {
		result["time_end"] = obj.TimeEnd.String()
	}

	if obj.TimeStart != nil {
		result["time_start"] = obj.TimeStart.String()
	}

	return result
}

func SubscriptionCurrencyToMap(obj *oci_onesubscription.SubscriptionCurrency) map[string]interface{} {
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

func SubscriptionProductToMap(obj *oci_onesubscription.SubscriptionProduct) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.PartNumber != nil {
		result["part_number"] = string(*obj.PartNumber)
	}

	if obj.ProvisioningGroup != nil {
		result["provisioning_group"] = string(*obj.ProvisioningGroup)
	}

	if obj.UnitOfMeasure != nil {
		result["unit_of_measure"] = string(*obj.UnitOfMeasure)
	}

	return result
}

func SubscriptionSubscribedServiceToMap(obj oci_onesubscription.SubscriptionSubscribedService) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailableAmount != nil {
		result["available_amount"] = string(*obj.AvailableAmount)
	}

	if obj.BookingOptyNumber != nil {
		result["booking_opty_number"] = string(*obj.BookingOptyNumber)
	}

	commitmentServices := []interface{}{}
	for _, item := range obj.CommitmentServices {
		commitmentServices = append(commitmentServices, CommitmentServiceToMap(item))
	}
	result["commitment_services"] = commitmentServices

	if obj.Csi != nil {
		result["csi"] = strconv.FormatInt(*obj.Csi, 10)
	}

	if obj.DataCenterRegion != nil {
		result["data_center_region"] = string(*obj.DataCenterRegion)
	}

	if obj.FundedAllocationValue != nil {
		result["funded_allocation_value"] = string(*obj.FundedAllocationValue)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsIntentToPay != nil {
		result["is_intent_to_pay"] = bool(*obj.IsIntentToPay)
	}

	if obj.NetUnitPrice != nil {
		result["net_unit_price"] = string(*obj.NetUnitPrice)
	}

	if obj.OperationType != nil {
		result["operation_type"] = string(*obj.OperationType)
	}

	if obj.OrderNumber != nil {
		result["order_number"] = strconv.FormatInt(*obj.OrderNumber, 10)
	}

	if obj.OriginalPromoAmount != nil {
		result["original_promo_amount"] = string(*obj.OriginalPromoAmount)
	}

	if obj.PartnerTransactionType != nil {
		result["partner_transaction_type"] = string(*obj.PartnerTransactionType)
	}

	if obj.PricingModel != nil {
		result["pricing_model"] = string(*obj.PricingModel)
	}

	if obj.Product != nil {
		result["product"] = []interface{}{SubscriptionProductToMap(obj.Product)}
	}

	if obj.ProgramType != nil {
		result["program_type"] = string(*obj.ProgramType)
	}

	if obj.PromoType != nil {
		result["promo_type"] = string(*obj.PromoType)
	}

	if obj.Quantity != nil {
		result["quantity"] = string(*obj.Quantity)
	}

	if obj.Status != nil {
		result["status"] = string(*obj.Status)
	}

	if obj.TermValue != nil {
		result["term_value"] = strconv.FormatInt(*obj.TermValue, 10)
	}

	if obj.TermValueUom != nil {
		result["term_value_uom"] = string(*obj.TermValueUom)
	}

	if obj.TimeEnd != nil {
		result["time_end"] = obj.TimeEnd.String()
	}

	if obj.TimeStart != nil {
		result["time_start"] = obj.TimeStart.String()
	}

	if obj.TotalValue != nil {
		result["total_value"] = string(*obj.TotalValue)
	}

	if obj.UsedAmount != nil {
		result["used_amount"] = string(*obj.UsedAmount)
	}

	return result
}
