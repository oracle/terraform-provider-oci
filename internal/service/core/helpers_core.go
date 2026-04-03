// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"
	"net"
	"reflect"
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

// normalizeByoipv6CidrDetailsForDiff reorders and replaces new detail entries
// so Terraform plans against stable state ordering rather than raw config
// ordering. Existing BYO IPv6 blocks keep their prior state entries and order;
// only genuinely new blocks are appended.
func normalizeByoipv6CidrDetailsForDiff(oldValue interface{}, newValue interface{}) ([]interface{}, bool) {
	oldItems, ok := oldValue.([]interface{})
	if !ok || len(oldItems) == 0 {
		return nil, false
	}

	newItems, ok := newValue.([]interface{})
	if !ok || len(newItems) == 0 {
		return nil, false
	}

	oldDetailsByBlock := make(map[string]map[string]interface{}, len(oldItems))
	oldOrder := make([]string, 0, len(oldItems))

	// Index the existing state entries by canonical IPv6 block for fast lookups.
	for _, item := range oldItems {
		data, block, ok := byoipv6DetailMapWithBlock(item)
		if !ok {
			continue
		}

		canonicalBlock := convertToCanonical(block)
		if _, exists := oldDetailsByBlock[canonicalBlock]; !exists {
			oldOrder = append(oldOrder, canonicalBlock)
		}
		oldDetailsByBlock[canonicalBlock] = data
	}

	if len(oldDetailsByBlock) == 0 {
		return nil, false
	}

	newDetailsByBlock := make(map[string]map[string]interface{}, len(newItems))

	for _, item := range newItems {
		data, block, ok := byoipv6DetailMapWithBlock(item)
		if !ok {
			continue
		}
		newDetailsByBlock[convertToCanonical(block)] = data
	}

	normalized := make([]interface{}, 0, len(newItems))
	matchedBlocks := make(map[string]struct{}, len(newItems))

	// Preserve the prior state order for blocks that still exist in config.
	for _, canonicalBlock := range oldOrder {
		if existing, ok := oldDetailsByBlock[canonicalBlock]; ok {
			if _, existsInConfig := newDetailsByBlock[canonicalBlock]; existsInConfig {
				normalized = append(normalized, existing)
				matchedBlocks[canonicalBlock] = struct{}{}
			}
		}
	}

	// Append only the genuinely new config entries in their declared order.
	for _, item := range newItems {
		_, block, ok := byoipv6DetailMapWithBlock(item)
		if !ok {
			normalized = append(normalized, item)
			continue
		}

		canonicalBlock := convertToCanonical(block)
		if _, alreadyPresent := matchedBlocks[canonicalBlock]; alreadyPresent {
			continue
		}

		normalized = append(normalized, item)
		matchedBlocks[canonicalBlock] = struct{}{}
	}

	if reflect.DeepEqual(normalized, newItems) {
		return nil, false
	}

	return normalized, true
}

// equalByoipv6DetailMaps reports whether two BYOIPv6 detail maps contain the
// same keys and values.
func equalByoipv6DetailMaps(left map[string]interface{}, right map[string]interface{}) bool {
	if len(left) != len(right) {
		return false
	}

	for key, leftValue := range left {
		rightValue, ok := right[key]
		if !ok || leftValue != rightValue {
			return false
		}
	}

	return true
}

func byoipv6DetailMapWithBlock(item interface{}) (map[string]interface{}, string, bool) {
	data, ok := item.(map[string]interface{})
	if !ok {
		return nil, "", false
	}

	block, ok := data["ipv6cidr_block"].(string)
	if !ok || block == "" {
		return nil, "", false
	}

	return data, block, true
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

func isEmptyString(input string) bool {
	return len(strings.TrimSpace(input)) == 0
}
