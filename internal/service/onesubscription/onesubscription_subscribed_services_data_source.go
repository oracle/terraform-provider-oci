// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package onesubscription

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_onesubscription "github.com/oracle/oci-go-sdk/v65/onesubscription"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OnesubscriptionSubscribedServicesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOnesubscriptionSubscribedServices,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"order_line_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subscribed_services": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"admin_email": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"agreement_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"agreement_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"agreement_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"available_amount": {
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
									"username": {
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
									"tca_cust_account_number": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"tca_customer_account_id": {
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
						"billing_frequency": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"booking_opty_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"buyer_email": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"commitment_schedule_id": {
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
						"created_by": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"credit_percentage": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"csi": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"customer_transaction_reference": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"data_center": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"data_center_region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"eligible_to_renew": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"end_user_address": {
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
						"end_user_contact": {
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
									"username": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"end_user_customer": {
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
									"tca_cust_account_number": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"tca_customer_account_id": {
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
						"fulfillment_set": {
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
						"is_allowance": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_cap_to_price_list": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_credit_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_having_usage": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_intent_to_pay": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_payg": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_single_rate_card": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_variable_commitment": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"line_net_amount": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"major_set": {
							Type:     schema.TypeString,
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
						"order_header_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"order_line_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"order_line_number": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"order_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"order_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"original_promo_amount": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"overage_bill_to": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"overage_discount_percentage": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"overage_policy": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"partner_credit_amount": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"partner_transaction_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"payg_policy": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"payment_method": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"payment_number": {
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
						"price_period": {
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
						"program_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"promo_order_line_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"promo_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"promotion_pricing_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"provisioning_source": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"quantity": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"rate_card_discount_percentage": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"rate_cards": {
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
									"discretionary_discount_percentage": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_tier": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"net_unit_price": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"overage_price": {
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
									"rate_card_tiers": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"net_unit_price": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"overage_price": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"up_to_quantity": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"subscribed_service_id": {
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
						"ratecard_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"renewal_opty_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"renewal_opty_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"renewal_opty_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"renewed_subscribed_service_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"reseller_address": {
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
						"reseller_contact": {
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
									"username": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"reseller_customer": {
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
									"tca_cust_account_number": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"tca_customer_account_id": {
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
						"revenue_line_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"revenue_line_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"revised_arr_in_lc": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"revised_arr_in_sc": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sales_account_party_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sales_channel": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"serial_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"service_to_address": {
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
						"service_to_contact": {
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
									"username": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"service_to_customer": {
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
									"tca_cust_account_number": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"tca_customer_account_id": {
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
						"sold_to_contact": {
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
									"username": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"sold_to_customer": {
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
									"tca_cust_account_number": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"tca_customer_account_id": {
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
						"start_date_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subscription_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subscription_source": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"system_arr_in_lc": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"system_arr_in_sc": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"system_atr_arr_in_lc": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"system_atr_arr_in_sc": {
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
						"time_agreement_end": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_customer_config": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_end": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_majorset_end": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_majorset_start": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_payment_expiry": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_provisioned": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_service_configuration_email_sent": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_start": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_welcome_email_sent": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"total_value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"transaction_extension_id": {
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

func readOnesubscriptionSubscribedServices(d *schema.ResourceData, m interface{}) error {
	sync := &OnesubscriptionSubscribedServicesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SubscribedServiceRegionalClient()

	return tfresource.ReadResource(sync)
}

type OnesubscriptionSubscribedServicesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_onesubscription.SubscribedServiceClient
	Res    *oci_onesubscription.ListSubscribedServicesResponse
}

func (s *OnesubscriptionSubscribedServicesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OnesubscriptionSubscribedServicesDataSourceCrud) Get() error {
	request := oci_onesubscription.ListSubscribedServicesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if orderLineId, ok := s.D.GetOkExists("order_line_id"); ok {
		tmp := orderLineId.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert orderLineId string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.OrderLineId = &tmpInt64
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		tmp := status.(string)
		request.Status = &tmp
	}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "onesubscription")

	response, err := s.Client.ListSubscribedServices(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSubscribedServices(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OnesubscriptionSubscribedServicesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OnesubscriptionSubscribedServicesDataSource-", OnesubscriptionSubscribedServicesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		subscribedService := map[string]interface{}{
			"subscription_id": *r.SubscriptionId,
		}

		if r.AdminEmail != nil {
			subscribedService["admin_email"] = *r.AdminEmail
		}

		if r.AgreementId != nil {
			subscribedService["agreement_id"] = strconv.FormatInt(*r.AgreementId, 10)
		}

		if r.AgreementName != nil {
			subscribedService["agreement_name"] = *r.AgreementName
		}

		if r.AgreementType != nil {
			subscribedService["agreement_type"] = *r.AgreementType
		}

		if r.AvailableAmount != nil {
			subscribedService["available_amount"] = *r.AvailableAmount
		}

		if r.BillToAddress != nil {
			subscribedService["bill_to_address"] = []interface{}{SubscribedServiceAddressToMap(r.BillToAddress)}
		} else {
			subscribedService["bill_to_address"] = nil
		}

		if r.BillToContact != nil {
			subscribedService["bill_to_contact"] = []interface{}{SubscribedServiceUserToMap(r.BillToContact)}
		} else {
			subscribedService["bill_to_contact"] = nil
		}

		if r.BillToCustomer != nil {
			subscribedService["bill_to_customer"] = []interface{}{SubscribedServiceBusinessPartnerToMap(r.BillToCustomer)}
		} else {
			subscribedService["bill_to_customer"] = nil
		}

		if r.BillingFrequency != nil {
			subscribedService["billing_frequency"] = *r.BillingFrequency
		}

		if r.BookingOptyNumber != nil {
			subscribedService["booking_opty_number"] = *r.BookingOptyNumber
		}

		if r.BuyerEmail != nil {
			subscribedService["buyer_email"] = *r.BuyerEmail
		}

		if r.CommitmentScheduleId != nil {
			subscribedService["commitment_schedule_id"] = *r.CommitmentScheduleId
		}

		if r.CreatedBy != nil {
			subscribedService["created_by"] = *r.CreatedBy
		}

		if r.CreditPercentage != nil {
			subscribedService["credit_percentage"] = *r.CreditPercentage
		}

		if r.Csi != nil {
			subscribedService["csi"] = strconv.FormatInt(*r.Csi, 10)
		}

		if r.CustomerTransactionReference != nil {
			subscribedService["customer_transaction_reference"] = *r.CustomerTransactionReference
		}

		if r.DataCenter != nil {
			subscribedService["data_center"] = *r.DataCenter
		}

		if r.DataCenterRegion != nil {
			subscribedService["data_center_region"] = *r.DataCenterRegion
		}

		if r.EligibleToRenew != nil {
			subscribedService["eligible_to_renew"] = *r.EligibleToRenew
		}

		if r.EndUserAddress != nil {
			subscribedService["end_user_address"] = []interface{}{SubscribedServiceAddressToMap(r.EndUserAddress)}
		} else {
			subscribedService["end_user_address"] = nil
		}

		if r.EndUserContact != nil {
			subscribedService["end_user_contact"] = []interface{}{SubscribedServiceUserToMap(r.EndUserContact)}
		} else {
			subscribedService["end_user_contact"] = nil
		}

		if r.EndUserCustomer != nil {
			subscribedService["end_user_customer"] = []interface{}{SubscribedServiceBusinessPartnerToMap(r.EndUserCustomer)}
		} else {
			subscribedService["end_user_customer"] = nil
		}

		if r.FulfillmentSet != nil {
			subscribedService["fulfillment_set"] = *r.FulfillmentSet
		}

		if r.FundedAllocationValue != nil {
			subscribedService["funded_allocation_value"] = *r.FundedAllocationValue
		}

		if r.Id != nil {
			subscribedService["id"] = *r.Id
		}

		if r.IsAllowance != nil {
			subscribedService["is_allowance"] = *r.IsAllowance
		}

		if r.IsCapToPriceList != nil {
			subscribedService["is_cap_to_price_list"] = *r.IsCapToPriceList
		}

		if r.IsCreditEnabled != nil {
			subscribedService["is_credit_enabled"] = *r.IsCreditEnabled
		}

		if r.IsHavingUsage != nil {
			subscribedService["is_having_usage"] = *r.IsHavingUsage
		}

		if r.IsIntentToPay != nil {
			subscribedService["is_intent_to_pay"] = *r.IsIntentToPay
		}

		if r.IsPayg != nil {
			subscribedService["is_payg"] = *r.IsPayg
		}

		if r.IsSingleRateCard != nil {
			subscribedService["is_single_rate_card"] = *r.IsSingleRateCard
		}

		if r.IsVariableCommitment != nil {
			subscribedService["is_variable_commitment"] = *r.IsVariableCommitment
		}

		if r.LineNetAmount != nil {
			subscribedService["line_net_amount"] = *r.LineNetAmount
		}

		if r.MajorSet != nil {
			subscribedService["major_set"] = strconv.FormatInt(*r.MajorSet, 10)
		}

		if r.NetUnitPrice != nil {
			subscribedService["net_unit_price"] = *r.NetUnitPrice
		}

		if r.OperationType != nil {
			subscribedService["operation_type"] = *r.OperationType
		}

		if r.OrderHeaderId != nil {
			subscribedService["order_header_id"] = strconv.FormatInt(*r.OrderHeaderId, 10)
		}

		if r.OrderLineId != nil {
			subscribedService["order_line_id"] = strconv.FormatInt(*r.OrderLineId, 10)
		}

		if r.OrderLineNumber != nil {
			subscribedService["order_line_number"] = *r.OrderLineNumber
		}

		if r.OrderNumber != nil {
			subscribedService["order_number"] = strconv.FormatInt(*r.OrderNumber, 10)
		}

		if r.OrderType != nil {
			subscribedService["order_type"] = *r.OrderType
		}

		if r.OriginalPromoAmount != nil {
			subscribedService["original_promo_amount"] = *r.OriginalPromoAmount
		}

		if r.OverageBillTo != nil {
			subscribedService["overage_bill_to"] = *r.OverageBillTo
		}

		if r.OverageDiscountPercentage != nil {
			subscribedService["overage_discount_percentage"] = *r.OverageDiscountPercentage
		}

		if r.OveragePolicy != nil {
			subscribedService["overage_policy"] = *r.OveragePolicy
		}

		if r.PartnerCreditAmount != nil {
			subscribedService["partner_credit_amount"] = *r.PartnerCreditAmount
		}

		if r.PartnerTransactionType != nil {
			subscribedService["partner_transaction_type"] = *r.PartnerTransactionType
		}

		if r.PaygPolicy != nil {
			subscribedService["payg_policy"] = *r.PaygPolicy
		}

		if r.PaymentMethod != nil {
			subscribedService["payment_method"] = *r.PaymentMethod
		}

		if r.PaymentNumber != nil {
			subscribedService["payment_number"] = *r.PaymentNumber
		}

		if r.PaymentTerm != nil {
			subscribedService["payment_term"] = []interface{}{SubscribedServicePaymentTermToMap(r.PaymentTerm)}
		} else {
			subscribedService["payment_term"] = nil
		}

		if r.PricePeriod != nil {
			subscribedService["price_period"] = *r.PricePeriod
		}

		if r.PricingModel != nil {
			subscribedService["pricing_model"] = *r.PricingModel
		}

		if r.Product != nil {
			subscribedService["product"] = []interface{}{RateCardProductToMap(r.Product)}
		} else {
			subscribedService["product"] = nil
		}

		if r.ProgramType != nil {
			subscribedService["program_type"] = *r.ProgramType
		}

		if r.PromoOrderLineId != nil {
			subscribedService["promo_order_line_id"] = strconv.FormatInt(*r.PromoOrderLineId, 10)
		}

		if r.PromoType != nil {
			subscribedService["promo_type"] = *r.PromoType
		}

		if r.PromotionPricingType != nil {
			subscribedService["promotion_pricing_type"] = *r.PromotionPricingType
		}

		if r.ProvisioningSource != nil {
			subscribedService["provisioning_source"] = *r.ProvisioningSource
		}

		if r.Quantity != nil {
			subscribedService["quantity"] = *r.Quantity
		}

		if r.RateCardDiscountPercentage != nil {
			subscribedService["rate_card_discount_percentage"] = *r.RateCardDiscountPercentage
		}

		if r.RatecardType != nil {
			subscribedService["ratecard_type"] = *r.RatecardType
		}

		if r.RenewalOptyId != nil {
			subscribedService["renewal_opty_id"] = strconv.FormatInt(*r.RenewalOptyId, 10)
		}

		if r.RenewalOptyNumber != nil {
			subscribedService["renewal_opty_number"] = *r.RenewalOptyNumber
		}

		if r.RenewalOptyType != nil {
			subscribedService["renewal_opty_type"] = *r.RenewalOptyType
		}

		if r.RenewedSubscribedServiceId != nil {
			subscribedService["renewed_subscribed_service_id"] = *r.RenewedSubscribedServiceId
		}

		if r.ResellerAddress != nil {
			subscribedService["reseller_address"] = []interface{}{SubscribedServiceAddressToMap(r.ResellerAddress)}
		} else {
			subscribedService["reseller_address"] = nil
		}

		if r.ResellerContact != nil {
			subscribedService["reseller_contact"] = []interface{}{SubscribedServiceUserToMap(r.ResellerContact)}
		} else {
			subscribedService["reseller_contact"] = nil
		}

		if r.ResellerCustomer != nil {
			subscribedService["reseller_customer"] = []interface{}{SubscribedServiceBusinessPartnerToMap(r.ResellerCustomer)}
		} else {
			subscribedService["reseller_customer"] = nil
		}

		if r.RevenueLineId != nil {
			subscribedService["revenue_line_id"] = strconv.FormatInt(*r.RevenueLineId, 10)
		}

		if r.RevenueLineNumber != nil {
			subscribedService["revenue_line_number"] = *r.RevenueLineNumber
		}

		if r.RevisedArrInLc != nil {
			subscribedService["revised_arr_in_lc"] = *r.RevisedArrInLc
		}

		if r.RevisedArrInSc != nil {
			subscribedService["revised_arr_in_sc"] = *r.RevisedArrInSc
		}

		if r.SalesAccountPartyId != nil {
			subscribedService["sales_account_party_id"] = strconv.FormatInt(*r.SalesAccountPartyId, 10)
		}

		if r.SalesChannel != nil {
			subscribedService["sales_channel"] = *r.SalesChannel
		}

		if r.SerialNumber != nil {
			subscribedService["serial_number"] = *r.SerialNumber
		}

		if r.ServiceToAddress != nil {
			subscribedService["service_to_address"] = []interface{}{SubscribedServiceAddressToMap(r.ServiceToAddress)}
		} else {
			subscribedService["service_to_address"] = nil
		}

		if r.ServiceToContact != nil {
			subscribedService["service_to_contact"] = []interface{}{SubscribedServiceUserToMap(r.ServiceToContact)}
		} else {
			subscribedService["service_to_contact"] = nil
		}

		if r.ServiceToCustomer != nil {
			subscribedService["service_to_customer"] = []interface{}{SubscribedServiceBusinessPartnerToMap(r.ServiceToCustomer)}
		} else {
			subscribedService["service_to_customer"] = nil
		}

		if r.SoldToContact != nil {
			subscribedService["sold_to_contact"] = []interface{}{SubscribedServiceUserToMap(r.SoldToContact)}
		} else {
			subscribedService["sold_to_contact"] = nil
		}

		if r.SoldToCustomer != nil {
			subscribedService["sold_to_customer"] = []interface{}{SubscribedServiceBusinessPartnerToMap(r.SoldToCustomer)}
		} else {
			subscribedService["sold_to_customer"] = nil
		}

		if r.StartDateType != nil {
			subscribedService["start_date_type"] = *r.StartDateType
		}

		if r.Status != nil {
			subscribedService["status"] = *r.Status
		}

		if r.SubscriptionSource != nil {
			subscribedService["subscription_source"] = *r.SubscriptionSource
		}

		if r.SystemArrInLc != nil {
			subscribedService["system_arr_in_lc"] = *r.SystemArrInLc
		}

		if r.SystemArrInSc != nil {
			subscribedService["system_arr_in_sc"] = *r.SystemArrInSc
		}

		if r.SystemAtrArrInLc != nil {
			subscribedService["system_atr_arr_in_lc"] = *r.SystemAtrArrInLc
		}

		if r.SystemAtrArrInSc != nil {
			subscribedService["system_atr_arr_in_sc"] = *r.SystemAtrArrInSc
		}

		if r.TermValue != nil {
			subscribedService["term_value"] = strconv.FormatInt(*r.TermValue, 10)
		}

		if r.TermValueUom != nil {
			subscribedService["term_value_uom"] = *r.TermValueUom
		}

		if r.TimeAgreementEnd != nil {
			subscribedService["time_agreement_end"] = r.TimeAgreementEnd.String()
		}

		if r.TimeCreated != nil {
			subscribedService["time_created"] = r.TimeCreated.String()
		}

		if r.TimeCustomerConfig != nil {
			subscribedService["time_customer_config"] = r.TimeCustomerConfig.String()
		}

		if r.TimeEnd != nil {
			subscribedService["time_end"] = r.TimeEnd.String()
		}

		if r.TimeMajorsetEnd != nil {
			subscribedService["time_majorset_end"] = r.TimeMajorsetEnd.String()
		}

		if r.TimeMajorsetStart != nil {
			subscribedService["time_majorset_start"] = r.TimeMajorsetStart.String()
		}

		if r.TimePaymentExpiry != nil {
			subscribedService["time_payment_expiry"] = r.TimePaymentExpiry.String()
		}

		if r.TimeProvisioned != nil {
			subscribedService["time_provisioned"] = r.TimeProvisioned.String()
		}

		if r.TimeServiceConfigurationEmailSent != nil {
			subscribedService["time_service_configuration_email_sent"] = r.TimeServiceConfigurationEmailSent.String()
		}

		if r.TimeStart != nil {
			subscribedService["time_start"] = r.TimeStart.String()
		}

		if r.TimeUpdated != nil {
			subscribedService["time_updated"] = r.TimeUpdated.String()
		}

		if r.TimeWelcomeEmailSent != nil {
			subscribedService["time_welcome_email_sent"] = r.TimeWelcomeEmailSent.String()
		}

		if r.TotalValue != nil {
			subscribedService["total_value"] = *r.TotalValue
		}

		if r.TransactionExtensionId != nil {
			subscribedService["transaction_extension_id"] = strconv.FormatInt(*r.TransactionExtensionId, 10)
		}

		if r.Type != nil {
			subscribedService["type"] = *r.Type
		}

		if r.UpdatedBy != nil {
			subscribedService["updated_by"] = *r.UpdatedBy
		}

		if r.UsedAmount != nil {
			subscribedService["used_amount"] = *r.UsedAmount
		}

		resources = append(resources, subscribedService)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, OnesubscriptionSubscribedServicesDataSource().Schema["subscribed_services"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("subscribed_services", resources); err != nil {
		return err
	}

	return nil
}

func SubscribedServiceBusinessPartnerToMap(obj *oci_onesubscription.SubscribedServiceBusinessPartner) map[string]interface{} {
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

	if obj.TcaCustAccountNumber != nil {
		result["tca_cust_account_number"] = string(*obj.TcaCustAccountNumber)
	}

	if obj.TcaCustomerAccountId != nil {
		result["tca_customer_account_id"] = strconv.FormatInt(*obj.TcaCustomerAccountId, 10)
	}

	if obj.TcaPartyId != nil {
		result["tca_party_id"] = strconv.FormatInt(*obj.TcaPartyId, 10)
	}

	if obj.TcaPartyNumber != nil {
		result["tca_party_number"] = string(*obj.TcaPartyNumber)
	}

	return result
}

func SubscribedServiceLocationToMap(obj *oci_onesubscription.SubscribedServiceLocation) map[string]interface{} {
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

func SubscribedServicePaymentTermToMap(obj *oci_onesubscription.SubscribedServicePaymentTerm) map[string]interface{} {
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

func SubscribedServiceUserToMap(obj *oci_onesubscription.SubscribedServiceUser) map[string]interface{} {
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

	if obj.Username != nil {
		result["username"] = string(*obj.Username)
	}

	return result
}

func RateCardProductToMap(obj *oci_onesubscription.RateCardProduct) map[string]interface{} {
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

func RateCardSummaryToMap(obj oci_onesubscription.RateCardSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Currency != nil {
		result["currency"] = []interface{}{SubscriptionCurrencyToMap(obj.Currency)}
	}

	if obj.DiscretionaryDiscountPercentage != nil {
		result["discretionary_discount_percentage"] = string(*obj.DiscretionaryDiscountPercentage)
	}

	if obj.IsTier != nil {
		result["is_tier"] = bool(*obj.IsTier)
	}

	if obj.NetUnitPrice != nil {
		result["net_unit_price"] = string(*obj.NetUnitPrice)
	}

	if obj.OveragePrice != nil {
		result["overage_price"] = string(*obj.OveragePrice)
	}

	if obj.Product != nil {
		result["product"] = []interface{}{RateCardProductToMap(obj.Product)}
	}

	rateCardTiers := []interface{}{}
	for _, item := range obj.RateCardTiers {
		rateCardTiers = append(rateCardTiers, RateCardTierToMap(item))
	}
	result["rate_card_tiers"] = rateCardTiers

	if obj.SubscribedServiceId != nil {
		result["subscribed_service_id"] = string(*obj.SubscribedServiceId)
	}

	if obj.TimeEnd != nil {
		result["time_end"] = obj.TimeEnd.String()
	}

	if obj.TimeStart != nil {
		result["time_start"] = obj.TimeStart.String()
	}

	return result
}

func RateCardTierToMap(obj oci_onesubscription.RateCardTier) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.NetUnitPrice != nil {
		result["net_unit_price"] = string(*obj.NetUnitPrice)
	}

	if obj.OveragePrice != nil {
		result["overage_price"] = string(*obj.OveragePrice)
	}

	if obj.UpToQuantity != nil {
		result["up_to_quantity"] = string(*obj.UpToQuantity)
	}

	return result
}

func SubscribedServiceAddressToMap(obj *oci_onesubscription.SubscribedServiceAddress) map[string]interface{} {
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
		result["location"] = []interface{}{SubscribedServiceLocationToMap(obj.Location)}
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
