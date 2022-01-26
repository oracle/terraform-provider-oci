// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package budget

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_budget "github.com/oracle/oci-go-sdk/v56/budget"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

func BudgetBudgetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createBudgetBudget,
		Read:     readBudgetBudget,
		Update:   updateBudgetBudget,
		Delete:   deleteBudgetBudget,
		Schema: map[string]*schema.Schema{
			// Required
			"amount": {
				Type:     schema.TypeInt, // Float per spec, but the service will only accept integers
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"reset_period": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"budget_processing_period_start_offset": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"target_compartment_id": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"target_type"},
				Deprecated:    tfresource.FieldDeprecatedForAnother("target_compartment_id", "target_type"),
			},
			"target_type": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"target_compartment_id"},
			},
			"targets": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"forecasted_spend": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"state": {
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

func createBudgetBudget(d *schema.ResourceData, m interface{}) error {
	sync := &BudgetBudgetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BudgetClient()

	return tfresource.CreateResource(d, sync)
}

func readBudgetBudget(d *schema.ResourceData, m interface{}) error {
	sync := &BudgetBudgetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BudgetClient()

	return tfresource.ReadResource(sync)
}

func updateBudgetBudget(d *schema.ResourceData, m interface{}) error {
	sync := &BudgetBudgetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BudgetClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteBudgetBudget(d *schema.ResourceData, m interface{}) error {
	sync := &BudgetBudgetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BudgetClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type BudgetBudgetResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_budget.BudgetClient
	Res                    *oci_budget.Budget
	DisableNotFoundRetries bool
}

func (s *BudgetBudgetResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *BudgetBudgetResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *BudgetBudgetResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_budget.BudgetLifecycleStateActive),
	}
}

func (s *BudgetBudgetResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *BudgetBudgetResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *BudgetBudgetResourceCrud) Create() error {
	request := oci_budget.CreateBudgetRequest{}

	if amount, ok := s.D.GetOkExists("amount"); ok {
		tmp := float32(amount.(int))
		request.Amount = &tmp
	}

	if budgetProcessingPeriodStartOffset, ok := s.D.GetOkExists("budget_processing_period_start_offset"); ok {
		tmp := budgetProcessingPeriodStartOffset.(int)
		request.BudgetProcessingPeriodStartOffset = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if resetPeriod, ok := s.D.GetOkExists("reset_period"); ok {
		request.ResetPeriod = oci_budget.ResetPeriodEnum(resetPeriod.(string))
	}

	if targetCompartmentId, ok := s.D.GetOkExists("target_compartment_id"); ok {
		tmp := targetCompartmentId.(string)
		request.TargetCompartmentId = &tmp
	}

	if targetType, ok := s.D.GetOkExists("target_type"); ok {
		request.TargetType = oci_budget.TargetTypeEnum(targetType.(string))
	}

	if targets, ok := s.D.GetOkExists("targets"); ok {
		interfaces := targets.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("targets") {
			request.Targets = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "budget")

	response, err := s.Client.CreateBudget(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Budget
	return nil
}

func (s *BudgetBudgetResourceCrud) Get() error {
	request := oci_budget.GetBudgetRequest{}

	tmp := s.D.Id()
	request.BudgetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "budget")

	response, err := s.Client.GetBudget(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Budget
	return nil
}

func (s *BudgetBudgetResourceCrud) Update() error {
	request := oci_budget.UpdateBudgetRequest{}

	if amount, ok := s.D.GetOkExists("amount"); ok {
		tmp := float32(amount.(int))
		request.Amount = &tmp
	}

	tmp := s.D.Id()
	request.BudgetId = &tmp

	if budgetProcessingPeriodStartOffset, ok := s.D.GetOkExists("budget_processing_period_start_offset"); ok {
		tmp := budgetProcessingPeriodStartOffset.(int)
		request.BudgetProcessingPeriodStartOffset = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if resetPeriod, ok := s.D.GetOkExists("reset_period"); ok {
		request.ResetPeriod = oci_budget.ResetPeriodEnum(resetPeriod.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "budget")

	response, err := s.Client.UpdateBudget(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Budget
	return nil
}

func (s *BudgetBudgetResourceCrud) Delete() error {
	request := oci_budget.DeleteBudgetRequest{}

	tmp := s.D.Id()
	request.BudgetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "budget")

	_, err := s.Client.DeleteBudget(context.Background(), request)
	return err
}

func (s *BudgetBudgetResourceCrud) SetData() error {
	if s.Res.ActualSpend != nil {
		s.D.Set("actual_spend", *s.Res.ActualSpend)
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
		s.D.Set("forecasted_spend", *s.Res.ForecastedSpend)
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
