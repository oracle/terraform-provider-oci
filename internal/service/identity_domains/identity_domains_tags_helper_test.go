// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity_domains

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func TestUnitMergeIgnoredStructuredDefinedTags(t *testing.T) {
	oldSuppressed := tfresource.DefinedTagsToSuppress
	defer func() {
		tfresource.DefinedTagsToSuppress = oldSuppressed
	}()

	tfresource.DefinedTagsToSuppress = []string{"Oracle-Tags.CreatedBy", "Oracle-Tags.CreatedOn"}

	oldRaw := []interface{}{
		map[string]interface{}{
			"defined_tags": schema.NewSet(definedTagsHashCodeForSets, []interface{}{
				map[string]interface{}{"namespace": "Oracle-Tags", "key": "CreatedBy", "value": "service"},
				map[string]interface{}{"namespace": "Oracle-Tags", "key": "CreatedOn", "value": "2026-03-16T20:45:36.983Z"},
				map[string]interface{}{"namespace": "MyNamespace", "key": "Managed", "value": "custom"},
			}),
		},
	}

	newRaw := []interface{}{
		map[string]interface{}{
			"defined_tags": schema.NewSet(definedTagsHashCodeForSets, []interface{}{
				map[string]interface{}{"namespace": "MyNamespace", "key": "Managed", "value": "custom"},
			}),
		},
	}

	configuredTagsByBlock := [][]map[string]interface{}{
		{
			{"namespace": "MyNamespace", "key": "Managed", "value": "custom"},
		},
	}

	merged, changed := mergeIgnoredStructuredDefinedTags(oldRaw, newRaw, configuredTagsByBlock)
	if !changed {
		t.Fatalf("expected mergeIgnoredStructuredDefinedTags to report a change")
	}

	mergedBlocks := normalizeOciTagsBlocks(merged)
	if len(mergedBlocks) != 1 {
		t.Fatalf("expected exactly one OCI tags block, got %d", len(mergedBlocks))
	}

	mergedTags := normalizeStructuredDefinedTags(mergedBlocks[0]["defined_tags"])
	if len(mergedTags) != 3 {
		t.Fatalf("expected three defined tags after merge, got %d", len(mergedTags))
	}

	names := map[string]bool{}
	for _, tag := range mergedTags {
		names[structuredDefinedTagName(tag)] = true
	}

	for _, expected := range []string{"Oracle-Tags.CreatedBy", "Oracle-Tags.CreatedOn", "MyNamespace.Managed"} {
		if !names[expected] {
			t.Fatalf("expected merged defined tags to include %s", expected)
		}
	}
}

func TestUnitMergeIgnoredStructuredDefinedTagsNoIgnoredTagsConfigured(t *testing.T) {
	oldSuppressed := tfresource.DefinedTagsToSuppress
	defer func() {
		tfresource.DefinedTagsToSuppress = oldSuppressed
	}()

	tfresource.DefinedTagsToSuppress = nil

	oldRaw := []interface{}{
		map[string]interface{}{
			"defined_tags": schema.NewSet(definedTagsHashCodeForSets, []interface{}{
				map[string]interface{}{"namespace": "Oracle-Tags", "key": "CreatedBy", "value": "service"},
			}),
		},
	}

	newRaw := []interface{}{
		map[string]interface{}{
			"defined_tags": schema.NewSet(definedTagsHashCodeForSets, []interface{}{}),
		},
	}

	merged, changed := mergeIgnoredStructuredDefinedTags(oldRaw, newRaw, nil)
	if changed {
		t.Fatalf("expected no merge when ignore_defined_tags is empty, got %#v", merged)
	}
}

func TestUnitMergeIgnoredStructuredDefinedTagsNoOciTagsBlockInConfig(t *testing.T) {
	oldSuppressed := tfresource.DefinedTagsToSuppress
	defer func() {
		tfresource.DefinedTagsToSuppress = oldSuppressed
	}()

	tfresource.DefinedTagsToSuppress = []string{"Oracle-Tags.CreatedBy", "Oracle-Tags.CreatedOn"}

	oldRaw := []interface{}{
		map[string]interface{}{
			"defined_tags": schema.NewSet(definedTagsHashCodeForSets, []interface{}{
				map[string]interface{}{"namespace": "Oracle-Tags", "key": "CreatedBy", "value": "service"},
				map[string]interface{}{"namespace": "Oracle-Tags", "key": "CreatedOn", "value": "2026-03-16T20:45:36.983Z"},
			}),
		},
	}

	newRaw := []interface{}{
		map[string]interface{}{
			"defined_tags": schema.NewSet(definedTagsHashCodeForSets, []interface{}{}),
		},
	}

	merged, changed := mergeIgnoredStructuredDefinedTags(oldRaw, newRaw, nil)
	if !changed {
		t.Fatalf("expected mergeIgnoredStructuredDefinedTags to report a change")
	}

	mergedBlocks := normalizeOciTagsBlocks(merged)
	if len(mergedBlocks) != 1 {
		t.Fatalf("expected exactly one OCI tags block, got %d", len(mergedBlocks))
	}

	mergedTags := normalizeStructuredDefinedTags(mergedBlocks[0]["defined_tags"])
	if len(mergedTags) != 2 {
		t.Fatalf("expected two ignored defined tags after merge, got %d: %v", len(mergedTags), mergedTags)
	}

	names := map[string]bool{}
	for _, tag := range mergedTags {
		names[structuredDefinedTagName(tag)] = true
	}

	for _, expected := range []string{"Oracle-Tags.CreatedBy", "Oracle-Tags.CreatedOn"} {
		if !names[expected] {
			t.Fatalf("expected merged defined tags to include %s, got %v", expected, names)
		}
	}
}

func TestUnitMergeIgnoredStructuredDefinedTagsArbitraryNamespaces(t *testing.T) {
	oldSuppressed := tfresource.DefinedTagsToSuppress
	defer func() {
		tfresource.DefinedTagsToSuppress = oldSuppressed
	}()

	tfresource.DefinedTagsToSuppress = []string{
		"Oracle-Tags.CreatedBy",
		"Oracle-Tags.CreatedOn",
		"A-Team.Creator",
		"Schedule.AnyDay",
		"Usage-Management.Expires",
		"Usage-Management.Creator",
	}

	oldRaw := []interface{}{
		map[string]interface{}{
			"defined_tags": schema.NewSet(definedTagsHashCodeForSets, []interface{}{
				map[string]interface{}{"namespace": "Oracle-Tags", "key": "CreatedBy", "value": "user"},
				map[string]interface{}{"namespace": "Oracle-Tags", "key": "CreatedOn", "value": "2026-01-01T00:00:00.000Z"},
				map[string]interface{}{"namespace": "A-Team", "key": "Creator", "value": "automation"},
				map[string]interface{}{"namespace": "Schedule", "key": "AnyDay", "value": "true"},
				map[string]interface{}{"namespace": "Usage-Management", "key": "Expires", "value": "never"},
				map[string]interface{}{"namespace": "Usage-Management", "key": "Creator", "value": "ops"},
			}),
		},
	}

	newRaw := []interface{}{
		map[string]interface{}{
			"defined_tags": schema.NewSet(definedTagsHashCodeForSets, []interface{}{}),
		},
	}

	merged, changed := mergeIgnoredStructuredDefinedTags(oldRaw, newRaw, nil)
	if !changed {
		t.Fatalf("expected mergeIgnoredStructuredDefinedTags to report a change")
	}

	mergedBlocks := normalizeOciTagsBlocks(merged)
	if len(mergedBlocks) != 1 {
		t.Fatalf("expected exactly one OCI tags block, got %d", len(mergedBlocks))
	}

	mergedTags := normalizeStructuredDefinedTags(mergedBlocks[0]["defined_tags"])
	if len(mergedTags) != 6 {
		t.Fatalf("expected six ignored defined tags after merge, got %d: %v", len(mergedTags), mergedTags)
	}

	names := map[string]bool{}
	for _, tag := range mergedTags {
		names[structuredDefinedTagName(tag)] = true
	}

	for _, expected := range []string{
		"Oracle-Tags.CreatedBy",
		"Oracle-Tags.CreatedOn",
		"A-Team.Creator",
		"Schedule.AnyDay",
		"Usage-Management.Expires",
		"Usage-Management.Creator",
	} {
		if !names[expected] {
			t.Fatalf("expected merged defined tags to include %s, got %v", expected, names)
		}
	}
}

func TestUnitMergeIgnoredStructuredDefinedTagsKeepsIgnoredButRemovesCustomTags(t *testing.T) {
	oldSuppressed := tfresource.DefinedTagsToSuppress
	defer func() {
		tfresource.DefinedTagsToSuppress = oldSuppressed
	}()

	tfresource.DefinedTagsToSuppress = []string{"Oracle-Tags.CreatedBy", "Oracle-Tags.CreatedOn"}

	oldRaw := []interface{}{
		map[string]interface{}{
			"defined_tags": schema.NewSet(definedTagsHashCodeForSets, []interface{}{
				map[string]interface{}{"namespace": "Oracle-Tags", "key": "CreatedBy", "value": "service"},
				map[string]interface{}{"namespace": "Oracle-Tags", "key": "CreatedOn", "value": "2026-03-16T20:45:36.983Z"},
				map[string]interface{}{"namespace": "testns", "key": "TestTag", "value": "TestValue"},
			}),
		},
	}

	// Simulate computed fallback in newRaw still containing the old custom tag.
	newRaw := []interface{}{
		map[string]interface{}{
			"defined_tags": schema.NewSet(definedTagsHashCodeForSets, []interface{}{
				map[string]interface{}{"namespace": "Oracle-Tags", "key": "CreatedBy", "value": "service"},
				map[string]interface{}{"namespace": "Oracle-Tags", "key": "CreatedOn", "value": "2026-03-16T20:45:36.983Z"},
				map[string]interface{}{"namespace": "testns", "key": "TestTag", "value": "TestValue"},
			}),
		},
	}

	merged, changed := mergeIgnoredStructuredDefinedTags(oldRaw, newRaw, [][]map[string]interface{}{{}})
	if !changed {
		t.Fatalf("expected mergeIgnoredStructuredDefinedTags to report a change")
	}

	mergedBlocks := normalizeOciTagsBlocks(merged)
	if len(mergedBlocks) != 1 {
		t.Fatalf("expected exactly one OCI tags block, got %d", len(mergedBlocks))
	}

	mergedTags := normalizeStructuredDefinedTags(mergedBlocks[0]["defined_tags"])
	if len(mergedTags) != 2 {
		t.Fatalf("expected only ignored defined tags after merge, got %d", len(mergedTags))
	}

	names := map[string]bool{}
	for _, tag := range mergedTags {
		names[structuredDefinedTagName(tag)] = true
	}

	for _, expected := range []string{"Oracle-Tags.CreatedBy", "Oracle-Tags.CreatedOn"} {
		if !names[expected] {
			t.Fatalf("expected merged defined tags to include %s", expected)
		}
	}

	if names["testns.TestTag"] {
		t.Fatalf("expected merged defined tags to remove testns.TestTag")
	}
}
