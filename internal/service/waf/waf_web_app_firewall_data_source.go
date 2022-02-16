// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waf

import (
	"context"
	"log"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_waf "github.com/oracle/oci-go-sdk/v58/waf"
)

func WafWebAppFirewallDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["web_app_firewall_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(WafWebAppFirewallResource(), fieldMap, readSingularWafWebAppFirewall)
}

func readSingularWafWebAppFirewall(d *schema.ResourceData, m interface{}) error {
	sync := &WafWebAppFirewallDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WafClient()

	return tfresource.ReadResource(sync)
}

type WafWebAppFirewallDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_waf.WafClient
	Res    *oci_waf.GetWebAppFirewallResponse
}

func (s *WafWebAppFirewallDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WafWebAppFirewallDataSourceCrud) Get() error {
	request := oci_waf.GetWebAppFirewallRequest{}

	if webAppFirewallId, ok := s.D.GetOkExists("web_app_firewall_id"); ok {
		tmp := webAppFirewallId.(string)
		request.WebAppFirewallId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "waf")

	response, err := s.Client.GetWebAppFirewall(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *WafWebAppFirewallDataSourceCrud) SetData() error {
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

	if s.Res.GetWebAppFirewallPolicyId() != nil {
		s.D.Set("web_app_firewall_policy_id", *s.Res.GetWebAppFirewallPolicyId())
	}

	switch v := (s.Res.WebAppFirewall).(type) {
	case oci_waf.WebAppFirewallLoadBalancer:
		s.D.Set("backend_type", "LOAD_BALANCER")

		if v.LoadBalancerId != nil {
			s.D.Set("load_balancer_id", *v.LoadBalancerId)
		}
	default:
		log.Printf("[WARN] Received unknown 'type': %v", s.Res.WebAppFirewall)
		return nil
	}

	return nil
}
