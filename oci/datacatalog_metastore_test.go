// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/oracle/oci-go-sdk/v54/common"
	oci_datacatalog "github.com/oracle/oci-go-sdk/v54/datacatalog"
	"github.com/oracle/oci-go-sdk/v54/objectstorage"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	MetastoreRequiredOnlyResource = MetastoreResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_datacatalog_metastore", "test_metastore", Required, Create, metastoreRepresentation)

	MetastoreResourceConfig = MetastoreResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_datacatalog_metastore", "test_metastore", Optional, Update, metastoreRepresentation)

	metastoreSingularDataSourceRepresentation = map[string]interface{}{
		"metastore_id": Representation{RepType: Required, Create: `${oci_datacatalog_metastore.test_metastore.id}`},
	}

	metastoreDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"state":          Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, metastoreDataSourceFilterRepresentation}}
	metastoreDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_datacatalog_metastore.test_metastore.id}`}},
	}

	//Changes for retrieving the ObjectStorageURIs in the runtime
	//Representation map to Create ObjectStorage bucket
	bucketRepresentationMetastore = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"name":           Representation{RepType: Required, Create: `tf_metastore_objectstoragebucket`},
		"namespace":      Representation{RepType: Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	}

	//Representation map to Create directories inside the ObjectStorage bucket above
	objectRepresentationDefault = map[string]interface{}{
		"bucket":    Representation{RepType: Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"namespace": Representation{RepType: Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"object":    Representation{RepType: Required, Create: "Default/"},
	}
	objectRepresentationExternal = map[string]interface{}{
		"bucket":    Representation{RepType: Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"namespace": Representation{RepType: Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"object":    Representation{RepType: Required, Create: "External/"},
	}

	//Object storage values required to Create URI i.e bucket name, namespace
	objectstoragenamespace = "${data.oci_objectstorage_namespace.test_namespace.namespace}"
	objectstoragebucket    = "tf_metastore_objectstoragebucket"

	//Trying to Create the URI(oci://bucket@namespace/sub-dir/) from objectstorage objects outputs
	defaultExternalTableLocationvar = "oci://" + objectstoragebucket + "@" + objectstoragenamespace + "/" + "External"
	defaultManagedTableLocationvar  = "oci://" + objectstoragebucket + "@" + objectstoragenamespace + "/" + "Default"

	metastoreRepresentation = map[string]interface{}{
		"compartment_id":                  Representation{RepType: Required, Create: `${var.compartment_id}`},
		"default_external_table_location": Representation{RepType: Required, Create: defaultExternalTableLocationvar},
		"default_managed_table_location":  Representation{RepType: Required, Create: defaultManagedTableLocationvar},
		"defined_tags":                    Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                    Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                   Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                       RepresentationGroup{Required, ignoreMetastoreDefinedTagsChangesRepresentation},
	}
	ignoreMetastoreDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": Representation{RepType: Required, Create: []string{`defined_tags`}},
	}
	//Changes made for Create dependency resources
	MetastoreResourceDependencies = GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Required, Create, bucketRepresentationMetastore) +
		GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation) +
		GenerateResourceFromRepresentationMap("oci_objectstorage_object", "object1", Required, Create, objectRepresentationDefault) +
		GenerateResourceFromRepresentationMap("oci_objectstorage_object", "object2", Required, Create, objectRepresentationExternal) +
		DefinedTagsDependencies
)

// issue-routing-tag: datacatalog/default
func TestDatacatalogMetastoreResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatacatalogMetastoreResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	//Create the ObjectStorageURIs
	objectstoragenamespace = getObjectStorageNamespace(compartmentId)
	defaultExternalTableLocation := "oci://" + objectstoragebucket + "@" + objectstoragenamespace + "/" + "External"
	defaultManagedTableLocation := "oci://" + objectstoragebucket + "@" + objectstoragenamespace + "/" + "Default"

	resourceName := "oci_datacatalog_metastore.test_metastore"
	datasourceName := "data.oci_datacatalog_metastores.test_metastores"
	singularDatasourceName := "data.oci_datacatalog_metastore.test_metastore"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+MetastoreResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_datacatalog_metastore", "test_metastore", Optional, Create, metastoreRepresentation), "datacatalog", "metastore", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatacatalogMetastoreDestroy,
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: config + compartmentIdVariableStr + MetastoreResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_datacatalog_metastore", "test_metastore", Required, Create, metastoreRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "default_external_table_location", defaultExternalTableLocation),
					resource.TestCheckResourceAttr(resourceName, "default_managed_table_location", defaultManagedTableLocation),

					func(s *terraform.State) (err error) {
						resId, err = FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// delete before next Create
			{
				Config: config + compartmentIdVariableStr + MetastoreResourceDependencies,
			},
			// verify Create with optionals
			{
				Config: config + compartmentIdVariableStr + MetastoreResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_datacatalog_metastore", "test_metastore", Optional, Create, metastoreRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "default_external_table_location", defaultExternalTableLocation),
					resource.TestCheckResourceAttr(resourceName, "default_managed_table_location", defaultManagedTableLocation),
					//resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),

					func(s *terraform.State) (err error) {
						resId, err = FromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify Update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + MetastoreResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_datacatalog_metastore", "test_metastore", Optional, Create,
						RepresentationCopyWithNewProperties(metastoreRepresentation, map[string]interface{}{
							"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
						})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "default_external_table_location", defaultExternalTableLocation),
					resource.TestCheckResourceAttr(resourceName, "default_managed_table_location", defaultManagedTableLocation),
					//resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),

					func(s *terraform.State) (err error) {
						resId2, err = FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + MetastoreResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_datacatalog_metastore", "test_metastore", Optional, Update, metastoreRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "default_external_table_location", defaultExternalTableLocation),
					resource.TestCheckResourceAttr(resourceName, "default_managed_table_location", defaultManagedTableLocation),
					//resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),

					func(s *terraform.State) (err error) {
						resId2, err = FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config +
					GenerateDataSourceFromRepresentationMap("oci_datacatalog_metastores", "test_metastores", Optional, Update, metastoreDataSourceRepresentation) +
					compartmentIdVariableStr + MetastoreResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_datacatalog_metastore", "test_metastore", Optional, Update, metastoreRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "metastores.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "metastores.0.compartment_id", compartmentId),
					//resource.TestCheckResourceAttr(datasourceName, "metastores.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "metastores.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "metastores.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "metastores.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "metastores.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "metastores.0.time_updated"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					GenerateDataSourceFromRepresentationMap("oci_datacatalog_metastore", "test_metastore", Required, Create, metastoreSingularDataSourceRepresentation) +
					compartmentIdVariableStr + MetastoreResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "metastore_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "default_external_table_location", defaultExternalTableLocation),
					resource.TestCheckResourceAttr(resourceName, "default_managed_table_location", defaultManagedTableLocation),
					//resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + MetastoreResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckDatacatalogMetastoreDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).dataCatalogClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datacatalog_metastore" {
			noResourceFound = false
			request := oci_datacatalog.GetMetastoreRequest{}

			tmp := rs.Primary.ID
			request.MetastoreId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "datacatalog")

			response, err := client.GetMetastore(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_datacatalog.LifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !InSweeperExcludeList("DatacatalogMetastore") {
		resource.AddTestSweepers("DatacatalogMetastore", &resource.Sweeper{
			Name:         "DatacatalogMetastore",
			Dependencies: DependencyGraph["metastore"],
			F:            sweepDatacatalogMetastoreResource,
		})
	}
}

func sweepDatacatalogMetastoreResource(compartment string) error {
	dataCatalogClient := GetTestClients(&schema.ResourceData{}).dataCatalogClient()
	metastoreIds, err := getMetastoreIds(compartment)
	if err != nil {
		return err
	}
	for _, metastoreId := range metastoreIds {
		if ok := SweeperDefaultResourceId[metastoreId]; !ok {
			deleteMetastoreRequest := oci_datacatalog.DeleteMetastoreRequest{}

			deleteMetastoreRequest.MetastoreId = &metastoreId

			deleteMetastoreRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "datacatalog")
			_, error := dataCatalogClient.DeleteMetastore(context.Background(), deleteMetastoreRequest)
			if error != nil {
				fmt.Printf("Error deleting Metastore %s %s, It is possible that the resource is already deleted. Please verify manually \n", metastoreId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &metastoreId, metastoreSweepWaitCondition, time.Duration(3*time.Minute),
				metastoreSweepResponseFetchOperation, "datacatalog", true)
		}
	}
	return nil
}

func getMetastoreIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "MetastoreId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataCatalogClient := GetTestClients(&schema.ResourceData{}).dataCatalogClient()

	listMetastoresRequest := oci_datacatalog.ListMetastoresRequest{}
	listMetastoresRequest.CompartmentId = &compartmentId
	listMetastoresRequest.LifecycleState = oci_datacatalog.ListMetastoresLifecycleStateActive
	listMetastoresResponse, err := dataCatalogClient.ListMetastores(context.Background(), listMetastoresRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Metastore list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, metastore := range listMetastoresResponse.Items {
		id := *metastore.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "MetastoreId", id)
		//SweeperDefaultResourceId[*metastore.DefaultExternalTableLocation] = true
		//SweeperDefaultResourceId[*metastore.DefaultManagedTableLocation] = true

	}
	return resourceIds, nil
}

func metastoreSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if metastoreResponse, ok := response.Response.(oci_datacatalog.GetMetastoreResponse); ok {
		return metastoreResponse.LifecycleState != oci_datacatalog.LifecycleStateDeleted
	}
	return false
}

func metastoreSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.dataCatalogClient().GetMetastore(context.Background(), oci_datacatalog.GetMetastoreRequest{
		MetastoreId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

//Function to get ObjectStorage Namespace
func getObjectStorageNamespace(compartmentId string) string {
	compartment := compartmentId
	objectStorageClient := GetTestClients(&schema.ResourceData{}).objectStorageClient()
	ctx := context.Background()
	request := objectstorage.GetNamespaceRequest{CompartmentId: common.String(compartment)}
	r, err := objectStorageClient.GetNamespace(ctx, request)
	if err != nil {
		err := fmt.Errorf("Error getting namespace : %v", err)
		fmt.Println(err.Error())
	}
	return *r.Value
}
