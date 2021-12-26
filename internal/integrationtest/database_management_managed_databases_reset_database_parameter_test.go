// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	managedDatabasesResetDatabaseParameterRepresentation = map[string]interface{}{
		"credentials":         acctest.RepresentationGroup{RepType: acctest.Required, Group: managedDatabasesResetDatabaseParameterCredentialsRepresentation},
		"managed_database_id": acctest.Representation{RepType: acctest.Required, Create: "ocid.database.testId"},
		"parameters":          acctest.Representation{RepType: acctest.Required, Create: []string{"open_cursors"}},
		"scope":               acctest.Representation{RepType: acctest.Required, Create: `BOTH`},
	}

	managedDatabasesResetDatabaseParameterCredentialsRepresentation = map[string]interface{}{
		"password":  acctest.Representation{RepType: acctest.Optional, Create: `system`},
		"role":      acctest.Representation{RepType: acctest.Optional, Create: `NORMAL`},
		"secret_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"user_name": acctest.Representation{RepType: acctest.Optional, Create: `system`},
	}

	ManagedDatabasesResetDatabaseParameterResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases", "test_managed_databases", acctest.Required, acctest.Create, managedDatabaseDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, vaultRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_vault_secrets", "test_secrets", acctest.Required, acctest.Create, secretDataSourceRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabasesResetDatabaseParameterResource_basic(t *testing.T) {
	t.Skip("Skip this test till Database Management service provides a better way of testing this. It requires a live managed database instance")
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabasesResetDatabaseParameterResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_management_managed_databases_reset_database_parameter.test_managed_databases_reset_database_parameter"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+acctest.GenerateResourceFromRepresentationMap("oci_database_management_managed_databases_reset_database_parameter", "test_managed_databases_reset_database_parameter", acctest.Required, acctest.Create, managedDatabasesResetDatabaseParameterRepresentation), "databasemanagement", "managedDatabasesResetDatabaseParameter", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + acctest.GenerateResourceFromRepresentationMap("oci_database_management_managed_databases_reset_database_parameter", "test_managed_databases_reset_database_parameter", acctest.Required, acctest.Create, managedDatabasesResetDatabaseParameterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "credentials.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "managed_database_id"),
				resource.TestCheckResourceAttr(resourceName, "parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scope", "BOTH"),
			),
		},
	})
}
