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

func DatabaseMigrationAssessmentAssessorCheckDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseMigrationAssessmentAssessorCheck,
		Schema: map[string]*schema.Schema{
			"assessment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"assessor_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"check_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
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
	}
}

func readSingularDatabaseMigrationAssessmentAssessorCheck(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationAssessmentAssessorCheckDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.ReadResource(sync)
}

type DatabaseMigrationAssessmentAssessorCheckDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_migration.DatabaseMigrationClient
	Res    *oci_database_migration.GetAssessorCheckResponse
}

func (s *DatabaseMigrationAssessmentAssessorCheckDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseMigrationAssessmentAssessorCheckDataSourceCrud) Get() error {
	request := oci_database_migration.GetAssessorCheckRequest{}

	if assessmentId, ok := s.D.GetOkExists("assessment_id"); ok {
		tmp := assessmentId.(string)
		request.AssessmentId = &tmp
	}

	if assessorName, ok := s.D.GetOkExists("assessor_name"); ok {
		tmp := assessorName.(string)
		request.AssessorName = &tmp
	}

	if checkName, ok := s.D.GetOkExists("check_name"); ok {
		tmp := checkName.(string)
		request.CheckName = &tmp
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

	response, err := s.Client.GetAssessorCheck(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseMigrationAssessmentAssessorCheckDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseMigrationAssessmentAssessorCheckDataSource-", DatabaseMigrationAssessmentAssessorCheckDataSource(), s.D))

	if s.Res.Action != nil {
		s.D.Set("action", *s.Res.Action)
	}

	if s.Res.AssessorCheckGroup != nil {
		s.D.Set("assessor_check_group", []interface{}{AssessorCheckGroupToMap(s.Res.AssessorCheckGroup)})
	} else {
		s.D.Set("assessor_check_group", nil)
	}

	s.D.Set("assessor_check_state", s.Res.AssessorCheckState)

	if s.Res.CheckAction != nil {
		s.D.Set("check_action", []interface{}{AssessorCheckActionToMap(s.Res.CheckAction)})
	} else {
		s.D.Set("check_action", nil)
	}

	columns := []interface{}{}
	for _, item := range s.Res.Columns {
		columns = append(columns, AdvisorReportCheckColumnToMap(item))
	}
	s.D.Set("columns", columns)

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.FixupScriptLocation != nil {
		s.D.Set("fixup_script_location", *s.Res.FixupScriptLocation)
	}

	if s.Res.HelpLinkText != nil {
		s.D.Set("help_link_text", *s.Res.HelpLinkText)
	}

	if s.Res.HelpLinkUrl != nil {
		s.D.Set("help_link_url", *s.Res.HelpLinkUrl)
	}

	if s.Res.Impact != nil {
		s.D.Set("impact", *s.Res.Impact)
	}

	if s.Res.IsExclusionAllowed != nil {
		s.D.Set("is_exclusion_allowed", *s.Res.IsExclusionAllowed)
	}

	if s.Res.Issue != nil {
		s.D.Set("issue", *s.Res.Issue)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.LogLocation != nil {
		s.D.Set("log_location", []interface{}{LogLocationBucketDetailsToMap(s.Res.LogLocation)})
	} else {
		s.D.Set("log_location", nil)
	}

	if s.Res.Metadata != nil {
		s.D.Set("metadata", []interface{}{ObjectMetadataToMap(s.Res.Metadata)})
	} else {
		s.D.Set("metadata", nil)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ObjectCount != nil {
		s.D.Set("object_count", *s.Res.ObjectCount)
	}

	if s.Res.ObjectsDisplayName != nil {
		s.D.Set("objects_display_name", *s.Res.ObjectsDisplayName)
	}

	return nil
}
