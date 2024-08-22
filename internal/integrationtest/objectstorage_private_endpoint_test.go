package integrationtest

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_object_storage "github.com/oracle/oci-go-sdk/v65/objectstorage"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

var (
	//ObjectStoragePrivateEndpointRequiredOnlyResource = ObjectStoragePrivateEndpointResourceDependencies +
	//	acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_private_endpoint", "test_pe", acctest.Required, acctest.Create, ObjectStoragePrivateEndpointRepresentation)
	//
	//ObjectStoragePrivateEndpointResourceConfig = ObjectStoragePrivateEndpointResourceDependencies +
	//	acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_private_endpoint", "test_pe", acctest.Optional, acctest.Update, ObjectStoragePrivateEndpointRepresentation)

	// Generate test values
	testPeName = "testPe1"
	//testPeNameUpdated             = testPeName + "2"
	testPrefix = "testPrefix1"
	//testPrefixUpdated             = "testPrefixUpdated"
	//testAdditionalPrefixes        = []string{utils.RandomStringOrHttpReplayValue(32, utils.Charset, "pePrefix")}
	//testAdditionalPrefixesUpdated = []string{utils.RandomStringOrHttpReplayValue(32, utils.Charset, "pePrefix")}
	//testAccessTargetsDetails      = oci_object_storage.AccessTargetDetails{
	//	Namespace:     String("*"),
	//	CompartmentId: String("*"),
	//	Bucket:        String("*"),
	//}
	//testAccessTargets = []oci_object_storage.AccessTargetDetails{
	//	testAccessTargetsDetails,
	//}
	//testAccessTargetsUpdated = []oci_object_storage.AccessTargetDetails{
	//	testAccessTargetsDetails,
	//}
	//
	//testIp        = "10.0.0.1"
	//testIpUpdated = "10.0.0.2"

	//ObjectStoragePrivateEndpointSingularDataSourceRepresentation = map[string]interface{}{
	//	"name":                acctest.Representation{RepType: acctest.Required, Create: testPeNameUpdated},
	//	"namespace":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	//	"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	//	"subnet_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_private_endpoint.test_pe.subnet_id}`},
	//	"prefix":              acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_private_endpoint.test_pe.prefix}`},
	//	"access_targets":      acctest.Representation{RepType: acctest.Required, Create: testAccessTargets, Update: testAccessTargetsUpdated},
	//	"private_endpoint_ip": acctest.Representation{RepType: acctest.Optional, Create: testIp, Update: testIpUpdated},
	//	"additional_prefixes": acctest.Representation{RepType: acctest.Optional, Create: testAdditionalPrefixes, Update: testAdditionalPrefixesUpdated},
	//	"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
	//	"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	//}
	//
	//ObjectStoragePrivateEndpointDataSourceRepresentation = map[string]interface{}{
	//	"name":                acctest.Representation{RepType: acctest.Required, Create: testPeNameUpdated},
	//	"namespace":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	//	"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	//	"subnet_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
	//	"prefix":              acctest.Representation{RepType: acctest.Required, Create: testPrefix, Update: testPrefixUpdated},
	//	"access_targets":      acctest.Representation{RepType: acctest.Required, Create: testAccessTargets, Update: testAccessTargetsUpdated},
	//	"private_endpoint_ip": acctest.Representation{RepType: acctest.Optional, Create: testIp, Update: testIpUpdated},
	//	"additional_prefixes": acctest.Representation{RepType: acctest.Optional, Create: testAdditionalPrefixes, Update: testAdditionalPrefixesUpdated},
	//	"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
	//	"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	//}
	//
	//ObjectStoragePrivateEndpointRepresentation = map[string]interface{}{
	//	"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_private_endpoint.test_pe.compartment_id}`},
	//	"name":                acctest.Representation{RepType: acctest.Required, Create: testPeName, Update: testPeNameUpdated},
	//	"namespace":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	//	"subnet_id":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_private_endpoint.test_pe.subnet_id}`},
	//	"prefix":              acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_private_endpoint.test_pe.prefix}`},
	//	"access_targets":      acctest.Representation{RepType: acctest.Required, Create: testAccessTargets, Update: testAccessTargetsUpdated},
	//	"private_endpoint_ip": acctest.Representation{RepType: acctest.Optional, Create: testIp, Update: testIpUpdated},
	//	"additional_prefixes": acctest.Representation{RepType: acctest.Optional, Create: testAdditionalPrefixes, Update: testAdditionalPrefixesUpdated},
	//	"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
	//	"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	//}
	//
	//ObjectStoragePrivateEndpointResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation) +
	//	acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
	//	acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) + DefinedTagsDependencies +
	//	KeyResourceDependencyConfig2
)

func TestObjectStoragePrivateEndpointResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestObjectStoragePrivateEndpointResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.LegacyTestProviderConfig()

	createConfig := config +
		`
			data "oci_identity_availability_domains" "ads1" {
				compartment_id = "${var.compartment_id}"
			}
	
			data "oci_objectstorage_namespace" "t1" {
				compartment_id = "${var.compartment_id}"
			}
	
			resource "oci_core_virtual_network" "test_vcn_1" {
				cidr_block     = "10.0.0.0/16"
				compartment_id = "${var.compartment_id}"
				display_name   = "network_name"
				dns_label      = "myvcn"
			}
	
			resource "oci_objectstorage_private_endpoint" "testPe1" {
				compartment_id = "${var.compartment_id}"
				namespace = "${data.oci_objectstorage_namespace.t1.namespace}"
				name = "testPe1"
				subnet_id = "${oci_core_subnet.test_subnet_1.id}"
				prefix = "testPrefix1"
				access_targets  {
					namespace = "*"
					compartment_id = "*"
					bucket = "*"
				  }
			}
	
			resource "oci_core_subnet" "test_subnet_1" {
				availability_domain = "${data.oci_identity_availability_domains.ads1.availability_domains.0.name}"
				cidr_block          = "10.0.1.0/24"
				display_name        = "-tf-subnet-1"
				compartment_id      = "${var.compartment_id}"
				vcn_id              = "${oci_core_virtual_network.test_vcn_1.id}"
				route_table_id      = "${oci_core_virtual_network.test_vcn_1.default_route_table_id}"
				dhcp_options_id     = "${oci_core_virtual_network.test_vcn_1.default_dhcp_options_id}"
				dns_label           = "testsubnet1"
			}
		`

	updateConfig := config +
		`
			data "oci_identity_availability_domains" "ads1" {
				compartment_id = "${var.compartment_id}"
			}
	
			data "oci_objectstorage_namespace" "t1" {
				compartment_id = "${var.compartment_id}"
			}
	
			resource "oci_core_virtual_network" "test_vcn_1" {
				cidr_block     = "10.0.0.0/16"
				compartment_id = "${var.compartment_id}"
				display_name   = "network_name"
				dns_label      = "myvcn"
			}
	
			resource "oci_objectstorage_private_endpoint" "testPe1" {
				compartment_id = "${var.compartment_id}"
				namespace = "${data.oci_objectstorage_namespace.t1.namespace}"
				name = "testPe1"
				subnet_id = "${oci_core_subnet.test_subnet_1.id}"
				prefix = "testPrefix1"
				access_targets  {
					namespace = "n1"
					compartment_id = "${var.compartment_id}"
					bucket = "b1"
				  }
			}
	
			resource "oci_core_subnet" "test_subnet_1" {
				availability_domain = "${data.oci_identity_availability_domains.ads1.availability_domains.0.name}"
				cidr_block          = "10.0.1.0/24"
				display_name        = "-tf-subnet-1"
				compartment_id      = "${var.compartment_id}"
				vcn_id              = "${oci_core_virtual_network.test_vcn_1.id}"
				route_table_id      = "${oci_core_virtual_network.test_vcn_1.default_route_table_id}"
				dhcp_options_id     = "${oci_core_virtual_network.test_vcn_1.default_dhcp_options_id}"
				dns_label           = "testsubnet1"
			}
		`

	listConfig := config +
		`
			data "oci_identity_availability_domains" "ads1" {
				compartment_id = "${var.compartment_id}"
			}
		
			data "oci_objectstorage_namespace" "t1" {
				compartment_id = "${var.compartment_id}"
			}
		
			resource "oci_core_virtual_network" "test_vcn_1" {
				cidr_block     = "10.0.0.0/16"
				compartment_id = "${var.compartment_id}"
				display_name   = "network_name"
				dns_label      = "myvcn"
			}
		
			resource "oci_objectstorage_private_endpoint" "testPe1" {
				compartment_id = "${var.compartment_id}"
				namespace = "${data.oci_objectstorage_namespace.t1.namespace}"
				name = "testPe1"
				subnet_id = "${oci_core_subnet.test_subnet_1.id}"
				prefix = "testPrefix1"
				access_targets  {
					namespace = "*"
					compartment_id = "*"
					bucket = "*"
				  }
			}

			data "oci_objectstorage_private_endpoint_summaries" "testlist" {
				compartment_id = "${var.compartment_id}"
				namespace = "${data.oci_objectstorage_namespace.t1.namespace}"
				filter {
					name   = "name"
    				values = [oci_objectstorage_private_endpoint.testPe1.name]
				}
			}
		
			resource "oci_core_subnet" "test_subnet_1" {
				availability_domain = "${data.oci_identity_availability_domains.ads1.availability_domains.0.name}"
				cidr_block          = "10.0.1.0/24"
				display_name        = "-tf-subnet-1"
				compartment_id      = "${var.compartment_id}"
				vcn_id              = "${oci_core_virtual_network.test_vcn_1.id}"
				route_table_id      = "${oci_core_virtual_network.test_vcn_1.default_route_table_id}"
				dhcp_options_id     = "${oci_core_virtual_network.test_vcn_1.default_dhcp_options_id}"
				dns_label           = "testsubnet1"
			}
		`

	//compartmentId2 := utils.GetEnvSettingWithBlankDefault("compartment_id_for_update")
	//compartmentId2VariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_objectstorage_private_endpoint.testPe1"
	datasourceName := "data.oci_objectstorage_private_endpoint_summaries.testlist"
	//singularDatasourceName := "oci_objectstorage_private_endpoint.testPe1"

	//var resId, resId2 string
	//acctest.SaveConfigContent(createConfig, "objectstorage", "private_endpoint", t)

	acctest.ResourceTest(t, testAccCheckObjectStoragePrivateEndpointDestroy, []resource.TestStep{
		// Verify Create
		{
			Config: createConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "name", testPeName),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "prefix", testPrefix),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "access_targets",
					map[string]string{
						"namespace":      "*",
						"compartment_id": "*",
						"bucket":         "*",
					},
					[]string{}),
			),
		},

		// delete before next Create
		{
			Config: config,
		},

		// VerifyUpdate
		{
			Config: updateConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "name", testPeName),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "prefix", testPrefix),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "access_targets",
					map[string]string{
						"namespace":      "n1",
						"compartment_id": compartmentId,
						"bucket":         "b1",
					},
					[]string{}),
			),
		},
		{
			Config: config,
		},
		//
		// verify datasource
		{
			Config: listConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "namespace"),
				resource.TestCheckResourceAttr(datasourceName, "private_endpoint_summaries.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "private_endpoint_summaries.0.created_by"),
				resource.TestCheckResourceAttrSet(datasourceName, "private_endpoint_summaries.0.etag"),
				resource.TestCheckResourceAttrSet(datasourceName, "private_endpoint_summaries.0.namespace"),
				resource.TestCheckResourceAttrSet(datasourceName, "private_endpoint_summaries.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "private_endpoint_summaries.0.prefix", testPrefix),
			),
		},
	})
}

func testAccCheckObjectStoragePrivateEndpointDestroy(state *terraform.State) error {
	noResourceFound := false
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ObjectStorageClient()
	for _, rs := range state.RootModule().Resources {
		if rs.Type == "object_storage_private_endpoint" {
			noResourceFound = true
			request := oci_object_storage.GetPrivateEndpointRequest{}

			if value, ok := rs.Primary.Attributes["name"]; ok {
				request.PeName = &value
			}

			if value, ok := rs.Primary.Attributes["namespace"]; ok {
				request.NamespaceName = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "object_storage")
			_, err := client.GetPrivateEndpoint(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("ObjectStoragePrivateEndpoint") {
		resource.AddTestSweepers("ObjectStoragePrivateEndpoint", &resource.Sweeper{
			Name:         "ObjectStoragePrivateEndpoint",
			Dependencies: acctest.DependencyGraph["private_endpoint"],
			F:            sweepObjectStoragePrivateEndpointResource,
		})
	}
}

func sweepObjectStoragePrivateEndpointResource(compartment string) error {
	objectStorageClient := acctest.GetTestClients(&schema.ResourceData{}).ObjectStorageClient()
	bucketIds, err := getObjectStoragePrivateEndpointIds(compartment)
	if err != nil {
		return err
	}
	for _, bucketId := range bucketIds {
		if ok := acctest.SweeperDefaultResourceId[bucketId]; !ok {
			deleteBucketRequest := oci_object_storage.DeleteBucketRequest{}

			deleteBucketRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "object_storage")
			_, error := objectStorageClient.DeleteBucket(context.Background(), deleteBucketRequest)
			if error != nil {
				fmt.Printf("Error deleting Bucket %s %s, It is possible that the resource is already deleted. Please verify manually \n", bucketId, error)
				continue
			}
		}
	}
	return nil
}

func getObjectStoragePrivateEndpointIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "peName")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	objectStorageClient := acctest.GetTestClients(&schema.ResourceData{}).ObjectStorageClient()

	listPrivateEndpointRequest := oci_object_storage.ListPrivateEndpointsRequest{}
	listPrivateEndpointRequest.CompartmentId = &compartmentId

	namespaces, error := getNamespaces(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting namespace required for PrivateEndpoint resource requests \n")
	}
	for _, namespace := range namespaces {
		listPrivateEndpointRequest.NamespaceName = &namespace

		listPrivateEndpointResponse, err := objectStorageClient.ListPrivateEndpoints(context.Background(), listPrivateEndpointRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting PrivateEndpoint list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, peItem := range listPrivateEndpointResponse.Items {
			id := *peItem.Name
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "peName", id)
		}

	}
	return resourceIds, nil
}
