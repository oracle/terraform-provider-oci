// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package budget

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	oci_budget "github.com/oracle/oci-go-sdk/v56/budget"
)

func BudgetAlertRuleResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createBudgetAlertRule,
		Read:     readBudgetAlertRule,
		Update:   updateBudgetAlertRule,
		Delete:   deleteBudgetAlertRule,
		Schema: map[string]*schema.Schema{
			// Required
			"budget_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"threshold": {
				Type:             schema.TypeFloat,
				Required:         true,
				DiffSuppressFunc: utils.MonetaryDiffSuppress,
			},
			"threshold_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
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
			"message": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"recipients": {
				Type:     schema.TypeString,
				Optional: true,
			},

			// Computed
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
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

func createBudgetAlertRule(d *schema.ResourceData, m interface{}) error {
	sync := &BudgetAlertRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BudgetClient()

	return tfresource.CreateResource(d, sync)
}

func readBudgetAlertRule(d *schema.ResourceData, m interface{}) error {
	sync := &BudgetAlertRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BudgetClient()

	return tfresource.ReadResource(sync)
}

func updateBudgetAlertRule(d *schema.ResourceData, m interface{}) error {
	sync := &BudgetAlertRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BudgetClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteBudgetAlertRule(d *schema.ResourceData, m interface{}) error {
	sync := &BudgetAlertRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BudgetClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type BudgetAlertRuleResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_budget.BudgetClient
	Res                    *oci_budget.AlertRule
	DisableNotFoundRetries bool
}

func (s *BudgetAlertRuleResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *BudgetAlertRuleResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *BudgetAlertRuleResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_budget.AlertRuleLifecycleStateActive),
	}
}

func (s *BudgetAlertRuleResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *BudgetAlertRuleResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *BudgetAlertRuleResourceCrud) Create() error {
	request := oci_budget.CreateAlertRuleRequest{}

	if budgetId, ok := s.D.GetOkExists("budget_id"); ok {
		tmp := budgetId.(string)
		request.BudgetId = &tmp
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

	if message, ok := s.D.GetOkExists("message"); ok {
		tmp := message.(string)
		request.Message = &tmp
	}

	if recipients, ok := s.D.GetOkExists("recipients"); ok {
		tmp := recipients.(string)
		request.Recipients = &tmp
	}

	if threshold, ok := s.D.GetOkExists("threshold"); ok {
		tmp := float32(threshold.(float64))
		request.Threshold = &tmp
	}

	if thresholdType, ok := s.D.GetOkExists("threshold_type"); ok {
		request.ThresholdType = oci_budget.ThresholdTypeEnum(thresholdType.(string))
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_budget.AlertTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "budget")

	response, err := s.Client.CreateAlertRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AlertRule
	return nil
}

func (s *BudgetAlertRuleResourceCrud) Get() error {
	request := oci_budget.GetAlertRuleRequest{}

	tmp := s.D.Id()
	request.AlertRuleId = &tmp

	if budgetId, ok := s.D.GetOkExists("budget_id"); ok {
		tmp := budgetId.(string)
		request.BudgetId = &tmp
	}

	alertRuleId, budgetId, err := parseAlertRuleCompositeId(s.D.Id())
	if err == nil {
		request.AlertRuleId = &alertRuleId
		request.BudgetId = &budgetId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "budget")

	response, err := s.Client.GetAlertRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AlertRule
	return nil
}

func (s *BudgetAlertRuleResourceCrud) Update() error {
	request := oci_budget.UpdateAlertRuleRequest{}

	tmp := s.D.Id()
	request.AlertRuleId = &tmp

	if budgetId, ok := s.D.GetOkExists("budget_id"); ok {
		tmp := budgetId.(string)
		request.BudgetId = &tmp
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

	if message, ok := s.D.GetOkExists("message"); ok {
		tmp := message.(string)
		request.Message = &tmp
	}

	if recipients, ok := s.D.GetOkExists("recipients"); ok {
		tmp := recipients.(string)
		request.Recipients = &tmp
	}

	if threshold, ok := s.D.GetOkExists("threshold"); ok {
		tmp := float32(threshold.(float64))
		request.Threshold = &tmp
	}

	if thresholdType, ok := s.D.GetOkExists("threshold_type"); ok {
		request.ThresholdType = oci_budget.ThresholdTypeEnum(thresholdType.(string))
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_budget.AlertTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "budget")

	response, err := s.Client.UpdateAlertRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AlertRule
	return nil
}

func (s *BudgetAlertRuleResourceCrud) Delete() error {
	request := oci_budget.DeleteAlertRuleRequest{}

	tmp := s.D.Id()
	request.AlertRuleId = &tmp

	if budgetId, ok := s.D.GetOkExists("budget_id"); ok {
		tmp := budgetId.(string)
		request.BudgetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "budget")

	_, err := s.Client.DeleteAlertRule(context.Background(), request)
	return err
}

func (s *BudgetAlertRuleResourceCrud) SetData() error {

	alertRuleId, budgetId, err := parseAlertRuleCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(alertRuleId)
		s.D.Set("budget_id", budgetId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.BudgetId != nil {
		s.D.Set("budget_id", *s.Res.BudgetId)
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

	if s.Res.Message != nil {
		s.D.Set("message", *s.Res.Message)
	}

	if s.Res.Recipients != nil {
		s.D.Set("recipients", *s.Res.Recipients)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.Threshold != nil {
		s.D.Set("threshold", *s.Res.Threshold)
	}

	s.D.Set("threshold_type", s.Res.ThresholdType)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("type", s.Res.Type)

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}

func GetAlertRuleCompositeId(alertRuleId string, budgetId string) string {
	alertRuleId = url.PathEscape(alertRuleId)
	budgetId = url.PathEscape(budgetId)
	compositeId := "budgets/" + budgetId + "/alertRules/" + alertRuleId
	return compositeId
}

func parseAlertRuleCompositeId(compositeId string) (alertRuleId string, budgetId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("budgets/.*/alertRules/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	budgetId, _ = url.PathUnescape(parts[1])
	alertRuleId, _ = url.PathUnescape(parts[3])

	return
}
