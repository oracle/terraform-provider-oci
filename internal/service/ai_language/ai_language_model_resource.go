// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_language

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_ai_language "github.com/oracle/oci-go-sdk/v65/ailanguage"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AiLanguageModelResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createAiLanguageModel,
		Read:     readAiLanguageModel,
		Update:   updateAiLanguageModel,
		Delete:   deleteAiLanguageModel,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"model_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"model_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"NAMED_ENTITY_RECOGNITION",
								"PRE_TRAINED_HEALTH_NLU",
								"PRE_TRAINED_KEYPHRASE_EXTRACTION",
								"PRE_TRAINED_LANGUAGE_DETECTION",
								"PRE_TRAINED_NAMED_ENTITY_RECOGNITION",
								"PRE_TRAINED_PHI",
								"PRE_TRAINED_PII",
								"PRE_TRAINED_SENTIMENT_ANALYSIS",
								"PRE_TRAINED_SUMMARIZATION",
								"PRE_TRAINED_TEXT_CLASSIFICATION",
								"PRE_TRAINED_UNIVERSAL",
								"TEXT_CLASSIFICATION",
							}, true),
						},

						// Optional
						"classification_mode": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"classification_mode": {
										Type:             schema.TypeString,
										Required:         true,
										ForceNew:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"MULTI_CLASS",
											"MULTI_LABEL",
										}, true),
									},

									// Optional
									"version": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"language_code": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
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
						"dataset_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"location_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
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
									"location_type": {
										Type:             schema.TypeString,
										Required:         true,
										ForceNew:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"OBJECT_LIST",
										}, true),
									},
									"namespace": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"object_names": {
										Type:     schema.TypeList,
										Required: true,
										ForceNew: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									// Optional

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
			"test_strategy": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"strategy_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"TEST_AND_VALIDATION_DATASET",
							}, true),
						},
						"testing_dataset": {
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
									"dataset_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"location_details": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
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
												"location_type": {
													Type:             schema.TypeString,
													Required:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"OBJECT_LIST",
													}, true),
												},
												"namespace": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"object_names": {
													Type:     schema.TypeList,
													Required: true,
													ForceNew: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												// Optional

												// Computed
											},
										},
									},

									// Computed
								},
							},
						},

						// Optional
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
									"dataset_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"location_details": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
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
												"location_type": {
													Type:             schema.TypeString,
													Required:         true,
													ForceNew:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"OBJECT_LIST",
													}, true),
												},
												"namespace": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"object_names": {
													Type:     schema.TypeList,
													Required: true,
													ForceNew: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												// Optional

												// Computed
											},
										},
									},

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},

			// Computed
			"evaluation_results": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"class_metrics": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"f1": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"label": {
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
									"support": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},
						"confusion_matrix": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"labels": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"entity_metrics": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"f1": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"label": {
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
								},
							},
						},
						"metrics": {
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
									"macro_f1": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"macro_precision": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"macro_recall": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"micro_f1": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"micro_precision": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"micro_recall": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"weighted_f1": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"weighted_precision": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"weighted_recall": {
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
					},
				},
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
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createAiLanguageModel(d *schema.ResourceData, m interface{}) error {
	sync := &AiLanguageModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceLanguageClient()

	return tfresource.CreateResource(d, sync)
}

func readAiLanguageModel(d *schema.ResourceData, m interface{}) error {
	sync := &AiLanguageModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceLanguageClient()

	return tfresource.ReadResource(sync)
}

func updateAiLanguageModel(d *schema.ResourceData, m interface{}) error {
	sync := &AiLanguageModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceLanguageClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteAiLanguageModel(d *schema.ResourceData, m interface{}) error {
	sync := &AiLanguageModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceLanguageClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type AiLanguageModelResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_ai_language.AIServiceLanguageClient
	Res                    *oci_ai_language.Model
	DisableNotFoundRetries bool
}

func (s *AiLanguageModelResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *AiLanguageModelResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_ai_language.ModelLifecycleStateCreating),
	}
}

func (s *AiLanguageModelResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_ai_language.ModelLifecycleStateActive),
	}
}

func (s *AiLanguageModelResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_ai_language.ModelLifecycleStateUpdating),
	}
}

func (s *AiLanguageModelResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_ai_language.ModelLifecycleStateActive),
	}
}

func (s *AiLanguageModelResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_ai_language.ModelLifecycleStateDeleting),
	}
}

func (s *AiLanguageModelResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_ai_language.ModelLifecycleStateDeleted),
	}
}

func (s *AiLanguageModelResourceCrud) Create() error {
	request := oci_ai_language.CreateModelRequest{}

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

	if modelDetails, ok := s.D.GetOkExists("model_details"); ok {
		if tmpList := modelDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "model_details", 0)
			tmp, err := s.mapToModelDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ModelDetails = tmp
		}
	}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	if testStrategy, ok := s.D.GetOkExists("test_strategy"); ok {
		if tmpList := testStrategy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "test_strategy", 0)
			tmp, err := s.mapToTestStrategy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TestStrategy = tmp
		}
	}

	if trainingDataset, ok := s.D.GetOkExists("training_dataset"); ok {
		if tmpList := trainingDataset.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "training_dataset", 0)
			tmp, err := s.mapToDatasetDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TrainingDataset = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_language")

	response, err := s.Client.CreateModel(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getModelFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "model"), oci_ai_language.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *AiLanguageModelResourceCrud) getModelFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_ai_language.ActionTypeEnum, timeout time.Duration) error {

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
		if tfresource.ShouldRetry(response, false, "ai_language", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_ai_language.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func modelWaitForWorkRequest(wId *string, entityType string, action oci_ai_language.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_ai_language.AIServiceLanguageClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "ai_language")
	retryPolicy.ShouldRetryOperation = modelWorkRequestShouldRetryFunc(timeout)

	response := oci_ai_language.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_ai_language.ActionTypeInProgress),
		},
		Target: []string{
			string("ACTIVE"),
			string("SUCCEEDED"),
			string("FAILED"),
			string(oci_ai_language.ActionTypeDeleted),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_ai_language.GetWorkRequestRequest{
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
	// if identifier == nil || response.Status == oci_ai_language.WorkRequestStatusFailed || response.Status == oci_ai_language.WorkRequestStatusCanceled {
	// 	return nil, getErrorFromAiLanguageModelWorkRequest(client, wId, retryPolicy, entityType, action)
	// }

	return identifier, nil
}

func getErrorFromAiLanguageModelWorkRequest(client *oci_ai_language.AIServiceLanguageClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_ai_language.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_ai_language.ListWorkRequestErrorsRequest{
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

func (s *AiLanguageModelResourceCrud) Get() error {
	request := oci_ai_language.GetModelRequest{}

	tmp := s.D.Id()
	request.ModelId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_language")

	response, err := s.Client.GetModel(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Model
	return nil
}

func (s *AiLanguageModelResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_ai_language.UpdateModelRequest{}

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_language")

	_, err := s.Client.UpdateModel(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		fmt.Printf("waitErr: %v\n", waitErr)
		return waitErr
	}

	return nil

	// workId := response.OpcWorkRequestId
	// return s.getModelFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_language"), oci_ai_language.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *AiLanguageModelResourceCrud) Delete() error {
	request := oci_ai_language.DeleteModelRequest{}

	tmp := s.D.Id()
	request.ModelId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_language")

	response, err := s.Client.DeleteModel(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := modelWaitForWorkRequest(workId, "model",
		oci_ai_language.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *AiLanguageModelResourceCrud) SetData() error {
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

	if s.Res.EvaluationResults != nil {
		evaluationResultsArray := []interface{}{}
		if evaluationResultsMap := EvaluationResultsToMap(&s.Res.EvaluationResults); evaluationResultsMap != nil {
			evaluationResultsArray = append(evaluationResultsArray, evaluationResultsMap)
		}
		s.D.Set("evaluation_results", evaluationResultsArray)
	} else {
		s.D.Set("evaluation_results", nil)
	}

	if s.Res.FreeformTags != nil {
		s.D.Set("freeform_tags", s.Res.FreeformTags)
	}

	// s.D.Set("freeform_tags", s.Res.FreeformTags)
	// s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ModelDetails != nil {
		modelDetailsArray := []interface{}{}
		if modelDetailsMap := ModelDetailsToMap(&s.Res.ModelDetails); modelDetailsMap != nil {
			modelDetailsArray = append(modelDetailsArray, modelDetailsMap)
		}
		s.D.Set("model_details", modelDetailsArray)
	} else {
		s.D.Set("model_details", nil)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", s.Res.SystemTags)
	}

	if s.Res.TestStrategy != nil {
		testStrategyArray := []interface{}{}
		if testStrategyMap := TestStrategyToMap(&s.Res.TestStrategy); testStrategyMap != nil {
			testStrategyArray = append(testStrategyArray, testStrategyMap)
		}
		s.D.Set("test_strategy", testStrategyArray)
	} else {
		s.D.Set("test_strategy", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TrainingDataset != nil {
		trainingDatasetArray := []interface{}{}
		if trainingDatasetMap := DatasetDetailsToMap(&s.Res.TrainingDataset); trainingDatasetMap != nil {
			trainingDatasetArray = append(trainingDatasetArray, trainingDatasetMap)
		}
		s.D.Set("training_dataset", trainingDatasetArray)
	} else {
		s.D.Set("training_dataset", nil)
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}

func (s *AiLanguageModelResourceCrud) mapToClassMetrics(fieldKeyFormat string) (oci_ai_language.ClassMetrics, error) {
	result := oci_ai_language.ClassMetrics{}

	if f1, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "f1")); ok {
		tmp := f1.(float32)
		result.F1 = &tmp
	}

	if label, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "label")); ok {
		tmp := label.(string)
		result.Label = &tmp
	}

	if precision, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "precision")); ok {
		tmp := precision.(float32)
		result.Precision = &tmp
	}

	if recall, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recall")); ok {
		tmp := recall.(float32)
		result.Recall = &tmp
	}

	if support, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "support")); ok {
		tmp := support.(float32)
		result.Support = &tmp
	}

	return result, nil
}

func ClassMetricsToMap(obj oci_ai_language.ClassMetrics) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.F1 != nil {
		result["f1"] = float32(*obj.F1)
	}

	if obj.Label != nil {
		result["label"] = string(*obj.Label)
	}

	if obj.Precision != nil {
		result["precision"] = float32(*obj.Precision)
	}

	if obj.Recall != nil {
		result["recall"] = float32(*obj.Recall)
	}

	if obj.Support != nil {
		result["support"] = float32(*obj.Support)
	}

	return result
}

func (s *AiLanguageModelResourceCrud) mapToClassificationType(fieldKeyFormat string) (oci_ai_language.ClassificationType, error) {
	var baseObject oci_ai_language.ClassificationType
	//discriminator
	classificationModeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "classification_mode"))
	var classificationMode string
	if ok {
		classificationMode = classificationModeRaw.(string)
	} else {
		classificationMode = "" // default value
	}
	switch strings.ToLower(classificationMode) {
	case strings.ToLower("MULTI_CLASS"):
		details := oci_ai_language.ClassificationMultiClassModeDetails{}
		if version, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version")); ok {
			tmp := version.(string)
			details.Version = &tmp
		}
		baseObject = details
	case strings.ToLower("MULTI_LABEL"):
		details := oci_ai_language.ClassificationMultiLabelModeDetails{}
		if version, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version")); ok {
			tmp := version.(string)
			details.Version = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown classification_mode '%v' was specified", classificationMode)
	}
	return baseObject, nil
}

func ClassificationTypeToMap(obj *oci_ai_language.ClassificationType) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_ai_language.ClassificationMultiClassModeDetails:
		result["classification_mode"] = "MULTI_CLASS"

		if v.Version != nil {
			result["version"] = string(*v.Version)
		}
	case oci_ai_language.ClassificationMultiLabelModeDetails:
		result["classification_mode"] = "MULTI_LABEL"

		if v.Version != nil {
			result["version"] = string(*v.Version)
		}
	default:
		log.Printf("[WARN] Received 'classification_mode' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *AiLanguageModelResourceCrud) mapToDatasetDetails(fieldKeyFormat string) (oci_ai_language.DatasetDetails, error) {
	var baseObject oci_ai_language.DatasetDetails
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
		details := oci_ai_language.DataScienceLabelingDataset{}
		if datasetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dataset_id")); ok {
			tmp := datasetId.(string)
			details.DatasetId = &tmp
		}
		baseObject = details
	case strings.ToLower("OBJECT_STORAGE"):
		details := oci_ai_language.ObjectStorageDataset{}
		if locationDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "location_details")); ok {
			if tmpList := locationDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "location_details"), 0)
				tmp, err := s.mapToLocationDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert location_details, encountered error: %v", err)
				}
				details.LocationDetails = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown dataset_type '%v' was specified", datasetType)
	}
	return baseObject, nil
}

func DatasetDetailsToMap(obj *oci_ai_language.DatasetDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_ai_language.DataScienceLabelingDataset:
		result["dataset_type"] = "DATA_SCIENCE_LABELING"

		if v.DatasetId != nil {
			result["dataset_id"] = string(*v.DatasetId)
		}
	case oci_ai_language.ObjectStorageDataset:
		result["dataset_type"] = "OBJECT_STORAGE"

		if v.LocationDetails != nil {
			locationDetailsArray := []interface{}{}
			if locationDetailsMap := LocationDetailsToMap(&v.LocationDetails); locationDetailsMap != nil {
				locationDetailsArray = append(locationDetailsArray, locationDetailsMap)
			}
			result["location_details"] = locationDetailsArray
		}
	default:
		log.Printf("[WARN] Received 'dataset_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *AiLanguageModelResourceCrud) mapToEntityMetrics(fieldKeyFormat string) (oci_ai_language.EntityMetrics, error) {
	result := oci_ai_language.EntityMetrics{}

	if f1, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "f1")); ok {
		tmp := f1.(float32)
		result.F1 = &tmp
	}

	if label, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "label")); ok {
		tmp := label.(string)
		result.Label = &tmp
	}

	if precision, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "precision")); ok {
		tmp := precision.(float32)
		result.Precision = &tmp
	}

	if recall, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recall")); ok {
		tmp := recall.(float32)
		result.Recall = &tmp
	}

	return result, nil
}

func EntityMetricsToMap(obj oci_ai_language.EntityMetrics) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.F1 != nil {
		result["f1"] = float32(*obj.F1)
	}

	if obj.Label != nil {
		result["label"] = string(*obj.Label)
	}

	if obj.Precision != nil {
		result["precision"] = float32(*obj.Precision)
	}

	if obj.Recall != nil {
		result["recall"] = float32(*obj.Recall)
	}

	return result
}

func EvaluationResultsToMap(obj *oci_ai_language.EvaluationResults) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_ai_language.NamedEntityRecognitionEvaluationResults:
		result["model_type"] = "NAMED_ENTITY_RECOGNITION"

		buf, err := json.Marshal(v.ConfusionMatrix)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ConfusionMatrix%s\n", string(buf))
		result["confusion_matrix"] = string(buf)

		// buf, err := json.Marshal(v.ConfusionMatrix)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Printf("ConfusionMatrix%s\n", buf)

		// result["confusion_matrix"] = v.ConfusionMatrix
		// result["confusion_matrix"] = v.ConfusionMatrix

		entityMetrics := []interface{}{}
		for _, item := range v.EntityMetrics {
			entityMetrics = append(entityMetrics, EntityMetricsToMap(item))
		}
		result["entity_metrics"] = entityMetrics

		result["labels"] = v.Labels
		result["labels"] = v.Labels

		if v.Metrics != nil {
			result["metrics"] = []interface{}{NamedEntityRecognitionModelMetricsToMap(v.Metrics)}
		}
	case oci_ai_language.TextClassificationEvaluationResults:
		result["model_type"] = "TEXT_CLASSIFICATION"

		classMetrics := []interface{}{}
		for _, item := range v.ClassMetrics {
			classMetrics = append(classMetrics, ClassMetricsToMap(item))
		}
		result["class_metrics"] = classMetrics

		buf, err := json.Marshal(v.ConfusionMatrix)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ConfusionMatrix%s\n", string(buf))
		result["confusion_matrix"] = string(buf)

		// result["confusion_matrix"] = v.ConfusionMatrix
		// result["confusion_matrix"] = v.ConfusionMatrix

		result["labels"] = v.Labels
		result["labels"] = v.Labels

		if v.Metrics != nil {
			result["metrics"] = []interface{}{TextClassificationModelMetricsToMap(v.Metrics)}
		}
	default:
		log.Printf("[WARN] Received 'model_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *AiLanguageModelResourceCrud) mapToLocationDetails(fieldKeyFormat string) (oci_ai_language.LocationDetails, error) {
	var baseObject oci_ai_language.LocationDetails
	//discriminator
	locationTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "location_type"))
	var locationType string
	if ok {
		locationType = locationTypeRaw.(string)
	} else {
		locationType = "" // default value
	}
	switch strings.ToLower(locationType) {
	case strings.ToLower("OBJECT_LIST"):
		details := oci_ai_language.ObjectListDataset{}
		if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.BucketName = &tmp
		}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.NamespaceName = &tmp
		}
		if objectNames, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_names")); ok {
			interfaces := objectNames.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "object_names")) {
				details.ObjectNames = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown location_type '%v' was specified", locationType)
	}
	return baseObject, nil
}

func LocationDetailsToMap(obj *oci_ai_language.LocationDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_ai_language.ObjectListDataset:
		result["location_type"] = "OBJECT_LIST"

		if v.BucketName != nil {
			result["bucket"] = string(*v.BucketName)
		}

		if v.NamespaceName != nil {
			result["namespace"] = string(*v.NamespaceName)
		}

		result["object_names"] = v.ObjectNames
	default:
		log.Printf("[WARN] Received 'location_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *AiLanguageModelResourceCrud) mapToModelDetails(fieldKeyFormat string) (oci_ai_language.ModelDetails, error) {
	var baseObject oci_ai_language.ModelDetails
	//discriminator
	modelTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_type"))
	var modelType string
	if ok {
		modelType = modelTypeRaw.(string)
	} else {
		modelType = "" // default value
	}
	switch strings.ToLower(modelType) {
	case strings.ToLower("NAMED_ENTITY_RECOGNITION"):
		details := oci_ai_language.NamedEntityRecognitionModelDetails{}
		if version, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version")); ok {
			tmp := version.(string)
			details.Version = &tmp
		}
		if languageCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "language_code")); ok {
			tmp := languageCode.(string)
			details.LanguageCode = &tmp
		}
		baseObject = details
	case strings.ToLower("PRE_TRAINED_HEALTH_NLU"):
		details := oci_ai_language.PreTrainedHealthNluModelDetails{}
		if version, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version")); ok {
			tmp := version.(string)
			details.Version = &tmp
		}
		if languageCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "language_code")); ok {
			tmp := languageCode.(string)
			details.LanguageCode = &tmp
		}
		baseObject = details
	case strings.ToLower("PRE_TRAINED_KEYPHRASE_EXTRACTION"):
		details := oci_ai_language.PreTrainedKeyPhraseExtractionModelDetails{}
		if version, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version")); ok {
			tmp := version.(string)
			details.Version = &tmp
		}
		if languageCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "language_code")); ok {
			tmp := languageCode.(string)
			details.LanguageCode = &tmp
		}
		baseObject = details
	case strings.ToLower("PRE_TRAINED_LANGUAGE_DETECTION"):
		details := oci_ai_language.PreTrainedLanguageDetectionModelDetails{}
		if version, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version")); ok {
			tmp := version.(string)
			details.Version = &tmp
		}
		if languageCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "language_code")); ok {
			tmp := languageCode.(string)
			details.LanguageCode = &tmp
		}
		baseObject = details
	case strings.ToLower("PRE_TRAINED_NAMED_ENTITY_RECOGNITION"):
		details := oci_ai_language.PreTrainedNamedEntityRecognitionModelDetails{}
		if version, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version")); ok {
			tmp := version.(string)
			details.Version = &tmp
		}
		if languageCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "language_code")); ok {
			tmp := languageCode.(string)
			details.LanguageCode = &tmp
		}
		baseObject = details
	// case strings.ToLower("PRE_TRAINED_PHI"):
	// 	details := oci_ai_language.PreTrainedPhiModelDetails{}
	// 	if version, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version")); ok {
	// 		tmp := version.(string)
	// 		details.Version = &tmp
	// 	}
	// 	if languageCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "language_code")); ok {
	// 		tmp := languageCode.(string)
	// 		details.LanguageCode = &tmp
	// 	}
	// 	baseObject = details
	case strings.ToLower("PRE_TRAINED_PII"):
		details := oci_ai_language.PreTrainedPiiModelDetails{}
		if version, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version")); ok {
			tmp := version.(string)
			details.Version = &tmp
		}
		if languageCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "language_code")); ok {
			tmp := languageCode.(string)
			details.LanguageCode = &tmp
		}
		baseObject = details
	case strings.ToLower("PRE_TRAINED_SENTIMENT_ANALYSIS"):
		details := oci_ai_language.PreTrainedSentimentAnalysisModelDetails{}
		if version, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version")); ok {
			tmp := version.(string)
			details.Version = &tmp
		}
		if languageCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "language_code")); ok {
			tmp := languageCode.(string)
			details.LanguageCode = &tmp
		}
		baseObject = details
	case strings.ToLower("PRE_TRAINED_SUMMARIZATION"):
		details := oci_ai_language.PreTrainedSummarization{}
		if version, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version")); ok {
			tmp := version.(string)
			details.Version = &tmp
		}
		if languageCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "language_code")); ok {
			tmp := languageCode.(string)
			details.LanguageCode = &tmp
		}
		baseObject = details
	case strings.ToLower("PRE_TRAINED_TEXT_CLASSIFICATION"):
		details := oci_ai_language.PreTrainedTextClassificationModelDetails{}
		if version, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version")); ok {
			tmp := version.(string)
			details.Version = &tmp
		}
		if languageCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "language_code")); ok {
			tmp := languageCode.(string)
			details.LanguageCode = &tmp
		}
		baseObject = details
	case strings.ToLower("PRE_TRAINED_UNIVERSAL"):
		details := oci_ai_language.PreTrainedUniversalModel{}
		if version, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version")); ok {
			tmp := version.(string)
			details.Version = &tmp
		}
		if languageCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "language_code")); ok {
			tmp := languageCode.(string)
			details.LanguageCode = &tmp
		}
		baseObject = details
	case strings.ToLower("TEXT_CLASSIFICATION"):
		details := oci_ai_language.TextClassificationModelDetails{}
		if classificationMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "classification_mode")); ok {
			if tmpList := classificationMode.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "classification_mode"), 0)
				tmp, err := s.mapToClassificationType(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert classification_mode, encountered error: %v", err)
				}
				details.ClassificationMode = tmp
			}
		}
		if languageCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "language_code")); ok {
			tmp := languageCode.(string)
			details.LanguageCode = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown model_type '%v' was specified", modelType)
	}
	return baseObject, nil
}

func ModelDetailsToMap(obj *oci_ai_language.ModelDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_ai_language.NamedEntityRecognitionModelDetails:
		result["model_type"] = "NAMED_ENTITY_RECOGNITION"

		if v.Version != nil {
			result["version"] = string(*v.Version)
		}

		if v.LanguageCode != nil {
			result["language_code"] = string(*v.LanguageCode)
		}
	case oci_ai_language.PreTrainedHealthNluModelDetails:
		result["model_type"] = "PRE_TRAINED_HEALTH_NLU"

		if v.Version != nil {
			result["version"] = string(*v.Version)
		}

		if v.LanguageCode != nil {
			result["language_code"] = string(*v.LanguageCode)
		}
	case oci_ai_language.PreTrainedKeyPhraseExtractionModelDetails:
		result["model_type"] = "PRE_TRAINED_KEYPHRASE_EXTRACTION"

		if v.Version != nil {
			result["version"] = string(*v.Version)
		}

		if v.LanguageCode != nil {
			result["language_code"] = string(*v.LanguageCode)
		}
	case oci_ai_language.PreTrainedLanguageDetectionModelDetails:
		result["model_type"] = "PRE_TRAINED_LANGUAGE_DETECTION"

		if v.Version != nil {
			result["version"] = string(*v.Version)
		}

		if v.LanguageCode != nil {
			result["language_code"] = string(*v.LanguageCode)
		}
	case oci_ai_language.PreTrainedNamedEntityRecognitionModelDetails:
		result["model_type"] = "PRE_TRAINED_NAMED_ENTITY_RECOGNITION"

		if v.Version != nil {
			result["version"] = string(*v.Version)
		}

		if v.LanguageCode != nil {
			result["language_code"] = string(*v.LanguageCode)
		}
	// case oci_ai_language.PreTrainedPhiModelDetails:
	// 	result["model_type"] = "PRE_TRAINED_PHI"

	// 	if v.Version != nil {
	// 		result["version"] = string(*v.Version)
	// 	}

	// 	if v.LanguageCode != nil {
	// 		result["language_code"] = string(*v.LanguageCode)
	// 	}
	case oci_ai_language.PreTrainedPiiModelDetails:
		result["model_type"] = "PRE_TRAINED_PII"

		if v.Version != nil {
			result["version"] = string(*v.Version)
		}

		if v.LanguageCode != nil {
			result["language_code"] = string(*v.LanguageCode)
		}
	case oci_ai_language.PreTrainedSentimentAnalysisModelDetails:
		result["model_type"] = "PRE_TRAINED_SENTIMENT_ANALYSIS"

		if v.Version != nil {
			result["version"] = string(*v.Version)
		}

		if v.LanguageCode != nil {
			result["language_code"] = string(*v.LanguageCode)
		}
	case oci_ai_language.PreTrainedSummarization:
		result["model_type"] = "PRE_TRAINED_SUMMARIZATION"

		if v.Version != nil {
			result["version"] = string(*v.Version)
		}

		if v.LanguageCode != nil {
			result["language_code"] = string(*v.LanguageCode)
		}
	case oci_ai_language.PreTrainedTextClassificationModelDetails:
		result["model_type"] = "PRE_TRAINED_TEXT_CLASSIFICATION"

		if v.Version != nil {
			result["version"] = string(*v.Version)
		}

		if v.LanguageCode != nil {
			result["language_code"] = string(*v.LanguageCode)
		}
	case oci_ai_language.PreTrainedUniversalModel:
		result["model_type"] = "PRE_TRAINED_UNIVERSAL"

		if v.Version != nil {
			result["version"] = string(*v.Version)
		}

		if v.LanguageCode != nil {
			result["language_code"] = string(*v.LanguageCode)
		}
	case oci_ai_language.TextClassificationModelDetails:
		result["model_type"] = "TEXT_CLASSIFICATION"

		if v.ClassificationMode != nil {
			classificationModeArray := []interface{}{}
			if classificationModeMap := ClassificationTypeToMap(&v.ClassificationMode); classificationModeMap != nil {
				classificationModeArray = append(classificationModeArray, classificationModeMap)
			}
			result["classification_mode"] = classificationModeArray
		}

		if v.LanguageCode != nil {
			result["language_code"] = string(*v.LanguageCode)
		}
	default:
		log.Printf("[WARN] Received 'model_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func ModelSummaryToMap(obj oci_ai_language.ModelSummary) map[string]interface{} {
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

	if obj.FreeformTags != nil {
		result["freeform_tags"] = obj.FreeformTags
	}

	// result["freeform_tags"] = obj.FreeformTags
	// result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.ModelDetails != nil {
		modelDetailsArray := []interface{}{}
		if modelDetailsMap := ModelDetailsToMap(&obj.ModelDetails); modelDetailsMap != nil {
			modelDetailsArray = append(modelDetailsArray, modelDetailsMap)
		}
		result["model_details"] = modelDetailsArray
	}

	if obj.ProjectId != nil {
		result["project_id"] = string(*obj.ProjectId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = obj.SystemTags
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}

func (s *AiLanguageModelResourceCrud) mapToNamedEntityRecognitionModelMetrics(fieldKeyFormat string) (oci_ai_language.NamedEntityRecognitionModelMetrics, error) {
	result := oci_ai_language.NamedEntityRecognitionModelMetrics{}

	if macroF1, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "macro_f1")); ok {
		tmp := macroF1.(float32)
		result.MacroF1 = &tmp
	}

	if macroPrecision, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "macro_precision")); ok {
		tmp := macroPrecision.(float32)
		result.MacroPrecision = &tmp
	}

	if macroRecall, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "macro_recall")); ok {
		tmp := macroRecall.(float32)
		result.MacroRecall = &tmp
	}

	if microF1, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "micro_f1")); ok {
		tmp := microF1.(float32)
		result.MicroF1 = &tmp
	}

	if microPrecision, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "micro_precision")); ok {
		tmp := microPrecision.(float32)
		result.MicroPrecision = &tmp
	}

	if microRecall, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "micro_recall")); ok {
		tmp := microRecall.(float32)
		result.MicroRecall = &tmp
	}

	if weightedF1, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "weighted_f1")); ok {
		tmp := weightedF1.(float32)
		result.WeightedF1 = &tmp
	}

	if weightedPrecision, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "weighted_precision")); ok {
		tmp := weightedPrecision.(float32)
		result.WeightedPrecision = &tmp
	}

	if weightedRecall, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "weighted_recall")); ok {
		tmp := weightedRecall.(float32)
		result.WeightedRecall = &tmp
	}

	return result, nil
}

func NamedEntityRecognitionModelMetricsToMap(obj *oci_ai_language.NamedEntityRecognitionModelMetrics) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MacroF1 != nil {
		result["macro_f1"] = float32(*obj.MacroF1)
	}

	if obj.MacroPrecision != nil {
		result["macro_precision"] = float32(*obj.MacroPrecision)
	}

	if obj.MacroRecall != nil {
		result["macro_recall"] = float32(*obj.MacroRecall)
	}

	if obj.MicroF1 != nil {
		result["micro_f1"] = float32(*obj.MicroF1)
	}

	if obj.MicroPrecision != nil {
		result["micro_precision"] = float32(*obj.MicroPrecision)
	}

	if obj.MicroRecall != nil {
		result["micro_recall"] = float32(*obj.MicroRecall)
	}

	if obj.WeightedF1 != nil {
		result["weighted_f1"] = float32(*obj.WeightedF1)
	}

	if obj.WeightedPrecision != nil {
		result["weighted_precision"] = float32(*obj.WeightedPrecision)
	}

	if obj.WeightedRecall != nil {
		result["weighted_recall"] = float32(*obj.WeightedRecall)
	}

	return result
}

func (s *AiLanguageModelResourceCrud) mapToTestStrategy(fieldKeyFormat string) (oci_ai_language.TestStrategy, error) {
	var baseObject oci_ai_language.TestStrategy
	//discriminator
	strategyTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "strategy_type"))
	var strategyType string
	if ok {
		strategyType = strategyTypeRaw.(string)
	} else {
		strategyType = "" // default value
	}
	switch strings.ToLower(strategyType) {
	case strings.ToLower("TEST_AND_VALIDATION_DATASET"):
		details := oci_ai_language.TestAndValidationDatasetStrategy{}
		if testingDataset, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "testing_dataset")); ok {
			if tmpList := testingDataset.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "testing_dataset"), 0)
				tmp, err := s.mapToDatasetDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert testing_dataset, encountered error: %v", err)
				}
				details.TestingDataset = tmp
			}
		}
		if validationDataset, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "validation_dataset")); ok {
			if tmpList := validationDataset.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "validation_dataset"), 0)
				tmp, err := s.mapToDatasetDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert validation_dataset, encountered error: %v", err)
				}
				details.ValidationDataset = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown strategy_type '%v' was specified", strategyType)
	}
	return baseObject, nil
}

func TestStrategyToMap(obj *oci_ai_language.TestStrategy) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_ai_language.TestAndValidationDatasetStrategy:
		result["strategy_type"] = "TEST_AND_VALIDATION_DATASET"

		if v.TestingDataset != nil {
			testingDatasetArray := []interface{}{}
			if testingDatasetMap := DatasetDetailsToMap(&v.TestingDataset); testingDatasetMap != nil {
				testingDatasetArray = append(testingDatasetArray, testingDatasetMap)
			}
			result["testing_dataset"] = testingDatasetArray
		}

		if v.ValidationDataset != nil {
			validationDatasetArray := []interface{}{}
			if validationDatasetMap := DatasetDetailsToMap(&v.ValidationDataset); validationDatasetMap != nil {
				validationDatasetArray = append(validationDatasetArray, validationDatasetMap)
			}
			result["validation_dataset"] = validationDatasetArray
		}
	default:
		log.Printf("[WARN] Received 'strategy_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *AiLanguageModelResourceCrud) mapToTextClassificationModelMetrics(fieldKeyFormat string) (oci_ai_language.TextClassificationModelMetrics, error) {
	result := oci_ai_language.TextClassificationModelMetrics{}

	if accuracy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "accuracy")); ok {
		tmp := accuracy.(float32)
		result.Accuracy = &tmp
	}

	if macroF1, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "macro_f1")); ok {
		tmp := macroF1.(float32)
		result.MacroF1 = &tmp
	}

	if macroPrecision, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "macro_precision")); ok {
		tmp := macroPrecision.(float32)
		result.MacroPrecision = &tmp
	}

	if macroRecall, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "macro_recall")); ok {
		tmp := macroRecall.(float32)
		result.MacroRecall = &tmp
	}

	if microF1, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "micro_f1")); ok {
		tmp := microF1.(float32)
		result.MicroF1 = &tmp
	}

	if microPrecision, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "micro_precision")); ok {
		tmp := microPrecision.(float32)
		result.MicroPrecision = &tmp
	}

	if microRecall, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "micro_recall")); ok {
		tmp := microRecall.(float32)
		result.MicroRecall = &tmp
	}

	if weightedF1, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "weighted_f1")); ok {
		tmp := weightedF1.(float32)
		result.WeightedF1 = &tmp
	}

	if weightedPrecision, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "weighted_precision")); ok {
		tmp := weightedPrecision.(float32)
		result.WeightedPrecision = &tmp
	}

	if weightedRecall, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "weighted_recall")); ok {
		tmp := weightedRecall.(float32)
		result.WeightedRecall = &tmp
	}

	return result, nil
}

func TextClassificationModelMetricsToMap(obj *oci_ai_language.TextClassificationModelMetrics) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Accuracy != nil {
		result["accuracy"] = float32(*obj.Accuracy)
	}

	if obj.MacroF1 != nil {
		result["macro_f1"] = float32(*obj.MacroF1)
	}

	if obj.MacroPrecision != nil {
		result["macro_precision"] = float32(*obj.MacroPrecision)
	}

	if obj.MacroRecall != nil {
		result["macro_recall"] = float32(*obj.MacroRecall)
	}

	if obj.MicroF1 != nil {
		result["micro_f1"] = float32(*obj.MicroF1)
	}

	if obj.MicroPrecision != nil {
		result["micro_precision"] = float32(*obj.MicroPrecision)
	}

	if obj.MicroRecall != nil {
		result["micro_recall"] = float32(*obj.MicroRecall)
	}

	if obj.WeightedF1 != nil {
		result["weighted_f1"] = float32(*obj.WeightedF1)
	}

	if obj.WeightedPrecision != nil {
		result["weighted_precision"] = float32(*obj.WeightedPrecision)
	}

	if obj.WeightedRecall != nil {
		result["weighted_recall"] = float32(*obj.WeightedRecall)
	}

	return result
}

func (s *AiLanguageModelResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_ai_language.ChangeModelCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ModelId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_language")

	_, err := s.Client.ChangeModelCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
