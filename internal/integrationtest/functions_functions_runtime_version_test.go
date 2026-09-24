// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	FunctionsFunctionsRuntimeVersionSingularDataSourceRepresentation = map[string]interface{}{
		"functions_runtime_version_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_functions_functions_runtime_versions.test_functions_runtime_versions.functions_runtime_version_collection.0.items.0.id}`},
	}

	FunctionsFunctionsRuntimeVersionDataSourceRepresentation = map[string]interface{}{
		"display_name":                 acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"functions_runtime_id":         acctest.Representation{RepType: acctest.Required, Create: `${data.oci_functions_functions_runtimes.test_functions_runtimes.functions_runtime_collection.0.items.0.id}`},
		"functions_runtime_name":       acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_functions_functions_runtimes.test_functions_runtimes.functions_runtime_collection.0.items.0.name}`},
		"functions_runtime_version_id": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_functions_functions_runtime_versions.test_functions_runtime_versions.functions_runtime_version_collection.0.items.0.id}`},
		"is_current_version":           acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"language_version":             acctest.Representation{RepType: acctest.Optional, Create: `languageVersion`},
		"os_version":                   acctest.Representation{RepType: acctest.Optional, Create: `osVersion`},
		"state":                        acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
	}

	FunctionsFunctionsRuntimeVersionResourceConfig = ""
)

// issue-routing-tag: functions/default
func TestFunctionsFunctionsRuntimeVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFunctionsFunctionsRuntimeVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_functions_functions_runtime_versions.test_functions_runtime_versions"
	singularDatasourceName := "data.oci_functions_functions_runtime_version.test_functions_runtime_version"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_functions_functions_runtimes", "test_functions_runtimes", acctest.Required, acctest.Create, FunctionsFunctionsRuntimeDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_functions_functions_runtime_versions", "test_functions_runtime_versions", acctest.Required, acctest.Create, FunctionsFunctionsRuntimeVersionDataSourceRepresentation) +
				compartmentIdVariableStr + FunctionsFunctionsRuntimeVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "functions_runtime_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "functions_runtime_version_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "functions_runtime_version_collection.0.items.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "functions_runtime_version_collection.0.items.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "functions_runtime_version_collection.0.items.0.functions_runtime_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "functions_runtime_version_collection.0.items.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "functions_runtime_version_collection.0.items.0.language_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "functions_runtime_version_collection.0.items.0.os_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "functions_runtime_version_collection.0.items.0.state"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_functions_functions_runtime_version", "test_functions_runtime_version", acctest.Required, acctest.Create, FunctionsFunctionsRuntimeVersionSingularDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_functions_functions_runtimes", "test_functions_runtimes", acctest.Required, acctest.Create, FunctionsFunctionsRuntimeDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_functions_functions_runtime_versions", "test_functions_runtime_versions", acctest.Required, acctest.Create, FunctionsFunctionsRuntimeVersionDataSourceRepresentation) +
				compartmentIdVariableStr + FunctionsFunctionsRuntimeVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "functions_runtime_version_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "language_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "metadata"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "os_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrWith(singularDatasourceName, "supported_architectures.#", func(value string) error {
					count, err := strconv.Atoi(value)
					if err != nil {
						return err
					}
					if count < 1 {
						return fmt.Errorf("expected at least one supported architecture, got %d", count)
					}
					return nil
				}),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
