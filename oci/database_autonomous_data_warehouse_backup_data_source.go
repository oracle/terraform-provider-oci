// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"
)

func init() {
	RegisterDatasource("oci_database_autonomous_data_warehouse_backup", DatabaseAutonomousDataWarehouseBackupDataSource())
}

func DatabaseAutonomousDataWarehouseBackupDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["autonomous_data_warehouse_backup_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(DatabaseAutonomousDataWarehouseBackupResource(), fieldMap, readSingularDatabaseAutonomousDataWarehouseBackup)
}

func readSingularDatabaseAutonomousDataWarehouseBackup(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDataWarehouseBackupDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return ReadResource(sync)
}

type DatabaseAutonomousDataWarehouseBackupDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetAutonomousDataWarehouseBackupResponse
}

func (s *DatabaseAutonomousDataWarehouseBackupDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousDataWarehouseBackupDataSourceCrud) Get() error {
	request := oci_database.GetAutonomousDataWarehouseBackupRequest{}

	if autonomousDataWarehouseBackupId, ok := s.D.GetOkExists("autonomous_data_warehouse_backup_id"); ok {
		tmp := autonomousDataWarehouseBackupId.(string)
		request.AutonomousDataWarehouseBackupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.GetAutonomousDataWarehouseBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseAutonomousDataWarehouseBackupDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AutonomousDataWarehouseId != nil {
		s.D.Set("autonomous_data_warehouse_id", *s.Res.AutonomousDataWarehouseId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.IsAutomatic != nil {
		s.D.Set("is_automatic", *s.Res.IsAutomatic)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeEnded != nil {
		s.D.Set("time_ended", s.Res.TimeEnded.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	s.D.Set("type", s.Res.Type)

	return nil
}
