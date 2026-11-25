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
	DistributedDatabaseDistributedDatabasesAddGdsControlNodeRequiredOnlyResource = DistributedDatabaseDistributedDatabasesAddGdsControlNodeResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_databases_add_gds_control_node", "test_distributed_databases_add_gds_control_node", acctest.Required, acctest.Create, DistributedDatabaseDistributedDatabasesAddGdsControlNodeRepresentation)

	DistributedDatabaseDistributedDatabasesAddGdsControlNodeRepresentation = map[string]interface{}{
		"distributed_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_distributed_database_distributed_database.test_distributed_database.id}`},
		"public_ssh_key":          acctest.Representation{RepType: acctest.Required, Create: `ssh-rsa CCCCB3NzaC1yc2EAAAADAQABAAABAQDO7gE9zAgsAlN/RyBV+OEdB6Pv8vTwdS5K3WvqiXK9hWCtoR2siam9jZ1MOueddedwwdwdwddwwdwdwdw`},
		"subnet_id":               acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
	}

	DistributedDatabaseDistributedDatabasesAddGdsControlNodeResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseExadataInfrastructureRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_exascale_db_storage_vault", "test_exascale_db_storage_vault", acctest.Required, acctest.Create, DatabaseExascaleDbStorageVaultRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Required, acctest.Create, DatabaseVmClusterNetworkRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Required, acctest.Create, DatabaseVmClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_database", "test_distributed_database", acctest.Required, acctest.Create, DistributedDatabaseDistributedDatabaseRepresentation) +
		AvailabilityDomainConfig +
		KeyResourceDependencyConfig
)

// issue-routing-tag: distributed_database/default
func TestDistributedDatabaseDistributedDatabasesAddGdsControlNodeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDistributedDatabaseDistributedDatabasesAddGdsControlNodeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_distributed_database_distributed_databases_add_gds_control_node.test_distributed_databases_add_gds_control_node"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DistributedDatabaseDistributedDatabasesAddGdsControlNodeResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_databases_add_gds_control_node", "test_distributed_databases_add_gds_control_node", acctest.Optional, acctest.Create, DistributedDatabaseDistributedDatabasesAddGdsControlNodeRepresentation), "distributeddatabase", "distributedDatabasesAddGdsControlNode", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DistributedDatabaseDistributedDatabasesAddGdsControlNodeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_databases_add_gds_control_node", "test_distributed_databases_add_gds_control_node", acctest.Required, acctest.Create, DistributedDatabaseDistributedDatabasesAddGdsControlNodeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "distributed_database_id"),
				resource.TestCheckResourceAttr(resourceName, "public_ssh_key", "ssh-rsa CCCCB3NzaC1yc2EAAAADAQABAAABAQDO7gE9zAgsAlN/RyBV+OEdB6Pv8vTwdS5K3WvqiXK9hWCtoR2siam9jZ1MOueddedwwdwdwddwwdwdwdw"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DistributedDatabaseDistributedDatabasesAddGdsControlNodeResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DistributedDatabaseDistributedDatabasesAddGdsControlNodeResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_databases_add_gds_control_node", "test_distributed_databases_add_gds_control_node", acctest.Optional, acctest.Create, DistributedDatabaseDistributedDatabasesAddGdsControlNodeRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "distributed_database_id"),
				resource.TestCheckResourceAttr(resourceName, "public_ssh_key", "ssh-rsa CCCCB3NzaC1yc2EAAAADAQABAAABAQDO7gE9zAgsAlN/RyBV+OEdB6Pv8vTwdS5K3WvqiXK9hWCtoR2siam9jZ1MOueddedwwdwdwddwwdwdwdw"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

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
