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

func NetworkFirewallNetworkFirewallPolicyUrlListResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createNetworkFirewallNetworkFirewallPolicyUrlList,
		Read:     readNetworkFirewallNetworkFirewallPolicyUrlList,
		Update:   updateNetworkFirewallNetworkFirewallPolicyUrlList,
		Delete:   deleteNetworkFirewallNetworkFirewallPolicyUrlList,
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
			"urls": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"pattern": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"SIMPLE",
							}, false),
						},
					},
				},
			},

			// Optional

			// Computed
			"parent_resource_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_urls": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func createNetworkFirewallNetworkFirewallPolicyUrlList(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyUrlListResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.CreateResource(d, sync)
}

func readNetworkFirewallNetworkFirewallPolicyUrlList(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyUrlListResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.ReadResource(sync)
}

func updateNetworkFirewallNetworkFirewallPolicyUrlList(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyUrlListResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteNetworkFirewallNetworkFirewallPolicyUrlList(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkFirewallNetworkFirewallPolicyUrlListResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkFirewallClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type NetworkFirewallNetworkFirewallPolicyUrlListResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_network_firewall.NetworkFirewallClient
	Res                    *oci_network_firewall.UrlList
	DisableNotFoundRetries bool
}

func (s *NetworkFirewallNetworkFirewallPolicyUrlListResourceCrud) ID() string {
	return GetNetworkFirewallPolicySubResourceCompositeId(s.D.Get("name").(string), s.D.Get("network_firewall_policy_id").(string), "urlLists")
}

func (s *NetworkFirewallNetworkFirewallPolicyUrlListResourceCrud) Create() error {
	request := oci_network_firewall.CreateUrlListRequest{}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	if urls, ok := s.D.GetOkExists("urls"); ok {
		interfaces := urls.([]interface{})
		tmp := make([]oci_network_firewall.UrlPattern, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urls", stateDataIndex)
			converted, err := s.mapToUrlPattern(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("urls") {
			request.Urls = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.CreateUrlList(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.UrlList
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyUrlListResourceCrud) Get() error {
	request := oci_network_firewall.GetUrlListRequest{}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	if urlListName, ok := s.D.GetOkExists("name"); ok {
		tmp := urlListName.(string)
		request.UrlListName = &tmp
	}

	urlListName, networkFirewallPolicyId, err := parseNetworkFirewallPolicySubResourceCompositeId(s.D.Id(), "urlLists")
	if err == nil {
		request.NetworkFirewallPolicyId = &networkFirewallPolicyId
		request.UrlListName = &urlListName
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.GetUrlList(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.UrlList
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyUrlListResourceCrud) Update() error {
	request := oci_network_firewall.UpdateUrlListRequest{}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	if urlListName, ok := s.D.GetOkExists("name"); ok {
		tmp := urlListName.(string)
		request.UrlListName = &tmp
	}

	if urls, ok := s.D.GetOkExists("urls"); ok {
		interfaces := urls.([]interface{})
		tmp := make([]oci_network_firewall.UrlPattern, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urls", stateDataIndex)
			converted, err := s.mapToUrlPattern(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("urls") {
			request.Urls = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	response, err := s.Client.UpdateUrlList(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.UrlList
	return nil
}

func (s *NetworkFirewallNetworkFirewallPolicyUrlListResourceCrud) Delete() error {
	request := oci_network_firewall.DeleteUrlListRequest{}

	if networkFirewallPolicyId, ok := s.D.GetOkExists("network_firewall_policy_id"); ok {
		tmp := networkFirewallPolicyId.(string)
		request.NetworkFirewallPolicyId = &tmp
	}

	if urlListName, ok := s.D.GetOkExists("name"); ok {
		tmp := urlListName.(string)
		request.UrlListName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_firewall")

	_, err := s.Client.DeleteUrlList(context.Background(), request)
	return err
}

func (s *NetworkFirewallNetworkFirewallPolicyUrlListResourceCrud) SetData() error {

	urlListName, networkFirewallPolicyId, err := parseNetworkFirewallPolicySubResourceCompositeId(s.D.Id(), "urlLists")
	if err == nil {
		s.D.Set("network_firewall_policy_id", &networkFirewallPolicyId)
		s.D.Set("name", &urlListName)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ParentResourceId != nil {
		s.D.Set("parent_resource_id", *s.Res.ParentResourceId)
	}

	if s.Res.TotalUrls != nil {
		s.D.Set("total_urls", *s.Res.TotalUrls)
	}

	urls := []interface{}{}
	for _, item := range s.Res.Urls {
		urls = append(urls, UrlPatternToMap(item))
	}
	s.D.Set("urls", urls)

	return nil
}

func UrlListSummaryToMap(obj oci_network_firewall.UrlListSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ParentResourceId != nil {
		result["parent_resource_id"] = string(*obj.ParentResourceId)
	}

	if obj.TotalUrls != nil {
		result["total_urls"] = int(*obj.TotalUrls)
	}

	return result
}

func (s *NetworkFirewallNetworkFirewallPolicyUrlListResourceCrud) mapToUrlPattern(fieldKeyFormat string) (oci_network_firewall.UrlPattern, error) {
	var baseObject oci_network_firewall.UrlPattern
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("SIMPLE"):
		details := oci_network_firewall.SimpleUrlPattern{}
		if pattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pattern")); ok {
			tmp := pattern.(string)
			details.Pattern = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func UrlPatternToMap(obj oci_network_firewall.UrlPattern) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_network_firewall.SimpleUrlPattern:
		result["type"] = "SIMPLE"

		if v.Pattern != nil {
			result["pattern"] = string(*v.Pattern)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}
