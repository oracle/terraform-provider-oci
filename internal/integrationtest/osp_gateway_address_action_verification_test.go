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
	OspGatewayAddressActionVerificationRequiredOnlyResource = OspGatewayAddressActionVerificationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_osp_gateway_address_action_verification", "test_address_action_verification", acctest.Required, acctest.Create, OspGatewayAddressActionVerificationRepresentation)

	OspGatewayAddressActionVerificationRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"osp_home_region": acctest.Representation{RepType: acctest.Required, Create: `${var.home_region}`},
		"address_key":     acctest.Representation{RepType: acctest.Optional, Create: `null`},
		"city":            acctest.Representation{RepType: acctest.Optional, Create: `Zapopan`},
		"country":         acctest.Representation{RepType: acctest.Optional, Create: `MX`},
		"county":          acctest.Representation{RepType: acctest.Optional, Create: `Zapopan`},
		"line1":           acctest.Representation{RepType: acctest.Optional, Create: `Blvd Puerta de Hierro 5065`},
		"line2":           acctest.Representation{RepType: acctest.Optional, Create: `Col Puerta de Hierro`},
		"line3":           acctest.Representation{RepType: acctest.Optional, Create: `null`},
		"postal_code":     acctest.Representation{RepType: acctest.Optional, Create: `45116`},
		"state":           acctest.Representation{RepType: acctest.Optional, Create: `JAL`},
	}

	OspGatewayAddressActionVerificationResourceDependencies = ""
)

// issue-routing-tag: osp_gateway/default
func TestOspGatewayAddressActionVerificationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOspGatewayAddressActionVerificationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	homeRegion := utils.GetEnvSettingWithBlankDefault("region")
	regionVariableStr := fmt.Sprintf("variable \"home_region\" { default = \"%s\" }\n", homeRegion)

	resourceName := "oci_osp_gateway_address_action_verification.test_address_action_verification"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OspGatewayAddressActionVerificationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_osp_gateway_address_action_verification", "test_address_action_verification", acctest.Optional, acctest.Create, OspGatewayAddressActionVerificationRepresentation), "ospgateway", "addressActionVerification", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: config + compartmentIdVariableStr + regionVariableStr + OspGatewayAddressActionVerificationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_osp_gateway_address_action_verification", "test_address_action_verification", acctest.Optional, acctest.Create, OspGatewayAddressActionVerificationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "address.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "city", "Zapopan"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "country", "MX"),
				resource.TestCheckResourceAttr(resourceName, "county", "Zapopan"),
				resource.TestCheckResourceAttr(resourceName, "line1", "Blvd Puerta de Hierro 5065"),
				resource.TestCheckResourceAttr(resourceName, "line2", "Col Puerta de Hierro"),
				resource.TestCheckResourceAttr(resourceName, "line3", "null"),
				resource.TestCheckResourceAttr(resourceName, "osp_home_region", homeRegion),
				resource.TestCheckResourceAttr(resourceName, "postal_code", "45116"),
				resource.TestCheckResourceAttr(resourceName, "state", "JAL"),

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
