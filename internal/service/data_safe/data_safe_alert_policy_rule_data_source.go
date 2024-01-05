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
	return &schema.Resource{
		Read: readSingularDataSafeAlertPolicyRule,
		Schema: map[string]*schema.Schema{
			"alert_policy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"expression": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
		DeprecationMessage: tfresource.DatasourceDeprecatedForAnother("oci_data_safe_alert_policy_rule", "oci_data_safe_alert_policy_rules"),
	}
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
	Res    *oci_data_safe.ListAlertPolicyRulesResponse
}

func (s *DataSafeAlertPolicyRuleDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeAlertPolicyRuleDataSourceCrud) Get() error {
	request := oci_data_safe.ListAlertPolicyRulesRequest{}

	if alertPolicyId, ok := s.D.GetOkExists("alert_policy_id"); ok {
		tmp := alertPolicyId.(string)
		request.AlertPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListAlertPolicyRules(context.Background(), request)
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

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AlertPolicyRuleSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}
