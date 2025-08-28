// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeSecurityAssessmentTemplateBaselineComparisonSingularDataSourceRepresentation = map[string]interface{}{
		"comparison_security_assessment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.security_assessment_baseline_ocid}`},
		"security_assessment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.security_assessment_ocid}`},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeSecurityAssessmentTemplateBaselineComparisonResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSecurityAssessmentTemplateBaselineComparisonResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compareSAId := utils.GetEnvSettingWithBlankDefault("security_assessment_baseline_ocid")
	compareSAIdStr := fmt.Sprintf("variable \"security_assessment_baseline_ocid\" { default = \"%s\" }\n", compareSAId)

	saId := utils.GetEnvSettingWithBlankDefault("security_assessment_ocid")
	saIdStr := fmt.Sprintf("variable \"security_assessment_ocid\" { default = \"%s\" }\n", saId)
	singularDatasourceName := "data.oci_data_safe_security_assessment_template_baseline_comparison.test_security_assessment_template_baseline_comparison"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_assessment_template_baseline_comparison", "test_security_assessment_template_baseline_comparison", acctest.Required, acctest.Create, DataSafeSecurityAssessmentTemplateBaselineComparisonSingularDataSourceRepresentation) +
				compartmentIdVariableStr + compareSAIdStr + saIdStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "comparison_security_assessment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "security_assessment_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "auditing.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "authorization_control.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_encryption.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_configuration.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fine_grained_access_control.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "privileges_and_roles.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "template_baseline_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "template_baseline_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_accounts.#"),
			),
		},
	})
}
