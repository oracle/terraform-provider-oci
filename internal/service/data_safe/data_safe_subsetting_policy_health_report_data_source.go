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

func DataSafeSubsettingPolicyHealthReportDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readSingularDataSafeSubsettingPolicyHealthReportWithContext,
		Schema: map[string]*schema.Schema{
			"subsetting_policy_health_report_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"error_count": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subsetting_policy_id": {
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"warning_count": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDataSafeSubsettingPolicyHealthReportWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingPolicyHealthReportDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataSafeSubsettingPolicyHealthReportDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetSubsettingPolicyHealthReportResponse
}

func (s *DataSafeSubsettingPolicyHealthReportDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSubsettingPolicyHealthReportDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.GetSubsettingPolicyHealthReportRequest{}

	if subsettingPolicyHealthReportId, ok := s.D.GetOkExists("subsetting_policy_health_report_id"); ok {
		tmp := subsettingPolicyHealthReportId.(string)
		request.SubsettingPolicyHealthReportId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetSubsettingPolicyHealthReport(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeSubsettingPolicyHealthReportDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ErrorCount != nil {
		s.D.Set("error_count", strconv.FormatInt(*s.Res.ErrorCount, 10))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubsettingPolicyId != nil {
		s.D.Set("subsetting_policy_id", *s.Res.SubsettingPolicyId)
	}

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.WarningCount != nil {
		s.D.Set("warning_count", strconv.FormatInt(*s.Res.WarningCount, 10))
	}

	return nil
}
