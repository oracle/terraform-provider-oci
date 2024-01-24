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
	DataSafeuserAssessmentUserAnalyticDataSourceRepresentation = map[string]interface{}{
		"user_assessment_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_user_assessment.test_user_assessment.id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"account_status":            acctest.Representation{RepType: acctest.Optional, Create: `accountStatus`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"target_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
		"time_last_login_greater_than_or_equal_to":            acctest.Representation{RepType: acctest.Optional, Create: `timeLastLoginGreaterThanOrEqualTo`},
		"time_last_login_less_than":                           acctest.Representation{RepType: acctest.Optional, Create: `timeLastLoginLessThan`},
		"time_password_last_changed_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `timePasswordLastChangedGreaterThanOrEqualTo`},
		"time_password_last_changed_less_than":                acctest.Representation{RepType: acctest.Optional, Create: `timePasswordLastChangedLessThan`},
		"time_user_created_greater_than_or_equal_to":          acctest.Representation{RepType: acctest.Optional, Create: `timeUserCreatedGreaterThanOrEqualTo`},
		"time_user_created_less_than":                         acctest.Representation{RepType: acctest.Optional, Create: `timeUserCreatedLessThan`},
		"user_category":                                       acctest.Representation{RepType: acctest.Optional, Create: `userCategory`},
		"user_key":                                            acctest.Representation{RepType: acctest.Optional, Create: `userKey`},
		"user_name":                                           acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_user.test_user.name}`},
	}

	DataSafeUserAssessmentUserAnalyticResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_user_assessment", "test_user_assessment", acctest.Required, acctest.Create, userAssessmentRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeUserAssessmentUserAnalyticResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeUserAssessmentUserAnalyticResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	datasourceName := "data.oci_data_safe_user_assessment_user_analytics.test_user_assessment_user_analytics"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_user_assessment_user_analytics", "test_user_assessment_user_analytics", acctest.Required, acctest.Create, DataSafeuserAssessmentUserAnalyticDataSourceRepresentation) +
				compartmentIdVariableStr + targetIdVariableStr + DataSafeUserAssessmentUserAnalyticResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "user_aggregations.#"),
				resource.TestCheckResourceAttr(datasourceName, "user_aggregations.0.items.#", "2"),
			),
		},
	})
}
