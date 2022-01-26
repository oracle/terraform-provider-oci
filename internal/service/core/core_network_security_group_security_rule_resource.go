// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	oci_core "github.com/oracle/oci-go-sdk/v56/core"
)

func CoreNetworkSecurityGroupSecurityRuleResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreNetworkSecurityGroupSecurityRule,
		Read:     readCoreNetworkSecurityGroupSecurityRule,
		Update:   updateCoreNetworkSecurityGroupSecurityRule,
		Delete:   deleteCoreNetworkSecurityGroupSecurityRule,
		Schema: map[string]*schema.Schema{
			// Required
			"network_security_group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"direction": {
				Type:     schema.TypeString,
				Required: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destination": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destination_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icmp_options": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:     schema.TypeInt,
							Required: true,
						},

						// Optional
						"code": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  -1,
						},

						// Computed
					},
				},
			},
			"source": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stateless": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"tcp_options": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"destination_port_range": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"max": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"min": {
										Type:     schema.TypeInt,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"source_port_range": {
							Type:     schema.TypeList,
							Optional: true,
							// @CODEGEN 2/2018: This should not be a computed field as generated, as it breaks how Terraform
							// considers diffs when the source_port_range is removed from config
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"max": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"min": {
										Type:     schema.TypeInt,
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
			"udp_options": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"destination_port_range": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"max": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"min": {
										Type:     schema.TypeInt,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"source_port_range": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"max": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"min": {
										Type:     schema.TypeInt,
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
			// Computed
			"is_valid": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCoreNetworkSecurityGroupSecurityRule(d *schema.ResourceData, m interface{}) error {
	sync := &CoreSecurityRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	return tfresource.CreateResource(d, sync)
}

func readCoreNetworkSecurityGroupSecurityRule(d *schema.ResourceData, m interface{}) error {
	sync := &CoreSecurityRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	return tfresource.ReadResource(sync)
}

func updateCoreNetworkSecurityGroupSecurityRule(d *schema.ResourceData, m interface{}) error {
	sync := &CoreSecurityRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	return tfresource.UpdateResource(d, sync)
}

func deleteCoreNetworkSecurityGroupSecurityRule(d *schema.ResourceData, m interface{}) error {
	sync := &CoreSecurityRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	return tfresource.DeleteResource(d, sync)
}

type CoreSecurityRuleResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.SecurityRule
	DisableNotFoundRetries bool
}

func (s *CoreSecurityRuleResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreSecurityRuleResourceCrud) Create() error {

	request := oci_core.AddNetworkSecurityGroupSecurityRulesRequest{}

	if networkSecurityGroupId, ok := s.D.GetOkExists("network_security_group_id"); ok {
		tmp := networkSecurityGroupId.(string)
		request.NetworkSecurityGroupId = &tmp
	}

	addSecurityRuleDetails := oci_core.AddSecurityRuleDetails{}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		addSecurityRuleDetails.Description = &tmp
	}

	if destination, ok := s.D.GetOkExists("destination"); ok {
		tmp := destination.(string)
		addSecurityRuleDetails.Destination = &tmp
	}

	if destinationType, ok := s.D.GetOkExists("destination_type"); ok {
		addSecurityRuleDetails.DestinationType = oci_core.AddSecurityRuleDetailsDestinationTypeEnum(destinationType.(string))
	}

	if direction, ok := s.D.GetOkExists("direction"); ok {
		addSecurityRuleDetails.Direction = oci_core.AddSecurityRuleDetailsDirectionEnum(direction.(string))
	}

	if icmpOptions, ok := s.D.GetOkExists("icmp_options"); ok {
		if tmpList := icmpOptions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "icmp_options", 0)
			tmp, err := s.mapToIcmpOptions(fieldKeyFormat)
			if err != nil {
				return fmt.Errorf("unable to convert icmp_options, encountered error: %v", err)
			}
			addSecurityRuleDetails.IcmpOptions = &tmp
		}
	}

	if protocol, ok := s.D.GetOkExists("protocol"); ok {
		tmp := protocol.(string)
		addSecurityRuleDetails.Protocol = &tmp
	}

	if source, ok := s.D.GetOkExists("source"); ok {
		tmp := source.(string)
		addSecurityRuleDetails.Source = &tmp
	}

	if sourceType, ok := s.D.GetOkExists("source_type"); ok {
		addSecurityRuleDetails.SourceType = oci_core.AddSecurityRuleDetailsSourceTypeEnum(sourceType.(string))
	}

	if stateless, ok := s.D.GetOkExists("stateless"); ok {
		tmp := stateless.(bool)
		addSecurityRuleDetails.IsStateless = &tmp
	}

	if tcpOptions, ok := s.D.GetOkExists("tcp_options"); ok {
		if tmpList := tcpOptions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tcp_options", 0)
			tmp, err := s.mapToTcpOptions(fieldKeyFormat)
			if err != nil {
				return fmt.Errorf("unable to convert tcp_options, encountered error: %v", err)
			}
			addSecurityRuleDetails.TcpOptions = &tmp
		}
	}

	if udpOptions, ok := s.D.GetOkExists("udp_options"); ok {
		if tmpList := udpOptions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "udp_options", 0)
			tmp, err := s.mapToUdpOptions(fieldKeyFormat)
			if err != nil {
				return fmt.Errorf("unable to convert udp_options, encountered error: %v", err)
			}
			addSecurityRuleDetails.UdpOptions = &tmp
		}
	}

	tmp := []oci_core.AddSecurityRuleDetails{addSecurityRuleDetails}
	request.SecurityRules = tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.AddNetworkSecurityGroupSecurityRules(context.Background(), request)
	if err != nil {
		return err
	}

	if len(response.SecurityRules) > 0 {
		s.Res = &response.SecurityRules[0]
	} else {
		return fmt.Errorf("security rule missing in response")
	}

	return nil
}

func (s *CoreSecurityRuleResourceCrud) Get() error {

	request := oci_core.ListNetworkSecurityGroupSecurityRulesRequest{}

	if networkSecurityGroupId, ok := s.D.GetOkExists("network_security_group_id"); ok {
		tmp := networkSecurityGroupId.(string)
		request.NetworkSecurityGroupId = &tmp
	}

	networkSecurityGroupId, securityRuleId, err := parseNetworkSecurityGroupSecurityRuleCompositeId(s.D.Id())
	if err == nil {
		request.NetworkSecurityGroupId = &networkSecurityGroupId
		s.D.Set("network_security_group_id", &networkSecurityGroupId)
		s.D.SetId(securityRuleId)
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListNetworkSecurityGroupSecurityRules(context.Background(), request)
	if err != nil {
		return err
	}
	var rules []oci_core.SecurityRule
	rules = response.Items
	request.Page = response.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListNetworkSecurityGroupSecurityRules(context.Background(), request)
		if err != nil {
			return err
		}

		rules = append(rules, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	for _, r := range rules {
		if *r.Id == s.D.Id() {
			s.Res = &r
			break
		}
	}

	if s.Res == nil {
		return fmt.Errorf("security rule not found in the list response")
	}

	return nil
}

func (s *CoreSecurityRuleResourceCrud) Update() error {
	request := oci_core.UpdateNetworkSecurityGroupSecurityRulesRequest{}

	if networkSecurityGroupId, ok := s.D.GetOkExists("network_security_group_id"); ok {
		tmp := networkSecurityGroupId.(string)
		request.NetworkSecurityGroupId = &tmp
	}
	updateSecurityRuleDetails := oci_core.UpdateSecurityRuleDetails{}

	ruleId := s.D.Id()
	updateSecurityRuleDetails.Id = &ruleId

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		updateSecurityRuleDetails.Description = &tmp
	}

	if destination, ok := s.D.GetOkExists("destination"); ok {
		tmp := destination.(string)
		updateSecurityRuleDetails.Destination = &tmp
	}

	if destinationType, ok := s.D.GetOkExists("destination_type"); ok {
		updateSecurityRuleDetails.DestinationType = oci_core.UpdateSecurityRuleDetailsDestinationTypeEnum(destinationType.(string))
	}

	if direction, ok := s.D.GetOkExists("direction"); ok {
		updateSecurityRuleDetails.Direction = oci_core.UpdateSecurityRuleDetailsDirectionEnum(direction.(string))
	}

	if icmpOptions, ok := s.D.GetOkExists("icmp_options"); ok {
		if tmpList := icmpOptions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "icmp_options", 0)
			tmp, err := s.mapToIcmpOptions(fieldKeyFormat)
			if err != nil {
				return fmt.Errorf("unable to convert icmp_options, encountered error: %v", err)
			}
			updateSecurityRuleDetails.IcmpOptions = &tmp
		}
	}

	if protocol, ok := s.D.GetOkExists("protocol"); ok {
		tmp := protocol.(string)
		updateSecurityRuleDetails.Protocol = &tmp
	}

	if source, ok := s.D.GetOkExists("source"); ok {
		tmp := source.(string)
		updateSecurityRuleDetails.Source = &tmp
	}

	if sourceType, ok := s.D.GetOkExists("source_type"); ok {
		updateSecurityRuleDetails.SourceType = oci_core.UpdateSecurityRuleDetailsSourceTypeEnum(sourceType.(string))
	}

	if stateless, ok := s.D.GetOkExists("stateless"); ok {
		tmp := stateless.(bool)
		updateSecurityRuleDetails.IsStateless = &tmp
	}

	if tcpOptions, ok := s.D.GetOkExists("tcp_options"); ok {
		if tmpList := tcpOptions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tcp_options", 0)
			tmp, err := s.mapToTcpOptions(fieldKeyFormat)
			if err != nil {
				return fmt.Errorf("unable to convert tcp_options, encountered error: %v", err)
			}
			updateSecurityRuleDetails.TcpOptions = &tmp
		}
	}

	if udpOptions, ok := s.D.GetOkExists("udp_options"); ok {
		if tmpList := udpOptions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "udp_options", 0)
			tmp, err := s.mapToUdpOptions(fieldKeyFormat)
			if err != nil {
				return fmt.Errorf("unable to convert udp_options, encountered error: %v", err)
			}
			updateSecurityRuleDetails.UdpOptions = &tmp
		}
	}
	tmp := []oci_core.UpdateSecurityRuleDetails{updateSecurityRuleDetails}
	request.SecurityRules = tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")
	response, err := s.Client.UpdateNetworkSecurityGroupSecurityRules(context.Background(), request)
	if err != nil {
		return fmt.Errorf("failed to update security rules, error: %v", err)
	}
	if response.SecurityRules != nil && len(response.SecurityRules) > 0 {
		s.Res = &response.SecurityRules[0]
	}

	return nil
}

func (s *CoreSecurityRuleResourceCrud) Delete() error {

	request := oci_core.RemoveNetworkSecurityGroupSecurityRulesRequest{}
	if networkSecurityGroupId, ok := s.D.GetOkExists("network_security_group_id"); ok {
		tmp := networkSecurityGroupId.(string)
		request.NetworkSecurityGroupId = &tmp
	}

	tmp := []string{s.D.Id()}
	request.SecurityRuleIds = tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")
	_, err := s.Client.RemoveNetworkSecurityGroupSecurityRules(context.Background(), request)
	if err != nil {
		return err
	}
	return nil
}

func (s *CoreSecurityRuleResourceCrud) SetData() error {

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.Destination != nil {
		s.D.Set("destination", *s.Res.Destination)
	}

	s.D.Set("destination_type", s.Res.DestinationType)

	s.D.Set("direction", s.Res.Direction)

	if s.Res.IcmpOptions != nil {
		s.D.Set("icmp_options", []interface{}{nsgIcmpOptionsToMap(s.Res.IcmpOptions)})
	} else {
		s.D.Set("icmp_options", nil)
	}

	if s.Res.IsValid != nil {
		s.D.Set("is_valid", *s.Res.IsValid)
	}

	if s.Res.Protocol != nil {
		s.D.Set("protocol", *s.Res.Protocol)
	}

	if s.Res.Source != nil {
		s.D.Set("source", *s.Res.Source)
	}

	s.D.Set("source_type", s.Res.SourceType)

	if s.Res.IsStateless != nil {
		s.D.Set("stateless", *s.Res.IsStateless)
	}

	if s.Res.TcpOptions != nil {
		s.D.Set("tcp_options", []interface{}{nsgTcpOptionsToMap(s.Res.TcpOptions)})
	} else {
		s.D.Set("tcp_options", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.UdpOptions != nil {
		s.D.Set("udp_options", []interface{}{nsgUdpOptionsToMap(s.Res.UdpOptions)})
	} else {
		s.D.Set("udp_options", nil)
	}

	return nil
}

func (s *CoreSecurityRuleResourceCrud) mapToIcmpOptions(fieldKeyFormat string) (oci_core.IcmpOptions, error) {
	result := oci_core.IcmpOptions{}

	if code, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "code")); ok {
		tmp := code.(int)
		if tmp > -1 {
			result.Code = &tmp
		}
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		tmp := type_.(int)
		result.Type = &tmp
	}

	return result, nil
}

func (s *CoreSecurityRuleResourceCrud) mapToPortRange(fieldKeyFormat string) (oci_core.PortRange, error) {
	result := oci_core.PortRange{}

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

func (s *CoreSecurityRuleResourceCrud) mapToTcpOptions(fieldKeyFormat string) (oci_core.TcpOptions, error) {
	result := oci_core.TcpOptions{}

	if destinationPortRange, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_port_range")); ok {
		if tmpList := destinationPortRange.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "destination_port_range"), 0)
			tmp, err := s.mapToPortRange(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert destination_port_range, encountered error: %v", err)
			}
			result.DestinationPortRange = &tmp
		}
	}

	if sourcePortRange, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_port_range")); ok {
		if tmpList := sourcePortRange.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "source_port_range"), 0)
			tmp, err := s.mapToPortRange(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert source_port_range, encountered error: %v", err)
			}
			result.SourcePortRange = &tmp
		}
	}

	return result, nil
}

func (s *CoreSecurityRuleResourceCrud) mapToUdpOptions(fieldKeyFormat string) (oci_core.UdpOptions, error) {
	result := oci_core.UdpOptions{}

	if destinationPortRange, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_port_range")); ok {
		if tmpList := destinationPortRange.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "destination_port_range"), 0)
			tmp, err := s.mapToPortRange(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert destination_port_range, encountered error: %v", err)
			}
			result.DestinationPortRange = &tmp
		}
	}

	if sourcePortRange, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_port_range")); ok {
		if tmpList := sourcePortRange.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "source_port_range"), 0)
			tmp, err := s.mapToPortRange(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert source_port_range, encountered error: %v", err)
			}
			result.SourcePortRange = &tmp
		}
	}

	return result, nil
}

func nsgIcmpOptionsToMap(obj *oci_core.IcmpOptions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Code != nil {
		result["code"] = int(*obj.Code)
	} else {
		result["code"] = -1
	}

	if obj.Type != nil {
		result["type"] = int(*obj.Type)
	}

	return result
}

func nsgTcpOptionsToMap(obj *oci_core.TcpOptions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DestinationPortRange != nil {
		result["destination_port_range"] = []interface{}{nsgPortRangeToMap(obj.DestinationPortRange)}
	}

	if obj.SourcePortRange != nil {
		result["source_port_range"] = []interface{}{nsgPortRangeToMap(obj.SourcePortRange)}
	}

	return result
}

func nsgUdpOptionsToMap(obj *oci_core.UdpOptions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DestinationPortRange != nil {
		result["destination_port_range"] = []interface{}{nsgPortRangeToMap(obj.DestinationPortRange)}
	}

	if obj.SourcePortRange != nil {
		result["source_port_range"] = []interface{}{nsgPortRangeToMap(obj.SourcePortRange)}
	}

	return result
}

func nsgPortRangeToMap(obj *oci_core.PortRange) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Max != nil {
		result["max"] = int(*obj.Max)
	}

	if obj.Min != nil {
		result["min"] = int(*obj.Min)
	}

	return result
}

func GetNetworkSecurityGroupSecurityRuleCompositeId(networkSecurityGroupId string, securityRuleId string) string {
	networkSecurityGroupId = url.PathEscape(networkSecurityGroupId)
	securityRuleId = url.PathEscape(securityRuleId)
	compositeId := "networkSecurityGroups/" + networkSecurityGroupId + "/securityRules/" + securityRuleId
	return compositeId
}

func parseNetworkSecurityGroupSecurityRuleCompositeId(compositeId string) (networkSecurityGroupId string, securityRuleId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("networkSecurityGroups/.*/securityRules/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	networkSecurityGroupId, _ = url.PathUnescape(parts[1])
	securityRuleId, _ = url.PathUnescape(parts[3])

	return
}
