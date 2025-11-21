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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_ai_vision "github.com/oracle/oci-go-sdk/v65/aivision"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AiVisionStreamJobResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createAiVisionStreamJob,
		Read:     readAiVisionStreamJob,
		Update:   updateAiVisionStreamJob,
		Delete:   deleteAiVisionStreamJob,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"features": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"feature_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"FACE_DETECTION",
								"OBJECT_TRACKING",
							}, true),
						},

						// Optional
						"max_results": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"should_return_landmarks": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"tracking_types": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"biometric_store_compartment_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"biometric_store_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"detection_model_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"max_results": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"objects": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"should_return_landmarks": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"tracking_model_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"stream_output_location": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"bucket": {
							Type:     schema.TypeString,
							Required: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Required: true,
						},
						"output_location_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"OBJECT_STORAGE",
							}, true),
						},
						"prefix": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"obo_token": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"stream_source_id": {
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
			"state": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_ai_vision.StreamJobLifecycleStateInactive),
					string(oci_ai_vision.StreamJobLifecycleStateActive),
				}, true),
			},

			// Computed
			"agent_participant_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
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

func createAiVisionStreamJob(d *schema.ResourceData, m interface{}) error {
	sync := &AiVisionStreamJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceVisionClient()
	var powerOff = false
	if powerState, ok := sync.D.GetOkExists("state"); ok {
		wantedPowerState := oci_ai_vision.StreamJobLifecycleStateEnum(strings.ToUpper(powerState.(string)))
		if wantedPowerState == oci_ai_vision.StreamJobLifecycleStateInactive {
			powerOff = true
		}
	}

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if powerOff {
		if err := sync.StopStreamJob(); err != nil {
			return err
		}
		sync.D.Set("state", oci_ai_vision.StreamJobLifecycleStateInactive)
	}
	return nil

}

func readAiVisionStreamJob(d *schema.ResourceData, m interface{}) error {
	sync := &AiVisionStreamJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceVisionClient()

	return tfresource.ReadResource(sync)
}

func updateAiVisionStreamJob(d *schema.ResourceData, m interface{}) error {
	sync := &AiVisionStreamJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceVisionClient()

	powerOn, powerOff := false, false

	if sync.D.HasChange("state") {
		wantedState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_ai_vision.StreamJobLifecycleStateActive == oci_ai_vision.StreamJobLifecycleStateEnum(wantedState) {
			powerOn = true
		} else if oci_ai_vision.StreamJobLifecycleStateInactive == oci_ai_vision.StreamJobLifecycleStateEnum(wantedState) {
			powerOff = true
		}
	}

	if powerOn {
		if err := sync.StartStreamJob(); err != nil {
			return err
		}
		sync.D.Set("state", oci_ai_vision.StreamJobLifecycleStateActive)
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	if powerOff {
		if err := sync.StopStreamJob(); err != nil {
			return err
		}
		sync.D.Set("state", oci_ai_vision.StreamJobLifecycleStateInactive)
	}

	return nil
}

func deleteAiVisionStreamJob(d *schema.ResourceData, m interface{}) error {
	sync := &AiVisionStreamJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceVisionClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type AiVisionStreamJobResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_ai_vision.AIServiceVisionClient
	Res                    *oci_ai_vision.StreamJob
	DisableNotFoundRetries bool
}

func (s *AiVisionStreamJobResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *AiVisionStreamJobResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_ai_vision.StreamJobLifecycleStateCreating),
	}
}

func (s *AiVisionStreamJobResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_ai_vision.StreamJobLifecycleStateInactive),
		string(oci_ai_vision.StreamJobLifecycleStateNeedsAttention),
	}
}

func (s *AiVisionStreamJobResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_ai_vision.StreamJobLifecycleStateDeleting),
	}
}

func (s *AiVisionStreamJobResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_ai_vision.StreamJobLifecycleStateDeleted),
	}
}

func (s *AiVisionStreamJobResourceCrud) Create() error {
	request := oci_ai_vision.CreateStreamJobRequest{}

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

	if features, ok := s.D.GetOkExists("features"); ok {
		interfaces := features.([]interface{})
		tmp := make([]oci_ai_vision.VideoStreamFeature, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "features", stateDataIndex)
			converted, err := s.mapToVideoStreamFeature(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("features") {
			request.Features = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if streamOutputLocation, ok := s.D.GetOkExists("stream_output_location"); ok {
		if tmpList := streamOutputLocation.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "stream_output_location", 0)
			tmp, err := s.mapToStreamOutputLocation(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.StreamOutputLocation = tmp
		}
	}

	if streamSourceId, ok := s.D.GetOkExists("stream_source_id"); ok {
		tmp := streamSourceId.(string)
		request.StreamSourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision")

	response, err := s.Client.CreateStreamJob(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getStreamJobFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision"), oci_ai_vision.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *AiVisionStreamJobResourceCrud) getStreamJobFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_ai_vision.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	streamJobId, err := streamJobWaitForWorkRequest(workId, "streamjob",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, streamJobId)
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
	s.D.SetId(*streamJobId)

	return s.Get()
}

func streamJobWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func streamJobWaitForWorkRequest(wId *string, entityType string, action oci_ai_vision.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_ai_vision.AIServiceVisionClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "ai_vision")
	retryPolicy.ShouldRetryOperation = streamJobWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromAiVisionStreamJobWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromAiVisionStreamJobWorkRequest(client *oci_ai_vision.AIServiceVisionClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_ai_vision.ActionTypeEnum) error {
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

func (s *AiVisionStreamJobResourceCrud) Get() error {
	request := oci_ai_vision.GetStreamJobRequest{}

	tmp := s.D.Id()
	request.StreamJobId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision")

	response, err := s.Client.GetStreamJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.StreamJob
	return nil
}

func (s *AiVisionStreamJobResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_ai_vision.UpdateStreamJobRequest{}

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

	if features, ok := s.D.GetOkExists("features"); ok {
		interfaces := features.([]interface{})
		tmp := make([]oci_ai_vision.VideoStreamFeature, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "features", stateDataIndex)
			converted, err := s.mapToVideoStreamFeature(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("features") {
			request.Features = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.StreamJobId = &tmp

	if streamOutputLocation, ok := s.D.GetOkExists("stream_output_location"); ok {
		if tmpList := streamOutputLocation.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "stream_output_location", 0)
			tmp, err := s.mapToStreamOutputLocation(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.StreamOutputLocation = tmp
		}
	}

	if streamSourceId, ok := s.D.GetOkExists("stream_source_id"); ok {
		tmp := streamSourceId.(string)
		request.StreamSourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision")

	response, err := s.Client.UpdateStreamJob(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getStreamJobFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision"), oci_ai_vision.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *AiVisionStreamJobResourceCrud) Delete() error {
	request := oci_ai_vision.DeleteStreamJobRequest{}

	tmp := s.D.Id()
	request.StreamJobId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision")

	response, err := s.Client.DeleteStreamJob(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := streamJobWaitForWorkRequest(workId, "streamjob",
		oci_ai_vision.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *AiVisionStreamJobResourceCrud) SetData() error {
	if s.Res.AgentParticipantId != nil {
		s.D.Set("agent_participant_id", *s.Res.AgentParticipantId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	features := []interface{}{}
	for _, item := range s.Res.Features {
		features = append(features, VideoStreamFeatureToMap(item))
	}
	s.D.Set("features", features)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StreamOutputLocation != nil {
		streamOutputLocationArray := []interface{}{}
		if streamOutputLocationMap := StreamOutputLocationToMap(&s.Res.StreamOutputLocation); streamOutputLocationMap != nil {
			streamOutputLocationArray = append(streamOutputLocationArray, streamOutputLocationMap)
		}
		s.D.Set("stream_output_location", streamOutputLocationArray)
	} else {
		s.D.Set("stream_output_location", nil)
	}

	if s.Res.StreamSourceId != nil {
		s.D.Set("stream_source_id", *s.Res.StreamSourceId)
	}

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

func (s *AiVisionStreamJobResourceCrud) StartStreamJob() error {
	request := oci_ai_vision.StartStreamJobRequest{}

	idTmp := s.D.Id()
	request.StreamJobId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision")

	_, err := s.Client.StartStreamJob(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_ai_vision.StreamJobLifecycleStateActive }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *AiVisionStreamJobResourceCrud) StopStreamJob() error {
	request := oci_ai_vision.StopStreamJobRequest{}

	idTmp := s.D.Id()
	request.StreamJobId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision")

	_, err := s.Client.StopStreamJob(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_ai_vision.StreamJobLifecycleStateInactive }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func StreamJobSummaryToMap(obj oci_ai_vision.StreamJobSummary) map[string]interface{} {
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

	features := []interface{}{}
	for _, item := range obj.Features {
		features = append(features, VideoStreamFeatureToMap(item))
	}
	result["features"] = features

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.StreamOutputLocation != nil {
		streamOutputLocationArray := []interface{}{}
		if streamOutputLocationMap := StreamOutputLocationToMap(&obj.StreamOutputLocation); streamOutputLocationMap != nil {
			streamOutputLocationArray = append(streamOutputLocationArray, streamOutputLocationMap)
		}
		result["stream_output_location"] = streamOutputLocationArray
	}

	if obj.StreamSourceId != nil {
		result["stream_source_id"] = string(*obj.StreamSourceId)
	}

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

func (s *AiVisionStreamJobResourceCrud) mapToStreamOutputLocation(fieldKeyFormat string) (oci_ai_vision.StreamOutputLocation, error) {
	var baseObject oci_ai_vision.StreamOutputLocation
	//discriminator
	outputLocationTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "output_location_type"))
	var outputLocationType string
	if ok {
		outputLocationType = outputLocationTypeRaw.(string)
	} else {
		outputLocationType = "" // default value
	}
	switch strings.ToLower(outputLocationType) {
	case strings.ToLower("OBJECT_STORAGE"):
		details := oci_ai_vision.ObjectStorageOutputLocation{}
		if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.BucketName = &tmp
		}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.NamespaceName = &tmp
		}
		if prefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "prefix")); ok {
			tmp := prefix.(string)
			details.Prefix = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown output_location_type '%v' was specified", outputLocationType)
	}
	return baseObject, nil
}

func StreamOutputLocationToMap(obj *oci_ai_vision.StreamOutputLocation) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_ai_vision.ObjectStorageOutputLocation:
		result["output_location_type"] = "OBJECT_STORAGE"

		if v.BucketName != nil {
			result["bucket"] = string(*v.BucketName)
		}

		if v.NamespaceName != nil {
			result["namespace"] = string(*v.NamespaceName)
		}

		if v.Prefix != nil {
			result["prefix"] = string(*v.Prefix)
		}
	default:
		log.Printf("[WARN] Received 'output_location_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *AiVisionStreamJobResourceCrud) mapToTrackingType(fieldKeyFormat string) (oci_ai_vision.TrackingType, error) {
	result := oci_ai_vision.TrackingType{}

	if detectionModelId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "detection_model_id")); ok {
		tmp := detectionModelId.(string)
		result.DetectionModelId = &tmp
	}

	if maxResults, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_results")); ok {
		tmp := maxResults.(int)
		result.MaxResults = &tmp
	}

	if objects, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "objects")); ok {
		interfaces := objects.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "objects")) {
			result.Objects = tmp
		}
	}

	if shouldReturnLandmarks, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_return_landmarks")); ok {
		tmp := shouldReturnLandmarks.(bool)
		result.ShouldReturnLandmarks = &tmp
	}

	if trackingModelId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tracking_model_id")); ok {
		tmp := trackingModelId.(string)
		result.TrackingModelId = &tmp
	}

	return result, nil
}

func TrackingTypeToMap(obj oci_ai_vision.TrackingType) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DetectionModelId != nil {
		result["detection_model_id"] = string(*obj.DetectionModelId)
	}

	if obj.MaxResults != nil {
		result["max_results"] = int(*obj.MaxResults)
	}

	result["objects"] = obj.Objects

	if obj.ShouldReturnLandmarks != nil {
		result["should_return_landmarks"] = bool(*obj.ShouldReturnLandmarks)
	}

	if obj.TrackingModelId != nil {
		result["tracking_model_id"] = string(*obj.TrackingModelId)
	}

	return result
}

func (s *AiVisionStreamJobResourceCrud) mapToVideoStreamFeature(fieldKeyFormat string) (oci_ai_vision.VideoStreamFeature, error) {
	var baseObject oci_ai_vision.VideoStreamFeature
	//discriminator
	featureTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "feature_type"))
	var featureType string
	if ok {
		featureType = featureTypeRaw.(string)
	} else {
		featureType = "" // default value
	}
	switch strings.ToLower(featureType) {
	case strings.ToLower("FACE_DETECTION"):
		details := oci_ai_vision.VideoStreamFaceDetectionFeature{}
		if maxResults, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_results")); ok {
			tmp := maxResults.(int)
			details.MaxResults = &tmp
		}
		if shouldReturnLandmarks, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_return_landmarks")); ok {
			tmp := shouldReturnLandmarks.(bool)
			details.ShouldReturnLandmarks = &tmp
		}
		baseObject = details
	case strings.ToLower("OBJECT_TRACKING"):
		details := oci_ai_vision.VideoStreamObjectTrackingFeature{}
		if trackingTypes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tracking_types")); ok {
			interfaces := trackingTypes.([]interface{})
			tmp := make([]oci_ai_vision.TrackingType, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "tracking_types"), stateDataIndex)
				converted, err := s.mapToTrackingType(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "tracking_types")) {
				details.TrackingTypes = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown feature_type '%v' was specified", featureType)
	}
	return baseObject, nil
}

func VideoStreamFeatureToMap(obj oci_ai_vision.VideoStreamFeature) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_ai_vision.VideoStreamFaceDetectionFeature:
		result["feature_type"] = "FACE_DETECTION"

		if v.MaxResults != nil {
			result["max_results"] = int(*v.MaxResults)
		}

		if v.ShouldReturnLandmarks != nil {
			result["should_return_landmarks"] = bool(*v.ShouldReturnLandmarks)
		}
	case oci_ai_vision.VideoStreamObjectTrackingFeature:
		result["feature_type"] = "OBJECT_TRACKING"

		trackingTypes := []interface{}{}
		for _, item := range v.TrackingTypes {
			trackingTypes = append(trackingTypes, TrackingTypeToMap(item))
		}
		result["tracking_types"] = trackingTypes
	default:
		log.Printf("[WARN] Received 'feature_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *AiVisionStreamJobResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_ai_vision.ChangeStreamJobCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.StreamJobId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision")

	response, err := s.Client.ChangeStreamJobCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getStreamJobFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision"), oci_ai_vision.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
