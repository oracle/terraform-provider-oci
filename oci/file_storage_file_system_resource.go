// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"

	oci_file_storage "github.com/oracle/oci-go-sdk/filestorage"
)

func init() {
	RegisterResource("oci_file_storage_file_system", FileStorageFileSystemResource())
}

func FileStorageFileSystemResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
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
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
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
			"kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
			},

			// Computed
			"metered_bytes": {
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

func createFileStorageFileSystem(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageFileSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient()

	return CreateResource(d, sync)
}

func readFileStorageFileSystem(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageFileSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient()

	return ReadResource(sync)
}

func updateFileStorageFileSystem(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageFileSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient()

	return UpdateResource(d, sync)
}

func deleteFileStorageFileSystem(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageFileSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type FileStorageFileSystemResourceCrud struct {
	BaseCrud
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

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
		tmp := kmsKeyId.(string)
		request.KmsKeyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.CreateFileSystem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.FileSystem
	return nil
}

func (s *FileStorageFileSystemResourceCrud) Get() error {
	request := oci_file_storage.GetFileSystemRequest{}

	tmp := s.D.Id()
	request.FileSystemId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "file_storage")

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
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
		tmp := kmsKeyId.(string)
		request.KmsKeyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.UpdateFileSystem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.FileSystem
	return nil
}

func (s *FileStorageFileSystemResourceCrud) Delete() error {
	request := oci_file_storage.DeleteFileSystemRequest{}

	tmp := s.D.Id()
	request.FileSystemId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	_, err := s.Client.DeleteFileSystem(context.Background(), request)
	return err
}

func (s *FileStorageFileSystemResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.KmsKeyId != nil {
		s.D.Set("kms_key_id", *s.Res.KmsKeyId)
	}

	if s.Res.MeteredBytes != nil {
		s.D.Set("metered_bytes", strconv.FormatInt(*s.Res.MeteredBytes, 10))
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *FileStorageFileSystemResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_file_storage.ChangeFileSystemCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.FileSystemId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	_, err := s.Client.ChangeFileSystemCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
