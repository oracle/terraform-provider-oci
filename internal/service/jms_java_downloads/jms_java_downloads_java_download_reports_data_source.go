// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms_java_downloads

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_jms_java_downloads "github.com/oracle/oci-go-sdk/v65/jmsjavadownloads"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsJavaDownloadsJavaDownloadReportsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readJmsJavaDownloadsJavaDownloadReports,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"java_download_report_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"java_download_report_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(JmsJavaDownloadsJavaDownloadReportResource()),
						},
					},
				},
			},
		},
	}
}

func readJmsJavaDownloadsJavaDownloadReports(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJavaDownloadsJavaDownloadReportsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaDownloadClient()

	return tfresource.ReadResource(sync)
}

type JmsJavaDownloadsJavaDownloadReportsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms_java_downloads.JavaDownloadClient
	Res    *oci_jms_java_downloads.ListJavaDownloadReportsResponse
}

func (s *JmsJavaDownloadsJavaDownloadReportsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsJavaDownloadsJavaDownloadReportsDataSourceCrud) Get() error {
	request := oci_jms_java_downloads.ListJavaDownloadReportsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if javaDownloadReportId, ok := s.D.GetOkExists("id"); ok {
		tmp := javaDownloadReportId.(string)
		request.JavaDownloadReportId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_jms_java_downloads.ListJavaDownloadReportsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms_java_downloads")

	response, err := s.Client.ListJavaDownloadReports(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListJavaDownloadReports(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *JmsJavaDownloadsJavaDownloadReportsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsJavaDownloadsJavaDownloadReportsDataSource-", JmsJavaDownloadsJavaDownloadReportsDataSource(), s.D))
	resources := []map[string]interface{}{}
	javaDownloadReport := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, JavaDownloadReportSummaryToMap(item))
	}
	javaDownloadReport["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, JmsJavaDownloadsJavaDownloadReportsDataSource().Schema["java_download_report_collection"].Elem.(*schema.Resource).Schema)
		javaDownloadReport["items"] = items
	}

	resources = append(resources, javaDownloadReport)
	if err := s.D.Set("java_download_report_collection", resources); err != nil {
		return err
	}

	return nil
}
