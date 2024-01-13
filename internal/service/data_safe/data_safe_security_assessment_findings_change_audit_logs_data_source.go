// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSecurityAssessmentFindingsChangeAuditLogsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeSecurityAssessmentFindingsChangeAuditLogs,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"finding_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"finding_title": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_risk_deferred": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"modified_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_assessment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"severity": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_updated_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_updated_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_valid_until_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_valid_until_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"findings_change_audit_log_collection": {
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
									"assessment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"finding_key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"finding_title": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_risk_deferred": {
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
									"modified_by": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"oracle_defined_severity": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"previous_severity": {
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
									"time_updated": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_valid_until": {
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

func readDataSafeSecurityAssessmentFindingsChangeAuditLogs(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeSecurityAssessmentFindingsChangeAuditLogsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeSecurityAssessmentFindingsChangeAuditLogsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListFindingsChangeAuditLogsResponse
}

func (s *DataSafeSecurityAssessmentFindingsChangeAuditLogsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSecurityAssessmentFindingsChangeAuditLogsDataSourceCrud) Get() error {
	request := oci_data_safe.ListFindingsChangeAuditLogsRequest{}

	if findingKey, ok := s.D.GetOkExists("finding_key"); ok {
		tmp := findingKey.(string)
		request.FindingKey = &tmp
	}

	if findingTitle, ok := s.D.GetOkExists("finding_title"); ok {
		tmp := findingTitle.(string)
		request.FindingTitle = &tmp
	}

	if isRiskDeferred, ok := s.D.GetOkExists("is_risk_deferred"); ok {
		tmp := isRiskDeferred.(bool)
		request.IsRiskDeferred = &tmp
	}

	if modifiedBy, ok := s.D.GetOkExists("modified_by"); ok {
		tmp := modifiedBy.(string)
		request.ModifiedBy = &tmp
	}

	if securityAssessmentId, ok := s.D.GetOkExists("security_assessment_id"); ok {
		tmp := securityAssessmentId.(string)
		request.SecurityAssessmentId = &tmp
	}

	if severity, ok := s.D.GetOkExists("severity"); ok {
		request.Severity = oci_data_safe.ListFindingsChangeAuditLogsSeverityEnum(severity.(string))
	}

	if timeUpdatedGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_updated_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeUpdatedGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeUpdatedGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeUpdatedLessThan, ok := s.D.GetOkExists("time_updated_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timeUpdatedLessThan.(string))
		if err != nil {
			return err
		}
		request.TimeUpdatedLessThan = &oci_common.SDKTime{Time: tmp}
	}

	if timeValidUntilGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_valid_until_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeValidUntilGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeValidUntilGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeValidUntilLessThan, ok := s.D.GetOkExists("time_valid_until_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timeValidUntilLessThan.(string))
		if err != nil {
			return err
		}
		request.TimeValidUntilLessThan = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListFindingsChangeAuditLogs(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListFindingsChangeAuditLogs(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSecurityAssessmentFindingsChangeAuditLogsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSecurityAssessmentFindingsChangeAuditLogsDataSource-", DataSafeSecurityAssessmentFindingsChangeAuditLogsDataSource(), s.D))
	resources := []map[string]interface{}{}
	securityAssessmentFindingsChangeAuditLog := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, FindingsChangeAuditLogSummaryToMap(item))
	}
	securityAssessmentFindingsChangeAuditLog["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSecurityAssessmentFindingsChangeAuditLogsDataSource().Schema["findings_change_audit_log_collection"].Elem.(*schema.Resource).Schema)
		securityAssessmentFindingsChangeAuditLog["items"] = items
	}

	resources = append(resources, securityAssessmentFindingsChangeAuditLog)
	if err := s.D.Set("findings_change_audit_log_collection", resources); err != nil {
		return err
	}

	return nil
}

func FindingsChangeAuditLogSummaryToMap(obj oci_data_safe.FindingsChangeAuditLogSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AssessmentId != nil {
		result["assessment_id"] = string(*obj.AssessmentId)
	}

	if obj.FindingKey != nil {
		result["finding_key"] = string(*obj.FindingKey)
	}

	if obj.FindingTitle != nil {
		result["finding_title"] = string(*obj.FindingTitle)
	}

	if obj.IsRiskDeferred != nil {
		result["is_risk_deferred"] = bool(*obj.IsRiskDeferred)
	}

	if obj.Justification != nil {
		result["justification"] = string(*obj.Justification)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.ModifiedBy != nil {
		result["modified_by"] = string(*obj.ModifiedBy)
	}

	result["oracle_defined_severity"] = string(obj.OracleDefinedSeverity)

	result["previous_severity"] = string(obj.PreviousSeverity)

	result["severity"] = string(obj.Severity)

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.TimeValidUntil != nil {
		result["time_valid_until"] = obj.TimeValidUntil.String()
	}

	return result
}
