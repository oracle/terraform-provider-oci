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
	BdsBdsInstanceGetOsPatchDataSourceRepresentation = map[string]interface{}{
		"bds_instance_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"os_patch_version": acctest.Representation{RepType: acctest.Required, Create: `${var.os_patch_version}`},
	}

	BdsBdsInstanceGetOsPatchResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Required, acctest.Create, bdsInstanceOdhRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation)
)

// issue-routing-tag: bds/default
func TestBdsBdsInstanceGetOsPatchResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBdsBdsInstanceGetOsPatchResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	osPatchVersion := utils.GetEnvSettingWithBlankDefault("os_patch_version")
	osPatchVersionVariableStr := fmt.Sprintf("variable \"os_patch_version\" { default = \"%s\" }\n", osPatchVersion)

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_ocid")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	datasourceName := "data.oci_bds_bds_instance_get_os_patch.test_bds_instance_get_os_patch"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_bds_instance_get_os_patch", "test_bds_instance_get_os_patch", acctest.Required, acctest.Create, BdsBdsInstanceGetOsPatchDataSourceRepresentation) +
				compartmentIdVariableStr + subnetIdVariableStr + osPatchVersionVariableStr + BdsBdsInstanceGetOsPatchResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(datasourceName, "os_patch_version", "ol7.9-x86_64-1.24.0.100-0.0"),
				resource.TestCheckResourceAttrSet(datasourceName, "min_bds_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "patch_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "release_date"),
			),
		},
	})
}
