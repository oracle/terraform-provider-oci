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

func DatabaseMigrationAssessmentAssessorChecksDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseMigrationAssessmentAssessorChecks,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"assessment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"assessor_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"assessor_check_summary_collection": {
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
									"action": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"assessor_check_group": {
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
												"is_expanded": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"assessor_check_state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"check_action": {
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
												"name": {
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
									"columns": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"key": {
													Type:     schema.TypeString,
													Computed: true,
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
									"fixup_script_location": {
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
									"impact": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_exclusion_allowed": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"issue": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"log_location": {
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
												"namespace": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"object": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"metadata": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"object_name_column": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"object_type_column": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"object_type_fixed": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"schema_owner_column": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"object_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"objects_display_name": {
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

func readDatabaseMigrationAssessmentAssessorChecks(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationAssessmentAssessorChecksDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.ReadResource(sync)
}

type DatabaseMigrationAssessmentAssessorChecksDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_migration.DatabaseMigrationClient
	Res    *oci_database_migration.ListAssessorChecksResponse
}

func (s *DatabaseMigrationAssessmentAssessorChecksDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseMigrationAssessmentAssessorChecksDataSourceCrud) Get() error {
	request := oci_database_migration.ListAssessorChecksRequest{}

	if assessmentId, ok := s.D.GetOkExists("assessment_id"); ok {
		tmp := assessmentId.(string)
		request.AssessmentId = &tmp
	}

	if assessorName, ok := s.D.GetOkExists("assessor_name"); ok {
		tmp := assessorName.(string)
		request.AssessorName = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_migration")

	response, err := s.Client.ListAssessorChecks(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAssessorChecks(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseMigrationAssessmentAssessorChecksDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseMigrationAssessmentAssessorChecksDataSource-", DatabaseMigrationAssessmentAssessorChecksDataSource(), s.D))
	resources := []map[string]interface{}{}
	assessmentAssessorCheck := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AssessorCheckSummaryToMap(item))
	}
	assessmentAssessorCheck["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseMigrationAssessmentAssessorChecksDataSource().Schema["assessor_check_summary_collection"].Elem.(*schema.Resource).Schema)
		assessmentAssessorCheck["items"] = items
	}

	resources = append(resources, assessmentAssessorCheck)
	if err := s.D.Set("assessor_check_summary_collection", resources); err != nil {
		return err
	}

	return nil
}

func AssessorCheckActionToMap(obj *oci_database_migration.AssessorCheckAction) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["name"] = string(obj.Name)

	if obj.UserDefinedProperties != nil {
		result["user_defined_properties"] = []interface{}{UserDefinedPropertiesToMap(obj.UserDefinedProperties)}
	}

	return result
}

func AssessorCheckGroupToMap(obj *oci_database_migration.AssessorCheckGroup) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.IsExpanded != nil {
		result["is_expanded"] = bool(*obj.IsExpanded)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func AssessorCheckSummaryToMap(obj oci_database_migration.AssessorCheckSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Action != nil {
		result["action"] = string(*obj.Action)
	}

	if obj.AssessorCheckGroup != nil {
		result["assessor_check_group"] = []interface{}{AssessorCheckGroupToMap(obj.AssessorCheckGroup)}
	}

	result["assessor_check_state"] = string(obj.AssessorCheckState)

	if obj.CheckAction != nil {
		result["check_action"] = []interface{}{AssessorCheckActionToMap(obj.CheckAction)}
	}

	columns := []interface{}{}
	for _, item := range obj.Columns {
		columns = append(columns, AdvisorReportCheckColumnToMap(item))
	}
	result["columns"] = columns

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.FixupScriptLocation != nil {
		result["fixup_script_location"] = string(*obj.FixupScriptLocation)
	}

	if obj.HelpLinkText != nil {
		result["help_link_text"] = string(*obj.HelpLinkText)
	}

	if obj.HelpLinkUrl != nil {
		result["help_link_url"] = string(*obj.HelpLinkUrl)
	}

	if obj.Impact != nil {
		result["impact"] = string(*obj.Impact)
	}

	if obj.IsExclusionAllowed != nil {
		result["is_exclusion_allowed"] = bool(*obj.IsExclusionAllowed)
	}

	if obj.Issue != nil {
		result["issue"] = string(*obj.Issue)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.LogLocation != nil {
		result["log_location"] = []interface{}{LogLocationBucketDetailsToMap(obj.LogLocation)}
	}

	if obj.Metadata != nil {
		result["metadata"] = []interface{}{ObjectMetadataToMap(obj.Metadata)}
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ObjectCount != nil {
		result["object_count"] = int(*obj.ObjectCount)
	}

	if obj.ObjectsDisplayName != nil {
		result["objects_display_name"] = string(*obj.ObjectsDisplayName)
	}

	return result
}
