// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseDatabaseSnapshotDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["database_snapshot_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(DatabaseDatabaseSnapshotResource(), fieldMap, readSingularDatabaseDatabaseSnapshotWithContext)
}

func readSingularDatabaseDatabaseSnapshotWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseDatabaseSnapshotDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatabaseDatabaseSnapshotDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetDatabaseSnapshotResponse
}

func (s *DatabaseDatabaseSnapshotDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDatabaseSnapshotDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database.GetDatabaseSnapshotRequest{}

	if databaseSnapshotId, ok := s.D.GetOkExists("database_snapshot_id"); ok {
		tmp := databaseSnapshotId.(string)
		request.DatabaseSnapshotId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetDatabaseSnapshot(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseDatabaseSnapshotDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ClusterId != nil {
		s.D.Set("cluster_id", *s.Res.ClusterId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DatabaseId != nil {
		s.D.Set("database_id", *s.Res.DatabaseId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
