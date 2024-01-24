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
	ApigatewayApiValidationSingularDataSourceRepresentation = map[string]interface{}{
		"api_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apigateway_api.test_api.id}`},
	}

	ApigatewayApiValidationResourceConfig = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_api", "test_api", acctest.Optional, acctest.Create, ApigatewayApiRepresentation)
)

// issue-routing-tag: apigateway/default
func TestApigatewayApiValidationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApigatewayApiValidationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_apigateway_api_validation.test_api_validation"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apigateway_api_validation", "test_api_validation", acctest.Required, acctest.Create, ApigatewayApiValidationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ApigatewayApiValidationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "api_id"),

				// there can be more validations done. Testing 1 which there will be always.
				resource.TestCheckResourceAttr(singularDatasourceName, "validations.0.name", "Schema"),
			),
		},
	})
}
