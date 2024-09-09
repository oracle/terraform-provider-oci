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

func FleetAppsManagementTaskRecordsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFleetAppsManagementTaskRecords,
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
			"platform": {
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
			"task_record_collection": {
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
									"details": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
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
											},
										},
									},
									"display_name": {
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
									"version": {
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

func readFleetAppsManagementTaskRecords(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementTaskRecordsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementRunbooksClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementTaskRecordsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementRunbooksClient
	Res    *oci_fleet_apps_management.ListTaskRecordsResponse
}

func (s *FleetAppsManagementTaskRecordsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementTaskRecordsDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.ListTaskRecordsRequest{}

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

	if platform, ok := s.D.GetOkExists("platform"); ok {
		tmp := platform.(string)
		request.Platform = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_fleet_apps_management.TaskRecordLifecycleStateEnum(state.(string))
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_fleet_apps_management.TaskRecordTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.ListTaskRecords(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListTaskRecords(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FleetAppsManagementTaskRecordsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FleetAppsManagementTaskRecordsDataSource-", FleetAppsManagementTaskRecordsDataSource(), s.D))
	resources := []map[string]interface{}{}
	taskRecord := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, TaskRecordSummaryToMap(item))
	}
	taskRecord["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, FleetAppsManagementTaskRecordsDataSource().Schema["task_record_collection"].Elem.(*schema.Resource).Schema)
		taskRecord["items"] = items
	}

	resources = append(resources, taskRecord)
	if err := s.D.Set("task_record_collection", resources); err != nil {
		return err
	}

	return nil
}

func ContentDetailsToMap(obj *oci_fleet_apps_management.ContentDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_fleet_apps_management.ObjectStorageBucketContentDetails:
		result["source_type"] = "OBJECT_STORAGE_BUCKET"

		if v.BucketName != nil {
			result["bucket"] = string(*v.BucketName)
		}

		if v.Checksum != nil {
			result["checksum"] = string(*v.Checksum)
		}

		if v.NamespaceName != nil {
			result["namespace"] = string(*v.NamespaceName)
		}

		if v.ObjectName != nil {
			result["object"] = string(*v.ObjectName)
		}
	default:
		log.Printf("[WARN] Received 'source_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func DetailsToMap(obj *oci_fleet_apps_management.Details) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ExecutionDetails != nil {
		executionDetailsArray := []interface{}{}
		if executionDetailsMap := ExecutionDetailsToMap(&obj.ExecutionDetails); executionDetailsMap != nil {
			executionDetailsArray = append(executionDetailsArray, executionDetailsMap)
		}
		result["execution_details"] = executionDetailsArray
	}

	result["os_type"] = string(obj.OsType)

	if obj.Platform != nil {
		result["platform"] = string(*obj.Platform)
	}

	if obj.Properties != nil {
		result["properties"] = []interface{}{PropertiesToMap(obj.Properties)}
	}

	result["scope"] = string(obj.Scope)

	return result
}

func ExecutionDetailsToMap(obj *oci_fleet_apps_management.ExecutionDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_fleet_apps_management.ApiBasedExecutionDetails:
		result["execution_type"] = "API"

		if v.Endpoint != nil {
			result["endpoint"] = string(*v.Endpoint)
		}
	case oci_fleet_apps_management.ScriptBasedExecutionDetails:
		result["execution_type"] = "SCRIPT"

		if v.Command != nil {
			result["command"] = string(*v.Command)
		}

		if v.Content != nil {
			contentArray := []interface{}{}
			if contentMap := ContentDetailsToMap(&v.Content); contentMap != nil {
				contentArray = append(contentArray, contentMap)
			}
			result["content"] = contentArray
		}

		if v.Variables != nil {
			result["variables"] = []interface{}{TaskVariableToMap(v.Variables)}
		}
	default:
		log.Printf("[WARN] Received 'execution_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func InputArgumentToMap(obj oci_fleet_apps_management.InputArgument) map[string]interface{} {
	result := map[string]interface{}{}
	switch (obj).(type) {
	case oci_fleet_apps_management.OutputVariableInputArgument:
		result["type"] = "OUTPUT_VARIABLE"
	case oci_fleet_apps_management.StringInputArgument:
		result["type"] = "STRING"
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func PropertiesToMap(obj *oci_fleet_apps_management.Properties) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.NumRetries != nil {
		result["num_retries"] = int(*obj.NumRetries)
	}

	if obj.TimeoutInSeconds != nil {
		result["timeout_in_seconds"] = int(*obj.TimeoutInSeconds)
	}

	return result
}

func TaskRecordSummaryToMap(obj oci_fleet_apps_management.TaskRecordSummary) map[string]interface{} {
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

	if obj.Details != nil {
		result["details"] = []interface{}{DetailsToMap(obj.Details)}
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

	if obj.ResourceRegion != nil {
		result["resource_region"] = string(*obj.ResourceRegion)
	}

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

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}

func TaskVariableToMap(obj *oci_fleet_apps_management.TaskVariable) map[string]interface{} {
	result := map[string]interface{}{}

	inputVariables := []interface{}{}
	for _, item := range obj.InputVariables {
		inputVariables = append(inputVariables, InputArgumentToMap(item))
	}
	result["input_variables"] = inputVariables

	result["output_variables"] = obj.OutputVariables

	return result
}
