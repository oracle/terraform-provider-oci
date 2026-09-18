// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License Version 2.0

package provider

import (
	"context"
	"maps"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	frameworktypes "github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	tfclient "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/globalvar"
)

func TestProviderConstructorsReturnFreshInstances(t *testing.T) {
	first := Provider()
	second := Provider()
	if first == second {
		t.Fatal("Provider returned a shared SDKv2 provider")
	}
	firstEmbedded := NewSDKv2ProviderForInProcess()
	secondEmbedded := NewSDKv2ProviderForInProcess()
	if firstEmbedded == secondEmbedded {
		t.Fatal("NewSDKv2ProviderForInProcess returned a shared SDKv2 provider")
	}
	assertSDKv2ResourceMapsIsolated(t, firstEmbedded.ResourcesMap, secondEmbedded.ResourcesMap)
	assertSDKv2ResourceMapsIsolated(t, firstEmbedded.DataSourcesMap, secondEmbedded.DataSourcesMap)
}

func TestSelectiveInProcessProvider(t *testing.T) {
	const resourceName = "oci_identity_tag_namespace"
	enabledServices := maps.Clone(oci_common.OciSdkEnabledServicesMap)
	first, err := NewSDKv2ProviderForInProcessResources(resourceName, resourceName)
	if err != nil {
		t.Fatalf("construct selective provider: %v", err)
	}
	second, err := NewSDKv2ProviderForInProcessResources(resourceName)
	if err != nil {
		t.Fatalf("construct second selective provider: %v", err)
	}
	if len(first.ResourcesMap) != 1 || first.ResourcesMap[resourceName] == nil {
		t.Fatalf("selective resource map = %v, want only %q", first.ResourcesMap, resourceName)
	}
	if len(first.DataSourcesMap) != 0 {
		t.Fatalf("selective provider retained %d data sources, want 0", len(first.DataSourcesMap))
	}
	if err := first.InternalValidate(); err != nil {
		t.Fatalf("selective provider validation failed: %v", err)
	}
	if first.ResourcesMap[resourceName] == second.ResourcesMap[resourceName] {
		t.Fatal("selective providers share a mutable resource schema")
	}

	if _, err := NewSDKv2ProviderForInProcessResources("oci_missing_resource"); err == nil {
		t.Fatal("selective provider accepted an unknown resource")
	}
	if !maps.Equal(enabledServices, oci_common.OciSdkEnabledServicesMap) {
		t.Fatal("selective schema construction changed OCI SDK enabled services")
	}
}

func TestConfigurationOnlyInProcessProvider(t *testing.T) {
	p := NewSDKv2ProviderForInProcessConfiguration()
	if len(p.Schema) == 0 {
		t.Fatal("configuration-only provider has no provider schema")
	}
	if p.ConfigureFunc == nil {
		t.Fatal("configuration-only provider has no ConfigureFunc")
	}
	if len(p.ResourcesMap) != 0 || len(p.DataSourcesMap) != 0 {
		t.Fatalf("configuration-only provider retained resources=%d dataSources=%d", len(p.ResourcesMap), len(p.DataSourcesMap))
	}
	if err := p.InternalValidate(); err != nil {
		t.Fatalf("configuration-only provider validation failed: %v", err)
	}
}

func TestClonedSDKv2ResourceReturnsLazyClientInitializationErrors(t *testing.T) {
	resource := cloneSDKv2Resource(&schema.Resource{
		Importer: &schema.ResourceImporter{
			StateContext: func(_ context.Context, _ *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				meta.(*tfclient.OracleClients).GetClient("oci_missing.Client")
				return nil, nil
			},
		},
		Create: func(_ *schema.ResourceData, meta interface{}) error {
			meta.(*tfclient.OracleClients).GetClient("oci_missing.Client")
			return nil
		},
		ReadContext: func(_ context.Context, _ *schema.ResourceData, meta interface{}) diag.Diagnostics {
			meta.(*tfclient.OracleClients).GetClient("oci_missing.Client")
			return nil
		},
	})
	clients := &tfclient.OracleClients{SdkClientMap: make(map[string]interface{})}

	err := resource.Create(nil, clients)
	if err == nil || !strings.Contains(err.Error(), "provider clients are not configured") {
		t.Fatalf("Create lazy-client error = %v", err)
	}
	diagnostics := resource.ReadContext(t.Context(), nil, clients)
	if !diagnostics.HasError() || !strings.Contains(diagnostics[0].Summary, "provider clients are not configured") {
		t.Fatalf("ReadContext lazy-client diagnostics = %v", diagnostics)
	}
	_, err = resource.Importer.StateContext(t.Context(), nil, clients)
	if err == nil || !strings.Contains(err.Error(), "provider clients are not configured") {
		t.Fatalf("import lazy-client error = %v", err)
	}
}

func TestClonedSDKv2ResourceDoesNotMaskUnrelatedPanics(t *testing.T) {
	resource := cloneSDKv2Resource(&schema.Resource{
		Create: func(*schema.ResourceData, interface{}) error {
			panic("programming error")
		},
	})

	defer func() {
		if recovered := recover(); recovered != "programming error" {
			t.Fatalf("recovered panic = %v", recovered)
		}
	}()
	_ = resource.Create(nil, nil)
}

func assertSDKv2ResourceMapsIsolated(t *testing.T, first, second map[string]*schema.Resource) {
	t.Helper()
	if len(first) != len(second) {
		t.Fatalf("resource map sizes differ: first=%d second=%d", len(first), len(second))
	}
	for name, firstResource := range first {
		secondResource, ok := second[name]
		if !ok {
			t.Fatalf("resource %q is missing from the second provider", name)
		}
		assertSDKv2ResourcesIsolated(t, name, firstResource, secondResource)
	}
}

func assertSDKv2ResourcesIsolated(t *testing.T, path string, first, second *schema.Resource) {
	t.Helper()
	if first == nil || second == nil {
		if first != second {
			t.Fatalf("resource %q differs between provider instances", path)
		}
		return
	}
	if first == second {
		t.Fatalf("resource %q is shared between provider instances", path)
	}
	assertSDKv2SchemaMapsIsolated(t, path, first.Schema, second.Schema)

	if first.Importer != nil && first.Importer == second.Importer {
		t.Fatalf("resource %q importer is shared between provider instances", path)
	}
	if first.Timeouts != nil && first.Timeouts == second.Timeouts {
		t.Fatalf("resource %q timeouts are shared between provider instances", path)
	}
	if first.Identity != nil && first.Identity == second.Identity {
		t.Fatalf("resource %q identity is shared between provider instances", path)
	}
}

func assertSDKv2SchemaMapsIsolated(t *testing.T, path string, first, second map[string]*schema.Schema) {
	t.Helper()
	if len(first) != len(second) {
		t.Fatalf("schema map sizes differ for %q: first=%d second=%d", path, len(first), len(second))
	}
	for name, firstSchema := range first {
		secondSchema, ok := second[name]
		if !ok {
			t.Fatalf("schema %q is missing from the second provider", path+"."+name)
		}
		assertSDKv2SchemasIsolated(t, path+"."+name, firstSchema, secondSchema)
	}
}

func assertSDKv2SchemasIsolated(t *testing.T, path string, first, second *schema.Schema) {
	t.Helper()
	if first == nil || second == nil {
		if first != second {
			t.Fatalf("schema %q differs between provider instances", path)
		}
		return
	}
	if first == second {
		t.Fatalf("schema %q is shared between provider instances", path)
	}

	switch firstElement := first.Elem.(type) {
	case *schema.Schema:
		secondElement, ok := second.Elem.(*schema.Schema)
		if !ok {
			t.Fatalf("schema element type differs for %q", path)
		}
		assertSDKv2SchemasIsolated(t, path+"[]", firstElement, secondElement)
	case *schema.Resource:
		secondElement, ok := second.Elem.(*schema.Resource)
		if !ok {
			t.Fatalf("schema element type differs for %q", path)
		}
		assertSDKv2ResourcesIsolated(t, path+"[]", firstElement, secondElement)
	}
}

func TestValidateInProcessProviderConfig(t *testing.T) {
	defaults := schema.TestResourceDataRaw(t, SchemaMap(), map[string]any{})
	if err := validateInProcessProviderConfig(defaults); err != nil {
		t.Fatalf("default provider configuration was rejected: %v", err)
	}

	tests := []struct {
		name   string
		field  string
		config map[string]any
	}{
		{
			name:   "defined tags",
			field:  globalvar.DefinedTagsToIgnore,
			config: map[string]any{globalvar.DefinedTagsToIgnore: []any{"Oracle-Tags.CreatedBy"}},
		},
		{
			name:   "realm-specific endpoint",
			field:  globalvar.RealmSpecificServiceEndpointTemplateEnabled,
			config: map[string]any{globalvar.RealmSpecificServiceEndpointTemplateEnabled: false},
		},
		{
			name:   "dual-stack endpoint",
			field:  globalvar.DualStackEndpointEnabled,
			config: map[string]any{globalvar.DualStackEndpointEnabled: false},
		},
		{
			name:   "disable automatic retries",
			field:  globalvar.DisableAutoRetriesAttrName,
			config: map[string]any{globalvar.DisableAutoRetriesAttrName: true},
		},
		{
			name:   "retry duration",
			field:  globalvar.RetryDurationSecondsAttrName,
			config: map[string]any{globalvar.RetryDurationSecondsAttrName: 30},
		},
		{
			name:   "retry configuration file",
			field:  globalvar.RetriesConfigFile,
			config: map[string]any{globalvar.RetriesConfigFile: "retries.json"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			configured := schema.TestResourceDataRaw(t, SchemaMap(), tt.config)
			err := validateInProcessProviderConfig(configured)
			if err == nil || !strings.Contains(err.Error(), tt.field) {
				t.Fatalf("validation error = %v, want unsupported option %q", err, tt.field)
			}
		})
	}
}

func TestFrameworkInProcessProviderRejectsIgnoreDefinedTags(t *testing.T) {
	p := &ociPluginProvider{inProcess: true}
	value := frameworktypes.ListValueMust(
		frameworktypes.StringType,
		[]attr.Value{frameworktypes.StringValue("Oracle-Tags.CreatedBy")},
	)
	if diags := p.setIgnoreDefinedTags(t.Context(), value); diags.HasError() {
		t.Fatalf("decode ignore_defined_tags: %v", diags)
	}
	if len(p.ignoreDefinedTags) != 1 || p.ignoreDefinedTags[0] != "Oracle-Tags.CreatedBy" {
		t.Fatalf("unexpected decoded ignore_defined_tags: %v", p.ignoreDefinedTags)
	}

	err := p.validateInProcessProviderConfig()
	if err == nil {
		t.Fatal("Framework in-process provider accepted ignore_defined_tags")
	}
	if !strings.Contains(err.Error(), globalvar.DefinedTagsToIgnore) {
		t.Fatalf("validation error does not identify ignore_defined_tags: %v", err)
	}
}

func TestFrameworkIgnoreDefinedTagsAcceptsUnknownValue(t *testing.T) {
	p := &ociPluginProvider{
		inProcess:         true,
		ignoreDefinedTags: []string{"stale-value"},
	}

	diags := p.setIgnoreDefinedTags(t.Context(), frameworktypes.ListUnknown(frameworktypes.StringType))
	if diags.HasError() {
		t.Fatalf("unknown ignore_defined_tags produced diagnostics: %v", diags)
	}
	if len(p.ignoreDefinedTags) != 0 {
		t.Fatalf("unknown ignore_defined_tags retained decoded values: %v", p.ignoreDefinedTags)
	}
}

func TestFrameworkEndpointOptionsPreserveConfiguredState(t *testing.T) {
	tests := []struct {
		name  string
		value frameworktypes.Bool
		want  string
	}{
		{name: "unset", value: frameworktypes.BoolNull(), want: ""},
		{name: "unknown", value: frameworktypes.BoolUnknown(), want: ""},
		{name: "false", value: frameworktypes.BoolValue(false), want: "false"},
		{name: "true", value: frameworktypes.BoolValue(true), want: "true"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ociPluginProvider{}
			p.SetDefaults(&ociProviderModel{
				RealmSpecificServiceEndpointTemplateEnabled: tt.value,
				DualStackEndpointEnabled:                    tt.value,
			})

			if p.realmSpecificServiceEndpointTemplateEnabled != tt.want {
				t.Fatalf("realm-specific endpoint option = %q, want %q", p.realmSpecificServiceEndpointTemplateEnabled, tt.want)
			}
			if p.dualStackEndpointEnabled != tt.want {
				t.Fatalf("dual-stack endpoint option = %q, want %q", p.dualStackEndpointEnabled, tt.want)
			}
		})
	}
}

func TestFrameworkInProcessProviderRejectsExplicitEndpointOptions(t *testing.T) {
	for _, tt := range []struct {
		name      string
		configure func(*ociPluginProvider)
		want      string
	}{
		{
			name: "realm-specific false",
			configure: func(p *ociPluginProvider) {
				p.realmSpecificServiceEndpointTemplateEnabled = "false"
			},
			want: globalvar.RealmSpecificServiceEndpointTemplateEnabled,
		},
		{
			name: "dual-stack false",
			configure: func(p *ociPluginProvider) {
				p.dualStackEndpointEnabled = "false"
			},
			want: globalvar.DualStackEndpointEnabled,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			p := &ociPluginProvider{inProcess: true}
			tt.configure(p)
			err := p.validateInProcessProviderConfig()
			if err == nil || !strings.Contains(err.Error(), tt.want) {
				t.Fatalf("validation error = %v, want unsupported option %q", err, tt.want)
			}
		})
	}
}

func TestFrameworkInProcessProviderRejectsRetryOptions(t *testing.T) {
	tests := []struct {
		name      string
		field     string
		configure func(*ociPluginProvider)
	}{
		{
			name:  "disable automatic retries",
			field: globalvar.DisableAutoRetriesAttrName,
			configure: func(p *ociPluginProvider) {
				p.disableAutoRetries = true
			},
		},
		{
			name:  "retry duration",
			field: globalvar.RetryDurationSecondsAttrName,
			configure: func(p *ociPluginProvider) {
				p.retryDurationSeconds = 30
			},
		},
		{
			name:  "retry configuration file",
			field: globalvar.RetriesConfigFile,
			configure: func(p *ociPluginProvider) {
				p.retriesConfigFile = "retries.json"
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ociPluginProvider{inProcess: true}
			tt.configure(p)
			err := p.validateInProcessProviderConfig()
			if err == nil || !strings.Contains(err.Error(), tt.field) {
				t.Fatalf("validation error = %v, want unsupported option %q", err, tt.field)
			}
		})
	}
}

func TestSetAvoidWaitingForDeleteTargetFromEnv(t *testing.T) {
	original := AvoidWaitingForDeleteTarget
	t.Cleanup(func() {
		AvoidWaitingForDeleteTarget = original
	})

	tests := []struct {
		name      string
		inProcess bool
		want      bool
	}{
		{
			name: "Terraform CLI",
			want: true,
		},
		{
			name:      "in-process",
			inProcess: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("avoid_waiting_for_delete_target", "true")
			AvoidWaitingForDeleteTarget = false
			setAvoidWaitingForDeleteTargetFromEnv(tt.inProcess)
			if AvoidWaitingForDeleteTarget != tt.want {
				t.Fatalf("AvoidWaitingForDeleteTarget = %t, want %t", AvoidWaitingForDeleteTarget, tt.want)
			}
		})
	}
}

func TestInternalAndEmbeddedSDKv2SchemasRemainCompatible(t *testing.T) {
	cli := Provider()
	embedded := NewSDKv2ProviderForInProcess()
	if len(cli.Schema) != len(embedded.Schema) {
		t.Fatalf("provider schema size differs: CLI=%d embedded=%d", len(cli.Schema), len(embedded.Schema))
	}
	if len(cli.ResourcesMap) != len(embedded.ResourcesMap) {
		t.Fatalf("resource schema size differs: CLI=%d embedded=%d", len(cli.ResourcesMap), len(embedded.ResourcesMap))
	}
	if len(cli.DataSourcesMap) != len(embedded.DataSourcesMap) {
		t.Fatalf("data-source schema size differs: CLI=%d embedded=%d", len(cli.DataSourcesMap), len(embedded.DataSourcesMap))
	}
}

func BenchmarkInProcessProviderConstruction(b *testing.B) {
	b.Run("full", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			_ = NewSDKv2ProviderForInProcess()
		}
	})
	b.Run("identity-resource", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			if _, err := NewSDKv2ProviderForInProcessResources("oci_identity_tag_namespace"); err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("configuration-only", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			_ = NewSDKv2ProviderForInProcessConfiguration()
		}
	})
}
