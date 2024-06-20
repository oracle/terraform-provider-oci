// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package mysql

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_mysql "github.com/oracle/oci-go-sdk/v65/mysql"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MysqlMysqlBackupDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["backup_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(MysqlMysqlBackupResource(), fieldMap, readSingularMysqlMysqlBackup)
}

func readSingularMysqlMysqlBackup(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlMysqlBackupDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbBackupsClient()

	return tfresource.ReadResource(sync)
}

type MysqlMysqlBackupDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_mysql.DbBackupsClient
	Res    *oci_mysql.GetBackupResponse
}

func (s *MysqlMysqlBackupDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MysqlMysqlBackupDataSourceCrud) Get() error {
	request := oci_mysql.GetBackupRequest{}

	if backupId, ok := s.D.GetOkExists("backup_id"); ok {
		tmp := backupId.(string)
		request.BackupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "mysql")

	response, err := s.Client.GetBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MysqlMysqlBackupDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

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

	if s.Res.DbSystemSnapshot != nil {
		s.D.Set("db_system_snapshot", []interface{}{DbSystemSnapshotToMap(s.Res.DbSystemSnapshot)})
	} else {
		s.D.Set("db_system_snapshot", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.ImmediateSourceBackupId != nil {
		s.D.Set("immediate_source_backup_id", *s.Res.ImmediateSourceBackupId)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MysqlVersion != nil {
		s.D.Set("mysql_version", *s.Res.MysqlVersion)
	}

	if s.Res.OriginalSourceBackupId != nil {
		s.D.Set("original_source_backup_id", *s.Res.OriginalSourceBackupId)
	}

	if s.Res.RetentionInDays != nil {
		s.D.Set("retention_in_days", *s.Res.RetentionInDays)
	}

	if s.Res.ShapeName != nil {
		s.D.Set("shape_name", *s.Res.ShapeName)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCopyCreated != nil {
		s.D.Set("time_copy_created", s.Res.TimeCopyCreated.String())
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
