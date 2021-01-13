// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	testVersioningBucketName = randomStringOrHttpReplayValue(32, charset, "bucket")

	objectVersionDataSourceRepresentation = map[string]interface{}{
		"bucket":      Representation{repType: Required, create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"namespace":   Representation{repType: Required, create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"delimiter":   Representation{repType: Optional, create: `/`},
		"end":         Representation{repType: Optional, create: `z`},
		"prefix":      Representation{repType: Optional, create: `${oci_objectstorage_object.test_object.object}`},
		"start":       Representation{repType: Optional, create: `${oci_objectstorage_object.test_object.object}`},
		"start_after": Representation{repType: Optional, create: `a`},
	}

	ObjectVersionResourceConfig = BucketResourceDependencies +
		generateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Required, Create, representationCopyWithNewProperties(bucketRepresentation, map[string]interface{}{"name": Representation{repType: Required, create: testVersioningBucketName}, "versioning": Representation{repType: Required, create: `Enabled`}})) +
		generateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", Optional, Update, objectRepresentation)
)

func TestObjectStorageObjectVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestObjectStorageObjectVersionResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_objectstorage_object_versions.test_object_versions"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_objectstorage_object_versions", "test_object_versions", Optional, Create, objectVersionDataSourceRepresentation) +
					compartmentIdVariableStr + ObjectVersionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "bucket", testVersioningBucketName),
					resource.TestCheckResourceAttr(datasourceName, "delimiter", "/"),
					resource.TestCheckResourceAttr(datasourceName, "end", "z"),
					resource.TestCheckResourceAttrSet(datasourceName, "namespace"),
					resource.TestCheckResourceAttr(datasourceName, "prefix", "my-test-object-2"),
					resource.TestCheckResourceAttr(datasourceName, "start", "my-test-object-2"),
					resource.TestCheckResourceAttr(datasourceName, "start_after", "a"),

					resource.TestCheckResourceAttr(datasourceName, "items.#", "1"),
				),
			},
		},
	})
}
