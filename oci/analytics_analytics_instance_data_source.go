// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_analytics "github.com/oracle/oci-go-sdk/analytics"
)

func init() {
	RegisterDatasource("oci_analytics_analytics_instance", AnalyticsAnalyticsInstanceDataSource())
}

func AnalyticsAnalyticsInstanceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["analytics_instance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(AnalyticsAnalyticsInstanceResource(), fieldMap, readSingularAnalyticsAnalyticsInstance)
}

func readSingularAnalyticsAnalyticsInstance(d *schema.ResourceData, m interface{}) error {
	sync := &AnalyticsAnalyticsInstanceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).analyticsClient

	return ReadResource(sync)
}

type AnalyticsAnalyticsInstanceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_analytics.AnalyticsClient
	Res    *oci_analytics.GetAnalyticsInstanceResponse
}

func (s *AnalyticsAnalyticsInstanceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AnalyticsAnalyticsInstanceDataSourceCrud) Get() error {
	request := oci_analytics.GetAnalyticsInstanceRequest{}

	if analyticsInstanceId, ok := s.D.GetOkExists("analytics_instance_id"); ok {
		tmp := analyticsInstanceId.(string)
		request.AnalyticsInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "analytics")

	response, err := s.Client.GetAnalyticsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AnalyticsAnalyticsInstanceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.Capacity != nil {
		s.D.Set("capacity", []interface{}{AnalyticsCapacityToMap(s.Res.Capacity)})
	} else {
		s.D.Set("capacity", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.EmailNotification != nil {
		s.D.Set("email_notification", *s.Res.EmailNotification)
	}

	s.D.Set("feature_set", s.Res.FeatureSet)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("license_type", s.Res.LicenseType)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ServiceUrl != nil {
		s.D.Set("service_url", *s.Res.ServiceUrl)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
