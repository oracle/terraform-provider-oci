// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeUserAssessmentProfileDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"user_assessment_id":                          acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_user_assessment.test_user_assessment.id}`},
		"access_level":                                acctest.Representation{RepType: acctest.Optional, Create: `RESTRICTED`},
		"compartment_id_in_subtree":                   acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"failed_login_attempts_greater_than_or_equal": acctest.Representation{RepType: acctest.Optional, Create: `failedLoginAttemptsGreaterThanOrEqual`},
		"failed_login_attempts_less_than":             acctest.Representation{RepType: acctest.Optional, Create: `failedLoginAttemptsLessThan`},
		"inactive_account_time_greater_than_or_equal": acctest.Representation{RepType: acctest.Optional, Create: `inactiveAccountTimeGreaterThanOrEqual`},
		"inactive_account_time_less_than":             acctest.Representation{RepType: acctest.Optional, Create: `inactiveAccountTimeLessThan`},
		"is_user_created":                             acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"password_lock_time_greater_than_or_equal":    acctest.Representation{RepType: acctest.Optional, Create: `passwordLockTimeGreaterThanOrEqual`},
		"password_lock_time_less_than":                acctest.Representation{RepType: acctest.Optional, Create: `passwordLockTimeLessThan`},
		"password_verification_function":              acctest.Representation{RepType: acctest.Optional, Create: `passwordVerificationFunction`},
		"profile_name":                                acctest.Representation{RepType: acctest.Optional, Create: `${oci_optimizer_profile.test_profile.name}`},
		"sessions_per_user_greater_than_or_equal":     acctest.Representation{RepType: acctest.Optional, Create: `sessionsPerUserGreaterThanOrEqual`},
		"sessions_per_user_less_than":                 acctest.Representation{RepType: acctest.Optional, Create: `sessionsPerUserLessThan`},
		"target_id":                                   acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_guard_target.test_target.id}`},
		"user_count_greater_than_or_equal":            acctest.Representation{RepType: acctest.Optional, Create: `userCountGreaterThanOrEqual`},
		"user_count_less_than":                        acctest.Representation{RepType: acctest.Optional, Create: `userCountLessThan`},
	}

	DataSafeUserAssessmentProfileResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_user_assessment", "test_user_assessment", acctest.Required, acctest.Create, userAssessmentRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeUserAssessmentProfileResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeUserAssessmentProfileResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	datasourceName := "data.oci_data_safe_user_assessment_profiles.test_user_assessment_profiles"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_user_assessment_profiles", "test_user_assessment_profiles", acctest.Required, acctest.Create, DataSafeUserAssessmentProfileDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeUserAssessmentProfileResourceConfig + targetIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "profiles.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "profiles.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "profiles.0.failed_login_attempts"),
				resource.TestCheckResourceAttrSet(datasourceName, "profiles.0.inactive_account_time"),
				resource.TestCheckResourceAttrSet(datasourceName, "profiles.0.is_user_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "profiles.0.password_lock_time"),
				resource.TestCheckResourceAttrSet(datasourceName, "profiles.0.password_verification_function"),
				resource.TestCheckResourceAttrSet(datasourceName, "profiles.0.profile_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "profiles.0.sessions_per_user"),
				resource.TestCheckResourceAttrSet(datasourceName, "profiles.0.target_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "profiles.0.user_assessment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "profiles.0.user_count"),
			),
		},
	})
}
