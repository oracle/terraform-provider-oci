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
	registryTypeSingularDataSourceRepresentation = map[string]interface{}{
		"registry_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_connectivity_registry.test_registry.id}`},
		"type_key":    acctest.Representation{RepType: acctest.Required, Create: `${data.oci_data_connectivity_registry_types.test_registry_types.types_summary_collection.0.items.0.key}`},
		"fields":      acctest.Representation{RepType: acctest.Optional, Create: []string{}},
	}

	registryTypeDataSourceRepresentation = map[string]interface{}{
		"registry_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_connectivity_registry.test_registry.id}`},
	}

	RegistryTypeResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_data_connectivity_registry", "test_registry", acctest.Required, acctest.Create, registryRepresentation)
)

// issue-routing-tag: data_connectivity/default
func TestDataConnectivityRegistryTypeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataConnectivityRegistryTypeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_data_connectivity_registry_types.test_registry_types"
	singularDatasourceName := "data.oci_data_connectivity_registry_type.test_registry_type"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_connectivity_registry_types", "test_registry_types", acctest.Required, acctest.Create, registryTypeDataSourceRepresentation) +
				compartmentIdVariableStr + RegistryTypeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "registry_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "types_summary_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "types_summary_collection.0.items.0.name"),
				resource.TestCheckResourceAttrSet(datasourceName, "types_summary_collection.0.items.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "types_summary_collection.0.items.0.key"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_connectivity_registry_types", "test_registry_types", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(registryTypeDataSourceRepresentation, map[string]interface{}{})) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_connectivity_registry_type", "test_registry_type", acctest.Required, acctest.Create, registryTypeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + RegistryTypeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "registry_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_asset_attributes.#"),
			),
		},
	})
}
