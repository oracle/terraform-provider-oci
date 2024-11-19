// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package recovery

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_recovery "github.com/oracle/oci-go-sdk/v65/recovery"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func RecoveryLongTermBackupDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["long_term_backup_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(RecoveryLongTermBackupResource(), fieldMap, readSingularRecoveryLongTermBackup)
}

func readSingularRecoveryLongTermBackup(d *schema.ResourceData, m interface{}) error {
	sync := &RecoveryLongTermBackupDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseRecoveryClient()

	return tfresource.ReadResource(sync)
}

type RecoveryLongTermBackupDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_recovery.DatabaseRecoveryClient
	Res    *oci_recovery.GetLongTermBackupResponse
}

func (s *RecoveryLongTermBackupDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *RecoveryLongTermBackupDataSourceCrud) Get() error {
	request := oci_recovery.GetLongTermBackupRequest{}

	if longTermBackupId, ok := s.D.GetOkExists("long_term_backup_id"); ok {
		tmp := longTermBackupId.(string)
		request.LongTermBackupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "recovery")

	response, err := s.Client.GetLongTermBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *RecoveryLongTermBackupDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DatabaseIdentifier != nil {
		s.D.Set("database_identifier", *s.Res.DatabaseIdentifier)
	}

	if s.Res.DatabaseSizeInGBs != nil {
		s.D.Set("database_size_in_gbs", *s.Res.DatabaseSizeInGBs)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("lifecycle_substate", s.Res.LifecycleSubstate)

	if s.Res.ProtectedDatabaseId != nil {
		s.D.Set("protected_database_id", *s.Res.ProtectedDatabaseId)
	}

	retentionPeriod := []interface{}{}
	for _, item := range s.Res.RetentionPeriod {
		retentionPeriod = append(retentionPeriod, RetentionPeriodValueToMap(item))
	}
	s.D.Set("retention_period", retentionPeriod)

	if s.Res.RetentionPointInTime != nil {
		s.D.Set("retention_point_in_time", s.Res.RetentionPointInTime.Format(time.RFC3339Nano))
	}

	if s.Res.RetentionScn != nil {
		s.D.Set("retention_scn", *s.Res.RetentionScn)
	}

	if s.Res.RetentionUntilDateTime != nil {
		s.D.Set("retention_until_date_time", s.Res.RetentionUntilDateTime.String())
	}

	if s.Res.RmanTag != nil {
		s.D.Set("rman_tag", *s.Res.RmanTag)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeBackupCompleted != nil {
		s.D.Set("time_backup_completed", s.Res.TimeBackupCompleted.String())
	}

	if s.Res.TimeBackupInitiated != nil {
		s.D.Set("time_backup_initiated", s.Res.TimeBackupInitiated.String())
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
