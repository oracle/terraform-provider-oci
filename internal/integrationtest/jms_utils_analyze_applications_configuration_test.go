// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	JmsUtilsAnalyzeApplicationsConfigurationCompartmentId = utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	JmsUtilsAnalyzeApplicationsConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: JmsUtilsAnalyzeApplicationsConfigurationCompartmentId},
	}
)

// issue-routing-tag: jms_utils/default
func TestJmsUtilsAnalyzeApplicationsConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsUtilsAnalyzeApplicationsConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	singularDatasourceName := "data.oci_jms_utils_analyze_applications_configuration.test_analyze_applications_configuration"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify create
		// note: we cannot write test for this case because
		// we don't have create API.

		// verify update
		// note: we cannot write test for this case because
		// we don't have update API.

		// verify singular datasource

		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_utils_analyze_applications_configuration",
					"test_analyze_applications_configuration",
					acctest.Optional,
					acctest.Create,
					JmsUtilsAnalyzeApplicationsConfigurationSingularDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				// check actual data matches data used for the GET API
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", JmsUtilsAnalyzeApplicationsConfigurationCompartmentId),
				// check actual data is set (doesn't make much sense to hardcode more values)
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bucket"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
			),
		},
	})
}
