// Copyright (c) 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package psql

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func newPsqlStorageCrud(t *testing.T, storage map[string]interface{}) *PsqlDbSystemResourceCrud {
	t.Helper()
	data := schema.TestResourceDataRaw(t, PsqlDbSystemResource().Schema, map[string]interface{}{
		"storage_details": []interface{}{storage},
	})
	return &PsqlDbSystemResourceCrud{BaseCrud: tfresource.BaseCrud{D: data}}
}

func assertPsqlStorageJSON(t *testing.T, details interface{}, want map[string]interface{}) {
	t.Helper()
	encoded, err := json.Marshal(details)
	if err != nil {
		t.Fatalf("marshal storage details: %v", err)
	}
	var got map[string]interface{}
	if err := json.Unmarshal(encoded, &got); err != nil {
		t.Fatalf("decode storage details: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("storage JSON = %s, want %v", encoded, want)
	}
}

func TestPsqlDbSystemCreateStorageDetails(t *testing.T) {
	for _, tc := range []struct {
		name    string
		storage map[string]interface{}
		want    map[string]interface{}
	}{
		{
			name: "AD-local storage preserves creation fields",
			storage: map[string]interface{}{
				"system_type":           "OCI_OPTIMIZED_STORAGE",
				"is_regionally_durable": false,
				"availability_domain":   "test:AD-1",
				"iops":                  "30000",
				"kms_key_id":            "ocid1.key.oc1..storage",
			},
			want: map[string]interface{}{
				"systemType":          "OCI_OPTIMIZED_STORAGE",
				"isRegionallyDurable": false,
				"availabilityDomain":  "test:AD-1",
				"iops":                float64(30000),
				"kmsKeyId":            "ocid1.key.oc1..storage",
			},
		},
		{
			name: "regional storage accepts lowercase discriminator",
			storage: map[string]interface{}{
				"system_type":           "oci_optimized_storage",
				"is_regionally_durable": true,
			},
			want: map[string]interface{}{
				"systemType":          "OCI_OPTIMIZED_STORAGE",
				"isRegionallyDurable": true,
				"availabilityDomain":  nil,
				"iops":                nil,
				"kmsKeyId":            nil,
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			details, err := newPsqlStorageCrud(t, tc.storage).mapToStorageDetails("storage_details.0.%s")
			if err != nil {
				t.Fatalf("mapToStorageDetails: %v", err)
			}
			assertPsqlStorageJSON(t, details, tc.want)
		})
	}
}

func TestPsqlDbSystemUpdateStorageDetails(t *testing.T) {
	for _, systemType := range []string{"OCI_OPTIMIZED_STORAGE", "oci_optimized_storage"} {
		t.Run(systemType, func(t *testing.T) {
			crud := newPsqlStorageCrud(t, map[string]interface{}{
				"system_type":           systemType,
				"is_regionally_durable": false,
				"availability_domain":   "test:AD-1",
				"iops":                  "45000",
				"kms_key_id":            "ocid1.key.oc1..updated",
			})
			details, err := crud.mapToUpdateStorageDetailsParams("storage_details.0.%s")
			if err != nil {
				t.Fatalf("mapToUpdateStorageDetailsParams: %v", err)
			}
			assertPsqlStorageJSON(t, details, map[string]interface{}{
				"iops":     float64(45000),
				"kmsKeyId": "ocid1.key.oc1..updated",
			})
		})
	}
}

func TestPsqlDbSystemStorageMappingErrors(t *testing.T) {
	for _, tc := range []struct {
		name       string
		systemType string
		iops       string
		wantError  string
	}{
		{"invalid IOPS", "OCI_OPTIMIZED_STORAGE", "invalid", "unable to convert iops"},
		{"overflow IOPS", "OCI_OPTIMIZED_STORAGE", "9223372036854775808", "unable to convert iops"},
		{"unknown storage type", "UNKNOWN", "30000", "unknown system_type"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			crud := newPsqlStorageCrud(t, map[string]interface{}{
				"system_type":           tc.systemType,
				"is_regionally_durable": true,
				"iops":                  tc.iops,
			})
			_, createErr := crud.mapToStorageDetails("storage_details.0.%s")
			_, updateErr := crud.mapToUpdateStorageDetailsParams("storage_details.0.%s")
			for name, err := range map[string]error{"create": createErr, "update": updateErr} {
				if err == nil || !strings.Contains(err.Error(), tc.wantError) {
					t.Errorf("%s storage mapping error = %v, want %q", name, err, tc.wantError)
				}
			}
		})
	}
}
