// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	NamespaceResourceConfig = NamespaceResourceDependencies + `
resource "oci_object_storage_namespace" "test_namespace" {
}
`
	NamespacePropertyVariables = `

`
	NamespaceResourceDependencies = ""
)

func TestObjectStorageNamespaceResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_object_storage_namespace.test_namespace"
	datasourceName := "data.oci_object_storage_namespaces.test_namespaces"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config + NamespacePropertyVariables + compartmentIdVariableStr + NamespaceResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to Force New parameters.
			{
				Config: config + `

                ` + compartmentIdVariableStr2 + NamespaceResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated but it wasn't.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config + `

data "oci_object_storage_namespaces" "test_namespaces" {

    filter {
    	name = "id"
    	values = ["${oci_object_storage_namespace.test_namespace.id}"]
    }
}
                ` + compartmentIdVariableStr2 + NamespaceResourceConfig,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttr(datasourceName, "namespaces.#", "1"),
				),
			},
		},
	})
}

func TestObjectStorageNamespaceResource_forcenew(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_object_storage_namespace.test_namespace"

	var resId string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + NamespacePropertyVariables + compartmentIdVariableStr + NamespaceResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// force new tests, test that changing a parameter would result in creation of a new resource.

		},
	})
}
