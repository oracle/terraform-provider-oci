// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	oci_file_storage "github.com/oracle/oci-go-sdk/filestorage"
)

func FileStorageSnapshotResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createFileStorageSnapshot,
		Read:     readFileStorageSnapshot,
		Update:   updateFileStorageSnapshot,
		Delete:   deleteFileStorageSnapshot,
		Schema: map[string]*schema.Schema{
			// Required
			"file_system_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createFileStorageSnapshot(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageSnapshotResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient

	return CreateResource(d, sync)
}

func readFileStorageSnapshot(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageSnapshotResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient

	return ReadResource(sync)
}

func updateFileStorageSnapshot(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageSnapshotResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient

	return UpdateResource(d, sync)
}

func deleteFileStorageSnapshot(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageSnapshotResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type FileStorageSnapshotResourceCrud struct {
	BaseCrud
	Client                 *oci_file_storage.FileStorageClient
	Res                    *oci_file_storage.Snapshot
	DisableNotFoundRetries bool
}

func (s *FileStorageSnapshotResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *FileStorageSnapshotResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_file_storage.SnapshotLifecycleStateCreating),
	}
}

func (s *FileStorageSnapshotResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_file_storage.SnapshotLifecycleStateActive),
	}
}

func (s *FileStorageSnapshotResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_file_storage.SnapshotLifecycleStateDeleting),
	}
}

func (s *FileStorageSnapshotResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_file_storage.SnapshotLifecycleStateDeleted),
	}
}

func (s *FileStorageSnapshotResourceCrud) Create() error {
	request := oci_file_storage.CreateSnapshotRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if fileSystemId, ok := s.D.GetOkExists("file_system_id"); ok {
		tmp := fileSystemId.(string)
		request.FileSystemId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.CreateSnapshot(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Snapshot
	return nil
}

func (s *FileStorageSnapshotResourceCrud) Get() error {
	request := oci_file_storage.GetSnapshotRequest{}

	tmp := s.D.Id()
	request.SnapshotId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.GetSnapshot(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Snapshot
	return nil
}

func (s *FileStorageSnapshotResourceCrud) Update() error {
	request := oci_file_storage.UpdateSnapshotRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.SnapshotId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.UpdateSnapshot(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Snapshot
	return nil
}

func (s *FileStorageSnapshotResourceCrud) Delete() error {
	request := oci_file_storage.DeleteSnapshotRequest{}

	tmp := s.D.Id()
	request.SnapshotId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	_, err := s.Client.DeleteSnapshot(context.Background(), request)
	return err
}

func (s *FileStorageSnapshotResourceCrud) SetData() error {
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
