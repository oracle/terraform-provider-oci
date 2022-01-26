// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package logging

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_logging "github.com/oracle/oci-go-sdk/v56/logging"
)

func LoggingLogDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["log_group_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["log_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(LoggingLogResource(), fieldMap, readSingularLoggingLog)
}

func readSingularLoggingLog(d *schema.ResourceData, m interface{}) error {
	sync := &LoggingLogDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoggingManagementClient()

	return tfresource.ReadResource(sync)
}

type LoggingLogDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_logging.LoggingManagementClient
	Res    *oci_logging.GetLogResponse
}

func (s *LoggingLogDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LoggingLogDataSourceCrud) Get() error {
	request := oci_logging.GetLogRequest{}

	if logGroupId, ok := s.D.GetOkExists("log_group_id"); ok {
		tmp := logGroupId.(string)
		request.LogGroupId = &tmp
	}

	if logId, ok := s.D.GetOkExists("log_id"); ok {
		tmp := logId.(string)
		request.LogId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "logging")

	response, err := s.Client.GetLog(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LoggingLogDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Configuration != nil {
		s.D.Set("configuration", []interface{}{LogConfigurationToMap(s.Res.Configuration)})
	} else {
		s.D.Set("configuration", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	s.D.Set("log_type", s.Res.LogType)

	if s.Res.RetentionDuration != nil {
		s.D.Set("retention_duration", *s.Res.RetentionDuration)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TenancyId != nil {
		s.D.Set("tenancy_id", *s.Res.TenancyId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastModified != nil {
		s.D.Set("time_last_modified", s.Res.TimeLastModified.String())
	}

	return nil
}
