// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	testVersioningBucketName = RandomStringOrHttpReplayValue(32, Charset, "bucketVersioning")

	objectVersionDataSourceRepresentation = map[string]interface{}{
		"bucket":      Representation{RepType: Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"namespace":   Representation{RepType: Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"delimiter":   Representation{RepType: Optional, Create: `/`},
		"end":         Representation{RepType: Optional, Create: `z`},
		"prefix":      Representation{RepType: Optional, Create: `${oci_objectstorage_object.test_object.object}`},
		"start":       Representation{RepType: Optional, Create: `${oci_objectstorage_object.test_object.object}`},
		"start_after": Representation{RepType: Optional, Create: `a`},
	}

	ObjectVersionResourceConfig = BucketResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Required, Create, RepresentationCopyWithNewProperties(bucketRepresentation, map[string]interface{}{"name": Representation{RepType: Required, Create: testVersioningBucketName}, "versioning": Representation{RepType: Required, Create: `Enabled`}})) +
		GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", Optional, Update, objectRepresentation)
)

// issue-routing-tag: object_storage/default
func TestObjectStorageObjectVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestObjectStorageObjectVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_objectstorage_object_versions.test_object_versions"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_objectstorage_object_versions", "test_object_versions", Optional, Create, objectVersionDataSourceRepresentation) +
				compartmentIdVariableStr + ObjectVersionResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "bucket", testVersioningBucketName),
				resource.TestCheckResourceAttr(datasourceName, "delimiter", "/"),
				resource.TestCheckResourceAttr(datasourceName, "end", "z"),
				resource.TestCheckResourceAttrSet(datasourceName, "namespace"),
				resource.TestCheckResourceAttr(datasourceName, "prefix", "my-test-object-2"),
				resource.TestCheckResourceAttr(datasourceName, "start", "my-test-object-2"),
				resource.TestCheckResourceAttr(datasourceName, "start_after", "a"),

				resource.TestCheckResourceAttr(datasourceName, "items.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "items.0.storage_tier", "InfrequentAccess"),
			),
		},
	})
}
