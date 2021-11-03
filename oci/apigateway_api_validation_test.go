// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	apiValidationSingularDataSourceRepresentation = map[string]interface{}{
		"api_id": Representation{RepType: Required, Create: `${oci_apigateway_api.test_api.id}`},
	}

	ApiValidationResourceConfig = DefinedTagsDependencies +
		GenerateResourceFromRepresentationMap("oci_apigateway_api", "test_api", Optional, Create, apiRepresentation)
)

// issue-routing-tag: apigateway/default
func TestApigatewayApiValidationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApigatewayApiValidationResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_apigateway_api_validation.test_api_validation"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_apigateway_api_validation", "test_api_validation", Required, Create, apiValidationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ApiValidationResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "api_id"),

				// there can be more validations done. Testing 1 which there will be always.
				resource.TestCheckResourceAttr(singularDatasourceName, "validations.0.name", "Schema"),
			),
		},
	})
}
