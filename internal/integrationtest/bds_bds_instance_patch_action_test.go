// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	//BdsBdsInstancePatchActionRequiredOnlyResource = BdsBdsInstancePatchActionResourceDependencies +
	//acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_patch_action", "test_bds_instance_patch_action", acctest.Required, acctest.Create, BdsBdsInstancePatchActionRepresentation)

	BdsBdsInstancePatchActionRepresentation = map[string]interface{}{
		"bds_instance_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_bds_bds_instance.test_bds_instance.id}`},
		"cluster_admin_password": acctest.Representation{RepType: acctest.Required, Create: `clusterAdminPassword`},
		"version":                acctest.Representation{RepType: acctest.Required, Create: `version`},
		"patching_config":        acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsBdsInstancePatchActionPatchingConfigRepresentation},
		"timeouts":               acctest.RepresentationGroup{RepType: acctest.Required, Group: PatchTimeoutsRepresentation},
	}
	BdsBdsInstancePatchActionPatchingConfigRepresentation = map[string]interface{}{
		"patching_config_strategy":            acctest.Representation{RepType: acctest.Required, Create: `DOWNTIME_BASED`},
		"batch_size":                          acctest.Representation{RepType: acctest.Required, Create: `3`},
		"wait_time_between_batch_in_seconds":  acctest.Representation{RepType: acctest.Required, Create: `600`},
		"wait_time_between_domain_in_seconds": acctest.Representation{RepType: acctest.Required, Create: `600`},
	}

	PatchTimeoutsRepresentation = map[string]interface{}{
		"create": acctest.Representation{RepType: acctest.Required, Create: `24h`},
		"update": acctest.Representation{RepType: acctest.Required, Create: `24h`},
		"delete": acctest.Representation{RepType: acctest.Required, Create: `24h`},
	}

	BdsBdsInstancePatchActionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance", "test_bds_instance", acctest.Required, acctest.Create, bdsInstanceOdhRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation)
)

// issue-routing-tag: bds/default
func TestBdsBdsInstancePatchActionResource_basic(t *testing.T) {
	//t.Skip("Run manual with an older cluster with patch available")
	httpreplay.SetScenario("TestBdsBdsInstancePatchActionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_bds_bds_instance_patch_action.test_bds_instance_patch_action"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.

	acctest.SaveConfigContent(config+compartmentIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_patch_action", "test_bds_instance_patch_action", acctest.Optional, acctest.Create, BdsBdsInstancePatchActionRepresentation), "bds", "bdsInstancePatchAction", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_patch_action", "test_bds_instance_patch_action", acctest.Required, acctest.Create, BdsBdsInstancePatchActionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BdsBdsInstancePatchActionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + BdsBdsInstancePatchActionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_patch_action", "test_bds_instance_patch_action", acctest.Optional, acctest.Create, BdsBdsInstancePatchActionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "clusterAdminPassword"),
				resource.TestCheckResourceAttr(resourceName, "patching_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "patching_config.0.batch_size", "3"),
				resource.TestCheckResourceAttr(resourceName, "patching_config.0.patching_config_strategy", "DOWNTIME_BASED"),
				resource.TestCheckResourceAttr(resourceName, "patching_config.0.wait_time_between_batch_in_seconds", "600"),
				resource.TestCheckResourceAttr(resourceName, "patching_config.0.wait_time_between_domain_in_seconds", "600"),
				resource.TestCheckResourceAttr(resourceName, "version", "version"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
	})
}
