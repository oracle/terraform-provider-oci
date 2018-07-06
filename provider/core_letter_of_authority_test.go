// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"strings"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	LetterOfAuthorityResourceConfig = LetterOfAuthorityResourceDependencies + `

`
	LetterOfAuthorityPropertyVariables = `

`
	LetterOfAuthorityResourceDependencies = CrossConnectPropertyVariables + CrossConnectRequiredOnlyResource
)

func TestCoreLetterOfAuthorityResource_basic(t *testing.T) {
	region := getRequiredEnvSetting("region")
	if !strings.EqualFold("r1", region) {
		t.Skip("FastConnect tests are not yet supported in production regions")
	}

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_letter_of_authority.test_letter_of_authority"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config + `

data "oci_core_letter_of_authority" "test_letter_of_authority" {
	#Required
	cross_connect_id = "${oci_core_cross_connect.test_cross_connect.id}"
}
                ` + compartmentIdVariableStr + LetterOfAuthorityResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "cross_connect_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "authorized_entity_name", "internalsdk"),
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
