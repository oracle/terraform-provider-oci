// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	letterOfAuthoritySingularDataSourceRepresentation = map[string]interface{}{
		"cross_connect_id": Representation{repType: Required, create: `${oci_core_cross_connect.test_cross_connect.id}`},
	}

	LetterOfAuthorityResourceConfig = CrossConnectResourceConfig
)

func TestCoreLetterOfAuthorityResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_letter_of_authority.test_letter_of_authority"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_letter_of_authority", "test_letter_of_authority", Required, Create, letterOfAuthoritySingularDataSourceRepresentation) +
					compartmentIdVariableStr + LetterOfAuthorityResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cross_connect_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "authorized_entity_name"),
					resource.TestCheckResourceAttr(singularDatasourceName, "circuit_type", "Single_mode_LC"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "facility_location"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "port_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_expires"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_issued"),
				),
			},
		},
	})
}
