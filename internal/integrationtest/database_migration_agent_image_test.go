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
	agentImageDataSourceRepresentation = map[string]interface{}{}

	AgentImageResourceConfig = ""
)

// issue-routing-tag: database_migration/default
func TestDatabaseMigrationAgentImageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseMigrationAgentImageResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_migration_agent_images.test_agent_images"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_agent_images", "test_agent_images", acctest.Required, acctest.Create, agentImageDataSourceRepresentation) +
				compartmentIdVariableStr + AgentImageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "agent_image_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "agent_image_collection.0.items.#", "1"),
			),
		},
	})
}
