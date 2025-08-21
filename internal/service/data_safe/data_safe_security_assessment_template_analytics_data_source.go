// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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

func DataSafeSecurityAssessmentTemplateAnalyticsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeSecurityAssessmentTemplateAnalytics,
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
			"is_compared": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_compliant": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_group": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"target_database_group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"template_assessment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"template_baseline_assessment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"template_analytics_collection": {
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
												"is_compared": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"is_compliant": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"is_group": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"target_database_group_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"target_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"template_assessment_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"template_baseline_assessment_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"time_last_compared": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"total_checks": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"total_checks_failed": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"total_non_compliant_targets": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"total_targets": {
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"metric_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"security_assessment_template_analytic_count": {
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

func readDataSafeSecurityAssessmentTemplateAnalytics(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityAssessmentTemplateAnalyticsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSecurityAssessmentTemplateAnalyticsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListTemplateAnalyticsResponse
}

func (s *DataSafeSecurityAssessmentTemplateAnalyticsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSecurityAssessmentTemplateAnalyticsDataSourceCrud) Get() error {
	request := oci_data_safe.ListTemplateAnalyticsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListTemplateAnalyticsAccessLevelEnum(accessLevel.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if isCompared, ok := s.D.GetOkExists("is_compared"); ok {
		tmp := isCompared.(bool)
		request.IsCompared = &tmp
	}

	if isCompliant, ok := s.D.GetOkExists("is_compliant"); ok {
		tmp := isCompliant.(bool)
		request.IsCompliant = &tmp
	}

	if isGroup, ok := s.D.GetOkExists("is_group"); ok {
		tmp := isGroup.(bool)
		request.IsGroup = &tmp
	}

	if targetDatabaseGroupId, ok := s.D.GetOkExists("target_database_group_id"); ok {
		tmp := targetDatabaseGroupId.(string)
		request.TargetDatabaseGroupId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	if templateAssessmentId, ok := s.D.GetOkExists("template_assessment_id"); ok {
		tmp := templateAssessmentId.(string)
		request.TemplateAssessmentId = &tmp
	}

	if templateBaselineAssessmentId, ok := s.D.GetOkExists("template_baseline_assessment_id"); ok {
		tmp := templateBaselineAssessmentId.(string)
		request.TemplateBaselineAssessmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListTemplateAnalytics(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListTemplateAnalytics(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSecurityAssessmentTemplateAnalyticsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSecurityAssessmentTemplateAnalyticsDataSource-", DataSafeSecurityAssessmentTemplateAnalyticsDataSource(), s.D))
	resources := []map[string]interface{}{}
	securityAssessmentTemplateAnalytic := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, TemplateAnalyticsSummaryToMap(item))
	}
	securityAssessmentTemplateAnalytic["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSecurityAssessmentTemplateAnalyticsDataSource().Schema["template_analytics_collection"].Elem.(*schema.Resource).Schema)
		securityAssessmentTemplateAnalytic["items"] = items
	}

	resources = append(resources, securityAssessmentTemplateAnalytic)
	if err := s.D.Set("template_analytics_collection", resources); err != nil {
		return err
	}

	return nil
}

func TemplateAnalyticsDimensionsToMap(obj *oci_data_safe.TemplateAnalyticsDimensions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsCompared != nil {
		result["is_compared"] = bool(*obj.IsCompared)
	}

	if obj.IsCompliant != nil {
		result["is_compliant"] = bool(*obj.IsCompliant)
	}

	if obj.IsGroup != nil {
		result["is_group"] = bool(*obj.IsGroup)
	}

	if obj.TargetDatabaseGroupId != nil {
		result["target_database_group_id"] = string(*obj.TargetDatabaseGroupId)
	}

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	if obj.TemplateAssessmentId != nil {
		result["template_assessment_id"] = string(*obj.TemplateAssessmentId)
	}

	if obj.TemplateBaselineAssessmentId != nil {
		result["template_baseline_assessment_id"] = string(*obj.TemplateBaselineAssessmentId)
	}

	if obj.TimeLastCompared != nil {
		result["time_last_compared"] = obj.TimeLastCompared.String()
	}

	if obj.TotalChecks != nil {
		result["total_checks"] = int(*obj.TotalChecks)
	}

	if obj.TotalChecksFailed != nil {
		result["total_checks_failed"] = int(*obj.TotalChecksFailed)
	}

	if obj.TotalNonCompliantTargets != nil {
		result["total_non_compliant_targets"] = int(*obj.TotalNonCompliantTargets)
	}

	if obj.TotalTargets != nil {
		result["total_targets"] = int(*obj.TotalTargets)
	}

	return result
}

func TemplateAnalyticsSummaryToMap(obj oci_data_safe.TemplateAnalyticsSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Dimensions != nil {
		result["dimensions"] = []interface{}{TemplateAnalyticsDimensionsToMap(obj.Dimensions)}
	}

	result["metric_name"] = string(obj.MetricName)

	if obj.Count != nil {
		result["security_assessment_template_analytic_count"] = strconv.FormatInt(*obj.Count, 10)
	}

	return result
}
