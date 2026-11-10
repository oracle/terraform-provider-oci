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
	DistributedDatabaseDistributedAutonomousDatabasesConfigureShardingRequiredOnlyResource = DistributedDatabaseDistributedAutonomousDatabasesConfigureShardingResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_autonomous_databases_configure_sharding", "test_distributed_autonomous_databases_configure_sharding", acctest.Required, acctest.Create, DistributedDatabaseDistributedAutonomousDatabasesConfigureShardingRepresentation)

	DistributedDatabaseDistributedAutonomousDatabasesConfigureShardingRepresentation = map[string]interface{}{
		"distributed_autonomous_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_distributed_database_distributed_autonomous_database.test_distributed_autonomous_database.id}`},
		"ca_bundle_id":                       acctest.Representation{RepType: acctest.Optional, Create: `${oci_certificates_management_ca_bundle.test_ca_bundle.id}`},
		"certificate_id":                     acctest.Representation{RepType: acctest.Optional, Create: `${oci_apigateway_certificate.test_certificate.id}`},
		"is_rebalance_required":              acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	DistributedDatabaseDistributedAutonomousDatabasesConfigureShardingResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_apigateway_certificate", "test_certificate", acctest.Required, acctest.Create, ApigatewayCertificateRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_ca_bundle", "test_ca_bundle", acctest.Required, acctest.Create, caBundleRepresentationRequired) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster", "test_cloud_autonomous_vm_cluster", acctest.Required, acctest.Create, DatabaseCloudAutonomousVmClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseCloudExadataInfrastructureRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_autonomous_database", "test_distributed_autonomous_database", acctest.Required, acctest.Create, DistributedDatabaseDistributedAutonomousDatabaseRepresentation) +
		AvailabilityDomainConfig
)

// issue-routing-tag: distributed_database/default
func TestDistributedDatabaseDistributedAutonomousDatabasesConfigureShardingResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDistributedDatabaseDistributedAutonomousDatabasesConfigureShardingResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_distributed_database_distributed_autonomous_databases_configure_sharding.test_distributed_autonomous_databases_configure_sharding"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DistributedDatabaseDistributedAutonomousDatabasesConfigureShardingResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_autonomous_databases_configure_sharding", "test_distributed_autonomous_databases_configure_sharding", acctest.Optional, acctest.Create, DistributedDatabaseDistributedAutonomousDatabasesConfigureShardingRepresentation), "distributeddatabase", "distributedAutonomousDatabasesConfigureSharding", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DistributedDatabaseDistributedAutonomousDatabasesConfigureShardingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_autonomous_databases_configure_sharding", "test_distributed_autonomous_databases_configure_sharding", acctest.Required, acctest.Create, DistributedDatabaseDistributedAutonomousDatabasesConfigureShardingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "distributed_autonomous_database_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DistributedDatabaseDistributedAutonomousDatabasesConfigureShardingResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DistributedDatabaseDistributedAutonomousDatabasesConfigureShardingResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_distributed_database_distributed_autonomous_databases_configure_sharding", "test_distributed_autonomous_databases_configure_sharding", acctest.Optional, acctest.Create, DistributedDatabaseDistributedAutonomousDatabasesConfigureShardingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "ca_bundle_id"),
				resource.TestCheckResourceAttrSet(resourceName, "certificate_id"),
				resource.TestCheckResourceAttrSet(resourceName, "distributed_autonomous_database_id"),
				resource.TestCheckResourceAttr(resourceName, "is_rebalance_required", "false"),

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
