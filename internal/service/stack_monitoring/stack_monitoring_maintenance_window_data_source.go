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

func StackMonitoringMaintenanceWindowDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["maintenance_window_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(StackMonitoringMaintenanceWindowResource(), fieldMap, readSingularStackMonitoringMaintenanceWindow)
}

func readSingularStackMonitoringMaintenanceWindow(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMaintenanceWindowDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

type StackMonitoringMaintenanceWindowDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_stack_monitoring.StackMonitoringClient
	Res    *oci_stack_monitoring.GetMaintenanceWindowResponse
}

func (s *StackMonitoringMaintenanceWindowDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *StackMonitoringMaintenanceWindowDataSourceCrud) Get() error {
	request := oci_stack_monitoring.GetMaintenanceWindowRequest{}

	if maintenanceWindowId, ok := s.D.GetOkExists("maintenance_window_id"); ok {
		tmp := maintenanceWindowId.(string)
		request.MaintenanceWindowId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "stack_monitoring")

	response, err := s.Client.GetMaintenanceWindow(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *StackMonitoringMaintenanceWindowDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("lifecycle_details", s.Res.LifecycleDetails)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	resources := []interface{}{}
	for _, item := range s.Res.Resources {
		resources = append(resources, CreateMaintenanceWindowResourceDetailsToMap(item))
	}
	s.D.Set("resources", resources)

	resourcesDetails := []interface{}{}
	for _, item := range s.Res.ResourcesDetails {
		resourcesDetails = append(resourcesDetails, MonitoredResourceDetailsToMap(item))
	}
	s.D.Set("resources_details", resourcesDetails)

	if s.Res.Schedule != nil {
		scheduleArray := []interface{}{}
		if scheduleMap := MaintenanceWindowScheduleToMap(&s.Res.Schedule); scheduleMap != nil {
			scheduleArray = append(scheduleArray, scheduleMap)
		}
		s.D.Set("schedule", scheduleArray)
	} else {
		s.D.Set("schedule", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
