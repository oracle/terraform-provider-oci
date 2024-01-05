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
	IdentityDomainsIdentityDomainsMyAppDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"my_app_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"my_app_filter": acctest.Representation{RepType: acctest.Optional, Create: ``},
		"start_index":   acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsMyAppResourceConfig = TestDomainForMyEndpointDependencies + GenerateMyAppTestApp("true") + MyAppTestGrantDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsMyAppResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsMyAppResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_identity_domains_my_apps.test_my_apps"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// create an App and assign to my user
		{
			Config: config + compartmentIdVariableStr + IdentityDomainsMyAppResourceConfig,
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_apps", "test_my_apps", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsMyAppDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsMyAppResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestCheckResourceAttrSet(datasourceName, "my_apps.#"),
				resource.TestCheckResourceAttr(datasourceName, "my_apps.0.schemas.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "total_results"),
			),
		},
		// clean up dependencies, set test app to inactive first before deletion
		{
			Config: config + compartmentIdVariableStr + TestDomainForMyEndpointDependencies + GenerateMyAppTestApp("false"),
		},
	})
}
