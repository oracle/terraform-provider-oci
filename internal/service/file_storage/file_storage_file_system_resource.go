// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package file_storage

import (
	"context"
	"fmt"
	"strconv"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_file_storage "github.com/oracle/oci-go-sdk/v65/filestorage"
)

func FileStorageFileSystemResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFileStorageFileSystem,
		Read:     readFileStorageFileSystem,
		Update:   updateFileStorageFileSystem,
		Delete:   deleteFileStorageFileSystem,
		Schema: map[string]*schema.Schema{
			// Required
			"availability_domain": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"clone_attach_status": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.AttachDiffSuppressFunction,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"filesystem_snapshot_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
				// Commenting out `Computed: true` to allow unset policy in update FS operation.
				// This should be ok since control-api doesn’t compute this field if it's not provided by user
				//  Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_snapshot_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"detach_clone_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			// Computed
			"clone_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"is_clone_parent": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_hydrated": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_targetable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"metered_bytes": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"replication_target_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"parent_file_system_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source_snapshot_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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

func createFileStorageFileSystem(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageFileSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if _, ok := sync.D.GetOkExists("detach_clone_trigger"); ok {
		err := sync.DetachClone()
		if err != nil {
			return err
		}
	}
	return nil

}

func readFileStorageFileSystem(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageFileSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.ReadResource(sync)
}

func updateFileStorageFileSystem(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageFileSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	if _, ok := sync.D.GetOkExists("detach_clone_trigger"); ok && sync.D.HasChange("detach_clone_trigger") {
		oldRaw, newRaw := sync.D.GetChange("detach_clone_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.DetachClone()

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("detach_clone_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	return nil
}

func deleteFileStorageFileSystem(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageFileSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type FileStorageFileSystemResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_file_storage.FileStorageClient
	Res                    *oci_file_storage.FileSystem
	DisableNotFoundRetries bool
}

func (s *FileStorageFileSystemResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *FileStorageFileSystemResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_file_storage.FileSystemLifecycleStateCreating),
	}
}

func (s *FileStorageFileSystemResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_file_storage.FileSystemLifecycleStateActive),
	}
}

func (s *FileStorageFileSystemResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_file_storage.FileSystemLifecycleStateDeleting),
	}
}

func (s *FileStorageFileSystemResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_file_storage.FileSystemLifecycleStateDeleted),
	}
}

func (s *FileStorageFileSystemResourceCrud) Create() error {
	request := oci_file_storage.CreateFileSystemRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if cloneAttachStatus, ok := s.D.GetOkExists("clone_attach_status"); ok {
		request.CloneAttachStatus = oci_file_storage.CreateFileSystemDetailsCloneAttachStatusEnum(cloneAttachStatus.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if filesystemSnapshotPolicyId, ok := s.D.GetOkExists("filesystem_snapshot_policy_id"); ok {
		tmp := filesystemSnapshotPolicyId.(string)
		request.FilesystemSnapshotPolicyId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
		tmp := kmsKeyId.(string)
		request.KmsKeyId = &tmp
	}

	if sourceSnapshotId, ok := s.D.GetOkExists("source_snapshot_id"); ok {
		tmp := sourceSnapshotId.(string)
		request.SourceSnapshotId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.CreateFileSystem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.FileSystem
	if waitErr := tfresource.WaitForCreatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *FileStorageFileSystemResourceCrud) Get() error {
	request := oci_file_storage.GetFileSystemRequest{}

	tmp := s.D.Id()
	request.FileSystemId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.GetFileSystem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.FileSystem
	return nil
}

func (s *FileStorageFileSystemResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_file_storage.UpdateFileSystemRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	tmp := s.D.Id()
	request.FileSystemId = &tmp

	if filesystemSnapshotPolicyId, ok := s.D.GetOkExists("filesystem_snapshot_policy_id"); ok {
		tmp := filesystemSnapshotPolicyId.(string)
		request.FilesystemSnapshotPolicyId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
		tmp := kmsKeyId.(string)
		request.KmsKeyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.UpdateFileSystem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.FileSystem
	return nil
}

func (s *FileStorageFileSystemResourceCrud) Delete() error {
	request := oci_file_storage.DeleteFileSystemRequest{}

	if canDetachChildFileSystem, ok := s.D.GetOkExists("can_detach_child_file_system"); ok {
		tmp := canDetachChildFileSystem.(bool)
		request.CanDetachChildFileSystem = &tmp
	}

	tmp := s.D.Id()
	request.FileSystemId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	_, err := s.Client.DeleteFileSystem(context.Background(), request)
	return err
}

func (s *FileStorageFileSystemResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	s.D.Set("clone_attach_status", s.Res.CloneAttachStatus)

	if s.Res.CloneCount != nil {
		s.D.Set("clone_count", *s.Res.CloneCount)
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

	if s.Res.FilesystemSnapshotPolicyId != nil {
		s.D.Set("filesystem_snapshot_policy_id", *s.Res.FilesystemSnapshotPolicyId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsCloneParent != nil {
		s.D.Set("is_clone_parent", *s.Res.IsCloneParent)
	}

	if s.Res.IsHydrated != nil {
		s.D.Set("is_hydrated", *s.Res.IsHydrated)
	}

	if s.Res.IsTargetable != nil {
		s.D.Set("is_targetable", *s.Res.IsTargetable)
	}

	if s.Res.KmsKeyId != nil {
		s.D.Set("kms_key_id", *s.Res.KmsKeyId)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MeteredBytes != nil {
		s.D.Set("metered_bytes", strconv.FormatInt(*s.Res.MeteredBytes, 10))
	}

	if s.Res.ReplicationTargetId != nil {
		s.D.Set("replication_target_id", *s.Res.ReplicationTargetId)
	}

	if s.Res.SourceDetails != nil {
		s.D.Set("source_details", []interface{}{FileSystemSourceDetailsToMap(s.Res.SourceDetails)})
	} else {
		s.D.Set("source_details", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *FileStorageFileSystemResourceCrud) DetachClone() error {
	request := oci_file_storage.DetachCloneRequest{}

	idTmp := s.D.Id()
	request.FileSystemId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	_, err := s.Client.DetachClone(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("detach_clone_trigger")
	s.D.Set("detach_clone_trigger", val)

	return nil
}

func FileSystemSourceDetailsToMap(obj *oci_file_storage.SourceDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ParentFileSystemId != nil {
		result["parent_file_system_id"] = string(*obj.ParentFileSystemId)
	}

	if obj.SourceSnapshotId != nil {
		result["source_snapshot_id"] = string(*obj.SourceSnapshotId)
	}

	return result
}

func (s *FileStorageFileSystemResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_file_storage.ChangeFileSystemCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.FileSystemId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	_, err := s.Client.ChangeFileSystemCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
