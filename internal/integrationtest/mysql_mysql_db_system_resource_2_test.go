// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.

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
	mysqlDbSystemRepresentationDataStorageUpdate = map[string]interface{}{
		"admin_password":          acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"admin_username":          acctest.Representation{RepType: acctest.Required, Create: `adminUser`},
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"configuration_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.MysqlConfigurationOCID[var.region]}`},
		"shape_name":              acctest.Representation{RepType: acctest.Required, Create: `MySQL.VM.Standard.E3.1.8GB`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Required, Create: `50`},
		"data_storage":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: MysqlMysqlDbSystemDataStorageRepresentation},
	}

	mysqlDbSystemRepresentationDbSystemModesUpdate = map[string]interface{}{
		"admin_password":          acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"admin_username":          acctest.Representation{RepType: acctest.Required, Create: `adminUser`},
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"configuration_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.MysqlConfigurationOCID[var.region]}`},
		"shape_name":              acctest.Representation{RepType: acctest.Required, Create: `MySQL.VM.Standard.E3.1.8GB`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Required, Create: `50`},
		"access_mode":             acctest.Representation{RepType: acctest.Optional, Create: `UNRESTRICTED`, Update: `RESTRICTED`},
		"database_mode":           acctest.Representation{RepType: acctest.Optional, Create: `READ_WRITE`, Update: `READ_ONLY`},
		"database_management":     acctest.Representation{RepType: acctest.Optional, Create: `DISABLED`},
	}

	MysqlMysqlDbSystemDataStorageRepresentation = map[string]interface{}{
		"is_auto_expand_storage_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"max_storage_size_in_gbs":        acctest.Representation{RepType: acctest.Optional, Create: `100`, Update: `150`},
	}
)

// issue-routing-tag: mysql/default
func TestMysqlMysqlDbSystemResource_dataStorageTest(t *testing.T) {
	httpreplay.SetScenario("TestMysqlMysqlDbSystemResource_dataStorageTest")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_mysql_db_system.test_mysql_db_system"

	var resId, resId2 string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create with data_storage
		{
			Config: config + compartmentIdVariableStr + MysqlMysqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Create, mysqlDbSystemRepresentationDataStorageUpdate),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "data_storage.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage.0.is_auto_expand_storage_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "data_storage.0.max_storage_size_in_gbs", "100"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_gb", "50"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify update to enable automatic data storage
		{
			Config: config + compartmentIdVariableStr + MysqlMysqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Update, mysqlDbSystemRepresentationDataStorageUpdate),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "data_storage.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage.0.is_auto_expand_storage_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "data_storage.0.max_storage_size_in_gbs", "150"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_gb", "50"),

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

// issue-routing-tag: mysql/default
func TestMysqlMysqlDbSystemResource_databaseModes(t *testing.T) {
	httpreplay.SetScenario("TestMysqlMysqlDbSystemResource_databaseModes")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_mysql_db_system.test_mysql_db_system"

	var resId, resId2 string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create with DbSystem modes
		{
			Config: config + compartmentIdVariableStr + MysqlMysqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Create, mysqlDbSystemRepresentationDbSystemModesUpdate),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "access_mode", "UNRESTRICTED"),
				resource.TestCheckResourceAttr(resourceName, "database_mode", "READ_WRITE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify update of DbSystem modes
		{
			Config: config + compartmentIdVariableStr + MysqlMysqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Update, mysqlDbSystemRepresentationDbSystemModesUpdate),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "access_mode", "RESTRICTED"),
				resource.TestCheckResourceAttr(resourceName, "database_mode", "READ_ONLY"),

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
