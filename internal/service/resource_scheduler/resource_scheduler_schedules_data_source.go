// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resource_scheduler

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_resource_scheduler "github.com/oracle/oci-go-sdk/v65/resourcescheduler"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ResourceSchedulerSchedulesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readResourceSchedulerSchedules,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"schedule_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"schedule_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(ResourceSchedulerScheduleResource()),
						},
					},
				},
			},
		},
	}
}

func readResourceSchedulerSchedules(d *schema.ResourceData, m interface{}) error {
	sync := &ResourceSchedulerSchedulesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ScheduleClient()

	return tfresource.ReadResource(sync)
}

type ResourceSchedulerSchedulesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_resource_scheduler.ScheduleClient
	Res    *oci_resource_scheduler.ListSchedulesResponse
}

func (s *ResourceSchedulerSchedulesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ResourceSchedulerSchedulesDataSourceCrud) Get() error {
	request := oci_resource_scheduler.ListSchedulesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if scheduleId, ok := s.D.GetOkExists("id"); ok {
		tmp := scheduleId.(string)
		request.ScheduleId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_resource_scheduler.ScheduleLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "resource_scheduler")
	response, err := s.Client.ListSchedules(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSchedules(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ResourceSchedulerSchedulesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ResourceSchedulerSchedulesDataSource-", ResourceSchedulerSchedulesDataSource(), s.D))
	resources := []map[string]interface{}{}
	schedule := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ScheduleSummaryToMap(item))
	}
	schedule["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ResourceSchedulerSchedulesDataSource().Schema["schedule_collection"].Elem.(*schema.Resource).Schema)
		schedule["items"] = items
	}

	resources = append(resources, schedule)
	if err := s.D.Set("schedule_collection", resources); err != nil {
		return err
	}

	return nil
}
