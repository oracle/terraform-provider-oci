// Copyright (c) 2017, 2025, Oracle and/or its affiliates. All rights reserved.

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
	mysqlDbSystemSecurityAttributes = map[string]interface{}{
		// standard required properties
		"admin_password":          acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"admin_username":          acctest.Representation{RepType: acctest.Required, Create: `adminUser`},
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"shape_name":              acctest.Representation{RepType: acctest.Required, Create: `MySQL.VM.Standard.E3.1.8GB`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Required, Create: `50`},

		// use an easier to track display name
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `TestDbSystemSecurityAttributes`},

		// avoid wasting time setting up DBM when that's not what we're testing here
		"database_management": acctest.Representation{RepType: acctest.Optional, Create: `DISABLED`},

		// disable backup policy to avoid wasting even more time and resources
		"backup_policy": acctest.RepresentationGroup{RepType: acctest.Required, Group: disabledBackupPolicy},

		// create with the default SAs and then clear SAs on update
		"security_attributes": acctest.Representation{RepType: acctest.Optional,
			Create: map[string]string{"oracle-zpr.sensitivity.value": "low", "oracle-zpr.sensitivity.mode": "enforce"},
			Update: map[string]string{}},
	}
)

// issue-routing-tag: mysql/default
func TestMysqlMysqlDbSystemResource_securityAttributes(t *testing.T) {
	httpreplay.SetScenario("TestMysqlMysqlDbSystemResource_securityAttributes")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_mysql_db_system.test_mysql_db_system"

	var resId, resId2 string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify create results in a DBS with SA
		{
			Config: config + compartmentIdVariableStr + MysqlMysqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Create, mysqlDbSystemSecurityAttributes),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "security_attributes.%", "2"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify update results in the removal of SA
		{
			Config: config + compartmentIdVariableStr + MysqlMysqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Update, mysqlDbSystemSecurityAttributes),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "security_attributes.%", "0"),

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
