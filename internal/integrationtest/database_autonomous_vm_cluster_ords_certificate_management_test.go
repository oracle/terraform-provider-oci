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
	certificateId                                                      = utils.GetEnvSettingWithBlankDefault("avm_certificate_id")
	certificateVariableStr                                             = fmt.Sprintf("variable \"avm_certificate_id\" { default = \"%s\" }\n", certificateId)
	DatabaseAutonomousVmClusterOrdsCertificateManagementRepresentation = map[string]interface{}{
		"autonomous_vm_cluster_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id}`},
		"certificate_generation_type": acctest.Representation{RepType: acctest.Required, Create: `BYOC`},
		"certificate_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.avm_certificate_id}`},
	}

	//DatabaseAutonomousVmClusterOrdsSystemCertificateManagementRepresentation = map[string]interface{}{
	//	"autonomous_vm_cluster_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id}`},
	//	"certificate_generation_type": acctest.Representation{RepType: acctest.Required, Create: `SYSTEM`},
	//}

	DatabaseAutonomousVmClusterOrdsManagementResourceDependencies = certificateVariableStr + DatabaseAutonomousVmClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", acctest.Required, acctest.Create, DatabaseAutonomousVmClusterRepresentation)
)

// issue-routing-tag: database/ExaCC
func TestDatabaseAutonomousVmClusterOrdsCertificateManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousVmClusterOrdsCertificateManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_vm_cluster_ords_certificate_management.test_autonomous_vm_cluster_ords_certificate_management"
	singularDatasourceName := "data.oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseAutonomousVmClusterOrdsManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster_ords_certificate_management", "test_autonomous_vm_cluster_ords_certificate_management", acctest.Required, acctest.Create, DatabaseAutonomousVmClusterOrdsCertificateManagementRepresentation), "database", "autonomousVmClusterOrdsCertificateManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousVmClusterOrdsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster_ords_certificate_management", "test_autonomous_vm_cluster_ords_certificate_management", acctest.Required, acctest.Create, DatabaseAutonomousVmClusterOrdsCertificateManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "certificate_generation_type", "BYOC"),
				resource.TestCheckResourceAttrSet(resourceName, "certificate_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousVmClusterOrdsManagementResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousVmClusterOrdsManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster_ords_certificate_management", "test_autonomous_vm_cluster_ords_certificate_management", acctest.Optional, acctest.Create, DatabaseAutonomousVmClusterOrdsCertificateManagementRepresentation),
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
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_ords_certificate_expires"),
			),
		},
	})
}
