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
	MysqlMysqlDbSystemRepresentationWithRest = map[string]interface{}{
		"mysql_version":           acctest.Representation{RepType: acctest.Required, Create: `9.3.1`},
		"admin_password":          acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"admin_username":          acctest.Representation{RepType: acctest.Required, Create: `adminUser`},
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"shape_name":              acctest.Representation{RepType: acctest.Required, Create: `MySQL.2`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Required, Create: `50`},
		"database_management":     acctest.Representation{RepType: acctest.Optional, Create: `DISABLED`},
		"rest":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: MysqlMysqlDbSystemRestRepresentationForUpdate},
	}

	MysqlMysqlDbSystemRestRepresentationForUpdate = map[string]interface{}{
		"configuration": acctest.Representation{RepType: acctest.Required, Create: `DBSYSTEM_ONLY`, Update: `DISABLED`},
		"port":          acctest.Representation{RepType: acctest.Optional, Create: `443`},
	}
)

// issue-routing-tag: mysql/default
func TestMysqlMysqlDbSystemResource_restTest(t *testing.T) {
	httpreplay.SetScenario("TestMysqlMysqlDbSystemResource_restTest")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_mysql_db_system.test_mysql_db_system"

	var resId, resId2 string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create with REST enabled
		{
			Config: config + compartmentIdVariableStr + MysqlMysqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Create, MysqlMysqlDbSystemRepresentationWithRest),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "rest.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rest.0.configuration", "DBSYSTEM_ONLY"),
				resource.TestCheckResourceAttr(resourceName, "rest.0.port", "443"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to REST disabled
		{
			Config: config + compartmentIdVariableStr + MysqlMysqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Update, MysqlMysqlDbSystemRepresentationWithRest),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "rest.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rest.0.configuration", "DISABLED"),

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
