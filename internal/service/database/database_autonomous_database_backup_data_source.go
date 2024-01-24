// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseAutonomousDatabaseBackupDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["autonomous_database_backup_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseAutonomousDatabaseBackupResource(), fieldMap, readSingularDatabaseAutonomousDatabaseBackup)
}

func readSingularDatabaseAutonomousDatabaseBackup(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseBackupDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousDatabaseBackupDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetAutonomousDatabaseBackupResponse
}

func (s *DatabaseAutonomousDatabaseBackupDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousDatabaseBackupDataSourceCrud) Get() error {
	request := oci_database.GetAutonomousDatabaseBackupRequest{}

	if autonomousDatabaseBackupId, ok := s.D.GetOkExists("autonomous_database_backup_id"); ok {
		tmp := autonomousDatabaseBackupId.(string)
		request.AutonomousDatabaseBackupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetAutonomousDatabaseBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseAutonomousDatabaseBackupDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AutonomousDatabaseId != nil {
		s.D.Set("autonomous_database_id", *s.Res.AutonomousDatabaseId)
	}

	if s.Res.BackupDestinationDetails != nil {
		s.D.Set("backup_destination_details", []interface{}{AutonomousBackupDestinationDetailsToMap(s.Res.BackupDestinationDetails)})
	} else {
		s.D.Set("backup_destination_details", nil)
	}

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
		s.D.Set("time_ended", s.Res.TimeEnded.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	s.D.Set("type", s.Res.Type)

	if s.Res.VaultId != nil {
		s.D.Set("vault_id", *s.Res.VaultId)
	}

	return nil
}
