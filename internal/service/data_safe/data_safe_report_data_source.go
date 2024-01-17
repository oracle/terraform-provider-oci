// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeReportDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["report_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataSafeReportResource(), fieldMap, readSingularDataSafeReport)
}

func readSingularDataSafeReport(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeReportDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeReportDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetReportResponse
}

func (s *DataSafeReportDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeReportDataSourceCrud) Get() error {
	request := oci_data_safe.GetReportRequest{}

	if reportId, ok := s.D.GetOkExists("report_id"); ok {
		tmp := reportId.(string)
		request.ReportId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetReport(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeReportDataSourceCrud) SetData() error {
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

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("mime_type", s.Res.MimeType)

	if s.Res.ReportDefinitionId != nil {
		s.D.Set("report_definition_id", *s.Res.ReportDefinitionId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeGenerated != nil {
		s.D.Set("time_generated", s.Res.TimeGenerated.String())
	}

	s.D.Set("type", s.Res.Type)

	return nil
}

/*
func ReportSummaryToMap(obj oci_data_safe.ReportSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = definedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["mime_type"] = string(obj.MimeType)

	if obj.ReportDefinitionId != nil {
		result["report_definition_id"] = string(*obj.ReportDefinitionId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeGenerated != nil {
		result["time_generated"] = obj.TimeGenerated.String()
	}

	return result
}
*/
