// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

// issue-routing-tag: object_storage/default
func TestResourceNamespaceMetadata_basic(t *testing.T) {
	httpreplay.SetScenario("TestObjectStorageNamespaceMetadataResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetRequiredEnvSetting("compartment_ocid")

	resourceName := "oci_objectstorage_namespace_metadata.test_namespace_metadata"
	datasourceName := "data.oci_objectstorage_namespace_metadata.test_namespace_metadata"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: config + `
data "oci_objectstorage_namespace" "t" {
}

resource "oci_objectstorage_namespace_metadata" "test_namespace_metadata" {
	namespace = "${data.oci_objectstorage_namespace.t.namespace}"
}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "default_s3compartment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "default_swift_compartment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
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
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "default_s3compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "default_swift_compartment_id", compartmentId),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(

					resource.TestCheckResourceAttrSet(datasourceName, "default_s3compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "default_swift_compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "namespace"),
				),
			},
		},
	})
}
