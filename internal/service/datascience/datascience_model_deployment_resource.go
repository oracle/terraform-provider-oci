// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_datascience "github.com/oracle/oci-go-sdk/v58/datascience"
)

func DatascienceModelDeploymentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: &tfresource.TwentyMinutes,
			Update: &tfresource.ThirtyMinutes,
			Delete: &tfresource.TwentyMinutes,
		},
		Create: createDatascienceModelDeployment,
		Read:   readDatascienceModelDeployment,
		Update: updateDatascienceModelDeployment,
		Delete: deleteDatascienceModelDeployment,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"model_deployment_configuration_details": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"deployment_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"SINGLE_MODEL",
							}, true),
						},
						"model_configuration_details": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"instance_configuration": {
										Type:     schema.TypeList,
										Required: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"instance_shape_name": {
													Type:     schema.TypeString,
													Required: true,
												},

												// Optional

												// Computed
											},
										},
									},
									"model_id": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"bandwidth_mbps": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"scaling_policy": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"instance_count": {
													Type:     schema.TypeInt,
													Required: true,
												},
												"policy_type": {
													Type:             schema.TypeString,
													Required:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"FIXED_SIZE",
													}, true),
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
			"category_log_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"access": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"log_group_id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"log_id": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"predict": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"log_group_id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"log_id": {
										Type:     schema.TypeString,
										Required: true,
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
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"model_deployment_url": {
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
		},
	}
}

func createDatascienceModelDeployment(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelDeploymentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.CreateResource(d, sync)
}

func readDatascienceModelDeployment(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelDeploymentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

func updateDatascienceModelDeployment(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelDeploymentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatascienceModelDeployment(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelDeploymentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatascienceModelDeploymentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_datascience.DataScienceClient
	Res                    *oci_datascience.ModelDeployment
	DisableNotFoundRetries bool
}

func (s *DatascienceModelDeploymentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatascienceModelDeploymentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_datascience.ModelDeploymentLifecycleStateCreating),
	}
}

func (s *DatascienceModelDeploymentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_datascience.ModelDeploymentLifecycleStateActive),
		string(oci_datascience.ModelDeploymentLifecycleStateFailed),
	}
}

func (s *DatascienceModelDeploymentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_datascience.ModelDeploymentLifecycleStateDeleting),
	}
}

func (s *DatascienceModelDeploymentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_datascience.ModelDeploymentLifecycleStateDeleted),
		string(oci_datascience.ModelDeploymentLifecycleStateFailed),
	}
}

func (s *DatascienceModelDeploymentResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_datascience.ModelDeploymentLifecycleStateUpdating),
	}
}

func (s *DatascienceModelDeploymentResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_datascience.ModelDeploymentLifecycleStateActive),
		string(oci_datascience.ModelDeploymentLifecycleStateInactive),
		string(oci_datascience.ModelDeploymentLifecycleStateFailed),
		string(oci_datascience.ModelDeploymentLifecycleStateNeedsAttention),
	}
}

func (s *DatascienceModelDeploymentResourceCrud) Create() error {
	request := oci_datascience.CreateModelDeploymentRequest{}

	if categoryLogDetails, ok := s.D.GetOkExists("category_log_details"); ok {
		if tmpList := categoryLogDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "category_log_details", 0)
			tmp, err := s.mapToCategoryLogDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CategoryLogDetails = &tmp
		}
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if modelDeploymentConfigurationDetails, ok := s.D.GetOkExists("model_deployment_configuration_details"); ok {
		if tmpList := modelDeploymentConfigurationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "model_deployment_configuration_details", 0)
			tmp, err := s.mapToModelDeploymentConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ModelDeploymentConfigurationDetails = tmp
		}
	}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.CreateModelDeployment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getModelDeploymentFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience"), oci_datascience.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DatascienceModelDeploymentResourceCrud) getModelDeploymentFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_datascience.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	modelDeploymentId, err := modelDeploymentWaitForWorkRequest(workId, "model-deployment",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, modelDeploymentId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_datascience.CancelWorkRequestRequest{
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
	s.D.SetId(*modelDeploymentId)

	return s.Get()
}

func modelDeploymentWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "datascience", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_datascience.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func modelDeploymentWaitForWorkRequest(wId *string, entityType string, action oci_datascience.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_datascience.DataScienceClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "datascience")
	retryPolicy.ShouldRetryOperation = modelDeploymentWorkRequestShouldRetryFunc(timeout)

	response := oci_datascience.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_datascience.WorkRequestStatusInProgress),
			string(oci_datascience.WorkRequestStatusAccepted),
			string(oci_datascience.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_datascience.WorkRequestStatusSucceeded),
			string(oci_datascience.WorkRequestStatusFailed),
			string(oci_datascience.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_datascience.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_datascience.WorkRequestStatusFailed || response.Status == oci_datascience.WorkRequestStatusCanceled {
		return nil, getErrorFromDatascienceModelDeploymentWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDatascienceModelDeploymentWorkRequest(client *oci_datascience.DataScienceClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_datascience.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_datascience.ListWorkRequestErrorsRequest{
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

func (s *DatascienceModelDeploymentResourceCrud) Get() error {
	request := oci_datascience.GetModelDeploymentRequest{}

	tmp := s.D.Id()
	request.ModelDeploymentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.GetModelDeployment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ModelDeployment
	return nil
}

func (s *DatascienceModelDeploymentResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_datascience.UpdateModelDeploymentRequest{}

	if categoryLogDetails, ok := s.D.GetOkExists("category_log_details"); ok {
		if tmpList := categoryLogDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "category_log_details", 0)
			tmp, err := s.mapToUpdateCategoryLogDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CategoryLogDetails = &tmp
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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if modelDeploymentConfigurationDetails, ok := s.D.GetOkExists("model_deployment_configuration_details"); ok {
		if tmpList := modelDeploymentConfigurationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "model_deployment_configuration_details", 0)
			tmp, err := s.mapToModelDeploymentConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ModelDeploymentConfigurationDetails = tmp
		}
	}

	tmp := s.D.Id()
	request.ModelDeploymentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.UpdateModelDeployment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getModelDeploymentFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience"), oci_datascience.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatascienceModelDeploymentResourceCrud) Delete() error {
	request := oci_datascience.DeleteModelDeploymentRequest{}

	tmp := s.D.Id()
	request.ModelDeploymentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.DeleteModelDeployment(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := modelDeploymentWaitForWorkRequest(workId, "model-deployment",
		oci_datascience.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DatascienceModelDeploymentResourceCrud) SetData() error {
	if s.Res.CategoryLogDetails != nil {
		s.D.Set("category_log_details", []interface{}{CategoryLogDetailsToMap(s.Res.CategoryLogDetails)})
	} else {
		s.D.Set("category_log_details", nil)
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

	if s.Res.ModelDeploymentConfigurationDetails != nil {
		modelDeploymentConfigurationDetailsArray := []interface{}{}
		if modelDeploymentConfigurationDetailsMap := ModelDeploymentConfigurationDetailsToMap(&s.Res.ModelDeploymentConfigurationDetails); modelDeploymentConfigurationDetailsMap != nil {
			modelDeploymentConfigurationDetailsArray = append(modelDeploymentConfigurationDetailsArray, modelDeploymentConfigurationDetailsMap)
		}
		s.D.Set("model_deployment_configuration_details", modelDeploymentConfigurationDetailsArray)
	} else {
		s.D.Set("model_deployment_configuration_details", nil)
	}

	if s.Res.ModelDeploymentUrl != nil {
		s.D.Set("model_deployment_url", *s.Res.ModelDeploymentUrl)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *DatascienceModelDeploymentResourceCrud) mapToCategoryLogDetails(fieldKeyFormat string) (oci_datascience.CategoryLogDetails, error) {
	result := oci_datascience.CategoryLogDetails{}

	if access, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "access")); ok {
		if tmpList := access.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "access"), 0)
			tmp, err := s.mapToLogDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert access, encountered error: %v", err)
			}
			result.Access = &tmp
		}
	}

	if predict, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "predict")); ok {
		if tmpList := predict.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "predict"), 0)
			tmp, err := s.mapToLogDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert predict, encountered error: %v", err)
			}
			result.Predict = &tmp
		}
	}

	return result, nil
}

func (s *DatascienceModelDeploymentResourceCrud) mapToUpdateCategoryLogDetails(fieldKeyFormat string) (oci_datascience.UpdateCategoryLogDetails, error) {
	result := oci_datascience.UpdateCategoryLogDetails{}

	if access, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "access")); ok {
		if tmpList := access.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "access"), 0)
			tmp, err := s.mapToLogDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert access, encountered error: %v", err)
			}
			result.Access = &tmp
		}
	}

	if predict, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "predict")); ok {
		if tmpList := predict.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "predict"), 0)
			tmp, err := s.mapToLogDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert predict, encountered error: %v", err)
			}
			result.Predict = &tmp
		}
	}

	return result, nil
}

func CategoryLogDetailsToMap(obj *oci_datascience.CategoryLogDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Access != nil {
		result["access"] = []interface{}{LogDetailsToMap(obj.Access)}
	}

	if obj.Predict != nil {
		result["predict"] = []interface{}{LogDetailsToMap(obj.Predict)}
	}

	return result
}

func (s *DatascienceModelDeploymentResourceCrud) mapToInstanceConfiguration(fieldKeyFormat string) (oci_datascience.InstanceConfiguration, error) {
	result := oci_datascience.InstanceConfiguration{}

	if instanceShapeName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_shape_name")); ok {
		tmp := instanceShapeName.(string)
		result.InstanceShapeName = &tmp
	}

	return result, nil
}

func InstanceConfigurationToMap(obj *oci_datascience.InstanceConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.InstanceShapeName != nil {
		result["instance_shape_name"] = string(*obj.InstanceShapeName)
	}

	return result
}

func (s *DatascienceModelDeploymentResourceCrud) mapToLogDetails(fieldKeyFormat string) (oci_datascience.LogDetails, error) {
	result := oci_datascience.LogDetails{}

	if logGroupId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "log_group_id")); ok {
		tmp := logGroupId.(string)
		result.LogGroupId = &tmp
	}

	if logId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "log_id")); ok {
		tmp := logId.(string)
		result.LogId = &tmp
	}

	return result, nil
}

func LogDetailsToMap(obj *oci_datascience.LogDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LogGroupId != nil {
		result["log_group_id"] = string(*obj.LogGroupId)
	}

	if obj.LogId != nil {
		result["log_id"] = string(*obj.LogId)
	}

	return result
}

func (s *DatascienceModelDeploymentResourceCrud) mapToModelConfigurationDetails(fieldKeyFormat string) (oci_datascience.ModelConfigurationDetails, error) {
	result := oci_datascience.ModelConfigurationDetails{}

	if bandwidthMbps, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bandwidth_mbps")); ok {
		tmp := bandwidthMbps.(int)
		result.BandwidthMbps = &tmp
	}

	if instanceConfiguration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_configuration")); ok {
		if tmpList := instanceConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "instance_configuration"), 0)
			tmp, err := s.mapToInstanceConfiguration(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert instance_configuration, encountered error: %v", err)
			}
			result.InstanceConfiguration = &tmp
		}
	}

	if modelId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_id")); ok {
		tmp := modelId.(string)
		result.ModelId = &tmp
	}

	if scalingPolicy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scaling_policy")); ok {
		if tmpList := scalingPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "scaling_policy"), 0)
			tmp, err := s.mapToScalingPolicy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert scaling_policy, encountered error: %v", err)
			}
			result.ScalingPolicy = tmp
		}
	}

	return result, nil
}

func ModelConfigurationDetailsToMap(obj *oci_datascience.ModelConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BandwidthMbps != nil {
		result["bandwidth_mbps"] = int(*obj.BandwidthMbps)
	}

	if obj.InstanceConfiguration != nil {
		result["instance_configuration"] = []interface{}{InstanceConfigurationToMap(obj.InstanceConfiguration)}
	}

	if obj.ModelId != nil {
		result["model_id"] = string(*obj.ModelId)
	}

	if obj.ScalingPolicy != nil {
		scalingPolicyArray := []interface{}{}
		if scalingPolicyMap := ScalingPolicyToMap(&obj.ScalingPolicy); scalingPolicyMap != nil {
			scalingPolicyArray = append(scalingPolicyArray, scalingPolicyMap)
		}
		result["scaling_policy"] = scalingPolicyArray
	}

	return result
}

func (s *DatascienceModelDeploymentResourceCrud) mapToModelDeploymentConfigurationDetails(fieldKeyFormat string) (oci_datascience.ModelDeploymentConfigurationDetails, error) {
	var baseObject oci_datascience.ModelDeploymentConfigurationDetails
	//discriminator
	deploymentTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "deployment_type"))
	var deploymentType string
	if ok {
		deploymentType = deploymentTypeRaw.(string)
	} else {
		deploymentType = "" // default value
	}
	switch strings.ToLower(deploymentType) {
	case strings.ToLower("SINGLE_MODEL"):
		details := oci_datascience.UpdateSingleModelDeploymentConfigurationDetails{}
		if modelConfigurationDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_configuration_details")); ok {
			if tmpList := modelConfigurationDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "model_configuration_details"), 0)
				tmp, err := s.mapToUpdateModelConfigurationDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert model_configuration_details, encountered error: %v", err)
				}
				details.ModelConfigurationDetails = &tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown deployment_type '%v' was specified", deploymentType)
	}
	return baseObject, nil
}

func ModelDeploymentConfigurationDetailsToMap(obj *oci_datascience.ModelDeploymentConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_datascience.SingleModelDeploymentConfigurationDetails:
		result["deployment_type"] = "SINGLE_MODEL"

		if v.ModelConfigurationDetails != nil {
			result["model_configuration_details"] = []interface{}{UpdateModelConfigurationDetailsToMap(v.ModelConfigurationDetails)}
		}
	default:
		log.Printf("[WARN] Received 'deployment_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatascienceModelDeploymentResourceCrud) mapToScalingPolicy(fieldKeyFormat string) (oci_datascience.ScalingPolicy, error) {
	var baseObject oci_datascience.ScalingPolicy
	//discriminator
	policyTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "policy_type"))
	var policyType string
	if ok {
		policyType = policyTypeRaw.(string)
	} else {
		policyType = "" // default value
	}
	switch strings.ToLower(policyType) {
	case strings.ToLower("FIXED_SIZE"):
		details := oci_datascience.FixedSizeScalingPolicy{}
		if instanceCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_count")); ok {
			tmp := instanceCount.(int)
			details.InstanceCount = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown policy_type '%v' was specified", policyType)
	}
	return baseObject, nil
}

func ScalingPolicyToMap(obj *oci_datascience.ScalingPolicy) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_datascience.FixedSizeScalingPolicy:
		result["policy_type"] = "FIXED_SIZE"

		if v.InstanceCount != nil {
			result["instance_count"] = int(*v.InstanceCount)
		}
	default:
		log.Printf("[WARN] Received 'policy_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatascienceModelDeploymentResourceCrud) mapToUpdateModelConfigurationDetails(fieldKeyFormat string) (oci_datascience.UpdateModelConfigurationDetails, error) {
	result := oci_datascience.UpdateModelConfigurationDetails{}

	if bandwidthMbps, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bandwidth_mbps")); ok {
		tmp := bandwidthMbps.(int)
		result.BandwidthMbps = &tmp
	}

	if instanceConfiguration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_configuration")); ok {
		if tmpList := instanceConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "instance_configuration"), 0)
			tmp, err := s.mapToInstanceConfiguration(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert instance_configuration, encountered error: %v", err)
			}
			result.InstanceConfiguration = &tmp
		}
	}

	if modelId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_id")); ok {
		tmp := modelId.(string)
		result.ModelId = &tmp
	}

	if scalingPolicy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scaling_policy")); ok {
		if tmpList := scalingPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "scaling_policy"), 0)
			tmp, err := s.mapToScalingPolicy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert scaling_policy, encountered error: %v", err)
			}
			result.ScalingPolicy = tmp
		}
	}

	return result, nil
}

func UpdateModelConfigurationDetailsToMap(obj *oci_datascience.ModelConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BandwidthMbps != nil {
		result["bandwidth_mbps"] = int(*obj.BandwidthMbps)
	}

	if obj.InstanceConfiguration != nil {
		result["instance_configuration"] = []interface{}{InstanceConfigurationToMap(obj.InstanceConfiguration)}
	}

	if obj.ModelId != nil {
		result["model_id"] = string(*obj.ModelId)
	}

	if obj.ScalingPolicy != nil {
		scalingPolicyArray := []interface{}{}
		if scalingPolicyMap := ScalingPolicyToMap(&obj.ScalingPolicy); scalingPolicyMap != nil {
			scalingPolicyArray = append(scalingPolicyArray, scalingPolicyMap)
		}
		result["scaling_policy"] = scalingPolicyArray
	}

	return result
}

func (s *DatascienceModelDeploymentResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_datascience.ChangeModelDeploymentCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ModelDeploymentId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.ChangeModelDeploymentCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
