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

func WaasProtectionRuleDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularWaasProtectionRule,
		Schema: map[string]*schema.Schema{
			"protection_rule_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"waas_policy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"action": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"exclusions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"exclusions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"target": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"labels": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"mod_security_rule_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularWaasProtectionRule(d *schema.ResourceData, m interface{}) error {
	sync := &WaasProtectionRuleDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WaasClient()

	return tfresource.ReadResource(sync)
}

type WaasProtectionRuleDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_waas.WaasClient
	Res    *oci_waas.GetProtectionRuleResponse
}

func (s *WaasProtectionRuleDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WaasProtectionRuleDataSourceCrud) Get() error {
	request := oci_waas.GetProtectionRuleRequest{}

	if protectionRuleKey, ok := s.D.GetOkExists("protection_rule_key"); ok {
		tmp := protectionRuleKey.(string)
		request.ProtectionRuleKey = &tmp
	}

	if waasPolicyId, ok := s.D.GetOkExists("waas_policy_id"); ok {
		tmp := waasPolicyId.(string)
		request.WaasPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "waas")

	response, err := s.Client.GetProtectionRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *WaasProtectionRuleDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceID())

	s.D.Set("action", s.Res.Action)

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	exclusions := []interface{}{}
	for _, item := range s.Res.Exclusions {
		exclusions = append(exclusions, ProtectionRuleExclusionToMap(item))
	}
	s.D.Set("exclusions", exclusions)

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	s.D.Set("labels", s.Res.Labels)

	s.D.Set("mod_security_rule_ids", s.Res.ModSecurityRuleIds)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	return nil
}
