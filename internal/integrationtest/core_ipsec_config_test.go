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
	CoreCoreIpSecConnectionDeviceConfigSingularDataSourceRepresentation = map[string]interface{}{
		"ipsec_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_ipsec.test_ip_sec_connection.id}`},
	}

	CoreIpSecConnectionDeviceConfigResourceConfig = CoreIpSecConnectionRequiredOnlyResource
)

// issue-routing-tag: core/default
func TestCoreIpSecConnectionDeviceConfigResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreIpSecConnectionDeviceConfigResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_ipsec_config.test_ip_sec_connection_device_config"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_ipsec_config", "test_ip_sec_connection_device_config", acctest.Required, acctest.Create, CoreCoreIpSecConnectionDeviceConfigSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CoreIpSecConnectionDeviceConfigResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ipsec_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tunnels.#", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tunnels.0.ip_address"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tunnels.0.shared_secret"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tunnels.0.time_created"),
			),
		},
	})
}
