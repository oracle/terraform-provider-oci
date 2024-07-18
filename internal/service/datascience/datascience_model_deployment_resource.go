// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"
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
												"model_deployment_instance_shape_config_details": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"cpu_baseline": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
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

															// Computed
														},
													},
												},
												"subnet_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

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
									"maximum_bandwidth_mbps": {
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
																					// Required

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

																					// Computed
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
																					// Required

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

																					// Computed
																				},
																			},
																		},

																		// Optional
																		"metric_type": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
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

												// Computed
											},
										},
									},

									// Computed
								},
							},
						},

						// Optional
						"environment_configuration_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"environment_configuration_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"DEFAULT",
											"OCIR_CONTAINER",
										}, true),
									},

									// Optional
									"cmd": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"entrypoint": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"environment_variables": {
										Type:     schema.TypeMap,
										Optional: true,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"health_check_port": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"image": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"image_digest": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"server_port": {
										Type:     schema.TypeInt,
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
			"opc_parent_rpt_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_datascience.ModelDeploymentLifecycleStateInactive),
					string(oci_datascience.ModelDeploymentLifecycleStateActive),
				}, true),
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
			"model_deployment_system_data": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"current_instance_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"system_infra_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"model_deployment_url": {
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
	var powerOff = false
	if powerState, ok := sync.D.GetOkExists("state"); ok {
		wantedPowerState := oci_datascience.ModelDeploymentLifecycleStateEnum(strings.ToUpper(powerState.(string)))
		if wantedPowerState == oci_datascience.ModelDeploymentLifecycleStateInactive {
			powerOff = true
		}
	}

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if powerOff {
		if err := sync.StopModelDeployment(); err != nil {
			return err
		}
		sync.D.Set("state", oci_datascience.ModelDeploymentLifecycleStateInactive)
	}
	return nil

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

	powerOn, powerOff := false, false

	if sync.D.HasChange("state") {
		wantedState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_datascience.ModelDeploymentLifecycleStateActive == oci_datascience.ModelDeploymentLifecycleStateEnum(wantedState) {
			powerOn = true
		} else if oci_datascience.ModelDeploymentLifecycleStateInactive == oci_datascience.ModelDeploymentLifecycleStateEnum(wantedState) {
			powerOff = true
		}
	}

	if powerOn {
		if err := sync.StartModelDeployment(); err != nil {
			return err
		}
		sync.D.Set("state", oci_datascience.ModelDeploymentLifecycleStateActive)
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	if powerOff {
		if err := sync.StopModelDeployment(); err != nil {
			return err
		}
		sync.D.Set("state", oci_datascience.ModelDeploymentLifecycleStateInactive)
	}

	return nil
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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
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

	if opcParentRptUrl, ok := s.D.GetOkExists("opc_parent_rpt_url"); ok {
		tmp := opcParentRptUrl.(string)
		request.OpcParentRptUrl = &tmp
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
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
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

	if s.Res.ModelDeploymentSystemData != nil {
		modelDeploymentSystemDataArray := []interface{}{}
		if modelDeploymentSystemDataMap := ModelDeploymentSystemDataToMap(&s.Res.ModelDeploymentSystemData); modelDeploymentSystemDataMap != nil {
			modelDeploymentSystemDataArray = append(modelDeploymentSystemDataArray, modelDeploymentSystemDataMap)
		}
		s.D.Set("model_deployment_system_data", modelDeploymentSystemDataArray)
	} else {
		s.D.Set("model_deployment_system_data", nil)
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

func (s *DatascienceModelDeploymentResourceCrud) StartModelDeployment() error {
	request := oci_datascience.ActivateModelDeploymentRequest{}

	idTmp := s.D.Id()
	request.ModelDeploymentId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.ActivateModelDeployment(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_datascience.ModelDeploymentLifecycleStateActive }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatascienceModelDeploymentResourceCrud) StopModelDeployment() error {
	request := oci_datascience.DeactivateModelDeploymentRequest{}

	idTmp := s.D.Id()
	request.ModelDeploymentId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.DeactivateModelDeployment(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_datascience.ModelDeploymentLifecycleStateInactive }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatascienceModelDeploymentResourceCrud) mapToAutoScalingPolicyDetails(fieldKeyFormat string) (oci_datascience.AutoScalingPolicyDetails, error) {
	var baseObject oci_datascience.AutoScalingPolicyDetails
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
		details := oci_datascience.ThresholdBasedAutoScalingPolicyDetails{}
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
			tmp := make([]oci_datascience.MetricExpressionRule, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "rules"), stateDataIndex)
				converted, err := s.mapToMetricExpressionRule(fieldKeyFormatNextLevel)
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

func AutoScalingPolicyDetailsToMap(obj oci_datascience.AutoScalingPolicyDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_datascience.ThresholdBasedAutoScalingPolicyDetails:
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
			rules = append(rules, MetricExpressionRuleToMap(item))
		}
		result["rules"] = rules
	default:
		log.Printf("[WARN] Received 'auto_scaling_policy_type' of unknown type %v", obj)
		return nil
	}

	return result
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

func (s *DatascienceModelDeploymentResourceCrud) mapToCustomExpressionQueryScalingConfiguration(fieldKeyFormat string) (oci_datascience.CustomExpressionQueryScalingConfiguration, error) {
	result := oci_datascience.CustomExpressionQueryScalingConfiguration{}

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

func CustomExpressionQueryScalingConfigurationToMap(obj *oci_datascience.CustomExpressionQueryScalingConfiguration) map[string]interface{} {
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

func (s *DatascienceModelDeploymentResourceCrud) mapToInstanceConfiguration(fieldKeyFormat string) (oci_datascience.InstanceConfiguration, error) {
	result := oci_datascience.InstanceConfiguration{}

	if instanceShapeName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_shape_name")); ok {
		tmp := instanceShapeName.(string)
		result.InstanceShapeName = &tmp
	}

	if modelDeploymentInstanceShapeConfigDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_deployment_instance_shape_config_details")); ok {
		if tmpList := modelDeploymentInstanceShapeConfigDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "model_deployment_instance_shape_config_details"), 0)
			tmp, err := s.mapToModelDeploymentInstanceShapeConfigDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert model_deployment_instance_shape_config_details, encountered error: %v", err)
			}
			result.ModelDeploymentInstanceShapeConfigDetails = &tmp
		}
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	return result, nil
}

func InstanceConfigurationToMap(obj *oci_datascience.InstanceConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.InstanceShapeName != nil {
		result["instance_shape_name"] = string(*obj.InstanceShapeName)
	}

	if obj.ModelDeploymentInstanceShapeConfigDetails != nil {
		result["model_deployment_instance_shape_config_details"] = []interface{}{ModelDeploymentInstanceShapeConfigDetailsToMap(obj.ModelDeploymentInstanceShapeConfigDetails)}
	}

	if obj.SubnetId != nil {
		if *obj.SubnetId == "" {
			result["subnet_id"] = nil
		} else {
			result["subnet_id"] = string(*obj.SubnetId)
		}
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

func (s *DatascienceModelDeploymentResourceCrud) mapToMetricExpressionRule(fieldKeyFormat string) (oci_datascience.MetricExpressionRule, error) {
	var baseObject oci_datascience.MetricExpressionRule
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
		details := oci_datascience.CustomMetricExpressionRule{}
		if scaleInConfiguration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scale_in_configuration")); ok {
			if tmpList := scaleInConfiguration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "scale_in_configuration"), 0)
				tmp, err := s.mapToCustomExpressionQueryScalingConfiguration(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert scale_in_configuration, encountered error: %v", err)
				}
				details.ScaleInConfiguration = &tmp
			}
		}
		if scaleOutConfiguration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scale_out_configuration")); ok {
			if tmpList := scaleOutConfiguration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "scale_out_configuration"), 0)
				tmp, err := s.mapToCustomExpressionQueryScalingConfiguration(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert scale_out_configuration, encountered error: %v", err)
				}
				details.ScaleOutConfiguration = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("PREDEFINED_EXPRESSION"):
		details := oci_datascience.PredefinedMetricExpressionRule{}
		if metricType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metric_type")); ok {
			details.MetricType = oci_datascience.PredefinedMetricExpressionRuleMetricTypeEnum(metricType.(string))
		}
		if scaleInConfiguration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scale_in_configuration")); ok {
			if tmpList := scaleInConfiguration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "scale_in_configuration"), 0)
				tmp, err := s.mapToPredefinedExpressionThresholdScalingConfiguration(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert scale_in_configuration, encountered error: %v", err)
				}
				details.ScaleInConfiguration = &tmp
			}
		}
		if scaleOutConfiguration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scale_out_configuration")); ok {
			if tmpList := scaleOutConfiguration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "scale_out_configuration"), 0)
				tmp, err := s.mapToPredefinedExpressionThresholdScalingConfiguration(fieldKeyFormatNextLevel)
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

func MetricExpressionRuleToMap(obj oci_datascience.MetricExpressionRule) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_datascience.CustomMetricExpressionRule:
		result["metric_expression_rule_type"] = "CUSTOM_EXPRESSION"

		if v.ScaleInConfiguration != nil {
			result["scale_in_configuration"] = []interface{}{CustomExpressionQueryScalingConfigurationToMap(v.ScaleInConfiguration)}
		}

		if v.ScaleOutConfiguration != nil {
			result["scale_out_configuration"] = []interface{}{CustomExpressionQueryScalingConfigurationToMap(v.ScaleOutConfiguration)}
		}
	case oci_datascience.PredefinedMetricExpressionRule:
		result["metric_expression_rule_type"] = "PREDEFINED_EXPRESSION"

		result["metric_type"] = string(v.MetricType)

		if v.ScaleInConfiguration != nil {
			result["scale_in_configuration"] = []interface{}{PredefinedExpressionThresholdScalingConfigurationToMap(v.ScaleInConfiguration)}
		}

		if v.ScaleOutConfiguration != nil {
			result["scale_out_configuration"] = []interface{}{PredefinedExpressionThresholdScalingConfigurationToMap(v.ScaleOutConfiguration)}
		}
	default:
		log.Printf("[WARN] Received 'metric_expression_rule_type' of unknown type %v", obj)
		return nil
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

	if maximumBandwidthMbps, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "maximum_bandwidth_mbps")); ok {
		tmp := maximumBandwidthMbps.(int)
		if tmp == 0 {
			result.MaximumBandwidthMbps = nil
		} else {
			result.MaximumBandwidthMbps = &tmp
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

	if maximumBandwidthMbps, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "maximum_bandwidth_mbps")); ok {
		tmp := maximumBandwidthMbps.(int)
		if tmp == 0 {
			result.MaximumBandwidthMbps = nil
		} else {
			result.MaximumBandwidthMbps = &tmp
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

	if obj.MaximumBandwidthMbps != nil {
		tmp := int(*obj.MaximumBandwidthMbps)
		if tmp == 0 {
			result["maximum_bandwidth_mbps"] = nil
		}
		result["maximum_bandwidth_mbps"] = tmp
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
		if environmentConfigurationDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "environment_configuration_details")); ok {
			if tmpList := environmentConfigurationDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "environment_configuration_details"), 0)
				tmp, err := s.mapToUpdateModelDeploymentEnvironmentConfigurationDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert environment_configuration_details, encountered error: %v", err)
				}
				details.EnvironmentConfigurationDetails = tmp
			}
		}
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

func (s *DatascienceModelDeploymentResourceCrud) mapToUpdateModelDeploymentConfigurationDetails(fieldKeyFormat string) (oci_datascience.UpdateModelDeploymentConfigurationDetails, error) {
	var baseObject oci_datascience.UpdateModelDeploymentConfigurationDetails
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
		if environmentConfigurationDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "environment_configuration_details")); ok {
			if tmpList := environmentConfigurationDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "environment_configuration_details"), 0)
				tmp, err := s.mapToUpdateModelDeploymentEnvironmentConfigurationDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert environment_configuration_details, encountered error: %v", err)
				}
				details.EnvironmentConfigurationDetails = tmp
			}
		}
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

		if v.EnvironmentConfigurationDetails != nil {
			environmentConfigurationDetailsArray := []interface{}{}
			if environmentConfigurationDetailsMap := ModelDeploymentEnvironmentConfigurationDetailsToMap(&v.EnvironmentConfigurationDetails); environmentConfigurationDetailsMap != nil {
				environmentConfigurationDetailsArray = append(environmentConfigurationDetailsArray, environmentConfigurationDetailsMap)
			}
			result["environment_configuration_details"] = environmentConfigurationDetailsArray
		}

		if v.ModelConfigurationDetails != nil {
			result["model_configuration_details"] = []interface{}{UpdateModelConfigurationDetailsToMap(v.ModelConfigurationDetails)}
		}
	default:
		log.Printf("[WARN] Received 'deployment_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatascienceModelDeploymentResourceCrud) mapToModelDeploymentEnvironmentConfigurationDetails(fieldKeyFormat string) (oci_datascience.ModelDeploymentEnvironmentConfigurationDetails, error) {
	var baseObject oci_datascience.ModelDeploymentEnvironmentConfigurationDetails
	//discriminator
	environmentConfigurationTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "environment_configuration_type"))
	var environmentConfigurationType string
	if ok {
		environmentConfigurationType = environmentConfigurationTypeRaw.(string)
	} else {
		environmentConfigurationType = "" // default value
	}
	switch strings.ToLower(environmentConfigurationType) {
	case strings.ToLower("DEFAULT"):
		details := oci_datascience.UpdateDefaultModelDeploymentEnvironmentConfigurationDetails{}
		if environmentVariables, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "environment_variables")); ok {
			details.EnvironmentVariables = tfresource.ObjectMapToStringMap(environmentVariables.(map[string]interface{}))
		}
		baseObject = details
	case strings.ToLower("OCIR_CONTAINER"):
		details := oci_datascience.UpdateOcirModelDeploymentEnvironmentConfigurationDetails{}
		if cmd, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cmd")); ok {
			interfaces := cmd.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "cmd")) {
				details.Cmd = tmp
			}
		}
		if entrypoint, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "entrypoint")); ok {
			interfaces := entrypoint.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "entrypoint")) {
				details.Entrypoint = tmp
			}
		}
		if environmentVariables, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "environment_variables")); ok {
			details.EnvironmentVariables = tfresource.ObjectMapToStringMap(environmentVariables.(map[string]interface{}))
		}
		if healthCheckPort, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "health_check_port")); ok {
			tmp := healthCheckPort.(int)
			details.HealthCheckPort = &tmp
		}
		if image, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "image")); ok {
			tmp := image.(string)
			details.Image = &tmp
		}
		if imageDigest, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "image_digest")); ok {
			tmp := imageDigest.(string)
			details.ImageDigest = &tmp
		}
		if serverPort, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "server_port")); ok {
			tmp := serverPort.(int)
			details.ServerPort = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown environment_configuration_type '%v' was specified", environmentConfigurationType)
	}
	return baseObject, nil
}

func (s *DatascienceModelDeploymentResourceCrud) mapToUpdateModelDeploymentEnvironmentConfigurationDetails(fieldKeyFormat string) (oci_datascience.UpdateModelDeploymentEnvironmentConfigurationDetails, error) {
	var baseObject oci_datascience.UpdateModelDeploymentEnvironmentConfigurationDetails
	//discriminator
	environmentConfigurationTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "environment_configuration_type"))
	var environmentConfigurationType string
	if ok {
		environmentConfigurationType = environmentConfigurationTypeRaw.(string)
	} else {
		environmentConfigurationType = "" // default value
	}
	switch strings.ToLower(environmentConfigurationType) {
	case strings.ToLower("DEFAULT"):
		details := oci_datascience.UpdateDefaultModelDeploymentEnvironmentConfigurationDetails{}
		if environmentVariables, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "environment_variables")); ok {
			details.EnvironmentVariables = tfresource.ObjectMapToStringMap(environmentVariables.(map[string]interface{}))
		}
		baseObject = details
	case strings.ToLower("OCIR_CONTAINER"):
		details := oci_datascience.UpdateOcirModelDeploymentEnvironmentConfigurationDetails{}
		if cmd, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cmd")); ok {
			interfaces := cmd.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "cmd")) {
				details.Cmd = tmp
			}
		}
		if entrypoint, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "entrypoint")); ok {
			interfaces := entrypoint.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "entrypoint")) {
				details.Entrypoint = tmp
			}
		}
		if environmentVariables, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "environment_variables")); ok {
			details.EnvironmentVariables = tfresource.ObjectMapToStringMap(environmentVariables.(map[string]interface{}))
		}
		if healthCheckPort, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "health_check_port")); ok {
			tmp := healthCheckPort.(int)
			details.HealthCheckPort = &tmp
		}
		if image, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "image")); ok {
			tmp := image.(string)
			details.Image = &tmp
		}
		if imageDigest, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "image_digest")); ok {
			tmp := imageDigest.(string)
			details.ImageDigest = &tmp
		}
		if serverPort, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "server_port")); ok {
			tmp := serverPort.(int)
			details.ServerPort = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown environment_configuration_type '%v' was specified", environmentConfigurationType)
	}
	return baseObject, nil
}

func ModelDeploymentEnvironmentConfigurationDetailsToMap(obj *oci_datascience.ModelDeploymentEnvironmentConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_datascience.DefaultModelDeploymentEnvironmentConfigurationDetails:
		result["environment_configuration_type"] = "DEFAULT"

		result["environment_variables"] = v.EnvironmentVariables
		result["environment_variables"] = v.EnvironmentVariables
	case oci_datascience.OcirModelDeploymentEnvironmentConfigurationDetails:
		result["environment_configuration_type"] = "OCIR_CONTAINER"

		result["cmd"] = v.Cmd
		result["cmd"] = v.Cmd

		result["entrypoint"] = v.Entrypoint
		result["entrypoint"] = v.Entrypoint

		result["environment_variables"] = v.EnvironmentVariables
		result["environment_variables"] = v.EnvironmentVariables

		if v.HealthCheckPort != nil {
			result["health_check_port"] = int(*v.HealthCheckPort)
		}

		if v.Image != nil {
			result["image"] = string(*v.Image)
		}

		if v.ImageDigest != nil {
			result["image_digest"] = string(*v.ImageDigest)
		}

		if v.ServerPort != nil {
			result["server_port"] = int(*v.ServerPort)
		}
	default:
		log.Printf("[WARN] Received 'environment_configuration_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatascienceModelDeploymentResourceCrud) mapToModelDeploymentInstanceShapeConfigDetails(fieldKeyFormat string) (oci_datascience.ModelDeploymentInstanceShapeConfigDetails, error) {
	result := oci_datascience.ModelDeploymentInstanceShapeConfigDetails{}

	if cpuBaseline, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cpu_baseline")); ok {
		result.CpuBaseline = oci_datascience.ModelDeploymentInstanceShapeConfigDetailsCpuBaselineEnum(cpuBaseline.(string))
	}

	memoryInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "memory_in_gbs"))
	if ok {
		tmp := float32(memoryInGBs.(float64))
		result.MemoryInGBs = &tmp
	} else {
		return result, fmt.Errorf("memory_in_gbs is required parameter")
	}

	ocpus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpus"))
	if ok {
		tmp := float32(ocpus.(float64))
		result.Ocpus = &tmp
	} else {
		return result, fmt.Errorf("ocpus is required parameter")
	}

	return result, nil
}

func ModelDeploymentInstanceShapeConfigDetailsToMap(obj *oci_datascience.ModelDeploymentInstanceShapeConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["cpu_baseline"] = string(obj.CpuBaseline)

	if obj.MemoryInGBs != nil {
		result["memory_in_gbs"] = float32(*obj.MemoryInGBs)
	}

	if obj.Ocpus != nil {
		result["ocpus"] = float32(*obj.Ocpus)
	}

	return result
}

func ModelDeploymentSystemDataToMap(obj *oci_datascience.ModelDeploymentSystemData) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_datascience.InstancePoolModelDeploymentSystemData:
		result["system_infra_type"] = "INSTANCE_POOL"

		if v.CurrentInstanceCount != nil {
			result["current_instance_count"] = int(*v.CurrentInstanceCount)
		}
	default:
		log.Printf("[WARN] Received 'system_infra_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatascienceModelDeploymentResourceCrud) mapToPredefinedExpressionThresholdScalingConfiguration(fieldKeyFormat string) (oci_datascience.PredefinedExpressionThresholdScalingConfiguration, error) {
	result := oci_datascience.PredefinedExpressionThresholdScalingConfiguration{}

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

func PredefinedExpressionThresholdScalingConfigurationToMap(obj *oci_datascience.PredefinedExpressionThresholdScalingConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.InstanceCountAdjustment != nil {
		result["instance_count_adjustment"] = int(*obj.InstanceCountAdjustment)
	}

	if obj.PendingDuration != nil {
		result["pending_duration"] = string(*obj.PendingDuration)
	}

	result["scaling_configuration_type"] = "THRESHOLD"

	if obj.Threshold != nil {
		result["threshold"] = int(*obj.Threshold)
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
	case strings.ToLower("AUTOSCALING"):
		details := oci_datascience.AutoScalingPolicy{}
		if autoScalingPolicies, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auto_scaling_policies")); ok {
			interfaces := autoScalingPolicies.([]interface{})
			tmp := make([]oci_datascience.AutoScalingPolicyDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "auto_scaling_policies"), stateDataIndex)
				converted, err := s.mapToAutoScalingPolicyDetails(fieldKeyFormatNextLevel)
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
	case oci_datascience.AutoScalingPolicy:
		result["policy_type"] = "AUTOSCALING"

		autoScalingPolicies := []interface{}{}
		for _, item := range v.AutoScalingPolicies {
			autoScalingPolicies = append(autoScalingPolicies, AutoScalingPolicyDetailsToMap(item))
		}
		result["auto_scaling_policies"] = autoScalingPolicies

		if v.CoolDownInSeconds != nil {
			result["cool_down_in_seconds"] = int(*v.CoolDownInSeconds)
		}

		if v.IsEnabled != nil {
			result["is_enabled"] = bool(*v.IsEnabled)
		}
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

func UpdateModelConfigurationDetailsToMap(obj *oci_datascience.ModelConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BandwidthMbps != nil {
		result["bandwidth_mbps"] = int(*obj.BandwidthMbps)
	}

	if obj.InstanceConfiguration != nil {
		result["instance_configuration"] = []interface{}{InstanceConfigurationToMap(obj.InstanceConfiguration)}
	}

	if obj.MaximumBandwidthMbps != nil {
		tmp := int(*obj.MaximumBandwidthMbps)
		if tmp == 0 {
			result["maximum_bandwidth_mbps"] = nil
		}
		result["maximum_bandwidth_mbps"] = tmp
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

func UpdateModelDeploymentEnvironmentConfigurationDetailsToMap(obj *oci_datascience.UpdateModelDeploymentEnvironmentConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_datascience.UpdateDefaultModelDeploymentEnvironmentConfigurationDetails:
		result["environment_configuration_type"] = "DEFAULT"

		result["environment_variables"] = v.EnvironmentVariables
		result["environment_variables"] = v.EnvironmentVariables
	case oci_datascience.UpdateOcirModelDeploymentEnvironmentConfigurationDetails:
		result["environment_configuration_type"] = "OCIR_CONTAINER"

		result["cmd"] = v.Cmd
		result["cmd"] = v.Cmd

		result["entrypoint"] = v.Entrypoint
		result["entrypoint"] = v.Entrypoint

		result["environment_variables"] = v.EnvironmentVariables
		result["environment_variables"] = v.EnvironmentVariables

		if v.HealthCheckPort != nil {
			result["health_check_port"] = int(*v.HealthCheckPort)
		}

		if v.Image != nil {
			result["image"] = string(*v.Image)
		}

		if v.ImageDigest != nil {
			result["image_digest"] = string(*v.ImageDigest)
		}

		if v.ServerPort != nil {
			result["server_port"] = int(*v.ServerPort)
		}
	default:
		log.Printf("[WARN] Received 'environment_configuration_type' of unknown type %v", *obj)
		return nil
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
