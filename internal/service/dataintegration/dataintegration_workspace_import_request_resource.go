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

func DataintegrationWorkspaceImportRequestResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataintegrationWorkspaceImportRequest,
		Read:     readDataintegrationWorkspaceImportRequest,
		Delete:   deleteDataintegrationWorkspaceImportRequest,
		Schema: map[string]*schema.Schema{
			// Required
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"file_name": {
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
			"are_data_asset_references_included": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"import_conflict_resolution": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"import_conflict_resolution_type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"duplicate_prefix": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"duplicate_suffix": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"object_key_for_import": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
			"imported_objects": {
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
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name_path": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"new_key": {
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
						"old_key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resolution_action": {
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
			"total_imported_object_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func createDataintegrationWorkspaceImportRequest(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceImportRequestResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.CreateResource(d, sync)
}

func readDataintegrationWorkspaceImportRequest(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceImportRequestResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.ReadResource(sync)
}

func deleteDataintegrationWorkspaceImportRequest(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceImportRequestResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataintegrationWorkspaceImportRequestResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_dataintegration.DataIntegrationClient
	Res                    *oci_dataintegration.ImportRequest
	DisableNotFoundRetries bool
}

func (s *DataintegrationWorkspaceImportRequestResourceCrud) ID() string {
	return GetWorkspaceImportRequestCompositeId(s.D.Get("key").(string), s.D.Get("workspace_id").(string))
}

func (s *DataintegrationWorkspaceImportRequestResourceCrud) Create() error {
	request := oci_dataintegration.CreateImportRequestRequest{}

	if areDataAssetReferencesIncluded, ok := s.D.GetOkExists("are_data_asset_references_included"); ok {
		tmp := areDataAssetReferencesIncluded.(bool)
		request.AreDataAssetReferencesIncluded = &tmp
	}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if fileName, ok := s.D.GetOkExists("file_name"); ok {
		tmp := fileName.(string)
		request.FileName = &tmp
	}

	if importConflictResolution, ok := s.D.GetOkExists("import_conflict_resolution"); ok {
		if tmpList := importConflictResolution.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "import_conflict_resolution", 0)
			tmp, err := s.mapToImportConflictResolution(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ImportConflictResolution = &tmp
		}
	}

	if objectKeyForImport, ok := s.D.GetOkExists("object_key_for_import"); ok {
		tmp := objectKeyForImport.(string)
		request.ObjectKeyForImport = &tmp
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

	response, err := s.Client.CreateImportRequest(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ImportRequest

	s.D.SetId(GetWorkspaceImportRequestCompositeId(*response.Key, *request.WorkspaceId))

	if setDataErr := s.SetData(); setDataErr != nil {
		log.Printf("[ERROR] error setting data before WaitForResourceCondition() error: %v", setDataErr)
	}

	retentionPolicyFunc := func() bool {
		return s.Res.Status == oci_dataintegration.ImportRequestStatusFailed || s.Res.Status == oci_dataintegration.ImportRequestStatusSuccessful
	}

	if err := tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate)); err != nil {
		return err
	}

	if s.Res.Status == oci_dataintegration.ImportRequestStatusFailed {
		log.Printf("[ERROR] import request failed with \"%s\"", s.Res.ErrorMessages[*s.Res.Key])
		return fmt.Errorf("import request failed with \"%s\"", s.Res.ErrorMessages[*s.Res.Key])
	}

	return nil
}

func (s *DataintegrationWorkspaceImportRequestResourceCrud) Get() error {
	request := oci_dataintegration.GetImportRequestRequest{}

	if importRequestKey, ok := s.D.GetOkExists("key"); ok {
		tmp := importRequestKey.(string)
		request.ImportRequestKey = &tmp
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	importRequestKey, workspaceId, err := parseWorkspaceImportRequestCompositeId(s.D.Id())
	if err == nil {
		request.ImportRequestKey = &importRequestKey
		request.WorkspaceId = &workspaceId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	response, err := s.Client.GetImportRequest(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ImportRequest
	return nil
}

func (s *DataintegrationWorkspaceImportRequestResourceCrud) Delete() error {
	request := oci_dataintegration.DeleteImportRequestRequest{}

	if importRequestKey, ok := s.D.GetOkExists("key"); ok {
		tmp := importRequestKey.(string)
		request.ImportRequestKey = &tmp
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	_, err := s.Client.DeleteImportRequest(context.Background(), request)
	return err
}

func (s *DataintegrationWorkspaceImportRequestResourceCrud) SetData() error {

	importRequestKey, workspaceId, err := parseWorkspaceImportRequestCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("key", &importRequestKey)
		s.D.Set("workspace_id", &workspaceId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.AreDataAssetReferencesIncluded != nil {
		s.D.Set("are_data_asset_references_included", *s.Res.AreDataAssetReferencesIncluded)
	}

	if s.Res.BucketName != nil {
		s.D.Set("bucket", *s.Res.BucketName)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	s.D.Set("error_messages", s.Res.ErrorMessages)

	if s.Res.FileName != nil {
		s.D.Set("file_name", *s.Res.FileName)
	}

	if s.Res.ImportConflictResolution != nil {
		s.D.Set("import_conflict_resolution", []interface{}{ImportConflictResolutionToMap(s.Res.ImportConflictResolution)})
	} else {
		s.D.Set("import_conflict_resolution", nil)
	}

	importedObjects := []interface{}{}
	for _, item := range s.Res.ImportedObjects {
		importedObjects = append(importedObjects, ImportObjectMetadataSummaryToMap(item))
	}
	s.D.Set("imported_objects", importedObjects)

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ObjectKeyForImport != nil {
		s.D.Set("object_key_for_import", *s.Res.ObjectKeyForImport)
	}

	if s.Res.ObjectStorageRegion != nil {
		s.D.Set("object_storage_region", *s.Res.ObjectStorageRegion)
	}

	if s.Res.ObjectStorageTenancyId != nil {
		s.D.Set("object_storage_tenancy_id", *s.Res.ObjectStorageTenancyId)
	}

	s.D.Set("status", s.Res.Status)

	if s.Res.TimeEndedInMillis != nil {
		s.D.Set("time_ended_in_millis", strconv.FormatInt(*s.Res.TimeEndedInMillis, 10))
	}

	if s.Res.TimeStartedInMillis != nil {
		s.D.Set("time_started_in_millis", strconv.FormatInt(*s.Res.TimeStartedInMillis, 10))
	}

	if s.Res.TotalImportedObjectCount != nil {
		s.D.Set("total_imported_object_count", *s.Res.TotalImportedObjectCount)
	}

	return nil
}

func GetWorkspaceImportRequestCompositeId(importRequestKey string, workspaceId string) string {
	importRequestKey = url.PathEscape(importRequestKey)
	workspaceId = url.PathEscape(workspaceId)
	compositeId := "workspaces/" + workspaceId + "/importRequests/" + importRequestKey
	return compositeId
}

func parseWorkspaceImportRequestCompositeId(compositeId string) (importRequestKey string, workspaceId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("workspaces/.*/importRequests/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	workspaceId, _ = url.PathUnescape(parts[1])
	importRequestKey, _ = url.PathUnescape(parts[3])

	return
}

func ImportConflictResolutionToMap(obj *oci_dataintegration.ImportConflictResolution) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DuplicatePrefix != nil {
		result["duplicate_prefix"] = string(*obj.DuplicatePrefix)
	}

	if obj.DuplicateSuffix != nil {
		result["duplicate_suffix"] = string(*obj.DuplicateSuffix)
	}

	result["import_conflict_resolution_type"] = string(obj.ImportConflictResolutionType)

	return result
}

func ImportObjectMetadataSummaryToMap(obj oci_dataintegration.ImportObjectMetadataSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AggregatorKey != nil {
		result["aggregator_key"] = string(*obj.AggregatorKey)
	}

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.NamePath != nil {
		result["name_path"] = string(*obj.NamePath)
	}

	if obj.NewKey != nil {
		result["new_key"] = string(*obj.NewKey)
	}

	if obj.ObjectType != nil {
		result["object_type"] = string(*obj.ObjectType)
	}

	if obj.ObjectVersion != nil {
		result["object_version"] = string(*obj.ObjectVersion)
	}

	if obj.OldKey != nil {
		result["old_key"] = string(*obj.OldKey)
	}

	result["resolution_action"] = string(obj.ResolutionAction)

	if obj.TimeUpdatedInMillis != nil {
		result["time_updated_in_millis"] = strconv.FormatInt(*obj.TimeUpdatedInMillis, 10)
	}

	return result
}

func ImportRequestSummaryToMap(obj oci_dataintegration.ImportRequestSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AreDataAssetReferencesIncluded != nil {
		result["are_data_asset_references_included"] = bool(*obj.AreDataAssetReferencesIncluded)
	}

	if obj.BucketName != nil {
		result["bucket"] = string(*obj.BucketName)
	}

	if obj.CreatedBy != nil {
		result["created_by"] = string(*obj.CreatedBy)
	}

	result["error_messages"] = obj.ErrorMessages

	if obj.FileName != nil {
		result["file_name"] = string(*obj.FileName)
	}

	if obj.ImportConflictResolution != nil {
		result["import_conflict_resolution"] = []interface{}{ImportConflictResolutionToMap(obj.ImportConflictResolution)}
	}

	importedObjects := []interface{}{}
	for _, item := range obj.ImportedObjects {
		importedObjects = append(importedObjects, ImportObjectMetadataSummaryToMap(item))
	}
	result["imported_objects"] = importedObjects

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ObjectKeyForImport != nil {
		result["object_key_for_import"] = string(*obj.ObjectKeyForImport)
	}

	if obj.ObjectStorageRegion != nil {
		result["object_storage_region"] = string(*obj.ObjectStorageRegion)
	}

	if obj.ObjectStorageTenancyId != nil {
		result["object_storage_tenancy_id"] = string(*obj.ObjectStorageTenancyId)
	}

	result["status"] = string(obj.Status)

	if obj.TimeEndedInMillis != nil {
		result["time_ended_in_millis"] = strconv.FormatInt(*obj.TimeEndedInMillis, 10)
	}

	if obj.TimeStartedInMillis != nil {
		result["time_started_in_millis"] = strconv.FormatInt(*obj.TimeStartedInMillis, 10)
	}

	if obj.TotalImportedObjectCount != nil {
		result["total_imported_object_count"] = int(*obj.TotalImportedObjectCount)
	}

	return result
}

func (s *DataintegrationWorkspaceImportRequestResourceCrud) mapToImportConflictResolution(fieldKeyFormat string) (oci_dataintegration.ImportConflictResolution, error) {
	result := oci_dataintegration.ImportConflictResolution{}

	if duplicatePrefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "duplicate_prefix")); ok {
		tmp := duplicatePrefix.(string)
		result.DuplicatePrefix = &tmp
	}

	if duplicateSuffix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "duplicate_suffix")); ok {
		tmp := duplicateSuffix.(string)
		result.DuplicateSuffix = &tmp
	}

	if importConflictResolutionType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "import_conflict_resolution_type")); ok {
		result.ImportConflictResolutionType = oci_dataintegration.ImportConflictResolutionImportConflictResolutionTypeEnum(importConflictResolutionType.(string))
	}

	return result, nil
}
