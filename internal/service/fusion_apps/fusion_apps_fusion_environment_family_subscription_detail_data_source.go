// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fusion_apps

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_fusion_apps "github.com/oracle/oci-go-sdk/v65/fusionapps"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FusionAppsFusionEnvironmentFamilySubscriptionDetailDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularFusionAppsFusionEnvironmentFamilySubscriptionDetail,
		Schema: map[string]*schema.Schema{
			"fusion_environment_family_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"subscriptions": {
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
						"id": {
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
								},
							},
						},
					},
				},
			},
		},
	}
}

func readSingularFusionAppsFusionEnvironmentFamilySubscriptionDetail(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentFamilySubscriptionDetailDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.ReadResource(sync)
}

type FusionAppsFusionEnvironmentFamilySubscriptionDetailDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fusion_apps.FusionApplicationsClient
	Res    *oci_fusion_apps.GetFusionEnvironmentFamilySubscriptionDetailResponse
}

func (s *FusionAppsFusionEnvironmentFamilySubscriptionDetailDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FusionAppsFusionEnvironmentFamilySubscriptionDetailDataSourceCrud) Get() error {
	request := oci_fusion_apps.GetFusionEnvironmentFamilySubscriptionDetailRequest{}

	if fusionEnvironmentFamilyId, ok := s.D.GetOkExists("fusion_environment_family_id"); ok {
		tmp := fusionEnvironmentFamilyId.(string)
		request.FusionEnvironmentFamilyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fusion_apps")

	response, err := s.Client.GetFusionEnvironmentFamilySubscriptionDetail(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FusionAppsFusionEnvironmentFamilySubscriptionDetailDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FusionAppsFusionEnvironmentFamilySubscriptionDetailDataSource-", FusionAppsFusionEnvironmentFamilySubscriptionDetailDataSource(), s.D))

	subscriptions := []interface{}{}
	for _, item := range s.Res.Subscriptions {
		subscriptions = append(subscriptions, SubscriptionToMap(item))
	}
	s.D.Set("subscriptions", subscriptions)

	return nil
}

func SubscriptionToMap(obj oci_fusion_apps.Subscription) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ClassicSubscriptionId != nil {
		result["classic_subscription_id"] = string(*obj.ClassicSubscriptionId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.ServiceName != nil {
		result["service_name"] = string(*obj.ServiceName)
	}

	skus := []interface{}{}
	for _, item := range obj.Skus {
		skus = append(skus, SubscriptionSkuToMap(item))
	}
	result["skus"] = skus

	return result
}

func SubscriptionSkuToMap(obj oci_fusion_apps.SubscriptionSku) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
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

	return result
}
