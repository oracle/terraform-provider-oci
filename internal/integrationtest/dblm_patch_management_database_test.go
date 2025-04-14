// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	oci_dblm "github.com/oracle/oci-go-sdk/v65/dblm"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DblmPatchManagementDatabaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_release": acctest.Representation{RepType: acctest.Optional, Create: `19.0.0.0.0`},
		"database_type":    acctest.Representation{RepType: acctest.Optional, Create: `SI`},
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"drifter_patch_id": acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"image_compliance": acctest.Representation{RepType: acctest.Optional, Create: `NOT_SUBSCRIBED`},
		"image_id":         acctest.Representation{RepType: acctest.Optional, Create: `imageId`},
		"severity_type":    acctest.Representation{RepType: acctest.Optional, Create: []oci_dblm.ResourcesSeveritiesEnum{`CRITICAL`}},
		"state":            acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
	}

	DblmPatchManagementDatabaseResourceConfig = ""
)

// issue-routing-tag: dblm/default
func TestDblmPatchManagementDatabaseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDblmPatchManagementDatabaseResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	// compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	// tenantId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	patchTypePlatformConfigurationId := utils.GetEnvSettingWithBlankDefault("test_patch_type_platform_configuration_id")
	productPlatformConfigurationId := utils.GetEnvSettingWithBlankDefault("test_product_platform_configuration_id")

	platformConfigurationsStr := fmt.Sprintf(
		"variable \"patch_type_platform_configuration_id\" { default = \"%s\" }\n"+
			"variable \"product_platform_configuration_id\" { default = \"%s\" }\n"+
			"variable \"compartment_id\" { default = \"%s\" }\n",
		patchTypePlatformConfigurationId, productPlatformConfigurationId, compartmentId)

	datasourceName := "data.oci_dblm_patch_management_databases.test_dblm_patch_management_databases"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dblm_patch_management_databases", "test_dblm_patch_management_databases", acctest.Optional, acctest.Create, DblmPatchManagementDatabaseDataSourceRepresentation) +
				platformConfigurationsStr + DblmPatchManagementDatabaseResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "database_release", "19.0.0.0.0"),
				resource.TestCheckResourceAttr(datasourceName, "database_type", "SI"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttrSet(datasourceName, "patch_databases_collection.#"),
			),
		},
	})
}
