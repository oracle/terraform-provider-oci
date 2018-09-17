// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"
)

func AutonomousDataWarehouseBackupDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularAutonomousDataWarehouseBackup,
		Schema: map[string]*schema.Schema{
			"autonomous_data_warehouse_backup_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"autonomous_data_warehouse_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
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

func readSingularAutonomousDataWarehouseBackup(d *schema.ResourceData, m interface{}) error {
	sync := &AutonomousDataWarehouseBackupDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return ReadResource(sync)
}

type AutonomousDataWarehouseBackupDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetAutonomousDataWarehouseBackupResponse
}

func (s *AutonomousDataWarehouseBackupDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AutonomousDataWarehouseBackupDataSourceCrud) Get() error {
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

func (s *AutonomousDataWarehouseBackupDataSourceCrud) SetData() error {
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
