// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package lustre_file_storage

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_lustre_file_storage "github.com/oracle/oci-go-sdk/v65/lustrefilestorage"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LustreFileStorageObjectStorageLinkSyncJobDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readSingularLustreFileStorageObjectStorageLinkSyncJobWithContext,
		Schema: map[string]*schema.Schema{
			"object_storage_link_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sync_job_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"bytes_transferred": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_overwrite": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"job_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lustre_file_system_path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"object_storage_path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"objects_transferred": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"skipped_error_count": {
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
			"time_finished": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_started": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_objects_scanned": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularLustreFileStorageObjectStorageLinkSyncJobWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &LustreFileStorageObjectStorageLinkSyncJobDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LustreFileStorageClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type LustreFileStorageObjectStorageLinkSyncJobDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_lustre_file_storage.LustreFileStorageClient
	Res    *oci_lustre_file_storage.GetSyncJobResponse
}

func (s *LustreFileStorageObjectStorageLinkSyncJobDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LustreFileStorageObjectStorageLinkSyncJobDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_lustre_file_storage.GetSyncJobRequest{}

	if objectStorageLinkId, ok := s.D.GetOkExists("object_storage_link_id"); ok {
		tmp := objectStorageLinkId.(string)
		request.ObjectStorageLinkId = &tmp
	}

	if syncJobId, ok := s.D.GetOkExists("sync_job_id"); ok {
		tmp := syncJobId.(string)
		request.SyncJobId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "lustre_file_storage")

	response, err := s.Client.GetSyncJob(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LustreFileStorageObjectStorageLinkSyncJobDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.BytesTransferred != nil {
		s.D.Set("bytes_transferred", strconv.FormatInt(*s.Res.BytesTransferred, 10))
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsOverwrite != nil {
		s.D.Set("is_overwrite", *s.Res.IsOverwrite)
	}

	s.D.Set("job_type", s.Res.JobType)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.LustreFileSystemPath != nil {
		s.D.Set("lustre_file_system_path", *s.Res.LustreFileSystemPath)
	}

	if s.Res.ObjectStoragePath != nil {
		s.D.Set("object_storage_path", *s.Res.ObjectStoragePath)
	}

	if s.Res.ObjectsTransferred != nil {
		s.D.Set("objects_transferred", strconv.FormatInt(*s.Res.ObjectsTransferred, 10))
	}

	if s.Res.ParentId != nil {
		s.D.Set("parent_id", *s.Res.ParentId)
	}

	if s.Res.SkippedErrorCount != nil {
		s.D.Set("skipped_error_count", strconv.FormatInt(*s.Res.SkippedErrorCount, 10))
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeFinished != nil {
		s.D.Set("time_finished", s.Res.TimeFinished.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	if s.Res.TotalObjectsScanned != nil {
		s.D.Set("total_objects_scanned", strconv.FormatInt(*s.Res.TotalObjectsScanned, 10))
	}

	return nil
}
