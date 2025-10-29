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

func LustreFileStorageObjectStorageLinkSyncJobsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readLustreFileStorageObjectStorageLinkSyncJobsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"object_storage_link_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sync_job_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

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
									"id": {
										Type:     schema.TypeString,
										Computed: true,
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
							},
						},
					},
				},
			},
		},
	}
}

func readLustreFileStorageObjectStorageLinkSyncJobsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &LustreFileStorageObjectStorageLinkSyncJobsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LustreFileStorageClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type LustreFileStorageObjectStorageLinkSyncJobsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_lustre_file_storage.LustreFileStorageClient
	Res    *oci_lustre_file_storage.ListSyncJobsResponse
}

func (s *LustreFileStorageObjectStorageLinkSyncJobsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LustreFileStorageObjectStorageLinkSyncJobsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_lustre_file_storage.ListSyncJobsRequest{}

	if objectStorageLinkId, ok := s.D.GetOkExists("object_storage_link_id"); ok {
		tmp := objectStorageLinkId.(string)
		request.ObjectStorageLinkId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_lustre_file_storage.SyncJobLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "lustre_file_storage")

	response, err := s.Client.ListSyncJobs(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSyncJobs(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *LustreFileStorageObjectStorageLinkSyncJobsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LustreFileStorageObjectStorageLinkSyncJobsDataSource-", LustreFileStorageObjectStorageLinkSyncJobsDataSource(), s.D))
	resources := []map[string]interface{}{}
	objectStorageLinkSyncJob := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SyncJobSummaryToMap(item))
	}
	objectStorageLinkSyncJob["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, LustreFileStorageObjectStorageLinkSyncJobsDataSource().Schema["sync_job_collection"].Elem.(*schema.Resource).Schema)
		objectStorageLinkSyncJob["items"] = items
	}

	resources = append(resources, objectStorageLinkSyncJob)
	if err := s.D.Set("sync_job_collection", resources); err != nil {
		return err
	}

	return nil
}

func SyncJobSummaryToMap(obj oci_lustre_file_storage.SyncJobSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BytesTransferred != nil {
		result["bytes_transferred"] = strconv.FormatInt(*obj.BytesTransferred, 10)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsOverwrite != nil {
		result["is_overwrite"] = bool(*obj.IsOverwrite)
	}

	result["job_type"] = string(obj.JobType)

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.LustreFileSystemPath != nil {
		result["lustre_file_system_path"] = string(*obj.LustreFileSystemPath)
	}

	if obj.ObjectStoragePath != nil {
		result["object_storage_path"] = string(*obj.ObjectStoragePath)
	}

	if obj.ObjectsTransferred != nil {
		result["objects_transferred"] = strconv.FormatInt(*obj.ObjectsTransferred, 10)
	}

	if obj.SkippedErrorCount != nil {
		result["skipped_error_count"] = strconv.FormatInt(*obj.SkippedErrorCount, 10)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeFinished != nil {
		result["time_finished"] = obj.TimeFinished.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	if obj.TotalObjectsScanned != nil {
		result["total_objects_scanned"] = strconv.FormatInt(*obj.TotalObjectsScanned, 10)
	}

	return result
}
