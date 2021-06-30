// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

func TestDatabasePluggableDatabasesLocalCloneResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabasePluggableDatabasesLocalCloneResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_pluggable_databases_local_clone.test_pluggable_databases_local_clone"

	// Save TF content to create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	saveConfigContent(config+compartmentIdVariableStr+PluggableDatabaseResourceDependencies+
		generateResourceFromRepresentationMap("oci_database_pluggable_databases_local_clone", "test_pluggable_databases_local_clone", Required, Create, pluggableDatabasesLocalCloneRepresentation), "database", "pluggableDatabasesLocalClone", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify local clone
			{
				Config: config + compartmentIdVariableStr + PluggableDatabaseResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", Required, Update, pluggableDatabaseRepresentation) +
					generateResourceFromRepresentationMap("oci_database_pluggable_databases_local_clone", "test_pluggable_databases_local_clone", Required, Create, pluggableDatabasesLocalCloneRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "cloned_pdb_name", "NewSalesPdb"),
					resource.TestCheckResourceAttr(resourceName, "pdb_admin_password", "BEstrO0ng_#11"),
					resource.TestCheckResourceAttrSet(resourceName, "pluggable_database_id"),
					resource.TestCheckResourceAttr(resourceName, "target_tde_wallet_password", "BEstrO0ng_#11"),
				),
			},
		},
	})
}
