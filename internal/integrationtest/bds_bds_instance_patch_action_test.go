// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	BdsBdsInstancePatchActionRepresentation = map[string]interface{}{
		"bds_instance_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.bdsinstance_id}`},
		"cluster_admin_password": acctest.Representation{RepType: acctest.Required, Create: `T3JhY2xlVGVhbVVTQSExMjM=`},
		"version":                acctest.Representation{RepType: acctest.Required, Create: `ODH-2.0.10.7-branch-ODH2-102024-7.tar.gz`},
		"patching_config":        acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsBdsInstancePatchActionPatchingConfigRepresentation},
		"timeouts":               acctest.RepresentationGroup{RepType: acctest.Required, Group: PatchTimeoutsRepresentation},
	}

	BdsBdsInstancePatchActionPatchingConfigRepresentation = map[string]interface{}{
		"patching_config_strategy":            acctest.Representation{RepType: acctest.Required, Create: `DOMAIN_BASED`},
		"tolerance_threshold_per_batch":       acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"tolerance_threshold_per_domain":      acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"batch_size":                          acctest.Representation{RepType: acctest.Required, Create: `5`},
		"wait_time_between_batch_in_seconds":  acctest.Representation{RepType: acctest.Required, Create: `0`},
		"wait_time_between_domain_in_seconds": acctest.Representation{RepType: acctest.Required, Create: `0`},
	}

	PatchTimeoutsRepresentation = map[string]interface{}{
		"create": acctest.Representation{RepType: acctest.Required, Create: `24h`},
		"update": acctest.Representation{RepType: acctest.Required, Create: `24h`},
		"delete": acctest.Representation{RepType: acctest.Required, Create: `24h`},
	}
)

// issue-routing-tag: bds/default
func TestBdsBdsInstancePatchActionResource_basic(t *testing.T) {
	//t.Skip("Run manual with an older cluster with patch available")
	httpreplay.SetScenario("TestBdsBdsInstancePatchActionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	bdsinstanceId := utils.GetEnvSettingWithBlankDefault("bdsinstance_ocid")
	bdsinstanceIdVariableStr := fmt.Sprintf("variable \"bdsinstance_id\" { default = \"%s\" }\n", bdsinstanceId)

	resourceName := "oci_bds_bds_instance_patch_action.test_bds_instance_patch_action"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.

	acctest.SaveConfigContent(config+compartmentIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_patch_action", "test_bds_instance_patch_action", acctest.Optional, acctest.Create, BdsBdsInstancePatchActionRepresentation), "bds", "bdsInstancePatchAction", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{

		// verify Create
		{
			Config: config + compartmentIdVariableStr + bdsinstanceIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_instance_patch_action", "test_bds_instance_patch_action", acctest.Optional, acctest.Create, BdsBdsInstancePatchActionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "bds_instance_id"),
				resource.TestCheckResourceAttr(resourceName, "cluster_admin_password", "T3JhY2xlVGVhbVVTQSExMjM="),
				resource.TestCheckResourceAttr(resourceName, "patching_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "patching_config.0.batch_size", "5"),
				resource.TestCheckResourceAttr(resourceName, "patching_config.0.patching_config_strategy", "DOMAIN_BASED"),
				resource.TestCheckResourceAttr(resourceName, "patching_config.0.tolerance_threshold_per_batch", "1"),
				resource.TestCheckResourceAttr(resourceName, "patching_config.0.tolerance_threshold_per_domain", "1"),
				resource.TestCheckResourceAttr(resourceName, "patching_config.0.wait_time_between_batch_in_seconds", "0"),
				resource.TestCheckResourceAttr(resourceName, "patching_config.0.wait_time_between_domain_in_seconds", "0"),
				resource.TestCheckResourceAttr(resourceName, "version", "ODH-2.0.10.7-branch-ODH2-102024-7.tar.gz"),

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
