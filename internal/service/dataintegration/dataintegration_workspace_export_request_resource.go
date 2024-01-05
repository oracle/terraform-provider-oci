// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dataintegration

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_dataintegration "github.com/oracle/oci-go-sdk/v65/dataintegration"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataintegrationWorkspaceExportRequestResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataintegrationWorkspaceExportRequest,
		Read:     readDataintegrationWorkspaceExportRequest,
		Delete:   deleteDataintegrationWorkspaceExportRequest,
		Schema: map[string]*schema.Schema{
			// Required
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"workspace_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"are_references_included": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"file_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"filters": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"is_object_overwrite_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"object_keys": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"object_storage_region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"object_storage_tenancy_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"error_messages": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"exported_items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"aggregator_key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"identifier": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name_path": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"object_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"object_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_updated_in_millis": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"referenced_items": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_ended_in_millis": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_started_in_millis": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_exported_object_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func createDataintegrationWorkspaceExportRequest(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceExportRequestResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.CreateResource(d, sync)
}

func readDataintegrationWorkspaceExportRequest(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceExportRequestResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.ReadResource(sync)
}

func deleteDataintegrationWorkspaceExportRequest(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceExportRequestResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataintegrationWorkspaceExportRequestResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_dataintegration.DataIntegrationClient
	Res                    *oci_dataintegration.ExportRequest
	DisableNotFoundRetries bool
}

func (s *DataintegrationWorkspaceExportRequestResourceCrud) ID() string {
	return GetWorkspaceExportRequestCompositeId(s.D.Get("key").(string), s.D.Get("workspace_id").(string))
}

func (s *DataintegrationWorkspaceExportRequestResourceCrud) Create() error {
	request := oci_dataintegration.CreateExportRequestRequest{}

	if areReferencesIncluded, ok := s.D.GetOkExists("are_references_included"); ok {
		tmp := areReferencesIncluded.(bool)
		request.AreReferencesIncluded = &tmp
	}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if fileName, ok := s.D.GetOkExists("file_name"); ok {
		tmp := fileName.(string)
		request.FileName = &tmp
	}

	if filters, ok := s.D.GetOkExists("filters"); ok {
		interfaces := filters.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("filters") {
			request.Filters = tmp
		}
	}

	if isObjectOverwriteEnabled, ok := s.D.GetOkExists("is_object_overwrite_enabled"); ok {
		tmp := isObjectOverwriteEnabled.(bool)
		request.IsObjectOverwriteEnabled = &tmp
	}

	if objectKeys, ok := s.D.GetOkExists("object_keys"); ok {
		interfaces := objectKeys.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("object_keys") {
			request.ObjectKeys = tmp
		}
	}

	if objectStorageRegion, ok := s.D.GetOkExists("object_storage_region"); ok {
		tmp := objectStorageRegion.(string)
		request.ObjectStorageRegion = &tmp
	}

	if objectStorageTenancyId, ok := s.D.GetOkExists("object_storage_tenancy_id"); ok {
		tmp := objectStorageTenancyId.(string)
		request.ObjectStorageTenancyId = &tmp
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	response, err := s.Client.CreateExportRequest(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExportRequest

	s.D.SetId(GetWorkspaceExportRequestCompositeId(*response.Key, *request.WorkspaceId))

	if setDataErr := s.SetData(); setDataErr != nil {
		log.Printf("[ERROR] error setting data before WaitForResourceCondition() error: %v", setDataErr)
	}

	retentionPolicyFunc := func() bool {
		return s.Res.Status == oci_dataintegration.ExportRequestStatusFailed || s.Res.Status == oci_dataintegration.ExportRequestStatusSuccessful
	}

	if err := tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate)); err != nil {
		return err
	}

	if s.Res.Status == oci_dataintegration.ExportRequestStatusFailed {
		log.Printf("[ERROR] export request failed with \"%s\"", s.Res.ErrorMessages[*s.Res.Key])
		return fmt.Errorf("export request failed with \"%s\"", s.Res.ErrorMessages[*s.Res.Key])
	}

	return nil
}

func (s *DataintegrationWorkspaceExportRequestResourceCrud) Get() error {
	request := oci_dataintegration.GetExportRequestRequest{}

	if exportRequestKey, ok := s.D.GetOkExists("key"); ok {
		tmp := exportRequestKey.(string)
		request.ExportRequestKey = &tmp
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	exportRequestKey, workspaceId, err := parseWorkspaceExportRequestCompositeId(s.D.Id())
	if err == nil {
		request.ExportRequestKey = &exportRequestKey
		request.WorkspaceId = &workspaceId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	response, err := s.Client.GetExportRequest(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExportRequest
	return nil
}

func (s *DataintegrationWorkspaceExportRequestResourceCrud) Delete() error {
	request := oci_dataintegration.DeleteExportRequestRequest{}

	if exportRequestKey, ok := s.D.GetOkExists("key"); ok {
		tmp := exportRequestKey.(string)
		request.ExportRequestKey = &tmp
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	_, err := s.Client.DeleteExportRequest(context.Background(), request)
	return err
}

func (s *DataintegrationWorkspaceExportRequestResourceCrud) SetData() error {

	exportRequestKey, workspaceId, err := parseWorkspaceExportRequestCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("key", &exportRequestKey)
		s.D.Set("workspace_id", &workspaceId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.AreReferencesIncluded != nil {
		s.D.Set("are_references_included", *s.Res.AreReferencesIncluded)
	}

	if s.Res.BucketName != nil {
		s.D.Set("bucket", *s.Res.BucketName)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	s.D.Set("error_messages", s.Res.ErrorMessages)

	exportedItems := []interface{}{}
	for _, item := range s.Res.ExportedItems {
		exportedItems = append(exportedItems, ExportObjectMetadataSummaryToMap(item))
	}
	s.D.Set("exported_items", exportedItems)

	if s.Res.FileName != nil {
		s.D.Set("file_name", *s.Res.FileName)
	}

	s.D.Set("filters", s.Res.Filters)

	if s.Res.IsObjectOverwriteEnabled != nil {
		s.D.Set("is_object_overwrite_enabled", *s.Res.IsObjectOverwriteEnabled)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("object_keys", s.Res.ObjectKeys)

	if s.Res.ObjectStorageRegion != nil {
		s.D.Set("object_storage_region", *s.Res.ObjectStorageRegion)
	}

	if s.Res.ObjectStorageTenancyId != nil {
		s.D.Set("object_storage_tenancy_id", *s.Res.ObjectStorageTenancyId)
	}

	if s.Res.ReferencedItems != nil {
		s.D.Set("referenced_items", *s.Res.ReferencedItems)
	}

	s.D.Set("status", s.Res.Status)

	if s.Res.TimeEndedInMillis != nil {
		s.D.Set("time_ended_in_millis", strconv.FormatInt(*s.Res.TimeEndedInMillis, 10))
	}

	if s.Res.TimeStartedInMillis != nil {
		s.D.Set("time_started_in_millis", strconv.FormatInt(*s.Res.TimeStartedInMillis, 10))
	}

	if s.Res.TotalExportedObjectCount != nil {
		s.D.Set("total_exported_object_count", *s.Res.TotalExportedObjectCount)
	}

	return nil
}

func GetWorkspaceExportRequestCompositeId(exportRequestKey string, workspaceId string) string {
	exportRequestKey = url.PathEscape(exportRequestKey)
	workspaceId = url.PathEscape(workspaceId)
	compositeId := "workspaces/" + workspaceId + "/exportRequests/" + exportRequestKey
	return compositeId
}

func parseWorkspaceExportRequestCompositeId(compositeId string) (exportRequestKey string, workspaceId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("workspaces/.*/exportRequests/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	workspaceId, _ = url.PathUnescape(parts[1])
	exportRequestKey, _ = url.PathUnescape(parts[3])

	return
}

func ExportObjectMetadataSummaryToMap(obj oci_dataintegration.ExportObjectMetadataSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AggregatorKey != nil {
		result["aggregator_key"] = string(*obj.AggregatorKey)
	}

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.NamePath != nil {
		result["name_path"] = string(*obj.NamePath)
	}

	if obj.ObjectType != nil {
		result["object_type"] = string(*obj.ObjectType)
	}

	if obj.ObjectVersion != nil {
		result["object_version"] = string(*obj.ObjectVersion)
	}

	if obj.TimeUpdatedInMillis != nil {
		result["time_updated_in_millis"] = strconv.FormatInt(*obj.TimeUpdatedInMillis, 10)
	}

	return result
}

func ExportRequestSummaryToMap(obj oci_dataintegration.ExportRequestSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AreReferencesIncluded != nil {
		result["are_references_included"] = bool(*obj.AreReferencesIncluded)
	}

	if obj.BucketName != nil {
		result["bucket"] = string(*obj.BucketName)
	}

	if obj.CreatedBy != nil {
		result["created_by"] = string(*obj.CreatedBy)
	}

	result["error_messages"] = obj.ErrorMessages

	exportedItems := []interface{}{}
	for _, item := range obj.ExportedItems {
		exportedItems = append(exportedItems, ExportObjectMetadataSummaryToMap(item))
	}
	result["exported_items"] = exportedItems

	if obj.FileName != nil {
		result["file_name"] = string(*obj.FileName)
	}

	result["filters"] = obj.Filters

	if obj.IsObjectOverwriteEnabled != nil {
		result["is_object_overwrite_enabled"] = bool(*obj.IsObjectOverwriteEnabled)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["object_keys"] = obj.ObjectKeys

	if obj.ObjectStorageRegion != nil {
		result["object_storage_region"] = string(*obj.ObjectStorageRegion)
	}

	if obj.ObjectStorageTenancyId != nil {
		result["object_storage_tenancy_id"] = string(*obj.ObjectStorageTenancyId)
	}

	if obj.ReferencedItems != nil {
		result["referenced_items"] = string(*obj.ReferencedItems)
	}

	result["status"] = string(obj.Status)

	if obj.TimeEndedInMillis != nil {
		result["time_ended_in_millis"] = strconv.FormatInt(*obj.TimeEndedInMillis, 10)
	}

	if obj.TimeStartedInMillis != nil {
		result["time_started_in_millis"] = strconv.FormatInt(*obj.TimeStartedInMillis, 10)
	}

	if obj.TotalExportedObjectCount != nil {
		result["total_exported_object_count"] = int(*obj.TotalExportedObjectCount)
	}

	return result
}
