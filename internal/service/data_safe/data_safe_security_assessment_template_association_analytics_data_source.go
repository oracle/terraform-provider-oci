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

func DataSafeSecurityAssessmentTemplateAssociationAnalyticsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeSecurityAssessmentTemplateAssociationAnalytics,
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
			"template_association_analytics_collection": {
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
											},
										},
									},
									"metric_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"security_assessment_template_association_analytic_count": {
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

func readDataSafeSecurityAssessmentTemplateAssociationAnalytics(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityAssessmentTemplateAssociationAnalyticsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSecurityAssessmentTemplateAssociationAnalyticsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListTemplateAssociationAnalyticsResponse
}

func (s *DataSafeSecurityAssessmentTemplateAssociationAnalyticsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSecurityAssessmentTemplateAssociationAnalyticsDataSourceCrud) Get() error {
	request := oci_data_safe.ListTemplateAssociationAnalyticsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListTemplateAssociationAnalyticsAccessLevelEnum(accessLevel.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
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

	response, err := s.Client.ListTemplateAssociationAnalytics(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListTemplateAssociationAnalytics(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSecurityAssessmentTemplateAssociationAnalyticsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSecurityAssessmentTemplateAssociationAnalyticsDataSource-", DataSafeSecurityAssessmentTemplateAssociationAnalyticsDataSource(), s.D))
	resources := []map[string]interface{}{}
	securityAssessmentTemplateAssociationAnalytic := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, TemplateAssociationAnalyticsSummaryToMap(item))
	}
	securityAssessmentTemplateAssociationAnalytic["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSecurityAssessmentTemplateAssociationAnalyticsDataSource().Schema["template_association_analytics_collection"].Elem.(*schema.Resource).Schema)
		securityAssessmentTemplateAssociationAnalytic["items"] = items
	}

	resources = append(resources, securityAssessmentTemplateAssociationAnalytic)
	if err := s.D.Set("template_association_analytics_collection", resources); err != nil {
		return err
	}

	return nil
}

func TemplateAssociationAnalyticsDimensionsToMap(obj *oci_data_safe.TemplateAssociationAnalyticsDimensions) map[string]interface{} {
	result := map[string]interface{}{}

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

	return result
}

func TemplateAssociationAnalyticsSummaryToMap(obj oci_data_safe.TemplateAssociationAnalyticsSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Dimensions != nil {
		result["dimensions"] = []interface{}{TemplateAssociationAnalyticsDimensionsToMap(obj.Dimensions)}
	}

	result["metric_name"] = string(obj.MetricName)

	if obj.Count != nil {
		result["security_assessment_template_association_analytic_count"] = strconv.FormatInt(*obj.Count, 10)
	}

	return result
}
