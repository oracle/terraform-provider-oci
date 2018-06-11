// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_file_storage "github.com/oracle/oci-go-sdk/filestorage"
)

func ExportSetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createExportSet,
		Read:     readExportSet,
		Update:   updateExportSet,
		Delete:   deleteExportSet,
		Schema: map[string]*schema.Schema{
			// Required
			"mount_target_id": {
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
			"max_fs_stat_bytes": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_fs_stat_files": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// Computed
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
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
			"vcn_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createExportSet(d *schema.ResourceData, m interface{}) error {
	sync := &ExportSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient

	return crud.CreateResource(d, sync)
}

func readExportSet(d *schema.ResourceData, m interface{}) error {
	sync := &ExportSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient

	return crud.ReadResource(sync)
}

func updateExportSet(d *schema.ResourceData, m interface{}) error {
	sync := &ExportSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient

	return crud.UpdateResource(d, sync)
}

func deleteExportSet(d *schema.ResourceData, m interface{}) error {
	// Export set is deleted when a mount target is deleted.
	return nil
}

type ExportSetResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_file_storage.FileStorageClient
	Res                    *oci_file_storage.ExportSet
	DisableNotFoundRetries bool
}

func (s *ExportSetResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ExportSetResourceCrud) Create() error {
	// We can't really create an ExportSet. We need to get the exportSetId from the MountTarget it is attached to.
	if mountTargetId, ok := s.D.GetOkExists("mount_target_id"); ok {
		tmp := mountTargetId.(string)

		request := oci_file_storage.GetMountTargetRequest{}
		request.MountTargetId = &tmp

		request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "file_storage")

		response, err := s.Client.GetMountTarget(context.Background(), request)
		if err != nil {
			return fmt.Errorf("getting mount target details failed with error: %s", err.Error())
		}

		exportSetId := response.MountTarget.ExportSetId

		if exportSetId == nil {
			return fmt.Errorf("export_set_id is not available in the mount target response")
		}

		s.D.SetId(*exportSetId)
		return s.Update()
	}

	return fmt.Errorf("no mount_target_id value could be found")
}

func (s *ExportSetResourceCrud) Get() error {
	request := oci_file_storage.GetExportSetRequest{}

	tmp := s.D.Id()
	request.ExportSetId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.GetExportSet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExportSet
	return nil
}

func (s *ExportSetResourceCrud) Update() error {
	request := oci_file_storage.UpdateExportSetRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	tmp := s.D.Id()
	request.ExportSetId = &tmp

	if maxFsStatBytes, ok := s.D.GetOkExists("max_fs_stat_bytes"); ok {
		tmp := maxFsStatBytes.(int)
		request.MaxFsStatBytes = &tmp
	}

	if maxFsStatFiles, ok := s.D.GetOkExists("max_fs_stat_files"); ok {
		tmp := maxFsStatFiles.(int)
		request.MaxFsStatFiles = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.UpdateExportSet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExportSet
	return nil
}

func (s *ExportSetResourceCrud) Delete() error {
	// Export set is deleted when a mount target is deleted.
	return nil
}

func (s *ExportSetResourceCrud) SetData() {
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

	if s.Res.MaxFsStatBytes != nil {
		s.D.Set("max_fs_stat_bytes", *s.Res.MaxFsStatBytes)
	}

	if s.Res.MaxFsStatFiles != nil {
		s.D.Set("max_fs_stat_files", *s.Res.MaxFsStatFiles)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}
}
