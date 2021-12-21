// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	summarizeResourceInventorySingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"time_end":       acctest.Representation{RepType: acctest.Optional, Create: `2021-11-20T01:00:00Z`},
		"time_start":     acctest.Representation{RepType: acctest.Optional, Create: `2021-11-01T01:00:00Z`},
	}
)

// issue-routing-tag: jms/default
func TestJmsSummarizeResourceInventoryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsSummarizeResourceInventoryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_jms_summarize_resource_inventory.test_summarize_resource_inventory"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_jms_summarize_resource_inventory", "test_summarize_resource_inventory", acctest.Optional, acctest.Create, summarizeResourceInventorySingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "time_end", "2021-11-20T01:00:00Z"),
				resource.TestCheckResourceAttr(singularDatasourceName, "time_start", "2021-11-01T01:00:00Z"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "active_fleet_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "application_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "installation_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "jre_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_instance_count"),
			),
		},
	})
}
