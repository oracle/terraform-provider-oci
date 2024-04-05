// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func StackMonitoringDiscoveryJobLogsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readStackMonitoringDiscoveryJobLogs,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"discovery_job_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"log_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"discovery_job_log_collection": {
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
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"log_message": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"log_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"system_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"time_created": {
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

func readStackMonitoringDiscoveryJobLogs(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringDiscoveryJobLogsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

type StackMonitoringDiscoveryJobLogsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_stack_monitoring.StackMonitoringClient
	Res    *oci_stack_monitoring.ListDiscoveryJobLogsResponse
}

func (s *StackMonitoringDiscoveryJobLogsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *StackMonitoringDiscoveryJobLogsDataSourceCrud) Get() error {
	request := oci_stack_monitoring.ListDiscoveryJobLogsRequest{}

	if discoveryJobId, ok := s.D.GetOkExists("discovery_job_id"); ok {
		tmp := discoveryJobId.(string)
		request.DiscoveryJobId = &tmp
	}

	if logType, ok := s.D.GetOkExists("log_type"); ok {
		request.LogType = oci_stack_monitoring.ListDiscoveryJobLogsLogTypeEnum(logType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "stack_monitoring")

	response, err := s.Client.ListDiscoveryJobLogs(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDiscoveryJobLogs(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *StackMonitoringDiscoveryJobLogsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("StackMonitoringDiscoveryJobLogsDataSource-", StackMonitoringDiscoveryJobLogsDataSource(), s.D))
	resources := []map[string]interface{}{}
	discoveryJobLog := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DiscoveryJobLogSummaryToMap(item))
	}
	discoveryJobLog["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, StackMonitoringDiscoveryJobLogsDataSource().Schema["discovery_job_log_collection"].Elem.(*schema.Resource).Schema)
		discoveryJobLog["items"] = items
	}

	resources = append(resources, discoveryJobLog)
	if err := s.D.Set("discovery_job_log_collection", resources); err != nil {
		return err
	}

	return nil
}

func DiscoveryJobLogSummaryToMap(obj oci_stack_monitoring.DiscoveryJobLogSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LogMessage != nil {
		result["log_message"] = string(*obj.LogMessage)
	}

	result["log_type"] = string(obj.LogType)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}
