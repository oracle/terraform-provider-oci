// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waf

import (
	"context"
	"log"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_waf "github.com/oracle/oci-go-sdk/v65/waf"
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
	switch v := (s.Res.WebAppFirewall).(type) {
	case oci_waf.WebAppFirewallLoadBalancer:
		s.D.Set("backend_type", "LOAD_BALANCER")

		if v.LoadBalancerId != nil {
			s.D.Set("load_balancer_id", *v.LoadBalancerId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.WebAppFirewallPolicyId != nil {
			s.D.Set("web_app_firewall_policy_id", *v.WebAppFirewallPolicyId)
		}
	default:
		log.Printf("[WARN] Received 'backend_type' of unknown type %v", s.Res.WebAppFirewall)
		return nil
	}

	return nil
}
