// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSubsettingReportDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readSingularDataSafeSubsettingReportWithContext,
		Schema: map[string]*schema.Schema{
			"subsetting_report_id": {
				Type:     schema.TypeString,
				Required: true,
			},
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
			"masking_status": {
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
	}
}

func readSingularDataSafeSubsettingReportWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingReportDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataSafeSubsettingReportDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetSubsettingReportResponse
}

func (s *DataSafeSubsettingReportDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSubsettingReportDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.GetSubsettingReportRequest{}

	if subsettingReportId, ok := s.D.GetOkExists("subsetting_report_id"); ok {
		tmp := subsettingReportId.(string)
		request.SubsettingReportId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetSubsettingReport(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeSubsettingReportDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DatabaseSizeAfterSubsettingInKBs != nil {
		s.D.Set("database_size_after_subsetting_in_kbs", *s.Res.DatabaseSizeAfterSubsettingInKBs)
	}

	if s.Res.DatabaseSizeBeforeSubsettingInKBs != nil {
		s.D.Set("database_size_before_subsetting_in_kbs", *s.Res.DatabaseSizeBeforeSubsettingInKBs)
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

	if s.Res.MaskingReportId != nil {
		s.D.Set("masking_report_id", *s.Res.MaskingReportId)
	}

	if s.Res.MaskingWorkRequestId != nil {
		s.D.Set("masking_work_request_id", *s.Res.MaskingWorkRequestId)
	}

	if maskingStatus, err := subsettingReportMaskingStatus(s.Res); err != nil {
		return err
	} else if maskingStatus != "" {
		s.D.Set("masking_status", maskingStatus)
	}

	if s.Res.ParallelDegree != nil {
		s.D.Set("parallel_degree", *s.Res.ParallelDegree)
	}

	if s.Res.Recompile != nil {
		s.D.Set("recompile", *s.Res.Recompile)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubsettingPolicyId != nil {
		s.D.Set("subsetting_policy_id", *s.Res.SubsettingPolicyId)
	}

	s.D.Set("subsetting_status", s.Res.SubsettingStatus)

	if s.Res.SubsettingWorkRequestId != nil {
		s.D.Set("subsetting_work_request_id", *s.Res.SubsettingWorkRequestId)
	}

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeSubsettingFinished != nil {
		s.D.Set("time_subsetting_finished", s.Res.TimeSubsettingFinished.String())
	}

	if s.Res.TimeSubsettingStarted != nil {
		s.D.Set("time_subsetting_started", s.Res.TimeSubsettingStarted.String())
	}

	if s.Res.TotalPostSubsettingScriptErrors != nil {
		s.D.Set("total_post_subsetting_script_errors", strconv.FormatInt(*s.Res.TotalPostSubsettingScriptErrors, 10))
	}

	if s.Res.TotalPreSubsettingScriptErrors != nil {
		s.D.Set("total_pre_subsetting_script_errors", strconv.FormatInt(*s.Res.TotalPreSubsettingScriptErrors, 10))
	}

	if s.Res.TotalSubsettedObjects != nil {
		s.D.Set("total_subsetted_objects", strconv.FormatInt(*s.Res.TotalSubsettedObjects, 10))
	}

	if s.Res.TotalSubsettedRows != nil {
		s.D.Set("total_subsetted_rows", strconv.FormatInt(*s.Res.TotalSubsettedRows, 10))
	}

	if s.Res.TotalSubsettedSchemas != nil {
		s.D.Set("total_subsetted_schemas", strconv.FormatInt(*s.Res.TotalSubsettedSchemas, 10))
	}

	return nil
}

// MaskingStatus is present in the service response but is not yet modeled by
// the vendored OCI SDK SubsettingReport type. Read it from the preserved raw
// response body until the SDK model includes the field.
func subsettingReportMaskingStatus(response *oci_data_safe.GetSubsettingReportResponse) (string, error) {
	if response == nil || response.RawResponse == nil || response.RawResponse.Body == nil {
		return "", nil
	}

	body, err := io.ReadAll(response.RawResponse.Body)
	if err != nil {
		return "", err
	}
	response.RawResponse.Body = io.NopCloser(bytes.NewReader(body))

	var report struct {
		MaskingStatus string `json:"maskingStatus"`
	}
	if err := json.Unmarshal(body, &report); err != nil {
		return "", err
	}
	return report.MaskingStatus, nil
}
