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
	DataSafeUserAssessmentProfileAnalyticDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"user_assessment_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_user_assessment.test_user_assessment.id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `RESTRICTED`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"target_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
	}

	DataSafeUserAssessmentProfileAnalyticResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_user_assessment", "test_user_assessment", acctest.Required, acctest.Create, userAssessmentRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeUserAssessmentProfileAnalyticResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeUserAssessmentProfileAnalyticResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_data_safe_user_assessment_profile_analytics.test_user_assessment_profile_analytics"

	acctest.SaveConfigContent("", "", "", t)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_user_assessment_profile_analytics", "test_user_assessment_profile_analytics", acctest.Required, acctest.Create, DataSafeUserAssessmentProfileAnalyticDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeUserAssessmentProfileAnalyticResourceConfig + targetIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "profile_aggregations.#"),
				resource.TestCheckResourceAttr(datasourceName, "profile_aggregations.0.items.#", "2"),
			),
		},
	})
}
