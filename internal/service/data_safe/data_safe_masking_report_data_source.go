// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v61/datasafe"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
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
			"masking_policy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"masking_work_request_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"target_id": {
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

	if s.Res.MaskingPolicyId != nil {
		s.D.Set("masking_policy_id", *s.Res.MaskingPolicyId)
	}

	if s.Res.MaskingWorkRequestId != nil {
		s.D.Set("masking_work_request_id", *s.Res.MaskingWorkRequestId)
	}

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
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

func MaskingReportSummaryToMap(obj oci_data_safe.MaskingReportSummary) map[string]interface{} {
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
