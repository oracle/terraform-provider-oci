// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resource_analytics

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_resource_analytics "github.com/oracle/oci-go-sdk/v65/resourceanalytics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ResourceAnalyticsTenancyAttachmentDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["tenancy_attachment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ResourceAnalyticsTenancyAttachmentResource(), fieldMap, readSingularResourceAnalyticsTenancyAttachment)
}

func readSingularResourceAnalyticsTenancyAttachment(d *schema.ResourceData, m interface{}) error {
	sync := &ResourceAnalyticsTenancyAttachmentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).TenancyAttachmentClient()

	return tfresource.ReadResource(sync)
}

type ResourceAnalyticsTenancyAttachmentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_resource_analytics.TenancyAttachmentClient
	Res    *oci_resource_analytics.GetTenancyAttachmentResponse
}

func (s *ResourceAnalyticsTenancyAttachmentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ResourceAnalyticsTenancyAttachmentDataSourceCrud) Get() error {
	request := oci_resource_analytics.GetTenancyAttachmentRequest{}

	if tenancyAttachmentId, ok := s.D.GetOkExists("tenancy_attachment_id"); ok {
		tmp := tenancyAttachmentId.(string)
		request.TenancyAttachmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "resource_analytics")

	response, err := s.Client.GetTenancyAttachment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ResourceAnalyticsTenancyAttachmentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.IsReportingTenancy != nil {
		s.D.Set("is_reporting_tenancy", *s.Res.IsReportingTenancy)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ResourceAnalyticsInstanceId != nil {
		s.D.Set("resource_analytics_instance_id", *s.Res.ResourceAnalyticsInstanceId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TenancyId != nil {
		s.D.Set("tenancy_id", *s.Res.TenancyId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
