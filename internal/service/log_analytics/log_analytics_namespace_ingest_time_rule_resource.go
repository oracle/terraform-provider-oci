// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LogAnalyticsNamespaceIngestTimeRuleResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createLogAnalyticsNamespaceIngestTimeRule,
		Read:     readLogAnalyticsNamespaceIngestTimeRule,
		Update:   updateLogAnalyticsNamespaceIngestTimeRule,
		Delete:   deleteLogAnalyticsNamespaceIngestTimeRule,
		Schema: map[string]*schema.Schema{
			// Required
			"actions": {
				Type:     schema.TypeSet,
				Required: true,
				Set:      ingestTimeRuleActionsHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"compartment_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"metric_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"METRIC_EXTRACTION",
							}, true),
						},

						// Optional
						"dimensions": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Set:      tfresource.LiteralTypeHashCodeForSets,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"resource_group": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"conditions": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"field_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"field_operator": {
							Type:     schema.TypeString,
							Required: true,
						},
						"field_value": {
							Type:     schema.TypeString,
							Required: true,
						},
						"kind": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"FIELD",
							}, true),
						},

						// Optional
						"additional_conditions": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Set:      ingestTimeRuleAdditionalConditionsHashCodeForSets,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"condition_field": {
										Type:     schema.TypeString,
										Required: true,
									},
									"condition_operator": {
										Type:     schema.TypeString,
										Required: true,
									},
									"condition_value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"is_enabled": {
				Type:     schema.TypeBool,
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ingest_time_rule_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createLogAnalyticsNamespaceIngestTimeRule(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceIngestTimeRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.CreateResource(d, sync)
}

func readLogAnalyticsNamespaceIngestTimeRule(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceIngestTimeRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

func updateLogAnalyticsNamespaceIngestTimeRule(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceIngestTimeRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteLogAnalyticsNamespaceIngestTimeRule(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceIngestTimeRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type LogAnalyticsNamespaceIngestTimeRuleResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_log_analytics.LogAnalyticsClient
	Res                    *oci_log_analytics.IngestTimeRule
	DisableNotFoundRetries bool
}

func (s *LogAnalyticsNamespaceIngestTimeRuleResourceCrud) ID() string {
	return GetNamespaceIngestTimeRuleCompositeId(*s.Res.Id, s.D.Get("namespace").(string))
}

func (s *LogAnalyticsNamespaceIngestTimeRuleResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *LogAnalyticsNamespaceIngestTimeRuleResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_log_analytics.ConfigLifecycleStateActive),
	}
}

func (s *LogAnalyticsNamespaceIngestTimeRuleResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *LogAnalyticsNamespaceIngestTimeRuleResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_log_analytics.ConfigLifecycleStateDeleted),
	}
}

func (s *LogAnalyticsNamespaceIngestTimeRuleResourceCrud) Create() error {
	request := oci_log_analytics.CreateIngestTimeRuleRequest{}

	if actions, ok := s.D.GetOkExists("actions"); ok {
		interfaces := actions.(*schema.Set).List()
		tmp := make([]oci_log_analytics.IngestTimeRuleAction, len(interfaces))
		for i := range interfaces {
			stateDataIndex := ingestTimeRuleActionsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "actions", stateDataIndex)
			converted, err := s.mapToIngestTimeRuleAction(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("actions") {
			request.Actions = tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if conditions, ok := s.D.GetOkExists("conditions"); ok {
		if tmpList := conditions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "conditions", 0)
			tmp, err := s.mapToIngestTimeRuleCondition(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Conditions = tmp
		}
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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.CreateIngestTimeRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IngestTimeRule
	return nil
}

func (s *LogAnalyticsNamespaceIngestTimeRuleResourceCrud) Get() error {
	request := oci_log_analytics.GetIngestTimeRuleRequest{}

	if ingestTimeRuleId, ok := s.D.GetOkExists("ingest_time_rule_id"); ok {
		tmp := ingestTimeRuleId.(string)
		request.IngestTimeRuleId = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	ingestTimeRuleId, namespace, err := parseNamespaceIngestTimeRuleCompositeId(s.D.Id())
	if err == nil {
		request.IngestTimeRuleId = &ingestTimeRuleId
		request.NamespaceName = &namespace
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.GetIngestTimeRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IngestTimeRule
	return nil
}

func (s *LogAnalyticsNamespaceIngestTimeRuleResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_log_analytics.UpdateIngestTimeRuleRequest{}

	if actions, ok := s.D.GetOkExists("actions"); ok {
		interfaces := actions.(*schema.Set).List()
		tmp := make([]oci_log_analytics.IngestTimeRuleAction, len(interfaces))
		for i := range interfaces {
			stateDataIndex := ingestTimeRuleActionsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "actions", stateDataIndex)
			converted, err := s.mapToIngestTimeRuleAction(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("actions") {
			request.UpdateIngestTimeRuleDetails.Actions = tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.UpdateIngestTimeRuleDetails.CompartmentId = &tmp
	}

	if conditions, ok := s.D.GetOkExists("conditions"); ok {
		if tmpList := conditions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "conditions", 0)
			tmp, err := s.mapToIngestTimeRuleCondition(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UpdateIngestTimeRuleDetails.Conditions = tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.UpdateIngestTimeRuleDetails.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.UpdateIngestTimeRuleDetails.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.UpdateIngestTimeRuleDetails.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.UpdateIngestTimeRuleDetails.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.UpdateIngestTimeRuleDetails.Id = &tmp

	if ingestTimeRuleId, ok := s.D.GetOkExists("ingest_time_rule_id"); ok {
		tmp := ingestTimeRuleId.(string)
		request.IngestTimeRuleId = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.UpdateIngestTimeRuleDetails.IsEnabled = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.UpdateIngestTimeRuleDetails.LifecycleState = oci_log_analytics.ConfigLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	response, err := s.Client.UpdateIngestTimeRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IngestTimeRule
	return nil
}

func (s *LogAnalyticsNamespaceIngestTimeRuleResourceCrud) Delete() error {
	request := oci_log_analytics.DeleteIngestTimeRuleRequest{}

	if ingestTimeRuleId, ok := s.D.GetOkExists("ingest_time_rule_id"); ok {
		tmp := ingestTimeRuleId.(string)
		request.IngestTimeRuleId = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	_, err := s.Client.DeleteIngestTimeRule(context.Background(), request)
	return err
}

func (s *LogAnalyticsNamespaceIngestTimeRuleResourceCrud) SetData() error {

	ingestTimeRuleId, namespace, err := parseNamespaceIngestTimeRuleCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(ingestTimeRuleId)
		s.D.Set("ingest_time_rule_id", &ingestTimeRuleId)
		s.D.Set("namespace", &namespace)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	actions := []interface{}{}
	for _, item := range s.Res.Actions {
		actions = append(actions, IngestTimeRuleActionToMap(item, false))
	}
	s.D.Set("actions", schema.NewSet(ingestTimeRuleActionsHashCodeForSets, actions))

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Conditions != nil {
		conditionsArray := []interface{}{}
		if conditionsMap := IngestTimeRuleConditionToMap(&s.Res.Conditions, false); conditionsMap != nil {
			conditionsArray = append(conditionsArray, conditionsMap)
		}
		s.D.Set("conditions", conditionsArray)
	} else {
		s.D.Set("conditions", nil)
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

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func ingestTimeRuleActionsHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})

	if compartmentId, ok := m["compartment_id"]; ok && compartmentId != "" {
		buf.WriteString(fmt.Sprintf("%v-", compartmentId))
	}

	if metricName, ok := m["metric_name"]; ok && metricName != "" {
		buf.WriteString(fmt.Sprintf("%v-", metricName))
	}

	if namespace, ok := m["namespace"]; ok && namespace != "" {
		buf.WriteString(fmt.Sprintf("%v-", namespace))
	}

	if actionsType, ok := m["type"]; ok && actionsType != "" {
		buf.WriteString(fmt.Sprintf("%v-", actionsType))
	}

	if resourceGroup, ok := m["resource_group"]; ok && resourceGroup != "" {
		buf.WriteString(fmt.Sprintf("%v-", resourceGroup))
	}

	if dimensions, ok := m["dimensions"]; ok && dimensions != "" {
		buf.WriteString(fmt.Sprintf("%v-", dimensions))
	}

	return utils.GetStringHashcode(buf.String())
}

func ingestTimeRuleAdditionalConditionsHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})

	if conditionField, ok := m["condition_field"]; ok && conditionField != "" {
		buf.WriteString(fmt.Sprintf("%v-", conditionField))
	}

	if conditionOperator, ok := m["condition_operator"]; ok && conditionOperator != "" {
		buf.WriteString(fmt.Sprintf("%v-", conditionOperator))
	}

	if conditionValue, ok := m["condition_value"]; ok && conditionValue != "" {
		buf.WriteString(fmt.Sprintf("%v-", conditionValue))
	}

	return utils.GetStringHashcode(buf.String())
}

func GetNamespaceIngestTimeRuleCompositeId(ingestTimeRuleId string, namespace string) string {
	ingestTimeRuleId = url.PathEscape(ingestTimeRuleId)
	namespace = url.PathEscape(namespace)
	compositeId := "namespaces/" + namespace + "/ingestTimeRules/" + ingestTimeRuleId
	return compositeId
}

func parseNamespaceIngestTimeRuleCompositeId(compositeId string) (ingestTimeRuleId string, namespace string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("namespaces/.*/ingestTimeRules/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	namespace, _ = url.PathUnescape(parts[1])
	ingestTimeRuleId, _ = url.PathUnescape(parts[3])

	return
}

func (s *LogAnalyticsNamespaceIngestTimeRuleResourceCrud) mapToIngestTimeRuleAction(fieldKeyFormat string) (oci_log_analytics.IngestTimeRuleAction, error) {
	var baseObject oci_log_analytics.IngestTimeRuleAction
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("METRIC_EXTRACTION"):
		details := oci_log_analytics.IngestTimeRuleMetricExtractionAction{}
		if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if dimensions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dimensions")); ok {
			interfaces := dimensions.(*schema.Set).List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "dimensions")) {
				details.Dimensions = tmp
			}
		}
		if metricName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metric_name")); ok {
			tmp := metricName.(string)
			details.MetricName = &tmp
		}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if resourceGroup, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_group")); ok {
			tmp := resourceGroup.(string)
			details.ResourceGroup = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func IngestTimeRuleActionToMap(obj oci_log_analytics.IngestTimeRuleAction, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_log_analytics.IngestTimeRuleMetricExtractionAction:
		result["type"] = "METRIC_EXTRACTION"

		if v.CompartmentId != nil {
			result["compartment_id"] = string(*v.CompartmentId)
		}

		dimensions := []interface{}{}
		for _, item := range v.Dimensions {
			dimensions = append(dimensions, item)
		}
		if datasource {
			result["dimensions"] = dimensions
		} else {
			result["dimensions"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, dimensions)
		}

		if v.MetricName != nil {
			result["metric_name"] = string(*v.MetricName)
		}

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.ResourceGroup != nil {
			result["resource_group"] = string(*v.ResourceGroup)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *LogAnalyticsNamespaceIngestTimeRuleResourceCrud) mapToIngestTimeRuleAdditionalFieldCondition(fieldKeyFormat string) (oci_log_analytics.IngestTimeRuleAdditionalFieldCondition, error) {
	result := oci_log_analytics.IngestTimeRuleAdditionalFieldCondition{}

	if conditionField, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "condition_field")); ok {
		tmp := conditionField.(string)
		result.ConditionField = &tmp
	}

	if conditionOperator, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "condition_operator")); ok {
		result.ConditionOperator = oci_log_analytics.IngestTimeRuleAdditionalFieldConditionConditionOperatorEnum(conditionOperator.(string))
	}

	if conditionValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "condition_value")); ok {
		tmp := conditionValue.(string)
		result.ConditionValue = &tmp
	}

	return result, nil
}

func IngestTimeRuleAdditionalFieldConditionToMap(obj oci_log_analytics.IngestTimeRuleAdditionalFieldCondition) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConditionField != nil {
		result["condition_field"] = string(*obj.ConditionField)
	}

	result["condition_operator"] = string(obj.ConditionOperator)

	if obj.ConditionValue != nil {
		result["condition_value"] = string(*obj.ConditionValue)
	}

	return result
}

func (s *LogAnalyticsNamespaceIngestTimeRuleResourceCrud) mapToIngestTimeRuleCondition(fieldKeyFormat string) (oci_log_analytics.IngestTimeRuleCondition, error) {
	var baseObject oci_log_analytics.IngestTimeRuleCondition
	//discriminator
	kindRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kind"))
	var kind string
	if ok {
		kind = kindRaw.(string)
	} else {
		kind = "" // default value
	}
	switch strings.ToLower(kind) {
	case strings.ToLower("FIELD"):
		details := oci_log_analytics.IngestTimeRuleFieldCondition{}
		if additionalConditions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "additional_conditions")); ok {
			interfaces := additionalConditions.(*schema.Set).List()
			tmp := make([]oci_log_analytics.IngestTimeRuleAdditionalFieldCondition, len(interfaces))
			for i := range interfaces {
				stateDataIndex := ingestTimeRuleAdditionalConditionsHashCodeForSets(interfaces[i])
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "additional_conditions"), stateDataIndex)
				converted, err := s.mapToIngestTimeRuleAdditionalFieldCondition(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "additional_conditions")) {
				details.AdditionalConditions = tmp
			}
		}
		if fieldName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_name")); ok {
			tmp := fieldName.(string)
			details.FieldName = &tmp
		}
		if fieldOperator, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_operator")); ok {
			details.FieldOperator = oci_log_analytics.IngestTimeRuleFieldConditionFieldOperatorEnum(fieldOperator.(string))
		}
		if fieldValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_value")); ok {
			tmp := fieldValue.(string)
			details.FieldValue = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown kind '%v' was specified", kind)
	}
	return baseObject, nil
}

func IngestTimeRuleConditionToMap(obj *oci_log_analytics.IngestTimeRuleCondition, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_log_analytics.IngestTimeRuleFieldCondition:
		result["kind"] = "FIELD"

		additionalConditions := []interface{}{}
		for _, item := range v.AdditionalConditions {
			additionalConditions = append(additionalConditions, IngestTimeRuleAdditionalFieldConditionToMap(item))
		}
		if datasource {
			result["additional_conditions"] = additionalConditions
		} else {
			result["additional_conditions"] = schema.NewSet(ingestTimeRuleAdditionalConditionsHashCodeForSets, additionalConditions)
		}

		if v.FieldName != nil {
			result["field_name"] = string(*v.FieldName)
		}

		result["field_operator"] = string(v.FieldOperator)

		if v.FieldValue != nil {
			result["field_value"] = string(*v.FieldValue)
		}
	default:
		log.Printf("[WARN] Received 'kind' of unknown type %v", *obj)
		return nil
	}

	return result
}

func IngestTimeRuleSummaryToMap(obj oci_log_analytics.IngestTimeRuleSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["condition_kind"] = string(obj.ConditionKind)

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.FieldName != nil {
		result["field_name"] = string(*obj.FieldName)
	}

	if obj.FieldValue != nil {
		result["field_value"] = string(*obj.FieldValue)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *LogAnalyticsNamespaceIngestTimeRuleResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_log_analytics.ChangeIngestTimeRuleCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	if ingestTimeRuleId, ok := s.D.GetOkExists("ingest_time_rule_id"); ok {
		tmp := ingestTimeRuleId.(string)
		changeCompartmentRequest.IngestTimeRuleId = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		changeCompartmentRequest.NamespaceName = &tmp
	}

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "log_analytics")

	_, err := s.Client.ChangeIngestTimeRuleCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
