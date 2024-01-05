// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

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

	//DatabaseAutonomousVmClusterSslCertificateManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_apigateway_certificate", "test_certificate", acctest.Required, acctest.Create, certificateRepresentation) +
	//	GenerateResourceFromRepresentationMap("oci_apigateway_certificate", "test_certificate", Required, Create, apiGatewaycertificateRepresentation) +
	//	acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_ca_bundle", "test_ca_bundle", acctest.Required, acctest.Create, caBundleRepresentation) +
	//	acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate_authority", "test_certificate_authority", acctest.Required, acctest.Create, certificateAuthorityRepresentation) +
	//	acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", acctest.Required, acctest.Create, autonomousVmClusterRepresentation) +
	//	acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Required, acctest.Create, exadataInfrastructureRepresentation) +
	//	acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Required, acctest.Create, vmClusterNetworkRepresentation) +
	//	KeyResourceDependencyConfig +
	//	acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, bucketRepresentation) +
	//	GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation)
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

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseAutonomousVmClusterSslManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster_ssl_certificate_management", "test_autonomous_vm_cluster_ssl_certificate_management", acctest.Required, acctest.Create, DatabaseAutonomousVmClusterSslCertificateManagementRepresentation), "database", "autonomousVmClusterSslCertificateManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousVmClusterSslManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster_ssl_certificate_management", "test_autonomous_vm_cluster_ssl_certificate_management", acctest.Required, acctest.Create, DatabaseAutonomousVmClusterSslCertificateManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "certificate_generation_type", "BYOC"),
				resource.TestCheckResourceAttrSet(resourceName, "certificate_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousVmClusterSslManagementResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousVmClusterSslManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster_ssl_certificate_management", "test_autonomous_vm_cluster_ssl_certificate_management", acctest.Optional, acctest.Create, DatabaseAutonomousVmClusterSslCertificateManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "certificate_generation_type", "BYOC"),
				resource.TestCheckResourceAttrSet(resourceName, "certificate_id"),

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
			Config: config + compartmentIdVariableStr + DatabaseAutonomousVmClusterOrdsManagementResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", acctest.Required, acctest.Create, DatabaseDatabaseAutonomousVmClusterSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_database_ssl_certificate_expires"),
			),
		},
	})
}
