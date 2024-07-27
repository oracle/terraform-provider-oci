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

func DataSafeAlertPolicyRulesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeAlertPolicyRules,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"alert_policy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"alert_policy_rule_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     DataSafeAlertPolicyRuleResource(),
						},
					},
				},
			},
		},
	}
}

func readDataSafeAlertPolicyRules(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAlertPolicyRulesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeAlertPolicyRulesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListAlertPolicyRulesResponse
}

func (s *DataSafeAlertPolicyRulesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeAlertPolicyRulesDataSourceCrud) Get() error {
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
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAlertPolicyRules(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeAlertPolicyRulesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeAlertPolicyRulesDataSource-", DataSafeAlertPolicyRulesDataSource(), s.D))
	resources := []map[string]interface{}{}
	alertPolicyRule := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AlertPolicyRuleSummaryToMap(item))
	}
	alertPolicyRule["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeAlertPolicyRulesDataSource().Schema["alert_policy_rule_collection"].Elem.(*schema.Resource).Schema)
		alertPolicyRule["items"] = items
	}

	resources = append(resources, alertPolicyRule)
	if err := s.D.Set("alert_policy_rule_collection", resources); err != nil {
		return err
	}

	return nil
}
func AlertPolicyRuleSummaryToMap(obj oci_data_safe.AlertPolicyRuleSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Expression != nil {
		result["expression"] = string(*obj.Expression)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	return result
}
