// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseBackupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("1h"),
			Update: tfresource.GetTimeoutDuration("1h"),
			Delete: tfresource.GetTimeoutDuration("1h"),
		},
		Create: createDatabaseBackup,
		Read:   readDatabaseBackup,
		Update: updateDatabaseBackup,
		Delete: deleteDatabaseBackup,
		Schema: map[string]*schema.Schema{
			// Required
			"database_id": {
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
			"retention_period_in_days": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"retention_period_in_years": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// Computed
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"backup_destination_type": {
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
			"is_using_oracle_managed_keys": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"encryption_key_location_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"hsm_password": {
							Type:      schema.TypeString,
							Computed:  true,
							Sensitive: true,
						},
						"provider_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"key_store_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key_store_wallet_name": {
				Type:     schema.TypeString,
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
			"secondary_kms_key_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"time_expiry_scheduled": {
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

func createDatabaseBackup(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseBackup(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseBackup(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseBackup(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type DatabaseBackupResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.Backup
	WorkRequestClient      *oci_work_requests.WorkRequestClient
	DisableNotFoundRetries bool
}

func (s *DatabaseBackupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseBackupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.BackupLifecycleStateCreating),
		string(oci_database.BackupLifecycleStateRestoring),
	}
}

func (s *DatabaseBackupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.BackupLifecycleStateActive),
	}
}

func (s *DatabaseBackupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.BackupLifecycleStateDeleting),
	}
}

func (s *DatabaseBackupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.BackupLifecycleStateDeleted),
	}
}

func (s *DatabaseBackupResourceCrud) Create() error {
	request := oci_database.CreateBackupRequest{}

	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if retentionPeriodInDays, ok := s.D.GetOkExists("retention_period_in_days"); ok {
		tmp := retentionPeriodInDays.(int)
		request.RetentionPeriodInDays = &tmp
	}

	if retentionPeriodInYears, ok := s.D.GetOkExists("retention_period_in_years"); ok {
		tmp := retentionPeriodInYears.(int)
		request.RetentionPeriodInYears = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateBackup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "backup", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	s.Res = &response.Backup
	return nil
}

func (s *DatabaseBackupResourceCrud) Get() error {
	request := oci_database.GetBackupRequest{}

	tmp := s.D.Id()
	request.BackupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Backup
	return nil
}

func (s *DatabaseBackupResourceCrud) Update() error {
	request := oci_database.UpdateBackupRequest{}

	tmp := s.D.Id()
	request.BackupId = &tmp

	if retentionPeriodInDays, ok := s.D.GetOkExists("retention_period_in_days"); ok {
		tmp := retentionPeriodInDays.(int)
		request.RetentionPeriodInDays = &tmp
	}

	if retentionPeriodInYears, ok := s.D.GetOkExists("retention_period_in_years"); ok {
		tmp := retentionPeriodInYears.(int)
		request.RetentionPeriodInYears = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateBackup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "backup", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return s.Get()
}

func (s *DatabaseBackupResourceCrud) Delete() error {
	request := oci_database.DeleteBackupRequest{}

	tmp := s.D.Id()
	request.BackupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.DeleteBackup(context.Background(), request)
	return err
}

func (s *DatabaseBackupResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	s.D.Set("backup_destination_type", s.Res.BackupDestinationType)

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

	if s.Res.IsUsingOracleManagedKeys != nil {
		s.D.Set("is_using_oracle_managed_keys", *s.Res.IsUsingOracleManagedKeys)
	}

	if s.Res.EncryptionKeyLocationDetails != nil {
		s.D.Set("encryption_key_location_details", []interface{}{EncryptionKeyLocationDetailsToMap(&s.Res.EncryptionKeyLocationDetails, "")})
	}

	if s.Res.KeyStoreId != nil {
		s.D.Set("key_store_id", *s.Res.KeyStoreId)
	}

	if s.Res.KeyStoreWalletName != nil {
		s.D.Set("key_store_wallet_name", *s.Res.KeyStoreWalletName)
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

	if s.Res.RetentionPeriodInDays != nil {
		s.D.Set("retention_period_in_days", *s.Res.RetentionPeriodInDays)
	}

	if s.Res.RetentionPeriodInYears != nil {
		s.D.Set("retention_period_in_years", *s.Res.RetentionPeriodInYears)
	}

	s.D.Set("secondary_kms_key_ids", s.Res.SecondaryKmsKeyIds)

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeEnded != nil {
		s.D.Set("time_ended", s.Res.TimeEnded.Format(time.RFC3339Nano))
	}

	if s.Res.TimeExpiryScheduled != nil {
		s.D.Set("time_expiry_scheduled", s.Res.TimeExpiryScheduled.String())
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
