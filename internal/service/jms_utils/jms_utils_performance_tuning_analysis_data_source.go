// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms_utils

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_jms_utils "github.com/oracle/oci-go-sdk/v65/jmsutils"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsUtilsPerformanceTuningAnalysisDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readJmsUtilsPerformanceTuningAnalysis,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"analysis_project_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"performance_tuning_analysis_result": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"performance_tuning_analysis_collection": {
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
									"analysis_project_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"artifact_object_storage_path": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"created_by": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"id": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"result": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"result_object_storage_path": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_finished": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_started": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"warning_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"work_request_id": {
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

func readJmsUtilsPerformanceTuningAnalysis(d *schema.ResourceData, m interface{}) error {
	sync := &JmsUtilsPerformanceTuningAnalysisDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JmsUtilsClient()

	return tfresource.ReadResource(sync)
}

type JmsUtilsPerformanceTuningAnalysisDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms_utils.JmsUtilsClient
	Res    *oci_jms_utils.ListPerformanceTuningAnalysisResponse
}

func (s *JmsUtilsPerformanceTuningAnalysisDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsUtilsPerformanceTuningAnalysisDataSourceCrud) Get() error {
	request := oci_jms_utils.ListPerformanceTuningAnalysisRequest{}

	if analysisProjectName, ok := s.D.GetOkExists("analysis_project_name"); ok {
		tmp := analysisProjectName.(string)
		request.AnalysisProjectName = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if performanceTuningAnalysisResult, ok := s.D.GetOkExists("performance_tuning_analysis_result"); ok {
		request.PerformanceTuningAnalysisResult = oci_jms_utils.PerformanceTuningAnalysisResultEnum(performanceTuningAnalysisResult.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms_utils")

	response, err := s.Client.ListPerformanceTuningAnalysis(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPerformanceTuningAnalysis(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *JmsUtilsPerformanceTuningAnalysisDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsUtilsPerformanceTuningAnalysisDataSource-", JmsUtilsPerformanceTuningAnalysisDataSource(), s.D))
	resources := []map[string]interface{}{}
	performanceTuningAnalysi := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, PerformanceTuningAnalysisSummaryToMap(item))
	}
	performanceTuningAnalysi["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, JmsUtilsPerformanceTuningAnalysisDataSource().Schema["performance_tuning_analysis_collection"].Elem.(*schema.Resource).Schema)
		performanceTuningAnalysi["items"] = items
	}

	resources = append(resources, performanceTuningAnalysi)
	if err := s.D.Set("performance_tuning_analysis_collection", resources); err != nil {
		return err
	}

	return nil
}

func PerformanceTuningAnalysisSummaryToMap(obj oci_jms_utils.PerformanceTuningAnalysisSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AnalysisProjectName != nil {
		result["analysis_project_name"] = string(*obj.AnalysisProjectName)
	}

	if obj.ArtifactObjectStoragePath != nil {
		result["artifact_object_storage_path"] = string(*obj.ArtifactObjectStoragePath)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.CreatedBy != nil {
		result["created_by"] = []interface{}{PrincipalToMap(obj.CreatedBy)}
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["result"] = string(obj.Result)

	if obj.ResultObjectStoragePath != nil {
		result["result_object_storage_path"] = string(*obj.ResultObjectStoragePath)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeFinished != nil {
		result["time_finished"] = obj.TimeFinished.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	if obj.WarningCount != nil {
		result["warning_count"] = int(*obj.WarningCount)
	}

	if obj.WorkRequestId != nil {
		result["work_request_id"] = string(*obj.WorkRequestId)
	}

	return result
}

func PrincipalToMap(obj *oci_jms_utils.Principal) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}
