// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatabaseDatabaseAutonomousExadataInfrastructureOcpuSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_exadata_infrastructure.test_autonomous_exadata_infrastructure.id}`},
	}

	DatabaseAutonomousExadataInfrastructureOcpuResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_exadata_infrastructure", "test_autonomous_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseAutonomousExadataInfrastructureRepresentation) +
		ExadataBaseDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, acctest.GetUpdatedRepresentationCopy("vcn_id", acctest.Representation{RepType: acctest.Required, Create: `${oci_core_virtual_network.t.id}`}, CoreNetworkSecurityGroupRepresentation)) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group2", acctest.Required, acctest.Create, acctest.GetUpdatedRepresentationCopy("vcn_id", acctest.Representation{RepType: acctest.Required, Create: `${oci_core_virtual_network.t.id}`}, CoreNetworkSecurityGroupRepresentation))
)

// issue-routing-tag: database/default
func TestDatabaseAutonomousExadataInfrastructureOcpuResource_basic(t *testing.T) {
	t.Skip("Skip this test as AEI and its api no longer exists.")

	httpreplay.SetScenario("TestDatabaseAutonomousExadataInfrastructureOcpuResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_database_autonomous_exadata_infrastructure_ocpu.test_autonomous_exadata_infrastructure_ocpu"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_exadata_infrastructure_ocpu", "test_autonomous_exadata_infrastructure_ocpu", acctest.Required, acctest.Create, DatabaseDatabaseAutonomousExadataInfrastructureOcpuSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousExadataInfrastructureOcpuResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_exadata_infrastructure_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "by_workload_type.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "consumed_cpu"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_cpu"),
			),
		},
	})
}
