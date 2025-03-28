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
	CoreCoreComputeGlobalImageCapabilitySchemasVersionSingularDataSourceRepresentation = map[string]interface{}{

		"compute_global_image_capability_schema_id":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_compute_global_image_capability_schemas.test_compute_global_image_capability_schemas.compute_global_image_capability_schemas.0.id}`},
		"compute_global_image_capability_schema_version_name": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_compute_global_image_capability_schemas_versions.test_compute_global_image_capability_schemas_versions.compute_global_image_capability_schema_versions.0.name}`},
	}

	CoreCoreComputeGlobalImageCapabilitySchemasVersionDataSourceRepresentation = map[string]interface{}{

		"compute_global_image_capability_schema_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_compute_global_image_capability_schemas.test_compute_global_image_capability_schemas.compute_global_image_capability_schemas.0.id}`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
	}

	CoreComputeGlobalImageCapabilitySchemasVersionResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_core_compute_global_image_capability_schemas", "test_compute_global_image_capability_schemas", acctest.Required, acctest.Create, CoreCoreComputeGlobalImageCapabilitySchemaDataSourceRepresentation)
)

// issue-routing-tag: core/computeImaging
func TestCoreComputeGlobalImageCapabilitySchemasVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreComputeGlobalImageCapabilitySchemasVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_compute_global_image_capability_schemas_versions.test_compute_global_image_capability_schemas_versions"
	singularDatasourceName := "data.oci_core_compute_global_image_capability_schemas_version.test_compute_global_image_capability_schemas_version"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_compute_global_image_capability_schemas_versions", "test_compute_global_image_capability_schemas_versions", acctest.Required, acctest.Create, CoreCoreComputeGlobalImageCapabilitySchemasVersionDataSourceRepresentation) +
				compartmentIdVariableStr + CoreComputeGlobalImageCapabilitySchemasVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "compute_global_image_capability_schema_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "compute_global_image_capability_schema_versions.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_global_image_capability_schema_versions.0.compute_global_image_capability_schema_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_global_image_capability_schema_versions.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_global_image_capability_schema_versions.0.name"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_global_image_capability_schema_versions.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_compute_global_image_capability_schemas_version", "test_compute_global_image_capability_schemas_version", acctest.Required, acctest.Create, CoreCoreComputeGlobalImageCapabilitySchemasVersionSingularDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_compute_global_image_capability_schemas_versions", "test_compute_global_image_capability_schemas_versions", acctest.Required, acctest.Create, CoreCoreComputeGlobalImageCapabilitySchemasVersionDataSourceRepresentation) +
				compartmentIdVariableStr + CoreComputeGlobalImageCapabilitySchemasVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_global_image_capability_schema_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_global_image_capability_schema_version_name"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
	})
}
