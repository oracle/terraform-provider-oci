// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package vn_monitoring

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_vn_monitoring "github.com/oracle/oci-go-sdk/v65/vnmonitoring"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func VnMonitoringPathAnalyzerTestResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createVnMonitoringPathAnalyzerTest,
		Read:     readVnMonitoringPathAnalyzerTest,
		Update:   updateVnMonitoringPathAnalyzerTest,
		Delete:   deleteVnMonitoringPathAnalyzerTest,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"destination_endpoint": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"COMPUTE_INSTANCE",
								"IP_ADDRESS",
								"LOAD_BALANCER",
								"LOAD_BALANCER_LISTENER",
								"NETWORK_LOAD_BALANCER",
								"NETWORK_LOAD_BALANCER_LISTENER",
								"SUBNET",
								"VLAN",
								"VNIC",
							}, true),
						},

						// Optional
						"address": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"instance_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"listener_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"load_balancer_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"network_load_balancer_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vnic_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						// internal for work request access
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"protocol": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"source_endpoint": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"COMPUTE_INSTANCE",
								"IP_ADDRESS",
								"LOAD_BALANCER",
								"LOAD_BALANCER_LISTENER",
								"NETWORK_LOAD_BALANCER",
								"NETWORK_LOAD_BALANCER_LISTENER",
								"SUBNET",
								"VLAN",
								"VNIC",
							}, true),
						},

						// Optional
						"address": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"instance_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"listener_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"load_balancer_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"network_load_balancer_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vnic_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						// internal for work request access
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"protocol_parameters": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"ICMP",
								"TCP",
								"UDP",
							}, true),
						},

						// Optional
						"destination_port": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"icmp_code": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"icmp_type": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"source_port": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"query_options": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"is_bi_directional_analysis": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
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
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createVnMonitoringPathAnalyzerTest(d *schema.ResourceData, m interface{}) error {
	sync := &VnMonitoringPathAnalyzerTestResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VnMonitoringClient()

	return tfresource.CreateResource(d, sync)
}

func readVnMonitoringPathAnalyzerTest(d *schema.ResourceData, m interface{}) error {
	sync := &VnMonitoringPathAnalyzerTestResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VnMonitoringClient()

	return tfresource.ReadResource(sync)
}

func updateVnMonitoringPathAnalyzerTest(d *schema.ResourceData, m interface{}) error {
	sync := &VnMonitoringPathAnalyzerTestResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VnMonitoringClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteVnMonitoringPathAnalyzerTest(d *schema.ResourceData, m interface{}) error {
	sync := &VnMonitoringPathAnalyzerTestResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VnMonitoringClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type VnMonitoringPathAnalyzerTestResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_vn_monitoring.VnMonitoringClient
	Res                    *oci_vn_monitoring.PathAnalyzerTest
	DisableNotFoundRetries bool
}

func (s *VnMonitoringPathAnalyzerTestResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *VnMonitoringPathAnalyzerTestResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *VnMonitoringPathAnalyzerTestResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_vn_monitoring.PathAnalyzerTestLifecycleStateActive),
	}
}

func (s *VnMonitoringPathAnalyzerTestResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *VnMonitoringPathAnalyzerTestResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_vn_monitoring.PathAnalyzerTestLifecycleStateDeleted),
	}
}

func (s *VnMonitoringPathAnalyzerTestResourceCrud) Create() error {
	request := oci_vn_monitoring.CreatePathAnalyzerTestRequest{}

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

	if destinationEndpoint, ok := s.D.GetOkExists("destination_endpoint"); ok {
		if tmpList := destinationEndpoint.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "destination_endpoint", 0)
			tmp, err := s.mapToEndpoint(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DestinationEndpoint = tmp
		}
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if protocol, ok := s.D.GetOkExists("protocol"); ok {
		tmp := protocol.(int)
		request.Protocol = &tmp
	}

	if protocolParameters, ok := s.D.GetOkExists("protocol_parameters"); ok {
		if tmpList := protocolParameters.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "protocol_parameters", 0)
			tmp, err := s.mapToProtocolParameters(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ProtocolParameters = tmp
		}
	}

	if queryOptions, ok := s.D.GetOkExists("query_options"); ok {
		if tmpList := queryOptions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "query_options", 0)
			tmp, err := s.mapToQueryOptions(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.QueryOptions = &tmp
		}
	}

	if sourceEndpoint, ok := s.D.GetOkExists("source_endpoint"); ok {
		if tmpList := sourceEndpoint.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source_endpoint", 0)
			tmp, err := s.mapToEndpoint(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SourceEndpoint = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "vn_monitoring")

	response, err := s.Client.CreatePathAnalyzerTest(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PathAnalyzerTest
	return nil
}

func (s *VnMonitoringPathAnalyzerTestResourceCrud) Get() error {
	request := oci_vn_monitoring.GetPathAnalyzerTestRequest{}

	tmp := s.D.Id()
	request.PathAnalyzerTestId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "vn_monitoring")

	response, err := s.Client.GetPathAnalyzerTest(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PathAnalyzerTest
	return nil
}

func (s *VnMonitoringPathAnalyzerTestResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_vn_monitoring.UpdatePathAnalyzerTestRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if destinationEndpoint, ok := s.D.GetOkExists("destination_endpoint"); ok {
		if tmpList := destinationEndpoint.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "destination_endpoint", 0)
			tmp, err := s.mapToEndpoint(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DestinationEndpoint = tmp
		}
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.PathAnalyzerTestId = &tmp

	if protocol, ok := s.D.GetOkExists("protocol"); ok {
		tmp := protocol.(int)
		request.Protocol = &tmp
	}

	if protocolParameters, ok := s.D.GetOkExists("protocol_parameters"); ok {
		if tmpList := protocolParameters.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "protocol_parameters", 0)
			tmp, err := s.mapToProtocolParameters(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ProtocolParameters = tmp
		}
	}

	if queryOptions, ok := s.D.GetOkExists("query_options"); ok {
		if tmpList := queryOptions.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "query_options", 0)
			tmp, err := s.mapToQueryOptions(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.QueryOptions = &tmp
		}
	}

	if sourceEndpoint, ok := s.D.GetOkExists("source_endpoint"); ok {
		if tmpList := sourceEndpoint.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source_endpoint", 0)
			tmp, err := s.mapToEndpoint(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SourceEndpoint = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "vn_monitoring")

	response, err := s.Client.UpdatePathAnalyzerTest(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PathAnalyzerTest
	return nil
}

func (s *VnMonitoringPathAnalyzerTestResourceCrud) Delete() error {
	request := oci_vn_monitoring.DeletePathAnalyzerTestRequest{}

	tmp := s.D.Id()
	request.PathAnalyzerTestId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "vn_monitoring")

	_, err := s.Client.DeletePathAnalyzerTest(context.Background(), request)
	return err
}

func (s *VnMonitoringPathAnalyzerTestResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DestinationEndpoint != nil {
		destinationEndpointArray := []interface{}{}
		if destinationEndpointMap := EndpointToMap(&s.Res.DestinationEndpoint); destinationEndpointMap != nil {
			destinationEndpointArray = append(destinationEndpointArray, destinationEndpointMap)
		}
		s.D.Set("destination_endpoint", destinationEndpointArray)
	} else {
		s.D.Set("destination_endpoint", nil)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Protocol != nil {
		s.D.Set("protocol", *s.Res.Protocol)
	}

	if s.Res.ProtocolParameters != nil {
		protocolParametersArray := []interface{}{}
		if protocolParametersMap := ProtocolParametersToMap(&s.Res.ProtocolParameters); protocolParametersMap != nil {
			protocolParametersArray = append(protocolParametersArray, protocolParametersMap)
		}
		s.D.Set("protocol_parameters", protocolParametersArray)
	} else {
		s.D.Set("protocol_parameters", nil)
	}

	if s.Res.QueryOptions != nil {
		s.D.Set("query_options", []interface{}{QueryOptionsToMap(s.Res.QueryOptions)})
	} else {
		s.D.Set("query_options", nil)
	}

	if s.Res.SourceEndpoint != nil {
		sourceEndpointArray := []interface{}{}
		if sourceEndpointMap := EndpointToMap(&s.Res.SourceEndpoint); sourceEndpointMap != nil {
			sourceEndpointArray = append(sourceEndpointArray, sourceEndpointMap)
		}
		s.D.Set("source_endpoint", sourceEndpointArray)
	} else {
		s.D.Set("source_endpoint", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *VnMonitoringPathAnalyzerTestResourceCrud) mapToEndpoint(fieldKeyFormat string) (oci_vn_monitoring.Endpoint, error) {
	var baseObject oci_vn_monitoring.Endpoint
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("COMPUTE_INSTANCE"):
		details := oci_vn_monitoring.ComputeInstanceEndpoint{}
		if address, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "address")); ok {
			tmp := address.(string)
			details.Address = &tmp
		}
		if instanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "instance_id")); ok {
			tmp := instanceId.(string)
			details.InstanceId = &tmp
		}
		if vnicId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vnic_id")); ok {
			tmp := vnicId.(string)
			details.VnicId = &tmp
		}
		baseObject = details
	case strings.ToLower("IP_ADDRESS"):
		details := oci_vn_monitoring.IpAddressEndpoint{}
		if address, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "address")); ok {
			tmp := address.(string)
			details.Address = &tmp
		}
		baseObject = details
	case strings.ToLower("LOAD_BALANCER"):
		details := oci_vn_monitoring.LoadBalancerEndpoint{}
		if loadBalancerId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "load_balancer_id")); ok {
			tmp := loadBalancerId.(string)
			details.LoadBalancerId = &tmp
		}
		baseObject = details
	case strings.ToLower("LOAD_BALANCER_LISTENER"):
		details := oci_vn_monitoring.LoadBalancerListenerEndpoint{}
		if listenerId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "listener_id")); ok {
			tmp := listenerId.(string)
			details.ListenerId = &tmp
		}
		if loadBalancerId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "load_balancer_id")); ok {
			tmp := loadBalancerId.(string)
			details.LoadBalancerId = &tmp
		}
		baseObject = details
	case strings.ToLower("NETWORK_LOAD_BALANCER"):
		details := oci_vn_monitoring.NetworkLoadBalancerEndpoint{}
		if networkLoadBalancerId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_load_balancer_id")); ok {
			tmp := networkLoadBalancerId.(string)
			details.NetworkLoadBalancerId = &tmp
		}
		baseObject = details
	case strings.ToLower("NETWORK_LOAD_BALANCER_LISTENER"):
		details := oci_vn_monitoring.NetworkLoadBalancerListenerEndpoint{}
		if listenerId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "listener_id")); ok {
			tmp := listenerId.(string)
			details.ListenerId = &tmp
		}
		if networkLoadBalancerId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_load_balancer_id")); ok {
			tmp := networkLoadBalancerId.(string)
			details.NetworkLoadBalancerId = &tmp
		}
		baseObject = details
	case strings.ToLower("SUBNET"):
		details := oci_vn_monitoring.SubnetEndpoint{}
		if address, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "address")); ok {
			tmp := address.(string)
			details.Address = &tmp
		}
		if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		baseObject = details
	case strings.ToLower("VLAN"):
		details := oci_vn_monitoring.VlanEndpoint{}
		if address, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "address")); ok {
			tmp := address.(string)
			details.Address = &tmp
		}
		if vlanId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vlan_id")); ok {
			tmp := vlanId.(string)
			details.VlanId = &tmp
		}
		baseObject = details
	case strings.ToLower("VNIC"):
		details := oci_vn_monitoring.VnicEndpoint{}
		if address, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "address")); ok {
			tmp := address.(string)
			details.Address = &tmp
		}
		if vnicId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vnic_id")); ok {
			tmp := vnicId.(string)
			details.VnicId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func EndpointToMap(obj *oci_vn_monitoring.Endpoint) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_vn_monitoring.ComputeInstanceEndpoint:
		result["type"] = "COMPUTE_INSTANCE"

		if v.Address != nil {
			result["address"] = string(*v.Address)
		}

		if v.InstanceId != nil {
			result["instance_id"] = string(*v.InstanceId)
		}

		if v.VnicId != nil {
			result["vnic_id"] = string(*v.VnicId)
		}
	case oci_vn_monitoring.IpAddressEndpoint:
		result["type"] = "IP_ADDRESS"

		if v.Address != nil {
			result["address"] = string(*v.Address)
		}
	case oci_vn_monitoring.LoadBalancerEndpoint:
		result["type"] = "LOAD_BALANCER"

		if v.LoadBalancerId != nil {
			result["load_balancer_id"] = string(*v.LoadBalancerId)
		}
	case oci_vn_monitoring.LoadBalancerListenerEndpoint:
		result["type"] = "LOAD_BALANCER_LISTENER"

		if v.ListenerId != nil {
			result["listener_id"] = string(*v.ListenerId)
		}

		if v.LoadBalancerId != nil {
			result["load_balancer_id"] = string(*v.LoadBalancerId)
		}
	case oci_vn_monitoring.NetworkLoadBalancerEndpoint:
		result["type"] = "NETWORK_LOAD_BALANCER"

		if v.NetworkLoadBalancerId != nil {
			result["network_load_balancer_id"] = string(*v.NetworkLoadBalancerId)
		}
	case oci_vn_monitoring.NetworkLoadBalancerListenerEndpoint:
		result["type"] = "NETWORK_LOAD_BALANCER_LISTENER"

		if v.ListenerId != nil {
			result["listener_id"] = string(*v.ListenerId)
		}

		if v.NetworkLoadBalancerId != nil {
			result["network_load_balancer_id"] = string(*v.NetworkLoadBalancerId)
		}
	case oci_vn_monitoring.SubnetEndpoint:
		result["type"] = "SUBNET"

		if v.Address != nil {
			result["address"] = string(*v.Address)
		}

		if v.SubnetId != nil {
			result["subnet_id"] = string(*v.SubnetId)
		}
	case oci_vn_monitoring.VlanEndpoint:
		result["type"] = "VLAN"

		if v.Address != nil {
			result["address"] = string(*v.Address)
		}

		if v.VlanId != nil {
			result["vlan_id"] = string(*v.VlanId)
		}
	case oci_vn_monitoring.VnicEndpoint:
		result["type"] = "VNIC"

		if v.Address != nil {
			result["address"] = string(*v.Address)
		}

		if v.VnicId != nil {
			result["vnic_id"] = string(*v.VnicId)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func PathAnalyzerTestSummaryToMap(obj oci_vn_monitoring.PathAnalyzerTestSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DestinationEndpoint != nil {
		destinationEndpointArray := []interface{}{}
		if destinationEndpointMap := EndpointToMap(&obj.DestinationEndpoint); destinationEndpointMap != nil {
			destinationEndpointArray = append(destinationEndpointArray, destinationEndpointMap)
		}
		result["destination_endpoint"] = destinationEndpointArray
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Protocol != nil {
		result["protocol"] = int(*obj.Protocol)
	}

	if obj.ProtocolParameters != nil {
		protocolParametersArray := []interface{}{}
		if protocolParametersMap := ProtocolParametersToMap(&obj.ProtocolParameters); protocolParametersMap != nil {
			protocolParametersArray = append(protocolParametersArray, protocolParametersMap)
		}
		result["protocol_parameters"] = protocolParametersArray
	}

	if obj.QueryOptions != nil {
		result["query_options"] = []interface{}{QueryOptionsToMap(obj.QueryOptions)}
	}

	if obj.SourceEndpoint != nil {
		sourceEndpointArray := []interface{}{}
		if sourceEndpointMap := EndpointToMap(&obj.SourceEndpoint); sourceEndpointMap != nil {
			sourceEndpointArray = append(sourceEndpointArray, sourceEndpointMap)
		}
		result["source_endpoint"] = sourceEndpointArray
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *VnMonitoringPathAnalyzerTestResourceCrud) mapToProtocolParameters(fieldKeyFormat string) (oci_vn_monitoring.ProtocolParameters, error) {
	var baseObject oci_vn_monitoring.ProtocolParameters
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("ICMP"):
		details := oci_vn_monitoring.IcmpProtocolParameters{}
		if icmpCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "icmp_code")); ok {
			tmp := icmpCode.(int)
			details.IcmpCode = &tmp
		}
		if icmpType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "icmp_type")); ok {
			tmp := icmpType.(int)
			details.IcmpType = &tmp
		}
		baseObject = details
	case strings.ToLower("TCP"):
		details := oci_vn_monitoring.TcpProtocolParameters{}
		if destinationPort, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_port")); ok {
			tmp := destinationPort.(int)
			details.DestinationPort = &tmp
		}
		if sourcePort, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_port")); ok {
			tmp := sourcePort.(int)
			details.SourcePort = &tmp
		}
		baseObject = details
	case strings.ToLower("UDP"):
		details := oci_vn_monitoring.UdpProtocolParameters{}
		if destinationPort, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_port")); ok {
			tmp := destinationPort.(int)
			details.DestinationPort = &tmp
		}
		if sourcePort, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_port")); ok {
			tmp := sourcePort.(int)
			details.SourcePort = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func ProtocolParametersToMap(obj *oci_vn_monitoring.ProtocolParameters) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_vn_monitoring.IcmpProtocolParameters:
		result["type"] = "ICMP"

		if v.IcmpCode != nil {
			result["icmp_code"] = int(*v.IcmpCode)
		}

		if v.IcmpType != nil {
			result["icmp_type"] = int(*v.IcmpType)
		}
	case oci_vn_monitoring.TcpProtocolParameters:
		result["type"] = "TCP"

		if v.DestinationPort != nil {
			result["destination_port"] = int(*v.DestinationPort)
		}

		if v.SourcePort != nil {
			result["source_port"] = int(*v.SourcePort)
		}
	case oci_vn_monitoring.UdpProtocolParameters:
		result["type"] = "UDP"

		if v.DestinationPort != nil {
			result["destination_port"] = int(*v.DestinationPort)
		}

		if v.SourcePort != nil {
			result["source_port"] = int(*v.SourcePort)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *VnMonitoringPathAnalyzerTestResourceCrud) mapToQueryOptions(fieldKeyFormat string) (oci_vn_monitoring.QueryOptions, error) {
	result := oci_vn_monitoring.QueryOptions{}

	if isBiDirectionalAnalysis, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_bi_directional_analysis")); ok {
		tmp := isBiDirectionalAnalysis.(bool)
		result.IsBiDirectionalAnalysis = &tmp
	}

	return result, nil
}

func QueryOptionsToMap(obj *oci_vn_monitoring.QueryOptions) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsBiDirectionalAnalysis != nil {
		result["is_bi_directional_analysis"] = bool(*obj.IsBiDirectionalAnalysis)
	}

	return result
}

func (s *VnMonitoringPathAnalyzerTestResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_vn_monitoring.ChangePathAnalyzerTestCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.PathAnalyzerTestId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "vn_monitoring")

	_, err := s.Client.ChangePathAnalyzerTestCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
