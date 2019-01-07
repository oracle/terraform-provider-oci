// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"regexp"
	"testing"

	"os"
	"strconv"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_object_storage "github.com/oracle/oci-go-sdk/objectstorage"
)

var (
	ObjectRequiredOnlyResource = ObjectResourceDependencies +
		generateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", Required, Create, objectRepresentation)

	ObjectResourceConfig = ObjectResourceDependencies +
		generateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", Optional, Create, objectRepresentation)

	objectDataSourceRepresentation = map[string]interface{}{
		"bucket":    Representation{repType: Required, create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"namespace": Representation{repType: Required, create: `${oci_objectstorage_bucket.test_bucket.namespace}`},
		"delimiter": Representation{repType: Optional, create: `delimiter`, update: `/`},
		"prefix":    Representation{repType: Optional, create: `prefix`, update: `my-test`},
		"filter":    RepresentationGroup{Required, objectDataSourceFilterRepresentation}}
	objectDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `name`},
		"values": Representation{repType: Required, create: []string{`${oci_objectstorage_object.test_object.object}`}},
	}

	objectRepresentation = map[string]interface{}{
		"bucket":           Representation{repType: Required, create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"content":          Representation{repType: Optional, create: `content`, update: `<a1>content</a1>`},
		"namespace":        Representation{repType: Required, create: `${oci_objectstorage_bucket.test_bucket.namespace}`},
		"object":           Representation{repType: Required, create: `my-test-object-1`, update: `my-test-object-2`},
		"content_encoding": Representation{repType: Optional, create: `identity`},
		"content_language": Representation{repType: Optional, create: `en-US`, update: `en-CA`},
		"content_type":     Representation{repType: Optional, create: `text/plain`, update: `text/xml`},
		"metadata":         Representation{repType: Optional, create: map[string]string{"content-type": "text/plain"}, update: map[string]string{"content-type": "text/xml"}},
	}

	ObjectResourceDependencies = BucketRequiredOnlyResource
)

func TestObjectStorageObjectResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_objectstorage_object.test_object"
	datasourceName := "data.oci_objectstorage_objects.test_objects"

	var resId, resId2 string
	hexSum := md5.Sum([]byte("content"))
	md5sum := hex.EncodeToString(hexSum[:])
	hexSum2 := md5.Sum([]byte("<a1>content</a1>"))
	md5sum2 := hex.EncodeToString(hexSum2[:])

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckObjectStorageObjectDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + ObjectResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", Required, Create, objectRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "bucket", "my-test-1"),
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
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ObjectResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ObjectResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", Optional, Create, objectRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "content_encoding", "identity"),
					resource.TestCheckResourceAttr(resourceName, "content_language", "en-US"),
					resource.TestCheckResourceAttr(resourceName, "content_length", "7"),
					resource.TestCheckResourceAttrSet(resourceName, "content_md5"),
					resource.TestCheckResourceAttr(resourceName, "content_type", "text/plain"),
					resource.TestCheckResourceAttr(resourceName, "bucket", "my-test-1"),
					resource.TestCheckResourceAttrSet(resourceName, "content"),
					resource.TestCheckResourceAttr(resourceName, "content", md5sum),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "metadata.content-type", "text/plain"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-1"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + ObjectResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", Optional, Update, objectRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "content_encoding", "identity"),
					resource.TestCheckResourceAttr(resourceName, "content_language", "en-CA"),
					resource.TestCheckResourceAttr(resourceName, "content_length", "16"),
					resource.TestCheckResourceAttrSet(resourceName, "content_md5"),
					resource.TestCheckResourceAttr(resourceName, "content_type", "text/xml"),
					resource.TestCheckResourceAttr(resourceName, "bucket", "my-test-1"),
					resource.TestCheckResourceAttrSet(resourceName, "content"),
					resource.TestCheckResourceAttr(resourceName, "content", md5sum2),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "metadata.content-type", "text/xml"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-2"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
				Config: config + compartmentIdVariableStr + ObjectResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", Optional, Update,
						getUpdatedRepresentationCopy("object", Representation{repType: Required, create: `my-test-object-1`, update: `my-test-object-3`}, objectRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "content_encoding", "identity"),
					resource.TestCheckResourceAttr(resourceName, "content_language", "en-CA"),
					resource.TestCheckResourceAttr(resourceName, "content_length", "16"),
					resource.TestCheckResourceAttrSet(resourceName, "content_md5"),
					resource.TestCheckResourceAttr(resourceName, "content_type", "text/xml"),
					resource.TestCheckResourceAttr(resourceName, "bucket", "my-test-1"),
					resource.TestCheckResourceAttrSet(resourceName, "content"),
					resource.TestCheckResourceAttr(resourceName, "content", md5sum2),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "metadata.content-type", "text/xml"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-3"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_objectstorage_objects", "test_objects", Required, Update, objectDataSourceRepresentation) +
					compartmentIdVariableStr + ObjectResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", Optional, Update, objectRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "bucket", "my-test-1"),
					resource.TestCheckResourceAttrSet(datasourceName, "namespace"),

					resource.TestCheckResourceAttr(datasourceName, "objects.#", "1"),
				),
			},
			// verify datasource for delimiter and prefix
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_objectstorage_objects", "test_objects", Optional, Update, objectDataSourceRepresentation) +
					compartmentIdVariableStr + ObjectResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", Optional, Update, objectRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "bucket", "my-test-1"),
					resource.TestCheckResourceAttrSet(datasourceName, "namespace"),

					resource.TestCheckResourceAttr(datasourceName, "objects.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "delimiter", "/"),
					resource.TestCheckResourceAttr(datasourceName, "prefix", "my-test"),
				),
			},
			// verify resource import
			{
				Config:            config,
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
				},
				ResourceName: resourceName,
			},
		},
	})
}

// This test is separated from the above test due to weird behavior from Terraform test framework.
// An test step that results in an error will result in the state being voided. Isolate such test steps to
// avoid interfering with regular tests that Create/Update resources.
func TestObjectStorageObjectResource_metadata(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckObjectStorageObjectDestroy,
		Steps: []resource.TestStep{
			// verify validations on metadata key
			{
				Config: config + compartmentIdVariableStr + ObjectResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", Optional, Update,
						getUpdatedRepresentationCopy("metadata", Representation{repType: Optional, create: map[string]string{"content-type": "text/plain"}, update: map[string]string{"CONTENT-TYPE": "text/xml"}}, objectRepresentation)),
				ExpectError: regexp.MustCompile("All 'metadata' keys must be lowercase"),
			},
		},
	})
}

func testAccCheckObjectStorageObjectDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).objectStorageClient

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

var (
	objectSourceRepresentation = map[string]interface{}{
		"bucket":           Representation{repType: Required, create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"namespace":        Representation{repType: Required, create: `${oci_objectstorage_bucket.test_bucket.namespace}`},
		"object":           Representation{repType: Required, create: `my-test-object-1`, update: `my-test-object-3`},
		"source":           Representation{repType: Optional, create: ``},
		"content_encoding": Representation{repType: Optional, create: `identity`},
		"content_language": Representation{repType: Optional, create: `en-US`, update: `en-CA`},
		"content_type":     Representation{repType: Optional, create: `text/plain`, update: `text/xml`},
		"metadata":         Representation{repType: Optional, create: map[string]string{"content-type": "text/plain"}, update: map[string]string{"content-type": "text/xml"}},
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

func TestObjectStorageObjectResource_multipartUpload(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_objectstorage_object.test_object"
	datasourceName := "data.oci_objectstorage_objects.test_objects"

	var resId, resId2 string

	singlePartFilePath, multiPartFilePath, err := createTmpFiles()
	if err != nil {
		t.Fatalf("Unable to create files to upload. Error: %q", err)
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckObjectStorageObjectDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + ObjectResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", Required, Create,
						getUpdatedRepresentationCopy("source", Representation{repType: Optional, create: singlePartFilePath}, objectSourceRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "bucket", "my-test-1"),
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
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ObjectResourceDependencies,
			},
			// verify create singlepart with optionals
			{
				Config: config + compartmentIdVariableStr + ObjectResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", Optional, Create,
						getUpdatedRepresentationCopy("source", Representation{repType: Optional, create: singlePartFilePath}, objectSourceRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "content_encoding", "identity"),
					resource.TestCheckResourceAttr(resourceName, "content_language", "en-US"),
					resource.TestCheckResourceAttr(resourceName, "content_length", strconv.Itoa(singlePartFileSize)),
					resource.TestCheckResourceAttr(resourceName, "content_md5", opcSingleMd5),
					resource.TestCheckResourceAttr(resourceName, "content_type", "text/plain"),
					resource.TestCheckResourceAttr(resourceName, "bucket", "my-test-1"),
					resource.TestCheckNoResourceAttr(resourceName, "content"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "metadata.content-type", "text/plain"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-1"),

					func(s *terraform.State) (err error) {

						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ObjectResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ObjectResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", Optional, Create,
						getUpdatedRepresentationCopy("source", Representation{repType: Optional, create: multiPartFilePath}, objectSourceRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "content_encoding", "identity"),
					resource.TestCheckResourceAttr(resourceName, "content_language", "en-US"),
					resource.TestCheckResourceAttr(resourceName, "content_length", strconv.Itoa(multiPartFileSize)),
					resource.TestCheckResourceAttr(resourceName, "content_md5", opcMultipartMd5),
					resource.TestCheckResourceAttr(resourceName, "content_type", "text/plain"),
					resource.TestCheckResourceAttr(resourceName, "bucket", "my-test-1"),
					resource.TestCheckNoResourceAttr(resourceName, "content"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "metadata.content-type", "text/plain"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-1"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify updates to name alone
			{
				Config: config + compartmentIdVariableStr + ObjectResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", Optional, Update,
						getUpdatedRepresentationCopy("source", Representation{repType: Optional, create: multiPartFilePath}, objectSourceRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "content_encoding", "identity"),
					resource.TestCheckResourceAttr(resourceName, "content_language", "en-CA"),
					resource.TestCheckResourceAttr(resourceName, "content_length", strconv.Itoa(multiPartFileSize)),
					resource.TestCheckResourceAttr(resourceName, "content_md5", opcMultipartMd5),
					resource.TestCheckResourceAttr(resourceName, "content_type", "text/xml"),
					resource.TestCheckResourceAttr(resourceName, "bucket", "my-test-1"),
					resource.TestCheckNoResourceAttr(resourceName, "content"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "metadata.content-type", "text/xml"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-3"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_objectstorage_objects", "test_objects", Required, Update, objectDataSourceRepresentation) +
					compartmentIdVariableStr + ObjectResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", Optional, Update,
						getUpdatedRepresentationCopy("source", Representation{repType: Optional, create: multiPartFilePath}, objectSourceRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "bucket", "my-test-1"),
					resource.TestCheckResourceAttrSet(datasourceName, "namespace"),

					resource.TestCheckResourceAttr(datasourceName, "objects.#", "1"),
				),
			},
			// verify datasource for delimiter and prefix
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_objectstorage_objects", "test_objects", Optional, Update, objectDataSourceRepresentation) +
					compartmentIdVariableStr + ObjectResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", Optional, Update,
						getUpdatedRepresentationCopy("source", Representation{repType: Optional, create: multiPartFilePath}, objectSourceRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "bucket", "my-test-1"),
					resource.TestCheckResourceAttrSet(datasourceName, "namespace"),

					resource.TestCheckResourceAttr(datasourceName, "objects.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "delimiter", "/"),
					resource.TestCheckResourceAttr(datasourceName, "prefix", "my-test"),
				),
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"source",
					"state",
					"work_request_id",
				},
				ResourceName: resourceName,
			},
		},
	})
}

var (
	ObjectResourceConfigWithoutContent = representationCopyWithRemovedProperties(objectRepresentation, []string{"content"})

	ObjectResourceConfigWithSourceSinglePart = ObjectResourceDependencies +
		generateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", Optional, Create, getUpdatedRepresentationCopy(
			"source", Representation{repType: Optional, create: ""}, objectSourceRepresentation))

	ObjectResourceConfigWithSourceURIFromContentObject = ObjectResourceConfigWithSourceURIFromContentObjectDependency +
		generateResourceFromRepresentationMap("oci_objectstorage_object", "test_object_copy", Optional, Create, representationCopyWithNewProperties(ObjectResourceConfigWithoutContent, map[string]interface{}{
			"source_uri_details": RepresentationGroup{Optional, objectSourceUriDetailsRepresentationSourceEtag},
			"object":             Representation{repType: Optional, create: `my-test-object-1-copy`},
		}))

	ObjectResourceConfigWithSourceURIFromContentObjectWithoutSourceEtag = generateResourceFromRepresentationMap("oci_objectstorage_object", "test_object_copy", Optional, Create, representationCopyWithNewProperties(ObjectResourceConfigWithoutContent, map[string]interface{}{
		"source_uri_details": RepresentationGroup{Optional, objectSourceUriDetailsRepresentation},
		"object":             Representation{repType: Optional, create: `my-test-object-1-copy`},
		"metadata":           Representation{repType: Optional, create: map[string]string{"content-type": "text/plain-copy"}},
	}))

	ObjectResourceConfigWithSourceURIFromCopyOfContentObject = generateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", Optional, Create, representationCopyWithNewProperties(ObjectResourceConfigWithoutContent, map[string]interface{}{
		"source_uri_details": RepresentationGroup{Optional, objectSourceUriDetailsRepresentationWithCopyObject},
		"metadata":           Representation{repType: Optional, create: map[string]string{"content-type": "text/plain-copy-copy"}},
	}))

	objectSourceUriDetailsRepresentation = map[string]interface{}{
		"region":    Representation{repType: Required, create: `${var.region}`},
		"namespace": Representation{repType: Required, create: `${oci_objectstorage_bucket.test_bucket.namespace}`},
		"bucket":    Representation{repType: Required, create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"object":    Representation{repType: Required, create: `my-test-object-1`},
	}
	objectSourceUriDetailsRepresentationSourceEtag = representationCopyWithNewProperties(objectSourceUriDetailsRepresentation, map[string]interface{}{
		"source_object_if_match_etag": Representation{repType: Optional, create: `${data.oci_objectstorage_object_head.object_head.etag}`},
	})
	objectSourceUriDetailsRepresentationWithCopyObject = representationCopyWithNewProperties(objectSourceUriDetailsRepresentation, map[string]interface{}{
		"object": Representation{repType: Optional, create: `my-test-object-1-copy`},
	})

	ObjectResourceConfigWithSourceURIFromContentObjectDependency = generateDataSourceFromRepresentationMap("oci_objectstorage_object_head", "object_head", Required, Create, objectHeadDatasourceRepresentation)

	objectHeadDatasourceRepresentation = map[string]interface{}{
		"namespace": Representation{repType: Required, create: `${oci_objectstorage_bucket.test_bucket.namespace}`},
		"bucket":    Representation{repType: Required, create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"object":    Representation{repType: Required, create: `my-test-object-1`},
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

func TestObjectStorageObjectResource_crossRegionCopy(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singlePartFilePath, err := createTmpObjectInOtherRegion()
	if err != nil {
		t.Fatalf("Unable to create file to upload. Error: %q", err)
	}

	resourceName := "oci_objectstorage_object.test_object"
	resourceNameCopy := "oci_objectstorage_object.test_object_copy"

	hexSum := md5.Sum([]byte("content"))
	md5sum := hex.EncodeToString(hexSum[:])

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckObjectStorageObjectDestroy,
		Steps: []resource.TestStep{
			// create from source with options to copy
			{
				Config: config + compartmentIdVariableStr + ObjectResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", Optional, Create, getUpdatedRepresentationCopy(
						"source", Representation{repType: Optional, create: singlePartFilePath}, objectSourceRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "content_encoding", "identity"),
					resource.TestCheckResourceAttr(resourceName, "content_language", "en-US"),
					resource.TestCheckResourceAttr(resourceName, "content_length", strconv.Itoa(singlePartFileSize)),
					resource.TestCheckResourceAttr(resourceName, "content_md5", opcSingleMd5),
					resource.TestCheckResourceAttr(resourceName, "content_type", "text/plain"),
					resource.TestCheckResourceAttr(resourceName, "bucket", "my-test-1"),
					resource.TestCheckNoResourceAttr(resourceName, "content"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "metadata.content-type", "text/plain"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-1"),
				),
			},
			// verify copy object copy of the source object
			{
				Config: config + compartmentIdVariableStr + ObjectResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", Optional, Create, getUpdatedRepresentationCopy(
						"source", Representation{repType: Optional, create: singlePartFilePath}, objectSourceRepresentation)) +
					ObjectResourceConfigWithSourceURIFromContentObjectWithoutSourceEtag,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceNameCopy, "namespace"),
					resource.TestCheckResourceAttr(resourceNameCopy, "bucket", "my-test-1"),
					resource.TestCheckResourceAttr(resourceNameCopy, "object", "my-test-object-1-copy"),
					//the values were not set for the object_copy, the source object are used
					resource.TestCheckResourceAttr(resourceNameCopy, "content_length", strconv.Itoa(singlePartFileSize)),
					resource.TestCheckResourceAttr(resourceNameCopy, "content_type", "text/plain"),
					resource.TestCheckResourceAttr(resourceNameCopy, "content_encoding", "identity"),
					resource.TestCheckResourceAttr(resourceNameCopy, "content_language", "en-US"),
					//the values were set for the object_copy
					resource.TestCheckResourceAttr(resourceNameCopy, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceNameCopy, "metadata.content-type", "text/plain-copy"),
				),
			},
			// delete before next create
			{
				Config: config + compartmentIdVariableStr + ObjectResourceDependencies,
			},
			// verify create content object with optionals
			{
				Config: config + compartmentIdVariableStr + ObjectResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "content_encoding", "identity"),
					resource.TestCheckResourceAttr(resourceName, "content_language", "en-US"),
					resource.TestCheckResourceAttr(resourceName, "content_length", "7"),
					resource.TestCheckResourceAttrSet(resourceName, "content_md5"),
					resource.TestCheckResourceAttr(resourceName, "content_type", "text/plain"),
					resource.TestCheckResourceAttr(resourceName, "bucket", "my-test-1"),
					resource.TestCheckResourceAttrSet(resourceName, "content"),
					resource.TestCheckResourceAttr(resourceName, "content", md5sum),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "metadata.content-type", "text/plain"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-1"),

					func(s *terraform.State) (err error) {
						_, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify copy content object in the same bucket with source etag
			{
				Config: config + compartmentIdVariableStr + ObjectResourceConfig + ObjectResourceConfigWithSourceURIFromContentObject,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceNameCopy, "namespace"),
					resource.TestCheckResourceAttr(resourceNameCopy, "bucket", "my-test-1"),
					resource.TestCheckResourceAttr(resourceNameCopy, "object", "my-test-object-1-copy"),
					//the values were not set for the object_copy, the source object are used
					resource.TestCheckResourceAttr(resourceNameCopy, "content_length", "7"),
					resource.TestCheckResourceAttr(resourceNameCopy, "content_type", "text/plain"),
					resource.TestCheckResourceAttr(resourceNameCopy, "content_encoding", "identity"),
					resource.TestCheckResourceAttr(resourceNameCopy, "content_language", "en-US"),
					//the values were set for the object_copy
					resource.TestCheckResourceAttr(resourceNameCopy, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceNameCopy, "metadata.content-type", "text/plain"),
				),
			},
			// verify recreate copy content object in the same bucket - remove source etag
			// metadata is updated
			{
				Config: config + compartmentIdVariableStr + ObjectResourceConfig + ObjectResourceConfigWithSourceURIFromContentObjectWithoutSourceEtag,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceNameCopy, "namespace"),
					resource.TestCheckResourceAttr(resourceNameCopy, "bucket", "my-test-1"),
					resource.TestCheckResourceAttr(resourceNameCopy, "object", "my-test-object-1-copy"),
					//the values were not set for the object_copy, the source object are used
					resource.TestCheckResourceAttr(resourceNameCopy, "content_length", "7"),
					resource.TestCheckResourceAttr(resourceNameCopy, "content_type", "text/plain"),
					resource.TestCheckResourceAttr(resourceNameCopy, "content_encoding", "identity"),
					resource.TestCheckResourceAttr(resourceNameCopy, "content_language", "en-US"),
					//the values were set for the object_copy
					resource.TestCheckResourceAttr(resourceNameCopy, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceNameCopy, "metadata.content-type", "text/plain-copy"),
				),
			},
			//  replace content object by the copy of the content object copy in the same bucket
			{
				Config: config + compartmentIdVariableStr + ObjectResourceConfigWithSourceURIFromContentObjectWithoutSourceEtag +
					ObjectResourceConfigWithSourceURIFromCopyOfContentObject + ObjectResourceConfigWithSourceURIFromContentObjectDependency +
					ObjectResourceDependencies,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "bucket", "my-test-1"),
					resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-1"),
					//the values were not set for the object_copy, the source object are used
					resource.TestCheckResourceAttr(resourceName, "content_length", "7"),
					resource.TestCheckResourceAttr(resourceName, "content_type", "text/plain"),
					resource.TestCheckResourceAttr(resourceName, "content_encoding", "identity"),
					resource.TestCheckResourceAttr(resourceName, "content_language", "en-US"),
					//the values were set for the object_copy
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "metadata.content-type", "text/plain-copy-copy"),
				),
			},
			// recreate copy of copy of content object by singlepart with optionals
			{
				Config: config + ObjectResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", Optional, Create, getUpdatedRepresentationCopy(
						"source", Representation{repType: Optional, create: singlePartFilePath}, objectSourceRepresentation)) +
					ObjectResourceConfigWithSourceURIFromContentObjectWithoutSourceEtag +
					compartmentIdVariableStr,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "content_encoding", "identity"),
					resource.TestCheckResourceAttr(resourceName, "content_language", "en-US"),
					resource.TestCheckResourceAttr(resourceName, "content_length", strconv.Itoa(singlePartFileSize)),
					resource.TestCheckResourceAttr(resourceName, "content_md5", opcSingleMd5),
					resource.TestCheckResourceAttr(resourceName, "content_type", "text/plain"),
					resource.TestCheckResourceAttr(resourceName, "bucket", "my-test-1"),
					resource.TestCheckNoResourceAttr(resourceName, "content"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "metadata.content-type", "text/plain"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-1"),
				),
			},
			// verify import copy of the content object
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"source_uri_details",
					"state",
					"work_request_id",
				},
				ResourceName: resourceNameCopy,
			},
		},
	})

}

func initObjectStorageObjectSweeper() {
	resource.AddTestSweepers("ObjectStorageObject", &resource.Sweeper{
		Name:         "ObjectStorageObject",
		Dependencies: DependencyGraph["object"],
		F:            sweepObjectStorageObjectResource,
	})
}

func sweepObjectStorageObjectResource(compartment string) error {
	return nil
}
