// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
)

func DataSafeSecurityAssessmentComparisonDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDataSafeSecurityAssessmentComparison,
		Schema: map[string]*schema.Schema{
			"comparison_security_assessment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"security_assessment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"baseline_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"targets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"auditing": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"added_items": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
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
									"current": {
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
									"modified_items": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"removed_items": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"severity": {
										Type:     schema.TypeString,
										Computed: true,
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
									"added_items": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
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
									"current": {
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
									"modified_items": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"removed_items": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"baseline_target_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"current_target_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"data_encryption": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"added_items": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
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
									"current": {
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
									"modified_items": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"removed_items": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"severity": {
										Type:     schema.TypeString,
										Computed: true,
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
									"added_items": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
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
									"current": {
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
									"modified_items": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"removed_items": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"severity": {
										Type:     schema.TypeString,
										Computed: true,
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
									"added_items": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
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
									"current": {
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
									"modified_items": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"removed_items": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"severity": {
										Type:     schema.TypeString,
										Computed: true,
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
									"added_items": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
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
									"current": {
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
									"modified_items": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"removed_items": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"user_accounts": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"added_items": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
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
									"current": {
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
									"modified_items": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"removed_items": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDataSafeSecurityAssessmentComparison(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityAssessmentComparisonDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSecurityAssessmentComparisonDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetSecurityAssessmentComparisonResponse
}

func (s *DataSafeSecurityAssessmentComparisonDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSecurityAssessmentComparisonDataSourceCrud) Get() error {
	request := oci_data_safe.GetSecurityAssessmentComparisonRequest{}

	if comparisonSecurityAssessmentId, ok := s.D.GetOkExists("comparison_security_assessment_id"); ok {
		tmp := comparisonSecurityAssessmentId.(string)
		request.ComparisonSecurityAssessmentId = &tmp
	}

	if securityAssessmentId, ok := s.D.GetOkExists("security_assessment_id"); ok {
		tmp := securityAssessmentId.(string)
		request.SecurityAssessmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetSecurityAssessmentComparison(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeSecurityAssessmentComparisonDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.BaselineId != nil {
		s.D.Set("baseline_id", *s.Res.BaselineId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	targets := []interface{}{}
	for _, item := range s.Res.Targets {
		targets = append(targets, SecurityAssessmentComparisonPerTargetToMap(item))
	}
	s.D.Set("targets", targets)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func DiffsToMap(obj oci_data_safe.Diffs) map[string]interface{} {
	result := map[string]interface{}{}

	result["added_items"] = obj.AddedItems

	if obj.Baseline != nil {
		result["baseline"] = []interface{}{FindingToMap(obj.Baseline)}
	}

	if obj.Current != nil {
		result["current"] = []interface{}{FindingToMap(obj.Current)}
	}

	result["modified_items"] = obj.ModifiedItems

	result["removed_items"] = obj.RemovedItems

	result["severity"] = string(obj.Severity)

	return result
}

func FindingToMap(obj *oci_data_safe.Finding) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AssessmentId != nil {
		result["assessment_id"] = string(*obj.AssessmentId)
	}

	if obj.Details != nil {
		result["details"] = []interface{}{}
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
		result["references"] = []interface{}{ReferencesToMap(obj.References)}
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

func ReferencesToMap(obj *oci_data_safe.References) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Cis != nil {
		result["cis"] = string(*obj.Cis)
	}

	if obj.Gdpr != nil {
		result["gdpr"] = string(*obj.Gdpr)
	}

	if obj.Stig != nil {
		result["stig"] = string(*obj.Stig)
	}

	return result
}

func SecurityAssessmentComparisonPerTargetToMap(obj oci_data_safe.SecurityAssessmentComparisonPerTarget) map[string]interface{} {
	result := map[string]interface{}{}

	auditing := []interface{}{}
	for _, item := range obj.Auditing {
		auditing = append(auditing, DiffsToMap(item))
	}
	result["auditing"] = auditing

	authorizationControl := []interface{}{}
	for _, item := range obj.AuthorizationControl {
		authorizationControl = append(authorizationControl, DiffsToMap(item))
	}
	result["authorization_control"] = authorizationControl

	if obj.BaselineTargetId != nil {
		result["baseline_target_id"] = string(*obj.BaselineTargetId)
	}

	if obj.CurrentTargetId != nil {
		result["current_target_id"] = string(*obj.CurrentTargetId)
	}

	dataEncryption := []interface{}{}
	for _, item := range obj.DataEncryption {
		dataEncryption = append(dataEncryption, DiffsToMap(item))
	}
	result["data_encryption"] = dataEncryption

	dbConfiguration := []interface{}{}
	for _, item := range obj.DbConfiguration {
		dbConfiguration = append(dbConfiguration, DiffsToMap(item))
	}
	result["db_configuration"] = dbConfiguration

	fineGrainedAccessControl := []interface{}{}
	for _, item := range obj.FineGrainedAccessControl {
		fineGrainedAccessControl = append(fineGrainedAccessControl, DiffsToMap(item))
	}
	result["fine_grained_access_control"] = fineGrainedAccessControl

	privilegesAndRoles := []interface{}{}
	for _, item := range obj.PrivilegesAndRoles {
		privilegesAndRoles = append(privilegesAndRoles, DiffsToMap(item))
	}
	result["privileges_and_roles"] = privilegesAndRoles

	userAccounts := []interface{}{}
	for _, item := range obj.UserAccounts {
		userAccounts = append(userAccounts, DiffsToMap(item))
	}
	result["user_accounts"] = userAccounts

	return result
}
