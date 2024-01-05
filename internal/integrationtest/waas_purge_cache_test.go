// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	WaasPurgeCacheRequiredOnlyResource = WaasPurgeCacheResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_waas_purge_cache", "test_purge_cache", acctest.Required, acctest.Create, WaasPurgeCacheRepresentation)

	WaasPurgeCacheRepresentation = map[string]interface{}{
		"waas_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_waas_waas_policy.test_scenario_waas_policy.id}`},
		"resources":      acctest.Representation{RepType: acctest.Optional, Create: []string{`/about`, `/home`}},
	}

	WaasPurgeCacheResourceDependencies = WaasPolicyResourceCachingOnlyConfig
)

// issue-routing-tag: waas/default
func TestWaasPurgeCacheResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWaasPurgeCacheResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_waas_purge_cache.test_purge_cache"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+WaasPurgeCacheResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_waas_purge_cache", "test_purge_cache", acctest.Optional, acctest.Create, WaasPurgeCacheRepresentation), "waas", "purgeCache", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify purge select resources
		{
			Config: config + compartmentIdVariableStr + WaasPurgeCacheResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waas_purge_cache", "test_purge_cache", acctest.Optional, acctest.Create, WaasPurgeCacheRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "resources.#", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "waas_policy_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + WaasPurgeCacheResourceDependencies,
		},
		// verify purge all resources
		{
			Config: config + compartmentIdVariableStr + WaasPurgeCacheResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waas_purge_cache", "test_purge_cache", acctest.Required, acctest.Create, WaasPurgeCacheRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "waas_policy_id"),

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
