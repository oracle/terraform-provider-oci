// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_apps_management

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
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetAppsManagementRunbookResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFleetAppsManagementRunbook,
		Read:     readFleetAppsManagementRunbook,
		Update:   updateFleetAppsManagementRunbook,
		Delete:   deleteFleetAppsManagementRunbook,
		Schema: map[string]*schema.Schema{
			// Required
			"associations": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"execution_workflow_details": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"workflow": {
										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"group_name": {
													Type:     schema.TypeString,
													Required: true,
												},
												"steps": {
													Type:     schema.TypeList,
													Required: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"type": {
																Type:             schema.TypeString,
																Required:         true,
																DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																ValidateFunc: validation.StringInSlice([]string{
																	"PARALLEL_TASK_GROUP",
																	"TASK",
																}, true),
															},

															// Optional
															"group_name": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"step_name": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"steps": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															// Computed
														},
													},
												},
												"type": {
													Type:     schema.TypeString,
													Required: true,
												},

												// Optional

												// Computed
											},
										},
									},

									// Optional

									// Computed
								},
							},
						},
						"groups": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"type": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"properties": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"action_on_failure": {
													Type:     schema.TypeString,
													Required: true,
												},

												// Optional
												"condition": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"notification_preferences": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"should_notify_on_pause": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
															},
															"should_notify_on_task_failure": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
															},
															"should_notify_on_task_success": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
															},

															// Computed
														},
													},
												},
												"pause_details": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"kind": {
																Type:             schema.TypeString,
																Required:         true,
																DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																ValidateFunc: validation.StringInSlice([]string{
																	"TIME_BASED",
																	"USER_ACTION",
																}, true),
															},

															// Optional
															"duration_in_minutes": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},

															// Computed
														},
													},
												},
												"run_on": {
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
						"tasks": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"association_type": {
										Type:     schema.TypeString,
										Required: true,
									},
									"step_name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"task_record_details": {
										Type:     schema.TypeList,
										Required: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"scope": {
													Type:             schema.TypeString,
													Required:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"LOCAL",
														"SHARED",
													}, true),
												},

												// Optional
												"description": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"execution_details": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"execution_type": {
																Type:             schema.TypeString,
																Required:         true,
																DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																ValidateFunc: validation.StringInSlice([]string{
																	"API",
																	"SCRIPT",
																}, true),
															},

															// Optional
															"command": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"content": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																MaxItems: 1,
																MinItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required
																		"bucket": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																		"checksum": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																		"namespace": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																		"object": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																		"source_type": {
																			Type:             schema.TypeString,
																			Required:         true,
																			DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																			ValidateFunc: validation.StringInSlice([]string{
																				"OBJECT_STORAGE_BUCKET",
																			}, true),
																		},

																		// Optional

																		// Computed
																	},
																},
															},
															"credentials": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional
																		"display_name": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																		"id": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},

																		// Computed
																	},
																},
															},
															"endpoint": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"variables": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																MaxItems: 1,
																MinItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional
																		"input_variables": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					// Required

																					// Optional
																					"description": {
																						Type:     schema.TypeString,
																						Optional: true,
																						Computed: true,
																					},
																					"name": {
																						Type:     schema.TypeString,
																						Optional: true,
																						Computed: true,
																					},
																					"type": {
																						Type:     schema.TypeString,
																						Optional: true,
																						Computed: true,
																					},

																					// Computed
																				},
																			},
																		},
																		"output_variables": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Computed: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
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
												"is_apply_subject_task": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"is_copy_to_library_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"is_discovery_output_task": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"os_type": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"platform": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"properties": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"num_retries": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
															"timeout_in_seconds": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},

															// Computed
														},
													},
												},
												"task_record_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},

									// Optional
									"output_variable_mappings": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"name": {
													Type:     schema.TypeString,
													Required: true,
												},
												"output_variable_details": {
													Type:     schema.TypeList,
													Required: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"output_variable_name": {
																Type:     schema.TypeString,
																Required: true,
															},
															"step_name": {
																Type:     schema.TypeString,
																Required: true,
															},

															// Optional

															// Computed
														},
													},
												},

												// Optional

												// Computed
											},
										},
									},
									"step_properties": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"action_on_failure": {
													Type:     schema.TypeString,
													Required: true,
												},

												// Optional
												"condition": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"notification_preferences": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"should_notify_on_pause": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
															},
															"should_notify_on_task_failure": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
															},
															"should_notify_on_task_success": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
															},

															// Computed
														},
													},
												},
												"pause_details": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"kind": {
																Type:             schema.TypeString,
																Required:         true,
																DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																ValidateFunc: validation.StringInSlice([]string{
																	"TIME_BASED",
																	"USER_ACTION",
																}, true),
															},

															// Optional
															"duration_in_minutes": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},

															// Computed
														},
													},
												},
												"run_on": {
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

						// Optional
						"rollback_workflow_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"scope": {
										Type:     schema.TypeString,
										Required: true,
									},
									"workflow": {
										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"group_name": {
													Type:     schema.TypeString,
													Required: true,
												},
												"steps": {
													Type:     schema.TypeList,
													Required: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"type": {
																Type:             schema.TypeString,
																Required:         true,
																DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																ValidateFunc: validation.StringInSlice([]string{
																	"PARALLEL_TASK_GROUP",
																	"TASK",
																}, true),
															},

															// Optional
															"group_name": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"step_name": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"steps": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															// Computed
														},
													},
												},
												"type": {
													Type:     schema.TypeString,
													Required: true,
												},

												// Optional

												// Computed
											},
										},
									},

									// Optional

									// Computed
								},
							},
						},
						"version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"operation": {
				Type:     schema.TypeString,
				Required: true,
			},
			"os_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"runbook_relevance": {
				Type:     schema.TypeString,
				Required: true,
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
			"estimated_time": {
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
			"is_default": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"platform": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_region": {
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
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createFleetAppsManagementRunbook(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementRunbookResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementRunbooksClient()
	sync.FleetClient = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readFleetAppsManagementRunbook(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementRunbookResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementRunbooksClient()
	sync.FleetClient = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.ReadResource(sync)
}

func updateFleetAppsManagementRunbook(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementRunbookResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementRunbooksClient()
	sync.FleetClient = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteFleetAppsManagementRunbook(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementRunbookResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementRunbooksClient()
	sync.DisableNotFoundRetries = true
	sync.FleetClient = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.DeleteResource(d, sync)
}

type FleetAppsManagementRunbookResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_fleet_apps_management.FleetAppsManagementRunbooksClient
	FleetClient            *oci_fleet_apps_management.FleetAppsManagementClient
	Res                    *oci_fleet_apps_management.Runbook
	DisableNotFoundRetries bool
}

func (s *FleetAppsManagementRunbookResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *FleetAppsManagementRunbookResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_fleet_apps_management.RunbookLifecycleStateCreating),
	}
}

func (s *FleetAppsManagementRunbookResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.RunbookLifecycleStateInactive),
	}
}

func (s *FleetAppsManagementRunbookResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_fleet_apps_management.RunbookLifecycleStateDeleting),
	}
}

func (s *FleetAppsManagementRunbookResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.RunbookLifecycleStateDeleted),
	}
}

func (s *FleetAppsManagementRunbookResourceCrud) Create() error {
	request := oci_fleet_apps_management.CreateRunbookRequest{}

	if associations, ok := s.D.GetOkExists("associations"); ok {
		if tmpList := associations.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "associations", 0)
			tmp, err := s.mapToAssociations(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Associations = &tmp
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

	if estimatedTime, ok := s.D.GetOkExists("estimated_time"); ok {
		tmp := estimatedTime.(string)
		request.EstimatedTime = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isDefault, ok := s.D.GetOkExists("is_default"); ok {
		tmp := isDefault.(bool)
		request.IsDefault = &tmp
	}

	if operation, ok := s.D.GetOkExists("operation"); ok {
		tmp := operation.(string)
		request.Operation = &tmp
	}

	if osType, ok := s.D.GetOkExists("os_type"); ok {
		request.OsType = oci_fleet_apps_management.OsTypeEnum(osType.(string))
	}

	if platform, ok := s.D.GetOkExists("platform"); ok {
		tmp := platform.(string)
		request.Platform = &tmp
	}

	if runbookRelevance, ok := s.D.GetOkExists("runbook_relevance"); ok {
		request.RunbookRelevance = oci_fleet_apps_management.RunbookRunbookRelevanceEnum(runbookRelevance.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.CreateRunbook(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getRunbookFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management"), oci_fleet_apps_management.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *FleetAppsManagementRunbookResourceCrud) getRunbookFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_fleet_apps_management.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	runbookId, err := runbookWaitForWorkRequest(workId, "runbook",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.FleetClient)

	if err != nil {
		return err
	}
	s.D.SetId(*runbookId)

	return s.Get()
}

func runbookWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "fleet_apps_management", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_fleet_apps_management.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func runbookWaitForWorkRequest(wId *string, entityType string, action oci_fleet_apps_management.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_fleet_apps_management.FleetAppsManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "fleet_apps_management")
	retryPolicy.ShouldRetryOperation = runbookWorkRequestShouldRetryFunc(timeout)

	response := oci_fleet_apps_management.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_fleet_apps_management.OperationStatusInProgress),
			string(oci_fleet_apps_management.OperationStatusAccepted),
			string(oci_fleet_apps_management.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_fleet_apps_management.OperationStatusSucceeded),
			string(oci_fleet_apps_management.OperationStatusFailed),
			string(oci_fleet_apps_management.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			log.Printf("[DEBUG] Waiting on WorkRequestId %s", *wId)
			response, err = client.GetWorkRequest(context.Background(),
				oci_fleet_apps_management.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_fleet_apps_management.OperationStatusFailed || response.Status == oci_fleet_apps_management.OperationStatusCanceled {
		return nil, getErrorFromFleetAppsManagementRunbookWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromFleetAppsManagementRunbookWorkRequest(client *oci_fleet_apps_management.FleetAppsManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_fleet_apps_management.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_fleet_apps_management.ListWorkRequestErrorsRequest{
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

func (s *FleetAppsManagementRunbookResourceCrud) Get() error {
	request := oci_fleet_apps_management.GetRunbookRequest{}

	tmp := s.D.Id()
	request.RunbookId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.GetRunbook(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Runbook
	return nil
}

func (s *FleetAppsManagementRunbookResourceCrud) Update() error {
	request := oci_fleet_apps_management.UpdateRunbookRequest{}

	if associations, ok := s.D.GetOkExists("associations"); ok {
		if tmpList := associations.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "associations", 0)
			tmp, err := s.mapToAssociations(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Associations = &tmp
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

	if estimatedTime, ok := s.D.GetOkExists("estimated_time"); ok {
		tmp := estimatedTime.(string)
		request.EstimatedTime = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isDefault, ok := s.D.GetOkExists("is_default"); ok {
		tmp := isDefault.(bool)
		request.IsDefault = &tmp
	}

	if operation, ok := s.D.GetOkExists("operation"); ok {
		tmp := operation.(string)
		request.Operation = &tmp
	}

	if osType, ok := s.D.GetOkExists("os_type"); ok {
		request.OsType = oci_fleet_apps_management.OsTypeEnum(osType.(string))
	}

	if platform, ok := s.D.GetOkExists("platform"); ok {
		tmp := platform.(string)
		request.Platform = &tmp
	}

	tmp := s.D.Id()
	request.RunbookId = &tmp

	if runbookRelevance, ok := s.D.GetOkExists("runbook_relevance"); ok {
		request.RunbookRelevance = oci_fleet_apps_management.RunbookRunbookRelevanceEnum(runbookRelevance.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.UpdateRunbook(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getRunbookFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management"), oci_fleet_apps_management.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *FleetAppsManagementRunbookResourceCrud) Delete() error {
	request := oci_fleet_apps_management.DeleteRunbookRequest{}

	tmp := s.D.Id()
	request.RunbookId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.DeleteRunbook(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := runbookWaitForWorkRequest(workId, "runbook",
		oci_fleet_apps_management.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.FleetClient)
	return delWorkRequestErr
}

func (s *FleetAppsManagementRunbookResourceCrud) SetData() error {
	if s.Res.Associations != nil {
		s.D.Set("associations", []interface{}{AssociationsToMap(s.Res.Associations)})
	} else {
		s.D.Set("associations", nil)
	}

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

	if s.Res.EstimatedTime != nil {
		s.D.Set("estimated_time", *s.Res.EstimatedTime)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsDefault != nil {
		s.D.Set("is_default", *s.Res.IsDefault)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Operation != nil {
		s.D.Set("operation", *s.Res.Operation)
	}

	s.D.Set("os_type", s.Res.OsType)

	if s.Res.Platform != nil {
		s.D.Set("platform", *s.Res.Platform)
	}

	if s.Res.ResourceRegion != nil {
		s.D.Set("resource_region", *s.Res.ResourceRegion)
	}

	s.D.Set("runbook_relevance", s.Res.RunbookRelevance)

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

	s.D.Set("type", s.Res.Type)

	return nil
}

func (s *FleetAppsManagementRunbookResourceCrud) mapToAssociatedTaskDetails(fieldKeyFormat string) (oci_fleet_apps_management.AssociatedTaskDetails, error) {
	var baseObject oci_fleet_apps_management.AssociatedTaskDetails
	//discriminator
	scopeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scope"))
	var scope string
	if ok {
		scope = scopeRaw.(string)
	} else {
		scope = "" // default value
	}
	switch strings.ToLower(scope) {
	case strings.ToLower("LOCAL"):
		details := oci_fleet_apps_management.AssociatedLocalTaskDetails{}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if executionDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "execution_details")); ok {
			if tmpList := executionDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "execution_details"), 0)
				tmp, err := s.mapToExecutionDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert execution_details, encountered error: %v", err)
				}
				details.ExecutionDetails = tmp
			}
		}
		if isApplySubjectTask, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_apply_subject_task")); ok {
			tmp := isApplySubjectTask.(bool)
			details.IsApplySubjectTask = &tmp
		}
		if isCopyToLibraryEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_copy_to_library_enabled")); ok {
			tmp := isCopyToLibraryEnabled.(bool)
			details.IsCopyToLibraryEnabled = &tmp
		}
		if isDiscoveryOutputTask, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_discovery_output_task")); ok {
			tmp := isDiscoveryOutputTask.(bool)
			details.IsDiscoveryOutputTask = &tmp
		}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if osType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "os_type")); ok {
			details.OsType = oci_fleet_apps_management.OsTypeEnum(osType.(string))
		}
		if platform, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "platform")); ok {
			tmp := platform.(string)
			details.Platform = &tmp
		}
		if properties, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "properties")); ok {
			if tmpList := properties.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "properties"), 0)
				tmp, err := s.mapToProperties(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert properties, encountered error: %v", err)
				}
				details.Properties = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("SHARED"):
		details := oci_fleet_apps_management.AssociatedSharedTaskDetails{}
		if taskRecordId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "task_record_id")); ok {
			tmp := taskRecordId.(string)
			details.TaskRecordId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown scope '%v' was specified", scope)
	}
	return baseObject, nil
}

func AssociatedTaskDetailsToMap(obj *oci_fleet_apps_management.AssociatedTaskDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_fleet_apps_management.AssociatedLocalTaskDetails:
		result["scope"] = "LOCAL"

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}

		if v.ExecutionDetails != nil {
			executionDetailsArray := []interface{}{}
			if executionDetailsMap := ExecutionDetailsToMap(&v.ExecutionDetails); executionDetailsMap != nil {
				executionDetailsArray = append(executionDetailsArray, executionDetailsMap)
			}
			result["execution_details"] = executionDetailsArray
		}

		if v.IsApplySubjectTask != nil {
			result["is_apply_subject_task"] = bool(*v.IsApplySubjectTask)
		}

		if v.IsCopyToLibraryEnabled != nil {
			result["is_copy_to_library_enabled"] = bool(*v.IsCopyToLibraryEnabled)
		}

		if v.IsDiscoveryOutputTask != nil {
			result["is_discovery_output_task"] = bool(*v.IsDiscoveryOutputTask)
		}

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}

		result["os_type"] = string(v.OsType)

		if v.Platform != nil {
			result["platform"] = string(*v.Platform)
		}

		if v.Properties != nil {
			result["properties"] = []interface{}{PropertiesToMap(v.Properties)}
		}
	case oci_fleet_apps_management.AssociatedSharedTaskDetails:
		result["scope"] = "SHARED"

		if v.TaskRecordId != nil {
			result["task_record_id"] = string(*v.TaskRecordId)
		}
	default:
		log.Printf("[WARN] Received 'scope' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *FleetAppsManagementRunbookResourceCrud) mapToAssociations(fieldKeyFormat string) (oci_fleet_apps_management.Associations, error) {
	result := oci_fleet_apps_management.Associations{}

	if executionWorkflowDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "execution_workflow_details")); ok {
		if tmpList := executionWorkflowDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "execution_workflow_details"), 0)
			tmp, err := s.mapToExecutionWorkflowDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert execution_workflow_details, encountered error: %v", err)
			}
			result.ExecutionWorkflowDetails = &tmp
		}
	}

	if groups, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "groups")); ok {
		interfaces := groups.([]interface{})
		tmp := make([]oci_fleet_apps_management.Group, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "groups"), stateDataIndex)
			converted, err := s.mapToGroup(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "groups")) {
			result.Groups = tmp
		}
	}

	if rollbackWorkflowDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rollback_workflow_details")); ok {
		if tmpList := rollbackWorkflowDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "rollback_workflow_details"), 0)
			tmp, err := s.mapToRollbackWorkflowDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert rollback_workflow_details, encountered error: %v", err)
			}
			result.RollbackWorkflowDetails = &tmp
		}
	}

	if tasks, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tasks")); ok {
		interfaces := tasks.([]interface{})
		tmp := make([]oci_fleet_apps_management.Task, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "tasks"), stateDataIndex)
			converted, err := s.mapToTask(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "tasks")) {
			result.Tasks = tmp
		}
	}

	if version, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version")); ok {
		tmp := version.(string)
		result.Version = &tmp
	}

	return result, nil
}

func AssociationsToMap(obj *oci_fleet_apps_management.Associations) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ExecutionWorkflowDetails != nil {
		result["execution_workflow_details"] = []interface{}{ExecutionWorkflowDetailsToMap(obj.ExecutionWorkflowDetails)}
	}

	groups := []interface{}{}
	for _, item := range obj.Groups {
		groups = append(groups, GroupToMap(item))
	}
	result["groups"] = groups

	if obj.RollbackWorkflowDetails != nil {
		result["rollback_workflow_details"] = []interface{}{RollbackWorkflowDetailsToMap(obj.RollbackWorkflowDetails)}
	}

	tasks := []interface{}{}
	for _, item := range obj.Tasks {
		tasks = append(tasks, TaskToMap(item))
	}
	result["tasks"] = tasks

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}

func (s *FleetAppsManagementRunbookResourceCrud) mapToComponentProperties(fieldKeyFormat string) (oci_fleet_apps_management.ComponentProperties, error) {
	result := oci_fleet_apps_management.ComponentProperties{}

	if actionOnFailure, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action_on_failure")); ok {
		result.ActionOnFailure = oci_fleet_apps_management.ComponentPropertiesActionOnFailureEnum(actionOnFailure.(string))
	}

	if condition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "condition")); ok {
		tmp := condition.(string)
		result.Condition = &tmp
	}

	if notificationPreferences, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "notification_preferences")); ok {
		if tmpList := notificationPreferences.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "notification_preferences"), 0)
			tmp, err := s.mapToTaskNotificationPreferences(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert notification_preferences, encountered error: %v", err)
			}
			result.NotificationPreferences = &tmp
		}
	}

	if pauseDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pause_details")); ok {
		if tmpList := pauseDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "pause_details"), 0)
			tmp, err := s.mapToPauseDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert pause_details, encountered error: %v", err)
			}
			result.PauseDetails = tmp
		}
	}

	if runOn, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "run_on")); ok {
		tmp := runOn.(string)
		result.RunOn = &tmp
	}

	return result, nil
}

func ComponentPropertiesToMap(obj *oci_fleet_apps_management.ComponentProperties) map[string]interface{} {
	result := map[string]interface{}{}

	result["action_on_failure"] = string(obj.ActionOnFailure)

	if obj.Condition != nil {
		result["condition"] = string(*obj.Condition)
	}

	if obj.NotificationPreferences != nil {
		result["notification_preferences"] = []interface{}{TaskNotificationPreferencesToMap(obj.NotificationPreferences)}
	}

	if obj.PauseDetails != nil {
		pauseDetailsArray := []interface{}{}
		if pauseDetailsMap := PauseDetailsToMap(&obj.PauseDetails); pauseDetailsMap != nil {
			pauseDetailsArray = append(pauseDetailsArray, pauseDetailsMap)
		}
		result["pause_details"] = pauseDetailsArray
	}

	if obj.RunOn != nil {
		result["run_on"] = string(*obj.RunOn)
	}

	return result
}

func (s *FleetAppsManagementRunbookResourceCrud) mapToConfigAssociationDetails(fieldKeyFormat string) (oci_fleet_apps_management.ConfigAssociationDetails, error) {
	result := oci_fleet_apps_management.ConfigAssociationDetails{}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	return result, nil
}

func (s *FleetAppsManagementRunbookResourceCrud) mapToContentDetails(fieldKeyFormat string) (oci_fleet_apps_management.ContentDetails, error) {
	var baseObject oci_fleet_apps_management.ContentDetails
	//discriminator
	sourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_type"))
	var sourceType string
	if ok {
		sourceType = sourceTypeRaw.(string)
	} else {
		sourceType = "" // default value
	}
	switch strings.ToLower(sourceType) {
	case strings.ToLower("OBJECT_STORAGE_BUCKET"):
		details := oci_fleet_apps_management.ObjectStorageBucketContentDetails{}
		if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.BucketName = &tmp
		}
		if checksum, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "checksum")); ok {
			tmp := checksum.(string)
			details.Checksum = &tmp
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
		return nil, fmt.Errorf("unknown source_type '%v' was specified", sourceType)
	}
	return baseObject, nil
}

func (s *FleetAppsManagementRunbookResourceCrud) mapToExecutionDetails(fieldKeyFormat string) (oci_fleet_apps_management.ExecutionDetails, error) {
	var baseObject oci_fleet_apps_management.ExecutionDetails
	//discriminator
	executionTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "execution_type"))
	var executionType string
	if ok {
		executionType = executionTypeRaw.(string)
	} else {
		executionType = "" // default value
	}
	switch strings.ToLower(executionType) {
	case strings.ToLower("API"):
		details := oci_fleet_apps_management.ApiBasedExecutionDetails{}
		if endpoint, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "endpoint")); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		baseObject = details
	case strings.ToLower("SCRIPT"):
		details := oci_fleet_apps_management.ScriptBasedExecutionDetails{}
		if command, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "command")); ok {
			tmp := command.(string)
			details.Command = &tmp
		}
		if content, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "content")); ok {
			if tmpList := content.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "content"), 0)
				tmp, err := s.mapToContentDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert content, encountered error: %v", err)
				}
				details.Content = tmp
			}
		}
		if credentials, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credentials")); ok {
			interfaces := credentials.([]interface{})
			tmp := make([]oci_fleet_apps_management.ConfigAssociationDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "credentials"), stateDataIndex)
				converted, err := s.mapToConfigAssociationDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "credentials")) {
				details.Credentials = tmp
			}
		}
		if variables, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "variables")); ok {
			if tmpList := variables.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "variables"), 0)
				tmp, err := s.mapToTaskVariable(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert variables, encountered error: %v", err)
				}
				details.Variables = &tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown execution_type '%v' was specified", executionType)
	}
	return baseObject, nil
}

func (s *FleetAppsManagementRunbookResourceCrud) mapToExecutionWorkflowDetails(fieldKeyFormat string) (oci_fleet_apps_management.ExecutionWorkflowDetails, error) {
	result := oci_fleet_apps_management.ExecutionWorkflowDetails{}

	if workflow, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "workflow")); ok {
		interfaces := workflow.([]interface{})
		tmp := make([]oci_fleet_apps_management.WorkflowGroup, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "workflow"), stateDataIndex)
			converted, err := s.mapToWorkflowGroup(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "workflow")) {
			result.Workflow = tmp
		}
	}

	return result, nil
}

func ExecutionWorkflowDetailsToMap(obj *oci_fleet_apps_management.ExecutionWorkflowDetails) map[string]interface{} {
	result := map[string]interface{}{}

	workflow := []interface{}{}
	for _, item := range obj.Workflow {
		workflow = append(workflow, WorkflowGroupToMap(item))
	}
	result["workflow"] = workflow

	return result
}

func (s *FleetAppsManagementRunbookResourceCrud) mapToGroup(fieldKeyFormat string) (oci_fleet_apps_management.Group, error) {
	result := oci_fleet_apps_management.Group{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if properties, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "properties")); ok {
		if tmpList := properties.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "properties"), 0)
			tmp, err := s.mapToComponentProperties(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert properties, encountered error: %v", err)
			}
			result.Properties = &tmp
		}
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_fleet_apps_management.GroupTypeEnum(type_.(string))
	}

	return result, nil
}

func GroupToMap(obj oci_fleet_apps_management.Group) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Properties != nil {
		result["properties"] = []interface{}{ComponentPropertiesToMap(obj.Properties)}
	}

	result["type"] = string(obj.Type)

	return result
}

func (s *FleetAppsManagementRunbookResourceCrud) mapToInputArgument(fieldKeyFormat string) (oci_fleet_apps_management.InputArgument, error) {
	var baseObject oci_fleet_apps_management.InputArgument
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("OUTPUT_VARIABLE"):
		details := oci_fleet_apps_management.OutputVariableInputArgument{}
		baseObject = details
	case strings.ToLower("STRING"):
		details := oci_fleet_apps_management.StringInputArgument{}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func (s *FleetAppsManagementRunbookResourceCrud) mapToOutputVariableDetails(fieldKeyFormat string) (oci_fleet_apps_management.OutputVariableDetails, error) {
	result := oci_fleet_apps_management.OutputVariableDetails{}

	if outputVariableName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "output_variable_name")); ok {
		tmp := outputVariableName.(string)
		result.OutputVariableName = &tmp
	}

	if stepName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_name")); ok {
		tmp := stepName.(string)
		result.StepName = &tmp
	}

	return result, nil
}

func OutputVariableDetailsToMap(obj *oci_fleet_apps_management.OutputVariableDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.OutputVariableName != nil {
		result["output_variable_name"] = string(*obj.OutputVariableName)
	}

	if obj.StepName != nil {
		result["step_name"] = string(*obj.StepName)
	}

	return result
}

func (s *FleetAppsManagementRunbookResourceCrud) mapToOutputVariableMapping(fieldKeyFormat string) (oci_fleet_apps_management.OutputVariableMapping, error) {
	result := oci_fleet_apps_management.OutputVariableMapping{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if outputVariableDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "output_variable_details")); ok {
		if tmpList := outputVariableDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "output_variable_details"), 0)
			tmp, err := s.mapToOutputVariableDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert output_variable_details, encountered error: %v", err)
			}
			result.OutputVariableDetails = &tmp
		}
	}

	return result, nil
}

func OutputVariableMappingToMap(obj oci_fleet_apps_management.OutputVariableMapping) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.OutputVariableDetails != nil {
		result["output_variable_details"] = []interface{}{OutputVariableDetailsToMap(obj.OutputVariableDetails)}
	}

	return result
}

func (s *FleetAppsManagementRunbookResourceCrud) mapToPauseDetails(fieldKeyFormat string) (oci_fleet_apps_management.PauseDetails, error) {
	var baseObject oci_fleet_apps_management.PauseDetails
	//discriminator
	kindRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kind"))
	var kind string
	if ok {
		kind = kindRaw.(string)
	} else {
		kind = "" // default value
	}
	switch strings.ToLower(kind) {
	case strings.ToLower("TIME_BASED"):
		details := oci_fleet_apps_management.TimeBasedPauseDetails{}
		if durationInMinutes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "duration_in_minutes")); ok {
			tmp := durationInMinutes.(int)
			details.DurationInMinutes = &tmp
		}
		baseObject = details
	case strings.ToLower("USER_ACTION"):
		details := oci_fleet_apps_management.UserActionBasedPauseDetails{}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown kind '%v' was specified", kind)
	}
	return baseObject, nil
}

func PauseDetailsToMap(obj *oci_fleet_apps_management.PauseDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_fleet_apps_management.TimeBasedPauseDetails:
		result["kind"] = "TIME_BASED"

		if v.DurationInMinutes != nil {
			result["duration_in_minutes"] = int(*v.DurationInMinutes)
		}
	case oci_fleet_apps_management.UserActionBasedPauseDetails:
		result["kind"] = "USER_ACTION"
	default:
		log.Printf("[WARN] Received 'kind' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *FleetAppsManagementRunbookResourceCrud) mapToProperties(fieldKeyFormat string) (oci_fleet_apps_management.Properties, error) {
	result := oci_fleet_apps_management.Properties{}

	if numRetries, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "num_retries")); ok {
		tmp := numRetries.(int)
		result.NumRetries = &tmp
	}

	if timeoutInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_seconds")); ok {
		tmp := timeoutInSeconds.(int)
		result.TimeoutInSeconds = &tmp
	}

	return result, nil
}

func (s *FleetAppsManagementRunbookResourceCrud) mapToRollbackWorkflowDetails(fieldKeyFormat string) (oci_fleet_apps_management.RollbackWorkflowDetails, error) {
	result := oci_fleet_apps_management.RollbackWorkflowDetails{}

	if scope, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scope")); ok {
		result.Scope = oci_fleet_apps_management.RollbackWorkflowDetailsScopeEnum(scope.(string))
	}

	if workflow, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "workflow")); ok {
		interfaces := workflow.([]interface{})
		tmp := make([]oci_fleet_apps_management.WorkflowGroup, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "workflow"), stateDataIndex)
			converted, err := s.mapToWorkflowGroup(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "workflow")) {
			result.Workflow = tmp
		}
	}

	return result, nil
}

func RollbackWorkflowDetailsToMap(obj *oci_fleet_apps_management.RollbackWorkflowDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["scope"] = string(obj.Scope)

	workflow := []interface{}{}
	for _, item := range obj.Workflow {
		workflow = append(workflow, WorkflowGroupToMap(item))
	}
	result["workflow"] = workflow

	return result
}

func RunbookSummaryToMap(obj oci_fleet_apps_management.RunbookSummary) map[string]interface{} {
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

	if obj.EstimatedTime != nil {
		result["estimated_time"] = string(*obj.EstimatedTime)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsDefault != nil {
		result["is_default"] = bool(*obj.IsDefault)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.Operation != nil {
		result["operation"] = string(*obj.Operation)
	}

	result["os_type"] = string(obj.OsType)

	if obj.Platform != nil {
		result["platform"] = string(*obj.Platform)
	}

	if obj.ResourceRegion != nil {
		result["resource_region"] = string(*obj.ResourceRegion)
	}

	result["runbook_relevance"] = string(obj.RunbookRelevance)

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	result["type"] = string(obj.Type)

	return result
}

func (s *FleetAppsManagementRunbookResourceCrud) mapToTask(fieldKeyFormat string) (oci_fleet_apps_management.Task, error) {
	result := oci_fleet_apps_management.Task{}

	if associationType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "association_type")); ok {
		result.AssociationType = oci_fleet_apps_management.TaskAssociationTypeEnum(associationType.(string))
	}

	if outputVariableMappings, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "output_variable_mappings")); ok {
		interfaces := outputVariableMappings.([]interface{})
		tmp := make([]oci_fleet_apps_management.OutputVariableMapping, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "output_variable_mappings"), stateDataIndex)
			converted, err := s.mapToOutputVariableMapping(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "output_variable_mappings")) {
			result.OutputVariableMappings = tmp
		}
	}

	if stepName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_name")); ok {
		tmp := stepName.(string)
		result.StepName = &tmp
	}

	if stepProperties, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_properties")); ok {
		if tmpList := stepProperties.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "step_properties"), 0)
			tmp, err := s.mapToComponentProperties(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert step_properties, encountered error: %v", err)
			}
			result.StepProperties = &tmp
		}
	}

	if taskRecordDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "task_record_details")); ok {
		if tmpList := taskRecordDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "task_record_details"), 0)
			tmp, err := s.mapToAssociatedTaskDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert task_record_details, encountered error: %v", err)
			}
			result.TaskRecordDetails = tmp
		}
	}

	return result, nil
}

func TaskToMap(obj oci_fleet_apps_management.Task) map[string]interface{} {
	result := map[string]interface{}{}

	result["association_type"] = string(obj.AssociationType)

	outputVariableMappings := []interface{}{}
	for _, item := range obj.OutputVariableMappings {
		outputVariableMappings = append(outputVariableMappings, OutputVariableMappingToMap(item))
	}
	result["output_variable_mappings"] = outputVariableMappings

	if obj.StepName != nil {
		result["step_name"] = string(*obj.StepName)
	}

	if obj.StepProperties != nil {
		result["step_properties"] = []interface{}{ComponentPropertiesToMap(obj.StepProperties)}
	}

	if obj.TaskRecordDetails != nil {
		taskRecordDetailsArray := []interface{}{}
		if taskRecordDetailsMap := AssociatedTaskDetailsToMap(&obj.TaskRecordDetails); taskRecordDetailsMap != nil {
			taskRecordDetailsArray = append(taskRecordDetailsArray, taskRecordDetailsMap)
		}
		result["task_record_details"] = taskRecordDetailsArray
	}

	return result
}

func (s *FleetAppsManagementRunbookResourceCrud) mapToTaskNotificationPreferences(fieldKeyFormat string) (oci_fleet_apps_management.TaskNotificationPreferences, error) {
	result := oci_fleet_apps_management.TaskNotificationPreferences{}

	if shouldNotifyOnPause, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_notify_on_pause")); ok {
		tmp := shouldNotifyOnPause.(bool)
		result.ShouldNotifyOnPause = &tmp
	}

	if shouldNotifyOnTaskFailure, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_notify_on_task_failure")); ok {
		tmp := shouldNotifyOnTaskFailure.(bool)
		result.ShouldNotifyOnTaskFailure = &tmp
	}

	if shouldNotifyOnTaskSuccess, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_notify_on_task_success")); ok {
		tmp := shouldNotifyOnTaskSuccess.(bool)
		result.ShouldNotifyOnTaskSuccess = &tmp
	}

	return result, nil
}

func TaskNotificationPreferencesToMap(obj *oci_fleet_apps_management.TaskNotificationPreferences) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ShouldNotifyOnPause != nil {
		result["should_notify_on_pause"] = bool(*obj.ShouldNotifyOnPause)
	}

	if obj.ShouldNotifyOnTaskFailure != nil {
		result["should_notify_on_task_failure"] = bool(*obj.ShouldNotifyOnTaskFailure)
	}

	if obj.ShouldNotifyOnTaskSuccess != nil {
		result["should_notify_on_task_success"] = bool(*obj.ShouldNotifyOnTaskSuccess)
	}

	return result
}

func (s *FleetAppsManagementRunbookResourceCrud) mapToTaskVariable(fieldKeyFormat string) (oci_fleet_apps_management.TaskVariable, error) {
	result := oci_fleet_apps_management.TaskVariable{}

	if inputVariables, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "input_variables")); ok {
		interfaces := inputVariables.([]interface{})
		tmp := make([]oci_fleet_apps_management.InputArgument, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "input_variables"), stateDataIndex)
			converted, err := s.mapToInputArgument(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "input_variables")) {
			result.InputVariables = tmp
		}
	}

	if outputVariables, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "output_variables")); ok {
		interfaces := outputVariables.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "output_variables")) {
			result.OutputVariables = tmp
		}
	}

	return result, nil
}

func (s *FleetAppsManagementRunbookResourceCrud) mapToWorkflowComponent(fieldKeyFormat string) (oci_fleet_apps_management.WorkflowComponent, error) {
	var baseObject oci_fleet_apps_management.WorkflowComponent
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("PARALLEL_TASK_GROUP"):
		details := oci_fleet_apps_management.WorkflowGroupComponent{}
		if groupName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "group_name")); ok {
			tmp := groupName.(string)
			details.GroupName = &tmp
		}
		if steps, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "steps")); ok {
			interfaces := steps.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "steps")) {
				details.Steps = tmp
			}
		}
		baseObject = details
	case strings.ToLower("TASK"):
		details := oci_fleet_apps_management.WorkflowTaskComponent{}
		if stepName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_name")); ok {
			tmp := stepName.(string)
			details.StepName = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func WorkflowComponentToMap(obj oci_fleet_apps_management.WorkflowComponent) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_fleet_apps_management.WorkflowGroupComponent:
		result["type"] = "PARALLEL_TASK_GROUP"

		if v.GroupName != nil {
			result["group_name"] = string(*v.GroupName)
		}

		result["steps"] = v.Steps
	case oci_fleet_apps_management.WorkflowTaskComponent:
		result["type"] = "TASK"

		if v.StepName != nil {
			result["step_name"] = string(*v.StepName)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *FleetAppsManagementRunbookResourceCrud) mapToWorkflowGroup(fieldKeyFormat string) (oci_fleet_apps_management.WorkflowGroup, error) {
	result := oci_fleet_apps_management.WorkflowGroup{}

	if groupName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "group_name")); ok {
		tmp := groupName.(string)
		result.GroupName = &tmp
	}

	if steps, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "steps")); ok {
		interfaces := steps.([]interface{})
		tmp := make([]oci_fleet_apps_management.WorkflowComponent, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "steps"), stateDataIndex)
			converted, err := s.mapToWorkflowComponent(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "steps")) {
			result.Steps = tmp
		}
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_fleet_apps_management.WorkflowGroupTypeEnum(type_.(string))
	}

	return result, nil
}

func WorkflowGroupToMap(obj oci_fleet_apps_management.WorkflowGroup) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.GroupName != nil {
		result["group_name"] = string(*obj.GroupName)
	}

	steps := []interface{}{}
	for _, item := range obj.Steps {
		steps = append(steps, WorkflowComponentToMap(item))
	}
	result["steps"] = steps

	result["type"] = string(obj.Type)

	return result
}
