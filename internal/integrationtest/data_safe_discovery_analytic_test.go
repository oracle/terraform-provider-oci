// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	discoveryAnalyticDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"group_by":                  acctest.Representation{RepType: acctest.Optional, Create: `targetId`},
	}

	DiscoveryAnalyticResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_discovery_analytics", "test_discovery_analytics", acctest.Required, acctest.Create, discoveryAnalyticDataSourceRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeDiscoveryAnalyticResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeDiscoveryAnalyticResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_data_safe_discovery_analytics.test_discovery_analytics"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + DiscoveryAnalyticResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "discovery_analytics_collection.#"),
			),
		},
	})
}
