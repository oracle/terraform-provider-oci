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

func DataSafeAuditTrailDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["audit_trail_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataSafeAuditTrailResource(), fieldMap, readSingularDataSafeAuditTrail)
}

func readSingularDataSafeAuditTrail(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditTrailDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeAuditTrailDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetAuditTrailResponse
}

func (s *DataSafeAuditTrailDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeAuditTrailDataSourceCrud) Get() error {
	request := oci_data_safe.GetAuditTrailRequest{}

	if auditTrailId, ok := s.D.GetOkExists("audit_trail_id"); ok {
		tmp := auditTrailId.(string)
		request.AuditTrailId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetAuditTrail(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeAuditTrailDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AuditCollectionStartTime != nil {
		s.D.Set("audit_collection_start_time", s.Res.AuditCollectionStartTime.String())
	}

	if s.Res.AuditProfileId != nil {
		s.D.Set("audit_profile_id", *s.Res.AuditProfileId)
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

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsAutoPurgeEnabled != nil {
		s.D.Set("is_auto_purge_enabled", *s.Res.IsAutoPurgeEnabled)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.PeerTargetDatabaseKey != nil {
		s.D.Set("peer_target_database_key", *s.Res.PeerTargetDatabaseKey)
	}

	if s.Res.PurgeJobDetails != nil {
		s.D.Set("purge_job_details", *s.Res.PurgeJobDetails)
	}

	s.D.Set("purge_job_status", s.Res.PurgeJobStatus)

	if s.Res.PurgeJobTime != nil {
		s.D.Set("purge_job_time", s.Res.PurgeJobTime.String())
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastCollected != nil {
		s.D.Set("time_last_collected", s.Res.TimeLastCollected.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TrailLocation != nil {
		s.D.Set("trail_location", *s.Res.TrailLocation)
	}

	s.D.Set("trail_source", s.Res.TrailSource)

	if s.Res.WorkRequestId != nil {
		s.D.Set("work_request_id", *s.Res.WorkRequestId)
	}

	return nil
}
