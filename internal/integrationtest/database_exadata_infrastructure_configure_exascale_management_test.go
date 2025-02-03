// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseExadataInfrastructureConfigureExascaleManagementRepresentation = map[string]interface{}{
		"exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"total_storage_in_gbs":      acctest.Representation{RepType: acctest.Required, Create: `4096`},
	}

	DatabaseExadataInfrastructureConfigureExascaleManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Optional, acctest.Update,
		acctest.RepresentationCopyWithNewProperties(exadataInfrastructureActivateRepresentation, map[string]interface{}{
			"activation_file":    acctest.Representation{RepType: acctest.Optional, Update: activationFilePath},
			"maintenance_window": acctest.RepresentationGroup{RepType: acctest.Optional, Group: exadataInfrastructureMaintenanceWindowRepresentationComplete},
		})) + DefinedTagsDependencies
)

// issue-routing-tag: database/ExaCC
func TestDatabaseExadataInfrastructureConfigureExascaleManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExadataInfrastructureConfigureExascaleManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_exadata_infrastructure_configure_exascale_management.test_exadata_infrastructure_configure_exascale_management"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseExadataInfrastructureConfigureExascaleManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure_configure_exascale_management", "test_exadata_infrastructure_configure_exascale_management", acctest.Required, acctest.Create, DatabaseExadataInfrastructureConfigureExascaleManagementRepresentation), "database", "exadataInfrastructureConfigureExascaleManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{

			Config: config + compartmentIdVariableStr + ExadataInfrastructureResourceActivateDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(exadataInfrastructureActivateRepresentation, map[string]interface{}{
						"activation_file":    acctest.Representation{RepType: acctest.Optional, Update: activationFilePath},
						"maintenance_window": acctest.RepresentationGroup{RepType: acctest.Optional, Group: exadataInfrastructureMaintenanceWindowRepresentationComplete},
					})) + acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure_configure_exascale_management", "test_exadata_infrastructure_configure_exascale_management", acctest.Required, acctest.Create, DatabaseExadataInfrastructureConfigureExascaleManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "total_storage_in_gbs", "4096"),

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
