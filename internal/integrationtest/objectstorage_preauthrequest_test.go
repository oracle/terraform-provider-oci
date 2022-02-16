// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_object_storage "github.com/oracle/oci-go-sdk/v58/objectstorage"

	tf_objectstorage "github.com/terraform-providers/terraform-provider-oci/internal/service/objectstorage"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	testPreAuthBucketName = utils.RandomStringOrHttpReplayValue(32, utils.Charset, "bucketPreAuth")

	expirationTimeForPar = time.Now().UTC().AddDate(0, 0, 1).Truncate(time.Millisecond)

	PreauthenticatedRequestRequiredOnlyResource = PreauthenticatedRequestResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_preauthrequest", "test_preauthenticated_request", acctest.Required, acctest.Create, preauthenticatedRequestRepresentation)

	PreauthenticatedRequestResourceConfig = PreauthenticatedRequestResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_preauthrequest", "test_preauthenticated_request", acctest.Optional, acctest.Update, preauthenticatedRequestRepresentation)

	preauthenticatedRequestSingularDataSourceRepresentation = map[string]interface{}{
		"bucket":    acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.namespace}`},
		"par_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_preauthrequest.test_preauthenticated_request.par_id}`},
	}

	preauthenticatedRequestDataSourceRepresentation = map[string]interface{}{
		"bucket":             acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"namespace":          acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.namespace}`},
		"object_name_prefix": acctest.Representation{RepType: acctest.Optional, Create: `my-test-object`},
		"filter":             acctest.RepresentationGroup{RepType: acctest.Required, Group: preauthenticatedRequestDataSourceFilterRepresentation}}
	preauthenticatedRequestDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_objectstorage_preauthrequest.test_preauthenticated_request.par_id}`}},
	}

	preauthenticatedRequestRepresentation = map[string]interface{}{
		"access_type":           acctest.Representation{RepType: acctest.Required, Create: `AnyObjectWrite`, Update: `ObjectRead`},
		"bucket":                acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"name":                  acctest.Representation{RepType: acctest.Required, Create: `-tf-par`},
		"namespace":             acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.namespace}`},
		"time_expires":          acctest.Representation{RepType: acctest.Required, Create: expirationTimeForPar.Format(time.RFC3339Nano)},
		"object":                acctest.Representation{RepType: acctest.Optional, Create: `my-test-object-1`},
		"bucket_listing_action": acctest.Representation{RepType: acctest.Optional, Create: ``},
	}

	PreauthenticatedRequestResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, acctest.GetUpdatedRepresentationCopy("name", acctest.Representation{RepType: acctest.Required, Create: testPreAuthBucketName}, bucketRepresentation)) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Optional, acctest.Create, namespaceSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Required, acctest.Create, objectRepresentation)
)

// issue-routing-tag: object_storage/default
func TestObjectStoragePreauthenticatedRequestResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestObjectStoragePreauthenticatedRequestResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_objectstorage_preauthrequest.test_preauthenticated_request"
	datasourceName := "data.oci_objectstorage_preauthrequests.test_preauthenticated_requests"
	singularDatasourceName := "data.oci_objectstorage_preauthrequest.test_preauthenticated_request"

	var resId string

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+PreauthenticatedRequestResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_preauthrequest", "test_preauthenticated_request", acctest.Optional, acctest.Create, preauthenticatedRequestRepresentation), "objectstorage", "preauthenticatedRequest", t)

	acctest.ResourceTest(t, testAccCheckObjectStoragePreauthenticatedRequestDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + PreauthenticatedRequestResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_preauthrequest", "test_preauthenticated_request", acctest.Required, acctest.Create, preauthenticatedRequestRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "access_type", "AnyObjectWrite"),
				resource.TestCheckResourceAttr(resourceName, "bucket", testPreAuthBucketName),
				resource.TestCheckResourceAttr(resourceName, "name", "-tf-par"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "time_expires", expirationTimeForPar.Format(time.RFC3339Nano)),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + PreauthenticatedRequestResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + PreauthenticatedRequestResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_preauthrequest", "test_preauthenticated_request", acctest.Optional, acctest.Update, preauthenticatedRequestRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "access_type", "ObjectRead"),
				resource.TestCheckResourceAttr(resourceName, "bucket_listing_action", ""),
				resource.TestCheckResourceAttrSet(resourceName, "access_uri"),
				resource.TestCheckResourceAttr(resourceName, "bucket", testPreAuthBucketName),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "-tf-par"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "time_expires", expirationTimeForPar.Format(time.RFC3339Nano)),

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

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_preauthrequests", "test_preauthenticated_requests", acctest.Optional, acctest.Update, preauthenticatedRequestDataSourceRepresentation) +
				compartmentIdVariableStr + PreauthenticatedRequestResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_preauthrequest", "test_preauthenticated_request", acctest.Optional, acctest.Update, preauthenticatedRequestRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "bucket", testPreAuthBucketName),
				resource.TestCheckResourceAttrSet(datasourceName, "namespace"),
				resource.TestCheckResourceAttr(datasourceName, "object_name_prefix", "my-test-object"),

				resource.TestCheckResourceAttr(datasourceName, "preauthenticated_requests.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "preauthenticated_requests.0.access_type", "ObjectRead"),
				resource.TestCheckResourceAttr(datasourceName, "preauthenticated_requests.0.bucket_listing_action", ""),
				resource.TestCheckResourceAttrSet(datasourceName, "preauthenticated_requests.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "preauthenticated_requests.0.name", "-tf-par"),
				resource.TestCheckResourceAttr(datasourceName, "preauthenticated_requests.0.object", "my-test-object-1"),
				resource.TestCheckResourceAttrSet(datasourceName, "preauthenticated_requests.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "preauthenticated_requests.0.time_expires", expirationTimeForPar.String()),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_preauthrequest", "test_preauthenticated_request", acctest.Required, acctest.Create, preauthenticatedRequestSingularDataSourceRepresentation) +
				compartmentIdVariableStr + PreauthenticatedRequestResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_preauthrequest", "test_preauthenticated_request", acctest.Optional, acctest.Update, preauthenticatedRequestRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "bucket", testPreAuthBucketName),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "par_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "access_type", "ObjectRead"),
				resource.TestCheckResourceAttr(singularDatasourceName, "bucket_listing_action", ""),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "-tf-par"),
				resource.TestCheckResourceAttr(singularDatasourceName, "object", "my-test-object-1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "time_expires", expirationTimeForPar.String()),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + PreauthenticatedRequestResourceConfig,
		},
		//verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"access_uri",
				"time_expires",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckObjectStoragePreauthenticatedRequestDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ObjectStorageClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_objectstorage_preauthrequest" {
			noResourceFound = false
			request := oci_object_storage.GetPreauthenticatedRequestRequest{}

			if value, ok := rs.Primary.Attributes["bucket"]; ok {
				request.BucketName = &value
			}

			if value, ok := rs.Primary.Attributes["namespace"]; ok {
				request.NamespaceName = &value
			}

			bucket, namespace, parId, er := tf_objectstorage.ParsePreauthenticatedRequestCompositeId(rs.Primary.ID)
			if er == nil {
				request.BucketName = &bucket
				request.NamespaceName = &namespace
				request.ParId = &parId
			} else {
				log.Printf("[WARN] Get() unable to parse current ID: %s", rs.Primary.ID)
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "object_storage")

			_, err := client.GetPreauthenticatedRequest(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("ObjectStoragePreauthenticatedRequest") {
		resource.AddTestSweepers("ObjectStoragePreauthenticatedRequest", &resource.Sweeper{
			Name:         "ObjectStoragePreauthenticatedRequest",
			Dependencies: acctest.DependencyGraph["preauthenticatedRequest"],
			F:            sweepObjectStoragePreauthenticatedRequestResource,
		})
	}
}

func sweepObjectStoragePreauthenticatedRequestResource(compartment string) error {
	objectStorageClient := acctest.GetTestClients(&schema.ResourceData{}).ObjectStorageClient()
	preauthenticatedRequestIds, err := getPreauthenticatedRequestIds(compartment)
	if err != nil {
		return err
	}
	for _, preauthenticatedRequestId := range preauthenticatedRequestIds {
		if ok := acctest.SweeperDefaultResourceId[preauthenticatedRequestId]; !ok {
			deletePreauthenticatedRequestRequest := oci_object_storage.DeletePreauthenticatedRequestRequest{}

			deletePreauthenticatedRequestRequest.ParId = &preauthenticatedRequestId

			deletePreauthenticatedRequestRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "object_storage")
			_, error := objectStorageClient.DeletePreauthenticatedRequest(context.Background(), deletePreauthenticatedRequestRequest)
			if error != nil {
				fmt.Printf("Error deleting PreauthenticatedRequest %s %s, It is possible that the resource is already deleted. Please verify manually \n", preauthenticatedRequestId, error)
				continue
			}
		}
	}
	return nil
}

func getPreauthenticatedRequestIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "PreauthenticatedRequestId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	objectStorageClient := acctest.GetTestClients(&schema.ResourceData{}).ObjectStorageClient()

	listPreauthenticatedRequestsRequest := oci_object_storage.ListPreauthenticatedRequestsRequest{}

	buckets, error := getBucketIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting bucket required for PreauthenticatedRequest resource requests \n")
	}
	for _, bucket := range buckets {
		listPreauthenticatedRequestsRequest.BucketName = &bucket

		namespaces, error := getNamespaces(compartment)
		if error != nil {
			return resourceIds, fmt.Errorf("Error getting namespace required for PreauthenticatedRequest resource requests \n")
		}
		for _, namespace := range namespaces {
			listPreauthenticatedRequestsRequest.NamespaceName = &namespace

			listPreauthenticatedRequestsResponse, err := objectStorageClient.ListPreauthenticatedRequests(context.Background(), listPreauthenticatedRequestsRequest)

			if err != nil {
				return resourceIds, fmt.Errorf("Error getting PreauthenticatedRequest list for compartment id : %s , %s \n", compartmentId, err)
			}
			for _, preauthenticatedRequest := range listPreauthenticatedRequestsResponse.Items {
				id := *preauthenticatedRequest.Id
				resourceIds = append(resourceIds, id)
				acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "PreauthenticatedRequestId", id)
			}

		}
	}
	return resourceIds, nil
}
