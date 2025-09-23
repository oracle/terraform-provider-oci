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

func ResourceAnalyticsResourceAnalyticsInstanceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["resource_analytics_instance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ResourceAnalyticsResourceAnalyticsInstanceResource(), fieldMap, readSingularResourceAnalyticsResourceAnalyticsInstance)
}

func readSingularResourceAnalyticsResourceAnalyticsInstance(d *schema.ResourceData, m interface{}) error {
	sync := &ResourceAnalyticsResourceAnalyticsInstanceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ResourceAnalyticsInstanceClient()

	return tfresource.ReadResource(sync)
}

type ResourceAnalyticsResourceAnalyticsInstanceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_resource_analytics.ResourceAnalyticsInstanceClient
	Res    *oci_resource_analytics.GetResourceAnalyticsInstanceResponse
}

func (s *ResourceAnalyticsResourceAnalyticsInstanceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ResourceAnalyticsResourceAnalyticsInstanceDataSourceCrud) Get() error {
	request := oci_resource_analytics.GetResourceAnalyticsInstanceRequest{}

	if resourceAnalyticsInstanceId, ok := s.D.GetOkExists("resource_analytics_instance_id"); ok {
		tmp := resourceAnalyticsInstanceId.(string)
		request.ResourceAnalyticsInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "resource_analytics")

	response, err := s.Client.GetResourceAnalyticsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ResourceAnalyticsResourceAnalyticsInstanceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AdwId != nil {
		s.D.Set("adw_id", *s.Res.AdwId)
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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.OacId != nil {
		s.D.Set("oac_id", *s.Res.OacId)
	}

	s.D.Set("state", s.Res.LifecycleState)

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
