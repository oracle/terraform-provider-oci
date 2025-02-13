// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatascienceScheduleResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatascienceSchedule,
		Read:     readDatascienceSchedule,
		Update:   updateDatascienceSchedule,
		Delete:   deleteDatascienceSchedule,
		Schema: map[string]*schema.Schema{
			// Required
			"action": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"action_details": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"http_action_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"CREATE_JOB_RUN",
											"CREATE_PIPELINE_RUN",
											"INVOKE_ML_APPLICATION_PROVIDER_TRIGGER",
										}, true),
									},

									// Optional
									"create_job_run_details": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"compartment_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
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
												"job_configuration_override_details": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"job_type": {
																Type:             schema.TypeString,
																Required:         true,
																DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																ValidateFunc: validation.StringInSlice([]string{
																	"DEFAULT",
																}, true),
															},

															// Optional
															"command_line_arguments": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"environment_variables": {
																Type:     schema.TypeMap,
																Optional: true,
																Computed: true,
																Elem:     schema.TypeString,
															},
															"maximum_runtime_in_minutes": {
																Type:             schema.TypeString,
																Optional:         true,
																Computed:         true,
																ValidateFunc:     tfresource.ValidateInt64TypeString,
																DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
															},

															// Computed
														},
													},
												},
												"job_environment_configuration_override_details": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"image": {
																Type:     schema.TypeString,
																Required: true,
															},
															"job_environment_type": {
																Type:             schema.TypeString,
																Required:         true,
																DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																ValidateFunc: validation.StringInSlice([]string{
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
															"image_digest": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"image_signature_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},

															// Computed
														},
													},
												},
												"job_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"job_log_configuration_override_details": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"enable_auto_log_creation": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
															},
															"enable_logging": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
															},
															"log_group_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"log_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},

															// Computed
														},
													},
												},
												"project_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"create_pipeline_run_details": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"compartment_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"configuration_override_details": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"type": {
																Type:             schema.TypeString,
																Required:         true,
																DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																ValidateFunc: validation.StringInSlice([]string{
																	"DEFAULT",
																}, true),
															},

															// Optional
															"command_line_arguments": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"environment_variables": {
																Type:     schema.TypeMap,
																Optional: true,
																Computed: true,
																Elem:     schema.TypeString,
															},
															"maximum_runtime_in_minutes": {
																Type:             schema.TypeString,
																Optional:         true,
																Computed:         true,
																ValidateFunc:     tfresource.ValidateInt64TypeString,
																DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
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
												"log_configuration_override_details": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"enable_auto_log_creation": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
															},
															"enable_logging": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
															},
															"log_group_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"log_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},

															// Computed
														},
													},
												},
												"pipeline_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"project_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"step_override_details": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"step_configuration_details": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																MaxItems: 1,
																MinItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional
																		"command_line_arguments": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"environment_variables": {
																			Type:     schema.TypeMap,
																			Optional: true,
																			Computed: true,
																			Elem:     schema.TypeString,
																		},
																		"maximum_runtime_in_minutes": {
																			Type:             schema.TypeString,
																			Optional:         true,
																			Computed:         true,
																			ValidateFunc:     tfresource.ValidateInt64TypeString,
																			DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
																		},

																		// Computed
																	},
																},
															},
															"step_container_configuration_details": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																MaxItems: 1,
																MinItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required
																		"container_type": {
																			Type:             schema.TypeString,
																			Required:         true,
																			DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																			ValidateFunc: validation.StringInSlice([]string{
																				"OCIR_CONTAINER",
																			}, true),
																		},
																		"image": {
																			Type:     schema.TypeString,
																			Required: true,
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
																		"image_digest": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"image_signature_id": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},

																		// Computed
																	},
																},
															},
															"step_name": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},

															// Computed
														},
													},
												},
												"system_tags": {
													Type:     schema.TypeMap,
													Optional: true,
													Computed: true,
													Elem:     schema.TypeString,
												},

												// Computed
											},
										},
									},
									"ml_application_instance_view_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"trigger_ml_application_instance_view_flow_details": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"parameters": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"name": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"value": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},

															// Computed
														},
													},
												},
												"trigger_name": {
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
						"action_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"HTTP",
							}, true),
						},

						// Optional

						// Computed
					},
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"trigger": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"trigger_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"CRON",
								"ICAL",
								"INTERVAL",
							}, true),
						},

						// Optional
						"cron_expression": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"frequency": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"interval": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"is_random_start_time": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: false,
							Default:  false,
						},
						"recurrence": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"time_end": {
							Type:             schema.TypeString,
							Optional:         true,
							Default:          "",
							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
						},
						"time_start": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"log_details": {
				Type:     schema.TypeList,
				Optional: true,
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
					},
				},
			},

			// Computed
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_schedule_run_details": {
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
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_schedule_run": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_next_scheduled_run": {
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

func createDatascienceSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.CreateResource(d, sync)
}

func readDatascienceSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

func updateDatascienceSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatascienceSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatascienceScheduleResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_datascience.DataScienceClient
	Res                    *oci_datascience.Schedule
	DisableNotFoundRetries bool
}

func (s *DatascienceScheduleResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatascienceScheduleResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_datascience.ScheduleLifecycleStateCreating),
	}
}

func (s *DatascienceScheduleResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_datascience.ScheduleLifecycleStateActive),
	}
}

func (s *DatascienceScheduleResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_datascience.ScheduleLifecycleStateDeleting),
	}
}

func (s *DatascienceScheduleResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_datascience.ScheduleLifecycleStateDeleted),
	}
}

func (s *DatascienceScheduleResourceCrud) Create() error {
	request := oci_datascience.CreateScheduleRequest{}

	if action, ok := s.D.GetOkExists("action"); ok {
		if tmpList := action.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "action", 0)
			tmp, err := s.mapToScheduleAction(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Action = tmp
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

	if logDetails, ok := s.D.GetOkExists("log_details"); ok {
		if tmpList := logDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "log_details", 0)
			tmp, err := s.mapToScheduleLogDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LogDetails = &tmp
		}
	}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	if trigger, ok := s.D.GetOkExists("trigger"); ok {
		if tmpList := trigger.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "trigger", 0)
			tmp, err := s.mapToScheduleTrigger(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Trigger = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.CreateSchedule(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getScheduleFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience"), oci_datascience.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DatascienceScheduleResourceCrud) getScheduleFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_datascience.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	scheduleId, err := scheduleWaitForWorkRequest(workId, "schedule",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, scheduleId)
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
	s.D.SetId(*scheduleId)

	return s.Get()
}

func scheduleWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func scheduleWaitForWorkRequest(wId *string, entityType string, action oci_datascience.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_datascience.DataScienceClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "datascience")
	retryPolicy.ShouldRetryOperation = scheduleWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromDatascienceScheduleWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDatascienceScheduleWorkRequest(client *oci_datascience.DataScienceClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_datascience.WorkRequestResourceActionTypeEnum) error {
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

func (s *DatascienceScheduleResourceCrud) Get() error {
	request := oci_datascience.GetScheduleRequest{}

	tmp := s.D.Id()
	request.ScheduleId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.GetSchedule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Schedule
	return nil
}

func (s *DatascienceScheduleResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_datascience.UpdateScheduleRequest{}

	if action, ok := s.D.GetOkExists("action"); ok {
		if tmpList := action.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "action", 0)
			tmp, err := s.mapToScheduleAction(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Action = tmp
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

	if logDetails, ok := s.D.GetOkExists("log_details"); ok {
		if tmpList := logDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "log_details", 0)
			tmp, err := s.mapToScheduleLogDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LogDetails = &tmp
		}
	}

	tmp := s.D.Id()
	request.ScheduleId = &tmp

	if trigger, ok := s.D.GetOkExists("trigger"); ok {
		if tmpList := trigger.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "trigger", 0)
			tmp, err := s.mapToScheduleTrigger(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Trigger = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.UpdateSchedule(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getScheduleFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience"), oci_datascience.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatascienceScheduleResourceCrud) Delete() error {
	request := oci_datascience.DeleteScheduleRequest{}

	tmp := s.D.Id()
	request.ScheduleId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.DeleteSchedule(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := scheduleWaitForWorkRequest(workId, "schedule",
		oci_datascience.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DatascienceScheduleResourceCrud) SetData() error {
	if s.Res.Action != nil {
		actionArray := []interface{}{}
		if actionMap := ScheduleActionToMap(&s.Res.Action); actionMap != nil {
			actionArray = append(actionArray, actionMap)
		}
		s.D.Set("action", actionArray)
	} else {
		s.D.Set("action", nil)
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

	if s.Res.LastScheduleRunDetails != nil {
		s.D.Set("last_schedule_run_details", *s.Res.LastScheduleRunDetails)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.LogDetails != nil {
		s.D.Set("log_details", []interface{}{ScheduleLogDetailsToMap(s.Res.LogDetails)})
	} else {
		s.D.Set("log_details", nil)
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

	if s.Res.TimeLastScheduleRun != nil {
		s.D.Set("time_last_schedule_run", s.Res.TimeLastScheduleRun.String())
	}

	if s.Res.TimeNextScheduledRun != nil {
		s.D.Set("time_next_scheduled_run", s.Res.TimeNextScheduledRun.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.Trigger != nil {
		triggerArray := []interface{}{}
		if triggerMap := ScheduleTriggerToMap(&s.Res.Trigger); triggerMap != nil {
			triggerArray = append(triggerArray, triggerMap)
		}
		s.D.Set("trigger", triggerArray)
	} else {
		s.D.Set("trigger", nil)
	}

	return nil
}

func (s *DatascienceScheduleResourceCrud) mapToCreateJobRunDetails(fieldKeyFormat string) (oci_datascience.CreateJobRunDetails, error) {
	result := oci_datascience.CreateJobRunDetails{}

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		result.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		tmp, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if jobConfigurationOverrideDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "job_configuration_override_details")); ok {
		if tmpList := jobConfigurationOverrideDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "job_configuration_override_details"), 0)
			tmp, err := s.mapToJobConfigurationDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert job_configuration_override_details, encountered error: %v", err)
			}
			result.JobConfigurationOverrideDetails = tmp
		}
	}

	if jobEnvironmentConfigurationOverrideDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "job_environment_configuration_override_details")); ok {
		if tmpList := jobEnvironmentConfigurationOverrideDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "job_environment_configuration_override_details"), 0)
			tmp, err := s.mapToJobEnvironmentConfigurationDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert job_environment_configuration_override_details, encountered error: %v", err)
			}
			result.JobEnvironmentConfigurationOverrideDetails = tmp
		}
	}

	if jobId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "job_id")); ok {
		tmp := jobId.(string)
		result.JobId = &tmp
	}

	if jobLogConfigurationOverrideDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "job_log_configuration_override_details")); ok {
		if tmpList := jobLogConfigurationOverrideDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "job_log_configuration_override_details"), 0)
			tmp, err := s.mapToJobLogConfigurationDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert job_log_configuration_override_details, encountered error: %v", err)
			}
			result.JobLogConfigurationOverrideDetails = &tmp
		}
	}

	if projectId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "project_id")); ok {
		tmp := projectId.(string)
		result.ProjectId = &tmp
	}

	return result, nil
}

func CreateJobRunDetailsToMap(obj *oci_datascience.CreateJobRunDetails) map[string]interface{} {
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

	if obj.JobConfigurationOverrideDetails != nil {
		jobConfigurationOverrideDetailsArray := []interface{}{}
		if jobConfigurationOverrideDetailsMap := JobConfigurationDetailsToMap(&obj.JobConfigurationOverrideDetails); jobConfigurationOverrideDetailsMap != nil {
			jobConfigurationOverrideDetailsArray = append(jobConfigurationOverrideDetailsArray, jobConfigurationOverrideDetailsMap)
		}
		result["job_configuration_override_details"] = jobConfigurationOverrideDetailsArray
	}

	if obj.JobEnvironmentConfigurationOverrideDetails != nil {
		jobEnvironmentConfigurationOverrideDetailsArray := []interface{}{}
		if jobEnvironmentConfigurationOverrideDetailsMap := JobEnvironmentConfigurationDetailsToMap(&obj.JobEnvironmentConfigurationOverrideDetails); jobEnvironmentConfigurationOverrideDetailsMap != nil {
			jobEnvironmentConfigurationOverrideDetailsArray = append(jobEnvironmentConfigurationOverrideDetailsArray, jobEnvironmentConfigurationOverrideDetailsMap)
		}
		result["job_environment_configuration_override_details"] = jobEnvironmentConfigurationOverrideDetailsArray
	}

	if obj.JobId != nil {
		result["job_id"] = string(*obj.JobId)
	}

	if obj.JobLogConfigurationOverrideDetails != nil {
		result["job_log_configuration_override_details"] = []interface{}{JobLogConfigurationDetailsToMap(obj.JobLogConfigurationOverrideDetails)}
	}

	if obj.ProjectId != nil {
		result["project_id"] = string(*obj.ProjectId)
	}

	return result
}

func (s *DatascienceScheduleResourceCrud) mapToCreatePipelineRunDetails(fieldKeyFormat string) (oci_datascience.CreatePipelineRunDetails, error) {
	result := oci_datascience.CreatePipelineRunDetails{}

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		result.CompartmentId = &tmp
	}

	if configurationOverrideDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "configuration_override_details")); ok {
		if tmpList := configurationOverrideDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "configuration_override_details"), 0)
			tmp, err := s.mapToPipelineConfigurationDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert configuration_override_details, encountered error: %v", err)
			}
			result.ConfigurationOverrideDetails = tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		tmp, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert defined_tags, encountered error: %v", err)
		}
		result.DefinedTags = tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		result.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if logConfigurationOverrideDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "log_configuration_override_details")); ok {
		if tmpList := logConfigurationOverrideDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "log_configuration_override_details"), 0)
			tmp, err := s.mapToPipelineLogConfigurationDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert log_configuration_override_details, encountered error: %v", err)
			}
			result.LogConfigurationOverrideDetails = &tmp
		}
	}

	if pipelineId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pipeline_id")); ok {
		tmp := pipelineId.(string)
		result.PipelineId = &tmp
	}

	if projectId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "project_id")); ok {
		tmp := projectId.(string)
		result.ProjectId = &tmp
	}

	if stepOverrideDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_override_details")); ok {
		interfaces := stepOverrideDetails.([]interface{})
		tmp := make([]oci_datascience.PipelineStepOverrideDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "step_override_details"), stateDataIndex)
			converted, err := s.mapToPipelineStepOverrideDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "step_override_details")) {
			result.StepOverrideDetails = tmp
		}
	}

	if systemTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "system_tags")); ok {
		tmp, err := tfresource.MapToSystemTags(systemTags.(map[string]interface{}))
		if err != nil {
			return result, fmt.Errorf("unable to convert system_tags, encountered error: %v", err)
		}
		result.SystemTags = tmp
	}

	return result, nil
}

func CreatePipelineRunDetailsToMap(obj *oci_datascience.CreatePipelineRunDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ConfigurationOverrideDetails != nil {
		configurationOverrideDetailsArray := []interface{}{}
		if configurationOverrideDetailsMap := PipelineConfigurationDetailsToMap(&obj.ConfigurationOverrideDetails); configurationOverrideDetailsMap != nil {
			configurationOverrideDetailsArray = append(configurationOverrideDetailsArray, configurationOverrideDetailsMap)
		}
		result["configuration_override_details"] = configurationOverrideDetailsArray
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.LogConfigurationOverrideDetails != nil {
		result["log_configuration_override_details"] = []interface{}{PipelineLogConfigurationDetailsToMap(obj.LogConfigurationOverrideDetails)}
	}

	if obj.PipelineId != nil {
		result["pipeline_id"] = string(*obj.PipelineId)
	}

	if obj.ProjectId != nil {
		result["project_id"] = string(*obj.ProjectId)
	}

	stepOverrideDetails := []interface{}{}
	for _, item := range obj.StepOverrideDetails {
		stepOverrideDetails = append(stepOverrideDetails, PipelineStepOverrideDetailsToMap(item))
	}
	result["step_override_details"] = stepOverrideDetails

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	return result
}

func (s *DatascienceScheduleResourceCrud) mapToJobConfigurationDetails(fieldKeyFormat string) (oci_datascience.JobConfigurationDetails, error) {
	var baseObject oci_datascience.JobConfigurationDetails
	//discriminator
	jobTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "job_type"))
	var jobType string
	if ok {
		jobType = jobTypeRaw.(string)
	} else {
		jobType = "" // default value
	}
	switch strings.ToLower(jobType) {
	case strings.ToLower("DEFAULT"):
		details := oci_datascience.DefaultJobConfigurationDetails{}
		if commandLineArguments, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "command_line_arguments")); ok {
			tmp := commandLineArguments.(string)
			details.CommandLineArguments = &tmp
		}
		if environmentVariables, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "environment_variables")); ok {
			details.EnvironmentVariables = tfresource.ObjectMapToStringMap(environmentVariables.(map[string]interface{}))
		}
		if maximumRuntimeInMinutes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "maximum_runtime_in_minutes")); ok {
			tmp := maximumRuntimeInMinutes.(string)
			tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
			if err != nil {
				return details, fmt.Errorf("unable to convert maximumRuntimeInMinutes string: %s to an int64 and encountered error: %v", tmp, err)
			}
			details.MaximumRuntimeInMinutes = &tmpInt64
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown job_type '%v' was specified", jobType)
	}
	return baseObject, nil
}

func (s *DatascienceScheduleResourceCrud) mapToJobEnvironmentConfigurationDetails(fieldKeyFormat string) (oci_datascience.JobEnvironmentConfigurationDetails, error) {
	var baseObject oci_datascience.JobEnvironmentConfigurationDetails
	//discriminator
	jobEnvironmentTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "job_environment_type"))
	var jobEnvironmentType string
	if ok {
		jobEnvironmentType = jobEnvironmentTypeRaw.(string)
	} else {
		jobEnvironmentType = "" // default value
	}
	switch strings.ToLower(jobEnvironmentType) {
	case strings.ToLower("OCIR_CONTAINER"):
		details := oci_datascience.OcirContainerJobEnvironmentConfigurationDetails{}
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
		if image, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "image")); ok {
			tmp := image.(string)
			details.Image = &tmp
		}
		if imageDigest, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "image_digest")); ok {
			tmp := imageDigest.(string)
			details.ImageDigest = &tmp
		}
		if imageSignatureId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "image_signature_id")); ok {
			tmp := imageSignatureId.(string)
			details.ImageSignatureId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown job_environment_type '%v' was specified", jobEnvironmentType)
	}
	return baseObject, nil
}

func (s *DatascienceScheduleResourceCrud) mapToJobLogConfigurationDetails(fieldKeyFormat string) (oci_datascience.JobLogConfigurationDetails, error) {
	result := oci_datascience.JobLogConfigurationDetails{}

	if enableAutoLogCreation, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "enable_auto_log_creation")); ok {
		tmp := enableAutoLogCreation.(bool)
		result.EnableAutoLogCreation = &tmp
	}

	if enableLogging, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "enable_logging")); ok {
		tmp := enableLogging.(bool)
		result.EnableLogging = &tmp
	}

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

func (s *DatascienceScheduleResourceCrud) mapToPipelineConfigurationDetails(fieldKeyFormat string) (oci_datascience.PipelineConfigurationDetails, error) {
	var baseObject oci_datascience.PipelineConfigurationDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("DEFAULT"):
		details := oci_datascience.PipelineDefaultConfigurationDetails{}
		if commandLineArguments, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "command_line_arguments")); ok {
			tmp := commandLineArguments.(string)
			details.CommandLineArguments = &tmp
		}
		if environmentVariables, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "environment_variables")); ok {
			details.EnvironmentVariables = tfresource.ObjectMapToStringMap(environmentVariables.(map[string]interface{}))
		}
		if maximumRuntimeInMinutes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "maximum_runtime_in_minutes")); ok {
			tmp := maximumRuntimeInMinutes.(string)
			tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
			if err != nil {
				return details, fmt.Errorf("unable to convert maximumRuntimeInMinutes string: %s to an int64 and encountered error: %v", tmp, err)
			}
			details.MaximumRuntimeInMinutes = &tmpInt64
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func (s *DatascienceScheduleResourceCrud) mapToPipelineContainerConfigurationDetails(fieldKeyFormat string) (oci_datascience.PipelineContainerConfigurationDetails, error) {
	var baseObject oci_datascience.PipelineContainerConfigurationDetails
	//discriminator
	containerTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "container_type"))
	var containerType string
	if ok {
		containerType = containerTypeRaw.(string)
	} else {
		containerType = "" // default value
	}
	switch strings.ToLower(containerType) {
	case strings.ToLower("OCIR_CONTAINER"):
		details := oci_datascience.PipelineOcirContainerConfigurationDetails{}
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
		if image, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "image")); ok {
			tmp := image.(string)
			details.Image = &tmp
		}
		if imageDigest, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "image_digest")); ok {
			tmp := imageDigest.(string)
			details.ImageDigest = &tmp
		}
		if imageSignatureId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "image_signature_id")); ok {
			tmp := imageSignatureId.(string)
			details.ImageSignatureId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown container_type '%v' was specified", containerType)
	}
	return baseObject, nil
}

func (s *DatascienceScheduleResourceCrud) mapToPipelineLogConfigurationDetails(fieldKeyFormat string) (oci_datascience.PipelineLogConfigurationDetails, error) {
	result := oci_datascience.PipelineLogConfigurationDetails{}

	if enableAutoLogCreation, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "enable_auto_log_creation")); ok {
		tmp := enableAutoLogCreation.(bool)
		result.EnableAutoLogCreation = &tmp
	}

	if enableLogging, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "enable_logging")); ok {
		tmp := enableLogging.(bool)
		result.EnableLogging = &tmp
	}

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

func (s *DatascienceScheduleResourceCrud) mapToPipelineStepConfigurationDetails(fieldKeyFormat string) (oci_datascience.PipelineStepConfigurationDetails, error) {
	result := oci_datascience.PipelineStepConfigurationDetails{}

	if commandLineArguments, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "command_line_arguments")); ok {
		tmp := commandLineArguments.(string)
		result.CommandLineArguments = &tmp
	}

	if environmentVariables, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "environment_variables")); ok {
		result.EnvironmentVariables = tfresource.ObjectMapToStringMap(environmentVariables.(map[string]interface{}))
	}

	if maximumRuntimeInMinutes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "maximum_runtime_in_minutes")); ok {
		tmp := maximumRuntimeInMinutes.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert maximumRuntimeInMinutes string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.MaximumRuntimeInMinutes = &tmpInt64
	}

	return result, nil
}

func (s *DatascienceScheduleResourceCrud) mapToPipelineStepOverrideDetails(fieldKeyFormat string) (oci_datascience.PipelineStepOverrideDetails, error) {
	result := oci_datascience.PipelineStepOverrideDetails{}

	if stepConfigurationDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_configuration_details")); ok {
		if tmpList := stepConfigurationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "step_configuration_details"), 0)
			tmp, err := s.mapToPipelineStepConfigurationDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert step_configuration_details, encountered error: %v", err)
			}
			result.StepConfigurationDetails = &tmp
		}
	}

	if stepContainerConfigurationDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_container_configuration_details")); ok {
		if tmpList := stepContainerConfigurationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "step_container_configuration_details"), 0)
			tmp, err := s.mapToPipelineContainerConfigurationDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert step_container_configuration_details, encountered error: %v", err)
			}
			result.StepContainerConfigurationDetails = tmp
		}
	}

	if stepName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_name")); ok {
		tmp := stepName.(string)
		result.StepName = &tmp
	}

	return result, nil
}

func (s *DatascienceScheduleResourceCrud) mapToScheduleAction(fieldKeyFormat string) (oci_datascience.ScheduleAction, error) {
	var baseObject oci_datascience.ScheduleAction
	//discriminator
	actionTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action_type"))
	var actionType string
	if ok {
		actionType = actionTypeRaw.(string)
	} else {
		actionType = "" // default value
	}
	switch strings.ToLower(actionType) {
	case strings.ToLower("HTTP"):
		details := oci_datascience.ScheduleHttpAction{}
		if actionDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action_details")); ok {
			if tmpList := actionDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "action_details"), 0)
				tmp, err := s.mapToScheduleHttpActionDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert action_details, encountered error: %v", err)
				}
				details.ActionDetails = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown action_type '%v' was specified", actionType)
	}
	return baseObject, nil
}

func ScheduleActionToMap(obj *oci_datascience.ScheduleAction) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_datascience.ScheduleHttpAction:
		result["action_type"] = "HTTP"

		if v.ActionDetails != nil {
			actionDetailsArray := []interface{}{}
			if actionDetailsMap := ScheduleHttpActionDetailsToMap(&v.ActionDetails); actionDetailsMap != nil {
				actionDetailsArray = append(actionDetailsArray, actionDetailsMap)
			}
			result["action_details"] = actionDetailsArray
		}
	default:
		log.Printf("[WARN] Received 'action_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatascienceScheduleResourceCrud) mapToScheduleHttpActionDetails(fieldKeyFormat string) (oci_datascience.ScheduleHttpActionDetails, error) {
	var baseObject oci_datascience.ScheduleHttpActionDetails
	//discriminator
	httpActionTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "http_action_type"))
	var httpActionType string
	if ok {
		httpActionType = httpActionTypeRaw.(string)
	} else {
		httpActionType = "" // default value
	}
	switch strings.ToLower(httpActionType) {
	case strings.ToLower("CREATE_JOB_RUN"):
		details := oci_datascience.CreateJobRunScheduleActionDetails{}
		if createJobRunDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "create_job_run_details")); ok {
			if tmpList := createJobRunDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "create_job_run_details"), 0)
				tmp, err := s.mapToCreateJobRunDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert create_job_run_details, encountered error: %v", err)
				}
				details.CreateJobRunDetails = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("CREATE_PIPELINE_RUN"):
		details := oci_datascience.CreatePipelineRunScheduleActionDetails{}
		if createPipelineRunDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "create_pipeline_run_details")); ok {
			if tmpList := createPipelineRunDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "create_pipeline_run_details"), 0)
				tmp, err := s.mapToCreatePipelineRunDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert create_pipeline_run_details, encountered error: %v", err)
				}
				details.CreatePipelineRunDetails = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("INVOKE_ML_APPLICATION_PROVIDER_TRIGGER"):
		details := oci_datascience.InvokeMlApplicationProviderTriggerScheduleActionDetails{}
		if mlApplicationInstanceViewId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ml_application_instance_view_id")); ok {
			tmp := mlApplicationInstanceViewId.(string)
			details.MlApplicationInstanceViewId = &tmp
		}
		if triggerMlApplicationInstanceViewFlowDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "trigger_ml_application_instance_view_flow_details")); ok {
			if tmpList := triggerMlApplicationInstanceViewFlowDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "trigger_ml_application_instance_view_flow_details"), 0)
				tmp, err := s.mapToTriggerMlApplicationInstanceViewFlowDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert trigger_ml_application_instance_view_flow_details, encountered error: %v", err)
				}
				details.TriggerMlApplicationInstanceViewFlowDetails = &tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown http_action_type '%v' was specified", httpActionType)
	}
	return baseObject, nil
}

func ScheduleHttpActionDetailsToMap(obj *oci_datascience.ScheduleHttpActionDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_datascience.CreateJobRunScheduleActionDetails:
		result["http_action_type"] = "CREATE_JOB_RUN"

		if v.CreateJobRunDetails != nil {
			result["create_job_run_details"] = []interface{}{CreateJobRunDetailsToMap(v.CreateJobRunDetails)}
		}
	case oci_datascience.CreatePipelineRunScheduleActionDetails:
		result["http_action_type"] = "CREATE_PIPELINE_RUN"

		if v.CreatePipelineRunDetails != nil {
			result["create_pipeline_run_details"] = []interface{}{CreatePipelineRunDetailsToMap(v.CreatePipelineRunDetails)}
		}
	case oci_datascience.InvokeMlApplicationProviderTriggerScheduleActionDetails:
		result["http_action_type"] = "INVOKE_ML_APPLICATION_PROVIDER_TRIGGER"

		if v.MlApplicationInstanceViewId != nil {
			result["ml_application_instance_view_id"] = string(*v.MlApplicationInstanceViewId)
		}

		if v.TriggerMlApplicationInstanceViewFlowDetails != nil {
			result["trigger_ml_application_instance_view_flow_details"] = []interface{}{TriggerMlApplicationInstanceViewFlowDetailsToMap(v.TriggerMlApplicationInstanceViewFlowDetails)}
		}
	default:
		log.Printf("[WARN] Received 'http_action_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatascienceScheduleResourceCrud) mapToScheduleLogDetails(fieldKeyFormat string) (oci_datascience.ScheduleLogDetails, error) {
	result := oci_datascience.ScheduleLogDetails{}

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

func ScheduleLogDetailsToMap(obj *oci_datascience.ScheduleLogDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LogGroupId != nil {
		result["log_group_id"] = string(*obj.LogGroupId)
	}

	if obj.LogId != nil {
		result["log_id"] = string(*obj.LogId)
	}

	return result
}

func (s *DatascienceScheduleResourceCrud) mapToScheduleTrigger(fieldKeyFormat string) (oci_datascience.ScheduleTrigger, error) {
	var baseObject oci_datascience.ScheduleTrigger
	//discriminator
	triggerTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "trigger_type"))
	var triggerType string
	if ok {
		triggerType = triggerTypeRaw.(string)
	} else {
		triggerType = "" // default value
	}

	switch strings.ToLower(triggerType) {

	case strings.ToLower("CRON"):
		details := oci_datascience.ScheduleCronTrigger{}
		if cronExpression, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cron_expression")); ok {
			tmp := cronExpression.(string)
			details.CronExpression = &tmp
		}
		if timeEnd, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_end")); ok {
			if len(timeEnd.(string)) == 0 {
				details.TimeEnd = nil
			} else {
				tmp, err := time.Parse(time.RFC3339, timeEnd.(string))
				if err != nil {
					return details, err
				}
				details.TimeEnd = &oci_common.SDKTime{Time: tmp}
			}
		}
		if timeStart, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_start")); ok {
			if len(timeStart.(string)) == 0 {
				details.TimeStart = nil
			} else {
				tmp, err := time.Parse(time.RFC3339, timeStart.(string))
				if err != nil {
					return details, err
				}
				details.TimeStart = &oci_common.SDKTime{Time: tmp}
			}
		}
		baseObject = details

	case strings.ToLower("ICAL"):
		details := oci_datascience.ScheduleICalTrigger{}
		if recurrence, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recurrence")); ok {
			tmp := recurrence.(string)
			details.Recurrence = &tmp
		}
		if timeEnd, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_end")); ok {
			if len(timeEnd.(string)) == 0 {
				details.TimeEnd = nil
			} else {
				tmp, err := time.Parse(time.RFC3339, timeEnd.(string))
				if err != nil {
					return details, err
				}
				details.TimeEnd = &oci_common.SDKTime{Time: tmp}
			}
		}
		if timeStart, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_start")); ok {
			if len(timeStart.(string)) == 0 {
				details.TimeStart = nil
			} else {
				tmp, err := time.Parse(time.RFC3339, timeStart.(string))
				if err != nil {
					return details, err
				}
				details.TimeStart = &oci_common.SDKTime{Time: tmp}
			}
		}
		baseObject = details

	case strings.ToLower("INTERVAL"):
		log.Printf("[DEBUG] INTERVAL TRIGGER STATE: %s", s.D.State())

		//if startTime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_start")); ok && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "time_start")) {
		//	startTimeOld, startTimeNew := s.D.GetChange(fmt.Sprintf(fieldKeyFormat, "time_start"))
		//	log.Printf("[DEBUG] START_TIME %s,  START_TIME_OLD %s,  START_TIME_NEW %s", startTime.(string), startTimeOld.(string), startTimeNew.(string))
		//}
		//
		//if endTime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_end")); ok && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "time_end")) {
		//	endTimeOld, endTimeNew := s.D.GetChange(fmt.Sprintf(fieldKeyFormat, "time_end"))
		//	log.Printf("[DEBUG] END_TIME %s,  END_TIME_OLD %s,  END_TIME_NEW %s", endTime.(string), endTimeOld.(string), endTimeNew.(string))
		//}
		//
		//if isRandomStartTime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_random_start_time")); ok && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "is_random_start_time")) {
		//	isRandomStartTimeOld, isRandomStartTimeNew := s.D.GetChange(fmt.Sprintf(fieldKeyFormat, "is_random_start_time"))
		//	log.Printf("[DEBUG] IS_RANDOM_START_TIME %s, IS_RANDOM_START_TIME_OLD %s,  IS_RANDOM_START_TIME_NEW %s", isRandomStartTime.(bool), isRandomStartTimeOld.(bool), isRandomStartTimeNew.(bool))
		//}
		//
		//if interval, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "interval")); ok && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "interval")) {
		//	intervalOld, intervalNew := s.D.GetChange(fmt.Sprintf(fieldKeyFormat, "interval"))
		//	log.Printf("[DEBUG] INTERVAL %s, INTERVAL_OLD %s,  INTERVAL_NEW %s", interval, intervalOld.(int), intervalNew.(int))
		//}
		//
		//if frequency, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "frequency")); ok && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "frequency")) {
		//	frequencyOld, frequencyNew := s.D.GetChange(fmt.Sprintf(fieldKeyFormat, "frequency"))
		//	log.Printf("[DEBUG] FREQUENCY %s, FREQUENCY_OLD %s,  FREQUENCY_NEW %s", frequency, frequencyOld.(string), frequencyNew.(string))
		//}

		details := oci_datascience.ScheduleIntervalTrigger{}

		if frequency, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "frequency")); ok {
			details.Frequency = oci_datascience.ScheduleIntervalTriggerFrequencyEnum(frequency.(string))
		}

		if interval, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "interval")); ok {
			tmp := interval.(int)
			details.Interval = &tmp
		}

		if isRandomStartTime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_random_start_time")); ok {
			tmp := isRandomStartTime.(bool)
			details.IsRandomStartTime = &tmp
		}

		if s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "time_start")) {
			if timeStart, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_start")); ok {
				if len(timeStart.(string)) == 0 {
					details.TimeStart = nil
				} else {
					tmp, err := time.Parse(time.RFC3339, timeStart.(string))
					if err != nil {
						return details, err
					}
					details.TimeStart = &oci_common.SDKTime{Time: tmp}
				}
			}
		}

		if timeEnd, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_end")); ok {
			if len(timeEnd.(string)) == 0 {
				details.TimeEnd = nil
			} else {
				tmp, err := time.Parse(time.RFC3339, timeEnd.(string))
				if err != nil {
					return details, err
				}
				details.TimeEnd = &oci_common.SDKTime{Time: tmp}
			}
		}

		baseObject = details
	default:
		return nil, fmt.Errorf("unknown trigger_type '%v' was specified", triggerType)
	}
	return baseObject, nil
}

func ScheduleTriggerToMap(obj *oci_datascience.ScheduleTrigger) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_datascience.ScheduleCronTrigger:
		result["trigger_type"] = "CRON"

		if v.CronExpression != nil {
			result["cron_expression"] = string(*v.CronExpression)
		}

		if v.TimeEnd != nil {
			result["time_end"] = v.TimeEnd.Format(time.RFC3339Nano)
		}

		if v.TimeStart != nil {
			result["time_start"] = v.TimeStart.Format(time.RFC3339Nano)
		}
	case oci_datascience.ScheduleICalTrigger:
		result["trigger_type"] = "ICAL"

		if v.Recurrence != nil {
			result["recurrence"] = string(*v.Recurrence)
		}

		if v.TimeEnd != nil {
			result["time_end"] = v.TimeEnd.Format(time.RFC3339Nano)
		}

		if v.TimeStart != nil {
			result["time_start"] = v.TimeStart.Format(time.RFC3339Nano)
		}
	case oci_datascience.ScheduleIntervalTrigger:
		result["trigger_type"] = "INTERVAL"

		result["frequency"] = string(v.Frequency)

		if v.Interval != nil {
			result["interval"] = int(*v.Interval)
		}

		if v.IsRandomStartTime != nil {
			result["is_random_start_time"] = bool(*v.IsRandomStartTime)
		}

		if v.TimeEnd != nil {
			result["time_end"] = v.TimeEnd.Format(time.RFC3339Nano)
		}

		if v.TimeStart != nil {
			result["time_start"] = v.TimeStart.Format(time.RFC3339Nano)
		}
	default:
		log.Printf("[WARN] Received 'trigger_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatascienceScheduleResourceCrud) mapToTriggerMlApplicationInstanceViewFlowDetails(fieldKeyFormat string) (oci_datascience.TriggerMlApplicationInstanceViewFlowDetails, error) {
	result := oci_datascience.TriggerMlApplicationInstanceViewFlowDetails{}

	if parameters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parameters")); ok {
		interfaces := parameters.([]interface{})
		tmp := make([]oci_datascience.TriggerParameter, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parameters"), stateDataIndex)
			converted, err := s.mapToTriggerParameter(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "parameters")) {
			result.Parameters = tmp
		}
	}

	if triggerName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "trigger_name")); ok {
		tmp := triggerName.(string)
		result.TriggerName = &tmp
	}

	return result, nil
}

func TriggerMlApplicationInstanceViewFlowDetailsToMap(obj *oci_datascience.TriggerMlApplicationInstanceViewFlowDetails) map[string]interface{} {
	result := map[string]interface{}{}

	parameters := []interface{}{}
	for _, item := range obj.Parameters {
		parameters = append(parameters, TriggerParameterToMap(item))
	}
	result["parameters"] = parameters

	if obj.TriggerName != nil {
		result["trigger_name"] = string(*obj.TriggerName)
	}

	return result
}

func (s *DatascienceScheduleResourceCrud) mapToTriggerParameter(fieldKeyFormat string) (oci_datascience.TriggerParameter, error) {
	result := oci_datascience.TriggerParameter{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func TriggerParameterToMap(obj oci_datascience.TriggerParameter) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *DatascienceScheduleResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_datascience.ChangeScheduleCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ScheduleId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.ChangeScheduleCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getScheduleFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience"), oci_datascience.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
