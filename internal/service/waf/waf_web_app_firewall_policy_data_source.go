// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waf

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_waf "github.com/oracle/oci-go-sdk/v56/waf"
)

func WafWebAppFirewallPolicyDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["web_app_firewall_policy_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(WafWebAppFirewallPolicyResource(), fieldMap, readSingularWafWebAppFirewallPolicy)
}

func readSingularWafWebAppFirewallPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &WafWebAppFirewallPolicyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WafClient()

	return tfresource.ReadResource(sync)
}

type WafWebAppFirewallPolicyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_waf.WafClient
	Res    *oci_waf.GetWebAppFirewallPolicyResponse
}

func (s *WafWebAppFirewallPolicyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WafWebAppFirewallPolicyDataSourceCrud) Get() error {
	request := oci_waf.GetWebAppFirewallPolicyRequest{}

	if webAppFirewallPolicyId, ok := s.D.GetOkExists("web_app_firewall_policy_id"); ok {
		tmp := webAppFirewallPolicyId.(string)
		request.WebAppFirewallPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "waf")

	response, err := s.Client.GetWebAppFirewallPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *WafWebAppFirewallPolicyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	actions := []interface{}{}
	for _, item := range s.Res.Actions {
		actions = append(actions, WafActionToMap(item))
	}
	s.D.Set("actions", actions)

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

	if s.Res.RequestAccessControl != nil {
		s.D.Set("request_access_control", []interface{}{RequestAccessControlToMap(s.Res.RequestAccessControl)})
	} else {
		s.D.Set("request_access_control", nil)
	}

	if s.Res.RequestProtection != nil {
		s.D.Set("request_protection", []interface{}{RequestProtectionToMap(s.Res.RequestProtection)})
	} else {
		s.D.Set("request_protection", nil)
	}

	if s.Res.RequestRateLimiting != nil {
		s.D.Set("request_rate_limiting", []interface{}{RequestRateLimitingToMap(s.Res.RequestRateLimiting)})
	} else {
		s.D.Set("request_rate_limiting", nil)
	}

	if s.Res.ResponseAccessControl != nil {
		s.D.Set("response_access_control", []interface{}{ResponseAccessControlToMap(s.Res.ResponseAccessControl)})
	} else {
		s.D.Set("response_access_control", nil)
	}

	if s.Res.ResponseProtection != nil {
		s.D.Set("response_protection", []interface{}{ResponseProtectionToMap(s.Res.ResponseProtection)})
	} else {
		s.D.Set("response_protection", nil)
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
