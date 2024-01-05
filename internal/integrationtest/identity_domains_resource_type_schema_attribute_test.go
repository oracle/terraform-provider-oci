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
	IdentityDomainsIdentityDomainsResourceTypeSchemaAttributeDataSourceRepresentation = map[string]interface{}{
		"idcs_endpoint":                         acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_domain.test_domain.url}`},
		"resource_type_schema_attribute_count":  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"resource_type_schema_attribute_filter": acctest.Representation{RepType: acctest.Optional, Create: `resourcetype eq \"User\" and mutability eq \"readWrite\" and idcsSearchable eq true`},
		"attribute_sets":                        acctest.Representation{RepType: acctest.Optional, Create: []string{`all`}},
		"start_index":                           acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	IdentityDomainsResourceTypeSchemaAttributeResourceConfig = TestDomainDependencies
)

// issue-routing-tag: identity_domains/default
func TestIdentityDomainsResourceTypeSchemaAttributeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityDomainsResourceTypeSchemaAttributeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_identity_domains_resource_type_schema_attributes.test_resource_type_schema_attributes"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_domains_resource_type_schema_attributes", "test_resource_type_schema_attributes", acctest.Optional, acctest.Create, IdentityDomainsIdentityDomainsResourceTypeSchemaAttributeDataSourceRepresentation) +
				compartmentIdVariableStr + IdentityDomainsResourceTypeSchemaAttributeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "resource_type_schema_attribute_count", "10"),
				resource.TestCheckResourceAttr(datasourceName, "resource_type_schema_attribute_filter", "resourcetype eq \"User\" and mutability eq \"readWrite\" and idcsSearchable eq true"),
				resource.TestCheckResourceAttr(datasourceName, "attribute_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "start_index", "1"),

				resource.TestCheckResourceAttrSet(datasourceName, "resource_type_schema_attributes.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "items_per_page"),
				resource.TestCheckResourceAttr(datasourceName, "schemas.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "start_index"),
				resource.TestCheckResourceAttrSet(datasourceName, "total_results"),
			),
		},
	})
}
