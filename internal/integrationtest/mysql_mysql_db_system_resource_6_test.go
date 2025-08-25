// Copyright (c) 2017, 2025, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	MysqlMysqlDbSystemRepresentationWithMydas = map[string]interface{}{
		"admin_password":          acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"admin_username":          acctest.Representation{RepType: acctest.Required, Create: `adminUser`},
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"shape_name":              acctest.Representation{RepType: acctest.Required, Create: `MySQL.2`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Required, Create: `50`},
		"database_console":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: MysqlMysqlDbSystemDatabaseConsoleRepresentationForUpdate},
		"database_management":     acctest.Representation{RepType: acctest.Optional, Create: `DISABLED`},
	}

	MysqlMysqlDbSystemDatabaseConsoleRepresentationForUpdate = map[string]interface{}{
		"status": acctest.Representation{RepType: acctest.Required, Create: `ENABLED`, Update: `DISABLED`},
		"port":   acctest.Representation{RepType: acctest.Optional, Create: `8443`},
	}
)

// issue-routing-tag: mysql/default
func TestMysqlMysqlDbSystemResource_mydasTest(t *testing.T) {
	httpreplay.SetScenario("TestMysqlMysqlDbSystemResource_mydasTest")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_mysql_db_system.test_mysql_db_system"

	var resId, resId2 string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create with MyDAS enabled
		{
			Config: config + compartmentIdVariableStr + MysqlMysqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Create, MysqlMysqlDbSystemRepresentationWithMydas),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "database_console.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "database_console.0.status", "ENABLED"),
				resource.TestCheckResourceAttr(resourceName, "database_console.0.port", "8443"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to MyDAS disabled
		{
			Config: config + compartmentIdVariableStr + MysqlMysqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Update, MysqlMysqlDbSystemRepresentationWithMydas),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "database_console.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "database_console.0.status", "DISABLED"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
	})
}
