// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeAuditArchiveRetrievalDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["audit_archive_retrieval_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataSafeAuditArchiveRetrievalResource(), fieldMap, readSingularDataSafeAuditArchiveRetrieval)
}

func readSingularDataSafeAuditArchiveRetrieval(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditArchiveRetrievalDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeAuditArchiveRetrievalDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetAuditArchiveRetrievalResponse
}

func (s *DataSafeAuditArchiveRetrievalDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeAuditArchiveRetrievalDataSourceCrud) Get() error {
	request := oci_data_safe.GetAuditArchiveRetrievalRequest{}

	if auditArchiveRetrievalId, ok := s.D.GetOkExists("audit_archive_retrieval_id"); ok {
		tmp := auditArchiveRetrievalId.(string)
		request.AuditArchiveRetrievalId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetAuditArchiveRetrieval(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeAuditArchiveRetrievalDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AuditEventCount != nil {
		s.D.Set("audit_event_count", strconv.FormatInt(*s.Res.AuditEventCount, 10))
	}

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

	if s.Res.EndDate != nil {
		s.D.Set("end_date", s.Res.EndDate.Format(time.RFC3339Nano))
	}

	if s.Res.ErrorInfo != nil {
		s.D.Set("error_info", *s.Res.ErrorInfo)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.StartDate != nil {
		s.D.Set("start_date", s.Res.StartDate.Format(time.RFC3339Nano))
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

	if s.Res.TimeCompleted != nil {
		s.D.Set("time_completed", s.Res.TimeCompleted.String())
	}

	if s.Res.TimeOfExpiry != nil {
		s.D.Set("time_of_expiry", s.Res.TimeOfExpiry.String())
	}

	if s.Res.TimeRequested != nil {
		s.D.Set("time_requested", s.Res.TimeRequested.String())
	}

	return nil
}
