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
	managementAgentCountSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"group_by":       Representation{repType: Required, create: []string{`version`}},
		"has_plugins":    Representation{repType: Required, create: `true`},
	}

	ManagementAgentCountResourceConfig = ""
)

// issue-routing-tag: management_agent/default
func TestManagementAgentManagementAgentCountResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestManagementAgentManagementAgentCountResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_management_agent_management_agent_count.test_management_agent_count"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_management_agent_management_agent_count", "test_management_agent_count", Required, Create, managementAgentCountSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ManagementAgentCountResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "has_plugins", "true"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "items.#"),
				),
			},
		},
	})
}
