// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package network_firewall

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_network_firewall "github.com/oracle/oci-go-sdk/v65/networkfirewall"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func NetworkFirewallNetworkFirewallPolicyServiceListResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createNetworkFirewallNetworkFirewallPolicyServiceList,
		Read:     readNetworkFirewallNetworkFirewallPolicyServiceList,
		Update:   updateNetworkFirewallNetworkFirewallPolicyServiceList,
		Delete:   deleteNetworkFirewallNetworkFirewallPolicyServiceList,
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
			"services": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 0,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional

			// Computed
			"parent_resource_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_services": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func createNetworkFirewallNetworkFirewallPolicyServiceList(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyServiceListResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.CreateResource(d, sync)
}

func readNetworkFirewallNetworkFirewallPolicyServiceList(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyServiceListResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.ReadResource(sync)
}

func updateNetworkFirewallNetworkFirewallPolicyServiceList(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyServiceListResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteNetworkFirewallNetworkFirewallPolicyServiceList(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyServiceListResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type NetworkFirewallNetworkFirewallPolicyServiceListResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_network_firewall.NetworkFirewallClient
	Res                    *oci_network_firewall.ServiceList
	DisableNotFoundRetries bool
}

func (s *NetworkFirewallNetworkFirewallPolicyServiceListResourceCrud) ID() string {
	return GetNetworkFirewallPolicySubResourceCompositeId(s.D.Get("name").(string), s.D.Get("network_firewall_policy_id").(string), "serviceLists")
}

func (s *NetworkFirewallNetworkFirewallPolicyServiceListResourceCrud) Create() error {
	request := oci_network_firewall.CreateServiceListRequest{}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	if services, ok := s.D.GetOkExists("services"); ok {
		interfaces := services.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.Services = tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.CreateServiceList(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ServiceList
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyServiceListResourceCrud) Get() error {
	request := oci_network_firewall.GetServiceListRequest{}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	if serviceListName, ok := s.D.GetOkExists("name"); ok {
		tmp := serviceListName.(string)
		request.ServiceListName = &tmp
	}

	serviceListName, networkFirewallPolicyId, err := parseNetworkFirewallPolicySubResourceCompositeId(s.D.Id(), "serviceLists")
	if err == nil {
		request.NetworkFirewallPolicyId = &networkFirewallPolicyId
		request.ServiceListName = &serviceListName
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.GetServiceList(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ServiceList
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyServiceListResourceCrud) Update() error {
	request := oci_network_firewall.UpdateServiceListRequest{}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	if serviceListName, ok := s.D.GetOkExists("name"); ok {
		tmp := serviceListName.(string)
		request.ServiceListName = &tmp
	}

	if services, ok := s.D.GetOkExists("services"); ok {
		interfaces := services.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.Services = tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.UpdateServiceList(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ServiceList
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyServiceListResourceCrud) Delete() error {
	request := oci_network_firewall.DeleteServiceListRequest{}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	if serviceListName, ok := s.D.GetOkExists("name"); ok {
		tmp := serviceListName.(string)
		request.ServiceListName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	_, err := s.Client.DeleteServiceList(context.Background(), request)
	return err
}

func (s *NetworkFirewallNetworkFirewallPolicyServiceListResourceCrud) SetData() error {

	serviceListName, networkFirewallPolicyId, err := parseNetworkFirewallPolicySubResourceCompositeId(s.D.Id(), "serviceLists")
	if err == nil {
		s.D.Set("network_firewall_policy_id", &networkFirewallPolicyId)
		s.D.Set("name", &serviceListName)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ParentResourceId != nil {
		s.D.Set("parent_resource_id", *s.Res.ParentResourceId)
	}

	s.D.Set("services", s.Res.Services)

	if s.Res.TotalServices != nil {
		s.D.Set("total_services", *s.Res.TotalServices)
	}

	return nil
}

func ServiceListSummaryToMap(obj oci_network_firewall.ServiceListSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ParentResourceId != nil {
		result["parent_resource_id"] = string(*obj.ParentResourceId)
	}

	if obj.TotalServices != nil {
		result["total_services"] = int(*obj.TotalServices)
	}

	return result
}
