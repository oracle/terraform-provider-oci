// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osp_gateway

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/ospgateway"
	oci_osp_gateway "github.com/oracle/oci-go-sdk/v65/ospgateway"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OspGatewayInvoicesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOspGatewayInvoices,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"invoice_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"osp_home_region": {
				Type:     schema.TypeString,
				Required: true,
			},
			"search_text": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"time_invoice_end": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_invoice_start": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_payment_end": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_payment_start": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"invoice_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"bill_to_address": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"address_line1": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"address_line2": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"address_line3": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"address_line4": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"city": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"company_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"contact_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"country": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"ascii3country_code": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"country_code": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"country_id": {
																Type:     schema.TypeFloat,
																Computed: true,
															},
															"country_name": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"language_id": {
																Type:     schema.TypeFloat,
																Computed: true,
															},
														},
													},
												},
												"county": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"postal_code": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"province": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"state": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"street_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"street_number": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"currency": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"currency_code": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"currency_symbol": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"round_decimal_point": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"usd_conversion": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
											},
										},
									},
									"internal_invoice_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"invoice_amount": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"invoice_amount_adjusted": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"invoice_amount_applied": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"invoice_amount_credited": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"invoice_amount_due": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"invoice_amount_in_dispute": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"invoice_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"invoice_number": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"invoice_po_number": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"invoice_ref_number": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"invoice_status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"invoice_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_credit_card_payable": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_display_download_pdf": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_display_view_pdf": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_paid": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_payable": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_payment_failed": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_pdf_email_available": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"last_payment_detail": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"account_number": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"amount_paid": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"card_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"credit_card_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"echeck_routing": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"last_digits": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"name_on_card": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"paid_by": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"payment_method": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"paypal_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"paypal_reference": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"routing_number": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"time_expiration": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"time_paid_on": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"payment_terms": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"party_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"preferred_email": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"subscription_ids": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"tax": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"time_invoice": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_invoice_due": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readOspGatewayInvoices(d *schema.ResourceData, m interface{}) error {
	sync := &OspGatewayInvoicesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).InvoiceServiceClient()

	return tfresource.ReadResource(sync)
}

type OspGatewayInvoicesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_osp_gateway.InvoiceServiceClient
	Res    *oci_osp_gateway.ListInvoicesResponse
}

func (s *OspGatewayInvoicesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OspGatewayInvoicesDataSourceCrud) Get() error {
	request := oci_osp_gateway.ListInvoicesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if invoiceId, ok := s.D.GetOkExists("invoice_id"); ok {
		tmp := invoiceId.(string)
		request.InvoiceId = &tmp
	}

	if ospHomeRegion, ok := s.D.GetOkExists("osp_home_region"); ok {
		tmp := ospHomeRegion.(string)
		request.OspHomeRegion = &tmp
	}

	if searchText, ok := s.D.GetOkExists("search_text"); ok {
		tmp := searchText.(string)
		request.SearchText = &tmp
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		interfaces := status.([]interface{})
		tmp := make([]ospgateway.ListInvoicesStatusEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = ospgateway.ListInvoicesStatusEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("status") {
			request.Status = tmp
		}
	}

	if timeInvoiceEnd, ok := s.D.GetOkExists("time_invoice_end"); ok {
		tmp, err := time.Parse(time.RFC3339, timeInvoiceEnd.(string))
		if err != nil {
			return err
		}
		request.TimeInvoiceEnd = &oci_common.SDKTime{Time: tmp}
	}

	if timeInvoiceStart, ok := s.D.GetOkExists("time_invoice_start"); ok {
		tmp, err := time.Parse(time.RFC3339, timeInvoiceStart.(string))
		if err != nil {
			return err
		}
		request.TimeInvoiceStart = &oci_common.SDKTime{Time: tmp}
	}

	if timePaymentEnd, ok := s.D.GetOkExists("time_payment_end"); ok {
		tmp, err := time.Parse(time.RFC3339, timePaymentEnd.(string))
		if err != nil {
			return err
		}
		request.TimePaymentEnd = &oci_common.SDKTime{Time: tmp}
	}

	if timePaymentStart, ok := s.D.GetOkExists("time_payment_start"); ok {
		tmp, err := time.Parse(time.RFC3339, timePaymentStart.(string))
		if err != nil {
			return err
		}
		request.TimePaymentStart = &oci_common.SDKTime{Time: tmp}
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		interfaces := type_.([]interface{})
		tmp := make([]ospgateway.ListInvoicesTypeEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = ospgateway.ListInvoicesTypeEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("type") {
			request.Type = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "osp_gateway")

	response, err := s.Client.ListInvoices(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListInvoices(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OspGatewayInvoicesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OspGatewayInvoicesDataSource-", OspGatewayInvoicesDataSource(), s.D))
	resources := []map[string]interface{}{}
	invoice := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, InvoiceSummaryToMap(item))
	}
	invoice["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OspGatewayInvoicesDataSource().Schema["invoice_collection"].Elem.(*schema.Resource).Schema)
		invoice["items"] = items
	}

	resources = append(resources, invoice)
	if err := s.D.Set("invoice_collection", resources); err != nil {
		return err
	}

	return nil
}

func InvoiceSummaryToMap(obj oci_osp_gateway.InvoiceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Currency != nil {
		result["currency"] = []interface{}{CurrencyToMap(obj.Currency)}
	}

	if obj.InternalInvoiceId != nil {
		result["internal_invoice_id"] = string(*obj.InternalInvoiceId)
	}

	if obj.InvoiceAmount != nil {
		result["invoice_amount"] = float32(*obj.InvoiceAmount)
	}

	if obj.InvoiceAmountAdjusted != nil {
		result["invoice_amount_adjusted"] = float32(*obj.InvoiceAmountAdjusted)
	}

	if obj.InvoiceAmountApplied != nil {
		result["invoice_amount_applied"] = float32(*obj.InvoiceAmountApplied)
	}

	if obj.InvoiceAmountCredited != nil {
		result["invoice_amount_credited"] = float32(*obj.InvoiceAmountCredited)
	}

	if obj.InvoiceAmountDue != nil {
		result["invoice_amount_due"] = float32(*obj.InvoiceAmountDue)
	}

	if obj.InvoiceAmountInDispute != nil {
		result["invoice_amount_in_dispute"] = float32(*obj.InvoiceAmountInDispute)
	}

	if obj.InvoiceId != nil {
		result["invoice_id"] = string(*obj.InvoiceId)
	}

	if obj.InvoiceNumber != nil {
		result["invoice_number"] = string(*obj.InvoiceNumber)
	}

	if obj.InvoicePoNumber != nil {
		result["invoice_po_number"] = string(*obj.InvoicePoNumber)
	}

	if obj.InvoiceRefNumber != nil {
		result["invoice_ref_number"] = string(*obj.InvoiceRefNumber)
	}

	result["invoice_status"] = string(obj.InvoiceStatus)

	result["invoice_type"] = string(obj.InvoiceType)

	if obj.IsCreditCardPayable != nil {
		result["is_credit_card_payable"] = bool(*obj.IsCreditCardPayable)
	}

	if obj.IsDisplayDownloadPdf != nil {
		result["is_display_download_pdf"] = bool(*obj.IsDisplayDownloadPdf)
	}

	if obj.IsDisplayViewPdf != nil {
		result["is_display_view_pdf"] = bool(*obj.IsDisplayViewPdf)
	}

	if obj.IsPaid != nil {
		result["is_paid"] = bool(*obj.IsPaid)
	}

	if obj.IsPayable != nil {
		result["is_payable"] = bool(*obj.IsPayable)
	}

	if obj.IsPaymentFailed != nil {
		result["is_payment_failed"] = bool(*obj.IsPaymentFailed)
	}

	if obj.IsPdfEmailAvailable != nil {
		result["is_pdf_email_available"] = bool(*obj.IsPdfEmailAvailable)
	}

	if obj.LastPaymentDetail != nil {
		lastPaymentDetailArray := []interface{}{}
		if lastPaymentDetailMap := PaymentDetailToMap(&obj.LastPaymentDetail); lastPaymentDetailMap != nil {
			lastPaymentDetailArray = append(lastPaymentDetailArray, lastPaymentDetailMap)
		}
		result["last_payment_detail"] = lastPaymentDetailArray
	}

	if obj.PartyName != nil {
		result["party_name"] = string(*obj.PartyName)
	}

	result["subscription_ids"] = obj.SubscriptionIds

	if obj.TimeInvoice != nil {
		result["time_invoice"] = obj.TimeInvoice.String()
	}

	if obj.TimeInvoiceDue != nil {
		result["time_invoice_due"] = obj.TimeInvoiceDue.String()
	}

	return result
}
