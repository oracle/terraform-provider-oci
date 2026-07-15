// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"
	"net"
	"reflect"
	"slices"
	"strconv"
	"strings"

	"github.com/hashicorp/go-cty/cty"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"
)

type coreWorkRequestClient interface {
	GetWorkRequest(context.Context, oci_work_requests.GetWorkRequestRequest) (oci_work_requests.GetWorkRequestResponse, error)
	ListWorkRequestErrors(context.Context, oci_work_requests.ListWorkRequestErrorsRequest) (oci_work_requests.ListWorkRequestErrorsResponse, error)
}

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

// This applies the differences between the regular DRG route table schema and
// the one we supply for default DRG route table management resources.
func ConvertToDefaultDrgRouteTableSchema(resourceSchema *schema.Resource) *schema.Resource {
	if resourceSchema == nil {
		return nil
	}

	resourceSchema.Importer = &schema.ResourceImporter{
		State: ImportDefaultDrgRouteTableResource,
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

	delete(resourceSchema.Schema, "drg_id")

	return resourceSchema
}

func ImportDefaultDrgRouteTableResource(d *schema.ResourceData, value interface{}) ([]*schema.ResourceData, error) {
	err := d.Set("manage_default_resource_id", d.Id())
	return []*schema.ResourceData{d}, err
}

// validateCoreWorkRequestStatus is intentionally scoped to Core VCN/Subnet
// callers. For these Core work requests, the top-level status is authoritative
// when it disagrees with a matching resource action that still reports
// IN_PROGRESS after the generic waiter returns.
func validateCoreWorkRequestStatus(ctx context.Context, workRequestClient coreWorkRequestClient, workRequestId *string, entityType string, disableFoundRetries bool) error {
	if workRequestId == nil {
		return nil
	}

	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "work_request")
	request := oci_work_requests.GetWorkRequestRequest{
		WorkRequestId: workRequestId,
	}
	request.RequestMetadata.RetryPolicy = retryPolicy

	response, err := workRequestClient.GetWorkRequest(ctx, request)
	if err != nil {
		return err
	}

	switch response.Status {
	case oci_work_requests.WorkRequestStatusFailed, oci_work_requests.WorkRequestStatusCanceled:
		return coreWorkRequestStatusError(ctx, workRequestClient, workRequestId, retryPolicy, entityType, response.Status)
	default:
		return nil
	}
}

// coreWorkRequestStatusError returns the service-provided error messages for a
// terminal failed/canceled Core work request, falling back to a deterministic
// message when the work request has no error entries.
func coreWorkRequestStatusError(ctx context.Context, workRequestClient coreWorkRequestClient, workRequestId *string, retryPolicy *oci_common.RetryPolicy, entityType string, status oci_work_requests.WorkRequestStatusEnum) error {
	request := oci_work_requests.ListWorkRequestErrorsRequest{
		WorkRequestId: workRequestId,
	}
	request.RequestMetadata.RetryPolicy = retryPolicy

	response, err := workRequestClient.ListWorkRequestErrors(ctx, request)
	if err != nil {
		return err
	}

	errorMessages := make([]string, 0, len(response.Items))
	for _, workRequestError := range response.Items {
		if workRequestError.Message != nil {
			errorMessages = append(errorMessages, *workRequestError.Message)
		}
	}

	errorMessage := strings.Join(errorMessages, "\n")
	if errorMessage == "" {
		errorMessage = "no work request error details were returned"
	}

	return fmt.Errorf("work request did not succeed, workId: %s, entity: %s, status: %s. Message: %s", *workRequestId, entityType, status, errorMessage)
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

	oldAddress, oldPrefix, oldIsCidr, oldOk := parseIpv6DiffValue(old)
	newAddress, newPrefix, newIsCidr, newOk := parseIpv6DiffValue(new)
	if !oldOk || !newOk || oldIsCidr != newIsCidr {
		return false
	}
	if oldAddress != newAddress {
		return false
	}
	if oldIsCidr {
		return oldPrefix == newPrefix
	}
	return true
}

func parseIpv6DiffValue(value string) (address string, prefixLength int, isCidr bool, ok bool) {
	ipPart := value
	prefixLength = -1

	if strings.Contains(value, "/") {
		parts := strings.Split(value, "/")
		if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
			return "", -1, false, false
		}

		prefix, err := strconv.Atoi(parts[1])
		if err != nil || prefix < 0 || prefix > 128 {
			return "", -1, false, false
		}

		ipPart = parts[0]
		prefixLength = prefix
		isCidr = true
	}

	parsedIP := net.ParseIP(ipPart)
	if parsedIP == nil || parsedIP.To4() != nil {
		return "", -1, false, false
	}

	normalizedIP := parsedIP.To16()
	if normalizedIP == nil {
		return "", -1, false, false
	}

	return normalizedIP.String(), prefixLength, isCidr, true
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

// normalizeByoipv6CidrDetailsForDiff replaces config entries for already-known
// BYO IPv6 blocks with their matching state entries while preserving config
// order. This suppresses diffs caused by computed range ids or canonical CIDR
// spelling without changing customer list intent.
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

	// Index the existing state entries by canonical IPv6 block for fast lookups.
	for _, item := range oldItems {
		data, block, ok := byoipv6DetailMapWithBlock(item)
		if !ok {
			continue
		}

		oldDetailsByBlock[convertToCanonical(block)] = data
	}

	if len(oldDetailsByBlock) == 0 {
		return nil, false
	}

	normalized := make([]interface{}, 0, len(newItems))

	// Preserve config order. For entries that already exist in state, reuse the
	// state map so computed fields do not create a needless diff.
	for _, item := range newItems {
		_, block, ok := byoipv6DetailMapWithBlock(item)
		if !ok {
			normalized = append(normalized, item)
			continue
		}

		if existing, exists := oldDetailsByBlock[convertToCanonical(block)]; exists {
			normalized = append(normalized, existing)
			continue
		}

		normalized = append(normalized, item)
	}

	if reflect.DeepEqual(normalized, newItems) {
		return nil, false
	}

	return normalized, true
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

// subnetIpv6PatchChangeSet is the normalized view of only the IPv6 fields that
// can participate in PatchSubnet for this resource.
//
// The key design constraint is that the IPv4 workflow must remain untouched.
// This structure therefore excludes cidr_block and ipv4cidr_blocks entirely and
// exists only to let the update path reason about the two IPv6 fields that now
// have conditional patch behavior.
type subnetIpv6PatchChangeSet struct {
	ipv6CidrBlockChanged  bool
	oldIpv6CidrBlock      string
	newIpv6CidrBlock      string
	currentIpv6CidrBlock  string
	ipv6CidrBlockInConfig bool
	ipv6CidrBlocksChanged bool
	oldIpv6CidrBlocks     []string
	newIpv6CidrBlocks     []string
}

// buildSubnetIpv6PatchChangeSet gathers the IPv6 patch-eligible diff once so
// the router can make a single decision and reuse the same normalized values in
// both the patch and legacy IPv6 helpers.
func buildSubnetIpv6PatchChangeSet(d *schema.ResourceData) subnetIpv6PatchChangeSet {
	changeSet := subnetIpv6PatchChangeSet{}

	if ipv6CidrBlock, ok := d.GetOkExists("ipv6cidr_block"); ok {
		changeSet.currentIpv6CidrBlock = ipv6CidrBlock.(string)
	}
	changeSet.ipv6CidrBlockInConfig = isAttributeConfigured(d, "ipv6cidr_block")

	if _, ok := d.GetOkExists("ipv6cidr_block"); ok && d.HasChange("ipv6cidr_block") {
		oldRaw, newRaw := d.GetChange("ipv6cidr_block")
		changeSet.ipv6CidrBlockChanged = true
		changeSet.oldIpv6CidrBlock = oldRaw.(string)
		changeSet.newIpv6CidrBlock = newRaw.(string)
	}

	if _, ok := d.GetOkExists("ipv6cidr_blocks"); ok && d.HasChange("ipv6cidr_blocks") {
		oldRaw, newRaw := d.GetChange("ipv6cidr_blocks")
		changeSet.ipv6CidrBlocksChanged = true
		changeSet.oldIpv6CidrBlocks = interfaceSliceToStringSlice(oldRaw)
		changeSet.newIpv6CidrBlocks = interfaceSliceToStringSlice(newRaw)
	}

	return changeSet
}

// changedFieldCount reports how many PatchSubnet-eligible IPv6 fields are
// changing in the current operation.
//
// Counting only these two fields is what keeps the patch logic from bleeding
// back into the IPv4 workflow or the regular UpdateSubnet cidr_block path.
func (c subnetIpv6PatchChangeSet) changedFieldCount() int {
	count := 0
	if c.ipv6CidrBlockChanged {
		count++
	}
	if c.ipv6CidrBlocksChanged {
		count++
	}
	return count
}

// shouldUsePatch decides whether the IPv6 part of the requested subnet update
// must move to PatchSubnet.
//
// The decision is intentionally conservative:
//   - if both ipv6cidr_block and ipv6cidr_blocks changed, PatchSubnet is required
//     because the legacy APIs cannot express that as one coherent update
//   - if ipv6cidr_block is a replacement, PatchSubnet is required even when it
//     is the only changed field because the legacy UpdateSubnet path only
//     models the simple "empty -> value" add case
//   - if ipv6cidr_blocks changes anywhere except a single add/remove at the end
//     of the list, PatchSubnet is required; the legacy helper is reserved for
//     append/remove-tail compatibility only
//
// Everything else stays on the existing code paths by design.
func (c subnetIpv6PatchChangeSet) shouldUsePatch() (bool, error) {
	if c.changedFieldCount() > 1 {
		return true, nil
	}

	if c.isIpv6CidrBlockReplacement() {
		return true, nil
	}

	if c.ipv6CidrBlocksChanged {
		return !isLegacySingleEndIpv6CidrBlockEdit(c.oldIpv6CidrBlocks, c.newIpv6CidrBlocks), nil
	}

	return false, nil
}

func (c subnetIpv6PatchChangeSet) isIpv6CidrBlockReplacement() bool {
	return c.ipv6CidrBlockChanged && c.oldIpv6CidrBlock != "" && c.newIpv6CidrBlock != ""
}

func (s *CoreSubnetResourceCrud) getSubnet() (oci_core.GetSubnetResponse, error) {
	request := oci_core.GetSubnetRequest{}

	tmp := s.D.Id()
	request.SubnetId = &tmp
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	return s.Client.GetSubnet(context.Background(), request)
}

func (c subnetIpv6PatchChangeSet) buildPatchInstructions() ([]oci_core.PatchSubnetInstruction, error) {
	var instructions []oci_core.PatchSubnetInstruction

	if c.ipv6CidrBlockChanged && !c.ipv6CidrBlocksChanged {
		instructions = append(instructions, buildSubnetReplaceInstruction("ipv6CidrBlock", map[string]interface{}{
			"cidr": c.newIpv6CidrBlock,
		}))
	}

	if c.ipv6CidrBlocksChanged {
		instructions = append(instructions, buildSubnetReplaceInstruction("ipv6CidrBlocks", map[string]interface{}{
			"cidrs": c.ipv6CidrBlocksForPatch(),
		}))
	}

	return instructions, nil
}

func (c subnetIpv6PatchChangeSet) ipv6CidrBlocksForPatch() []string {
	cidrs := append([]string{}, c.newIpv6CidrBlocks...)

	if !c.ipv6CidrBlockChanged && !c.ipv6CidrBlockInConfig {
		return cidrs
	}

	cidrBlock := c.currentIpv6CidrBlock
	if c.ipv6CidrBlockChanged {
		cidrBlock = c.newIpv6CidrBlock
	}

	if cidrBlock == "" {
		return cidrs
	}

	if slices.ContainsFunc(cidrs, isIPv6CidrIdentical(cidrBlock)) {
		return cidrs
	}

	return append(cidrs, cidrBlock)
}

func isAttributeConfigured(d *schema.ResourceData, attributeName string) bool {
	configValue, diagnostics := d.GetRawConfigAt(cty.GetAttrPath(attributeName))
	if diagnostics.HasError() {
		return false
	}
	return !configValue.IsNull()
}

func buildSubnetReplaceInstruction(selection string, value interface{}) oci_core.PatchSubnetInstruction {
	valueCopy := value
	return oci_core.PatchSubnetReplaceInstruction{
		Selection: &selection,
		Value:     &valueCopy,
	}
}

func interfaceSliceToStringSlice(raw interface{}) []string {
	interfaces, ok := raw.([]interface{})
	if !ok {
		return []string{}
	}

	result := make([]string, len(interfaces))
	for i := range interfaces {
		if interfaces[i] != nil {
			result[i] = interfaces[i].(string)
		}
	}

	return result
}

// vcnIpv6PatchChangeSet is the normalized view of the VCN IPv6 fields that can
// participate in PatchVcn.
//
// The update router uses this to keep the legacy single add/remove APIs in
// place for narrow changes while switching broader IPv6 mutations to PatchVcn.
type vcnIpv6PatchChangeSet struct {
	byoipv6CidrDetailsChanged bool
	oldByoipv6CidrDetails     []oci_core.Byoipv6CidrDetails
	newByoipv6CidrDetails     []oci_core.Byoipv6CidrDetails
	byoipv6CidrDetailsPresent bool
	currentByoipv6CidrDetails []oci_core.Byoipv6CidrDetails
	ipv6PrivateCidrChanged    bool
	oldIpv6PrivateCidrs       []string
	newIpv6PrivateCidrs       []string
	ipv6PrivateCidrPresent    bool
	currentIpv6PrivateCidrs   []string
}

func buildVcnIpv6PatchChangeSet(d *schema.ResourceData) (vcnIpv6PatchChangeSet, error) {
	changeSet := vcnIpv6PatchChangeSet{}

	if byoipv6CidrDetails, ok := d.GetOkExists("byoipv6cidr_details"); ok {
		currentDetails, err := interfaceToByoipv6CidrDetailsSlice(byoipv6CidrDetails)
		if err != nil {
			return changeSet, err
		}
		changeSet.byoipv6CidrDetailsPresent = true
		changeSet.currentByoipv6CidrDetails = currentDetails
	}

	if changeSet.byoipv6CidrDetailsPresent && d.HasChange("byoipv6cidr_details") {
		oldRaw, newRaw := d.GetChange("byoipv6cidr_details")
		oldDetails, err := interfaceToByoipv6CidrDetailsSlice(oldRaw)
		if err != nil {
			return changeSet, err
		}
		newDetails, err := interfaceToByoipv6CidrDetailsSlice(newRaw)
		if err != nil {
			return changeSet, err
		}

		changeSet.byoipv6CidrDetailsChanged = true
		changeSet.oldByoipv6CidrDetails = oldDetails
		changeSet.newByoipv6CidrDetails = newDetails
	}

	if ipv6PrivateCidrBlocks, ok := d.GetOkExists("ipv6private_cidr_blocks"); ok {
		changeSet.ipv6PrivateCidrPresent = true
		changeSet.currentIpv6PrivateCidrs = interfaceSliceToStringSlice(ipv6PrivateCidrBlocks)
	}

	if changeSet.ipv6PrivateCidrPresent && d.HasChange("ipv6private_cidr_blocks") {
		oldRaw, newRaw := d.GetChange("ipv6private_cidr_blocks")
		changeSet.ipv6PrivateCidrChanged = true
		changeSet.oldIpv6PrivateCidrs = interfaceSliceToStringSlice(oldRaw)
		changeSet.newIpv6PrivateCidrs = interfaceSliceToStringSlice(newRaw)
	}

	return changeSet, nil
}

func (c vcnIpv6PatchChangeSet) changedFieldCount() int {
	count := 0
	if c.byoipv6CidrDetailsChanged {
		count++
	}
	if c.ipv6PrivateCidrChanged {
		count++
	}
	return count
}

func (c vcnIpv6PatchChangeSet) shouldUsePatch() (bool, error) {
	if c.changedFieldCount() > 1 {
		return true, nil
	}

	if c.byoipv6CidrDetailsChanged {
		return !isLegacySingleEndByoipv6CidrDetailsEdit(c.oldByoipv6CidrDetails, c.newByoipv6CidrDetails), nil
	}

	if c.ipv6PrivateCidrChanged {
		return !isLegacySingleEndIpv6CidrBlockEdit(c.oldIpv6PrivateCidrs, c.newIpv6PrivateCidrs), nil
	}

	return false, nil
}

func isLegacySingleEndIpv6CidrBlockEdit(oldBlocks []string, newBlocks []string) bool {
	if len(newBlocks) == len(oldBlocks)+1 {
		return ipv6CidrBlockListsEqual(oldBlocks, newBlocks[:len(oldBlocks)])
	}
	if len(oldBlocks) == len(newBlocks)+1 {
		return ipv6CidrBlockListsEqual(oldBlocks[:len(newBlocks)], newBlocks)
	}
	return false
}

func ipv6CidrBlockListsEqual(left []string, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	for i := range left {
		if !isIPv6CidrIdentical(left[i])(right[i]) {
			return false
		}
	}
	return true
}

func isLegacySingleEndByoipv6CidrDetailsEdit(oldDetails []oci_core.Byoipv6CidrDetails, newDetails []oci_core.Byoipv6CidrDetails) bool {
	if len(newDetails) == len(oldDetails)+1 {
		return byoipv6CidrDetailsListsEqual(oldDetails, newDetails[:len(oldDetails)])
	}
	if len(oldDetails) == len(newDetails)+1 {
		return byoipv6CidrDetailsListsEqual(oldDetails[:len(newDetails)], newDetails)
	}
	return false
}

func byoipv6CidrDetailsListsEqual(left []oci_core.Byoipv6CidrDetails, right []oci_core.Byoipv6CidrDetails) bool {
	if len(left) != len(right) {
		return false
	}
	for i := range left {
		if !equalByoipv6PatchDetails(left[i], right[i]) {
			return false
		}
	}
	return true
}

func (c vcnIpv6PatchChangeSet) buildPatchInstructions() ([]oci_core.PatchVcnInstruction, error) {
	var instructions []oci_core.PatchVcnInstruction

	if c.byoipv6CidrDetailsPresent {
		instructions = append(instructions, buildVcnReplaceInstruction("byoipv6CidrDetails", map[string]interface{}{
			"cidrs": byoipv6CidrDetailsToPatchPayload(c.currentByoipv6CidrDetails),
		}))
	} else if c.byoipv6CidrDetailsChanged {
		instructions = append(instructions, buildVcnReplaceInstruction("byoipv6CidrDetails", map[string]interface{}{
			"cidrs": byoipv6CidrDetailsToPatchPayload(c.newByoipv6CidrDetails),
		}))
	}

	if c.ipv6PrivateCidrPresent {
		instructions = append(instructions, buildVcnReplaceInstruction("ipv6PrivateCidrBlocks", map[string]interface{}{
			"cidrs": c.currentIpv6PrivateCidrs,
		}))
	} else if c.ipv6PrivateCidrChanged {
		instructions = append(instructions, buildVcnReplaceInstruction("ipv6PrivateCidrBlocks", map[string]interface{}{
			"cidrs": c.newIpv6PrivateCidrs,
		}))
	}

	return instructions, nil
}

func buildVcnReplaceInstruction(selection string, value interface{}) oci_core.PatchVcnInstruction {
	valueCopy := value
	return oci_core.PatchVcnReplaceInstruction{
		Selection: &selection,
		Value:     &valueCopy,
	}
}

func interfaceToByoipv6CidrDetailsSlice(raw interface{}) ([]oci_core.Byoipv6CidrDetails, error) {
	interfaces, ok := raw.([]interface{})
	if !ok {
		return []oci_core.Byoipv6CidrDetails{}, nil
	}

	result := make([]oci_core.Byoipv6CidrDetails, 0, len(interfaces))
	for _, item := range interfaces {
		detailMap, ok := item.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("expected byoipv6cidr_details entry to be a map, got %T", item)
		}

		detail := oci_core.Byoipv6CidrDetails{}

		rangeIDRaw, ok := detailMap["byoipv6range_id"]
		if !ok {
			return nil, fmt.Errorf("missing byoipv6range_id in byoipv6cidr_details entry")
		}
		rangeID, ok := rangeIDRaw.(string)
		if !ok {
			return nil, fmt.Errorf("expected byoipv6range_id to be a string, got %T", rangeIDRaw)
		}
		detail.Byoipv6RangeId = &rangeID

		cidrRaw, ok := detailMap["ipv6cidr_block"]
		if !ok {
			return nil, fmt.Errorf("missing ipv6cidr_block in byoipv6cidr_details entry")
		}
		cidr, ok := cidrRaw.(string)
		if !ok {
			return nil, fmt.Errorf("expected ipv6cidr_block to be a string, got %T", cidrRaw)
		}
		detail.Ipv6CidrBlock = &cidr

		result = append(result, detail)
	}

	return result, nil
}

func byoipv6CidrDetailsToPatchPayload(details []oci_core.Byoipv6CidrDetails) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(details))
	for _, detail := range details {
		payload := map[string]interface{}{}
		if detail.Byoipv6RangeId != nil {
			payload["byoipv6RangeId"] = *detail.Byoipv6RangeId
		}
		if detail.Ipv6CidrBlock != nil {
			payload["ipv6CidrBlock"] = *detail.Ipv6CidrBlock
		}
		result = append(result, payload)
	}
	return result
}

func equalByoipv6PatchDetails(left oci_core.Byoipv6CidrDetails, right oci_core.Byoipv6CidrDetails) bool {
	leftCidr := ""
	if left.Ipv6CidrBlock != nil {
		leftCidr = convertToCanonical(*left.Ipv6CidrBlock)
	}
	rightCidr := ""
	if right.Ipv6CidrBlock != nil {
		rightCidr = convertToCanonical(*right.Ipv6CidrBlock)
	}

	return leftCidr == rightCidr
}
