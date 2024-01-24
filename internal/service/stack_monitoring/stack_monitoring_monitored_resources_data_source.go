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

func StackMonitoringMonitoredResourcesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readStackMonitoringMonitoredResources,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"work_request_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"monitored_resource_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(StackMonitoringMonitoredResourceResource()),
						},
					},
				},
			},
		},
	}
}

func readStackMonitoringMonitoredResources(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoredResourcesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

type StackMonitoringMonitoredResourcesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_stack_monitoring.StackMonitoringClient
	Res    *oci_stack_monitoring.ListMonitoredResourcesResponse
}

func (s *StackMonitoringMonitoredResourcesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *StackMonitoringMonitoredResourcesDataSourceCrud) Get() error {
	request := oci_stack_monitoring.ListMonitoredResourcesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if workRequestId, ok := s.D.GetOkExists("work_request_id"); ok {
		tmp := workRequestId.(string)
		request.WorkRequestId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "stack_monitoring")

	response, err := s.Client.ListMonitoredResources(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMonitoredResources(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *StackMonitoringMonitoredResourcesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("StackMonitoringMonitoredResourcesDataSource-", StackMonitoringMonitoredResourcesDataSource(), s.D))
	resources := []map[string]interface{}{}
	monitoredResource := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MonitoredResourceSummaryToMap(item))
	}
	monitoredResource["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, StackMonitoringMonitoredResourcesDataSource().Schema["monitored_resource_collection"].Elem.(*schema.Resource).Schema)
		monitoredResource["items"] = items
	}

	resources = append(resources, monitoredResource)
	if err := s.D.Set("monitored_resource_collection", resources); err != nil {
		return err
	}

	return nil
}
