// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opensearch

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_opensearch "github.com/oracle/oci-go-sdk/v65/opensearch"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func TestOpensearchRemoteReindexUpdateCondition(t *testing.T) {
	const prefix = "remote_reindex_trust_config.0.remote_sources."
	for _, tc := range []struct {
		name string
		diff map[string]*terraform.ResourceAttrDiff
		want bool
	}{
		{name: "unchanged", diff: map[string]*terraform.ResourceAttrDiff{}, want: false},
		{name: "empty_sources", diff: map[string]*terraform.ResourceAttrDiff{
			prefix + "#": {Old: "1", New: "0"},
		}, want: true},
		{name: "endpoint_only", diff: map[string]*terraform.ResourceAttrDiff{
			prefix + "0.endpoint": {Old: "https://source-a.example.com:9200", New: "https://source-b.example.com:9200"},
		}, want: true},
		{name: "ip_only", diff: map[string]*terraform.ResourceAttrDiff{
			prefix + "0.endpoint_resolution_ip": {Old: "192.0.2.10", New: "192.0.2.11"},
		}, want: true},
		{name: "ca_bundle_only", diff: map[string]*terraform.ResourceAttrDiff{
			prefix + "0.ca_bundle_id": {Old: "original-ca", New: "replacement-ca"},
		}, want: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			data, err := schema.InternalMap(OpensearchOpensearchClusterResource().Schema).Data(&terraform.InstanceState{
				ID: "test-cluster",
				Attributes: map[string]string{
					"remote_reindex_trust_config.#":     "1",
					prefix + "#":                        "1",
					prefix + "0.endpoint":               "https://source-a.example.com:9200",
					prefix + "0.endpoint_resolution_ip": "192.0.2.10",
					prefix + "0.ca_bundle_id":           "original-ca",
				},
			}, &terraform.InstanceDiff{Attributes: tc.diff})
			if err != nil {
				t.Fatal(err)
			}
			crud := OpensearchOpensearchClusterResourceCrud{BaseCrud: tfresource.BaseCrud{D: data}}
			if got := crud.UpdateConditionMet(); got != tc.want {
				t.Fatalf("UpdateConditionMet() = %t, want %t for a remote-reindex-only change", got, tc.want)
			}
		})
	}
}

func TestOpensearchRemoteReindexSourceMapping(t *testing.T) {
	for _, tc := range []struct {
		name string
		ip   interface{}
	}{
		{name: "omitted"},
		{name: "ipv4", ip: "192.0.2.10"},
		{name: "ipv6", ip: "2001:db8::10"},
		// The legacy SDK normalizes this Optional+Computed zero value to omission.
		{name: "empty_is_omitted", ip: ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			source := map[string]interface{}{
				"endpoint": "https://source.example.com:9200", "ca_bundle_id": "test-ca-bundle",
			}
			if tc.ip != nil {
				source["endpoint_resolution_ip"] = tc.ip
			}
			data := schema.TestResourceDataRaw(t, OpensearchOpensearchClusterResource().Schema, map[string]interface{}{
				"remote_reindex_trust_config": []interface{}{map[string]interface{}{
					"remote_sources": []interface{}{source},
				}},
			})
			crud := OpensearchOpensearchClusterResourceCrud{BaseCrud: tfresource.BaseCrud{D: data}}
			trust, err := crud.mapToRemoteReindexTrustConfig("remote_reindex_trust_config.0.%s")
			if err != nil {
				t.Fatal(err)
			}
			if len(trust.RemoteSources) != 1 {
				t.Fatalf("expected one mapped source, got %d", len(trust.RemoteSources))
			}
			got := trust.RemoteSources[0]
			if got.Endpoint == nil || *got.Endpoint != source["endpoint"] || got.CaBundleId == nil || *got.CaBundleId != source["ca_bundle_id"] {
				t.Fatalf("required remote source fields were not preserved: %+v", got)
			}
			if tc.ip == nil || tc.ip == "" {
				if got.EndpointResolutionIp != nil {
					t.Fatalf("omitted IP must map to nil, got %q", *got.EndpointResolutionIp)
				}
			} else if got.EndpointResolutionIp == nil || *got.EndpointResolutionIp != tc.ip {
				t.Fatalf("resolution IP was not preserved: %+v", got)
			}
		})
	}
}

func TestOpensearchRemoteReindexUpdateRequest(t *testing.T) {
	const prefix = "remote_reindex_trust_config.0.remote_sources."
	for _, tc := range []struct {
		name string
		diff map[string]*terraform.ResourceAttrDiff
		want string // Empty means the JSON property must be absent, not null.
	}{
		{
			name: "trust_only_update",
			diff: map[string]*terraform.ResourceAttrDiff{
				prefix + "0.endpoint": {Old: "https://source-a.example.com:9200", New: "https://source-b.example.com:9200"},
			},
			want: `{"remoteSources":[{"endpoint":"https://source-b.example.com:9200","caBundleId":"test-ca","endpointResolutionIp":"192.0.2.10"}]}`,
		},
		{
			name: "empty_sources_removes_trust",
			diff: map[string]*terraform.ResourceAttrDiff{
				prefix + "#":                        {Old: "1", New: "0"},
				prefix + "0.endpoint":               {Old: "https://source-a.example.com:9200", NewRemoved: true},
				prefix + "0.endpoint_resolution_ip": {Old: "192.0.2.10", NewRemoved: true},
				prefix + "0.ca_bundle_id":           {Old: "test-ca", NewRemoved: true},
			},
			// This must be an explicit empty array, not a missing property or null.
			want: `{"remoteSources":[]}`,
		},
		{
			name: "unrelated_update_omits_unchanged_trust",
			diff: map[string]*terraform.ResourceAttrDiff{
				"display_name": {Old: "test-cluster", New: "renamed-cluster"},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			data, err := schema.InternalMap(OpensearchOpensearchClusterResource().Schema).Data(&terraform.InstanceState{
				ID: "test-cluster",
				Attributes: map[string]string{
					"display_name":                      "test-cluster",
					"remote_reindex_trust_config.#":     "1",
					prefix + "#":                        "1",
					prefix + "0.endpoint":               "https://source-a.example.com:9200",
					prefix + "0.endpoint_resolution_ip": "192.0.2.10",
					prefix + "0.ca_bundle_id":           "test-ca",
				},
			}, &terraform.InstanceDiff{Attributes: tc.diff})
			if err != nil {
				t.Fatal(err)
			}
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			var payload []byte
			client := &oci_opensearch.OpensearchClusterClient{BaseClient: common.BaseClient{
				Host: "https://example.invalid", UserAgent: "remote-reindex-unit-test",
				// Capture the actual SDK-serialized request, then stop before
				// signing or network I/O. No credentials or OCI access are used.
				Interceptor: func(request *http.Request) error {
					if request.Method != http.MethodPut {
						t.Fatalf("unexpected request method: %s", request.Method)
					}
					payload, err = io.ReadAll(request.Body)
					if err != nil {
						t.Fatal(err)
					}
					cancel()
					return context.Canceled
				},
			}}
			crud := OpensearchOpensearchClusterResourceCrud{BaseCrud: tfresource.BaseCrud{D: data}, Client: client}
			if err := crud.UpdateWithContext(ctx); err == nil {
				t.Fatal("expected the request interceptor to stop the update")
			}
			if len(payload) == 0 {
				t.Fatal("the provider did not build an update request")
			}
			var body map[string]json.RawMessage
			if err := json.Unmarshal(payload, &body); err != nil {
				t.Fatal(err)
			}
			actual, exists := body["remoteReindexTrustConfig"]
			if tc.want == "" {
				if exists {
					t.Fatalf("unchanged trust must be omitted, got %s", actual)
				}
				return
			}
			var got, want interface{}
			if err := json.Unmarshal(actual, &got); err != nil {
				t.Fatalf("missing or invalid remote-reindex payload: %s: %v", actual, err)
			}
			if err := json.Unmarshal([]byte(tc.want), &want); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("remote-reindex JSON = %s, want %s", actual, tc.want)
			}
		})
	}
}

func TestOpensearchRemoteReindexReadBack(t *testing.T) {
	emptySources := []interface{}{}
	populatedSources := []interface{}{map[string]interface{}{
		"endpoint": "https://source.example.com:9200", "ca_bundle_id": "test-ca",
		"certificate_id":         "",
		"endpoint_resolution_ip": "192.0.2.10",
	}}
	apiTrust := &oci_opensearch.RemoteReindexTrustConfig{
		RemoteSources: []oci_opensearch.RemoteReindexSourceDetails{{
			Endpoint: common.String("https://source.example.com:9200"), CaBundleId: common.String("test-ca"),
			EndpointResolutionIp: common.String("192.0.2.10"),
		}},
	}
	for _, tc := range []struct {
		name        string
		prior       []interface{} // nil means the parent block was absent.
		api         *oci_opensearch.RemoteReindexTrustConfig
		wantBlocks  int
		wantSources []interface{}
	}{
		{name: "absent_stays_absent"},
		{name: "cleared_stays_empty", prior: emptySources, wantBlocks: 1, wantSources: emptySources},
		{name: "external_removal_is_not_masked", prior: populatedSources},
		{name: "import_populated_trust", api: apiTrust, wantBlocks: 1, wantSources: populatedSources},
		{name: "external_addition_is_not_masked", prior: emptySources, api: apiTrust, wantBlocks: 1, wantSources: populatedSources},
		{name: "populated_trust_is_preserved", prior: populatedSources, api: apiTrust, wantBlocks: 1, wantSources: populatedSources},
		{name: "empty_api_object", prior: populatedSources, api: &oci_opensearch.RemoteReindexTrustConfig{}, wantBlocks: 1, wantSources: emptySources},
	} {
		t.Run(tc.name, func(t *testing.T) {
			raw := map[string]interface{}{}
			if tc.prior != nil {
				raw["remote_reindex_trust_config"] = []interface{}{map[string]interface{}{"remote_sources": tc.prior}}
			}
			resourceSchema := OpensearchOpensearchClusterResource().Schema
			data := schema.TestResourceDataRaw(t, resourceSchema, raw)
			data.SetId("offline-test-cluster")
			// Check apply-time read-back, then a refresh from saved state without
			// raw configuration. The representation must remain stable in both.
			for read := 0; read < 2; read++ {
				crud := OpensearchOpensearchClusterResourceCrud{
					BaseCrud: tfresource.BaseCrud{D: data},
					Res:      &oci_opensearch.OpensearchCluster{RemoteReindexTrustConfig: tc.api},
				}
				if err := crud.SetData(); err != nil {
					t.Fatal(err)
				}
				if got := data.Get("remote_reindex_trust_config.#"); got != tc.wantBlocks {
					t.Fatalf("read %d: parent count = %v, want %d", read, got, tc.wantBlocks)
				}
				if tc.wantBlocks > 0 {
					if got := data.Get("remote_reindex_trust_config.0.remote_sources"); !reflect.DeepEqual(got, tc.wantSources) {
						t.Fatalf("read %d: sources = %#v, want %#v", read, got, tc.wantSources)
					}
				}
				var err error
				data, err = schema.InternalMap(resourceSchema).Data(data.State(), nil)
				if err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}

// Use the production trust schema, mapper, and read-back with an in-memory API.
// The API returns nil trust after clearing, as the live service does. These tests
// exercise real Terraform parsing/planning, but never contact OCI.
func testOpensearchRemoteReindexTerraformProvider() *schema.Provider {
	var apiTrust *oci_opensearch.RemoteReindexTrustConfig
	createCount := 0
	read := func(_ context.Context, d *schema.ResourceData, _ interface{}) diag.Diagnostics {
		crud := OpensearchOpensearchClusterResourceCrud{
			BaseCrud: tfresource.BaseCrud{D: d},
			Res:      &oci_opensearch.OpensearchCluster{RemoteReindexTrustConfig: apiTrust},
		}
		return diag.FromErr(crud.setRemoteReindexTrustConfig())
	}
	apply := func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
		if d.HasChange("remote_reindex_trust_config") && len(d.Get("remote_reindex_trust_config").([]interface{})) > 0 {
			crud := OpensearchOpensearchClusterResourceCrud{BaseCrud: tfresource.BaseCrud{D: d}}
			trust, err := crud.mapToRemoteReindexTrustConfig("remote_reindex_trust_config.0.%s")
			if err != nil {
				return diag.FromErr(err)
			}
			apiTrust = nil
			if len(trust.RemoteSources) > 0 {
				apiTrust = &trust
			}
		}
		return read(ctx, d, meta)
	}
	testResource := &schema.Resource{
		CustomizeDiff: validateOpensearchRemoteReindexAliases,
		Schema: map[string]*schema.Schema{
			"remote_reindex_trust_config": OpensearchOpensearchClusterResource().Schema["remote_reindex_trust_config"],
		},
		CreateContext: func(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
			createCount++
			d.SetId("offline-test-cluster")
			if createCount > 1 {
				d.SetId(fmt.Sprintf("offline-test-cluster-%d", createCount))
			}
			return apply(ctx, d, meta)
		},
		ReadContext: read, UpdateContext: apply,
		DeleteContext: func(_ context.Context, d *schema.ResourceData, _ interface{}) diag.Diagnostics {
			apiTrust = nil
			d.SetId("")
			return nil
		},
	}
	return &schema.Provider{ResourcesMap: map[string]*schema.Resource{
		"test_remote_reindex": testResource,
		// Provides an unknown CA bundle ID during planning without contacting OCI.
		"test_ca_bundle": {
			Schema: map[string]*schema.Schema{"value": {Type: schema.TypeString, Computed: true}},
			CreateContext: func(_ context.Context, d *schema.ResourceData, _ interface{}) diag.Diagnostics {
				d.SetId("offline-ca-bundle")
				return diag.FromErr(d.Set("value", "test-ca"))
			},
			ReadContext: func(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics { return nil },
			DeleteContext: func(_ context.Context, d *schema.ResourceData, _ interface{}) diag.Diagnostics {
				d.SetId("")
				return nil
			},
		},
	}}
}

func TestOpensearchRemoteReindexAliasMapping(t *testing.T) {
	for _, tc := range []struct {
		name, legacy, current string
		wantError             bool
	}{
		{name: "legacy", legacy: "test-ca"},
		{name: "current", current: "test-ca"},
		{name: "both", legacy: "test-ca", current: "test-ca", wantError: true},
		{name: "neither", wantError: true},
		{name: "blank", legacy: " ", wantError: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			data := schema.TestResourceDataRaw(t, OpensearchOpensearchClusterResource().Schema, map[string]interface{}{
				"remote_reindex_trust_config": []interface{}{map[string]interface{}{
					"remote_sources": []interface{}{map[string]interface{}{
						"endpoint": "https://source.example.com:9200", "certificate_id": tc.legacy, "ca_bundle_id": tc.current,
					}},
				}},
			})
			crud := OpensearchOpensearchClusterResourceCrud{BaseCrud: tfresource.BaseCrud{D: data}}
			trust, err := crud.mapToRemoteReindexTrustConfig("remote_reindex_trust_config.0.%s")
			if tc.wantError {
				if err == nil || !strings.Contains(err.Error(), "exactly one") {
					t.Fatalf("expected alias validation error, got %v", err)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			body, err := json.Marshal(trust)
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(body), `"caBundleId":"test-ca"`) || strings.Contains(string(body), "certificateId") {
				t.Fatalf("both aliases must serialize as caBundleId: %s", body)
			}
		})
	}
}

func TestOpensearchRemoteReindexAliasDeprecation(t *testing.T) {
	p := testOpensearchRemoteReindexTerraformProvider()
	for _, attribute := range []string{"certificate_id", "ca_bundle_id"} {
		t.Run(attribute, func(t *testing.T) {
			diagnostics := p.ValidateResource("test_remote_reindex", terraform.NewResourceConfigRaw(map[string]interface{}{
				"remote_reindex_trust_config": []interface{}{map[string]interface{}{
					"remote_sources": []interface{}{map[string]interface{}{
						"endpoint": "https://source.example.com", attribute: "test-ca",
					}},
				}},
			}))
			if diagnostics.HasError() {
				t.Fatalf("alias configuration rejected: %v", diagnostics)
			}
			warned := false
			for _, diagnostic := range diagnostics {
				if diagnostic.Severity == diag.Warning && strings.Contains(diagnostic.Detail, "Use ca_bundle_id instead") {
					warned = true
				}
			}
			if warned != (attribute == "certificate_id") {
				t.Fatalf("deprecation warning for %s = %t; diagnostics: %v", attribute, warned, diagnostics)
			}
		})
	}
}

func TestOpensearchRemoteReindexAliasReadBack(t *testing.T) {
	const prefix = "remote_reindex_trust_config.0.remote_sources."
	sources := []interface{}{
		map[string]interface{}{"endpoint": "https://old.example.com", "certificate_id": "old-ca"},
		map[string]interface{}{"endpoint": "https://new.example.com", "ca_bundle_id": "new-ca"},
	}
	resourceSchema := OpensearchOpensearchClusterResource().Schema
	data := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{
		"remote_reindex_trust_config": []interface{}{map[string]interface{}{"remote_sources": sources}},
	})
	data.SetId("offline-cluster")
	trust := &oci_opensearch.RemoteReindexTrustConfig{RemoteSources: []oci_opensearch.RemoteReindexSourceDetails{
		{Endpoint: common.String("https://old.example.com"), CaBundleId: common.String("rotated-old-ca")},
		{Endpoint: common.String("https://new.example.com"), CaBundleId: common.String("rotated-new-ca")},
	}}
	for read := 0; read < 2; read++ {
		crud := OpensearchOpensearchClusterResourceCrud{BaseCrud: tfresource.BaseCrud{D: data}, Res: &oci_opensearch.OpensearchCluster{RemoteReindexTrustConfig: trust}}
		if err := crud.setRemoteReindexTrustConfig(); err != nil {
			t.Fatal(err)
		}
		for field, want := range map[string]string{
			"0.certificate_id": "rotated-old-ca", "0.ca_bundle_id": "",
			"1.certificate_id": "", "1.ca_bundle_id": "rotated-new-ca",
		} {
			if got := data.Get(prefix + field); got != want {
				t.Fatalf("read %d: %s = %v, want %s", read, field, got, want)
			}
		}
		var err error
		data, err = schema.InternalMap(resourceSchema).Data(data.State(), nil)
		if err != nil {
			t.Fatal(err)
		}
	}
	// Computed data source outputs keep existing references to either name working.
	flat := RemoteReindexTrustConfigToMap(trust)["remote_sources"].([]interface{})
	for i, source := range flat {
		values := source.(map[string]interface{})
		if values["certificate_id"] != *trust.RemoteSources[i].CaBundleId || values["ca_bundle_id"] != values["certificate_id"] {
			t.Fatalf("data source aliases differ: %#v", values)
		}
	}
}

func TestOpensearchRemoteReindexTerraformAliasMigration(t *testing.T) {
	const name = "test_remote_reindex.cluster"
	const prefix = "remote_reindex_trust_config.0.remote_sources.0."
	config := func(attribute string) string {
		return fmt.Sprintf(`resource "test_remote_reindex" "cluster" {
  remote_reindex_trust_config {
    remote_sources {
      endpoint = "https://source.example.com"
      %s = "test-ca"
    }
  }
}`, attribute)
	}
	resource.UnitTest(t, resource.TestCase{
		Providers: map[string]*schema.Provider{"test": testOpensearchRemoteReindexTerraformProvider()},
		Steps: []resource.TestStep{
			{Config: config("certificate_id"), Check: resource.TestCheckResourceAttr(name, prefix+"certificate_id", "test-ca")},
			{Config: config("certificate_id"), PlanOnly: true, ExpectNonEmptyPlan: false},
			{Config: `resource "test_remote_reindex" "cluster" {
  remote_reindex_trust_config { remote_sources = [] }
}`, Check: resource.TestCheckResourceAttr(name, "remote_reindex_trust_config.0.remote_sources.#", "0")},
			{Config: config("certificate_id"), Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(name, prefix+"certificate_id", "test-ca"),
				resource.TestCheckResourceAttr(name, "id", "offline-test-cluster"),
			)},
			{Config: config("ca_bundle_id"), Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(name, prefix+"ca_bundle_id", "test-ca"),
				resource.TestCheckResourceAttr(name, prefix+"certificate_id", ""),
				resource.TestCheckResourceAttr(name, "id", "offline-test-cluster"),
			)},
			{Config: config("ca_bundle_id"), PlanOnly: true, ExpectNonEmptyPlan: false},
			{Config: config("certificate_id"), Check: resource.TestCheckResourceAttr(name, prefix+"certificate_id", "test-ca")},
			{Config: config("certificate_id"), PlanOnly: true, ExpectNonEmptyPlan: false},
		},
	})
}

func TestOpensearchRemoteReindexTerraformAliasValidation(t *testing.T) {
	for _, tc := range []struct {
		name, fields string
	}{
		{name: "both", fields: "certificate_id = \"ca\"\nca_bundle_id = \"ca\""},
		{name: "neither"},
		{name: "empty", fields: "certificate_id = \"\""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			resource.UnitTest(t, resource.TestCase{
				Providers: map[string]*schema.Provider{"test": testOpensearchRemoteReindexTerraformProvider()},
				Steps: []resource.TestStep{{Config: fmt.Sprintf(`resource "test_remote_reindex" "cluster" {
  remote_reindex_trust_config {
    remote_sources {
      endpoint = "https://source.example.com"
      %s
    }
  }
}`, tc.fields), PlanOnly: true, ExpectError: regexp.MustCompile("exactly one of certificate_id or ca_bundle_id")}},
			})
		})
	}
}

func TestOpensearchRemoteReindexTerraformMixedAliases(t *testing.T) {
	const config = `resource "test_ca_bundle" "bundle" {}
resource "test_remote_reindex" "cluster" {
  remote_reindex_trust_config {
    remote_sources {
      endpoint = "https://legacy.example.com"
      certificate_id = test_ca_bundle.bundle.value
    }
    remote_sources {
      endpoint = "https://current.example.com"
      ca_bundle_id = test_ca_bundle.bundle.value
    }
  }
}`
	resource.UnitTest(t, resource.TestCase{
		Providers: map[string]*schema.Provider{"test": testOpensearchRemoteReindexTerraformProvider()},
		Steps: []resource.TestStep{
			{Config: config, Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("test_remote_reindex.cluster", "remote_reindex_trust_config.0.remote_sources.0.certificate_id", "test-ca"),
				resource.TestCheckResourceAttr("test_remote_reindex.cluster", "remote_reindex_trust_config.0.remote_sources.1.ca_bundle_id", "test-ca"),
			)},
			{Config: config, PlanOnly: true, ExpectNonEmptyPlan: false},
		},
	})
}

func TestOpensearchRemoteReindexTerraformPreservation(t *testing.T) {
	const name = "test_remote_reindex.cluster"
	const configured = `resource "test_remote_reindex" "cluster" {
  remote_reindex_trust_config {
    remote_sources {
      endpoint = "https://source.example.com:9200"
      ca_bundle_id = "test-ca"
    }
  }
}`
	resource.UnitTest(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"test": testOpensearchRemoteReindexTerraformProvider(),
		},
		Steps: []resource.TestStep{
			{Config: configured, Check: resource.TestCheckResourceAttr(name, "remote_reindex_trust_config.#", "1")},
			{Config: `resource "test_remote_reindex" "cluster" {}`, Check: resource.TestCheckResourceAttr(name, "remote_reindex_trust_config.#", "1")},
			{Config: configured, Check: resource.TestCheckResourceAttr(name, "remote_reindex_trust_config.#", "1")},
		},
	})
}

// Catch schema rejection and read-back drift for explicit [] without provisioning
// a live cluster. The acceptance test independently checks the real API response.
func TestOpensearchRemoteReindexTerraformRemoval(t *testing.T) {
	const name = "test_remote_reindex.cluster"
	const sourceCount = "remote_reindex_trust_config.0.remote_sources.#"
	const configured = `resource "test_remote_reindex" "cluster" {
  remote_reindex_trust_config {
    remote_sources {
      endpoint = "https://source.example.com:9200"
      ca_bundle_id = "test-ca"
    }
  }
}`
	const cleared = `resource "test_remote_reindex" "cluster" {
  remote_reindex_trust_config {
    remote_sources = []
  }
}`
	resource.UnitTest(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"test": testOpensearchRemoteReindexTerraformProvider(),
		},
		Steps: []resource.TestStep{
			{Config: configured, Check: resource.TestCheckResourceAttr(name, sourceCount, "1")},
			{Config: cleared, Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(name, "id", "offline-test-cluster"),
				resource.TestCheckResourceAttr(name, "remote_reindex_trust_config.#", "1"),
				resource.TestCheckResourceAttr(name, sourceCount, "0"),
			)},
			{Config: cleared, PlanOnly: true, ExpectNonEmptyPlan: false},
			{Config: `resource "test_remote_reindex" "cluster" {}`, Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(name, "remote_reindex_trust_config.#", "1"),
				resource.TestCheckResourceAttr(name, sourceCount, "0"),
			)},
			{Config: cleared, PlanOnly: true, ExpectNonEmptyPlan: false},
			{Config: configured, Check: resource.TestCheckResourceAttr(name, sourceCount, "1")},
			{Config: configured, PlanOnly: true, ExpectNonEmptyPlan: false},
		},
	})
}

func TestOpensearchRemoteReindexTerraformEmptyTrust(t *testing.T) {
	const name = "test_remote_reindex.cluster"
	for _, tc := range []struct {
		name       string
		config     string
		wantBlocks string
	}{
		{name: "omitted", config: `resource "test_remote_reindex" "cluster" {}`, wantBlocks: "0"},
		{name: "explicit_empty", config: `resource "test_remote_reindex" "cluster" {
  remote_reindex_trust_config {
    remote_sources = []
  }
}`, wantBlocks: "1"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			resource.UnitTest(t, resource.TestCase{
				Providers: map[string]*schema.Provider{
					"test": testOpensearchRemoteReindexTerraformProvider(),
				},
				Steps: []resource.TestStep{
					{Config: tc.config, Check: resource.TestCheckResourceAttr(name, "remote_reindex_trust_config.#", tc.wantBlocks)},
					{Config: tc.config, PlanOnly: true, ExpectNonEmptyPlan: false},
				},
			})
		})
	}
}
