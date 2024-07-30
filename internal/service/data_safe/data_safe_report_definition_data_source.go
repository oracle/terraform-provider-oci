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

func DataSafeReportDefinitionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["report_definition_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataSafeReportDefinitionResource(), fieldMap, readSingularDataSafeReportDefinition)
}

func readSingularDataSafeReportDefinition(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeReportDefinitionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeReportDefinitionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetReportDefinitionResponse
}

func (s *DataSafeReportDefinitionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeReportDefinitionDataSourceCrud) Get() error {
	request := oci_data_safe.GetReportDefinitionRequest{}

	if reportDefinitionId, ok := s.D.GetOkExists("report_definition_id"); ok {
		tmp := reportDefinitionId.(string)
		request.ReportDefinitionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetReportDefinition(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeReportDefinitionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("category", s.Res.Category)

	columnFilters := []interface{}{}
	for _, item := range s.Res.ColumnFilters {
		columnFilters = append(columnFilters, columnFilterToMap(item))
	}
	s.D.Set("column_filters", columnFilters)

	columnInfo := []interface{}{}
	for _, item := range s.Res.ColumnInfo {
		columnInfo = append(columnInfo, columnToMap(item))
	}
	s.D.Set("column_info", columnInfo)

	columnSortings := []interface{}{}
	for _, item := range s.Res.ColumnSortings {
		columnSortings = append(columnSortings, columnSortingToMap(item))
	}
	s.D.Set("column_sortings", columnSortings)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("compliance_standards", s.Res.ComplianceStandards)
	s.D.Set("compliance_standards", s.Res.ComplianceStandards)

	s.D.Set("data_source", s.Res.DataSource)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DisplayOrder != nil {
		s.D.Set("display_order", *s.Res.DisplayOrder)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsSeeded != nil {
		s.D.Set("is_seeded", *s.Res.IsSeeded)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ParentId != nil {
		s.D.Set("parent_id", *s.Res.ParentId)
	}

	if s.Res.RecordTimeSpan != nil {
		s.D.Set("record_time_span", *s.Res.RecordTimeSpan)
	}

	if s.Res.Schedule != nil {
		s.D.Set("schedule", *s.Res.Schedule)
	}

	if s.Res.ScheduledReportCompartmentId != nil {
		s.D.Set("scheduled_report_compartment_id", *s.Res.ScheduledReportCompartmentId)
	}

	s.D.Set("scheduled_report_mime_type", s.Res.ScheduledReportMimeType)

	if s.Res.ScheduledReportName != nil {
		s.D.Set("scheduled_report_name", *s.Res.ScheduledReportName)
	}

	if s.Res.ScheduledReportRowLimit != nil {
		s.D.Set("scheduled_report_row_limit", *s.Res.ScheduledReportRowLimit)
	}

	if s.Res.ScimFilter != nil {
		s.D.Set("scim_filter", *s.Res.ScimFilter)
	}

	s.D.Set("state", s.Res.LifecycleState)

	summary := []interface{}{}
	for _, item := range s.Res.Summary {
		summary = append(summary, summaryToMap(item))
	}
	s.D.Set("summary", summary)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
