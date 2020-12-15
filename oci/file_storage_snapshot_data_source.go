// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_file_storage "github.com/oracle/oci-go-sdk/v31/filestorage"
)

func init() {
	RegisterDatasource("oci_file_storage_snapshot", FileStorageSnapshotDataSource())
}

func FileStorageSnapshotDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["snapshot_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(FileStorageSnapshotResource(), fieldMap, readSingularFileStorageSnapshot)
}

func readSingularFileStorageSnapshot(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageSnapshotDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient()

	return ReadResource(sync)
}

type FileStorageSnapshotDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_file_storage.FileStorageClient
	Res    *oci_file_storage.GetSnapshotResponse
}

func (s *FileStorageSnapshotDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FileStorageSnapshotDataSourceCrud) Get() error {
	request := oci_file_storage.GetSnapshotRequest{}

	if snapshotId, ok := s.D.GetOkExists("snapshot_id"); ok {
		tmp := snapshotId.(string)
		request.SnapshotId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "file_storage")

	response, err := s.Client.GetSnapshot(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FileStorageSnapshotDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.FileSystemId != nil {
		s.D.Set("file_system_id", *s.Res.FileSystemId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
