// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseManagementExternalExadataInfrastructureExadataManagementRepresentation = map[string]interface{}{
		"external_exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_external_exadata_infrastructure.test_external_exadata_infrastructure.id}`},
		"license_model":                      acctest.Representation{RepType: acctest.Required, Create: `LICENSE_INCLUDED`},
		"enable_exadata":                     acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
	}

	ExternalExadataInfrastructureExadataManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_exadata_infrastructure", "test_external_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseManagementExternalExadataInfrastructureRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExternalExadataInfrastructureExadataManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementExternalExadataInfrastructureExadataManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	dbSystemId := utils.GetEnvSettingWithBlankDefault("exa_db_system_id")
	dbSystemIdStr := fmt.Sprintf("variable \"db_system_id\" { default = \"%s\" }\n", dbSystemId)

	resourceName := "oci_database_management_external_exadata_infrastructure_exadata_management.test_external_exadata_infrastructure_exadata_management"
	parentResourceName := "oci_database_management_external_exadata_infrastructure_exadata_management.test_external_exadata_infrastructure_exadata_management"
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ExternalExadataInfrastructureExadataManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_exadata_infrastructure_exadata_management", "test_external_exadata_infrastructure_exadata_management", acctest.Required, acctest.Create, DatabaseManagementExternalExadataInfrastructureExadataManagementRepresentation), "databasemanagement", "externalExadataInfrastructureExadataManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// create with disable
		{
			Config: config + dbSystemIdStr + compartmentIdVariableStr + ExternalExadataInfrastructureExadataManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_exadata_infrastructure_exadata_management", "test_external_exadata_infrastructure_exadata_management", acctest.Required, acctest.Create, DatabaseManagementExternalExadataInfrastructureExadataManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
			),
		},
		// verify enable
		{
			Config: config + dbSystemIdStr + compartmentIdVariableStr + ExternalExadataInfrastructureExadataManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_exadata_infrastructure_exadata_management", "test_external_exadata_infrastructure_exadata_management", acctest.Required, acctest.Update, DatabaseManagementExternalExadataInfrastructureExadataManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_exadata", "true"),
			),
		},
	})
}
