// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waas

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_waas "github.com/oracle/oci-go-sdk/v58/waas"
)

func WaasProtectionRulesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readWaasProtectionRules,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"action": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"mod_security_rule_id": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"waas_policy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"protection_rules": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     WaasProtectionRuleResource(),
			},
		},
	}
}

func readWaasProtectionRules(d *schema.ResourceData, m interface{}) error {
	sync := &WaasProtectionRulesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WaasClient()

	return tfresource.ReadResource(sync)
}

type WaasProtectionRulesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_waas.WaasClient
	Res    *oci_waas.ListProtectionRulesResponse
}

func (s *WaasProtectionRulesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WaasProtectionRulesDataSourceCrud) Get() error {
	request := oci_waas.ListProtectionRulesRequest{}

	if action, ok := s.D.GetOkExists("action"); ok {
		interfaces := action.([]interface{})
		tmp := make([]oci_waas.ListProtectionRulesActionEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_waas.ListProtectionRulesActionEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("action") {
			request.Action = tmp
		}
	}

	if modSecurityRuleId, ok := s.D.GetOkExists("mod_security_rule_id"); ok {
		interfaces := modSecurityRuleId.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("mod_security_rule_id") {
			request.ModSecurityRuleId = tmp
		}
	}

	if waasPolicyId, ok := s.D.GetOkExists("waas_policy_id"); ok {
		tmp := waasPolicyId.(string)
		request.WaasPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "waas")

	response, err := s.Client.ListProtectionRules(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListProtectionRules(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *WaasProtectionRulesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("WaasProtectionRulesDataSource-", WaasProtectionRulesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		protectionRule := map[string]interface{}{}

		protectionRule["action"] = r.Action

		if r.Description != nil {
			protectionRule["description"] = *r.Description
		}

		exclusions := []interface{}{}
		for _, item := range r.Exclusions {
			exclusions = append(exclusions, ProtectionRuleExclusionToMap(item))
		}
		protectionRule["exclusions"] = exclusions

		if r.Key != nil {
			protectionRule["key"] = *r.Key
		}

		protectionRule["labels"] = r.Labels

		protectionRule["mod_security_rule_ids"] = r.ModSecurityRuleIds

		if r.Name != nil {
			protectionRule["name"] = *r.Name
		}

		resources = append(resources, protectionRule)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, WaasProtectionRulesDataSource().Schema["protection_rules"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("protection_rules", resources); err != nil {
		return err
	}

	return nil
}
