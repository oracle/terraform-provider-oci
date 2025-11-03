// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_apps_management

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetAppsManagementRunbookVersionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFleetAppsManagementRunbookVersion,
		Read:     readFleetAppsManagementRunbookVersion,
		Update:   updateFleetAppsManagementRunbookVersion,
		Delete:   deleteFleetAppsManagementRunbookVersion,
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
									"pre_condition": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"run_on": {
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
														"PREVIOUS_TASK_INSTANCES",
														"SCHEDULED_INSTANCES",
														"SELF_HOSTED_INSTANCES",
													}, true),
												},

												// Optional
												"condition": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"host": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"previous_task_instance_details": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"output_variable_details": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																MaxItems: 1,
																MinItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional
																		"output_variable_name": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
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
															"resource_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"resource_type": {
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

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"runbook_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tasks": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
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
														"TERRAFORM",
													}, true),
												},

												// Optional
												"catalog_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"command": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"config_file": {
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
															"source_type": {
																Type:             schema.TypeString,
																Required:         true,
																DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
																ValidateFunc: validation.StringInSlice([]string{
																	"CATALOG",
																	"OBJECT_STORAGE_BUCKET",
																}, true),
															},

															// Optional
															"bucket": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"catalog_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"checksum": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"namespace": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"object": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},

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
												"is_executable_content": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"is_locked": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"is_read_output_variable_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"target_compartment_id": {
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
												"system_variables": {
													Type:     schema.TypeList,
													Computed: true,
													Elem:     &schema.Schema{Type: schema.TypeString},
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
									"pre_condition": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"run_on": {
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
														"PREVIOUS_TASK_INSTANCES",
														"SCHEDULED_INSTANCES",
														"SELF_HOSTED_INSTANCES",
													}, true),
												},

												// Optional
												"condition": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"host": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"previous_task_instance_details": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"output_variable_details": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																MaxItems: 1,
																MinItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional
																		"output_variable_name": {
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
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
															"resource_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"resource_type": {
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
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				ForceNew: true,
				// DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				// DiffSuppressFunc: suppressRunbookTransientDiffs,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// k looks like "defined_tags.%", "defined_tags.<key>"
					if strings.HasPrefix(k, "defined_tags.Oracle-Tags.CreatedBy") ||
						strings.HasPrefix(k, "defined_tags.Oracle-Tags.CreatedOn") {
						return true
					}
					return false
				},
				Elem: schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},
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

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_latest": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
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
		},
	}
}

func createFleetAppsManagementRunbookVersion(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementRunbookVersionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementRunbooksClient()
	sync.WorkRequestClient = m.(*client.OracleClients).FleetAppsManagementFleetAppsManagementWorkRequestClient()

	return tfresource.CreateResource(d, sync)
}

func readFleetAppsManagementRunbookVersion(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementRunbookVersionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementRunbooksClient()

	return tfresource.ReadResource(sync)
}

func updateFleetAppsManagementRunbookVersion(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementRunbookVersionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementRunbooksClient()
	sync.WorkRequestClient = m.(*client.OracleClients).FleetAppsManagementFleetAppsManagementWorkRequestClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteFleetAppsManagementRunbookVersion(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementRunbookVersionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementRunbooksClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).FleetAppsManagementFleetAppsManagementWorkRequestClient()

	return tfresource.DeleteResource(d, sync)
}

type FleetAppsManagementRunbookVersionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_fleet_apps_management.FleetAppsManagementRunbooksClient
	Res                    *oci_fleet_apps_management.RunbookVersion
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_fleet_apps_management.FleetAppsManagementWorkRequestClient
}

func (s *FleetAppsManagementRunbookVersionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *FleetAppsManagementRunbookVersionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_fleet_apps_management.RunbookVersionLifecycleStateCreating),
	}
}

func (s *FleetAppsManagementRunbookVersionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.RunbookVersionLifecycleStateActive),
		string(oci_fleet_apps_management.RunbookVersionLifecycleStateNeedsAttention),
		string(oci_fleet_apps_management.RunbookVersionLifecycleStateInactive),
	}
}

func (s *FleetAppsManagementRunbookVersionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_fleet_apps_management.RunbookVersionLifecycleStateDeleting),
	}
}

func (s *FleetAppsManagementRunbookVersionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.RunbookVersionLifecycleStateDeleted),
	}
}

func (s *FleetAppsManagementRunbookVersionResourceCrud) Create() error {
	request := oci_fleet_apps_management.CreateRunbookVersionRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if executionWorkflowDetails, ok := s.D.GetOkExists("execution_workflow_details"); ok {
		if tmpList := executionWorkflowDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "execution_workflow_details", 0)
			tmp, err := s.mapToExecutionWorkflowDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ExecutionWorkflowDetails = &tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if groups, ok := s.D.GetOkExists("groups"); ok {
		interfaces := groups.([]interface{})
		tmp := make([]oci_fleet_apps_management.Group, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "groups", stateDataIndex)
			converted, err := s.mapToGroup(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("groups") {
			request.Groups = tmp
		}
	}

	if rollbackWorkflowDetails, ok := s.D.GetOkExists("rollback_workflow_details"); ok {
		if tmpList := rollbackWorkflowDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rollback_workflow_details", 0)
			tmp, err := s.mapToRollbackWorkflowDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.RollbackWorkflowDetails = &tmp
		}
	}

	if runbookId, ok := s.D.GetOkExists("runbook_id"); ok {
		tmp := runbookId.(string)
		request.RunbookId = &tmp
	}

	if tasks, ok := s.D.GetOkExists("tasks"); ok {
		interfaces := tasks.([]interface{})
		tmp := make([]oci_fleet_apps_management.Task, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tasks", stateDataIndex)
			converted, err := s.mapToTask(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("tasks") {
			request.Tasks = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.CreateRunbookVersion(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getRunbookVersionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management"), oci_fleet_apps_management.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *FleetAppsManagementRunbookVersionResourceCrud) getRunbookVersionFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_fleet_apps_management.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	runbookVersionId, err := runbookVersionWaitForWorkRequest(workId, "runbookversion",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*runbookVersionId)

	return s.Get()
}

func runbookVersionWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func runbookVersionWaitForWorkRequest(wId *string, entityType string, action oci_fleet_apps_management.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_fleet_apps_management.FleetAppsManagementWorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "fleet_apps_management")
	retryPolicy.ShouldRetryOperation = runbookVersionWorkRequestShouldRetryFunc(timeout)

	response := oci_fleet_apps_management.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
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
		return nil, getErrorFromFleetAppsManagementRunbookVersionWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromFleetAppsManagementRunbookVersionWorkRequest(client *oci_fleet_apps_management.FleetAppsManagementWorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_fleet_apps_management.ActionTypeEnum) error {
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

func (s *FleetAppsManagementRunbookVersionResourceCrud) Get() error {
	request := oci_fleet_apps_management.GetRunbookVersionRequest{}

	tmp := s.D.Id()
	request.RunbookVersionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.GetRunbookVersion(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.RunbookVersion
	return nil
}

func (s *FleetAppsManagementRunbookVersionResourceCrud) Update() error {
	request := oci_fleet_apps_management.UpdateRunbookVersionRequest{}

	if executionWorkflowDetails, ok := s.D.GetOkExists("execution_workflow_details"); ok {
		if tmpList := executionWorkflowDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "execution_workflow_details", 0)
			tmp, err := s.mapToExecutionWorkflowDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ExecutionWorkflowDetails = &tmp
		}
	}

	if groups, ok := s.D.GetOkExists("groups"); ok {
		interfaces := groups.([]interface{})
		tmp := make([]oci_fleet_apps_management.Group, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "groups", stateDataIndex)
			converted, err := s.mapToGroup(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("groups") {
			request.Groups = tmp
		}
	}

	if rollbackWorkflowDetails, ok := s.D.GetOkExists("rollback_workflow_details"); ok {
		if tmpList := rollbackWorkflowDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rollback_workflow_details", 0)
			tmp, err := s.mapToRollbackWorkflowDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.RollbackWorkflowDetails = &tmp
		}
	}

	tmp := s.D.Id()
	request.RunbookVersionId = &tmp

	if tasks, ok := s.D.GetOkExists("tasks"); ok {
		interfaces := tasks.([]interface{})
		tmp := make([]oci_fleet_apps_management.Task, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tasks", stateDataIndex)
			converted, err := s.mapToTask(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("tasks") {
			request.Tasks = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.UpdateRunbookVersion(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getRunbookVersionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management"), oci_fleet_apps_management.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *FleetAppsManagementRunbookVersionResourceCrud) Delete() error {
	request := oci_fleet_apps_management.DeleteRunbookVersionRequest{}

	tmp := s.D.Id()
	request.RunbookVersionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.DeleteRunbookVersion(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := runbookVersionWaitForWorkRequest(workId, "runbookversion",
		oci_fleet_apps_management.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *FleetAppsManagementRunbookVersionResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.ExecutionWorkflowDetails != nil {
		s.D.Set("execution_workflow_details", []interface{}{ExecutionWorkflowDetailsToMap(s.Res.ExecutionWorkflowDetails)})
	} else {
		s.D.Set("execution_workflow_details", nil)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	groups := []interface{}{}
	for _, item := range s.Res.Groups {
		groups = append(groups, GroupToMap(item))
	}
	s.D.Set("groups", groups)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.RollbackWorkflowDetails != nil {
		s.D.Set("rollback_workflow_details", []interface{}{RollbackWorkflowDetailsToMap(s.Res.RollbackWorkflowDetails)})
	} else {
		s.D.Set("rollback_workflow_details", nil)
	}

	if s.Res.RunbookId != nil {
		s.D.Set("runbook_id", *s.Res.RunbookId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	} else {
		// FAMS API sometimes returns null rather than {} for empty system_tags.
		systemTags := map[string]interface{}{}
		s.D.Set("system_tags", systemTags)
	}

	tasks := []interface{}{}
	for _, item := range s.Res.Tasks {
		tasks = append(tasks, TaskToMap(item))
	}
	s.D.Set("tasks", tasks)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *FleetAppsManagementRunbookVersionResourceCrud) mapToAssociatedTaskDetails(fieldKeyFormat string) (oci_fleet_apps_management.AssociatedTaskDetails, error) {
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

func (s *FleetAppsManagementRunbookVersionResourceCrud) mapToComponentProperties(fieldKeyFormat string) (oci_fleet_apps_management.ComponentProperties, error) {
	result := oci_fleet_apps_management.ComponentProperties{}

	if actionOnFailure, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action_on_failure")); ok {
		result.ActionOnFailure = oci_fleet_apps_management.ComponentPropertiesActionOnFailureEnum(actionOnFailure.(string))
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

	if preCondition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pre_condition")); ok {
		tmp := preCondition.(string)
		result.PreCondition = &tmp
	}

	if runOn, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "run_on")); ok {
		if tmpList := runOn.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "run_on"), 0)
			tmp, err := s.mapToRunOnDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert run_on, encountered error: %v", err)
			}
			result.RunOn = tmp
		}
	}

	return result, nil
}

func (s *FleetAppsManagementRunbookVersionResourceCrud) mapToConfigAssociationDetails(fieldKeyFormat string) (oci_fleet_apps_management.ConfigAssociationDetails, error) {
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

func (s *FleetAppsManagementRunbookVersionResourceCrud) mapToContentDetails(fieldKeyFormat string) (oci_fleet_apps_management.ContentDetails, error) {
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
	case strings.ToLower("CATALOG"):
		details := oci_fleet_apps_management.CatalogContentDetails{}
		if catalogId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "catalog_id")); ok {
			tmp := catalogId.(string)
			details.CatalogId = &tmp
		}
		baseObject = details
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

func (s *FleetAppsManagementRunbookVersionResourceCrud) mapToExecutionDetails(fieldKeyFormat string) (oci_fleet_apps_management.ExecutionDetails, error) {
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
		if isExecutableContent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_executable_content")); ok {
			tmp := isExecutableContent.(bool)
			details.IsExecutableContent = &tmp
		}
		if isLocked, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_locked")); ok {
			tmp := isLocked.(bool)
			details.IsLocked = &tmp
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
	case strings.ToLower("TERRAFORM"):
		details := oci_fleet_apps_management.TerraformBasedExecutionDetails{}
		if catalogId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "catalog_id")); ok {
			tmp := catalogId.(string)
			details.CatalogId = &tmp
		}
		if configFile, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_file")); ok {
			tmp := configFile.(string)
			details.ConfigFile = &tmp
		}
		if isReadOutputVariableEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_read_output_variable_enabled")); ok {
			tmp := isReadOutputVariableEnabled.(bool)
			details.IsReadOutputVariableEnabled = &tmp
		}
		if targetCompartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_compartment_id")); ok {
			tmp := targetCompartmentId.(string)
			details.TargetCompartmentId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown execution_type '%v' was specified", executionType)
	}
	return baseObject, nil
}

func (s *FleetAppsManagementRunbookVersionResourceCrud) mapToExecutionWorkflowDetails(fieldKeyFormat string) (oci_fleet_apps_management.ExecutionWorkflowDetails, error) {
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

func (s *FleetAppsManagementRunbookVersionResourceCrud) mapToGroup(fieldKeyFormat string) (oci_fleet_apps_management.Group, error) {
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

func (s *FleetAppsManagementRunbookVersionResourceCrud) mapToInputArgument(fieldKeyFormat string) (oci_fleet_apps_management.InputArgument, error) {
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
	case strings.ToLower("FILE"):
		details := oci_fleet_apps_management.FileInputArgument{}
		baseObject = details
	case strings.ToLower("OUTPUT_VARIABLE"):
		details := oci_fleet_apps_management.OutputVariableInputArgument{}
		baseObject = details
	case strings.ToLower("STRING"):
		details := oci_fleet_apps_management.StringInputArgument{}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func (s *FleetAppsManagementRunbookVersionResourceCrud) mapToOutputVariableDetails(fieldKeyFormat string) (oci_fleet_apps_management.OutputVariableDetails, error) {
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

func (s *FleetAppsManagementRunbookVersionResourceCrud) mapToOutputVariableMapping(fieldKeyFormat string) (oci_fleet_apps_management.OutputVariableMapping, error) {
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

func (s *FleetAppsManagementRunbookVersionResourceCrud) mapToPauseDetails(fieldKeyFormat string) (oci_fleet_apps_management.PauseDetails, error) {
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

func (s *FleetAppsManagementRunbookVersionResourceCrud) mapToPreviousTaskInstanceDetails(fieldKeyFormat string) (oci_fleet_apps_management.PreviousTaskInstanceDetails, error) {
	result := oci_fleet_apps_management.PreviousTaskInstanceDetails{}

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

	if resourceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_id")); ok {
		tmp := resourceId.(string)
		result.ResourceId = &tmp
	}

	if resourceType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_type")); ok {
		tmp := resourceType.(string)
		result.ResourceType = &tmp
	}

	return result, nil
}

func (s *FleetAppsManagementRunbookVersionResourceCrud) mapToProperties(fieldKeyFormat string) (oci_fleet_apps_management.Properties, error) {
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

func (s *FleetAppsManagementRunbookVersionResourceCrud) mapToRollbackWorkflowDetails(fieldKeyFormat string) (oci_fleet_apps_management.RollbackWorkflowDetails, error) {
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

func (s *FleetAppsManagementRunbookVersionResourceCrud) mapToRunOnDetails(fieldKeyFormat string) (oci_fleet_apps_management.RunOnDetails, error) {
	var baseObject oci_fleet_apps_management.RunOnDetails
	//discriminator
	kindRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kind"))
	var kind string
	if ok {
		kind = kindRaw.(string)
	} else {
		kind = "" // default value
	}
	switch strings.ToLower(kind) {
	case strings.ToLower("PREVIOUS_TASK_INSTANCES"):
		details := oci_fleet_apps_management.PreviousTaskInstanceRunOnDetails{}
		if previousTaskInstanceDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "previous_task_instance_details")); ok {
			interfaces := previousTaskInstanceDetails.([]interface{})
			tmp := make([]oci_fleet_apps_management.PreviousTaskInstanceDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "previous_task_instance_details"), stateDataIndex)
				converted, err := s.mapToPreviousTaskInstanceDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "previous_task_instance_details")) {
				details.PreviousTaskInstanceDetails = tmp
			}
		}
		baseObject = details
	case strings.ToLower("SCHEDULED_INSTANCES"):
		details := oci_fleet_apps_management.ScheduleInstanceRunOnDetails{}
		if condition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "condition")); ok {
			tmp := condition.(string)
			details.Condition = &tmp
		}
		baseObject = details
	case strings.ToLower("SELF_HOSTED_INSTANCES"):
		details := oci_fleet_apps_management.SelfHostedInstanceRunOnDetails{}
		if host, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "host")); ok {
			tmp := host.(string)
			details.Host = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown kind '%v' was specified", kind)
	}
	return baseObject, nil
}

func RunbookVersionSummaryToMap(obj oci_fleet_apps_management.RunbookVersionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.ExecutionWorkflowDetails != nil {
		result["execution_workflow_details"] = []interface{}{ExecutionWorkflowDetailsToMap(obj.ExecutionWorkflowDetails)}
	}

	result["freeform_tags"] = obj.FreeformTags

	groups := []interface{}{}
	for _, item := range obj.Groups {
		groups = append(groups, GroupToMap(item))
	}
	result["groups"] = groups

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsLatest != nil {
		result["is_latest"] = bool(*obj.IsLatest)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.RollbackWorkflowDetails != nil {
		result["rollback_workflow_details"] = []interface{}{RollbackWorkflowDetailsToMap(obj.RollbackWorkflowDetails)}
	}

	if obj.RunbookId != nil {
		result["runbook_id"] = string(*obj.RunbookId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	tasks := []interface{}{}
	for _, item := range obj.Tasks {
		tasks = append(tasks, TaskToMap(item))
	}
	result["tasks"] = tasks

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *FleetAppsManagementRunbookVersionResourceCrud) mapToTask(fieldKeyFormat string) (oci_fleet_apps_management.Task, error) {
	result := oci_fleet_apps_management.Task{}

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

func (s *FleetAppsManagementRunbookVersionResourceCrud) mapToTaskNotificationPreferences(fieldKeyFormat string) (oci_fleet_apps_management.TaskNotificationPreferences, error) {
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

func (s *FleetAppsManagementRunbookVersionResourceCrud) mapToTaskVariable(fieldKeyFormat string) (oci_fleet_apps_management.TaskVariable, error) {
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

func (s *FleetAppsManagementRunbookVersionResourceCrud) mapToWorkflowComponent(fieldKeyFormat string) (oci_fleet_apps_management.WorkflowComponent, error) {
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

func (s *FleetAppsManagementRunbookVersionResourceCrud) mapToWorkflowGroup(fieldKeyFormat string) (oci_fleet_apps_management.WorkflowGroup, error) {
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
