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
	agentImageDataSourceRepresentation = map[string]interface{}{}

	AgentImageResourceConfig = ""
)

// issue-routing-tag: database_migration/default
func TestDatabaseMigrationAgentImageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseMigrationAgentImageResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_migration_agent_images.test_agent_images"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_migration_agent_images", "test_agent_images", Required, Create, agentImageDataSourceRepresentation) +
					compartmentIdVariableStr + AgentImageResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(

					resource.TestCheckResourceAttrSet(datasourceName, "agent_image_collection.#"),
					resource.TestCheckResourceAttr(datasourceName, "agent_image_collection.0.items.#", "1"),
				),
			},
		},
	})
}
