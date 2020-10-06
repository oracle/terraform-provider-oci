// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_logging "github.com/oracle/oci-go-sdk/v26/logging"
)

func init() {
	RegisterDatasource("oci_logging_log_group", LoggingLogGroupDataSource())
}

func LoggingLogGroupDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["log_group_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(LoggingLogGroupResource(), fieldMap, readSingularLoggingLogGroup)
}

func readSingularLoggingLogGroup(d *schema.ResourceData, m interface{}) error {
	sync := &LoggingLogGroupDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loggingManagementClient()

	return ReadResource(sync)
}

type LoggingLogGroupDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_logging.LoggingManagementClient
	Res    *oci_logging.GetLogGroupResponse
}

func (s *LoggingLogGroupDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LoggingLogGroupDataSourceCrud) Get() error {
	request := oci_logging.GetLogGroupRequest{}

	if logGroupId, ok := s.D.GetOkExists("log_group_id"); ok {
		tmp := logGroupId.(string)
		request.LogGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "logging")

	response, err := s.Client.GetLogGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LoggingLogGroupDataSourceCrud) SetData() error {
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

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastModified != nil {
		s.D.Set("time_last_modified", s.Res.TimeLastModified.String())
	}

	return nil
}
