// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osmanagement

import (
	"context"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_osmanagement "github.com/oracle/oci-go-sdk/v65/osmanagement"
)

func OsmanagementManagedInstanceEventReportDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularOsmanagementManagedInstanceEventReport,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"latest_timestamp_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"latest_timestamp_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"counts": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func readSingularOsmanagementManagedInstanceEventReport(d *schema.ResourceData, m interface{}) error {
	sync := &OsmanagementManagedInstanceEventReportDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EventClient()

	return tfresource.ReadResource(sync)
}

type OsmanagementManagedInstanceEventReportDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_osmanagement.EventClient
	Res    *oci_osmanagement.GetEventReportResponse
}

func (s *OsmanagementManagedInstanceEventReportDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsmanagementManagedInstanceEventReportDataSourceCrud) Get() error {
	request := oci_osmanagement.GetEventReportRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if latestTimestampGreaterThanOrEqualTo, ok := s.D.GetOkExists("latest_timestamp_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, latestTimestampGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.LatestTimestampGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if latestTimestampLessThan, ok := s.D.GetOkExists("latest_timestamp_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, latestTimestampLessThan.(string))
		if err != nil {
			return err
		}
		request.LatestTimestampLessThan = &oci_common.SDKTime{Time: tmp}
	}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "osmanagement")

	response, err := s.Client.GetEventReport(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsmanagementManagedInstanceEventReportDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsmanagementManagedInstanceEventReportDataSource-", OsmanagementManagedInstanceEventReportDataSource(), s.D))

	if s.Res.Count != nil {
		s.D.Set("counts", *s.Res.Count)
	}

	return nil
}
