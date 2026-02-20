// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"
	"net"
	"strings"

	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"
)

// This applies the differences between the regular schema and the one
// we supply for default resources, and returns the schema for a default resource
func ConvertToDefaultVcnResourceSchema(resourceSchema *schema.Resource) *schema.Resource {
	if resourceSchema == nil {
		return nil
	}

	resourceSchema.Importer = &schema.ResourceImporter{
		State: ImportDefaultVcnResource,
	}

	resourceSchema.Schema["manage_default_resource_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	}

	resourceSchema.Schema["compartment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	}

	delete(resourceSchema.Schema, "vcn_id")

	return resourceSchema
}

func ImportDefaultVcnResource(d *schema.ResourceData, value interface{}) ([]*schema.ResourceData, error) {
	err := d.Set("manage_default_resource_id", d.Id())
	return []*schema.ResourceData{d}, err
}

func LaunchOptionsToMap(obj *oci_core.LaunchOptions) map[string]interface{} {
	result := map[string]interface{}{}

	result["boot_volume_type"] = string(obj.BootVolumeType)

	result["firmware"] = string(obj.Firmware)

	result["network_type"] = string(obj.NetworkType)

	result["remote_data_volume_type"] = string(obj.RemoteDataVolumeType)

	if obj.IsConsistentVolumeNamingEnabled != nil {
		result["is_consistent_volume_naming_enabled"] = bool(*obj.IsConsistentVolumeNamingEnabled)
	}

	if obj.IsPvEncryptionInTransitEnabled != nil {
		result["is_pv_encryption_in_transit_enabled"] = bool(*obj.IsPvEncryptionInTransitEnabled)
	}

	return result
}

func getBackupPolicyId(assetId *string, client *oci_core.BlockstorageClient) (*string, error) {
	request := oci_core.GetVolumeBackupPolicyAssetAssignmentRequest{}
	request.AssetId = assetId
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := client.GetVolumeBackupPolicyAssetAssignment(context.Background(), request)
	if err != nil {
		return nil, err
	}

	if len(response.Items) > 0 {
		policyAssignment := response.Items[0]
		return policyAssignment.PolicyId, nil
	} else {
		return nil, nil
	}
}

func (s *CoreVolumeBackupResourceCrud) createBlockStorageSourceRegionClient(region string) error {
	if s.SourceRegionClient == nil {
		sourceBlockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(*s.Client.ConfigurationProvider())
		if err != nil {
			return fmt.Errorf("cannot Create client for the source region: %v", err)
		}
		err = tf_client.ConfigureClientVar(&sourceBlockStorageClient.BaseClient)
		if err != nil {
			return fmt.Errorf("cannot configure client for the source region: %v", err)
		}
		s.SourceRegionClient = &sourceBlockStorageClient
	}
	s.SourceRegionClient.SetRegion(region)

	return nil
}

func (s *CoreVolumeGroupBackupResourceCrud) createBlockStorageSourceRegionClient(region string) error {
	if s.SourceRegionClient == nil {
		sourceBlockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(*s.Client.ConfigurationProvider())
		if err != nil {
			return fmt.Errorf("cannot Create client for the source region: %v", err)
		}
		err = tf_client.ConfigureClientVar(&sourceBlockStorageClient.BaseClient)
		if err != nil {
			return fmt.Errorf("cannot configure client for the source region: %v", err)
		}
		s.SourceRegionClient = &sourceBlockStorageClient
	}
	s.SourceRegionClient.SetRegion(region)

	return nil
}

func (s *CoreBootVolumeBackupResourceCrud) createBlockStorageSourceRegionClient(region string) error {
	if s.SourceRegionClient == nil {
		sourceBlockStorageClient, err := oci_core.NewBlockstorageClientWithConfigurationProvider(*s.Client.ConfigurationProvider())
		if err != nil {
			return fmt.Errorf("cannot Create client for the source region: %v", err)
		}
		err = tf_client.ConfigureClientVar(&sourceBlockStorageClient.BaseClient)
		if err != nil {
			return fmt.Errorf("cannot configure client for the source region: %v", err)
		}
		s.SourceRegionClient = &sourceBlockStorageClient
	}
	s.SourceRegionClient.SetRegion(region)

	return nil
}

// This before suppression is required because
// `fd00:aaaa:0123::/48` in request comes back as `fd00:aaaa:123::/48` in response
func ipv6CompressionDiffSuppressFunction(key string, old string, new string, d *schema.ResourceData) bool {
	if old == "" || new == "" {
		return false
	}
	oldIp := strings.Split(old, "/")
	newIp := strings.Split(new, "/")
	if len(oldIp) < 2 || len(newIp) < 2 {
		return false
	}
	oldParsedIp := net.ParseIP(oldIp[0])
	oldSubnetMask := oldIp[1]
	newParsedIp := net.ParseIP(newIp[0])
	newSubnetMask := newIp[1]
	return strings.EqualFold(oldParsedIp.String(), newParsedIp.String()) && strings.EqualFold(oldSubnetMask, newSubnetMask)
}

func ipv6Cidr_blocksSuppressFunction(key string, old string, new string, d *schema.ResourceData) bool {
	if key == "ipv6cidr_blocks.#" {
		if old == "" || new == "" {
			return false
		}

		// old and new should represent size of list
		// if there is a difference in size between old and new values, we have a diff
		if old != new {
			return false
		}

		return true
	}

	return ipv6CompressionDiffSuppressFunction(key, old, new, d)
}

func suppressDiffIfOldIsEmptyString(key string, old string, new string, d *schema.ResourceData) bool {
	return old == "" && new != ""
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// This function computes the list of byoIpv6CidrBlocks present in a list of byoIpV6CidrDetails from config
func computeIPv6BlocksFromBYOIPv6Details(byoIpV6CidrDetails interface{}) []string {
	// Handle nil
	if byoIpV6CidrDetails == nil {
		return []string{}
	}

	// Convert the input interface to a slice of interfaces
	items, ok := byoIpV6CidrDetails.([]interface{})
	if !ok {
		// If it's not a list, return empty
		return []string{}
	}
	byoipv6Cidrs := make([]string, 0, len(items))

	// Iterate over the slice
	for _, item := range items {
		// Assert that each item is a map
		data, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		if val, ok := data["ipv6cidr_block"]; ok {
			// Assert the value is a string
			if blockStr, ok := val.(string); ok && blockStr != "" {
				byoipv6Cidrs = append(byoipv6Cidrs, blockStr)
			}
		}
	}
	return byoipv6Cidrs
}

func isIPv6CidrIdentical(elementToFind string) func(currentElement string) bool {
	return func(currentElement string) bool {
		return convertToCanonical(elementToFind) == convertToCanonical(currentElement)
	}
}

func convertToCanonical(block string) string {
	splitString := strings.Split(block, ":")

	// Separate out prefix from the last octet
	prefixSplit := strings.Split(splitString[len(splitString)-1], "/")

	splitString[len(splitString)-1] = prefixSplit[0]

	final := []string{"0000", "0000", "0000", "0000", "0000", "0000", "0000", "0000"}

	for i := 0; i < len(splitString); i++ {

		// append 4 - len(i) 0's to the left, and add it to string along with :
		final[i] = strings.Repeat("0", 4-len(splitString[i])) + splitString[i]
	}
	result := strings.Join(final, ":")

	return result + "/" + prefixSplit[1]
}
