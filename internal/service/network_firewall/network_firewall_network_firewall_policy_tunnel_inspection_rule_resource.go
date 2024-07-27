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

func NetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createNetworkFirewallNetworkFirewallPolicyTunnelInspectionRule,
		Read:     readNetworkFirewallNetworkFirewallPolicyTunnelInspectionRule,
		Update:   updateNetworkFirewallNetworkFirewallPolicyTunnelInspectionRule,
		Delete:   deleteNetworkFirewallNetworkFirewallPolicyTunnelInspectionRule,
		Schema: map[string]*schema.Schema{
			// Required
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
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"source_address": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
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
			"protocol": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"VXLAN",
				}, true),
			},

			// Optional
			"action": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"position": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
					},
				},
			},
			"profile": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"must_return_traffic_to_source": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
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

func createNetworkFirewallNetworkFirewallPolicyTunnelInspectionRule(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.CreateResource(d, sync)
}

func readNetworkFirewallNetworkFirewallPolicyTunnelInspectionRule(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.ReadResource(sync)
}

func updateNetworkFirewallNetworkFirewallPolicyTunnelInspectionRule(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteNetworkFirewallNetworkFirewallPolicyTunnelInspectionRule(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type NetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_network_firewall.NetworkFirewallClient
	Res                    *oci_network_firewall.TunnelInspectionRule
	DisableNotFoundRetries bool
}

func (s *NetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleResourceCrud) ID() string {
	return GetNetworkFirewallPolicySubResourceCompositeId(s.D.Get("name").(string), s.D.Get("network_firewall_policy_id").(string), "tunnelInspectionRules")
}

func (s *NetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleResourceCrud) Create() error {
	request := oci_network_firewall.CreateTunnelInspectionRuleRequest{}
	err := s.populateTopLevelPolymorphicCreateTunnelInspectionRuleRequest(&request)
	if err != nil {
		return err
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.CreateTunnelInspectionRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.TunnelInspectionRule
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleResourceCrud) Get() error {
	request := oci_network_firewall.GetTunnelInspectionRuleRequest{}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	if tunnelInspectionRuleName, ok := s.D.GetOkExists("name"); ok {
		tmp := tunnelInspectionRuleName.(string)
		request.TunnelInspectionRuleName = &tmp
	}

	networkFirewallPolicyId, tunnelInspectionRuleName, err := parseNetworkFirewallPolicyTunnelInspectionRuleCompositeId(s.D.Id())
	if err == nil {
		request.NetworkFirewallPolicyId = &networkFirewallPolicyId
		request.TunnelInspectionRuleName = &tunnelInspectionRuleName
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.GetTunnelInspectionRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.TunnelInspectionRule
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleResourceCrud) Update() error {
	request := oci_network_firewall.UpdateTunnelInspectionRuleRequest{}
	err := s.populateTopLevelPolymorphicUpdateTunnelInspectionRuleRequest(&request)
	if err != nil {
		return err
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	if ruleName, ok := s.D.GetOkExists("name"); ok {
		tmp := ruleName.(string)
		request.TunnelInspectionRuleName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.UpdateTunnelInspectionRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.TunnelInspectionRule
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleResourceCrud) Delete() error {
	request := oci_network_firewall.DeleteTunnelInspectionRuleRequest{}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	if tunnelInspectionRuleName, ok := s.D.GetOkExists("name"); ok {
		tmp := tunnelInspectionRuleName.(string)
		request.TunnelInspectionRuleName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	_, err := s.Client.DeleteTunnelInspectionRule(context.Background(), request)
	return err
}

func (s *NetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleResourceCrud) SetData() error {

	networkFirewallPolicyId, tunnelInspectionRuleName, err := parseNetworkFirewallPolicyTunnelInspectionRuleCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("network_firewall_policy_id", &networkFirewallPolicyId)
		s.D.Set("name", &tunnelInspectionRuleName)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	switch v := (*s.Res).(type) {
	case oci_network_firewall.VxlanInspectionRule:
		s.D.Set("protocol", "VXLAN")

		if v.Condition != nil {
			s.D.Set("condition", []interface{}{VxlanInspectionRuleMatchCriteriaToMap(v.Condition)})
		} else {
			s.D.Set("condition", nil)
		}

		if v.Profile != nil {
			s.D.Set("profile", []interface{}{VxlanInspectionRuleProfileToMap(v.Profile)})
		} else {
			s.D.Set("profile", nil)
		}

		s.D.Set("action", v.Action)

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		if v.ParentResourceId != nil {
			s.D.Set("parent_resource_id", *v.ParentResourceId)
		}

		if v.Position != nil {
			s.D.Set("position", []interface{}{RulePositionToMapTunnelRules(v.Position)})
		} else {
			s.D.Set("position", nil)
		}

		if v.PriorityOrder != nil {
			s.D.Set("priority_order", strconv.FormatInt(*v.PriorityOrder, 10))
		}
	default:
		log.Printf("[WARN] Received 'protocol' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func GetNetworkFirewallPolicyTunnelInspectionRuleCompositeId(networkFirewallPolicyId string, tunnelInspectionRuleName string) string {
	networkFirewallPolicyId = url.PathEscape(networkFirewallPolicyId)
	tunnelInspectionRuleName = url.PathEscape(tunnelInspectionRuleName)
	compositeId := "networkFirewallPolicies/" + networkFirewallPolicyId + "/tunnelInspectionRules/" + tunnelInspectionRuleName
	return compositeId
}

func parseNetworkFirewallPolicyTunnelInspectionRuleCompositeId(compositeId string) (networkFirewallPolicyId string, tunnelInspectionRuleName string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("networkFirewallPolicies/.*/tunnelInspectionRules/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	networkFirewallPolicyId, _ = url.PathUnescape(parts[1])
	tunnelInspectionRuleName, _ = url.PathUnescape(parts[3])

	return
}

func (s *NetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleResourceCrud) mapToRulePositionUpdate() (oci_network_firewall.RulePosition, error) {
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

func (s *NetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleResourceCrud) mapToRulePositionCreate(fieldKeyFormat string) (oci_network_firewall.RulePosition, error) {
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

func RulePositionToMapTunnelRules(obj *oci_network_firewall.RulePosition) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AfterRule != nil {
		result["after_rule"] = string(*obj.AfterRule)
	}

	if obj.BeforeRule != nil {
		result["before_rule"] = string(*obj.BeforeRule)
	}

	return result
}

func TunnelInspectionRuleSummaryToMap(obj oci_network_firewall.TunnelInspectionRuleSummary) map[string]interface{} {
	result := map[string]interface{}{}

	switch v := (obj).(type) {
	case oci_network_firewall.VxlanInspectionRuleSummary:
		if v.Name != nil {
			result["name"] = string(*v.Name)
		}

		result["action"] = string(v.Action)

		if v.ParentResourceId != nil {
			result["parent_resource_id"] = string(*v.ParentResourceId)
		}

		if v.PriorityOrder != nil {
			result["priority_order"] = strconv.FormatInt(*v.PriorityOrder, 10)
		}
		result["protocol"] = "VXLAN"

		if v.Condition != nil {
			result["condition"] = []interface{}{VxlanInspectionRuleMatchCriteriaToMap(v.Condition)}
		}

		if v.Profile != nil {
			result["profile"] = []interface{}{VxlanInspectionRuleProfileToMap(v.Profile)}
		}
	default:
		log.Printf("[WARN] Received 'protocol' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *NetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleResourceCrud) mapToVxlanInspectionRuleMatchCriteria(fieldKeyFormat string) (oci_network_firewall.VxlanInspectionRuleMatchCriteria, error) {
	result := oci_network_firewall.VxlanInspectionRuleMatchCriteria{}

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

func VxlanInspectionRuleMatchCriteriaToMap(obj *oci_network_firewall.VxlanInspectionRuleMatchCriteria) map[string]interface{} {
	result := map[string]interface{}{}

	result["destination_address"] = obj.DestinationAddress

	result["source_address"] = obj.SourceAddress

	return result
}

func (s *NetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleResourceCrud) mapToVxlanInspectionRuleProfile(fieldKeyFormat string) (oci_network_firewall.VxlanInspectionRuleProfile, error) {
	result := oci_network_firewall.VxlanInspectionRuleProfile{}

	if mustReturnTrafficToSource, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "must_return_traffic_to_source")); ok {
		tmp := mustReturnTrafficToSource.(bool)
		result.MustReturnTrafficToSource = &tmp
	}

	return result, nil
}

func VxlanInspectionRuleProfileToMap(obj *oci_network_firewall.VxlanInspectionRuleProfile) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MustReturnTrafficToSource != nil {
		result["must_return_traffic_to_source"] = bool(*obj.MustReturnTrafficToSource)
	}

	return result
}

func (s *NetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleResourceCrud) populateTopLevelPolymorphicCreateTunnelInspectionRuleRequest(request *oci_network_firewall.CreateTunnelInspectionRuleRequest) error {
	//discriminator
	protocolRaw, ok := s.D.GetOkExists("protocol")
	var protocol string
	if ok {
		protocol = protocolRaw.(string)
	} else {
		protocol = "" // default value
	}
	switch strings.ToLower(protocol) {
	case strings.ToLower("VXLAN"):
		details := oci_network_firewall.CreateVxlanInspectionRuleDetails{}
		if condition, ok := s.D.GetOkExists("condition"); ok {
			if tmpList := condition.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "condition", 0)
				tmp, err := s.mapToVxlanInspectionRuleMatchCriteria(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Condition = &tmp
			}
		}
		if profile, ok := s.D.GetOkExists("profile"); ok {
			if tmpList := profile.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "profile", 0)
				tmp, err := s.mapToVxlanInspectionRuleProfile(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Profile = &tmp
			}
		}
		if action, ok := s.D.GetOkExists("action"); ok {
			details.Action = oci_network_firewall.InspectActionTypeEnum(action.(string))
		}
		if name, ok := s.D.GetOkExists("name"); ok {
			tmp := name.(string)
			details.Name = &tmp
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
		request.CreateTunnelInspectionRuleDetails = details

	default:
		return fmt.Errorf("unknown protocol '%v' was specified", protocol)
	}
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleResourceCrud) populateTopLevelPolymorphicUpdateTunnelInspectionRuleRequest(request *oci_network_firewall.UpdateTunnelInspectionRuleRequest) error {
	//discriminator
	protocolRaw, ok := s.D.GetOkExists("protocol")
	var protocol string
	if ok {
		protocol = protocolRaw.(string)
	} else {
		protocol = "" // default value
	}
	switch strings.ToLower(protocol) {
	case strings.ToLower("VXLAN"):
		details := oci_network_firewall.UpdateVxlanInspectionRuleDetails{}
		if condition, ok := s.D.GetOkExists("condition"); ok {
			if tmpList := condition.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "condition", 0)
				tmp, err := s.mapToVxlanInspectionRuleMatchCriteria(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Condition = &tmp
			}
		}
		if profile, ok := s.D.GetOkExists("profile"); ok {
			if tmpList := profile.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "profile", 0)
				tmp, err := s.mapToVxlanInspectionRuleProfile(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Profile = &tmp
			}
		}
		if action, ok := s.D.GetOkExists("action"); ok {
			details.Action = oci_network_firewall.InspectActionTypeEnum(action.(string))
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
		request.UpdateTunnelInspectionRuleDetails = details
	default:
		return fmt.Errorf("unknown protocol '%v' was specified", protocol)
	}
	return nil
}
