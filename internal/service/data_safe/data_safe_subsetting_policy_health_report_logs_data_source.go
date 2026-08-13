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

func DataSafeSubsettingPolicyHealthReportLogsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDataSafeSubsettingPolicyHealthReportLogsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"message_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subsetting_policy_health_report_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subsetting_policy_health_report_log_collection": {
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
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"health_check_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"message": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"message_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"remediation": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"timestamp": {
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

func readDataSafeSubsettingPolicyHealthReportLogsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingPolicyHealthReportLogsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataSafeSubsettingPolicyHealthReportLogsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListSubsettingPolicyHealthReportLogsResponse
}

func (s *DataSafeSubsettingPolicyHealthReportLogsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSubsettingPolicyHealthReportLogsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.ListSubsettingPolicyHealthReportLogsRequest{}

	if messageType, ok := s.D.GetOkExists("message_type"); ok {
		request.MessageType = oci_data_safe.ListSubsettingPolicyHealthReportLogsMessageTypeEnum(messageType.(string))
	}

	if subsettingPolicyHealthReportId, ok := s.D.GetOkExists("subsetting_policy_health_report_id"); ok {
		tmp := subsettingPolicyHealthReportId.(string)
		request.SubsettingPolicyHealthReportId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListSubsettingPolicyHealthReportLogs(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSubsettingPolicyHealthReportLogs(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSubsettingPolicyHealthReportLogsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSubsettingPolicyHealthReportLogsDataSource-", DataSafeSubsettingPolicyHealthReportLogsDataSource(), s.D))
	resources := []map[string]interface{}{}
	subsettingPolicyHealthReportLog := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SubsettingPolicyHealthReportLogSummaryToMap(item))
	}
	subsettingPolicyHealthReportLog["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSubsettingPolicyHealthReportLogsDataSource().Schema["subsetting_policy_health_report_log_collection"].Elem.(*schema.Resource).Schema)
		subsettingPolicyHealthReportLog["items"] = items
	}

	resources = append(resources, subsettingPolicyHealthReportLog)
	if err := s.D.Set("subsetting_policy_health_report_log_collection", resources); err != nil {
		return err
	}

	return nil
}

func SubsettingPolicyHealthReportLogSummaryToMap(obj oci_data_safe.SubsettingPolicyHealthReportLogSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["health_check_type"] = string(obj.HealthCheckType)

	if obj.Message != nil {
		result["message"] = string(*obj.Message)
	}

	result["message_type"] = string(obj.MessageType)

	if obj.Remediation != nil {
		result["remediation"] = string(*obj.Remediation)
	}

	if obj.Timestamp != nil {
		result["timestamp"] = obj.Timestamp.String()
	}

	return result
}
