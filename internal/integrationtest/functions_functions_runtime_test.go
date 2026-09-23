// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	FunctionsFunctionsRuntimeSingularDataSourceRepresentation = map[string]interface{}{
		"functions_runtime_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_functions_functions_runtimes.test_functions_runtimes.functions_runtime_collection.0.items.0.id}`},
	}

	FunctionsFunctionsRuntimeDataSourceRepresentation = map[string]interface{}{
		"functions_runtime_id": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_functions_functions_runtimes.test_functions_runtimes.functions_runtime_collection.0.items.0.id}`},
		"language":             acctest.Representation{RepType: acctest.Optional, Create: `language`},
		"name":                 acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"name_contains":        acctest.Representation{RepType: acctest.Optional, Create: `nameContains`},
		"name_starts_with":     acctest.Representation{RepType: acctest.Optional, Create: `nameStartsWith`},
		"os":                   acctest.Representation{RepType: acctest.Optional, Create: `os`},
		"state":                acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
	}

	FunctionsFunctionsRuntimeResourceConfig = ""
)

// issue-routing-tag: functions/default
func TestFunctionsFunctionsRuntimeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFunctionsFunctionsRuntimeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_functions_functions_runtimes.test_functions_runtimes"
	singularDatasourceName := "data.oci_functions_functions_runtime.test_functions_runtime"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_functions_functions_runtimes", "test_functions_runtimes", acctest.Required, acctest.Create, FunctionsFunctionsRuntimeDataSourceRepresentation) +
				compartmentIdVariableStr + FunctionsFunctionsRuntimeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "functions_runtime_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "functions_runtime_collection.0.items.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "functions_runtime_collection.0.items.0.current_functions_runtime_version_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "functions_runtime_collection.0.items.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "functions_runtime_collection.0.items.0.language"),
				resource.TestCheckResourceAttrSet(datasourceName, "functions_runtime_collection.0.items.0.name"),
				resource.TestCheckResourceAttrSet(datasourceName, "functions_runtime_collection.0.items.0.os"),
				resource.TestCheckResourceAttrSet(datasourceName, "functions_runtime_collection.0.items.0.state"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_functions_functions_runtime", "test_functions_runtime", acctest.Required, acctest.Create, FunctionsFunctionsRuntimeSingularDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_functions_functions_runtimes", "test_functions_runtimes", acctest.Required, acctest.Create, FunctionsFunctionsRuntimeDataSourceRepresentation) +
				compartmentIdVariableStr + FunctionsFunctionsRuntimeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "functions_runtime_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "current_functions_runtime_version_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "language"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "os"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_decommissioned"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_deprecated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
