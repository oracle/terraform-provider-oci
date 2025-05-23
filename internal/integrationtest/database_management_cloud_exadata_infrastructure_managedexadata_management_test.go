// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	_ "github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseManagementCloudExadataInfrastructureManagedexadataManagementRepresentation = map[string]interface{}{
		"cloud_exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure.id}`},
		"license_model":                   acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"enable_managedexadata":           acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
	}

	CloudExadataInfrastructureManagedexadataManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseManagementCloudExadataInfrastructureRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementCloudExadataInfrastructureManagedexadataManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementCloudExadataInfrastructureManagedexadataManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	discoveryKey := utils.GetEnvSettingWithBlankDefault("discovery_key")
	discoveryKeyStr := fmt.Sprintf("variable \"discovery_key\" { default = \"%s\" }\n", discoveryKey)

	resourceName := "oci_database_management_cloud_exadata_infrastructure_managedexadata_management.test_cloud_exadata_infrastructure_managedexadata_management"
	parentResourceName := "oci_database_management_cloud_exadata_infrastructure_managedexadata_management.test_cloud_exadata_infrastructure_managedexadata_management"
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CloudExadataInfrastructureManagedexadataManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_exadata_infrastructure_managedexadata_management", "test_cloud_exadata_infrastructure_managedexadata_management", acctest.Required, acctest.Create, DatabaseManagementCloudExadataInfrastructureManagedexadataManagementRepresentation), "databasemanagement", "cloudExadataInfrastructureManagedexadataManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// create with disable
		{
			Config: config + compartmentIdVariableStr + discoveryKeyStr +
				CloudExadataInfrastructureManagedexadataManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_exadata_infrastructure_managedexadata_management", "test_cloud_exadata_infrastructure_managedexadata_management", acctest.Optional, acctest.Create, DatabaseManagementCloudExadataInfrastructureManagedexadataManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_exadata_infrastructure_id"),
			),
		},
		// verify enable
		{
			Config: config + compartmentIdVariableStr + CloudExadataInfrastructureManagedexadataManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_cloud_exadata_infrastructure_managedexadata_management", "test_cloud_exadata_infrastructure_managedexadata_management", acctest.Optional, acctest.Update, DatabaseManagementCloudExadataInfrastructureManagedexadataManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_managedexadata", "true"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
			),
		},
	})
}
