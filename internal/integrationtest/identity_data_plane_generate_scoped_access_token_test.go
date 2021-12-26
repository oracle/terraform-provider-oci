// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	generateScopedAccessTokenRepresentation = map[string]interface{}{
		"public_key": acctest.Representation{RepType: acctest.Required, Create: `-----BEGIN PUBLIC KEY-----MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuYNxKqyNSTPApIVh1xiR3914Q8Ex+goi8kbMUjMa/b47A12SGdh18SAsZTTkld09MGhIswyv2Eln5MQKyupf646zk0E0kxH4llpfSAtUEaa5bxRXhko5BejvimMy4hCMn+kYkzAre7CoAw97rZ96L+TgkqdtwYXl0JzE4xYwfM7OqkH9/3TIeiX4q8kVDi0CsHMGbBo4gMIIunLoEn27ej/Vm6Nbkgl8AnJaWZq8gG8y6ojDLrJhnTK4IVYZ3XYx2uxz/E5VcjMaTdWVjKVCS4F2yK9hFbL1G2KDDh8k3G7dFDFwGI6qxwidbZW7JtcXQWu0Qx0tBNdB28VlsDWZEQIDAQAB-----END PUBLIC KEY-----`},
		"scope":      acctest.Representation{RepType: acctest.Required, Create: `urn:oracle:db::id::*`},
	}

	GenerateScopedAccessTokenResourceDependencies = ""
)

// issue-routing-tag: identity_data_plane/default
func TestIdentityDataPlaneGenerateScopedAccessTokenResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDataPlaneGenerateScopedAccessTokenResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_data_plane_generate_scoped_access_token.test_generate_scoped_access_token"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+GenerateScopedAccessTokenResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_data_plane_generate_scoped_access_token", "test_generate_scoped_access_token", acctest.Required, acctest.Create, generateScopedAccessTokenRepresentation), "identitydataplane", "generateScopedAccessToken", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + GenerateScopedAccessTokenResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_data_plane_generate_scoped_access_token", "test_generate_scoped_access_token", acctest.Required, acctest.Create, generateScopedAccessTokenRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "public_key", "-----BEGIN PUBLIC KEY-----MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuYNxKqyNSTPApIVh1xiR3914Q8Ex+goi8kbMUjMa/b47A12SGdh18SAsZTTkld09MGhIswyv2Eln5MQKyupf646zk0E0kxH4llpfSAtUEaa5bxRXhko5BejvimMy4hCMn+kYkzAre7CoAw97rZ96L+TgkqdtwYXl0JzE4xYwfM7OqkH9/3TIeiX4q8kVDi0CsHMGbBo4gMIIunLoEn27ej/Vm6Nbkgl8AnJaWZq8gG8y6ojDLrJhnTK4IVYZ3XYx2uxz/E5VcjMaTdWVjKVCS4F2yK9hFbL1G2KDDh8k3G7dFDFwGI6qxwidbZW7JtcXQWu0Qx0tBNdB28VlsDWZEQIDAQAB-----END PUBLIC KEY-----"),
				resource.TestCheckResourceAttr(resourceName, "scope", "urn:oracle:db::id::*"),

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
	})
}
