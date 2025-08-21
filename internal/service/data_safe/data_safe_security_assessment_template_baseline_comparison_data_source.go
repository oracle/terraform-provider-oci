// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSecurityAssessmentTemplateBaselineComparisonDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDataSafeSecurityAssessmentTemplateBaselineComparison,
		Schema: map[string]*schema.Schema{
			"category": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"comparison_security_assessment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"finding_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_assessment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"auditing": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"baseline": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"assessment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"details": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"has_target_db_risk_level_changed": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_risk_modified": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"justification": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"oracle_defined_severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"references": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"cis": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"gdpr": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"obp": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"stig": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"remarks": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"summary": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_updated": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_valid_until": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"title": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"targets": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"title": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"authorization_control": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"baseline": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"assessment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"details": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"has_target_db_risk_level_changed": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_risk_modified": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"justification": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"oracle_defined_severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"references": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"cis": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"gdpr": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"obp": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"stig": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"remarks": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"summary": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_updated": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_valid_until": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"title": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"targets": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"title": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"data_encryption": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"baseline": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"assessment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"details": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"has_target_db_risk_level_changed": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_risk_modified": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"justification": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"oracle_defined_severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"references": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"cis": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"gdpr": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"obp": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"stig": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"remarks": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"summary": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_updated": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_valid_until": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"title": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"targets": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"title": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"db_configuration": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"baseline": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"assessment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"details": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"has_target_db_risk_level_changed": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_risk_modified": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"justification": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"oracle_defined_severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"references": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"cis": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"gdpr": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"obp": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"stig": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"remarks": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"summary": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_updated": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_valid_until": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"title": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"targets": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"title": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"fine_grained_access_control": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"baseline": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"assessment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"details": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"has_target_db_risk_level_changed": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_risk_modified": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"justification": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"oracle_defined_severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"references": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"cis": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"gdpr": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"obp": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"stig": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"remarks": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"summary": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_updated": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_valid_until": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"title": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"targets": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"title": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"privileges_and_roles": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"baseline": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"assessment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"details": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"has_target_db_risk_level_changed": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_risk_modified": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"justification": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"oracle_defined_severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"references": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"cis": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"gdpr": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"obp": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"stig": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"remarks": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"summary": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_updated": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_valid_until": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"title": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"targets": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"title": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"template_baseline_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"template_baseline_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_accounts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"baseline": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"assessment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"details": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"has_target_db_risk_level_changed": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_risk_modified": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"justification": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"oracle_defined_severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"references": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"cis": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"gdpr": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"obp": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"stig": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"remarks": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"summary": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_updated": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_valid_until": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"title": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"targets": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"title": {
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

func readSingularDataSafeSecurityAssessmentTemplateBaselineComparison(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityAssessmentTemplateBaselineComparisonDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSecurityAssessmentTemplateBaselineComparisonDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetTemplateBaselineComparisonResponse
}

func (s *DataSafeSecurityAssessmentTemplateBaselineComparisonDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSecurityAssessmentTemplateBaselineComparisonDataSourceCrud) Get() error {
	request := oci_data_safe.GetTemplateBaselineComparisonRequest{}

	if category, ok := s.D.GetOkExists("category"); ok {
		tmp := category.(string)
		request.Category = &tmp
	}

	if comparisonSecurityAssessmentId, ok := s.D.GetOkExists("comparison_security_assessment_id"); ok {
		tmp := comparisonSecurityAssessmentId.(string)
		request.ComparisonSecurityAssessmentId = &tmp
	}

	if findingKey, ok := s.D.GetOkExists("finding_key"); ok {
		tmp := findingKey.(string)
		request.FindingKey = &tmp
	}

	if securityAssessmentId, ok := s.D.GetOkExists("security_assessment_id"); ok {
		tmp := securityAssessmentId.(string)
		request.SecurityAssessmentId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetTemplateBaselineComparison(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeSecurityAssessmentTemplateBaselineComparisonDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	auditing := []interface{}{}
	for _, item := range s.Res.Auditing {
		auditing = append(auditing, TemplateBaselineDiffsToMap(item))
	}
	s.D.Set("auditing", auditing)

	authorizationControl := []interface{}{}
	for _, item := range s.Res.AuthorizationControl {
		authorizationControl = append(authorizationControl, TemplateBaselineDiffsToMap(item))
	}
	s.D.Set("authorization_control", authorizationControl)

	dataEncryption := []interface{}{}
	for _, item := range s.Res.DataEncryption {
		dataEncryption = append(dataEncryption, TemplateBaselineDiffsToMap(item))
	}
	s.D.Set("data_encryption", dataEncryption)

	dbConfiguration := []interface{}{}
	for _, item := range s.Res.DbConfiguration {
		dbConfiguration = append(dbConfiguration, TemplateBaselineDiffsToMap(item))
	}
	s.D.Set("db_configuration", dbConfiguration)

	fineGrainedAccessControl := []interface{}{}
	for _, item := range s.Res.FineGrainedAccessControl {
		fineGrainedAccessControl = append(fineGrainedAccessControl, TemplateBaselineDiffsToMap(item))
	}
	s.D.Set("fine_grained_access_control", fineGrainedAccessControl)

	privilegesAndRoles := []interface{}{}
	for _, item := range s.Res.PrivilegesAndRoles {
		privilegesAndRoles = append(privilegesAndRoles, TemplateBaselineDiffsToMap(item))
	}
	s.D.Set("privileges_and_roles", privilegesAndRoles)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TemplateBaselineId != nil {
		s.D.Set("template_baseline_id", *s.Res.TemplateBaselineId)
	}

	if s.Res.TemplateBaselineName != nil {
		s.D.Set("template_baseline_name", *s.Res.TemplateBaselineName)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	userAccounts := []interface{}{}
	for _, item := range s.Res.UserAccounts {
		userAccounts = append(userAccounts, TemplateBaselineDiffsToMap(item))
	}
	s.D.Set("user_accounts", userAccounts)

	return nil
}

func FindingToMap2(obj *oci_data_safe.Finding) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AssessmentId != nil {
		result["assessment_id"] = string(*obj.AssessmentId)
	}

	if obj.Details != nil {
		result["details"] = []interface{}{
			objectToMap(
				(*obj.Details).(map[string]interface{}),
			),
		}
	}

	if obj.HasTargetDbRiskLevelChanged != nil {
		result["has_target_db_risk_level_changed"] = bool(*obj.HasTargetDbRiskLevelChanged)
	}

	if obj.IsRiskModified != nil {
		result["is_risk_modified"] = bool(*obj.IsRiskModified)
	}

	if obj.Justification != nil {
		result["justification"] = string(*obj.Justification)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["oracle_defined_severity"] = string(obj.OracleDefinedSeverity)

	if obj.References != nil {
		result["references"] = []interface{}{ReferencesToMap2(obj.References)}
	}

	if obj.Remarks != nil {
		result["remarks"] = string(*obj.Remarks)
	}

	result["severity"] = string(obj.Severity)

	result["state"] = string(obj.LifecycleState)

	if obj.Summary != nil {
		result["summary"] = string(*obj.Summary)
	}

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.TimeValidUntil != nil {
		result["time_valid_until"] = obj.TimeValidUntil.String()
	}

	if obj.Title != nil {
		result["title"] = string(*obj.Title)
	}

	return result
}

func ReferencesToMap2(obj *oci_data_safe.References) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Cis != nil {
		result["cis"] = string(*obj.Cis)
	}

	if obj.Gdpr != nil {
		result["gdpr"] = string(*obj.Gdpr)
	}

	if obj.Obp != nil {
		result["obp"] = string(*obj.Obp)
	}

	if obj.Stig != nil {
		result["stig"] = string(*obj.Stig)
	}

	return result
}

func TemplateBaselineDiffsToMap(obj oci_data_safe.TemplateBaselineDiffs) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Baseline != nil {
		result["baseline"] = []interface{}{FindingToMap2(obj.Baseline)}
	}

	targets := []interface{}{}
	for _, item := range obj.Targets {
		targets = append(targets, TemplateBaselineDiffsPerTargetToMap(item))
	}
	result["targets"] = targets

	return result
}

func TemplateBaselineDiffsPerTargetToMap(obj oci_data_safe.TemplateBaselineDiffsPerTarget) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	result["severity"] = string(obj.Severity)

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	if obj.Title != nil {
		result["title"] = string(*obj.Title)
	}

	return result
}
