// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	listUserGrantDataSourceRepresentation = map[string]interface{}{
		"user_assessment_id":                   Representation{RepType: Required, Create: `${oci_data_safe_user_assessment.test_user_assessment.id}`},
		"user_key":                             Representation{RepType: Required, Create: `${lookup(data.oci_data_safe_user_assessment_users.test_user_assessment_users.users[0], "key")}`},
		"depth_level":                          Representation{RepType: Optional, Create: `10`},
		"depth_level_greater_than_or_equal_to": Representation{RepType: Optional, Create: `10`},
		"depth_level_less_than":                Representation{RepType: Optional, Create: `10`},
		"grant_key":                            Representation{RepType: Optional, Create: `grantKey`},
		"grant_name":                           Representation{RepType: Optional, Create: `grantName`},
		"privilege_category":                   Representation{RepType: Optional, Create: `privilegeCategory`},
		"privilege_type":                       Representation{RepType: Optional, Create: `privilegeType`},
	}

	ListUserGrantResourceConfig = GenerateResourceFromRepresentationMap("oci_data_safe_user_assessment", "test_user_assessment", Required, Create, userAssessmentRepresentation) +
		GenerateDataSourceFromRepresentationMap("oci_data_safe_user_assessment_users", "test_user_assessment_users", Required, Create, userAssessmentUserDataSourceRepresentation)
)

func TestDataSafeListUserGrantResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeListUserGrantResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_data_safe_list_user_grants.test_list_user_grants"

	SaveConfigContent(config+compartmentIdVariableStr+ListUserGrantResourceConfig+
		GenerateDataSourceFromRepresentationMap("oci_data_safe_list_user_grants", "test_list_user_grants", Required, Create, listUserGrantDataSourceRepresentation), "datasafe", "listUserGrants", t)
	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_data_safe_list_user_grants", "test_list_user_grants", Required, Create, listUserGrantDataSourceRepresentation) +
				compartmentIdVariableStr + ListUserGrantResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
