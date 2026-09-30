// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeSubsettingPolicySubsettingRuleDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["subsetting_policy_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["subsetting_rule_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(DataSafeSubsettingPolicySubsettingRuleResource(), fieldMap, readSingularDataSafeSubsettingPolicySubsettingRuleWithContext)
}

func readSingularDataSafeSubsettingPolicySubsettingRuleWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataSafeSubsettingPolicySubsettingRuleDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataSafeSubsettingPolicySubsettingRuleDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetSubsettingRuleResponse
}

func (s *DataSafeSubsettingPolicySubsettingRuleDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeSubsettingPolicySubsettingRuleDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_data_safe.GetSubsettingRuleRequest{}

	if subsettingPolicyId, ok := s.D.GetOkExists("subsetting_policy_id"); ok {
		tmp := subsettingPolicyId.(string)
		request.SubsettingPolicyId = &tmp
	}

	if subsettingRuleKey, ok := s.D.GetOkExists("subsetting_rule_key"); ok {
		tmp := subsettingRuleKey.(string)
		request.SubsettingRuleKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetSubsettingRule(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeSubsettingPolicySubsettingRuleDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeSubsettingPolicySubsettingRuleDataSource-", DataSafeSubsettingPolicySubsettingRuleDataSource(), s.D))

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	s.D.Set("peer_tables_action", s.Res.PeerTablesAction)

	s.D.Set("related_tables_propagation", s.Res.RelatedTablesPropagation)

	s.D.Set("rule_combination_mode", s.Res.RuleCombinationMode)

	if s.Res.Scope != nil {
		scopeArray := []interface{}{}
		if scopeMap := SubsetScopeToMap(&s.Res.Scope); scopeMap != nil {
			scopeArray = append(scopeArray, scopeMap)
		}
		s.D.Set("scope", scopeArray)
	} else {
		s.D.Set("scope", nil)
	}

	if s.Res.SubsetRuleEntry != nil {
		subsetRuleEntryArray := []interface{}{}
		if subsetRuleEntryMap := SubsetRuleEntryToMap(&s.Res.SubsetRuleEntry); subsetRuleEntryMap != nil {
			subsetRuleEntryArray = append(subsetRuleEntryArray, subsetRuleEntryMap)
		}
		s.D.Set("subset_rule_entry", subsetRuleEntryArray)
	} else {
		s.D.Set("subset_rule_entry", nil)
	}

	return nil
}
