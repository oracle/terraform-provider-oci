// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_document

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_ai_document "github.com/oracle/oci-go-sdk/v65/aidocument"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AiDocumentModelResource() *schema.Resource {
	var (
		FifteenMinutes = 15 * time.Minute
		TwentyMinutes  = 20 * time.Minute
		OneHour        = 60 * time.Minute
	)
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: &OneHour,
			Update: &FifteenMinutes,
			Delete: &TwentyMinutes,
		},
		Create: createAiDocumentModel,
		Read:   readAiDocumentModel,
		Update: updateAiDocumentModel,
		Delete: deleteAiDocumentModel,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"model_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
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

			// Optional
			"component_models": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"model_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
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
			"max_training_time_in_hours": {
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
						"namespace": {
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
			"training_dataset": {
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
						"namespace": {
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
						"namespace": {
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
			"is_composed_model": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"labels": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"metrics": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"dataset_summary": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"test_sample_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"training_sample_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"validation_sample_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"label_metrics_report": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"confidence_entries": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"accuracy": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"f1score": {
													Type:     schema.TypeFloat,
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
												"threshold": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
											},
										},
									},
									"document_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"label": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"mean_average_precision": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},
						"model_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"overall_metrics_report": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"confidence_entries": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"accuracy": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"f1score": {
													Type:     schema.TypeFloat,
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
												"threshold": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
											},
										},
									},
									"document_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"mean_average_precision": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
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
			"tenancy_id": {
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
			"trained_time_in_hours": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
		},
	}
}

func createAiDocumentModel(d *schema.ResourceData, m interface{}) error {
	sync := &AiDocumentModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceDocumentClient()

	return tfresource.CreateResource(d, sync)
}

func readAiDocumentModel(d *schema.ResourceData, m interface{}) error {
	sync := &AiDocumentModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceDocumentClient()

	return tfresource.ReadResource(sync)
}

func updateAiDocumentModel(d *schema.ResourceData, m interface{}) error {
	sync := &AiDocumentModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceDocumentClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteAiDocumentModel(d *schema.ResourceData, m interface{}) error {
	sync := &AiDocumentModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceDocumentClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type AiDocumentModelResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_ai_document.AIServiceDocumentClient
	Res                    *oci_ai_document.Model
	DisableNotFoundRetries bool
}

func (s *AiDocumentModelResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *AiDocumentModelResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_ai_document.ModelLifecycleStateCreating),
	}
}

func (s *AiDocumentModelResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_ai_document.ModelLifecycleStateActive),
	}
}

func (s *AiDocumentModelResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_ai_document.ModelLifecycleStateDeleting),
	}
}

func (s *AiDocumentModelResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_ai_document.ModelLifecycleStateDeleted),
	}
}

func (s *AiDocumentModelResourceCrud) Create() error {
	request := oci_ai_document.CreateModelRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if componentModels, ok := s.D.GetOkExists("component_models"); ok {
		interfaces := componentModels.([]interface{})
		tmp := make([]oci_ai_document.ComponentModel, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "component_models", stateDataIndex)
			converted, err := s.mapToComponentModel(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("component_models") {
			request.ComponentModels = tmp
		}
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

	if maxTrainingTimeInHours, ok := s.D.GetOkExists("max_training_time_in_hours"); ok {
		tmp := maxTrainingTimeInHours.(float64)
		request.MaxTrainingTimeInHours = &tmp
	}

	if modelType, ok := s.D.GetOkExists("model_type"); ok {
		request.ModelType = oci_ai_document.ModelModelTypeEnum(modelType.(string))
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_document")

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

	err = s.getModelFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_document"), oci_ai_document.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
	if err != nil {
		return err
	}

	return nil
}

func (s *AiDocumentModelResourceCrud) getModelFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_ai_document.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	modelId, err := modelWaitForWorkRequest(workId, "model",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, modelId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_ai_document.CancelWorkRequestRequest{
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
		if tfresource.ShouldRetry(response, false, "ai_document", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_ai_document.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func modelWaitForWorkRequest(wId *string, entityType string, action oci_ai_document.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_ai_document.AIServiceDocumentClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "ai_document")
	retryPolicy.ShouldRetryOperation = modelWorkRequestShouldRetryFunc(timeout)

	response := oci_ai_document.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_ai_document.OperationStatusInProgress),
			string(oci_ai_document.OperationStatusAccepted),
			string(oci_ai_document.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_ai_document.OperationStatusSucceeded),
			string(oci_ai_document.OperationStatusFailed),
			string(oci_ai_document.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_ai_document.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_ai_document.OperationStatusFailed || response.Status == oci_ai_document.OperationStatusCanceled {
		return nil, getErrorFromAiDocumentModelWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromAiDocumentModelWorkRequest(client *oci_ai_document.AIServiceDocumentClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_ai_document.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_ai_document.ListWorkRequestErrorsRequest{
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

func (s *AiDocumentModelResourceCrud) Get() error {
	request := oci_ai_document.GetModelRequest{}

	tmp := s.D.Id()
	request.ModelId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_document")

	response, err := s.Client.GetModel(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Model
	return nil
}

func (s *AiDocumentModelResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_ai_document.UpdateModelRequest{}

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_document")

	response, err := s.Client.UpdateModel(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getModelFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_document"), oci_ai_document.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}

	return nil
}

func (s *AiDocumentModelResourceCrud) Delete() error {
	request := oci_ai_document.DeleteModelRequest{}

	tmp := s.D.Id()
	request.ModelId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_document")

	response, err := s.Client.DeleteModel(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := modelWaitForWorkRequest(workId, "model",
		oci_ai_document.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *AiDocumentModelResourceCrud) SetData() error {

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	componentModels := []interface{}{}
	for _, item := range s.Res.ComponentModels {
		componentModels = append(componentModels, ComponentModelToMap(item))
	}
	s.D.Set("component_models", componentModels)

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
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsComposedModel != nil {
		s.D.Set("is_composed_model", *s.Res.IsComposedModel)
	}

	if s.Res.IsQuickMode != nil {
		s.D.Set("is_quick_mode", *s.Res.IsQuickMode)
	}

	s.D.Set("labels", s.Res.Labels)
	s.D.Set("labels", s.Res.Labels)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MaxTrainingTimeInHours != nil {
		s.D.Set("max_training_time_in_hours", *s.Res.MaxTrainingTimeInHours)
	}

	if s.Res.Metrics != nil {
		metricsArray := []interface{}{}
		if metricsMap := ModelMetricsToMap(&s.Res.Metrics); metricsMap != nil {
			metricsArray = append(metricsArray, metricsMap)
		}
		s.D.Set("metrics", metricsArray)
	} else {
		s.D.Set("metrics", nil)
	}

	s.D.Set("model_type", s.Res.ModelType)

	if s.Res.ModelVersion != nil {
		s.D.Set("model_version", *s.Res.ModelVersion)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TenancyId != nil {
		s.D.Set("tenancy_id", *s.Res.TenancyId)
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

	if s.Res.TrainedTimeInHours != nil {
		s.D.Set("trained_time_in_hours", *s.Res.TrainedTimeInHours)
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

func ComponentModelToMap(obj oci_ai_document.ComponentModel) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ModelId != nil {
		result["model_id"] = string(*obj.ModelId)
	}

	return result
}

func (s *AiDocumentModelResourceCrud) mapToComponentModel(fieldKeyFormat string) (oci_ai_document.ComponentModel, error) {
	var componentModel oci_ai_document.ComponentModel

	if modelId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_id")); ok {
		tmp := modelId.(string)
		componentModel.ModelId = &tmp
	}

	return componentModel, nil
}

func (s *AiDocumentModelResourceCrud) mapToDataset(fieldKeyFormat string) (oci_ai_document.Dataset, error) {
	var baseObject oci_ai_document.Dataset
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
		details := oci_ai_document.DataScienceLabelingDataset{}
		if datasetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dataset_id")); ok {
			tmp := datasetId.(string)
			details.DatasetId = &tmp
		}
		baseObject = details
	case strings.ToLower("OBJECT_STORAGE"):
		details := oci_ai_document.ObjectStorageDataset{}
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

func DatasetToMap(obj *oci_ai_document.Dataset) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_ai_document.DataScienceLabelingDataset:
		result["dataset_type"] = "DATA_SCIENCE_LABELING"

		if v.DatasetId != nil {
			result["dataset_id"] = string(*v.DatasetId)
		}
	case oci_ai_document.ObjectStorageDataset:
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

func DatasetSummaryToMap(obj *oci_ai_document.DatasetSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.TestSampleCount != nil {
		result["test_sample_count"] = int(*obj.TestSampleCount)
	}

	if obj.TrainingSampleCount != nil {
		result["training_sample_count"] = int(*obj.TrainingSampleCount)
	}

	if obj.ValidationSampleCount != nil {
		result["validation_sample_count"] = int(*obj.ValidationSampleCount)
	}

	return result
}

func (s *AiDocumentModelResourceCrud) mapToDocumentClassificationConfidenceEntry(fieldKeyFormat string) (oci_ai_document.DocumentClassificationConfidenceEntry, error) {
	result := oci_ai_document.DocumentClassificationConfidenceEntry{}

	if f1Score, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "f1score")); ok {
		tmp := f1Score.(float32)
		result.F1Score = &tmp
	}

	if precision, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "precision")); ok {
		tmp := precision.(float32)
		result.Precision = &tmp
	}

	if recall, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recall")); ok {
		tmp := recall.(float32)
		result.Recall = &tmp
	}

	if threshold, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "threshold")); ok {
		tmp := threshold.(float32)
		result.Threshold = &tmp
	}

	return result, nil
}

func DocumentClassificationConfidenceEntryToMap(obj oci_ai_document.DocumentClassificationConfidenceEntry) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.F1Score != nil {
		result["f1score"] = float32(*obj.F1Score)
	}

	if obj.Precision != nil {
		result["precision"] = float32(*obj.Precision)
	}

	if obj.Recall != nil {
		result["recall"] = float32(*obj.Recall)
	}

	if obj.Threshold != nil {
		result["threshold"] = float32(*obj.Threshold)
	}

	return result
}

func (s *AiDocumentModelResourceCrud) mapToDocumentClassificationLabelMetricsReport(fieldKeyFormat string) (oci_ai_document.DocumentClassificationLabelMetricsReport, error) {
	result := oci_ai_document.DocumentClassificationLabelMetricsReport{}

	if confidenceEntries, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "confidence_entries")); ok {
		interfaces := confidenceEntries.([]interface{})
		tmp := make([]oci_ai_document.DocumentClassificationConfidenceEntry, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "confidence_entries"), stateDataIndex)
			converted, err := s.mapToDocumentClassificationConfidenceEntry(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "confidence_entries")) {
			result.ConfidenceEntries = tmp
		}
	}

	if label, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "label")); ok {
		tmp := label.(string)
		result.Label = &tmp
	}

	if meanAveragePrecision, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mean_average_precision")); ok {
		tmp := meanAveragePrecision.(float32)
		result.MeanAveragePrecision = &tmp
	}

	return result, nil
}

func DocumentClassificationLabelMetricsReportToMap(obj oci_ai_document.DocumentClassificationLabelMetricsReport) map[string]interface{} {
	result := map[string]interface{}{}

	confidenceEntries := []interface{}{}
	for _, item := range obj.ConfidenceEntries {
		confidenceEntries = append(confidenceEntries, DocumentClassificationConfidenceEntryToMap(item))
	}
	result["confidence_entries"] = confidenceEntries

	if obj.Label != nil {
		result["label"] = string(*obj.Label)
	}

	if obj.MeanAveragePrecision != nil {
		result["mean_average_precision"] = float32(*obj.MeanAveragePrecision)
	}

	return result
}

func (s *AiDocumentModelResourceCrud) mapToDocumentClassificationOverallMetricsReport(fieldKeyFormat string) (oci_ai_document.DocumentClassificationOverallMetricsReport, error) {
	result := oci_ai_document.DocumentClassificationOverallMetricsReport{}

	if confidenceEntries, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "confidence_entries")); ok {
		interfaces := confidenceEntries.([]interface{})
		tmp := make([]oci_ai_document.DocumentClassificationConfidenceEntry, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "confidence_entries"), stateDataIndex)
			converted, err := s.mapToDocumentClassificationConfidenceEntry(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "confidence_entries")) {
			result.ConfidenceEntries = tmp
		}
	}

	if meanAveragePrecision, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mean_average_precision")); ok {
		tmp := meanAveragePrecision.(float32)
		result.MeanAveragePrecision = &tmp
	}

	return result, nil
}

func DocumentClassificationOverallMetricsReportToMap(obj *oci_ai_document.DocumentClassificationOverallMetricsReport) map[string]interface{} {
	result := map[string]interface{}{}

	confidenceEntries := []interface{}{}
	for _, item := range obj.ConfidenceEntries {
		confidenceEntries = append(confidenceEntries, DocumentClassificationConfidenceEntryToMap(item))
	}
	result["confidence_entries"] = confidenceEntries

	if obj.MeanAveragePrecision != nil {
		result["mean_average_precision"] = float32(*obj.MeanAveragePrecision)
	}

	return result
}

func (s *AiDocumentModelResourceCrud) mapToKeyValueDetectionConfidenceEntry(fieldKeyFormat string) (oci_ai_document.KeyValueDetectionConfidenceEntry, error) {
	result := oci_ai_document.KeyValueDetectionConfidenceEntry{}

	if accuracy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "accuracy")); ok {
		tmp := accuracy.(float32)
		result.Accuracy = &tmp
	}

	if f1Score, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "f1score")); ok {
		tmp := f1Score.(float32)
		result.F1Score = &tmp
	}

	if precision, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "precision")); ok {
		tmp := precision.(float32)
		result.Precision = &tmp
	}

	if recall, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recall")); ok {
		tmp := recall.(float32)
		result.Recall = &tmp
	}

	if threshold, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "threshold")); ok {
		tmp := threshold.(float32)
		result.Threshold = &tmp
	}

	return result, nil
}

func KeyValueDetectionConfidenceEntryToMap(obj oci_ai_document.KeyValueDetectionConfidenceEntry) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Accuracy != nil {
		result["accuracy"] = float32(*obj.Accuracy)
	}

	if obj.F1Score != nil {
		result["f1score"] = float32(*obj.F1Score)
	}

	if obj.Precision != nil {
		result["precision"] = float32(*obj.Precision)
	}

	if obj.Recall != nil {
		result["recall"] = float32(*obj.Recall)
	}

	if obj.Threshold != nil {
		result["threshold"] = float32(*obj.Threshold)
	}

	return result
}

func (s *AiDocumentModelResourceCrud) mapToKeyValueDetectionLabelMetricsReport(fieldKeyFormat string) (oci_ai_document.KeyValueDetectionLabelMetricsReport, error) {
	result := oci_ai_document.KeyValueDetectionLabelMetricsReport{}

	if confidenceEntries, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "confidence_entries")); ok {
		interfaces := confidenceEntries.([]interface{})
		tmp := make([]oci_ai_document.KeyValueDetectionConfidenceEntry, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "confidence_entries"), stateDataIndex)
			converted, err := s.mapToKeyValueDetectionConfidenceEntry(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "confidence_entries")) {
			result.ConfidenceEntries = tmp
		}
	}

	if documentCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "document_count")); ok {
		tmp := documentCount.(int)
		result.DocumentCount = &tmp
	}

	if label, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "label")); ok {
		tmp := label.(string)
		result.Label = &tmp
	}

	if meanAveragePrecision, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mean_average_precision")); ok {
		tmp := meanAveragePrecision.(float32)
		result.MeanAveragePrecision = &tmp
	}

	return result, nil
}

func KeyValueDetectionLabelMetricsReportToMap(obj oci_ai_document.KeyValueDetectionLabelMetricsReport) map[string]interface{} {
	result := map[string]interface{}{}

	confidenceEntries := []interface{}{}
	for _, item := range obj.ConfidenceEntries {
		confidenceEntries = append(confidenceEntries, KeyValueDetectionConfidenceEntryToMap(item))
	}
	result["confidence_entries"] = confidenceEntries

	if obj.DocumentCount != nil {
		result["document_count"] = int(*obj.DocumentCount)
	}

	if obj.Label != nil {
		result["label"] = string(*obj.Label)
	}

	if obj.MeanAveragePrecision != nil {
		result["mean_average_precision"] = float32(*obj.MeanAveragePrecision)
	}

	return result
}

func (s *AiDocumentModelResourceCrud) mapToKeyValueDetectionOverallMetricsReport(fieldKeyFormat string) (oci_ai_document.KeyValueDetectionOverallMetricsReport, error) {
	result := oci_ai_document.KeyValueDetectionOverallMetricsReport{}

	if confidenceEntries, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "confidence_entries")); ok {
		interfaces := confidenceEntries.([]interface{})
		tmp := make([]oci_ai_document.KeyValueDetectionConfidenceEntry, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "confidence_entries"), stateDataIndex)
			converted, err := s.mapToKeyValueDetectionConfidenceEntry(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "confidence_entries")) {
			result.ConfidenceEntries = tmp
		}
	}

	if documentCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "document_count")); ok {
		tmp := documentCount.(int)
		result.DocumentCount = &tmp
	}

	if meanAveragePrecision, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mean_average_precision")); ok {
		tmp := meanAveragePrecision.(float32)
		result.MeanAveragePrecision = &tmp
	}

	return result, nil
}

func KeyValueDetectionOverallMetricsReportToMap(obj *oci_ai_document.KeyValueDetectionOverallMetricsReport) map[string]interface{} {
	result := map[string]interface{}{}

	confidenceEntries := []interface{}{}
	for _, item := range obj.ConfidenceEntries {
		confidenceEntries = append(confidenceEntries, KeyValueDetectionConfidenceEntryToMap(item))
	}
	result["confidence_entries"] = confidenceEntries

	if obj.DocumentCount != nil {
		result["document_count"] = int(*obj.DocumentCount)
	}

	if obj.MeanAveragePrecision != nil {
		result["mean_average_precision"] = float32(*obj.MeanAveragePrecision)
	}

	return result
}

func ModelMetricsToMap(obj *oci_ai_document.ModelMetrics) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_ai_document.DocumentClassificationModelMetrics:
		result["model_type"] = "DOCUMENT_CLASSIFICATION"

		labelMetricsReport := []interface{}{}
		for _, item := range v.LabelMetricsReport {
			labelMetricsReport = append(labelMetricsReport, DocumentClassificationLabelMetricsReportToMap(item))
		}
		result["label_metrics_report"] = labelMetricsReport

		if v.OverallMetricsReport != nil {
			result["overall_metrics_report"] = []interface{}{DocumentClassificationOverallMetricsReportToMap(v.OverallMetricsReport)}
		}

		if v.DatasetSummary != nil {
			result["dataset_summary"] = []interface{}{DatasetSummaryToMap(v.DatasetSummary)}
		}
	case oci_ai_document.KeyValueDetectionModelMetrics:
		result["model_type"] = "KEY_VALUE_EXTRACTION"

		labelMetricsReport := []interface{}{}
		for _, item := range v.LabelMetricsReport {
			labelMetricsReport = append(labelMetricsReport, KeyValueDetectionLabelMetricsReportToMap(item))
		}
		result["label_metrics_report"] = labelMetricsReport

		if v.OverallMetricsReport != nil {
			result["overall_metrics_report"] = []interface{}{KeyValueDetectionOverallMetricsReportToMap(v.OverallMetricsReport)}
		}

		if v.DatasetSummary != nil {
			result["dataset_summary"] = []interface{}{DatasetSummaryToMap(v.DatasetSummary)}
		}
	default:
		log.Printf("[WARN] Received 'model_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func ModelSummaryToMap(obj oci_ai_document.ModelSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	componentModels := []interface{}{}
	for _, item := range obj.ComponentModels {
		componentModels = append(componentModels, ComponentModelToMap(item))
	}
	result["component_models"] = componentModels

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
	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsComposedModel != nil {
		result["is_composed_model"] = bool(*obj.IsComposedModel)
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

	if obj.TenancyId != nil {
		result["tenancy_id"] = string(*obj.TenancyId)
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

func (s *AiDocumentModelResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_ai_document.ChangeModelCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ModelId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_document")

	_, err := s.Client.ChangeModelCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
