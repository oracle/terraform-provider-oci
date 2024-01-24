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
	CoreCoreConsoleHistoryContentSingularDataSourceRepresentation = map[string]interface{}{
		"console_history_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_console_history.test_console_history.id}`},
		"length":             acctest.Representation{RepType: acctest.Optional, Create: `10240`},
		"offset":             acctest.Representation{RepType: acctest.Optional, Create: `0`},
	}

	CoreConsoleHistoryContentResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_console_history", "test_console_history", acctest.Required, acctest.Create, CoreConsoleHistoryRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreInstanceRepresentation) +
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_console_history_data", "test_console_history_content", acctest.Optional, acctest.Create, CoreCoreConsoleHistoryContentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CoreConsoleHistoryContentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "console_history_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "length", "10240"),
				resource.TestCheckResourceAttr(singularDatasourceName, "offset", "0"),
			),
		},
	})
}
