// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
)

// These tests are offline: no TF_ACC, Terraform process, or OCI credentials needed.
func TestUnitOpensearchRemoteReindexConfig(t *testing.T) {
	source := opensearchRemoteReindexSource{
		caBundleResource: opensearchRemoteReindexCaResource,
		endpoint:         "https://source.example.com:9200",
		resolutionIP:     common.String("192.0.2.10"),
	}
	withoutIP := source
	withoutIP.resolutionIP = nil
	other := source
	other.endpoint = "https://other.example.com:9200"
	other.caBundleResource = opensearchRemoteReindexOtherCaResource
	other.resolutionIP = common.String("2001:db8::10")
	for _, tc := range []struct {
		name    string
		options opensearchRemoteReindexConfigOptions
	}{
		{name: "omitted"},
		{name: "empty_sources_omitted", options: opensearchRemoteReindexConfigOptions{sources: []opensearchRemoteReindexSource{}}},
		{name: "explicit_empty_sources", options: opensearchRemoteReindexConfigOptions{clearSources: true}},
		{name: "one_source", options: opensearchRemoteReindexConfigOptions{sources: []opensearchRemoteReindexSource{source}}},
		{name: "optional_ip_omitted", options: opensearchRemoteReindexConfigOptions{sources: []opensearchRemoteReindexSource{withoutIP}}},
		{name: "multiple_bundles", options: opensearchRemoteReindexConfigOptions{sources: []opensearchRemoteReindexSource{source, other}, replacementCA: true}},
		{name: "resize", options: opensearchRemoteReindexConfigOptions{sources: []opensearchRemoteReindexSource{source}, tagged: true, dataNodeCount: 2, dataNodeOCPUs: 4, dataNodeMemoryGB: 40}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			config := testAccOpensearchRemoteReindexConfig("test-compartment", "test-cluster", tc.options)
			file, diags := hclsyntax.ParseConfig([]byte(config), "remote-reindex.tf", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatal(diags.Error())
			}
			resources := map[string]*hclsyntax.Block{}
			for _, block := range file.Body.(*hclsyntax.Body).Blocks {
				if block.Type == "resource" {
					resources[block.Labels[0]+"."+block.Labels[1]] = block
				}
			}
			wantResources := 4
			if tc.options.replacementCA {
				wantResources++
			}
			if len(resources) != wantResources {
				t.Fatalf("expected %d resources, got %d", wantResources, len(resources))
			}
			cluster := resources[opensearchRemoteReindexClusterResource]
			if cluster == nil || cluster.Body.Attributes["depends_on"] == nil {
				t.Fatal("cluster must keep explicit CA dependencies for cleanup ordering")
			}
			var trust *hclsyntax.Block
			for _, block := range cluster.Body.Blocks {
				if block.Type == "remote_reindex_trust_config" {
					trust = block
				}
			}
			if tc.options.clearSources {
				if trust == nil || len(trust.Body.Blocks) != 0 {
					t.Fatal("clear must retain the trust block without any remote source blocks")
				}
				sources := trust.Body.Attributes["remote_sources"]
				if sources == nil {
					t.Fatal("clear must explicitly set remote_sources = []")
				}
				value, diags := sources.Expr.Value(nil)
				if diags.HasErrors() || value.IsNull() || !value.IsKnown() || !value.Type().IsTupleType() || value.LengthInt() != 0 {
					t.Fatal("clear must use an empty array, not null or omission")
				}
				return
			}
			if len(tc.options.sources) == 0 {
				if trust != nil || cluster.Body.Attributes["remote_reindex_trust_config"] != nil {
					t.Fatal("no sources must omit trust, not clear it")
				}
				return
			}
			if trust == nil || len(trust.Body.Blocks) != len(tc.options.sources) {
				t.Fatal("wrong number of remote source blocks")
			}
			for i, want := range tc.options.sources {
				block := trust.Body.Blocks[i]
				caBundle := block.Body.Attributes["ca_bundle_id"]
				if caBundle == nil || block.Body.Attributes["certificate_id"] != nil {
					t.Fatal("remote sources must use ca_bundle_id, not the old certificate_id field")
				}
				bundleRange := caBundle.Expr.Range()
				if got := config[bundleRange.Start.Byte:bundleRange.End.Byte]; got != want.caBundleResource+".id" {
					t.Fatalf("CA bundle reference = %q, want %q", got, want.caBundleResource+".id")
				}
				endpoint, diags := block.Body.Attributes["endpoint"].Expr.Value(nil)
				if diags.HasErrors() || endpoint.AsString() != want.endpoint {
					t.Fatal("endpoint was not preserved")
				}
				ip := block.Body.Attributes["endpoint_resolution_ip"]
				if want.resolutionIP == nil {
					if ip != nil {
						t.Fatal("an omitted resolution IP must not render as an empty string")
					}
				} else {
					if ip == nil {
						t.Fatal("resolution IP is missing")
					}
					value, diags := ip.Expr.Value(nil)
					if diags.HasErrors() || value.AsString() != *want.resolutionIP {
						t.Fatal("resolution IP was not preserved")
					}
				}
			}
		})
	}
}

func TestUnitOpensearchRemoteReindexClearedState(t *testing.T) {
	const trustCount = "remote_reindex_trust_config.#"
	const sourceCount = "remote_reindex_trust_config.0.remote_sources.#"
	const oldCA = "remote_reindex_trust_config.0.remote_sources.0.ca_bundle_id"
	for _, tc := range []struct {
		name    string
		attrs   map[string]string
		wantErr bool
	}{
		{name: "absent_trust", attrs: map[string]string{trustCount: "0"}},
		{name: "empty_sources", attrs: map[string]string{trustCount: "1", sourceCount: "0"}},
		{name: "remaining_source", attrs: map[string]string{trustCount: "1", sourceCount: "1"}, wantErr: true},
		{name: "stale_ca", attrs: map[string]string{trustCount: "1", sourceCount: "0", oldCA: "old-ca"}, wantErr: true},
		{name: "absent_trust_stale_sources", attrs: map[string]string{trustCount: "0", sourceCount: "1"}, wantErr: true},
		{name: "missing_source_count", attrs: map[string]string{trustCount: "1"}, wantErr: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tc.attrs["state"] = "ACTIVE"
			state := terraform.NewState()
			state.RootModule().Resources[opensearchRemoteReindexClusterResource] = &terraform.ResourceState{
				Type: "oci_opensearch_opensearch_cluster",
				Primary: &terraform.InstanceState{
					ID: "test-cluster", Attributes: tc.attrs,
				},
			}
			err := testAccCheckOpensearchRemoteReindexClearedState(opensearchRemoteReindexClusterResource)(state)
			if (err != nil) != tc.wantErr {
				t.Fatalf("cleared-state check error = %v, want error = %t", err, tc.wantErr)
			}
		})
	}
}

func TestUnitOpensearchRemoteReindexNodeHostShape(t *testing.T) {
	for _, tc := range []struct {
		name  string
		shape string
		want  string
	}{
		{name: "default", want: "VM.Standard.E3.Flex"},
		{name: "dev_mumbai", shape: "InternalShapePriorityList", want: "InternalShapePriorityList"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("TF_VAR_opensearch_node_host_shape", tc.shape)
			t.Setenv("OCI_opensearch_node_host_shape", "")
			t.Setenv("opensearch_node_host_shape", "")
			config := testAccOpensearchRemoteReindexConfig("test-compartment", "test-cluster", opensearchRemoteReindexConfigOptions{})
			file, diags := hclsyntax.ParseConfig([]byte(config), "remote-reindex.tf", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatal(diags.Error())
			}
			for _, block := range file.Body.(*hclsyntax.Body).Blocks {
				if block.Type != "resource" || block.Labels[0] != "oci_opensearch_opensearch_cluster" {
					continue
				}
				for _, name := range []string{"master_node_host_shape", "data_node_host_shape", "opendashboard_node_host_shape"} {
					attribute := block.Body.Attributes[name]
					if attribute == nil {
						t.Fatalf("missing %s", name)
					}
					value, diags := attribute.Expr.Value(nil)
					if diags.HasErrors() || value.AsString() != tc.want {
						t.Fatalf("%s must use shape %q", name, tc.want)
					}
				}
				return
			}
			t.Fatal("cluster resource is missing")
		})
	}
}

func TestUnitOpensearchRemoteReindexCAFixtures(t *testing.T) {
	var certificates []*x509.Certificate
	for _, fixture := range []string{opensearchRemoteReindexTestCaPEM, opensearchRemoteReindexReplacementCaPEM} {
		block, rest := pem.Decode([]byte(fixture))
		if block == nil || block.Type != "CERTIFICATE" || len(bytes.TrimSpace(rest)) != 0 {
			t.Fatal("each fixture must contain exactly one public certificate")
		}
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			t.Fatal(err)
		}
		if !cert.IsCA || !cert.BasicConstraintsValid || cert.KeyUsage&x509.KeyUsageCertSign == 0 {
			t.Fatal("fixture must be a valid signing CA")
		}
		if err := cert.CheckSignatureFrom(cert); err != nil {
			t.Fatal(err)
		}
		certificates = append(certificates, cert)
	}
	if bytes.Equal(certificates[0].Raw, certificates[1].Raw) {
		t.Fatal("replacement CA must have different certificate content")
	}
}
