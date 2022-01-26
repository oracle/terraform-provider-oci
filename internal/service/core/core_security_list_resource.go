// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"bytes"
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	oci_core "github.com/oracle/oci-go-sdk/v56/core"
)

func CoreSecurityListResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreSecurityList,
		Read:     readCoreSecurityList,
		Update:   updateCoreSecurityList,
		Delete:   deleteCoreSecurityList,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
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
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"egress_security_rules": {
				Type: schema.TypeSet,
				// Code-gen and specs say this should be required and has a max item limit
				// Keep it optional to continue to allow empty security rules and avoid a breaking change.
				// Also remove the max item limit, to avoid a potential breaking change.
				Optional: true,
				Set:      egressSecurityRulesHashCodeForSets,
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
						"description": {
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"ingress_security_rules": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      ingressSecurityRulesHashCodeForSets,
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
						"description": {
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
										// @CODEGEN 2/2018: This is a workaround for Terraform setting this to 0 if not specified.
										// Since 0 is a valid 'code', we will define our own value (-1) to represent it
										// as being unset. This should ensure that not setting it here will also not set it
										// in the SDK request.
									},

									// Computed
								},
							},
						},
						"source_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"stateless": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
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

			// Computed
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

func createCoreSecurityList(d *schema.ResourceData, m interface{}) error {
	sync := &CoreSecurityListResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.CreateResource(d, sync)
}

func readCoreSecurityList(d *schema.ResourceData, m interface{}) error {
	sync := &CoreSecurityListResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

func updateCoreSecurityList(d *schema.ResourceData, m interface{}) error {
	sync := &CoreSecurityListResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreSecurityList(d *schema.ResourceData, m interface{}) error {
	sync := &CoreSecurityListResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CoreSecurityListResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.SecurityList
	DisableNotFoundRetries bool
}

func (s *CoreSecurityListResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreSecurityListResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.SecurityListLifecycleStateProvisioning),
	}
}

func (s *CoreSecurityListResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.SecurityListLifecycleStateAvailable),
	}
}

func (s *CoreSecurityListResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.SecurityListLifecycleStateTerminating),
	}
}

func (s *CoreSecurityListResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.SecurityListLifecycleStateTerminated),
	}
}

func (s *CoreSecurityListResourceCrud) Create() error {
	request := oci_core.CreateSecurityListRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
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
		set := egressSecurityRules.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_core.EgressSecurityRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := egressSecurityRulesHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "egress_security_rules", stateDataIndex)
			converted, err := s.mapToEgressSecurityRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		request.EgressSecurityRules = tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.IngressSecurityRules = []oci_core.IngressSecurityRule{}
	if ingressSecurityRules, ok := s.D.GetOkExists("ingress_security_rules"); ok {
		set := ingressSecurityRules.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_core.IngressSecurityRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := ingressSecurityRulesHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "ingress_security_rules", stateDataIndex)
			converted, err := s.mapToIngressSecurityRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		request.IngressSecurityRules = tmp
	}

	if vcnId, ok := s.D.GetOkExists("vcn_id"); ok {
		tmp := vcnId.(string)
		request.VcnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateSecurityList(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SecurityList
	return nil
}

func (s *CoreSecurityListResourceCrud) Get() error {
	request := oci_core.GetSecurityListRequest{}

	tmp := s.D.Id()
	request.SecurityListId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetSecurityList(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SecurityList
	return nil
}

func (s *CoreSecurityListResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateSecurityListRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
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
		set := egressSecurityRules.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_core.EgressSecurityRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := egressSecurityRulesHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "egress_security_rules", stateDataIndex)
			converted, err := s.mapToEgressSecurityRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		request.EgressSecurityRules = tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.IngressSecurityRules = []oci_core.IngressSecurityRule{}
	if ingressSecurityRules, ok := s.D.GetOkExists("ingress_security_rules"); ok {
		set := ingressSecurityRules.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_core.IngressSecurityRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := ingressSecurityRulesHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "ingress_security_rules", stateDataIndex)
			converted, err := s.mapToIngressSecurityRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		request.IngressSecurityRules = tmp
	}

	tmp := s.D.Id()
	request.SecurityListId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateSecurityList(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SecurityList
	return nil
}

func (s *CoreSecurityListResourceCrud) Delete() error {
	request := oci_core.DeleteSecurityListRequest{}

	tmp := s.D.Id()
	request.SecurityListId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteSecurityList(context.Background(), request)
	return err
}

func (s *CoreSecurityListResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	egressSecurityRules := []interface{}{}
	for _, item := range s.Res.EgressSecurityRules {
		egressSecurityRules = append(egressSecurityRules, EgressSecurityRuleToMap(item))
	}
	s.D.Set("egress_security_rules", schema.NewSet(egressSecurityRulesHashCodeForSets, egressSecurityRules))

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	ingressSecurityRules := []interface{}{}
	for _, item := range s.Res.IngressSecurityRules {
		ingressSecurityRules = append(ingressSecurityRules, IngressSecurityRuleToMap(item))
	}
	s.D.Set("ingress_security_rules", schema.NewSet(ingressSecurityRulesHashCodeForSets, ingressSecurityRules))

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	return nil
}

func (s *CoreSecurityListResourceCrud) mapToEgressSecurityRule(fieldKeyFormat string) (oci_core.EgressSecurityRule, error) {
	result := oci_core.EgressSecurityRule{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok && description != "" {
		tmp := description.(string)
		result.Description = &tmp
	}

	if destination, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination")); ok && destination != "" {
		tmp := destination.(string)
		result.Destination = &tmp
	}

	if destinationType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_type")); ok && destinationType != "" {
		result.DestinationType = oci_core.EgressSecurityRuleDestinationTypeEnum(destinationType.(string))
	}

	if icmpOptions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "icmp_options")); ok {
		if tmpList := icmpOptions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "icmp_options"), 0)
			tmp, err := s.mapToIcmpOptions(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert icmp_options, encountered error: %v", err)
			}
			result.IcmpOptions = &tmp
		}
	}

	if protocol, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protocol")); ok && protocol != "" {
		tmp := protocol.(string)
		result.Protocol = &tmp
	}

	if stateless, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "stateless")); ok {
		tmp := stateless.(bool)
		result.IsStateless = &tmp
	}

	if tcpOptions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tcp_options")); ok {
		if tmpList := tcpOptions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "tcp_options"), 0)
			tmp, err := s.mapToTcpOptions(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert tcp_options, encountered error: %v", err)
			}
			result.TcpOptions = &tmp
		}
	}

	if udpOptions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "udp_options")); ok {
		if tmpList := udpOptions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "udp_options"), 0)
			tmp, err := s.mapToUdpOptions(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert udp_options, encountered error: %v", err)
			}
			result.UdpOptions = &tmp
		}
	}

	return result, nil
}

func EgressSecurityRuleToMap(obj oci_core.EgressSecurityRule) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Destination != nil {
		result["destination"] = string(*obj.Destination)
	}

	result["destination_type"] = string(obj.DestinationType)

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

func (s *CoreSecurityListResourceCrud) mapToIcmpOptions(fieldKeyFormat string) (oci_core.IcmpOptions, error) {
	result := oci_core.IcmpOptions{}

	if code, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "code")); ok {
		tmp := code.(int)
		if tmp != -1 {
			result.Code = &tmp
		}
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		tmp := type_.(int)
		result.Type = &tmp
	}

	return result, nil
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

func (s *CoreSecurityListResourceCrud) mapToIngressSecurityRule(fieldKeyFormat string) (oci_core.IngressSecurityRule, error) {
	result := oci_core.IngressSecurityRule{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok && description != "" {
		tmp := description.(string)
		result.Description = &tmp
	}

	if icmpOptions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "icmp_options")); ok {
		if tmpList := icmpOptions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "icmp_options"), 0)
			tmp, err := s.mapToIcmpOptions(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert icmp_options, encountered error: %v", err)
			}
			result.IcmpOptions = &tmp
		}
	}

	if protocol, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protocol")); ok {
		tmp := protocol.(string)
		result.Protocol = &tmp
	}

	if source, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source")); ok {
		tmp := source.(string)
		result.Source = &tmp
	}

	if sourceType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_type")); ok {
		result.SourceType = oci_core.IngressSecurityRuleSourceTypeEnum(sourceType.(string))
	}

	if stateless, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "stateless")); ok {
		tmp := stateless.(bool)
		result.IsStateless = &tmp
	}

	if tcpOptions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tcp_options")); ok {
		if tmpList := tcpOptions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "tcp_options"), 0)
			tmp, err := s.mapToTcpOptions(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert tcp_options, encountered error: %v", err)
			}
			result.TcpOptions = &tmp
		}
	}

	if udpOptions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "udp_options")); ok {
		if tmpList := udpOptions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "udp_options"), 0)
			tmp, err := s.mapToUdpOptions(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert udp_options, encountered error: %v", err)
			}
			result.UdpOptions = &tmp
		}
	}

	return result, nil
}

func IngressSecurityRuleToMap(obj oci_core.IngressSecurityRule) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.IcmpOptions != nil {
		result["icmp_options"] = []interface{}{IcmpOptionsToMap(obj.IcmpOptions)}
	}

	if obj.Protocol != nil {
		result["protocol"] = string(*obj.Protocol)
	}

	if obj.Source != nil {
		result["source"] = string(*obj.Source)
	}

	result["source_type"] = string(obj.SourceType)

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

func (s *CoreSecurityListResourceCrud) mapToPortRange(fieldKeyFormat string) (oci_core.PortRange, error) {
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

func (s *CoreSecurityListResourceCrud) mapToTcpOptions(fieldKeyFormat string) (oci_core.TcpOptions, error) {
	result := oci_core.TcpOptions{}

	max, maxExists := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max"))
	min, minExists := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "min"))
	if (maxExists && max.(int) != 0) || (minExists && min.(int) != 0) {
		tmp, err := s.mapToPortRange(fieldKeyFormat)
		if err != nil {
			return result, fmt.Errorf("unable to convert destination_port_range, encountered error: %v", err)
		}
		result.DestinationPortRange = &tmp
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

func (s *CoreSecurityListResourceCrud) mapToUdpOptions(fieldKeyFormat string) (oci_core.UdpOptions, error) {
	result := oci_core.UdpOptions{}

	max, maxExists := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max"))
	min, minExists := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "min"))
	if (maxExists && max.(int) != 0) || (minExists && min.(int) != 0) {
		tmp, err := s.mapToPortRange(fieldKeyFormat)
		if err != nil {
			return result, fmt.Errorf("unable to convert destination_port_range, encountered error: %v", err)
		}
		result.DestinationPortRange = &tmp
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

func egressSecurityRulesHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if description, ok := m["description"]; ok && description != "" {
		buf.WriteString(fmt.Sprintf("%v-", description))
	}
	if destination, ok := m["destination"]; ok && destination != "" {
		buf.WriteString(fmt.Sprintf("%v-", destination))
	}
	if destinationType, ok := m["destination_type"]; ok && destinationType != "" {
		buf.WriteString(fmt.Sprintf("%v-", destinationType))
	} else {
		buf.WriteString(fmt.Sprintf("%v-", oci_core.EgressSecurityRuleDestinationTypeCidrBlock))
	}
	if icmpOptions, ok := m["icmp_options"]; ok {
		if tmpList := icmpOptions.([]interface{}); len(tmpList) > 0 && tmpList[0] != nil {
			buf.WriteString("icmp_options-")
			icmpOptionsRaw := tmpList[0].(map[string]interface{})
			if code, ok := icmpOptionsRaw["code"]; ok {
				buf.WriteString(fmt.Sprintf("%v-", code))
			}
			if type_, ok := icmpOptionsRaw["type"]; ok {
				buf.WriteString(fmt.Sprintf("%v-", type_))
			}
		}
	}
	if protocol, ok := m["protocol"]; ok && protocol != "" {
		buf.WriteString(fmt.Sprintf("%v-", protocol))
	}
	if stateless, ok := m["stateless"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", stateless))
	} else {
		buf.WriteString(fmt.Sprintf("%v-", "false"))
	}
	if tcpOptions, ok := m["tcp_options"]; ok {
		if tmpList := tcpOptions.([]interface{}); len(tmpList) > 0 && tmpList[0] != nil {
			buf.WriteString("tcp_options-")
			tcpOptionsRaw := tmpList[0].(map[string]interface{})
			if max, ok := tcpOptionsRaw["max"]; ok {
				buf.WriteString(fmt.Sprintf("%v-", max))
			}
			if min, ok := tcpOptionsRaw["min"]; ok {
				buf.WriteString(fmt.Sprintf("%v-", min))
			}

			if sourcePortRange, ok := tcpOptionsRaw["source_port_range"]; ok {
				if tmpList := sourcePortRange.([]interface{}); len(tmpList) > 0 {
					buf.WriteString("source_port_range-")
					sourcePortRangeRaw := tmpList[0].(map[string]interface{})
					if max, ok := sourcePortRangeRaw["max"]; ok {
						buf.WriteString(fmt.Sprintf("%v-", max))
					}
					if min, ok := sourcePortRangeRaw["min"]; ok {
						buf.WriteString(fmt.Sprintf("%v-", min))
					}
				}
			}
		}
	}
	if udpOptions, ok := m["udp_options"]; ok {
		if tmpList := udpOptions.([]interface{}); len(tmpList) > 0 && tmpList[0] != nil {
			buf.WriteString("udp_options-")
			udpOptionsRaw := tmpList[0].(map[string]interface{})
			if max, ok := udpOptionsRaw["max"]; ok && max != 0 {
				buf.WriteString(fmt.Sprintf("%v-", max))
			}
			if min, ok := udpOptionsRaw["min"]; ok && min != 0 {
				buf.WriteString(fmt.Sprintf("%v-", min))
			}

			if sourcePortRange, ok := udpOptionsRaw["source_port_range"]; ok {
				if tmpList := sourcePortRange.([]interface{}); len(tmpList) > 0 {
					buf.WriteString("source_port_range-")
					sourcePortRangeRaw := tmpList[0].(map[string]interface{})
					if max, ok := sourcePortRangeRaw["max"]; ok {
						buf.WriteString(fmt.Sprintf("%v-", max))
					}
					if min, ok := sourcePortRangeRaw["min"]; ok {
						buf.WriteString(fmt.Sprintf("%v-", min))
					}
				}
			}
		}
	}
	return hashcode.String(buf.String())
}

func ingressSecurityRulesHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if description, ok := m["description"]; ok && description != "" {
		buf.WriteString(fmt.Sprintf("%v-", description))
	}
	if icmpOptions, ok := m["icmp_options"]; ok {
		if tmpList := icmpOptions.([]interface{}); len(tmpList) > 0 && tmpList[0] != nil {
			buf.WriteString("icmp_options-")
			icmpOptionsRaw := tmpList[0].(map[string]interface{})
			if code, ok := icmpOptionsRaw["code"]; ok {
				buf.WriteString(fmt.Sprintf("%v-", code))
			}
			if type_, ok := icmpOptionsRaw["type"]; ok {
				buf.WriteString(fmt.Sprintf("%v-", type_))
			}
		}
	}
	if protocol, ok := m["protocol"]; ok && protocol != "" {
		buf.WriteString(fmt.Sprintf("%v-", protocol))
	}
	if source, ok := m["source"]; ok && source != "" {
		buf.WriteString(fmt.Sprintf("%v-", source))
	}
	if sourceType, ok := m["source_type"]; ok && sourceType != "" {
		buf.WriteString(fmt.Sprintf("%v-", sourceType))
	} else {
		buf.WriteString(fmt.Sprintf("%v-", oci_core.IngressSecurityRuleSourceTypeCidrBlock))
	}
	if stateless, ok := m["stateless"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", stateless))
	} else {
		buf.WriteString(fmt.Sprintf("%v-", "false"))
	}
	if tcpOptions, ok := m["tcp_options"]; ok {
		if tmpList := tcpOptions.([]interface{}); len(tmpList) > 0 && tmpList[0] != nil {
			buf.WriteString("tcp_options-")
			tcpOptionsRaw := tmpList[0].(map[string]interface{})
			if max, ok := tcpOptionsRaw["max"]; ok {
				buf.WriteString(fmt.Sprintf("%v-", max))
			}
			if min, ok := tcpOptionsRaw["min"]; ok {
				buf.WriteString(fmt.Sprintf("%v-", min))
			}
			if sourcePortRange, ok := tcpOptionsRaw["source_port_range"]; ok {
				if tmpList := sourcePortRange.([]interface{}); len(tmpList) > 0 {
					buf.WriteString("source_port_range-")
					sourcePortRangeRaw := tmpList[0].(map[string]interface{})
					if max, ok := sourcePortRangeRaw["max"]; ok {
						buf.WriteString(fmt.Sprintf("%v-", max))
					}
					if min, ok := sourcePortRangeRaw["min"]; ok {
						buf.WriteString(fmt.Sprintf("%v-", min))
					}
				}
			}
		}
	}
	if udpOptions, ok := m["udp_options"]; ok {
		if tmpList := udpOptions.([]interface{}); len(tmpList) > 0 && tmpList[0] != nil {
			buf.WriteString("udp_options-")
			udpOptionsRaw := tmpList[0].(map[string]interface{})
			if max, ok := udpOptionsRaw["max"]; ok {
				buf.WriteString(fmt.Sprintf("%v-", max))
			}
			if min, ok := udpOptionsRaw["min"]; ok {
				buf.WriteString(fmt.Sprintf("%v-", min))
			}

			if sourcePortRange, ok := udpOptionsRaw["source_port_range"]; ok {
				if tmpList := sourcePortRange.([]interface{}); len(tmpList) > 0 {
					buf.WriteString("source_port_range-")
					sourcePortRangeRaw := tmpList[0].(map[string]interface{})
					if max, ok := sourcePortRangeRaw["max"]; ok {
						buf.WriteString(fmt.Sprintf("%v-", max))
					}
					if min, ok := sourcePortRangeRaw["min"]; ok {
						buf.WriteString(fmt.Sprintf("%v-", min))
					}
				}
			}
		}
	}
	return hashcode.String(buf.String())
}
func (s *CoreSecurityListResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeSecurityListCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.SecurityListId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.ChangeSecurityListCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
