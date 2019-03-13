// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_budget "github.com/oracle/oci-go-sdk/budget"
)

func BudgetBudgetDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularBudgetBudget,
		Schema: map[string]*schema.Schema{
			"budget_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"actual_spend": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"alert_rule_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"amount": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"forecasted_spend": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"reset_period": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"target_compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_spend_computed": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func readSingularBudgetBudget(d *schema.ResourceData, m interface{}) error {
	sync := &BudgetBudgetDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).budgetClient

	return ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "budget")

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

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
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

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeSpendComputed != nil {
		s.D.Set("time_spend_computed", *s.Res.TimeSpendComputed)
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}
