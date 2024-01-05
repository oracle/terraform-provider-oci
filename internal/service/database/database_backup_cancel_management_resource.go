// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseBackupCancelManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseBackupCancelManagement,
		Read:     readDatabaseBackupCancelManagement,
		Delete:   deleteDatabaseBackupCancelManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"backup_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"cancel_backup_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},

			// Optional

			// Computed
		},
	}
}

func createDatabaseBackupCancelManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseBackupCancelManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseBackupCancelManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDatabaseBackupCancelManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseBackupCancelManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.Backup
	DisableNotFoundRetries bool
}

func (s *DatabaseBackupCancelManagementResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("DatabaseBackupCancelManagementResource-", DatabaseBackupCancelManagementResource(), s.D)
}

func (s *DatabaseBackupCancelManagementResourceCrud) Create() error {
	request := oci_database.CancelBackupRequest{}

	if backupId, ok := s.D.GetOkExists("backup_id"); ok {
		err := s.getBackupInfo(backupId.(string))
		if err != nil {
			return fmt.Errorf("[ERROR] Could not get Backup for the : %v", err)
		}

		backupType := string(s.Res.Type)

		if !s.IsAutomaticBackup(backupType) {
			return fmt.Errorf("[ERROR] Cancel backup only supported on automatic backups")
		}

		tmp := backupId.(string)
		request.BackupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")
	_, err := s.Client.CancelBackup(context.Background(), request)

	if err != nil {
		log.Printf("Error occurred during cancel backup: %s", err)
	}
	trigger := s.D.Get("cancel_backup_trigger")
	s.D.Set("cancel_backup_trigger", trigger)
	val := s.D.Get("backup_id")
	s.D.Set("backup_id", val)

	return nil
}

func (s *DatabaseBackupCancelManagementResourceCrud) SetData() error {
	return nil
}

func (s *DatabaseBackupCancelManagementResourceCrud) getBackupInfo(backupId string) error {
	request := oci_database.GetBackupRequest{}

	request.BackupId = &backupId

	response, err := s.Client.GetBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Backup

	return nil
}

func (s *DatabaseBackupCancelManagementResourceCrud) IsAutomaticBackup(backupType string) bool {
	// Only allowing cancelation of automatic backups
	return strings.EqualFold(backupType, "INCREMENTAL")
}
