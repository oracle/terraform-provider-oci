// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_anomaly_detection

import (
	"context"
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_ai_anomaly_detection "github.com/oracle/oci-go-sdk/v56/aianomalydetection"
	oci_common "github.com/oracle/oci-go-sdk/v56/common"
)

func AiAnomalyDetectionModelResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createAiAnomalyDetectionModel,
		Read:     readAiAnomalyDetectionModel,
		Update:   updateAiAnomalyDetectionModel,
		Delete:   deleteAiAnomalyDetectionModel,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"model_training_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"data_asset_ids": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Optional
						"target_fap": {
							Type:             schema.TypeFloat,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.AdDiffSuppress,
						},
						"training_fraction": {
							Type:             schema.TypeFloat,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.AdDiffSuppress,
						},

						// Computed
					},
				},
			},
			"project_id": {
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

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"model_training_results": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"fap": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"is_training_goal_achieved": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"multivariate_fap": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"row_reduction_details": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"is_reduction_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"reduction_method": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"reduction_percentage": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},
						"signal_details": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"fap": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"is_quantized": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"max": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"min": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"mvi_ratio": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"signal_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"std": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},
						"warning": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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

func createAiAnomalyDetectionModel(d *schema.ResourceData, m interface{}) error {
	sync := &AiAnomalyDetectionModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnomalyDetectionClient()

	return tfresource.CreateResource(d, sync)
}

func readAiAnomalyDetectionModel(d *schema.ResourceData, m interface{}) error {
	sync := &AiAnomalyDetectionModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnomalyDetectionClient()

	return tfresource.ReadResource(sync)
}

func updateAiAnomalyDetectionModel(d *schema.ResourceData, m interface{}) error {
	sync := &AiAnomalyDetectionModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnomalyDetectionClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteAiAnomalyDetectionModel(d *schema.ResourceData, m interface{}) error {
	sync := &AiAnomalyDetectionModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnomalyDetectionClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type AiAnomalyDetectionModelResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_ai_anomaly_detection.AnomalyDetectionClient
	Res                    *oci_ai_anomaly_detection.Model
	DisableNotFoundRetries bool
}

func (s *AiAnomalyDetectionModelResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *AiAnomalyDetectionModelResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_ai_anomaly_detection.ModelLifecycleStateCreating),
	}
}

func (s *AiAnomalyDetectionModelResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_ai_anomaly_detection.ModelLifecycleStateActive),
	}
}

func (s *AiAnomalyDetectionModelResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_ai_anomaly_detection.ModelLifecycleStateDeleting),
	}
}

func (s *AiAnomalyDetectionModelResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_ai_anomaly_detection.ModelLifecycleStateDeleted),
	}
}

func (s *AiAnomalyDetectionModelResourceCrud) Create() error {
	request := oci_ai_anomaly_detection.CreateModelRequest{}

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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if modelTrainingDetails, ok := s.D.GetOkExists("model_training_details"); ok {
		if tmpList := modelTrainingDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "model_training_details", 0)
			tmp, err := s.mapToModelTrainingDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ModelTrainingDetails = &tmp
		}
	}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_anomaly_detection")

	response, err := s.Client.CreateModel(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getModelFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_anomaly_detection"), oci_ai_anomaly_detection.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *AiAnomalyDetectionModelResourceCrud) getModelFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_ai_anomaly_detection.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	modelId, err := modelWaitForWorkRequest(workId, "model",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, modelId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_ai_anomaly_detection.CancelWorkRequestRequest{
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
	s.D.SetId(*modelId)

	return s.Get()
}

func modelWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "ai_anomaly_detection", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_ai_anomaly_detection.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func modelWaitForWorkRequest(wId *string, entityType string, action oci_ai_anomaly_detection.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_ai_anomaly_detection.AnomalyDetectionClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "ai_anomaly_detection")
	retryPolicy.ShouldRetryOperation = modelWorkRequestShouldRetryFunc(timeout)

	response := oci_ai_anomaly_detection.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_ai_anomaly_detection.OperationStatusInProgress),
			string(oci_ai_anomaly_detection.OperationStatusAccepted),
			string(oci_ai_anomaly_detection.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_ai_anomaly_detection.ActionTypeCreated),
			string(oci_ai_anomaly_detection.OperationStatusSucceeded),
			string(oci_ai_anomaly_detection.OperationStatusFailed),
			string(oci_ai_anomaly_detection.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_ai_anomaly_detection.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_ai_anomaly_detection.OperationStatusFailed || response.Status == oci_ai_anomaly_detection.OperationStatusCanceled {
		return nil, getErrorFromAiAnomalyDetectionModelWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromAiAnomalyDetectionModelWorkRequest(client *oci_ai_anomaly_detection.AnomalyDetectionClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_ai_anomaly_detection.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_ai_anomaly_detection.ListWorkRequestErrorsRequest{
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

func (s *AiAnomalyDetectionModelResourceCrud) Get() error {
	request := oci_ai_anomaly_detection.GetModelRequest{}

	tmp := s.D.Id()
	request.ModelId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_anomaly_detection")

	response, err := s.Client.GetModel(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Model
	return nil
}

func (s *AiAnomalyDetectionModelResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_ai_anomaly_detection.UpdateModelRequest{}

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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.ModelId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_anomaly_detection")

	response, err := s.Client.UpdateModel(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getModelFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_anomaly_detection"), oci_ai_anomaly_detection.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *AiAnomalyDetectionModelResourceCrud) Delete() error {
	request := oci_ai_anomaly_detection.DeleteModelRequest{}

	tmp := s.D.Id()
	request.ModelId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_anomaly_detection")

	response, err := s.Client.DeleteModel(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := modelWaitForWorkRequest(workId, "model",
		oci_ai_anomaly_detection.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *AiAnomalyDetectionModelResourceCrud) SetData() error {
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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ModelTrainingDetails != nil {
		s.D.Set("model_training_details", []interface{}{ModelTrainingDetailsToMap(s.Res.ModelTrainingDetails)})
	} else {
		s.D.Set("model_training_details", nil)
	}

	if s.Res.ModelTrainingResults != nil {
		s.D.Set("model_training_results", []interface{}{ModelTrainingResultsToMap(s.Res.ModelTrainingResults)})
	} else {
		s.D.Set("model_training_results", nil)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
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

func ModelSummaryToMap(obj oci_ai_anomaly_detection.ModelSummary) map[string]interface{} {
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

	if obj.ModelTrainingDetails != nil {
		result["model_training_details"] = []interface{}{ModelTrainingDetailsToMap(obj.ModelTrainingDetails)}
	}

	if obj.ModelTrainingResults != nil {
		result["model_training_results"] = []interface{}{ModelTrainingResultsToMap(obj.ModelTrainingResults)}
	}

	if obj.ProjectId != nil {
		result["project_id"] = string(*obj.ProjectId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}

func (s *AiAnomalyDetectionModelResourceCrud) mapToModelTrainingDetails(fieldKeyFormat string) (oci_ai_anomaly_detection.ModelTrainingDetails, error) {
	result := oci_ai_anomaly_detection.ModelTrainingDetails{}

	if dataAssetIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_asset_ids")); ok {
		interfaces := dataAssetIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "data_asset_ids")) {
			result.DataAssetIds = tmp
		}
	}

	if targetFap, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_fap")); ok {
		tmp := targetFap.(float64)
		r64 := math.Round(tmp*100) / 100
		f32 := float32(r64)
		result.TargetFap = &f32
	}

	if trainingFraction, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "training_fraction")); ok {
		tmp := trainingFraction.(float64)
		r64 := math.Round(tmp*100) / 100
		f32 := float32(r64)
		result.TrainingFraction = &f32
	}

	return result, nil
}

func ModelTrainingDetailsToMap(obj *oci_ai_anomaly_detection.ModelTrainingDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["data_asset_ids"] = obj.DataAssetIds

	if obj.TargetFap != nil {
		result["target_fap"] = float32(*obj.TargetFap)
	}

	if obj.TrainingFraction != nil {
		result["training_fraction"] = float32(*obj.TrainingFraction)
	}

	return result
}

func ModelTrainingResultsToMap(obj *oci_ai_anomaly_detection.ModelTrainingResults) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Fap != nil {
		result["fap"] = float32(*obj.Fap)
	}

	if obj.IsTrainingGoalAchieved != nil {
		result["is_training_goal_achieved"] = bool(*obj.IsTrainingGoalAchieved)
	}

	if obj.MultivariateFap != nil {
		result["multivariate_fap"] = float32(*obj.MultivariateFap)
	}

	if obj.RowReductionDetails != nil {
		result["row_reduction_details"] = []interface{}{RowReductionDetailsToMap(obj.RowReductionDetails)}
	}

	signalDetails := []interface{}{}
	for _, item := range obj.SignalDetails {
		signalDetails = append(signalDetails, PerSignalDetailsToMap(item))
	}
	result["signal_details"] = signalDetails

	if obj.Warning != nil {
		result["warning"] = string(*obj.Warning)
	}

	return result
}

func PerSignalDetailsToMap(obj oci_ai_anomaly_detection.PerSignalDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Details != nil {
		result["details"] = string(*obj.Details)
	}

	if obj.Fap != nil {
		result["fap"] = float32(*obj.Fap)
	}

	if obj.IsQuantized != nil {
		result["is_quantized"] = bool(*obj.IsQuantized)
	}

	if obj.Max != nil {
		result["max"] = float64(*obj.Max)
	}

	if obj.Min != nil {
		result["min"] = float64(*obj.Min)
	}

	if obj.MviRatio != nil {
		result["mvi_ratio"] = float64(*obj.MviRatio)
	}

	if obj.SignalName != nil {
		result["signal_name"] = string(*obj.SignalName)
	}

	result["status"] = string(obj.Status)

	if obj.Std != nil {
		result["std"] = float64(*obj.Std)
	}

	return result
}

func RowReductionDetailsToMap(obj *oci_ai_anomaly_detection.RowReductionDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsReductionEnabled != nil {
		result["is_reduction_enabled"] = bool(*obj.IsReductionEnabled)
	}

	result["reduction_method"] = string(obj.ReductionMethod)

	if obj.ReductionPercentage != nil {
		result["reduction_percentage"] = float64(*obj.ReductionPercentage)
	}

	return result
}

func (s *AiAnomalyDetectionModelResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_ai_anomaly_detection.ChangeModelCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ModelId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_anomaly_detection")

	_, err := s.Client.ChangeModelCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
