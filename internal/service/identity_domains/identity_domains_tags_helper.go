// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity_domains

import (
	"context"
	"strings"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

const identityDomainsOciTagsField = "urnietfparamsscimschemasoracleidcsextension_oci_tags"

func suppressIgnoredStructuredDefinedTagsForOciTags(field string) schema.CustomizeDiffFunc {
	return func(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
		if len(tfresource.DefinedTagsToSuppress) == 0 {
			return nil
		}

		configuredTagsByBlock := configuredStructuredDefinedTagsFromRawConfig(d, field)
		oldRaw, newRaw := d.GetChange(field)
		merged, changed := mergeIgnoredStructuredDefinedTags(oldRaw, newRaw, configuredTagsByBlock)
		if !changed {
			return nil
		}

		return d.SetNew(field, merged)
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

	// If the user did not configure the OCI tags block at all, preserve the
	// planned tags Terraform is already carrying and only merge ignored tags on top.
	userConfiguredField := configuredTagsByBlock != nil

	changed := false
	for i, oldBlock := range oldBlocks {
		if len(newBlocks) <= i {
			newBlocks = append(newBlocks, map[string]interface{}{})
			changed = true
		}

		oldTags := normalizeStructuredDefinedTags(oldBlock["defined_tags"])
		if len(oldTags) == 0 {
			continue
		}

		var newTags []map[string]interface{}
		if userConfiguredField && len(configuredTagsByBlock) > i {
			newTags = cloneTags(configuredTagsByBlock[i])
		} else {
			newTags = normalizeStructuredDefinedTags(newBlocks[i]["defined_tags"])
		}

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
