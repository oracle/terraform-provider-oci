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
	pluggableDatabasesLocalCloneRepresentation = map[string]interface{}{
		"cloned_pdb_name":                    Representation{RepType: Required, Create: `NewSalesPdb`},
		"pdb_admin_password":                 Representation{RepType: Required, Create: `BEstrO0ng_#11`},
		"pluggable_database_id":              Representation{RepType: Required, Create: `${oci_database_pluggable_database.test_pluggable_database.id}`},
		"target_tde_wallet_password":         Representation{RepType: Required, Create: `BEstrO0ng_#11`},
		"should_pdb_admin_account_be_locked": Representation{RepType: Optional, Create: `false`},
		"lifecycle":                          RepresentationGroup{Required, ignoreChangesLBRepresentation},
	}
)

// issue-routing-tag: database/default
func TestDatabasePluggableDatabasesLocalCloneResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabasePluggableDatabasesLocalCloneResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_pluggable_databases_local_clone.test_pluggable_databases_local_clone"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+PluggableDatabaseResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_database_pluggable_databases_local_clone", "test_pluggable_databases_local_clone", Optional, Create, pluggableDatabasesLocalCloneRepresentation), "database", "pluggableDatabasesLocalClone", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify local clone
		{
			Config: config + compartmentIdVariableStr + PluggableDatabaseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", Optional, Update, pluggableDatabaseRepresentation) +
				GenerateResourceFromRepresentationMap("oci_database_pluggable_databases_local_clone", "test_pluggable_databases_local_clone", Optional, Create, pluggableDatabasesLocalCloneRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cloned_pdb_name", "NewSalesPdb"),
				resource.TestCheckResourceAttrSet(resourceName, "pluggable_database_id"),
				resource.TestCheckResourceAttr(resourceName, "pdb_admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "target_tde_wallet_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "should_pdb_admin_account_be_locked", "false"),
			),
		},
	})
}
