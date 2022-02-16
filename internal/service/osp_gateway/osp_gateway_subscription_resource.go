// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osp_gateway

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_osp_gateway "github.com/oracle/oci-go-sdk/v58/ospgateway"
)

func OspGatewaySubscriptionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOspGatewaySubscription,
		Read:     readOspGatewaySubscription,
		Update:   updateOspGatewaySubscription,
		Delete:   deleteOspGatewaySubscription,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"osp_home_region": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subscription": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"subscription_plan_number": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"bill_to_cust_account_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"billing_address": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"address_key": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"city": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"company_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"country": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"email_address": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"first_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"last_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"line1": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"line2": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"postal_code": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"currency_code": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"gsi_org_code": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"is_intent_to_pay": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"language_code": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"organization_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"payment_gateway": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"merchant_defined_data": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"cloud_account_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"promo_type": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},

									// Computed
								},
							},
						},
						"payment_options": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"payment_method": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"wallet_instrument_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"wallet_transaction_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"plan_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ship_to_cust_acct_role_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ship_to_cust_acct_site_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tax_info": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"no_tax_reason_code": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"no_tax_reason_code_details": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tax_cnpj": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tax_payer_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tax_reg_number": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"time_plan_upgrade": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: utils.TimeDiffSuppressFunction,
						},
						"time_start": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: utils.TimeDiffSuppressFunction,
						},
						"upgrade_state": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"upgrade_state_details": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"bill_to_cust_account_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"billing_address": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"address_key": {
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
						"country": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"email_address": {
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
						"line1": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"line2": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"postal_code": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"currency_code": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gsi_org_code": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_intent_to_pay": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"language_code": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"organization_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"payment_gateway": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"merchant_defined_data": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"cloud_account_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"promo_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"payment_options": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"payment_method": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"wallet_instrument_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"wallet_transaction_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"plan_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ship_to_cust_acct_role_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ship_to_cust_acct_site_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subscription_plan_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tax_info": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"no_tax_reason_code": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"no_tax_reason_code_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tax_cnpj": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tax_payer_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tax_reg_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"time_plan_upgrade": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_start": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"upgrade_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"upgrade_state_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createOspGatewaySubscription(d *schema.ResourceData, m interface{}) error {
	sync := &OspGatewaySubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SubscriptionServiceClient()

	return tfresource.CreateResource(d, sync)
}

func readOspGatewaySubscription(d *schema.ResourceData, m interface{}) error {
	sync := &OspGatewaySubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SubscriptionServiceClient()

	return tfresource.ReadResource(sync)
}

func updateOspGatewaySubscription(d *schema.ResourceData, m interface{}) error {
	sync := &OspGatewaySubscriptionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SubscriptionServiceClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOspGatewaySubscription(d *schema.ResourceData, m interface{}) error {
	return nil
}

type OspGatewaySubscriptionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_osp_gateway.SubscriptionServiceClient
	Res                    *oci_osp_gateway.Subscription
	DisableNotFoundRetries bool
}

func (s *OspGatewaySubscriptionResourceCrud) ID() string {
	return GetSubscriptionCompositeId(s.D.Get("id").(string))
}

func (s *OspGatewaySubscriptionResourceCrud) Create() error {
	request := oci_osp_gateway.UpdateSubscriptionRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if email, ok := s.D.GetOkExists("email"); ok {
		tmp := email.(string)
		request.Email = &tmp
	}

	if ospHomeRegion, ok := s.D.GetOkExists("osp_home_region"); ok {
		tmp := ospHomeRegion.(string)
		request.OspHomeRegion = &tmp
	}

	if subscription, ok := s.D.GetOkExists("subscription"); ok {
		if tmpList := subscription.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "subscription", 0)
			tmp, err := s.mapToSubscription(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Subscription = &tmp
		}
	}

	if subscriptionId, ok := s.D.GetOkExists("id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "osp_gateway")

	response, err := s.Client.UpdateSubscription(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Subscription
	return nil
}

func (s *OspGatewaySubscriptionResourceCrud) Get() error {
	request := oci_osp_gateway.GetSubscriptionRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if ospHomeRegion, ok := s.D.GetOkExists("osp_home_region"); ok {
		tmp := ospHomeRegion.(string)
		request.OspHomeRegion = &tmp
	}

	tmp := s.D.Id()
	request.SubscriptionId = &tmp

	subscriptionId, err := parseSubscriptionCompositeId(s.D.Id())
	if err == nil {
		request.SubscriptionId = &subscriptionId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "osp_gateway")

	response, err := s.Client.GetSubscription(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Subscription
	return nil
}

func (s *OspGatewaySubscriptionResourceCrud) Update() error {
	request := oci_osp_gateway.UpdateSubscriptionRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if email, ok := s.D.GetOkExists("email"); ok {
		tmp := email.(string)
		request.Email = &tmp
	}

	if ospHomeRegion, ok := s.D.GetOkExists("osp_home_region"); ok {
		tmp := ospHomeRegion.(string)
		request.OspHomeRegion = &tmp
	}

	if subscription, ok := s.D.GetOkExists("subscription"); ok {
		if tmpList := subscription.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "subscription", 0)
			tmp, err := s.mapToSubscription(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Subscription = &tmp
		}
	}

	tmp := s.D.Id()
	request.SubscriptionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "osp_gateway")

	response, err := s.Client.UpdateSubscription(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Subscription
	return nil
}

func (s *OspGatewaySubscriptionResourceCrud) SetData() error {

	subscriptionId, err := parseSubscriptionCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(subscriptionId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.BillToCustAccountId != nil {
		s.D.Set("bill_to_cust_account_id", *s.Res.BillToCustAccountId)
	}

	if s.Res.BillingAddress != nil {
		s.D.Set("billing_address", []interface{}{BillingAddressToMap(s.Res.BillingAddress)})
	} else {
		s.D.Set("billing_address", nil)
	}

	if s.Res.CurrencyCode != nil {
		s.D.Set("currency_code", *s.Res.CurrencyCode)
	}

	if s.Res.GsiOrgCode != nil {
		s.D.Set("gsi_org_code", *s.Res.GsiOrgCode)
	}

	if s.Res.IsIntentToPay != nil {
		s.D.Set("is_intent_to_pay", *s.Res.IsIntentToPay)
	}

	if s.Res.LanguageCode != nil {
		s.D.Set("language_code", *s.Res.LanguageCode)
	}

	if s.Res.OrganizationId != nil {
		s.D.Set("organization_id", *s.Res.OrganizationId)
	}

	if s.Res.PaymentGateway != nil {
		s.D.Set("payment_gateway", []interface{}{PaymentGatewayToMap(s.Res.PaymentGateway)})
	} else {
		s.D.Set("payment_gateway", nil)
	}

	paymentOptions := []interface{}{}
	for _, item := range s.Res.PaymentOptions {
		paymentOptions = append(paymentOptions, PaymentOptionToMap(item))
	}
	s.D.Set("payment_options", paymentOptions)

	s.D.Set("plan_type", s.Res.PlanType)

	if s.Res.ShipToCustAcctRoleId != nil {
		s.D.Set("ship_to_cust_acct_role_id", *s.Res.ShipToCustAcctRoleId)
	}

	if s.Res.ShipToCustAcctSiteId != nil {
		s.D.Set("ship_to_cust_acct_site_id", *s.Res.ShipToCustAcctSiteId)
	}

	if s.Res.SubscriptionPlanNumber != nil {
		s.D.Set("subscription_plan_number", *s.Res.SubscriptionPlanNumber)
	}

	if s.Res.TaxInfo != nil {
		s.D.Set("tax_info", []interface{}{TaxInfoToMap(s.Res.TaxInfo)})
	} else {
		s.D.Set("tax_info", nil)
	}

	if s.Res.TimePlanUpgrade != nil {
		s.D.Set("time_plan_upgrade", s.Res.TimePlanUpgrade.String())
	}

	if s.Res.TimeStart != nil {
		s.D.Set("time_start", s.Res.TimeStart.String())
	}

	s.D.Set("upgrade_state", s.Res.UpgradeState)

	s.D.Set("upgrade_state_details", s.Res.UpgradeStateDetails)

	return nil
}

func GetSubscriptionCompositeId(subscriptionId string) string {
	subscriptionId = url.PathEscape(subscriptionId)
	compositeId := "subscriptions/" + subscriptionId
	return compositeId
}

func parseSubscriptionCompositeId(compositeId string) (subscriptionId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("subscriptions/.*/compartmentId/.*/ospHomeRegion/.*", compositeId)
	if !match || len(parts) != 6 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	subscriptionId, _ = url.PathUnescape(parts[1])

	return
}

func (s *OspGatewaySubscriptionResourceCrud) mapToBillingAddress(fieldKeyFormat string) (oci_osp_gateway.BillingAddress, error) {
	result := oci_osp_gateway.BillingAddress{}

	if addressKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "address_key")); ok {
		tmp := addressKey.(string)
		result.AddressKey = &tmp
	}

	if city, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "city")); ok {
		tmp := city.(string)
		result.City = &tmp
	}

	if companyName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "company_name")); ok {
		tmp := companyName.(string)
		result.CompanyName = &tmp
	}

	if country, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "country")); ok {
		tmp := country.(string)
		result.Country = &tmp
	}

	if emailAddress, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "email_address")); ok {
		tmp := emailAddress.(string)
		result.EmailAddress = &tmp
	}

	if firstName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "first_name")); ok {
		tmp := firstName.(string)
		result.FirstName = &tmp
	}

	if lastName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "last_name")); ok {
		tmp := lastName.(string)
		result.LastName = &tmp
	}

	if line1, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "line1")); ok {
		tmp := line1.(string)
		result.Line1 = &tmp
	}

	if line2, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "line2")); ok {
		tmp := line2.(string)
		result.Line2 = &tmp
	}

	if postalCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "postal_code")); ok {
		tmp := postalCode.(string)
		result.PostalCode = &tmp
	}

	if state, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "state")); ok {
		tmp := state.(string)
		result.State = &tmp
	}

	return result, nil
}

func BillingAddressToMap(obj *oci_osp_gateway.BillingAddress) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AddressKey != nil {
		result["address_key"] = string(*obj.AddressKey)
	}

	if obj.City != nil {
		result["city"] = string(*obj.City)
	}

	if obj.CompanyName != nil {
		result["company_name"] = string(*obj.CompanyName)
	}

	if obj.Country != nil {
		result["country"] = string(*obj.Country)
	}

	if obj.EmailAddress != nil {
		result["email_address"] = string(*obj.EmailAddress)
	}

	if obj.FirstName != nil {
		result["first_name"] = string(*obj.FirstName)
	}

	if obj.LastName != nil {
		result["last_name"] = string(*obj.LastName)
	}

	if obj.Line1 != nil {
		result["line1"] = string(*obj.Line1)
	}

	if obj.Line2 != nil {
		result["line2"] = string(*obj.Line2)
	}

	if obj.PostalCode != nil {
		result["postal_code"] = string(*obj.PostalCode)
	}

	if obj.State != nil {
		result["state"] = string(*obj.State)
	}

	return result
}

func (s *OspGatewaySubscriptionResourceCrud) mapToMerchantDefinedData(fieldKeyFormat string) (oci_osp_gateway.MerchantDefinedData, error) {
	result := oci_osp_gateway.MerchantDefinedData{}

	if cloudAccountName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cloud_account_name")); ok {
		tmp := cloudAccountName.(string)
		result.CloudAccountName = &tmp
	}

	if promoType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "promo_type")); ok {
		tmp := promoType.(string)
		result.PromoType = &tmp
	}

	return result, nil
}

func MerchantDefinedDataToMap(obj *oci_osp_gateway.MerchantDefinedData) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CloudAccountName != nil {
		result["cloud_account_name"] = string(*obj.CloudAccountName)
	}

	if obj.PromoType != nil {
		result["promo_type"] = string(*obj.PromoType)
	}

	return result
}

func (s *OspGatewaySubscriptionResourceCrud) mapToPaymentGateway(fieldKeyFormat string) (oci_osp_gateway.PaymentGateway, error) {
	result := oci_osp_gateway.PaymentGateway{}

	if merchantDefinedData, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "merchant_defined_data")); ok {
		if tmpList := merchantDefinedData.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "merchant_defined_data"), 0)
			tmp, err := s.mapToMerchantDefinedData(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert merchant_defined_data, encountered error: %v", err)
			}
			result.MerchantDefinedData = &tmp
		}
	}

	return result, nil
}

func PaymentGatewayToMap(obj *oci_osp_gateway.PaymentGateway) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MerchantDefinedData != nil {
		result["merchant_defined_data"] = []interface{}{MerchantDefinedDataToMap(obj.MerchantDefinedData)}
	}

	return result
}

func (s *OspGatewaySubscriptionResourceCrud) mapToPaymentOption(fieldKeyFormat string) (oci_osp_gateway.PaymentOption, error) {
	var baseObject oci_osp_gateway.PaymentOption
	//discriminator
	paymentMethodRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "payment_method"))
	var paymentMethod string
	if ok {
		paymentMethod = paymentMethodRaw.(string)
	} else {
		paymentMethod = "" // default value
	}
	switch strings.ToLower(paymentMethod) {
	case strings.ToLower("CREDIT_CARD"):
		details := oci_osp_gateway.CreditCardPaymentOption{}
		baseObject = details
	case strings.ToLower("PAYPAL"):
		details := oci_osp_gateway.PaypalPaymentOption{}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown payment_method '%v' was specified", paymentMethod)
	}
	return baseObject, nil
}

func PaymentOptionToMap(obj oci_osp_gateway.PaymentOption) map[string]interface{} {
	result := map[string]interface{}{}
	switch (obj).(type) {
	case oci_osp_gateway.CreditCardPaymentOption:
		result["payment_method"] = "CREDIT_CARD"
	case oci_osp_gateway.PaypalPaymentOption:
		result["payment_method"] = "PAYPAL"
	default:
		log.Printf("[WARN] Received 'payment_method' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *OspGatewaySubscriptionResourceCrud) mapToSubscription(fieldKeyFormat string) (oci_osp_gateway.Subscription, error) {
	result := oci_osp_gateway.Subscription{}

	if billToCustAccountId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bill_to_cust_account_id")); ok {
		tmp := billToCustAccountId.(string)
		result.BillToCustAccountId = &tmp
	}

	if billingAddress, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "billing_address")); ok {
		if tmpList := billingAddress.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "billing_address"), 0)
			tmp, err := s.mapToBillingAddress(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert billing_address, encountered error: %v", err)
			}
			result.BillingAddress = &tmp
		}
	}

	if currencyCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "currency_code")); ok {
		tmp := currencyCode.(string)
		result.CurrencyCode = &tmp
	}

	if gsiOrgCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "gsi_org_code")); ok {
		tmp := gsiOrgCode.(string)
		result.GsiOrgCode = &tmp
	}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	if isIntentToPay, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_intent_to_pay")); ok {
		tmp := isIntentToPay.(bool)
		result.IsIntentToPay = &tmp
	}

	if languageCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "language_code")); ok {
		tmp := languageCode.(string)
		result.LanguageCode = &tmp
	}

	if organizationId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "organization_id")); ok {
		tmp := organizationId.(string)
		result.OrganizationId = &tmp
	}

	if paymentGateway, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "payment_gateway")); ok {
		if tmpList := paymentGateway.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "payment_gateway"), 0)
			tmp, err := s.mapToPaymentGateway(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert payment_gateway, encountered error: %v", err)
			}
			result.PaymentGateway = &tmp
		}
	}

	if paymentOptions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "payment_options")); ok {
		interfaces := paymentOptions.([]interface{})
		tmp := make([]oci_osp_gateway.PaymentOption, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "payment_options"), stateDataIndex)
			converted, err := s.mapToPaymentOption(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "payment_options")) {
			result.PaymentOptions = tmp
		}
	}

	if planType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "plan_type")); ok {
		result.PlanType = oci_osp_gateway.SubscriptionPlanTypeEnum(planType.(string))
	}

	if shipToCustAcctRoleId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ship_to_cust_acct_role_id")); ok {
		tmp := shipToCustAcctRoleId.(string)
		result.ShipToCustAcctRoleId = &tmp
	}

	if shipToCustAcctSiteId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ship_to_cust_acct_site_id")); ok {
		tmp := shipToCustAcctSiteId.(string)
		result.ShipToCustAcctSiteId = &tmp
	}

	if subscriptionPlanNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subscription_plan_number")); ok {
		tmp := subscriptionPlanNumber.(string)
		result.SubscriptionPlanNumber = &tmp
	}

	if taxInfo, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tax_info")); ok {
		if tmpList := taxInfo.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "tax_info"), 0)
			tmp, err := s.mapToTaxInfo(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert tax_info, encountered error: %v", err)
			}
			result.TaxInfo = &tmp
		}
	}

	if timePlanUpgrade, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_plan_upgrade")); ok {
		tmp, err := time.Parse(time.RFC3339, timePlanUpgrade.(string))
		if err != nil {
			return result, err
		}
		result.TimePlanUpgrade = &oci_common.SDKTime{Time: tmp}
	}

	if timeStart, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_start")); ok {
		tmp, err := time.Parse(time.RFC3339, timeStart.(string))
		if err != nil {
			return result, err
		}
		result.TimeStart = &oci_common.SDKTime{Time: tmp}
	}

	if upgradeState, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "upgrade_state")); ok {
		result.UpgradeState = oci_osp_gateway.SubscriptionUpgradeStateEnum(upgradeState.(string))
	}

	if upgradeStateDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "upgrade_state_details")); ok {
		result.UpgradeStateDetails = oci_osp_gateway.SubscriptionUpgradeStateDetailsEnum(upgradeStateDetails.(string))
	}

	return result, nil
}

func (s *OspGatewaySubscriptionResourceCrud) mapToTaxInfo(fieldKeyFormat string) (oci_osp_gateway.TaxInfo, error) {
	result := oci_osp_gateway.TaxInfo{}

	if noTaxReasonCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "no_tax_reason_code")); ok {
		tmp := noTaxReasonCode.(string)
		result.NoTaxReasonCode = &tmp
	}

	if noTaxReasonCodeDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "no_tax_reason_code_details")); ok {
		tmp := noTaxReasonCodeDetails.(string)
		result.NoTaxReasonCodeDetails = &tmp
	}

	if taxCnpj, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tax_cnpj")); ok {
		tmp := taxCnpj.(string)
		result.TaxCnpj = &tmp
	}

	if taxPayerId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tax_payer_id")); ok {
		tmp := taxPayerId.(string)
		result.TaxPayerId = &tmp
	}

	if taxRegNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tax_reg_number")); ok {
		tmp := taxRegNumber.(string)
		result.TaxRegNumber = &tmp
	}

	return result, nil
}

func SubscriptionToMap(obj *oci_osp_gateway.Subscription) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BillToCustAccountId != nil {
		result["bill_to_cust_account_id"] = string(*obj.BillToCustAccountId)
	}

	if obj.BillingAddress != nil {
		result["billing_address"] = []interface{}{BillingAddressToMap(obj.BillingAddress)}
	}

	if obj.CurrencyCode != nil {
		result["currency_code"] = string(*obj.CurrencyCode)
	}

	if obj.GsiOrgCode != nil {
		result["gsi_org_code"] = string(*obj.GsiOrgCode)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsIntentToPay != nil {
		result["is_intent_to_pay"] = bool(*obj.IsIntentToPay)
	}

	if obj.LanguageCode != nil {
		result["language_code"] = string(*obj.LanguageCode)
	}

	if obj.OrganizationId != nil {
		result["organization_id"] = string(*obj.OrganizationId)
	}

	if obj.PaymentGateway != nil {
		result["payment_gateway"] = []interface{}{PaymentGatewayToMap(obj.PaymentGateway)}
	}

	paymentOptions := []interface{}{}
	for _, item := range obj.PaymentOptions {
		paymentOptions = append(paymentOptions, PaymentOptionToMap(item))
	}
	result["payment_options"] = paymentOptions

	result["plan_type"] = string(obj.PlanType)

	if obj.ShipToCustAcctRoleId != nil {
		result["ship_to_cust_acct_role_id"] = string(*obj.ShipToCustAcctRoleId)
	}

	if obj.ShipToCustAcctSiteId != nil {
		result["ship_to_cust_acct_site_id"] = string(*obj.ShipToCustAcctSiteId)
	}

	if obj.SubscriptionPlanNumber != nil {
		result["subscription_plan_number"] = string(*obj.SubscriptionPlanNumber)
	}

	if obj.TaxInfo != nil {
		result["tax_info"] = []interface{}{TaxInfoToMap(obj.TaxInfo)}
	}

	if obj.TimePlanUpgrade != nil {
		result["time_plan_upgrade"] = obj.TimePlanUpgrade.Format(time.RFC3339Nano)
	}

	if obj.TimeStart != nil {
		result["time_start"] = obj.TimeStart.Format(time.RFC3339Nano)
	}

	result["upgrade_state"] = string(obj.UpgradeState)

	result["upgrade_state_details"] = string(obj.UpgradeStateDetails)

	return result
}

func SubscriptionSummaryToMap(obj oci_osp_gateway.SubscriptionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BillToCustAccountId != nil {
		result["bill_to_cust_account_id"] = string(*obj.BillToCustAccountId)
	}

	if obj.BillingAddress != nil {
		result["billing_address"] = []interface{}{BillingAddressToMap(obj.BillingAddress)}
	}

	if obj.CurrencyCode != nil {
		result["currency_code"] = string(*obj.CurrencyCode)
	}

	if obj.GsiOrgCode != nil {
		result["gsi_org_code"] = string(*obj.GsiOrgCode)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsIntentToPay != nil {
		result["is_intent_to_pay"] = bool(*obj.IsIntentToPay)
	}

	if obj.LanguageCode != nil {
		result["language_code"] = string(*obj.LanguageCode)
	}

	if obj.OrganizationId != nil {
		result["organization_id"] = string(*obj.OrganizationId)
	}

	if obj.PaymentGateway != nil {
		result["payment_gateway"] = []interface{}{PaymentGatewayToMap(obj.PaymentGateway)}
	}

	paymentOptions := []interface{}{}
	for _, item := range obj.PaymentOptions {
		paymentOptions = append(paymentOptions, PaymentOptionToMap(item))
	}
	result["payment_options"] = paymentOptions

	result["plan_type"] = string(obj.PlanType)

	if obj.ShipToCustAcctRoleId != nil {
		result["ship_to_cust_acct_role_id"] = string(*obj.ShipToCustAcctRoleId)
	}

	if obj.ShipToCustAcctSiteId != nil {
		result["ship_to_cust_acct_site_id"] = string(*obj.ShipToCustAcctSiteId)
	}

	if obj.SubscriptionPlanNumber != nil {
		result["subscription_plan_number"] = string(*obj.SubscriptionPlanNumber)
	}

	if obj.TaxInfo != nil {
		result["tax_info"] = []interface{}{TaxInfoToMap(obj.TaxInfo)}
	}

	if obj.TimePlanUpgrade != nil {
		result["time_plan_upgrade"] = obj.TimePlanUpgrade.String()
	}

	if obj.TimeStart != nil {
		result["time_start"] = obj.TimeStart.String()
	}

	result["upgrade_state"] = string(obj.UpgradeState)

	result["upgrade_state_details"] = string(obj.UpgradeStateDetails)

	return result
}

func TaxInfoToMap(obj *oci_osp_gateway.TaxInfo) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.NoTaxReasonCode != nil {
		result["no_tax_reason_code"] = string(*obj.NoTaxReasonCode)
	}

	if obj.NoTaxReasonCodeDetails != nil {
		result["no_tax_reason_code_details"] = string(*obj.NoTaxReasonCodeDetails)
	}

	if obj.TaxCnpj != nil {
		result["tax_cnpj"] = string(*obj.TaxCnpj)
	}

	if obj.TaxPayerId != nil {
		result["tax_payer_id"] = string(*obj.TaxPayerId)
	}

	if obj.TaxRegNumber != nil {
		result["tax_reg_number"] = string(*obj.TaxRegNumber)
	}

	return result
}
