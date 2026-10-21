// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package golden_gate

import (
	"testing"

	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"
)

func TestDeploymentCertificateKeyFromWorkRequestResources(t *testing.T) {
	tests := []struct {
		name      string
		resources []oci_golden_gate.WorkRequestResource
		wantKey   string
	}{
		{
			name: "certificate entity",
			resources: []oci_golden_gate.WorkRequestResource{
				workRequestResource("certificate", "EltostratusRSARoot1", oci_golden_gate.ActionTypeCreated),
			},
			wantKey: "EltostratusRSARoot1",
		},
		{
			name: "deployment certificate entity does not become the deployment ID",
			resources: []oci_golden_gate.WorkRequestResource{
				workRequestResource("deploymentCertificate", "EltostratusRSARoot1", oci_golden_gate.ActionTypeCreated),
				workRequestResource("deployment", "ocid1.goldengatedeployment.oc1..example", oci_golden_gate.ActionTypeCreated),
			},
			wantKey: "EltostratusRSARoot1",
		},
		{
			name: "ignores deployment-only resources",
			resources: []oci_golden_gate.WorkRequestResource{
				workRequestResource("deployment", "ocid1.goldengatedeployment.oc1..example", oci_golden_gate.ActionTypeCreated),
			},
		},
		{
			name: "ignores resources for another action",
			resources: []oci_golden_gate.WorkRequestResource{
				workRequestResource("certificate", "EltostratusRSARoot1", oci_golden_gate.ActionTypeDeleted),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := deploymentCertificateKeyFromWorkRequestResources(test.resources, oci_golden_gate.ActionTypeCreated)
			if test.wantKey == "" {
				if got != nil {
					t.Fatalf("expected no certificate key, got %q", *got)
				}
				return
			}
			if got == nil {
				t.Fatalf("expected certificate key %q, got nil", test.wantKey)
			}
			if *got != test.wantKey {
				t.Fatalf("expected certificate key %q, got %q", test.wantKey, *got)
			}
		})
	}
}

func workRequestResource(entityType, identifier string, action oci_golden_gate.ActionTypeEnum) oci_golden_gate.WorkRequestResource {
	return oci_golden_gate.WorkRequestResource{
		EntityType: &entityType,
		Identifier: &identifier,
		ActionType: action,
	}
}
