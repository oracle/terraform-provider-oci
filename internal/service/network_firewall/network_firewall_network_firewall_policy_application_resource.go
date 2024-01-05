// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package network_firewall

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_network_firewall "github.com/oracle/oci-go-sdk/v65/networkfirewall"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func NetworkFirewallNetworkFirewallPolicyApplicationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createNetworkFirewallNetworkFirewallPolicyApplication,
		Read:     readNetworkFirewallNetworkFirewallPolicyApplication,
		Update:   updateNetworkFirewallNetworkFirewallPolicyApplication,
		Delete:   deleteNetworkFirewallNetworkFirewallPolicyApplication,
		Schema: map[string]*schema.Schema{
			// Required
			"icmp_type": {
				Type:     schema.TypeInt,
				Required: true,
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
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"ICMP",
					"ICMP_V6",
				}, false),
			},

			// Optional
			"icmp_code": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// Computed
			"parent_resource_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createNetworkFirewallNetworkFirewallPolicyApplication(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.CreateResource(d, sync)
}

func readNetworkFirewallNetworkFirewallPolicyApplication(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.ReadResource(sync)
}

func updateNetworkFirewallNetworkFirewallPolicyApplication(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteNetworkFirewallNetworkFirewallPolicyApplication(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyApplicationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type NetworkFirewallNetworkFirewallPolicyApplicationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_network_firewall.NetworkFirewallClient
	Res                    *oci_network_firewall.Application
	DisableNotFoundRetries bool
}

func (s *NetworkFirewallNetworkFirewallPolicyApplicationResourceCrud) ID() string {
	return GetNetworkFirewallPolicySubResourceCompositeId(s.D.Get("name").(string), s.D.Get("network_firewall_policy_id").(string), "applications")
}

func (s *NetworkFirewallNetworkFirewallPolicyApplicationResourceCrud) Create() error {
	request := oci_network_firewall.CreateApplicationRequest{}
	err := s.populateTopLevelPolymorphicCreateApplicationRequest(&request)
	if err != nil {
		return err
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.CreateApplication(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Application
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyApplicationResourceCrud) Get() error {
	request := oci_network_firewall.GetApplicationRequest{}

	if applicationName, ok := s.D.GetOkExists("name"); ok {
		tmp := applicationName.(string)
		request.ApplicationName = &tmp
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	applicationName, networkFirewallPolicyId, err := parseNetworkFirewallPolicySubResourceCompositeId(s.D.Id(), "applications")
	if err == nil {
		request.ApplicationName = &applicationName
		request.NetworkFirewallPolicyId = &networkFirewallPolicyId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.GetApplication(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Application
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyApplicationResourceCrud) Update() error {
	request := oci_network_firewall.UpdateApplicationRequest{}
	err := s.populateTopLevelPolymorphicUpdateApplicationRequest(&request)
	if err != nil {
		return err
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	if applicationName, ok := s.D.GetOkExists("name"); ok {
		tmp := applicationName.(string)
		request.ApplicationName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.UpdateApplication(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Application
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyApplicationResourceCrud) Delete() error {
	request := oci_network_firewall.DeleteApplicationRequest{}

	if applicationName, ok := s.D.GetOkExists("name"); ok {
		tmp := applicationName.(string)
		request.ApplicationName = &tmp
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	_, err := s.Client.DeleteApplication(context.Background(), request)
	return err
}

func (s *NetworkFirewallNetworkFirewallPolicyApplicationResourceCrud) SetData() error {

	applicationName, networkFirewallPolicyId, err := parseNetworkFirewallPolicySubResourceCompositeId(s.D.Id(), "applications")
	if err == nil {
		s.D.Set("name", &applicationName)
		s.D.Set("network_firewall_policy_id", &networkFirewallPolicyId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	switch v := (*s.Res).(type) {
	case oci_network_firewall.IcmpApplication:
		s.D.Set("type", "ICMP")

		if v.IcmpCode != nil {
			s.D.Set("icmp_code", *v.IcmpCode)
		}

		if v.IcmpType != nil {
			s.D.Set("icmp_type", *v.IcmpType)
		}

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		if v.ParentResourceId != nil {
			s.D.Set("parent_resource_id", *v.ParentResourceId)
		}
	case oci_network_firewall.Icmp6Application:
		s.D.Set("type", "ICMP_V6")

		if v.IcmpCode != nil {
			s.D.Set("icmp_code", *v.IcmpCode)
		}

		if v.IcmpType != nil {
			s.D.Set("icmp_type", *v.IcmpType)
		}

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

func ApplicationSummaryToMap(obj oci_network_firewall.ApplicationSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_network_firewall.IcmpApplicationSummary:
		result["type"] = "ICMP"

		if v.IcmpCode != nil {
			result["icmp_code"] = int(*v.IcmpCode)
		}

		if v.IcmpType != nil {
			result["icmp_type"] = int(*v.IcmpType)
		}
	case oci_network_firewall.Icmp6ApplicationSummary:
		result["type"] = "ICMP_V6"

		if v.IcmpCode != nil {
			result["icmp_code"] = int(*v.IcmpCode)
		}

		if v.IcmpType != nil {
			result["icmp_type"] = int(*v.IcmpType)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *NetworkFirewallNetworkFirewallPolicyApplicationResourceCrud) populateTopLevelPolymorphicCreateApplicationRequest(request *oci_network_firewall.CreateApplicationRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("ICMP"):
		details := oci_network_firewall.CreateIcmpApplicationDetails{}
		if icmpCode, ok := s.D.GetOkExists("icmp_code"); ok {
			tmp := icmpCode.(int)
			details.IcmpCode = &tmp
		}
		if icmpType, ok := s.D.GetOkExists("icmp_type"); ok {
			tmp := icmpType.(int)
			details.IcmpType = &tmp
		}
		if name, ok := s.D.GetOkExists("name"); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		request.CreateApplicationDetails = details
	case strings.ToLower("ICMP_V6"):
		details := oci_network_firewall.CreateIcmp6ApplicationDetails{}
		if icmpCode, ok := s.D.GetOkExists("icmp_code"); ok {
			tmp := icmpCode.(int)
			details.IcmpCode = &tmp
		}
		if icmpType, ok := s.D.GetOkExists("icmp_type"); ok {
			tmp := icmpType.(int)
			details.IcmpType = &tmp
		}
		if name, ok := s.D.GetOkExists("name"); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		request.CreateApplicationDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyApplicationResourceCrud) populateTopLevelPolymorphicUpdateApplicationRequest(request *oci_network_firewall.UpdateApplicationRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("ICMP"):
		details := oci_network_firewall.UpdateIcmpApplicationDetails{}
		if icmpCode, ok := s.D.GetOkExists("icmp_code"); ok {
			tmp := icmpCode.(int)
			details.IcmpCode = &tmp
		}
		if icmpType, ok := s.D.GetOkExists("icmp_type"); ok {
			tmp := icmpType.(int)
			details.IcmpType = &tmp
		}
		request.UpdateApplicationDetails = details
	case strings.ToLower("ICMP_V6"):
		details := oci_network_firewall.UpdateIcmp6ApplicationDetails{}
		if icmpCode, ok := s.D.GetOkExists("icmp_code"); ok {
			tmp := icmpCode.(int)
			details.IcmpCode = &tmp
		}
		if icmpType, ok := s.D.GetOkExists("icmp_type"); ok {
			tmp := icmpType.(int)
			details.IcmpType = &tmp
		}
		request.UpdateApplicationDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}
