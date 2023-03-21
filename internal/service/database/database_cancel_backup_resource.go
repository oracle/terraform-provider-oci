// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseCancelBackupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCancelBackup,
		Read:     readCancelBackup,
		Delete:   deleteCancelBackup,
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

			// Computed
			"database_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_edition": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_size_in_gbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"kms_key_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"kms_key_version_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"shape": {
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
			"vault_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCancelBackup(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCancelBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.CreateResource(d, sync)
}

func readCancelBackup(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteCancelBackup(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseCancelBackupResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.Backup
	DisableNotFoundRetries bool
}

func (s *DatabaseCancelBackupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseCancelBackupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.BackupLifecycleStateCanceling),
	}
}

func (s *DatabaseCancelBackupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.BackupLifecycleStateCanceling),
		string(oci_database.BackupLifecycleStateActive),
		string(oci_database.BackupLifecycleStateFailed),
	}
}

func (s *DatabaseCancelBackupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.BackupLifecycleStateDeleting),
	}
}

func (s *DatabaseCancelBackupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.BackupLifecycleStateDeleted),
	}
}

func (s *DatabaseCancelBackupResourceCrud) Create() error {
	request := oci_database.CancelBackupRequest{}

	if backupId, ok := s.D.GetOkExists("backup_id"); ok {
		err := s.getBackupInfo(backupId.(string))
		if err != nil {
			return fmt.Errorf("[ERROR] Could not get Backup for the : %v", err)
		}

		backupType := string(s.Res.Type)
		fmt.Printf("backupType = %s \n", backupType)
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

func (s *DatabaseCancelBackupResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("database_edition", s.Res.DatabaseEdition)

	if s.Res.DatabaseId != nil {
		s.D.Set("database_id", *s.Res.DatabaseId)
	}

	if s.Res.DatabaseSizeInGBs != nil {
		s.D.Set("database_size_in_gbs", *s.Res.DatabaseSizeInGBs)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.KmsKeyId != nil {
		s.D.Set("kms_key_id", *s.Res.KmsKeyId)
	}

	if s.Res.KmsKeyVersionId != nil {
		s.D.Set("kms_key_version_id", *s.Res.KmsKeyVersionId)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeEnded != nil {
		s.D.Set("time_ended", s.Res.TimeEnded.Format(time.RFC3339Nano))
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.Format(time.RFC3339Nano))
	}

	s.D.Set("type", s.Res.Type)

	if s.Res.VaultId != nil {
		s.D.Set("vault_id", *s.Res.VaultId)
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}

func (s *DatabaseCancelBackupResourceCrud) getBackupInfo(backupId string) error {
	request := oci_database.GetBackupRequest{}

	request.BackupId = &backupId

	response, err := s.Client.GetBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Backup

	return nil
}

func (s *DatabaseCancelBackupResourceCrud) IsAutomaticBackup(backupType string) bool {
	return strings.EqualFold(backupType, "INCREMENTAL")
}
