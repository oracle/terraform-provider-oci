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
func TestBdsBdsInstanceReplaceNodeActionResource(t *testing.T) {
	//t.Skip("Run manual with an older cluster with patch available")
	httpreplay.SetScenario("TestBdsBdsInstanceReplaceNodeActionResource")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_ocid")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	bdsInstanceId := utils.GetEnvSettingWithBlankDefault("bds_instance_ocid")
	bdsInstanceIdVariableStr := fmt.Sprintf("variable \"bds_instance_id\" { default = \"%s\" }\n", bdsInstanceId)

	nodeBackupId := utils.GetEnvSettingWithBlankDefault("node_backup_ocid")
	nodeBackupIdVariableStr := fmt.Sprintf("variable \"node_backup_id\" { default = \"%s\" }\n", nodeBackupId)

	nodeHostName := utils.GetEnvSettingWithBlankDefault("node_host_name")
	nodeHostNameVariableStr := fmt.Sprintf("variable \"node_host_name\" { default = \"%s\" }\n", nodeHostName)

	// To use default patching strategy (nodes will be patched and rebooted AD/FD by AD/FD), comment patching_configs & above config representation
	var (
		BdsBdsInstanceReplaceNodeActionRepresentation = map[string]interface{}{
			"cluster_admin_password": acctest.Representation{RepType: acctest.Required, Create: `T3JhY2xlVGVhbVVTQSExMjM=`},
			"bds_instance_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.bds_instance_id}`},
			"node_backup_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.node_backup_id}`},
			"node_host_name":         acctest.Representation{RepType: acctest.Required, Create: `${var.node_host_name}`},
		}

		BdsBdsInstanceReplaceNodeActionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Required, acctest.Create, bdsInstanceOdhRepresentation) +
			acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
			acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation)
	)

	resourceName := "oci_bds_bds_instance_replace_node_action.test_bds_instance_replace_node_action"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+subnetIdVariableStr+BdsBdsInstanceReplaceNodeActionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_replace_node_action", "test_bds_instance_replace_node_action", acctest.Required, acctest.Create, BdsBdsInstanceReplaceNodeActionRepresentation), "bds", "bdsInstanceReplaceNodeAction", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + bdsInstanceIdVariableStr + nodeBackupIdVariableStr + nodeHostNameVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_replace_node_action", "test_bds_instance_replace_node_action", acctest.Required, acctest.Create, BdsBdsInstanceReplaceNodeActionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "T3JhY2xlVGVhbVVTQSExMjM="),
			),
		},
	})

}
