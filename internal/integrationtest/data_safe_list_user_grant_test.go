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
	listUserGrantDataSourceRepresentation = map[string]interface{}{
		"user_assessment_id":                   acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_user_assessment.test_user_assessment.id}`},
		"user_key":                             acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_data_safe_user_assessment_users.test_user_assessment_users.users[0], "key")}`},
		"depth_level":                          acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"depth_level_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"depth_level_less_than":                acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"grant_key":                            acctest.Representation{RepType: acctest.Optional, Create: `grantKey`},
		"grant_name":                           acctest.Representation{RepType: acctest.Optional, Create: `grantName`},
		"privilege_category":                   acctest.Representation{RepType: acctest.Optional, Create: `privilegeCategory`},
		"privilege_type":                       acctest.Representation{RepType: acctest.Optional, Create: `privilegeType`},
	}

	ListUserGrantResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_user_assessment", "test_user_assessment", acctest.Required, acctest.Create, userAssessmentRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_user_assessment_users", "test_user_assessment_users", acctest.Required, acctest.Create, userAssessmentUserDataSourceRepresentation)
)

func TestDataSafeListUserGrantResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeListUserGrantResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_data_safe_list_user_grants.test_list_user_grants"

	acctest.SaveConfigContent(config+compartmentIdVariableStr+ListUserGrantResourceConfig+
		acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_list_user_grants", "test_list_user_grants", acctest.Required, acctest.Create, listUserGrantDataSourceRepresentation), "datasafe", "listUserGrants", t)
	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_list_user_grants", "test_list_user_grants", acctest.Required, acctest.Create, listUserGrantDataSourceRepresentation) +
				compartmentIdVariableStr + ListUserGrantResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "user_assessment_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "grants.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "grants.0.depth_level"),
				resource.TestCheckResourceAttrSet(datasourceName, "grants.0.grant_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "grants.0.key"),
				resource.TestCheckResourceAttrSet(datasourceName, "grants.0.privilege_type"),
			),
		},
	})
}
