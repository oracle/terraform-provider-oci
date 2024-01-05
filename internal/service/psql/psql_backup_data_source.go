// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package psql

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_psql "github.com/oracle/oci-go-sdk/v65/psql"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func PsqlBackupDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["backup_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(PsqlBackupResource(), fieldMap, readSingularPsqlBackup)
}

func readSingularPsqlBackup(d *schema.ResourceData, m interface{}) error {
	sync := &PsqlBackupDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PostgresqlClient()

	return tfresource.ReadResource(sync)
}

type PsqlBackupDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_psql.PostgresqlClient
	Res    *oci_psql.GetBackupResponse
}

func (s *PsqlBackupDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *PsqlBackupDataSourceCrud) Get() error {
	request := oci_psql.GetBackupRequest{}

	if backupId, ok := s.D.GetOkExists("backup_id"); ok {
		tmp := backupId.(string)
		request.BackupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "psql")

	response, err := s.Client.GetBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *PsqlBackupDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.BackupSize != nil {
		s.D.Set("backup_size", *s.Res.BackupSize)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DbSystemDetails != nil {
		s.D.Set("db_system_details", []interface{}{DbSystemDetailsToMap(s.Res.DbSystemDetails)})
	} else {
		s.D.Set("db_system_details", nil)
	}

	if s.Res.DbSystemId != nil {
		s.D.Set("db_system_id", *s.Res.DbSystemId)
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
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LastAcceptedRequestToken != nil {
		s.D.Set("last_accepted_request_token", *s.Res.LastAcceptedRequestToken)
	}

	if s.Res.LastCompletedRequestToken != nil {
		s.D.Set("last_completed_request_token", *s.Res.LastCompletedRequestToken)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.RetentionPeriod != nil {
		s.D.Set("retention_period", *s.Res.RetentionPeriod)
	}

	s.D.Set("source_type", s.Res.SourceType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
