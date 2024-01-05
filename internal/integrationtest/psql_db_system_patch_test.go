// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	PsqlDbSystemInsertPatchOperationsRepresentation = map[string]interface{}{
		"operation": acctest.Representation{RepType: acctest.Required, Create: `INSERT`},
		"selection": acctest.Representation{RepType: acctest.Required, Create: `instances`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: map[string]string{"displayName": "patch-instance"}},
	}

	PsqlPatchDbSystemDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `test-terraform`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
	}

	PsqlPatchDbSystemSingularDataSourceRepresentation = map[string]interface{}{
		"db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_psql_db_systems.test_db_systems.db_system_collection.0.items.0.id}`},
	}

	PsqlDbSystemRemovePatchOperationsRepresentation = map[string]interface{}{
		"operation": acctest.Representation{RepType: acctest.Required, Create: `REMOVE`},
		"selection": acctest.Representation{RepType: acctest.Required, Create: `instances[?id == '${data.oci_psql_db_system.test_db_system.instances.1.id}']`},
	}
)

// issue-routing-tag: psql/default
func TestPsqlDbSystemPatchResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestPsqlDbSystemPatchResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_ocid")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	resourceName := "oci_psql_db_system.test_db_system"

	var resId string
	var resId2 string
	// Save TF content to Create resource with required properties.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+PsqlDbSystemResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_psql_db_system", "test_db_system", acctest.Required, acctest.Create, PsqlDbSystemRepresentation), "psql", "dbSystem", t)

	acctest.ResourceTest(t, testAccCheckPsqlDbSystemDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + PsqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_psql_db_system", "test_db_system", acctest.Required, acctest.Create, PsqlDbSystemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_version", "14"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test-terraform"),
				resource.TestCheckResourceAttr(resourceName, "network_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_details.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "shape", "PostgreSQL.VM.Standard.E4.Flex.2.32GB"),
				resource.TestCheckResourceAttr(resourceName, "storage_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "instance_count", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "storage_details.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "storage_details.0.is_regionally_durable", "false"),
				resource.TestCheckResourceAttr(resourceName, "storage_details.0.system_type", "OCI_OPTIMIZED_STORAGE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// Verify adding read replica with Patch operation
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + PsqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_psql_db_system", "test_db_system", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(PsqlDbSystemRepresentation, map[string]interface{}{
						"patch_operations": acctest.RepresentationGroup{RepType: acctest.Required, Group: PsqlDbSystemInsertPatchOperationsRepresentation},
						"instance_count":   acctest.Representation{RepType: acctest.Required, Create: `2`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_version", "14"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test-terraform"),
				resource.TestCheckResourceAttr(resourceName, "network_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_details.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "instance_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "shape", "PostgreSQL.VM.Standard.E4.Flex.2.32GB"),
				resource.TestCheckResourceAttr(resourceName, "storage_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "storage_details.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "storage_details.0.is_regionally_durable", "false"),
				resource.TestCheckResourceAttr(resourceName, "storage_details.0.system_type", "OCI_OPTIMIZED_STORAGE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		//Verify after patch operation is removed
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + PsqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_psql_db_system", "test_db_system", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(PsqlDbSystemRepresentation, map[string]interface{}{
						"instance_count": acctest.Representation{RepType: acctest.Required, Create: `2`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_version", "14"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test-terraform"),
				resource.TestCheckResourceAttr(resourceName, "instance_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "network_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_details.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "shape", "PostgreSQL.VM.Standard.E4.Flex.2.32GB"),
				resource.TestCheckResourceAttr(resourceName, "storage_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "storage_details.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "storage_details.0.is_regionally_durable", "false"),
				resource.TestCheckResourceAttr(resourceName, "storage_details.0.system_type", "OCI_OPTIMIZED_STORAGE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// Verify removing read replica with Patch operation
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + PsqlDbSystemResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_psql_db_systems", "test_db_systems", acctest.Optional, acctest.Create, PsqlPatchDbSystemDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_psql_db_system", "test_db_system", acctest.Optional, acctest.Create, PsqlPatchDbSystemSingularDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_psql_db_system", "test_db_system", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(PsqlDbSystemRepresentation, map[string]interface{}{
						"patch_operations": acctest.RepresentationGroup{RepType: acctest.Required, Group: PsqlDbSystemRemovePatchOperationsRepresentation},
						"instance_count":   acctest.Representation{RepType: acctest.Required, Create: `1`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_version", "14"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test-terraform"),
				resource.TestCheckResourceAttr(resourceName, "instance_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "network_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_details.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "shape", "PostgreSQL.VM.Standard.E4.Flex.2.32GB"),
				resource.TestCheckResourceAttr(resourceName, "storage_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "storage_details.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "storage_details.0.is_regionally_durable", "false"),
				resource.TestCheckResourceAttr(resourceName, "storage_details.0.system_type", "OCI_OPTIMIZED_STORAGE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		//Verify after patch operation is removed
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + PsqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_psql_db_system", "test_db_system", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(PsqlDbSystemRepresentation, map[string]interface{}{
						"instance_count": acctest.Representation{RepType: acctest.Required, Create: `1`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_version", "14"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "test-terraform"),
				resource.TestCheckResourceAttr(resourceName, "network_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_details.0.subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "shape", "PostgreSQL.VM.Standard.E4.Flex.2.32GB"),
				resource.TestCheckResourceAttr(resourceName, "storage_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "storage_details.0.availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "storage_details.0.is_regionally_durable", "false"),
				resource.TestCheckResourceAttr(resourceName, "storage_details.0.system_type", "OCI_OPTIMIZED_STORAGE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("PsqlDbSystemPatch") {
		resource.AddTestSweepers("PsqlDbSystemPatch", &resource.Sweeper{
			Name:         "PsqlDbSystemPatch",
			Dependencies: acctest.DependencyGraph["dbSystem"],
			F:            sweepPsqlDbSystemResource,
		})
	}
}
