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
	IdentityIdentityCostTrackingTagDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	IdentityCostTrackingTagResourceConfig = ""
)

// issue-routing-tag: identity/default
func TestIdentityCostTrackingTagResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityCostTrackingTagResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_identity_cost_tracking_tags.test_cost_tracking_tags"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_cost_tracking_tags", "test_cost_tracking_tags", acctest.Required, acctest.Create, IdentityIdentityCostTrackingTagDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityCostTrackingTagResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "tags.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "tags.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "tags.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "tags.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "tags.0.is_cost_tracking"),
				resource.TestCheckResourceAttrSet(datasourceName, "tags.0.is_retired"),
				resource.TestCheckResourceAttrSet(datasourceName, "tags.0.name"),
				resource.TestCheckResourceAttrSet(datasourceName, "tags.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "tags.0.tag_namespace_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "tags.0.tag_namespace_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "tags.0.time_created"),
			),
		},
	})
}
