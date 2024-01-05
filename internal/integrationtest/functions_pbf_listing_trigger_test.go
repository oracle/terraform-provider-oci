// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	FunctionsFunctionsPbfListingTriggerDataSourceRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Optional, Create: `HTTP`},
	}
)

// issue-routing-tag: functions/default
func TestFunctionsPbfListingTriggerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestFunctionsPbfListingTriggerResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_functions_pbf_listing_triggers.test_pbf_listing_triggers"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_functions_pbf_listing_triggers", "test_pbf_listing_triggers", acctest.Required, acctest.Create, FunctionsFunctionsPbfListingTriggerDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckNoResourceAttr(datasourceName, "name"),

				resource.TestCheckResourceAttr(datasourceName, "triggers_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "triggers_collection.0.items.#"),
			),
		},
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_functions_pbf_listing_triggers", "test_pbf_listing_triggers", acctest.Optional, acctest.Create, FunctionsFunctionsPbfListingTriggerDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "name", "HTTP"),

				resource.TestCheckResourceAttr(datasourceName, "triggers_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "triggers_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "triggers_collection.0.items.0.name", "HTTP"),
			),
		},
	})
}
