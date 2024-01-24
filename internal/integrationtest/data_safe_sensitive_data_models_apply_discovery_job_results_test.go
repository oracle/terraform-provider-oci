// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	SensitiveDataModelsApplyDiscoveryJobResultsRepresentation = map[string]interface{}{
		"sensitive_data_model_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_sensitive_data_model.test_sensitive_data_model.id}`},
		"discovery_job_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_discovery_job.test_discovery_job.id}`},
	}

	SensitiveDataModelsApplyDiscoveryJobResultsDependencies = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_model", "test_sensitive_data_model", acctest.Required, acctest.Create, sensitiveDataModelRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_discovery_job", "test_discovery_job", acctest.Required, acctest.Create, discoveryJobRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeSensitiveDataModelsApplyDiscoveryJobResultsResource_basic(t *testing.T) {
	httpreplay.SetScenario("DataSafeSensitiveDataModelsApplyDiscoveryJobResultsResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	resourceName := "oci_data_safe_sensitive_data_models_apply_discovery_job_results.test_sensitive_data_models_apply_discovery_job_results"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+SensitiveDataModelsApplyDiscoveryJobResultsDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_models_apply_discovery_job_results", "test_sensitive_data_models_apply_discovery_job_results", acctest.Required, acctest.Create, SensitiveDataModelsApplyDiscoveryJobResultsRepresentation), "datasafe", "sensitiveDataModel", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + SensitiveDataModelsApplyDiscoveryJobResultsDependencies + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_models_apply_discovery_job_results", "test_sensitive_data_models_apply_discovery_job_results", acctest.Required, acctest.Create, SensitiveDataModelsApplyDiscoveryJobResultsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

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
	})
}
