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
	PsqlInsightCapabilityDataSourceRepresentation = map[string]interface{}{}

	PsqlInsightCapabilityResourceConfig = ""
)

// issue-routing-tag: psql/default
func TestPsqlInsightCapabilityResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestPsqlInsightCapabilityResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_psql_insight_capabilities.test_insight_capabilities"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_psql_insight_capabilities", "test_insight_capabilities", acctest.Required, acctest.Create, PsqlInsightCapabilityDataSourceRepresentation) +
				compartmentIdVariableStr + PsqlInsightCapabilityResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "insight_capability_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "insight_capability_collection.0.items.#", "1"),
			),
		},
	})
}
