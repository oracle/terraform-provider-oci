// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"
)

func LogAnalyticsNamespaceScheduledTasksDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLogAnalyticsNamespaceScheduledTasks,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"target_service": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"task_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"template_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scheduled_task_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(LogAnalyticsNamespaceScheduledTaskResource()),
						},
					},
				},
			},
		},
	}
}

func readLogAnalyticsNamespaceScheduledTasks(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceScheduledTasksDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

type LogAnalyticsNamespaceScheduledTasksDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.ListScheduledTasksResponse
}

func (s *LogAnalyticsNamespaceScheduledTasksDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsNamespaceScheduledTasksDataSourceCrud) Get() error {
	request := oci_log_analytics.ListScheduledTasksRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if targetService, ok := s.D.GetOkExists("target_service"); ok {
		tmp := targetService.(string)
		request.TargetService = &tmp
	}

	if taskType, ok := s.D.GetOkExists("task_type"); ok {
		request.TaskType = oci_log_analytics.ListScheduledTasksTaskTypeEnum(taskType.(string))
	}

	if templateId, ok := s.D.GetOkExists("template_id"); ok {
		tmp := templateId.(string)
		request.TemplateId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "log_analytics")

	response, err := s.Client.ListScheduledTasks(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListScheduledTasks(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *LogAnalyticsNamespaceScheduledTasksDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LogAnalyticsNamespaceScheduledTasksDataSource-", LogAnalyticsNamespaceScheduledTasksDataSource(), s.D))
	resources := []map[string]interface{}{}
	namespaceScheduledTask := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ScheduledTaskSummaryToMap(item))
	}
	namespaceScheduledTask["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, LogAnalyticsNamespaceScheduledTasksDataSource().Schema["scheduled_task_collection"].Elem.(*schema.Resource).Schema)
		namespaceScheduledTask["items"] = items
	}

	resources = append(resources, namespaceScheduledTask)
	if err := s.D.Set("scheduled_task_collection", resources); err != nil {
		return err
	}

	return nil
}
