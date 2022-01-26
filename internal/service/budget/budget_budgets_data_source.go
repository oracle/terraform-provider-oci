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

func BudgetBudgetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readBudgetBudgets,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
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
			"target_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"budgets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(BudgetBudgetResource()),
			},
		},
	}
}

func readBudgetBudgets(d *schema.ResourceData, m interface{}) error {
	sync := &BudgetBudgetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BudgetClient()

	return tfresource.ReadResource(sync)
}

type BudgetBudgetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_budget.BudgetClient
	Res    *oci_budget.ListBudgetsResponse
}

func (s *BudgetBudgetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BudgetBudgetsDataSourceCrud) Get() error {
	request := oci_budget.ListBudgetsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_budget.ListBudgetsLifecycleStateEnum(state.(string))
	}

	if targetType, ok := s.D.GetOkExists("target_type"); ok {
		request.TargetType = oci_budget.ListBudgetsTargetTypeEnum(targetType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "budget")

	response, err := s.Client.ListBudgets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBudgets(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BudgetBudgetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("BudgetBudgetsDataSource-", BudgetBudgetsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		budget := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.ActualSpend != nil {
			budget["actual_spend"] = *r.ActualSpend
		}

		if r.AlertRuleCount != nil {
			budget["alert_rule_count"] = *r.AlertRuleCount
		}

		if r.Amount != nil {
			budget["amount"] = *r.Amount
		}

		if r.BudgetProcessingPeriodStartOffset != nil {
			budget["budget_processing_period_start_offset"] = *r.BudgetProcessingPeriodStartOffset
		}

		if r.DefinedTags != nil {
			budget["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			budget["description"] = *r.Description
		}

		if r.DisplayName != nil {
			budget["display_name"] = *r.DisplayName
		}

		if r.ForecastedSpend != nil {
			budget["forecasted_spend"] = *r.ForecastedSpend
		}

		budget["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			budget["id"] = *r.Id
		}

		budget["reset_period"] = r.ResetPeriod

		budget["state"] = r.LifecycleState

		if r.TargetCompartmentId != nil {
			budget["target_compartment_id"] = *r.TargetCompartmentId
		}

		budget["target_type"] = r.TargetType

		budget["targets"] = r.Targets

		if r.TimeCreated != nil {
			budget["time_created"] = r.TimeCreated.String()
		}

		if r.TimeSpendComputed != nil {
			budget["time_spend_computed"] = r.TimeSpendComputed.String()
		}

		if r.TimeUpdated != nil {
			budget["time_updated"] = r.TimeUpdated.String()
		}

		if r.Version != nil {
			budget["version"] = *r.Version
		}

		resources = append(resources, budget)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, BudgetBudgetsDataSource().Schema["budgets"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("budgets", resources); err != nil {
		return err
	}

	return nil
}
