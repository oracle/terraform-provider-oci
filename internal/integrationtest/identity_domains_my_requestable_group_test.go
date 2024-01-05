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
	IdentityDomainsIdentityDomainsMyRequestableGroupDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":               acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"my_requestable_group_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"my_requestable_group_filter": acctest.Representation{RepType: acctest.Optional, Create: `id eq \"${oci_identity_domains_group.test_group.id}\"`},
		"start_index":                 acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsMyRequestableGroupResourceConfig = TestDomainForMyEndpointDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domains_group", "test_group", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(IdentityDomainsGroupRepresentation, map[string]interface{}{"attribute_sets": acctest.Representation{RepType: acctest.Required, Create: []string{`all`}}, "urnietfparamsscimschemasoracleidcsextensionrequestable_group": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{"requestable": acctest.Representation{RepType: acctest.Required, Create: `true`}}}}))
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsMyRequestableGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsMyRequestableGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_identity_domains_my_requestable_groups.test_my_requestable_groups"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_my_requestable_groups", "test_my_requestable_groups", acctest.Optional, acctest.Create, IdentityDomainsIdentityDomainsMyRequestableGroupDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsMyRequestableGroupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "my_requestable_group_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestCheckResourceAttrSet(datasourceName, "my_requestable_groups.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "total_results"),
			),
		},
	})
}
