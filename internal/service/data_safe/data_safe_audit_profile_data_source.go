// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeAuditProfileDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["audit_profile_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataSafeAuditProfileResource(), fieldMap, readSingularDataSafeAuditProfile)
}

func readSingularDataSafeAuditProfile(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditProfileDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeAuditProfileDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetAuditProfileResponse
}

func (s *DataSafeAuditProfileDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeAuditProfileDataSourceCrud) Get() error {
	request := oci_data_safe.GetAuditProfileRequest{}

	if auditProfileId, ok := s.D.GetOkExists("audit_profile_id"); ok {
		tmp := auditProfileId.(string)
		request.AuditProfileId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetAuditProfile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeAuditProfileDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AuditCollectedVolume != nil {
		s.D.Set("audit_collected_volume", strconv.FormatInt(*s.Res.AuditCollectedVolume, 10))
	}

	auditTrails := []interface{}{}
	for _, item := range s.Res.AuditTrails {
		auditTrails = append(auditTrails, AuditTrailToMap(item))
	}
	s.D.Set("audit_trails", auditTrails)

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

	if s.Res.IsOverrideGlobalRetentionSetting != nil {
		s.D.Set("is_override_global_retention_setting", *s.Res.IsOverrideGlobalRetentionSetting)
	}

	if s.Res.IsPaidUsageEnabled != nil {
		s.D.Set("is_paid_usage_enabled", *s.Res.IsPaidUsageEnabled)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.OfflineMonths != nil {
		s.D.Set("offline_months", *s.Res.OfflineMonths)
	}

	if s.Res.OnlineMonths != nil {
		s.D.Set("online_months", *s.Res.OnlineMonths)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
