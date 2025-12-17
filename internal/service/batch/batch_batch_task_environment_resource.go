// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package batch

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_batch "github.com/oracle/oci-go-sdk/v65/batch"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func BatchBatchTaskEnvironmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createBatchBatchTaskEnvironmentWithContext,
		ReadContext:   readBatchBatchTaskEnvironmentWithContext,
		UpdateContext: updateBatchBatchTaskEnvironmentWithContext,
		DeleteContext: deleteBatchBatchTaskEnvironmentWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"image_url": {
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
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"security_context": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"fs_group": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"run_as_group": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"run_as_user": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"volumes": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"local_mount_directory_path": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"mount_target_export_path": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"mount_target_fqdn": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"NFS",
							}, true),
						},

						// Optional

						// Computed
					},
				},
			},
			"working_directory": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
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

func createBatchBatchTaskEnvironmentWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BatchBatchTaskEnvironmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BatchComputingClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readBatchBatchTaskEnvironmentWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BatchBatchTaskEnvironmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BatchComputingClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateBatchBatchTaskEnvironmentWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BatchBatchTaskEnvironmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BatchComputingClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteBatchBatchTaskEnvironmentWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BatchBatchTaskEnvironmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BatchComputingClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type BatchBatchTaskEnvironmentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_batch.BatchComputingClient
	Res                    *oci_batch.BatchTaskEnvironment
	DisableNotFoundRetries bool
}

func (s *BatchBatchTaskEnvironmentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *BatchBatchTaskEnvironmentResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *BatchBatchTaskEnvironmentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_batch.BatchTaskEnvironmentLifecycleStateActive),
	}
}

func (s *BatchBatchTaskEnvironmentResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *BatchBatchTaskEnvironmentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_batch.BatchTaskEnvironmentLifecycleStateDeleted),
	}
}

func (s *BatchBatchTaskEnvironmentResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_batch.CreateBatchTaskEnvironmentRequest{}

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

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if imageUrl, ok := s.D.GetOkExists("image_url"); ok {
		tmp := imageUrl.(string)
		request.ImageUrl = &tmp
	}

	if securityContext, ok := s.D.GetOkExists("security_context"); ok {
		if tmpList := securityContext.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "security_context", 0)
			tmp, err := s.mapToSecurityContext(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SecurityContext = &tmp
		}
	}

	if volumes, ok := s.D.GetOkExists("volumes"); ok {
		interfaces := volumes.([]interface{})
		tmp := make([]oci_batch.BatchTaskEnvironmentVolume, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "volumes", stateDataIndex)
			converted, err := s.mapToBatchTaskEnvironmentVolume(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("volumes") {
			request.Volumes = tmp
		}
	}

	if workingDirectory, ok := s.D.GetOkExists("working_directory"); ok {
		tmp := workingDirectory.(string)
		request.WorkingDirectory = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "batch")

	response, err := s.Client.CreateBatchTaskEnvironment(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.BatchTaskEnvironment
	return nil
}

func (s *BatchBatchTaskEnvironmentResourceCrud) getBatchTaskEnvironmentFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_batch.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	batchTaskEnvironmentId, err := batchTaskEnvironmentWaitForWorkRequest(ctx, workId, "batchtaskenvironment",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*batchTaskEnvironmentId)

	return s.GetWithContext(ctx)
}

func batchTaskEnvironmentWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "batch", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_batch.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func batchTaskEnvironmentWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_batch.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_batch.BatchComputingClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "batch")
	retryPolicy.ShouldRetryOperation = batchTaskEnvironmentWorkRequestShouldRetryFunc(timeout)

	response := oci_batch.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_batch.OperationStatusInProgress),
			string(oci_batch.OperationStatusAccepted),
			string(oci_batch.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_batch.OperationStatusSucceeded),
			string(oci_batch.OperationStatusFailed),
			string(oci_batch.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_batch.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_batch.OperationStatusFailed || response.Status == oci_batch.OperationStatusCanceled {
		return nil, getErrorFromBatchBatchTaskEnvironmentWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromBatchBatchTaskEnvironmentWorkRequest(ctx context.Context, client *oci_batch.BatchComputingClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_batch.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_batch.ListWorkRequestErrorsRequest{
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

func (s *BatchBatchTaskEnvironmentResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_batch.GetBatchTaskEnvironmentRequest{}

	tmp := s.D.Id()
	request.BatchTaskEnvironmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "batch")

	response, err := s.Client.GetBatchTaskEnvironment(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.BatchTaskEnvironment
	return nil
}

func (s *BatchBatchTaskEnvironmentResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_batch.UpdateBatchTaskEnvironmentRequest{}

	tmp := s.D.Id()
	request.BatchTaskEnvironmentId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "batch")

	response, err := s.Client.UpdateBatchTaskEnvironment(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.BatchTaskEnvironment
	return nil
}

func (s *BatchBatchTaskEnvironmentResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_batch.DeleteBatchTaskEnvironmentRequest{}

	tmp := s.D.Id()
	request.BatchTaskEnvironmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "batch")

	_, err := s.Client.DeleteBatchTaskEnvironment(ctx, request)
	return err
}

func (s *BatchBatchTaskEnvironmentResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.ImageUrl != nil {
		s.D.Set("image_url", *s.Res.ImageUrl)
	}

	if s.Res.SecurityContext != nil {
		s.D.Set("security_context", []interface{}{SecurityContextToMap(s.Res.SecurityContext)})
	} else {
		s.D.Set("security_context", nil)
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

	volumes := []interface{}{}
	for _, item := range s.Res.Volumes {
		volumes = append(volumes, BatchTaskEnvironmentVolumeToMap(item))
	}
	s.D.Set("volumes", volumes)

	if s.Res.WorkingDirectory != nil {
		s.D.Set("working_directory", *s.Res.WorkingDirectory)
	}

	return nil
}

func BatchTaskEnvironmentSummaryToMap(obj oci_batch.BatchTaskEnvironmentSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.ImageUrl != nil {
		result["image_url"] = string(*obj.ImageUrl)
	}

	if obj.SecurityContext != nil {
		result["security_context"] = []interface{}{SecurityContextToMap(obj.SecurityContext)}
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

	if obj.WorkingDirectory != nil {
		result["working_directory"] = string(*obj.WorkingDirectory)
	}

	return result
}

func (s *BatchBatchTaskEnvironmentResourceCrud) mapToBatchTaskEnvironmentVolume(fieldKeyFormat string) (oci_batch.BatchTaskEnvironmentVolume, error) {
	var baseObject oci_batch.BatchTaskEnvironmentVolume
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("NFS"):
		details := oci_batch.NfsVolume{}
		if localMountDirectoryPath, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "local_mount_directory_path")); ok {
			tmp := localMountDirectoryPath.(string)
			details.LocalMountDirectoryPath = &tmp
		}
		if mountTargetExportPath, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mount_target_export_path")); ok {
			tmp := mountTargetExportPath.(string)
			details.MountTargetExportPath = &tmp
		}
		if mountTargetFqdn, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mount_target_fqdn")); ok {
			tmp := mountTargetFqdn.(string)
			details.MountTargetFqdn = &tmp
		}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func BatchTaskEnvironmentVolumeToMap(obj oci_batch.BatchTaskEnvironmentVolume) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_batch.NfsVolume:
		result["type"] = "NFS"

		if v.LocalMountDirectoryPath != nil {
			result["local_mount_directory_path"] = string(*v.LocalMountDirectoryPath)
		}

		if v.MountTargetExportPath != nil {
			result["mount_target_export_path"] = string(*v.MountTargetExportPath)
		}

		if v.MountTargetFqdn != nil {
			result["mount_target_fqdn"] = string(*v.MountTargetFqdn)
		}

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *BatchBatchTaskEnvironmentResourceCrud) mapToSecurityContext(fieldKeyFormat string) (oci_batch.SecurityContext, error) {
	result := oci_batch.SecurityContext{}

	if fsGroup, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "fs_group")); ok {
		tmp := fsGroup.(int)
		result.FsGroup = &tmp
	}

	if runAsGroup, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "run_as_group")); ok {
		tmp := runAsGroup.(int)
		result.RunAsGroup = &tmp
	}

	if runAsUser, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "run_as_user")); ok {
		tmp := runAsUser.(int)
		result.RunAsUser = &tmp
	}

	return result, nil
}

func SecurityContextToMap(obj *oci_batch.SecurityContext) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FsGroup != nil {
		result["fs_group"] = int(*obj.FsGroup)
	}

	if obj.RunAsGroup != nil {
		result["run_as_group"] = int(*obj.RunAsGroup)
	}

	if obj.RunAsUser != nil {
		result["run_as_user"] = int(*obj.RunAsUser)
	}

	return result
}

func (s *BatchBatchTaskEnvironmentResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_batch.ChangeBatchTaskEnvironmentCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.BatchTaskEnvironmentId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "batch")

	response, err := s.Client.ChangeBatchTaskEnvironmentCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getBatchTaskEnvironmentFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "batch"), oci_batch.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
