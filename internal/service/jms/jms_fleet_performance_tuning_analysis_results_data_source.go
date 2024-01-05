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

func JmsFleetPerformanceTuningAnalysisResultsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readJmsFleetPerformanceTuningAnalysisResults,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"application_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fleet_id": {
				Type:     schema.TypeString,
				Required: true,
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
			"performance_tuning_analysis_result_collection": {
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
									"application_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"application_installation_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"application_installation_path": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"application_name": {
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
									"namespace": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"object": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"result": {
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

func readJmsFleetPerformanceTuningAnalysisResults(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetPerformanceTuningAnalysisResultsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsFleetPerformanceTuningAnalysisResultsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.ListPerformanceTuningAnalysisResultsResponse
}

func (s *JmsFleetPerformanceTuningAnalysisResultsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsFleetPerformanceTuningAnalysisResultsDataSourceCrud) Get() error {
	request := oci_jms.ListPerformanceTuningAnalysisResultsRequest{}

	if applicationId, ok := s.D.GetOkExists("application_id"); ok {
		tmp := applicationId.(string)
		request.ApplicationId = &tmp
	}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
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

	response, err := s.Client.ListPerformanceTuningAnalysisResults(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPerformanceTuningAnalysisResults(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *JmsFleetPerformanceTuningAnalysisResultsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsFleetPerformanceTuningAnalysisResultsDataSource-", JmsFleetPerformanceTuningAnalysisResultsDataSource(), s.D))
	resources := []map[string]interface{}{}
	fleetPerformanceTuningAnalysisResult := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, PerformanceTuningAnalysisResultSummaryToMap(item))
	}
	fleetPerformanceTuningAnalysisResult["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, JmsFleetPerformanceTuningAnalysisResultsDataSource().Schema["performance_tuning_analysis_result_collection"].Elem.(*schema.Resource).Schema)
		fleetPerformanceTuningAnalysisResult["items"] = items
	}

	resources = append(resources, fleetPerformanceTuningAnalysisResult)
	if err := s.D.Set("performance_tuning_analysis_result_collection", resources); err != nil {
		return err
	}

	return nil
}

func PerformanceTuningAnalysisResultSummaryToMap(obj oci_jms.PerformanceTuningAnalysisResultSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApplicationId != nil {
		result["application_id"] = string(*obj.ApplicationId)
	}

	if obj.ApplicationInstallationId != nil {
		result["application_installation_id"] = string(*obj.ApplicationInstallationId)
	}

	if obj.ApplicationInstallationPath != nil {
		result["application_installation_path"] = string(*obj.ApplicationInstallationPath)
	}

	if obj.ApplicationName != nil {
		result["application_name"] = string(*obj.ApplicationName)
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

	if obj.Namespace != nil {
		result["namespace"] = string(*obj.Namespace)
	}

	if obj.ObjectName != nil {
		result["object"] = string(*obj.ObjectName)
	}

	result["result"] = string(obj.Result)

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
