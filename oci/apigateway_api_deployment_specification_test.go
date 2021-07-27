// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	apiDeploymentSpecificationSingularDataSourceRepresentation = map[string]interface{}{
		"api_id": Representation{repType: Required, create: `${oci_apigateway_api.test_api.id}`},
	}

	ApiDeploymentSpecificationResourceConfig = DefinedTagsDependencies +
		generateResourceFromRepresentationMap("oci_apigateway_api", "test_api", Optional, Create, apiRepresentation)
)

// issue-routing-tag: apigateway/default
func TestApigatewayApiDeploymentSpecificationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApigatewayApiDeploymentSpecificationResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_apigateway_api_deployment_specification.test_api_deployment_specification"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_apigateway_api_deployment_specification", "test_api_deployment_specification", Required, Create, apiDeploymentSpecificationSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ApiDeploymentSpecificationResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "api_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "logging_policies.#", "0"),
					resource.TestCheckResourceAttr(singularDatasourceName, "request_policies.#", "0"),
					resource.TestCheckResourceAttr(singularDatasourceName, "routes.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "routes.0.path", "/ping"),
					resource.TestCheckResourceAttr(singularDatasourceName, "routes.0.backend.0.status", "200"),
					resource.TestCheckResourceAttr(singularDatasourceName, "routes.0.methods.0", "GET"),
				),
			},
		},
	})
}
