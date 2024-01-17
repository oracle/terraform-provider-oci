// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSecurityAssessmentSecurityFeatureAnalyticsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeSecurityAssessmentSecurityFeatureAnalytics,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_feature_analytics_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"dimensions": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"security_feature": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"metric_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"security_assessment_security_feature_analytic_count": {
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

func readDataSafeSecurityAssessmentSecurityFeatureAnalytics(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityAssessmentSecurityFeatureAnalyticsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSecurityAssessmentSecurityFeatureAnalyticsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListSecurityFeatureAnalyticsResponse
}

func (s *DataSafeSecurityAssessmentSecurityFeatureAnalyticsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSecurityAssessmentSecurityFeatureAnalyticsDataSourceCrud) Get() error {
	request := oci_data_safe.ListSecurityFeatureAnalyticsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListSecurityFeatureAnalyticsAccessLevelEnum(accessLevel.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListSecurityFeatureAnalytics(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeSecurityAssessmentSecurityFeatureAnalyticsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSecurityAssessmentSecurityFeatureAnalyticsDataSource-", DataSafeSecurityAssessmentSecurityFeatureAnalyticsDataSource(), s.D))
	resources := []map[string]interface{}{}
	securityAssessmentSecurityFeatureAnalytic := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SecurityFeatureAnalyticsSummaryToMap(item))
	}
	securityAssessmentSecurityFeatureAnalytic["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSecurityAssessmentSecurityFeatureAnalyticsDataSource().Schema["security_feature_analytics_collection"].Elem.(*schema.Resource).Schema)
		securityAssessmentSecurityFeatureAnalytic["items"] = items
	}

	resources = append(resources, securityAssessmentSecurityFeatureAnalytic)
	if err := s.D.Set("security_feature_analytics_collection", resources); err != nil {
		return err
	}

	return nil
}

func SecurityFeatureAnalyticsDimensionsToMap(obj *oci_data_safe.SecurityFeatureAnalyticsDimensions) map[string]interface{} {
	result := map[string]interface{}{}

	result["security_feature"] = string(obj.SecurityFeature)

	return result
}

func SecurityFeatureAnalyticsSummaryToMap(obj oci_data_safe.SecurityFeatureAnalyticsSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Dimensions != nil {
		result["dimensions"] = []interface{}{SecurityFeatureAnalyticsDimensionsToMap(obj.Dimensions)}
	}

	result["metric_name"] = string(obj.MetricName)

	if obj.Count != nil {
		result["security_assessment_security_feature_analytic_count"] = strconv.FormatInt(*obj.Count, 10)
	}

	return result
}
