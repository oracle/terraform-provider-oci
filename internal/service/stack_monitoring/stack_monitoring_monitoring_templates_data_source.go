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

func StackMonitoringMonitoringTemplatesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readStackMonitoringMonitoringTemplates,
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
			"metric_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"monitoring_template_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"namespace": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"resource_types": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"monitoring_template_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(StackMonitoringMonitoringTemplateResource()),
						},
					},
				},
			},
		},
	}
}

func readStackMonitoringMonitoringTemplates(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoringTemplatesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

type StackMonitoringMonitoringTemplatesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_stack_monitoring.StackMonitoringClient
	Res    *oci_stack_monitoring.ListMonitoringTemplatesResponse
}

func (s *StackMonitoringMonitoringTemplatesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *StackMonitoringMonitoringTemplatesDataSourceCrud) Get() error {
	request := oci_stack_monitoring.ListMonitoringTemplatesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if metricName, ok := s.D.GetOkExists("metric_name"); ok {
		interfaces := metricName.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("metric_name") {
			request.MetricName = tmp
		}
	}

	if monitoringTemplateId, ok := s.D.GetOkExists("id"); ok {
		tmp := monitoringTemplateId.(string)
		request.MonitoringTemplateId = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		interfaces := namespace.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("namespace") {
			request.Namespace = tmp
		}
	}

	if resourceTypes, ok := s.D.GetOkExists("resource_types"); ok {
		interfaces := resourceTypes.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("resource_types") {
			request.ResourceTypes = tmp
		}
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_stack_monitoring.ListMonitoringTemplatesLifecycleStateEnum(state.(string))
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_stack_monitoring.ListMonitoringTemplatesStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "stack_monitoring")

	response, err := s.Client.ListMonitoringTemplates(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMonitoringTemplates(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *StackMonitoringMonitoringTemplatesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("StackMonitoringMonitoringTemplatesDataSource-", StackMonitoringMonitoringTemplatesDataSource(), s.D))
	resources := []map[string]interface{}{}
	monitoringTemplate := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MonitoringTemplateSummaryToMap(item))
	}
	monitoringTemplate["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, StackMonitoringMonitoringTemplatesDataSource().Schema["monitoring_template_collection"].Elem.(*schema.Resource).Schema)
		monitoringTemplate["items"] = items
	}

	resources = append(resources, monitoringTemplate)
	if err := s.D.Set("monitoring_template_collection", resources); err != nil {
		return err
	}

	return nil
}
