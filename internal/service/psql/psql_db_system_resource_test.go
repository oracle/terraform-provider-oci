// Copyright (c) 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package psql

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_psql "github.com/oracle/oci-go-sdk/v65/psql"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

const psqlTestDbSystemID = "ocid1.postgresqldbsystem.oc1..example"

func psqlTestConfig(version string) map[string]interface{} {
	return map[string]interface{}{
		"compartment_id": "ocid1.compartment.oc1..example",
		"db_version":     version,
		"display_name":   "example",
		"shape":          "VM.Standard.E4.Flex",
		"instances_details": []interface{}{map[string]interface{}{
			"display_name": "example-node", "description": "example", "private_ip": "10.0.0.2",
		}},
		"network_details": []interface{}{map[string]interface{}{"subnet_id": "ocid1.subnet.oc1..example"}},
		"storage_details": []interface{}{map[string]interface{}{
			"is_regionally_durable": true,
			"system_type":           "OCI_OPTIMIZED_STORAGE",
		}},
	}
}

func TestPsqlDbSystemVersionDiff(t *testing.T) {
	for _, tc := range []struct {
		name, oldVersion, newVersion string
		oldAllowed                   bool
		newAllowed                   interface{}
		legacyState                  bool
		wantError                    bool
		wantReplacement              bool
	}{
		{name: "opt in from legacy state", oldVersion: "16", newVersion: "17", newAllowed: true, legacyState: true},
		{name: "enable upgrades", oldVersion: "16", newVersion: "17", newAllowed: true},
		{name: "upgrades remain enabled", oldVersion: "16", newVersion: "17", oldAllowed: true, newAllowed: true},
		{name: "numeric major comparison", oldVersion: "9", newVersion: "10", newAllowed: true},
		{name: "reported minor version", oldVersion: "16.4", newVersion: "17", newAllowed: true},
		{name: "unchanged version", oldVersion: "17", newVersion: "17", newAllowed: true},
		{name: "create with upgrades enabled", newVersion: "17", newAllowed: true},
		{name: "downgrade opted in", oldVersion: "17", newVersion: "16", newAllowed: true, wantError: true},
		{name: "minor change opted in", oldVersion: "17.1", newVersion: "17.2", newAllowed: true, wantError: true},
		{name: "invalid target opted in", oldVersion: "16", newVersion: "invalid", newAllowed: true, wantError: true},
		{name: "permission omitted replaces", oldVersion: "16", newVersion: "17", wantReplacement: true},
		{name: "upgrades disabled replaces", oldVersion: "16", newVersion: "17", newAllowed: false, wantReplacement: true},
		{name: "disable upgrades replaces", oldVersion: "16", newVersion: "17", oldAllowed: true, newAllowed: false, wantReplacement: true},
		{name: "permission removed replaces", oldVersion: "16", newVersion: "17", oldAllowed: true, wantReplacement: true},
		{name: "legacy downgrade replaces", oldVersion: "17", newVersion: "16", wantReplacement: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			resource := PsqlDbSystemResource()
			var state *terraform.InstanceState
			if tc.oldVersion != "" {
				oldConfig := psqlTestConfig(tc.oldVersion)
				oldConfig["allow_major_version_upgrade"] = tc.oldAllowed
				data := schema.TestResourceDataRaw(t, resource.Schema, oldConfig)
				data.SetId(psqlTestDbSystemID)
				state = data.State()
				if tc.legacyState {
					delete(state.Attributes, "allow_major_version_upgrade")
				}
			}
			config := psqlTestConfig(tc.newVersion)
			if tc.newAllowed != nil {
				config["allow_major_version_upgrade"] = tc.newAllowed
			}
			diff, err := resource.Diff(context.Background(), state, terraform.NewResourceConfigRaw(config), nil)
			if (err != nil) != tc.wantError {
				t.Fatalf("Diff() error = %v, want error %v", err, tc.wantError)
			}
			if err == nil && state != nil {
				replaces := diff != nil && diff.RequiresNew()
				if replaces != tc.wantReplacement {
					t.Fatalf("Diff() replacement = %t, want %t", replaces, tc.wantReplacement)
				}
			}
		})
	}
}

func TestPsqlDbSystemUpgradeConfigDiff(t *testing.T) {
	for _, explicitConfig := range []bool{false, true} {
		t.Run(fmt.Sprintf("explicit_config_%t", explicitConfig), func(t *testing.T) {
			resource := PsqlDbSystemResource()
			old := psqlTestConfig("16")
			old["config_id"] = "ocid1.postgresqlconfiguration.oc1..old"
			data := schema.TestResourceDataRaw(t, resource.Schema, old)
			data.SetId(psqlTestDbSystemID)
			config := psqlTestConfig("17")
			config["allow_major_version_upgrade"] = true
			if explicitConfig {
				config["config_id"] = "ocid1.postgresqlconfiguration.oc1..target"
			}
			state := data.State()
			configID := cty.NullVal(cty.String)
			if explicitConfig {
				configID = cty.StringVal(config["config_id"].(string))
			}
			// The SDK carries the current protocol configuration alongside prior state.
			state.RawConfig = cty.ObjectVal(map[string]cty.Value{"config_id": configID})
			diff, err := resource.Diff(context.Background(), state, terraform.NewResourceConfigRaw(config), nil)
			if err != nil {
				t.Fatal(err)
			}
			configDiff := diff.Attributes["config_id"]
			if configDiff == nil {
				t.Fatal("upgrade must plan a target configuration or a new service default")
			}
			if configDiff.NewComputed != !explicitConfig {
				t.Fatalf("config_id computed = %t, want %t", configDiff.NewComputed, !explicitConfig)
			}
			if explicitConfig && configDiff.New != config["config_id"] {
				t.Fatalf("target config = %q, want %q", configDiff.New, config["config_id"])
			}
		})
	}
}

func TestPsqlDbSystemUpgradeOptionsValidation(t *testing.T) {
	for _, tc := range []struct {
		name      string
		options   []interface{}
		wantError bool
	}{
		{name: "omitted"},
		{name: "empty", options: []interface{}{}},
		{name: "verification", options: []interface{}{"RUN_POST_UPGRADE_VERIFICATION"}},
		{name: "vacuum", options: []interface{}{"RUN_POST_UPGRADE_VACUUM"}},
		{name: "both", options: []interface{}{"RUN_POST_UPGRADE_VERIFICATION", "RUN_POST_UPGRADE_VACUUM"}},
		{name: "unknown", options: []interface{}{"RUN_SOMETHING_ELSE"}, wantError: true},
		{name: "wrong case", options: []interface{}{"run_post_upgrade_vacuum"}, wantError: true},
		{name: "empty entry", options: []interface{}{""}, wantError: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			config := psqlTestConfig("17")
			if tc.options != nil {
				config["upgrade_options"] = tc.options
			}
			diagnostics := PsqlDbSystemResource().Validate(terraform.NewResourceConfigRaw(config))
			if diagnostics.HasError() != tc.wantError {
				t.Fatalf("Validate() = %v, want error %t", diagnostics, tc.wantError)
			}
		})
	}
}

func TestPsqlDbSystemUpgradePermissionDefaultsToFalse(t *testing.T) {
	data := schema.TestResourceDataRaw(t, PsqlDbSystemResource().Schema, psqlTestConfig("17"))
	if data.Get("allow_major_version_upgrade") != false {
		t.Fatal("major version upgrades must require explicit opt-in")
	}
}

func TestPsqlDbSystemConfigurationIDValidation(t *testing.T) {
	for _, tc := range []struct {
		name      string
		configID  interface{}
		wantError bool
	}{
		{name: "omitted"},
		{name: "empty", configID: "", wantError: true},
		{name: "minimum length", configID: "x"},
		{name: "maximum length", configID: strings.Repeat("x", 255)},
		{name: "too long", configID: strings.Repeat("x", 256), wantError: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			config := psqlTestConfig("17")
			if tc.configID != nil {
				config["config_id"] = tc.configID
			}
			diagnostics := PsqlDbSystemResource().Validate(terraform.NewResourceConfigRaw(config))
			if diagnostics.HasError() != tc.wantError {
				t.Fatalf("Validate() = %v, want error %t", diagnostics, tc.wantError)
			}
		})
	}
}

type psqlTestPriorState struct {
	allowed  bool
	configID string
	options  []string
}

func psqlTestUpdateData(t *testing.T, changes map[string]*terraform.ResourceAttrDiff, rawConfig ...cty.Value) *schema.ResourceData {
	t.Helper()
	return psqlTestUpdateDataWithState(t, psqlTestPriorState{}, changes, rawConfig...)
}

func psqlTestUpdateDataWithState(t *testing.T, prior psqlTestPriorState, changes map[string]*terraform.ResourceAttrDiff, rawConfig ...cty.Value) *schema.ResourceData {
	t.Helper()
	diff := &terraform.InstanceDiff{Attributes: changes}
	if len(rawConfig) != 0 {
		diff.RawConfig = rawConfig[0]
	}
	if prior.configID == "" {
		prior.configID = "ocid1.postgresqlconfiguration.oc1..old"
	}
	attributes := map[string]string{
		"db_version":                  "16",
		"allow_major_version_upgrade": strconv.FormatBool(prior.allowed),
		"config_id":                   prior.configID,
		"display_name":                "example",
		"patch_operations.#":          "0",
	}
	if prior.options != nil {
		attributes["upgrade_options.#"] = strconv.Itoa(len(prior.options))
		for i, option := range prior.options {
			attributes[fmt.Sprintf("upgrade_options.%d", i)] = option
		}
	}
	data, err := schema.InternalMap(PsqlDbSystemResource().Schema).Data(&terraform.InstanceState{
		ID:         psqlTestDbSystemID,
		Attributes: attributes,
	}, diff)
	if err != nil {
		t.Fatal(err)
	}
	return data
}

type psqlTestDispatcher func(*http.Request) (*http.Response, error)

func (f psqlTestDispatcher) Do(request *http.Request) (*http.Response, error) { return f(request) }

type psqlTestSigner struct{}

func (psqlTestSigner) Sign(*http.Request) error { return nil }

func psqlTestClient(dispatch psqlTestDispatcher) *oci_psql.PostgresqlClient {
	return &oci_psql.PostgresqlClient{BaseClient: oci_common.BaseClient{
		Host:       "https://postgresql.example.invalid",
		BasePath:   "20220915",
		HTTPClient: dispatch,
		Signer:     psqlTestSigner{},
		UserAgent:  "terraform-provider-oci-unit-test",
	}}
}

func psqlTestResponse(request *http.Request, status int, body string, workID string) *http.Response {
	header := make(http.Header)
	header.Set("Content-Type", "application/json")
	if workID != "" {
		header.Set("opc-work-request-id", workID)
	}
	return &http.Response{StatusCode: status, Header: header, Body: io.NopCloser(strings.NewReader(body)), Request: request}
}

func psqlTestWorkRequest(status string) string {
	return fmt.Sprintf(`{"id":"upgrade-work","status":%q,"timeFinished":"2026-01-01T00:00:00Z","resources":[{"entityType":"dbsystem","identifier":%q,"actionType":"UPDATED"}]}`, status, psqlTestDbSystemID)
}

func TestPsqlDbSystemUpgradeRequestAndRefresh(t *testing.T) {
	for _, tc := range []struct {
		name            string
		options         []string
		configID        string
		rename          bool
		resize          bool
		unchangedConfig bool
		alreadyEnabled  bool
		computedConfig  bool
		previousOptions []string
	}{
		{name: "omitted options and previous default configuration"},
		{name: "empty options", options: []string{}},
		{name: "previous options removed", previousOptions: []string{"RUN_POST_UPGRADE_VERIFICATION", "RUN_POST_UPGRADE_VACUUM"}},
		{name: "previous options explicitly emptied", options: []string{}, previousOptions: []string{"RUN_POST_UPGRADE_VERIFICATION", "RUN_POST_UPGRADE_VACUUM"}},
		{name: "verification only", options: []string{"RUN_POST_UPGRADE_VERIFICATION"}},
		{name: "vacuum only", options: []string{"RUN_POST_UPGRADE_VACUUM"}},
		{name: "target configuration and both operations", configID: "ocid1.postgresqlconfiguration.oc1..target", options: []string{"RUN_POST_UPGRADE_VERIFICATION", "RUN_POST_UPGRADE_VACUUM"}},
		{name: "upgrade and rename", configID: "ocid1.postgresqlconfiguration.oc1..target", rename: true},
		{name: "upgrade and resize uses new default configuration", resize: true},
		{name: "computed default configuration is omitted", computedConfig: true},
		{name: "upgrade and resize with computed default configuration", resize: true, computedConfig: true},
		{name: "explicit unchanged configuration reaches service validation", configID: "ocid1.postgresqlconfiguration.oc1..old", unchangedConfig: true},
		{name: "upgrades already enabled without a permission diff", alreadyEnabled: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			changes := map[string]*terraform.ResourceAttrDiff{
				"db_version": {Old: "16", New: "17"},
			}
			oldConfigID := "ocid1.postgresqlconfiguration.oc1..old"
			targetConfigID := tc.configID
			if tc.configID == "" {
				oldConfigID = "ocid1.postgresqldefaultconfiguration.oc1..old"
				targetConfigID = "ocid1.postgresqldefaultconfiguration.oc1..target"
			}
			if !tc.alreadyEnabled {
				changes["allow_major_version_upgrade"] = &terraform.ResourceAttrDiff{Old: "false", New: "true"}
			}
			if tc.configID != "" && !tc.unchangedConfig {
				changes["config_id"] = &terraform.ResourceAttrDiff{Old: oldConfigID, New: tc.configID}
			}
			if tc.computedConfig {
				changes["config_id"] = &terraform.ResourceAttrDiff{Old: oldConfigID, NewComputed: true}
			}
			if tc.options != nil || tc.previousOptions != nil {
				changes["upgrade_options.#"] = &terraform.ResourceAttrDiff{Old: strconv.Itoa(len(tc.previousOptions)), New: strconv.Itoa(len(tc.options))}
				for i, option := range tc.previousOptions {
					changes[fmt.Sprintf("upgrade_options.%d", i)] = &terraform.ResourceAttrDiff{Old: option, NewRemoved: true}
				}
				for i, option := range tc.options {
					changes[fmt.Sprintf("upgrade_options.%d", i)] = &terraform.ResourceAttrDiff{New: option}
				}
			}
			if tc.rename {
				changes["display_name"] = &terraform.ResourceAttrDiff{Old: "example", New: "renamed"}
			}
			if tc.resize {
				changes["shape"] = &terraform.ResourceAttrDiff{Old: "VM.Standard.E4.Flex", New: "VM.Standard.E5.Flex"}
			}
			var calls []string
			rawConfigID := cty.NullVal(cty.String)
			if tc.configID != "" {
				rawConfigID = cty.StringVal(tc.configID)
			}
			data := psqlTestUpdateDataWithState(t, psqlTestPriorState{allowed: tc.alreadyEnabled, configID: oldConfigID, options: tc.previousOptions}, changes, cty.ObjectVal(map[string]cty.Value{"config_id": rawConfigID}))
			if tc.alreadyEnabled && data.HasChange("allow_major_version_upgrade") {
				t.Fatal("test fixture must keep the permission enabled without changing it")
			}
			crud := &PsqlDbSystemResourceCrud{BaseCrud: tfresource.BaseCrud{D: data}}
			crud.Client = psqlTestClient(func(request *http.Request) (*http.Response, error) {
				path := strings.TrimPrefix(request.URL.Path, "/20220915")
				calls = append(calls, request.Method+" "+path)
				switch {
				case request.Method == http.MethodPost && path == "/dbSystems/"+psqlTestDbSystemID+"/actions/upgrade":
					var payload struct {
						DBVersion      string          `json:"dbVersion"`
						ConfigID       *string         `json:"configId"`
						Options        []string        `json:"options"`
						UpgradeOptions json.RawMessage `json:"upgrade_options"`
					}
					if err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
						t.Fatal(err)
					}
					if len(payload.UpgradeOptions) != 0 {
						t.Fatal("Terraform upgrade_options must use the API options field in the request body")
					}
					if payload.DBVersion != "17" {
						t.Fatalf("target version = %q", payload.DBVersion)
					}
					if tc.configID == "" && payload.ConfigID != nil {
						t.Fatalf("previous configuration leaked into upgrade: %q", *payload.ConfigID)
					}
					if tc.configID != "" && (payload.ConfigID == nil || *payload.ConfigID != tc.configID) {
						t.Fatalf("target configuration = %v", payload.ConfigID)
					}
					if len(payload.Options) != len(tc.options) || (len(tc.options) > 0 && !reflect.DeepEqual(payload.Options, tc.options)) {
						t.Fatalf("options = %v, want %v", payload.Options, tc.options)
					}
					return psqlTestResponse(request, http.StatusAccepted, "", "upgrade-work"), nil
				case request.Method == http.MethodGet && path == "/workRequests/upgrade-work":
					return psqlTestResponse(request, http.StatusOK, psqlTestWorkRequest("SUCCEEDED"), ""), nil
				case request.Method == http.MethodGet && path == "/dbSystems/"+psqlTestDbSystemID:
					return psqlTestResponse(request, http.StatusOK, fmt.Sprintf(`{"id":%q,"dbVersion":"17","configId":%q,"lifecycleState":"ACTIVE"}`, psqlTestDbSystemID, targetConfigID), ""), nil
				case (tc.rename || tc.resize) && request.Method == http.MethodPut && path == "/dbSystems/"+psqlTestDbSystemID:
					var payload map[string]interface{}
					if err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
						t.Fatal(err)
					}
					if tc.rename && payload["displayName"] != "renamed" {
						t.Fatalf("rename payload = %v", payload)
					}
					if tc.resize {
						configuration, ok := payload["dbConfigurationParams"].(map[string]interface{})
						if !ok || configuration["configId"] != targetConfigID {
							t.Fatalf("resize must use the upgraded configuration: %v", payload)
						}
					} else if _, exists := payload["dbConfigurationParams"]; exists {
						t.Fatalf("upgrade configuration reapplied by regular update: %v", payload)
					}
					return psqlTestResponse(request, http.StatusAccepted, "", "upgrade-work"), nil
				default:
					t.Fatalf("unexpected API call: %s %s", request.Method, path)
					return nil, fmt.Errorf("unexpected API call")
				}
			})
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := crud.UpdateWithContext(ctx); err != nil {
				t.Fatal(err)
			}
			wantCalls := []string{"POST /dbSystems/" + psqlTestDbSystemID + "/actions/upgrade", "GET /workRequests/upgrade-work", "GET /dbSystems/" + psqlTestDbSystemID}
			if tc.rename || tc.resize {
				wantCalls = append(wantCalls, "PUT /dbSystems/"+psqlTestDbSystemID, "GET /workRequests/upgrade-work", "GET /dbSystems/"+psqlTestDbSystemID)
			} else {
				wantCalls = append(wantCalls, "GET /dbSystems/"+psqlTestDbSystemID)
			}
			if !reflect.DeepEqual(calls, wantCalls) {
				t.Fatalf("API calls = %v, want %v", calls, wantCalls)
			}
			if crud.Res == nil || crud.Res.DbVersion == nil || *crud.Res.DbVersion != "17" {
				t.Fatalf("upgraded DB system was not refreshed: %v", crud.Res)
			}
			if err := crud.SetData(); err != nil {
				t.Fatal(err)
			}
			state := data.State()
			if state.Attributes["db_version"] != "17" || state.Attributes["config_id"] != targetConfigID {
				t.Fatalf("refreshed state version/config = %q/%q, want 17/%q", state.Attributes["db_version"], state.Attributes["config_id"], targetConfigID)
			}
		})
	}
}

func TestPsqlDbSystemUpgradeControlsAloneDoNotUpgrade(t *testing.T) {
	for name, changes := range map[string]map[string]*terraform.ResourceAttrDiff{
		"upgrade options only": {
			"upgrade_options.#": {Old: "0", New: "1"},
			"upgrade_options.0": {New: "RUN_POST_UPGRADE_VACUUM"},
		},
		"enabling upgrades only": {"allow_major_version_upgrade": {Old: "false", New: "true"}},
	} {
		t.Run(name, func(t *testing.T) {
			data := psqlTestUpdateData(t, changes)
			crud := &PsqlDbSystemResourceCrud{BaseCrud: tfresource.BaseCrud{D: data}, Client: psqlTestClient(func(request *http.Request) (*http.Response, error) {
				if request.Method != http.MethodGet {
					t.Fatalf("control-only change sent %s %s", request.Method, request.URL.Path)
				}
				return psqlTestResponse(request, http.StatusOK, fmt.Sprintf(`{"id":%q,"dbVersion":"16","lifecycleState":"ACTIVE"}`, psqlTestDbSystemID), ""), nil
			})}
			if err := crud.UpdateWithContext(context.Background()); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestPsqlDbSystemVersionUpdateWithoutOptInIsRejected(t *testing.T) {
	for _, operation := range []string{"update", "upgrade"} {
		t.Run(operation, func(t *testing.T) {
			data := psqlTestUpdateData(t, map[string]*terraform.ResourceAttrDiff{"db_version": {Old: "16", New: "17"}})
			crud := &PsqlDbSystemResourceCrud{BaseCrud: tfresource.BaseCrud{D: data}, Client: psqlTestClient(func(request *http.Request) (*http.Response, error) {
				t.Fatalf("version update without opt-in sent %s %s", request.Method, request.URL.Path)
				return nil, fmt.Errorf("unexpected API call")
			})}
			var err error
			if operation == "update" {
				err = crud.UpdateWithContext(context.Background())
			} else {
				err = crud.UpgradeDbSystem(context.Background())
			}
			if err == nil || !strings.Contains(err.Error(), "allow_major_version_upgrade") {
				t.Fatalf("expected allow_major_version_upgrade error, got %v", err)
			}
		})
	}
}

func TestPsqlDbSystemUpgradeFailure(t *testing.T) {
	for _, tc := range []struct {
		name, workStatus, wantError string
		status                      int
		workID                      string
		configID                    string
		serviceMessage              string
	}{
		{name: "API rejects upgrade", status: http.StatusBadRequest, wantError: "InvalidParameter"},
		{name: "custom configuration requires target configuration", status: http.StatusBadRequest, serviceMessage: "configId is required when upgrading a custom configuration", wantError: "configId is required when upgrading a custom configuration"},
		{name: "target configuration incompatible", status: http.StatusBadRequest, configID: "ocid1.postgresqlconfiguration.oc1..incompatible", serviceMessage: "configId is incompatible with target dbVersion", wantError: "configId is incompatible with target dbVersion"},
		{name: "missing work request ID", status: http.StatusAccepted, wantError: "opc-work-request-id"},
		{name: "failed work request", status: http.StatusAccepted, workID: "upgrade-work", workStatus: "FAILED", wantError: "upgrade verification failed"},
		{name: "canceled work request", status: http.StatusAccepted, workID: "upgrade-work", workStatus: "CANCELED", wantError: "upgrade verification failed"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			changes := map[string]*terraform.ResourceAttrDiff{
				"db_version":                  {Old: "16", New: "17"},
				"allow_major_version_upgrade": {Old: "false", New: "true"},
			}
			if tc.configID != "" {
				changes["config_id"] = &terraform.ResourceAttrDiff{Old: "ocid1.postgresqlconfiguration.oc1..old", New: tc.configID}
			}
			data := psqlTestUpdateData(t, changes)
			postCount := 0
			crud := &PsqlDbSystemResourceCrud{BaseCrud: tfresource.BaseCrud{D: data}, Client: psqlTestClient(func(request *http.Request) (*http.Response, error) {
				switch {
				case request.Method == http.MethodPost:
					postCount++
					if postCount > 1 {
						t.Fatal("failed upgrade must not send subsequent writes")
					}
					message := tc.serviceMessage
					if message == "" {
						message = "unsupported upgrade"
					} else {
						var payload struct {
							ConfigID *string `json:"configId"`
						}
						if err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
							t.Fatal(err)
						}
						if tc.configID == "" && payload.ConfigID != nil {
							t.Fatal("old custom configuration must not be implicitly forwarded")
						}
						if tc.configID != "" && (payload.ConfigID == nil || *payload.ConfigID != tc.configID) {
							t.Fatal("target custom configuration was not forwarded for service validation")
						}
					}
					return psqlTestResponse(request, tc.status, fmt.Sprintf(`{"code":"InvalidParameter","message":%q}`, message), tc.workID), nil
				case request.Method == http.MethodGet && strings.HasSuffix(request.URL.Path, "/errors"):
					return psqlTestResponse(request, http.StatusOK, `{"items":[{"message":"upgrade verification failed"}]}`, ""), nil
				case request.Method == http.MethodGet && strings.HasSuffix(request.URL.Path, "/workRequests/upgrade-work"):
					return psqlTestResponse(request, http.StatusOK, psqlTestWorkRequest(tc.workStatus), ""), nil
				default:
					t.Fatalf("upgrade failure must stop subsequent API calls: %s %s", request.Method, request.URL.Path)
					return nil, fmt.Errorf("unexpected API call")
				}
			})}
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			err := crud.UpdateWithContext(ctx)
			if err == nil || !strings.Contains(err.Error(), tc.wantError) {
				t.Fatalf("UpdateWithContext() error = %v, want %q", err, tc.wantError)
			}
		})
	}
}

func TestPsqlDbSystemStartBeforeUpgradeUsesCurrentConfiguration(t *testing.T) {
	data := psqlTestUpdateData(t, map[string]*terraform.ResourceAttrDiff{
		"db_version":                  {Old: "16", New: "17"},
		"allow_major_version_upgrade": {Old: "false", New: "true"},
		"config_id":                   {Old: "ocid1.postgresqlconfiguration.oc1..old", New: "ocid1.postgresqlconfiguration.oc1..target"},
	})
	started := false
	crud := &PsqlDbSystemResourceCrud{BaseCrud: tfresource.BaseCrud{D: data}, Client: psqlTestClient(func(request *http.Request) (*http.Response, error) {
		if request.Method == http.MethodPost && strings.HasSuffix(request.URL.Path, "/actions/start") {
			var payload struct {
				ConfigID string `json:"configId"`
			}
			if err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
				t.Fatal(err)
			}
			if payload.ConfigID != "ocid1.postgresqlconfiguration.oc1..old" {
				t.Fatalf("start must use the current-version configuration, got %q", payload.ConfigID)
			}
			started = true
			return psqlTestResponse(request, http.StatusAccepted, "", ""), nil
		}
		if started && request.Method == http.MethodGet && strings.HasSuffix(request.URL.Path, "/dbSystems/"+psqlTestDbSystemID) {
			return psqlTestResponse(request, http.StatusOK, fmt.Sprintf(`{"id":%q,"dbVersion":"16","lifecycleState":"ACTIVE"}`, psqlTestDbSystemID), ""), nil
		}
		t.Fatalf("unexpected API call: %s %s", request.Method, request.URL.Path)
		return nil, fmt.Errorf("unexpected API call")
	})}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := crud.StartDbSystem(ctx); err != nil {
		t.Fatal(err)
	}
	if !started {
		t.Fatal("DB system was not started")
	}
}

func TestPsqlDbSystemUnknownUpgradeDiff(t *testing.T) {
	for _, tc := range []struct {
		name             string
		version, allowed cty.Value
		wantReplacement  bool
	}{
		{name: "unknown version with explicit opt-in", version: cty.UnknownVal(cty.String), allowed: cty.True},
		{name: "unknown permission", version: cty.StringVal("17"), allowed: cty.UnknownVal(cty.Bool), wantReplacement: true},
		{name: "both unknown", version: cty.UnknownVal(cty.String), allowed: cty.UnknownVal(cty.Bool), wantReplacement: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			fullResource := PsqlDbSystemResource()
			resource := &schema.Resource{
				Schema:        map[string]*schema.Schema{"db_version": fullResource.Schema["db_version"], "config_id": fullResource.Schema["config_id"], "allow_major_version_upgrade": fullResource.Schema["allow_major_version_upgrade"]},
				CustomizeDiff: fullResource.CustomizeDiff,
			}
			config := terraform.NewResourceConfigShimmed(cty.ObjectVal(map[string]cty.Value{
				"db_version": tc.version, "config_id": cty.NullVal(cty.String), "allow_major_version_upgrade": tc.allowed,
			}), resource.CoreConfigSchema())
			diff, err := resource.Diff(context.Background(), &terraform.InstanceState{
				ID: psqlTestDbSystemID, Attributes: map[string]string{"db_version": "16", "config_id": "current-config", "allow_major_version_upgrade": "true"},
			}, config, nil)
			if err != nil {
				t.Fatal(err)
			}
			if diff == nil || diff.Attributes["db_version"] == nil {
				t.Fatal("expected a db_version diff")
			}
			if diff.Attributes["db_version"].NewComputed != !tc.version.IsKnown() {
				t.Fatal("unknown target version must remain unknown until apply")
			}
			if diff.RequiresNew() != tc.wantReplacement {
				t.Fatalf("Diff() replacement = %t, want %t", diff.RequiresNew(), tc.wantReplacement)
			}
		})
	}
}
