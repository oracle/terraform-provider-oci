// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	oci_common "github.com/oracle/oci-go-sdk/v58/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	securityAssessmentComparisonSingularDataSourceRepresentation = map[string]interface{}{
		"comparison_security_assessment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_security_assessment.test_security_assessment3.id}`},
		"security_assessment_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_security_assessment.test_security_assessment4.id}`},
	}

	SecurityAssessmentComparisonResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create, autonomousDatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_database", "test_target_database", acctest.Required, acctest.Create, targetDatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_assessment", "test_security_assessment3", acctest.Required, acctest.Create, securityAssessmentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_security_assessment", "test_security_assessment4", acctest.Required, acctest.Create, securityAssessmentRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeSecurityAssessmentComparisonResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSecurityAssessmentComparisonResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + SecurityAssessmentComparisonResourceConfig + compartmentIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_assessment_comparison", "test_security_assessment_comparison", acctest.Required, acctest.Create, securityAssessmentComparisonSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					if failure, isServiceError := oci_common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
						return err
					}
					return nil
				},
			),
		},
	})
}
