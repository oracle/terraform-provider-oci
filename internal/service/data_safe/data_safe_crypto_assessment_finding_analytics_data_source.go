// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeCryptoAssessmentFindingAnalyticsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDataSafeCryptoAssessmentFindingAnalyticsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"category": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"finding_key": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"is_quantum_readiness_check": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"crypto_assessment_finding_analytics_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
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
									"finding_key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"priority": {
										Type:     schema.TypeInt,
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
									"target_count": {
										Type:     schema.TypeInt,
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

func readDataSafeCryptoAssessmentFindingAnalyticsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeCryptoAssessmentFindingAnalyticsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataSafeCryptoAssessmentFindingAnalyticsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListCryptoAssessmentFindingAnalyticsResponse
}

func (s *DataSafeCryptoAssessmentFindingAnalyticsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeCryptoAssessmentFindingAnalyticsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.ListCryptoAssessmentFindingAnalyticsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListCryptoAssessmentFindingAnalyticsAccessLevelEnum(accessLevel.(string))
	}

	if category, ok := s.D.GetOkExists("category"); ok {
		request.Category = oci_data_safe.ListCryptoAssessmentFindingAnalyticsCategoryEnum(category.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if findingKey, ok := s.D.GetOkExists("finding_key"); ok {
		interfaces := findingKey.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("finding_key") {
			request.FindingKey = tmp
		}
	}

	if isQuantumReadinessCheck, ok := s.D.GetOkExists("is_quantum_readiness_check"); ok {
		tmp := isQuantumReadinessCheck.(bool)
		request.IsQuantumReadinessCheck = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListCryptoAssessmentFindingAnalytics(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCryptoAssessmentFindingAnalytics(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeCryptoAssessmentFindingAnalyticsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeCryptoAssessmentFindingAnalyticsDataSource-", DataSafeCryptoAssessmentFindingAnalyticsDataSource(), s.D))
	resources := []map[string]interface{}{}
	cryptoAssessmentFindingAnalytic := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, CryptoAssessmentFindingAnalyticsSummaryToMap(item))
	}
	cryptoAssessmentFindingAnalytic["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeCryptoAssessmentFindingAnalyticsDataSource().Schema["crypto_assessment_finding_analytics_collection"].Elem.(*schema.Resource).Schema)
		cryptoAssessmentFindingAnalytic["items"] = items
	}

	resources = append(resources, cryptoAssessmentFindingAnalytic)
	if err := s.D.Set("crypto_assessment_finding_analytics_collection", resources); err != nil {
		return err
	}

	return nil
}

func CryptoAssessmentFindingAnalyticsSummaryToMap(obj oci_data_safe.CryptoAssessmentFindingAnalyticsSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["category"] = string(obj.Category)

	if obj.FindingKey != nil {
		result["finding_key"] = string(*obj.FindingKey)
	}

	if obj.Priority != nil {
		result["priority"] = int(*obj.Priority)
	}

	result["severity"] = string(obj.Severity)

	if obj.ShortRemediation != nil {
		result["short_remediation"] = string(*obj.ShortRemediation)
	}

	if obj.ShortSummary != nil {
		result["short_summary"] = string(*obj.ShortSummary)
	}

	if obj.TargetCount != nil {
		result["target_count"] = int(*obj.TargetCount)
	}

	if obj.Title != nil {
		result["title"] = string(*obj.Title)
	}

	return result
}
