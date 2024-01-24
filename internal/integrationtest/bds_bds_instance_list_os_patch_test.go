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
	BdsBdsInstanceListOsPatchDataSourceRepresentation = map[string]interface{}{
		"bds_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
	}

	BdsBdsInstanceListOsPatchResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Required, acctest.Create, bdsInstanceOdhRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation)
)

// issue-routing-tag: bds/default
func TestBdsBdsInstanceListOsPatchResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBdsBdsInstanceListOsPatchResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_ocid")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	datasourceName := "data.oci_bds_bds_instance_list_os_patches.test_bds_instance_list_os_patches"
	acctest.SaveConfigContent("", "", "", t)
	//	acctest.SaveConfigContent(config+compartmentIdVariableStr+BdsBdsInstanceListOsPatchResourceConfig+
	//		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Required, acctest.Create, BdsBdsInstanceListOsPatchDataSourceRepresentation), "bds", "bdsInstanceListOsPatch", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_bds_instance_list_os_patches", "test_bds_instance_list_os_patches", acctest.Required, acctest.Create, BdsBdsInstanceListOsPatchDataSourceRepresentation) +
				compartmentIdVariableStr + subnetIdVariableStr + BdsBdsInstanceListOsPatchResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "bds_instance_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "os_patches.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "os_patches.0.os_patch_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "os_patches.0.release_date"),
			),
		},
	})
}
