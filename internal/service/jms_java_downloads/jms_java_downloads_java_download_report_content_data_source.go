// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms_java_downloads

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_jms_java_downloads "github.com/oracle/oci-go-sdk/v65/jmsjavadownloads"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsJavaDownloadsJavaDownloadReportContentDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularJmsJavaDownloadsJavaDownloadReportContent,
		Schema: map[string]*schema.Schema{
			"java_download_report_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
		},
	}
}

func readSingularJmsJavaDownloadsJavaDownloadReportContent(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJavaDownloadsJavaDownloadReportContentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaDownloadClient()

	return tfresource.ReadResource(sync)
}

type JmsJavaDownloadsJavaDownloadReportContentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms_java_downloads.JavaDownloadClient
	Res    *oci_jms_java_downloads.GetJavaDownloadReportContentResponse
}

func (s *JmsJavaDownloadsJavaDownloadReportContentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsJavaDownloadsJavaDownloadReportContentDataSourceCrud) Get() error {
	request := oci_jms_java_downloads.GetJavaDownloadReportContentRequest{}

	if javaDownloadReportId, ok := s.D.GetOkExists("java_download_report_id"); ok {
		tmp := javaDownloadReportId.(string)
		request.JavaDownloadReportId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms_java_downloads")

	response, err := s.Client.GetJavaDownloadReportContent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *JmsJavaDownloadsJavaDownloadReportContentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsJavaDownloadsJavaDownloadReportContentDataSource-", JmsJavaDownloadsJavaDownloadReportContentDataSource(), s.D))

	return nil
}
