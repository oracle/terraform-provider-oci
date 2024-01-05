// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package onesubscription

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_onesubscription "github.com/oracle/oci-go-sdk/v65/onesubscription"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OnesubscriptionInvoicesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOnesubscriptionInvoices,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"ar_customer_transaction_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"fields": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"time_from": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"invoices": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"ar_invoices": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"bill_to_address": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"bill_site_use_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_bill_to": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_ship_to": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"location": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"address1": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"address2": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"city": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"country": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"postal_code": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"region": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"tca_location_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"phone": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"service2site_use_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"tca_cust_acct_site_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"tca_party_site_number": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"bill_to_contact": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"email": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"first_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"last_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"tca_contact_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"tca_cust_accnt_site_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"tca_party_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"user_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"bill_to_customer": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"customer_chain_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_chain_customer": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_public_sector": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name_phonetic": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"tca_customer_account_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"tca_customer_account_number": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"tca_party_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"tca_party_number": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"created_by": {
							Type:     schema.TypeString,
							Computed: true,
						},
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
						"invoice_lines": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"ar_invoice_number": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"data_center": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
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
												"billing_category": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"part_number": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"product_category": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"ucm_rate_card_part_type": {
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
						"organization": {
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
									"number": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},
						"payment_method": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"payment_term": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"created_by": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_active": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_updated": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"updated_by": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"receipt_method": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"spm_invoice_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subscription_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_invoice_date": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"updated_by": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readOnesubscriptionInvoices(d *schema.ResourceData, m interface{}) error {
	sync := &OnesubscriptionInvoicesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).InvoiceSummaryRegionalClient()

	return tfresource.ReadResource(sync)
}

type OnesubscriptionInvoicesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_onesubscription.InvoiceSummaryClient
	Res    *oci_onesubscription.ListInvoicesResponse
}

func (s *OnesubscriptionInvoicesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OnesubscriptionInvoicesDataSourceCrud) Get() error {
	request := oci_onesubscription.ListInvoicesRequest{}

	if arCustomerTransactionId, ok := s.D.GetOkExists("ar_customer_transaction_id"); ok {
		tmp := arCustomerTransactionId.(string)
		request.ArCustomerTransactionId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if fields, ok := s.D.GetOkExists("fields"); ok {
		interfaces := fields.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("fields") {
			request.Fields = tmp
		}
	}

	if timeFrom, ok := s.D.GetOkExists("time_from"); ok {
		tmp, err := time.Parse(time.RFC3339, timeFrom.(string))
		if err != nil {
			return err
		}
		request.TimeFrom = &oci_common.SDKTime{Time: tmp}
	}

	if timeTo, ok := s.D.GetOkExists("time_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeTo.(string))
		if err != nil {
			return err
		}
		request.TimeTo = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "onesubscription")

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

func (s *OnesubscriptionInvoicesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OnesubscriptionInvoicesDataSource-", OnesubscriptionInvoicesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		invoice := map[string]interface{}{}

		if r.ArInvoices != nil {
			invoice["ar_invoices"] = *r.ArInvoices
		}

		if r.BillToAddress != nil {
			invoice["bill_to_address"] = []interface{}{InvoicingAddressToMap(r.BillToAddress)}
		} else {
			invoice["bill_to_address"] = nil
		}

		if r.BillToContact != nil {
			invoice["bill_to_contact"] = []interface{}{InvoicingUserToMap(r.BillToContact)}
		} else {
			invoice["bill_to_contact"] = nil
		}

		if r.BillToCustomer != nil {
			invoice["bill_to_customer"] = []interface{}{InvoicingBusinessPartnerToMap(r.BillToCustomer)}
		} else {
			invoice["bill_to_customer"] = nil
		}

		if r.CreatedBy != nil {
			invoice["created_by"] = *r.CreatedBy
		}

		if r.Currency != nil {
			invoice["currency"] = []interface{}{InvoicingCurrencyToMap(r.Currency)}
		} else {
			invoice["currency"] = nil
		}

		invoiceLines := []interface{}{}
		for _, item := range r.InvoiceLines {
			invoiceLines = append(invoiceLines, InvoiceLineSummaryToMap(item))
		}
		invoice["invoice_lines"] = invoiceLines

		if r.Organization != nil {
			invoice["organization"] = []interface{}{InvoicingOrganizationToMap(r.Organization)}
		} else {
			invoice["organization"] = nil
		}

		if r.PaymentMethod != nil {
			invoice["payment_method"] = *r.PaymentMethod
		}

		if r.PaymentTerm != nil {
			invoice["payment_term"] = []interface{}{InvoicingPaymentTermToMap(r.PaymentTerm)}
		} else {
			invoice["payment_term"] = nil
		}

		if r.ReceiptMethod != nil {
			invoice["receipt_method"] = *r.ReceiptMethod
		}

		if r.SpmInvoiceNumber != nil {
			invoice["spm_invoice_number"] = *r.SpmInvoiceNumber
		}

		if r.Status != nil {
			invoice["status"] = *r.Status
		}

		if r.SubscriptionNumber != nil {
			invoice["subscription_number"] = *r.SubscriptionNumber
		}

		if r.TimeCreated != nil {
			invoice["time_created"] = r.TimeCreated.String()
		}

		if r.TimeInvoiceDate != nil {
			invoice["time_invoice_date"] = r.TimeInvoiceDate.String()
		}

		if r.TimeUpdated != nil {
			invoice["time_updated"] = r.TimeUpdated.String()
		}

		if r.Type != nil {
			invoice["type"] = *r.Type
		}

		if r.UpdatedBy != nil {
			invoice["updated_by"] = *r.UpdatedBy
		}

		resources = append(resources, invoice)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, OnesubscriptionInvoicesDataSource().Schema["invoices"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("invoices", resources); err != nil {
		return err
	}

	return nil
}

func InvoiceLineSummaryToMap(obj oci_onesubscription.InvoiceLineSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ArInvoiceNumber != nil {
		result["ar_invoice_number"] = string(*obj.ArInvoiceNumber)
	}

	if obj.DataCenter != nil {
		result["data_center"] = string(*obj.DataCenter)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Product != nil {
		result["product"] = []interface{}{InvoicingProductToMap(obj.Product)}
	}

	if obj.TimeEnd != nil {
		result["time_end"] = obj.TimeEnd.String()
	}

	if obj.TimeStart != nil {
		result["time_start"] = obj.TimeStart.String()
	}

	return result
}

func InvoicingAddressToMap(obj *oci_onesubscription.InvoicingAddress) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BillSiteUseId != nil {
		result["bill_site_use_id"] = strconv.FormatInt(*obj.BillSiteUseId, 10)
	}

	if obj.IsBillTo != nil {
		result["is_bill_to"] = bool(*obj.IsBillTo)
	}

	if obj.IsShipTo != nil {
		result["is_ship_to"] = bool(*obj.IsShipTo)
	}

	if obj.Location != nil {
		result["location"] = []interface{}{InvoicingLocationToMap(obj.Location)}
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Phone != nil {
		result["phone"] = string(*obj.Phone)
	}

	if obj.Service2SiteUseId != nil {
		result["service2site_use_id"] = strconv.FormatInt(*obj.Service2SiteUseId, 10)
	}

	if obj.TcaCustAcctSiteId != nil {
		result["tca_cust_acct_site_id"] = strconv.FormatInt(*obj.TcaCustAcctSiteId, 10)
	}

	if obj.TcaPartySiteNumber != nil {
		result["tca_party_site_number"] = string(*obj.TcaPartySiteNumber)
	}

	return result
}

func InvoicingBusinessPartnerToMap(obj *oci_onesubscription.InvoicingBusinessPartner) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CustomerChainType != nil {
		result["customer_chain_type"] = string(*obj.CustomerChainType)
	}

	if obj.IsChainCustomer != nil {
		result["is_chain_customer"] = bool(*obj.IsChainCustomer)
	}

	if obj.IsPublicSector != nil {
		result["is_public_sector"] = bool(*obj.IsPublicSector)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.NamePhonetic != nil {
		result["name_phonetic"] = string(*obj.NamePhonetic)
	}

	if obj.TcaCustomerAccountId != nil {
		result["tca_customer_account_id"] = strconv.FormatInt(*obj.TcaCustomerAccountId, 10)
	}

	if obj.TcaCustomerAccountNumber != nil {
		result["tca_customer_account_number"] = string(*obj.TcaCustomerAccountNumber)
	}

	if obj.TcaPartyId != nil {
		result["tca_party_id"] = strconv.FormatInt(*obj.TcaPartyId, 10)
	}

	if obj.TcaPartyNumber != nil {
		result["tca_party_number"] = string(*obj.TcaPartyNumber)
	}

	return result
}

func InvoicingCurrencyToMap(obj *oci_onesubscription.InvoicingCurrency) map[string]interface{} {
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

func InvoicingLocationToMap(obj *oci_onesubscription.InvoicingLocation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Address1 != nil {
		result["address1"] = string(*obj.Address1)
	}

	if obj.Address2 != nil {
		result["address2"] = string(*obj.Address2)
	}

	if obj.City != nil {
		result["city"] = string(*obj.City)
	}

	if obj.Country != nil {
		result["country"] = string(*obj.Country)
	}

	if obj.PostalCode != nil {
		result["postal_code"] = string(*obj.PostalCode)
	}

	if obj.Region != nil {
		result["region"] = string(*obj.Region)
	}

	if obj.TcaLocationId != nil {
		result["tca_location_id"] = strconv.FormatInt(*obj.TcaLocationId, 10)
	}

	return result
}

func InvoicingOrganizationToMap(obj *oci_onesubscription.InvoicingOrganization) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Number != nil {
		result["number"] = float64(*obj.Number)
	}

	return result
}

func InvoicingPaymentTermToMap(obj *oci_onesubscription.InvoicingPaymentTerm) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CreatedBy != nil {
		result["created_by"] = string(*obj.CreatedBy)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.IsActive != nil {
		result["is_active"] = bool(*obj.IsActive)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.UpdatedBy != nil {
		result["updated_by"] = string(*obj.UpdatedBy)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func InvoicingProductToMap(obj *oci_onesubscription.InvoicingProduct) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BillingCategory != nil {
		result["billing_category"] = string(*obj.BillingCategory)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.PartNumber != nil {
		result["part_number"] = string(*obj.PartNumber)
	}

	if obj.ProductCategory != nil {
		result["product_category"] = string(*obj.ProductCategory)
	}

	if obj.UcmRateCardPartType != nil {
		result["ucm_rate_card_part_type"] = string(*obj.UcmRateCardPartType)
	}

	if obj.UnitOfMeasure != nil {
		result["unit_of_measure"] = string(*obj.UnitOfMeasure)
	}

	return result
}

func InvoicingUserToMap(obj *oci_onesubscription.InvoicingUser) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Email != nil {
		result["email"] = string(*obj.Email)
	}

	if obj.FirstName != nil {
		result["first_name"] = string(*obj.FirstName)
	}

	if obj.LastName != nil {
		result["last_name"] = string(*obj.LastName)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.TcaContactId != nil {
		result["tca_contact_id"] = strconv.FormatInt(*obj.TcaContactId, 10)
	}

	if obj.TcaCustAccntSiteId != nil {
		result["tca_cust_accnt_site_id"] = strconv.FormatInt(*obj.TcaCustAccntSiteId, 10)
	}

	if obj.TcaPartyId != nil {
		result["tca_party_id"] = strconv.FormatInt(*obj.TcaPartyId, 10)
	}

	if obj.UserName != nil {
		result["user_name"] = string(*obj.UserName)
	}

	return result
}
