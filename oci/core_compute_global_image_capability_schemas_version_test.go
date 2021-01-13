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
	computeGlobalImageCapabilitySchemasVersionSingularDataSourceRepresentation = map[string]interface{}{
		"compute_global_image_capability_schema_id":           Representation{repType: Required, create: `${data.oci_core_compute_global_image_capability_schemas.test_compute_global_image_capability_schemas.compute_global_image_capability_schemas.0.id}`},
		"compute_global_image_capability_schema_version_name": Representation{repType: Required, create: `${data.oci_core_compute_global_image_capability_schemas_versions.test_compute_global_image_capability_schemas_versions.compute_global_image_capability_schema_versions.0.name}`},
	}

	computeGlobalImageCapabilitySchemasVersionDataSourceRepresentation = map[string]interface{}{
		"compute_global_image_capability_schema_id": Representation{repType: Required, create: `${data.oci_core_compute_global_image_capability_schemas.test_compute_global_image_capability_schemas.compute_global_image_capability_schemas.0.id}`},
		"display_name": Representation{repType: Optional, create: `displayName`},
	}

	ComputeGlobalImageCapabilitySchemasVersionResourceConfig = generateDataSourceFromRepresentationMap("oci_core_compute_global_image_capability_schemas", "test_compute_global_image_capability_schemas", Required, Create, computeGlobalImageCapabilitySchemaDataSourceRepresentation)
)

func TestCoreComputeGlobalImageCapabilitySchemasVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreComputeGlobalImageCapabilitySchemasVersionResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_compute_global_image_capability_schemas_versions.test_compute_global_image_capability_schemas_versions"
	singularDatasourceName := "data.oci_core_compute_global_image_capability_schemas_version.test_compute_global_image_capability_schemas_version"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_compute_global_image_capability_schemas_versions", "test_compute_global_image_capability_schemas_versions", Required, Create, computeGlobalImageCapabilitySchemasVersionDataSourceRepresentation) +
					compartmentIdVariableStr + ComputeGlobalImageCapabilitySchemasVersionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_core_compute_global_image_capability_schemas_version", "test_compute_global_image_capability_schemas_version", Required, Create, computeGlobalImageCapabilitySchemasVersionSingularDataSourceRepresentation) +
					generateDataSourceFromRepresentationMap("oci_core_compute_global_image_capability_schemas_versions", "test_compute_global_image_capability_schemas_versions", Required, Create, computeGlobalImageCapabilitySchemasVersionDataSourceRepresentation) +
					compartmentIdVariableStr + ComputeGlobalImageCapabilitySchemasVersionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_global_image_capability_schema_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_global_image_capability_schema_version_name"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
		},
	})
}
