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
	BdsBdsInstancePatchHistoryDataSourceRepresentation = map[string]interface{}{
		"bds_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"patch_type":      acctest.Representation{RepType: acctest.Optional, Create: `ODH`},
		"patch_version":   acctest.Representation{RepType: acctest.Optional, Create: `patchVersion`},
		"state":           acctest.Representation{RepType: acctest.Optional, Create: `INSTALLED`},
	}

	BdsBdsInstancePatchHistoryResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Required, acctest.Create, bdsInstanceOdhRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation)
)

// issue-routing-tag: bds/default
func TestBdsBdsInstancePatchHistoryResource_basic(t *testing.T) {
	t.Skip("Run manual with an older cluster with patch available")
	httpreplay.SetScenario("TestBdsBdsInstancePatchHistoryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_ocid")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	datasourceName := "data.oci_bds_bds_instance_patch_histories.test_bds_instance_patch_histories"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_bds_bds_instance_patch_histories", "test_bds_instance_patch_histories", acctest.Required, acctest.Create, BdsBdsInstancePatchHistoryDataSourceRepresentation) +
				compartmentIdVariableStr + subnetIdVariableStr + BdsBdsInstancePatchHistoryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(datasourceName, "patch_type", "ODH"),
				resource.TestCheckResourceAttr(datasourceName, "patch_version", "patchVersion"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttrSet(datasourceName, "patch_histories.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "patch_histories.0.patch_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "patch_histories.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "patch_histories.0.time_updated"),
				resource.TestCheckResourceAttrSet(datasourceName, "patch_histories.0.version"),
			),
		},
	})
}
