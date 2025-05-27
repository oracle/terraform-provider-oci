// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	wlmsManagedInstanceOcid = utils.GetEnvSettingWithBlankDefault("wlms_mi_ocid")
	wlmsServerName          = utils.GetEnvSettingWithBlankDefault("wlms_server_name")
	wlsDomainOcid           = utils.GetEnvSettingWithBlankDefault("wlms_wls_domain_ocid")

	WlmsManagedInstanceScanResultDataSourceRepresentation = map[string]interface{}{
		"managed_instance_id": acctest.Representation{RepType: acctest.Required, Create: wlmsManagedInstanceOcid},
		"server_name":         acctest.Representation{RepType: acctest.Optional, Create: wlmsServerName},
		"wls_domain_id":       acctest.Representation{RepType: acctest.Optional, Create: wlsDomainOcid},
	}
)

// issue-routing-tag: wlms/default
func TestWlmsManagedInstanceScanResultResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWlmsManagedInstanceScanResultResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	// compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	// compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_wlms_managed_instance_scan_results.test_managed_instance_scan_results"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_wlms_managed_instance_scan_results", "test_managed_instance_scan_results", acctest.Required, acctest.Create, WlmsManagedInstanceScanResultDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "scan_result_collection.#"),
			),
		},
	})
}
