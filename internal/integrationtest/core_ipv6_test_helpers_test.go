// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

type subnetIpv6PlanExpectation struct {
	blockFieldChanges       int
	blockFieldReplacements  int
	blocksFieldChanges      int
	blocksAdditions         int
	blocksRemovals          int
	blocksReplacementGroups int
}

type subnetIpv6PlanCheck struct {
	resourceAddress string
	expectation     subnetIpv6PlanExpectation
}

func (c subnetIpv6PlanCheck) CheckPlan(_ context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	change, err := findResourceChange(req.Plan, c.resourceAddress)
	if err != nil {
		resp.Error = err
		return
	}

	summary, err := summarizeSubnetIpv6PlanChange(change)
	if err != nil {
		resp.Error = err
		return
	}

	if summary.blockFieldChanges != c.expectation.blockFieldChanges {
		resp.Error = fmt.Errorf("resource %s ipv6cidr_block changes = %d, want %d", c.resourceAddress, summary.blockFieldChanges, c.expectation.blockFieldChanges)
		return
	}
	if summary.blockFieldReplacements != c.expectation.blockFieldReplacements {
		resp.Error = fmt.Errorf("resource %s ipv6cidr_block replacements = %d, want %d", c.resourceAddress, summary.blockFieldReplacements, c.expectation.blockFieldReplacements)
		return
	}
	if summary.blocksFieldChanges != c.expectation.blocksFieldChanges {
		resp.Error = fmt.Errorf("resource %s ipv6cidr_blocks field changes = %d, want %d", c.resourceAddress, summary.blocksFieldChanges, c.expectation.blocksFieldChanges)
		return
	}
	if summary.blocksAdditions != c.expectation.blocksAdditions {
		resp.Error = fmt.Errorf("resource %s ipv6cidr_blocks additions = %d, want %d", c.resourceAddress, summary.blocksAdditions, c.expectation.blocksAdditions)
		return
	}
	if summary.blocksRemovals != c.expectation.blocksRemovals {
		resp.Error = fmt.Errorf("resource %s ipv6cidr_blocks removals = %d, want %d", c.resourceAddress, summary.blocksRemovals, c.expectation.blocksRemovals)
		return
	}
	if summary.blocksReplacementGroups != c.expectation.blocksReplacementGroups {
		resp.Error = fmt.Errorf("resource %s ipv6cidr_blocks replacements = %d, want %d", c.resourceAddress, summary.blocksReplacementGroups, c.expectation.blocksReplacementGroups)
	}
}

type subnetIpv6PlanSummary struct {
	blockFieldChanges       int
	blockFieldReplacements  int
	blocksFieldChanges      int
	blocksAdditions         int
	blocksRemovals          int
	blocksReplacementGroups int
}

func findResourceChange(plan *tfjson.Plan, resourceAddress string) (*tfjson.ResourceChange, error) {
	if plan == nil {
		return nil, fmt.Errorf("plan is nil")
	}

	for _, change := range plan.ResourceChanges {
		if change != nil && change.Address == resourceAddress {
			return change, nil
		}
	}

	return nil, fmt.Errorf("resource change not found for %s", resourceAddress)
}

func summarizeSubnetIpv6PlanChange(change *tfjson.ResourceChange) (subnetIpv6PlanSummary, error) {
	summary := subnetIpv6PlanSummary{}
	if change == nil || change.Change == nil {
		return summary, fmt.Errorf("resource change or nested change is nil")
	}

	before, err := planValueMap(change.Change.Before)
	if err != nil {
		return summary, err
	}
	after, err := planValueMap(change.Change.After)
	if err != nil {
		return summary, err
	}

	beforeBlock, err := canonicalIpv6CidrLiteral(planStringValue(before["ipv6cidr_block"]))
	if err != nil {
		return summary, err
	}
	afterBlock, err := canonicalIpv6CidrLiteral(planStringValue(after["ipv6cidr_block"]))
	if err != nil {
		return summary, err
	}

	if beforeBlock != afterBlock {
		summary.blockFieldChanges = 1
		if beforeBlock != "" && afterBlock != "" {
			summary.blockFieldReplacements = 1
		}
	}

	beforeBlocks, err := canonicalIpv6CidrList(before["ipv6cidr_blocks"])
	if err != nil {
		return summary, err
	}
	afterBlocks, err := canonicalIpv6CidrList(after["ipv6cidr_blocks"])
	if err != nil {
		return summary, err
	}

	removed := ipv6CidrsMissingFrom(beforeBlocks, afterBlocks)
	added := ipv6CidrsMissingFrom(afterBlocks, beforeBlocks)

	if len(removed) > 0 || len(added) > 0 {
		summary.blocksFieldChanges = 1
	}
	summary.blocksAdditions = len(added)
	summary.blocksRemovals = len(removed)
	if len(added) < len(removed) {
		summary.blocksReplacementGroups = len(added)
	} else {
		summary.blocksReplacementGroups = len(removed)
	}

	return summary, nil
}

func planValueMap(value interface{}) (map[string]interface{}, error) {
	if value == nil {
		return map[string]interface{}{}, nil
	}

	valueMap, ok := value.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("expected plan object map, got %T", value)
	}

	return valueMap, nil
}

func planStringValue(value interface{}) string {
	if value == nil {
		return ""
	}

	if stringValue, ok := value.(string); ok {
		return stringValue
	}

	return ""
}

func canonicalIpv6CidrList(value interface{}) ([]string, error) {
	if value == nil {
		return []string{}, nil
	}

	switch typed := value.(type) {
	case []interface{}:
		result := make([]string, 0, len(typed))
		for _, entry := range typed {
			stringValue, ok := entry.(string)
			if !ok {
				return nil, fmt.Errorf("expected ipv6 cidr string entry, got %T", entry)
			}
			canonical, err := canonicalIpv6CidrLiteral(stringValue)
			if err != nil {
				return nil, err
			}
			result = append(result, canonical)
		}
		return result, nil
	case []string:
		result := make([]string, 0, len(typed))
		for _, entry := range typed {
			canonical, err := canonicalIpv6CidrLiteral(entry)
			if err != nil {
				return nil, err
			}
			result = append(result, canonical)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("expected ipv6 cidr list, got %T", value)
	}
}

func canonicalIpv6CidrLiteral(value string) (string, error) {
	if value == "" {
		return "", nil
	}

	ip, network, err := net.ParseCIDR(value)
	if err != nil {
		return "", fmt.Errorf("invalid ipv6 cidr %q: %w", value, err)
	}

	ip = ip.To16()
	if ip == nil {
		return "", fmt.Errorf("expected IPv6 cidr, got %q", value)
	}

	maskSize, _ := network.Mask.Size()
	parts := make([]string, 0, 8)
	for i := 0; i < len(ip); i += 2 {
		parts = append(parts, fmt.Sprintf("%02x%02x", ip[i], ip[i+1]))
	}

	return strings.Join(parts, ":") + "/" + strconv.Itoa(maskSize), nil
}

func ipv6CidrsMissingFrom(blocks []string, referenceBlocks []string) []string {
	missing := make([]string, 0)

	reference := make(map[string]struct{}, len(referenceBlocks))
	for _, block := range referenceBlocks {
		reference[block] = struct{}{}
	}

	for _, block := range blocks {
		if _, ok := reference[block]; !ok {
			missing = append(missing, block)
		}
	}

	return missing
}

func testCheckCanonicalResourceAttrEqualsLiteral(resourceName, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found in state: %s", resourceName)
		}

		got, ok := rs.Primary.Attributes[attr]
		if !ok {
			return fmt.Errorf("attribute not found in state: %s.%s", resourceName, attr)
		}

		gotCanonical, err := canonicalIpv6CidrLiteral(got)
		if err != nil {
			return err
		}
		wantCanonical, err := canonicalIpv6CidrLiteral(want)
		if err != nil {
			return err
		}

		if gotCanonical != wantCanonical {
			return fmt.Errorf("expected canonical %s (%s) to equal %s", attr, got, want)
		}

		return nil
	}
}

func testCheckCanonicalTypeSetContains(resourceName, attr string, expected []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found in state: %s", resourceName)
		}

		expectedCanonical := make(map[string]struct{}, len(expected))
		for _, block := range expected {
			canonical, err := canonicalIpv6CidrLiteral(block)
			if err != nil {
				return err
			}
			expectedCanonical[canonical] = struct{}{}
		}

		actualCanonical := make(map[string]struct{})
		prefix := attr + "."
		for key, value := range rs.Primary.Attributes {
			if !strings.HasPrefix(key, prefix) || key == attr+".#" {
				continue
			}

			canonical, err := canonicalIpv6CidrLiteral(value)
			if err != nil {
				return err
			}
			actualCanonical[canonical] = struct{}{}
		}

		for block := range expectedCanonical {
			if _, ok := actualCanonical[block]; !ok {
				return fmt.Errorf("expected %s to contain canonical block %s, actual blocks: %#v", attr, block, actualCanonical)
			}
		}

		return nil
	}
}

func testCheckCanonicalListEquals(resourceName, attr string, expected []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found in state: %s", resourceName)
		}

		countValue, ok := rs.Primary.Attributes[attr+".#"]
		if !ok {
			return fmt.Errorf("attribute count not found in state: %s.%s.#", resourceName, attr)
		}

		count, err := strconv.Atoi(countValue)
		if err != nil {
			return fmt.Errorf("invalid %s.%s.# value %q: %w", resourceName, attr, countValue, err)
		}
		if count != len(expected) {
			return fmt.Errorf("expected %s.%s count %d, got %d", resourceName, attr, len(expected), count)
		}

		for i, want := range expected {
			key := fmt.Sprintf("%s.%d", attr, i)
			got, ok := rs.Primary.Attributes[key]
			if !ok {
				return fmt.Errorf("attribute not found in state: %s.%s", resourceName, key)
			}

			gotCanonical, err := canonicalIpv6CidrLiteral(got)
			if err != nil {
				return err
			}
			wantCanonical, err := canonicalIpv6CidrLiteral(want)
			if err != nil {
				return err
			}
			if gotCanonical != wantCanonical {
				return fmt.Errorf("expected %s.%s canonical value %s, got %s", resourceName, key, wantCanonical, gotCanonical)
			}
		}

		return nil
	}
}

func testCheckCanonicalNestedIpv6CidrBlockListEquals(resourceName, attr string, expected []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found in state: %s", resourceName)
		}

		countValue, ok := rs.Primary.Attributes[attr+".#"]
		if !ok {
			return fmt.Errorf("attribute count not found in state: %s.%s.#", resourceName, attr)
		}

		count, err := strconv.Atoi(countValue)
		if err != nil {
			return fmt.Errorf("invalid %s.%s.# value %q: %w", resourceName, attr, countValue, err)
		}
		if count != len(expected) {
			return fmt.Errorf("expected %s.%s count %d, got %d", resourceName, attr, len(expected), count)
		}

		for i, want := range expected {
			key := fmt.Sprintf("%s.%d.ipv6cidr_block", attr, i)
			got, ok := rs.Primary.Attributes[key]
			if !ok {
				return fmt.Errorf("attribute not found in state: %s.%s", resourceName, key)
			}

			gotCanonical, err := canonicalIpv6CidrLiteral(got)
			if err != nil {
				return err
			}
			wantCanonical, err := canonicalIpv6CidrLiteral(want)
			if err != nil {
				return err
			}
			if gotCanonical != wantCanonical {
				return fmt.Errorf("expected %s.%s canonical value %s, got %s", resourceName, key, wantCanonical, gotCanonical)
			}
		}

		return nil
	}
}

func testCheckMissingOrZeroResourceAttr(resourceName, attr string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found in state: %s", resourceName)
		}

		value, ok := rs.Primary.Attributes[attr]
		if !ok || value == "0" || value == "" {
			return nil
		}

		return fmt.Errorf("expected %s.%s to be missing or zero, got %q", resourceName, attr, value)
	}
}
