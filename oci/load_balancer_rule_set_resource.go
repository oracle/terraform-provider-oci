// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"
)

func LoadBalancerRuleSetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createLoadBalancerRuleSet,
		Read:     readLoadBalancerRuleSet,
		Update:   updateLoadBalancerRuleSet,
		Delete:   deleteLoadBalancerRuleSet,
		Schema: map[string]*schema.Schema{
			// Required
			"items": {
				Type:     schema.TypeSet,
				Required: true,
				Set:      itemsHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"action": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"ADD_HTTP_REQUEST_HEADER",
								"ADD_HTTP_RESPONSE_HEADER",
								"ALLOW",
								"CONTROL_ACCESS_USING_HTTP_METHODS",
								"EXTEND_HTTP_REQUEST_HEADER_VALUE",
								"EXTEND_HTTP_RESPONSE_HEADER_VALUE",
								"REMOVE_HTTP_REQUEST_HEADER",
								"REMOVE_HTTP_RESPONSE_HEADER",
							}, true),
						},

						// Optional
						"allowed_methods": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"conditions": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"attribute_name": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"SOURCE_IP_ADDRESS",
											"SOURCE_VCN_ID",
											"SOURCE_VCN_IP_ADDRESS",
										}, true),
									},
									"attribute_value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"header": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"prefix": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status_code": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"suffix": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			// internal for work request access
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createLoadBalancerRuleSet(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerRuleSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return CreateResource(d, sync)
}

func readLoadBalancerRuleSet(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerRuleSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return ReadResource(sync)
}

func updateLoadBalancerRuleSet(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerRuleSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return UpdateResource(d, sync)
}

func deleteLoadBalancerRuleSet(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerRuleSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type LoadBalancerRuleSetResourceCrud struct {
	BaseCrud
	Client                 *oci_load_balancer.LoadBalancerClient
	Res                    *oci_load_balancer.RuleSet
	DisableNotFoundRetries bool
	WorkRequest            *oci_load_balancer.WorkRequest
}

func (s *LoadBalancerRuleSetResourceCrud) ID() string {
	if s.WorkRequest != nil {
		if s.WorkRequest.LifecycleState == oci_load_balancer.WorkRequestLifecycleStateSucceeded {
			return getRuleSetCompositeId(s.D.Get("load_balancer_id").(string), s.D.Get("name").(string))
		} else {
			return *s.WorkRequest.Id
		}
	}
	return ""
}

func (s *LoadBalancerRuleSetResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *LoadBalancerRuleSetResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *LoadBalancerRuleSetResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *LoadBalancerRuleSetResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *LoadBalancerRuleSetResourceCrud) Create() error {
	request := oci_load_balancer.CreateRuleSetRequest{}

	request.Items = []oci_load_balancer.Rule{}
	if items, ok := s.D.GetOkExists("items"); ok {
		set := items.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_load_balancer.Rule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := itemsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "items", stateDataIndex)
			converted, err := s.mapToRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		request.Items = tmp
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.CreateRuleSet(context.Background(), request)
	if err != nil {
		return err
	}

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	err = LoadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, getRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	return nil
}

func (s *LoadBalancerRuleSetResourceCrud) Get() error {
	_, stillWorking, err := LoadBalancerResourceGet(s.Client, s.D, s.WorkRequest, getRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	if stillWorking {
		return nil
	}
	request := oci_load_balancer.GetRuleSetRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.RuleSetName = &tmp
	}

	if !strings.HasPrefix(s.D.Id(), "ocid1.loadbalancerworkrequest.") {
		loadBalancerId, name, err := parseRuleSetCompositeId(s.D.Id())
		if err == nil {
			request.LoadBalancerId = &loadBalancerId
			request.RuleSetName = &name
		} else {
			log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
		}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.GetRuleSet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.RuleSet
	return nil
}

func (s *LoadBalancerRuleSetResourceCrud) Update() error {
	request := oci_load_balancer.UpdateRuleSetRequest{}

	request.Items = []oci_load_balancer.Rule{}
	if items, ok := s.D.GetOkExists("items"); ok {
		set := items.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_load_balancer.Rule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := itemsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "items", stateDataIndex)
			converted, err := s.mapToRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		request.Items = tmp
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.RuleSetName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.UpdateRuleSet(context.Background(), request)
	if err != nil {
		return err
	}

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	err = LoadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, getRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}

	return s.Get()
}

func (s *LoadBalancerRuleSetResourceCrud) Delete() error {
	request := oci_load_balancer.DeleteRuleSetRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.RuleSetName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.DeleteRuleSet(context.Background(), request)
	if err != nil {
		return err
	}

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	err = LoadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, getRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	return nil
}

func (s *LoadBalancerRuleSetResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	loadBalancerId, name, err := parseRuleSetCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("load_balancer_id", &loadBalancerId)
		s.D.Set("name", &name)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RuleToMap(item))
	}
	s.D.Set("items", schema.NewSet(itemsHashCodeForSets, items))

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	return nil
}

func getRuleSetCompositeId(loadBalancerId string, name string) string {
	loadBalancerId = url.PathEscape(loadBalancerId)
	name = url.PathEscape(name)
	compositeId := "loadBalancers/" + loadBalancerId + "/ruleSets/" + name
	return compositeId
}

func parseRuleSetCompositeId(compositeId string) (loadBalancerId string, name string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("loadBalancers/.*/ruleSets/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	loadBalancerId, _ = url.PathUnescape(parts[1])
	name, _ = url.PathUnescape(parts[3])

	return
}

func (s *LoadBalancerRuleSetResourceCrud) mapToRule(fieldKeyFormat string) (oci_load_balancer.Rule, error) {
	var baseObject oci_load_balancer.Rule
	//discriminator
	actionRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action"))
	var action string
	if ok {
		action = actionRaw.(string)
	} else {
		action = "" // default value
	}
	switch strings.ToLower(action) {
	case strings.ToLower("ADD_HTTP_REQUEST_HEADER"):
		details := oci_load_balancer.AddHttpRequestHeaderRule{}
		if header, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "header")); ok {
			tmp := header.(string)
			details.Header = &tmp
		}
		if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
			tmp := value.(string)
			details.Value = &tmp
		}
		baseObject = details
	case strings.ToLower("ADD_HTTP_RESPONSE_HEADER"):
		details := oci_load_balancer.AddHttpResponseHeaderRule{}
		if header, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "header")); ok {
			tmp := header.(string)
			details.Header = &tmp
		}
		if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
			tmp := value.(string)
			details.Value = &tmp
		}
		baseObject = details
	case strings.ToLower("ALLOW"):
		details := oci_load_balancer.AllowRule{}
		details.Conditions = []oci_load_balancer.RuleCondition{}
		if conditions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "conditions")); ok {
			interfaces := conditions.([]interface{})
			tmp := make([]oci_load_balancer.RuleCondition, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "conditions"), stateDataIndex)
				converted, err := s.mapToRuleCondition(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			details.Conditions = tmp
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		baseObject = details
	case strings.ToLower("CONTROL_ACCESS_USING_HTTP_METHODS"):
		details := oci_load_balancer.ControlAccessUsingHttpMethodsRule{}
		details.AllowedMethods = []string{}
		if allowedMethods, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allowed_methods")); ok {
			interfaces := allowedMethods.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			details.AllowedMethods = tmp
		}
		if statusCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "status_code")); ok {
			tmp := statusCode.(int)
			details.StatusCode = &tmp
		}
		baseObject = details
	case strings.ToLower("EXTEND_HTTP_REQUEST_HEADER_VALUE"):
		details := oci_load_balancer.ExtendHttpRequestHeaderValueRule{}
		if header, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "header")); ok {
			tmp := header.(string)
			details.Header = &tmp
		}
		if prefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "prefix")); ok {
			tmp := prefix.(string)
			details.Prefix = &tmp
		}
		if suffix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "suffix")); ok {
			tmp := suffix.(string)
			details.Suffix = &tmp
		}
		baseObject = details
	case strings.ToLower("EXTEND_HTTP_RESPONSE_HEADER_VALUE"):
		details := oci_load_balancer.ExtendHttpResponseHeaderValueRule{}
		if header, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "header")); ok {
			tmp := header.(string)
			details.Header = &tmp
		}
		if prefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "prefix")); ok {
			tmp := prefix.(string)
			details.Prefix = &tmp
		}
		if suffix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "suffix")); ok {
			tmp := suffix.(string)
			details.Suffix = &tmp
		}
		baseObject = details
	case strings.ToLower("REMOVE_HTTP_REQUEST_HEADER"):
		details := oci_load_balancer.RemoveHttpRequestHeaderRule{}
		if header, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "header")); ok {
			tmp := header.(string)
			details.Header = &tmp
		}
		baseObject = details
	case strings.ToLower("REMOVE_HTTP_RESPONSE_HEADER"):
		details := oci_load_balancer.RemoveHttpResponseHeaderRule{}
		if header, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "header")); ok {
			tmp := header.(string)
			details.Header = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown action '%v' was specified", action)
	}
	return baseObject, nil
}

func RuleToMap(obj oci_load_balancer.Rule) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_load_balancer.AddHttpRequestHeaderRule:
		result["action"] = "ADD_HTTP_REQUEST_HEADER"

		if v.Header != nil {
			result["header"] = string(*v.Header)
		}

		if v.Value != nil {
			result["value"] = string(*v.Value)
		}
	case oci_load_balancer.AddHttpResponseHeaderRule:
		result["action"] = "ADD_HTTP_RESPONSE_HEADER"

		if v.Header != nil {
			result["header"] = string(*v.Header)
		}

		if v.Value != nil {
			result["value"] = string(*v.Value)
		}
	case oci_load_balancer.AllowRule:
		result["action"] = "ALLOW"

		conditions := []interface{}{}
		for _, item := range v.Conditions {
			conditions = append(conditions, RuleConditionToMap(item))
		}
		result["conditions"] = conditions

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}
	case oci_load_balancer.ControlAccessUsingHttpMethodsRule:
		result["action"] = "CONTROL_ACCESS_USING_HTTP_METHODS"

		result["allowed_methods"] = v.AllowedMethods

		if v.StatusCode != nil {
			result["status_code"] = int(*v.StatusCode)
		}
	case oci_load_balancer.ExtendHttpRequestHeaderValueRule:
		result["action"] = "EXTEND_HTTP_REQUEST_HEADER_VALUE"

		if v.Header != nil {
			result["header"] = string(*v.Header)
		}

		if v.Prefix != nil {
			result["prefix"] = string(*v.Prefix)
		}

		if v.Suffix != nil {
			result["suffix"] = string(*v.Suffix)
		}
	case oci_load_balancer.ExtendHttpResponseHeaderValueRule:
		result["action"] = "EXTEND_HTTP_RESPONSE_HEADER_VALUE"

		if v.Header != nil {
			result["header"] = string(*v.Header)
		}

		if v.Prefix != nil {
			result["prefix"] = string(*v.Prefix)
		}

		if v.Suffix != nil {
			result["suffix"] = string(*v.Suffix)
		}
	case oci_load_balancer.RemoveHttpRequestHeaderRule:
		result["action"] = "REMOVE_HTTP_REQUEST_HEADER"

		if v.Header != nil {
			result["header"] = string(*v.Header)
		}
	case oci_load_balancer.RemoveHttpResponseHeaderRule:
		result["action"] = "REMOVE_HTTP_RESPONSE_HEADER"

		if v.Header != nil {
			result["header"] = string(*v.Header)
		}
	default:
		log.Printf("[WARN] Received 'action' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *LoadBalancerRuleSetResourceCrud) mapToRuleCondition(fieldKeyFormat string) (oci_load_balancer.RuleCondition, error) {
	var baseObject oci_load_balancer.RuleCondition
	//discriminator
	attributeNameRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "attribute_name"))
	var attributeName string
	if ok {
		attributeName = attributeNameRaw.(string)
	} else {
		attributeName = "" // default value
	}
	switch strings.ToLower(attributeName) {
	case strings.ToLower("SOURCE_IP_ADDRESS"):
		details := oci_load_balancer.SourceIpAddressCondition{}
		if attributeValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "attribute_value")); ok {
			tmp := attributeValue.(string)
			details.AttributeValue = &tmp
		}
		baseObject = details
	case strings.ToLower("SOURCE_VCN_ID"):
		details := oci_load_balancer.SourceVcnIdCondition{}
		if attributeValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "attribute_value")); ok {
			tmp := attributeValue.(string)
			details.AttributeValue = &tmp
		}
		baseObject = details
	case strings.ToLower("SOURCE_VCN_IP_ADDRESS"):
		details := oci_load_balancer.SourceVcnIpAddressCondition{}
		if attributeValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "attribute_value")); ok {
			tmp := attributeValue.(string)
			details.AttributeValue = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown attribute_name '%v' was specified", attributeName)
	}
	return baseObject, nil
}

func RuleConditionToMap(obj oci_load_balancer.RuleCondition) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_load_balancer.SourceIpAddressCondition:
		result["attribute_name"] = "SOURCE_IP_ADDRESS"

		if v.AttributeValue != nil {
			result["attribute_value"] = string(*v.AttributeValue)
		}
	case oci_load_balancer.SourceVcnIdCondition:
		result["attribute_name"] = "SOURCE_VCN_ID"

		if v.AttributeValue != nil {
			result["attribute_value"] = string(*v.AttributeValue)
		}
	case oci_load_balancer.SourceVcnIpAddressCondition:
		result["attribute_name"] = "SOURCE_VCN_IP_ADDRESS"

		if v.AttributeValue != nil {
			result["attribute_value"] = string(*v.AttributeValue)
		}
	default:
		log.Printf("[WARN] Received 'attribute_name' of unknown type %v", obj)
		return nil
	}

	return result
}

func itemsHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if action, ok := m["action"]; ok && action != "" {
		buf.WriteString(fmt.Sprintf("%v-", action))
		if action == "CONTROL_ACCESS_USING_HTTP_METHODS" {
			if statusCode, ok := m["status_code"]; ok && statusCode != 0 {
				buf.WriteString(fmt.Sprintf("%v-", statusCode))
			} else {
				buf.WriteString(fmt.Sprintf("%v-", 405))
			}
			if allowedMethods, ok := m["allowed_methods"]; ok && allowedMethods != "" {
				buf.WriteString(fmt.Sprintf("%v-", allowedMethods))
			}
		}
	}
	if conditions, ok := m["conditions"]; ok && conditions != nil {
		if tmpList := conditions.([]interface{}); len(tmpList) > 0 {
			for _, conditionsRaw := range tmpList {
				buf.WriteString("conditions-")
				if conditionsRaw != nil {
					tmpMap := conditionsRaw.(map[string]interface{})
					if name, ok := tmpMap["attribute_name"]; ok {
						buf.WriteString(fmt.Sprintf("%v-", name))
					}
					if value, ok := tmpMap["attribute_value"]; ok {
						buf.WriteString(fmt.Sprintf("%v-", value))
					}
				}
			}
		}
	}
	if description, ok := m["description"]; ok && description != "" {
		buf.WriteString(fmt.Sprintf("%v-", description))
	}
	if header, ok := m["header"]; ok && header != "" {
		buf.WriteString(fmt.Sprintf("%v-", header))
	}
	if prefix, ok := m["prefix"]; ok && prefix != "" {
		buf.WriteString(fmt.Sprintf("%v-", prefix))
	}
	if suffix, ok := m["suffix"]; ok && suffix != "" {
		buf.WriteString(fmt.Sprintf("%v-", suffix))
	}
	if value, ok := m["value"]; ok && value != "" {
		buf.WriteString(fmt.Sprintf("%v-", value))
	}
	return hashcode.String(buf.String())
}
