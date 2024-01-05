// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package load_balancer

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

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_load_balancer "github.com/oracle/oci-go-sdk/v65/loadbalancer"
)

func LoadBalancerRuleSetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
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
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"ADD_HTTP_REQUEST_HEADER",
								"ADD_HTTP_RESPONSE_HEADER",
								"ALLOW",
								"CONTROL_ACCESS_USING_HTTP_METHODS",
								"EXTEND_HTTP_REQUEST_HEADER_VALUE",
								"EXTEND_HTTP_RESPONSE_HEADER_VALUE",
								"HTTP_HEADER",
								"REDIRECT",
								"REMOVE_HTTP_REQUEST_HEADER",
								"REMOVE_HTTP_RESPONSE_HEADER",
							}, true),
						},

						// Optional
						"allowed_methods": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Set:      tfresource.LiteralTypeHashCodeForSets,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"are_invalid_characters_allowed": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
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
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"PATH",
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
									"operator": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

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
						"http_large_header_size_in_kb": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"prefix": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"redirect_uri": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"host": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"path": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"port": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"protocol": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"query": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"response_code": {
							Type:     schema.TypeInt,
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
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.CreateResource(d, sync)
}

func readLoadBalancerRuleSet(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerRuleSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.ReadResource(sync)
}

func updateLoadBalancerRuleSet(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerRuleSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteLoadBalancerRuleSet(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerRuleSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type LoadBalancerRuleSetResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_load_balancer.LoadBalancerClient
	Res                    *oci_load_balancer.RuleSet
	DisableNotFoundRetries bool
	WorkRequest            *oci_load_balancer.WorkRequest
}

func (s *LoadBalancerRuleSetResourceCrud) ID() string {
	if s.WorkRequest != nil {
		if s.WorkRequest.LifecycleState == oci_load_balancer.WorkRequestLifecycleStateSucceeded {
			return GetRuleSetCompositeId(s.D.Get("load_balancer_id").(string), s.D.Get("name").(string))
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
		if len(tmp) != 0 || s.D.HasChange("items") {
			request.Items = tmp
		}
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.CreateRuleSet(context.Background(), request)
	if err != nil {
		return err
	}

	var compositeId string
	compositeId = GetRuleSetCompositeId(s.D.Get("load_balancer_id").(string), s.D.Get("name").(string))
	s.D.SetId(compositeId)
	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	err = loadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	return nil
}

func (s *LoadBalancerRuleSetResourceCrud) Get() error {
	_, stillWorking, err := loadBalancerResourceGet(s.Client, s.D, s.WorkRequest, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.GetRuleSet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.RuleSet
	return nil
}

func (s *LoadBalancerRuleSetResourceCrud) Update() error {
	request := oci_load_balancer.UpdateRuleSetRequest{}

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
		if len(tmp) != 0 || s.D.HasChange("items") {
			request.Items = tmp
		}
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.RuleSetName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.UpdateRuleSet(context.Background(), request)
	if err != nil {
		return err
	}

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	err = loadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.DeleteRuleSet(context.Background(), request)
	if err != nil {
		return err
	}

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	err = loadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
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
		items = append(items, RuleToMap(item, false))
	}
	s.D.Set("items", schema.NewSet(itemsHashCodeForSets, items))

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	return nil
}

func GetRuleSetCompositeId(loadBalancerId string, name string) string {
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

func (s *LoadBalancerRuleSetResourceCrud) mapToRedirectUri(fieldKeyFormat string) (oci_load_balancer.RedirectUri, error) {
	result := oci_load_balancer.RedirectUri{}

	if host, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "host")); ok {
		tmp := host.(string)
		result.Host = &tmp
	}

	if path, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "path")); ok {
		tmp := path.(string)
		result.Path = &tmp
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(int)
		result.Port = &tmp
	}

	if protocol, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protocol")); ok {
		tmp := protocol.(string)
		result.Protocol = &tmp
	}

	if query, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "query")); ok {
		tmp := query.(string)
		result.Query = &tmp
	}

	return result, nil
}

func RedirectUriToMap(obj *oci_load_balancer.RedirectUri) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Host != nil {
		result["host"] = string(*obj.Host)
	}

	if obj.Path != nil {
		result["path"] = string(*obj.Path)
	}

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	if obj.Protocol != nil {
		result["protocol"] = string(*obj.Protocol)
	}

	if obj.Query != nil {
		result["query"] = string(*obj.Query)
	}

	return result
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
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "conditions")) {
				details.Conditions = tmp
			}
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		baseObject = details
	case strings.ToLower("CONTROL_ACCESS_USING_HTTP_METHODS"):
		details := oci_load_balancer.ControlAccessUsingHttpMethodsRule{}
		if allowedMethods, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allowed_methods")); ok {
			tmpInterfaces := allowedMethods.(*schema.Set)
			interfaces := tmpInterfaces.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "allowed_methods")) {
				details.AllowedMethods = tmp
			}
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
	case strings.ToLower("HTTP_HEADER"):
		details := oci_load_balancer.HttpHeaderRule{}
		if areInvalidCharactersAllowed, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "are_invalid_characters_allowed")); ok {
			tmp := areInvalidCharactersAllowed.(bool)
			details.AreInvalidCharactersAllowed = &tmp
		}
		if httpLargeHeaderSizeInKB, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "http_large_header_size_in_kb")); ok {
			tmp := httpLargeHeaderSizeInKB.(int)
			details.HttpLargeHeaderSizeInKB = &tmp
		}
		baseObject = details
	case strings.ToLower("REDIRECT"):
		details := oci_load_balancer.RedirectRule{}
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
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "conditions")) {
				details.Conditions = tmp
			}
		}
		if redirectUri, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "redirect_uri")); ok {
			if tmpList := redirectUri.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "redirect_uri"), 0)
				tmp, err := s.mapToRedirectUri(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert redirect_uri, encountered error: %v", err)
				}
				details.RedirectUri = &tmp
			}
		}
		if responseCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "response_code")); ok {
			tmp := responseCode.(int)
			details.ResponseCode = &tmp
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

func RuleToMap(obj oci_load_balancer.Rule, datasource bool) map[string]interface{} {
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

		if datasource {
			result["allowed_methods"] = v.AllowedMethods
		} else {
			if v.AllowedMethods != nil {
				allowedMethods := []interface{}{}
				for _, item := range v.AllowedMethods {
					allowedMethods = append(allowedMethods, item)
				}
				result["allowed_methods"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, allowedMethods)
			}
		}

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
	case oci_load_balancer.HttpHeaderRule:
		result["action"] = "HTTP_HEADER"

		if v.AreInvalidCharactersAllowed != nil {
			result["are_invalid_characters_allowed"] = bool(*v.AreInvalidCharactersAllowed)
		}

		if v.HttpLargeHeaderSizeInKB != nil {
			result["http_large_header_size_in_kb"] = int(*v.HttpLargeHeaderSizeInKB)
		}
	case oci_load_balancer.RedirectRule:
		result["action"] = "REDIRECT"

		conditions := []interface{}{}
		for _, item := range v.Conditions {
			conditions = append(conditions, RuleConditionToMap(item))
		}
		result["conditions"] = conditions

		if v.RedirectUri != nil {
			result["redirect_uri"] = []interface{}{RedirectUriToMap(v.RedirectUri)}
		}

		if v.ResponseCode != nil {
			result["response_code"] = int(*v.ResponseCode)
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
	case strings.ToLower("PATH"):
		details := oci_load_balancer.PathMatchCondition{}
		if attributeValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "attribute_value")); ok {
			tmp := attributeValue.(string)
			details.AttributeValue = &tmp
		}
		if operator, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operator")); ok {
			details.Operator = oci_load_balancer.PathMatchConditionOperatorEnum(operator.(string))
		}
		baseObject = details
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
	case oci_load_balancer.PathMatchCondition:
		result["attribute_name"] = "PATH"

		if v.AttributeValue != nil {
			result["attribute_value"] = string(*v.AttributeValue)
		}

		result["operator"] = string(v.Operator)
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
			if allowedMethods, ok := m["allowed_methods"]; ok && allowedMethods != "" && allowedMethods != nil {
				buf.WriteString(fmt.Sprintf("%v-", allowedMethods.(*schema.Set).List()))
			}
		}
	}
	if areInvalidCharactersAllowed, ok := m["are_invalid_characters_allowed"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", areInvalidCharactersAllowed))
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
	if httpLargeHeaderSizeInKB, ok := m["http_large_header_size_in_kb"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", httpLargeHeaderSizeInKB))
	}
	if prefix, ok := m["prefix"]; ok && prefix != "" {
		buf.WriteString(fmt.Sprintf("%v-", prefix))
	}
	if redirectUri, ok := m["redirect_uri"]; ok {
		if tmpList := redirectUri.([]interface{}); len(tmpList) > 0 {
			buf.WriteString("redirect_uri-")
			redirectUriRaw := tmpList[0].(map[string]interface{})
			if host, ok := redirectUriRaw["host"]; ok && host != "" {
				buf.WriteString(fmt.Sprintf("%v-", host))
			}
			if path, ok := redirectUriRaw["path"]; ok && path != "" {
				buf.WriteString(fmt.Sprintf("%v-", path))
			}
			if port, ok := redirectUriRaw["port"]; ok {
				buf.WriteString(fmt.Sprintf("%v-", port))
			}
			if protocol, ok := redirectUriRaw["protocol"]; ok && protocol != "" {
				buf.WriteString(fmt.Sprintf("%v-", protocol))
			}
			if query, ok := redirectUriRaw["query"]; ok && query != "" {
				buf.WriteString(fmt.Sprintf("%v-", query))
			}
		}
	}
	if responseCode, ok := m["response_code"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", responseCode))
	}
	if suffix, ok := m["suffix"]; ok && suffix != "" {
		buf.WriteString(fmt.Sprintf("%v-", suffix))
	}
	if value, ok := m["value"]; ok && value != "" {
		buf.WriteString(fmt.Sprintf("%v-", value))
	}
	return utils.GetStringHashcode(buf.String())
}
