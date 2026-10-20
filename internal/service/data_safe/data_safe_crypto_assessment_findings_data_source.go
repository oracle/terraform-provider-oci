// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeCryptoAssessmentFindingsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDataSafeCryptoAssessmentFindingsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"category": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"crypto_assessment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"finding_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_quantum_readiness_check": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"title": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"crypto_assessment_finding_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"assessment_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"database_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"category": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compliance": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"expected_value": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"finding_key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_quantum_readiness_check": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"observed_value": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"priority": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"recommended_value": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"remediation": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"short_remediation": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"short_summary": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"summary": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"title": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"url": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"summary": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"backup_status": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"findings": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"pass_checks": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"total_checks": {
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"critical": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"data_encryption_status": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"findings": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"pass_checks": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"total_checks": {
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"high": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"low": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"med": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"network_encryption_status": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"findings": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"pass_checks": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"total_checks": {
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"status_counts": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"total_checks": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"total_findings": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"wallet_status": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"findings": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"pass_checks": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"total_checks": {
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"target_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDataSafeCryptoAssessmentFindingsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeCryptoAssessmentFindingsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataSafeCryptoAssessmentFindingsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListCryptoAssessmentFindingsResponse
}

func (s *DataSafeCryptoAssessmentFindingsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeCryptoAssessmentFindingsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.ListCryptoAssessmentFindingsRequest{}

	if category, ok := s.D.GetOkExists("category"); ok {
		request.Category = oci_data_safe.ListCryptoAssessmentFindingsCategoryEnum(category.(string))
	}

	if cryptoAssessmentId, ok := s.D.GetOkExists("crypto_assessment_id"); ok {
		tmp := cryptoAssessmentId.(string)
		request.CryptoAssessmentId = &tmp
	}

	if findingKey, ok := s.D.GetOkExists("finding_key"); ok {
		tmp := findingKey.(string)
		request.FindingKey = &tmp
	}

	if isQuantumReadinessCheck, ok := s.D.GetOkExists("is_quantum_readiness_check"); ok {
		tmp := isQuantumReadinessCheck.(bool)
		request.IsQuantumReadinessCheck = &tmp
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_data_safe.ListCryptoAssessmentFindingsStatusEnum(status.(string))
	}

	if title, ok := s.D.GetOkExists("title"); ok {
		tmp := title.(string)
		request.Title = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListCryptoAssessmentFindings(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCryptoAssessmentFindings(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeCryptoAssessmentFindingsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeCryptoAssessmentFindingsDataSource-", DataSafeCryptoAssessmentFindingsDataSource(), s.D))
	resources := []map[string]interface{}{}
	cryptoAssessmentFinding := map[string]interface{}{}

	cryptoAssessmentFinding["assessment_type"] = s.Res.AssessmentType

	if s.Res.DatabaseVersion != nil {
		cryptoAssessmentFinding["database_version"] = *s.Res.DatabaseVersion
	}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CryptoAssessmentFindingSummaryToMap(item))
	}
	cryptoAssessmentFinding["items"] = items

	if s.Res.Summary != nil {
		cryptoAssessmentFinding["summary"] = []interface{}{CryptoAssessmentFindingSummaryMetricsToMap(s.Res.Summary)}
	} else {
		cryptoAssessmentFinding["summary"] = nil
	}

	if s.Res.TargetId != nil {
		cryptoAssessmentFinding["target_id"] = *s.Res.TargetId
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeCryptoAssessmentFindingsDataSource().Schema["crypto_assessment_finding_collection"].Elem.(*schema.Resource).Schema)
		cryptoAssessmentFinding["items"] = items
	}

	resources = append(resources, cryptoAssessmentFinding)
	if err := s.D.Set("crypto_assessment_finding_collection", resources); err != nil {
		return err
	}

	return nil
}

func CryptoAssessmentFindingCategorySummaryToMap(obj *oci_data_safe.CryptoAssessmentFindingCategorySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Findings != nil {
		result["findings"] = int(*obj.Findings)
	}

	if obj.PassChecks != nil {
		result["pass_checks"] = int(*obj.PassChecks)
	}

	if obj.TotalChecks != nil {
		result["total_checks"] = int(*obj.TotalChecks)
	}

	return result
}

func CryptoAssessmentFindingSummaryToMap(obj oci_data_safe.CryptoAssessmentFindingSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Category != nil {
		result["category"] = string(*obj.Category)
	}

	if obj.Compliance != nil {
		result["compliance"] = string(*obj.Compliance)
	}

	if obj.ExpectedValue != nil {
		result["expected_value"] = string(*obj.ExpectedValue)
	}

	if obj.FindingKey != nil {
		result["finding_key"] = string(*obj.FindingKey)
	}

	if obj.IsQuantumReadinessCheck != nil {
		result["is_quantum_readiness_check"] = bool(*obj.IsQuantumReadinessCheck)
	}

	result["observed_value"] = obj.ObservedValue

	if obj.Priority != nil {
		result["priority"] = int(*obj.Priority)
	}

	if obj.RecommendedValue != nil {
		result["recommended_value"] = string(*obj.RecommendedValue)
	}

	if obj.Remediation != nil {
		result["remediation"] = string(*obj.Remediation)
	}

	result["severity"] = string(obj.Severity)

	if obj.ShortRemediation != nil {
		result["short_remediation"] = string(*obj.ShortRemediation)
	}

	if obj.ShortSummary != nil {
		result["short_summary"] = string(*obj.ShortSummary)
	}

	result["status"] = string(obj.Status)

	if obj.Summary != nil {
		result["summary"] = string(*obj.Summary)
	}

	if obj.Title != nil {
		result["title"] = string(*obj.Title)
	}

	if obj.Url != nil {
		result["url"] = string(*obj.Url)
	}

	return result
}

func CryptoAssessmentFindingSummaryMetricsToMap(obj *oci_data_safe.CryptoAssessmentFindingSummaryMetrics) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BackupStatus != nil {
		result["backup_status"] = []interface{}{CryptoAssessmentFindingCategorySummaryToMap(obj.BackupStatus)}
	}

	if obj.Critical != nil {
		result["critical"] = int(*obj.Critical)
	}

	if obj.DataEncryptionStatus != nil {
		result["data_encryption_status"] = []interface{}{CryptoAssessmentFindingCategorySummaryToMap(obj.DataEncryptionStatus)}
	}

	if obj.High != nil {
		result["high"] = int(*obj.High)
	}

	if obj.Low != nil {
		result["low"] = int(*obj.Low)
	}

	if obj.Med != nil {
		result["med"] = int(*obj.Med)
	}

	if obj.NetworkEncryptionStatus != nil {
		result["network_encryption_status"] = []interface{}{CryptoAssessmentFindingCategorySummaryToMap(obj.NetworkEncryptionStatus)}
	}

	statusCounts := map[string]string{}
	for key, value := range obj.StatusCounts {
		statusCounts[key] = strconv.Itoa(value)
	}
	result["status_counts"] = statusCounts

	if obj.TotalChecks != nil {
		result["total_checks"] = int(*obj.TotalChecks)
	}

	if obj.TotalFindings != nil {
		result["total_findings"] = int(*obj.TotalFindings)
	}

	if obj.WalletStatus != nil {
		result["wallet_status"] = []interface{}{CryptoAssessmentFindingCategorySummaryToMap(obj.WalletStatus)}
	}

	return result
}
