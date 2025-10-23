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
	JmsUtilsAnalyzeApplicationsConfigurationCompartmentId = utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	JmsUtilsAnalyzeApplicationsConfigurationBucket        = utils.GetEnvSettingWithBlankDefault("bucket")
	JmsUtilsAnalyzeApplicationsConfigurationNamespace     = utils.GetEnvSettingWithBlankDefault("namespace")

	JmsUtilsAnalyzeApplicationsConfigurationRequiredOnlyResource = JmsUtilsAnalyzeApplicationsConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_jms_utils_analyze_applications_configuration", "test_analyze_applications_configuration", acctest.Required, acctest.Create, JmsUtilsAnalyzeApplicationsConfigurationRepresentation)

	JmsUtilsAnalyzeApplicationsConfigurationResourceConfig = JmsUtilsAnalyzeApplicationsConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_jms_utils_analyze_applications_configuration", "test_analyze_applications_configuration", acctest.Optional, acctest.Update, JmsUtilsAnalyzeApplicationsConfigurationRepresentation)

	JmsUtilsAnalyzeApplicationsConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: JmsUtilsAnalyzeApplicationsConfigurationCompartmentId},
	}

	JmsUtilsAnalyzeApplicationsConfigurationRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: JmsUtilsAnalyzeApplicationsConfigurationCompartmentId},
		"bucket":         acctest.Representation{RepType: acctest.Optional, Create: JmsUtilsAnalyzeApplicationsConfigurationBucket, Update: JmsUtilsAnalyzeApplicationsConfigurationBucket},
		"namespace":      acctest.Representation{RepType: acctest.Optional, Create: JmsUtilsAnalyzeApplicationsConfigurationNamespace, Update: JmsUtilsAnalyzeApplicationsConfigurationNamespace},
	}

	JmsUtilsAnalyzeApplicationsConfigurationResourceDependencies = ""
)

// issue-routing-tag: jms_utils/default
func TestJmsUtilsAnalyzeApplicationsConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsUtilsAnalyzeApplicationsConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := JmsUtilsAnalyzeApplicationsConfigurationCompartmentId
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_jms_utils_analyze_applications_configuration.test_analyze_applications_configuration"

	singularDatasourceName := "data.oci_jms_utils_analyze_applications_configuration.test_analyze_applications_configuration"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+JmsUtilsAnalyzeApplicationsConfigurationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_jms_utils_analyze_applications_configuration", "test_analyze_applications_configuration", acctest.Optional, acctest.Create, JmsUtilsAnalyzeApplicationsConfigurationRepresentation), "jmsutils", "analyzeApplicationsConfiguration", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + JmsUtilsAnalyzeApplicationsConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_jms_utils_analyze_applications_configuration", "test_analyze_applications_configuration", acctest.Required, acctest.Create, JmsUtilsAnalyzeApplicationsConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + JmsUtilsAnalyzeApplicationsConfigurationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + JmsUtilsAnalyzeApplicationsConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_jms_utils_analyze_applications_configuration", "test_analyze_applications_configuration", acctest.Optional, acctest.Create, JmsUtilsAnalyzeApplicationsConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bucket", JmsUtilsAnalyzeApplicationsConfigurationBucket),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "namespace", JmsUtilsAnalyzeApplicationsConfigurationNamespace),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + JmsUtilsAnalyzeApplicationsConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_jms_utils_analyze_applications_configuration", "test_analyze_applications_configuration", acctest.Optional, acctest.Update, JmsUtilsAnalyzeApplicationsConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bucket", JmsUtilsAnalyzeApplicationsConfigurationBucket),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "namespace", JmsUtilsAnalyzeApplicationsConfigurationNamespace),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_jms_utils_analyze_applications_configuration", "test_analyze_applications_configuration", acctest.Required, acctest.Create, JmsUtilsAnalyzeApplicationsConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + JmsUtilsAnalyzeApplicationsConfigurationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(singularDatasourceName, "bucket", JmsUtilsAnalyzeApplicationsConfigurationBucket),
				resource.TestCheckResourceAttr(singularDatasourceName, "namespace", JmsUtilsAnalyzeApplicationsConfigurationNamespace),
			),
		},
		// verify resource import
		{
			Config:            config + JmsUtilsAnalyzeApplicationsConfigurationRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"compartment_id",
			},
			ResourceName: resourceName,
		},
	})
}
