// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityDomainsIdentityDomainsMyGroupDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"my_group_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"my_group_filter": acctest.Representation{RepType: acctest.Optional, Create: ``},
		"attribute_sets":  acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"start_index":     acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsMyGroupResourceConfig = TestDomainDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsMyGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsMyGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_identity_domains_my_groups.test_my_groups"

	acctest.SaveConfigContent("", "", "", t)

	print(config +
		acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_groups", "test_my_groups", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsMyGroupDataSourceRepresentation) +
		compartmentIdVariableStr + IdentityDomainsMyGroupResourceConfig)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_groups", "test_my_groups", acctest.Required, acctest.Create, IdentityDomainsIdentityDomainsMyGroupDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsMyGroupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestMatchResourceAttr(datasourceName, "my_groups.#", regexp.MustCompile("[1-9]+")),
				resource.TestCheckResourceAttr(datasourceName, "my_groups.0.meta.0.resource_type", "MyGroup"),
			),
		},
	})
}
