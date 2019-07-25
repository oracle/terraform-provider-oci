// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

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

	ConsoleHistoryContentResourceConfig = ConsoleHistoryRequiredOnlyResource
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
