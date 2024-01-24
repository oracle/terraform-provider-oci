// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms_java_downloads

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_jms_java_downloads "github.com/oracle/oci-go-sdk/v65/jmsjavadownloads"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsJavaDownloadsJavaDownloadReportDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["java_download_report_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(JmsJavaDownloadsJavaDownloadReportResource(), fieldMap, readSingularJmsJavaDownloadsJavaDownloadReport)
}

func readSingularJmsJavaDownloadsJavaDownloadReport(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJavaDownloadsJavaDownloadReportDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaDownloadClient()

	return tfresource.ReadResource(sync)
}

type JmsJavaDownloadsJavaDownloadReportDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms_java_downloads.JavaDownloadClient
	Res    *oci_jms_java_downloads.GetJavaDownloadReportResponse
}

func (s *JmsJavaDownloadsJavaDownloadReportDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsJavaDownloadsJavaDownloadReportDataSourceCrud) Get() error {
	request := oci_jms_java_downloads.GetJavaDownloadReportRequest{}

	if javaDownloadReportId, ok := s.D.GetOkExists("java_download_report_id"); ok {
		tmp := javaDownloadReportId.(string)
		request.JavaDownloadReportId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms_java_downloads")

	response, err := s.Client.GetJavaDownloadReport(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *JmsJavaDownloadsJavaDownloadReportDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("checksum_type", s.Res.ChecksumType)

	if s.Res.ChecksumValue != nil {
		s.D.Set("checksum_value", *s.Res.ChecksumValue)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", []interface{}{PrincipalToMap(s.Res.CreatedBy)})
	} else {
		s.D.Set("created_by", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.FileSizeInBytes != nil {
		s.D.Set("file_size_in_bytes", strconv.FormatInt(*s.Res.FileSizeInBytes, 10))
	}

	s.D.Set("format", s.Res.Format)

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
