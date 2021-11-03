// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	autonomousExadataInfrastructureOcpuSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_exadata_infrastructure_id": Representation{RepType: Required, Create: `${oci_database_autonomous_exadata_infrastructure.test_autonomous_exadata_infrastructure.id}`},
	}

	AutonomousExadataInfrastructureOcpuResourceConfig = GenerateResourceFromRepresentationMap("oci_database_autonomous_exadata_infrastructure", "test_autonomous_exadata_infrastructure", Required, Create, autonomousExadataInfrastructureRepresentation) +
		ExadataBaseDependencies +
		GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Required, Create, GetUpdatedRepresentationCopy("vcn_id", Representation{RepType: Required, Create: `${oci_core_virtual_network.t.id}`}, networkSecurityGroupRepresentation)) +
		GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group2", Required, Create, GetUpdatedRepresentationCopy("vcn_id", Representation{RepType: Required, Create: `${oci_core_virtual_network.t.id}`}, networkSecurityGroupRepresentation))
)

// issue-routing-tag: database/default
func TestDatabaseAutonomousExadataInfrastructureOcpuResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousExadataInfrastructureOcpuResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_database_autonomous_exadata_infrastructure_ocpu.test_autonomous_exadata_infrastructure_ocpu"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_autonomous_exadata_infrastructure_ocpu", "test_autonomous_exadata_infrastructure_ocpu", Required, Create, autonomousExadataInfrastructureOcpuSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousExadataInfrastructureOcpuResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_exadata_infrastructure_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "by_workload_type.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "consumed_cpu"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_cpu"),
			),
		},
	})
}
