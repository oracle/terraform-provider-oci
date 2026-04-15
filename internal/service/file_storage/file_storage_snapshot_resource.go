// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package file_storage

import (
	"context"
	"fmt"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_file_storage "github.com/oracle/oci-go-sdk/v65/filestorage"
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
			"expiration_time": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"lock_duration_details": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"lock_duration": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"lock_mode": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"cool_off_duration": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"locks": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"message": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"related_resource_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"time_created": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
						},

						// Computed
					},
				},
			},
			"is_lock_override": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			// Computed
			"filesystem_snapshot_policy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
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
			"snapshot_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"snapshot_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_locked": {
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

	if expirationTime, ok := s.D.GetOkExists("expiration_time"); ok {
		tmp, err := time.Parse(time.RFC3339, expirationTime.(string))
		if err != nil {
			return err
		}
		request.ExpirationTime = &oci_common.SDKTime{Time: tmp}
	}

	if fileSystemId, ok := s.D.GetOkExists("file_system_id"); ok {
		tmp := fileSystemId.(string)
		request.FileSystemId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if lockDurationDetails, ok := s.D.GetOkExists("lock_duration_details"); ok {
		if tmpList := lockDurationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "lock_duration_details", 0)
			tmp, err := s.mapToLockDurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LockDurationDetails = &tmp
		} else {
			request.LockDurationDetails = nil
		}
	} else {
		request.LockDurationDetails = nil
	}

	if locks, ok := s.D.GetOkExists("locks"); ok {
		interfaces := locks.([]interface{})
		tmp := make([]oci_file_storage.ResourceLock, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
			converted, err := s.mapToResourceLock(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("locks") {
			request.Locks = tmp
		}
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

	if expirationTime, ok := s.D.GetOkExists("expiration_time"); ok {
		tmp, err := time.Parse(time.RFC3339, expirationTime.(string))
		if err != nil {
			return err
		}
		request.ExpirationTime = &oci_common.SDKTime{Time: tmp}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		request.IsLockOverride = &tmp
	}

	if lockDurationDetails, ok := s.D.GetOkExists("lock_duration_details"); ok {
		if tmpList := lockDurationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "lock_duration_details", 0)
			tmp, err := s.mapToLockDurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LockDurationDetails = &tmp
		} else {
			request.LockDurationDetails = nil
		}
	} else {
		request.LockDurationDetails = nil
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

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		request.IsLockOverride = &tmp
	}

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

	if s.Res.LockDurationDetails != nil {
		s.D.Set("lock_duration_details", []interface{}{LockDurationDetailsToMap(s.Res.LockDurationDetails)})
	} else {
		s.D.Set("lock_duration_details", []interface{}{})
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
		s.D.Set("snapshot_time", s.Res.SnapshotTime.Format(time.RFC3339))
	}

	s.D.Set("snapshot_type", s.Res.SnapshotType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.Format(time.RFC3339))
	}

	if s.Res.TimeLocked != nil {
		s.D.Set("time_locked", s.Res.TimeLocked.Format(time.RFC3339))
	}

	return nil
}

func (s *FileStorageSnapshotResourceCrud) mapToLockDurationDetails(fieldKeyFormat string) (oci_file_storage.LockDurationDetails, error) {
	result := oci_file_storage.LockDurationDetails{}

	if coolOffDuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cool_off_duration")); ok {
		tmp := coolOffDuration.(int)
		result.CoolOffDuration = &tmp
	}

	if lockDuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "lock_duration")); ok {
		tmp := lockDuration.(int)
		result.LockDuration = &tmp
	}

	if lockMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "lock_mode")); ok {
		result.LockMode = oci_file_storage.LockDurationDetailsLockModeEnum(lockMode.(string))
	}

	return result, nil
}

func LockDurationDetailsToMap(obj *oci_file_storage.LockDurationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CoolOffDuration != nil {
		result["cool_off_duration"] = int(*obj.CoolOffDuration)
	}

	if obj.LockDuration != nil {
		result["lock_duration"] = int(*obj.LockDuration)
	}

	result["lock_mode"] = string(obj.LockMode)

	return result
}

func (s *FileStorageSnapshotResourceCrud) mapToResourceLock(fieldKeyFormat string) (oci_file_storage.ResourceLock, error) {
	result := oci_file_storage.ResourceLock{}

	if message, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "message")); ok {
		tmp := message.(string)
		result.Message = &tmp
	}

	if relatedResourceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "related_resource_id")); ok {
		tmp := relatedResourceId.(string)
		result.RelatedResourceId = &tmp
	}

	if timeCreated, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_created")); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreated.(string))
		if err != nil {
			return result, err
		}
		result.TimeCreated = &oci_common.SDKTime{Time: tmp}
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_file_storage.ResourceLockTypeEnum(type_.(string))
	}

	return result, nil
}
