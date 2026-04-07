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

// issue-routing-tag: bds/default
func TestBdsBdsInstanceExecuteBootstrapScriptActionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBdsBdsInstanceExecuteBootstrapScriptActionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_ocid")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	// Pass cluster ocid as variable to directly apply OS patch to existing cluster
	bdsInstanceId := utils.GetEnvSettingWithBlankDefault("bds_instance_ocid")
	bdsInstanceIdVariableStr := fmt.Sprintf("variable \"bds_instance_id\" { default = \"%s\" }\n", bdsInstanceId)

	bootstrapScriptUrl := utils.GetEnvSettingWithBlankDefault("bootstrap_script_url")
	bootstrapScriptUrlVariableStr := fmt.Sprintf("variable \"bootstrap_script_url\" { default = \"%s\" }\n", bootstrapScriptUrl)

	// To use default patching strategy (nodes will be patched and rebooted AD/FD by AD/FD), comment patching_configs & above config representation
	var (
		BdsBdsInstanceExecuteBootstrapScriptActionRepresentation = map[string]interface{}{
			"bds_instance_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.bds_instance_ocid}`},
			"cluster_admin_password": acctest.Representation{RepType: acctest.Required, Create: `T3JhY2xlVGVhbVVTQSExMjM=`},
			"bootstrap_script_url":   acctest.Representation{RepType: acctest.Optional, Create: `${var.bootstrap_script_url}`},
		}
	)

	resourceName := "oci_bds_bds_instance_execute_bootstrap_script_action.test_bds_instance_execute_bootstrap_script_action"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+subnetIdVariableStr+bdsInstanceIdVariableStr+bootstrapScriptUrlVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_execute_bootstrap_script_action", "test_bds_instance_execute_bootstrap_script_action", acctest.Required, acctest.Create, BdsBdsInstanceExecuteBootstrapScriptActionRepresentation), "bds", "bdsInstanceExecuteBootstrapScriptAction", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + bdsInstanceIdVariableStr + bootstrapScriptUrlVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_execute_bootstrap_script_action", "test_bds_instance_execute_bootstrap_script_action", acctest.Optional, acctest.Create, BdsBdsInstanceExecuteBootstrapScriptActionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "T3JhY2xlVGVhbVVTQSExMjM="),
				resource.TestCheckResourceAttr(resourceName, "bootstrap_script_url", bootstrapScriptUrl),
			),
		},
	})

}
