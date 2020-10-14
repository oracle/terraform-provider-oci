// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_logging "github.com/oracle/oci-go-sdk/v27/logging"
)

func init() {
	RegisterDatasource("oci_logging_log_saved_search", LoggingLogSavedSearchDataSource())
}

func LoggingLogSavedSearchDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["log_saved_search_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(LoggingLogSavedSearchResource(), fieldMap, readSingularLoggingLogSavedSearch)
}

func readSingularLoggingLogSavedSearch(d *schema.ResourceData, m interface{}) error {
	sync := &LoggingLogSavedSearchDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loggingManagementClient()

	return ReadResource(sync)
}

type LoggingLogSavedSearchDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_logging.LoggingManagementClient
	Res    *oci_logging.GetLogSavedSearchResponse
}

func (s *LoggingLogSavedSearchDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LoggingLogSavedSearchDataSourceCrud) Get() error {
	request := oci_logging.GetLogSavedSearchRequest{}

	if logSavedSearchId, ok := s.D.GetOkExists("log_saved_search_id"); ok {
		tmp := logSavedSearchId.(string)
		request.LogSavedSearchId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "logging")

	response, err := s.Client.GetLogSavedSearch(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LoggingLogSavedSearchDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Query != nil {
		s.D.Set("query", *s.Res.Query)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastModified != nil {
		s.D.Set("time_last_modified", s.Res.TimeLastModified.String())
	}

	return nil
}
