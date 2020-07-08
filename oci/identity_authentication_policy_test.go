// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
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
		generateResourceFromRepresentationMap("oci_identity_authentication_policy", "test_authentication_policy", Required, Create, authenticationPolicyRepresentation)

	AuthenticationPolicyResourceConfig = AuthenticationPolicyResourceDependencies +
		generateResourceFromRepresentationMap("oci_identity_authentication_policy", "test_authentication_policy", Optional, Update, authenticationPolicyRepresentation)

	authenticationPolicySingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.tenancy_ocid}`},
	}

	authenticationPolicyRepresentation = map[string]interface{}{
		"compartment_id":  Representation{repType: Required, create: `${var.tenancy_ocid}`},
		"password_policy": RepresentationGroup{Optional, authenticationPolicyPasswordPolicyRepresentation},
	}
	authenticationPolicyPasswordPolicyRepresentation = map[string]interface{}{
		"is_lowercase_characters_required": Representation{repType: Optional, create: `true`, update: `false`},
		"is_numeric_characters_required":   Representation{repType: Optional, create: `true`, update: `false`},
		"is_special_characters_required":   Representation{repType: Optional, create: `true`, update: `false`},
		"is_uppercase_characters_required": Representation{repType: Optional, create: `true`, update: `false`},
		"is_username_containment_allowed":  Representation{repType: Optional, create: `false`},
		"minimum_password_length":          Representation{repType: Optional, create: `11`, update: `8`},
	}

	AuthenticationPolicyResourceDependencies = ""
)

func TestIdentityAuthenticationPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityAuthenticationPolicyResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := getEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_identity_authentication_policy.test_authentication_policy"

	singularDatasourceName := "data.oci_identity_authentication_policy.test_authentication_policy"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + AuthenticationPolicyResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_authentication_policy", "test_authentication_policy", Optional, Create, authenticationPolicyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + AuthenticationPolicyResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + AuthenticationPolicyResourceDependencies +
					generateResourceFromRepresentationMap("oci_identity_authentication_policy", "test_authentication_policy", Optional, Create, authenticationPolicyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "password_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "password_policy.0.is_lowercase_characters_required", "true"),
					resource.TestCheckResourceAttr(resourceName, "password_policy.0.is_numeric_characters_required", "true"),
					resource.TestCheckResourceAttr(resourceName, "password_policy.0.is_special_characters_required", "true"),
					resource.TestCheckResourceAttr(resourceName, "password_policy.0.is_uppercase_characters_required", "true"),
					resource.TestCheckResourceAttr(resourceName, "password_policy.0.is_username_containment_allowed", "false"),
					resource.TestCheckResourceAttr(resourceName, "password_policy.0.minimum_password_length", "11"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
					generateResourceFromRepresentationMap("oci_identity_authentication_policy", "test_authentication_policy", Optional, Update, authenticationPolicyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
					resource.TestCheckResourceAttr(resourceName, "password_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "password_policy.0.is_lowercase_characters_required", "false"),
					resource.TestCheckResourceAttr(resourceName, "password_policy.0.is_numeric_characters_required", "false"),
					resource.TestCheckResourceAttr(resourceName, "password_policy.0.is_special_characters_required", "false"),
					resource.TestCheckResourceAttr(resourceName, "password_policy.0.is_uppercase_characters_required", "false"),
					resource.TestCheckResourceAttr(resourceName, "password_policy.0.is_username_containment_allowed", "false"),
					resource.TestCheckResourceAttr(resourceName, "password_policy.0.minimum_password_length", "8"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_identity_authentication_policy", "test_authentication_policy", Required, Create, authenticationPolicySingularDataSourceRepresentation) +
					compartmentIdVariableStr + AuthenticationPolicyResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", tenancyId),
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
		},
	})
}
