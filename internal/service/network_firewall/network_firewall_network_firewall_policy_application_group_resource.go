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

func NetworkFirewallNetworkFirewallPolicyApplicationGroupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createNetworkFirewallNetworkFirewallPolicyApplicationGroup,
		Read:     readNetworkFirewallNetworkFirewallPolicyApplicationGroup,
		Update:   updateNetworkFirewallNetworkFirewallPolicyApplicationGroup,
		Delete:   deleteNetworkFirewallNetworkFirewallPolicyApplicationGroup,
		Schema: map[string]*schema.Schema{
			// Required
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"apps": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				MinItems: 0,
			},
			"network_firewall_policy_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Computed
			"parent_resource_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_apps": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func createNetworkFirewallNetworkFirewallPolicyApplicationGroup(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyApplicationGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.CreateResource(d, sync)
}

func readNetworkFirewallNetworkFirewallPolicyApplicationGroup(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyApplicationGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.ReadResource(sync)
}

func updateNetworkFirewallNetworkFirewallPolicyApplicationGroup(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyApplicationGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteNetworkFirewallNetworkFirewallPolicyApplicationGroup(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyApplicationGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type NetworkFirewallNetworkFirewallPolicyApplicationGroupResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_network_firewall.NetworkFirewallClient
	Res                    *oci_network_firewall.ApplicationGroup
	DisableNotFoundRetries bool
}

func (s *NetworkFirewallNetworkFirewallPolicyApplicationGroupResourceCrud) ID() string {
	return GetNetworkFirewallPolicySubResourceCompositeId(s.D.Get("name").(string), s.D.Get("network_firewall_policy_id").(string), "applicationGroups")
}

func (s *NetworkFirewallNetworkFirewallPolicyApplicationGroupResourceCrud) Create() error {
	request := oci_network_firewall.CreateApplicationGroupRequest{}

	if apps, ok := s.D.GetOkExists("apps"); ok {
		interfaces := apps.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.Apps = tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.CreateApplicationGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ApplicationGroup
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyApplicationGroupResourceCrud) Get() error {
	request := oci_network_firewall.GetApplicationGroupRequest{}

	if applicationGroupName, ok := s.D.GetOkExists("name"); ok {
		tmp := applicationGroupName.(string)
		request.ApplicationGroupName = &tmp
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	applicationGroupName, networkFirewallPolicyId, err := parseNetworkFirewallPolicySubResourceCompositeId(s.D.Id(), "applicationGroups")
	if err == nil {
		request.ApplicationGroupName = &applicationGroupName
		request.NetworkFirewallPolicyId = &networkFirewallPolicyId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.GetApplicationGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ApplicationGroup
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyApplicationGroupResourceCrud) Update() error {
	request := oci_network_firewall.UpdateApplicationGroupRequest{}

	if applicationGroupName, ok := s.D.GetOkExists("name"); ok {
		tmp := applicationGroupName.(string)
		request.ApplicationGroupName = &tmp
	}

	if apps, ok := s.D.GetOkExists("apps"); ok {
		interfaces := apps.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.Apps = tmp

	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.UpdateApplicationGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ApplicationGroup
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyApplicationGroupResourceCrud) Delete() error {
	request := oci_network_firewall.DeleteApplicationGroupRequest{}

	if applicationGroupName, ok := s.D.GetOkExists("name"); ok {
		tmp := applicationGroupName.(string)
		request.ApplicationGroupName = &tmp
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	_, err := s.Client.DeleteApplicationGroup(context.Background(), request)
	return err
}

func (s *NetworkFirewallNetworkFirewallPolicyApplicationGroupResourceCrud) SetData() error {

	applicationGroupName, networkFirewallPolicyId, err := parseNetworkFirewallPolicySubResourceCompositeId(s.D.Id(), "applicationGroups")
	if err == nil {
		s.D.Set("name", &applicationGroupName)
		s.D.Set("network_firewall_policy_id", &networkFirewallPolicyId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	s.D.Set("apps", s.Res.Apps)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ParentResourceId != nil {
		s.D.Set("parent_resource_id", *s.Res.ParentResourceId)
	}

	if s.Res.TotalApps != nil {
		s.D.Set("total_apps", *s.Res.TotalApps)
	}

	return nil
}

func ApplicationGroupSummaryToMap(obj oci_network_firewall.ApplicationGroupSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ParentResourceId != nil {
		result["parent_resource_id"] = string(*obj.ParentResourceId)
	}

	if obj.TotalApps != nil {
		result["total_apps"] = int(*obj.TotalApps)
	}

	return result
}
