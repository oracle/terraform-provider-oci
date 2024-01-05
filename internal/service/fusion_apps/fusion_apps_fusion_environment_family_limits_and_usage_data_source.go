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

func FusionAppsFusionEnvironmentFamilyLimitsAndUsageDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularFusionAppsFusionEnvironmentFamilyLimitsAndUsage,
		Schema: map[string]*schema.Schema{
			"fusion_environment_family_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"development_limit_and_usage": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"limit": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"usage": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"production_limit_and_usage": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"limit": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"usage": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"test_limit_and_usage": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"limit": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"usage": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularFusionAppsFusionEnvironmentFamilyLimitsAndUsage(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentFamilyLimitsAndUsageDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.ReadResource(sync)
}

type FusionAppsFusionEnvironmentFamilyLimitsAndUsageDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fusion_apps.FusionApplicationsClient
	Res    *oci_fusion_apps.GetFusionEnvironmentFamilyLimitsAndUsageResponse
}

func (s *FusionAppsFusionEnvironmentFamilyLimitsAndUsageDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FusionAppsFusionEnvironmentFamilyLimitsAndUsageDataSourceCrud) Get() error {
	request := oci_fusion_apps.GetFusionEnvironmentFamilyLimitsAndUsageRequest{}

	if fusionEnvironmentFamilyId, ok := s.D.GetOkExists("fusion_environment_family_id"); ok {
		tmp := fusionEnvironmentFamilyId.(string)
		request.FusionEnvironmentFamilyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fusion_apps")

	response, err := s.Client.GetFusionEnvironmentFamilyLimitsAndUsage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FusionAppsFusionEnvironmentFamilyLimitsAndUsageDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FusionAppsFusionEnvironmentFamilyLimitsAndUsageDataSource-", FusionAppsFusionEnvironmentFamilyLimitsAndUsageDataSource(), s.D))

	if s.Res.DevelopmentLimitAndUsage != nil {
		s.D.Set("development_limit_and_usage", []interface{}{LimitAndUsageToMap(s.Res.DevelopmentLimitAndUsage)})
	} else {
		s.D.Set("development_limit_and_usage", nil)
	}

	if s.Res.ProductionLimitAndUsage != nil {
		s.D.Set("production_limit_and_usage", []interface{}{LimitAndUsageToMap(s.Res.ProductionLimitAndUsage)})
	} else {
		s.D.Set("production_limit_and_usage", nil)
	}

	if s.Res.TestLimitAndUsage != nil {
		s.D.Set("test_limit_and_usage", []interface{}{LimitAndUsageToMap(s.Res.TestLimitAndUsage)})
	} else {
		s.D.Set("test_limit_and_usage", nil)
	}

	return nil
}

func LimitAndUsageToMap(obj *oci_fusion_apps.LimitAndUsage) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Limit != nil {
		result["limit"] = int(*obj.Limit)
	}

	if obj.Usage != nil {
		result["usage"] = int(*obj.Usage)
	}

	return result
}
