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

func TenantmanagercontrolplaneAssignedSubscriptionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readTenantmanagercontrolplaneAssignedSubscriptions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"entity_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"assigned_subscription_collection": {
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
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_classic_subscription": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_government_subscription": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"managed_by": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"order_ids": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
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
							},
						},
					},
				},
			},
		},
	}
}

func readTenantmanagercontrolplaneAssignedSubscriptions(d *schema.ResourceData, m interface{}) error {
	sync := &TenantmanagercontrolplaneAssignedSubscriptionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OrganizationsSubscriptionClient()

	return tfresource.ReadResource(sync)
}

type TenantmanagercontrolplaneAssignedSubscriptionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_tenantmanagercontrolplane.SubscriptionClient
	Res    *oci_tenantmanagercontrolplane.ListAssignedSubscriptionsResponse
}

func (s *TenantmanagercontrolplaneAssignedSubscriptionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *TenantmanagercontrolplaneAssignedSubscriptionsDataSourceCrud) Get() error {
	request := oci_tenantmanagercontrolplane.ListAssignedSubscriptionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if entityVersion, ok := s.D.GetOkExists("entity_version"); ok {
		request.EntityVersion = oci_tenantmanagercontrolplane.ListAssignedSubscriptionsEntityVersionEnum(entityVersion.(string))
	}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "tenantmanagercontrolplane")

	response, err := s.Client.ListAssignedSubscriptions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAssignedSubscriptions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *TenantmanagercontrolplaneAssignedSubscriptionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("TenantmanagercontrolplaneAssignedSubscriptionsDataSource-", TenantmanagercontrolplaneAssignedSubscriptionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	assignedSubscription := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AssignedSubscriptionSummaryToMap(item))
	}
	assignedSubscription["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, TenantmanagercontrolplaneAssignedSubscriptionsDataSource().Schema["assigned_subscription_collection"].Elem.(*schema.Resource).Schema)
		assignedSubscription["items"] = items
	}

	resources = append(resources, assignedSubscription)
	if err := s.D.Set("assigned_subscription_collection", resources); err != nil {
		return err
	}

	return nil
}

func AssignedSubscriptionSummaryToMap(obj oci_tenantmanagercontrolplane.AssignedSubscriptionSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_tenantmanagercontrolplane.ClassicAssignedSubscriptionSummary:
		result["entity_version"] = "V1"

		if v.ClassicSubscriptionId != nil {
			result["classic_subscription_id"] = string(*v.ClassicSubscriptionId)
		}

		if v.CsiNumber != nil {
			result["csi_number"] = string(*v.CsiNumber)
		}

		if v.EndDate != nil {
			result["end_date"] = v.EndDate.Format(time.RFC3339Nano)
		}

		if v.IsClassicSubscription != nil {
			result["is_classic_subscription"] = bool(*v.IsClassicSubscription)
		}

		result["managed_by"] = string(v.ManagedBy)

		if v.RegionAssignment != nil {
			result["region_assignment"] = string(*v.RegionAssignment)
		}

		if v.StartDate != nil {
			result["start_date"] = v.StartDate.Format(time.RFC3339Nano)
		}

		result["state"] = string(v.LifecycleState)
	case oci_tenantmanagercontrolplane.CloudAssignedSubscriptionSummary:
		result["entity_version"] = "V2"

		if v.CurrencyCode != nil {
			result["currency_code"] = string(*v.CurrencyCode)
		}

		result["state"] = string(v.LifecycleState)

		if v.SubscriptionNumber != nil {
			result["subscription_number"] = string(*v.SubscriptionNumber)
		}
	default:
		log.Printf("[WARN] Received 'entity_version' of unknown type %v", obj)
		return nil
	}

	return result
}

func AssignedSubscriptionPromotionToMap(obj oci_tenantmanagercontrolplane.Promotion) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Amount != nil {
		result["amount"] = float32(*obj.Amount)
	}

	if obj.CurrencyUnit != nil {
		result["currency_unit"] = string(*obj.CurrencyUnit)
	}

	if obj.Duration != nil {
		result["duration"] = int(*obj.Duration)
	}

	if obj.DurationUnit != nil {
		result["duration_unit"] = string(*obj.DurationUnit)
	}

	if obj.IsIntentToPay != nil {
		result["is_intent_to_pay"] = bool(*obj.IsIntentToPay)
	}

	result["status"] = string(obj.Status)

	if obj.TimeExpired != nil {
		result["time_expired"] = obj.TimeExpired.Format(time.RFC3339Nano)
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.Format(time.RFC3339Nano)
	}

	return result
}

func AssignedSubscriptionSkuToMap(obj oci_tenantmanagercontrolplane.SubscriptionSku) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.EndDate != nil {
		result["end_date"] = obj.EndDate.Format(time.RFC3339Nano)
	}

	if obj.GsiOrderLineId != nil {
		result["gsi_order_line_id"] = string(*obj.GsiOrderLineId)
	}

	if obj.IsAdditionalInstance != nil {
		result["is_additional_instance"] = bool(*obj.IsAdditionalInstance)
	}

	if obj.IsBaseServiceComponent != nil {
		result["is_base_service_component"] = bool(*obj.IsBaseServiceComponent)
	}

	if obj.LicensePartDescription != nil {
		result["license_part_description"] = string(*obj.LicensePartDescription)
	}

	if obj.MetricName != nil {
		result["metric_name"] = string(*obj.MetricName)
	}

	if obj.Quantity != nil {
		result["quantity"] = int(*obj.Quantity)
	}

	if obj.Sku != nil {
		result["sku"] = string(*obj.Sku)
	}

	if obj.StartDate != nil {
		result["start_date"] = obj.StartDate.Format(time.RFC3339Nano)
	}

	return result
}
