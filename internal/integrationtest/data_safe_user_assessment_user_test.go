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
	userAssessmentUserDataSourceRepresentation = map[string]interface{}{
		"user_assessment_id":                                  acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_user_assessment.test_user_assessment.id}`},
		"access_level":                                        acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree":                           acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"target_id":                                           acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_target_database.test_target_database.id}`},
		"time_last_login_greater_than_or_equal_to":            acctest.Representation{RepType: acctest.Optional, Create: `timeLastLoginGreaterThanOrEqualTo`},
		"time_last_login_less_than":                           acctest.Representation{RepType: acctest.Optional, Create: `timeLastLoginLessThan`},
		"time_password_last_changed_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `timePasswordLastChangedGreaterThanOrEqualTo`},
		"time_password_last_changed_less_than":                acctest.Representation{RepType: acctest.Optional, Create: `timePasswordLastChangedLessThan`},
		"time_user_created_greater_than_or_equal_to":          acctest.Representation{RepType: acctest.Optional, Create: `timeUserCreatedGreaterThanOrEqualTo`},
		"time_user_created_less_than":                         acctest.Representation{RepType: acctest.Optional, Create: `timeUserCreatedLessThan`},
		"user_name":                                           acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_user.test_user.name}`},
	}

	UserAssessmentUserResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create, autonomousDatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_database", "test_target_database", acctest.Required, acctest.Create, targetDatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_user_assessment", "test_user_assessment", acctest.Required, acctest.Create, userAssessmentRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeUserAssessmentUserResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeUserAssessmentUserResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_data_safe_user_assessment_users.test_user_assessment_users"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_user_assessment_users", "test_user_assessment_users", acctest.Required, acctest.Create, userAssessmentUserDataSourceRepresentation) +
				compartmentIdVariableStr + UserAssessmentUserResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "user_assessment_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "users.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "users.0.account_status"),
				resource.TestCheckResourceAttrSet(datasourceName, "users.0.admin_roles.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "users.0.authentication_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "users.0.key"),
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
