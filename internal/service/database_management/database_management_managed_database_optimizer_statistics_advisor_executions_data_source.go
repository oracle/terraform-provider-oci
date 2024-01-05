// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"end_time_less_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"start_time_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"optimizer_statistics_advisor_executions_collection": {
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
									"database": {
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
												"db_deployment_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"db_sub_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"db_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"db_version": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"id": {
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
									"error_message": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"execution_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"findings": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"report": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"rules": {
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
															"findings": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional

																		// Computed
																		"details": {
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					// Required

																					// Optional

																					// Computed
																					"operations": {
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},
																					"schemas": {
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
																								"objects": {
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
																		"message": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"recommendations": {
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					// Required

																					// Optional

																					// Computed
																					"example": {
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								// Required

																								// Optional

																								// Computed
																								"lines": {
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											// Required

																											// Optional

																											// Computed
																											"comment": {
																												Type:     schema.TypeString,
																												Computed: true,
																											},
																											"operation": {
																												Type:     schema.TypeString,
																												Computed: true,
																											},
																										},
																									},
																								},
																							},
																						},
																					},
																					"message": {
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																					"rationales": {
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								// Required

																								// Optional

																								// Computed
																								"message": {
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
															"name": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"summary": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status_message": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"task_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_end": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_start": {
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

func readDatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutions(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListOptimizerStatisticsAdvisorExecutionsResponse
}

func (s *DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionsDataSourceCrud) Get() error {
	request := oci_database_management.ListOptimizerStatisticsAdvisorExecutionsRequest{}

	if endTimeLessThanOrEqualTo, ok := s.D.GetOkExists("end_time_less_than_or_equal_to"); ok {
		tmp := endTimeLessThanOrEqualTo.(string)
		request.EndTimeLessThanOrEqualTo = &tmp
	}

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	if startTimeGreaterThanOrEqualTo, ok := s.D.GetOkExists("start_time_greater_than_or_equal_to"); ok {
		tmp := startTimeGreaterThanOrEqualTo.(string)
		request.StartTimeGreaterThanOrEqualTo = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListOptimizerStatisticsAdvisorExecutions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionsDataSource-", DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedDatabaseOptimizerStatisticsAdvisorExecution := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OptimizerStatisticsAdvisorExecutionSummaryToMap(item))
	}
	managedDatabaseOptimizerStatisticsAdvisorExecution["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionsDataSource().Schema["optimizer_statistics_advisor_executions_collection"].Elem.(*schema.Resource).Schema)
		managedDatabaseOptimizerStatisticsAdvisorExecution["items"] = items
	}

	resources = append(resources, managedDatabaseOptimizerStatisticsAdvisorExecution)
	if err := s.D.Set("optimizer_statistics_advisor_executions_collection", resources); err != nil {
		return err
	}

	return nil
}

func AdvisorRuleToMap(obj oci_database_management.AdvisorRule) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	findings := []interface{}{}
	for _, item := range obj.Findings {
		findings = append(findings, RuleFindingToMap(item))
	}
	result["findings"] = findings

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func FindingSchemaOrOperationToMap(obj oci_database_management.FindingSchemaOrOperation) map[string]interface{} {
	result := map[string]interface{}{}

	result["operations"] = obj.Operations

	schemas := []interface{}{}
	for _, item := range obj.Schemas {
		schemas = append(schemas, SchemaDefinitionToMap(item))
	}
	result["schemas"] = schemas

	return result
}

func ExecutionsOptimizerDatabaseToMap(obj *oci_database_management.OptimizerDatabase) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["db_deployment_type"] = string(obj.DbDeploymentType)

	result["db_sub_type"] = string(obj.DbSubType)

	result["db_type"] = string(obj.DbType)

	if obj.DbVersion != nil {
		result["db_version"] = string(*obj.DbVersion)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func OptimizerStatisticsAdvisorExecutionReportToMap(obj *oci_database_management.OptimizerStatisticsAdvisorExecutionReport) map[string]interface{} {
	result := map[string]interface{}{}

	rules := []interface{}{}
	for _, item := range obj.Rules {
		rules = append(rules, AdvisorRuleToMap(item))
	}
	result["rules"] = rules

	if obj.Summary != nil {
		result["summary"] = string(*obj.Summary)
	}

	return result
}

func OptimizerStatisticsAdvisorExecutionSummaryToMap(obj oci_database_management.OptimizerStatisticsAdvisorExecutionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ErrorMessage != nil {
		result["error_message"] = string(*obj.ErrorMessage)
	}

	if obj.ExecutionName != nil {
		result["execution_name"] = string(*obj.ExecutionName)
	}

	if obj.Findings != nil {
		result["findings"] = int(*obj.Findings)
	}

	result["status"] = string(obj.Status)

	if obj.StatusMessage != nil {
		result["status_message"] = string(*obj.StatusMessage)
	}

	if obj.TaskName != nil {
		result["task_name"] = string(*obj.TaskName)
	}

	if obj.TimeEnd != nil {
		result["time_end"] = obj.TimeEnd.String()
	}

	if obj.TimeStart != nil {
		result["time_start"] = obj.TimeStart.String()
	}

	return result
}

func RecommendationToMap(obj oci_database_management.Recommendation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Example != nil {
		result["example"] = []interface{}{RecommendationExampleToMap(obj.Example)}
	}

	if obj.Message != nil {
		result["message"] = string(*obj.Message)
	}

	rationales := []interface{}{}
	for _, item := range obj.Rationales {
		rationales = append(rationales, RecommendationRationaleToMap(item))
	}
	result["rationales"] = rationales

	return result
}

func RecommendationExampleToMap(obj *oci_database_management.RecommendationExample) map[string]interface{} {
	result := map[string]interface{}{}

	lines := []interface{}{}
	for _, item := range obj.Lines {
		lines = append(lines, RecommendationExampleLineToMap(item))
	}
	result["lines"] = lines

	return result
}

func RecommendationExampleLineToMap(obj oci_database_management.RecommendationExampleLine) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Comment != nil {
		result["comment"] = string(*obj.Comment)
	}

	if obj.Operation != nil {
		result["operation"] = string(*obj.Operation)
	}

	return result
}

func RecommendationRationaleToMap(obj oci_database_management.RecommendationRationale) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Message != nil {
		result["message"] = string(*obj.Message)
	}

	return result
}

func RuleFindingToMap(obj oci_database_management.RuleFinding) map[string]interface{} {
	result := map[string]interface{}{}

	details := []interface{}{}
	for _, item := range obj.Details {
		details = append(details, FindingSchemaOrOperationToMap(item))
	}
	result["details"] = details

	if obj.Message != nil {
		result["message"] = string(*obj.Message)
	}

	recommendations := []interface{}{}
	for _, item := range obj.Recommendations {
		recommendations = append(recommendations, RecommendationToMap(item))
	}
	result["recommendations"] = recommendations

	return result
}

func SchemaDefinitionToMap(obj oci_database_management.SchemaDefinition) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["objects"] = obj.Objects

	return result
}
