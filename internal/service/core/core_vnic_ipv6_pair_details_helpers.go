// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"
	"log"
	"net"
	"sort"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func isIpv6PairDetailsExplicitlyConfigured(d *schema.ResourceData, createVnicDetails map[string]interface{}) bool {
	if isIpv6PairDetailsExplicitlyConfiguredInRawConfig(d) {
		return true
	}

	// Fallback to preserving the prior state shape when it strongly resembles user-specified input.
	return hasLikelyExplicitIpv6PairDetails(createVnicDetails)
}

func isIpv6PairDetailsExplicitlyConfiguredInRawConfig(d *schema.ResourceData) bool {
	if d == nil {
		return false
	}

	path := cty.Path{
		cty.GetAttrStep{Name: "create_vnic_details"},
		cty.IndexStep{Key: cty.NumberIntVal(0)},
		cty.GetAttrStep{Name: "ipv6address_ipv6subnet_cidr_pair_details"},
	}

	value, diags := d.GetRawConfigAt(path)
	if diags.HasError() {
		return false
	}

	if value.IsNull() {
		return false
	}

	if !value.IsKnown() {
		return true
	}

	if !value.CanIterateElements() {
		return true
	}

	// Treat only non-empty lists as explicitly user configured.
	if value.LengthInt() == 0 {
		return false
	}

	return true
}

func hasLikelyExplicitIpv6PairDetails(createVnicDetails map[string]interface{}) bool {
	if createVnicDetails == nil {
		return false
	}

	rawPairDetails, exists := createVnicDetails["ipv6address_ipv6subnet_cidr_pair_details"]
	if !exists || rawPairDetails == nil {
		return false
	}

	pairDetails, ok := rawPairDetails.([]interface{})
	if !ok || len(pairDetails) == 0 {
		return false
	}

	hasExplicitSubnetCidr := false
	for _, pairDetail := range pairDetails {
		pairMap, ok := pairDetail.(map[string]interface{})
		if !ok || pairMap == nil {
			continue
		}

		if ipv6IDValue, exists := pairMap["ipv6id"]; exists {
			if ipv6ID, ok := ipv6IDValue.(string); ok && ipv6ID != "" {
				// Entries with ipv6id are typically API-derived; do not treat as explicit config.
				return false
			}
		}

		if ipv6SubnetCidrValue, exists := pairMap["ipv6subnet_cidr"]; exists {
			if ipv6SubnetCidr, ok := ipv6SubnetCidrValue.(string); ok && ipv6SubnetCidr != "" {
				hasExplicitSubnetCidr = true
			}
		}
	}

	return hasExplicitSubnetCidr
}

func deriveIpv6PairDetailsFromVnic(vnic *oci_core.Vnic, virtualNetworkClient *oci_core.VirtualNetworkClient, disableNotFoundRetries bool) []interface{} {
	if vnic == nil {
		return []interface{}{}
	}

	type ipv6Details struct {
		Ipv6Id         string
		Ipv6SubnetCidr string
	}
	ipv6DetailsByAddress := map[string]ipv6Details{}

	for _, ipv6Address := range vnic.Ipv6Addresses {
		if ipv6Address == "" {
			continue
		}
		if _, exists := ipv6DetailsByAddress[ipv6Address]; !exists {
			ipv6DetailsByAddress[ipv6Address] = ipv6Details{}
		}
	}

	if virtualNetworkClient != nil && vnic.Id != nil {
		listRequest := oci_core.ListIpv6sRequest{
			VnicId: vnic.Id,
		}
		listRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(disableNotFoundRetries, "core")

		for {
			listResponse, err := virtualNetworkClient.ListIpv6s(context.Background(), listRequest)
			if err != nil {
				log.Printf("[DEBUG] Unable to list IPv6 details for VNIC during refresh. (VNIC ID: %q, Error: %q)", *vnic.Id, err)
				break
			}

			for _, ipv6 := range listResponse.Items {
				if ipv6.IpAddress == nil || *ipv6.IpAddress == "" {
					continue
				}

				ipv6Address := *ipv6.IpAddress
				details := ipv6DetailsByAddress[ipv6Address]

				if ipv6.Id != nil && *ipv6.Id != "" {
					details.Ipv6Id = *ipv6.Id
				}
				if ipv6.CidrPrefixLength != nil {
					if ipv6SubnetCidr := deriveIpv6SubnetCidrFromAddress(ipv6Address, *ipv6.CidrPrefixLength); ipv6SubnetCidr != "" {
						details.Ipv6SubnetCidr = ipv6SubnetCidr
					}
				}

				ipv6DetailsByAddress[ipv6Address] = details
			}

			if listResponse.OpcNextPage == nil {
				break
			}
			listRequest.Page = listResponse.OpcNextPage
		}
	}

	orderedIpv6Addresses := make([]string, 0, len(ipv6DetailsByAddress))
	for ipv6Address := range ipv6DetailsByAddress {
		orderedIpv6Addresses = append(orderedIpv6Addresses, ipv6Address)
	}
	sort.Strings(orderedIpv6Addresses)

	result := make([]interface{}, 0, len(orderedIpv6Addresses))
	for _, ipv6Address := range orderedIpv6Addresses {
		details := ipv6DetailsByAddress[ipv6Address]
		item := map[string]interface{}{
			"ipv6address": ipv6Address,
		}
		if details.Ipv6Id != "" {
			item["ipv6id"] = details.Ipv6Id
		}
		if details.Ipv6SubnetCidr != "" {
			item["ipv6subnet_cidr"] = details.Ipv6SubnetCidr
		}
		result = append(result, item)
	}

	return result
}

func deriveIpv6SubnetCidrFromAddress(ipv6Address string, cidrPrefixLength int) string {
	if cidrPrefixLength < 0 || cidrPrefixLength > 128 {
		return ""
	}

	parsedAddress := net.ParseIP(ipv6Address)
	if parsedAddress == nil {
		return ""
	}

	normalizedIpv6 := parsedAddress.To16()
	if normalizedIpv6 == nil {
		return ""
	}

	subnetMask := net.CIDRMask(cidrPrefixLength, 128)
	return fmt.Sprintf("%s/%d", normalizedIpv6.Mask(subnetMask).String(), cidrPrefixLength)
}
