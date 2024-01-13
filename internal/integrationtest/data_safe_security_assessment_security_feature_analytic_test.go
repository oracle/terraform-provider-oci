// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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
	DataSafeSecurityAssessmentSecurityFeatureAnalyticDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Required, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Required, Create: `true`},
	}

	DataSafeSecurityAssessmentSecurityFeatureAnalyticResourceConfig = ""
)

// issue-routing-tag: data_safe/default
func TestDataSafeSecurityAssessmentSecurityFeatureAnalyticResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSecurityAssessmentSecurityFeatureAnalyticResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_data_safe_security_assessment_security_feature_analytics.test_security_assessment_security_feature_analytics"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_assessment_security_feature_analytics", "test_security_assessment_security_feature_analytics", acctest.Required, acctest.Create, DataSafeSecurityAssessmentSecurityFeatureAnalyticDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeSecurityAssessmentSecurityFeatureAnalyticResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "security_feature_analytics_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_feature_analytics_collection.0.items.#"),
			),
		},
	})
}
