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
								"EXTEND_HTTP_REQUEST_HEADER_VALUE",
								"EXTEND_HTTP_RESPONSE_HEADER_VALUE",
								"REMOVE_HTTP_REQUEST_HEADER",
								"REMOVE_HTTP_RESPONSE_HEADER",
							}, true),
						},
						"header": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"prefix": {
							Type:     schema.TypeString,
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
	if strings.Contains(s.D.Id(), "ocid1.loadbalancerworkrequest") {
		return nil
	}
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

func itemsHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if action, ok := m["action"]; ok && action != "" {
		buf.WriteString(fmt.Sprintf("%v-", action))
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
