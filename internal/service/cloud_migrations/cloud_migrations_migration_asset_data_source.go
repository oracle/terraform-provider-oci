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

func CloudMigrationsMigrationAssetDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["migration_asset_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CloudMigrationsMigrationAssetResource(), fieldMap, readSingularCloudMigrationsMigrationAsset)
}

func readSingularCloudMigrationsMigrationAsset(d *schema.ResourceData, m interface{}) error {
	sync := &CloudMigrationsMigrationAssetDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MigrationClient()

	return tfresource.ReadResource(sync)
}

type CloudMigrationsMigrationAssetDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_migrations.MigrationClient
	Res    *oci_cloud_migrations.GetMigrationAssetResponse
}

func (s *CloudMigrationsMigrationAssetDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudMigrationsMigrationAssetDataSourceCrud) Get() error {
	request := oci_cloud_migrations.GetMigrationAssetRequest{}

	if migrationAssetId, ok := s.D.GetOkExists("migration_asset_id"); ok {
		tmp := migrationAssetId.(string)
		request.MigrationAssetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_migrations")

	response, err := s.Client.GetMigrationAsset(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CloudMigrationsMigrationAssetDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("depended_on_by", s.Res.DependedOnBy)
	s.D.Set("depended_on_by", s.Res.DependedOnBy)

	s.D.Set("migration_asset_depends_on", s.Res.DependsOn)
	s.D.Set("migration_asset_depends_on", s.Res.DependsOn)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MigrationId != nil {
		s.D.Set("migration_id", *s.Res.MigrationId)
	}

	s.D.Set("notifications", s.Res.Notifications)
	s.D.Set("notifications", s.Res.Notifications)

	if s.Res.ParentSnapshot != nil {
		s.D.Set("parent_snapshot", *s.Res.ParentSnapshot)
	}

	if s.Res.ReplicationCompartmentId != nil {
		s.D.Set("replication_compartment_id", *s.Res.ReplicationCompartmentId)
	}

	if s.Res.ReplicationScheduleId != nil {
		s.D.Set("replication_schedule_id", *s.Res.ReplicationScheduleId)
	}

	if s.Res.SnapShotBucketName != nil {
		s.D.Set("snap_shot_bucket_name", *s.Res.SnapShotBucketName)
	}

	s.D.Set("snapshots", s.Res.Snapshots)
	s.D.Set("snapshots", s.Res.Snapshots)

	if s.Res.SourceAssetId != nil {
		s.D.Set("source_asset_id", *s.Res.SourceAssetId)
	}

	//fake test
	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TenancyId != nil {
		s.D.Set("tenancy_id", *s.Res.TenancyId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.Type != nil {
		s.D.Set("type", *s.Res.Type)
	}

	return nil
}
