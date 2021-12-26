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
	managedDatabasesChangeDatabaseParameterRepresentation = map[string]interface{}{
		"credentials":         acctest.RepresentationGroup{RepType: acctest.Required, Group: managedDatabasesChangeDatabaseParameterCredentialsRepresentation},
		"managed_database_id": acctest.Representation{RepType: acctest.Required, Create: "ocid.database.testId"},
		"parameters":          acctest.RepresentationGroup{RepType: acctest.Required, Group: managedDatabasesChangeDatabaseParameterParametersRepresentation},
		"scope":               acctest.Representation{RepType: acctest.Required, Create: `BOTH`},
	}

	managedDatabasesChangeDatabaseParameterCredentialsRepresentation = map[string]interface{}{
		"password":  acctest.Representation{RepType: acctest.Optional, Create: `system`},
		"role":      acctest.Representation{RepType: acctest.Optional, Create: `NORMAL`},
		"secret_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"user_name": acctest.Representation{RepType: acctest.Optional, Create: `system`},
	}
	managedDatabasesChangeDatabaseParameterParametersRepresentation = map[string]interface{}{
		"name":           acctest.Representation{RepType: acctest.Required, Create: `open_cursors`},
		"value":          acctest.Representation{RepType: acctest.Required, Create: `305`},
		"update_comment": acctest.Representation{RepType: acctest.Required, Create: `Terraform provision of open cursors`},
	}

	ManagedDatabasesChangeDatabaseParameterResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases", "test_managed_databases", acctest.Required, acctest.Create, managedDatabaseDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, vaultRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_vault_secrets", "test_secrets", acctest.Required, acctest.Create, secretDataSourceRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabasesChangeDatabaseParameterResource_basic(t *testing.T) {
	t.Skip("Skip this test till Database Management service provides a better way of testing this. It requires a live managed database instance")
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabasesChangeDatabaseParameterResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_management_managed_databases_change_database_parameter.test_managed_databases_change_database_parameter"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_managed_databases_change_database_parameter", "test_managed_databases_change_database_parameter", acctest.Required, acctest.Create, managedDatabasesChangeDatabaseParameterRepresentation), "databasemanagement", "managedDatabasesChangeDatabaseParameter", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_managed_databases_change_database_parameter", "test_managed_databases_change_database_parameter", acctest.Required, acctest.Create, managedDatabasesChangeDatabaseParameterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "credentials.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "managed_database_id"),
				resource.TestCheckResourceAttr(resourceName, "parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.name", "open_cursors"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.value", "305"),
				resource.TestCheckResourceAttr(resourceName, "scope", "BOTH"),
			),
		},
	})
}
