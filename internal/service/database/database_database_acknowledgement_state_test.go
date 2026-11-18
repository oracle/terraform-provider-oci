// Copyright (c) 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

const databaseAcknowledgementStatePath = "database.0.managed_software_update_details.0.preference_details.0.version_scheme_details.0.is_older_version_acknowledged"

func databaseAcknowledgementTestData(t *testing.T, prior *bool) *schema.ResourceData {
	t.Helper()
	if prior == nil {
		data := DatabaseDatabaseResource().Data(nil)
		data.SetId("D1")
		return data
	}
	config := map[string]interface{}{"source": "NONE"}
	if prior != nil {
		config["database"] = []interface{}{map[string]interface{}{
			"db_name": "ACKTEST",
			"managed_software_update_details": []interface{}{map[string]interface{}{
				"is_enrolled": true,
				"preference_details": []interface{}{map[string]interface{}{
					"version_scheme_details": []interface{}{map[string]interface{}{
						"source": "VERSION_SERIES", "major_version": "19.0.0.0", "version_preference": "ORACLE_DB_N",
						"is_older_version_acknowledged": *prior,
					}},
				}},
			}},
		}}
	}
	data := schema.TestResourceDataRaw(t, DatabaseDatabaseResource().Schema, config)
	data.SetId("D1")
	return data
}

func databaseAcknowledgementTestResponse(apiValue *bool) *oci_database.Database {
	return &oci_database.Database{
		Id: oci_common.String("D1"), DbName: oci_common.String("ACKTEST"),
		LifecycleState: oci_database.DatabaseLifecycleStateAvailable,
		ManagedSoftwareUpdateDetails: &oci_database.ManagedSoftwareUpdateDetails{
			IsEnrolled: oci_common.Bool(true),
			PreferenceDetails: &oci_database.ManagedSoftwareUpdatePreferenceDetails{
				VersionSchemeDetails: oci_database.VersionSchemeFromVersionSeriesDetails{
					MajorVersion:               oci_common.String("19.0.0.0"),
					VersionPreference:          oci_database.VersionSchemeFromVersionSeriesDetailsVersionPreferenceEnum("ORACLE_DB_N"),
					IsOlderVersionAcknowledged: apiValue,
				},
			},
		},
	}
}

func databaseAcknowledgementMappedScheme(t *testing.T, mapped map[string]interface{}) map[string]interface{} {
	t.Helper()
	managed := mapped["managed_software_update_details"].([]interface{})[0].(map[string]interface{})
	preferences := managed["preference_details"].([]interface{})[0].(map[string]interface{})
	return preferences["version_scheme_details"].([]interface{})[0].(map[string]interface{})
}

func TestDatabaseDatabaseAcknowledgementStateRefresh(t *testing.T) {
	for _, tc := range []struct {
		name             string
		prior, api, want *bool
	}{
		{"preserve_true_when_api_omits", oci_common.Bool(true), nil, oci_common.Bool(true)},
		{"preserve_false_when_api_omits", oci_common.Bool(false), nil, oci_common.Bool(false)},
		{"import_without_prior_does_not_invent_value", nil, nil, nil},
		{"api_false_overrides_prior_true", oci_common.Bool(true), oci_common.Bool(false), oci_common.Bool(false)},
		{"api_true_overrides_prior_false", oci_common.Bool(false), oci_common.Bool(true), oci_common.Bool(true)},
		{"import_honors_api_true", nil, oci_common.Bool(true), oci_common.Bool(true)},
		{"import_honors_api_false", nil, oci_common.Bool(false), oci_common.Bool(false)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			data := databaseAcknowledgementTestData(t, tc.prior)
			response := databaseAcknowledgementTestResponse(tc.api)
			crud := &DatabaseDatabaseResourceCrud{BaseCrud: tfresource.BaseCrud{D: data}, Res: response}
			mapped := crud.DatabaseToMap(response)
			scheme := databaseAcknowledgementMappedScheme(t, mapped)
			got, exists := scheme["is_older_version_acknowledged"]
			if tc.want == nil {
				if exists {
					t.Fatalf("mapped acknowledgment=%v; expected absent without API/prior value", got)
				}
			} else if !exists || got != *tc.want {
				t.Fatalf("mapped acknowledgment=%v, present=%t; want %t (including explicit false)", got, exists, *tc.want)
			}
			if scheme["source"] != "VERSION_SERIES" || scheme["major_version"] != "19.0.0.0" || scheme["version_preference"] != "ORACLE_DB_N" {
				t.Fatalf("unrelated version-series fields changed: %v", scheme)
			}
			// A read fallback must not mutate the actual API model.
			apiSeries := response.ManagedSoftwareUpdateDetails.PreferenceDetails.VersionSchemeDetails.(oci_database.VersionSchemeFromVersionSeriesDetails)
			if apiSeries.IsOlderVersionAcknowledged != tc.api {
				t.Fatal("API response acknowledgment pointer was mutated")
			}
			if err := crud.SetData(); err != nil {
				t.Fatalf("SetData: %v", err)
			}
			if tc.want != nil {
				if got := data.State().Attributes[databaseAcknowledgementStatePath]; got != strconv.FormatBool(*tc.want) {
					t.Fatalf("serialized nested state acknowledgment=%q; want %t", got, *tc.want)
				}
			}
		})
	}
}

func TestDatabaseDatabaseAcknowledgementStateRepeatedRefresh(t *testing.T) {
	data := databaseAcknowledgementTestData(t, oci_common.Bool(true))
	response := databaseAcknowledgementTestResponse(nil)
	crud := &DatabaseDatabaseResourceCrud{BaseCrud: tfresource.BaseCrud{D: data}, Res: response}
	for i := 0; i < 2; i++ {
		if err := crud.SetData(); err != nil {
			t.Fatalf("refresh %d: %v", i, err)
		}
		state := data.State()
		if got := state.Attributes[databaseAcknowledgementStatePath]; got != "true" {
			t.Fatalf("refresh %d lost serialized true acknowledgment: %q", i, got)
		}
		data = DatabaseDatabaseResource().Data(state)
		crud.D = data
	}
}

func TestDatabaseDatabaseAcknowledgementStateDoesNotRestoreMissingBlocks(t *testing.T) {
	for _, missing := range []string{"managed_details", "preference_details", "version_scheme_details"} {
		t.Run(missing, func(t *testing.T) {
			response := databaseAcknowledgementTestResponse(nil)
			switch missing {
			case "managed_details":
				response.ManagedSoftwareUpdateDetails = nil
			case "preference_details":
				response.ManagedSoftwareUpdateDetails.PreferenceDetails = nil
			case "version_scheme_details":
				response.ManagedSoftwareUpdateDetails.PreferenceDetails.VersionSchemeDetails = nil
			}
			data := databaseAcknowledgementTestData(t, oci_common.Bool(true))
			crud := &DatabaseDatabaseResourceCrud{BaseCrud: tfresource.BaseCrud{D: data}, Res: response}
			mapped := crud.DatabaseToMap(response)
			if missing == "managed_details" {
				if _, exists := mapped["managed_software_update_details"]; exists {
					t.Fatal("invented missing managed-software block")
				}
				return
			}
			managed := mapped["managed_software_update_details"].([]interface{})[0].(map[string]interface{})
			if missing == "preference_details" {
				if _, exists := managed["preference_details"]; exists {
					t.Fatal("invented missing preference block")
				}
				return
			}
			preferences := managed["preference_details"].([]interface{})[0].(map[string]interface{})
			if _, exists := preferences["version_scheme_details"]; exists {
				t.Fatal("invented missing version-scheme block")
			}
		})
	}
}
