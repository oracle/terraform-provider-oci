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
	managedDatabasesResetDatabaseParameterRepresentation = map[string]interface{}{
		"credentials":         RepresentationGroup{Required, managedDatabasesResetDatabaseParameterCredentialsRepresentation},
		"managed_database_id": Representation{repType: Required, create: "ocid.database.testId"},
		"parameters":          Representation{repType: Required, create: []string{"open_cursors"}},
		"scope":               Representation{repType: Required, create: `BOTH`},
	}

	managedDatabasesResetDatabaseParameterCredentialsRepresentation = map[string]interface{}{
		"password":  Representation{repType: Optional, create: `system`},
		"role":      Representation{repType: Optional, create: `NORMAL`},
		"secret_id": Representation{repType: Optional, create: `${oci_vault_secret.test_secret.id}`},
		"user_name": Representation{repType: Optional, create: `system`},
	}

	ManagedDatabasesResetDatabaseParameterResourceDependencies = generateDataSourceFromRepresentationMap("oci_database_management_managed_databases", "test_managed_databases", Required, Create, managedDatabaseDataSourceRepresentation) +
		generateResourceFromRepresentationMap("oci_kms_vault", "test_vault", Required, Create, vaultRepresentation) +
		generateDataSourceFromRepresentationMap("oci_vault_secrets", "test_secrets", Required, Create, secretDataSourceRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabasesResetDatabaseParameterResource_basic(t *testing.T) {
	t.Skip("Skip this test till Database Management service provides a better way of testing this. It requires a live managed database instance")
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabasesResetDatabaseParameterResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_management_managed_databases_reset_database_parameter.test_managed_databases_reset_database_parameter"

	// Save TF content to create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	saveConfigContent(config+compartmentIdVariableStr+generateResourceFromRepresentationMap("oci_database_management_managed_databases_reset_database_parameter", "test_managed_databases_reset_database_parameter", Required, Create, managedDatabasesResetDatabaseParameterRepresentation), "databasemanagement", "managedDatabasesResetDatabaseParameter", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + generateResourceFromRepresentationMap("oci_database_management_managed_databases_reset_database_parameter", "test_managed_databases_reset_database_parameter", Required, Create, managedDatabasesResetDatabaseParameterRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "credentials.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "managed_database_id"),
				resource.TestCheckResourceAttr(resourceName, "parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scope", "BOTH"),
			),
		},
	})
}
