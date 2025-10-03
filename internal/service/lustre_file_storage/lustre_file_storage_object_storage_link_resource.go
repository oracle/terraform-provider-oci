// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package lustre_file_storage

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_lustre_file_storage "github.com/oracle/oci-go-sdk/v65/lustrefilestorage"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LustreFileStorageObjectStorageLinkResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createLustreFileStorageObjectStorageLinkWithContext,
		ReadContext:   readLustreFileStorageObjectStorageLinkWithContext,
		UpdateContext: updateLustreFileStorageObjectStorageLinkWithContext,
		DeleteContext: deleteLustreFileStorageObjectStorageLinkWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"availability_domain": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"file_system_path": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"is_overwrite": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"lustre_file_system_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"object_storage_prefix": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
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
			"start_export_to_object_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"start_import_from_object_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"stop_export_to_object_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"stop_import_from_object_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			// Computed
			"current_job_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_job_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
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
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createLustreFileStorageObjectStorageLinkWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &LustreFileStorageObjectStorageLinkResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LustreFileStorageClient()

	if e := tfresource.CreateResourceWithContext(ctx, d, sync); e != nil {
		return tfresource.HandleDiagError(m, e)
	}

	if _, ok := sync.D.GetOkExists("start_export_to_object_trigger"); ok {
		err := sync.StartExportToObject()
		if err != nil {
			return tfresource.HandleDiagError(m, err)
		}
	}

	if _, ok := sync.D.GetOkExists("start_import_from_object_trigger"); ok {
		err := sync.StartImportFromObject()
		if err != nil {
			return tfresource.HandleDiagError(m, err)
		}
	}

	if _, ok := sync.D.GetOkExists("stop_export_to_object_trigger"); ok {
		err := sync.StopExportToObject()
		if err != nil {
			return tfresource.HandleDiagError(m, err)
		}
	}

	if _, ok := sync.D.GetOkExists("stop_import_from_object_trigger"); ok {
		err := sync.StopImportFromObject()
		if err != nil {
			return tfresource.HandleDiagError(m, err)
		}
	}
	return nil

}

func readLustreFileStorageObjectStorageLinkWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &LustreFileStorageObjectStorageLinkResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LustreFileStorageClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateLustreFileStorageObjectStorageLinkWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &LustreFileStorageObjectStorageLinkResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LustreFileStorageClient()

	if _, ok := sync.D.GetOkExists("start_export_to_object_trigger"); ok && sync.D.HasChange("start_export_to_object_trigger") {
		oldRaw, newRaw := sync.D.GetChange("start_export_to_object_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.StartExportToObject()

			if err != nil {
				return tfresource.HandleDiagError(m, err)
			}
		} else {
			sync.D.Set("start_export_to_object_trigger", oldRaw)
			return diag.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if _, ok := sync.D.GetOkExists("start_import_from_object_trigger"); ok && sync.D.HasChange("start_import_from_object_trigger") {
		oldRaw, newRaw := sync.D.GetChange("start_import_from_object_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.StartImportFromObject()

			if err != nil {
				return tfresource.HandleDiagError(m, err)
			}
		} else {
			sync.D.Set("start_import_from_object_trigger", oldRaw)
			return diag.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if _, ok := sync.D.GetOkExists("stop_export_to_object_trigger"); ok && sync.D.HasChange("stop_export_to_object_trigger") {
		oldRaw, newRaw := sync.D.GetChange("stop_export_to_object_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.StopExportToObject()

			if err != nil {
				return tfresource.HandleDiagError(m, err)
			}
		} else {
			sync.D.Set("stop_export_to_object_trigger", oldRaw)
			return diag.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if _, ok := sync.D.GetOkExists("stop_import_from_object_trigger"); ok && sync.D.HasChange("stop_import_from_object_trigger") {
		oldRaw, newRaw := sync.D.GetChange("stop_import_from_object_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.StopImportFromObject()

			if err != nil {
				return tfresource.HandleDiagError(m, err)
			}
		} else {
			sync.D.Set("stop_import_from_object_trigger", oldRaw)
			return diag.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if err := tfresource.UpdateResourceWithContext(ctx, d, sync); err != nil {
		return tfresource.HandleDiagError(m, err)
	}

	return nil
}

func deleteLustreFileStorageObjectStorageLinkWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &LustreFileStorageObjectStorageLinkResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LustreFileStorageClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type LustreFileStorageObjectStorageLinkResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_lustre_file_storage.LustreFileStorageClient
	Res                    *oci_lustre_file_storage.ObjectStorageLink
	DisableNotFoundRetries bool
}

func (s *LustreFileStorageObjectStorageLinkResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *LustreFileStorageObjectStorageLinkResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_lustre_file_storage.ObjectStorageLinkLifecycleStateCreating),
	}
}

func (s *LustreFileStorageObjectStorageLinkResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_lustre_file_storage.ObjectStorageLinkLifecycleStateActive),
	}
}

func (s *LustreFileStorageObjectStorageLinkResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_lustre_file_storage.ObjectStorageLinkLifecycleStateDeleting),
	}
}

func (s *LustreFileStorageObjectStorageLinkResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_lustre_file_storage.ObjectStorageLinkLifecycleStateDeleted),
	}
}

func (s *LustreFileStorageObjectStorageLinkResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_lustre_file_storage.CreateObjectStorageLinkRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if fileSystemPath, ok := s.D.GetOkExists("file_system_path"); ok {
		tmp := fileSystemPath.(string)
		request.FileSystemPath = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isOverwrite, ok := s.D.GetOkExists("is_overwrite"); ok {
		tmp := isOverwrite.(bool)
		request.IsOverwrite = &tmp
	}

	if lustreFileSystemId, ok := s.D.GetOkExists("lustre_file_system_id"); ok {
		tmp := lustreFileSystemId.(string)
		request.LustreFileSystemId = &tmp
	}

	if objectStoragePrefix, ok := s.D.GetOkExists("object_storage_prefix"); ok {
		tmp := objectStoragePrefix.(string)
		request.ObjectStoragePrefix = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "lustre_file_storage")

	response, err := s.Client.CreateObjectStorageLink(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.ObjectStorageLink
	return nil
}

func (s *LustreFileStorageObjectStorageLinkResourceCrud) getObjectStorageLinkFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_lustre_file_storage.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	objectStorageLinkId, err := objectStorageLinkWaitForWorkRequest(ctx, workId, "objectstoragelink",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, objectStorageLinkId)
		_, cancelErr := s.Client.CancelWorkRequest(ctx,
			oci_lustre_file_storage.CancelWorkRequestRequest{
				WorkRequestId: workId,
				RequestMetadata: oci_common.RequestMetadata{
					RetryPolicy: retryPolicy,
				},
			})
		if cancelErr != nil {
			log.Printf("[DEBUG] cleanup cancelWorkRequest failed with the error: %v\n", cancelErr)
		}
		return err
	}
	s.D.SetId(*objectStorageLinkId)

	return s.GetWithContext(ctx)
}

func objectStorageLinkWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "lustre_file_storage", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_lustre_file_storage.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func objectStorageLinkWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_lustre_file_storage.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_lustre_file_storage.LustreFileStorageClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "lustre_file_storage")
	retryPolicy.ShouldRetryOperation = objectStorageLinkWorkRequestShouldRetryFunc(timeout)

	response := oci_lustre_file_storage.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_lustre_file_storage.OperationStatusInProgress),
			string(oci_lustre_file_storage.OperationStatusAccepted),
			string(oci_lustre_file_storage.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_lustre_file_storage.OperationStatusSucceeded),
			string(oci_lustre_file_storage.OperationStatusFailed),
			string(oci_lustre_file_storage.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_lustre_file_storage.GetWorkRequestRequest{
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.WorkRequest
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_lustre_file_storage.OperationStatusFailed || response.Status == oci_lustre_file_storage.OperationStatusCanceled {
		return nil, getErrorFromLustreFileStorageObjectStorageLinkWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromLustreFileStorageObjectStorageLinkWorkRequest(ctx context.Context, client *oci_lustre_file_storage.LustreFileStorageClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_lustre_file_storage.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_lustre_file_storage.ListWorkRequestErrorsRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		})
	if err != nil {
		return err
	}

	allErrs := make([]string, 0)
	for _, wrkErr := range response.Items {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *LustreFileStorageObjectStorageLinkResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_lustre_file_storage.GetObjectStorageLinkRequest{}

	tmp := s.D.Id()
	request.ObjectStorageLinkId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "lustre_file_storage")

	response, err := s.Client.GetObjectStorageLink(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.ObjectStorageLink
	return nil
}

func (s *LustreFileStorageObjectStorageLinkResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_lustre_file_storage.UpdateObjectStorageLinkRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isOverwrite, ok := s.D.GetOkExists("is_overwrite"); ok {
		tmp := isOverwrite.(bool)
		request.IsOverwrite = &tmp
	}

	tmp := s.D.Id()
	request.ObjectStorageLinkId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "lustre_file_storage")

	response, err := s.Client.UpdateObjectStorageLink(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.ObjectStorageLink
	return nil
}

func (s *LustreFileStorageObjectStorageLinkResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_lustre_file_storage.DeleteObjectStorageLinkRequest{}

	tmp := s.D.Id()
	request.ObjectStorageLinkId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "lustre_file_storage")

	response, err := s.Client.DeleteObjectStorageLink(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := objectStorageLinkWaitForWorkRequest(ctx, workId, "objectstoragelink",
		oci_lustre_file_storage.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *LustreFileStorageObjectStorageLinkResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CurrentJobId != nil {
		s.D.Set("current_job_id", *s.Res.CurrentJobId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.FileSystemPath != nil {
		s.D.Set("file_system_path", *s.Res.FileSystemPath)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsOverwrite != nil {
		s.D.Set("is_overwrite", *s.Res.IsOverwrite)
	}

	if s.Res.LastJobId != nil {
		s.D.Set("last_job_id", *s.Res.LastJobId)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.LustreFileSystemId != nil {
		s.D.Set("lustre_file_system_id", *s.Res.LustreFileSystemId)
	}

	if s.Res.ObjectStoragePrefix != nil {
		s.D.Set("object_storage_prefix", *s.Res.ObjectStoragePrefix)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *LustreFileStorageObjectStorageLinkResourceCrud) StartExportToObject() error {
	request := oci_lustre_file_storage.StartExportToObjectRequest{}

	idTmp := s.D.Id()
	request.ObjectStorageLinkId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "lustre_file_storage")

	_, err := s.Client.StartExportToObject(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("start_export_to_object_trigger")
	s.D.Set("start_export_to_object_trigger", val)

	// s.Res = &response.ObjectStorageLink
	return nil
}

func (s *LustreFileStorageObjectStorageLinkResourceCrud) StartImportFromObject() error {
	request := oci_lustre_file_storage.StartImportFromObjectRequest{}

	idTmp := s.D.Id()
	request.ObjectStorageLinkId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "lustre_file_storage")

	_, err := s.Client.StartImportFromObject(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("start_import_from_object_trigger")
	s.D.Set("start_import_from_object_trigger", val)

	// s.Res = &response.ObjectStorageLink
	return nil
}

func (s *LustreFileStorageObjectStorageLinkResourceCrud) StopExportToObject() error {
	request := oci_lustre_file_storage.StopExportToObjectRequest{}

	idTmp := s.D.Id()
	request.ObjectStorageLinkId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "lustre_file_storage")

	_, err := s.Client.StopExportToObject(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("stop_export_to_object_trigger")
	s.D.Set("stop_export_to_object_trigger", val)

	// s.Res = &response.ObjectStorageLink
	return nil
}

func (s *LustreFileStorageObjectStorageLinkResourceCrud) StopImportFromObject() error {
	request := oci_lustre_file_storage.StopImportFromObjectRequest{}

	idTmp := s.D.Id()
	request.ObjectStorageLinkId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "lustre_file_storage")

	_, err := s.Client.StopImportFromObject(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("stop_import_from_object_trigger")
	s.D.Set("stop_import_from_object_trigger", val)

	// s.Res = &response.ObjectStorageLink
	return nil
}

func ObjectStorageLinkSummaryToMap(obj oci_lustre_file_storage.ObjectStorageLinkSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.CurrentJobId != nil {
		result["current_job_id"] = string(*obj.CurrentJobId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.FileSystemPath != nil {
		result["file_system_path"] = string(*obj.FileSystemPath)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsOverwrite != nil {
		result["is_overwrite"] = bool(*obj.IsOverwrite)
	}

	if obj.LastJobId != nil {
		result["last_job_id"] = string(*obj.LastJobId)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.LustreFileSystemId != nil {
		result["lustre_file_system_id"] = string(*obj.LustreFileSystemId)
	}

	if obj.ObjectStoragePrefix != nil {
		result["object_storage_prefix"] = string(*obj.ObjectStoragePrefix)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *LustreFileStorageObjectStorageLinkResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_lustre_file_storage.ChangeObjectStorageLinkCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ObjectStorageLinkId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "lustre_file_storage")

	_, err := s.Client.ChangeObjectStorageLinkCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
