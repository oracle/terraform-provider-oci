// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseAutonomousDatabaseBackupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseAutonomousDatabaseBackup,
		Read:     readDatabaseAutonomousDatabaseBackup,
		Update:   updateDatabaseAutonomousDatabaseBackup,
		Delete:   deleteDatabaseAutonomousDatabaseBackup,
		Schema: map[string]*schema.Schema{
			// Required
			"autonomous_database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"is_long_term_backup": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"retention_period_in_days": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// Computed
			"backup_destination_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"internet_proxy": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"vpc_password": {
							Type:      schema.TypeString,
							Optional:  true,
							Computed:  true,
							ForceNew:  true,
							Sensitive: true,
						},
						"vpc_user": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
					},
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_size_in_tbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"db_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_automatic": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_restorable": {
				Type:     schema.TypeBool,
				Computed: true,
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
			"size_in_tbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_available_till": {
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
		},
	}
}

func createDatabaseAutonomousDatabaseBackup(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseAutonomousDatabaseBackup(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.ReadResource(sync)
}

func updateDatabaseAutonomousDatabaseBackup(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseAutonomousDatabaseBackup(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type DatabaseAutonomousDatabaseBackupResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	WorkRequestClient      *oci_work_requests.WorkRequestClient
	Res                    *oci_database.AutonomousDatabaseBackup
	DisableNotFoundRetries bool
}

func (s *DatabaseAutonomousDatabaseBackupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseAutonomousDatabaseBackupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.AutonomousDatabaseBackupLifecycleStateCreating),
	}
}

func (s *DatabaseAutonomousDatabaseBackupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.AutonomousDatabaseBackupLifecycleStateActive),
	}
}

func (s *DatabaseAutonomousDatabaseBackupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.AutonomousDatabaseBackupLifecycleStateDeleting),
	}
}

func (s *DatabaseAutonomousDatabaseBackupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.AutonomousDatabaseBackupLifecycleStateDeleted),
	}
}

func (s *DatabaseAutonomousDatabaseBackupResourceCrud) Create() error {
	request := oci_database.CreateAutonomousDatabaseBackupRequest{}

	if autonomousDatabaseId, ok := s.D.GetOkExists("autonomous_database_id"); ok {
		tmp := autonomousDatabaseId.(string)
		request.AutonomousDatabaseId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if isLongTermBackup, ok := s.D.GetOkExists("is_long_term_backup"); ok {
		tmp := isLongTermBackup.(bool)
		request.IsLongTermBackup = &tmp
	}

	if retentionPeriodInDays, ok := s.D.GetOkExists("retention_period_in_days"); ok {
		tmp := retentionPeriodInDays.(int)
		request.RetentionPeriodInDays = &tmp
	}

	if backupDestinationDetails, ok := s.D.GetOkExists("backup_destination_details"); ok {
		if tmpList := backupDestinationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "backup_destination_details", 0)
			tmp, err := s.mapToAutonomousBackupDestinationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.BackupDestinationDetails = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateAutonomousDatabaseBackup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		var identifier *string
		var err error
		identifier = response.Id
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	s.Res = &response.AutonomousDatabaseBackup
	return nil
}

func (s *DatabaseAutonomousDatabaseBackupResourceCrud) Get() error {
	request := oci_database.GetAutonomousDatabaseBackupRequest{}

	tmp := s.D.Id()
	request.AutonomousDatabaseBackupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetAutonomousDatabaseBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousDatabaseBackup
	return nil
}

func (s *DatabaseAutonomousDatabaseBackupResourceCrud) Update() error {
	request := oci_database.UpdateAutonomousDatabaseBackupRequest{}

	tmp := s.D.Id()
	request.AutonomousDatabaseBackupId = &tmp

	if retentionPeriodInDays, ok := s.D.GetOkExists("retention_period_in_days"); ok {
		tmp := retentionPeriodInDays.(int)
		request.RetentionPeriodInDays = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateAutonomousDatabaseBackup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "autonomousDatabaseBackup", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return s.Get()
}

func (s *DatabaseAutonomousDatabaseBackupResourceCrud) Delete() error {
	request := oci_database.DeleteAutonomousDatabaseBackupRequest{}

	tmp := s.D.Id()
	request.AutonomousDatabaseBackupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.DeleteAutonomousDatabaseBackup(context.Background(), request)
	return err
}

func (s *DatabaseAutonomousDatabaseBackupResourceCrud) SetData() error {
	if s.Res.AutonomousDatabaseId != nil {
		s.D.Set("autonomous_database_id", *s.Res.AutonomousDatabaseId)
	}

	if s.Res.BackupDestinationDetails != nil {
		s.D.Set("backup_destination_details", []interface{}{AutonomousBackupDestinationDetailsToMap(s.Res.BackupDestinationDetails)})
	} else {
		s.D.Set("backup_destination_details", nil)
	}

	//if s.Res.DbVersion != nil {
	//	s.D.Set("db_version", *s.Res.DbVersion)
	//}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DatabaseSizeInTBs != nil {
		s.D.Set("database_size_in_tbs", *s.Res.DatabaseSizeInTBs)
	}

	if s.Res.DbVersion != nil {
		s.D.Set("db_version", *s.Res.DbVersion)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.IsAutomatic != nil {
		s.D.Set("is_automatic", *s.Res.IsAutomatic)
	}

	if s.Res.IsRestorable != nil {
		s.D.Set("is_restorable", *s.Res.IsRestorable)
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

	if s.Res.SizeInTBs != nil {
		s.D.Set("size_in_tbs", *s.Res.SizeInTBs)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeAvailableTill != nil {
		s.D.Set("time_available_till", s.Res.TimeAvailableTill.String())
	}

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

	return nil
}

func (s *DatabaseAutonomousDatabaseBackupResourceCrud) mapToAutonomousBackupDestinationDetails(fieldKeyFormat string) (oci_database.BackupDestinationDetails, error) {
	result := oci_database.BackupDestinationDetails{}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	if internetProxy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "internet_proxy")); ok {
		tmp := internetProxy.(string)
		result.InternetProxy = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_database.BackupDestinationDetailsTypeEnum(type_.(string))
	}

	if vpcPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vpc_password")); ok {
		tmp := vpcPassword.(string)
		result.VpcPassword = &tmp
	}

	if vpcUser, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vpc_user")); ok {
		tmp := vpcUser.(string)
		result.VpcUser = &tmp
	}

	return result, nil
}

func AutonomousBackupDestinationDetailsToMap(obj *oci_database.BackupDestinationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.InternetProxy != nil {
		result["internet_proxy"] = string(*obj.InternetProxy)
	}

	result["type"] = string(obj.Type)

	if obj.VpcPassword != nil {
		result["vpc_password"] = string(*obj.VpcPassword)
	}

	if obj.VpcUser != nil {
		result["vpc_user"] = string(*obj.VpcUser)
	}

	return result
}
