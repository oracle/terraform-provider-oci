// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package network_firewall

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_network_firewall "github.com/oracle/oci-go-sdk/v65/networkfirewall"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func NetworkFirewallNetworkFirewallPolicyServiceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createNetworkFirewallNetworkFirewallPolicyService,
		Read:     readNetworkFirewallNetworkFirewallPolicyService,
		Update:   updateNetworkFirewallNetworkFirewallPolicyService,
		Delete:   deleteNetworkFirewallNetworkFirewallPolicyService,
		Schema: map[string]*schema.Schema{
			// Required
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
			"port_ranges": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"minimum_port": {
							Type:     schema.TypeInt,
							Required: true,
						},

						// Optional
						"maximum_port": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"type": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"TCP_SERVICE",
					"UDP_SERVICE",
				}, false),
			},

			// Computed
			"parent_resource_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createNetworkFirewallNetworkFirewallPolicyService(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyServiceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.CreateResource(d, sync)
}

func readNetworkFirewallNetworkFirewallPolicyService(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyServiceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.ReadResource(sync)
}

func updateNetworkFirewallNetworkFirewallPolicyService(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyServiceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteNetworkFirewallNetworkFirewallPolicyService(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyServiceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type NetworkFirewallNetworkFirewallPolicyServiceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_network_firewall.NetworkFirewallClient
	Res                    *oci_network_firewall.Service
	DisableNotFoundRetries bool
}

func (s *NetworkFirewallNetworkFirewallPolicyServiceResourceCrud) ID() string {
	return GetNetworkFirewallPolicySubResourceCompositeId(s.D.Get("name").(string), s.D.Get("network_firewall_policy_id").(string), "services")
}

func (s *NetworkFirewallNetworkFirewallPolicyServiceResourceCrud) Create() error {
	request := oci_network_firewall.CreateServiceRequest{}
	err := s.populateTopLevelPolymorphicCreateServiceRequest(&request)
	if err != nil {
		return err
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.CreateService(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Service
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyServiceResourceCrud) Get() error {
	request := oci_network_firewall.GetServiceRequest{}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	if serviceName, ok := s.D.GetOkExists("name"); ok {
		tmp := serviceName.(string)
		request.ServiceName = &tmp
	}

	serviceName, networkFirewallPolicyId, err := parseNetworkFirewallPolicySubResourceCompositeId(s.D.Id(), "services")

	if err == nil {
		request.NetworkFirewallPolicyId = &networkFirewallPolicyId
		request.ServiceName = &serviceName
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.GetService(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Service
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyServiceResourceCrud) Update() error {
	request := oci_network_firewall.UpdateServiceRequest{}
	err := s.populateTopLevelPolymorphicUpdateServiceRequest(&request)
	if err != nil {
		return err
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	if serviceName, ok := s.D.GetOkExists("name"); ok {
		tmp := serviceName.(string)
		request.ServiceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.UpdateService(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Service
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyServiceResourceCrud) Delete() error {
	request := oci_network_firewall.DeleteServiceRequest{}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	if serviceName, ok := s.D.GetOkExists("name"); ok {
		tmp := serviceName.(string)
		request.ServiceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	_, err := s.Client.DeleteService(context.Background(), request)
	return err
}

func (s *NetworkFirewallNetworkFirewallPolicyServiceResourceCrud) SetData() error {

	serviceName, networkFirewallPolicyId, err := parseNetworkFirewallPolicySubResourceCompositeId(s.D.Id(), "services")
	if err == nil {
		s.D.Set("network_firewall_policy_id", &networkFirewallPolicyId)
		s.D.Set("name", &serviceName)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	switch v := (*s.Res).(type) {
	case oci_network_firewall.TcpService:
		s.D.Set("type", "TCP_SERVICE")

		portRanges := []interface{}{}
		for _, item := range v.PortRanges {
			portRanges = append(portRanges, PortRangeToMap(item))
		}
		s.D.Set("port_ranges", portRanges)

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		if v.ParentResourceId != nil {
			s.D.Set("parent_resource_id", *v.ParentResourceId)
		}
	case oci_network_firewall.UdpService:
		s.D.Set("type", "UDP_SERVICE")

		portRanges := []interface{}{}
		for _, item := range v.PortRanges {
			portRanges = append(portRanges, PortRangeToMap(item))
		}
		s.D.Set("port_ranges", portRanges)

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		if v.ParentResourceId != nil {
			s.D.Set("parent_resource_id", *v.ParentResourceId)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyServiceResourceCrud) mapToPortRange(fieldKeyFormat string) (oci_network_firewall.PortRange, error) {
	result := oci_network_firewall.PortRange{}

	if maximumPort, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "maximum_port")); ok {
		tmp := maximumPort.(int)
		result.MaximumPort = &tmp
	}

	if minimumPort, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "minimum_port")); ok {
		tmp := minimumPort.(int)
		result.MinimumPort = &tmp
	}

	return result, nil
}

func PortRangeToMap(obj oci_network_firewall.PortRange) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MaximumPort != nil {
		result["maximum_port"] = int(*obj.MaximumPort)
	}

	if obj.MinimumPort != nil {
		result["minimum_port"] = int(*obj.MinimumPort)
	}

	return result
}

func ServiceSummaryToMap(obj oci_network_firewall.ServiceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ParentResourceId != nil {
		result["parent_resource_id"] = string(*obj.ParentResourceId)
	}

	result["type"] = string(obj.Type)

	return result
}

func (s *NetworkFirewallNetworkFirewallPolicyServiceResourceCrud) populateTopLevelPolymorphicCreateServiceRequest(request *oci_network_firewall.CreateServiceRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("TCP_SERVICE"):
		details := oci_network_firewall.CreateTcpServiceDetails{}
		if portRanges, ok := s.D.GetOkExists("port_ranges"); ok {
			interfaces := portRanges.([]interface{})
			tmp := make([]oci_network_firewall.PortRange, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "port_ranges", stateDataIndex)
				converted, err := s.mapToPortRange(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("port_ranges") {
				details.PortRanges = tmp
			}
		}
		if name, ok := s.D.GetOkExists("name"); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		request.CreateServiceDetails = details
	case strings.ToLower("UDP_SERVICE"):
		details := oci_network_firewall.CreateUdpServiceDetails{}
		if portRanges, ok := s.D.GetOkExists("port_ranges"); ok {
			interfaces := portRanges.([]interface{})
			tmp := make([]oci_network_firewall.PortRange, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "port_ranges", stateDataIndex)
				converted, err := s.mapToPortRange(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("port_ranges") {
				details.PortRanges = tmp
			}
		}
		if name, ok := s.D.GetOkExists("name"); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		request.CreateServiceDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyServiceResourceCrud) populateTopLevelPolymorphicUpdateServiceRequest(request *oci_network_firewall.UpdateServiceRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("TCP_SERVICE"):
		details := oci_network_firewall.UpdateTcpServiceDetails{}
		if portRanges, ok := s.D.GetOkExists("port_ranges"); ok {
			interfaces := portRanges.([]interface{})
			tmp := make([]oci_network_firewall.PortRange, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "port_ranges", stateDataIndex)
				converted, err := s.mapToPortRange(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("port_ranges") {
				details.PortRanges = tmp
			}
		}
		request.UpdateServiceDetails = details
	case strings.ToLower("UDP_SERVICE"):
		details := oci_network_firewall.UpdateUdpServiceDetails{}
		if portRanges, ok := s.D.GetOkExists("port_ranges"); ok {
			interfaces := portRanges.([]interface{})
			tmp := make([]oci_network_firewall.PortRange, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "port_ranges", stateDataIndex)
				converted, err := s.mapToPortRange(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("port_ranges") {
				details.PortRanges = tmp
			}
		}
		request.UpdateServiceDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}
