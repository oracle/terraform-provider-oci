// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
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
	DataConnectivityDataConnectivityRegistryTypeSingularDataSourceRepresentation = map[string]interface{}{
		"registry_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_connectivity_registry.test_registry.id}`},
		"type_key":    acctest.Representation{RepType: acctest.Required, Create: `${data.oci_data_connectivity_registry_types.test_registry_types.types_summary_collection.0.items.0.key}`},
		"fields":      acctest.Representation{RepType: acctest.Optional, Create: []string{}},
	}

	DataConnectivityDataConnectivityRegistryTypeDataSourceRepresentation = map[string]interface{}{
		"registry_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_connectivity_registry.test_registry.id}`},
	}

	DataConnectivityRegistryTypeResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_data_connectivity_registry", "test_registry", acctest.Required, acctest.Create, DataConnectivityRegistryRepresentation)
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_connectivity_registry_types", "test_registry_types", acctest.Required, acctest.Create, DataConnectivityDataConnectivityRegistryTypeDataSourceRepresentation) +
				compartmentIdVariableStr + DataConnectivityRegistryTypeResourceConfig,
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_connectivity_registry_types", "test_registry_types", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(DataConnectivityDataConnectivityRegistryTypeDataSourceRepresentation, map[string]interface{}{})) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_connectivity_registry_type", "test_registry_type", acctest.Required, acctest.Create, DataConnectivityDataConnectivityRegistryTypeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataConnectivityRegistryTypeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "registry_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_asset_attributes.#"),
			),
		},
	})
}
