// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package metering_computation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_metering_computation "github.com/oracle/oci-go-sdk/v65/usageapi"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MeteringComputationScheduledRunsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMeteringComputationScheduledRuns,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"schedule_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"scheduled_run_collection": {
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
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"schedule_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
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
								},
							},
						},
					},
				},
			},
		},
	}
}

func readMeteringComputationScheduledRuns(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationScheduledRunsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()

	return tfresource.ReadResource(sync)
}

type MeteringComputationScheduledRunsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_metering_computation.UsageapiClient
	Res    *oci_metering_computation.ListScheduledRunsResponse
}

func (s *MeteringComputationScheduledRunsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MeteringComputationScheduledRunsDataSourceCrud) Get() error {
	request := oci_metering_computation.ListScheduledRunsRequest{}

	if scheduleId, ok := s.D.GetOkExists("schedule_id"); ok {
		tmp := scheduleId.(string)
		request.ScheduleId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "metering_computation")

	response, err := s.Client.ListScheduledRuns(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListScheduledRuns(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MeteringComputationScheduledRunsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MeteringComputationScheduledRunsDataSource-", MeteringComputationScheduledRunsDataSource(), s.D))
	resources := []map[string]interface{}{}
	scheduledRun := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ScheduledRunSummaryToMap(item))
	}
	scheduledRun["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, MeteringComputationScheduledRunsDataSource().Schema["scheduled_run_collection"].Elem.(*schema.Resource).Schema)
		scheduledRun["items"] = items
	}

	resources = append(resources, scheduledRun)
	if err := s.D.Set("scheduled_run_collection", resources); err != nil {
		return err
	}

	return nil
}

func ScheduledRunSummaryToMap(obj oci_metering_computation.ScheduledRunSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.ScheduleId != nil {
		result["schedule_id"] = string(*obj.ScheduleId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeFinished != nil {
		result["time_finished"] = obj.TimeFinished.String()
	}

	return result
}
