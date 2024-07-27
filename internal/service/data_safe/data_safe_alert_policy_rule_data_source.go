// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeAlertPolicyRuleDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["alert_policy_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["rule_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataSafeAlertPolicyRuleResource(), fieldMap, readSingularDataSafeAlertPolicyRule)
}

func readSingularDataSafeAlertPolicyRule(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAlertPolicyRuleDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeAlertPolicyRuleDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.GetAlertPolicyRuleResponse
}

func (s *DataSafeAlertPolicyRuleDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeAlertPolicyRuleDataSourceCrud) Get() error {
	request := oci_data_safe.GetAlertPolicyRuleRequest{}

	if alertPolicyId, ok := s.D.GetOkExists("alert_policy_id"); ok {
		tmp := alertPolicyId.(string)
		request.AlertPolicyId = &tmp
	}

	if ruleKey, ok := s.D.GetOkExists("rule_key"); ok {
		tmp := ruleKey.(string)
		request.RuleKey = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.GetAlertPolicyRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeAlertPolicyRuleDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeAlertPolicyRuleDataSource-", DataSafeAlertPolicyRuleDataSource(), s.D))

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Expression != nil {
		s.D.Set("expression", *s.Res.Expression)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
