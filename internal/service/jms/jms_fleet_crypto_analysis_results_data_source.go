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

func JmsFleetCryptoAnalysisResultsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readJmsFleetCryptoAnalysisResults,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"aggregation_mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"finding_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"finding_count_greater_than": {
				Type:     schema.TypeInt,
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
			"non_compliant_finding_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"non_compliant_finding_count_greater_than": {
				Type:     schema.TypeInt,
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
			"crypto_analysis_result_collection": {
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
									"aggregation_mode": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"bucket": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"crypto_roadmap_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"finding_count": {
										Type:     schema.TypeInt,
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
									"non_compliant_finding_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"object": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"summarized_event_count": {
										Type:     schema.TypeInt,
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
									"time_first_event": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_last_event": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_started": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"total_event_count": {
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

func readJmsFleetCryptoAnalysisResults(d *schema.ResourceData, m interface{}) error {
	sync := &JmsFleetCryptoAnalysisResultsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsFleetCryptoAnalysisResultsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.ListCryptoAnalysisResultsResponse
}

func (s *JmsFleetCryptoAnalysisResultsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsFleetCryptoAnalysisResultsDataSourceCrud) Get() error {
	request := oci_jms.ListCryptoAnalysisResultsRequest{}

	if aggregationMode, ok := s.D.GetOkExists("aggregation_mode"); ok {
		request.AggregationMode = oci_jms.ListCryptoAnalysisResultsAggregationModeEnum(aggregationMode.(string))
	}

	if findingCount, ok := s.D.GetOkExists("finding_count"); ok {
		tmp := findingCount.(int)
		request.FindingCount = &tmp
	}

	if findingCountGreaterThan, ok := s.D.GetOkExists("finding_count_greater_than"); ok {
		tmp := findingCountGreaterThan.(int)
		request.FindingCountGreaterThan = &tmp
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

	if nonCompliantFindingCount, ok := s.D.GetOkExists("non_compliant_finding_count"); ok {
		tmp := nonCompliantFindingCount.(int)
		request.NonCompliantFindingCount = &tmp
	}

	if nonCompliantFindingCountGreaterThan, ok := s.D.GetOkExists("non_compliant_finding_count_greater_than"); ok {
		tmp := nonCompliantFindingCountGreaterThan.(int)
		request.NonCompliantFindingCountGreaterThan = &tmp
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

	response, err := s.Client.ListCryptoAnalysisResults(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCryptoAnalysisResults(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *JmsFleetCryptoAnalysisResultsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsFleetCryptoAnalysisResultsDataSource-", JmsFleetCryptoAnalysisResultsDataSource(), s.D))
	resources := []map[string]interface{}{}
	fleetCryptoAnalysisResult := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CryptoAnalysisResultSummaryToMap(item))
	}
	fleetCryptoAnalysisResult["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, JmsFleetCryptoAnalysisResultsDataSource().Schema["crypto_analysis_result_collection"].Elem.(*schema.Resource).Schema)
		fleetCryptoAnalysisResult["items"] = items
	}

	resources = append(resources, fleetCryptoAnalysisResult)
	if err := s.D.Set("crypto_analysis_result_collection", resources); err != nil {
		return err
	}

	return nil
}

func CryptoAnalysisResultSummaryToMap(obj oci_jms.CryptoAnalysisResultSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["aggregation_mode"] = string(obj.AggregationMode)

	if obj.BucketName != nil {
		result["bucket"] = string(*obj.BucketName)
	}

	if obj.CryptoRoadmapVersion != nil {
		result["crypto_roadmap_version"] = string(*obj.CryptoRoadmapVersion)
	}

	if obj.FindingCount != nil {
		result["finding_count"] = int(*obj.FindingCount)
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

	if obj.NonCompliantFindingCount != nil {
		result["non_compliant_finding_count"] = int(*obj.NonCompliantFindingCount)
	}

	if obj.ObjectName != nil {
		result["object"] = string(*obj.ObjectName)
	}

	if obj.SummarizedEventCount != nil {
		result["summarized_event_count"] = int(*obj.SummarizedEventCount)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeFinished != nil {
		result["time_finished"] = obj.TimeFinished.String()
	}

	if obj.TimeFirstEvent != nil {
		result["time_first_event"] = obj.TimeFirstEvent.String()
	}

	if obj.TimeLastEvent != nil {
		result["time_last_event"] = obj.TimeLastEvent.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	if obj.TotalEventCount != nil {
		result["total_event_count"] = int(*obj.TotalEventCount)
	}

	if obj.WorkRequestId != nil {
		result["work_request_id"] = string(*obj.WorkRequestId)
	}

	return result
}
