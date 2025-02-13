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
	MysqlMysqlDbSystemRepresentationWithReadEndpoint = map[string]interface{}{
		"admin_password":          acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"admin_username":          acctest.Representation{RepType: acctest.Required, Create: `adminUser`},
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"shape_name":              acctest.Representation{RepType: acctest.Required, Create: `MySQL.VM.Standard3.4.64GB`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Required, Create: `50`},
		"database_management":     acctest.Representation{RepType: acctest.Optional, Create: `DISABLED`},
		"backup_policy":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: MysqlMysqlDbSystemBackupPolicyDisabledRepresentation},
		"read_endpoint":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: MysqlMysqlDbSystemReadEndpointRepresentationForUpdate},
		"lifecycle":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: ignoreMysqlMysqlDbSystemDefinedTagsChangesRepresentation},
	}

	ignoreMysqlMysqlDbSystemDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Optional, Create: []string{`defined_tags`}},
	}

	MysqlMysqlDbSystemReadEndpointRepresentationForUpdate = map[string]interface{}{
		"is_enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
	}

	MysqlMysqlDbSystemBackupPolicyDisabledRepresentation = map[string]interface{}{
		"is_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
	}
)

// issue-routing-tag: mysql/default
func TestMysqlMysqlDbSystemResource_readEndpointTest(t *testing.T) {
	httpreplay.SetScenario("TestMysqlMysqlDbSystemResource_readEndpointTest")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_mysql_db_system.test_mysql_db_system"

	var resId, resId2 string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create with read endpoint
		{
			Config: config + compartmentIdVariableStr + MysqlMysqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Create, MysqlMysqlDbSystemRepresentationWithReadEndpoint),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "read_endpoint.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "read_endpoint.0.is_enabled", "true"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to read endpoint
		{
			Config: config + compartmentIdVariableStr + MysqlMysqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Update, MysqlMysqlDbSystemRepresentationWithReadEndpoint),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "read_endpoint.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "read_endpoint.0.is_enabled", "false"),

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
