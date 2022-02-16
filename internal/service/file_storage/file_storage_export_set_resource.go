// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package file_storage

import (
	"context"
	"fmt"
	"strconv"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_file_storage "github.com/oracle/oci-go-sdk/v58/filestorage"
)

func FileStorageExportSetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFileStorageExportSet,
		Read:     readFileStorageExportSet,
		Update:   updateFileStorageExportSet,
		Delete:   deleteFileStorageExportSet,
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
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     utils.ValidateInt64TypeString,
				DiffSuppressFunc: utils.Int64StringDiffSuppressFunction,
			},
			"max_fs_stat_files": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     utils.ValidateInt64TypeString,
				DiffSuppressFunc: utils.Int64StringDiffSuppressFunction,
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

func createFileStorageExportSet(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageExportSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.CreateResource(d, sync)
}

func readFileStorageExportSet(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageExportSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.ReadResource(sync)
}

func updateFileStorageExportSet(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageExportSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteFileStorageExportSet(d *schema.ResourceData, m interface{}) error {
	// Export set is deleted when a mount target is deleted.
	return nil
}

type FileStorageExportSetResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_file_storage.FileStorageClient
	Res                    *oci_file_storage.ExportSet
	DisableNotFoundRetries bool
}

func (s *FileStorageExportSetResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *FileStorageExportSetResourceCrud) Create() error {
	// We can't really Create an ExportSet. We need to get the exportSetId from the MountTarget it is attached to.
	if mountTargetId, ok := s.D.GetOkExists("mount_target_id"); ok {
		tmp := mountTargetId.(string)

		request := oci_file_storage.GetMountTargetRequest{}
		request.MountTargetId = &tmp

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

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

func (s *FileStorageExportSetResourceCrud) Get() error {
	request := oci_file_storage.GetExportSetRequest{}

	tmp := s.D.Id()
	request.ExportSetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.GetExportSet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExportSet
	return nil
}

func (s *FileStorageExportSetResourceCrud) Update() error {
	request := oci_file_storage.UpdateExportSetRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	tmp := s.D.Id()
	request.ExportSetId = &tmp

	if maxFsStatBytes, ok := s.D.GetOkExists("max_fs_stat_bytes"); ok {
		tmp := maxFsStatBytes.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert maxFsStatBytes string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.MaxFsStatBytes = &tmpInt64
	}

	if maxFsStatFiles, ok := s.D.GetOkExists("max_fs_stat_files"); ok {
		tmp := maxFsStatFiles.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert maxFsStatFiles string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.MaxFsStatFiles = &tmpInt64
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.UpdateExportSet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExportSet
	return nil
}

func (s *FileStorageExportSetResourceCrud) Delete() error {
	// Export set is deleted when a mount target is deleted.
	return nil
}

func (s *FileStorageExportSetResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.MaxFsStatBytes != nil {
		s.D.Set("max_fs_stat_bytes", strconv.FormatInt(*s.Res.MaxFsStatBytes, 10))
	}

	if s.Res.MaxFsStatFiles != nil {
		s.D.Set("max_fs_stat_files", strconv.FormatInt(*s.Res.MaxFsStatFiles, 10))
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	return nil
}
