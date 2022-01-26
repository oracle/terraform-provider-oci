// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package budget

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_budget "github.com/oracle/oci-go-sdk/v56/budget"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func BudgetAlertRulesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readBudgetAlertRules,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"budget_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"alert_rules": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(BudgetAlertRuleResource()),
			},
		},
	}
}

func readBudgetAlertRules(d *schema.ResourceData, m interface{}) error {
	sync := &BudgetAlertRulesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BudgetClient()

	return tfresource.ReadResource(sync)
}

type BudgetAlertRulesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_budget.BudgetClient
	Res    *oci_budget.ListAlertRulesResponse
}

func (s *BudgetAlertRulesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BudgetAlertRulesDataSourceCrud) Get() error {
	request := oci_budget.ListAlertRulesRequest{}

	if budgetId, ok := s.D.GetOkExists("budget_id"); ok {
		tmp := budgetId.(string)
		request.BudgetId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_budget.ListAlertRulesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "budget")

	response, err := s.Client.ListAlertRules(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAlertRules(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BudgetAlertRulesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("BudgetAlertRulesDataSource-", BudgetAlertRulesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		alertRule := map[string]interface{}{
			"budget_id": *r.BudgetId,
		}

		if r.DefinedTags != nil {
			alertRule["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			alertRule["description"] = *r.Description
		}

		if r.DisplayName != nil {
			alertRule["display_name"] = *r.DisplayName
		}

		alertRule["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			alertRule["id"] = *r.Id
		}

		if r.Message != nil {
			alertRule["message"] = *r.Message
		}

		if r.Recipients != nil {
			alertRule["recipients"] = *r.Recipients
		}

		alertRule["state"] = r.LifecycleState

		if r.Threshold != nil {
			alertRule["threshold"] = *r.Threshold
		}

		alertRule["threshold_type"] = r.ThresholdType

		if r.TimeCreated != nil {
			alertRule["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			alertRule["time_updated"] = r.TimeUpdated.String()
		}

		alertRule["type"] = r.Type

		if r.Version != nil {
			alertRule["version"] = *r.Version
		}

		resources = append(resources, alertRule)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, BudgetAlertRulesDataSource().Schema["alert_rules"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("alert_rules", resources); err != nil {
		return err
	}

	return nil
}
