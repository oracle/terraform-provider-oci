// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v58/database"
)

func DatabaseMaintenanceRunDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["maintenance_run_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseMaintenanceRunResource(), fieldMap, readSingularDatabaseMaintenanceRun)
}

func readSingularDatabaseMaintenanceRun(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMaintenanceRunDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

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

	if s.Res.PatchFailureCount != nil {
		s.D.Set("patch_failure_count", *s.Res.PatchFailureCount)
	}

	if s.Res.PatchId != nil {
		s.D.Set("patch_id", *s.Res.PatchId)
	}

	s.D.Set("patching_mode", s.Res.PatchingMode)

	if s.Res.PeerMaintenanceRunId != nil {
		s.D.Set("peer_maintenance_run_id", *s.Res.PeerMaintenanceRunId)
	}

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
