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
	DatabaseThalesKeyStoreResourceConfig = DatabaseKeyStoreWithoutTagsResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_key_store", "test_key_store", acctest.Optional, acctest.Create, DatabaseThalesKeyStoreRepresentation)

	DatabaseAutonomousVmClusterThalesKMSManagementRepresentation = map[string]interface{}{
		"autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id}`},
		"key_management_system":    acctest.Representation{RepType: acctest.Required, Create: `THALES`},
		"is_exclusive":             acctest.Representation{RepType: acctest.Required, Create: `false`},
		"key_store_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_database_key_store.test_key_store.id}`},
		"thales_hsm_port":          acctest.Representation{RepType: acctest.Required, Create: `9000`},
	}

	DatabaseAutonomousVmClusterOkvKMSManagementRepresentation = map[string]interface{}{
		"autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id}`},
		"key_management_system":    acctest.Representation{RepType: acctest.Required, Create: `ORACLE_KEY_VAULT`},
		"is_exclusive":             acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}

	DatabaseAutonomousVmClusterOmkKMSManagementRepresentation = map[string]interface{}{
		"autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id}`},
		"key_management_system":    acctest.Representation{RepType: acctest.Required, Create: `ORACLE_MANAGED_KEYS`},
	}

	DatabaseAutonomousVmClusterKMSManagementResourceDependencies = DatabaseAutonomousVmClusterResourceDependencies + DatabaseThalesKeyStoreResourceConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", acctest.Required, acctest.Create, DatabaseAutonomousVmClusterRepresentation)
)

// issue-routing-tag: database/ExaCC
func TestDatabaseAutonomousVmClusterKMSManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousVmClusterKMSManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_vm_cluster_kms_management.test_autonomous_vm_cluster_kms_management"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseAutonomousVmClusterKMSManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster_kms_management", "test_autonomous_vm_cluster_kms_management", acctest.Required, acctest.Create, DatabaseAutonomousVmClusterThalesKMSManagementRepresentation), "database", "autonomousVmClusterKMSManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousVmClusterKMSManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster_kms_management", "test_autonomous_vm_cluster_kms_management", acctest.Required, acctest.Create, DatabaseAutonomousVmClusterThalesKMSManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "key_management_system", "THALES"),
				resource.TestCheckResourceAttr(resourceName, "thales_hsm_port", "9000"),
				resource.TestCheckResourceAttr(resourceName, "is_exclusive", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_id"),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousVmClusterKMSManagementResourceDependencies,
		},
		// verify enable OMK
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousVmClusterKMSManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster_kms_management", "test_autonomous_vm_cluster_kms_management", acctest.Required, acctest.Create, DatabaseAutonomousVmClusterOmkKMSManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "key_management_system", "ORACLE_MANAGED_KEYS"),
			),
		},
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousVmClusterKMSManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster_kms_management", "test_autonomous_vm_cluster_kms_management", acctest.Optional, acctest.Create, DatabaseAutonomousVmClusterOkvKMSManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_vm_cluster_id"),
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
