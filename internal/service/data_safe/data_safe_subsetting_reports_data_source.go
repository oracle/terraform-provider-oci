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

func DataSafeSubsettingReportsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDataSafeSubsettingReportsWithContext,
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
			"subsetting_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_database_group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subsetting_report_collection": {
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
									"database_size_after_subsetting_in_kbs": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"database_size_before_subsetting_in_kbs": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
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
									"masking_report_id": {
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
									"subsetting_policy_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"subsetting_status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"subsetting_work_request_id": {
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
									"time_subsetting_finished": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_subsetting_started": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"total_post_subsetting_script_errors": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"total_pre_subsetting_script_errors": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"total_subsetted_objects": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"total_subsetted_rows": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"total_subsetted_schemas": {
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

func readDataSafeSubsettingReportsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingReportsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataSafeSubsettingReportsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListSubsettingReportsResponse
}

func (s *DataSafeSubsettingReportsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSubsettingReportsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.ListSubsettingReportsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListSubsettingReportsAccessLevelEnum(accessLevel.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if subsettingPolicyId, ok := s.D.GetOkExists("subsetting_policy_id"); ok {
		tmp := subsettingPolicyId.(string)
		request.SubsettingPolicyId = &tmp
	}

	if targetDatabaseGroupId, ok := s.D.GetOkExists("target_database_group_id"); ok {
		tmp := targetDatabaseGroupId.(string)
		request.TargetDatabaseGroupId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListSubsettingReports(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSubsettingReports(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeSubsettingReportsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSubsettingReportsDataSource-", DataSafeSubsettingReportsDataSource(), s.D))
	resources := []map[string]interface{}{}
	subsettingReport := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SubsettingReportSummaryToMap(item))
	}
	subsettingReport["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeSubsettingReportsDataSource().Schema["subsetting_report_collection"].Elem.(*schema.Resource).Schema)
		subsettingReport["items"] = items
	}

	resources = append(resources, subsettingReport)
	if err := s.D.Set("subsetting_report_collection", resources); err != nil {
		return err
	}

	return nil
}

func SubsettingReportSummaryToMap(obj oci_data_safe.SubsettingReportSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DatabaseSizeAfterSubsettingInKBs != nil {
		result["database_size_after_subsetting_in_kbs"] = string(*obj.DatabaseSizeAfterSubsettingInKBs)
	}

	if obj.DatabaseSizeBeforeSubsettingInKBs != nil {
		result["database_size_before_subsetting_in_kbs"] = string(*obj.DatabaseSizeBeforeSubsettingInKBs)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
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

	if obj.MaskingReportId != nil {
		result["masking_report_id"] = string(*obj.MaskingReportId)
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

	if obj.SubsettingPolicyId != nil {
		result["subsetting_policy_id"] = string(*obj.SubsettingPolicyId)
	}

	result["subsetting_status"] = string(obj.SubsettingStatus)

	if obj.SubsettingWorkRequestId != nil {
		result["subsetting_work_request_id"] = string(*obj.SubsettingWorkRequestId)
	}

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeSubsettingFinished != nil {
		result["time_subsetting_finished"] = obj.TimeSubsettingFinished.String()
	}

	if obj.TimeSubsettingStarted != nil {
		result["time_subsetting_started"] = obj.TimeSubsettingStarted.String()
	}

	if obj.TotalPostSubsettingScriptErrors != nil {
		result["total_post_subsetting_script_errors"] = strconv.FormatInt(*obj.TotalPostSubsettingScriptErrors, 10)
	}

	if obj.TotalPreSubsettingScriptErrors != nil {
		result["total_pre_subsetting_script_errors"] = strconv.FormatInt(*obj.TotalPreSubsettingScriptErrors, 10)
	}

	if obj.TotalSubsettedObjects != nil {
		result["total_subsetted_objects"] = strconv.FormatInt(*obj.TotalSubsettedObjects, 10)
	}

	if obj.TotalSubsettedRows != nil {
		result["total_subsetted_rows"] = strconv.FormatInt(*obj.TotalSubsettedRows, 10)
	}

	if obj.TotalSubsettedSchemas != nil {
		result["total_subsetted_schemas"] = strconv.FormatInt(*obj.TotalSubsettedSchemas, 10)
	}

	return result
}
