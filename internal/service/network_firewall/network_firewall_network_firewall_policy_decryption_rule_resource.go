// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package network_firewall

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_network_firewall "github.com/oracle/oci-go-sdk/v65/networkfirewall"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func NetworkFirewallNetworkFirewallPolicyDecryptionRuleResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createNetworkFirewallNetworkFirewallPolicyDecryptionRule,
		Read:     readNetworkFirewallNetworkFirewallPolicyDecryptionRule,
		Update:   updateNetworkFirewallNetworkFirewallPolicyDecryptionRule,
		Delete:   deleteNetworkFirewallNetworkFirewallPolicyDecryptionRule,
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
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"source_address": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
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

			// Optional
			"decryption_profile": {
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
					},
				},
			},
			"secret": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority_order": {
				Type:     schema.TypeString,
				Optional: true,
			},

			// Computed
			"parent_resource_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createNetworkFirewallNetworkFirewallPolicyDecryptionRule(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyDecryptionRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.CreateResource(d, sync)
}

func readNetworkFirewallNetworkFirewallPolicyDecryptionRule(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyDecryptionRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.ReadResource(sync)
}

func updateNetworkFirewallNetworkFirewallPolicyDecryptionRule(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyDecryptionRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteNetworkFirewallNetworkFirewallPolicyDecryptionRule(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyDecryptionRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type NetworkFirewallNetworkFirewallPolicyDecryptionRuleResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_network_firewall.NetworkFirewallClient
	Res                    *oci_network_firewall.DecryptionRule
	DisableNotFoundRetries bool
}

func (s *NetworkFirewallNetworkFirewallPolicyDecryptionRuleResourceCrud) ID() string {
	return GetNetworkFirewallPolicySubResourceCompositeId(s.D.Get("name").(string), s.D.Get("network_firewall_policy_id").(string), "decryptionRules")
}

func (s *NetworkFirewallNetworkFirewallPolicyDecryptionRuleResourceCrud) Create() error {
	request := oci_network_firewall.CreateDecryptionRuleRequest{}

	if action, ok := s.D.GetOkExists("action"); ok {
		request.Action = oci_network_firewall.DecryptionActionTypeEnum(action.(string))
	}

	if condition, ok := s.D.GetOkExists("condition"); ok {
		if tmpList := condition.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "condition", 0)
			tmp, err := s.mapToDecryptionRuleMatchCriteria(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Condition = &tmp
		}
	}

	if decryptionProfile, ok := s.D.GetOkExists("decryption_profile"); ok {
		tmp := decryptionProfile.(string)
		request.DecryptionProfile = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
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
			request.Position = &tmp
		}
	}

	if secret, ok := s.D.GetOkExists("secret"); ok {
		tmp := secret.(string)
		request.Secret = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.CreateDecryptionRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DecryptionRule
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyDecryptionRuleResourceCrud) Get() error {
	request := oci_network_firewall.GetDecryptionRuleRequest{}

	if decryptionRuleName, ok := s.D.GetOkExists("name"); ok {
		tmp := decryptionRuleName.(string)
		request.DecryptionRuleName = &tmp
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	decryptionRuleName, networkFirewallPolicyId, err := parseNetworkFirewallPolicySubResourceCompositeId(s.D.Id(), "decryptionRules")
	if err == nil {
		request.DecryptionRuleName = &decryptionRuleName
		request.NetworkFirewallPolicyId = &networkFirewallPolicyId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.GetDecryptionRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DecryptionRule
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyDecryptionRuleResourceCrud) Update() error {
	request := oci_network_firewall.UpdateDecryptionRuleRequest{}

	if action, ok := s.D.GetOkExists("action"); ok {
		request.Action = oci_network_firewall.DecryptionActionTypeEnum(action.(string))
	}

	if condition, ok := s.D.GetOkExists("condition"); ok {
		if tmpList := condition.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "condition", 0)
			tmp, err := s.mapToDecryptionRuleMatchCriteria(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Condition = &tmp
		}
	}

	if decryptionProfile, ok := s.D.GetOkExists("decryption_profile"); ok {
		tmp := decryptionProfile.(string)
		request.DecryptionProfile = &tmp
	}

	if decryptionRuleName, ok := s.D.GetOkExists("name"); ok {
		tmp := decryptionRuleName.(string)
		request.DecryptionRuleName = &tmp
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
			request.Position = &tmp
		}
	}

	if secret, ok := s.D.GetOkExists("secret"); ok {
		tmp := secret.(string)
		request.Secret = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.UpdateDecryptionRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DecryptionRule
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyDecryptionRuleResourceCrud) Delete() error {
	request := oci_network_firewall.DeleteDecryptionRuleRequest{}

	if decryptionRuleName, ok := s.D.GetOkExists("name"); ok {
		tmp := decryptionRuleName.(string)
		request.DecryptionRuleName = &tmp
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	_, err := s.Client.DeleteDecryptionRule(context.Background(), request)
	return err
}

func (s *NetworkFirewallNetworkFirewallPolicyDecryptionRuleResourceCrud) SetData() error {

	decryptionRuleName, networkFirewallPolicyId, err := parseNetworkFirewallPolicySubResourceCompositeId(s.D.Id(), "decryptionRules")
	if err == nil {
		s.D.Set("name", &decryptionRuleName)
		s.D.Set("network_firewall_policy_id", &networkFirewallPolicyId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	s.D.Set("action", s.Res.Action)

	if s.Res.Condition != nil {
		s.D.Set("condition", []interface{}{DecryptionRuleMatchCriteriaToMap(s.Res.Condition)})
	} else {
		s.D.Set("condition", nil)
	}

	if s.Res.DecryptionProfile != nil {
		s.D.Set("decryption_profile", *s.Res.DecryptionProfile)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ParentResourceId != nil {
		s.D.Set("parent_resource_id", *s.Res.ParentResourceId)
	}

	if s.Res.Position != nil {
		s.D.Set("position", []interface{}{RulePositionToMapDecRule(s.Res.Position)})
	} else {
		s.D.Set("position", nil)
	}

	if s.Res.Secret != nil {
		s.D.Set("secret", *s.Res.Secret)
	}

	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyDecryptionRuleResourceCrud) mapToDecryptionRuleMatchCriteria(fieldKeyFormat string) (oci_network_firewall.DecryptionRuleMatchCriteria, error) {
	result := oci_network_firewall.DecryptionRuleMatchCriteria{}

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

func DecryptionRuleMatchCriteriaToMap(obj *oci_network_firewall.DecryptionRuleMatchCriteria) map[string]interface{} {
	result := map[string]interface{}{}

	result["destination_address"] = obj.DestinationAddress
	result["destination_address"] = obj.DestinationAddress

	result["source_address"] = obj.SourceAddress
	result["source_address"] = obj.SourceAddress

	return result
}

func DecryptionRuleSummaryToMap(obj oci_network_firewall.DecryptionRuleSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["action"] = string(obj.Action)

	if obj.DecryptionProfile != nil {
		result["decryption_profile"] = string(*obj.DecryptionProfile)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ParentResourceId != nil {
		result["parent_resource_id"] = string(*obj.ParentResourceId)
	}

	if obj.PriorityOrder != nil {
		result["priority_order"] = strconv.FormatInt(*obj.PriorityOrder, 10)
	}

	if obj.Secret != nil {
		result["secret"] = string(*obj.Secret)
	}

	return result
}

func (s *NetworkFirewallNetworkFirewallPolicyDecryptionRuleResourceCrud) mapToRulePositionUpdate() (oci_network_firewall.RulePosition, error) {
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
		actualBeforePosition := s.Res.Position.BeforeRule
		if actualBeforePosition != nil {
			result.BeforeRule = actualBeforePosition
			return result, nil
		}

		actualAfterPosition := s.Res.Position.BeforeRule
		if actualAfterPosition != nil {
			result.AfterRule = actualAfterPosition
			return result, nil
		}
	}
	return result, nil
}

func (s *NetworkFirewallNetworkFirewallPolicyDecryptionRuleResourceCrud) mapToRulePositionCreate(fieldKeyFormat string) (oci_network_firewall.RulePosition, error) {
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

func RulePositionToMapDecRule(obj *oci_network_firewall.RulePosition) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AfterRule != nil {
		result["after_rule"] = string(*obj.AfterRule)
	}

	if obj.BeforeRule != nil {
		result["before_rule"] = string(*obj.BeforeRule)
	}

	return result
}
