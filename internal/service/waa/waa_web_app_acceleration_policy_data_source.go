// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waa

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_waa "github.com/oracle/oci-go-sdk/v65/waa"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func WaaWebAppAccelerationPolicyDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["web_app_acceleration_policy_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(WaaWebAppAccelerationPolicyResource(), fieldMap, readSingularWaaWebAppAccelerationPolicy)
}

func readSingularWaaWebAppAccelerationPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &WaaWebAppAccelerationPolicyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WaaClient()

	return tfresource.ReadResource(sync)
}

type WaaWebAppAccelerationPolicyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_waa.WaaClient
	Res    *oci_waa.GetWebAppAccelerationPolicyResponse
}

func (s *WaaWebAppAccelerationPolicyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WaaWebAppAccelerationPolicyDataSourceCrud) Get() error {
	request := oci_waa.GetWebAppAccelerationPolicyRequest{}

	if webAppAccelerationPolicyId, ok := s.D.GetOkExists("web_app_acceleration_policy_id"); ok {
		tmp := webAppAccelerationPolicyId.(string)
		request.WebAppAccelerationPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "waa")

	response, err := s.Client.GetWebAppAccelerationPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *WaaWebAppAccelerationPolicyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ResponseCachingPolicy != nil {
		s.D.Set("response_caching_policy", []interface{}{ResponseCachingPolicyToMap(s.Res.ResponseCachingPolicy)})
	} else {
		s.D.Set("response_caching_policy", nil)
	}

	if s.Res.ResponseCompressionPolicy != nil {
		s.D.Set("response_compression_policy", []interface{}{ResponseCompressionPolicyToMap(s.Res.ResponseCompressionPolicy)})
	} else {
		s.D.Set("response_compression_policy", nil)
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
