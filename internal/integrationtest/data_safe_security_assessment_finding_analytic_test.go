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
	DataSafeSecurityAssessmentFindingAnalyticDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `RESTRICTED`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"finding_key":               acctest.Representation{RepType: acctest.Optional, Create: `findingKey`},
		"group_by":                  acctest.Representation{RepType: acctest.Optional, Create: `findingKeyAndTopFindingStatus`},
		"is_top_finding":            acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"severity":                  acctest.Representation{RepType: acctest.Optional, Create: `HIGH`},
		"top_finding_status":        acctest.Representation{RepType: acctest.Optional, Create: `RISK`},
	}

	DataSafeSecurityAssessmentFindingAnalyticResourceConfig = ""
)

// issue-routing-tag: data_safe/default
func TestDataSafeSecurityAssessmentFindingAnalyticResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSecurityAssessmentFindingAnalyticResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_data_safe_security_assessment_finding_analytics.test_security_assessment_finding_analytics"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_assessment_finding_analytics", "test_security_assessment_finding_analytics", acctest.Required, acctest.Create, DataSafeSecurityAssessmentFindingAnalyticDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeSecurityAssessmentFindingAnalyticResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "finding_analytics_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "finding_analytics_collection.0.items.#"),
			),
		},
	})
}
