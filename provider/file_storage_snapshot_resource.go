// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_file_storage "github.com/oracle/oci-go-sdk/filestorage"
)

func SnapshotResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createSnapshot,
		Read:     readSnapshot,
		Delete:   deleteSnapshot,
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

func createSnapshot(d *schema.ResourceData, m interface{}) error {
	sync := &SnapshotResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient

	return crud.CreateResource(d, sync)
}

func readSnapshot(d *schema.ResourceData, m interface{}) error {
	sync := &SnapshotResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient

	return crud.ReadResource(sync)
}

func deleteSnapshot(d *schema.ResourceData, m interface{}) error {
	sync := &SnapshotResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type SnapshotResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_file_storage.FileStorageClient
	Res                    *oci_file_storage.Snapshot
	DisableNotFoundRetries bool
}

func (s *SnapshotResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *SnapshotResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_file_storage.SnapshotLifecycleStateCreating),
	}
}

func (s *SnapshotResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_file_storage.SnapshotLifecycleStateActive),
	}
}

func (s *SnapshotResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_file_storage.SnapshotLifecycleStateDeleting),
	}
}

func (s *SnapshotResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_file_storage.SnapshotLifecycleStateDeleted),
	}
}

func (s *SnapshotResourceCrud) Create() error {
	request := oci_file_storage.CreateSnapshotRequest{}

	if fileSystemId, ok := s.D.GetOkExists("file_system_id"); ok {
		tmp := fileSystemId.(string)
		request.FileSystemId = &tmp
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

func (s *SnapshotResourceCrud) Get() error {
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

func (s *SnapshotResourceCrud) Delete() error {
	request := oci_file_storage.DeleteSnapshotRequest{}

	tmp := s.D.Id()
	request.SnapshotId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	_, err := s.Client.DeleteSnapshot(context.Background(), request)
	return err
}

func (s *SnapshotResourceCrud) SetData() {
	if s.Res.FileSystemId != nil {
		s.D.Set("file_system_id", *s.Res.FileSystemId)
	}

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

}
