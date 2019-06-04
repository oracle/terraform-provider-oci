// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	oci_autoscaling "github.com/oracle/oci-go-sdk/autoscaling"
)

func AutoscalingAutoScalingConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createAutoscalingAutoScalingConfiguration,
		Read:     readAutoscalingAutoScalingConfiguration,
		Update:   updateAutoscalingAutoScalingConfiguration,
		Delete:   deleteAutoscalingAutoScalingConfiguration,
		Schema: map[string]*schema.Schema{
			// Required
			"auto_scaling_resources": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"policies": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"capacity": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"initial": {
										Type:     schema.TypeInt,
										Required: true,
										ForceNew: true,
									},
									"max": {
										Type:     schema.TypeInt,
										Required: true,
										ForceNew: true,
									},
									"min": {
										Type:     schema.TypeInt,
										Required: true,
										ForceNew: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"policy_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"threshold",
							}, true),
						},
						"rules": {
							Type:     schema.TypeSet,
							Required: true,
							ForceNew: true,
							Set:      autoScalingConfigurationPolicyRulesHashCodeForSets,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"action": {
										Type:     schema.TypeList,
										Required: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"type": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"value": {
													Type:     schema.TypeInt,
													Required: true,
													ForceNew: true,
												},

												// Optional

												// Computed
											},
										},
									},
									"metric": {
										Type:     schema.TypeList,
										Required: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"metric_type": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"threshold": {
													Type:     schema.TypeList,
													Required: true,
													ForceNew: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"operator": {
																Type:     schema.TypeString,
																Required: true,
																ForceNew: true,
															},
															"value": {
																Type:     schema.TypeInt,
																Required: true,
																ForceNew: true,
															},

															// Optional

															// Computed
														},
													},
												},

												// Optional

												// Computed
											},
										},
									},
									// Modifying to Required, since we do not have a good work around for calculating hashcode for optional computed fields
									"display_name": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},

									// Computed
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						// Optional
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			// Optional
			"cool_down_in_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
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
			"is_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// Computed
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createAutoscalingAutoScalingConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &AutoscalingAutoScalingConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).autoScalingClient

	return CreateResource(d, sync)
}

func readAutoscalingAutoScalingConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &AutoscalingAutoScalingConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).autoScalingClient

	return ReadResource(sync)
}

func updateAutoscalingAutoScalingConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &AutoscalingAutoScalingConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).autoScalingClient

	return UpdateResource(d, sync)
}

func deleteAutoscalingAutoScalingConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &AutoscalingAutoScalingConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).autoScalingClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type AutoscalingAutoScalingConfigurationResourceCrud struct {
	BaseCrud
	Client                 *oci_autoscaling.AutoScalingClient
	Res                    *oci_autoscaling.AutoScalingConfiguration
	DisableNotFoundRetries bool
}

func (s *AutoscalingAutoScalingConfigurationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *AutoscalingAutoScalingConfigurationResourceCrud) Create() error {
	request := oci_autoscaling.CreateAutoScalingConfigurationRequest{}

	if autoScalingResources, ok := s.D.GetOkExists("auto_scaling_resources"); ok {
		if tmpList := autoScalingResources.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "auto_scaling_resources", 0)
			tmp, err := s.mapToResource(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Resource = tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if coolDownInSeconds, ok := s.D.GetOkExists("cool_down_in_seconds"); ok {
		tmp := coolDownInSeconds.(int)
		request.CoolDownInSeconds = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	request.Policies = []oci_autoscaling.CreateAutoScalingPolicyDetails{}
	if policies, ok := s.D.GetOkExists("policies"); ok {
		interfaces := policies.([]interface{})
		tmp := make([]oci_autoscaling.CreateAutoScalingPolicyDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "policies", stateDataIndex)
			converted, err := s.mapToCreateAutoScalingPolicyDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		request.Policies = tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "autoscaling")

	response, err := s.Client.CreateAutoScalingConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutoScalingConfiguration
	return nil
}

func (s *AutoscalingAutoScalingConfigurationResourceCrud) Get() error {
	request := oci_autoscaling.GetAutoScalingConfigurationRequest{}

	tmp := s.D.Id()
	request.AutoScalingConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "autoscaling")

	response, err := s.Client.GetAutoScalingConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutoScalingConfiguration
	return nil
}

func (s *AutoscalingAutoScalingConfigurationResourceCrud) Update() error {
	request := oci_autoscaling.UpdateAutoScalingConfigurationRequest{}

	tmp := s.D.Id()
	request.AutoScalingConfigurationId = &tmp

	if coolDownInSeconds, ok := s.D.GetOkExists("cool_down_in_seconds"); ok {
		tmp := coolDownInSeconds.(int)
		request.CoolDownInSeconds = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "autoscaling")

	response, err := s.Client.UpdateAutoScalingConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutoScalingConfiguration
	return nil
}

func (s *AutoscalingAutoScalingConfigurationResourceCrud) Delete() error {
	request := oci_autoscaling.DeleteAutoScalingConfigurationRequest{}

	tmp := s.D.Id()
	request.AutoScalingConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "autoscaling")

	_, err := s.Client.DeleteAutoScalingConfiguration(context.Background(), request)
	return err
}

func (s *AutoscalingAutoScalingConfigurationResourceCrud) SetData() error {
	if s.Res.Resource != nil {
		autoScalingResourcesArray := []interface{}{}
		if autoScalingResourcesMap := ResourceToMap(&s.Res.Resource); autoScalingResourcesMap != nil {
			autoScalingResourcesArray = append(autoScalingResourcesArray, autoScalingResourcesMap)
		}
		s.D.Set("auto_scaling_resources", autoScalingResourcesArray)
	} else {
		s.D.Set("auto_scaling_resources", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CoolDownInSeconds != nil {
		s.D.Set("cool_down_in_seconds", *s.Res.CoolDownInSeconds)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	policies := []interface{}{}
	for _, item := range s.Res.Policies {
		policies = append(policies, AutoScalingPolicyToMap(item, false))
	}
	s.D.Set("policies", policies)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *AutoscalingAutoScalingConfigurationResourceCrud) mapToAction(fieldKeyFormat string) (oci_autoscaling.Action, error) {
	result := oci_autoscaling.Action{}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_autoscaling.ActionTypeEnum(type_.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(int)
		result.Value = &tmp
	}

	return result, nil
}

func ActionToMap(obj *oci_autoscaling.Action) map[string]interface{} {
	result := map[string]interface{}{}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = int(*obj.Value)
	}

	return result
}

func (s *AutoscalingAutoScalingConfigurationResourceCrud) mapToCapacity(fieldKeyFormat string) (oci_autoscaling.Capacity, error) {
	result := oci_autoscaling.Capacity{}

	if initial, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "initial")); ok {
		tmp := initial.(int)
		result.Initial = &tmp
	}

	if max, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max")); ok {
		tmp := max.(int)
		result.Max = &tmp
	}

	if min, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "min")); ok {
		tmp := min.(int)
		result.Min = &tmp
	}

	return result, nil
}

func CapacityToMap(obj *oci_autoscaling.Capacity) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Initial != nil {
		result["initial"] = int(*obj.Initial)
	}

	if obj.Max != nil {
		result["max"] = int(*obj.Max)
	}

	if obj.Min != nil {
		result["min"] = int(*obj.Min)
	}

	return result
}

func (s *AutoscalingAutoScalingConfigurationResourceCrud) mapToCreateAutoScalingPolicyDetails(fieldKeyFormat string) (oci_autoscaling.CreateAutoScalingPolicyDetails, error) {
	var baseObject oci_autoscaling.CreateAutoScalingPolicyDetails
	//discriminator
	policyTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "policy_type"))
	var policyType string
	if ok {
		policyType = policyTypeRaw.(string)
	} else {
		policyType = "" // default value
	}
	switch strings.ToLower(policyType) {
	case strings.ToLower("threshold"):
		details := oci_autoscaling.CreateThresholdPolicyDetails{}
		details.Rules = []oci_autoscaling.CreateConditionDetails{}
		if rules, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rules")); ok {
			set := rules.(*schema.Set)
			interfaces := set.List()
			tmp := make([]oci_autoscaling.CreateConditionDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := autoScalingConfigurationPolicyRulesHashCodeForSets(interfaces[i])
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "rules"), stateDataIndex)
				converted, err := s.mapToCreateConditionDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			details.Rules = tmp
		}
		if capacity, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "capacity")); ok {
			if tmpList := capacity.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "capacity"), 0)
				tmp, err := s.mapToCapacity(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert capacity, encountered error: %v", err)
				}
				details.Capacity = &tmp
			}
		}
		if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown policy_type '%v' was specified", policyType)
	}
	return baseObject, nil
}

func AutoScalingPolicyToMap(obj oci_autoscaling.AutoScalingPolicy, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_autoscaling.ThresholdPolicy:
		result["policy_type"] = "threshold"

		if v.Rules != nil {
			rules := []interface{}{}
			for _, item := range v.Rules {
				rules = append(rules, CreateConditionDetailsToMap(item))
			}
			if datasource {
				result["rules"] = rules
			} else {
				result["rules"] = schema.NewSet(autoScalingConfigurationPolicyRulesHashCodeForSets, rules)
			}
		}

		if v.Capacity != nil {
			result["capacity"] = []interface{}{CapacityToMap(v.Capacity)}
		}

		if v.DisplayName != nil {
			result["display_name"] = string(*v.DisplayName)
		}

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}

		if v.TimeCreated != nil {
			result["time_created"] = v.TimeCreated.String()
		}
	default:
		log.Printf("[WARN] Received 'policy_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *AutoscalingAutoScalingConfigurationResourceCrud) mapToCreateConditionDetails(fieldKeyFormat string) (oci_autoscaling.CreateConditionDetails, error) {
	result := oci_autoscaling.CreateConditionDetails{}

	if action, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action")); ok {
		if tmpList := action.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "action"), 0)
			tmp, err := s.mapToAction(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert action, encountered error: %v", err)
			}
			result.Action = &tmp
		}
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if metric, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metric")); ok {
		if tmpList := metric.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "metric"), 0)
			tmp, err := s.mapToMetric(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert metric, encountered error: %v", err)
			}
			result.Metric = &tmp
		}
	}

	return result, nil
}

func CreateConditionDetailsToMap(obj oci_autoscaling.Condition) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Action != nil {
		result["action"] = []interface{}{ActionToMap(obj.Action)}
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Metric != nil {
		result["metric"] = []interface{}{MetricToMap(obj.Metric)}
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}

func (s *AutoscalingAutoScalingConfigurationResourceCrud) mapToMetric(fieldKeyFormat string) (oci_autoscaling.Metric, error) {
	result := oci_autoscaling.Metric{}

	if metricType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metric_type")); ok {
		result.MetricType = oci_autoscaling.MetricMetricTypeEnum(metricType.(string))
	}

	if threshold, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "threshold")); ok {
		if tmpList := threshold.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "threshold"), 0)
			tmp, err := s.mapToThreshold(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert threshold, encountered error: %v", err)
			}
			result.Threshold = &tmp
		}
	}

	return result, nil
}

func MetricToMap(obj *oci_autoscaling.Metric) map[string]interface{} {
	result := map[string]interface{}{}

	result["metric_type"] = string(obj.MetricType)

	if obj.Threshold != nil {
		result["threshold"] = []interface{}{ThresholdToMap(obj.Threshold)}
	}

	return result
}

func (s *AutoscalingAutoScalingConfigurationResourceCrud) mapToResource(fieldKeyFormat string) (oci_autoscaling.Resource, error) {
	var baseObject oci_autoscaling.Resource
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("instancePool"):
		details := oci_autoscaling.InstancePoolResource{}
		if idRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
			var id_ string
			id_ = idRaw.(string)
			details.Id = &id_
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func ResourceToMap(obj *oci_autoscaling.Resource) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_autoscaling.InstancePoolResource:
		result["type"] = "instancePool"
		result["id"] = string(*v.Id)
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *AutoscalingAutoScalingConfigurationResourceCrud) mapToThreshold(fieldKeyFormat string) (oci_autoscaling.Threshold, error) {
	result := oci_autoscaling.Threshold{}

	if operator, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operator")); ok {
		result.Operator = oci_autoscaling.ThresholdOperatorEnum(operator.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(int)
		result.Value = &tmp
	}

	return result, nil
}

func ThresholdToMap(obj *oci_autoscaling.Threshold) map[string]interface{} {
	result := map[string]interface{}{}

	result["operator"] = string(obj.Operator)

	if obj.Value != nil {
		result["value"] = int(*obj.Value)
	}

	return result
}

func autoScalingConfigurationPolicyRulesHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if action, ok := m["action"]; ok {
		if tmpList := action.([]interface{}); len(tmpList) > 0 {
			buf.WriteString("action-")
			actionRaw := tmpList[0].(map[string]interface{})
			if type_, ok := actionRaw["type"]; ok && type_ != "" {
				buf.WriteString(fmt.Sprintf("%v-", type_))
			}
			if value, ok := actionRaw["value"]; ok {
				buf.WriteString(fmt.Sprintf("%v-", value))
			}
		}
	}
	if displayName, ok := m["display_name"]; ok && displayName != "" {
		buf.WriteString(fmt.Sprintf("%v-", displayName))
	}
	if metric, ok := m["metric"]; ok {
		if tmpList := metric.([]interface{}); len(tmpList) > 0 {
			buf.WriteString("metric-")
			metricRaw := tmpList[0].(map[string]interface{})
			if metricType, ok := metricRaw["metric_type"]; ok && metricType != "" {
				buf.WriteString(fmt.Sprintf("%v-", metricType))
			}
			if threshold, ok := metricRaw["threshold"]; ok {
				if tmpList := threshold.([]interface{}); len(tmpList) > 0 {
					buf.WriteString("threshold-")
					thresholdRaw := tmpList[0].(map[string]interface{})
					if operator, ok := thresholdRaw["operator"]; ok && operator != "" {
						buf.WriteString(fmt.Sprintf("%v-", operator))
					}
					if value, ok := thresholdRaw["value"]; ok {
						buf.WriteString(fmt.Sprintf("%v-", value))
					}
				}
			}
		}
	}
	return hashcode.String(buf.String())
}
