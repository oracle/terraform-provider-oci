// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_vision

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_ai_vision "github.com/oracle/oci-go-sdk/v65/aivision"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AiVisionStreamGroupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createAiVisionStreamGroup,
		Read:     readAiVisionStreamGroup,
		Update:   updateAiVisionStreamGroup,
		Delete:   deleteAiVisionStreamGroup,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
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
			"is_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"stream_overlaps": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"overlapping_streams": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
					},
				},
			},
			"stream_source_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

func createAiVisionStreamGroup(d *schema.ResourceData, m interface{}) error {
	sync := &AiVisionStreamGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceVisionClient()

	return tfresource.CreateResource(d, sync)
}

func readAiVisionStreamGroup(d *schema.ResourceData, m interface{}) error {
	sync := &AiVisionStreamGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceVisionClient()

	return tfresource.ReadResource(sync)
}

func updateAiVisionStreamGroup(d *schema.ResourceData, m interface{}) error {
	sync := &AiVisionStreamGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceVisionClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteAiVisionStreamGroup(d *schema.ResourceData, m interface{}) error {
	sync := &AiVisionStreamGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceVisionClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type AiVisionStreamGroupResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_ai_vision.AIServiceVisionClient
	Res                    *oci_ai_vision.StreamGroup
	DisableNotFoundRetries bool
}

func (s *AiVisionStreamGroupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *AiVisionStreamGroupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_ai_vision.StreamGroupLifecycleStateCreating),
	}
}

func (s *AiVisionStreamGroupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_ai_vision.StreamGroupLifecycleStateActive),
	}
}

func (s *AiVisionStreamGroupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_ai_vision.StreamGroupLifecycleStateDeleting),
	}
}

func (s *AiVisionStreamGroupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_ai_vision.StreamGroupLifecycleStateDeleted),
	}
}

func (s *AiVisionStreamGroupResourceCrud) Create() error {
	request := oci_ai_vision.CreateStreamGroupRequest{}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if streamOverlaps, ok := s.D.GetOkExists("stream_overlaps"); ok {
		interfaces := streamOverlaps.([]interface{})
		tmp := make([]oci_ai_vision.StreamGroupOverlap, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "stream_overlaps", stateDataIndex)
			converted, err := s.mapToStreamGroupOverlap(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("stream_overlaps") {
			request.StreamOverlaps = tmp
		}
	}

	if streamSourceIds, ok := s.D.GetOkExists("stream_source_ids"); ok {
		interfaces := streamSourceIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("stream_source_ids") {
			request.StreamSourceIds = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision")

	response, err := s.Client.CreateStreamGroup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getStreamGroupFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision"), oci_ai_vision.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *AiVisionStreamGroupResourceCrud) getStreamGroupFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_ai_vision.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	streamGroupId, err := streamGroupWaitForWorkRequest(workId, "streamgroup",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, streamGroupId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_ai_vision.CancelWorkRequestRequest{
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
	s.D.SetId(*streamGroupId)

	return s.Get()
}

func streamGroupWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "ai_vision", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_ai_vision.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func streamGroupWaitForWorkRequest(wId *string, entityType string, action oci_ai_vision.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_ai_vision.AIServiceVisionClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "ai_vision")
	retryPolicy.ShouldRetryOperation = streamGroupWorkRequestShouldRetryFunc(timeout)

	response := oci_ai_vision.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_ai_vision.OperationStatusInProgress),
			string(oci_ai_vision.OperationStatusAccepted),
			string(oci_ai_vision.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_ai_vision.OperationStatusSucceeded),
			string(oci_ai_vision.OperationStatusFailed),
			string(oci_ai_vision.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_ai_vision.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_ai_vision.OperationStatusFailed || response.Status == oci_ai_vision.OperationStatusCanceled {
		return nil, getErrorFromAiVisionStreamGroupWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromAiVisionStreamGroupWorkRequest(client *oci_ai_vision.AIServiceVisionClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_ai_vision.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_ai_vision.ListWorkRequestErrorsRequest{
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

func (s *AiVisionStreamGroupResourceCrud) Get() error {
	request := oci_ai_vision.GetStreamGroupRequest{}

	tmp := s.D.Id()
	request.StreamGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision")

	response, err := s.Client.GetStreamGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.StreamGroup
	return nil
}

func (s *AiVisionStreamGroupResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_ai_vision.UpdateStreamGroupRequest{}

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

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	tmp := s.D.Id()
	request.StreamGroupId = &tmp

	if streamOverlaps, ok := s.D.GetOkExists("stream_overlaps"); ok {
		interfaces := streamOverlaps.([]interface{})
		tmp := make([]oci_ai_vision.StreamGroupOverlap, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "stream_overlaps", stateDataIndex)
			converted, err := s.mapToStreamGroupOverlap(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("stream_overlaps") {
			request.StreamOverlaps = tmp
		}
	}

	if streamSourceIds, ok := s.D.GetOkExists("stream_source_ids"); ok {
		interfaces := streamSourceIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("stream_source_ids") {
			request.StreamSourceIds = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision")

	response, err := s.Client.UpdateStreamGroup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getStreamGroupFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision"), oci_ai_vision.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *AiVisionStreamGroupResourceCrud) Delete() error {
	request := oci_ai_vision.DeleteStreamGroupRequest{}

	tmp := s.D.Id()
	request.StreamGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision")

	response, err := s.Client.DeleteStreamGroup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := streamGroupWaitForWorkRequest(workId, "streamgroup",
		oci_ai_vision.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *AiVisionStreamGroupResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	s.D.Set("state", s.Res.LifecycleState)

	streamOverlaps := []interface{}{}
	for _, item := range s.Res.StreamOverlaps {
		streamOverlaps = append(streamOverlaps, StreamGroupOverlapToMap(item))
	}
	s.D.Set("stream_overlaps", streamOverlaps)

	s.D.Set("stream_source_ids", s.Res.StreamSourceIds)

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

func (s *AiVisionStreamGroupResourceCrud) mapToStreamGroupOverlap(fieldKeyFormat string) (oci_ai_vision.StreamGroupOverlap, error) {
	result := oci_ai_vision.StreamGroupOverlap{}

	if overlappingStreams, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "overlapping_streams")); ok {
		interfaces := overlappingStreams.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "overlapping_streams")) {
			result.OverlappingStreams = tmp
		}
	}

	return result, nil
}

func StreamGroupOverlapToMap(obj oci_ai_vision.StreamGroupOverlap) map[string]interface{} {
	result := map[string]interface{}{}

	result["overlapping_streams"] = obj.OverlappingStreams

	return result
}

func StreamGroupSummaryToMap(obj oci_ai_vision.StreamGroupSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
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

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	result["state"] = string(obj.LifecycleState)

	streamOverlaps := []interface{}{}
	for _, item := range obj.StreamOverlaps {
		streamOverlaps = append(streamOverlaps, StreamGroupOverlapToMap(item))
	}
	result["stream_overlaps"] = streamOverlaps

	result["stream_source_ids"] = obj.StreamSourceIds

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

func (s *AiVisionStreamGroupResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_ai_vision.ChangeStreamGroupCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.StreamGroupId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision")

	response, err := s.Client.ChangeStreamGroupCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getStreamGroupFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision"), oci_ai_vision.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
