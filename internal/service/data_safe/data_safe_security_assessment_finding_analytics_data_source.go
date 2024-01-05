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

func DataSafeSecurityAssessmentFindingAnalyticsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeSecurityAssessmentFindingAnalytics,
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
			"finding_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_top_finding": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"severity": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"top_finding_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"finding_analytics_collection": {
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
												"key": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"severity": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"target_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"title": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"top_finding_category": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"top_finding_status": {
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
									"security_assessment_finding_analytic_count": {
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

func readDataSafeSecurityAssessmentFindingAnalytics(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityAssessmentFindingAnalyticsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSecurityAssessmentFindingAnalyticsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListFindingAnalyticsResponse
}

func (s *DataSafeSecurityAssessmentFindingAnalyticsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSecurityAssessmentFindingAnalyticsDataSourceCrud) Get() error {
	request := oci_data_safe.ListFindingAnalyticsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListFindingAnalyticsAccessLevelEnum(accessLevel.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if findingKey, ok := s.D.GetOkExists("finding_key"); ok {
		tmp := findingKey.(string)
		request.FindingKey = &tmp
	}

	if groupBy, ok := s.D.GetOkExists("group_by"); ok {
		request.GroupBy = oci_data_safe.ListFindingAnalyticsGroupByEnum(groupBy.(string))
	}

	if isTopFinding, ok := s.D.GetOkExists("is_top_finding"); ok {
		tmp := isTopFinding.(bool)
		request.IsTopFinding = &tmp
	}

	if severity, ok := s.D.GetOkExists("severity"); ok {
		request.Severity = oci_data_safe.ListFindingAnalyticsSeverityEnum(severity.(string))
	}

	if topFindingStatus, ok := s.D.GetOkExists("top_finding_status"); ok {
		request.TopFindingStatus = oci_data_safe.FindingAnalyticsDimensionsTopFindingStatusEnum(topFindingStatus.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListFindingAnalytics(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListFindingAnalytics(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSecurityAssessmentFindingAnalyticsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSecurityAssessmentFindingAnalyticsDataSource-", DataSafeSecurityAssessmentFindingAnalyticsDataSource(), s.D))
	resources := []map[string]interface{}{}
	securityAssessmentFindingAnalytic := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, FindingAnalyticsSummaryToMap(item))
	}
	securityAssessmentFindingAnalytic["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSecurityAssessmentFindingAnalyticsDataSource().Schema["finding_analytics_collection"].Elem.(*schema.Resource).Schema)
		securityAssessmentFindingAnalytic["items"] = items
	}

	resources = append(resources, securityAssessmentFindingAnalytic)
	if err := s.D.Set("finding_analytics_collection", resources); err != nil {
		return err
	}

	return nil
}

func FindingAnalyticsDimensionsToMap(obj *oci_data_safe.FindingAnalyticsDimensions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	result["severity"] = string(obj.Severity)

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	if obj.Title != nil {
		result["title"] = string(*obj.Title)
	}

	if obj.TopFindingCategory != nil {
		result["top_finding_category"] = string(*obj.TopFindingCategory)
	}

	result["top_finding_status"] = string(obj.TopFindingStatus)

	return result
}

func FindingAnalyticsSummaryToMap(obj oci_data_safe.FindingAnalyticsSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Dimensions != nil {
		result["dimensions"] = []interface{}{FindingAnalyticsDimensionsToMap(obj.Dimensions)}
	}

	result["metric_name"] = string(obj.MetricName)

	if obj.Count != nil {
		result["security_assessment_finding_analytic_count"] = strconv.FormatInt(*obj.Count, 10)
	}

	return result
}
