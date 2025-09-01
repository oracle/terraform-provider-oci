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

func JmsUtilsJavaMigrationAnalysisDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readJmsUtilsJavaMigrationAnalysis,
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
			"java_migration_analysis_collection": {
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
									"analysis_result_files": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"analysis_result_object_storage_path": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"bucket": {
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
									"input_applications_object_storage_paths": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"metadata": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_jdk_version": {
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

func readJmsUtilsJavaMigrationAnalysis(d *schema.ResourceData, m interface{}) error {
	sync := &JmsUtilsJavaMigrationAnalysisDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JmsUtilsClient()

	return tfresource.ReadResource(sync)
}

type JmsUtilsJavaMigrationAnalysisDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms_utils.JmsUtilsClient
	Res    *oci_jms_utils.ListJavaMigrationAnalysisResponse
}

func (s *JmsUtilsJavaMigrationAnalysisDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsUtilsJavaMigrationAnalysisDataSourceCrud) Get() error {
	request := oci_jms_utils.ListJavaMigrationAnalysisRequest{}

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms_utils")

	response, err := s.Client.ListJavaMigrationAnalysis(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListJavaMigrationAnalysis(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *JmsUtilsJavaMigrationAnalysisDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsUtilsJavaMigrationAnalysisDataSource-", JmsUtilsJavaMigrationAnalysisDataSource(), s.D))
	resources := []map[string]interface{}{}
	javaMigrationAnalysi := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, JavaMigrationAnalysisSummaryToMap(item))
	}
	javaMigrationAnalysi["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, JmsUtilsJavaMigrationAnalysisDataSource().Schema["java_migration_analysis_collection"].Elem.(*schema.Resource).Schema)
		javaMigrationAnalysi["items"] = items
	}

	resources = append(resources, javaMigrationAnalysi)
	if err := s.D.Set("java_migration_analysis_collection", resources); err != nil {
		return err
	}

	return nil
}

func JavaMigrationAnalysisSummaryToMap(obj oci_jms_utils.JavaMigrationAnalysisSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AnalysisProjectName != nil {
		result["analysis_project_name"] = string(*obj.AnalysisProjectName)
	}

	result["analysis_result_files"] = obj.AnalysisResultFiles

	if obj.AnalysisResultObjectStoragePath != nil {
		result["analysis_result_object_storage_path"] = string(*obj.AnalysisResultObjectStoragePath)
	}

	if obj.BucketName != nil {
		result["bucket"] = string(*obj.BucketName)
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

	result["input_applications_object_storage_paths"] = obj.InputApplicationsObjectStoragePaths

	if obj.Metadata != nil {
		result["metadata"] = string(*obj.Metadata)
	}

	if obj.NamespaceName != nil {
		result["namespace"] = string(*obj.NamespaceName)
	}

	if obj.TargetJdkVersion != nil {
		result["target_jdk_version"] = string(*obj.TargetJdkVersion)
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

	if obj.WorkRequestId != nil {
		result["work_request_id"] = string(*obj.WorkRequestId)
	}

	return result
}
