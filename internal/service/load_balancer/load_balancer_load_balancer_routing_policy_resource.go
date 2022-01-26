// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package load_balancer

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	oci_load_balancer "github.com/oracle/oci-go-sdk/v56/loadbalancer"
)

func LoadBalancerLoadBalancerRoutingPolicyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createLoadBalancerLoadBalancerRoutingPolicy,
		Read:     readLoadBalancerLoadBalancerRoutingPolicy,
		Update:   updateLoadBalancerLoadBalancerRoutingPolicy,
		Delete:   deleteLoadBalancerLoadBalancerRoutingPolicy,
		Schema: map[string]*schema.Schema{
			// Required
			"condition_language_version": {
				Type:     schema.TypeString,
				Required: true,
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
			"rules": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"actions": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"name": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"FORWARD_TO_BACKENDSET",
										}, true),
									},

									// Optional
									"backend_set_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"condition": {
							Type:     schema.TypeString,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
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

func createLoadBalancerLoadBalancerRoutingPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerLoadBalancerRoutingPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.CreateResource(d, sync)
}

func readLoadBalancerLoadBalancerRoutingPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerLoadBalancerRoutingPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.ReadResource(sync)
}

func updateLoadBalancerLoadBalancerRoutingPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerLoadBalancerRoutingPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteLoadBalancerLoadBalancerRoutingPolicy(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerLoadBalancerRoutingPolicyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type LoadBalancerLoadBalancerRoutingPolicyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_load_balancer.LoadBalancerClient
	Res                    *oci_load_balancer.RoutingPolicy
	DisableNotFoundRetries bool
	WorkRequest            *oci_load_balancer.WorkRequest
}

func (s *LoadBalancerLoadBalancerRoutingPolicyResourceCrud) ID() string {
	if s.WorkRequest != nil {
		if s.WorkRequest.LifecycleState == oci_load_balancer.WorkRequestLifecycleStateSucceeded {
			return GetLoadBalancerRoutingPolicyCompositeId(s.D.Get("load_balancer_id").(string), s.D.Get("name").(string))
		} else {
			return *s.WorkRequest.Id
		}
	}
	return ""
}

func (s *LoadBalancerLoadBalancerRoutingPolicyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *LoadBalancerLoadBalancerRoutingPolicyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *LoadBalancerLoadBalancerRoutingPolicyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *LoadBalancerLoadBalancerRoutingPolicyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *LoadBalancerLoadBalancerRoutingPolicyResourceCrud) Create() error {
	request := oci_load_balancer.CreateRoutingPolicyRequest{}

	if conditionLanguageVersion, ok := s.D.GetOkExists("condition_language_version"); ok {
		request.ConditionLanguageVersion = oci_load_balancer.CreateRoutingPolicyDetailsConditionLanguageVersionEnum(conditionLanguageVersion.(string))
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if rules, ok := s.D.GetOkExists("rules"); ok {
		interfaces := rules.([]interface{})
		tmp := make([]oci_load_balancer.RoutingRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rules", stateDataIndex)
			converted, err := s.mapToRoutingRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("rules") {
			request.Rules = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.CreateRoutingPolicy(context.Background(), request)
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
	err = tfresource.LoadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	return nil
}

func (s *LoadBalancerLoadBalancerRoutingPolicyResourceCrud) Get() error {
	_, stillWorking, err := tfresource.LoadBalancerResourceGet(s.Client, s.D, s.WorkRequest, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	if stillWorking {
		return nil
	}
	request := oci_load_balancer.GetRoutingPolicyRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if routingPolicyName, ok := s.D.GetOkExists("name"); ok {
		tmp := routingPolicyName.(string)
		request.RoutingPolicyName = &tmp
	}

	if !strings.HasPrefix(s.D.Id(), "ocid1.loadbalancerworkrequest.") {
		loadBalancerId, routingPolicyName, err := parseLoadBalancerRoutingPolicyCompositeId(s.D.Id())
		if err == nil {
			request.LoadBalancerId = &loadBalancerId
			request.RoutingPolicyName = &routingPolicyName
		} else {
			log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.GetRoutingPolicy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.RoutingPolicy
	return nil
}

func (s *LoadBalancerLoadBalancerRoutingPolicyResourceCrud) Update() error {
	request := oci_load_balancer.UpdateRoutingPolicyRequest{}

	if conditionLanguageVersion, ok := s.D.GetOkExists("condition_language_version"); ok {
		request.ConditionLanguageVersion = oci_load_balancer.UpdateRoutingPolicyDetailsConditionLanguageVersionEnum(conditionLanguageVersion.(string))
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if routingPolicyName, ok := s.D.GetOkExists("name"); ok {
		tmp := routingPolicyName.(string)
		request.RoutingPolicyName = &tmp
	}

	if rules, ok := s.D.GetOkExists("rules"); ok {
		interfaces := rules.([]interface{})
		tmp := make([]oci_load_balancer.RoutingRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rules", stateDataIndex)
			converted, err := s.mapToRoutingRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("rules") {
			request.Rules = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.UpdateRoutingPolicy(context.Background(), request)
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
	err = tfresource.LoadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}

	return s.Get()
}

func (s *LoadBalancerLoadBalancerRoutingPolicyResourceCrud) Delete() error {
	request := oci_load_balancer.DeleteRoutingPolicyRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if routingPolicyName, ok := s.D.GetOkExists("name"); ok {
		tmp := routingPolicyName.(string)
		request.RoutingPolicyName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.DeleteRoutingPolicy(context.Background(), request)
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
	err = tfresource.LoadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	return nil
}

func (s *LoadBalancerLoadBalancerRoutingPolicyResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	loadBalancerId, routingPolicyName, err := parseLoadBalancerRoutingPolicyCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("load_balancer_id", &loadBalancerId)
		s.D.Set("name", &routingPolicyName)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	s.D.Set("condition_language_version", s.Res.ConditionLanguageVersion)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	rules := []interface{}{}
	for _, item := range s.Res.Rules {
		rules = append(rules, RoutingRuleToMap(item))
	}
	s.D.Set("rules", rules)

	return nil
}

func GetLoadBalancerRoutingPolicyCompositeId(loadBalancerId string, routingPolicyName string) string {
	loadBalancerId = url.PathEscape(loadBalancerId)
	routingPolicyName = url.PathEscape(routingPolicyName)
	compositeId := "loadBalancers/" + loadBalancerId + "/routingPolicies/" + routingPolicyName
	return compositeId
}

func parseLoadBalancerRoutingPolicyCompositeId(compositeId string) (loadBalancerId string, routingPolicyName string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("loadBalancers/.*/routingPolicies/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	loadBalancerId, _ = url.PathUnescape(parts[1])
	routingPolicyName, _ = url.PathUnescape(parts[3])

	return
}

func (s *LoadBalancerLoadBalancerRoutingPolicyResourceCrud) mapToAction(fieldKeyFormat string) (oci_load_balancer.Action, error) {
	var baseObject oci_load_balancer.Action
	//discriminator
	nameRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name"))
	var name string
	if ok {
		name = nameRaw.(string)
	} else {
		name = "" // default value
	}
	switch strings.ToLower(name) {
	case strings.ToLower("FORWARD_TO_BACKENDSET"):
		details := oci_load_balancer.ForwardToBackendSet{}
		if backendSetName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backend_set_name")); ok {
			tmp := backendSetName.(string)
			details.BackendSetName = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown name '%v' was specified", name)
	}
	return baseObject, nil
}

func RoutingRuleActionToMap(obj oci_load_balancer.Action) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_load_balancer.ForwardToBackendSet:
		result["name"] = "FORWARD_TO_BACKENDSET"

		if v.BackendSetName != nil {
			result["backend_set_name"] = string(*v.BackendSetName)
		}
	default:
		log.Printf("[WARN] Received 'name' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *LoadBalancerLoadBalancerRoutingPolicyResourceCrud) mapToRoutingRule(fieldKeyFormat string) (oci_load_balancer.RoutingRule, error) {
	result := oci_load_balancer.RoutingRule{}

	if actions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "actions")); ok {
		interfaces := actions.([]interface{})
		tmp := make([]oci_load_balancer.Action, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "actions"), stateDataIndex)
			converted, err := s.mapToAction(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "actions")) {
			result.Actions = tmp
		}
	}

	if condition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "condition")); ok {
		tmp := condition.(string)
		result.Condition = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}

func RoutingRuleToMap(obj oci_load_balancer.RoutingRule) map[string]interface{} {
	result := map[string]interface{}{}

	actions := []interface{}{}
	for _, item := range obj.Actions {
		actions = append(actions, RoutingRuleActionToMap(item))
	}
	result["actions"] = actions

	if obj.Condition != nil {
		result["condition"] = string(*obj.Condition)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}
