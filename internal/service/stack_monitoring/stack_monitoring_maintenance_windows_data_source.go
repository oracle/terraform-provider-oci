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

func StackMonitoringMaintenanceWindowsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readStackMonitoringMaintenanceWindows,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"maintenance_window_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(StackMonitoringMaintenanceWindowSummaryResponse()),
						},
					},
				},
			},
		},
	}
}

func readStackMonitoringMaintenanceWindows(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMaintenanceWindowsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

type StackMonitoringMaintenanceWindowsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_stack_monitoring.StackMonitoringClient
	Res    *oci_stack_monitoring.ListMaintenanceWindowsResponse
}

func (s *StackMonitoringMaintenanceWindowsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *StackMonitoringMaintenanceWindowsDataSourceCrud) Get() error {
	request := oci_stack_monitoring.ListMaintenanceWindowsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if lifecycleDetails, ok := s.D.GetOkExists("lifecycle_details"); ok {
		request.LifecycleDetails = oci_stack_monitoring.ListMaintenanceWindowsLifecycleDetailsEnum(lifecycleDetails.(string))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_stack_monitoring.ListMaintenanceWindowsStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "stack_monitoring")

	response, err := s.Client.ListMaintenanceWindows(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMaintenanceWindows(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *StackMonitoringMaintenanceWindowsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("StackMonitoringMaintenanceWindowsDataSource-", StackMonitoringMaintenanceWindowsDataSource(), s.D))
	resources := []map[string]interface{}{}
	maintenanceWindow := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MaintenanceWindowSummaryToMap(item))
	}
	maintenanceWindow["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, StackMonitoringMaintenanceWindowsDataSource().Schema["maintenance_window_collection"].Elem.(*schema.Resource).Schema)
		maintenanceWindow["items"] = items
	}

	resources = append(resources, maintenanceWindow)
	if err := s.D.Set("maintenance_window_collection", resources); err != nil {
		return err
	}

	return nil
}
