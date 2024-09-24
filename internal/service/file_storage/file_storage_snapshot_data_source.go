// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package file_storage

import (
	"context"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_file_storage "github.com/oracle/oci-go-sdk/v65/filestorage"
)

func FileStorageSnapshotDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["snapshot_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(FileStorageSnapshotResource(), fieldMap, readSingularFileStorageSnapshot)
}

func readSingularFileStorageSnapshot(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageSnapshotDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "file_storage")

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
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.ExpirationTime != nil {
		s.D.Set("expiration_time", s.Res.ExpirationTime.Format(time.RFC3339Nano))
	}

	if s.Res.FileSystemId != nil {
		s.D.Set("file_system_id", *s.Res.FileSystemId)
	}

	if s.Res.FilesystemSnapshotPolicyId != nil {
		s.D.Set("filesystem_snapshot_policy_id", *s.Res.FilesystemSnapshotPolicyId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsCloneSource != nil {
		s.D.Set("is_clone_source", *s.Res.IsCloneSource)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	locks := []interface{}{}
	for _, item := range s.Res.Locks {
		locks = append(locks, ResourceLockToMap(item))
	}
	s.D.Set("locks", locks)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ProvenanceId != nil {
		s.D.Set("provenance_id", *s.Res.ProvenanceId)
	}

	if s.Res.SnapshotTime != nil {
		s.D.Set("snapshot_time", s.Res.SnapshotTime.String())
	}

	s.D.Set("snapshot_type", s.Res.SnapshotType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
