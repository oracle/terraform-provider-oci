// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package file_storage

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_file_storage "github.com/oracle/oci-go-sdk/v65/filestorage"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FileStorageFilesystemSnapshotPolicyDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["filesystem_snapshot_policy_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(FileStorageFilesystemSnapshotPolicyResource(), fieldMap, readSingularFileStorageFilesystemSnapshotPolicy)
}

func readSingularFileStorageFilesystemSnapshotPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageFilesystemSnapshotPolicyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.ReadResource(sync)
}

type FileStorageFilesystemSnapshotPolicyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_file_storage.FileStorageClient
	Res    *oci_file_storage.GetFilesystemSnapshotPolicyResponse
}

func (s *FileStorageFilesystemSnapshotPolicyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FileStorageFilesystemSnapshotPolicyDataSourceCrud) Get() error {
	request := oci_file_storage.GetFilesystemSnapshotPolicyRequest{}

	if filesystemSnapshotPolicyId, ok := s.D.GetOkExists("filesystem_snapshot_policy_id"); ok {
		tmp := filesystemSnapshotPolicyId.(string)
		request.FilesystemSnapshotPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "file_storage")

	response, err := s.Client.GetFilesystemSnapshotPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FileStorageFilesystemSnapshotPolicyDataSourceCrud) SetData() error {
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

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.PolicyPrefix != nil {
		s.D.Set("policy_prefix", *s.Res.PolicyPrefix)
	}

	schedules := []interface{}{}
	for _, item := range s.Res.Schedules {
		schedules = append(schedules, SnapshotScheduleToMap(item))
	}
	s.D.Set("schedules", schedules)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
