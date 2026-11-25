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
	DatabaseCloudAutonomousVmClusterThalesKMSManagementRepresentation = map[string]interface{}{
		"cloud_autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id}`},
		"key_management_system":          acctest.Representation{RepType: acctest.Required, Create: `THALES`},
		"is_exclusive":                   acctest.Representation{RepType: acctest.Required, Create: `false`},
		"key_store_id":                   acctest.Representation{RepType: acctest.Required, Create: `${oci_database_key_store.test_key_store.id}`},
		"thales_hsm_port":                acctest.Representation{RepType: acctest.Required, Create: `9000`},
	}

	DatabaseCloudAutonomousVmClusterOkvKMSManagementRepresentation = map[string]interface{}{
		"cloud_autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id}`},
		"key_management_system":          acctest.Representation{RepType: acctest.Required, Create: `ORACLE_KEY_VAULT`},
		"is_exclusive":                   acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	DatabaseCloudAutonomousVmClusterOmkKMSManagementRepresentation = map[string]interface{}{
		"cloud_autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id}`},
		"key_management_system":          acctest.Representation{RepType: acctest.Required, Create: `ORACLE_MANAGED_KEYS`},
	}

	DatabaseCloudAutonomousVmClusterKMSManagementResourceDependencies = DatabaseCloudAutonomousVmClusterResourceDependencies + DatabaseThalesKeyStoreResourceConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster", "test_cloud_autonomous_vm_cluster", acctest.Required, acctest.Create, DatabaseCloudAutonomousVmClusterRepresentation)
)

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseCloudAutonomousVmClusterKMSManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseCloudAutonomousVmClusterKMSManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_cloud_autonomous_vm_cluster_kms_management.test_cloud_autonomous_vm_cluster_kms_management"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseCloudAutonomousVmClusterKMSManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster_kms_management", "test_cloud_autonomous_vm_cluster_kms_management", acctest.Optional, acctest.Create, DatabaseCloudAutonomousVmClusterThalesKMSManagementRepresentation), "database", "cloudAutonomousVmClusterKMSManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseCloudAutonomousVmClusterKMSManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster_kms_management", "test_cloud_autonomous_vm_cluster_kms_management", acctest.Required, acctest.Create, DatabaseCloudAutonomousVmClusterThalesKMSManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "key_management_system", "THALES"),
				resource.TestCheckResourceAttr(resourceName, "thales_hsm_port", "9000"),
				resource.TestCheckResourceAttr(resourceName, "is_exclusive", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseCloudAutonomousVmClusterKMSManagementResourceDependencies,
		},
		// verify enable OMK
		{
			Config: config + compartmentIdVariableStr + DatabaseCloudAutonomousVmClusterKMSManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster_kms_management", "test_cloud_autonomous_vm_cluster_kms_management", acctest.Required, acctest.Create, DatabaseCloudAutonomousVmClusterOmkKMSManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "key_management_system", "ORACLE_MANAGED_KEYS"),
			),
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseCloudAutonomousVmClusterKMSManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster_kms_management", "test_cloud_autonomous_vm_cluster_kms_management", acctest.Optional, acctest.Create, DatabaseCloudAutonomousVmClusterOkvKMSManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "is_exclusive", "false"),
				resource.TestCheckResourceAttr(resourceName, "key_management_system", "ORACLE_KEY_VAULT"),

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
