// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_data_platform

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ai_data_platform "github.com/oracle/oci-go-sdk/v65/aidataplatform"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AiDataPlatformAiDataPlatformResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},
		CreateContext: createAiDataPlatformAiDataPlatformWithContext,
		ReadContext:   readAiDataPlatformAiDataPlatformWithContext,
		UpdateContext: updateAiDataPlatformAiDataPlatformWithContext,
		DeleteContext: deleteAiDataPlatformAiDataPlatformWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"ai_data_platform_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_workspace_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
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
			"system_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"alias_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_by": {
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
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"web_socket_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createAiDataPlatformAiDataPlatformWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &AiDataPlatformAiDataPlatformResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiDataPlatformClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readAiDataPlatformAiDataPlatformWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &AiDataPlatformAiDataPlatformResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiDataPlatformClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateAiDataPlatformAiDataPlatformWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &AiDataPlatformAiDataPlatformResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiDataPlatformClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteAiDataPlatformAiDataPlatformWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &AiDataPlatformAiDataPlatformResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiDataPlatformClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type AiDataPlatformAiDataPlatformResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_ai_data_platform.AiDataPlatformClient
	Res                    *oci_ai_data_platform.AiDataPlatform
	DisableNotFoundRetries bool
}

func (s *AiDataPlatformAiDataPlatformResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *AiDataPlatformAiDataPlatformResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_ai_data_platform.AiDataPlatformLifecycleStateCreating),
	}
}

func (s *AiDataPlatformAiDataPlatformResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_ai_data_platform.AiDataPlatformLifecycleStateActive),
	}
}

func (s *AiDataPlatformAiDataPlatformResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_ai_data_platform.AiDataPlatformLifecycleStateDeleting),
	}
}

func (s *AiDataPlatformAiDataPlatformResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_ai_data_platform.AiDataPlatformLifecycleStateDeleted),
	}
}

func (s *AiDataPlatformAiDataPlatformResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_ai_data_platform.CreateAiDataPlatformRequest{}

	if aiDataPlatformType, ok := s.D.GetOkExists("ai_data_platform_type"); ok {
		tmp := aiDataPlatformType.(string)
		request.AiDataPlatformType = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if defaultWorkspaceName, ok := s.D.GetOkExists("default_workspace_name"); ok {
		tmp := defaultWorkspaceName.(string)
		request.DefaultWorkspaceName = &tmp
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if systemTags, ok := s.D.GetOkExists("system_tags"); ok {
		convertedSystemTags, err := tfresource.MapToSystemTags(systemTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.SystemTags = convertedSystemTags
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_data_platform")

	response, err := s.Client.CreateAiDataPlatform(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getAiDataPlatformFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_data_platform"), oci_ai_data_platform.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *AiDataPlatformAiDataPlatformResourceCrud) getAiDataPlatformFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_ai_data_platform.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	aiDataPlatformId, err := aiDataPlatformWaitForWorkRequest(ctx, workId, "datahub",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, aiDataPlatformId)
		_, cancelErr := s.Client.CancelWorkRequest(ctx,
			oci_ai_data_platform.CancelWorkRequestRequest{
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
	s.D.SetId(*aiDataPlatformId)

	return s.GetWithContext(ctx)
}

func aiDataPlatformWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "ai_data_platform", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_ai_data_platform.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func aiDataPlatformWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_ai_data_platform.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_ai_data_platform.AiDataPlatformClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "ai_data_platform")
	retryPolicy.ShouldRetryOperation = aiDataPlatformWorkRequestShouldRetryFunc(timeout)

	response := oci_ai_data_platform.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_ai_data_platform.OperationStatusInProgress),
			string(oci_ai_data_platform.OperationStatusAccepted),
			string(oci_ai_data_platform.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_ai_data_platform.OperationStatusSucceeded),
			string(oci_ai_data_platform.OperationStatusFailed),
			string(oci_ai_data_platform.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_ai_data_platform.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_ai_data_platform.OperationStatusFailed || response.Status == oci_ai_data_platform.OperationStatusCanceled {
		return nil, getErrorFromAiDataPlatformAiDataPlatformWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromAiDataPlatformAiDataPlatformWorkRequest(ctx context.Context, client *oci_ai_data_platform.AiDataPlatformClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_ai_data_platform.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_ai_data_platform.ListWorkRequestErrorsRequest{
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

func (s *AiDataPlatformAiDataPlatformResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_ai_data_platform.GetAiDataPlatformRequest{}

	tmp := s.D.Id()
	request.AiDataPlatformId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_data_platform")

	response, err := s.Client.GetAiDataPlatform(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.AiDataPlatform
	return nil
}

func (s *AiDataPlatformAiDataPlatformResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_ai_data_platform.UpdateAiDataPlatformRequest{}

	tmp := s.D.Id()
	request.AiDataPlatformId = &tmp

	if aiDataPlatformType, ok := s.D.GetOkExists("ai_data_platform_type"); ok {
		tmp := aiDataPlatformType.(string)
		request.AiDataPlatformType = &tmp
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if systemTags, ok := s.D.GetOkExists("system_tags"); ok {
		convertedSystemTags, err := tfresource.MapToSystemTags(systemTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.SystemTags = convertedSystemTags
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_data_platform")

	response, err := s.Client.UpdateAiDataPlatform(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAiDataPlatformFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_data_platform"), oci_ai_data_platform.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *AiDataPlatformAiDataPlatformResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_ai_data_platform.DeleteAiDataPlatformRequest{}

	tmp := s.D.Id()
	request.AiDataPlatformId = &tmp

	if isForceDelete, ok := s.D.GetOkExists("is_force_delete"); ok {
		tmp := isForceDelete.(bool)
		request.IsForceDelete = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_data_platform")

	response, err := s.Client.DeleteAiDataPlatform(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := aiDataPlatformWaitForWorkRequest(ctx, workId, "datahub",
		oci_ai_data_platform.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *AiDataPlatformAiDataPlatformResourceCrud) SetData() error {
	if s.Res.AiDataPlatformType != nil {
		s.D.Set("ai_data_platform_type", *s.Res.AiDataPlatformType)
	}

	if s.Res.AliasKey != nil {
		s.D.Set("alias_key", *s.Res.AliasKey)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
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

	if s.Res.WebSocketEndpoint != nil {
		s.D.Set("web_socket_endpoint", *s.Res.WebSocketEndpoint)
	}

	return nil
}

func AiDataPlatformSummaryToMap(obj oci_ai_data_platform.AiDataPlatformSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AiDataPlatformType != nil {
		result["ai_data_platform_type"] = string(*obj.AiDataPlatformType)
	}

	if obj.AliasKey != nil {
		result["alias_key"] = string(*obj.AliasKey)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.CreatedBy != nil {
		result["created_by"] = string(*obj.CreatedBy)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
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

func (s *AiDataPlatformAiDataPlatformResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_ai_data_platform.ChangeAiDataPlatformCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.AiDataPlatformId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_data_platform")

	response, err := s.Client.ChangeAiDataPlatformCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAiDataPlatformFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_data_platform"), oci_ai_data_platform.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
