// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_generative_ai "github.com/oracle/oci-go-sdk/v65/generativeai"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GenerativeAiModelResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createGenerativeAiModel,
		Read:     readGenerativeAiModel,
		Update:   updateGenerativeAiModel,
		Delete:   deleteGenerativeAiModel,
		Schema: map[string]*schema.Schema{
			// Required
			"base_model_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"fine_tune_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"dedicated_ai_cluster_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"training_dataset": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"bucket": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"dataset_type": {
										Type:             schema.TypeString,
										Required:         true,
										ForceNew:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"OBJECT_STORAGE",
										}, true),
									},
									"namespace": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"object": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},

									// Optional

									// Computed
								},
							},
						},

						// Optional
						"training_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"training_config_type": {
										Type:             schema.TypeString,
										Required:         true,
										ForceNew:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"TFEW_TRAINING_CONFIG",
											"VANILLA_TRAINING_CONFIG",
										}, true),
									},

									// Optional
									"early_stopping_patience": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"early_stopping_threshold": {
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"learning_rate": {
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"log_model_metrics_interval_in_steps": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"num_of_last_layers": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"total_training_epochs": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"training_batch_size": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
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
			"vendor": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"capabilities": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"is_long_term_supported": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"model_metrics": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"final_accuracy": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"final_loss": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"model_metrics_type": {
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
			"time_deprecated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createGenerativeAiModel(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.CreateResource(d, sync)
}

func readGenerativeAiModel(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.ReadResource(sync)
}

func updateGenerativeAiModel(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteGenerativeAiModel(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type GenerativeAiModelResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_generative_ai.GenerativeAiClient
	Res                    *oci_generative_ai.Model
	DisableNotFoundRetries bool
}

func (s *GenerativeAiModelResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *GenerativeAiModelResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_generative_ai.ModelLifecycleStateCreating),
	}
}

func (s *GenerativeAiModelResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_generative_ai.ModelLifecycleStateActive),
	}
}

func (s *GenerativeAiModelResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_generative_ai.ModelLifecycleStateDeleting),
	}
}

func (s *GenerativeAiModelResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_generative_ai.ModelLifecycleStateDeleted),
	}
}

func (s *GenerativeAiModelResourceCrud) Create() error {
	request := oci_generative_ai.CreateModelRequest{}

	if baseModelId, ok := s.D.GetOkExists("base_model_id"); ok {
		tmp := baseModelId.(string)
		request.BaseModelId = &tmp
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

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if fineTuneDetails, ok := s.D.GetOkExists("fine_tune_details"); ok {
		if tmpList := fineTuneDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "fine_tune_details", 0)
			tmp, err := s.mapToFineTuneDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.FineTuneDetails = &tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if vendor, ok := s.D.GetOkExists("vendor"); ok {
		tmp := vendor.(string)
		request.Vendor = &tmp
	}

	if version, ok := s.D.GetOkExists("version"); ok {
		tmp := version.(string)
		request.Version = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.CreateModel(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getModelFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai"), oci_generative_ai.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *GenerativeAiModelResourceCrud) getModelFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_generative_ai.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	modelId, err := modelWaitForWorkRequest(workId, "model",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
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
		if tfresource.ShouldRetry(response, false, "generative_ai", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_generative_ai.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func modelWaitForWorkRequest(wId *string, entityType string, action oci_generative_ai.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_generative_ai.GenerativeAiClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "generative_ai")
	retryPolicy.ShouldRetryOperation = modelWorkRequestShouldRetryFunc(timeout)

	response := oci_generative_ai.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_generative_ai.OperationStatusInProgress),
			string(oci_generative_ai.OperationStatusAccepted),
			string(oci_generative_ai.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_generative_ai.OperationStatusSucceeded),
			string(oci_generative_ai.OperationStatusFailed),
			string(oci_generative_ai.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_generative_ai.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_generative_ai.OperationStatusFailed || response.Status == oci_generative_ai.OperationStatusCanceled {
		return nil, getErrorFromGenerativeAiModelWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromGenerativeAiModelWorkRequest(client *oci_generative_ai.GenerativeAiClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_generative_ai.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_generative_ai.ListWorkRequestErrorsRequest{
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

func (s *GenerativeAiModelResourceCrud) Get() error {
	request := oci_generative_ai.GetModelRequest{}

	tmp := s.D.Id()
	request.ModelId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.GetModel(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Model
	return nil
}

func (s *GenerativeAiModelResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_generative_ai.UpdateModelRequest{}

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

	tmp := s.D.Id()
	request.ModelId = &tmp

	if vendor, ok := s.D.GetOkExists("vendor"); ok {
		tmp := vendor.(string)
		request.Vendor = &tmp
	}

	if version, ok := s.D.GetOkExists("version"); ok {
		tmp := version.(string)
		request.Version = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.UpdateModel(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Model
	return nil
}

func (s *GenerativeAiModelResourceCrud) Delete() error {
	request := oci_generative_ai.DeleteModelRequest{}

	tmp := s.D.Id()
	request.ModelId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.DeleteModel(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := modelWaitForWorkRequest(workId, "model",
		oci_generative_ai.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *GenerativeAiModelResourceCrud) SetData() error {
	if s.Res.BaseModelId != nil {
		s.D.Set("base_model_id", *s.Res.BaseModelId)
	}

	s.D.Set("capabilities", s.Res.Capabilities)

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

	if s.Res.FineTuneDetails != nil {
		s.D.Set("fine_tune_details", []interface{}{FineTuneDetailsToMap(s.Res.FineTuneDetails)})
	} else {
		s.D.Set("fine_tune_details", nil)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsLongTermSupported != nil {
		s.D.Set("is_long_term_supported", *s.Res.IsLongTermSupported)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ModelMetrics != nil {
		modelMetricsArray := []interface{}{}
		if modelMetricsMap := ModelMetricsToMap(&s.Res.ModelMetrics); modelMetricsMap != nil {
			modelMetricsArray = append(modelMetricsArray, modelMetricsMap)
		}
		s.D.Set("model_metrics", modelMetricsArray)
	} else {
		s.D.Set("model_metrics", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeDeprecated != nil {
		s.D.Set("time_deprecated", s.Res.TimeDeprecated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("type", s.Res.Type)

	if s.Res.Vendor != nil {
		s.D.Set("vendor", *s.Res.Vendor)
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}

func (s *GenerativeAiModelResourceCrud) mapToDataset(fieldKeyFormat string) (oci_generative_ai.Dataset, error) {
	var baseObject oci_generative_ai.Dataset
	//discriminator
	datasetTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dataset_type"))
	var datasetType string
	if ok {
		datasetType = datasetTypeRaw.(string)
	} else {
		datasetType = "" // default value
	}
	switch strings.ToLower(datasetType) {
	case strings.ToLower("OBJECT_STORAGE"):
		details := oci_generative_ai.ObjectStorageDataset{}
		if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.BucketName = &tmp
		}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.NamespaceName = &tmp
		}
		if object, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object")); ok {
			tmp := object.(string)
			details.ObjectName = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown dataset_type '%v' was specified", datasetType)
	}
	return baseObject, nil
}

func DatasetToMap(obj *oci_generative_ai.Dataset) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_generative_ai.ObjectStorageDataset:
		result["dataset_type"] = "OBJECT_STORAGE"

		if v.BucketName != nil {
			result["bucket"] = string(*v.BucketName)
		}

		if v.NamespaceName != nil {
			result["namespace"] = string(*v.NamespaceName)
		}

		if v.ObjectName != nil {
			result["object"] = string(*v.ObjectName)
		}
	default:
		log.Printf("[WARN] Received 'dataset_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GenerativeAiModelResourceCrud) mapToFineTuneDetails(fieldKeyFormat string) (oci_generative_ai.FineTuneDetails, error) {
	result := oci_generative_ai.FineTuneDetails{}

	if dedicatedAiClusterId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dedicated_ai_cluster_id")); ok {
		tmp := dedicatedAiClusterId.(string)
		result.DedicatedAiClusterId = &tmp
	}

	if trainingConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "training_config")); ok {
		if tmpList := trainingConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "training_config"), 0)
			tmp, err := s.mapToTrainingConfig(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert training_config, encountered error: %v", err)
			}
			result.TrainingConfig = tmp
		}
	}

	if trainingDataset, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "training_dataset")); ok {
		if tmpList := trainingDataset.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "training_dataset"), 0)
			tmp, err := s.mapToDataset(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert training_dataset, encountered error: %v", err)
			}
			result.TrainingDataset = tmp
		}
	}

	return result, nil
}

func FineTuneDetailsToMap(obj *oci_generative_ai.FineTuneDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DedicatedAiClusterId != nil {
		result["dedicated_ai_cluster_id"] = string(*obj.DedicatedAiClusterId)
	}

	if obj.TrainingConfig != nil {
		trainingConfigArray := []interface{}{}
		if trainingConfigMap := TrainingConfigToMap(&obj.TrainingConfig); trainingConfigMap != nil {
			trainingConfigArray = append(trainingConfigArray, trainingConfigMap)
		}
		result["training_config"] = trainingConfigArray
	}

	if obj.TrainingDataset != nil {
		trainingDatasetArray := []interface{}{}
		if trainingDatasetMap := DatasetToMap(&obj.TrainingDataset); trainingDatasetMap != nil {
			trainingDatasetArray = append(trainingDatasetArray, trainingDatasetMap)
		}
		result["training_dataset"] = trainingDatasetArray
	}

	return result
}

func ModelMetricsToMap(obj *oci_generative_ai.ModelMetrics) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_generative_ai.TextGenerationModelMetrics:
		result["model_metrics_type"] = "TEXT_GENERATION_MODEL_METRICS"

		if v.FinalAccuracy != nil {
			result["final_accuracy"] = float64(*v.FinalAccuracy)
		}

		if v.FinalLoss != nil {
			result["final_loss"] = float64(*v.FinalLoss)
		}
	default:
		log.Printf("[WARN] Received 'model_metrics_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func ModelSummaryToMap(obj oci_generative_ai.ModelSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BaseModelId != nil {
		result["base_model_id"] = string(*obj.BaseModelId)
	}

	result["capabilities"] = obj.Capabilities

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.FineTuneDetails != nil {
		result["fine_tune_details"] = []interface{}{FineTuneDetailsToMap(obj.FineTuneDetails)}
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsLongTermSupported != nil {
		result["is_long_term_supported"] = bool(*obj.IsLongTermSupported)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.ModelMetrics != nil {
		modelMetricsArray := []interface{}{}
		if modelMetricsMap := ModelMetricsToMap(&obj.ModelMetrics); modelMetricsMap != nil {
			modelMetricsArray = append(modelMetricsArray, modelMetricsMap)
		}
		result["model_metrics"] = modelMetricsArray
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeDeprecated != nil {
		result["time_deprecated"] = obj.TimeDeprecated.String()
	}

	result["type"] = string(obj.Type)

	if obj.Vendor != nil {
		result["vendor"] = string(*obj.Vendor)
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}

func (s *GenerativeAiModelResourceCrud) mapToTrainingConfig(fieldKeyFormat string) (oci_generative_ai.TrainingConfig, error) {
	var baseObject oci_generative_ai.TrainingConfig
	//discriminator
	trainingConfigTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "training_config_type"))
	var trainingConfigType string
	if ok {
		trainingConfigType = trainingConfigTypeRaw.(string)
	} else {
		trainingConfigType = "TFEW_TRAINING_CONFIG" // default value
	}
	switch strings.ToLower(trainingConfigType) {
	case strings.ToLower("TFEW_TRAINING_CONFIG"):
		details := oci_generative_ai.TFewTrainingConfig{}
		if earlyStoppingPatience, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "early_stopping_patience")); ok {
			tmp := earlyStoppingPatience.(int)
			details.EarlyStoppingPatience = &tmp
		}
		if earlyStoppingThreshold, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "early_stopping_threshold")); ok {
			tmp := earlyStoppingThreshold.(float64)
			details.EarlyStoppingThreshold = &tmp
		}
		if learningRate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "learning_rate")); ok {
			tmp := learningRate.(float64)
			details.LearningRate = &tmp
		}
		if logModelMetricsIntervalInSteps, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "log_model_metrics_interval_in_steps")); ok {
			tmp := logModelMetricsIntervalInSteps.(int)
			details.LogModelMetricsIntervalInSteps = &tmp
		}
		if totalTrainingEpochs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "total_training_epochs")); ok {
			tmp := totalTrainingEpochs.(int)
			details.TotalTrainingEpochs = &tmp
		}
		if trainingBatchSize, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "training_batch_size")); ok {
			tmp := trainingBatchSize.(int)
			details.TrainingBatchSize = &tmp
		}
		baseObject = details
	case strings.ToLower("VANILLA_TRAINING_CONFIG"):
		details := oci_generative_ai.VanillaTrainingConfig{}
		if numOfLastLayers, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "num_of_last_layers")); ok {
			tmp := numOfLastLayers.(int)
			details.NumOfLastLayers = &tmp
		}
		if earlyStoppingPatience, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "early_stopping_patience")); ok {
			tmp := earlyStoppingPatience.(int)
			details.EarlyStoppingPatience = &tmp
		}
		if earlyStoppingThreshold, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "early_stopping_threshold")); ok {
			tmp := earlyStoppingThreshold.(float64)
			details.EarlyStoppingThreshold = &tmp
		}
		if learningRate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "learning_rate")); ok {
			tmp := learningRate.(float64)
			details.LearningRate = &tmp
		}
		if logModelMetricsIntervalInSteps, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "log_model_metrics_interval_in_steps")); ok {
			tmp := logModelMetricsIntervalInSteps.(int)
			details.LogModelMetricsIntervalInSteps = &tmp
		}
		if totalTrainingEpochs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "total_training_epochs")); ok {
			tmp := totalTrainingEpochs.(int)
			details.TotalTrainingEpochs = &tmp
		}
		if trainingBatchSize, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "training_batch_size")); ok {
			tmp := trainingBatchSize.(int)
			details.TrainingBatchSize = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown training_config_type '%v' was specified", trainingConfigType)
	}
	return baseObject, nil
}

func TrainingConfigToMap(obj *oci_generative_ai.TrainingConfig) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_generative_ai.TFewTrainingConfig:
		result["training_config_type"] = "TFEW_TRAINING_CONFIG"

		if v.EarlyStoppingPatience != nil {
			result["early_stopping_patience"] = int(*v.EarlyStoppingPatience)
		}

		if v.EarlyStoppingThreshold != nil {
			result["early_stopping_threshold"] = float64(*v.EarlyStoppingThreshold)
		}

		if v.LearningRate != nil {
			result["learning_rate"] = float64(*v.LearningRate)
		}

		if v.LogModelMetricsIntervalInSteps != nil {
			result["log_model_metrics_interval_in_steps"] = int(*v.LogModelMetricsIntervalInSteps)
		}

		if v.TotalTrainingEpochs != nil {
			result["total_training_epochs"] = int(*v.TotalTrainingEpochs)
		}

		if v.TrainingBatchSize != nil {
			result["training_batch_size"] = int(*v.TrainingBatchSize)
		}
	case oci_generative_ai.VanillaTrainingConfig:
		result["training_config_type"] = "VANILLA_TRAINING_CONFIG"

		if v.NumOfLastLayers != nil {
			result["num_of_last_layers"] = int(*v.NumOfLastLayers)
		}

		if v.EarlyStoppingPatience != nil {
			result["early_stopping_patience"] = int(*v.EarlyStoppingPatience)
		}

		if v.EarlyStoppingThreshold != nil {
			result["early_stopping_threshold"] = float64(*v.EarlyStoppingThreshold)
		}

		if v.LearningRate != nil {
			result["learning_rate"] = float64(*v.LearningRate)
		}

		if v.LogModelMetricsIntervalInSteps != nil {
			result["log_model_metrics_interval_in_steps"] = int(*v.LogModelMetricsIntervalInSteps)
		}

		if v.TotalTrainingEpochs != nil {
			result["total_training_epochs"] = int(*v.TotalTrainingEpochs)
		}

		if v.TrainingBatchSize != nil {
			result["training_batch_size"] = int(*v.TrainingBatchSize)
		}
	default:
		log.Printf("[WARN] Received 'training_config_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GenerativeAiModelResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_generative_ai.ChangeModelCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ModelId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	_, err := s.Client.ChangeModelCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
