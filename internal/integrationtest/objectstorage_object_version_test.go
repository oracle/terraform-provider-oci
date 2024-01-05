// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	testVersioningBucketName = utils.RandomStringOrHttpReplayValue(32, utils.Charset, "bucketVersioning")

	ObjectStorageObjectStorageObjectVersionDataSourceRepresentation = map[string]interface{}{
		"bucket":      acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"namespace":   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"delimiter":   acctest.Representation{RepType: acctest.Optional, Create: `/`},
		"end":         acctest.Representation{RepType: acctest.Optional, Create: `z`},
		"prefix":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_objectstorage_object.test_object.object}`},
		"start":       acctest.Representation{RepType: acctest.Optional, Create: `${oci_objectstorage_object.test_object.object}`},
		"start_after": acctest.Representation{RepType: acctest.Optional, Create: `a`},
	}

	ObjectStorageObjectVersionResourceConfig = ObjectStorageBucketResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(ObjectStorageBucketRepresentation, map[string]interface{}{"name": acctest.Representation{RepType: acctest.Required, Create: testVersioningBucketName}, "versioning": acctest.Representation{RepType: acctest.Required, Create: `Enabled`}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Optional, acctest.Update, ObjectStorageObjectRepresentation)
)

// issue-routing-tag: object_storage/default
func TestObjectStorageObjectVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestObjectStorageObjectVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_objectstorage_object_versions.test_object_versions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_object_versions", "test_object_versions", acctest.Optional, acctest.Create, ObjectStorageObjectStorageObjectVersionDataSourceRepresentation) +
				compartmentIdVariableStr + ObjectStorageObjectVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
