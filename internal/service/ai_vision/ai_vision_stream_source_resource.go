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

func AiVisionStreamSourceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createAiVisionStreamSource,
		Read:     readAiVisionStreamSource,
		Update:   updateAiVisionStreamSource,
		Delete:   deleteAiVisionStreamSource,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"stream_source_details": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"camera_url": {
							Type:     schema.TypeString,
							Required: true,
						},
						"source_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"RTSP",
							}, true),
						},
						"stream_network_access_details": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"private_endpoint_id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"stream_access_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"PRIVATE",
										}, true),
									},

									// Optional

									// Computed
								},
							},
						},

						// Optional
						"secret_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
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

func createAiVisionStreamSource(d *schema.ResourceData, m interface{}) error {
	sync := &AiVisionStreamSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceVisionClient()

	return tfresource.CreateResource(d, sync)
}

func readAiVisionStreamSource(d *schema.ResourceData, m interface{}) error {
	sync := &AiVisionStreamSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceVisionClient()

	return tfresource.ReadResource(sync)
}

func updateAiVisionStreamSource(d *schema.ResourceData, m interface{}) error {
	sync := &AiVisionStreamSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceVisionClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteAiVisionStreamSource(d *schema.ResourceData, m interface{}) error {
	sync := &AiVisionStreamSourceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceVisionClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type AiVisionStreamSourceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_ai_vision.AIServiceVisionClient
	Res                    *oci_ai_vision.StreamSource
	DisableNotFoundRetries bool
}

func (s *AiVisionStreamSourceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *AiVisionStreamSourceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_ai_vision.StreamSourceLifecycleStateCreating),
	}
}

func (s *AiVisionStreamSourceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_ai_vision.StreamSourceLifecycleStateActive),
	}
}

func (s *AiVisionStreamSourceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_ai_vision.StreamSourceLifecycleStateDeleting),
	}
}

func (s *AiVisionStreamSourceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_ai_vision.StreamSourceLifecycleStateDeleted),
	}
}

func (s *AiVisionStreamSourceResourceCrud) Create() error {
	request := oci_ai_vision.CreateStreamSourceRequest{}

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

	if streamSourceDetails, ok := s.D.GetOkExists("stream_source_details"); ok {
		if tmpList := streamSourceDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "stream_source_details", 0)
			tmp, err := s.mapToStreamSourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.StreamSourceDetails = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision")

	response, err := s.Client.CreateStreamSource(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getStreamSourceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision"), oci_ai_vision.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *AiVisionStreamSourceResourceCrud) getStreamSourceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_ai_vision.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	streamSourceId, err := streamSourceWaitForWorkRequest(workId, "streamsource",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, streamSourceId)
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
	s.D.SetId(*streamSourceId)

	return s.Get()
}

func streamSourceWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func streamSourceWaitForWorkRequest(wId *string, entityType string, action oci_ai_vision.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_ai_vision.AIServiceVisionClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "ai_vision")
	retryPolicy.ShouldRetryOperation = streamSourceWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromAiVisionStreamSourceWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromAiVisionStreamSourceWorkRequest(client *oci_ai_vision.AIServiceVisionClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_ai_vision.ActionTypeEnum) error {
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

func (s *AiVisionStreamSourceResourceCrud) Get() error {
	request := oci_ai_vision.GetStreamSourceRequest{}

	tmp := s.D.Id()
	request.StreamSourceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision")

	response, err := s.Client.GetStreamSource(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.StreamSource
	return nil
}

func (s *AiVisionStreamSourceResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_ai_vision.UpdateStreamSourceRequest{}

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

	if streamSourceDetails, ok := s.D.GetOkExists("stream_source_details"); ok {
		if tmpList := streamSourceDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "stream_source_details", 0)
			tmp, err := s.mapToStreamSourceDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.StreamSourceDetails = tmp
		}
	}

	tmp := s.D.Id()
	request.StreamSourceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision")

	response, err := s.Client.UpdateStreamSource(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getStreamSourceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision"), oci_ai_vision.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *AiVisionStreamSourceResourceCrud) Delete() error {
	request := oci_ai_vision.DeleteStreamSourceRequest{}

	tmp := s.D.Id()
	request.StreamSourceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision")

	response, err := s.Client.DeleteStreamSource(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := streamSourceWaitForWorkRequest(workId, "streamsource",
		oci_ai_vision.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *AiVisionStreamSourceResourceCrud) SetData() error {
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

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StreamSourceDetails != nil {
		streamSourceDetailsArray := []interface{}{}
		if streamSourceDetailsMap := StreamSourceDetailsToMap(&s.Res.StreamSourceDetails); streamSourceDetailsMap != nil {
			streamSourceDetailsArray = append(streamSourceDetailsArray, streamSourceDetailsMap)
		}
		s.D.Set("stream_source_details", streamSourceDetailsArray)
	} else {
		s.D.Set("stream_source_details", nil)
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

func (s *AiVisionStreamSourceResourceCrud) mapToStreamNetworkAccessDetails(fieldKeyFormat string) (oci_ai_vision.StreamNetworkAccessDetails, error) {
	var baseObject oci_ai_vision.StreamNetworkAccessDetails
	//discriminator
	streamAccessTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "stream_access_type"))
	var streamAccessType string
	if ok {
		streamAccessType = streamAccessTypeRaw.(string)
	} else {
		streamAccessType = "" // default value
	}
	switch strings.ToLower(streamAccessType) {
	case strings.ToLower("PRIVATE"):
		details := oci_ai_vision.PrivateStreamNetworkAccessDetails{}
		if privateEndpointId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_endpoint_id")); ok {
			tmp := privateEndpointId.(string)
			details.PrivateEndpointId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown stream_access_type '%v' was specified", streamAccessType)
	}
	return baseObject, nil
}

func StreamNetworkAccessDetailsToMap(obj *oci_ai_vision.StreamNetworkAccessDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_ai_vision.PrivateStreamNetworkAccessDetails:
		result["stream_access_type"] = "PRIVATE"

		if v.PrivateEndpointId != nil {
			result["private_endpoint_id"] = string(*v.PrivateEndpointId)
		}
	default:
		log.Printf("[WARN] Received 'stream_access_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *AiVisionStreamSourceResourceCrud) mapToStreamSourceDetails(fieldKeyFormat string) (oci_ai_vision.StreamSourceDetails, error) {
	var baseObject oci_ai_vision.StreamSourceDetails
	//discriminator
	sourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_type"))
	var sourceType string
	if ok {
		sourceType = sourceTypeRaw.(string)
	} else {
		sourceType = "" // default value
	}
	switch strings.ToLower(sourceType) {
	case strings.ToLower("RTSP"):
		details := oci_ai_vision.RtspSourceDetails{}
		if cameraUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "camera_url")); ok {
			tmp := cameraUrl.(string)
			details.CameraUrl = &tmp
		}
		if secretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_id")); ok {
			tmp := secretId.(string)
			details.SecretId = &tmp
		}
		if streamNetworkAccessDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "stream_network_access_details")); ok {
			if tmpList := streamNetworkAccessDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "stream_network_access_details"), 0)
				tmp, err := s.mapToStreamNetworkAccessDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert stream_network_access_details, encountered error: %v", err)
				}
				details.StreamNetworkAccessDetails = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown source_type '%v' was specified", sourceType)
	}
	return baseObject, nil
}

func StreamSourceDetailsToMap(obj *oci_ai_vision.StreamSourceDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_ai_vision.RtspSourceDetails:
		result["source_type"] = "RTSP"

		if v.CameraUrl != nil {
			result["camera_url"] = string(*v.CameraUrl)
		}

		if v.SecretId != nil {
			result["secret_id"] = string(*v.SecretId)
		}

		if v.StreamNetworkAccessDetails != nil {
			streamNetworkAccessDetailsArray := []interface{}{}
			if streamNetworkAccessDetailsMap := StreamNetworkAccessDetailsToMap(&v.StreamNetworkAccessDetails); streamNetworkAccessDetailsMap != nil {
				streamNetworkAccessDetailsArray = append(streamNetworkAccessDetailsArray, streamNetworkAccessDetailsMap)
			}
			result["stream_network_access_details"] = streamNetworkAccessDetailsArray
		}
	default:
		log.Printf("[WARN] Received 'source_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func StreamSourceSummaryToMap(obj oci_ai_vision.StreamSourceSummary) map[string]interface{} {
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

	result["state"] = string(obj.LifecycleState)

	if obj.StreamSourceDetails != nil {
		streamSourceDetailsArray := []interface{}{}
		if streamSourceDetailsMap := StreamSourceDetailsToMap(&obj.StreamSourceDetails); streamSourceDetailsMap != nil {
			streamSourceDetailsArray = append(streamSourceDetailsArray, streamSourceDetailsMap)
		}
		result["stream_source_details"] = streamSourceDetailsArray
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

func (s *AiVisionStreamSourceResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_ai_vision.ChangeStreamSourceCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.StreamSourceId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision")

	response, err := s.Client.ChangeStreamSourceCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getStreamSourceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision"), oci_ai_vision.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
