// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestObjectStorageNamespaceMetadataResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")

	resourceName := "oci_objectstorage_namespace_metadata.test_namespace_metadata"
	datasourceName := "data.oci_objectstorage_namespace_metadata.test_namespace_metadata"

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
				Config: config + `
data "oci_objectstorage_namespace" "t" {
}

resource "oci_objectstorage_namespace_metadata" "test_namespace_metadata" {
	namespace = "${data.oci_objectstorage_namespace.t.namespace}"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "default_s3compartment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "default_swift_compartment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
data "oci_objectstorage_namespace" "t" {
}

resource "oci_objectstorage_namespace_metadata" "test_namespace_metadata" {
	namespace = "${data.oci_objectstorage_namespace.t.namespace}"
  	default_s3compartment_id = "` + compartmentId + `"
  	default_swift_compartment_id = "` + compartmentId + `"
}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "default_s3compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "default_swift_compartment_id", compartmentId),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config + `
data "oci_objectstorage_namespace" "t" {
}

resource "oci_objectstorage_namespace_metadata" "test_namespace_metadata" {
	namespace = "${data.oci_objectstorage_namespace.t.namespace}"
  	default_s3compartment_id = "` + compartmentId + `"
  	default_swift_compartment_id = "` + compartmentId + `"
}

data "oci_objectstorage_namespace_metadata" "test_namespace_metadata" {
	namespace = "${data.oci_objectstorage_namespace.t.namespace}"
}
                `,
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttrSet(datasourceName, "default_s3compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "default_swift_compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "namespace"),
				),
			},
		},
	})
}
