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

func DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecution,
		Schema: map[string]*schema.Schema{
			"execution_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"managed_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"task_name": {
				Type:     schema.TypeString,
				Required: true,
			},
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
			"time_end": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_start": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecution(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetOptimizerStatisticsAdvisorExecutionResponse
}

func (s *DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionDataSourceCrud) Get() error {
	request := oci_database_management.GetOptimizerStatisticsAdvisorExecutionRequest{}

	if executionName, ok := s.D.GetOkExists("execution_name"); ok {
		tmp := executionName.(string)
		request.ExecutionName = &tmp
	}

	if managedDatabaseId, ok := s.D.GetOkExists("managed_database_id"); ok {
		tmp := managedDatabaseId.(string)
		request.ManagedDatabaseId = &tmp
	}

	if taskName, ok := s.D.GetOkExists("task_name"); ok {
		tmp := taskName.(string)
		request.TaskName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetOptimizerStatisticsAdvisorExecution(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionDataSource-", DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionDataSource(), s.D))

	if s.Res.Database != nil {
		s.D.Set("database", []interface{}{ExecutionOptimizerDatabaseToMap(s.Res.Database)})
	} else {
		s.D.Set("database", nil)
	}

	if s.Res.ErrorMessage != nil {
		s.D.Set("error_message", *s.Res.ErrorMessage)
	}

	if s.Res.Findings != nil {
		s.D.Set("findings", *s.Res.Findings)
	}

	if s.Res.Report != nil {
		s.D.Set("report", []interface{}{ExecutionOptimizerStatisticsAdvisorExecutionReportToMap(s.Res.Report)})
	} else {
		s.D.Set("report", nil)
	}

	s.D.Set("status", s.Res.Status)

	if s.Res.StatusMessage != nil {
		s.D.Set("status_message", *s.Res.StatusMessage)
	}

	if s.Res.TimeEnd != nil {
		s.D.Set("time_end", s.Res.TimeEnd.String())
	}

	if s.Res.TimeStart != nil {
		s.D.Set("time_start", s.Res.TimeStart.String())
	}

	return nil
}

func ExecutionAdvisorRuleToMap(obj oci_database_management.AdvisorRule) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	findings := []interface{}{}
	for _, item := range obj.Findings {
		findings = append(findings, ExecutionRuleFindingToMap(item))
	}
	result["findings"] = findings

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func ExecutionFindingSchemaOrOperationToMap(obj oci_database_management.FindingSchemaOrOperation) map[string]interface{} {
	result := map[string]interface{}{}

	result["operations"] = obj.Operations

	schemas := []interface{}{}
	for _, item := range obj.Schemas {
		schemas = append(schemas, ExecutionSchemaDefinitionToMap(item))
	}
	result["schemas"] = schemas

	return result
}

func ExecutionOptimizerDatabaseToMap(obj *oci_database_management.OptimizerDatabase) map[string]interface{} {
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

func ExecutionOptimizerStatisticsAdvisorExecutionReportToMap(obj *oci_database_management.OptimizerStatisticsAdvisorExecutionReport) map[string]interface{} {
	result := map[string]interface{}{}

	rules := []interface{}{}
	for _, item := range obj.Rules {
		rules = append(rules, ExecutionAdvisorRuleToMap(item))
	}
	result["rules"] = rules

	if obj.Summary != nil {
		result["summary"] = string(*obj.Summary)
	}

	return result
}

func ExecutionOptimizerStatisticsAdvisorExecutionSummaryToMap(obj oci_database_management.OptimizerStatisticsAdvisorExecutionSummary) map[string]interface{} {
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

func ExecutionRecommendationToMap(obj oci_database_management.Recommendation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Example != nil {
		result["example"] = []interface{}{ExecutionRecommendationExampleToMap(obj.Example)}
	}

	if obj.Message != nil {
		result["message"] = string(*obj.Message)
	}

	rationales := []interface{}{}
	for _, item := range obj.Rationales {
		rationales = append(rationales, ExecutionRecommendationRationaleToMap(item))
	}
	result["rationales"] = rationales

	return result
}

func ExecutionRecommendationExampleToMap(obj *oci_database_management.RecommendationExample) map[string]interface{} {
	result := map[string]interface{}{}

	lines := []interface{}{}
	for _, item := range obj.Lines {
		lines = append(lines, RecommendationExampleLineToMap(item))
	}
	result["lines"] = lines

	return result
}

func ExecutionRecommendationExampleLineToMap(obj oci_database_management.RecommendationExampleLine) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Comment != nil {
		result["comment"] = string(*obj.Comment)
	}

	if obj.Operation != nil {
		result["operation"] = string(*obj.Operation)
	}

	return result
}

func ExecutionRecommendationRationaleToMap(obj oci_database_management.RecommendationRationale) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Message != nil {
		result["message"] = string(*obj.Message)
	}

	return result
}

func ExecutionRuleFindingToMap(obj oci_database_management.RuleFinding) map[string]interface{} {
	result := map[string]interface{}{}

	details := []interface{}{}
	for _, item := range obj.Details {
		details = append(details, ExecutionFindingSchemaOrOperationToMap(item))
	}
	result["details"] = details

	if obj.Message != nil {
		result["message"] = string(*obj.Message)
	}

	recommendations := []interface{}{}
	for _, item := range obj.Recommendations {
		recommendations = append(recommendations, ExecutionRecommendationToMap(item))
	}
	result["recommendations"] = recommendations

	return result
}

func ExecutionSchemaDefinitionToMap(obj oci_database_management.SchemaDefinition) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["objects"] = obj.Objects

	return result
}
