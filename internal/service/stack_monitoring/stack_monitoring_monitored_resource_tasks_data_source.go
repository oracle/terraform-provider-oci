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

func StackMonitoringMonitoredResourceTasksDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readStackMonitoringMonitoredResourceTasks,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"monitored_resource_tasks_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(StackMonitoringMonitoredResourceTaskResource()),
						},
					},
				},
			},
		},
	}
}

func readStackMonitoringMonitoredResourceTasks(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoredResourceTasksDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

type StackMonitoringMonitoredResourceTasksDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_stack_monitoring.StackMonitoringClient
	Res    *oci_stack_monitoring.ListMonitoredResourceTasksResponse
}

func (s *StackMonitoringMonitoredResourceTasksDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *StackMonitoringMonitoredResourceTasksDataSourceCrud) Get() error {
	request := oci_stack_monitoring.ListMonitoredResourceTasksRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_stack_monitoring.ListMonitoredResourceTasksStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "stack_monitoring")

	response, err := s.Client.ListMonitoredResourceTasks(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMonitoredResourceTasks(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *StackMonitoringMonitoredResourceTasksDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("StackMonitoringMonitoredResourceTasksDataSource-", StackMonitoringMonitoredResourceTasksDataSource(), s.D))
	resources := []map[string]interface{}{}
	monitoredResourceTask := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MonitoredResourceTaskSummaryToMap(item))
	}
	monitoredResourceTask["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, StackMonitoringMonitoredResourceTasksDataSource().Schema["monitored_resource_tasks_collection"].Elem.(*schema.Resource).Schema)
		monitoredResourceTask["items"] = items
	}

	resources = append(resources, monitoredResourceTask)
	if err := s.D.Set("monitored_resource_tasks_collection", resources); err != nil {
		return err
	}

	return nil
}
