// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeMaskingReportsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeMaskingReports,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"access_level": {
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
			"masking_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"masking_report_collection": {
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
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_drop_temp_tables_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_redo_logging_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_refresh_stats_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"masking_policy_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"masking_work_request_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"parallel_degree": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"recompile": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_masking_finished": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_masking_started": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"total_masked_columns": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"total_masked_objects": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"total_masked_schemas": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"total_masked_sensitive_types": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"total_masked_values": {
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

func readDataSafeMaskingReports(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeMaskingReportsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeMaskingReportsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListMaskingReportsResponse
}

func (s *DataSafeMaskingReportsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeMaskingReportsDataSourceCrud) Get() error {
	request := oci_data_safe.ListMaskingReportsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListMaskingReportsAccessLevelEnum(accessLevel.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if maskingPolicyId, ok := s.D.GetOkExists("masking_policy_id"); ok {
		tmp := maskingPolicyId.(string)
		request.MaskingPolicyId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListMaskingReports(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMaskingReports(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeMaskingReportsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeMaskingReportsDataSource-", DataSafeMaskingReportsDataSource(), s.D))
	resources := []map[string]interface{}{}
	maskingReport := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MaskingReportSummaryToMap(item))
	}
	maskingReport["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeMaskingReportsDataSource().Schema["masking_report_collection"].Elem.(*schema.Resource).Schema)
		maskingReport["items"] = items
	}

	resources = append(resources, maskingReport)
	if err := s.D.Set("masking_report_collection", resources); err != nil {
		return err
	}

	return nil
}

func DataSafeMaskingReportSummaryToMap(obj oci_data_safe.MaskingReportSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.MaskingPolicyId != nil {
		result["masking_policy_id"] = string(*obj.MaskingPolicyId)
	}

	if obj.MaskingWorkRequestId != nil {
		result["masking_work_request_id"] = string(*obj.MaskingWorkRequestId)
	}

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	if obj.TimeMaskingFinished != nil {
		result["time_masking_finished"] = obj.TimeMaskingFinished.String()
	}

	if obj.TimeMaskingStarted != nil {
		result["time_masking_started"] = obj.TimeMaskingStarted.String()
	}

	if obj.TotalMaskedColumns != nil {
		result["total_masked_columns"] = strconv.FormatInt(*obj.TotalMaskedColumns, 10)
	}

	if obj.TotalMaskedObjects != nil {
		result["total_masked_objects"] = strconv.FormatInt(*obj.TotalMaskedObjects, 10)
	}

	if obj.TotalMaskedSchemas != nil {
		result["total_masked_schemas"] = strconv.FormatInt(*obj.TotalMaskedSchemas, 10)
	}

	if obj.TotalMaskedSensitiveTypes != nil {
		result["total_masked_sensitive_types"] = strconv.FormatInt(*obj.TotalMaskedSensitiveTypes, 10)
	}

	if obj.TotalMaskedValues != nil {
		result["total_masked_values"] = strconv.FormatInt(*obj.TotalMaskedValues, 10)
	}

	return result
}

func MaskingReportSummaryToMap(obj oci_data_safe.MaskingReportSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsDropTempTablesEnabled != nil {
		result["is_drop_temp_tables_enabled"] = bool(*obj.IsDropTempTablesEnabled)
	}

	if obj.IsRedoLoggingEnabled != nil {
		result["is_redo_logging_enabled"] = bool(*obj.IsRedoLoggingEnabled)
	}

	if obj.IsRefreshStatsEnabled != nil {
		result["is_refresh_stats_enabled"] = bool(*obj.IsRefreshStatsEnabled)
	}

	if obj.MaskingPolicyId != nil {
		result["masking_policy_id"] = string(*obj.MaskingPolicyId)
	}

	if obj.MaskingWorkRequestId != nil {
		result["masking_work_request_id"] = string(*obj.MaskingWorkRequestId)
	}

	if obj.ParallelDegree != nil {
		result["parallel_degree"] = string(*obj.ParallelDegree)
	}

	if obj.Recompile != nil {
		result["recompile"] = string(*obj.Recompile)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeMaskingFinished != nil {
		result["time_masking_finished"] = obj.TimeMaskingFinished.String()
	}

	if obj.TimeMaskingStarted != nil {
		result["time_masking_started"] = obj.TimeMaskingStarted.String()
	}

	if obj.TotalMaskedColumns != nil {
		result["total_masked_columns"] = strconv.FormatInt(*obj.TotalMaskedColumns, 10)
	}

	if obj.TotalMaskedObjects != nil {
		result["total_masked_objects"] = strconv.FormatInt(*obj.TotalMaskedObjects, 10)
	}

	if obj.TotalMaskedSchemas != nil {
		result["total_masked_schemas"] = strconv.FormatInt(*obj.TotalMaskedSchemas, 10)
	}

	if obj.TotalMaskedSensitiveTypes != nil {
		result["total_masked_sensitive_types"] = strconv.FormatInt(*obj.TotalMaskedSensitiveTypes, 10)
	}

	if obj.TotalMaskedValues != nil {
		result["total_masked_values"] = strconv.FormatInt(*obj.TotalMaskedValues, 10)
	}

	return result
}
