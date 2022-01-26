// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dns

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_dns "github.com/oracle/oci-go-sdk/v56/dns"
)

func DnsSteeringPolicyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDnsSteeringPolicy,
		Read:     readDnsSteeringPolicy,
		Update:   updateDnsSteeringPolicy,
		Delete:   deleteDnsSteeringPolicy,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"template": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"answers": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"rdata": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"rtype": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"is_disabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"pool": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"health_check_monitor_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rules": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"rule_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"FILTER",
								"HEALTH",
								"LIMIT",
								"PRIORITY",
								"WEIGHTED",
							}, true),
						},

						// Optional
						"cases": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"answer_data": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"answer_condition": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"should_keep": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"value": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
											},
										},
									},
									"case_condition": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"count": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"default_answer_data": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"answer_condition": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"should_keep": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"value": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"default_count": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// Computed
			"self": {
				Type:     schema.TypeString,
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
		},
	}
}

func createDnsSteeringPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &DnsSteeringPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.CreateResource(d, sync)
}

func readDnsSteeringPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &DnsSteeringPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.ReadResource(sync)
}

func updateDnsSteeringPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &DnsSteeringPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDnsSteeringPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &DnsSteeringPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DnsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DnsSteeringPolicyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_dns.DnsClient
	Res                    *oci_dns.SteeringPolicy
	DisableNotFoundRetries bool
}

func (s *DnsSteeringPolicyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DnsSteeringPolicyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_dns.SteeringPolicyLifecycleStateCreating),
	}
}

func (s *DnsSteeringPolicyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_dns.SteeringPolicyLifecycleStateActive),
	}
}

func (s *DnsSteeringPolicyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_dns.SteeringPolicyLifecycleStateDeleting),
	}
}

func (s *DnsSteeringPolicyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_dns.SteeringPolicyLifecycleStateDeleted),
	}
}

func (s *DnsSteeringPolicyResourceCrud) Create() error {
	request := oci_dns.CreateSteeringPolicyRequest{}

	if answers, ok := s.D.GetOkExists("answers"); ok {
		interfaces := answers.([]interface{})
		tmp := make([]oci_dns.SteeringPolicyAnswer, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "answers", stateDataIndex)
			converted, err := s.mapToSteeringPolicyAnswer(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("answers") {
			request.Answers = tmp
		}
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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if healthCheckMonitorId, ok := s.D.GetOkExists("health_check_monitor_id"); ok {
		tmp := healthCheckMonitorId.(string)
		request.HealthCheckMonitorId = &tmp
	}

	if rules, ok := s.D.GetOkExists("rules"); ok {
		interfaces := rules.([]interface{})
		tmp := make([]oci_dns.SteeringPolicyRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rules", stateDataIndex)
			converted, err := s.mapToSteeringPolicyRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("rules") {
			request.Rules = tmp
		}
	}

	if template, ok := s.D.GetOkExists("template"); ok {
		request.Template = oci_dns.CreateSteeringPolicyDetailsTemplateEnum(template.(string))
	}

	if ttl, ok := s.D.GetOkExists("ttl"); ok {
		tmp := ttl.(int)
		request.Ttl = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	response, err := s.Client.CreateSteeringPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SteeringPolicy

	if waitErr := tfresource.WaitForCreatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *DnsSteeringPolicyResourceCrud) Get() error {
	request := oci_dns.GetSteeringPolicyRequest{}

	tmp := s.D.Id()
	request.SteeringPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	response, err := s.Client.GetSteeringPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SteeringPolicy
	return nil
}

func (s *DnsSteeringPolicyResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_dns.UpdateSteeringPolicyRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if healthCheckMonitorId, ok := s.D.GetOkExists("health_check_monitor_id"); ok {
		tmp := healthCheckMonitorId.(string)
		request.HealthCheckMonitorId = &tmp
	}

	tmp := s.D.Id()
	request.SteeringPolicyId = &tmp

	if template, ok := s.D.GetOkExists("template"); ok {
		request.Template = oci_dns.UpdateSteeringPolicyDetailsTemplateEnum(template.(string))
	}

	if ttl, ok := s.D.GetOkExists("ttl"); ok {
		tmp := ttl.(int)
		request.Ttl = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	response, err := s.Client.UpdateSteeringPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SteeringPolicy

	// This update does not support work-request
	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *DnsSteeringPolicyResourceCrud) Delete() error {
	request := oci_dns.DeleteSteeringPolicyRequest{}

	tmp := s.D.Id()
	request.SteeringPolicyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	_, err := s.Client.DeleteSteeringPolicy(context.Background(), request)
	return err
}

func (s *DnsSteeringPolicyResourceCrud) SetData() error {
	answers := []interface{}{}
	for _, item := range s.Res.Answers {
		answers = append(answers, SteeringPolicyAnswerToMap(item))
	}
	s.D.Set("answers", answers)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HealthCheckMonitorId != nil {
		s.D.Set("health_check_monitor_id", *s.Res.HealthCheckMonitorId)
	}

	rules := []interface{}{}
	for _, item := range s.Res.Rules {
		rules = append(rules, SteeringPolicyRuleToMap(item))
	}
	s.D.Set("rules", rules)

	if s.Res.Self != nil {
		s.D.Set("self", *s.Res.Self)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("template", s.Res.Template)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.Ttl != nil {
		s.D.Set("ttl", *s.Res.Ttl)
	}

	return nil
}

func (s *DnsSteeringPolicyResourceCrud) mapToSteeringPolicyAnswer(fieldKeyFormat string) (oci_dns.SteeringPolicyAnswer, error) {
	result := oci_dns.SteeringPolicyAnswer{}

	if isDisabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_disabled")); ok {
		tmp := isDisabled.(bool)
		result.IsDisabled = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if pool, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pool")); ok {
		tmp := pool.(string)
		result.Pool = &tmp
	}

	if rdata, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rdata")); ok {
		tmp := rdata.(string)
		result.Rdata = &tmp
	}

	if rtype, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rtype")); ok {
		tmp := rtype.(string)
		result.Rtype = &tmp
	}

	return result, nil
}

func SteeringPolicyAnswerToMap(obj oci_dns.SteeringPolicyAnswer) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsDisabled != nil {
		result["is_disabled"] = bool(*obj.IsDisabled)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Pool != nil {
		result["pool"] = string(*obj.Pool)
	}

	if obj.Rdata != nil {
		result["rdata"] = string(*obj.Rdata)
	}

	if obj.Rtype != nil {
		result["rtype"] = string(*obj.Rtype)
	}

	return result
}

func (s *DnsSteeringPolicyResourceCrud) mapToSteeringPolicyFilterAnswerData(fieldKeyFormat string) (oci_dns.SteeringPolicyFilterAnswerData, error) {
	result := oci_dns.SteeringPolicyFilterAnswerData{}

	if answerCondition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "answer_condition")); ok {
		tmp := answerCondition.(string)
		result.AnswerCondition = &tmp
	}

	if shouldKeep, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_keep")); ok {
		tmp := shouldKeep.(bool)
		result.ShouldKeep = &tmp
	}

	return result, nil
}

func SteeringPolicyFilterAnswerDataToMap(obj oci_dns.SteeringPolicyFilterAnswerData) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AnswerCondition != nil {
		result["answer_condition"] = string(*obj.AnswerCondition)
	}

	if obj.ShouldKeep != nil {
		result["should_keep"] = bool(*obj.ShouldKeep)
	}

	return result
}

func (s *DnsSteeringPolicyResourceCrud) mapToSteeringPolicyFilterRuleCase(fieldKeyFormat string) (oci_dns.SteeringPolicyFilterRuleCase, error) {
	result := oci_dns.SteeringPolicyFilterRuleCase{}

	if answerData, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "answer_data")); ok {
		interfaces := answerData.([]interface{})
		tmp := make([]oci_dns.SteeringPolicyFilterAnswerData, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "answer_data"), stateDataIndex)
			converted, err := s.mapToSteeringPolicyFilterAnswerData(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "answer_data")) {
			result.AnswerData = tmp
		}
	}

	if caseCondition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "case_condition")); ok {
		tmp := caseCondition.(string)
		result.CaseCondition = &tmp
	}

	return result, nil
}

func SteeringPolicyFilterRuleCaseToMap(obj oci_dns.SteeringPolicyFilterRuleCase) map[string]interface{} {
	result := map[string]interface{}{}

	answerData := []interface{}{}
	for _, item := range obj.AnswerData {
		answerData = append(answerData, SteeringPolicyFilterAnswerDataToMap(item))
	}
	result["answer_data"] = answerData

	if obj.CaseCondition != nil {
		result["case_condition"] = string(*obj.CaseCondition)
	}

	return result
}

func (s *DnsSteeringPolicyResourceCrud) mapToSteeringPolicyHealthRuleCase(fieldKeyFormat string) (oci_dns.SteeringPolicyHealthRuleCase, error) {
	result := oci_dns.SteeringPolicyHealthRuleCase{}

	if caseCondition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "case_condition")); ok {
		tmp := caseCondition.(string)
		result.CaseCondition = &tmp
	}

	return result, nil
}

func SteeringPolicyHealthRuleCaseToMap(obj oci_dns.SteeringPolicyHealthRuleCase) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CaseCondition != nil {
		result["case_condition"] = string(*obj.CaseCondition)
	}

	return result
}

func (s *DnsSteeringPolicyResourceCrud) mapToSteeringPolicyLimitRuleCase(fieldKeyFormat string) (oci_dns.SteeringPolicyLimitRuleCase, error) {
	result := oci_dns.SteeringPolicyLimitRuleCase{}

	if caseCondition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "case_condition")); ok {
		tmp := caseCondition.(string)
		result.CaseCondition = &tmp
	}

	if count, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "count")); ok {
		tmp := count.(int)
		result.Count = &tmp
	}

	return result, nil
}

func SteeringPolicyLimitRuleCaseToMap(obj oci_dns.SteeringPolicyLimitRuleCase) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CaseCondition != nil {
		result["case_condition"] = string(*obj.CaseCondition)
	}

	if obj.Count != nil {
		result["count"] = int(*obj.Count)
	}

	return result
}

func (s *DnsSteeringPolicyResourceCrud) mapToSteeringPolicyPriorityAnswerData(fieldKeyFormat string) (oci_dns.SteeringPolicyPriorityAnswerData, error) {
	result := oci_dns.SteeringPolicyPriorityAnswerData{}

	if answerCondition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "answer_condition")); ok {
		tmp := answerCondition.(string)
		result.AnswerCondition = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(int)
		result.Value = &tmp
	}

	return result, nil
}

func SteeringPolicyPriorityAnswerDataToMap(obj oci_dns.SteeringPolicyPriorityAnswerData) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AnswerCondition != nil {
		result["answer_condition"] = string(*obj.AnswerCondition)
	}

	if obj.Value != nil {
		result["value"] = int(*obj.Value)
	}

	return result
}

func (s *DnsSteeringPolicyResourceCrud) mapToSteeringPolicyPriorityRuleCase(fieldKeyFormat string) (oci_dns.SteeringPolicyPriorityRuleCase, error) {
	result := oci_dns.SteeringPolicyPriorityRuleCase{}

	if answerData, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "answer_data")); ok {
		interfaces := answerData.([]interface{})
		tmp := make([]oci_dns.SteeringPolicyPriorityAnswerData, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "answer_data"), stateDataIndex)
			converted, err := s.mapToSteeringPolicyPriorityAnswerData(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "answer_data")) {
			result.AnswerData = tmp
		}
	}

	if caseCondition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "case_condition")); ok {
		tmp := caseCondition.(string)
		result.CaseCondition = &tmp
	}

	return result, nil
}

func SteeringPolicyPriorityRuleCaseToMap(obj oci_dns.SteeringPolicyPriorityRuleCase) map[string]interface{} {
	result := map[string]interface{}{}

	answerData := []interface{}{}
	for _, item := range obj.AnswerData {
		answerData = append(answerData, SteeringPolicyPriorityAnswerDataToMap(item))
	}
	result["answer_data"] = answerData

	if obj.CaseCondition != nil {
		result["case_condition"] = string(*obj.CaseCondition)
	}

	return result
}

func (s *DnsSteeringPolicyResourceCrud) mapToSteeringPolicyRule(fieldKeyFormat string) (oci_dns.SteeringPolicyRule, error) {
	var baseObject oci_dns.SteeringPolicyRule
	//discriminator
	ruleTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rule_type"))
	var ruleType string
	if ok {
		ruleType = ruleTypeRaw.(string)
	} else {
		ruleType = "" // default value
	}
	switch strings.ToLower(ruleType) {
	case strings.ToLower("FILTER"):
		details := oci_dns.SteeringPolicyFilterRule{}
		if cases, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cases")); ok {
			interfaces := cases.([]interface{})
			tmp := make([]oci_dns.SteeringPolicyFilterRuleCase, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "cases"), stateDataIndex)
				converted, err := s.mapToSteeringPolicyFilterRuleCase(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "cases")) {
				details.Cases = tmp
			}
		}
		if defaultAnswerData, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "default_answer_data")); ok {
			interfaces := defaultAnswerData.([]interface{})
			tmp := make([]oci_dns.SteeringPolicyFilterAnswerData, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "default_answer_data"), stateDataIndex)
				converted, err := s.mapToSteeringPolicyFilterAnswerData(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "default_answer_data")) {
				details.DefaultAnswerData = tmp
			}
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		baseObject = details
	case strings.ToLower("HEALTH"):
		details := oci_dns.SteeringPolicyHealthRule{}
		if cases, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cases")); ok {
			interfaces := cases.([]interface{})
			tmp := make([]oci_dns.SteeringPolicyHealthRuleCase, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "cases"), stateDataIndex)
				converted, err := s.mapToSteeringPolicyHealthRuleCase(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "cases")) {
				details.Cases = tmp
			}
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		baseObject = details
	case strings.ToLower("LIMIT"):
		details := oci_dns.SteeringPolicyLimitRule{}
		if cases, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cases")); ok {
			interfaces := cases.([]interface{})
			tmp := make([]oci_dns.SteeringPolicyLimitRuleCase, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "cases"), stateDataIndex)
				converted, err := s.mapToSteeringPolicyLimitRuleCase(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "cases")) {
				details.Cases = tmp
			}
		}
		if defaultCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "default_count")); ok {
			tmp := defaultCount.(int)
			details.DefaultCount = &tmp
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		baseObject = details
	case strings.ToLower("PRIORITY"):
		details := oci_dns.SteeringPolicyPriorityRule{}
		if cases, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cases")); ok {
			interfaces := cases.([]interface{})
			tmp := make([]oci_dns.SteeringPolicyPriorityRuleCase, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "cases"), stateDataIndex)
				converted, err := s.mapToSteeringPolicyPriorityRuleCase(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "cases")) {
				details.Cases = tmp
			}
		}
		if defaultAnswerData, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "default_answer_data")); ok {
			interfaces := defaultAnswerData.([]interface{})
			tmp := make([]oci_dns.SteeringPolicyPriorityAnswerData, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "default_answer_data"), stateDataIndex)
				converted, err := s.mapToSteeringPolicyPriorityAnswerData(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "default_answer_data")) {
				details.DefaultAnswerData = tmp
			}
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		baseObject = details
	case strings.ToLower("WEIGHTED"):
		details := oci_dns.SteeringPolicyWeightedRule{}
		if cases, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cases")); ok {
			interfaces := cases.([]interface{})
			tmp := make([]oci_dns.SteeringPolicyWeightedRuleCase, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "cases"), stateDataIndex)
				converted, err := s.mapToSteeringPolicyWeightedRuleCase(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "cases")) {
				details.Cases = tmp
			}
		}
		if defaultAnswerData, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "default_answer_data")); ok {
			interfaces := defaultAnswerData.([]interface{})
			tmp := make([]oci_dns.SteeringPolicyWeightedAnswerData, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "default_answer_data"), stateDataIndex)
				converted, err := s.mapToSteeringPolicyWeightedAnswerData(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "default_answer_data")) {
				details.DefaultAnswerData = tmp
			}
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown rule_type '%v' was specified", ruleType)
	}
	return baseObject, nil
}

func SteeringPolicyRuleToMap(obj oci_dns.SteeringPolicyRule) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_dns.SteeringPolicyFilterRule:
		result["rule_type"] = "FILTER"

		cases := []interface{}{}
		for _, item := range v.Cases {
			cases = append(cases, SteeringPolicyFilterRuleCaseToMap(item))
		}
		result["cases"] = cases

		defaultAnswerData := []interface{}{}
		for _, item := range v.DefaultAnswerData {
			defaultAnswerData = append(defaultAnswerData, SteeringPolicyFilterAnswerDataToMap(item))
		}
		result["default_answer_data"] = defaultAnswerData

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}
	case oci_dns.SteeringPolicyHealthRule:
		result["rule_type"] = "HEALTH"

		cases := []interface{}{}
		for _, item := range v.Cases {
			cases = append(cases, SteeringPolicyHealthRuleCaseToMap(item))
		}
		result["cases"] = cases

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}
	case oci_dns.SteeringPolicyLimitRule:
		result["rule_type"] = "LIMIT"

		cases := []interface{}{}
		for _, item := range v.Cases {
			cases = append(cases, SteeringPolicyLimitRuleCaseToMap(item))
		}
		result["cases"] = cases

		if v.DefaultCount != nil {
			result["default_count"] = int(*v.DefaultCount)
		}

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}
	case oci_dns.SteeringPolicyPriorityRule:
		result["rule_type"] = "PRIORITY"

		cases := []interface{}{}
		for _, item := range v.Cases {
			cases = append(cases, SteeringPolicyPriorityRuleCaseToMap(item))
		}
		result["cases"] = cases

		defaultAnswerData := []interface{}{}
		for _, item := range v.DefaultAnswerData {
			defaultAnswerData = append(defaultAnswerData, SteeringPolicyPriorityAnswerDataToMap(item))
		}
		result["default_answer_data"] = defaultAnswerData

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}
	case oci_dns.SteeringPolicyWeightedRule:
		result["rule_type"] = "WEIGHTED"

		cases := []interface{}{}
		for _, item := range v.Cases {
			cases = append(cases, SteeringPolicyWeightedRuleCaseToMap(item))
		}
		result["cases"] = cases

		defaultAnswerData := []interface{}{}
		for _, item := range v.DefaultAnswerData {
			defaultAnswerData = append(defaultAnswerData, SteeringPolicyWeightedAnswerDataToMap(item))
		}
		result["default_answer_data"] = defaultAnswerData

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}
	default:
		log.Printf("[WARN] Received 'rule_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *DnsSteeringPolicyResourceCrud) mapToSteeringPolicyWeightedAnswerData(fieldKeyFormat string) (oci_dns.SteeringPolicyWeightedAnswerData, error) {
	result := oci_dns.SteeringPolicyWeightedAnswerData{}

	if answerCondition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "answer_condition")); ok {
		tmp := answerCondition.(string)
		result.AnswerCondition = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(int)
		result.Value = &tmp
	}

	return result, nil
}

func SteeringPolicyWeightedAnswerDataToMap(obj oci_dns.SteeringPolicyWeightedAnswerData) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AnswerCondition != nil {
		result["answer_condition"] = string(*obj.AnswerCondition)
	}

	if obj.Value != nil {
		result["value"] = int(*obj.Value)
	}

	return result
}

func (s *DnsSteeringPolicyResourceCrud) mapToSteeringPolicyWeightedRuleCase(fieldKeyFormat string) (oci_dns.SteeringPolicyWeightedRuleCase, error) {
	result := oci_dns.SteeringPolicyWeightedRuleCase{}

	if answerData, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "answer_data")); ok {
		interfaces := answerData.([]interface{})
		tmp := make([]oci_dns.SteeringPolicyWeightedAnswerData, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "answer_data"), stateDataIndex)
			converted, err := s.mapToSteeringPolicyWeightedAnswerData(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "answer_data")) {
			result.AnswerData = tmp
		}
	}

	if caseCondition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "case_condition")); ok {
		tmp := caseCondition.(string)
		result.CaseCondition = &tmp
	}

	return result, nil
}

func SteeringPolicyWeightedRuleCaseToMap(obj oci_dns.SteeringPolicyWeightedRuleCase) map[string]interface{} {
	result := map[string]interface{}{}

	answerData := []interface{}{}
	for _, item := range obj.AnswerData {
		answerData = append(answerData, SteeringPolicyWeightedAnswerDataToMap(item))
	}
	result["answer_data"] = answerData

	if obj.CaseCondition != nil {
		result["case_condition"] = string(*obj.CaseCondition)
	}

	return result
}

func (s *DnsSteeringPolicyResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_dns.ChangeSteeringPolicyCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.SteeringPolicyId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dns")

	_, err := s.Client.ChangeSteeringPolicyCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
