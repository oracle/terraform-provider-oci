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
	CoreCoreComputeGlobalImageCapabilitySchemaSingularDataSourceRepresentation = map[string]interface{}{

		"compute_global_image_capability_schema_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_core_compute_global_image_capability_schemas.test_compute_global_image_capability_schemas.compute_global_image_capability_schemas.0.id}`},
	}

	CoreCoreComputeGlobalImageCapabilitySchemaDataSourceRepresentation = map[string]interface{}{}

	CoreComputeGlobalImageCapabilitySchemaResourceConfig = ""
)

// issue-routing-tag: core/computeImaging
func TestCoreComputeGlobalImageCapabilitySchemaResource_basic(t *testing.T) {

	httpreplay.SetScenario("TestCoreComputeGlobalImageCapabilitySchemaResource_basic")

	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_compute_global_image_capability_schemas.test_compute_global_image_capability_schemas"
	singularDatasourceName := "data.oci_core_compute_global_image_capability_schema.test_compute_global_image_capability_schema"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_compute_global_image_capability_schemas", "test_compute_global_image_capability_schemas", acctest.Required, acctest.Create, CoreCoreComputeGlobalImageCapabilitySchemaDataSourceRepresentation) +
				compartmentIdVariableStr + CoreComputeGlobalImageCapabilitySchemaResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "compute_global_image_capability_schemas.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_global_image_capability_schemas.0.current_version_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_global_image_capability_schemas.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_global_image_capability_schemas.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_global_image_capability_schemas.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_compute_global_image_capability_schema", "test_compute_global_image_capability_schema", acctest.Required, acctest.Create, CoreCoreComputeGlobalImageCapabilitySchemaSingularDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_compute_global_image_capability_schemas", "test_compute_global_image_capability_schemas", acctest.Required, acctest.Create, CoreCoreComputeGlobalImageCapabilitySchemaDataSourceRepresentation) +
				compartmentIdVariableStr + CoreComputeGlobalImageCapabilitySchemaResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_global_image_capability_schema_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "current_version_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
	})
}
