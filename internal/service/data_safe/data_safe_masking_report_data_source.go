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

func DataSafeMaskingReportDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDataSafeMaskingReport,
		Schema: map[string]*schema.Schema{
			"masking_report_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
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
	}
}

func readSingularDataSafeMaskingReport(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeMaskingReportDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeMaskingReportDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetMaskingReportResponse
}

func (s *DataSafeMaskingReportDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeMaskingReportDataSourceCrud) Get() error {
	request := oci_data_safe.GetMaskingReportRequest{}

	if maskingReportId, ok := s.D.GetOkExists("masking_report_id"); ok {
		tmp := maskingReportId.(string)
		request.MaskingReportId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetMaskingReport(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeMaskingReportDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
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

	if s.Res.MaskingPolicyId != nil {
		s.D.Set("masking_policy_id", *s.Res.MaskingPolicyId)
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

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

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
