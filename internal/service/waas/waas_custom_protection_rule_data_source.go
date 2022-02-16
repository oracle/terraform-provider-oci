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

func WaasCustomProtectionRuleDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["custom_protection_rule_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(WaasCustomProtectionRuleResource(), fieldMap, readSingularWaasCustomProtectionRule)
}

func readSingularWaasCustomProtectionRule(d *schema.ResourceData, m interface{}) error {
	sync := &WaasCustomProtectionRuleDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WaasClient()

	return tfresource.ReadResource(sync)
}

type WaasCustomProtectionRuleDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_waas.WaasClient
	Res    *oci_waas.GetCustomProtectionRuleResponse
}

func (s *WaasCustomProtectionRuleDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WaasCustomProtectionRuleDataSourceCrud) Get() error {
	request := oci_waas.GetCustomProtectionRuleRequest{}

	if customProtectionRuleId, ok := s.D.GetOkExists("custom_protection_rule_id"); ok {
		tmp := customProtectionRuleId.(string)
		request.CustomProtectionRuleId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "waas")

	response, err := s.Client.GetCustomProtectionRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *WaasCustomProtectionRuleDataSourceCrud) SetData() error {
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

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("mod_security_rule_ids", s.Res.ModSecurityRuleIds)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.Template != nil {
		s.D.Set("template", *s.Res.Template)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
