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
	DataSafeuserAssessmentUserDataSourceRepresentation = map[string]interface{}{
		"user_assessment_id":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_user_assessment.test_user_assessment.id}`},
		"access_level":                             acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree":                acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"account_status":                           acctest.Representation{RepType: acctest.Optional, Create: `accountStatus`},
		"are_all_schemas_accessible":               acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"authentication_type":                      acctest.Representation{RepType: acctest.Optional, Create: `authenticationType`},
		"schema_list":                              acctest.Representation{RepType: acctest.Optional, Create: []string{`schemaList`}},
		"target_id":                                acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
		"time_last_login_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `timeLastLoginGreaterThanOrEqualTo`},
		"time_last_login_less_than":                acctest.Representation{RepType: acctest.Optional, Create: `timeLastLoginLessThan`},
		"time_password_last_changed_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `timePasswordLastChangedGreaterThanOrEqualTo`},
		"time_password_last_changed_less_than":                acctest.Representation{RepType: acctest.Optional, Create: `timePasswordLastChangedLessThan`},
		"time_user_created_greater_than_or_equal_to":          acctest.Representation{RepType: acctest.Optional, Create: `timeUserCreatedGreaterThanOrEqualTo`},
		"time_user_created_less_than":                         acctest.Representation{RepType: acctest.Optional, Create: `timeUserCreatedLessThan`},
		"user_name":                                           acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_user.test_user.name}`},
		"user_profile":                                        acctest.Representation{RepType: acctest.Optional, Create: `userProfile`},
		"user_role":                                           acctest.Representation{RepType: acctest.Optional, Create: `userRole`},
		"user_type":                                           acctest.Representation{RepType: acctest.Optional, Create: `userType`},
	}

	DataSafeUserAssessmentUserResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_user_assessment", "test_user_assessment", acctest.Required, acctest.Create, userAssessmentRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeUserAssessmentUserResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeUserAssessmentUserResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	datasourceName := "data.oci_data_safe_user_assessment_users.test_user_assessment_users"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_user_assessment_users", "test_user_assessment_users", acctest.Required, acctest.Create, DataSafeuserAssessmentUserDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeUserAssessmentUserResourceConfig + targetIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				//resource.TestCheckResourceAttrSet(datasourceName, "schema_list.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "target_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "user_assessment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "users.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "users.0.account_status"),
				resource.TestCheckResourceAttr(datasourceName, "users.0.admin_roles.#", "0"),
				resource.TestCheckResourceAttrSet(datasourceName, "users.0.are_all_schemas_accessible"),
				resource.TestCheckResourceAttrSet(datasourceName, "users.0.admin_roles.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "users.0.authentication_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "users.0.key"),
				resource.TestCheckResourceAttrSet(datasourceName, "users.0.schema_list.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "users.0.target_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "users.0.time_user_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "users.0.user_category"),
				resource.TestCheckResourceAttrSet(datasourceName, "users.0.user_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "users.0.user_profile"),
				resource.TestCheckResourceAttrSet(datasourceName, "users.0.user_types.#"),
			),
		},
	})
}
