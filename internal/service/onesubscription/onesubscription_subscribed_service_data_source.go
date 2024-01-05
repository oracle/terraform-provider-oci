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

func OnesubscriptionSubscribedServiceDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularOnesubscriptionSubscribedService,
		Schema: map[string]*schema.Schema{
			"fields": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"subscribed_service_id": {
				Type:     schema.TypeString,
				Required: true,
			},
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
	}
}

func readSingularOnesubscriptionSubscribedService(d *schema.ResourceData, m interface{}) error {
	sync := &OnesubscriptionSubscribedServiceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SubscribedServiceRegionalClient()

	return tfresource.ReadResource(sync)
}

type OnesubscriptionSubscribedServiceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_onesubscription.SubscribedServiceClient
	Res    *oci_onesubscription.GetSubscribedServiceResponse
}

func (s *OnesubscriptionSubscribedServiceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OnesubscriptionSubscribedServiceDataSourceCrud) Get() error {
	request := oci_onesubscription.GetSubscribedServiceRequest{}

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

	if subscribedServiceId, ok := s.D.GetOkExists("subscribed_service_id"); ok {
		tmp := subscribedServiceId.(string)
		request.SubscribedServiceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "onesubscription")

	response, err := s.Client.GetSubscribedService(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OnesubscriptionSubscribedServiceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AdminEmail != nil {
		s.D.Set("admin_email", *s.Res.AdminEmail)
	}

	if s.Res.AgreementId != nil {
		s.D.Set("agreement_id", strconv.FormatInt(*s.Res.AgreementId, 10))
	}

	if s.Res.AgreementName != nil {
		s.D.Set("agreement_name", *s.Res.AgreementName)
	}

	if s.Res.AgreementType != nil {
		s.D.Set("agreement_type", *s.Res.AgreementType)
	}

	if s.Res.AvailableAmount != nil {
		s.D.Set("available_amount", *s.Res.AvailableAmount)
	}

	if s.Res.BillToAddress != nil {
		s.D.Set("bill_to_address", []interface{}{SubscribedServiceAddressToMap(s.Res.BillToAddress)})
	} else {
		s.D.Set("bill_to_address", nil)
	}

	if s.Res.BillToContact != nil {
		s.D.Set("bill_to_contact", []interface{}{SubscribedServiceUserToMap(s.Res.BillToContact)})
	} else {
		s.D.Set("bill_to_contact", nil)
	}

	if s.Res.BillToCustomer != nil {
		s.D.Set("bill_to_customer", []interface{}{SubscribedServiceBusinessPartnerToMap(s.Res.BillToCustomer)})
	} else {
		s.D.Set("bill_to_customer", nil)
	}

	if s.Res.BillingFrequency != nil {
		s.D.Set("billing_frequency", *s.Res.BillingFrequency)
	}

	if s.Res.BookingOptyNumber != nil {
		s.D.Set("booking_opty_number", *s.Res.BookingOptyNumber)
	}

	if s.Res.BuyerEmail != nil {
		s.D.Set("buyer_email", *s.Res.BuyerEmail)
	}

	if s.Res.CommitmentScheduleId != nil {
		s.D.Set("commitment_schedule_id", *s.Res.CommitmentScheduleId)
	}

	commitmentServices := []interface{}{}
	for _, item := range s.Res.CommitmentServices {
		commitmentServices = append(commitmentServices, CommitmentServiceToMap(item))
	}
	s.D.Set("commitment_services", commitmentServices)

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.CreditPercentage != nil {
		s.D.Set("credit_percentage", *s.Res.CreditPercentage)
	}

	if s.Res.Csi != nil {
		s.D.Set("csi", strconv.FormatInt(*s.Res.Csi, 10))
	}

	if s.Res.CustomerTransactionReference != nil {
		s.D.Set("customer_transaction_reference", *s.Res.CustomerTransactionReference)
	}

	if s.Res.DataCenter != nil {
		s.D.Set("data_center", *s.Res.DataCenter)
	}

	if s.Res.DataCenterRegion != nil {
		s.D.Set("data_center_region", *s.Res.DataCenterRegion)
	}

	if s.Res.EligibleToRenew != nil {
		s.D.Set("eligible_to_renew", *s.Res.EligibleToRenew)
	}

	if s.Res.EndUserAddress != nil {
		s.D.Set("end_user_address", []interface{}{SubscribedServiceAddressToMap(s.Res.EndUserAddress)})
	} else {
		s.D.Set("end_user_address", nil)
	}

	if s.Res.EndUserContact != nil {
		s.D.Set("end_user_contact", []interface{}{SubscribedServiceUserToMap(s.Res.EndUserContact)})
	} else {
		s.D.Set("end_user_contact", nil)
	}

	if s.Res.EndUserCustomer != nil {
		s.D.Set("end_user_customer", []interface{}{SubscribedServiceBusinessPartnerToMap(s.Res.EndUserCustomer)})
	} else {
		s.D.Set("end_user_customer", nil)
	}

	if s.Res.FulfillmentSet != nil {
		s.D.Set("fulfillment_set", *s.Res.FulfillmentSet)
	}

	if s.Res.FundedAllocationValue != nil {
		s.D.Set("funded_allocation_value", *s.Res.FundedAllocationValue)
	}

	if s.Res.IsAllowance != nil {
		s.D.Set("is_allowance", *s.Res.IsAllowance)
	}

	if s.Res.IsCapToPriceList != nil {
		s.D.Set("is_cap_to_price_list", *s.Res.IsCapToPriceList)
	}

	if s.Res.IsCreditEnabled != nil {
		s.D.Set("is_credit_enabled", *s.Res.IsCreditEnabled)
	}

	if s.Res.IsHavingUsage != nil {
		s.D.Set("is_having_usage", *s.Res.IsHavingUsage)
	}

	if s.Res.IsIntentToPay != nil {
		s.D.Set("is_intent_to_pay", *s.Res.IsIntentToPay)
	}

	if s.Res.IsPayg != nil {
		s.D.Set("is_payg", *s.Res.IsPayg)
	}

	if s.Res.IsSingleRateCard != nil {
		s.D.Set("is_single_rate_card", *s.Res.IsSingleRateCard)
	}

	if s.Res.IsVariableCommitment != nil {
		s.D.Set("is_variable_commitment", *s.Res.IsVariableCommitment)
	}

	if s.Res.LineNetAmount != nil {
		s.D.Set("line_net_amount", *s.Res.LineNetAmount)
	}

	if s.Res.MajorSet != nil {
		s.D.Set("major_set", strconv.FormatInt(*s.Res.MajorSet, 10))
	}

	if s.Res.NetUnitPrice != nil {
		s.D.Set("net_unit_price", *s.Res.NetUnitPrice)
	}

	if s.Res.OperationType != nil {
		s.D.Set("operation_type", *s.Res.OperationType)
	}

	if s.Res.OrderHeaderId != nil {
		s.D.Set("order_header_id", strconv.FormatInt(*s.Res.OrderHeaderId, 10))
	}

	if s.Res.OrderLineId != nil {
		s.D.Set("order_line_id", strconv.FormatInt(*s.Res.OrderLineId, 10))
	}

	if s.Res.OrderLineNumber != nil {
		s.D.Set("order_line_number", *s.Res.OrderLineNumber)
	}

	if s.Res.OrderNumber != nil {
		s.D.Set("order_number", strconv.FormatInt(*s.Res.OrderNumber, 10))
	}

	if s.Res.OrderType != nil {
		s.D.Set("order_type", *s.Res.OrderType)
	}

	if s.Res.OriginalPromoAmount != nil {
		s.D.Set("original_promo_amount", *s.Res.OriginalPromoAmount)
	}

	if s.Res.OverageBillTo != nil {
		s.D.Set("overage_bill_to", *s.Res.OverageBillTo)
	}

	if s.Res.OverageDiscountPercentage != nil {
		s.D.Set("overage_discount_percentage", *s.Res.OverageDiscountPercentage)
	}

	if s.Res.OveragePolicy != nil {
		s.D.Set("overage_policy", *s.Res.OveragePolicy)
	}

	if s.Res.PartnerCreditAmount != nil {
		s.D.Set("partner_credit_amount", *s.Res.PartnerCreditAmount)
	}

	if s.Res.PartnerTransactionType != nil {
		s.D.Set("partner_transaction_type", *s.Res.PartnerTransactionType)
	}

	if s.Res.PaygPolicy != nil {
		s.D.Set("payg_policy", *s.Res.PaygPolicy)
	}

	if s.Res.PaymentMethod != nil {
		s.D.Set("payment_method", *s.Res.PaymentMethod)
	}

	if s.Res.PaymentNumber != nil {
		s.D.Set("payment_number", *s.Res.PaymentNumber)
	}

	if s.Res.PaymentTerm != nil {
		s.D.Set("payment_term", []interface{}{SubscribedServicePaymentTermToMap(s.Res.PaymentTerm)})
	} else {
		s.D.Set("payment_term", nil)
	}

	if s.Res.PricePeriod != nil {
		s.D.Set("price_period", *s.Res.PricePeriod)
	}

	if s.Res.PricingModel != nil {
		s.D.Set("pricing_model", *s.Res.PricingModel)
	}

	if s.Res.Product != nil {
		s.D.Set("product", []interface{}{RateCardProductToMap(s.Res.Product)})
	} else {
		s.D.Set("product", nil)
	}

	if s.Res.ProgramType != nil {
		s.D.Set("program_type", *s.Res.ProgramType)
	}

	if s.Res.PromoOrderLineId != nil {
		s.D.Set("promo_order_line_id", strconv.FormatInt(*s.Res.PromoOrderLineId, 10))
	}

	if s.Res.PromoType != nil {
		s.D.Set("promo_type", *s.Res.PromoType)
	}

	if s.Res.PromotionPricingType != nil {
		s.D.Set("promotion_pricing_type", *s.Res.PromotionPricingType)
	}

	if s.Res.ProvisioningSource != nil {
		s.D.Set("provisioning_source", *s.Res.ProvisioningSource)
	}

	if s.Res.Quantity != nil {
		s.D.Set("quantity", *s.Res.Quantity)
	}

	if s.Res.RateCardDiscountPercentage != nil {
		s.D.Set("rate_card_discount_percentage", *s.Res.RateCardDiscountPercentage)
	}

	rateCards := []interface{}{}
	for _, item := range s.Res.RateCards {
		rateCards = append(rateCards, RateCardSummaryToMap(item))
	}
	s.D.Set("rate_cards", rateCards)

	if s.Res.RatecardType != nil {
		s.D.Set("ratecard_type", *s.Res.RatecardType)
	}

	if s.Res.RenewalOptyId != nil {
		s.D.Set("renewal_opty_id", strconv.FormatInt(*s.Res.RenewalOptyId, 10))
	}

	if s.Res.RenewalOptyNumber != nil {
		s.D.Set("renewal_opty_number", *s.Res.RenewalOptyNumber)
	}

	if s.Res.RenewalOptyType != nil {
		s.D.Set("renewal_opty_type", *s.Res.RenewalOptyType)
	}

	if s.Res.RenewedSubscribedServiceId != nil {
		s.D.Set("renewed_subscribed_service_id", *s.Res.RenewedSubscribedServiceId)
	}

	if s.Res.ResellerAddress != nil {
		s.D.Set("reseller_address", []interface{}{SubscribedServiceAddressToMap(s.Res.ResellerAddress)})
	} else {
		s.D.Set("reseller_address", nil)
	}

	if s.Res.ResellerContact != nil {
		s.D.Set("reseller_contact", []interface{}{SubscribedServiceUserToMap(s.Res.ResellerContact)})
	} else {
		s.D.Set("reseller_contact", nil)
	}

	if s.Res.ResellerCustomer != nil {
		s.D.Set("reseller_customer", []interface{}{SubscribedServiceBusinessPartnerToMap(s.Res.ResellerCustomer)})
	} else {
		s.D.Set("reseller_customer", nil)
	}

	if s.Res.RevenueLineId != nil {
		s.D.Set("revenue_line_id", strconv.FormatInt(*s.Res.RevenueLineId, 10))
	}

	if s.Res.RevenueLineNumber != nil {
		s.D.Set("revenue_line_number", *s.Res.RevenueLineNumber)
	}

	if s.Res.RevisedArrInLc != nil {
		s.D.Set("revised_arr_in_lc", *s.Res.RevisedArrInLc)
	}

	if s.Res.RevisedArrInSc != nil {
		s.D.Set("revised_arr_in_sc", *s.Res.RevisedArrInSc)
	}

	if s.Res.SalesAccountPartyId != nil {
		s.D.Set("sales_account_party_id", strconv.FormatInt(*s.Res.SalesAccountPartyId, 10))
	}

	if s.Res.SalesChannel != nil {
		s.D.Set("sales_channel", *s.Res.SalesChannel)
	}

	if s.Res.SerialNumber != nil {
		s.D.Set("serial_number", *s.Res.SerialNumber)
	}

	if s.Res.ServiceToAddress != nil {
		s.D.Set("service_to_address", []interface{}{SubscribedServiceAddressToMap(s.Res.ServiceToAddress)})
	} else {
		s.D.Set("service_to_address", nil)
	}

	if s.Res.ServiceToContact != nil {
		s.D.Set("service_to_contact", []interface{}{SubscribedServiceUserToMap(s.Res.ServiceToContact)})
	} else {
		s.D.Set("service_to_contact", nil)
	}

	if s.Res.ServiceToCustomer != nil {
		s.D.Set("service_to_customer", []interface{}{SubscribedServiceBusinessPartnerToMap(s.Res.ServiceToCustomer)})
	} else {
		s.D.Set("service_to_customer", nil)
	}

	if s.Res.SoldToContact != nil {
		s.D.Set("sold_to_contact", []interface{}{SubscribedServiceUserToMap(s.Res.SoldToContact)})
	} else {
		s.D.Set("sold_to_contact", nil)
	}

	if s.Res.SoldToCustomer != nil {
		s.D.Set("sold_to_customer", []interface{}{SubscribedServiceBusinessPartnerToMap(s.Res.SoldToCustomer)})
	} else {
		s.D.Set("sold_to_customer", nil)
	}

	if s.Res.StartDateType != nil {
		s.D.Set("start_date_type", *s.Res.StartDateType)
	}

	if s.Res.Status != nil {
		s.D.Set("status", *s.Res.Status)
	}

	if s.Res.SubscriptionId != nil {
		s.D.Set("subscription_id", *s.Res.SubscriptionId)
	}

	if s.Res.SubscriptionSource != nil {
		s.D.Set("subscription_source", *s.Res.SubscriptionSource)
	}

	if s.Res.SystemArrInLc != nil {
		s.D.Set("system_arr_in_lc", *s.Res.SystemArrInLc)
	}

	if s.Res.SystemArrInSc != nil {
		s.D.Set("system_arr_in_sc", *s.Res.SystemArrInSc)
	}

	if s.Res.SystemAtrArrInLc != nil {
		s.D.Set("system_atr_arr_in_lc", *s.Res.SystemAtrArrInLc)
	}

	if s.Res.SystemAtrArrInSc != nil {
		s.D.Set("system_atr_arr_in_sc", *s.Res.SystemAtrArrInSc)
	}

	if s.Res.TermValue != nil {
		s.D.Set("term_value", strconv.FormatInt(*s.Res.TermValue, 10))
	}

	if s.Res.TermValueUom != nil {
		s.D.Set("term_value_uom", *s.Res.TermValueUom)
	}

	if s.Res.TimeAgreementEnd != nil {
		s.D.Set("time_agreement_end", s.Res.TimeAgreementEnd.String())
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeCustomerConfig != nil {
		s.D.Set("time_customer_config", s.Res.TimeCustomerConfig.String())
	}

	if s.Res.TimeEnd != nil {
		s.D.Set("time_end", s.Res.TimeEnd.String())
	}

	if s.Res.TimeMajorsetEnd != nil {
		s.D.Set("time_majorset_end", s.Res.TimeMajorsetEnd.String())
	}

	if s.Res.TimeMajorsetStart != nil {
		s.D.Set("time_majorset_start", s.Res.TimeMajorsetStart.String())
	}

	if s.Res.TimePaymentExpiry != nil {
		s.D.Set("time_payment_expiry", s.Res.TimePaymentExpiry.String())
	}

	if s.Res.TimeProvisioned != nil {
		s.D.Set("time_provisioned", s.Res.TimeProvisioned.String())
	}

	if s.Res.TimeServiceConfigurationEmailSent != nil {
		s.D.Set("time_service_configuration_email_sent", s.Res.TimeServiceConfigurationEmailSent.String())
	}

	if s.Res.TimeStart != nil {
		s.D.Set("time_start", s.Res.TimeStart.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TimeWelcomeEmailSent != nil {
		s.D.Set("time_welcome_email_sent", s.Res.TimeWelcomeEmailSent.String())
	}

	if s.Res.TotalValue != nil {
		s.D.Set("total_value", *s.Res.TotalValue)
	}

	if s.Res.TransactionExtensionId != nil {
		s.D.Set("transaction_extension_id", strconv.FormatInt(*s.Res.TransactionExtensionId, 10))
	}

	if s.Res.Type != nil {
		s.D.Set("type", *s.Res.Type)
	}

	if s.Res.UpdatedBy != nil {
		s.D.Set("updated_by", *s.Res.UpdatedBy)
	}

	if s.Res.UsedAmount != nil {
		s.D.Set("used_amount", *s.Res.UsedAmount)
	}

	return nil
}
