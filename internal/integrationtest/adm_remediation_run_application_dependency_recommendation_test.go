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
	AdmRemediationRunApplicationDependencyRecommendationDataSourceRepresentation = map[string]interface{}{
		"remediation_run_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_adm_remediation_run.test_remediation_run.id}`},
		"gav":                acctest.Representation{RepType: acctest.Optional, Create: `gav`},
		"purl":               acctest.Representation{RepType: acctest.Optional, Create: `purl`},
	}

	AdmRemediationRunApplicationDependencyRecommendationResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_adm_knowledge_base", "test_knowledge_base", acctest.Required, acctest.Create, knowledgeBaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_adm_remediation_recipe", "test_remediation_recipe", acctest.Required, acctest.Create, AdmRemediationRecipeRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_adm_remediation_run", "test_remediation_run", acctest.Required, acctest.Create, AdmRemediationRunRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) + DefinedTagsDependencies
)

// issue-routing-tag: adm/default
func TestAdmRemediationRunApplicationDependencyRecommendationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAdmRemediationRunApplicationDependencyRecommendationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	secretId := utils.GetEnvSettingWithBlankDefault("kms_secret_ocid")
	secretIdVariableStr := fmt.Sprintf("variable \"kms_secret_ocid\" { default = \"%s\" }\n", secretId)

	codeRepositoryId := utils.GetEnvSettingWithBlankDefault("devops_code_repository_ocid")
	codeRepositoryIdStr := fmt.Sprintf("variable \"devops_code_repository_ocid\" { default = \"%s\" }\n", codeRepositoryId)

	pipelineId := utils.GetEnvSettingWithBlankDefault("devops_build_pipeline_ocid")
	pipelineIdStr := fmt.Sprintf("variable \"devops_build_pipeline_ocid\" { default = \"%s\" }\n", pipelineId)

	datasourceName := "data.oci_adm_remediation_run_application_dependency_recommendations.test_remediation_run_application_dependency_recommendations"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + secretIdVariableStr + codeRepositoryIdStr + pipelineIdStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_adm_remediation_run_application_dependency_recommendations", "test_remediation_run_application_dependency_recommendations", acctest.Optional, acctest.Create, AdmRemediationRunApplicationDependencyRecommendationDataSourceRepresentation) +
				compartmentIdVariableStr + AdmRemediationRunApplicationDependencyRecommendationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "gav", "gav"),
				resource.TestCheckResourceAttr(datasourceName, "purl", "purl"),
				resource.TestCheckResourceAttrSet(datasourceName, "remediation_run_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "application_dependency_recommendation_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "application_dependency_recommendation_collection.0.items.#"),
			),
		},
	})
}
