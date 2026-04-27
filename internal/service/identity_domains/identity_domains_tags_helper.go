// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity_domains

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

const identityDomainsOciTagsField = "urnietfparamsscimschemasoracleidcsextension_oci_tags"

func suppressIgnoredStructuredDefinedTagsForOciTags(field string) schema.CustomizeDiffFunc {
	return func(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
		log.Printf("[DEBUG] Identity Domains ignore_defined_tags suppressor invoked: field=%q ignored_tags=%v", field, tfresource.DefinedTagsToSuppress)
		if len(tfresource.DefinedTagsToSuppress) == 0 {
			log.Printf("[DEBUG] Identity Domains ignore_defined_tags suppressor skipping: field=%q ignored_tags is empty", field)
			return nil
		}

		configuredTagsByBlock := configuredStructuredDefinedTagsFromRawConfig(d, field)
		oldRaw, newRaw := d.GetChange(field)
		log.Printf(
			"[DEBUG] Identity Domains ignore_defined_tags suppressor inputs: field=%q old=%s new=%s configured=%s",
			field,
			debugSummaryForStructuredDefinedTagsRaw(oldRaw),
			debugSummaryForStructuredDefinedTagsRaw(newRaw),
			debugSummaryForConfiguredStructuredDefinedTags(configuredTagsByBlock),
		)
		merged, changed := mergeIgnoredStructuredDefinedTags(oldRaw, newRaw, configuredTagsByBlock)
		log.Printf(
			"[DEBUG] Identity Domains ignore_defined_tags suppressor merge result: field=%q changed=%t merged=%s",
			field,
			changed,
			debugSummaryForStructuredDefinedTagsRaw(merged),
		)
		if !changed {
			log.Printf("[DEBUG] Identity Domains ignore_defined_tags suppressor no-op: field=%q", field)
			return nil
		}

		if err := d.SetNew(field, merged); err != nil {
			log.Printf("[DEBUG] Identity Domains ignore_defined_tags suppressor SetNew failed: field=%q err=%v", field, err)
			return err
		}

		log.Printf("[DEBUG] Identity Domains ignore_defined_tags suppressor SetNew applied: field=%q", field)
		return nil
	}
}

func identityDomainsIgnoreDefinedTagsCustomizeDiff(resourceName string, field string) schema.CustomizeDiffFunc {
	suppressor := suppressIgnoredStructuredDefinedTagsForOciTags(field)
	return func(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
		log.Printf(
			"[DEBUG] Identity Domains resource CustomizeDiff invoked: resource=%q field=%q id=%q has_change=%t ignored_tags=%v",
			resourceName,
			field,
			d.Id(),
			d.HasChange(field),
			tfresource.DefinedTagsToSuppress,
		)
		return suppressor(ctx, d, meta)
	}
}

func mergeIgnoredStructuredDefinedTags(oldRaw interface{}, newRaw interface{}, configuredTagsByBlock [][]map[string]interface{}) ([]interface{}, bool) {
	oldBlocks := normalizeOciTagsBlocks(oldRaw)
	if len(oldBlocks) == 0 {
		return toInterfaceSlice(normalizeOciTagsBlocks(newRaw)), false
	}

	newBlocks := normalizeOciTagsBlocks(newRaw)
	if len(newBlocks) == 0 {
		newBlocks = []map[string]interface{}{{}}
	}

	changed := false
	for i, oldBlock := range oldBlocks {
		var configuredTags []map[string]interface{}
		if len(configuredTagsByBlock) > i {
			configuredTags = cloneTags(configuredTagsByBlock[i])
		}

		if len(newBlocks) <= i {
			newBlocks = append(newBlocks, map[string]interface{}{})
			changed = true
		}

		oldTags := normalizeStructuredDefinedTags(oldBlock["defined_tags"])
		if len(oldTags) == 0 {
			if len(configuredTags) > 0 {
				newBlocks[i]["defined_tags"] = schema.NewSet(definedTagsHashCodeForSets, toInterfaceSlice(configuredTags))
			}
			continue
		}

		newTags := configuredTags
		existing := map[string]bool{}
		for _, tag := range newTags {
			if tagName := structuredDefinedTagName(tag); tagName != "" {
				existing[strings.ToLower(tagName)] = true
			}
		}

		for _, tag := range oldTags {
			tagName := structuredDefinedTagName(tag)
			if tagName == "" || !isIgnoredDefinedTag(tagName) || existing[strings.ToLower(tagName)] {
				continue
			}

			newTags = append(newTags, cloneMap(tag))
			existing[strings.ToLower(tagName)] = true
			changed = true
		}

		if len(newTags) > 0 {
			newBlocks[i]["defined_tags"] = schema.NewSet(definedTagsHashCodeForSets, toInterfaceSlice(newTags))
		} else if _, ok := newBlocks[i]["defined_tags"]; ok {
			newBlocks[i]["defined_tags"] = schema.NewSet(definedTagsHashCodeForSets, []interface{}{})
		}
	}

	return toInterfaceSlice(newBlocks), changed
}

func configuredStructuredDefinedTagsFromRawConfig(d *schema.ResourceDiff, field string) [][]map[string]interface{} {
	rawConfig := d.GetRawConfig()
	if rawConfig.IsNull() || !rawConfig.IsKnown() || !rawConfig.Type().IsObjectType() {
		return nil
	}

	fieldValue, ok := rawConfig.AsValueMap()[field]
	if !ok || fieldValue.IsNull() || !fieldValue.IsKnown() {
		return nil
	}

	return structuredDefinedTagsFromCtyBlocks(fieldValue)
}

func structuredDefinedTagsFromCtyBlocks(blocksValue cty.Value) [][]map[string]interface{} {
	if blocksValue.IsNull() || !blocksValue.IsKnown() || !canIterateCtyCollection(blocksValue) {
		return nil
	}

	result := make([][]map[string]interface{}, 0, blocksValue.LengthInt())
	blocksValue.ForEachElement(func(_ cty.Value, blockValue cty.Value) bool {
		if blockValue.IsNull() || !blockValue.IsKnown() || !blockValue.Type().IsObjectType() {
			result = append(result, nil)
			return false
		}

		blockAttrs := blockValue.AsValueMap()
		definedTagsValue, ok := blockAttrs["defined_tags"]
		if !ok {
			result = append(result, nil)
			return false
		}

		result = append(result, structuredDefinedTagsFromCtyValue(definedTagsValue))
		return false
	})

	return result
}

func structuredDefinedTagsFromCtyValue(value cty.Value) []map[string]interface{} {
	if value.IsNull() || !value.IsKnown() || !canIterateCtyCollection(value) {
		return nil
	}

	result := make([]map[string]interface{}, 0, value.LengthInt())
	value.ForEachElement(func(_ cty.Value, tagValue cty.Value) bool {
		if tagValue.IsNull() || !tagValue.IsKnown() || !tagValue.Type().IsObjectType() {
			return false
		}

		tagAttrs := tagValue.AsValueMap()
		tag := map[string]interface{}{}

		if namespace, ok := tagAttrs["namespace"]; ok && namespace.IsKnown() && !namespace.IsNull() {
			tag["namespace"] = namespace.AsString()
		}
		if key, ok := tagAttrs["key"]; ok && key.IsKnown() && !key.IsNull() {
			tag["key"] = key.AsString()
		}
		if val, ok := tagAttrs["value"]; ok && val.IsKnown() && !val.IsNull() {
			tag["value"] = val.AsString()
		}

		if structuredDefinedTagName(tag) != "" {
			result = append(result, tag)
		}
		return false
	})

	return result
}

func canIterateCtyCollection(value cty.Value) bool {
	return value.Type().IsTupleType() || value.Type().IsListType() || value.Type().IsSetType()
}

func normalizeOciTagsBlocks(raw interface{}) []map[string]interface{} {
	items, ok := raw.([]interface{})
	if !ok {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(items))
	for _, item := range items {
		if itemMap, ok := item.(map[string]interface{}); ok {
			result = append(result, cloneMap(itemMap))
		}
	}

	return result
}

func normalizeStructuredDefinedTags(raw interface{}) []map[string]interface{} {
	var items []interface{}
	switch v := raw.(type) {
	case nil:
		return nil
	case []interface{}:
		items = v
	case *schema.Set:
		items = v.List()
	default:
		return nil
	}

	result := make([]map[string]interface{}, 0, len(items))
	for _, item := range items {
		if itemMap, ok := item.(map[string]interface{}); ok {
			result = append(result, cloneMap(itemMap))
		}
	}

	return result
}

func structuredDefinedTagName(tag map[string]interface{}) string {
	namespace, _ := tag["namespace"].(string)
	key, _ := tag["key"].(string)
	if namespace == "" || key == "" {
		return ""
	}

	return namespace + "." + key
}

func isIgnoredDefinedTag(tagName string) bool {
	for _, ignoredTag := range tfresource.DefinedTagsToSuppress {
		if strings.EqualFold(tagName, ignoredTag) {
			return true
		}
	}

	return false
}

func cloneMap(src map[string]interface{}) map[string]interface{} {
	if src == nil {
		return nil
	}

	dst := make(map[string]interface{}, len(src))
	for key, value := range src {
		dst[key] = value
	}

	return dst
}

func cloneTags(src []map[string]interface{}) []map[string]interface{} {
	if src == nil {
		return nil
	}

	dst := make([]map[string]interface{}, 0, len(src))
	for _, item := range src {
		dst = append(dst, cloneMap(item))
	}

	return dst
}

func toInterfaceSlice[T any](items []T) []interface{} {
	result := make([]interface{}, 0, len(items))
	for _, item := range items {
		result = append(result, item)
	}

	return result
}

func debugSummaryForStructuredDefinedTagsRaw(raw interface{}) string {
	if raw == nil {
		return "<nil>"
	}

	blocks := normalizeOciTagsBlocks(raw)
	if len(blocks) == 0 {
		if rawBlocks, ok := raw.([]interface{}); ok {
			return fmt.Sprintf("blocks=%d", len(rawBlocks))
		}
		return fmt.Sprintf("type=%T", raw)
	}

	summaries := make([]string, 0, len(blocks))
	for i, block := range blocks {
		summaries = append(summaries, fmt.Sprintf("block[%d]=%v", i, structuredDefinedTagDetails(normalizeStructuredDefinedTags(block["defined_tags"]))))
	}

	return strings.Join(summaries, "; ")
}

func debugSummaryForConfiguredStructuredDefinedTags(blocks [][]map[string]interface{}) string {
	if len(blocks) == 0 {
		return "blocks=0"
	}

	summaries := make([]string, 0, len(blocks))
	for i, block := range blocks {
		summaries = append(summaries, fmt.Sprintf("block[%d]=%v", i, structuredDefinedTagDetails(block)))
	}

	return strings.Join(summaries, "; ")
}

func structuredDefinedTagDetails(tags []map[string]interface{}) []string {
	if len(tags) == 0 {
		return []string{}
	}

	details := make([]string, 0, len(tags))
	for _, tag := range tags {
		if name := structuredDefinedTagName(tag); name != "" {
			value, _ := tag["value"].(string)
			details = append(details, fmt.Sprintf("%s=%q", name, value))
		}
	}

	sort.Strings(details)
	return details
}
