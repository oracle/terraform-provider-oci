// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_vision

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_ai_vision "github.com/oracle/oci-go-sdk/v65/aivision"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AiVisionModelResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: &tfresource.OneHour,
			Update: &tfresource.TwentyMinutes,
			Delete: &tfresource.TwentyMinutes,
		},
		Create: createAiVisionModel,
		Read:   readAiVisionModel,
		Update: updateAiVisionModel,
		Delete: deleteAiVisionModel,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"model_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"project_id": {
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
						"dataset_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"DATA_SCIENCE_LABELING",
								"OBJECT_STORAGE",
							}, true),
						},

						// Optional
						"bucket": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"dataset_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"namespace_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"object": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
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
			"is_quick_mode": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"max_training_duration_in_hours": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"model_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"testing_dataset": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"dataset_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"DATA_SCIENCE_LABELING",
								"OBJECT_STORAGE",
							}, true),
						},

						// Optional
						"bucket": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"dataset_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"namespace_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"object": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"validation_dataset": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"dataset_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"DATA_SCIENCE_LABELING",
								"OBJECT_STORAGE",
							}, true),
						},

						// Optional
						"bucket": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"dataset_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"namespace_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"object": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},

			// Computed
			"average_precision": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"confidence_threshold": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"metrics": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"precision": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"recall": {
				Type:     schema.TypeFloat,
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
			"test_image_count": {
				Type:     schema.TypeInt,
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
			"total_image_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"trained_duration_in_hours": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
		},
	}
}

func createAiVisionModel(d *schema.ResourceData, m interface{}) error {
	sync := &AiVisionModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceVisionClient()

	return tfresource.CreateResource(d, sync)
}

func readAiVisionModel(d *schema.ResourceData, m interface{}) error {
	sync := &AiVisionModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceVisionClient()

	return tfresource.ReadResource(sync)
}

func updateAiVisionModel(d *schema.ResourceData, m interface{}) error {
	sync := &AiVisionModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceVisionClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteAiVisionModel(d *schema.ResourceData, m interface{}) error {
	sync := &AiVisionModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceVisionClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type AiVisionModelResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_ai_vision.AIServiceVisionClient
	Res                    *oci_ai_vision.Model
	DisableNotFoundRetries bool
}

func (s *AiVisionModelResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *AiVisionModelResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_ai_vision.ModelLifecycleStateCreating),
	}
}

func (s *AiVisionModelResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_ai_vision.ModelLifecycleStateActive),
	}
}

func (s *AiVisionModelResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_ai_vision.ModelLifecycleStateDeleting),
	}
}

func (s *AiVisionModelResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_ai_vision.ModelLifecycleStateDeleted),
	}
}

func (s *AiVisionModelResourceCrud) Create() error {
	request := oci_ai_vision.CreateModelRequest{}

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

	if isQuickMode, ok := s.D.GetOkExists("is_quick_mode"); ok {
		tmp := isQuickMode.(bool)
		request.IsQuickMode = &tmp
	}

	if maxTrainingDurationInHours, ok := s.D.GetOkExists("max_training_duration_in_hours"); ok {
		tmp := maxTrainingDurationInHours.(float64)
		request.MaxTrainingDurationInHours = &tmp
	}

	if modelType, ok := s.D.GetOkExists("model_type"); ok {
		request.ModelType = oci_ai_vision.ModelModelTypeEnum(modelType.(string))
	}

	if modelVersion, ok := s.D.GetOkExists("model_version"); ok {
		tmp := modelVersion.(string)
		request.ModelVersion = &tmp
	}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	if testingDataset, ok := s.D.GetOkExists("testing_dataset"); ok {
		if tmpList := testingDataset.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "testing_dataset", 0)
			tmp, err := s.mapToDataset(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TestingDataset = tmp
		}
	}

	if trainingDataset, ok := s.D.GetOkExists("training_dataset"); ok {
		if tmpList := trainingDataset.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "training_dataset", 0)
			tmp, err := s.mapToDataset(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TrainingDataset = tmp
		}
	}

	if validationDataset, ok := s.D.GetOkExists("validation_dataset"); ok {
		if tmpList := validationDataset.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "validation_dataset", 0)
			tmp, err := s.mapToDataset(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ValidationDataset = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision")

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
	return s.getModelFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision"), oci_ai_vision.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *AiVisionModelResourceCrud) getModelFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_ai_vision.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	modelId, err := modelWaitForWorkRequest(workId, "model",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, modelId)
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

func modelWaitForWorkRequest(wId *string, entityType string, action oci_ai_vision.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_ai_vision.AIServiceVisionClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "ai_vision")
	retryPolicy.ShouldRetryOperation = modelWorkRequestShouldRetryFunc(timeout)

	response := oci_ai_vision.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
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
		return nil, getErrorFromAiVisionModelWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromAiVisionModelWorkRequest(client *oci_ai_vision.AIServiceVisionClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_ai_vision.ActionTypeEnum) error {
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

func (s *AiVisionModelResourceCrud) Get() error {
	request := oci_ai_vision.GetModelRequest{}

	tmp := s.D.Id()
	request.ModelId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision")

	response, err := s.Client.GetModel(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Model
	return nil
}

func (s *AiVisionModelResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_ai_vision.UpdateModelRequest{}

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision")

	response, err := s.Client.UpdateModel(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getModelFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision"), oci_ai_vision.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *AiVisionModelResourceCrud) Delete() error {
	request := oci_ai_vision.DeleteModelRequest{}

	tmp := s.D.Id()
	request.ModelId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision")

	response, err := s.Client.DeleteModel(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := modelWaitForWorkRequest(workId, "model",
		oci_ai_vision.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *AiVisionModelResourceCrud) SetData() error {
	if s.Res.AveragePrecision != nil {
		s.D.Set("average_precision", *s.Res.AveragePrecision)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConfidenceThreshold != nil {
		s.D.Set("confidence_threshold", *s.Res.ConfidenceThreshold)
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

	if s.Res.IsQuickMode != nil {
		s.D.Set("is_quick_mode", *s.Res.IsQuickMode)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MaxTrainingDurationInHours != nil {
		s.D.Set("max_training_duration_in_hours", *s.Res.MaxTrainingDurationInHours)
	}

	if s.Res.Metrics != nil {
		s.D.Set("metrics", *s.Res.Metrics)
	}

	s.D.Set("model_type", s.Res.ModelType)

	if s.Res.ModelVersion != nil {
		s.D.Set("model_version", *s.Res.ModelVersion)
	}

	if s.Res.Precision != nil {
		s.D.Set("precision", *s.Res.Precision)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	if s.Res.Recall != nil {
		s.D.Set("recall", *s.Res.Recall)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TestImageCount != nil {
		s.D.Set("test_image_count", *s.Res.TestImageCount)
	}

	if s.Res.TestingDataset != nil {
		testingDatasetArray := []interface{}{}
		if testingDatasetMap := DatasetToMap(&s.Res.TestingDataset); testingDatasetMap != nil {
			testingDatasetArray = append(testingDatasetArray, testingDatasetMap)
		}
		s.D.Set("testing_dataset", testingDatasetArray)
	} else {
		s.D.Set("testing_dataset", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TotalImageCount != nil {
		s.D.Set("total_image_count", *s.Res.TotalImageCount)
	}

	if s.Res.TrainedDurationInHours != nil {
		s.D.Set("trained_duration_in_hours", *s.Res.TrainedDurationInHours)
	}

	if s.Res.TrainingDataset != nil {
		trainingDatasetArray := []interface{}{}
		if trainingDatasetMap := DatasetToMap(&s.Res.TrainingDataset); trainingDatasetMap != nil {
			trainingDatasetArray = append(trainingDatasetArray, trainingDatasetMap)
		}
		s.D.Set("training_dataset", trainingDatasetArray)
	} else {
		s.D.Set("training_dataset", nil)
	}

	if s.Res.ValidationDataset != nil {
		validationDatasetArray := []interface{}{}
		if validationDatasetMap := DatasetToMap(&s.Res.ValidationDataset); validationDatasetMap != nil {
			validationDatasetArray = append(validationDatasetArray, validationDatasetMap)
		}
		s.D.Set("validation_dataset", validationDatasetArray)
	} else {
		s.D.Set("validation_dataset", nil)
	}

	return nil
}

func (s *AiVisionModelResourceCrud) mapToDataset(fieldKeyFormat string) (oci_ai_vision.Dataset, error) {
	var baseObject oci_ai_vision.Dataset
	//discriminator
	datasetTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dataset_type"))
	var datasetType string
	if ok {
		datasetType = datasetTypeRaw.(string)
	} else {
		datasetType = "" // default value
	}
	switch strings.ToLower(datasetType) {
	case strings.ToLower("DATA_SCIENCE_LABELING"):
		details := oci_ai_vision.DataScienceLabelingDataset{}
		if datasetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dataset_id")); ok {
			tmp := datasetId.(string)
			details.DatasetId = &tmp
		}
		baseObject = details
	case strings.ToLower("OBJECT_STORAGE"):
		details := oci_ai_vision.ObjectStorageDataset{}
		if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.BucketName = &tmp
		}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace_name")); ok {
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

func DatasetToMap(obj *oci_ai_vision.Dataset) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_ai_vision.DataScienceLabelingDataset:
		result["dataset_type"] = "DATA_SCIENCE_LABELING"

		if v.DatasetId != nil {
			result["dataset_id"] = string(*v.DatasetId)
		}
	case oci_ai_vision.ObjectStorageDataset:
		result["dataset_type"] = "OBJECT_STORAGE"

		if v.BucketName != nil {
			result["bucket"] = string(*v.BucketName)
		}

		if v.NamespaceName != nil {
			result["namespace_name"] = string(*v.NamespaceName)
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

func ModelSummaryToMap(obj oci_ai_vision.ModelSummary) map[string]interface{} {
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

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["model_type"] = string(obj.ModelType)

	if obj.ModelVersion != nil {
		result["model_version"] = string(*obj.ModelVersion)
	}

	if obj.Precision != nil {
		result["precision"] = float32(*obj.Precision)
	}

	if obj.ProjectId != nil {
		result["project_id"] = string(*obj.ProjectId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TestingDataset != nil {
		testingDatasetArray := []interface{}{}
		if testingDatasetMap := DatasetToMap(&obj.TestingDataset); testingDatasetMap != nil {
			testingDatasetArray = append(testingDatasetArray, testingDatasetMap)
		}
		result["testing_dataset"] = testingDatasetArray
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.TrainingDataset != nil {
		trainingDatasetArray := []interface{}{}
		if trainingDatasetMap := DatasetToMap(&obj.TrainingDataset); trainingDatasetMap != nil {
			trainingDatasetArray = append(trainingDatasetArray, trainingDatasetMap)
		}
		result["training_dataset"] = trainingDatasetArray
	}

	if obj.ValidationDataset != nil {
		validationDatasetArray := []interface{}{}
		if validationDatasetMap := DatasetToMap(&obj.ValidationDataset); validationDatasetMap != nil {
			validationDatasetArray = append(validationDatasetArray, validationDatasetMap)
		}
		result["validation_dataset"] = validationDatasetArray
	}

	return result
}

func (s *AiVisionModelResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_ai_vision.ChangeModelCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ModelId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_vision")

	_, err := s.Client.ChangeModelCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
