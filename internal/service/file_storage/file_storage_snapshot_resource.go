// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package file_storage

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_file_storage "github.com/oracle/oci-go-sdk/v56/filestorage"
)

func FileStorageSnapshotResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
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
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"is_clone_source": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"provenance_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
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
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.CreateResource(d, sync)
}

func readFileStorageSnapshot(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageSnapshotResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.ReadResource(sync)
}

func updateFileStorageSnapshot(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageSnapshotResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteFileStorageSnapshot(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageSnapshotResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type FileStorageSnapshotResourceCrud struct {
	tfresource.BaseCrud
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
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.CreateSnapshot(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Snapshot
	if waitErr := tfresource.WaitForCreatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *FileStorageSnapshotResourceCrud) Get() error {
	request := oci_file_storage.GetSnapshotRequest{}

	tmp := s.D.Id()
	request.SnapshotId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

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
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.SnapshotId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	_, err := s.Client.DeleteSnapshot(context.Background(), request)
	return err
}

func (s *FileStorageSnapshotResourceCrud) SetData() error {
	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.FileSystemId != nil {
		s.D.Set("file_system_id", *s.Res.FileSystemId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsCloneSource != nil {
		s.D.Set("is_clone_source", *s.Res.IsCloneSource)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ProvenanceId != nil {
		s.D.Set("provenance_id", *s.Res.ProvenanceId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
