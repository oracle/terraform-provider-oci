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

// issue-routing-tag: bds/default
func TestBdsBdsInstanceOSPatchActionResource(t *testing.T) {
	//t.Skip("Run manual with an older cluster with patch available")
	httpreplay.SetScenario("TestBdsBdsInstanceOSPatchActionResource")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_ocid")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	// Pass cluster ocid as variable to directly apply OS patch to existing cluster
	//bdsinstanceId := utils.GetEnvSettingWithBlankDefault("bdsinstance_ocid")
	//bdsinstanceIdVariableStr := fmt.Sprintf("variable \"bdsinstance_id\" { default = \"%s\" }\n", bdsinstanceId)

	// Passing patching configs strategy & parameters in this representation
	BdsBdsInstanceOSPatchConfigRepresentation := map[string]interface{}{
		"patching_config_strategy":           acctest.Representation{RepType: acctest.Required, Create: "PATCHING_BASED"},
		"batch_size":                         acctest.Representation{RepType: acctest.Required, Create: "3"},
		"wait_time_between_batch_in_seconds": acctest.Representation{RepType: acctest.Required, Create: "60"},
		"tolerance_threshold_per_batch":      acctest.Representation{RepType: acctest.Required, Create: "0"},
	}

	// Passing timeout representation since OS patching may take more time then default values
	PatchTimeoutsRepresentation := map[string]interface{}{
		"create": acctest.Representation{RepType: acctest.Required, Create: `4h`},
		"update": acctest.Representation{RepType: acctest.Required, Create: `4h`},
		"delete": acctest.Representation{RepType: acctest.Required, Create: `4h`},
	}

	// To use default patching strategy (nodes will be patched and rebooted AD/FD by AD/FD), comment patching_configs & above config representation
	var (
		BdsBdsInstanceOSPatchActionRepresentation = map[string]interface{}{
			"bds_instance_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
			"cluster_admin_password": acctest.Representation{RepType: acctest.Required, Create: `T3JhY2xlVGVhbVVTQSExMjM=`},
			"os_patch_version":       acctest.Representation{RepType: acctest.Required, Create: "ol7.9-x86_64-1.28.0.619-0.0"},
			"patching_configs":       acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsBdsInstanceOSPatchConfigRepresentation},
			"timeouts":               acctest.RepresentationGroup{RepType: acctest.Required, Group: PatchTimeoutsRepresentation},
		}

		BdsBdsInstanceOSPatchActionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Required, acctest.Create, bdsInstanceOdhRepresentation) +
			acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
			acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation)
	)

	resourceName := "oci_bds_bds_instance_os_patch_action.test_bds_instance_os_patch_action"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+subnetIdVariableStr+BdsBdsInstanceOSPatchActionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_os_patch_action", "test_bds_instance_os_patch_action", acctest.Required, acctest.Create, BdsBdsInstanceOSPatchActionRepresentation), "bds", "bdsInstanceOSPatchAction", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + BdsBdsInstanceOSPatchActionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_os_patch_action", "test_bds_instance_os_patch_action", acctest.Required, acctest.Create, BdsBdsInstanceOSPatchActionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "T3JhY2xlVGVhbVVTQSExMjM="),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
			),
		},
	})

}
