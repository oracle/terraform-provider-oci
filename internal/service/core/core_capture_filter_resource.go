// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreCaptureFilterResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCoreCaptureFilter,
		Read:     readCoreCaptureFilter,
		Update:   updateCoreCaptureFilter,
		Delete:   deleteCoreCaptureFilter,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"filter_type": {
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
			"flow_log_capture_filter_rules": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"destination_cidr": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"flow_log_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"icmp_options": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
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
										Computed: true,
									},

									// Computed
								},
							},
						},
						"is_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"priority": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"protocol": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rule_action": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sampling_rate": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"source_cidr": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tcp_options": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
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
						"udp_options": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
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
					},
				},
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"vtap_capture_filter_rules": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"traffic_direction": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"destination_cidr": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"icmp_options": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
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
										Computed: true,
									},

									// Computed
								},
							},
						},
						"protocol": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rule_action": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"source_cidr": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tcp_options": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
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
						"udp_options": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
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

func createCoreCaptureFilter(d *schema.ResourceData, m interface{}) error {
	sync := &CoreCaptureFilterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readCoreCaptureFilter(d *schema.ResourceData, m interface{}) error {
	sync := &CoreCaptureFilterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

func updateCoreCaptureFilter(d *schema.ResourceData, m interface{}) error {
	sync := &CoreCaptureFilterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteCoreCaptureFilter(d *schema.ResourceData, m interface{}) error {
	sync := &CoreCaptureFilterResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type CoreCaptureFilterResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_core.VirtualNetworkClient
	Res                    *oci_core.CaptureFilter
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *CoreCaptureFilterResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CoreCaptureFilterResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_core.CaptureFilterLifecycleStateProvisioning),
	}
}

func (s *CoreCaptureFilterResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_core.CaptureFilterLifecycleStateAvailable),
	}
}

func (s *CoreCaptureFilterResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_core.CaptureFilterLifecycleStateTerminating),
	}
}

func (s *CoreCaptureFilterResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_core.CaptureFilterLifecycleStateTerminated),
	}
}

func (s *CoreCaptureFilterResourceCrud) Create() error {
	request := oci_core.CreateCaptureFilterRequest{}

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

	if filterType, ok := s.D.GetOkExists("filter_type"); ok {
		request.FilterType = oci_core.CreateCaptureFilterDetailsFilterTypeEnum(filterType.(string))
	}

	if flowLogCaptureFilterRules, ok := s.D.GetOkExists("flow_log_capture_filter_rules"); ok {
		interfaces := flowLogCaptureFilterRules.([]interface{})
		tmp := make([]oci_core.FlowLogCaptureFilterRuleDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "flow_log_capture_filter_rules", stateDataIndex)
			converted, err := s.mapToFlowLogCaptureFilterRuleDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("flow_log_capture_filter_rules") {
			request.FlowLogCaptureFilterRules = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if vtapCaptureFilterRules, ok := s.D.GetOkExists("vtap_capture_filter_rules"); ok {
		interfaces := vtapCaptureFilterRules.([]interface{})
		tmp := make([]oci_core.VtapCaptureFilterRuleDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "vtap_capture_filter_rules", stateDataIndex)
			converted, err := s.mapToVtapCaptureFilterRuleDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("vtap_capture_filter_rules") {
			request.VtapCaptureFilterRules = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.CreateCaptureFilter(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CaptureFilter
	return nil
}

func (s *CoreCaptureFilterResourceCrud) Get() error {
	request := oci_core.GetCaptureFilterRequest{}

	tmp := s.D.Id()
	request.CaptureFilterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.GetCaptureFilter(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CaptureFilter
	return nil
}

func (s *CoreCaptureFilterResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_core.UpdateCaptureFilterRequest{}

	tmp := s.D.Id()
	request.CaptureFilterId = &tmp

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

	if flowLogCaptureFilterRules, ok := s.D.GetOkExists("flow_log_capture_filter_rules"); ok {
		interfaces := flowLogCaptureFilterRules.([]interface{})
		tmp := make([]oci_core.FlowLogCaptureFilterRuleDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "flow_log_capture_filter_rules", stateDataIndex)
			converted, err := s.mapToFlowLogCaptureFilterRuleDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("flow_log_capture_filter_rules") {
			request.FlowLogCaptureFilterRules = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if vtapCaptureFilterRules, ok := s.D.GetOkExists("vtap_capture_filter_rules"); ok {
		interfaces := vtapCaptureFilterRules.([]interface{})
		tmp := make([]oci_core.VtapCaptureFilterRuleDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "vtap_capture_filter_rules", stateDataIndex)
			converted, err := s.mapToVtapCaptureFilterRuleDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("vtap_capture_filter_rules") {
			request.VtapCaptureFilterRules = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateCaptureFilter(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.CaptureFilter
	return nil
}

func (s *CoreCaptureFilterResourceCrud) Delete() error {
	request := oci_core.DeleteCaptureFilterRequest{}

	tmp := s.D.Id()
	request.CaptureFilterId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	_, err := s.Client.DeleteCaptureFilter(context.Background(), request)
	return err
}

func (s *CoreCaptureFilterResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("filter_type", s.Res.FilterType)

	flowLogCaptureFilterRules := []interface{}{}
	for _, item := range s.Res.FlowLogCaptureFilterRules {
		flowLogCaptureFilterRules = append(flowLogCaptureFilterRules, FlowLogCaptureFilterRuleDetailsToMap(item))
	}
	s.D.Set("flow_log_capture_filter_rules", flowLogCaptureFilterRules)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	vtapCaptureFilterRules := []interface{}{}
	for _, item := range s.Res.VtapCaptureFilterRules {
		vtapCaptureFilterRules = append(vtapCaptureFilterRules, VtapCaptureFilterRuleDetailsToMap(item))
	}
	s.D.Set("vtap_capture_filter_rules", vtapCaptureFilterRules)

	return nil
}

func (s *CoreCaptureFilterResourceCrud) mapToFlowLogCaptureFilterRuleDetails(fieldKeyFormat string) (oci_core.FlowLogCaptureFilterRuleDetails, error) {
	result := oci_core.FlowLogCaptureFilterRuleDetails{}

	if destinationCidr, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_cidr")); ok {
		tmp := destinationCidr.(string)
		result.DestinationCidr = &tmp
	}

	if flowLogType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "flow_log_type")); ok {
		result.FlowLogType = oci_core.FlowLogCaptureFilterRuleDetailsFlowLogTypeEnum(flowLogType.(string))
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

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	if priority, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "priority")); ok {
		tmp := priority.(int)
		result.Priority = &tmp
	}

	if protocol, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protocol")); ok {
		tmp := protocol.(string)
		result.Protocol = &tmp
	}

	if ruleAction, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rule_action")); ok {
		result.RuleAction = oci_core.FlowLogCaptureFilterRuleDetailsRuleActionEnum(ruleAction.(string))
	}

	if samplingRate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sampling_rate")); ok {
		tmp := samplingRate.(int)
		result.SamplingRate = &tmp
	}

	if sourceCidr, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_cidr")); ok {
		tmp := sourceCidr.(string)
		result.SourceCidr = &tmp
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

func FlowLogCaptureFilterRuleDetailsToMap(obj oci_core.FlowLogCaptureFilterRuleDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DestinationCidr != nil {
		result["destination_cidr"] = string(*obj.DestinationCidr)
	}

	result["flow_log_type"] = string(obj.FlowLogType)

	if obj.IcmpOptions != nil {
		result["icmp_options"] = []interface{}{IcmpOptionsToMap(obj.IcmpOptions)}
	}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	if obj.Priority != nil {
		result["priority"] = int(*obj.Priority)
	}

	if obj.Protocol != nil {
		result["protocol"] = string(*obj.Protocol)
	}

	result["rule_action"] = string(obj.RuleAction)

	if obj.SamplingRate != nil {
		result["sampling_rate"] = int(*obj.SamplingRate)
	}

	if obj.SourceCidr != nil {
		result["source_cidr"] = string(*obj.SourceCidr)
	}

	if obj.TcpOptions != nil {
		result["tcp_options"] = []interface{}{TcpOptionsToMap(obj.TcpOptions)}
	}

	if obj.UdpOptions != nil {
		result["udp_options"] = []interface{}{UdpOptionsToMap(obj.UdpOptions)}
	}

	return result
}

func (s *CoreCaptureFilterResourceCrud) mapToIcmpOptions(fieldKeyFormat string) (oci_core.IcmpOptions, error) {
	result := oci_core.IcmpOptions{}

	if code, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "code")); ok {
		tmp := code.(int)
		result.Code = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		tmp := type_.(int)
		result.Type = &tmp
	}

	return result, nil
}

func captureFilterIcmpOptionsToMap(obj *oci_core.IcmpOptions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Code != nil {
		result["code"] = int(*obj.Code)
	}

	if obj.Type != nil {
		result["type"] = int(*obj.Type)
	}

	return result
}

func (s *CoreCaptureFilterResourceCrud) mapToPortRange(fieldKeyFormat string) (oci_core.PortRange, error) {
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

func captureFilterPortRangeToMap(obj *oci_core.PortRange) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Max != nil {
		result["max"] = int(*obj.Max)
	}

	if obj.Min != nil {
		result["min"] = int(*obj.Min)
	}

	return result
}

func (s *CoreCaptureFilterResourceCrud) mapToTcpOptions(fieldKeyFormat string) (oci_core.TcpOptions, error) {
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

func captureFilterTcpOptionsToMap(obj *oci_core.TcpOptions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DestinationPortRange != nil {
		result["destination_port_range"] = []interface{}{captureFilterPortRangeToMap(obj.DestinationPortRange)}
	}

	if obj.SourcePortRange != nil {
		result["source_port_range"] = []interface{}{captureFilterPortRangeToMap(obj.SourcePortRange)}
	}

	return result
}

func (s *CoreCaptureFilterResourceCrud) mapToUdpOptions(fieldKeyFormat string) (oci_core.UdpOptions, error) {
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

func captureFilterUdpOptionsToMap(obj *oci_core.UdpOptions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DestinationPortRange != nil {
		result["destination_port_range"] = []interface{}{captureFilterPortRangeToMap(obj.DestinationPortRange)}
	}

	if obj.SourcePortRange != nil {
		result["source_port_range"] = []interface{}{captureFilterPortRangeToMap(obj.SourcePortRange)}
	}

	return result
}

func (s *CoreCaptureFilterResourceCrud) mapToVtapCaptureFilterRuleDetails(fieldKeyFormat string) (oci_core.VtapCaptureFilterRuleDetails, error) {
	result := oci_core.VtapCaptureFilterRuleDetails{}

	if destinationCidr, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_cidr")); ok {
		tmp := destinationCidr.(string)
		result.DestinationCidr = &tmp
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

	if ruleAction, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rule_action")); ok {
		result.RuleAction = oci_core.VtapCaptureFilterRuleDetailsRuleActionEnum(ruleAction.(string))
	}

	if sourceCidr, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_cidr")); ok {
		tmp := sourceCidr.(string)
		result.SourceCidr = &tmp
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

	if trafficDirection, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "traffic_direction")); ok {
		result.TrafficDirection = oci_core.VtapCaptureFilterRuleDetailsTrafficDirectionEnum(trafficDirection.(string))
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

func VtapCaptureFilterRuleDetailsToMap(obj oci_core.VtapCaptureFilterRuleDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DestinationCidr != nil {
		result["destination_cidr"] = string(*obj.DestinationCidr)
	}

	if obj.IcmpOptions != nil {
		result["icmp_options"] = []interface{}{captureFilterIcmpOptionsToMap(obj.IcmpOptions)}
	}

	if obj.Protocol != nil {
		result["protocol"] = string(*obj.Protocol)
	}

	result["rule_action"] = string(obj.RuleAction)

	if obj.SourceCidr != nil {
		result["source_cidr"] = string(*obj.SourceCidr)
	}

	if obj.TcpOptions != nil {
		result["tcp_options"] = []interface{}{captureFilterTcpOptionsToMap(obj.TcpOptions)}
	}

	result["traffic_direction"] = string(obj.TrafficDirection)

	if obj.UdpOptions != nil {
		result["udp_options"] = []interface{}{captureFilterUdpOptionsToMap(obj.UdpOptions)}
	}

	return result
}

func (s *CoreCaptureFilterResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_core.ChangeCaptureFilterCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.CaptureFilterId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.ChangeCaptureFilterCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "captureFilter", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	return nil
}
