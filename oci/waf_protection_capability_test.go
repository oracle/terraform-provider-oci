// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	protectionCapabilityDataSourceRepresentation = map[string]interface{}{
		"compartment_id":    Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":      Representation{RepType: Optional, Create: `SQL Injection (SQLi) basic SQL auth bypass attempts`},
		"group_tag":         Representation{RepType: Optional, Create: []string{`OWASP`}},
		"is_latest_version": Representation{RepType: Optional, Create: []string{`true`}},
		"key":               Representation{RepType: Optional, Create: `942340`},
		"type":              Representation{RepType: Optional, Create: `REQUEST_PROTECTION_CAPABILITY`},
	}

	ProtectionCapabilityResourceConfig = ""
)

// issue-routing-tag: waf/default
func TestWafProtectionCapabilityResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWafProtectionCapabilityResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_waf_protection_capabilities.test_protection_capabilities"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_waf_protection_capabilities", "test_protection_capabilities", Optional, Create, protectionCapabilityDataSourceRepresentation) +
				compartmentIdVariableStr + ProtectionCapabilityResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
