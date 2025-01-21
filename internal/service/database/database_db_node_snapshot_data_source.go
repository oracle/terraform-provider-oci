// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseDbNodeSnapshotDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["dbnode_snapshot_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseDbNodeSnapshotResource(), fieldMap, readSingularDatabaseDbNodeSnapshot)
}

func readSingularDatabaseDbNodeSnapshot(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbNodeSnapshotDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseDbNodeSnapshotDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetDbnodeSnapshotResponse
}

func (s *DatabaseDbNodeSnapshotDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDbNodeSnapshotDataSourceCrud) Get() error {
	request := oci_database.GetDbnodeSnapshotRequest{}

	if dbnodeSnapshotId, ok := s.D.GetOkExists("dbnode_snapshot_id"); ok {
		tmp := dbnodeSnapshotId.(string)
		request.DbnodeSnapshotId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetDbnodeSnapshot(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseDbNodeSnapshotDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.Id != nil {
		s.D.Set("dbnode_snapshot_id", *s.Res.Id)
	}

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

	mountPoints := []interface{}{}
	for _, item := range s.Res.MountPoints {
		mountPoints = append(mountPoints, MountPointDetailsToMap(item))
	}
	s.D.Set("mount_points", mountPoints)

	if len(mountPoints) > 0 {
		mountPoint0 := mountPoints[0].(map[string]interface{})
		if tmp, exist := mountPoint0["db_node_id"]; exist {
			dbNodeId := tmp.(string)
			s.D.Set("mount_dbnode_id", dbNodeId)
		}
	} else {
		s.D.Set("mount_dbnode_id", "null")
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.SourceDbnodeId != nil {
		s.D.Set("source_dbnode_id", *s.Res.SourceDbnodeId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	volumes := []interface{}{}
	for _, item := range s.Res.Volumes {
		volumes = append(volumes, VolumeDetailsToMap(item))
	}
	s.D.Set("volumes", volumes)

	return nil
}
