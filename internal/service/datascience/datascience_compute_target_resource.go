// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatascienceComputeTargetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDatascienceComputeTargetWithContext,
		ReadContext:   readDatascienceComputeTargetWithContext,
		UpdateContext: updateDatascienceComputeTargetWithContext,
		DeleteContext: deleteDatascienceComputeTargetWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compute_configuration_details": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"compute_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"MANAGED_COMPUTE_CLUSTER",
							}, true),
						},
						"instance_configuration": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"instance_shape": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"boot_volume_size_in_gbs": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"instance_shape_details": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Optional
												"memory_in_gbs": {
													Type:     schema.TypeFloat,
													Optional: true,
													Computed: true,
												},
												"ocpus": {
													Type:     schema.TypeFloat,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},

						// Optional
						"scaling_policy": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"policy_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"AUTOSCALING",
											"FIXED_SIZE",
										}, true),
									},

									// Optional
									"auto_scaling_policies": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"auto_scaling_policy_type": {
													Type:             schema.TypeString,
													Required:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"THRESHOLD",
													}, true),
												},
												"initial_instance_count": {
													Type:     schema.TypeInt,
													Required: true,
												},
												"maximum_instance_count": {
													Type:     schema.TypeInt,
													Required: true,
												},
												"minimum_instance_count": {
													Type:     schema.TypeInt,
													Required: true,
												},
												"rules": {
													Type:     schema.TypeList,
													Required: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"metric_expression_rule_type": {
																Type:             schema.TypeString,
																Required:         true,
																DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																ValidateFunc: validation.StringInSlice([]string{
																	"CUSTOM_EXPRESSION",
																	"PREDEFINED_EXPRESSION",
																}, true),
															},
															"scale_in_configuration": {
																Type:     schema.TypeList,
																Required: true,
																MaxItems: 1,
																MinItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Optional
																		"instance_count_adjustment": {
																			Type:     schema.TypeInt,
																			Optional: true,
																			Computed: true,
																		},
																		"pending_duration": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"query": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"scaling_configuration_type": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"threshold": {
																			Type:     schema.TypeInt,
																			Optional: true,
																			Computed: true,
																		},
																	},
																},
															},
															"scale_out_configuration": {
																Type:     schema.TypeList,
																Required: true,
																MaxItems: 1,
																MinItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Optional
																		"instance_count_adjustment": {
																			Type:     schema.TypeInt,
																			Optional: true,
																			Computed: true,
																		},
																		"pending_duration": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"query": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"scaling_configuration_type": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"threshold": {
																			Type:     schema.TypeInt,
																			Optional: true,
																			Computed: true,
																		},
																	},
																},
															},

															// Optional
															"metric_type": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"cool_down_in_seconds": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"instance_count": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"is_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
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
			"metadata": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"compute_target_system_data": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Computed
						"compute_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"current_instance_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
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
		},
	}
}

func createDatascienceComputeTargetWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatascienceComputeTargetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDatascienceComputeTargetWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatascienceComputeTargetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDatascienceComputeTargetWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatascienceComputeTargetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteDatascienceComputeTargetWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatascienceComputeTargetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DatascienceComputeTargetResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_datascience.DataScienceClient
	Res                    *oci_datascience.ComputeTarget
	DisableNotFoundRetries bool
}

func (s *DatascienceComputeTargetResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatascienceComputeTargetResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_datascience.ComputeTargetLifecycleStateCreating),
	}
}

func (s *DatascienceComputeTargetResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_datascience.ComputeTargetLifecycleStateActive),
		string(oci_datascience.ComputeTargetLifecycleStateNeedsAttention),
	}
}

func (s *DatascienceComputeTargetResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_datascience.ComputeTargetLifecycleStateDeleting),
	}
}

func (s *DatascienceComputeTargetResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_datascience.ComputeTargetLifecycleStateDeleted),
	}
}

func (s *DatascienceComputeTargetResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_datascience.CreateComputeTargetRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if computeConfigurationDetails, ok := s.D.GetOkExists("compute_configuration_details"); ok {
		if tmpList := computeConfigurationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "compute_configuration_details", 0)
			tmp, err := s.mapToComputeConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ComputeConfigurationDetails = tmp
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

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		request.Metadata = tfresource.ObjectMapToStringMap(metadata.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.CreateComputeTarget(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getComputeTargetFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience"), oci_datascience.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DatascienceComputeTargetResourceCrud) getComputeTargetFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_datascience.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	computeTargetId, err := computeTargetWaitForWorkRequest(ctx, workId, "compute-target",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, computeTargetId)
		_, cancelErr := s.Client.CancelWorkRequest(ctx,
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
	s.D.SetId(*computeTargetId)

	return s.GetWithContext(ctx)
}

func computeTargetWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func computeTargetWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_datascience.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_datascience.DataScienceClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "datascience")
	retryPolicy.ShouldRetryOperation = computeTargetWorkRequestShouldRetryFunc(timeout)

	response := oci_datascience.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
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
			response, err = client.GetWorkRequest(ctx,
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
		return nil, getErrorFromDatascienceComputeTargetWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDatascienceComputeTargetWorkRequest(ctx context.Context, client *oci_datascience.DataScienceClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_datascience.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
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

func (s *DatascienceComputeTargetResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_datascience.GetComputeTargetRequest{}

	tmp := s.D.Id()
	request.ComputeTargetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.GetComputeTarget(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.ComputeTarget
	return nil
}

func (s *DatascienceComputeTargetResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_datascience.UpdateComputeTargetRequest{}

	if computeConfigurationDetails, ok := s.D.GetOkExists("compute_configuration_details"); ok {
		if tmpList := computeConfigurationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "compute_configuration_details", 0)
			tmp, err := s.mapToUpdateComputeConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ComputeConfigurationDetails = tmp
		}
	}

	tmp := s.D.Id()
	request.ComputeTargetId = &tmp

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

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		request.Metadata = tfresource.ObjectMapToStringMap(metadata.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.UpdateComputeTarget(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getComputeTargetFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience"), oci_datascience.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatascienceComputeTargetResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_datascience.DeleteComputeTargetRequest{}

	tmp := s.D.Id()
	request.ComputeTargetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.DeleteComputeTarget(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := computeTargetWaitForWorkRequest(ctx, workId, "compute-target",
		oci_datascience.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DatascienceComputeTargetResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComputeConfigurationDetails != nil {
		computeConfigurationDetailsArray := []interface{}{}
		if computeConfigurationDetailsMap := ComputeConfigurationDetailsToMap(&s.Res.ComputeConfigurationDetails); computeConfigurationDetailsMap != nil {
			computeConfigurationDetailsArray = append(computeConfigurationDetailsArray, computeConfigurationDetailsMap)
		}
		s.D.Set("compute_configuration_details", computeConfigurationDetailsArray)
	} else {
		s.D.Set("compute_configuration_details", nil)
	}

	if s.Res.ComputeTargetSystemData != nil {
		computeTargetSystemDataArray := []interface{}{}
		if computeTargetSystemDataMap := ComputeTargetSystemDataToMap(&s.Res.ComputeTargetSystemData); computeTargetSystemDataMap != nil {
			computeTargetSystemDataArray = append(computeTargetSystemDataArray, computeTargetSystemDataMap)
		}
		s.D.Set("compute_target_system_data", computeTargetSystemDataArray)
	} else {
		s.D.Set("compute_target_system_data", nil)
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

	s.D.Set("metadata", s.Res.Metadata)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *DatascienceComputeTargetResourceCrud) mapToComputeConfigurationDetails(fieldKeyFormat string) (oci_datascience.ComputeConfigurationDetails, error) {
	var baseObject oci_datascience.ComputeConfigurationDetails
	//discriminator
	computeTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compute_type"))
	var computeType string
	if ok {
		computeType = computeTypeRaw.(string)
	} else {
		computeType = "" // default value
	}
	switch strings.ToLower(computeType) {
	case strings.ToLower("MANAGED_COMPUTE_CLUSTER"):
		details := oci_datascience.UpdateManagedComputeClusterConfigurationDetails{}
		if instanceConfiguration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_configuration")); ok {
			if tmpList := instanceConfiguration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "instance_configuration"), 0)
				tmp, err := s.mapToUpdateManagedComputeClusterInstanceConfigurationDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert instance_configuration, encountered error: %v", err)
				}
				details.InstanceConfiguration = &tmp
			}
		}
		if scalingPolicy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scaling_policy")); ok {
			if tmpList := scalingPolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "scaling_policy"), 0)
				tmp, err := s.mapToManagedComputeClusterScalingPolicy(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert scaling_policy, encountered error: %v", err)
				}
				details.ScalingPolicy = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown compute_type '%v' was specified", computeType)
	}
	return baseObject, nil
}

func (s *DatascienceComputeTargetResourceCrud) mapToUpdateComputeConfigurationDetails(fieldKeyFormat string) (oci_datascience.UpdateComputeConfigurationDetails, error) {
	var baseObject oci_datascience.UpdateComputeConfigurationDetails
	//discriminator
	computeTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compute_type"))
	var computeType string
	if ok {
		computeType = computeTypeRaw.(string)
	} else {
		computeType = "" // default value
	}
	switch strings.ToLower(computeType) {
	case strings.ToLower("MANAGED_COMPUTE_CLUSTER"):
		details := oci_datascience.UpdateManagedComputeClusterConfigurationDetails{}
		if instanceConfiguration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_configuration")); ok {
			if tmpList := instanceConfiguration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "instance_configuration"), 0)
				tmp, err := s.mapToUpdateManagedComputeClusterInstanceConfigurationDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert instance_configuration, encountered error: %v", err)
				}
				details.InstanceConfiguration = &tmp
			}
		}
		if scalingPolicy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scaling_policy")); ok {
			if tmpList := scalingPolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "scaling_policy"), 0)
				tmp, err := s.mapToManagedComputeClusterScalingPolicy(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert scaling_policy, encountered error: %v", err)
				}
				details.ScalingPolicy = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown compute_type '%v' was specified", computeType)
	}
	return baseObject, nil
}

func ComputeConfigurationDetailsToMap(obj *oci_datascience.ComputeConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_datascience.ManagedComputeClusterComputeConfigurationDetails:
		result["compute_type"] = "MANAGED_COMPUTE_CLUSTER"

		if v.InstanceConfiguration != nil {
			result["instance_configuration"] = []interface{}{
				ManagedComputeClusterInstanceConfigurationDetailsToMap(v.InstanceConfiguration),
			}
		}

		if v.ScalingPolicy != nil {
			scalingPolicyArray := []interface{}{}
			if scalingPolicyMap := ManagedComputeClusterScalingPolicyToMap(&v.ScalingPolicy); scalingPolicyMap != nil {
				scalingPolicyArray = append(scalingPolicyArray, scalingPolicyMap)
			}
			result["scaling_policy"] = scalingPolicyArray
		}
	case oci_datascience.UpdateManagedComputeClusterConfigurationDetails:
		result["compute_type"] = "MANAGED_COMPUTE_CLUSTER"

		if v.InstanceConfiguration != nil {
			result["instance_configuration"] = []interface{}{UpdateManagedComputeClusterInstanceConfigurationDetailsToMap(v.InstanceConfiguration)}
		}

		if v.ScalingPolicy != nil {
			scalingPolicyArray := []interface{}{}
			if scalingPolicyMap := ManagedComputeClusterScalingPolicyToMap(&v.ScalingPolicy); scalingPolicyMap != nil {
				scalingPolicyArray = append(scalingPolicyArray, scalingPolicyMap)
			}
			result["scaling_policy"] = scalingPolicyArray
		}
	default:
		log.Printf("[WARN] Received 'compute_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func ComputeTargetSystemDataToMap(obj *oci_datascience.ComputeTargetSystemData) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_datascience.ManagedComputeClusterSystemData:
		result["compute_type"] = "MANAGED_COMPUTE_CLUSTER"

		if v.CurrentInstanceCount != nil {
			result["current_instance_count"] = int(*v.CurrentInstanceCount)
		}
	default:
		log.Printf("[WARN] Received 'compute_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatascienceComputeTargetResourceCrud) mapToManagedComputeClusterAutoScalingPolicyDetails(fieldKeyFormat string) (oci_datascience.ManagedComputeClusterAutoScalingPolicyDetails, error) {
	var baseObject oci_datascience.ManagedComputeClusterAutoScalingPolicyDetails
	//discriminator
	autoScalingPolicyTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auto_scaling_policy_type"))
	var autoScalingPolicyType string
	if ok {
		autoScalingPolicyType = autoScalingPolicyTypeRaw.(string)
	} else {
		autoScalingPolicyType = "" // default value
	}
	switch strings.ToLower(autoScalingPolicyType) {
	case strings.ToLower("THRESHOLD"):
		details := oci_datascience.ManagedComputeClusterThresholdBasedAutoScalingPolicyDetails{}
		if initialInstanceCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "initial_instance_count")); ok {
			tmp := initialInstanceCount.(int)
			details.InitialInstanceCount = &tmp
		}
		if maximumInstanceCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "maximum_instance_count")); ok {
			tmp := maximumInstanceCount.(int)
			details.MaximumInstanceCount = &tmp
		}
		if minimumInstanceCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "minimum_instance_count")); ok {
			tmp := minimumInstanceCount.(int)
			details.MinimumInstanceCount = &tmp
		}
		if rules, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rules")); ok {
			interfaces := rules.([]interface{})
			tmp := make([]oci_datascience.ManagedComputeClusterMetricExpressionRule, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "rules"), stateDataIndex)
				converted, err := s.mapToManagedComputeClusterMetricExpressionRule(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "rules")) {
				details.Rules = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown auto_scaling_policy_type '%v' was specified", autoScalingPolicyType)
	}
	return baseObject, nil
}

func ManagedComputeClusterAutoScalingPolicyDetailsToMap(obj oci_datascience.ManagedComputeClusterAutoScalingPolicyDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_datascience.ManagedComputeClusterThresholdBasedAutoScalingPolicyDetails:
		result["auto_scaling_policy_type"] = "THRESHOLD"

		if v.InitialInstanceCount != nil {
			result["initial_instance_count"] = int(*v.InitialInstanceCount)
		}

		if v.MaximumInstanceCount != nil {
			result["maximum_instance_count"] = int(*v.MaximumInstanceCount)
		}

		if v.MinimumInstanceCount != nil {
			result["minimum_instance_count"] = int(*v.MinimumInstanceCount)
		}

		rules := []interface{}{}
		for _, item := range v.Rules {
			rules = append(rules, ManagedComputeClusterMetricExpressionRuleToMap(item))
		}
		result["rules"] = rules
	default:
		log.Printf("[WARN] Received 'auto_scaling_policy_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *DatascienceComputeTargetResourceCrud) mapToManagedComputeClusterCustomExpressionQueryScalingConfiguration(fieldKeyFormat string) (oci_datascience.ManagedComputeClusterCustomExpressionQueryScalingConfiguration, error) {
	result := oci_datascience.ManagedComputeClusterCustomExpressionQueryScalingConfiguration{}

	if instanceCountAdjustment, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_count_adjustment")); ok {
		tmp := instanceCountAdjustment.(int)
		result.InstanceCountAdjustment = &tmp
	}

	if pendingDuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pending_duration")); ok {
		tmp := pendingDuration.(string)
		result.PendingDuration = &tmp
	}

	if query, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "query")); ok {
		tmp := query.(string)
		result.Query = &tmp
	}

	return result, nil
}

func ManagedComputeClusterCustomExpressionQueryScalingConfigurationToMap(obj *oci_datascience.ManagedComputeClusterCustomExpressionQueryScalingConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.InstanceCountAdjustment != nil {
		result["instance_count_adjustment"] = int(*obj.InstanceCountAdjustment)
	}

	if obj.PendingDuration != nil {
		result["pending_duration"] = string(*obj.PendingDuration)
	}

	if obj.Query != nil {
		result["query"] = string(*obj.Query)
	}

	result["scaling_configuration_type"] = "QUERY"

	return result
}

func (s *DatascienceComputeTargetResourceCrud) mapToManagedComputeClusterInstanceConfigurationDetails(fieldKeyFormat string) (oci_datascience.ManagedComputeClusterInstanceConfigurationDetails, error) {
	result := oci_datascience.ManagedComputeClusterInstanceConfigurationDetails{}

	if bootVolumeSizeInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "boot_volume_size_in_gbs")); ok {
		tmp := bootVolumeSizeInGBs.(int)
		result.BootVolumeSizeInGBs = &tmp
	}

	if instanceShape, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_shape")); ok {
		tmp := instanceShape.(string)
		result.InstanceShape = &tmp
	}

	if instanceShapeDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_shape_details")); ok {
		if tmpList := instanceShapeDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "instance_shape_details"), 0)
			tmp, err := s.mapToManagedComputeClusterInstanceShapeDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert instance_shape_details, encountered error: %v", err)
			}
			result.InstanceShapeDetails = &tmp
		}
	}

	return result, nil
}

func (s *DatascienceComputeTargetResourceCrud) mapToUpdateManagedComputeClusterInstanceConfigurationDetails(fieldKeyFormat string) (oci_datascience.UpdateManagedComputeClusterInstanceConfigurationDetails, error) {
	result := oci_datascience.UpdateManagedComputeClusterInstanceConfigurationDetails{}

	if bootVolumeSizeInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "boot_volume_size_in_gbs")); ok {
		tmp := bootVolumeSizeInGBs.(int)
		result.BootVolumeSizeInGBs = &tmp
	}

	if instanceShape, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_shape")); ok {
		tmp := instanceShape.(string)
		result.InstanceShape = &tmp
	}

	if instanceShapeDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_shape_details")); ok {
		if tmpList := instanceShapeDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "instance_shape_details"), 0)
			tmp, err := s.mapToManagedComputeClusterInstanceShapeDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert instance_shape_details, encountered error: %v", err)
			}
			result.InstanceShapeDetails = &tmp
		}
	}

	return result, nil
}

func ManagedComputeClusterInstanceConfigurationDetailsToMap(obj *oci_datascience.ManagedComputeClusterInstanceConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BootVolumeSizeInGBs != nil {
		result["boot_volume_size_in_gbs"] = int(*obj.BootVolumeSizeInGBs)
	}

	if obj.InstanceShape != nil {
		result["instance_shape"] = string(*obj.InstanceShape)
	}

	if obj.InstanceShapeDetails != nil {
		result["instance_shape_details"] = []interface{}{ManagedComputeClusterInstanceShapeDetailsToMap(obj.InstanceShapeDetails)}
	}

	return result
}

func (s *DatascienceComputeTargetResourceCrud) mapToManagedComputeClusterInstanceShapeDetails(fieldKeyFormat string) (oci_datascience.ManagedComputeClusterInstanceShapeDetails, error) {
	result := oci_datascience.ManagedComputeClusterInstanceShapeDetails{}

	if memoryInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "memory_in_gbs")); ok {
		tmp := float32(memoryInGBs.(float64))
		result.MemoryInGBs = &tmp
	}

	if ocpus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpus")); ok {
		tmp := float32(ocpus.(float64))
		result.Ocpus = &tmp
	}

	return result, nil
}

func ManagedComputeClusterInstanceShapeDetailsToMap(obj *oci_datascience.ManagedComputeClusterInstanceShapeDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MemoryInGBs != nil {
		result["memory_in_gbs"] = float32(*obj.MemoryInGBs)
	}

	if obj.Ocpus != nil {
		result["ocpus"] = float32(*obj.Ocpus)
	}

	return result
}

func (s *DatascienceComputeTargetResourceCrud) mapToManagedComputeClusterMetricExpressionRule(fieldKeyFormat string) (oci_datascience.ManagedComputeClusterMetricExpressionRule, error) {
	var baseObject oci_datascience.ManagedComputeClusterMetricExpressionRule
	//discriminator
	metricExpressionRuleTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metric_expression_rule_type"))
	var metricExpressionRuleType string
	if ok {
		metricExpressionRuleType = metricExpressionRuleTypeRaw.(string)
	} else {
		metricExpressionRuleType = "" // default value
	}
	switch strings.ToLower(metricExpressionRuleType) {
	case strings.ToLower("CUSTOM_EXPRESSION"):
		details := oci_datascience.ManagedComputeClusterCustomMetricExpressionRule{}
		if scaleInConfiguration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scale_in_configuration")); ok {
			if tmpList := scaleInConfiguration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "scale_in_configuration"), 0)
				tmp, err := s.mapToManagedComputeClusterCustomExpressionQueryScalingConfiguration(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert scale_in_configuration, encountered error: %v", err)
				}
				details.ScaleInConfiguration = &tmp
			}
		}
		if scaleOutConfiguration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scale_out_configuration")); ok {
			if tmpList := scaleOutConfiguration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "scale_out_configuration"), 0)
				tmp, err := s.mapToManagedComputeClusterCustomExpressionQueryScalingConfiguration(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert scale_out_configuration, encountered error: %v", err)
				}
				details.ScaleOutConfiguration = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("PREDEFINED_EXPRESSION"):
		details := oci_datascience.ManagedComputeClusterPredefinedMetricExpressionRule{}
		if metricType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metric_type")); ok {
			details.MetricType = oci_datascience.ManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeEnum(metricType.(string))
		}
		if scaleInConfiguration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scale_in_configuration")); ok {
			if tmpList := scaleInConfiguration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "scale_in_configuration"), 0)
				tmp, err := s.mapToManagedComputeClusterPredefinedExpressionThresholdScalingConfiguration(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert scale_in_configuration, encountered error: %v", err)
				}
				details.ScaleInConfiguration = &tmp
			}
		}
		if scaleOutConfiguration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scale_out_configuration")); ok {
			if tmpList := scaleOutConfiguration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "scale_out_configuration"), 0)
				tmp, err := s.mapToManagedComputeClusterPredefinedExpressionThresholdScalingConfiguration(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert scale_out_configuration, encountered error: %v", err)
				}
				details.ScaleOutConfiguration = &tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown metric_expression_rule_type '%v' was specified", metricExpressionRuleType)
	}
	return baseObject, nil
}

func ManagedComputeClusterMetricExpressionRuleToMap(obj oci_datascience.ManagedComputeClusterMetricExpressionRule) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_datascience.ManagedComputeClusterCustomMetricExpressionRule:
		result["metric_expression_rule_type"] = "CUSTOM_EXPRESSION"

		if v.ScaleInConfiguration != nil {
			result["scale_in_configuration"] = []interface{}{ManagedComputeClusterCustomExpressionQueryScalingConfigurationToMap(v.ScaleInConfiguration)}
		}

		if v.ScaleOutConfiguration != nil {
			result["scale_out_configuration"] = []interface{}{ManagedComputeClusterCustomExpressionQueryScalingConfigurationToMap(v.ScaleOutConfiguration)}
		}
	case oci_datascience.ManagedComputeClusterPredefinedMetricExpressionRule:
		result["metric_expression_rule_type"] = "PREDEFINED_EXPRESSION"

		result["metric_type"] = string(v.MetricType)

		if v.ScaleInConfiguration != nil {
			result["scale_in_configuration"] = []interface{}{ManagedComputeClusterPredefinedExpressionThresholdScalingConfigurationToMap(v.ScaleInConfiguration)}
		}

		if v.ScaleOutConfiguration != nil {
			result["scale_out_configuration"] = []interface{}{ManagedComputeClusterPredefinedExpressionThresholdScalingConfigurationToMap(v.ScaleOutConfiguration)}
		}
	default:
		log.Printf("[WARN] Received 'metric_expression_rule_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *DatascienceComputeTargetResourceCrud) mapToManagedComputeClusterPredefinedExpressionThresholdScalingConfiguration(fieldKeyFormat string) (oci_datascience.ManagedComputeClusterPredefinedExpressionThresholdScalingConfiguration, error) {
	result := oci_datascience.ManagedComputeClusterPredefinedExpressionThresholdScalingConfiguration{}

	if instanceCountAdjustment, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_count_adjustment")); ok {
		tmp := instanceCountAdjustment.(int)
		result.InstanceCountAdjustment = &tmp
	}

	if pendingDuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pending_duration")); ok {
		tmp := pendingDuration.(string)
		result.PendingDuration = &tmp
	}

	if threshold, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "threshold")); ok {
		tmp := threshold.(int)
		result.Threshold = &tmp
	}

	return result, nil
}

func ManagedComputeClusterPredefinedExpressionThresholdScalingConfigurationToMap(obj *oci_datascience.ManagedComputeClusterPredefinedExpressionThresholdScalingConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.InstanceCountAdjustment != nil {
		result["instance_count_adjustment"] = int(*obj.InstanceCountAdjustment)
	}

	if obj.PendingDuration != nil {
		result["pending_duration"] = string(*obj.PendingDuration)
	}

	if obj.Threshold != nil {
		result["threshold"] = int(*obj.Threshold)
	}

	return result
}

func (s *DatascienceComputeTargetResourceCrud) mapToManagedComputeClusterScalingPolicy(fieldKeyFormat string) (oci_datascience.ManagedComputeClusterScalingPolicy, error) {
	var baseObject oci_datascience.ManagedComputeClusterScalingPolicy
	//discriminator
	policyTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "policy_type"))
	var policyType string
	if ok {
		policyType = policyTypeRaw.(string)
	} else {
		policyType = "" // default value
	}
	switch strings.ToLower(policyType) {
	case strings.ToLower("AUTOSCALING"):
		details := oci_datascience.ManagedComputeClusterAutoScalingPolicy{}
		if autoScalingPolicies, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auto_scaling_policies")); ok {
			interfaces := autoScalingPolicies.([]interface{})
			tmp := make([]oci_datascience.ManagedComputeClusterAutoScalingPolicyDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "auto_scaling_policies"), stateDataIndex)
				converted, err := s.mapToManagedComputeClusterAutoScalingPolicyDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "auto_scaling_policies")) {
				details.AutoScalingPolicies = tmp
			}
		}
		if coolDownInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cool_down_in_seconds")); ok {
			tmp := coolDownInSeconds.(int)
			details.CoolDownInSeconds = &tmp
		}
		if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
			tmp := isEnabled.(bool)
			details.IsEnabled = &tmp
		}
		baseObject = details
	case strings.ToLower("FIXED_SIZE"):
		details := oci_datascience.ManagedComputeClusterFixedSizeScalingPolicy{}
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

func ManagedComputeClusterScalingPolicyToMap(obj *oci_datascience.ManagedComputeClusterScalingPolicy) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_datascience.ManagedComputeClusterAutoScalingPolicy:
		result["policy_type"] = "AUTOSCALING"

		autoScalingPolicies := []interface{}{}
		for _, item := range v.AutoScalingPolicies {
			autoScalingPolicies = append(autoScalingPolicies, ManagedComputeClusterAutoScalingPolicyDetailsToMap(item))
		}
		result["auto_scaling_policies"] = autoScalingPolicies

		if v.CoolDownInSeconds != nil {
			result["cool_down_in_seconds"] = int(*v.CoolDownInSeconds)
		}

		if v.IsEnabled != nil {
			result["is_enabled"] = bool(*v.IsEnabled)
		}
	case oci_datascience.ManagedComputeClusterFixedSizeScalingPolicy:
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

func UpdateManagedComputeClusterInstanceConfigurationDetailsToMap(obj *oci_datascience.UpdateManagedComputeClusterInstanceConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BootVolumeSizeInGBs != nil {
		result["boot_volume_size_in_gbs"] = int(*obj.BootVolumeSizeInGBs)
	}

	if obj.InstanceShape != nil {
		result["instance_shape"] = string(*obj.InstanceShape)
	}

	if obj.InstanceShapeDetails != nil {
		result["instance_shape_details"] = []interface{}{ManagedComputeClusterInstanceShapeDetailsToMap(obj.InstanceShapeDetails)}
	}

	return result
}

func (s *DatascienceComputeTargetResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_datascience.ChangeComputeTargetCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ComputeTargetId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.ChangeComputeTargetCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
