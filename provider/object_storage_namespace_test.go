// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestObjectStorageNamespaceResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	singularDatasourceName := "data.oci_objectstorage_namespace.test_namespace"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config + `

data "oci_objectstorage_namespace" "test_namespace" {
}
                `,
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
				),
			},
		},
	})
}
