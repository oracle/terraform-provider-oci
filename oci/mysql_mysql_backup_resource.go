// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_mysql "github.com/oracle/oci-go-sdk/mysql"
)

func init() {
	RegisterResource("oci_mysql_mysql_backup", MysqlMysqlBackupResource())
}

func MysqlMysqlBackupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createMysqlMysqlBackup,
		Read:     readMysqlMysqlBackup,
		Update:   updateMysqlMysqlBackup,
		Delete:   deleteMysqlMysqlBackup,
		Schema: map[string]*schema.Schema{
			// Required
			"db_system_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"backup_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"retention_in_days": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// Computed
			"backup_size_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"creation_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_storage_size_in_gb": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mysql_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"shape_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createMysqlMysqlBackup(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlMysqlBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dbBackupsClient()

	return CreateResource(d, sync)
}

func readMysqlMysqlBackup(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlMysqlBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dbBackupsClient()

	return ReadResource(sync)
}

func updateMysqlMysqlBackup(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlMysqlBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dbBackupsClient()

	return UpdateResource(d, sync)
}

func deleteMysqlMysqlBackup(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlMysqlBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dbBackupsClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type MysqlMysqlBackupResourceCrud struct {
	BaseCrud
	Client                 *oci_mysql.DbBackupsClient
	Res                    *oci_mysql.Backup
	DisableNotFoundRetries bool
}

func (s *MysqlMysqlBackupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *MysqlMysqlBackupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_mysql.BackupLifecycleStateCreating),
	}
}

func (s *MysqlMysqlBackupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_mysql.BackupLifecycleStateActive),
	}
}

func (s *MysqlMysqlBackupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_mysql.BackupLifecycleStateDeleting),
	}
}

func (s *MysqlMysqlBackupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_mysql.BackupLifecycleStateDeleted),
	}
}

func (s *MysqlMysqlBackupResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_mysql.BackupLifecycleStateUpdating),
	}
}

func (s *MysqlMysqlBackupResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_mysql.BackupLifecycleStateActive),
	}
}

func (s *MysqlMysqlBackupResourceCrud) Create() error {
	request := oci_mysql.CreateBackupRequest{}

	if backupType, ok := s.D.GetOkExists("backup_type"); ok {
		request.BackupType = oci_mysql.CreateBackupDetailsBackupTypeEnum(backupType.(string))
	}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if retentionInDays, ok := s.D.GetOkExists("retention_in_days"); ok {
		tmp := retentionInDays.(int)
		request.RetentionInDays = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "mysql")

	response, err := s.Client.CreateBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Backup
	return nil
}

func (s *MysqlMysqlBackupResourceCrud) Get() error {
	request := oci_mysql.GetBackupRequest{}

	tmp := s.D.Id()
	request.BackupId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "mysql")

	response, err := s.Client.GetBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Backup
	return nil
}

func (s *MysqlMysqlBackupResourceCrud) Update() error {
	request := oci_mysql.UpdateBackupRequest{}

	tmp := s.D.Id()
	request.BackupId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if retentionInDays, ok := s.D.GetOkExists("retention_in_days"); ok {
		tmp := retentionInDays.(int)
		request.RetentionInDays = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.UpdateBackup(context.Background(), request)
	if err != nil {
		return err
	}

	return s.Get()
}

func (s *MysqlMysqlBackupResourceCrud) Delete() error {
	request := oci_mysql.DeleteBackupRequest{}

	tmp := s.D.Id()
	request.BackupId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.DeleteBackup(context.Background(), request)
	return err
}

func (s *MysqlMysqlBackupResourceCrud) SetData() error {
	if s.Res.BackupSizeInGBs != nil {
		s.D.Set("backup_size_in_gbs", *s.Res.BackupSizeInGBs)
	}

	s.D.Set("backup_type", s.Res.BackupType)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("creation_type", s.Res.CreationType)

	if s.Res.DataStorageSizeInGBs != nil {
		s.D.Set("data_storage_size_in_gb", *s.Res.DataStorageSizeInGBs)
	}

	if s.Res.DbSystemId != nil {
		s.D.Set("db_system_id", *s.Res.DbSystemId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MysqlVersion != nil {
		s.D.Set("mysql_version", *s.Res.MysqlVersion)
	}

	if s.Res.RetentionInDays != nil {
		s.D.Set("retention_in_days", *s.Res.RetentionInDays)
	}

	if s.Res.ShapeName != nil {
		s.D.Set("shape_name", *s.Res.ShapeName)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
