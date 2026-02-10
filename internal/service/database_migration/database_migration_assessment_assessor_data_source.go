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

func DatabaseMigrationAssessmentAssessorDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseMigrationAssessmentAssessor,
		Schema: map[string]*schema.Schema{
			"assessment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"assessor_name": {
				Type:     schema.TypeString,
				Required: true,
			},
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
	}
}

func readSingularDatabaseMigrationAssessmentAssessor(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationAssessmentAssessorDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.ReadResource(sync)
}

type DatabaseMigrationAssessmentAssessorDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_migration.DatabaseMigrationClient
	Res    *oci_database_migration.GetAssessorResponse
}

func (s *DatabaseMigrationAssessmentAssessorDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseMigrationAssessmentAssessorDataSourceCrud) Get() error {
	request := oci_database_migration.GetAssessorRequest{}

	if assessmentId, ok := s.D.GetOkExists("assessment_id"); ok {
		tmp := assessmentId.(string)
		request.AssessmentId = &tmp
	}

	if assessorName, ok := s.D.GetOkExists("assessor_name"); ok {
		tmp := assessorName.(string)
		request.AssessorName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_migration")

	response, err := s.Client.GetAssessor(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseMigrationAssessmentAssessorDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseMigrationAssessmentAssessorDataSource-", DatabaseMigrationAssessmentAssessorDataSource(), s.D))

	actions := []interface{}{}
	for _, item := range s.Res.Actions {
		actions = append(actions, AssessorActionToMap(item))
	}
	s.D.Set("actions", actions)

	if s.Res.AssessorGroup != nil {
		s.D.Set("assessor_group", []interface{}{AssessorGroupToMap(s.Res.AssessorGroup)})
	} else {
		s.D.Set("assessor_group", nil)
	}

	if s.Res.AssessorResult != nil {
		s.D.Set("assessor_result", *s.Res.AssessorResult)
	}

	if s.Res.ChecksSummary != nil {
		s.D.Set("checks_summary", *s.Res.ChecksSummary)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DoesScriptRequireRestart != nil {
		s.D.Set("does_script_require_restart", *s.Res.DoesScriptRequireRestart)
	}

	if s.Res.HasScript != nil {
		s.D.Set("has_script", *s.Res.HasScript)
	}

	if s.Res.HelpLinkText != nil {
		s.D.Set("help_link_text", *s.Res.HelpLinkText)
	}

	if s.Res.HelpLinkUrl != nil {
		s.D.Set("help_link_url", *s.Res.HelpLinkUrl)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Script != nil {
		s.D.Set("script", *s.Res.Script)
	}

	s.D.Set("state", s.Res.LifecycleState)

	return nil
}
