// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_file_storage "github.com/oracle/oci-go-sdk/filestorage"
)

func FileSystemResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createFileSystem,
		Read:     readFileSystem,
		Update:   updateFileSystem,
		Delete:   deleteFileSystem,
		Schema: map[string]*schema.Schema{
			// Required
			"availability_domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"metered_bytes": {
				Type:     schema.TypeInt,
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

func createFileSystem(d *schema.ResourceData, m interface{}) error {
	sync := &FileSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient

	return crud.CreateResource(d, sync)
}

func readFileSystem(d *schema.ResourceData, m interface{}) error {
	sync := &FileSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient

	return crud.ReadResource(sync)
}

func updateFileSystem(d *schema.ResourceData, m interface{}) error {
	sync := &FileSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient

	return crud.UpdateResource(d, sync)
}

func deleteFileSystem(d *schema.ResourceData, m interface{}) error {
	sync := &FileSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type FileSystemResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_file_storage.FileStorageClient
	Res                    *oci_file_storage.FileSystem
	DisableNotFoundRetries bool
}

func (s *FileSystemResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *FileSystemResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_file_storage.FileSystemLifecycleStateCreating),
	}
}

func (s *FileSystemResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_file_storage.FileSystemLifecycleStateActive),
	}
}

func (s *FileSystemResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_file_storage.FileSystemLifecycleStateDeleting),
	}
}

func (s *FileSystemResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_file_storage.FileSystemLifecycleStateDeleted),
	}
}

func (s *FileSystemResourceCrud) Create() error {
	request := oci_file_storage.CreateFileSystemRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.CreateFileSystem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.FileSystem
	return nil
}

func (s *FileSystemResourceCrud) Get() error {
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

func (s *FileSystemResourceCrud) Update() error {
	request := oci_file_storage.UpdateFileSystemRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	tmp := s.D.Id()
	request.FileSystemId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.UpdateFileSystem(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.FileSystem
	return nil
}

func (s *FileSystemResourceCrud) Delete() error {
	request := oci_file_storage.DeleteFileSystemRequest{}

	tmp := s.D.Id()
	request.FileSystemId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	_, err := s.Client.DeleteFileSystem(context.Background(), request)
	return err
}

func (s *FileSystemResourceCrud) SetData() {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	if s.Res.MeteredBytes != nil {
		s.D.Set("metered_bytes", *s.Res.MeteredBytes)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

}
