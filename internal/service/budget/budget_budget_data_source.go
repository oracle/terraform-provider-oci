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

func BudgetBudgetDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["budget_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(BudgetBudgetResource(), fieldMap, readSingularBudgetBudget)
}

func readSingularBudgetBudget(d *schema.ResourceData, m interface{}) error {
	sync := &BudgetBudgetDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BudgetClient()

	return tfresource.ReadResource(sync)
}

type BudgetBudgetDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_budget.BudgetClient
	Res    *oci_budget.GetBudgetResponse
}

func (s *BudgetBudgetDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BudgetBudgetDataSourceCrud) Get() error {
	request := oci_budget.GetBudgetRequest{}

	if budgetId, ok := s.D.GetOkExists("budget_id"); ok {
		tmp := budgetId.(string)
		request.BudgetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "budget")

	response, err := s.Client.GetBudget(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *BudgetBudgetDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ActualSpend != nil {
		s.D.Set("actual_spend", int(*s.Res.ActualSpend))
	}

	if s.Res.AlertRuleCount != nil {
		s.D.Set("alert_rule_count", *s.Res.AlertRuleCount)
	}

	if s.Res.Amount != nil {
		s.D.Set("amount", *s.Res.Amount)
	}

	if s.Res.BudgetProcessingPeriodStartOffset != nil {
		s.D.Set("budget_processing_period_start_offset", *s.Res.BudgetProcessingPeriodStartOffset)
	}

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

	if s.Res.ForecastedSpend != nil {
		s.D.Set("forecasted_spend", int(*s.Res.ForecastedSpend))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("reset_period", s.Res.ResetPeriod)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TargetCompartmentId != nil {
		s.D.Set("target_compartment_id", *s.Res.TargetCompartmentId)
	}

	s.D.Set("target_type", s.Res.TargetType)

	s.D.Set("targets", s.Res.Targets)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeSpendComputed != nil {
		s.D.Set("time_spend_computed", s.Res.TimeSpendComputed.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}
