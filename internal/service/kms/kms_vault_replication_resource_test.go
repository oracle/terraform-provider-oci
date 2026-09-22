// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License Version 2.0

package kms

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestUnitParseVaultReplicationId(t *testing.T) {
	tests := map[string]struct {
		id            string
		wantVaultId   string
		wantRegion    string
		wantErrSubstr string
	}{
		"valid": {
			id:          "ocid1.vault.oc1.iad.example:us-phoenix-1",
			wantVaultId: "ocid1.vault.oc1.iad.example",
			wantRegion:  "us-phoenix-1",
		},
		"missing separator": {
			id:            "bad-id",
			wantErrSubstr: "expected format",
		},
		"empty vault ID": {
			id:            ":us-phoenix-1",
			wantErrSubstr: "expected format",
		},
		"empty replica region": {
			id:            "ocid1.vault.oc1.iad.example:",
			wantErrSubstr: "expected format",
		},
		"additional separator": {
			id:            "ocid1.vault.oc1.iad.example:us-phoenix-1:extra",
			wantErrSubstr: "expected format",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			vaultId, region, err := parseVaultReplicationId(tt.id)
			if tt.wantErrSubstr != "" {
				if err == nil || !strings.Contains(err.Error(), tt.wantErrSubstr) {
					t.Fatalf("parseVaultReplicationId(%q) error = %v, want error containing %q", tt.id, err, tt.wantErrSubstr)
				}
				return
			}
			if err != nil {
				t.Fatalf("parseVaultReplicationId(%q) unexpected error: %v", tt.id, err)
			}
			if vaultId != tt.wantVaultId || region != tt.wantRegion {
				t.Fatalf("parseVaultReplicationId(%q) = (%q, %q), want (%q, %q)", tt.id, vaultId, region, tt.wantVaultId, tt.wantRegion)
			}
		})
	}
}

func TestUnitKmsVaultReplicationGetMalformedIdDoesNotExit(t *testing.T) {
	const childEnv = "TF_OCI_TEST_VAULT_REPLICATION_MALFORMED_ID"
	if os.Getenv(childEnv) == "1" {
		d := schema.TestResourceDataRaw(t, KmsVaultReplicationResource().Schema, map[string]any{})
		d.SetId("bad-id")
		crud := &KmsVaultReplicaResourceCrud{}
		crud.D = d
		err := crud.Get()
		if err == nil || !strings.Contains(err.Error(), "expected format") {
			t.Fatalf("Get() error = %v, want malformed ID error", err)
		}
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=^TestUnitKmsVaultReplicationGetMalformedIdDoesNotExit$")
	cmd.Env = append(os.Environ(), childEnv+"=1")
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("Get() terminated its host process: %v\n%s", err, output)
	}
}
