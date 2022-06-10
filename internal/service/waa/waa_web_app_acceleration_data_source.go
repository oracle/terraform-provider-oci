// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waa

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_waa "github.com/oracle/oci-go-sdk/v65/waa"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func WaaWebAppAccelerationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["web_app_acceleration_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(WaaWebAppAccelerationResource(), fieldMap, readSingularWaaWebAppAcceleration)
}

func readSingularWaaWebAppAcceleration(d *schema.ResourceData, m interface{}) error {
	sync := &WaaWebAppAccelerationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WaaClient()

	return tfresource.ReadResource(sync)
}

type WaaWebAppAccelerationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_waa.WaaClient
	Res    *oci_waa.GetWebAppAccelerationResponse
}

func (s *WaaWebAppAccelerationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WaaWebAppAccelerationDataSourceCrud) Get() error {
	request := oci_waa.GetWebAppAccelerationRequest{}

	if webAppAccelerationId, ok := s.D.GetOkExists("web_app_acceleration_id"); ok {
		tmp := webAppAccelerationId.(string)
		request.WebAppAccelerationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "waa")

	response, err := s.Client.GetWebAppAcceleration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *WaaWebAppAccelerationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.GetId())

	if s.Res.GetCompartmentId() != nil {
		s.D.Set("compartment_id", *s.Res.GetCompartmentId())
	}

	if s.Res.GetDefinedTags() != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.GetDefinedTags()))
	}

	if s.Res.GetDisplayName() != nil {
		s.D.Set("display_name", *s.Res.GetDisplayName())
	}

	s.D.Set("freeform_tags", s.Res.GetFreeformTags())

	if s.Res.GetLifecycleDetails() != nil {
		s.D.Set("lifecycle_details", *s.Res.GetLifecycleDetails())
	}

	switch v := (s.Res.WebAppAcceleration).(type) {
	case oci_waa.WebAppAccelerationLoadBalancer:
		s.D.Set("backend_type", "LOAD_BALANCER")

		if v.LoadBalancerId != nil {
			s.D.Set("load_balancer_id", *v.LoadBalancerId)
		}
	default:
		log.Printf("[WARN] Received unknown 'type': %v", s.Res.WebAppAcceleration)
		return nil
	}

	s.D.Set("state", s.Res.GetLifecycleState())

	if s.Res.GetSystemTags() != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.GetSystemTags()))
	}

	if s.Res.GetTimeCreated() != nil {
		s.D.Set("time_created", s.Res.GetTimeCreated().String())
	}

	if s.Res.GetTimeUpdated() != nil {
		s.D.Set("time_updated", s.Res.GetTimeUpdated().String())
	}

	if s.Res.GetWebAppAccelerationPolicyId() != nil {
		s.D.Set("web_app_acceleration_policy_id", *s.Res.GetWebAppAccelerationPolicyId())
	}

	return nil
}
