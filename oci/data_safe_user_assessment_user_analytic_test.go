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
	userAssessmentUserAnalyticDataSourceRepresentation = map[string]interface{}{
		"user_assessment_id":        Representation{RepType: Required, Create: `${oci_data_safe_user_assessment.test_user_assessment.id}`},
		"access_level":              Representation{RepType: Optional, Create: `ACCESSIBLE`},
		"account_status":            Representation{RepType: Optional, Create: `accountStatus`},
		"compartment_id_in_subtree": Representation{RepType: Optional, Create: `true`},
		"target_id":                 Representation{RepType: Required, Create: `${oci_data_safe_target_database.test_target_database.id}`},
		"time_last_login_greater_than_or_equal_to":            Representation{RepType: Optional, Create: `timeLastLoginGreaterThanOrEqualTo`},
		"time_last_login_less_than":                           Representation{RepType: Optional, Create: `timeLastLoginLessThan`},
		"time_password_last_changed_greater_than_or_equal_to": Representation{RepType: Optional, Create: `timePasswordLastChangedGreaterThanOrEqualTo`},
		"time_password_last_changed_less_than":                Representation{RepType: Optional, Create: `timePasswordLastChangedLessThan`},
		"time_user_created_greater_than_or_equal_to":          Representation{RepType: Optional, Create: `timeUserCreatedGreaterThanOrEqualTo`},
		"time_user_created_less_than":                         Representation{RepType: Optional, Create: `timeUserCreatedLessThan`},
		"user_category":                                       Representation{RepType: Optional, Create: `userCategory`},
		"user_key":                                            Representation{RepType: Optional, Create: `userKey`},
		"user_name":                                           Representation{RepType: Optional, Create: `${oci_identity_user.test_user.name}`},
	}

	UserAssessmentUserAnalyticResourceConfig = GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", Required, Create, autonomousDatabaseRepresentation) +
		GenerateResourceFromRepresentationMap("oci_data_safe_target_database", "test_target_database", Required, Create, targetDatabaseRepresentation) +
		GenerateResourceFromRepresentationMap("oci_data_safe_user_assessment", "test_user_assessment", Required, Create, userAssessmentRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeUserAssessmentUserAnalyticResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeUserAssessmentUserAnalyticResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_data_safe_user_assessment_user_analytics.test_user_assessment_user_analytics"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_data_safe_user_assessment_user_analytics", "test_user_assessment_user_analytics", Required, Create, userAssessmentUserAnalyticDataSourceRepresentation) +
				compartmentIdVariableStr + UserAssessmentUserAnalyticResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "user_aggregation.#"),
				resource.TestCheckResourceAttr(datasourceName, "user_aggregation.0.items.#", "0"),
			),
		},
	})
}
