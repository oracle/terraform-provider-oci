// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_migration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_migration "github.com/oracle/oci-go-sdk/v65/databasemigration"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseMigrationAssessmentAssessorsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseMigrationAssessmentAssessors,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"assessment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"assessor_summary_collection": {
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
									"actions": {
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
												"display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"is_disabled": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"resource_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"title": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"user_defined_properties": {
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
															"display_name": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"help_link_text": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"help_link_url": {
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
																		"default_value": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"description": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"display_name": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"is_required": {
																			Type:     schema.TypeBool,
																			Computed: true,
																		},
																		"max_length": {
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"min_length": {
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"name": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"options": {
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
																					"display_name": {
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"value": {
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
																		"value": {
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
										},
									},
									"assessment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"assessor_group": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"actions": {
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
															"display_name": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"is_disabled": {
																Type:     schema.TypeBool,
																Computed: true,
															},
															"name": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"resource_id": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"title": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"user_defined_properties": {
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
																		"display_name": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"help_link_text": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"help_link_url": {
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
																					"default_value": {
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"description": {
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"display_name": {
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"is_required": {
																						Type:     schema.TypeBool,
																						Computed: true,
																					},
																					"max_length": {
																						Type:     schema.TypeInt,
																						Computed: true,
																					},
																					"min_length": {
																						Type:     schema.TypeInt,
																						Computed: true,
																					},
																					"name": {
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"options": {
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
																								"display_name": {
																									Type:     schema.TypeString,
																									Computed: true,
																								},
																								"value": {
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
																					"value": {
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
													},
												},
												"description": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"assessor_result": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"checks_summary": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"does_script_require_restart": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"has_script": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"help_link_text": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"help_link_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"script": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
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

func readDatabaseMigrationAssessmentAssessors(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationAssessmentAssessorsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.ReadResource(sync)
}

type DatabaseMigrationAssessmentAssessorsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_migration.DatabaseMigrationClient
	Res    *oci_database_migration.ListAssessorsResponse
}

func (s *DatabaseMigrationAssessmentAssessorsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseMigrationAssessmentAssessorsDataSourceCrud) Get() error {
	request := oci_database_migration.ListAssessorsRequest{}

	if assessmentId, ok := s.D.GetOkExists("assessment_id"); ok {
		tmp := assessmentId.(string)
		request.AssessmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database_migration.ListAssessorsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_migration")

	response, err := s.Client.ListAssessors(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAssessors(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseMigrationAssessmentAssessorsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseMigrationAssessmentAssessorsDataSource-", DatabaseMigrationAssessmentAssessorsDataSource(), s.D))
	resources := []map[string]interface{}{}
	assessmentAssessor := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AssessorSummaryToMap(item))
	}
	assessmentAssessor["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseMigrationAssessmentAssessorsDataSource().Schema["assessor_summary_collection"].Elem.(*schema.Resource).Schema)
		assessmentAssessor["items"] = items
	}

	resources = append(resources, assessmentAssessor)
	if err := s.D.Set("assessor_summary_collection", resources); err != nil {
		return err
	}

	return nil
}

func AssessorActionToMap(obj oci_database_migration.AssessorAction) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.IsDisabled != nil {
		result["is_disabled"] = bool(*obj.IsDisabled)
	}

	result["name"] = string(obj.Name)

	if obj.ResourceId != nil {
		result["resource_id"] = string(*obj.ResourceId)
	}

	if obj.Title != nil {
		result["title"] = string(*obj.Title)
	}

	if obj.UserDefinedProperties != nil {
		result["user_defined_properties"] = []interface{}{UserDefinedPropertiesToMap(obj.UserDefinedProperties)}
	}

	return result
}

func AssessorGroupToMap(obj *oci_database_migration.AssessorGroup) map[string]interface{} {
	result := map[string]interface{}{}

	actions := []interface{}{}
	for _, item := range obj.Actions {
		actions = append(actions, AssessorActionToMap(item))
	}
	result["actions"] = actions

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func AssessorSummaryToMap(obj oci_database_migration.AssessorSummary) map[string]interface{} {
	result := map[string]interface{}{}

	actions := []interface{}{}
	for _, item := range obj.Actions {
		actions = append(actions, AssessorActionToMap(item))
	}
	result["actions"] = actions

	if obj.AssessmentId != nil {
		result["assessment_id"] = string(*obj.AssessmentId)
	}

	if obj.AssessorGroup != nil {
		result["assessor_group"] = []interface{}{AssessorGroupToMap(obj.AssessorGroup)}
	}

	if obj.AssessorResult != nil {
		result["assessor_result"] = string(*obj.AssessorResult)
	}

	if obj.ChecksSummary != nil {
		result["checks_summary"] = string(*obj.ChecksSummary)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.DoesScriptRequireRestart != nil {
		result["does_script_require_restart"] = bool(*obj.DoesScriptRequireRestart)
	}

	if obj.HasScript != nil {
		result["has_script"] = bool(*obj.HasScript)
	}

	if obj.HelpLinkText != nil {
		result["help_link_text"] = string(*obj.HelpLinkText)
	}

	if obj.HelpLinkUrl != nil {
		result["help_link_url"] = string(*obj.HelpLinkUrl)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Script != nil {
		result["script"] = string(*obj.Script)
	}

	result["state"] = string(obj.LifecycleState)

	return result
}

func UserDefinedPropertiesToMap(obj *oci_database_migration.UserDefinedProperties) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.HelpLinkText != nil {
		result["help_link_text"] = string(*obj.HelpLinkText)
	}

	if obj.HelpLinkUrl != nil {
		result["help_link_url"] = string(*obj.HelpLinkUrl)
	}

	properties := []interface{}{}
	for _, item := range obj.Properties {
		properties = append(properties, UserDefinedPropertyToMap(item))
	}
	result["properties"] = properties

	return result
}

func UserDefinedPropertyToMap(obj oci_database_migration.UserDefinedProperty) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefaultValue != nil {
		result["default_value"] = string(*obj.DefaultValue)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.IsRequired != nil {
		result["is_required"] = bool(*obj.IsRequired)
	}

	if obj.MaxLength != nil {
		result["max_length"] = int(*obj.MaxLength)
	}

	if obj.MinLength != nil {
		result["min_length"] = int(*obj.MinLength)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	options := []interface{}{}
	for _, item := range obj.Options {
		options = append(options, UserDefinedPropertyOptionToMap(item))
	}
	result["options"] = options

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func UserDefinedPropertyOptionToMap(obj oci_database_migration.UserDefinedPropertyOption) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}
