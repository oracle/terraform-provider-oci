// Copyright (c) 2017, 2024, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	ManagedKafkaAddonOptionDataSourceRepresentation = map[string]interface{}{}
)

// issue-routing-tag: managed_kafka/default
func TestManagedKafkaAddonOptionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestManagedKafkaAddonOptionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_managed_kafka_addon_options.test_addon_options"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_managed_kafka_addon_options",
					"test_addon_options",
					acctest.Required,
					acctest.Create,
					ManagedKafkaAddonOptionDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "addon_option_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "addon_option_collection.0.items.#", "1"),
			),
		},
	})
}
