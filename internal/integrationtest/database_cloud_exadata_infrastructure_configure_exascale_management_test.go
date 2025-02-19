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
	DatabaseCloudExadataInfrastructureConfigureExascaleManagementRepresentation = map[string]interface{}{
		"cloud_exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure.id}`},
		"total_storage_in_gbs":            acctest.Representation{RepType: acctest.Required, Create: `4096`},
	}

	DatabaseCloudExadataInfrastructureConfigureExascaleManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseCloudExadataInfrastructureRepresentation) +
		AvailabilityDomainConfig
)

// issue-routing-tag: database/ExaCS
func TestDatabaseCloudExadataInfrastructureConfigureExascaleManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseCloudExadataInfrastructureConfigureExascaleManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_cloud_exadata_infrastructure_configure_exascale_management.test_cloud_exadata_infrastructure_configure_exascale_management"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseCloudExadataInfrastructureConfigureExascaleManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure_configure_exascale_management", "test_cloud_exadata_infrastructure_configure_exascale_management", acctest.Required, acctest.Create, DatabaseCloudExadataInfrastructureConfigureExascaleManagementRepresentation), "database", "cloudExadataInfrastructureConfigureExascaleManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseCloudExadataInfrastructureConfigureExascaleManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure_configure_exascale_management", "test_cloud_exadata_infrastructure_configure_exascale_management", acctest.Required, acctest.Create, DatabaseCloudExadataInfrastructureConfigureExascaleManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_exadata_infrastructure_id"),
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
