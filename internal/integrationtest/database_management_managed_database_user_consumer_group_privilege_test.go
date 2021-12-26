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
	managedDatabaseUserConsumerGroupPrivilegeSingularDataSourceRepresentation = map[string]interface{}{
		"managed_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_managed_database.test_managed_database.id}`},
		"user_name":           acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_user.test_user.name}`},
		"name":                acctest.Representation{RepType: acctest.Optional, Create: `name`},
	}

	managedDatabaseUserConsumerGroupPrivilegeDataSourceRepresentation = map[string]interface{}{
		"managed_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_managed_database.test_managed_database.id}`},
		"user_name":           acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_user.test_user.name}`},
		"name":                acctest.Representation{RepType: acctest.Optional, Create: `name`},
	}

	ManagedDatabaseUserConsumerGroupPrivilegeResourceConfig = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseUserConsumerGroupPrivilegeResource_basic(t *testing.T) {
	t.Skip("Skip this test till Database Management service provides a better way of testing this. It requires a live managed database instance")
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseUserConsumerGroupPrivilegeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_management_managed_database_user_consumer_group_privileges.test_managed_database_user_consumer_group_privileges"
	singularDatasourceName := "data.oci_database_management_managed_database_user_consumer_group_privilege.test_managed_database_user_consumer_group_privilege"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_user_consumer_group_privileges", "test_managed_database_user_consumer_group_privileges", acctest.Required, acctest.Create, managedDatabaseUserConsumerGroupPrivilegeDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedDatabaseUserConsumerGroupPrivilegeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "user_name"),

				resource.TestCheckResourceAttrSet(datasourceName, "consumer_group_privilege_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "consumer_group_privilege_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_user_consumer_group_privilege", "test_managed_database_user_consumer_group_privilege", acctest.Required, acctest.Create, managedDatabaseUserConsumerGroupPrivilegeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedDatabaseUserConsumerGroupPrivilegeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_name"),

				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "1"),
			),
		},
	})
}
