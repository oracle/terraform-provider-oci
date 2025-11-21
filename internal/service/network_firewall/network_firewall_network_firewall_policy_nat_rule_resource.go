// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package network_firewall

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_network_firewall "github.com/oracle/oci-go-sdk/v65/networkfirewall"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func NetworkFirewallNetworkFirewallPolicyNatRuleResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createNetworkFirewallNetworkFirewallPolicyNatRule,
		Read:     readNetworkFirewallNetworkFirewallPolicyNatRule,
		Update:   updateNetworkFirewallNetworkFirewallPolicyNatRule,
		Delete:   deleteNetworkFirewallNetworkFirewallPolicyNatRule,
		Schema: map[string]*schema.Schema{
			// Required
			"action": {
				Type:     schema.TypeString,
				Required: true,
			},
			"condition": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"destination_address": {
							Type:     schema.TypeList,
							Optional: true,
							MinItems: 0,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"service": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"source_address": {
							Type:     schema.TypeList,
							Optional: true,
							MinItems: 0,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"network_firewall_policy_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"NATV4",
				}, true),
			},

			// Optional
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"position": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"after_rule": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"before_rule": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Computed
			"parent_resource_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"priority_order": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createNetworkFirewallNetworkFirewallPolicyNatRule(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyNatRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.CreateResource(d, sync)
}

func readNetworkFirewallNetworkFirewallPolicyNatRule(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyNatRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.ReadResource(sync)
}

func updateNetworkFirewallNetworkFirewallPolicyNatRule(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyNatRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteNetworkFirewallNetworkFirewallPolicyNatRule(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyNatRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type NetworkFirewallNetworkFirewallPolicyNatRuleResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_network_firewall.NetworkFirewallClient
	Res                    *oci_network_firewall.NatRule
	DisableNotFoundRetries bool
}

func (s *NetworkFirewallNetworkFirewallPolicyNatRuleResourceCrud) ID() string {
	return GetNetworkFirewallPolicySubResourceCompositeId(s.D.Get("name").(string), s.D.Get("network_firewall_policy_id").(string), "natRules")
}
func (s *NetworkFirewallNetworkFirewallPolicyNatRuleResourceCrud) Create() error {
	request := oci_network_firewall.CreateNatRuleRequest{}
	err := s.populateTopLevelPolymorphicCreateNatRuleRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.CreateNatRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NatRule
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyNatRuleResourceCrud) Get() error {
	request := oci_network_firewall.GetNatRuleRequest{}

	if natRuleName, ok := s.D.GetOkExists("name"); ok {
		tmp := natRuleName.(string)
		request.NatRuleName = &tmp
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	natRuleName, networkFirewallPolicyId, err := parseNetworkFirewallPolicyNatRuleCompositeId(s.D.Id())
	if err == nil {
		request.NatRuleName = &natRuleName
		request.NetworkFirewallPolicyId = &networkFirewallPolicyId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.GetNatRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NatRule
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyNatRuleResourceCrud) Update() error {
	request := oci_network_firewall.UpdateNatRuleRequest{}
	err := s.populateTopLevelPolymorphicUpdateNatRuleRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.UpdateNatRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NatRule
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyNatRuleResourceCrud) Delete() error {
	request := oci_network_firewall.DeleteNatRuleRequest{}

	if natRuleName, ok := s.D.GetOkExists("name"); ok {
		tmp := natRuleName.(string)
		request.NatRuleName = &tmp
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	_, err := s.Client.DeleteNatRule(context.Background(), request)
	return err
}

func (s *NetworkFirewallNetworkFirewallPolicyNatRuleResourceCrud) SetData() error {

	natRuleName, networkFirewallPolicyId, err := parseNetworkFirewallPolicyNatRuleCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("name", &natRuleName)
		s.D.Set("network_firewall_policy_id", &networkFirewallPolicyId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	switch v := (*s.Res).(type) {
	case oci_network_firewall.NatV4NatRule:
		s.D.Set("type", "NATV4")

		s.D.Set("action", v.Action)

		if v.Condition != nil {
			s.D.Set("condition", []interface{}{NatRuleMatchCriteriaToMap(v.Condition)})
		} else {
			s.D.Set("condition", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		if v.ParentResourceId != nil {
			s.D.Set("parent_resource_id", *v.ParentResourceId)
		}

		if v.Position != nil {
			s.D.Set("position", []interface{}{RulePositionToMap(v.Position)})
		} else {
			s.D.Set("position", nil)
		}

		if v.PriorityOrder != nil {
			s.D.Set("priority_order", strconv.FormatInt(*v.PriorityOrder, 10))
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func GetNetworkFirewallPolicyNatRuleCompositeId(natRuleName string, networkFirewallPolicyId string) string {
	natRuleName = url.PathEscape(natRuleName)
	networkFirewallPolicyId = url.PathEscape(networkFirewallPolicyId)
	compositeId := "networkFirewallPolicies/" + networkFirewallPolicyId + "/natRules/" + natRuleName
	return compositeId
}

func parseNetworkFirewallPolicyNatRuleCompositeId(compositeId string) (natRuleName string, networkFirewallPolicyId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("networkFirewallPolicies/.*/natRules/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	networkFirewallPolicyId, _ = url.PathUnescape(parts[1])
	natRuleName, _ = url.PathUnescape(parts[3])

	return
}

func (s *NetworkFirewallNetworkFirewallPolicyNatRuleResourceCrud) mapToNatRuleMatchCriteria(fieldKeyFormat string) (oci_network_firewall.NatRuleMatchCriteria, error) {
	result := oci_network_firewall.NatRuleMatchCriteria{}

	if destinationAddress, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_address")); ok {
		interfaces := destinationAddress.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "destination_address")) {
			result.DestinationAddress = tmp
		}
	}

	if service, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service")); ok {
		tmp := service.(string)
		result.Service = &tmp
	}

	if sourceAddress, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_address")); ok {
		interfaces := sourceAddress.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "source_address")) {
			result.SourceAddress = tmp
		}
	}

	return result, nil
}

func NatRuleMatchCriteriaToMap(obj *oci_network_firewall.NatRuleMatchCriteria) map[string]interface{} {
	result := map[string]interface{}{}

	result["destination_address"] = obj.DestinationAddress

	if obj.Service != nil {
		result["service"] = string(*obj.Service)
	}

	result["source_address"] = obj.SourceAddress

	return result
}

func NatRuleSummaryToMap(obj oci_network_firewall.NatRuleSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_network_firewall.NatV4NatSummary:
		if v.Name != nil {
			result["name"] = string(*v.Name)
		}

		result["type"] = "NATV4"

		result["action"] = string(v.Action)

		if v.ParentResourceId != nil {
			result["parent_resource_id"] = string(*v.ParentResourceId)
		}

		if v.Condition != nil {
			result["condition"] = []interface{}{NatRuleMatchCriteriaToMap(v.Condition)}
		}

		if v.PriorityOrder != nil {
			result["priority_order"] = strconv.FormatInt(*v.PriorityOrder, 10)
		}

	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *NetworkFirewallNetworkFirewallPolicyNatRuleResourceCrud) mapToRulePositionCreate(fieldKeyFormat string) (oci_network_firewall.RulePosition, error) {
	result := oci_network_firewall.RulePosition{}

	beforeRule, beforeRuleOk := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "before_rule"))
	afterRule, afterRuleOk := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "after_rule"))

	if beforeRule != "" && afterRule != "" {
		return result, fmt.Errorf("rule position cannot be ambiguous, provide one of before_rule or after_rule rule positions")
	}

	if beforeRuleOk {
		tmp := beforeRule.(string)
		if beforeRule != "" {
			result.BeforeRule = &tmp
			return result, nil
		}
	}

	if afterRuleOk {
		tmp := afterRule.(string)
		if afterRule != "" {
			result.AfterRule = &tmp
			return result, nil
		}
	}

	return result, nil
}

func RulePositionToMapNatRule(obj *oci_network_firewall.RulePosition) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AfterRule != nil {
		result["after_rule"] = string(*obj.AfterRule)
	}

	if obj.BeforeRule != nil {
		result["before_rule"] = string(*obj.BeforeRule)
	}

	return result
}

func (s *NetworkFirewallNetworkFirewallPolicyNatRuleResourceCrud) mapToRulePositionUpdate() (oci_network_firewall.RulePosition, error) {
	result := oci_network_firewall.RulePosition{}

	if s.D.HasChange("position") {
		oldPosition, newPosition := s.D.GetChange("position")

		oldPositionMap := oldPosition.([]interface{})[0].(map[string]interface{})
		oldBeforeRule := fmt.Sprintf("%v", oldPositionMap["before_rule"])
		oldAfterRule := fmt.Sprintf("%v", oldPositionMap["after_rule"])

		newPositionMap := newPosition.([]interface{})[0].(map[string]interface{})
		newBeforeRule := fmt.Sprintf("%v", newPositionMap["before_rule"])
		newAfterRule := fmt.Sprintf("%v", newPositionMap["after_rule"])

		if newBeforeRule != oldBeforeRule && newAfterRule != oldAfterRule && newBeforeRule != "" && newAfterRule != "" {
			return result, fmt.Errorf("rule position cannot be ambiguous, provide one of before_rule or after_rule rule positions")
		}

		if newBeforeRule != "" && newBeforeRule != oldBeforeRule {
			result.BeforeRule = &newBeforeRule
			return result, nil
		}

		if newAfterRule != "" && newAfterRule != oldAfterRule {
			result.AfterRule = &newAfterRule
			return result, nil
		}
	}

	err := s.Get()

	if err == nil {
		actualBeforePosition := (*s.Res).GetPosition().BeforeRule
		if actualBeforePosition != nil {
			result.BeforeRule = actualBeforePosition
			return result, nil
		}

		actualAfterPosition := (*s.Res).GetPosition().AfterRule
		if actualAfterPosition != nil {
			result.AfterRule = actualAfterPosition
			return result, nil
		}
	}
	return result, nil
}

func (s *NetworkFirewallNetworkFirewallPolicyNatRuleResourceCrud) populateTopLevelPolymorphicCreateNatRuleRequest(request *oci_network_firewall.CreateNatRuleRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("NATV4"):
		details := oci_network_firewall.CreateNatV4RuleDetails{}
		if action, ok := s.D.GetOkExists("action"); ok {
			details.Action = oci_network_firewall.NatV4ActionTypeEnum(action.(string))
		}
		if condition, ok := s.D.GetOkExists("condition"); ok {
			if tmpList := condition.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "condition", 0)
				tmp, err := s.mapToNatRuleMatchCriteria(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Condition = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if name, ok := s.D.GetOkExists("name"); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
			tmp := networkFirewallPolicyId.(string)
			request.NetworkFirewallPolicyId = &tmp
		}
		if position, ok := s.D.GetOkExists("position"); ok {
			if tmpList := position.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "position", 0)
				tmp, err := s.mapToRulePositionCreate(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Position = &tmp
			}
		}
		request.CreateNatRuleDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyNatRuleResourceCrud) populateTopLevelPolymorphicUpdateNatRuleRequest(request *oci_network_firewall.UpdateNatRuleRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("NATV4"):
		details := oci_network_firewall.UpdateNatV4RuleDetails{}
		if action, ok := s.D.GetOkExists("action"); ok {
			details.Action = oci_network_firewall.NatV4ActionTypeEnum(action.(string))
		}
		if condition, ok := s.D.GetOkExists("condition"); ok {
			if tmpList := condition.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "condition", 0)
				tmp, err := s.mapToNatRuleMatchCriteria(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Condition = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if natRuleName, ok := s.D.GetOkExists("name"); ok {
			tmp := natRuleName.(string)
			request.NatRuleName = &tmp
		}

		if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
			tmp := networkFirewallPolicyId.(string)
			request.NetworkFirewallPolicyId = &tmp
		}
		if position, ok := s.D.GetOkExists("position"); ok {
			if tmpList := position.([]interface{}); len(tmpList) > 0 {
				tmp, err := s.mapToRulePositionUpdate()
				if err != nil {
					return err
				}
				details.Position = &tmp
			}
		}
		request.UpdateNatRuleDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}
