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
	MysqlDbSystemRepresentationCopyPoliciesUpdate = map[string]interface{}{
		"admin_password":      acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"admin_username":      acctest.Representation{RepType: acctest.Required, Create: `adminUser`},
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"backup_policy":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: mysqlMysqlDbSystemBackupPolicyRepresentation},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"configuration_id":    acctest.Representation{RepType: acctest.Optional, Create: `${var.MysqlConfigurationOCID[var.region]}`},
		"shape_name":          acctest.Representation{RepType: acctest.Required, Create: `MySQL.VM.Standard.E3.1.8GB`},
		"subnet_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
	}

	mysqlMysqlDbSystemBackupPolicyRepresentation = map[string]interface{}{
		"copy_policies":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: mysqlMysqlDbSystemBackupPolicyCopyPoliciesRepresentation},
		"defined_tags":      acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"freeform_tags":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
		"is_enabled":        acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"pitr_policy":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: mysqlDbSystemBackupPolicyPitrPolicyNotUpdateableRepresentation},
		"retention_in_days": acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}
	mysqlMysqlDbSystemBackupPolicyCopyPoliciesRepresentation = map[string]interface{}{
		"copy_to_region":                acctest.Representation{RepType: acctest.Required, Create: `us-phoenix-1`, Update: `eu-frankfurt-1`},
		"backup_copy_retention_in_days": acctest.Representation{RepType: acctest.Optional, Create: `4`, Update: `5`},
	}
)

// issue-routing-tag: mysql/default
func TestMysqlMysqlDbSystemResource_copyPoliciesTest(t *testing.T) {
	httpreplay.SetScenario("TestMysqlMysqlDbSystemResource_copyPoliciesTest")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_mysql_db_system.test_mysql_db_system"

	var resId, resId2 string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify create with copy policy.
		{
			Config: config + compartmentIdVariableStr + MysqlMysqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Create, MysqlDbSystemRepresentationCopyPoliciesUpdate),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backup_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.copy_policies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.copy_policies.0.backup_copy_retention_in_days", "4"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.copy_policies.0.copy_to_region", "us-phoenix-1"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.is_enabled", "true"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify update to copy policy.
		{
			Config: config + compartmentIdVariableStr + MysqlMysqlDbSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Optional, acctest.Update, MysqlDbSystemRepresentationCopyPoliciesUpdate),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backup_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.copy_policies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.copy_policies.0.backup_copy_retention_in_days", "5"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.copy_policies.0.copy_to_region", "eu-frankfurt-1"),
				resource.TestCheckResourceAttr(resourceName, "backup_policy.0.is_enabled", "true"),

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
