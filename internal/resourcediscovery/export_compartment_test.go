// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resourcediscovery

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"regexp"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/globalvar"
	tf_provider "github.com/terraform-providers/terraform-provider-oci/internal/provider"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	resourceDiscoveryTestCompartmentOcid   = "ocid1.testcompartment.abc"
	resourceDiscoveryTestTenancyOcid       = "ocid1.testtenancy.xyz"
	resourceDiscoveryTestActiveLifecycle   = "ACTIVE"
	resourceDiscoveryTestInactiveLifecycle = "INACTIVE"
	resourceIdFor404ErrorResource          = "ocid1.child.abcdefghiklmnop.1"
)

var exportParentDefinition = &TerraformResourceHints{
	resourceClass:               "oci_test_parent",
	datasourceClass:             "oci_test_parents",
	resourceAbbreviation:        "parent",
	datasourceItemsAttr:         "items",
	discoverableLifecycleStates: []string{resourceDiscoveryTestActiveLifecycle},
	alwaysExportable:            true,
}

var exportChildDefinition = &TerraformResourceHints{
	resourceClass:               "oci_test_child",
	datasourceClass:             "oci_test_children",
	resourceAbbreviation:        "child",
	datasourceItemsAttr:         "item_summaries",
	discoverableLifecycleStates: []string{resourceDiscoveryTestActiveLifecycle},
	requireResourceRefresh:      true,
}

var exportParentDefinitionWithFaultyDatasource = &TerraformResourceHints{
	resourceClass:               "oci_test_parent",
	datasourceClass:             "oci_test_error_parents",
	resourceAbbreviation:        "parent",
	datasourceItemsAttr:         "items",
	discoverableLifecycleStates: []string{resourceDiscoveryTestActiveLifecycle},
	alwaysExportable:            true,
}

var exportChildDefinitionWithFaultyDatasource = &TerraformResourceHints{
	resourceClass:               "oci_test_error_child",
	datasourceClass:             "oci_test_children",
	resourceAbbreviation:        "child",
	datasourceItemsAttr:         "item_summaries",
	discoverableLifecycleStates: []string{resourceDiscoveryTestActiveLifecycle},
	requireResourceRefresh:      true,
}

var exportResourceDefinitionWith404Error = &TerraformResourceHints{
	resourceClass:               "oci_test_404_error_child",
	datasourceClass:             "oci_test_children",
	resourceAbbreviation:        "child",
	datasourceItemsAttr:         "item_summaries",
	discoverableLifecycleStates: []string{resourceDiscoveryTestActiveLifecycle},
	requireResourceRefresh:      true,
}

var exportResourceDefinitionWithPanic = &TerraformResourceHints{
	resourceClass:               "oci_test_child",
	datasourceClass:             "oci_test_panic_children",
	resourceAbbreviation:        "child",
	datasourceItemsAttr:         "item_summaries",
	discoverableLifecycleStates: []string{resourceDiscoveryTestActiveLifecycle},
}

var tenancyTestingResourceGraph = TerraformResourceGraph{
	"oci_identity_tenancy": {
		{
			TerraformResourceHints: exportParentDefinition,
		},
	},
	"oci_test_parent": {
		{
			TerraformResourceHints: exportChildDefinition,
			datasourceQueryParams:  map[string]string{"parent_id": "id"},
		},
	},
}

var compartmentTestingResourceGraph = TerraformResourceGraph{
	"oci_identity_compartment": {
		{
			TerraformResourceHints: exportParentDefinition,
		},
	},
	"oci_test_parent": {
		{
			TerraformResourceHints: exportChildDefinition,
			datasourceQueryParams:  map[string]string{"parent_id": "id"},
		},
	},
}

var compartmentTestingResourceGraphWithFaultyParentResource = TerraformResourceGraph{
	"oci_identity_compartment": {
		{
			TerraformResourceHints: exportParentDefinitionWithFaultyDatasource,
		},
	},
}

var compartmentTestingResourceGraphWithFaultyChildResource = TerraformResourceGraph{
	"oci_identity_compartment": {
		{
			TerraformResourceHints: exportParentDefinition,
		},
	},
	"oci_test_parent": {
		{
			TerraformResourceHints: exportChildDefinitionWithFaultyDatasource,
			datasourceQueryParams:  map[string]string{"parent_id": "id"},
		},
	},
}

var compartmentTestingResourceGraphWith404ErrorResource = TerraformResourceGraph{
	"oci_identity_compartment": {
		{
			TerraformResourceHints: exportParentDefinition,
		},
	},
	"oci_test_parent": {
		{
			TerraformResourceHints: exportResourceDefinitionWith404Error,
			datasourceQueryParams:  map[string]string{"parent_id": "id"},
		},
	},
}

var compartmentTestingResourceGraphWithPanicResource = TerraformResourceGraph{
	"oci_identity_compartment": {
		{
			TerraformResourceHints: exportParentDefinition,
		},
	},
	"oci_test_parent": {
		{
			TerraformResourceHints: exportResourceDefinitionWithPanic,
			datasourceQueryParams:  map[string]string{"parent_id": "id"},
		},
		{
			TerraformResourceHints: exportChildDefinition,
			datasourceQueryParams:  map[string]string{"parent_id": "id"},
		},
	},
}

var childrenResources map[string]map[string]interface{}
var parentResources map[string]map[string]interface{}

// Test resources used by resource discovery tests
func testParentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Create: createTestParent,
		Read:   readTestParent,
		Delete: deleteTestParent,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"a_map": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"a_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"a_bool": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"a_int": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"a_float": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"a_list": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
				ForceNew: true,
			},
			"a_set": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
				ForceNew: true,
				Set:      utils.LiteralTypeHashCodeForSets,
			},
			"a_nested": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nested_string": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"nested_bool": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"nested_int": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"nested_float": {
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
					},
				},
			},

			// Computed
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func testChildResource() *schema.Resource {
	// Reuse the parent schema and add a parent dependency attribute
	childResourceSchema := testParentResource().Schema
	childResourceSchema["parent_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}

	// Don't have a display_name attribute so a different name can be generated
	delete(childResourceSchema, "display_name")

	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Create: createTestChild,
		Read:   readTestChild,
		Delete: deleteTestChild,
		Schema: childResourceSchema,
	}
}

func testParentsDatasource() *schema.Resource {
	return &schema.Resource{
		Read: listTestParents,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(testParentResource()),
			},
		},
	}
}

func testChildrenDatasource() *schema.Resource {
	// Convert child resource schema to datasource schema
	childDatasourceSchema := tfresource.GetDataSourceItemSchema(testChildResource())

	// Remove some attributes from datasource (i.e. treat the datasource results as incomplete representations of the resource)
	delete(childDatasourceSchema.Schema, "a_nested")

	return &schema.Resource{
		Read: listTestChildren,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"parent_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"item_summaries": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     childDatasourceSchema,
			},
		},
	}
}

func testParentsDatasourceWithError() *schema.Resource {
	return &schema.Resource{
		Read: listTestParentsWithError,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(testParentResource()),
			},
		},
	}
}

func testParentsDatasourceWithPanic() *schema.Resource {
	return &schema.Resource{
		Read: listTestParentsWithPanic,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(testParentResource()),
			},
		},
	}
}

func testChildResourceWithError() *schema.Resource {
	// Reuse the parent schema and add a parent dependency attribute
	childResourceSchema := testParentResource().Schema
	childResourceSchema["parent_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}

	// Don't have a display_name attribute so a different name can be generated
	delete(childResourceSchema, "display_name")

	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Create: createTestChild,
		Read:   readTestChildWithError,
		Delete: deleteTestChild,
		Schema: childResourceSchema,
	}
}

func createTestParent(d *schema.ResourceData, m interface{}) error {
	return nil
}

func readTestParent(d *schema.ResourceData, m interface{}) error {
	if resource, exists := parentResources[d.Id()]; exists {
		for key, value := range resource {
			d.Set(key, value)
		}
	} else {
		return fmt.Errorf("could not find parent with id %s", d.Id())
	}
	return nil
}

// TestUnitRunExportCommand_Parallel is reusing the same test resource CRUD definitions for different entries
// in export graph for the test services, hence listTestParents get called by multiple threads
// need a lock to modify the `resource` map concurrently
var modifyParentLock sync.Mutex

func listTestParents(d *schema.ResourceData, m interface{}) error {
	results := make([]interface{}, len(parentResources))
	modifyParentLock.Lock()
	for i := 0; i < len(parentResources); i++ {
		id := getTestResourceId("parent", i)
		resource := parentResources[id]
		resource["id"] = id
		results[i] = resource
	}
	d.Set("items", results)
	modifyParentLock.Unlock()
	return nil
}

func listTestParentsWithError(d *schema.ResourceData, m interface{}) error {
	return fmt.Errorf("could not find resources: error in listTestParentsWithError")
}

func listTestParentsWithPanic(d *schema.ResourceData, m interface{}) error {
	panic("panic from listTestParentsWithPanic")
}

func deleteTestParent(d *schema.ResourceData, m interface{}) error {
	return nil
}

func createTestChild(d *schema.ResourceData, m interface{}) error {
	return nil
}

func readTestChild(d *schema.ResourceData, m interface{}) error {
	modifyChildLock.RLock()
	if resource, exists := childrenResources[d.Id()]; exists {
		for key, value := range resource {
			d.Set(key, value)
		}
	} else {
		modifyChildLock.RUnlock()
		return fmt.Errorf("could not find child with id %s", d.Id())
	}
	modifyChildLock.RUnlock()
	return nil
}

func readTestChildWithError(d *schema.ResourceData, m interface{}) error {
	return fmt.Errorf("could not find child with id %s", d.Id())
}

// TestUnitRunExportCommand_Parallel is reusing the same test resource CRUD definitions for different entries
// in export graph for the test services, hence listTestChildren get called by multiple threads
// need a lock to modify the `resource` map concurrently
var modifyChildLock sync.RWMutex

func listTestChildren(d *schema.ResourceData, m interface{}) error {

	parentId, parentIdExists := d.GetOkExists("parent_id")
	results := []interface{}{}
	modifyChildLock.Lock()
	for i := 0; i < len(childrenResources); i++ {
		id := getTestResourceId("child", i)
		resource := childrenResources[id]
		resource["id"] = id

		if !parentIdExists || resource["parent_id"] == parentId {
			copyResource := map[string]interface{}{}
			for key, val := range resource {
				if key == "a_nested" {
					continue
				}
				copyResource[key] = val
			}
			results = append(results, copyResource)
		}
	}
	if err := d.Set("item_summaries", results); err != nil {
		return err
	}
	modifyChildLock.Unlock()
	return nil
}

func deleteTestChild(d *schema.ResourceData, m interface{}) error {
	return nil
}

func testChildResourceWith404Error() *schema.Resource {
	// Reuse the parent schema and add a parent dependency attribute
	childResourceSchema := testParentResource().Schema
	childResourceSchema["parent_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	}

	// Don't have a display_name attribute so a different name can be generated
	delete(childResourceSchema, "display_name")

	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Create: createTestChild,
		Read:   readTestChildWith404Error,
		Delete: deleteTestChild,
		Schema: childResourceSchema,
	}
}

func readTestChildWith404Error(d *schema.ResourceData, m interface{}) error {
	sync := &TestChildWith404ErrorResourceCrud{}
	sync.D = d

	return tfresource.ReadResource(sync)
}

type TestChildWith404ErrorResourceCrud struct {
	tfresource.BaseCrud
}

func (s TestChildWith404ErrorResourceCrud) Get() error {
	if s.D.Id() == resourceIdFor404ErrorResource {
		return fmt.Errorf("404 not found")
	} else {
		return readTestChild(s.D, nil)
	}

}

func (s TestChildWith404ErrorResourceCrud) SetData() error {
	return nil
}

func initResourceDiscoveryTests() {
	resourceNameCount = map[string]int{}
	resourcesMap = tf_provider.ResourcesMap()
	datasourcesMap = tf_provider.DataSourcesMap()
	tfHclVersion = &TfHclVersion12{}

	resourcesMap["oci_test_parent"] = testParentResource()
	resourcesMap["oci_test_child"] = testChildResource()
	resourcesMap["oci_test_error_child"] = testChildResourceWithError()
	resourcesMap["oci_test_404_error_child"] = testChildResourceWith404Error()

	datasourcesMap["oci_test_parents"] = testParentsDatasource()
	datasourcesMap["oci_test_children"] = testChildrenDatasource()
	datasourcesMap["oci_test_error_parents"] = testParentsDatasourceWithError()

	datasourcesMap["oci_test_panic_children"] = testParentsDatasourceWithPanic()

	tenancyResourceGraphs["tenancy_testing"] = tenancyTestingResourceGraph
	compartmentResourceGraphs["compartment_testing"] = compartmentTestingResourceGraph

	initTestResources()
}

func cleanupResourceDiscoveryTests() {
	delete(resourcesMap, "oci_test_parent")
	delete(resourcesMap, "oci_test_child")
	delete(resourcesMap, "oci_test_error_child")
	delete(datasourcesMap, "oci_test_parents")
	delete(datasourcesMap, "oci_test_children")
	delete(datasourcesMap, "oci_test_error_children")
	delete(datasourcesMap, "oci_test_panic_children")
	delete(tenancyResourceGraphs, "tenancy_testing")
	delete(compartmentResourceGraphs, "compartment_testing")
}

func initTestResources() {
	numParentResources := 4
	if parentResources == nil || len(parentResources) != numParentResources {
		parentResources = make(map[string]map[string]interface{}, numParentResources)
		for i := 0; i < numParentResources; i++ {
			parentResources[getTestResourceId("parent", i)] = generateTestResourceFromSchema(i, resourcesMap["oci_test_parent"].Schema)
		}
	}

	numChildrenResourcesPerParent := 2
	numChildrenResources := numParentResources * numChildrenResourcesPerParent
	if childrenResources == nil || len(childrenResources) != numChildrenResources {
		childrenResources = make(map[string]map[string]interface{}, numParentResources*numChildrenResourcesPerParent)
		childCount := 0
		for i := 0; i < len(parentResources); i++ {
			parentId := getTestResourceId("parent", i)
			for j := 0; j < numChildrenResourcesPerParent; j++ {
				childResource := generateTestResourceFromSchema(i, resourcesMap["oci_test_child"].Schema)
				childResource["parent_id"] = parentId

				childrenResources[getTestResourceId("child", childCount)] = childResource
				childCount++
			}
		}
	}
}

func getRootCompartmentResource() *OCIResource {
	return &OCIResource{
		compartmentId: resourceDiscoveryTestCompartmentOcid,
		TerraformResource: TerraformResource{
			id:             resourceDiscoveryTestCompartmentOcid,
			terraformClass: "oci_identity_compartment",
			terraformName:  "export",
		},
	}
}

func getTestResourceId(resourceType string, id int) string {
	return fmt.Sprintf("ocid1.%s.abcdefghiklmnop.%d", resourceType, id)
}

func generatePrimitiveValue(id int, valueType schema.ValueType) interface{} {
	switch valueType {
	case schema.TypeInt:
		return id
	case schema.TypeBool:
		return true
	case schema.TypeFloat:
		res, _ := strconv.ParseFloat(fmt.Sprintf("%d.%d", id, id), 64)
		return res
	case schema.TypeString:
		return fmt.Sprintf("string%d", id)
	}
	return nil
}

func generateTestResourceFromSchema(id int, resourceSchemaMap map[string]*schema.Schema) map[string]interface{} {
	result := map[string]interface{}{}
	for resourceAttribute, resourceSchema := range resourceSchemaMap {
		switch resourceAttribute {
		case "state":
			result[resourceAttribute] = resourceDiscoveryTestActiveLifecycle
			continue
		case "time_created":
			result[resourceAttribute] = time.Now().Format(time.RFC3339)
			continue
		}

		switch resourceSchema.Type {
		case schema.TypeInt, schema.TypeBool, schema.TypeFloat, schema.TypeString:
			result[resourceAttribute] = generatePrimitiveValue(id, resourceSchema.Type)
		case schema.TypeMap:
			mapResult := map[string]interface{}{}
			if elemType, ok := resourceSchema.Elem.(schema.ValueType); ok {
				for i := 0; i < id; i++ {
					mapKey := fmt.Sprintf("key%d", i)
					mapResult[mapKey] = generatePrimitiveValue(i, elemType)
				}
			}
			result[resourceAttribute] = mapResult
		case schema.TypeList, schema.TypeSet:
			listResult := make([]interface{}, id)
			switch elemType := resourceSchema.Elem.(type) {
			case schema.ValueType:
				for i := 0; i < id; i++ {
					listResult[i] = generatePrimitiveValue(i, elemType)
				}
			case *schema.Resource:
				for i := 0; i < id; i++ {
					listResult[i] = generateTestResourceFromSchema(i, elemType.Schema)
				}
			}
			result[resourceAttribute] = listResult
		}
	}
	return result
}

// Basic test to ensure that RunExportCommand generates TF artifacts
// issue-routing-tag: terraform/default
func TestUnitRunExportCommand_basic(t *testing.T) {
	initResourceDiscoveryTests()
	defer cleanupResourceDiscoveryTests()
	compartmentId := resourceDiscoveryTestCompartmentOcid
	if err := os.Setenv("export_tenancy_id", resourceDiscoveryTestTenancyOcid); err != nil {
		t.Logf("unable to set export_tenancy_id. err: %v", err)
		t.Fail()
	}
	outputDir, err := os.Getwd()
	outputDir = fmt.Sprintf("%s%sdiscoveryTest-%d", outputDir, string(os.PathSeparator), time.Now().Nanosecond())
	if err = os.Mkdir(outputDir, os.ModePerm); err != nil {
		t.Logf("unable to mkdir %s. err: %v", outputDir, err)
		t.Fail()
	}

	tfHclVersions := []TfHclVersion{&TfHclVersion11{}, &TfHclVersion12{}}
	for _, tfVersion := range tfHclVersions {
		tfHclVersion = tfVersion
		args := &ExportCommandArgs{
			CompartmentId: &compartmentId,
			Services:      []string{"compartment_testing", "tenancy_testing"},
			OutputDir:     &outputDir,
			GenerateState: false,
			TFVersion:     &tfHclVersion,
			Parallelism:   1,
		}

		if err, _ = RunExportCommand(args); err != nil {
			t.Logf("(TF version %s) export command failed due to err: %v", tfHclVersion.toString(), err)
			t.Fail()
		}

		if _, err = os.Stat(fmt.Sprintf("%s%stenancy_testing.tf", outputDir, string(os.PathSeparator))); !os.IsNotExist(err) {
			t.Logf("(TF version %s) tenancy_testing.tf file generated even though it wasn't expected", tfHclVersion.toString())
			t.Fail()
		}

		if _, err = os.Stat(fmt.Sprintf("%s%scompartment_testing.tf", outputDir, string(os.PathSeparator))); os.IsNotExist(err) {
			t.Logf("(TF version %s) no compartment_testing.tf file generated", tfHclVersion.toString())
			t.Fail()
		}

		if _, err = os.Stat(fmt.Sprintf("%s%sterraform.tfstate", outputDir, string(os.PathSeparator))); !os.IsNotExist(err) {
			t.Logf("(TF version %s) found terraform.tfstate even though it wasn't expected", tfHclVersion.toString())
		}
	}

	os.RemoveAll(outputDir)
}

// issue-routing-tag: terraform/default
func TestUnitRunExportCommand_Parallel(t *testing.T) {
	initResourceDiscoveryTests()
	// add more services to compartment graphs
	compartmentResourceGraphs["compartment_testing_1"] = compartmentTestingResourceGraph
	compartmentResourceGraphs["compartment_testing_2"] = compartmentTestingResourceGraph
	compartmentResourceGraphs["compartment_testing_3"] = compartmentTestingResourceGraph
	compartmentResourceGraphs["compartment_testing_4"] = compartmentTestingResourceGraph
	compartmentResourceGraphs["compartment_testing_5"] = compartmentTestingResourceGraph
	compartmentResourceGraphs["compartment_testing_6"] = compartmentTestingResourceGraph
	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack()
			t.Logf("[ERROR] panic in RunExportCommand: unknown error occurred in test export")
			t.Fail()
		}
		cleanupResourceDiscoveryTests()
	}()
	compartmentId := resourceDiscoveryTestCompartmentOcid
	if err := os.Setenv("export_tenancy_id", resourceDiscoveryTestTenancyOcid); err != nil {
		t.Logf("unable to set export_tenancy_id. err: %v", err)
		t.Fail()
	}
	outputDir, err := os.Getwd()
	outputDir = fmt.Sprintf("%s%sdiscoveryTest-%d", outputDir, string(os.PathSeparator), time.Now().Nanosecond())
	if err = os.Mkdir(outputDir, os.ModePerm); err != nil {
		t.Logf("unable to mkdir %s. err: %v", outputDir, err)
		t.Fail()
	}

	tfHclVersions := []TfHclVersion{&TfHclVersion11{}, &TfHclVersion12{}}
	for _, tfVersion := range tfHclVersions {
		tfHclVersion = tfVersion
		args := &ExportCommandArgs{
			CompartmentId: &compartmentId,
			Services:      []string{"compartment_testing", "compartment_testing_1", "compartment_testing_2", "compartment_testing_3", "compartment_testing_4", "compartment_testing_5", "compartment_testing_6", "tenancy_testing"},
			OutputDir:     &outputDir,
			GenerateState: false,
			TFVersion:     &tfHclVersion,
			Parallelism:   4,
		}

		if err, _ = RunExportCommand(args); err != nil {
			t.Logf("(TF version %s) export command failed due to err: %v", tfHclVersion.toString(), err)
			t.Fail()
		}

		if _, err = os.Stat(fmt.Sprintf("%s%stenancy_testing.tf", outputDir, string(os.PathSeparator))); !os.IsNotExist(err) {
			t.Logf("(TF version %s) tenancy_testing.tf file generated even though it wasn't expected", tfHclVersion.toString())
			t.Fail()
		}

		if _, err = os.Stat(fmt.Sprintf("%s%scompartment_testing.tf", outputDir, string(os.PathSeparator))); os.IsNotExist(err) {
			t.Logf("(TF version %s) no compartment_testing.tf file generated", tfHclVersion.toString())
			t.Fail()
		}

		if _, err = os.Stat(fmt.Sprintf("%s%sterraform.tfstate", outputDir, string(os.PathSeparator))); !os.IsNotExist(err) {
			t.Logf("(TF version %s) found terraform.tfstate even though it wasn't expected", tfHclVersion.toString())
		}
	}

	os.RemoveAll(outputDir)
}

// issue-routing-tag: terraform/default
func TestUnitRunExportCommand_ParallelNegative(t *testing.T) {
	initResourceDiscoveryTests()

	defer cleanupResourceDiscoveryTests()
	compartmentId := resourceDiscoveryTestCompartmentOcid
	if err := os.Setenv("export_tenancy_id", resourceDiscoveryTestTenancyOcid); err != nil {
		t.Logf("unable to set export_tenancy_id. err: %v", err)
		t.Fail()
	}
	outputDir, err := os.Getwd()
	outputDir = fmt.Sprintf("%s%sdiscoveryTest-%d", outputDir, string(os.PathSeparator), time.Now().Nanosecond())
	if err = os.Mkdir(outputDir, os.ModePerm); err != nil {
		t.Logf("unable to mkdir %s. err: %v", outputDir, err)
		t.Fail()
	}

	tfHclVersions := []TfHclVersion{&TfHclVersion11{}, &TfHclVersion12{}}
	for _, tfVersion := range tfHclVersions {
		tfHclVersion = tfVersion
		args := &ExportCommandArgs{
			CompartmentId: &compartmentId,
			Services:      []string{"compartment_testing", "compartment_testing_1", "compartment_testing_2", "compartment_testing_3", "compartment_testing_4", "compartment_testing_5", "compartment_testing_6", "tenancy_testing"},
			OutputDir:     &outputDir,
			GenerateState: false,
			TFVersion:     &tfHclVersion,
			Parallelism:   -1,
		}

		if err, _ = RunExportCommand(args); err == nil {
			t.Logf("expected error but found none")
			t.Fail()
		} else {
			assert.Equal(t, "[ERROR] invalid value for arument parallelism, specify a value >= 1", err.Error())
		}
	}

	os.RemoveAll(outputDir)
}

// issue-routing-tag: terraform/default
func TestUnitRunExportCommand_error(t *testing.T) {
	initResourceDiscoveryTests()
	defer cleanupResourceDiscoveryTests()
	compartmentId := resourceDiscoveryTestCompartmentOcid
	outputDir, err := os.Getwd()
	outputDir = fmt.Sprintf("%s%sdiscoveryTest-%d", outputDir, string(os.PathSeparator), time.Now().Nanosecond())
	if err = os.Mkdir(outputDir, os.ModePerm); err != nil {
		t.Logf("unable to mkdir %s. err: %v", outputDir, err)
		t.Fail()
	}

	nonexistentOutputDir := fmt.Sprintf("%s%s%s", outputDir, string(os.PathSeparator), "baddirectory")
	tfHclVersion = &TfHclVersion12{}
	args := &ExportCommandArgs{
		CompartmentId: &compartmentId,
		Services:      []string{"compartment_testing", "tenancy_testing"},
		OutputDir:     &nonexistentOutputDir,
		GenerateState: false,
		TFVersion:     &tfHclVersion,
	}
	if err, _ = RunExportCommand(args); err == nil {
		t.Logf("export command expected to fail due to non-existent path, but it succeeded")
		t.Fail()
	}

	os.RemoveAll(outputDir)
}

// issue-routing-tag: terraform/default
func TestUnitRunExportCommand_panic(t *testing.T) {
	//compartmentId := resourceDiscoveryTestCompartmentOcid
	outputDir, err := os.Getwd()
	outputDir = fmt.Sprintf("%s%sdiscoveryTest-%d", outputDir, string(os.PathSeparator), time.Now().Nanosecond())
	if err = os.Mkdir(outputDir, os.ModePerm); err != nil {
		t.Logf("unable to mkdir %s. err: %v", outputDir, err)
		t.Fail()
	}

	tfHclVersion = &TfHclVersion12{}

	// nil args will cause panic and if panic is not handled test will fail else it will pass
	_, _ = RunExportCommand(nil)

	os.RemoveAll(outputDir)
}

// Test exit status in case of partial success
// issue-routing-tag: terraform/default
func TestUnitRunExportCommand_exitStatusForPartialSuccess(t *testing.T) {
	initResourceDiscoveryTests()
	// Replace compartmentResourceGraphs with the one having resource that has error in read
	// Status returned should be StatusPartialSuccess
	compartmentResourceGraphs["compartment_testing"] = compartmentTestingResourceGraphWithFaultyParentResource

	defer cleanupResourceDiscoveryTests()
	compartmentId := resourceDiscoveryTestCompartmentOcid
	if err := os.Setenv("export_tenancy_id", resourceDiscoveryTestTenancyOcid); err != nil {
		t.Logf("unable to set export_tenancy_id. err: %v", err)
		t.Fail()
	}
	outputDir, err := os.Getwd()
	outputDir = fmt.Sprintf("%s%sdiscoveryTest-%d", outputDir, string(os.PathSeparator), time.Now().Nanosecond())
	if err = os.Mkdir(outputDir, os.ModePerm); err != nil {
		t.Logf("unable to mkdir %s. err: %v", outputDir, err)
		t.Fail()
	}

	tfHclVersion = &TfHclVersion12{}
	args := &ExportCommandArgs{
		CompartmentId: &compartmentId,
		Services:      []string{"compartment_testing", "tenancy_testing"},
		OutputDir:     &outputDir,
		GenerateState: false,
		TFVersion:     &tfHclVersion,
		Parallelism:   1,
	}

	if err, status := RunExportCommand(args); err != nil {
		t.Logf("(TF version %s) export command failed due to err: %v", tfHclVersion.toString(), err)
		t.Fail()
	} else if status != StatusPartialSuccess {
		t.Logf("(TF version %s) export command returned unexpected Exit Status: %v", tfHclVersion.toString(), status)
		t.Fail()
	}

	if _, err = os.Stat(fmt.Sprintf("%s%scompartment_testing.tf", outputDir, string(os.PathSeparator))); os.IsNotExist(err) {
		t.Logf("(TF version %s) no compartment_testing.tf file generated", tfHclVersion.toString())
		t.Fail()
	}

	if _, err = os.Stat(fmt.Sprintf("%s%sterraform.tfstate", outputDir, string(os.PathSeparator))); !os.IsNotExist(err) {
		t.Logf("(TF version %s) found terraform.tfstate even though it wasn't expected", tfHclVersion.toString())
	}

	os.RemoveAll(outputDir)
}

// Test that resources can be found using a resource dependency graph
// issue-routing-tag: terraform/default
func TestUnitFindResources_basic(t *testing.T) {
	initResourceDiscoveryTests()
	defer cleanupResourceDiscoveryTests()
	rootResource := getRootCompartmentResource()

	ctx := &resourceDiscoveryContext{
		errorList: ErrorList{},
	}
	results, err := findResources(ctx, rootResource, compartmentTestingResourceGraph)
	if err != nil {
		t.Logf("got error from findResources: %v", err)
		t.Fail()
	}

	if len(results) != len(childrenResources)+len(parentResources) {
		t.Logf("got %d results but expected %d results", len(results), len(childrenResources)+len(parentResources))
		t.Fail()
	}

	for _, foundResource := range results {
		if foundResource.terraformClass == "oci_test_child" {
			if _, resourceRefreshAttributeExists := foundResource.sourceAttributes["a_nested"]; !resourceRefreshAttributeExists {
				t.Logf("child resource is missing an expected attribute that should have been filled by a resource refresh")
				t.Fail()
			}

			expectedTfNamePrefix := fmt.Sprintf("%s_%s", foundResource.parent.terraformName, exportChildDefinition.resourceAbbreviation)
			if !strings.HasPrefix(foundResource.terraformName, expectedTfNamePrefix) {
				t.Logf("child resource should have a name with prefix '%s' but name is '%s' instead", expectedTfNamePrefix, foundResource.terraformName)
				t.Fail()
			}
		}
	}
}

// Test that resource with 404 Not found error do not show up in results
// issue-routing-tag: terraform/default
func TestUnitFindResources_404Error(t *testing.T) {
	initResourceDiscoveryTests()
	// Replace compartmentResourceGraphs with the one having resource that has 404 error in read
	// Resource with 404 error should be skipped

	defer cleanupResourceDiscoveryTests()
	rootResource := getRootCompartmentResource()

	ctx := &resourceDiscoveryContext{
		errorList: ErrorList{},
	}
	results, err := findResources(ctx, rootResource, compartmentTestingResourceGraphWith404ErrorResource)
	if err != nil {
		t.Logf("got error from findResources: %v", err)
		t.Fail()
	}

	for _, resource := range results {
		if resource.sourceAttributes["id"] == "" {
			// State is voided but resource still showed up in results
			t.Logf("got resource with 404 not found error in results when not expected")
			t.Fail()
		}
	}

	// Check that we got all child resources except 1 that had 404 error
	if len(results) != len(parentResources)+len(childrenResources)-1 {
		t.Logf("got %d results but expected %d results", len(results), len(parentResources)+len(childrenResources)-1)
		t.Fail()
	}
}

// Test that discovery continues after panic
// issue-routing-tag: terraform/default
func TestUnitFindResources_panic(t *testing.T) {
	// env var export_enable_tenancy_lookup=false needed for this test
	initResourceDiscoveryTests()
	// Replace compartmentResourceGraphs with the one having resource that will panic

	defer cleanupResourceDiscoveryTests()
	rootResource := getRootCompartmentResource()

	ctx := &resourceDiscoveryContext{
		errorList: ErrorList{},
	}
	results, err := findResources(ctx, rootResource, compartmentTestingResourceGraphWithPanicResource)
	if err != nil {
		t.Logf("got error from findResources: %v", err)
		t.Fail()
	}

	// Check that we got all child resources except the ones that panicked
	// since we don't actually need the panic resources in result as we will never get to that point so we are not initializing those
	if len(results) != len(parentResources)+len(childrenResources) {
		t.Logf("got %d results but expected %d results", len(results), len(parentResources)+len(childrenResources)-1)
		t.Fail()
	}
}

// Test that errorList has errors if resources are not found
// issue-routing-tag: terraform/default
func TestUnitFindResources_errorList(t *testing.T) {
	initResourceDiscoveryTests()
	defer cleanupResourceDiscoveryTests()
	rootResource := getRootCompartmentResource()

	ctx := &resourceDiscoveryContext{
		errorList: ErrorList{},
	}
	_, err := findResources(ctx, rootResource, compartmentTestingResourceGraphWithFaultyChildResource)
	if err != nil {
		t.Logf("got error from findResources: %v", err)
		t.Fail()
	}
	if len(ctx.errorList.errors) == 0 {
		t.Logf("expected errors for failed resources in resourceDiscoveryContext errorList but found none")
		t.Fail()
	}
}

// Test that only targeted ocid resources are exportable
// issue-routing-tag: terraform/default
func TestUnitFindResources_restrictedOcids(t *testing.T) {
	initResourceDiscoveryTests()
	defer cleanupResourceDiscoveryTests()
	rootResource := getRootCompartmentResource()

	// Parent resources are defined as alwaysExportable. So even if it's not specified in the ocids, it should be exported.
	restrictedOcidTests := []map[string]interface{}{
		{
			"ocids":                map[string]bool{getTestResourceId("parent", 0): false, getTestResourceId("child", 0): false},
			"numExpectedResources": len(parentResources) + 1,
		},
		{
			"ocids":                map[string]bool{getTestResourceId("parent", 0): false, getTestResourceId("child", 3): false},
			"numExpectedResources": len(parentResources) + 1,
		},
		{
			"ocids":                map[string]bool{getTestResourceId("parent", 0): false, getTestResourceId("child", 0): false, "nonexistentID": false},
			"numExpectedResources": len(parentResources) + 1,
		},
		{
			"ocids":                map[string]bool{getTestResourceId("child", 0): false, getTestResourceId("child", 3): false, "nonexistentID": false},
			"numExpectedResources": len(parentResources) + 2,
		},
	}

	for idx, testCase := range restrictedOcidTests {
		t.Logf("running test #%d with following ocids: ", idx)
		restrictedOcids := testCase["ocids"].(map[string]bool)
		for ocid := range restrictedOcids {
			t.Logf(ocid)
		}

		ctx := &resourceDiscoveryContext{
			expectedResourceIds: restrictedOcids,
			errorList:           ErrorList{},
		}

		results, err := findResources(ctx, rootResource, compartmentTestingResourceGraph)
		if err != nil {
			t.Logf("got error from findResources: %v", err)
			t.Fail()
		}

		exportResourceCount := 0
		for _, resource := range results {
			if !resource.omitFromExport {
				exportResourceCount++
			}
		}
		if exportResourceCount != testCase["numExpectedResources"].(int) {
			t.Logf("expected %d resources to be exported, but got %d", testCase["numExpectedResources"].(int), len(results))
			t.Fail()
		}
	}
}

// Test that overriden find function is invoked if a resource has one
// issue-routing-tag: terraform/default
func TestUnitFindResources_overrideFn(t *testing.T) {
	initResourceDiscoveryTests()
	defer cleanupResourceDiscoveryTests()
	rootResource := getRootCompartmentResource()

	// Create an override function that returns nothing when discovering child test resources
	exportChildDefinition.findResourcesOverrideFn = func(*resourceDiscoveryContext, *TerraformResourceAssociation, *OCIResource, *TerraformResourceGraph) ([]*OCIResource, error) {
		return []*OCIResource{}, nil
	}
	defer func() { exportChildDefinition.findResourcesOverrideFn = nil }()

	ctx := &resourceDiscoveryContext{
		errorList: ErrorList{},
	}
	results, err := findResources(ctx, rootResource, compartmentTestingResourceGraph)
	if err != nil {
		t.Logf("got error from findResources: %v", err)
		t.Fail()
	}

	// Check that we got only parent resources, because child resources have been overridden to return nothing in their discovery function
	if len(results) != len(parentResources) {
		t.Logf("got %d results but expected %d results", len(results), len(parentResources))
		t.Fail()
	}

	for _, foundResource := range results {
		if foundResource.terraformClass == "oci_test_child" {
			t.Logf("oci_test_child resource was returned when not expected")
			t.Fail()
		}
	}
}

// Test that process resource function is invoked if a resource has one
// issue-routing-tag: terraform/default
func TestUnitFindResources_processResourceFn(t *testing.T) {
	initResourceDiscoveryTests()
	defer cleanupResourceDiscoveryTests()
	rootResource := getRootCompartmentResource()

	// Create a processing function that adds a new attribute to every discovered child resource
	exportChildDefinition.processDiscoveredResourcesFn = func(ctx *resourceDiscoveryContext, resources []*OCIResource) ([]*OCIResource, error) {
		for _, resource := range resources {
			resource.sourceAttributes["added_by_process_function"] = true
		}
		return resources, nil
	}
	defer func() { exportChildDefinition.processDiscoveredResourcesFn = nil }()

	ctx := &resourceDiscoveryContext{
		errorList: ErrorList{},
	}

	results, err := findResources(ctx, rootResource, compartmentTestingResourceGraph)
	if err != nil {
		t.Logf("got error from findResources: %v", err)
		t.Fail()
	}

	// Check that we got only parent resources, because child resources have been overridden to return nothing in their discovery function
	if len(results) != len(parentResources)+len(childrenResources) {
		t.Logf("got %d results but expected %d results", len(results), len(childrenResources)+len(parentResources))
		t.Fail()
	}

	for _, foundResource := range results {
		if foundResource.terraformClass == "oci_test_child" {
			if _, ok := foundResource.sourceAttributes["added_by_process_function"]; !ok {
				t.Logf("oci_test_child resource was returned when not expected")
				t.Fail()
			}
		}
	}
}

// Test that Terraform names can be generated from discovered resources
// issue-routing-tag: terraform/default
func TestUnitGenerateTerraformNameFromResource_basic(t *testing.T) {
	type testCase struct {
		resource     map[string]interface{}
		schema       map[string]*schema.Schema
		expectError  bool
		expectedName string
	}

	testResourceSchema := testParentResource().Schema
	testCases := []testCase{
		{
			resource:     map[string]interface{}{"display_name": "abc"},
			schema:       testResourceSchema,
			expectedName: "export_abc",
		},
		{
			// Repeating it should result in a different name than the previous test
			resource:     map[string]interface{}{"display_name": "abc"},
			schema:       testResourceSchema,
			expectedName: "export_abc_1",
		},
		{
			// Non-alphanumeric or non-hyphen/underscore characters in resource name should be removed
			resource:     map[string]interface{}{"display_name": "?!@#$%^ABC:&*()def-+ghi123_"},
			schema:       testResourceSchema,
			expectedName: "export_-ABC-def--ghi123_",
		},
		{
			// Resources without display_name attribute should result in error
			resource:    map[string]interface{}{},
			schema:      testResourceSchema,
			expectError: true,
		},
		{
			// Resources with display_name attribute should result in error, because it's not part of the schema
			resource:    map[string]interface{}{"display_name": "abc"},
			schema:      map[string]*schema.Schema{},
			expectError: true,
		},
	}

	for idx, test := range testCases {
		t.Logf("Running test case %d", idx)
		result, err := generateTerraformNameFromResource(test.resource, test.schema)
		if (err != nil) != test.expectError {
			t.Logf("expect error was '%v' but got err '%v'", test.expectError, err)
			t.Fail()
		}

		if result != test.expectedName {
			t.Logf("expect generated TF name to be %s but got %s", test.expectedName, result)
			t.Fail()
		}
	}
}

// Test that correct HCL is generated from a discovered test resource
// issue-routing-tag: terraform/default
func TestUnitGetHCLString_basic(t *testing.T) {
	initResourceDiscoveryTests()
	defer cleanupResourceDiscoveryTests()
	rootResource := getRootCompartmentResource()

	ctx := &resourceDiscoveryContext{
		errorList: ErrorList{},
	}
	results, err := findResources(ctx, rootResource, compartmentTestingResourceGraph)
	if err != nil {
		t.Logf("got error from findResources: %v", err)
		t.Fail()
	}

	targetResourceOcid := getTestResourceId("child", len(childrenResources)-1)
	testStringBuilder := &strings.Builder{}
	var targetResource *OCIResource
	for _, resource := range results {
		if resource.id == targetResourceOcid {
			targetResource = resource
			break
		}
	}

	if err := targetResource.getHCLString(testStringBuilder, nil); err != nil {
		t.Logf("got error '%v' when trying to get HCL string", err)
		t.Fail()
	}
	resultHcl := testStringBuilder.String()

	expectedHclResult := `resource oci_test_child export_string3_child_2 {
a_bool = "true"
a_float = "3.3"
a_int = "3"
a_list = [
"string0",
"string1",
"string2",
]
a_map = {
"key0" = "string0"
"key1" = "string1"
"key2" = "string2"
}
a_nested {
nested_bool = "true"
nested_float = "0"
nested_int = "0"
nested_string = "string0"
}
a_nested {
nested_bool = "true"
nested_float = "1.1"
nested_int = "1"
nested_string = "string1"
}
a_nested {
nested_bool = "true"
nested_float = "2.2"
nested_int = "2"
nested_string = "string2"
}
a_set = [
"string0",
"string2",
"string1",
]
a_string = "string3"
compartment_id = "string3"
parent_id = "ocid1.parent.abcdefghiklmnop.3"
}

`
	if expectedHclResult != resultHcl {
		t.Log("resulting Hcl does not match expected Hcl")
		t.Fail()
	}
}

// Test that HCL can be generated when optional or required fields are missing
// issue-routing-tag: terraform/default
func TestUnitGetHCLString_missingFields(t *testing.T) {
	initResourceDiscoveryTests()
	defer cleanupResourceDiscoveryTests()
	rootResource := getRootCompartmentResource()

	ctx := &resourceDiscoveryContext{
		errorList: ErrorList{},
	}
	results, err := findResources(ctx, rootResource, compartmentTestingResourceGraph)
	if err != nil {
		t.Logf("got error from findResources: %v", err)
		t.Fail()
	}

	targetResourceOcid := getTestResourceId("child", len(childrenResources)-1)
	testStringBuilder := &strings.Builder{}
	var targetResource *OCIResource
	for _, resource := range results {
		if resource.id == targetResourceOcid {
			targetResource = resource
			break
		}
	}

	delete(targetResource.sourceAttributes, "compartment_id")
	delete(targetResource.sourceAttributes, "a_string")
	targetResource.sourceAttributes["a_map"] = nil
	if err := targetResource.getHCLString(testStringBuilder, nil); err != nil {
		t.Logf("got error '%v' when trying to get HCL string", err)
		t.Fail()
	}
	resultHcl := testStringBuilder.String()

	if !strings.Contains(resultHcl, "compartment_id = \"<placeholder for missing required attribute>\"\t#Required") || !strings.Contains(resultHcl, "#a_string = <<Optional") {
		t.Logf("expected 'Required' compartment_id to have a placeholder value with comment and 'Optional' a_string field to be commented out, but they weren't")
		t.Fail()
	}

	if strings.Contains(resultHcl, "a_map") {
		t.Logf("a_map was set to nil but it still shows up in result Hcl")
		t.Fail()
	}
}

// Test that HCL can be generated with values replaced by interpolation syntax
// issue-routing-tag: terraform/default
func TestUnitGetHCLString_interpolationMap(t *testing.T) {
	initResourceDiscoveryTests()
	defer cleanupResourceDiscoveryTests()
	rootResource := getRootCompartmentResource()

	ctx := &resourceDiscoveryContext{
		errorList: ErrorList{},
	}
	results, err := findResources(ctx, rootResource, compartmentTestingResourceGraph)
	if err != nil {
		t.Logf("got error from findResources: %v", err)
		t.Fail()
	}

	targetResourceOcid := getTestResourceId("child", len(childrenResources)-1)
	testStringBuilder := &strings.Builder{}
	var targetResource *OCIResource
	for _, resource := range results {
		if resource.id == targetResourceOcid {
			targetResource = resource
			break
		}
	}

	// Test that ocids can be replaced with parent ID references
	interpolationMap := map[string]string{targetResource.parent.id: targetResource.parent.getHclReferenceIdString()}
	if err := targetResource.getHCLString(testStringBuilder, interpolationMap); err != nil {
		t.Logf("got error '%v' when trying to get HCL string", err)
		t.Fail()
	}
	resultHcl := testStringBuilder.String()

	if !strings.Contains(resultHcl, targetResource.parent.getHclReferenceIdString()) || strings.Contains(resultHcl, targetResource.parent.id) {
		t.Logf("expected hcl to replace parent ocid '%s' with '%s', but it wasn't", targetResource.parent.id, targetResource.parent.getHclReferenceIdString())
		t.Fail()
	}

	// Test that self-referencing IDs are ignored and do not show up in result hcl
	interpolationMap = map[string]string{targetResource.parent.id: targetResource.getHclReferenceIdString()}
	if err := targetResource.getHCLString(testStringBuilder, interpolationMap); err != nil {
		t.Logf("got error '%v' when trying to get HCL string", err)
		t.Fail()
	}
	resultHcl = testStringBuilder.String()

	if strings.Contains(resultHcl, targetResource.getHclReferenceIdString()) || !strings.Contains(resultHcl, targetResource.parent.id) {
		t.Logf("expected hcl to avoid cyclical reference '%s' but found one", targetResource.getHclReferenceIdString())
		t.Fail()
	}
}

// issue-routing-tag: terraform/default
func TestUnitGetHCLString_tfSyntaxVersion(t *testing.T) {
	initResourceDiscoveryTests()
	defer cleanupResourceDiscoveryTests()
	rootResource := getRootCompartmentResource()

	ctx := &resourceDiscoveryContext{}
	results, err := findResources(ctx, rootResource, compartmentTestingResourceGraph)
	if err != nil {
		t.Logf("got error from findResources: %v", err)
		t.Fail()
	}

	targetResourceOcid := getTestResourceId("child", len(childrenResources)-1)
	var targetResource *OCIResource
	for _, resource := range results {
		if resource.id == targetResourceOcid {
			targetResource = resource
			break
		}
	}

	// Test the syntax version generated
	tfHclVersion = &TfHclVersion11{}
	interpolationMap := map[string]string{targetResource.parent.id: targetResource.parent.getHclReferenceIdString()}
	r, _ := regexp.Compile("\\$\\{.*}")
	if !r.MatchString(interpolationMap[targetResource.parent.id]) {
		t.Logf("incorrect syntax generated for version %v", tfHclVersion.toString())
		t.Fail()
	}

	tfHclVersion = &TfHclVersion12{}
	interpolationMap = map[string]string{targetResource.parent.id: targetResource.parent.getHclReferenceIdString()}
	r, _ = regexp.Compile("[^${}]")
	if !r.MatchString(interpolationMap[targetResource.parent.id]) {
		t.Logf("incorrect syntax generated for version %v", tfHclVersion.toString())
		t.Fail()
	}

}

// issue-routing-tag: terraform/default
func TestUnitGetExportConfig(t *testing.T) {
	if os.Getenv("TF_HOME_OVERRIDE") == "" {
		t.Skip("This run requires you to set TF_HOME_OVERRIDE")
	}

	acctest.ProviderConfigTest(t, true, true, globalvar.AuthAPIKeySetting, "", getExportConfig)              // ApiKey with required fields + disable auto-retries
	acctest.ProviderConfigTest(t, false, true, globalvar.AuthAPIKeySetting, "", getExportConfig)             // ApiKey without required fields
	acctest.ProviderConfigTest(t, false, false, globalvar.AuthInstancePrincipalSetting, "", getExportConfig) // InstancePrincipal
	acctest.ProviderConfigTest(t, true, false, "invalid-auth-setting", "", getExportConfig)                  // Invalid auth + disable auto-retries
	configFile, keyFile, err := acctest.WriteConfigFile()
	assert.Nil(t, err)
	acctest.ProviderConfigTest(t, true, true, globalvar.AuthAPIKeySetting, "DEFAULT", getExportConfig)              // ApiKey with required fields + disable auto-retries
	acctest.ProviderConfigTest(t, false, true, globalvar.AuthAPIKeySetting, "DEFAULT", getExportConfig)             // ApiKey without required fields
	acctest.ProviderConfigTest(t, false, false, globalvar.AuthInstancePrincipalSetting, "DEFAULT", getExportConfig) // InstancePrincipal
	acctest.ProviderConfigTest(t, true, false, "invalid-auth-setting", "DEFAULT", getExportConfig)                  // Invalid auth + disable auto-retries
	acctest.ProviderConfigTest(t, false, false, globalvar.AuthAPIKeySetting, "PROFILE1", getExportConfig)           // correct profileName
	acctest.ProviderConfigTest(t, false, false, globalvar.AuthAPIKeySetting, "wrongProfile", getExportConfig)       // Invalid profileName
	//acctest.ProviderConfigTest(t, false, false, globalvar.AuthAPIKeySetting, "PROFILE2", getExportConfig)           // correct profileName with mix and match, disable for TC
	acctest.ProviderConfigTest(t, false, false, globalvar.AuthAPIKeySetting, "PROFILE3", getExportConfig) // correct profileName with mix and match & env
	defer utils.RemoveFile(configFile)
	defer utils.RemoveFile(keyFile)
	defer os.RemoveAll(path.Join(utils.GetHomeFolder(), globalvar.DefaultConfigDirName))
}

/*
This test is used to Create or destroy resources in a compartment using ORM stack
Parameter:
enable_create_destroy_rd_resources: true/false. Enable this run
stack_id: stack for the job
job_operation: APPLY/DESTROY. Job operation for the stack
DO NOT RUN THIS TEST LOCALLY AS IT WILL DESTROY INFRASTRUCTURE
*/
// issue-routing-tag: terraform/default
/*


func TestResourceDiscoveryApplyOrDestroyResourcesUsingStack(t *testing.T) {
	// env var check so as to prevent local run of this test.
	if reCreateResourceDiscoveryResources, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_create_destroy_rd_resources", "false")); !reCreateResourceDiscoveryResources {
		t.Skip("This run is used to apply/destroy resource for RD")
	}
	resourceManagerClient := acctest.GetTestClients(&schema.ResourceData{}).resourceManagerClient()
	stackId := GetEnvSettingWithBlankDefault("stack_id")
	if stackId == "" {
		t.Skip("Dependency stack_id not defined for test")
	}
	jobOperation := GetEnvSettingWithBlankDefault("job_operation")
	operation := oci_resourcemanager.JobOperationEnum(jobOperation)
	// Create resources using stack Create job
	isAutoApproved := true

	createJobRequest := oci_resourcemanager.CreateJobRequest{
		CreateJobDetails: oci_resourcemanager.CreateJobDetails{
			StackId:   &stackId,
			Operation: operation,
			ApplyJobPlanResolution: &oci_resourcemanager.ApplyJobPlanResolution{
				IsAutoApproved: &isAutoApproved,
			},
		},
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: GetRetryPolicy(false, "resourcemanager"),
		},
	}
	job_timeout_in_minutes, err := strconv.Atoi(GetEnvSettingWithDefault("job_timeout_in_minutes", "120"))
	assert.NoError(t, err)
	timeout := time.Duration(job_timeout_in_minutes) * time.Minute
	// Many resources require long time to Create/destroy
	createJobRequest.RequestMetadata.RetryPolicy.ShouldRetryOperation = ConditionShouldRetry(timeout, jobSuccessWaitCondition, "resourcemanager", false)

	createJobResponse, err := resourceManagerClient.CreateJob(context.Background(), createJobRequest)

	if err != nil {
		log.Fatalf("[ERROR] error in destroy job for stack: %v", err)
	}
	assert.NoError(t, err)

	retryPolicy := GetRetryPolicy(false, "resourcemanager")
	retryPolicy.ShouldRetryOperation = ConditionShouldRetry(time.Duration(15*time.Minute), jobSuccessWaitCondition, "resourcemanager", false)

	_, err = resourceManagerClient.GetJob(context.Background(), oci_resourcemanager.GetJobRequest{
		JobId: createJobResponse.Id,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	if err != nil {
		log.Fatalf("[WARN] wait for jobSuccessWaitCondition failed for %s resource with error %v", *createJobResponse.Id, err)
	} else {
		log.Printf("[INFO] end of jobSuccessWaitCondition for resource %s ", *createJobResponse.Id)
	}
	assert.NoError(t, err)
}

// issue-routing-tag: terraform/default
func TestResourceDiscoveryUpdateStack(t *testing.T) {
	stackId := GetEnvSettingWithBlankDefault("stack_id")
	if stackId == "" {
		t.Skip("Dependency stack_id not defined for test")
	}
	resourceType := GetEnvSettingWithBlankDefault("resource_type")
	if resourceType == "" {
		t.Skip("Dependency resource_type not defined for test")
	}
	resourceManagerClient := GetTestClients(&schema.ResourceData{}).resourceManagerClient()
	basePath := "../infrastructure/resource_discovery/" + resourceType + "/"
	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		log.Fatalf("[ERROR] unable to read files from resource manager example (%s): %v", basePath, err)
	}
	//Add files to zip
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".tf") {
			continue
		}
		data, err := ioutil.ReadFile(basePath + file.Name())
		if err != nil {
			log.Fatalf("[ERROR] read config file: %v", err)
		}

		f, err := zipWriter.Create(file.Name())
		if err != nil {
			log.Fatalf("[ERROR] cannot Create file for zip configuration: %v", err)
		}
		_, err = f.Write(data)
		if err != nil {
			log.Fatalf("[ERROR] cannot write tf configuration to zip archive: %v", err)
		}
	}
	// close zip writer
	err = zipWriter.Close()
	if err != nil {
		log.Fatalf("[ERROR] cannot close zip writer: %v", err)
	}

	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())
	terraformVersion := GetEnvSettingWithDefault("terraform_version", "0.11.x")

	updateStackRequest := oci_resourcemanager.UpdateStackRequest{
		StackId: &stackId,
		UpdateStackDetails: oci_resourcemanager.UpdateStackDetails{
			TerraformVersion: &terraformVersion,
			ConfigSource: oci_resourcemanager.UpdateZipUploadConfigSourceDetails{
				ZipFileBase64Encoded: &encoded,
			},
		},
	}
	_, err = resourceManagerClient.UpdateStack(context.Background(), updateStackRequest)

	assert.NoError(t, err)
}
func jobSuccessWaitCondition(response oci_common.OCIOperationResponse) bool {
	if jobResponse, ok := response.Response.(oci_resourcemanager.GetJobResponse); ok {
		return jobResponse.LifecycleState != oci_resourcemanager.JobLifecycleStateSucceeded
	}
	return false
}

// issue-routing-tag: terraform/default
func TestResourceDiscoveryOnCompartment(t *testing.T) {

	var exportCommandArgs ExportCommandArgs
	for serviceName, _ := range tenancyResourceGraphs {
		exportCommandArgs.Services = append(exportCommandArgs.Services, serviceName)
	}
	for serviceName, _ := range compartmentResourceGraphs {
		exportCommandArgs.Services = append(exportCommandArgs.Services, serviceName)
	}
	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	exportCommandArgs.GenerateState = true
	err := testExportCompartment(&compartmentId, &exportCommandArgs)
	assert.NoError(t, err)
}
*/
// issue-routing-tag: terraform/default
func TestExportCommandArgs_finalizeServices(t *testing.T) {
	compartmentResourceGraphs["compartment_testing"] = compartmentTestingResourceGraph
	compartmentResourceGraphs["compartment_testing_2"] = compartmentTestingResourceGraph
	tenancyResourceGraphs["tenancy_testing"] = tenancyTestingResourceGraph
	tenancyResourceGraphs["tenancy_testing_2"] = tenancyTestingResourceGraph

	defer func() {
		delete(compartmentResourceGraphs, "compartment_testing")
		delete(compartmentResourceGraphs, "compartment_testing_2")
		delete(compartmentResourceGraphs, "tenancy_testing")
		delete(compartmentResourceGraphs, "tenancy_testing_2")
	}()

	compartmentScopeServices = []string{"compartment_testing", "compartment_testing_2"}
	tenancyScopeServices = []string{"tenancy_testing", "tenancy_testing_2"}
	tenancyOcid := resourceDiscoveryTestTenancyOcid
	compartmentId := resourceDiscoveryTestCompartmentOcid

	type fields struct {
		FinalizedServices []string
	}
	type args struct {
		ctx *resourceDiscoveryContext
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "compartment_without_exclude",
			fields: fields{
				FinalizedServices: []string{"compartment_testing", "compartment_testing_2", "tenancy_testing"},
			},
			args: args{
				ctx: &resourceDiscoveryContext{
					tenancyOcid: tenancyOcid,
					ExportCommandArgs: &ExportCommandArgs{
						CompartmentId:   &compartmentId,
						Services:        []string{"compartment_testing", "compartment_testing_2", "tenancy_testing"},
						ExcludeServices: []string{},
					},
				},
			},
		},
		{
			name: "compartment_with_exclude",
			fields: fields{
				FinalizedServices: []string{"compartment_testing_2", "tenancy_testing"},
			},
			args: args{
				ctx: &resourceDiscoveryContext{
					tenancyOcid: tenancyOcid,
					ExportCommandArgs: &ExportCommandArgs{
						CompartmentId:   &compartmentId,
						Services:        []string{"compartment_testing", "compartment_testing_2", "tenancy_testing"},
						ExcludeServices: []string{"compartment_testing"},
					},
				},
			},
		},
		{
			name: "root_compartment_without_services_with_exclude",
			fields: fields{
				FinalizedServices: []string{"compartment_testing", "tenancy_testing", "tenancy_testing_2"},
			},
			args: args{
				ctx: &resourceDiscoveryContext{
					tenancyOcid: tenancyOcid,
					ExportCommandArgs: &ExportCommandArgs{
						CompartmentId:   &tenancyOcid,
						ExcludeServices: []string{"compartment_testing_2"},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.args.ctx.ExportCommandArgs.finalizeServices(tt.args.ctx)
			if !reflect.DeepEqual(tt.args.ctx.ExportCommandArgs.Services, tt.fields.FinalizedServices) {
				t.Logf("incorrect services list, expected: %v actual: %v", tt.fields.FinalizedServices, tt.args.ctx.ExportCommandArgs.Services)
				t.Fail()
			}
		})
	}
}

// issue-routing-tag: terraform/default
func TestUnitGetHCLString_logging(t *testing.T) {

	initResourceDiscoveryTests()
	defer cleanupResourceDiscoveryTests()
	rootResource := getRootCompartmentResource()
	outputDir, err := os.Getwd()
	outputDir = fmt.Sprintf("%s%sdiscoveryTest-%d", outputDir, string(os.PathSeparator), time.Now().Nanosecond())
	logFilePath := fmt.Sprintf("%s%sresource_discovery.log", outputDir, string(os.PathSeparator))

	if err = os.Mkdir(outputDir, os.ModePerm); err != nil {
		t.Logf("unable to mkdir %s. err: %v", outputDir, err)
		t.Fail()
	}
	os.Setenv("OCI_TF_LOG_PATH", logFilePath)
	l, _ := utils.NewTFProviderLogger()
	utils.SetTFProviderLogger(l)

	ctx := &resourceDiscoveryContext{
		errorList: ErrorList{},
	}
	_, err = findResources(ctx, rootResource, compartmentTestingResourceGraph)
	if err != nil {
		t.Logf("got error from findResources: %v", err)
		t.Fail()
	}
	if f, err := os.Stat(logFilePath); os.IsNotExist(err) || f.Size() == 0 {
		t.Logf("resource discovery logs not redirected to path specified by OCI_TF_LOG_PATH")
		t.Fail()
	}

	os.RemoveAll(outputDir)
}

// issue-routing-tag: terraform/default
func TestRunListExportableServicesCommand(t *testing.T) {

	outputDir, _ := os.Getwd()
	outputDir = fmt.Sprintf("%s%sdiscoveryTest-%d", outputDir, string(os.PathSeparator), time.Now().Nanosecond())

	if err := os.Mkdir(outputDir, os.ModePerm); err != nil {
		t.Logf("unable to mkdir %s. err: %v", outputDir, err)
		t.Fail()
	}
	servicesJsonPath := fmt.Sprintf("%s/services.json", outputDir)
	if err := RunListExportableServicesCommand(servicesJsonPath); err != nil {
		t.Errorf("RunListExportableServicesCommand() error = %v", err)
	}

	if f, err := os.Stat(servicesJsonPath); os.IsNotExist(err) || f.Size() == 0 {
		t.Logf("resource discovery services json not exported to path specified")
		t.Fail()
	}
	os.RemoveAll(outputDir)
}

// deleteInvalidReferences removes invalid reference from referenceMap if import fails for any resource
// issue-routing-tag: terraform/default
func Test_deleteInvalidReferences(t *testing.T) {
	discoveredResources := []*OCIResource{
		{
			compartmentId: resourceDiscoveryTestCompartmentOcid,
			TerraformResource: TerraformResource{
				id:             "ocid1.a.b.c",
				terraformClass: "oci_resource_type1",
				terraformName:  "type1_res1",
			},
		},
		{
			// resource with import failure
			compartmentId: resourceDiscoveryTestCompartmentOcid,
			TerraformResource: TerraformResource{
				id:             "ocid1.d.e.f",
				terraformClass: "oci_resource_type2",
				terraformName:  "type2_res1",
			},
			isErrorResource: true,
		},
		{
			compartmentId: resourceDiscoveryTestCompartmentOcid,
			TerraformResource: TerraformResource{
				id:             "ocid1.g.h.i",
				terraformClass: "oci_resource_type2",
				terraformName:  "type2_res2",
			},
		},
	}

	referenceMap := map[string]string{
		"ocid1.a.b.c": "oci_resource_type1.type1_res1",
		"ocid1.d.e.f": "oci_resource_type2.type2_res1", // failed resource
		"ocid1.g.h.i": "oci_resource_type2.type2_res2",
		"ocid1.j.k.l": "oci_resource_type2.type2_res1.attribute",  // reference to failed resource
		"ocid1.m.n.o": "oci_resource_type2.type2_res11.attribute", // similar name to failed resource
	}

	deleteInvalidReferences(referenceMap, discoveredResources)
	if _, ok := referenceMap["ocid1.d.e.f"]; ok {
		t.Logf("failed resource entry not removed from reference map")
		t.Fail()
	}

	if _, ok := referenceMap["ocid1.j.k.l"]; ok {
		t.Logf("reference to failed resource not removed from reference map")
		t.Fail()
	}
}

// issue-routing-tag: terraform/default
func Test_createTerraformStruct(t *testing.T) {

	_ = os.Unsetenv(globalvar.TerraformBinPathName)
	outputDir, err := os.Getwd()
	outputDir = fmt.Sprintf("%s%sdiscoveryTest-%d", outputDir, string(os.PathSeparator), time.Now().Nanosecond())
	if err = os.Mkdir(outputDir, os.ModePerm); err != nil {
		t.Logf("unable to mkdir %s. err: %v", outputDir, err)
		t.Fail()
	}
	args := &ExportCommandArgs{
		OutputDir: &outputDir,
	}
	tfHclVersion = &TfHclVersion12{}
	// verify executable from system path
	if _, _, err := createTerraformStruct(args); err != nil {
		t.Errorf("createTerraformStruct() error = %v", err)
		t.Fail()
	}

	// verify executable from env var
	// if invalid path is specified
	_ = os.Setenv(globalvar.TerraformBinPathName, "invalidPath")

	if _, _, err := createTerraformStruct(args); err == nil {
		t.Errorf("createTerraformStruct() expected error but succeeded")
		t.Fail()
	}

	// if path specified is a directory
	_ = os.Setenv(globalvar.TerraformBinPathName, "./")

	if _, _, err := createTerraformStruct(args); err == nil {
		t.Errorf("createTerraformStruct() expected error but succeeded")
		t.Fail()
	}

}
