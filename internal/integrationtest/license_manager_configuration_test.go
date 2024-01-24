// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	LicenseManagerConfigurationResourceConfig = LicenseManagerConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_license_manager_configuration", "test_configuration", acctest.Optional, acctest.Update, LicenseManagerConfigurationRepresentation)

	LicenseManagerLicenseManagerConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
	}

	LicenseManagerConfigurationRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"email_ids":      acctest.Representation{RepType: acctest.Required, Create: []string{`test_create@oracle.com`}, Update: []string{`test_update@oracle.com`}},
	}

	LicenseManagerConfigurationResourceDependencies = ""
)

// issue-routing-tag: license_manager/default
func TestLicenseManagerConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLicenseManagerConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_license_manager_configuration.test_configuration"

	singularDatasourceName := "data.oci_license_manager_configuration.test_configuration"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+LicenseManagerConfigurationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_license_manager_configuration", "test_configuration", acctest.Required, acctest.Create, LicenseManagerConfigurationRepresentation), "licensemanager", "configuration", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify updates to updatable parameters
		{

			Config: config + compartmentIdVariableStr + LicenseManagerConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_license_manager_configuration", "test_configuration", acctest.Optional, acctest.Update, LicenseManagerConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "email_ids.#", "1"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if compartmentId != resId {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		//verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_license_manager_configuration", "test_configuration", acctest.Required, acctest.Create, LicenseManagerLicenseManagerConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + LicenseManagerConfigurationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "email_ids.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		//verify resource import
		{
			Config:                  config + LicenseManagerConfigurationResourceConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
