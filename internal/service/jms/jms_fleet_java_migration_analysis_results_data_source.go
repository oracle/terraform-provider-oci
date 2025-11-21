// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_jms "github.com/oracle/oci-go-sdk/v65/jms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsFleetJavaMigrationAnalysisResultsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readJmsFleetJavaMigrationAnalysisResults,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"application_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fleet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"host_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_instance_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_end": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_start": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"java_migration_analysis_result_collection": {
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
									"application_execution_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"application_key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"application_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"application_path": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"bucket": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"fleet_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"host_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"managed_instance_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"metadata": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"object_list": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"object_storage_upload_dir_path": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"source_jdk_version": {
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

func readJmsFleetJavaMigrationAnalysisResults(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetJavaMigrationAnalysisResultsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsFleetJavaMigrationAnalysisResultsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.ListJavaMigrationAnalysisResultsResponse
}

func (s *JmsFleetJavaMigrationAnalysisResultsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsFleetJavaMigrationAnalysisResultsDataSourceCrud) Get() error {
	request := oci_jms.ListJavaMigrationAnalysisResultsRequest{}

	if applicationName, ok := s.D.GetOkExists("application_name"); ok {
		tmp := applicationName.(string)
		request.ApplicationName = &tmp
	}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if hostName, ok := s.D.GetOkExists("host_name"); ok {
		tmp := hostName.(string)
		request.HostName = &tmp
	}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	if timeEnd, ok := s.D.GetOkExists("time_end"); ok {
		tmp, err := time.Parse(time.RFC3339, timeEnd.(string))
		if err != nil {
			return err
		}
		request.TimeEnd = &oci_common.SDKTime{Time: tmp}
	}

	if timeStart, ok := s.D.GetOkExists("time_start"); ok {
		tmp, err := time.Parse(time.RFC3339, timeStart.(string))
		if err != nil {
			return err
		}
		request.TimeStart = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms")

	response, err := s.Client.ListJavaMigrationAnalysisResults(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListJavaMigrationAnalysisResults(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *JmsFleetJavaMigrationAnalysisResultsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsFleetJavaMigrationAnalysisResultsDataSource-", JmsFleetJavaMigrationAnalysisResultsDataSource(), s.D))
	resources := []map[string]interface{}{}
	fleetJavaMigrationAnalysisResult := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, JavaMigrationAnalysisResultSummaryToMap(item))
	}
	fleetJavaMigrationAnalysisResult["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, JmsFleetJavaMigrationAnalysisResultsDataSource().Schema["java_migration_analysis_result_collection"].Elem.(*schema.Resource).Schema)
		fleetJavaMigrationAnalysisResult["items"] = items
	}

	resources = append(resources, fleetJavaMigrationAnalysisResult)
	if err := s.D.Set("java_migration_analysis_result_collection", resources); err != nil {
		return err
	}

	return nil
}

func JavaMigrationAnalysisResultSummaryToMap(obj oci_jms.JavaMigrationAnalysisResultSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["application_execution_type"] = string(obj.ApplicationExecutionType)

	if obj.ApplicationKey != nil {
		result["application_key"] = string(*obj.ApplicationKey)
	}

	if obj.ApplicationName != nil {
		result["application_name"] = string(*obj.ApplicationName)
	}

	if obj.ApplicationPath != nil {
		result["application_path"] = string(*obj.ApplicationPath)
	}

	if obj.BucketName != nil {
		result["bucket"] = string(*obj.BucketName)
	}

	if obj.FleetId != nil {
		result["fleet_id"] = string(*obj.FleetId)
	}

	if obj.HostName != nil {
		result["host_name"] = string(*obj.HostName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.ManagedInstanceId != nil {
		result["managed_instance_id"] = string(*obj.ManagedInstanceId)
	}

	if obj.Metadata != nil {
		result["metadata"] = string(*obj.Metadata)
	}

	if obj.Namespace != nil {
		result["namespace"] = string(*obj.Namespace)
	}

	result["object_list"] = obj.ObjectList

	if obj.ObjectStorageUploadDirPath != nil {
		result["object_storage_upload_dir_path"] = string(*obj.ObjectStorageUploadDirPath)
	}

	if obj.SourceJdkVersion != nil {
		result["source_jdk_version"] = string(*obj.SourceJdkVersion)
	}

	if obj.TargetJdkVersion != nil {
		result["target_jdk_version"] = string(*obj.TargetJdkVersion)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.WorkRequestId != nil {
		result["work_request_id"] = string(*obj.WorkRequestId)
	}

	return result
}
