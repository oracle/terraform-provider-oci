// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/database"
)

func init() {
	RegisterResource("oci_database_autonomous_data_warehouse_backup", DatabaseAutonomousDataWarehouseBackupResource())
}

func DatabaseAutonomousDataWarehouseBackupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createDatabaseAutonomousDataWarehouseBackup,
		Read:     readDatabaseAutonomousDataWarehouseBackup,
		Delete:   deleteDatabaseAutonomousDataWarehouseBackup,
		Schema: map[string]*schema.Schema{
			// Required
			"autonomous_data_warehouse_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_automatic": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_ended": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_started": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseAutonomousDataWarehouseBackup(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDataWarehouseBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return CreateResource(d, sync)
}

func readDatabaseAutonomousDataWarehouseBackup(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDataWarehouseBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return ReadResource(sync)
}

func deleteDatabaseAutonomousDataWarehouseBackup(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseAutonomousDataWarehouseBackupResourceCrud struct {
	BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.AutonomousDataWarehouseBackup
	DisableNotFoundRetries bool
}

func (s *DatabaseAutonomousDataWarehouseBackupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseAutonomousDataWarehouseBackupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.AutonomousDataWarehouseBackupLifecycleStateCreating),
	}
}

func (s *DatabaseAutonomousDataWarehouseBackupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.AutonomousDataWarehouseBackupLifecycleStateActive),
	}
}

func (s *DatabaseAutonomousDataWarehouseBackupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.AutonomousDataWarehouseBackupLifecycleStateDeleting),
	}
}

func (s *DatabaseAutonomousDataWarehouseBackupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.AutonomousDataWarehouseBackupLifecycleStateDeleted),
	}
}

func (s *DatabaseAutonomousDataWarehouseBackupResourceCrud) Create() error {
	request := oci_database.CreateAutonomousDataWarehouseBackupRequest{}

	if autonomousDataWarehouseId, ok := s.D.GetOkExists("autonomous_data_warehouse_id"); ok {
		tmp := autonomousDataWarehouseId.(string)
		request.AutonomousDataWarehouseId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateAutonomousDataWarehouseBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousDataWarehouseBackup
	return nil
}

func (s *DatabaseAutonomousDataWarehouseBackupResourceCrud) Get() error {
	request := oci_database.GetAutonomousDataWarehouseBackupRequest{}

	tmp := s.D.Id()
	request.AutonomousDataWarehouseBackupId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetAutonomousDataWarehouseBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousDataWarehouseBackup
	return nil
}

func (s *DatabaseAutonomousDataWarehouseBackupResourceCrud) SetData() error {
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
