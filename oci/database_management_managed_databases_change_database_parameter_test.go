// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	managedDatabasesChangeDatabaseParameterRepresentation = map[string]interface{}{
		"credentials":         RepresentationGroup{Required, managedDatabasesChangeDatabaseParameterCredentialsRepresentation},
		"managed_database_id": Representation{RepType: Required, Create: "ocid.database.testId"},
		"parameters":          RepresentationGroup{Required, managedDatabasesChangeDatabaseParameterParametersRepresentation},
		"scope":               Representation{RepType: Required, Create: `BOTH`},
	}

	managedDatabasesChangeDatabaseParameterCredentialsRepresentation = map[string]interface{}{
		"password":  Representation{RepType: Optional, Create: `system`},
		"role":      Representation{RepType: Optional, Create: `NORMAL`},
		"secret_id": Representation{RepType: Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"user_name": Representation{RepType: Optional, Create: `system`},
	}
	managedDatabasesChangeDatabaseParameterParametersRepresentation = map[string]interface{}{
		"name":           Representation{RepType: Required, Create: `open_cursors`},
		"value":          Representation{RepType: Required, Create: `305`},
		"update_comment": Representation{RepType: Required, Create: `Terraform provision of open cursors`},
	}

	ManagedDatabasesChangeDatabaseParameterResourceDependencies = GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases", "test_managed_databases", Required, Create, managedDatabaseDataSourceRepresentation) +
		GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", Required, Create, vaultRepresentation) +
		GenerateDataSourceFromRepresentationMap("oci_vault_secrets", "test_secrets", Required, Create, secretDataSourceRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabasesChangeDatabaseParameterResource_basic(t *testing.T) {
	t.Skip("Skip this test till Database Management service provides a better way of testing this. It requires a live managed database instance")
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabasesChangeDatabaseParameterResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_management_managed_databases_change_database_parameter.test_managed_databases_change_database_parameter"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+
		GenerateResourceFromRepresentationMap("oci_database_management_managed_databases_change_database_parameter", "test_managed_databases_change_database_parameter", Required, Create, managedDatabasesChangeDatabaseParameterRepresentation), "databasemanagement", "managedDatabasesChangeDatabaseParameter", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr +
				GenerateResourceFromRepresentationMap("oci_database_management_managed_databases_change_database_parameter", "test_managed_databases_change_database_parameter", Required, Create, managedDatabasesChangeDatabaseParameterRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
