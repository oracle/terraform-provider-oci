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
	DatabaseAutonomousVmClusterSslCertificateManagementRepresentation = map[string]interface{}{
		"autonomous_vm_cluster_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id}`},
		"certificate_generation_type": acctest.Representation{RepType: acctest.Required, Create: `BYOC`},
		"certificate_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.avm_certificate_id}`},
	}

	DatabaseAutonomousVmClusterSslManagementResourceDependencies = DatabaseAutonomousVmClusterResourceDependencies + certificateVariableStr +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", acctest.Required, acctest.Create, DatabaseAutonomousVmClusterRepresentation)
)

// issue-routing-tag: database/ExaCC
func TestDatabaseAutonomousVmClusterSslCertificateManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousVmClusterSslCertificateManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_vm_cluster_ssl_certificate_management.test_autonomous_vm_cluster_ssl_certificate_management"
	singularDatasourceName := "data.oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster"
	simulateDb, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("simulate_db", "false"))

	dependencyConfig := DatabaseAutonomousVmClusterSslManagementResourceDependencies
	representation := DatabaseAutonomousVmClusterSslCertificateManagementRepresentation
	singularDatasourceRepresentation := DatabaseDatabaseAutonomousVmClusterSingularDataSourceRepresentation

	if simulateDb {
		acctest.PreCheck(t)

		sharedAutonomousVmClusterID := utils.GetEnvSettingWithBlankDefault("autonomous_vm_cluster_id")
		sharedDependencyAddresses := []string{
			"oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster",
		}
		sharedDependencyIDs, cleanup := ResolveOrCreateSharedDependenciesFromConfig(
			t,
			map[string]string{
				"oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster": sharedAutonomousVmClusterID,
			},
			config+compartmentIdVariableStr+DatabaseAVMClusterWithSingleNetworkResourceDependencies+
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", acctest.Required, acctest.Create, DatabaseAutonomousVmClusterRepresentation),
			sharedDependencyAddresses,
		)
		sharedAutonomousVmClusterID = sharedDependencyIDs["oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster"]
		if cleanup != nil {
			t.Cleanup(cleanup)
		}
		t.Logf("[SHARED_DEP_SETUP] autonomous_vm_cluster_id=%s", sharedAutonomousVmClusterID)

		dependencyConfig = certificateVariableStr + fmt.Sprintf("variable \"autonomous_vm_cluster_id\" { default = \"%s\" }\n", sharedAutonomousVmClusterID)
		representation = acctest.RepresentationCopyWithNewProperties(DatabaseAutonomousVmClusterSslCertificateManagementRepresentation, map[string]interface{}{
			"autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: sharedAutonomousVmClusterID},
		})
		singularDatasourceRepresentation = acctest.RepresentationCopyWithNewProperties(DatabaseDatabaseAutonomousVmClusterSingularDataSourceRepresentation, map[string]interface{}{
			"autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: sharedAutonomousVmClusterID},
		})
	}

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+dependencyConfig+getExaccTagDependency()+
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster_ssl_certificate_management", "test_autonomous_vm_cluster_ssl_certificate_management", acctest.Required, acctest.Create, representation), "database", "autonomousVmClusterSslCertificateManagement", t)

	t.Run("CreateAVM-Rotate-SSL-Certificate-Delete-AVM", func(t *testing.T) {
		acctest.ResourceTest(t, nil, []resource.TestStep{
			// verify Create
			{
				Config: config + compartmentIdVariableStr + dependencyConfig + getExaccTagDependency() +
					acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster_ssl_certificate_management", "test_autonomous_vm_cluster_ssl_certificate_management", acctest.Required, acctest.Create, representation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "autonomous_vm_cluster_id"),
					resource.TestCheckResourceAttr(resourceName, "certificate_generation_type", "BYOC"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate_id"),
					exaccMainResourceLog(t, "create autonomous VM cluster database SSL certificate management", resourceName, nil, &resId,
						"autonomous_vm_cluster_id", "certificate_generation_type", "certificate_id"),
				),
			},

			// delete before next Create
			{
				Config: config + compartmentIdVariableStr + dependencyConfig + getExaccTagDependency(),
			},
			// verify Create with optionals
			{
				Config: config + compartmentIdVariableStr + dependencyConfig + getExaccTagDependency() +
					acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster_ssl_certificate_management", "test_autonomous_vm_cluster_ssl_certificate_management", acctest.Optional, acctest.Create, representation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "autonomous_vm_cluster_id"),
					resource.TestCheckResourceAttr(resourceName, "certificate_generation_type", "BYOC"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate_id"),
					exaccMainResourceLog(t, "recreate autonomous VM cluster database SSL certificate management with optionals", resourceName, nil, &resId,
						"autonomous_vm_cluster_id", "certificate_generation_type", "certificate_id"),

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
			{
				Config: config + compartmentIdVariableStr + dependencyConfig + getExaccTagDependency() +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", acctest.Required, acctest.Create, singularDatasourceRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_vm_cluster_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_database_ssl_certificate_expires"),
				),
			},
		})
	})
}
