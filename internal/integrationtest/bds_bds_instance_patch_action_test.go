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
	BdsBdsInstancePatchActionRepresentation = map[string]interface{}{
		"bds_instance_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"cluster_admin_password": acctest.Representation{RepType: acctest.Required, Create: `V2VsY29tZTE=`},
		"version":                acctest.Representation{RepType: acctest.Required, Create: "ODH-1.1.0.379"},
	}

	BdsBdsInstancePatchActionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Required, acctest.Create, bdsInstanceOdhRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation)
)

// issue-routing-tag: bds/default
func TestBdsBdsInstancePatchActionResource_basic(t *testing.T) {
	t.Skip("Run manual with an older cluster with patch available")
	httpreplay.SetScenario("TestBdsBdsInstancePatchActionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_bds_bds_instance_patch_action.test_bds_instance_patch_action"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+BdsBdsInstancePatchActionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_patch_action", "test_bds_instance_patch_action", acctest.Required, acctest.Create, BdsBdsInstancePatchActionRepresentation), "bds", "bdsInstancePatchAction", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + BdsBdsInstancePatchActionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_patch_action", "test_bds_instance_patch_action", acctest.Required, acctest.Create, BdsBdsInstancePatchActionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
			),
		},
	})
}
