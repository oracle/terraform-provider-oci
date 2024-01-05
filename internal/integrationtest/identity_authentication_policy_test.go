// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityAuthenticationPolicyRequiredOnlyResource = IdentityAuthenticationPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_authentication_policy", "test_authentication_policy", acctest.Required, acctest.Create, IdentityAuthenticationPolicyRepresentation)

	IdentityAuthenticationPolicyResourceConfig = IdentityAuthenticationPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_authentication_policy", "test_authentication_policy", acctest.Optional, acctest.Update, IdentityAuthenticationPolicyRepresentation)

	IdentityIdentityAuthenticationPolicySingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
	}

	IdentityAuthenticationPolicyRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"network_policy":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityAuthenticationPolicyNetworkPolicyRepresentation},
		"password_policy": acctest.RepresentationGroup{RepType: acctest.Optional, Group: IdentityAuthenticationPolicyPasswordPolicyRepresentation},
	}
	IdentityAuthenticationPolicyNetworkPolicyRepresentation = map[string]interface{}{
		"network_source_ids": acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_identity_network_source.test_network_source.id}`}, Update: []string{}},
	}
	IdentityAuthenticationPolicyPasswordPolicyRepresentation = map[string]interface{}{
		"is_lowercase_characters_required": acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
		"is_numeric_characters_required":   acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
		"is_special_characters_required":   acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
		"is_uppercase_characters_required": acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
		"is_username_containment_allowed":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"minimum_password_length":          acctest.Representation{RepType: acctest.Optional, Create: `11`, Update: `8`},
	}

	IdentityAuthenticationPolicyResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_identity_network_source", "test_network_source", acctest.Required, acctest.Create, IdentityNetworkSourceRepresentation)
)

// issue-routing-tag: identity/default
func TestIdentityAuthenticationPolicyResource_basic(t *testing.T) {
	t.Skip("Skip this test as this might lock users out of the tenancy and they will not be able to login through console")

	httpreplay.SetScenario("TestIdentityAuthenticationPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_identity_authentication_policy.test_authentication_policy"

	singularDatasourceName := "data.oci_identity_authentication_policy.test_authentication_policy"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityAuthenticationPolicyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_authentication_policy", "test_authentication_policy", acctest.Optional, acctest.Create, IdentityAuthenticationPolicyRepresentation), "identity", "authenticationPolicy", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityAuthenticationPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_authentication_policy", "test_authentication_policy", acctest.Optional, acctest.Create, IdentityAuthenticationPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + IdentityAuthenticationPolicyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + IdentityAuthenticationPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_authentication_policy", "test_authentication_policy", acctest.Optional, acctest.Create, IdentityAuthenticationPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + IdentityAuthenticationPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_authentication_policy", "test_authentication_policy", acctest.Optional, acctest.Update, IdentityAuthenticationPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_authentication_policy", "test_authentication_policy", acctest.Required, acctest.Create, IdentityIdentityAuthenticationPolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityAuthenticationPolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
		// verify resource import
		{
			Config:                  config + IdentityAuthenticationPolicyRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
