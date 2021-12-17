// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	consoleHistoryContentSingularDataSourceRepresentation = map[string]interface{}{
		"console_history_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_console_history.test_console_history.id}`},
		"length":             acctest.Representation{RepType: acctest.Optional, Create: `10240`},
		"offset":             acctest.Representation{RepType: acctest.Optional, Create: `0`},
	}

	ConsoleHistoryContentResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_console_history", "test_console_history", acctest.Required, acctest.Create, consoleHistoryRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceRepresentation) +
		AvailabilityDomainConfig
)

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreConsoleHistoryContentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreConsoleHistoryContentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_console_history_data.test_console_history_content"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_console_history_data", "test_console_history_content", acctest.Optional, acctest.Create, consoleHistoryContentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ConsoleHistoryContentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "console_history_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "length", "10240"),
				resource.TestCheckResourceAttr(singularDatasourceName, "offset", "0"),
			),
		},
	})
}
