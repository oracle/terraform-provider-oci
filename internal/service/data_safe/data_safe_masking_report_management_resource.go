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

func DataSafeMaskingReportManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeMaskingReportManagement,
		Read:     readDataSafeMaskingReportManagement,
		Update:   updateDataSafeMaskingReportManagement,
		Delete:   deleteDataSafeMaskingReportManagement,
		Schema: map[string]*schema.Schema{
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"masking_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
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
	}
}

func createDataSafeMaskingReportManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeMaskingReportManagementResourceCurd{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	err := sync.getMaskingReportWorkReq()
	if err != nil {
		return err
	}

	err1 := sync.Get()
	if err1 != nil {
		return err1
	}

	err = sync.SetData()
	if err != nil {
		return err
	}

	return nil
}

func readDataSafeMaskingReportManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeMaskingReportManagementResourceCurd{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeMaskingReportManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDataSafeMaskingReportManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DataSafeMaskingReportManagementResourceCurd struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.MaskingReport
	DisableNotFoundRetries bool
}

func (s *DataSafeMaskingReportManagementResourceCurd) ID() string {
	return *s.Res.Id
}

func (s *DataSafeMaskingReportManagementResourceCurd) getMaskingReportWorkReq() error {
	// Masking report will be in same compartment as of target
	getTargetDatabaseRequest := oci_data_safe.GetTargetDatabaseRequest{}
	var compartmentId *string

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		getTargetDatabaseRequest.TargetDatabaseId = &tmp
	}

	getTargetDatabaseResponse, err := s.Client.GetTargetDatabase(context.Background(), getTargetDatabaseRequest)
	if err != nil {
		return err
	}

	compartmentId = getTargetDatabaseResponse.CompartmentId

	err = s.D.Set("compartment_id", compartmentId)
	if err != nil {
		return err
	}

	// List all masking reports for given target and masking policy ID
	err = s.GetMaskingReportList()
	if err != nil {
		return err
	}

	// check if masking report id is set and masking report already exists
	if s.D.Id() != "" {
		return nil
	}

	// Mask target to generate masking report
	maskTargetDatabaseRequest := oci_data_safe.MaskDataRequest{}
	if maskingPolicyId, ok := s.D.GetOkExists("masking_policy_id"); ok {
		tmp := maskingPolicyId.(string)
		maskTargetDatabaseRequest.MaskingPolicyId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		maskTargetDatabaseRequest.TargetId = &tmp
	}

	maskTargetDatabaseRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.MaskData(context.Background(), maskTargetDatabaseRequest)
	if err != nil {
		return err
	}
	workId := response.OpcWorkRequestId
	_, err = maskDataWaitForWorkRequest(workId, "maskingpolicy",
		oci_data_safe.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate),
		s.DisableNotFoundRetries, s.Client)
	if err != nil {
		return err
	}

	return s.GetMaskingReportList()
}

func (s *DataSafeMaskingReportManagementResourceCurd) GetMaskingReportList() error {
	request := oci_data_safe.ListMaskingReportsRequest{}
	var maskingReport = new(oci_data_safe.MaskingReport)
	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
		request.SortOrder = oci_data_safe.ListMaskingReportsSortOrderDesc
		request.SortBy = oci_data_safe.ListMaskingReportsSortByTimemaskingfinished
	}

	if maskingPolicyId, ok := s.D.GetOkExists("masking_policy_id"); ok {
		tmp := maskingPolicyId.(string)
		request.MaskingPolicyId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ListMaskingReports(context.Background(), request)
	if err != nil {
		return err
	}
	if response.MaskingReportCollection.Items != nil && len(response.MaskingReportCollection.Items) > 0 {
		temp1 := response.MaskingReportCollection.Items[0]
		maskingReport.Id = temp1.Id
	}

	if maskingReport.Id == nil {
		return nil
	}

	s.D.SetId(*maskingReport.Id)
	return nil
}

func (s *DataSafeMaskingReportManagementResourceCurd) Get() error {
	request := oci_data_safe.GetMaskingReportRequest{}

	tmp := s.D.Id()
	request.MaskingReportId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetMaskingReport(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MaskingReport
	return nil
}

func (s *DataSafeMaskingReportManagementResourceCurd) SetData() error {

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.MaskingPolicyId != nil {
		s.D.Set("masking_policy_id", *s.Res.MaskingPolicyId)
	}

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

	if s.Res.IsDropTempTablesEnabled != nil {
		s.D.Set("is_drop_temp_tables_enabled", *s.Res.IsDropTempTablesEnabled)
	}

	if s.Res.IsRedoLoggingEnabled != nil {
		s.D.Set("is_redo_logging_enabled", *s.Res.IsRedoLoggingEnabled)
	}

	if s.Res.IsRefreshStatsEnabled != nil {
		s.D.Set("is_refresh_stats_enabled", *s.Res.IsRefreshStatsEnabled)
	}

	if s.Res.MaskingWorkRequestId != nil {
		s.D.Set("masking_work_request_id", *s.Res.MaskingWorkRequestId)
	}

	if s.Res.ParallelDegree != nil {
		s.D.Set("parallel_degree", *s.Res.ParallelDegree)
	}

	if s.Res.Recompile != nil {
		s.D.Set("recompile", *s.Res.Recompile)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeMaskingFinished != nil {
		s.D.Set("time_masking_finished", s.Res.TimeMaskingFinished.String())
	}

	if s.Res.TimeMaskingStarted != nil {
		s.D.Set("time_masking_started", s.Res.TimeMaskingStarted.String())
	}

	if s.Res.TotalMaskedColumns != nil {
		s.D.Set("total_masked_columns", strconv.FormatInt(*s.Res.TotalMaskedColumns, 10))
	}

	if s.Res.TotalMaskedObjects != nil {
		s.D.Set("total_masked_objects", strconv.FormatInt(*s.Res.TotalMaskedObjects, 10))
	}

	if s.Res.TotalMaskedSchemas != nil {
		s.D.Set("total_masked_schemas", strconv.FormatInt(*s.Res.TotalMaskedSchemas, 10))
	}

	if s.Res.TotalMaskedSensitiveTypes != nil {
		s.D.Set("total_masked_sensitive_types", strconv.FormatInt(*s.Res.TotalMaskedSensitiveTypes, 10))
	}

	if s.Res.TotalMaskedValues != nil {
		s.D.Set("total_masked_values", strconv.FormatInt(*s.Res.TotalMaskedValues, 10))
	}

	return nil
}
