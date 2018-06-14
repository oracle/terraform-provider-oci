// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_file_storage "github.com/oracle/oci-go-sdk/filestorage"
)

func ExportResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createExport,
		Read:     readExport,
		Delete:   deleteExport,
		Schema: map[string]*schema.Schema{
			// Required
			"export_set_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"file_system_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"path": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"id": {
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

func createExport(d *schema.ResourceData, m interface{}) error {
	sync := &ExportResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient

	return crud.CreateResource(d, sync)
}

func readExport(d *schema.ResourceData, m interface{}) error {
	sync := &ExportResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient

	return crud.ReadResource(sync)
}

func deleteExport(d *schema.ResourceData, m interface{}) error {
	sync := &ExportResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type ExportResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_file_storage.FileStorageClient
	Res                    *oci_file_storage.Export
	DisableNotFoundRetries bool
}

func (s *ExportResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ExportResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_file_storage.ExportLifecycleStateCreating),
	}
}

func (s *ExportResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_file_storage.ExportLifecycleStateActive),
	}
}

func (s *ExportResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_file_storage.ExportLifecycleStateDeleting),
	}
}

func (s *ExportResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_file_storage.ExportLifecycleStateDeleted),
	}
}

func (s *ExportResourceCrud) Create() error {
	request := oci_file_storage.CreateExportRequest{}

	if exportSetId, ok := s.D.GetOkExists("export_set_id"); ok {
		tmp := exportSetId.(string)
		request.ExportSetId = &tmp
	}

	if fileSystemId, ok := s.D.GetOkExists("file_system_id"); ok {
		tmp := fileSystemId.(string)
		request.FileSystemId = &tmp
	}

	if path, ok := s.D.GetOkExists("path"); ok {
		tmp := path.(string)
		request.Path = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.CreateExport(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Export
	return nil
}

func (s *ExportResourceCrud) Get() error {
	request := oci_file_storage.GetExportRequest{}

	tmp := s.D.Id()
	request.ExportId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.GetExport(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Export
	return nil
}

func (s *ExportResourceCrud) Delete() error {
	request := oci_file_storage.DeleteExportRequest{}

	tmp := s.D.Id()
	request.ExportId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	_, err := s.Client.DeleteExport(context.Background(), request)
	return err
}

func (s *ExportResourceCrud) SetData() {
	if s.Res.ExportSetId != nil {
		s.D.Set("export_set_id", *s.Res.ExportSetId)
	}

	if s.Res.FileSystemId != nil {
		s.D.Set("file_system_id", *s.Res.FileSystemId)
	}

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	if s.Res.Path != nil {
		s.D.Set("path", *s.Res.Path)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

}
