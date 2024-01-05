// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package file_storage

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_file_storage "github.com/oracle/oci-go-sdk/v65/filestorage"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FileStorageReplicationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFileStorageReplication,
		Read:     readFileStorageReplication,
		Update:   updateFileStorageReplication,
		Delete:   deleteFileStorageReplication,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"source_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"target_id": {
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
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"replication_interval": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     tfresource.ValidateInt64TypeString,
				DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
			},

			// Computed
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"delta_progress": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"delta_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_snapshot_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"recovery_point_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"replication_target_id": {
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

func createFileStorageReplication(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageReplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.CreateResource(d, sync)
}

func readFileStorageReplication(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageReplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.ReadResource(sync)
}

func updateFileStorageReplication(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageReplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteFileStorageReplication(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageReplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type FileStorageReplicationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_file_storage.FileStorageClient
	Res                    *oci_file_storage.Replication
	DisableNotFoundRetries bool
}

func (s *FileStorageReplicationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *FileStorageReplicationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_file_storage.ReplicationLifecycleStateCreating),
	}
}

func (s *FileStorageReplicationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_file_storage.ReplicationLifecycleStateActive),
	}
}

func (s *FileStorageReplicationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_file_storage.ReplicationLifecycleStateDeleting),
	}
}

func (s *FileStorageReplicationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_file_storage.ReplicationLifecycleStateDeleted),
	}
}

func (s *FileStorageReplicationResourceCrud) Create() error {
	request := oci_file_storage.CreateReplicationRequest{}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if replicationInterval, ok := s.D.GetOkExists("replication_interval"); ok {
		tmp := replicationInterval.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert replicationInterval string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.ReplicationInterval = &tmpInt64
	}

	if sourceId, ok := s.D.GetOkExists("source_id"); ok {
		tmp := sourceId.(string)
		request.SourceId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.CreateReplication(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Replication
	return nil
}

func (s *FileStorageReplicationResourceCrud) Get() error {
	request := oci_file_storage.GetReplicationRequest{}

	tmp := s.D.Id()
	request.ReplicationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.GetReplication(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Replication
	return nil
}

func (s *FileStorageReplicationResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_file_storage.UpdateReplicationRequest{}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.ReplicationId = &tmp

	if replicationInterval, ok := s.D.GetOkExists("replication_interval"); ok {
		tmp := replicationInterval.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert replicationInterval string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.ReplicationInterval = &tmpInt64
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.UpdateReplication(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Replication
	return nil
}

func (s *FileStorageReplicationResourceCrud) Delete() error {
	request := oci_file_storage.DeleteReplicationRequest{}

	if deleteMode, ok := s.D.GetOkExists("delete_mode"); ok {
		request.DeleteMode = oci_file_storage.DeleteReplicationDeleteModeEnum(deleteMode.(string))
	}

	tmp := s.D.Id()
	request.ReplicationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	_, err := s.Client.DeleteReplication(context.Background(), request)
	return err
}

func (s *FileStorageReplicationResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DeltaProgress != nil {
		s.D.Set("delta_progress", strconv.FormatInt(*s.Res.DeltaProgress, 10))
	}

	s.D.Set("delta_status", s.Res.DeltaStatus)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LastSnapshotId != nil {
		s.D.Set("last_snapshot_id", *s.Res.LastSnapshotId)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.RecoveryPointTime != nil {
		s.D.Set("recovery_point_time", s.Res.RecoveryPointTime.String())
	}

	if s.Res.ReplicationInterval != nil {
		s.D.Set("replication_interval", strconv.FormatInt(*s.Res.ReplicationInterval, 10))
	}

	if s.Res.ReplicationTargetId != nil {
		s.D.Set("replication_target_id", *s.Res.ReplicationTargetId)
	}

	if s.Res.SourceId != nil {
		s.D.Set("source_id", *s.Res.SourceId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *FileStorageReplicationResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_file_storage.ChangeReplicationCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ReplicationId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	_, err := s.Client.ChangeReplicationCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
