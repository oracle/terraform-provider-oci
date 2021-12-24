// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	protectionCapabilityDataSourceRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":      acctest.Representation{RepType: acctest.Optional, Create: `SQL Injection (SQLi) basic SQL auth bypass attempts`},
		"group_tag":         acctest.Representation{RepType: acctest.Optional, Create: []string{`OWASP`}},
		"is_latest_version": acctest.Representation{RepType: acctest.Optional, Create: []string{`true`}},
		"key":               acctest.Representation{RepType: acctest.Optional, Create: `942340`},
		"type":              acctest.Representation{RepType: acctest.Optional, Create: `REQUEST_PROTECTION_CAPABILITY`},
	}

	ProtectionCapabilityResourceConfig = ""
)

// issue-routing-tag: waf/default
func TestWafProtectionCapabilityResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWafProtectionCapabilityResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_waf_protection_capabilities.test_protection_capabilities"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_waf_protection_capabilities", "test_protection_capabilities", acctest.Optional, acctest.Create, protectionCapabilityDataSourceRepresentation) +
				compartmentIdVariableStr + ProtectionCapabilityResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "SQL Injection (SQLi) basic SQL auth bypass attempts"),
				resource.TestCheckResourceAttr(datasourceName, "group_tag.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "is_latest_version.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "key", "942340"),
				resource.TestCheckResourceAttr(datasourceName, "type", "REQUEST_PROTECTION_CAPABILITY"),

				resource.TestCheckResourceAttrSet(datasourceName, "protection_capability_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "protection_capability_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "protection_capability_collection.0.items.0.key", "942340"),
			),
		},
	})
}
