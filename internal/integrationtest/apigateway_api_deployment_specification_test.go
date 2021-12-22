// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	apiDeploymentSpecificationSingularDataSourceRepresentation = map[string]interface{}{
		"api_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apigateway_api.test_api.id}`},
	}

	ApiDeploymentSpecificationResourceConfig = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_api", "test_api", acctest.Optional, acctest.Create, apiRepresentation)
)

// issue-routing-tag: apigateway/default
func TestApigatewayApiDeploymentSpecificationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApigatewayApiDeploymentSpecificationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_apigateway_api_deployment_specification.test_api_deployment_specification"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apigateway_api_deployment_specification", "test_api_deployment_specification", acctest.Required, acctest.Create, apiDeploymentSpecificationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ApiDeploymentSpecificationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "api_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "logging_policies.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_policies.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "routes.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "routes.0.path", "/ping"),
				resource.TestCheckResourceAttr(singularDatasourceName, "routes.0.backend.0.status", "200"),
				resource.TestCheckResourceAttr(singularDatasourceName, "routes.0.methods.0", "GET"),
			),
		},
	})
}
