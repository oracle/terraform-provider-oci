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

func NetworkFirewallNetworkFirewallPolicyAddressListResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createNetworkFirewallNetworkFirewallPolicyAddressList,
		Read:     readNetworkFirewallNetworkFirewallPolicyAddressList,
		Update:   updateNetworkFirewallNetworkFirewallPolicyAddressList,
		Delete:   deleteNetworkFirewallNetworkFirewallPolicyAddressList,
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
			"type": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"FQDN",
					"IP",
				}, true),
			},
			"addresses": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Computed
			"parent_resource_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_addresses": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func createNetworkFirewallNetworkFirewallPolicyAddressList(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyAddressListResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.CreateResource(d, sync)
}

func readNetworkFirewallNetworkFirewallPolicyAddressList(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyAddressListResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.ReadResource(sync)
}

func updateNetworkFirewallNetworkFirewallPolicyAddressList(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyAddressListResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteNetworkFirewallNetworkFirewallPolicyAddressList(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyAddressListResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type NetworkFirewallNetworkFirewallPolicyAddressListResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_network_firewall.NetworkFirewallClient
	Res                    *oci_network_firewall.AddressList
	DisableNotFoundRetries bool
}

func (s *NetworkFirewallNetworkFirewallPolicyAddressListResourceCrud) ID() string {
	return GetNetworkFirewallPolicySubResourceCompositeId(s.D.Get("name").(string), s.D.Get("network_firewall_policy_id").(string), "addressLists")
}

func (s *NetworkFirewallNetworkFirewallPolicyAddressListResourceCrud) Create() error {
	request := oci_network_firewall.CreateAddressListRequest{}

	if addresses, ok := s.D.GetOkExists("addresses"); ok {
		interfaces := addresses.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("addresses") {
			request.Addresses = tmp
		}
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_network_firewall.AddressListTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.CreateAddressList(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AddressList
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyAddressListResourceCrud) Get() error {
	request := oci_network_firewall.GetAddressListRequest{}

	if addressListName, ok := s.D.GetOkExists("name"); ok {
		tmp := addressListName.(string)
		request.AddressListName = &tmp
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	addressListName, networkFirewallPolicyId, err := parseNetworkFirewallPolicySubResourceCompositeId(s.D.Id(), "addressLists")
	if err == nil {
		request.AddressListName = &addressListName
		request.NetworkFirewallPolicyId = &networkFirewallPolicyId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.GetAddressList(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AddressList
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyAddressListResourceCrud) Update() error {
	request := oci_network_firewall.UpdateAddressListRequest{}
	err := s.populateTopLevelPolymorphicUpdateAddressListRequest(&request)
	if err != nil {
		return err
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	if addressListName, ok := s.D.GetOkExists("name"); ok {
		tmp := addressListName.(string)
		request.AddressListName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.UpdateAddressList(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AddressList
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyAddressListResourceCrud) Delete() error {
	request := oci_network_firewall.DeleteAddressListRequest{}

	if addressListName, ok := s.D.GetOkExists("name"); ok {
		tmp := addressListName.(string)
		request.AddressListName = &tmp
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	_, err := s.Client.DeleteAddressList(context.Background(), request)
	return err
}

func (s *NetworkFirewallNetworkFirewallPolicyAddressListResourceCrud) SetData() error {

	addressListName, networkFirewallPolicyId, err := parseNetworkFirewallPolicySubResourceCompositeId(s.D.Id(), "addressLists")
	if err == nil {
		s.D.Set("name", &addressListName)
		s.D.Set("network_firewall_policy_id", &networkFirewallPolicyId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	s.D.Set("addresses", s.Res.Addresses)

	if s.Res.ParentResourceId != nil {
		s.D.Set("parent_resource_id", *s.Res.ParentResourceId)
	}

	if s.Res.TotalAddresses != nil {
		s.D.Set("total_addresses", *s.Res.TotalAddresses)
	}

	s.D.Set("type", s.Res.Type)

	return nil
}

func AddressListSummaryToMap(obj oci_network_firewall.AddressListSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ParentResourceId != nil {
		result["parent_resource_id"] = string(*obj.ParentResourceId)
	}

	if obj.TotalAddresses != nil {
		result["total_addresses"] = int(*obj.TotalAddresses)
	}

	result["type"] = string(obj.Type)

	return result
}

func (s *NetworkFirewallNetworkFirewallPolicyAddressListResourceCrud) populateTopLevelPolymorphicUpdateAddressListRequest(request *oci_network_firewall.UpdateAddressListRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("FQDN"):
		details := oci_network_firewall.UpdateFqdnAddressListDetails{}
		if addresses, ok := s.D.GetOkExists("addresses"); ok {
			interfaces := addresses.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("addresses") {
				details.Addresses = tmp
			}
		}
		if addresses, ok := s.D.GetOkExists("addresses"); ok {
			interfaces := addresses.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("addresses") {
				details.Addresses = tmp
			}
		}
		request.UpdateAddressListDetails = details
	case strings.ToLower("IP"):
		details := oci_network_firewall.UpdateIpAddressListDetails{}
		if addresses, ok := s.D.GetOkExists("addresses"); ok {
			interfaces := addresses.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("addresses") {
				details.Addresses = tmp
			}
		}
		if addresses, ok := s.D.GetOkExists("addresses"); ok {
			interfaces := addresses.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("addresses") {
				details.Addresses = tmp
			}
		}
		request.UpdateAddressListDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}
