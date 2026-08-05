// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseAutonomousContainerDatabaseBackupDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseAutonomousContainerDatabaseBackup,
		Schema: map[string]*schema.Schema{
			"autonomous_container_database_backup_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"acd_display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"autonomous_container_database_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"autonomous_databases": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"backup_destination_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"backup_retention_policy_on_terminate": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"dbrs_policy_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"internet_proxy": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_remote": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_retention_lock_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_zero_data_loss_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"remote_region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vpc_password": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vpc_user": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"infrastructure_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_automatic": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_remote_backup": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"retention_period_in_days": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
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

func readSingularDatabaseAutonomousContainerDatabaseBackup(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerDatabaseBackupDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousContainerDatabaseBackupDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetAutonomousContainerDatabaseBackupResponse
}

func (s *DatabaseAutonomousContainerDatabaseBackupDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousContainerDatabaseBackupDataSourceCrud) Get() error {
	request := oci_database.GetAutonomousContainerDatabaseBackupRequest{}

	if autonomousContainerDatabaseBackupId, ok := s.D.GetOkExists("autonomous_container_database_backup_id"); ok {
		tmp := autonomousContainerDatabaseBackupId.(string)
		request.AutonomousContainerDatabaseBackupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetAutonomousContainerDatabaseBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseAutonomousContainerDatabaseBackupDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AcdDisplayName != nil {
		s.D.Set("acd_display_name", *s.Res.AcdDisplayName)
	}

	if s.Res.AutonomousContainerDatabaseId != nil {
		s.D.Set("autonomous_container_database_id", *s.Res.AutonomousContainerDatabaseId)
	}

	autonomousDatabases := []interface{}{}
	for _, item := range s.Res.AutonomousDatabases {
		autonomousDatabases = append(autonomousDatabases, AutonomousDatabaseInBackupToMap(item))
	}
	s.D.Set("autonomous_databases", autonomousDatabases)

	if s.Res.BackupDestinationDetails != nil {
		s.D.Set("backup_destination_details", []interface{}{BackupDestinationDetailsToMap(*s.Res.BackupDestinationDetails)})
	} else {
		s.D.Set("backup_destination_details", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DbVersion != nil {
		s.D.Set("db_version", *s.Res.DbVersion)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("infrastructure_type", s.Res.InfrastructureType)

	if s.Res.IsAutomatic != nil {
		s.D.Set("is_automatic", *s.Res.IsAutomatic)
	}

	if s.Res.IsRemoteBackup != nil {
		s.D.Set("is_remote_backup", *s.Res.IsRemoteBackup)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.RetentionPeriodInDays != nil {
		s.D.Set("retention_period_in_days", *s.Res.RetentionPeriodInDays)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeEnded != nil {
		s.D.Set("time_ended", s.Res.TimeEnded.Format(time.RFC3339Nano))
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.Format(time.RFC3339Nano))
	}

	s.D.Set("type", s.Res.Type)

	return nil
}
