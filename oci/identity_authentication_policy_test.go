// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	AuthenticationPolicyRequiredOnlyResource = AuthenticationPolicyResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_identity_authentication_policy", "test_authentication_policy", Required, Create, authenticationPolicyRepresentation)

	AuthenticationPolicyResourceConfig = AuthenticationPolicyResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_identity_authentication_policy", "test_authentication_policy", Optional, Update, authenticationPolicyRepresentation)

	authenticationPolicySingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.tenancy_ocid}`},
	}

	authenticationPolicyRepresentation = map[string]interface{}{
		"compartment_id":  Representation{RepType: Required, Create: `${var.tenancy_ocid}`},
		"network_policy":  RepresentationGroup{Optional, authenticationPolicyNetworkPolicyRepresentation},
		"password_policy": RepresentationGroup{Optional, authenticationPolicyPasswordPolicyRepresentation},
	}
	authenticationPolicyNetworkPolicyRepresentation = map[string]interface{}{
		"network_source_ids": Representation{RepType: Optional, Create: []string{`${oci_identity_network_source.test_network_source.id}`}, Update: []string{}},
	}
	authenticationPolicyPasswordPolicyRepresentation = map[string]interface{}{
		"is_lowercase_characters_required": Representation{RepType: Optional, Create: `true`, Update: `false`},
		"is_numeric_characters_required":   Representation{RepType: Optional, Create: `true`, Update: `false`},
		"is_special_characters_required":   Representation{RepType: Optional, Create: `true`, Update: `false`},
		"is_uppercase_characters_required": Representation{RepType: Optional, Create: `true`, Update: `false`},
		"is_username_containment_allowed":  Representation{RepType: Optional, Create: `false`},
		"minimum_password_length":          Representation{RepType: Optional, Create: `11`, Update: `8`},
	}

	AuthenticationPolicyResourceDependencies = GenerateResourceFromRepresentationMap("oci_identity_network_source", "test_network_source", Required, Create, networkSourceRepresentation)
)

// issue-routing-tag: identity/default
func TestIdentityAuthenticationPolicyResource_basic(t *testing.T) {
	t.Skip("Skip this test as this might lock users out of the tenancy and they will not be able to login through console")

	httpreplay.SetScenario("TestIdentityAuthenticationPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_identity_authentication_policy.test_authentication_policy"

	singularDatasourceName := "data.oci_identity_authentication_policy.test_authentication_policy"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+AuthenticationPolicyResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_identity_authentication_policy", "test_authentication_policy", Optional, Create, authenticationPolicyRepresentation), "identity", "authenticationPolicy", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AuthenticationPolicyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_identity_authentication_policy", "test_authentication_policy", Optional, Create, authenticationPolicyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AuthenticationPolicyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AuthenticationPolicyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_identity_authentication_policy", "test_authentication_policy", Optional, Create, authenticationPolicyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "network_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_policy.0.network_source_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "password_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "password_policy.0.is_lowercase_characters_required", "true"),
				resource.TestCheckResourceAttr(resourceName, "password_policy.0.is_numeric_characters_required", "true"),
				resource.TestCheckResourceAttr(resourceName, "password_policy.0.is_special_characters_required", "true"),
				resource.TestCheckResourceAttr(resourceName, "password_policy.0.is_uppercase_characters_required", "true"),
				resource.TestCheckResourceAttr(resourceName, "password_policy.0.is_username_containment_allowed", "false"),
				resource.TestCheckResourceAttr(resourceName, "password_policy.0.minimum_password_length", "11"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + AuthenticationPolicyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_identity_authentication_policy", "test_authentication_policy", Optional, Update, authenticationPolicyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "network_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_policy.0.network_source_ids.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "password_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "password_policy.0.is_lowercase_characters_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "password_policy.0.is_numeric_characters_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "password_policy.0.is_special_characters_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "password_policy.0.is_uppercase_characters_required", "false"),
				resource.TestCheckResourceAttr(resourceName, "password_policy.0.is_username_containment_allowed", "false"),
				resource.TestCheckResourceAttr(resourceName, "password_policy.0.minimum_password_length", "8"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_identity_authentication_policy", "test_authentication_policy", Required, Create, authenticationPolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + AuthenticationPolicyResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_policy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_policy.0.network_source_ids.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "password_policy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "password_policy.0.is_lowercase_characters_required", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "password_policy.0.is_numeric_characters_required", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "password_policy.0.is_special_characters_required", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "password_policy.0.is_uppercase_characters_required", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "password_policy.0.is_username_containment_allowed", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "password_policy.0.minimum_password_length", "8"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + AuthenticationPolicyResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
