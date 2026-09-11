// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeCryptoAssessmentSingularDataSourceRepresentation = map[string]interface{}{
		"crypto_assessment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.crypto_assessment_id}`},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeCryptoAssessmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeCryptoAssessmentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	cryptoAssessmentId := utils.GetEnvSettingWithBlankDefault("crypto_assessment_id")
	cryptoAssessmentIdVariableStr := fmt.Sprintf("variable \"crypto_assessment_id\" { default = \"%s\" }\n", cryptoAssessmentId)

	datasourceName := "data.oci_data_safe_crypto_assessment.test_crypto_assessment"
	resourceName := "oci_data_safe_crypto_assessment.test_crypto_assessment"
	var resId string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_crypto_assessment", "test_crypto_assessment", acctest.Required, acctest.Create, DataSafeCryptoAssessmentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + cryptoAssessmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "crypto_assessment_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "display_name"),
				resource.TestCheckResourceAttr(datasourceName, "type", "LATEST"),
				resource.TestCheckResourceAttr(datasourceName, "target_type", "TARGET_DATABASE"),
				resource.TestCheckResourceAttrSet(datasourceName, "database_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "posture_category"),
				resource.TestCheckResourceAttrSet(datasourceName, "issue_count"),
				resource.TestCheckResourceAttrSet(datasourceName, "targets_with_issues_count"),
				resource.TestCheckResourceAttrSet(datasourceName, "is_assessment_scheduled"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, datasourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
	})
}
