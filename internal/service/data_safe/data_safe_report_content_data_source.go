// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeReportContentDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDataSafeReportContent,
		Schema: map[string]*schema.Schema{
			"report_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
		},
	}
}

func readSingularDataSafeReportContent(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeReportContentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeReportContentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetReportContentResponse
}

func (s *DataSafeReportContentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeReportContentDataSourceCrud) Get() error {
	request := oci_data_safe.GetReportContentRequest{}

	if reportId, ok := s.D.GetOkExists("report_id"); ok {
		tmp := reportId.(string)
		request.ReportId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetReportContent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeReportContentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeReportContentDataSource-", DataSafeReportContentDataSource(), s.D))

	return nil
}
