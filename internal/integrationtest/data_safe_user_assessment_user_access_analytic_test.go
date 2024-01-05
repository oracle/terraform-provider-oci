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
	DataSafeUserAssessmentUserAccessAnalyticDataSourceRepresentation = map[string]interface{}{
		"user_assessment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.user_assessment_id}`},
	}

	DataSafeUserAssessmentUserAccessAnalyticResourceConfig = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeUserAssessmentUserAccessAnalyticResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeUserAssessmentUserAccessAnalyticResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	userAssessmentId := utils.GetEnvSettingWithBlankDefault("user_assessment_ocid")
	userAssessmentIdVariableStr := fmt.Sprintf("variable \"user_assessment_id\" { default = \"%s\" }\n", userAssessmentId)

	datasourceName := "data.oci_data_safe_user_assessment_user_access_analytics.test_user_assessment_user_access_analytics"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_user_assessment_user_access_analytics", "test_user_assessment_user_access_analytics", acctest.Required, acctest.Create, DataSafeUserAssessmentUserAccessAnalyticDataSourceRepresentation) +
				compartmentIdVariableStr + userAssessmentIdVariableStr + DataSafeUserAssessmentUserAccessAnalyticResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "user_assessment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "user_access_analytics_collection.#"),
				//resource.TestCheckResourceAttr(datasourceName, "user_access_analytics_collection.0.items.#", "1"),
			),
		},
	})
}
