// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	pluggableDatabasesLocalCloneRepresentation = map[string]interface{}{
		"cloned_pdb_name":                    acctest.Representation{RepType: acctest.Required, Create: `NewSalesPdb`},
		"pdb_admin_password":                 acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"pluggable_database_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_database_pluggable_database.test_pluggable_database.id}`},
		"target_tde_wallet_password":         acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"should_pdb_admin_account_be_locked": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"lifecycle":                          acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesLBRepresentation},
	}
)

// issue-routing-tag: database/default
func TestDatabasePluggableDatabasesLocalCloneResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabasePluggableDatabasesLocalCloneResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_pluggable_databases_local_clone.test_pluggable_databases_local_clone"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+PluggableDatabaseResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_databases_local_clone", "test_pluggable_databases_local_clone", acctest.Optional, acctest.Create, pluggableDatabasesLocalCloneRepresentation), "database", "pluggableDatabasesLocalClone", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify local clone
		{
			Config: config + compartmentIdVariableStr + PluggableDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Optional, acctest.Update, pluggableDatabaseRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_databases_local_clone", "test_pluggable_databases_local_clone", acctest.Optional, acctest.Create, pluggableDatabasesLocalCloneRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cloned_pdb_name", "NewSalesPdb"),
				resource.TestCheckResourceAttrSet(resourceName, "pluggable_database_id"),
				resource.TestCheckResourceAttr(resourceName, "pdb_admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "target_tde_wallet_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "should_pdb_admin_account_be_locked", "false"),
			),
		},
	})
}
