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

func DataSafeMaskingPolicyHealthReportLogsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeMaskingPolicyHealthReportLogs,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"masking_policy_health_report_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"message_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"masking_policy_health_report_log_collection": {
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

func readDataSafeMaskingPolicyHealthReportLogs(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeMaskingPolicyHealthReportLogsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeMaskingPolicyHealthReportLogsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListMaskingPolicyHealthReportLogsResponse
}

func (s *DataSafeMaskingPolicyHealthReportLogsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeMaskingPolicyHealthReportLogsDataSourceCrud) Get() error {
	request := oci_data_safe.ListMaskingPolicyHealthReportLogsRequest{}

	if maskingPolicyHealthReportId, ok := s.D.GetOkExists("masking_policy_health_report_id"); ok {
		tmp := maskingPolicyHealthReportId.(string)
		request.MaskingPolicyHealthReportId = &tmp
	}

	if messageType, ok := s.D.GetOkExists("message_type"); ok {
		request.MessageType = oci_data_safe.ListMaskingPolicyHealthReportLogsMessageTypeEnum(messageType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListMaskingPolicyHealthReportLogs(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMaskingPolicyHealthReportLogs(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeMaskingPolicyHealthReportLogsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeMaskingPolicyHealthReportLogsDataSource-", DataSafeMaskingPolicyHealthReportLogsDataSource(), s.D))
	resources := []map[string]interface{}{}
	maskingPolicyHealthReportLog := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MaskingPolicyHealthReportLogSummaryToMap(item))
	}
	maskingPolicyHealthReportLog["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeMaskingPolicyHealthReportLogsDataSource().Schema["masking_policy_health_report_log_collection"].Elem.(*schema.Resource).Schema)
		maskingPolicyHealthReportLog["items"] = items
	}

	resources = append(resources, maskingPolicyHealthReportLog)
	if err := s.D.Set("masking_policy_health_report_log_collection", resources); err != nil {
		return err
	}

	return nil
}

func MaskingPolicyHealthReportLogSummaryToMap(obj oci_data_safe.MaskingPolicyHealthReportLogSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

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
