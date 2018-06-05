// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func SecurityListResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createSecurityList,
		Read:     readSecurityList,
		Update:   updateSecurityList,
		Delete:   deleteSecurityList,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"egress_security_rules": {
				Type: schema.TypeList,
				// Code-gen and specs say this should be required and has a max item limit
				// Keep it optional to continue to allow empty security rules and avoid a breaking change.
				// Also remove the max item limit, to avoid a potential breaking change.
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"destination": {
							Type:     schema.TypeString,
							Required: true,
						},
						"protocol": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"icmp_options": {
							Type:     schema.TypeList,
							Optional: true,
							// @CODEGEN 2/2018: This should not be a computed field as generated, as it breaks how Terraform
							// considers diffs when the number of icmp_options, tcp_options, and udp_options change.
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
										// @CODEGEN 2/2018: This is a workaround for Terraform setting this to 0 if not specified.
										// Since 0 is a valid 'code', we will define our own value (-1) to represent it
										// as being unset. This should ensure that not setting it here will also not set it
										// in the SDK request.
										Default: -1,
									},

									// Computed
								},
							},
						},
						"stateless": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"tcp_options": {
							Type:     schema.TypeList,
							Optional: true,
							// @CODEGEN 2/2018: This should not be a computed field as generated, as it breaks how Terraform
							// considers diffs when the number of icmp_options, tcp_options, and udp_options change.
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
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
									// Code-gen and specs say the following max and min should be under a destination_port_range schema
									// similar to source_port_range above.
									// We promoted it to the tcp_options schema to avoid a breaking change to how this is configured.
									// This is applied everywhere else in the schema where "max"/"min" should normally fall under destination_port_range.
									"max": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"min": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									// Computed
								},
							},
						},
						"udp_options": {
							Type:     schema.TypeList,
							Optional: true,
							// @CODEGEN 2/2018: This should not be a computed field as generated, as it breaks how Terraform
							// considers diffs when the number of icmp_options, tcp_options, and udp_options change.
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
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
									"max": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"min": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"ingress_security_rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"protocol": {
							Type:     schema.TypeString,
							Required: true,
						},
						"source": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"icmp_options": {
							Type:     schema.TypeList,
							Optional: true,
							// @CODEGEN 2/2018: This should not be a computed field as generated, as it breaks how Terraform
							// considers diffs when the number of icmp_options, tcp_options, and udp_options change.
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
										// @CODEGEN 2/2018: This is a workaround for Terraform setting this to 0 if not specified.
										// Since 0 is a valid 'code', we will define our own value (-1) to represent it
										// as being unset. This should ensure that not setting it here will also not set it
										// in the SDK request.
										Default: -1,
									},

									// Computed
								},
							},
						},
						"stateless": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"tcp_options": {
							Type:     schema.TypeList,
							Optional: true,
							// @CODEGEN 2/2018: This should not be a computed field as generated, as it breaks how Terraform
							// considers diffs when the number of icmp_options, tcp_options, and udp_options change.
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
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
									"max": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"min": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									// Computed
								},
							},
						},
						"udp_options": {
							Type:     schema.TypeList,
							Optional: true,
							// @CODEGEN 2/2018: This should not be a computed field as generated, as it breaks how Terraform
							// considers diffs when the number of icmp_options, tcp_options, and udp_options change.
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"source_port_range": {
										Type:     schema.TypeList,
										Optional: true,
										// @CODEGEN 2/2018: This should not be a computed field as generated, as it breaks how Terraform
										// considers diffs when the source_port_range is removed from config.
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
									"max": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"min": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
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
			"id": {
				Type:     schema.TypeString,
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
		},
	}
}

func createSecurityList(d *schema.ResourceData, m interface{}) error {
	sync := &SecurityListResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.CreateResource(d, sync)
}

func readSecurityList(d *schema.ResourceData, m interface{}) error {
	sync := &SecurityListResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

func updateSecurityList(d *schema.ResourceData, m interface{}) error {
	sync := &SecurityListResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.UpdateResource(d, sync)
}

func deleteSecurityList(d *schema.ResourceData, m interface{}) error {
	sync := &SecurityListResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type SecurityListResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.SecurityList
	DisableNotFoundRetries bool
}

func (s *SecurityListResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *SecurityListResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.SecurityListLifecycleStateProvisioning),
	}
}

func (s *SecurityListResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.SecurityListLifecycleStateAvailable),
	}
}

func (s *SecurityListResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.SecurityListLifecycleStateTerminating),
	}
}

func (s *SecurityListResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.SecurityListLifecycleStateTerminated),
	}
}

func (s *SecurityListResourceCrud) Create() error {
	request := oci_core.CreateSecurityListRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.EgressSecurityRules = []oci_core.EgressSecurityRule{}
	if egressSecurityRules, ok := s.D.GetOkExists("egress_security_rules"); ok {
		interfaces := egressSecurityRules.([]interface{})
		tmp := make([]oci_core.EgressSecurityRule, len(interfaces))
		for i, toBeConverted := range interfaces {
			tmp[i] = mapToEgressSecurityRule(toBeConverted.(map[string]interface{}))
		}
		request.EgressSecurityRules = tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.IngressSecurityRules = []oci_core.IngressSecurityRule{}
	if ingressSecurityRules, ok := s.D.GetOkExists("ingress_security_rules"); ok {
		interfaces := ingressSecurityRules.([]interface{})
		tmp := make([]oci_core.IngressSecurityRule, len(interfaces))
		for i, toBeConverted := range interfaces {
			tmp[i] = mapToIngressSecurityRule(toBeConverted.(map[string]interface{}))
		}
		request.IngressSecurityRules = tmp
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateSecurityList(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SecurityList
	return nil
}

func (s *SecurityListResourceCrud) Get() error {
	request := oci_core.GetSecurityListRequest{}

	tmp := s.D.Id()
	request.SecurityListId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetSecurityList(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SecurityList
	return nil
}

func (s *SecurityListResourceCrud) Update() error {
	request := oci_core.UpdateSecurityListRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.EgressSecurityRules = []oci_core.EgressSecurityRule{}
	if egressSecurityRules, ok := s.D.GetOkExists("egress_security_rules"); ok {
		interfaces := egressSecurityRules.([]interface{})
		tmp := make([]oci_core.EgressSecurityRule, len(interfaces))
		for i, toBeConverted := range interfaces {
			tmp[i] = mapToEgressSecurityRule(toBeConverted.(map[string]interface{}))
		}
		request.EgressSecurityRules = tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.IngressSecurityRules = []oci_core.IngressSecurityRule{}
	if ingressSecurityRules, ok := s.D.GetOkExists("ingress_security_rules"); ok {
		interfaces := ingressSecurityRules.([]interface{})
		tmp := make([]oci_core.IngressSecurityRule, len(interfaces))
		for i, toBeConverted := range interfaces {
			tmp[i] = mapToIngressSecurityRule(toBeConverted.(map[string]interface{}))
		}
		request.IngressSecurityRules = tmp
	}

	tmp := s.D.Id()
	request.SecurityListId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateSecurityList(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SecurityList
	return nil
}

func (s *SecurityListResourceCrud) Delete() error {
	request := oci_core.DeleteSecurityListRequest{}

	tmp := s.D.Id()
	request.SecurityListId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteSecurityList(context.Background(), request)
	return err
}

func (s *SecurityListResourceCrud) SetData() {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	egressSecurityRules := []interface{}{}
	for _, item := range s.Res.EgressSecurityRules {
		egressSecurityRules = append(egressSecurityRules, EgressSecurityRuleToMap(item))
	}
	s.D.Set("egress_security_rules", egressSecurityRules)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	ingressSecurityRules := []interface{}{}
	for _, item := range s.Res.IngressSecurityRules {
		ingressSecurityRules = append(ingressSecurityRules, IngressSecurityRuleToMap(item))
	}
	s.D.Set("ingress_security_rules", ingressSecurityRules)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

}

func mapToEgressSecurityRule(raw map[string]interface{}) oci_core.EgressSecurityRule {
	result := oci_core.EgressSecurityRule{}

	if destination, ok := raw["destination"]; ok && destination != "" {
		tmp := destination.(string)
		result.Destination = &tmp
	}

	if icmpOptions, ok := raw["icmp_options"]; ok {
		if tmpList := icmpOptions.([]interface{}); len(tmpList) > 0 {
			tmp := mapToIcmpOptions(tmpList[0].(map[string]interface{}))
			result.IcmpOptions = &tmp
		}
	}

	if protocol, ok := raw["protocol"]; ok && protocol != "" {
		tmp := protocol.(string)
		result.Protocol = &tmp
	}

	if stateless, ok := raw["stateless"]; ok {
		tmp := stateless.(bool)
		result.IsStateless = &tmp
	}

	if tcpOptions, ok := raw["tcp_options"]; ok {
		if tmpList := tcpOptions.([]interface{}); len(tmpList) > 0 {
			tmp := mapToTcpOptions(tmpList[0].(map[string]interface{}))
			result.TcpOptions = &tmp
		}
	}

	if udpOptions, ok := raw["udp_options"]; ok {
		if tmpList := udpOptions.([]interface{}); len(tmpList) > 0 {
			tmp := mapToUdpOptions(tmpList[0].(map[string]interface{}))
			result.UdpOptions = &tmp
		}
	}

	return result
}

func EgressSecurityRuleToMap(obj oci_core.EgressSecurityRule) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Destination != nil {
		result["destination"] = string(*obj.Destination)
	}

	if obj.IcmpOptions != nil {
		result["icmp_options"] = []interface{}{IcmpOptionsToMap(obj.IcmpOptions)}
	}

	if obj.Protocol != nil {
		result["protocol"] = string(*obj.Protocol)
	}

	if obj.IsStateless != nil {
		result["stateless"] = bool(*obj.IsStateless)
	}

	if obj.TcpOptions != nil {
		result["tcp_options"] = []interface{}{TcpOptionsToMap(obj.TcpOptions)}
	}

	if obj.UdpOptions != nil {
		result["udp_options"] = []interface{}{UdpOptionsToMap(obj.UdpOptions)}
	}

	return result
}

func mapToIcmpOptions(raw map[string]interface{}) oci_core.IcmpOptions {
	result := oci_core.IcmpOptions{}

	if code, ok := raw["code"]; ok {
		tmp := code.(int)
		if tmp != -1 {
			result.Code = &tmp
		}
	}

	if type_, ok := raw["type"]; ok {
		tmp := type_.(int)
		result.Type = &tmp
	}

	return result
}

func IcmpOptionsToMap(obj *oci_core.IcmpOptions) map[string]interface{} {
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

func mapToIngressSecurityRule(raw map[string]interface{}) oci_core.IngressSecurityRule {
	result := oci_core.IngressSecurityRule{}

	if icmpOptions, ok := raw["icmp_options"]; ok {
		if tmpList := icmpOptions.([]interface{}); len(tmpList) > 0 {
			tmp := mapToIcmpOptions(tmpList[0].(map[string]interface{}))
			result.IcmpOptions = &tmp
		}
	}

	if protocol, ok := raw["protocol"]; ok && protocol != "" {
		tmp := protocol.(string)
		result.Protocol = &tmp
	}

	if source, ok := raw["source"]; ok && source != "" {
		tmp := source.(string)
		result.Source = &tmp
	}

	if stateless, ok := raw["stateless"]; ok {
		tmp := stateless.(bool)
		result.IsStateless = &tmp
	}

	if tcpOptions, ok := raw["tcp_options"]; ok {
		if tmpList := tcpOptions.([]interface{}); len(tmpList) > 0 {
			tmp := mapToTcpOptions(tmpList[0].(map[string]interface{}))
			result.TcpOptions = &tmp
		}
	}

	if udpOptions, ok := raw["udp_options"]; ok {
		if tmpList := udpOptions.([]interface{}); len(tmpList) > 0 {
			tmp := mapToUdpOptions(tmpList[0].(map[string]interface{}))
			result.UdpOptions = &tmp
		}
	}

	return result
}

func IngressSecurityRuleToMap(obj oci_core.IngressSecurityRule) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IcmpOptions != nil {
		result["icmp_options"] = []interface{}{IcmpOptionsToMap(obj.IcmpOptions)}
	}

	if obj.Protocol != nil {
		result["protocol"] = string(*obj.Protocol)
	}

	if obj.Source != nil {
		result["source"] = string(*obj.Source)
	}

	if obj.IsStateless != nil {
		result["stateless"] = bool(*obj.IsStateless)
	}

	if obj.TcpOptions != nil {
		result["tcp_options"] = []interface{}{TcpOptionsToMap(obj.TcpOptions)}
	}

	if obj.UdpOptions != nil {
		result["udp_options"] = []interface{}{UdpOptionsToMap(obj.UdpOptions)}
	}

	return result
}

func mapToPortRange(raw map[string]interface{}) oci_core.PortRange {
	result := oci_core.PortRange{}

	if max, ok := raw["max"]; ok {
		tmp := max.(int)
		result.Max = &tmp
	}

	if min, ok := raw["min"]; ok {
		tmp := min.(int)
		result.Min = &tmp
	}

	return result
}

func PortRangeToMap(obj *oci_core.PortRange) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Max != nil {
		result["max"] = int(*obj.Max)
	}

	if obj.Min != nil {
		result["min"] = int(*obj.Min)
	}

	return result
}

func mapToTcpOptions(raw map[string]interface{}) oci_core.TcpOptions {
	result := oci_core.TcpOptions{}

	// Max and Min default to 0, and that is not a valid port number, so we can assume that if
	// the value is 0 then the user has not set the port number.
	// Also, note that if either max or min is set, then the service will return an error if both are not
	// set. However, we want to create the PortRange if either is set and let the service return the error.
	if raw["max"].(int) != 0 || raw["min"].(int) != 0 {
		tmp := mapToPortRange(raw)
		result.DestinationPortRange = &tmp
	}

	if sourcePortRange, ok := raw["source_port_range"]; ok {
		if tmpList := sourcePortRange.([]interface{}); len(tmpList) > 0 {
			tmp := mapToPortRange(tmpList[0].(map[string]interface{}))
			result.SourcePortRange = &tmp
		}
	}

	return result
}

func TcpOptionsToMap(obj *oci_core.TcpOptions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DestinationPortRange != nil {
		if obj.DestinationPortRange.Max != nil {
			result["max"] = *obj.DestinationPortRange.Max
		}

		if obj.DestinationPortRange.Min != nil {
			result["min"] = *obj.DestinationPortRange.Min
		}
	}

	if obj.SourcePortRange != nil {
		result["source_port_range"] = []interface{}{PortRangeToMap(obj.SourcePortRange)}
	}

	return result
}

func mapToUdpOptions(raw map[string]interface{}) oci_core.UdpOptions {
	result := oci_core.UdpOptions{}

	// Max and Min default to 0, and that is not a valid port number, so we can assume that if
	// the value is 0 then the user has not set the port number.
	// Also, note that if either max or min is set, then the service will return an error if both are not
	// set. However, we want to create the PortRange if either is set and let the service return the error.
	if raw["max"].(int) != 0 || raw["min"].(int) != 0 {
		tmp := mapToPortRange(raw)
		result.DestinationPortRange = &tmp
	}

	if sourcePortRange, ok := raw["source_port_range"]; ok {
		if tmpList := sourcePortRange.([]interface{}); len(tmpList) > 0 {
			tmp := mapToPortRange(tmpList[0].(map[string]interface{}))
			result.SourcePortRange = &tmp
		}
	}

	return result
}

func UdpOptionsToMap(obj *oci_core.UdpOptions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DestinationPortRange != nil {
		if obj.DestinationPortRange.Max != nil {
			result["max"] = *obj.DestinationPortRange.Max
		}

		if obj.DestinationPortRange.Min != nil {
			result["min"] = *obj.DestinationPortRange.Min
		}
	}

	if obj.SourcePortRange != nil {
		result["source_port_range"] = []interface{}{PortRangeToMap(obj.SourcePortRange)}
	}

	return result
}
