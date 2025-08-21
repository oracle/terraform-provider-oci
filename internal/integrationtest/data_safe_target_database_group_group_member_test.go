// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeTargetDatabaseGroupGroupMemberSingularDataSourceRepresentation = map[string]interface{}{
		"target_database_group_id": acctest.Representation{RepType: acctest.Required, Create: `${var.target_database_group_id}`},
		"target_database_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.target_database_id}`},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeTargetDatabaseGroupGroupMemberResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeTargetDatabaseGroupGroupMemberResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetDatabaseGroupId := utils.GetEnvSettingWithBlankDefault("target_database_group_id")
	targetDatabaseGroupIdVariableStr := fmt.Sprintf("variable \"target_database_group_id\" { default = \"%s\" }\n", targetDatabaseGroupId)

	targetDatabaseId := utils.GetEnvSettingWithBlankDefault("target_database_id")
	targetDatabaseIdVariableStr := fmt.Sprintf("variable \"target_database_id\" { default = \"%s\" }\n", targetDatabaseId)

	singularDatasourceName := "data.oci_data_safe_target_database_group_group_member.test_target_database_group_group_member"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_target_database_group_group_member", "test_target_database_group_group_member", acctest.Required, acctest.Create, DataSafeTargetDatabaseGroupGroupMemberSingularDataSourceRepresentation) +
				compartmentIdVariableStr + targetDatabaseGroupIdVariableStr + targetDatabaseIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_database_group_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_databases.#", "1"),
			),
		},
	})
}
