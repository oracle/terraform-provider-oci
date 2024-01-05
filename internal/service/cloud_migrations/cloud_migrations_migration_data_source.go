// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_migrations

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cloud_migrations "github.com/oracle/oci-go-sdk/v65/cloudmigrations"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudMigrationsMigrationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["migration_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CloudMigrationsMigrationResource(), fieldMap, readSingularCloudMigrationsMigration)
}

func readSingularCloudMigrationsMigration(d *schema.ResourceData, m interface{}) error {
	sync := &CloudMigrationsMigrationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MigrationClient()

	return tfresource.ReadResource(sync)
}

type CloudMigrationsMigrationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_migrations.MigrationClient
	Res    *oci_cloud_migrations.GetMigrationResponse
}

func (s *CloudMigrationsMigrationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudMigrationsMigrationDataSourceCrud) Get() error {
	request := oci_cloud_migrations.GetMigrationRequest{}

	if migrationId, ok := s.D.GetOkExists("migration_id"); ok {
		tmp := migrationId.(string)
		request.MigrationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_migrations")

	response, err := s.Client.GetMigration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CloudMigrationsMigrationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsCompleted != nil {
		s.D.Set("is_completed", *s.Res.IsCompleted)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ReplicationScheduleId != nil {
		s.D.Set("replication_schedule_id", *s.Res.ReplicationScheduleId)
	}

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
