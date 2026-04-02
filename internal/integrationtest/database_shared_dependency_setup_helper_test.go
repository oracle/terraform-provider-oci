// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import "testing"

func TestTerraformResourceIDFromState(t *testing.T) {
	t.Run("finds resource id in root module", func(t *testing.T) {
		state := terraformStateFile{
			Resources: []terraformStateResource{
				{
					Type: "oci_database_autonomous_container_database",
					Name: "test_autonomous_container_database",
					Instances: []terraformStateInstance{
						{
							Attributes: map[string]any{
								"id": "ocid1.autonomouscontainerdatabase.oc1..example",
							},
						},
					},
				},
			},
		}

		resourceID, found := terraformResourceIDFromState(state, "oci_database_autonomous_container_database.test_autonomous_container_database")
		if !found {
			t.Fatal("expected resource to be found")
		}
		if resourceID != "ocid1.autonomouscontainerdatabase.oc1..example" {
			t.Fatalf("unexpected resource id: %s", resourceID)
		}
	})

	t.Run("finds resource id in module-qualified resource", func(t *testing.T) {
		state := terraformStateFile{
			Resources: []terraformStateResource{
				{
					Module: "module.shared",
					Type:   "oci_database_autonomous_container_database",
					Name:   "test_autonomous_container_database",
					Instances: []terraformStateInstance{
						{
							Attributes: map[string]any{
								"id": "ocid1.autonomouscontainerdatabase.oc1..child",
							},
						},
					},
				},
			},
		}

		resourceID, found := terraformResourceIDFromState(state, "module.shared.oci_database_autonomous_container_database.test_autonomous_container_database")
		if !found {
			t.Fatal("expected module-qualified resource to be found")
		}
		if resourceID != "ocid1.autonomouscontainerdatabase.oc1..child" {
			t.Fatalf("unexpected module-qualified resource id: %s", resourceID)
		}
	})

	t.Run("returns empty id when id attribute is missing", func(t *testing.T) {
		state := terraformStateFile{
			Resources: []terraformStateResource{
				{
					Type: "oci_database_autonomous_container_database",
					Name: "test_autonomous_container_database",
					Instances: []terraformStateInstance{
						{
							Attributes: map[string]any{},
						},
					},
				},
			},
		}

		resourceID, found := terraformResourceIDFromState(state, "oci_database_autonomous_container_database.test_autonomous_container_database")
		if !found {
			t.Fatal("expected resource to be found")
		}
		if resourceID != "" {
			t.Fatalf("expected empty resource id, got %s", resourceID)
		}
	})
}

func TestIsUsableSharedDependencyID(t *testing.T) {
	testCases := []struct {
		name       string
		resourceID string
		expected   bool
	}{
		{
			name:       "empty string",
			resourceID: "",
			expected:   false,
		},
		{
			name:       "whitespace string",
			resourceID: "   ",
			expected:   false,
		},
		{
			name:       "real ocid",
			resourceID: "ocid1.autonomouscontainerdatabase.oc1..example",
			expected:   true,
		},
		{
			name:       "env placeholder expression",
			resourceID: "${TF_VAR_autonomous_container_database_id}",
			expected:   false,
		},
		{
			name:       "terraform var expression",
			resourceID: "${var.autonomous_container_database_id}",
			expected:   false,
		},
		{
			name:       "bare tf var token",
			resourceID: "TF_VAR_autonomous_container_database_id",
			expected:   false,
		},
		{
			name:       "trimmed ocid",
			resourceID: "  ocid1.autonomouscontainerdatabase.oc1..example  ",
			expected:   true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := isUsableSharedDependencyID(testCase.resourceID)
			if actual != testCase.expected {
				t.Fatalf("unexpected usability for %q: got %t want %t", testCase.resourceID, actual, testCase.expected)
			}
		})
	}
}
