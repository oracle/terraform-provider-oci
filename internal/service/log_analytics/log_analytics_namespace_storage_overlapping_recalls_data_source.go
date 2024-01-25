// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LogAnalyticsNamespaceStorageOverlappingRecallsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLogAnalyticsNamespaceStorageOverlappingRecalls,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"time_data_ended": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_data_started": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"overlapping_recall_collection": {
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
									"collection_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"created_by": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"log_sets": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"purpose": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"query_string": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"recall_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_data_ended": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_data_started": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_started": {
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

func readLogAnalyticsNamespaceStorageOverlappingRecalls(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceStorageOverlappingRecallsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

type LogAnalyticsNamespaceStorageOverlappingRecallsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.ListOverlappingRecallsResponse
}

func (s *LogAnalyticsNamespaceStorageOverlappingRecallsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsNamespaceStorageOverlappingRecallsDataSourceCrud) Get() error {
	request := oci_log_analytics.ListOverlappingRecallsRequest{}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if timeDataEnded, ok := s.D.GetOkExists("time_data_ended"); ok {
		tmp, err := time.Parse(time.RFC3339, timeDataEnded.(string))
		if err != nil {
			return err
		}
		request.TimeDataEnded = &oci_common.SDKTime{Time: tmp}
	}

	if timeDataStarted, ok := s.D.GetOkExists("time_data_started"); ok {
		tmp, err := time.Parse(time.RFC3339, timeDataStarted.(string))
		if err != nil {
			return err
		}
		request.TimeDataStarted = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "log_analytics")

	response, err := s.Client.ListOverlappingRecalls(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOverlappingRecalls(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *LogAnalyticsNamespaceStorageOverlappingRecallsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LogAnalyticsNamespaceStorageOverlappingRecallsDataSource-", LogAnalyticsNamespaceStorageOverlappingRecallsDataSource(), s.D))
	resources := []map[string]interface{}{}
	namespaceStorageOverlappingRecall := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OverlappingRecallSummaryToMap(item))
	}
	namespaceStorageOverlappingRecall["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, LogAnalyticsNamespaceStorageOverlappingRecallsDataSource().Schema["overlapping_recall_collection"].Elem.(*schema.Resource).Schema)
		namespaceStorageOverlappingRecall["items"] = items
	}

	resources = append(resources, namespaceStorageOverlappingRecall)
	if err := s.D.Set("overlapping_recall_collection", resources); err != nil {
		return err
	}

	return nil
}

func OverlappingRecallSummaryToMap(obj oci_log_analytics.OverlappingRecallSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CollectionId != nil {
		result["collection_id"] = strconv.FormatInt(*obj.CollectionId, 10)
	}

	if obj.CreatedBy != nil {
		result["created_by"] = string(*obj.CreatedBy)
	}

	if obj.LogSets != nil {
		result["log_sets"] = string(*obj.LogSets)
	}

	if obj.Purpose != nil {
		result["purpose"] = string(*obj.Purpose)
	}

	if obj.QueryString != nil {
		result["query_string"] = string(*obj.QueryString)
	}

	if obj.RecallId != nil {
		result["recall_id"] = strconv.FormatInt(*obj.RecallId, 10)
	}

	result["status"] = string(obj.Status)

	if obj.TimeDataEnded != nil {
		result["time_data_ended"] = obj.TimeDataEnded.String()
	}

	if obj.TimeDataStarted != nil {
		result["time_data_started"] = obj.TimeDataStarted.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	return result
}
