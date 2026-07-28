// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabasePluggableDatabaseSnapshotDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["pluggable_database_snapshot_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(DatabasePluggableDatabaseSnapshotResource(), fieldMap, readSingularDatabasePluggableDatabaseSnapshotWithContext)
}

func readSingularDatabasePluggableDatabaseSnapshotWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabasePluggableDatabaseSnapshotDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatabasePluggableDatabaseSnapshotDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetPluggableDatabaseSnapshotResponse
}

func (s *DatabasePluggableDatabaseSnapshotDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabasePluggableDatabaseSnapshotDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_database.GetPluggableDatabaseSnapshotRequest{}

	if pluggableDatabaseSnapshotId, ok := s.D.GetOkExists("pluggable_database_snapshot_id"); ok {
		tmp := pluggableDatabaseSnapshotId.(string)
		request.PluggableDatabaseSnapshotId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetPluggableDatabaseSnapshot(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabasePluggableDatabaseSnapshotDataSourceCrud) SetData() error {
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

	if s.Res.PluggableDatabaseId != nil {
		s.D.Set("pluggable_database_id", *s.Res.PluggableDatabaseId)
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
