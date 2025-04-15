// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LogAnalyticsNamespaceLookupDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["lookup_name"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["namespace"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(LogAnalyticsNamespaceLookupResource(), fieldMap, readSingularLogAnalyticsNamespaceLookup)
}

func readSingularLogAnalyticsNamespaceLookup(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceLookupDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

type LogAnalyticsNamespaceLookupDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.GetLookupResponse
}

func (s *LogAnalyticsNamespaceLookupDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsNamespaceLookupDataSourceCrud) Get() error {
	request := oci_log_analytics.GetLookupRequest{}

	if lookupName, ok := s.D.GetOkExists("lookup_name"); ok {
		tmp := lookupName.(string)
		request.LookupName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "log_analytics")

	response, err := s.Client.GetLookup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LogAnalyticsNamespaceLookupDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ActiveEditVersion != nil {
		s.D.Set("active_edit_version", strconv.FormatInt(*s.Res.ActiveEditVersion, 10))
	}

	if s.Res.CanonicalLink != nil {
		s.D.Set("canonical_link", *s.Res.CanonicalLink)
	}

	categories := []interface{}{}
	for _, item := range s.Res.Categories {
		categories = append(categories, LogAnalyticsCategoryToMap(item))
	}
	s.D.Set("categories", categories)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.EditVersion != nil {
		s.D.Set("edit_version", strconv.FormatInt(*s.Res.EditVersion, 10))
	}

	fields := []interface{}{}
	for _, item := range s.Res.Fields {
		fields = append(fields, LookupFieldToMap(item))
	}
	s.D.Set("fields", fields)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsBuiltIn != nil {
		s.D.Set("is_built_in", strconv.FormatInt(*s.Res.IsBuiltIn, 10))
	}

	if s.Res.IsHidden != nil {
		s.D.Set("is_hidden", *s.Res.IsHidden)
	}

	if s.Res.Id != nil {
		s.D.Set("lookup_id", *s.Res.Id)
	}

	if s.Res.LookupDisplayName != nil {
		s.D.Set("lookup_display_name", *s.Res.LookupDisplayName)
	}

	if s.Res.LookupReference != nil {
		s.D.Set("lookup_reference", strconv.FormatInt(*s.Res.LookupReference, 10))
	}

	if s.Res.LookupReferenceString != nil {
		s.D.Set("lookup_reference_string", *s.Res.LookupReferenceString)
	}

	if s.Res.ReferringSources != nil {
		s.D.Set("referring_sources", []interface{}{AutoLookupsToMap(s.Res.ReferringSources)})
	} else {
		s.D.Set("referring_sources", nil)
	}

	if s.Res.StatusSummary.Filename != nil {
		s.D.Set("register_lookup_file", *s.Res.StatusSummary.Filename)
	}

	if s.Res.StatusSummary != nil {
		s.D.Set("status_summary", []interface{}{StatusSummaryToMap(s.Res.StatusSummary)})
	} else {
		s.D.Set("status_summary", nil)
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("type", s.Res.Type)

	return nil
}
