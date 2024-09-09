// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_apps_management

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetAppsManagementRunbooksDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFleetAppsManagementRunbooks,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"operation": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"platform": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"runbook_relevance": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"runbook_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"associations": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"execution_workflow_details": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"workflow": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional

																		// Computed
																		"group_name": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"steps": {
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					// Required

																					// Optional

																					// Computed
																					"group_name": {
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"step_name": {
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"steps": {
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},
																					"type": {
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"type": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
														},
													},
												},
												"groups": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"name": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"properties": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional

																		// Computed
																		"action_on_failure": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"condition": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"run_on": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"type": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"tasks": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"association_type": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"output_variable_mappings": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional

																		// Computed
																		"name": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"output_variable_details": {
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					// Required

																					// Optional

																					// Computed
																					"output_variable_name": {
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"step_name": {
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"step_name": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"step_properties": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional

																		// Computed
																		"action_on_failure": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"condition": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"run_on": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"task_record_details": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional

																		// Computed
																		"description": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"execution_details": {
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					// Required

																					// Optional

																					// Computed
																					"command": {
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"content": {
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								// Required

																								// Optional

																								// Computed
																								"bucket": {
																									Type:     schema.TypeString,
																									Computed: true,
																								},
																								"checksum": {
																									Type:     schema.TypeString,
																									Computed: true,
																								},
																								"namespace": {
																									Type:     schema.TypeString,
																									Computed: true,
																								},
																								"object": {
																									Type:     schema.TypeString,
																									Computed: true,
																								},
																								"source_type": {
																									Type:     schema.TypeString,
																									Computed: true,
																								},
																							},
																						},
																					},
																					"endpoint": {
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"execution_type": {
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"variables": {
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								// Required

																								// Optional

																								// Computed
																								"input_variables": {
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											// Required

																											// Optional

																											// Computed
																											"description": {
																												Type:     schema.TypeString,
																												Computed: true,
																											},
																											"name": {
																												Type:     schema.TypeString,
																												Computed: true,
																											},
																											"type": {
																												Type:     schema.TypeString,
																												Computed: true,
																											},
																										},
																									},
																								},
																								"output_variables": {
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																		"is_copy_to_library_enabled": {
																			Type:     schema.TypeBool,
																			Computed: true,
																		},
																		"name": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"os_type": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"platform": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"properties": {
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					// Required

																					// Optional

																					// Computed
																					"num_retries": {
																						Type:     schema.TypeInt,
																						Computed: true,
																					},
																					"timeout_in_seconds": {
																						Type:     schema.TypeInt,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"scope": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"task_record_id": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
														},
													},
												},
												"version": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"estimated_time": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_default": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"operation": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"os_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"platform": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_region": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"runbook_relevance": {
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
							},
						},
					},
				},
			},
		},
	}
}

func readFleetAppsManagementRunbooks(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementRunbooksDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementRunbooksClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementRunbooksDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementRunbooksClient
	Res    *oci_fleet_apps_management.ListRunbooksResponse
}

func (s *FleetAppsManagementRunbooksDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementRunbooksDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.ListRunbooksRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if operation, ok := s.D.GetOkExists("operation"); ok {
		tmp := operation.(string)
		request.Operation = &tmp
	}

	if platform, ok := s.D.GetOkExists("platform"); ok {
		tmp := platform.(string)
		request.Platform = &tmp
	}

	if runbookRelevance, ok := s.D.GetOkExists("runbook_relevance"); ok {
		request.RunbookRelevance = oci_fleet_apps_management.RunbookRunbookRelevanceEnum(runbookRelevance.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_fleet_apps_management.RunbookLifecycleStateEnum(state.(string))
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_fleet_apps_management.RunbookTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.ListRunbooks(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListRunbooks(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FleetAppsManagementRunbooksDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FleetAppsManagementRunbooksDataSource-", FleetAppsManagementRunbooksDataSource(), s.D))
	resources := []map[string]interface{}{}
	runbook := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RunbookSummaryToMap(item))
	}
	runbook["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FleetAppsManagementRunbooksDataSource().Schema["runbook_collection"].Elem.(*schema.Resource).Schema)
		runbook["items"] = items
	}

	resources = append(resources, runbook)
	if err := s.D.Set("runbook_collection", resources); err != nil {
		return err
	}

	return nil
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

		if v.IsCopyToLibraryEnabled != nil {
			result["is_copy_to_library_enabled"] = bool(*v.IsCopyToLibraryEnabled)
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

func ComponentPropertiesToMap(obj *oci_fleet_apps_management.ComponentProperties) map[string]interface{} {
	result := map[string]interface{}{}

	result["action_on_failure"] = string(obj.ActionOnFailure)

	if obj.Condition != nil {
		result["condition"] = string(*obj.Condition)
	}

	if obj.RunOn != nil {
		result["run_on"] = string(*obj.RunOn)
	}

	return result
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
