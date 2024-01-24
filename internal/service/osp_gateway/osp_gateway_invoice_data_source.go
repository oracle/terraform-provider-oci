// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osp_gateway

import (
	"context"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_osp_gateway "github.com/oracle/oci-go-sdk/v65/ospgateway"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OspGatewayInvoiceDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularOspGatewayInvoice,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"internal_invoice_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"osp_home_region": {
				Type:     schema.TypeString,
				Required: true,
			},
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
			"is_payable": {
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
	}
}

func readSingularOspGatewayInvoice(d *schema.ResourceData, m interface{}) error {
	sync := &OspGatewayInvoiceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).InvoiceServiceClient()

	return tfresource.ReadResource(sync)
}

type OspGatewayInvoiceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_osp_gateway.InvoiceServiceClient
	Res    *oci_osp_gateway.GetInvoiceResponse
}

func (s *OspGatewayInvoiceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OspGatewayInvoiceDataSourceCrud) Get() error {
	request := oci_osp_gateway.GetInvoiceRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if internalInvoiceId, ok := s.D.GetOkExists("internal_invoice_id"); ok {
		tmp := internalInvoiceId.(string)
		request.InternalInvoiceId = &tmp
	}

	if ospHomeRegion, ok := s.D.GetOkExists("osp_home_region"); ok {
		tmp := ospHomeRegion.(string)
		request.OspHomeRegion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "osp_gateway")

	response, err := s.Client.GetInvoice(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OspGatewayInvoiceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OspGatewayInvoiceDataSource-", OspGatewayInvoiceDataSource(), s.D))

	if s.Res.BillToAddress != nil {
		s.D.Set("bill_to_address", []interface{}{BillToAddressToMap(s.Res.BillToAddress)})
	} else {
		s.D.Set("bill_to_address", nil)
	}

	if s.Res.Currency != nil {
		s.D.Set("currency", []interface{}{CurrencyToMap(s.Res.Currency)})
	} else {
		s.D.Set("currency", nil)
	}

	if s.Res.InvoiceAmount != nil {
		s.D.Set("invoice_amount", *s.Res.InvoiceAmount)
	}

	if s.Res.InvoiceAmountAdjusted != nil {
		s.D.Set("invoice_amount_adjusted", *s.Res.InvoiceAmountAdjusted)
	}

	if s.Res.InvoiceAmountApplied != nil {
		s.D.Set("invoice_amount_applied", *s.Res.InvoiceAmountApplied)
	}

	if s.Res.InvoiceAmountCredited != nil {
		s.D.Set("invoice_amount_credited", *s.Res.InvoiceAmountCredited)
	}

	if s.Res.InvoiceAmountDue != nil {
		s.D.Set("invoice_amount_due", *s.Res.InvoiceAmountDue)
	}

	if s.Res.InvoiceId != nil {
		s.D.Set("invoice_id", *s.Res.InvoiceId)
	}

	if s.Res.InvoiceNumber != nil {
		s.D.Set("invoice_number", *s.Res.InvoiceNumber)
	}

	if s.Res.InvoicePoNumber != nil {
		s.D.Set("invoice_po_number", *s.Res.InvoicePoNumber)
	}

	if s.Res.InvoiceRefNumber != nil {
		s.D.Set("invoice_ref_number", *s.Res.InvoiceRefNumber)
	}

	s.D.Set("invoice_status", s.Res.InvoiceStatus)

	s.D.Set("invoice_type", s.Res.InvoiceType)

	if s.Res.IsCreditCardPayable != nil {
		s.D.Set("is_credit_card_payable", *s.Res.IsCreditCardPayable)
	}

	if s.Res.IsDisplayDownloadPdf != nil {
		s.D.Set("is_display_download_pdf", *s.Res.IsDisplayDownloadPdf)
	}

	if s.Res.IsPayable != nil {
		s.D.Set("is_payable", *s.Res.IsPayable)
	}

	if s.Res.IsPdfEmailAvailable != nil {
		s.D.Set("is_pdf_email_available", *s.Res.IsPdfEmailAvailable)
	}

	if s.Res.LastPaymentDetail != nil {
		lastPaymentDetailArray := []interface{}{}
		if lastPaymentDetailMap := PaymentDetailToMap(&s.Res.LastPaymentDetail); lastPaymentDetailMap != nil {
			lastPaymentDetailArray = append(lastPaymentDetailArray, lastPaymentDetailMap)
		}
		s.D.Set("last_payment_detail", lastPaymentDetailArray)
	} else {
		s.D.Set("last_payment_detail", nil)
	}

	if s.Res.PaymentTerms != nil {
		s.D.Set("payment_terms", *s.Res.PaymentTerms)
	}

	if s.Res.PreferredEmail != nil {
		s.D.Set("preferred_email", *s.Res.PreferredEmail)
	}

	s.D.Set("subscription_ids", s.Res.SubscriptionIds)

	if s.Res.Tax != nil {
		s.D.Set("tax", *s.Res.Tax)
	}

	if s.Res.TimeInvoice != nil {
		s.D.Set("time_invoice", s.Res.TimeInvoice.String())
	}

	if s.Res.TimeInvoiceDue != nil {
		s.D.Set("time_invoice_due", s.Res.TimeInvoiceDue.String())
	}

	return nil
}

func BillToAddressToMap(obj *oci_osp_gateway.BillToAddress) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AddressLine1 != nil {
		result["address_line1"] = string(*obj.AddressLine1)
	}

	if obj.AddressLine2 != nil {
		result["address_line2"] = string(*obj.AddressLine2)
	}

	if obj.AddressLine3 != nil {
		result["address_line3"] = string(*obj.AddressLine3)
	}

	if obj.AddressLine4 != nil {
		result["address_line4"] = string(*obj.AddressLine4)
	}

	if obj.City != nil {
		result["city"] = string(*obj.City)
	}

	if obj.CompanyName != nil {
		result["company_name"] = string(*obj.CompanyName)
	}

	if obj.ContactName != nil {
		result["contact_name"] = string(*obj.ContactName)
	}

	if obj.Country != nil {
		result["country"] = []interface{}{CountryToMap(obj.Country)}
	}

	if obj.County != nil {
		result["county"] = string(*obj.County)
	}

	if obj.PostalCode != nil {
		result["postal_code"] = string(*obj.PostalCode)
	}

	if obj.Province != nil {
		result["province"] = string(*obj.Province)
	}

	if obj.State != nil {
		result["state"] = string(*obj.State)
	}

	if obj.StreetName != nil {
		result["street_name"] = string(*obj.StreetName)
	}

	if obj.StreetNumber != nil {
		result["street_number"] = string(*obj.StreetNumber)
	}

	return result
}

func CountryToMap(obj *oci_osp_gateway.Country) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ascii3CountryCode != nil {
		result["ascii3country_code"] = string(*obj.Ascii3CountryCode)
	}

	if obj.CountryCode != nil {
		result["country_code"] = string(*obj.CountryCode)
	}

	if obj.CountryId != nil {
		result["country_id"] = float32(*obj.CountryId)
	}

	if obj.CountryName != nil {
		result["country_name"] = string(*obj.CountryName)
	}

	if obj.LanguageId != nil {
		result["language_id"] = float32(*obj.LanguageId)
	}

	return result
}

func PaymentDetailToMap(obj *oci_osp_gateway.PaymentDetail) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_osp_gateway.CreditCardPaymentDetail:
		result["payment_method"] = "CREDIT_CARD"

		result["credit_card_type"] = string(v.CreditCardType)

		if v.LastDigits != nil {
			result["last_digits"] = string(*v.LastDigits)
		}

		if v.NameOnCard != nil {
			result["name_on_card"] = string(*v.NameOnCard)
		}

		if v.TimeExpiration != nil {
			result["time_expiration"] = v.TimeExpiration.Format(time.RFC3339Nano)
		}

		if v.AmountPaid != nil {
			result["amount_paid"] = float32(*v.AmountPaid)
		}

		if v.PaidBy != nil {
			result["paid_by"] = string(*v.PaidBy)
		}

		if v.TimePaidOn != nil {
			result["time_paid_on"] = v.TimePaidOn.String()
		}
	case oci_osp_gateway.EcheckPaymentDetail:
		result["payment_method"] = "ECHECK"

		if v.AccountNumber != nil {
			result["account_number"] = string(*v.AccountNumber)
		}

		result["card_type"] = string(v.CardType)

		if v.NameOnCard != nil {
			result["name_on_card"] = string(*v.NameOnCard)
		}

		if v.RoutingNumber != nil {
			result["routing_number"] = string(*v.RoutingNumber)
		}

		if v.AmountPaid != nil {
			result["amount_paid"] = float32(*v.AmountPaid)
		}

		if v.PaidBy != nil {
			result["paid_by"] = string(*v.PaidBy)
		}

		if v.TimePaidOn != nil {
			result["time_paid_on"] = v.TimePaidOn.String()
		}
	case oci_osp_gateway.OtherPaymentDetail:
		result["payment_method"] = "OTHER"

		result["credit_card_type"] = string(v.CreditCardType)

		if v.EcheckRouting != nil {
			result["echeck_routing"] = string(*v.EcheckRouting)
		}

		if v.LastDigits != nil {
			result["last_digits"] = string(*v.LastDigits)
		}

		if v.NameOnCard != nil {
			result["name_on_card"] = string(*v.NameOnCard)
		}

		if v.TimeExpiration != nil {
			result["time_expiration"] = v.TimeExpiration.Format(time.RFC3339Nano)
		}

		if v.AmountPaid != nil {
			result["amount_paid"] = float32(*v.AmountPaid)
		}

		if v.PaidBy != nil {
			result["paid_by"] = string(*v.PaidBy)
		}

		if v.TimePaidOn != nil {
			result["time_paid_on"] = v.TimePaidOn.String()
		}
	case oci_osp_gateway.PaypalPaymentDetail:
		result["payment_method"] = "PAYPAL"

		if v.PaypalId != nil {
			result["paypal_id"] = string(*v.PaypalId)
		}

		if v.PaypalReference != nil {
			result["paypal_reference"] = string(*v.PaypalReference)
		}

		if v.AmountPaid != nil {
			result["amount_paid"] = float32(*v.AmountPaid)
		}

		if v.PaidBy != nil {
			result["paid_by"] = string(*v.PaidBy)
		}

		if v.TimePaidOn != nil {
			result["time_paid_on"] = v.TimePaidOn.String()
		}
	default:
		log.Printf("[WARN] Received 'payment_method' of unknown type %v", *obj)
		return nil
	}

	return result
}
