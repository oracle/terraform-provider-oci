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
	DataSafeGenerateOnPremConnectorConfigurationRepresentation = map[string]interface{}{
		"on_prem_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_on_prem_connector.test_on_prem_connector.id}`},
		"password":             acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#1111`},
	}

	DataSafeGenerateOnPremConnectorConfigurationResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_on_prem_connector", "test_on_prem_connector", acctest.Required, acctest.Create, onPremConnectorRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeGenerateOnPremConnectorConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeGenerateOnPremConnectorConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_data_safe_generate_on_prem_connector_configuration.test_generate_on_prem_connector_configuration"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataSafeGenerateOnPremConnectorConfigurationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_generate_on_prem_connector_configuration", "test_generate_on_prem_connector_configuration", acctest.Required, acctest.Create, DataSafeGenerateOnPremConnectorConfigurationRepresentation), "datasafe", "generateOnPremConnectorConfiguration", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataSafeGenerateOnPremConnectorConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_generate_on_prem_connector_configuration", "test_generate_on_prem_connector_configuration", acctest.Required, acctest.Create, DataSafeGenerateOnPremConnectorConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "on_prem_connector_id"),
				resource.TestCheckResourceAttr(resourceName, "password", "BEstrO0ng_#1111"),

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
