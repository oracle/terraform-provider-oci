// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_jms "github.com/oracle/oci-go-sdk/v65/jms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsTaskSchedulesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readJmsTaskSchedules,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"fleet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_instance_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"task_schedule_name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"task_schedule_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(JmsTaskScheduleResource()),
						},
					},
				},
			},
		},
	}
}

func readJmsTaskSchedules(d *schema.ResourceData, m interface{}) error {
	sync := &JmsTaskSchedulesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsTaskSchedulesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.ListTaskSchedulesResponse
}

func (s *JmsTaskSchedulesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsTaskSchedulesDataSourceCrud) Get() error {
	request := oci_jms.ListTaskSchedulesRequest{}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if taskScheduleNameContains, ok := s.D.GetOkExists("task_schedule_name_contains"); ok {
		tmp := taskScheduleNameContains.(string)
		request.TaskScheduleNameContains = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms")

	response, err := s.Client.ListTaskSchedules(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListTaskSchedules(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *JmsTaskSchedulesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsTaskSchedulesDataSource-", JmsTaskSchedulesDataSource(), s.D))
	resources := []map[string]interface{}{}
	taskSchedule := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, TaskScheduleSummaryToMap(item))
	}
	taskSchedule["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, JmsTaskSchedulesDataSource().Schema["task_schedule_collection"].Elem.(*schema.Resource).Schema)
		taskSchedule["items"] = items
	}

	resources = append(resources, taskSchedule)
	if err := s.D.Set("task_schedule_collection", resources); err != nil {
		return err
	}

	return nil
}
