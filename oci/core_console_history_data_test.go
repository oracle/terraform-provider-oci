// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	consoleHistoryContentSingularDataSourceRepresentation = map[string]interface{}{
		"console_history_id": Representation{repType: Required, create: `${oci_core_console_history.test_console_history.id}`},
		"length":             Representation{repType: Optional, create: `10240`},
		"offset":             Representation{repType: Optional, create: `0`},
	}

	ConsoleHistoryContentResourceConfig = generateResourceFromRepresentationMap("oci_core_console_history", "test_console_history", Required, Create, consoleHistoryRepresentation) +
		OciImageIdsVariable +
		generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceRepresentation) +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		AvailabilityDomainConfig
)

func TestCoreConsoleHistoryContentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreConsoleHistoryContentResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_console_history_data.test_console_history_content"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_console_history_data", "test_console_history_content", Optional, Create, consoleHistoryContentSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ConsoleHistoryContentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "console_history_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "length", "10240"),
					resource.TestCheckResourceAttr(singularDatasourceName, "offset", "0"),
				),
			},
		},
	})
}
