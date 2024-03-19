// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_object_storage "github.com/oracle/oci-go-sdk/v65/objectstorage"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	ObjectStorageObjectLifecyclePolicyRequiredOnlyResource = ObjectStorageObjectLifecyclePolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", acctest.Required, acctest.Create, ObjectStorageObjectLifecyclePolicyRepresentation)

	ObjectStorageObjectLifecyclePolicyResourceConfig = ObjectStorageObjectLifecyclePolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", acctest.Optional, acctest.Update, ObjectStorageObjectLifecyclePolicyRepresentation)

	bucketName  = utils.RandomStringOrHttpReplayValue(32, utils.Charset, "bucket1")
	bucketName2 = utils.RandomStringOrHttpReplayValue(32, utils.Charset, "bucket2")
	bucketName3 = utils.RandomStringOrHttpReplayValue(32, utils.Charset, "bucket3")

	ObjectStorageObjectStorageObjectLifecyclePolicySingularDataSourceRepresentation = map[string]interface{}{
		"bucket":    acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	}

	ObjectStorageObjectLifecyclePolicyRepresentation = map[string]interface{}{
		"bucket":    acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"rules":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: ObjectStorageObjectLifecyclePolicyRulesRepresentation},
	}
	ObjectStorageObjectLifecyclePolicyRulesRepresentation = map[string]interface{}{
		"action":             acctest.Representation{RepType: acctest.Required, Create: `ARCHIVE`, Update: `DELETE`},
		"is_enabled":         acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"name":               acctest.Representation{RepType: acctest.Required, Create: `sampleRule`, Update: `name2`},
		"time_amount":        acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"time_unit":          acctest.Representation{RepType: acctest.Required, Create: `DAYS`, Update: `YEARS`},
		"object_name_filter": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ObjectStorageObjectLifecyclePolicyRulesObjectNameFilterRepresentation},
		"target":             acctest.Representation{RepType: acctest.Optional, Create: `objects`},
	}
	ObjectStorageObjectLifecyclePolicyRulesObjectNameFilterRepresentation = map[string]interface{}{
		"exclusion_patterns": acctest.Representation{RepType: acctest.Optional, Create: []string{`exclusionPattern1`, `exclusionPattern2`}, Update: []string{`exclusionPattern1`, `exclusionPattern2`, `exclusionPattern3`}},
		"inclusion_patterns": acctest.Representation{RepType: acctest.Optional, Create: []string{`inclusionPattern1`, `inclusionPattern2`}, Update: []string{`inclusionPattern1`, `inclusionPattern2`, `inclusionPattern3`}},
		"inclusion_prefixes": acctest.Representation{RepType: acctest.Optional, Create: []string{bucketName, bucketName2}, Update: []string{bucketName, bucketName2, bucketName3}},
	}

	ObjectStorageObjectLifecyclePolicyResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, acctest.GetUpdatedRepresentationCopy("name", acctest.Representation{RepType: acctest.Required, Create: bucketName}, ObjectStorageBucketRepresentation)) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Optional, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: object_storage/default
func TestObjectStorageObjectLifecyclePolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestObjectStorageObjectLifecyclePolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_objectstorage_object_lifecycle_policy.test_object_lifecycle_policy"

	singularDatasourceName := "data.oci_objectstorage_object_lifecycle_policy.test_object_lifecycle_policy"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ObjectStorageObjectLifecyclePolicyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", acctest.Optional, acctest.Create, ObjectStorageObjectLifecyclePolicyRepresentation), "objectstorage", "objectLifecyclePolicy", t)

	acctest.ResourceTest(t, testAccCheckObjectStorageObjectLifecyclePolicyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ObjectStorageObjectLifecyclePolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", acctest.Required, acctest.Create, ObjectStorageObjectLifecyclePolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bucket", bucketName),
				resource.TestCheckResourceAttr(resourceName, "rules.#", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ObjectStorageObjectLifecyclePolicyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ObjectStorageObjectLifecyclePolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", acctest.Optional, acctest.Create, ObjectStorageObjectLifecyclePolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bucket", bucketName),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "rules", map[string]string{
					"action":               "ARCHIVE",
					"is_enabled":           "false",
					"name":                 "sampleRule",
					"object_name_filter.#": "1",
					"object_name_filter.0.inclusion_prefixes.#": "2",
					"object_name_filter.0.exclusion_patterns.#": "2",
					"object_name_filter.0.inclusion_patterns.#": "2",
					"target":      "objects",
					"time_amount": "10",
					"time_unit":   "DAYS",
				},
					[]string{}),

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
			Config: config + compartmentIdVariableStr + ObjectStorageObjectLifecyclePolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", acctest.Optional, acctest.Update, ObjectStorageObjectLifecyclePolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bucket", bucketName),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "rules", map[string]string{
					"action":               "DELETE",
					"is_enabled":           "true",
					"name":                 "name2",
					"object_name_filter.#": "1",
					"object_name_filter.0.inclusion_prefixes.#": "3",
					"object_name_filter.0.exclusion_patterns.#": "3",
					"object_name_filter.0.inclusion_patterns.#": "3",
					"target":      "objects",
					"time_amount": "11",
					"time_unit":   "YEARS",
				},
					[]string{}),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", acctest.Required, acctest.Create, ObjectStorageObjectStorageObjectLifecyclePolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + ObjectStorageObjectLifecyclePolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "bucket", bucketName),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(singularDatasourceName, "rules", map[string]string{},
					[]string{}),

				resource.TestCheckResourceAttr(singularDatasourceName, "rules.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(singularDatasourceName, "rules", map[string]string{
					"action":               "DELETE",
					"is_enabled":           "true",
					"name":                 "name2",
					"object_name_filter.#": "1",
					"object_name_filter.0.inclusion_prefixes.#": "3",
					"object_name_filter.0.exclusion_patterns.#": "3",
					"object_name_filter.0.inclusion_patterns.#": "3",
					"target":      "objects",
					"time_amount": "11",
					"time_unit":   "YEARS",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + ObjectStorageObjectLifecyclePolicyRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckObjectStorageObjectLifecyclePolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ObjectStorageClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_objectstorage_object_lifecycle_policy" {
			noResourceFound = false
			request := oci_object_storage.GetObjectLifecyclePolicyRequest{}

			if value, ok := rs.Primary.Attributes["bucket"]; ok {
				request.BucketName = &value
			}

			if value, ok := rs.Primary.Attributes["namespace"]; ok {
				request.NamespaceName = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "object_storage")

			_, err := client.GetObjectLifecyclePolicy(context.Background(), request)

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
