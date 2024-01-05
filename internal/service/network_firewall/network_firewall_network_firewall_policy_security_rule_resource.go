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

func NetworkFirewallNetworkFirewallPolicySecurityRuleResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createNetworkFirewallNetworkFirewallPolicySecurityRule,
		Read:     readNetworkFirewallNetworkFirewallPolicySecurityRule,
		Update:   updateNetworkFirewallNetworkFirewallPolicySecurityRule,
		Delete:   deleteNetworkFirewallNetworkFirewallPolicySecurityRule,
		Schema: map[string]*schema.Schema{
			// Required
			"action": {
				Type:     schema.TypeString,
				Required: true,
			},
			"condition": {
				Type:     schema.TypeList,
				MinItems: 1,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"application": {
							Type:     schema.TypeList,
							Optional: true,
							MinItems: 0,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"destination_address": {
							Type:     schema.TypeList,
							Optional: true,
							MinItems: 0,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"service": {
							Type:     schema.TypeList,
							Optional: true,
							MinItems: 0,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"source_address": {
							Type:     schema.TypeList,
							Optional: true,
							MinItems: 0,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"url": {
							Type:     schema.TypeList,
							Optional: true,
							MinItems: 0,
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
			"position": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Optional
						"after_rule": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"before_rule": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			// Optional
			"inspection": {
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

func createNetworkFirewallNetworkFirewallPolicySecurityRule(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicySecurityRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.CreateResource(d, sync)
}

func readNetworkFirewallNetworkFirewallPolicySecurityRule(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicySecurityRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.ReadResource(sync)
}

func updateNetworkFirewallNetworkFirewallPolicySecurityRule(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicySecurityRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteNetworkFirewallNetworkFirewallPolicySecurityRule(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicySecurityRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type NetworkFirewallNetworkFirewallPolicySecurityRuleResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_network_firewall.NetworkFirewallClient
	Res                    *oci_network_firewall.SecurityRule
	DisableNotFoundRetries bool
}

func (s *NetworkFirewallNetworkFirewallPolicySecurityRuleResourceCrud) ID() string {
	return GetNetworkFirewallPolicySubResourceCompositeId(s.D.Get("name").(string), s.D.Get("network_firewall_policy_id").(string), "securityRules")
}

func (s *NetworkFirewallNetworkFirewallPolicySecurityRuleResourceCrud) Create() error {
	request := oci_network_firewall.CreateSecurityRuleRequest{}

	if action, ok := s.D.GetOkExists("action"); ok {
		request.Action = oci_network_firewall.TrafficActionTypeEnum(action.(string))
	}

	if condition, ok := s.D.GetOkExists("condition"); ok {
		if tmpList := condition.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "condition", 0)
			tmp, err := s.mapToSecurityRuleMatchCriteria(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Condition = &tmp
		}
	}

	if inspection, ok := s.D.GetOkExists("inspection"); ok {
		request.Inspection = oci_network_firewall.TrafficInspectionTypeEnum(inspection.(string))
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
			tmp, err := s.mapToRulePosition(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Position = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.CreateSecurityRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SecurityRule
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicySecurityRuleResourceCrud) Get() error {
	request := oci_network_firewall.GetSecurityRuleRequest{}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	if securityRuleName, ok := s.D.GetOkExists("name"); ok {
		tmp := securityRuleName.(string)
		request.SecurityRuleName = &tmp
	}

	securityRuleName, networkFirewallPolicyId, err := parseNetworkFirewallPolicySubResourceCompositeId(s.D.Id(), "securityRules")
	if err == nil {
		request.NetworkFirewallPolicyId = &networkFirewallPolicyId
		request.SecurityRuleName = &securityRuleName
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.GetSecurityRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SecurityRule
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicySecurityRuleResourceCrud) Update() error {
	request := oci_network_firewall.UpdateSecurityRuleRequest{}

	if action, ok := s.D.GetOkExists("action"); ok {
		request.Action = oci_network_firewall.TrafficActionTypeEnum(action.(string))
	}

	if condition, ok := s.D.GetOkExists("condition"); ok {
		if tmpList := condition.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "condition", 0)
			tmp, err := s.mapToSecurityRuleMatchCriteria(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Condition = &tmp
		}
	}

	if inspection, ok := s.D.GetOkExists("inspection"); ok {
		request.Inspection = oci_network_firewall.TrafficInspectionTypeEnum(inspection.(string))
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	if position, ok := s.D.GetOkExists("position"); ok {
		if tmpList := position.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "position", 0)
			tmp, err := s.mapToRulePosition(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Position = &tmp
		}
	}

	if securityRuleName, ok := s.D.GetOkExists("name"); ok {
		tmp := securityRuleName.(string)
		request.SecurityRuleName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.UpdateSecurityRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SecurityRule
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicySecurityRuleResourceCrud) Delete() error {
	request := oci_network_firewall.DeleteSecurityRuleRequest{}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	if securityRuleName, ok := s.D.GetOkExists("name"); ok {
		tmp := securityRuleName.(string)
		request.SecurityRuleName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	_, err := s.Client.DeleteSecurityRule(context.Background(), request)
	return err
}

func (s *NetworkFirewallNetworkFirewallPolicySecurityRuleResourceCrud) SetData() error {

	securityRuleName, networkFirewallPolicyId, err := parseNetworkFirewallPolicySubResourceCompositeId(s.D.Id(), "securityRules")
	if err == nil {
		s.D.Set("network_firewall_policy_id", &networkFirewallPolicyId)
		s.D.Set("name", &securityRuleName)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	s.D.Set("action", s.Res.Action)

	if s.Res.Condition != nil {
		s.D.Set("condition", []interface{}{SecurityRuleMatchCriteriaToMap(s.Res.Condition)})
	} else {
		s.D.Set("condition", nil)
	}

	s.D.Set("inspection", s.Res.Inspection)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ParentResourceId != nil {
		s.D.Set("parent_resource_id", *s.Res.ParentResourceId)
	}

	if s.Res.Position != nil {
		s.D.Set("position", []interface{}{RulePositionToMap(s.Res.Position)})
	} else {
		s.D.Set("position", nil)
	}

	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicySecurityRuleResourceCrud) mapToRulePosition(fieldKeyFormat string) (oci_network_firewall.RulePosition, error) {
	result := oci_network_firewall.RulePosition{}

	if afterRule, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "after_rule")); ok {
		tmp := afterRule.(string)
		result.AfterRule = &tmp
	}

	if beforeRule, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "before_rule")); ok {
		tmp := beforeRule.(string)
		result.BeforeRule = &tmp
	}

	return result, nil
}

func RulePositionToMap(obj *oci_network_firewall.RulePosition) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AfterRule != nil {
		result["after_rule"] = string(*obj.AfterRule)
	}

	if obj.BeforeRule != nil {
		result["before_rule"] = string(*obj.BeforeRule)
	}

	return result
}

func (s *NetworkFirewallNetworkFirewallPolicySecurityRuleResourceCrud) mapToSecurityRuleMatchCriteria(fieldKeyFormat string) (oci_network_firewall.SecurityRuleMatchCriteria, error) {
	result := oci_network_firewall.SecurityRuleMatchCriteria{}

	if application, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "application")); ok {
		interfaces := application.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		result.Application = tmp
	}

	if destinationAddress, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_address")); ok {
		interfaces := destinationAddress.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		result.DestinationAddress = tmp
	}

	if service, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service")); ok {
		interfaces := service.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		result.Service = tmp
	}

	if sourceAddress, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_address")); ok {
		interfaces := sourceAddress.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		result.SourceAddress = tmp
	}

	if url, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "url")); ok {
		interfaces := url.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		result.Url = tmp
	}

	return result, nil
}

func SecurityRuleMatchCriteriaToMap(obj *oci_network_firewall.SecurityRuleMatchCriteria) map[string]interface{} {
	result := map[string]interface{}{}

	result["application"] = obj.Application

	result["destination_address"] = obj.DestinationAddress

	result["service"] = obj.Service

	result["source_address"] = obj.SourceAddress

	result["url"] = obj.Url

	return result
}

func SecurityRuleSummaryToMap(obj oci_network_firewall.SecurityRuleSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["action"] = string(obj.Action)

	result["inspection"] = string(obj.Inspection)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ParentResourceId != nil {
		result["parent_resource_id"] = string(*obj.ParentResourceId)
	}

	if obj.PriorityOrder != nil {
		result["priority_order"] = strconv.FormatInt(*obj.PriorityOrder, 10)
	}

	return result
}
