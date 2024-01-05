// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"encoding/hex"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"testing"

	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_object_storage "github.com/oracle/oci-go-sdk/v65/objectstorage"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	ObjectStorageObjectRequiredOnlyResource = ObjectStorageObjectResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Required, acctest.Create, ObjectStorageObjectRepresentation)

	ObjectStorageObjectResourceConfig = ObjectStorageObjectResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Optional, acctest.Create, ObjectStorageObjectRepresentation)

	ObjectStorageObjectStorageObjectDataSourceRepresentation = map[string]interface{}{
		"bucket":      acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"namespace":   acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.namespace}`},
		"delimiter":   acctest.Representation{RepType: acctest.Optional, Create: `delimiter`, Update: `/`},
		"end":         acctest.Representation{RepType: acctest.Optional, Create: `x`},
		"prefix":      acctest.Representation{RepType: acctest.Optional, Create: `prefix`, Update: `my-test`},
		"start":       acctest.Representation{RepType: acctest.Optional, Create: `${oci_objectstorage_object.test_object.object}`},
		"start_after": acctest.Representation{RepType: acctest.Optional, Create: `a`},
		"filter":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ObjectStorageObjectDataSourceFilterRepresentation}}
	ObjectStorageObjectDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_objectstorage_object.test_object.object}`}},
	}

	objectSingularDataSourceRepresentation = map[string]interface{}{
		"bucket":                            acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"namespace":                         acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.namespace}`},
		"object":                            acctest.Representation{RepType: acctest.Required, Create: `my-test-object-3`},
		"content_length_limit":              acctest.Representation{RepType: acctest.Optional, Create: `17`, Update: `16`},
		"base64_encode_content":             acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"version_id":                        acctest.Representation{RepType: acctest.Optional, Create: `${oci_objectstorage_object.test_object.version_id}`},
		"http_response_cache_control":       acctest.Representation{RepType: acctest.Optional, Create: `no-cache`, Update: `no-store`},
		"http_response_content_disposition": acctest.Representation{RepType: acctest.Optional, Create: `inline`, Update: `inline`},
		"http_response_content_encoding":    acctest.Representation{RepType: acctest.Optional, Create: `identity`, Update: `identity`},
		"http_response_content_language":    acctest.Representation{RepType: acctest.Optional, Create: `en-US`, Update: `en-US`},
		"http_response_content_type":        acctest.Representation{RepType: acctest.Optional, Create: `text/plain`, Update: `text/plain`},
		"http_response_expires":             acctest.Representation{RepType: acctest.Optional, Create: expirationTimeForPar.Format(time.RFC3339Nano), Update: expirationTimeForPar.Format(time.RFC3339Nano)},
	}

	ObjectStorageObjectRepresentation = map[string]interface{}{
		"bucket":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"content":                    acctest.Representation{RepType: acctest.Optional, Create: `content`, Update: `<a1>content</a1>`},
		"namespace":                  acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.namespace}`},
		"object":                     acctest.Representation{RepType: acctest.Required, Create: `my-test-object-1`, Update: `my-test-object-2`},
		"cache_control":              acctest.Representation{RepType: acctest.Optional, Create: `no-cache`, Update: `no-store`},
		"content_disposition":        acctest.Representation{RepType: acctest.Optional, Create: `inline`, Update: `attachment; filename=\"filename.html\"`},
		"content_encoding":           acctest.Representation{RepType: acctest.Optional, Create: `identity`},
		"content_language":           acctest.Representation{RepType: acctest.Optional, Create: `en-US`, Update: `en-CA`},
		"content_md5":                acctest.Representation{RepType: acctest.Optional, Create: `${md5("content")}`, Update: Md5Base64Encoded2},
		"content_type":               acctest.Representation{RepType: acctest.Optional, Create: `text/plain`, Update: `text/xml`},
		"storage_tier":               acctest.Representation{RepType: acctest.Optional, Create: `Standard`, Update: `InfrequentAccess`},
		"opc_sse_kms_key_id":         acctest.Representation{RepType: acctest.Optional, Create: utils.GetEnvSettingWithBlankDefault("kms_key_ocid")},
		"delete_all_object_versions": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"metadata":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"content-type": "text/plain"}, Update: map[string]string{"content-type": "text/xml"}},
	}

	objectEmptyRepresentation = map[string]interface{}{
		"bucket":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"content":                    acctest.Representation{RepType: acctest.Optional, Create: ``},
		"namespace":                  acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.namespace}`},
		"object":                     acctest.Representation{RepType: acctest.Required, Create: `my-test-empty-object`},
		"cache_control":              acctest.Representation{RepType: acctest.Optional, Create: `no-cache`},
		"content_disposition":        acctest.Representation{RepType: acctest.Optional, Create: `inline`},
		"content_encoding":           acctest.Representation{RepType: acctest.Optional, Create: `identity`},
		"content_language":           acctest.Representation{RepType: acctest.Optional, Create: `en-US`},
		"content_type":               acctest.Representation{RepType: acctest.Optional, Create: `text/plain`},
		"delete_all_object_versions": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"metadata":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"content-type": "text/plain"}},
	}

	ObjectStorageObjectResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, ObjectStorageBucketRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation)

	Md5Base64Encoded2, _ = tfresource.HexToB64(tfresource.GetMd5Hash("<a1>content</a1>"))
)

// issue-routing-tag: object_storage/default
func TestObjectStorageObjectResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestObjectStorageObjectResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_objectstorage_object.test_object"
	datasourceName := "data.oci_objectstorage_objects.test_objects"
	singularDatasourceName := "data.oci_objectstorage_object.test_object"

	var resId, resId2 string
	hexSum := md5.Sum([]byte("content"))
	md5sum := hex.EncodeToString(hexSum[:])
	hexSum2 := md5.Sum([]byte("<a1>content</a1>"))
	md5sum2 := hex.EncodeToString(hexSum2[:])
	md5B64Encode, _ := tfresource.HexToB64(md5sum)
	md5B64Encode2, _ := tfresource.HexToB64(md5sum2)

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ObjectStorageObjectResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Optional, acctest.Create, ObjectStorageObjectRepresentation), "objectstorage", "object", t)

	acctest.ResourceTest(t, testAccCheckObjectStorageObjectDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ObjectStorageObjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Required, acctest.Create, ObjectStorageObjectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bucket", testBucketName),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-1"),
				resource.TestCheckResourceAttr(resourceName, "content_type", "application/octet-stream"),
				// New SDK doesn't set omitted values from response, check they are missing from state.
				resource.TestCheckNoResourceAttr(resourceName, "content"),
				resource.TestCheckNoResourceAttr(resourceName, "content_language"),
				resource.TestCheckNoResourceAttr(resourceName, "content_encoding"),
				resource.TestCheckResourceAttr(resourceName, "content_length", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "content_md5"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ObjectStorageObjectResourceDependencies,
		},

		// verify Create empty
		{
			Config: config + compartmentIdVariableStr + ObjectStorageObjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Required, acctest.Create, objectEmptyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bucket", testBucketName),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "object", "my-test-empty-object"),
				resource.TestCheckResourceAttr(resourceName, "content_type", "application/octet-stream"),
				// New SDK doesn't set omitted values from response, check they are missing from state.
				resource.TestCheckNoResourceAttr(resourceName, "content"),
				resource.TestCheckNoResourceAttr(resourceName, "content_language"),
				resource.TestCheckNoResourceAttr(resourceName, "content_encoding"),
				resource.TestCheckResourceAttr(resourceName, "content_length", "0"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ObjectStorageObjectResourceDependencies,
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ObjectStorageObjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Optional, acctest.Create, ObjectStorageObjectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cache_control", "no-cache"),
				resource.TestCheckResourceAttr(resourceName, "content_disposition", "inline"),
				resource.TestCheckResourceAttr(resourceName, "content_encoding", "identity"),
				resource.TestCheckResourceAttr(resourceName, "content_language", "en-US"),
				resource.TestCheckResourceAttr(resourceName, "content_length", "7"),
				resource.TestCheckResourceAttr(resourceName, "content_md5", *md5B64Encode),
				resource.TestCheckResourceAttr(resourceName, "content_type", "text/plain"),
				resource.TestCheckResourceAttr(resourceName, "delete_all_object_versions", "false"),
				resource.TestCheckResourceAttr(resourceName, "bucket", testBucketName),
				resource.TestCheckResourceAttrSet(resourceName, "content"),
				resource.TestCheckResourceAttr(resourceName, "content", md5sum),
				resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "metadata.content-type", "text/plain"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-1"),
				resource.TestCheckResourceAttr(resourceName, "storage_tier", "Standard"),
				resource.TestCheckResourceAttr(resourceName, "opc_sse_kms_key_id", utils.GetEnvSettingWithBlankDefault("kms_key_ocid")),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ObjectStorageObjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Optional, acctest.Update, ObjectStorageObjectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cache_control", "no-store"),
				resource.TestCheckResourceAttr(resourceName, "content_disposition", "attachment; filename=\"filename.html\""),
				resource.TestCheckResourceAttr(resourceName, "content_encoding", "identity"),
				resource.TestCheckResourceAttr(resourceName, "content_language", "en-CA"),
				resource.TestCheckResourceAttr(resourceName, "content_length", "16"),
				resource.TestCheckResourceAttr(resourceName, "content_md5", *md5B64Encode2),
				resource.TestCheckResourceAttr(resourceName, "content_type", "text/xml"),
				resource.TestCheckResourceAttr(resourceName, "storage_tier", "InfrequentAccess"),
				resource.TestCheckResourceAttr(resourceName, "bucket", testBucketName),
				resource.TestCheckResourceAttr(resourceName, "delete_all_object_versions", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "content"),
				resource.TestCheckResourceAttr(resourceName, "content", md5sum2),
				resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "metadata.content-type", "text/xml"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-2"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					// @CODEGEN 06/2018: Name is part of the id, and hence id will be updated
					if resId == resId2 {
						return fmt.Errorf("Resource updated when it was supposed to be recreated.")
					}
					return err
				},
			),
		},
		// verify either a hex or a base64 equivalent content_md5 makes no diff
		{
			Config: config + compartmentIdVariableStr + ObjectStorageObjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Optional, acctest.Update,
					acctest.GetUpdatedRepresentationCopy("content_md5", acctest.Representation{RepType: acctest.Optional, Create: Md5Base64Encoded2, Update: `${md5("<a1>content</a1>")}`}, ObjectStorageObjectRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cache_control", "no-store"),
				resource.TestCheckResourceAttr(resourceName, "content_disposition", "attachment; filename=\"filename.html\""),
				resource.TestCheckResourceAttr(resourceName, "content_encoding", "identity"),
				resource.TestCheckResourceAttr(resourceName, "content_language", "en-CA"),
				resource.TestCheckResourceAttr(resourceName, "content_length", "16"),
				resource.TestCheckResourceAttr(resourceName, "content_md5", *md5B64Encode2),
				resource.TestCheckResourceAttr(resourceName, "content_type", "text/xml"),
				resource.TestCheckResourceAttr(resourceName, "bucket", testBucketName),
				resource.TestCheckResourceAttr(resourceName, "delete_all_object_versions", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "content"),
				resource.TestCheckResourceAttr(resourceName, "content", md5sum2),
				resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "metadata.content-type", "text/xml"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-2"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					// @CODEGEN 06/2018: Name is part of the id, and hence id will be updated
					if resId == resId2 {
						return fmt.Errorf("Resource updated when it was supposed to be recreated.")
					}
					return err
				},
			),
		},
		// verify updates to name alone
		{
			Config: config + compartmentIdVariableStr + ObjectStorageObjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Optional, acctest.Update,
					acctest.GetUpdatedRepresentationCopy("object", acctest.Representation{RepType: acctest.Required, Create: `my-test-object-1`, Update: `my-test-object-3`}, ObjectStorageObjectRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cache_control", "no-store"),
				resource.TestCheckResourceAttr(resourceName, "content_disposition", "attachment; filename=\"filename.html\""),
				resource.TestCheckResourceAttr(resourceName, "content_encoding", "identity"),
				resource.TestCheckResourceAttr(resourceName, "content_language", "en-CA"),
				resource.TestCheckResourceAttr(resourceName, "content_length", "16"),
				resource.TestCheckResourceAttr(resourceName, "content_md5", *md5B64Encode2),
				resource.TestCheckResourceAttr(resourceName, "content_type", "text/xml"),
				resource.TestCheckResourceAttr(resourceName, "bucket", testBucketName),
				resource.TestCheckResourceAttrSet(resourceName, "content"),
				resource.TestCheckResourceAttr(resourceName, "content", md5sum2),
				resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "metadata.content-type", "text/xml"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-3"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					// @CODEGEN 06/2018: Name is part of the id, and hence id will be updated
					if resId == resId2 {
						return fmt.Errorf("Resource updated when it was supposed to be recreated.")
					}
					return err
				},
			),
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + ObjectStorageObjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Optional, acctest.Update,
					acctest.GetUpdatedRepresentationCopy("object", acctest.Representation{RepType: acctest.Required, Create: `my-test-object-1`, Update: `my-test-object-3`}, ObjectStorageObjectRepresentation)) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Required, acctest.Create, objectSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "base64_encode_content", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cache_control", "no-store"),
				resource.TestCheckResourceAttr(singularDatasourceName, "content_disposition", "attachment; filename=\"filename.html\""),
				resource.TestCheckResourceAttr(singularDatasourceName, "content_encoding", "identity"),
				resource.TestCheckResourceAttr(singularDatasourceName, "content_language", "en-CA"),
				resource.TestCheckResourceAttr(singularDatasourceName, "content_length", "16"),
				resource.TestCheckResourceAttr(singularDatasourceName, "content_md5", *md5B64Encode2),
				resource.TestCheckResourceAttr(singularDatasourceName, "content_type", "text/xml"),
				resource.TestCheckResourceAttr(singularDatasourceName, "bucket", testBucketName),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "content"),
				resource.TestCheckResourceAttr(singularDatasourceName, "content", "<a1>content</a1>"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.content-type", "text/xml"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
				resource.TestCheckResourceAttr(singularDatasourceName, "object", "my-test-object-3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "version_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "storage_tier", "InfrequentAccess"),
			),
		},
		// verify base64 encoding in singular datasource
		{
			Config: config + compartmentIdVariableStr + ObjectStorageObjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Optional, acctest.Update,
					acctest.GetUpdatedRepresentationCopy("object", acctest.Representation{RepType: acctest.Required, Create: `my-test-object-1`, Update: `my-test-object-3`}, ObjectStorageObjectRepresentation)) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Optional, acctest.Update, objectSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "base64_encode_content", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cache_control", "no-store"),
				resource.TestCheckResourceAttr(singularDatasourceName, "content_disposition", "inline"),
				resource.TestCheckResourceAttr(singularDatasourceName, "content_encoding", "identity"),
				resource.TestCheckResourceAttr(singularDatasourceName, "content_language", "en-US"),
				resource.TestCheckResourceAttr(singularDatasourceName, "content_length", "16"),
				resource.TestCheckResourceAttr(singularDatasourceName, "content_md5", *md5B64Encode2),
				resource.TestCheckResourceAttr(singularDatasourceName, "content_type", "text/plain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "http_response_expires", expirationTimeForPar.Format(time.RFC3339Nano)),
				resource.TestCheckResourceAttr(singularDatasourceName, "bucket", testBucketName),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "content"),
				resource.TestCheckResourceAttr(singularDatasourceName, "content", base64.StdEncoding.EncodeToString([]byte("<a1>content</a1>"))),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.content-type", "text/xml"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
				resource.TestCheckResourceAttr(singularDatasourceName, "object", "my-test-object-3"),
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_objects", "test_objects", acctest.Required, acctest.Update, ObjectStorageObjectStorageObjectDataSourceRepresentation) +
				compartmentIdVariableStr + ObjectStorageObjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Optional, acctest.Update, ObjectStorageObjectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "bucket", testBucketName),
				resource.TestCheckResourceAttrSet(datasourceName, "namespace"),

				resource.TestCheckResourceAttr(datasourceName, "objects.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "objects.0.etag"),
				resource.TestCheckResourceAttr(datasourceName, "objects.0.storage_tier", "InfrequentAccess"),
			),
		},
		// verify datasource for delimiter and prefix
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_objects", "test_objects", acctest.Optional, acctest.Update, ObjectStorageObjectStorageObjectDataSourceRepresentation) +
				compartmentIdVariableStr + ObjectStorageObjectResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Optional, acctest.Update, ObjectStorageObjectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "bucket", testBucketName),
				resource.TestCheckResourceAttrSet(datasourceName, "namespace"),
				resource.TestCheckResourceAttr(datasourceName, "objects.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "delimiter", "/"),
				resource.TestCheckResourceAttr(datasourceName, "end", "x"),
				resource.TestCheckResourceAttr(datasourceName, "prefix", "my-test"),
				resource.TestCheckResourceAttrSet(datasourceName, "start"),
				resource.TestCheckResourceAttr(datasourceName, "start_after", "a"),
			),
		},
		// verify resource import
		{
			Config:            config + ObjectStorageObjectRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				// TODO: Terraform exhibits abnormal behavior when importing fields that need to be converted via StateFunc
				// before storing in state.
				//
				// In this case, we were able to retrieve the content and set it using ResourceData.Set. But when converting
				// ResourceData to a state, Terraform strips it (possibly because ResourceData.Set stores it as a byte
				// array, while the schema expects a string?) Ignore this check as part of import tests for now.
				"content",
				"state",
				"work_request_id",
				"delete_all_object_versions",
				"metadata",
				"storage_tier",
				"opc_sse_kms_key_id",
			},
			ResourceName: resourceName,
		},
	})
}

// issue-routing-tag: object_storage/default
func TestObjectStorageObjectResource_failContentLengthLimit(t *testing.T) {
	httpreplay.SetScenario("TestObjectStorageObjectResource_failContentLengthLimit")
	defer httpreplay.SaveScenario()
	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	var resourceName = "oci_objectstorage_object.test_object"
	var failObjectName, failBucketName, failNamespaceName string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			{
				Config: config + compartmentIdVariableStr + ObjectStorageObjectResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Optional, acctest.Update,
						acctest.GetUpdatedRepresentationCopy("object", acctest.Representation{RepType: acctest.Required, Create: `my-test-object-1`, Update: `my-test-object-3`}, ObjectStorageObjectRepresentation)),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					func(s *terraform.State) (err error) {
						failObjectName, err = acctest.FromInstanceState(s, resourceName, "object")
						if err != nil {
							return err
						}
						failBucketName, err = acctest.FromInstanceState(s, resourceName, "bucket")
						if err != nil {
							return err
						}
						failNamespaceName, err = acctest.FromInstanceState(s, resourceName, "namespace")
						return err
					}),
			},
			{
				Config: config + compartmentIdVariableStr + ObjectStorageObjectResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Optional, acctest.Update,
						acctest.GetUpdatedRepresentationCopy("object", acctest.Representation{RepType: acctest.Required, Create: `my-test-object-1`, Update: `my-test-object-3`}, ObjectStorageObjectRepresentation)) +
					acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Optional, acctest.Update,
						acctest.GetUpdatedRepresentationCopy("content_length_limit", acctest.Representation{RepType: acctest.Optional, Create: `17`, Update: `15`}, objectSingularDataSourceRepresentation)),
				ExpectError: regexp.MustCompile("the requested object's content length is 16 the limit is set to 15"),
			},
		},
	})

	//destroy test will be skipped since there is no state after the error in Get
	if failObjectName != "" && failBucketName != "" && failNamespaceName != "" {
		client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ObjectStorageClient()
		_, objectErr := client.DeleteObject(context.Background(), oci_object_storage.DeleteObjectRequest{
			NamespaceName: &failNamespaceName,
			BucketName:    &failBucketName,
			ObjectName:    &failObjectName,
		})

		_, bucketErr := client.DeleteBucket(context.Background(), oci_object_storage.DeleteBucketRequest{
			NamespaceName: &failNamespaceName,
			BucketName:    &failBucketName,
		})

		if objectErr != nil || bucketErr != nil {
			t.Errorf("failed to delete resources for the test: %v, %v", objectErr, bucketErr)
		}
	}
}

// This test is separated from the above test due to weird behavior from Terraform test framework.
// An test step that results in an error will result in the state being voided. Isolate such test steps to
// avoid interfering with regular tests that Create/Update resources.
// issue-routing-tag: object_storage/default
func TestObjectStorageObjectResource_metadata(t *testing.T) {
	httpreplay.SetScenario("TestObjectStorageObjectResource_metadata")
	defer httpreplay.SaveScenario()
	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckObjectStorageObjectDestroy,
		Steps: []resource.TestStep{
			// verify validations on metadata key
			{
				Config: config + compartmentIdVariableStr + ObjectStorageObjectResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Optional, acctest.Update,
						acctest.GetUpdatedRepresentationCopy("metadata", acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"content-type": "text/plain"}, Update: map[string]string{"CONTENT-TYPE": "text/xml"}}, ObjectStorageObjectRepresentation)),
				ExpectError: regexp.MustCompile("All 'metadata' keys must be lowercase"),
			},
		},
	})
}

func testAccCheckObjectStorageObjectDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ObjectStorageClient()

	if singlePartFile != nil {
		if _, err := os.Stat(singlePartFile.Name()); err == nil {
			os.Remove(singlePartFile.Name())
		}
	}
	if multiPartFile != nil {
		if _, err := os.Stat(multiPartFile.Name()); err == nil {
			os.Remove(multiPartFile.Name())
		}
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_objectstorage_object" {
			noResourceFound = false
			request := oci_object_storage.HeadObjectRequest{}

			if value, ok := rs.Primary.Attributes["bucket"]; ok {
				request.BucketName = &value
			}

			if value, ok := rs.Primary.Attributes["namespace"]; ok {
				request.NamespaceName = &value
			}

			if value, ok := rs.Primary.Attributes["object"]; ok {
				request.ObjectName = &value
			}

			if value, ok := rs.Primary.Attributes["version_id"]; ok {
				request.VersionId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "object_storage")

			_, err := client.HeadObject(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("ObjectStorageObject") {
		resource.AddTestSweepers("ObjectStorageObject", &resource.Sweeper{
			Name:         "ObjectStorageObject",
			Dependencies: acctest.DependencyGraph["object"],
			F:            sweepObjectStorageObjectResource,
		})
	}
}

func sweepObjectStorageObjectResource(compartment string) error {
	objectStorageClient := acctest.GetTestClients(&schema.ResourceData{}).ObjectStorageClient()
	objectIds, err := getObjectStorageObjectIds(compartment)
	if err != nil {
		return err
	}
	for _, objectId := range objectIds {
		if ok := acctest.SweeperDefaultResourceId[objectId]; !ok {
			deleteObjectRequest := oci_object_storage.DeleteObjectRequest{}

			deleteObjectRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "object_storage")
			_, error := objectStorageClient.DeleteObject(context.Background(), deleteObjectRequest)
			if error != nil {
				fmt.Printf("Error deleting Object %s %s, It is possible that the resource is already deleted. Please verify manually \n", objectId, error)
				continue
			}
		}
	}
	return nil
}

func getObjectStorageObjectIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ObjectId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	objectStorageClient := acctest.GetTestClients(&schema.ResourceData{}).ObjectStorageClient()

	listObjectsRequest := oci_object_storage.ListObjectsRequest{}

	buckets, error := getObjectStorageBucketIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting bucket required for Object resource requests \n")
	}
	for _, bucket := range buckets {
		listObjectsRequest.BucketName = &bucket

		namespaces, error := getNamespaces(compartment)
		if error != nil {
			return resourceIds, fmt.Errorf("Error getting namespace required for Object resource requests \n")
		}
		for _, namespace := range namespaces {
			listObjectsRequest.NamespaceName = &namespace

			listObjectsResponse, err := objectStorageClient.ListObjects(context.Background(), listObjectsRequest)

			if err != nil {
				return resourceIds, fmt.Errorf("Error getting Object list for compartment id : %s , %s \n", compartmentId, err)
			}
			for _, object := range listObjectsResponse.Objects {
				id := *object.Name
				resourceIds = append(resourceIds, id)
				acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ObjectId", id)
			}

		}
	}
	return resourceIds, nil
}

var (
	objectSourceRepresentation = map[string]interface{}{
		"bucket":              acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"namespace":           acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.namespace}`},
		"object":              acctest.Representation{RepType: acctest.Required, Create: `my-test-object-1`, Update: `my-test-object-3`},
		"source":              acctest.Representation{RepType: acctest.Optional, Create: ``},
		"cache_control":       acctest.Representation{RepType: acctest.Optional, Create: `no-cache`},
		"content_disposition": acctest.Representation{RepType: acctest.Optional, Create: `inline`},
		"content_encoding":    acctest.Representation{RepType: acctest.Optional, Create: `identity`},
		"content_language":    acctest.Representation{RepType: acctest.Optional, Create: `en-US`, Update: `en-CA`},
		"content_type":        acctest.Representation{RepType: acctest.Optional, Create: `text/plain`, Update: `text/xml`},
		"storage_tier":        acctest.Representation{RepType: acctest.Optional, Create: `InfrequentAccess`},
		"metadata":            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"content-type": "text/plain"}, Update: map[string]string{"content-type": "text/xml"}},
	}
)

const (
	//the object size is less than default part value, single part upload
	singlePartFilePrefix = "small-"
	singlePartFileSize   = 42e6
	opcSingleMd5         = "iMBtc3kGpfXuMgOX9sVm0Q=="

	//the object will be split on 3 parts
	multiPartFilePrefix = "large-"
	multiPartFileSize   = 300e6
	opcMultipartMd5     = "leCtKnqvcLbdMUeTbjnKnA==-3"
)

var (
	singlePartFile *os.File
	multiPartFile  *os.File
)

func createTmpFiles() (string, string, error) {
	singlePartFile, err := ioutil.TempFile(os.TempDir(), singlePartFilePrefix)
	if err != nil {
		return "", "", err
	}
	if err := singlePartFile.Truncate(singlePartFileSize); err != nil {
		return "", "", err
	}

	multiPartFile, err = ioutil.TempFile(os.TempDir(), multiPartFilePrefix)
	if err != nil {
		return "", "", err
	}
	if err := multiPartFile.Truncate(multiPartFileSize); err != nil {
		return "", "", err
	}

	return singlePartFile.Name(), multiPartFile.Name(), nil
}

// issue-routing-tag: object_storage/default
func TestObjectStorageObjectResource_multipartUpload(t *testing.T) {
	httpreplay.SetScenario("TestObjectStorageObjectResource_multipartUpload")
	defer httpreplay.SaveScenario()
	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_objectstorage_object.test_object"
	datasourceName := "data.oci_objectstorage_objects.test_objects"

	var resId, resId2 string

	singlePartFilePath, multiPartFilePath, err := createTmpFiles()
	if err != nil {
		t.Fatalf("Unable to Create files to upload. Error: %q", err)
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckObjectStorageObjectDestroy,
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: config + compartmentIdVariableStr + ObjectStorageObjectResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Required, acctest.Create,
						acctest.GetUpdatedRepresentationCopy("source", acctest.Representation{RepType: acctest.Optional, Create: singlePartFilePath}, objectSourceRepresentation)),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "bucket", testBucketName),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-1"),
					resource.TestCheckResourceAttr(resourceName, "content_type", "application/octet-stream"),
					// New SDK doesn't set omitted values from response, check they are missing from state.
					resource.TestCheckNoResourceAttr(resourceName, "cache_control"),
					resource.TestCheckNoResourceAttr(resourceName, "content_disposition"),
					resource.TestCheckNoResourceAttr(resourceName, "content"),
					resource.TestCheckNoResourceAttr(resourceName, "content_language"),
					resource.TestCheckNoResourceAttr(resourceName, "content_encoding"),
					resource.TestCheckResourceAttr(resourceName, "content_length", "0"),
					resource.TestCheckResourceAttrSet(resourceName, "content_md5"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// delete before next Create
			{
				Config: config + compartmentIdVariableStr + ObjectStorageObjectResourceDependencies,
			},
			// verify Create singlepart with optionals
			{
				Config: config + compartmentIdVariableStr + ObjectStorageObjectResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Optional, acctest.Create,
						acctest.GetUpdatedRepresentationCopy("source", acctest.Representation{RepType: acctest.Optional, Create: singlePartFilePath}, objectSourceRepresentation)),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "cache_control", "no-cache"),
					resource.TestCheckResourceAttr(resourceName, "content_disposition", "inline"),
					resource.TestCheckResourceAttr(resourceName, "content_encoding", "identity"),
					resource.TestCheckResourceAttr(resourceName, "content_language", "en-US"),
					resource.TestCheckResourceAttr(resourceName, "content_length", strconv.Itoa(singlePartFileSize)),
					resource.TestCheckResourceAttr(resourceName, "content_md5", opcSingleMd5),
					resource.TestCheckResourceAttr(resourceName, "content_type", "text/plain"),
					resource.TestCheckResourceAttr(resourceName, "bucket", testBucketName),
					resource.TestCheckNoResourceAttr(resourceName, "content"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "metadata.content-type", "text/plain"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-1"),
					resource.TestCheckResourceAttr(resourceName, "storage_tier", "InfrequentAccess"),

					func(s *terraform.State) (err error) {

						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// delete before next Create
			{
				Config: config + compartmentIdVariableStr + ObjectStorageObjectResourceDependencies,
			},
			// verify Create with optionals
			{
				Config: config + compartmentIdVariableStr + ObjectStorageObjectResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Optional, acctest.Create,
						acctest.GetUpdatedRepresentationCopy("source", acctest.Representation{RepType: acctest.Optional, Create: multiPartFilePath}, objectSourceRepresentation)),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "cache_control", "no-cache"),
					resource.TestCheckResourceAttr(resourceName, "content_disposition", "inline"),
					resource.TestCheckResourceAttr(resourceName, "content_encoding", "identity"),
					resource.TestCheckResourceAttr(resourceName, "content_language", "en-US"),
					resource.TestCheckResourceAttr(resourceName, "content_length", strconv.Itoa(multiPartFileSize)),
					resource.TestCheckResourceAttr(resourceName, "content_md5", opcMultipartMd5),
					resource.TestCheckResourceAttr(resourceName, "content_type", "text/plain"),
					resource.TestCheckResourceAttr(resourceName, "bucket", testBucketName),
					resource.TestCheckNoResourceAttr(resourceName, "content"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "metadata.content-type", "text/plain"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-1"),
					resource.TestCheckResourceAttr(resourceName, "storage_tier", "InfrequentAccess"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify updates to name alone
			{
				Config: config + compartmentIdVariableStr + ObjectStorageObjectResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Optional, acctest.Update,
						acctest.GetUpdatedRepresentationCopy("source", acctest.Representation{RepType: acctest.Optional, Create: multiPartFilePath}, objectSourceRepresentation)),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "cache_control", "no-cache"),
					resource.TestCheckResourceAttr(resourceName, "content_disposition", "inline"),
					resource.TestCheckResourceAttr(resourceName, "content_encoding", "identity"),
					resource.TestCheckResourceAttr(resourceName, "content_language", "en-CA"),
					resource.TestCheckResourceAttr(resourceName, "content_length", strconv.Itoa(multiPartFileSize)),
					resource.TestCheckResourceAttr(resourceName, "content_md5", opcMultipartMd5),
					resource.TestCheckResourceAttr(resourceName, "content_type", "text/xml"),
					resource.TestCheckResourceAttr(resourceName, "bucket", testBucketName),
					resource.TestCheckNoResourceAttr(resourceName, "content"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "metadata.content-type", "text/xml"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-3"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						// @CODEGEN 06/2018: Name is part of the id, and hence id will be updated
						if resId == resId2 {
							return fmt.Errorf("Resource updated when it was supposed to be recreated.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_objects", "test_objects", acctest.Required, acctest.Update, ObjectStorageObjectStorageObjectDataSourceRepresentation) +
					compartmentIdVariableStr + ObjectStorageObjectResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Optional, acctest.Update,
						acctest.GetUpdatedRepresentationCopy("source", acctest.Representation{RepType: acctest.Optional, Create: multiPartFilePath}, objectSourceRepresentation)),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "bucket", testBucketName),
					resource.TestCheckResourceAttrSet(datasourceName, "namespace"),

					resource.TestCheckResourceAttr(datasourceName, "objects.#", "1"),
				),
			},
			// verify datasource for delimiter and prefix
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_objects", "test_objects", acctest.Optional, acctest.Update, ObjectStorageObjectStorageObjectDataSourceRepresentation) +
					compartmentIdVariableStr + ObjectStorageObjectResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Optional, acctest.Update,
						acctest.GetUpdatedRepresentationCopy("source", acctest.Representation{RepType: acctest.Optional, Create: multiPartFilePath}, objectSourceRepresentation)),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "bucket", testBucketName),
					resource.TestCheckResourceAttrSet(datasourceName, "namespace"),

					resource.TestCheckResourceAttr(datasourceName, "objects.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "delimiter", "/"),
					resource.TestCheckResourceAttr(datasourceName, "prefix", "my-test"),
				),
			},
			// verify resource import
			{
				Config:            config + ObjectStorageObjectRequiredOnlyResource,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"source",
					"state",
					"work_request_id",
					"delete_all_object_versions",
				},
				ResourceName: resourceName,
			},
		},
	})
}

var (
	ObjectResourceConfigWithoutContent = acctest.RepresentationCopyWithRemovedProperties(ObjectStorageObjectRepresentation, []string{"content", "content_md5"})

	ObjectResourceConfigWithSourceSinglePart = ObjectStorageObjectResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Optional, acctest.Create, acctest.GetUpdatedRepresentationCopy(
			"source", acctest.Representation{RepType: acctest.Optional, Create: ""}, objectSourceRepresentation))

	ObjectCopyRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object_copy", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(ObjectResourceConfigWithoutContent, map[string]interface{}{
		"source_uri_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: objectSourceUriDetailsRepresentationSourceEtag},
		"object":             acctest.Representation{RepType: acctest.Optional, Create: `my-test-object-1-copy`},
	}))

	ObjectResourceConfigWithSourceURIFromContentObject = ObjectResourceConfigWithSourceURIFromContentObjectDependency +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object_copy", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(ObjectResourceConfigWithoutContent, map[string]interface{}{
			"source_uri_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: objectSourceUriDetailsRepresentationSourceEtag},
			"object":             acctest.Representation{RepType: acctest.Optional, Create: `my-test-object-1-copy`},
		}))

	ObjectResourceConfigWithSourceURIFromContentObjectWithoutSourceEtag = acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object_copy", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(ObjectResourceConfigWithoutContent, map[string]interface{}{
		"source_uri_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: objectSourceUriDetailsRepresentation},
		"object":             acctest.Representation{RepType: acctest.Optional, Create: `my-test-object-1-copy`},
		"metadata":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"content-type": "text/plain-copy"}},
		"storage_tier":       acctest.Representation{RepType: acctest.Optional, Create: "InfrequentAccess"},
	}))

	ObjectResourceConfigWithSourceURIFromCopyOfContentObject = acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(ObjectResourceConfigWithoutContent, map[string]interface{}{
		"source_uri_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: objectSourceUriDetailsRepresentationWithCopyObject},
		"metadata":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"content-type": "text/plain-copy-copy"}},
	}))

	ObjectResourceConfigWithSourceURIWithVersionId = acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(ObjectResourceConfigWithoutContent, map[string]interface{}{
		"source_uri_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: objectSourceUriDetailsRepresentationWithVersionId},
		"metadata":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"content-type": "text/plain-copy-copy"}},
	}))

	objectSourceUriDetailsRepresentation = map[string]interface{}{
		"region":    acctest.Representation{RepType: acctest.Required, Create: `${var.region}`},
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.namespace}`},
		"bucket":    acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"object":    acctest.Representation{RepType: acctest.Required, Create: `my-test-object-1`},
	}
	objectSourceUriDetailsRepresentationSourceEtag = acctest.RepresentationCopyWithNewProperties(objectSourceUriDetailsRepresentation, map[string]interface{}{
		"source_object_if_match_etag": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_objectstorage_object_head.object_head.etag}`},
	})
	objectSourceUriDetailsRepresentationWithCopyObject = acctest.RepresentationCopyWithNewProperties(objectSourceUriDetailsRepresentation, map[string]interface{}{
		"object": acctest.Representation{RepType: acctest.Optional, Create: `my-test-object-1-copy`},
	})
	objectSourceUriDetailsRepresentationWithVersionId = acctest.RepresentationCopyWithNewProperties(objectSourceUriDetailsRepresentation, map[string]interface{}{
		"object":            acctest.Representation{RepType: acctest.Optional, Create: `my-test-object-1-copy`},
		"source_version_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_objectstorage_object.test_object_copy.version_id}`},
	})

	ObjectResourceConfigWithSourceURIFromContentObjectDependency = acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_object_head", "object_head", acctest.Required, acctest.Create, objectHeadDatasourceRepresentation)

	objectHeadDatasourceRepresentation = map[string]interface{}{
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.namespace}`},
		"bucket":    acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"object":    acctest.Representation{RepType: acctest.Required, Create: `my-test-object-1`},
	}
)

func createTmpObjectInOtherRegion() (string, error) {
	// now running tests in one region
	singlePartFile, err := ioutil.TempFile(os.TempDir(), singlePartFilePrefix)
	if err != nil {
		return "", err
	}
	if err := singlePartFile.Truncate(singlePartFileSize); err != nil {
		return "", err
	}

	return singlePartFile.Name(), nil
}

// issue-routing-tag: object_storage/default
func TestObjectStorageObjectResource_crossRegionCopy(t *testing.T) {
	httpreplay.SetScenario("TestObjectStorageObjectResource_crossRegionCopy")
	defer httpreplay.SaveScenario()
	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singlePartFilePath, err := createTmpObjectInOtherRegion()
	if err != nil {
		t.Fatalf("Unable to Create file to upload. Error: %q", err)
	}

	resourceName := "oci_objectstorage_object.test_object"
	resourceNameCopy := "oci_objectstorage_object.test_object_copy"

	hexSum := md5.Sum([]byte("content"))
	md5sum := hex.EncodeToString(hexSum[:])

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckObjectStorageObjectDestroy,
		Steps: []resource.TestStep{
			// Create from source with options to copy
			{
				Config: config + compartmentIdVariableStr + ObjectStorageObjectResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Optional, acctest.Create, acctest.GetUpdatedRepresentationCopy(
						"source", acctest.Representation{RepType: acctest.Optional, Create: singlePartFilePath}, objectSourceRepresentation)),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "cache_control", "no-cache"),
					resource.TestCheckResourceAttr(resourceName, "content_disposition", "inline"),
					resource.TestCheckResourceAttr(resourceName, "content_encoding", "identity"),
					resource.TestCheckResourceAttrSet(resourceName, "content_language"),
					resource.TestCheckResourceAttr(resourceName, "content_length", strconv.Itoa(singlePartFileSize)),
					resource.TestCheckResourceAttr(resourceName, "content_md5", opcSingleMd5),
					resource.TestCheckResourceAttr(resourceName, "content_type", "text/plain"),
					resource.TestCheckResourceAttr(resourceName, "bucket", testBucketName),
					resource.TestCheckNoResourceAttr(resourceName, "content"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "metadata.content-type", "text/plain"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-1"),
				),
			},
			// verify copy object copy of the source object
			{
				Config: config + compartmentIdVariableStr + ObjectStorageObjectResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Optional, acctest.Create, acctest.GetUpdatedRepresentationCopy(
						"source", acctest.Representation{RepType: acctest.Optional, Create: singlePartFilePath}, objectSourceRepresentation)) +
					ObjectResourceConfigWithSourceURIFromContentObjectWithoutSourceEtag,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceNameCopy, "namespace"),
					resource.TestCheckResourceAttr(resourceNameCopy, "bucket", testBucketName),
					resource.TestCheckResourceAttr(resourceNameCopy, "object", "my-test-object-1-copy"),
					//the values were not set for the object_copy, the source object are used
					resource.TestCheckResourceAttr(resourceNameCopy, "cache_control", "no-cache"),
					resource.TestCheckResourceAttr(resourceNameCopy, "content_disposition", "inline"),
					resource.TestCheckResourceAttr(resourceNameCopy, "content_length", strconv.Itoa(singlePartFileSize)),
					resource.TestCheckResourceAttr(resourceNameCopy, "content_type", "text/plain"),
					resource.TestCheckResourceAttr(resourceNameCopy, "content_encoding", "identity"),
					resource.TestCheckResourceAttrSet(resourceName, "content_language"),
					//the values were set for the object_copy
					resource.TestCheckResourceAttr(resourceNameCopy, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceNameCopy, "metadata.content-type", "text/plain-copy"),
					resource.TestCheckResourceAttr(resourceNameCopy, "storage_tier", "InfrequentAccess"),
				),
			},
			// delete before next Create
			{
				Config: config + compartmentIdVariableStr + ObjectStorageObjectResourceDependencies,
			},
			// verify Create content object with optionals
			{
				Config: config + compartmentIdVariableStr + ObjectStorageObjectResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "cache_control", "no-cache"),
					resource.TestCheckResourceAttr(resourceName, "content_disposition", "inline"),
					resource.TestCheckResourceAttr(resourceName, "content_encoding", "identity"),
					resource.TestCheckResourceAttrSet(resourceName, "content_language"),
					resource.TestCheckResourceAttr(resourceName, "content_length", "7"),
					resource.TestCheckResourceAttrSet(resourceName, "content_md5"),
					resource.TestCheckResourceAttr(resourceName, "content_type", "text/plain"),
					resource.TestCheckResourceAttr(resourceName, "bucket", testBucketName),
					resource.TestCheckResourceAttrSet(resourceName, "content"),
					resource.TestCheckResourceAttr(resourceName, "content", md5sum),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "metadata.content-type", "text/plain"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-1"),

					func(s *terraform.State) (err error) {
						_, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify copy content object in the same bucket with source etag
			{
				Config: config + compartmentIdVariableStr + ObjectStorageObjectResourceConfig + ObjectResourceConfigWithSourceURIFromContentObject,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceNameCopy, "namespace"),
					resource.TestCheckResourceAttr(resourceNameCopy, "bucket", testBucketName),
					resource.TestCheckResourceAttr(resourceNameCopy, "object", "my-test-object-1-copy"),
					//the values were not set for the object_copy, the source object are used
					resource.TestCheckResourceAttr(resourceNameCopy, "cache_control", "no-cache"),
					resource.TestCheckResourceAttr(resourceNameCopy, "content_disposition", "inline"),
					resource.TestCheckResourceAttr(resourceNameCopy, "content_length", "7"),
					resource.TestCheckResourceAttr(resourceNameCopy, "content_type", "text/plain"),
					resource.TestCheckResourceAttr(resourceNameCopy, "content_encoding", "identity"),
					resource.TestCheckResourceAttrSet(resourceName, "content_language"),
					//the values were set for the object_copy
					resource.TestCheckResourceAttr(resourceNameCopy, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceNameCopy, "metadata.content-type", "text/plain"),
				),
			},
			// verify recreate copy content object in the same bucket - remove source etag
			// metadata is updated
			{
				Config: config + compartmentIdVariableStr + ObjectStorageObjectResourceConfig + ObjectResourceConfigWithSourceURIFromContentObjectWithoutSourceEtag,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceNameCopy, "namespace"),
					resource.TestCheckResourceAttr(resourceNameCopy, "bucket", testBucketName),
					resource.TestCheckResourceAttr(resourceNameCopy, "object", "my-test-object-1-copy"),
					//the values were not set for the object_copy, the source object are used
					resource.TestCheckResourceAttr(resourceNameCopy, "cache_control", "no-cache"),
					resource.TestCheckResourceAttr(resourceNameCopy, "content_disposition", "inline"),
					resource.TestCheckResourceAttr(resourceNameCopy, "content_length", "7"),
					resource.TestCheckResourceAttr(resourceNameCopy, "content_type", "text/plain"),
					resource.TestCheckResourceAttr(resourceNameCopy, "content_encoding", "identity"),
					resource.TestCheckResourceAttrSet(resourceName, "content_language"),
					//the values were set for the object_copy
					resource.TestCheckResourceAttr(resourceNameCopy, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceNameCopy, "metadata.content-type", "text/plain-copy"),
				),
			},
			//  replace content object by the copy of the content object copy in the same bucket
			{
				Config: config + compartmentIdVariableStr + ObjectResourceConfigWithSourceURIFromContentObjectWithoutSourceEtag +
					ObjectResourceConfigWithSourceURIFromCopyOfContentObject + ObjectResourceConfigWithSourceURIFromContentObjectDependency +
					ObjectStorageObjectResourceDependencies,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "bucket", testBucketName),
					resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-1"),
					//the values were not set for the object_copy, the source object are used
					resource.TestCheckResourceAttr(resourceName, "cache_control", "no-cache"),
					resource.TestCheckResourceAttr(resourceName, "content_disposition", "inline"),
					resource.TestCheckResourceAttr(resourceName, "content_length", "7"),
					resource.TestCheckResourceAttr(resourceName, "content_type", "text/plain"),
					resource.TestCheckResourceAttr(resourceName, "content_encoding", "identity"),
					resource.TestCheckResourceAttrSet(resourceName, "content_language"),
					//the values were set for the object_copy
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "metadata.content-type", "text/plain-copy-copy"),
				),
			},

			//  restore object version using source_uri_details
			{
				Config: config + compartmentIdVariableStr + ObjectResourceConfigWithSourceURIFromContentObjectWithoutSourceEtag +
					ObjectResourceConfigWithSourceURIWithVersionId + ObjectResourceConfigWithSourceURIFromContentObjectDependency +
					ObjectStorageObjectResourceDependencies,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "bucket", testBucketName),
					resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-1"),
					//the values were not set for the object_copy, the source object are used
					resource.TestCheckResourceAttr(resourceName, "cache_control", "no-cache"),
					resource.TestCheckResourceAttr(resourceName, "content_disposition", "inline"),
					resource.TestCheckResourceAttr(resourceName, "content_length", "7"),
					resource.TestCheckResourceAttr(resourceName, "content_type", "text/plain"),
					resource.TestCheckResourceAttr(resourceName, "content_encoding", "identity"),
					resource.TestCheckResourceAttrSet(resourceName, "content_language"),
					//the values were set for the object_copy
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "metadata.content-type", "text/plain-copy-copy"),
				),
			},
			// recreate copy of copy of content object by singlepart with optionals
			{
				Config: config + ObjectStorageObjectResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Optional, acctest.Create, acctest.GetUpdatedRepresentationCopy(
						"source", acctest.Representation{RepType: acctest.Optional, Create: singlePartFilePath}, objectSourceRepresentation)) +
					ObjectResourceConfigWithSourceURIFromContentObjectWithoutSourceEtag +
					compartmentIdVariableStr,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "cache_control", "no-cache"),
					resource.TestCheckResourceAttr(resourceName, "content_disposition", "inline"),
					resource.TestCheckResourceAttr(resourceName, "content_encoding", "identity"),
					resource.TestCheckResourceAttr(resourceName, "content_length", strconv.Itoa(singlePartFileSize)),
					resource.TestCheckResourceAttr(resourceName, "content_md5", opcSingleMd5),
					resource.TestCheckResourceAttr(resourceName, "content_type", "text/plain"),
					resource.TestCheckResourceAttr(resourceName, "bucket", testBucketName),
					resource.TestCheckNoResourceAttr(resourceName, "content"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "metadata.content-type", "text/plain"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-1"),
				),
			},
			// verify import copy of the content object
			{
				Config:            config + ObjectCopyRequiredOnlyResource,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"source_uri_details",
					"state",
					"work_request_id",
					"delete_all_object_versions",
					"storage_tier",
					"opc_sse_kms_key_id",
				},
				ResourceName: resourceNameCopy,
			},
		},
	})

}
