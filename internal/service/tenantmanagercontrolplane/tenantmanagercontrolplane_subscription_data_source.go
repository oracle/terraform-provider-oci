// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package tenantmanagercontrolplane

import (
	"context"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_tenantmanagercontrolplane "github.com/oracle/oci-go-sdk/v65/tenantmanagercontrolplane"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func TenantmanagercontrolplaneSubscriptionDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularTenantmanagercontrolplaneSubscription,
		Schema: map[string]*schema.Schema{
			"subscription_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"classic_subscription_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloud_amount_currency": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"csi_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"currency_code": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"customer_country_code": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"end_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_classic_subscription": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_government_subscription": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"payment_model": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"program_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"promotion": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"amount": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"currency_unit": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"duration": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"duration_unit": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_intent_to_pay": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_expired": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_started": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"purchase_entitlement_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"region_assignment": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"skus": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"end_date": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"gsi_order_line_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_additional_instance": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_base_service_component": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"license_part_description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"metric_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"quantity": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"sku": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"start_date": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"start_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subscription_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subscription_tier": {
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
		},
	}
}

func readSingularTenantmanagercontrolplaneSubscription(d *schema.ResourceData, m interface{}) error {
	sync := &TenantmanagercontrolplaneSubscriptionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OrganizationsSubscriptionClient()

	return tfresource.ReadResource(sync)
}

type TenantmanagercontrolplaneSubscriptionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_tenantmanagercontrolplane.SubscriptionClient
	Res    *oci_tenantmanagercontrolplane.GetSubscriptionResponse
}

func (s *TenantmanagercontrolplaneSubscriptionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *TenantmanagercontrolplaneSubscriptionDataSourceCrud) Get() error {
	request := oci_tenantmanagercontrolplane.GetSubscriptionRequest{}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "tenantmanagercontrolplane")

	response, err := s.Client.GetSubscription(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *TenantmanagercontrolplaneSubscriptionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.GetId())
	switch v := (s.Res.Subscription).(type) {
	case oci_tenantmanagercontrolplane.ClassicSubscription:
		s.D.Set("entity_version", "V1")

		if v.ClassicSubscriptionId != nil {
			s.D.Set("classic_subscription_id", *v.ClassicSubscriptionId)
		}

		if v.CloudAmountCurrency != nil {
			s.D.Set("cloud_amount_currency", *v.CloudAmountCurrency)
		}

		if v.CsiNumber != nil {
			s.D.Set("csi_number", *v.CsiNumber)
		}

		if v.CustomerCountryCode != nil {
			s.D.Set("customer_country_code", *v.CustomerCountryCode)
		}

		if v.EndDate != nil {
			s.D.Set("end_date", v.EndDate.Format(time.RFC3339Nano))
		}

		if v.IsClassicSubscription != nil {
			s.D.Set("is_classic_subscription", *v.IsClassicSubscription)
		}

		if v.IsGovernmentSubscription != nil {
			s.D.Set("is_government_subscription", *v.IsGovernmentSubscription)
		}

		if v.PaymentModel != nil {
			s.D.Set("payment_model", *v.PaymentModel)
		}

		if v.ProgramType != nil {
			s.D.Set("program_type", *v.ProgramType)
		}

		promotion := []interface{}{}
		for _, item := range v.Promotion {
			promotion = append(promotion, PromotionToMap(item))
		}
		s.D.Set("promotion", promotion)

		if v.PurchaseEntitlementId != nil {
			s.D.Set("purchase_entitlement_id", *v.PurchaseEntitlementId)
		}

		if v.RegionAssignment != nil {
			s.D.Set("region_assignment", *v.RegionAssignment)
		}

		skus := []interface{}{}
		for _, item := range v.Skus {
			skus = append(skus, SubscriptionSkuToMap(item))
		}
		s.D.Set("skus", skus)

		if v.StartDate != nil {
			s.D.Set("start_date", v.StartDate.Format(time.RFC3339Nano))
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubscriptionTier != nil {
			s.D.Set("subscription_tier", *v.SubscriptionTier)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.ServiceName != nil {
			s.D.Set("service_name", *v.ServiceName)
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_tenantmanagercontrolplane.CloudSubscription:
		s.D.Set("entity_version", "V2")

		if v.CurrencyCode != nil {
			s.D.Set("currency_code", *v.CurrencyCode)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubscriptionNumber != nil {
			s.D.Set("subscription_number", *v.SubscriptionNumber)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.ServiceName != nil {
			s.D.Set("service_name", *v.ServiceName)
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'entity_version' of unknown type %v", s.Res.Subscription)
		return nil
	}

	return nil
}
