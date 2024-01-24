// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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
	ApiGatewayApiSingularDatasourceRepresentation = map[string]interface{}{
		"api_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apigateway_api.test_api.id}`},
	}

	ApiGatewayApiContentResourceConfig = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_api", "test_api", acctest.Optional, acctest.Create, ApigatewayApiRepresentation)
)

// issue-routing-tag: apigateway/default
func TestApigatewayApiContentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApigatewayApiContentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_apigateway_api_content.test_api_content"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apigateway_api_content", "test_api_content", acctest.Required, acctest.Create, ApiGatewayApiSingularDatasourceRepresentation) +
				compartmentIdVariableStr + ApiGatewayApiContentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "api_id"),
			),
		},
	})
}
