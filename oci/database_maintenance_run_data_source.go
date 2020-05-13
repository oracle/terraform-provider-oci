// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"
)

func init() {
	RegisterDatasource("oci_database_maintenance_run", DatabaseMaintenanceRunDataSource())
}

func DatabaseMaintenanceRunDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["maintenance_run_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(DatabaseMaintenanceRunResource(), fieldMap, readSingularDatabaseMaintenanceRun)
}

func readSingularDatabaseMaintenanceRun(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMaintenanceRunDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return ReadResource(sync)
}

type DatabaseMaintenanceRunDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetMaintenanceRunResponse
}

func (s *DatabaseMaintenanceRunDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseMaintenanceRunDataSourceCrud) Get() error {
	request := oci_database.GetMaintenanceRunRequest{}

	if maintenanceRunId, ok := s.D.GetOkExists("maintenance_run_id"); ok {
		tmp := maintenanceRunId.(string)
		request.MaintenanceRunId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.GetMaintenanceRun(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseMaintenanceRunDataSourceCrud) SetData() error {
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

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("maintenance_subtype", s.Res.MaintenanceSubtype)

	s.D.Set("maintenance_type", s.Res.MaintenanceType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TargetResourceId != nil {
		s.D.Set("target_resource_id", *s.Res.TargetResourceId)
	}

	s.D.Set("target_resource_type", s.Res.TargetResourceType)

	if s.Res.TimeEnded != nil {
		s.D.Set("time_ended", s.Res.TimeEnded.String())
	}

	if s.Res.TimeScheduled != nil {
		s.D.Set("time_scheduled", s.Res.TimeScheduled.Format(time.RFC3339Nano))
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	return nil
}
