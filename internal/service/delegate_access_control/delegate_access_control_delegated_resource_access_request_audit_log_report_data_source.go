// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package delegate_access_control

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_delegate_access_control "github.com/oracle/oci-go-sdk/v65/delegateaccesscontrol"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DelegateAccessControlDelegatedResourceAccessRequestAuditLogReportDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDelegateAccessControlDelegatedResourceAccessRequestAuditLogReport,
		Schema: map[string]*schema.Schema{
			"delegated_resource_access_request_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_process_tree_enabled": {
				Type:     schema.TypeBool,
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
			"time_report_generated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDelegateAccessControlDelegatedResourceAccessRequestAuditLogReport(d *schema.ResourceData, m interface{}) error {
	sync := &DelegateAccessControlDelegatedResourceAccessRequestAuditLogReportDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DelegateAccessControlClient()

	return tfresource.ReadResource(sync)
}

type DelegateAccessControlDelegatedResourceAccessRequestAuditLogReportDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_delegate_access_control.DelegateAccessControlClient
	Res    *oci_delegate_access_control.GetDelegatedResourceAccessRequestAuditLogReportResponse
}

func (s *DelegateAccessControlDelegatedResourceAccessRequestAuditLogReportDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DelegateAccessControlDelegatedResourceAccessRequestAuditLogReportDataSourceCrud) Get() error {
	request := oci_delegate_access_control.GetDelegatedResourceAccessRequestAuditLogReportRequest{}

	if delegatedResourceAccessRequestId, ok := s.D.GetOkExists("delegated_resource_access_request_id"); ok {
		tmp := delegatedResourceAccessRequestId.(string)
		request.DelegatedResourceAccessRequestId = &tmp
	}

	if isProcessTreeEnabled, ok := s.D.GetOkExists("is_process_tree_enabled"); ok {
		tmp := isProcessTreeEnabled.(bool)
		request.IsProcessTreeEnabled = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "delegate_access_control")

	response, err := s.Client.GetDelegatedResourceAccessRequestAuditLogReport(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DelegateAccessControlDelegatedResourceAccessRequestAuditLogReportDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DelegateAccessControlDelegatedResourceAccessRequestAuditLogReportDataSource-", DelegateAccessControlDelegatedResourceAccessRequestAuditLogReportDataSource(), s.D))

	s.D.Set("audit_report_status", s.Res.AuditReportStatus)

	if s.Res.ProcessTree != nil {
		s.D.Set("process_tree", *s.Res.ProcessTree)
	}

	if s.Res.Report != nil {
		s.D.Set("report", *s.Res.Report)
	}

	if s.Res.TimeReportGenerated != nil {
		s.D.Set("time_report_generated", s.Res.TimeReportGenerated.String())
	}

	return nil
}
