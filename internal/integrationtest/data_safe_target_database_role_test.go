// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DataSafetargetDatabaseRoleDataSourceRepresentation = map[string]interface{}{
		"target_database_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
		"authentication_type":  acctest.Representation{RepType: acctest.Optional, Create: `authenticationType`},
		"is_oracle_maintained": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"role_name":            acctest.Representation{RepType: acctest.Optional, Create: []string{`roleName`}},
		"role_name_contains":   acctest.Representation{RepType: acctest.Optional, Create: `roleNameContains`},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeTargetDatabaseRoleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeTargetDatabaseRoleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	datasourceName := "data.oci_data_safe_target_database_roles.test_target_database_roles"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_target_database_roles", "test_target_database_roles", acctest.Required, acctest.Create, DataSafetargetDatabaseRoleDataSourceRepresentation) +
				compartmentIdVariableStr + targetIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "target_database_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "roles.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "roles.0.authentication_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "roles.0.is_common"),
				resource.TestCheckResourceAttrSet(datasourceName, "roles.0.is_implicit"),
				resource.TestCheckResourceAttrSet(datasourceName, "roles.0.is_inherited"),
				resource.TestCheckResourceAttrSet(datasourceName, "roles.0.is_oracle_maintained"),
				resource.TestCheckResourceAttrSet(datasourceName, "roles.0.is_password_required"),
				resource.TestCheckResourceAttrSet(datasourceName, "roles.0.role_name"),
			),
		},
	})
}
