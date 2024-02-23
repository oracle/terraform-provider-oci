// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package operator_access_control

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_operator_access_control "github.com/oracle/oci-go-sdk/v65/operatoraccesscontrol"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OperatorAccessControlAccessRequestAuditLogReportDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularOperatorAccessControlAccessRequestAuditLogReport,
		Schema: map[string]*schema.Schema{
			"access_request_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"enable_process_tree": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			// Computed
			"audit_report_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"process_tree": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"report": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_of_report_generation": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularOperatorAccessControlAccessRequestAuditLogReport(d *schema.ResourceData, m interface{}) error {
	sync := &OperatorAccessControlAccessRequestAuditLogReportDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AccessRequestsClient()

	return tfresource.ReadResource(sync)
}

type OperatorAccessControlAccessRequestAuditLogReportDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_operator_access_control.AccessRequestsClient
	Res    *oci_operator_access_control.GetAuditLogReportResponse
}

func (s *OperatorAccessControlAccessRequestAuditLogReportDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OperatorAccessControlAccessRequestAuditLogReportDataSourceCrud) Get() error {
	request := oci_operator_access_control.GetAuditLogReportRequest{}

	if accessRequestId, ok := s.D.GetOkExists("access_request_id"); ok {
		tmp := accessRequestId.(string)
		request.AccessRequestId = &tmp
	}

	if enableProcessTree, ok := s.D.GetOkExists("enable_process_tree"); ok {
		tmp := enableProcessTree.(int)
		request.EnableProcessTree = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "operator_access_control")

	response, err := s.Client.GetAuditLogReport(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OperatorAccessControlAccessRequestAuditLogReportDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OperatorAccessControlAccessRequestAuditLogReportDataSource-", OperatorAccessControlAccessRequestAuditLogReportDataSource(), s.D))

	s.D.Set("audit_report_status", s.Res.AuditReportStatus)

	if s.Res.ProcessTree != nil {
		s.D.Set("process_tree", *s.Res.ProcessTree)
	}

	if s.Res.Report != nil {
		s.D.Set("report", *s.Res.Report)
	}

	if s.Res.TimeOfReportGeneration != nil {
		s.D.Set("time_of_report_generation", s.Res.TimeOfReportGeneration.String())
	}

	return nil
}
