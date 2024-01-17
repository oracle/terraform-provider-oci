// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resourcediscovery

import (
	"context"
	"errors"
	"fmt"
	"net/http"
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

	"github.com/hashicorp/go-version"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"

	"github.com/hashicorp/terraform-exec/tfinstall"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/hashicorp/terraform-exec/tfexec"
	oci_identity "github.com/oracle/oci-go-sdk/v65/identity"

	tf_client "github.com/oracle/terraform-provider-oci/internal/client"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/globalvar"
	tf_provider "github.com/oracle/terraform-provider-oci/internal/provider"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	resourceDiscoveryTestCompartmentOcid   = "ocid1.testcompartment.abc"
	resourceDiscoveryTestTenancyOcid       = "ocid1.testtenancy.xyz"
	resourceDiscoveryTestActiveLifecycle   = "ACTIVE"
	resourceDiscoveryTestInactiveLifecycle = "INACTIVE"
	resourceIdFor404ErrorResource          = "ocid1.child.abcdefghiklmnop.1"
)

var exportParentDefinition = &tf_export.TerraformResourceHints{
	ResourceClass:               "oci_test_parent",
	DatasourceClass:             "oci_test_parents",
	ResourceAbbreviation:        "parent",
	DatasourceItemsAttr:         "items",
	DiscoverableLifecycleStates: []string{resourceDiscoveryTestActiveLifecycle},
	AlwaysExportable:            true,
}

var exportChildDefinition = &tf_export.TerraformResourceHints{
	ResourceClass:               "oci_test_child",
	DatasourceClass:             "oci_test_children",
	ResourceAbbreviation:        "child",
	DatasourceItemsAttr:         "item_summaries",
	DiscoverableLifecycleStates: []string{resourceDiscoveryTestActiveLifecycle},
	RequireResourceRefresh:      true,
}

var exportChildDefinitionInactive = &tf_export.TerraformResourceHints{
	ResourceClass:               "oci_test_child_inactive",
	DatasourceClass:             "oci_test_children",
	ResourceAbbreviation:        "child",
	DatasourceItemsAttr:         "item_summaries",
	DiscoverableLifecycleStates: []string{resourceDiscoveryTestInactiveLifecycle},
}

var exportParentDefinitionWithFaultyDatasource = &tf_export.TerraformResourceHints{
	ResourceClass:               "oci_test_parent",
	DatasourceClass:             "oci_test_error_parents",
	ResourceAbbreviation:        "parent",
	DatasourceItemsAttr:         "items",
	DiscoverableLifecycleStates: []string{resourceDiscoveryTestActiveLifecycle},
	AlwaysExportable:            true,
}

var exportChildDefinitionWithFaultyDatasource = &tf_export.TerraformResourceHints{
	ResourceClass:               "oci_test_error_child",
	DatasourceClass:             "oci_test_children",
	ResourceAbbreviation:        "child",
	DatasourceItemsAttr:         "item_summaries",
	DiscoverableLifecycleStates: []string{resourceDiscoveryTestActiveLifecycle},
	RequireResourceRefresh:      true,
}

var exportResourceDefinitionWith404Error = &tf_export.TerraformResourceHints{
	ResourceClass:               "oci_test_404_error_child",
	DatasourceClass:             "oci_test_children",
	ResourceAbbreviation:        "child",
	DatasourceItemsAttr:         "item_summaries",
	DiscoverableLifecycleStates: []string{resourceDiscoveryTestActiveLifecycle},
	RequireResourceRefresh:      true,
}

var exportResourceDefinitionWithPanic = &tf_export.TerraformResourceHints{
	ResourceClass:               "oci_test_child",
	DatasourceClass:             "oci_test_panic_children",
	ResourceAbbreviation:        "child",
	DatasourceItemsAttr:         "item_summaries",
	DiscoverableLifecycleStates: []string{resourceDiscoveryTestActiveLifecycle},
}

var tenancyTestingResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_tenancy": {
		{
			TerraformResourceHints: exportParentDefinition,
		},
	},
	"oci_test_parent": {
		{
			TerraformResourceHints: exportChildDefinition,
			DatasourceQueryParams:  map[string]string{"parent_id": "id"},
		},
	},
}

var compartmentTestingResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{
			TerraformResourceHints: exportParentDefinition,
		},
	},
	"oci_test_parent": {
		{
			TerraformResourceHints: exportChildDefinition,
			DatasourceQueryParams:  map[string]string{"parent_id": "id"},
		},
	},
}

var compartmentTestingResourceInactiveLifeCycleGraphGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{
			TerraformResourceHints: exportParentDefinition,
		},
	},
	"oci_test_parent": {
		{
			TerraformResourceHints: exportChildDefinitionInactive,
			DatasourceQueryParams:  map[string]string{"parent_id": "id"},
		},
		{
			TerraformResourceHints: exportChildDefinition,
			DatasourceQueryParams:  map[string]string{"parent_id": "id"},
		},
	},
}

var compartmentTestingResourceGraphWithFaultyParentResource = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{
			TerraformResourceHints: exportParentDefinitionWithFaultyDatasource,
		},
	},
}

var compartmentTestingResourceGraphWithFaultyChildResource = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{
			TerraformResourceHints: exportParentDefinition,
		},
	},
	"oci_test_parent": {
		{
			TerraformResourceHints: exportChildDefinitionWithFaultyDatasource,
			DatasourceQueryParams:  map[string]string{"parent_id": "id"},
		},
	},
}

var compartmentTestingResourceGraphWith404ErrorResource = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{
			TerraformResourceHints: exportParentDefinition,
		},
	},
	"oci_test_parent": {
		{
			TerraformResourceHints: exportResourceDefinitionWith404Error,
			DatasourceQueryParams:  map[string]string{"parent_id": "id"},
		},
	},
}

var compartmentTestingResourceGraphWithPanicResource = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{
			TerraformResourceHints: exportParentDefinition,
		},
	},
	"oci_test_parent": {
		{
			TerraformResourceHints: exportResourceDefinitionWithPanic,
			DatasourceQueryParams:  map[string]string{"parent_id": "id"},
		},
		{
			TerraformResourceHints: exportChildDefinition,
			DatasourceQueryParams:  map[string]string{"parent_id": "id"},
		},
	},
}

var childrenResources map[string]map[string]interface{}
var parentResources map[string]map[string]interface{}

func getTestClients() *tf_client.OracleClients {
	return &tf_client.OracleClients{}
}
func getOutputDir() string {
	outputDir, err := os.Getwd()
	if err != nil {
		return "tmp/"
	}
	outputDir = fmt.Sprintf("%s%sunit-test", outputDir, string(os.PathSeparator))
	return outputDir
}
func createOutputDir() (string, error) {
	outputDir := getOutputDir()
	os.Mkdir(outputDir, os.ModePerm)
	return outputDir, nil
}
func getTestCtx() *tf_export.ResourceDiscoveryContext {
	clients := getTestClients()
	compartmentId := "dummy_compartment_id"
	outputDir, _ := createOutputDir()
	args := &tf_export.ExportCommandArgs{
		CompartmentId: &compartmentId,
		Services:      []string{"compartment_testing", "tenancy_testing"},
		OutputDir:     &outputDir,
		GenerateState: false,
		TFVersion:     &tf_export.TfHclVersionvar,
		Parallelism:   1,
	}

	ctx := &tf_export.ResourceDiscoveryContext{
		Clients:             clients,
		ExportCommandArgs:   args,
		TenancyOcid:         "tenancyOcid",
		DiscoveredResources: []*tf_export.OCIResource{},
		SummaryStatements:   []string{},
		ErrorList: tf_export.ErrorList{
			Errors: []*tf_export.ResourceDiscoveryError{},
		},
		TargetSpecificResources: false,
		ResourceHintsLookup:     createResourceHintsLookupMap(),
	}
	return ctx
}

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
				Set:      tfresource.LiteralTypeHashCodeForSets,
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
func getTypeSetResourceSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
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
	tf_export.ResourceNameCount = map[string]int{}
	tf_export.ResourcesMap = tf_provider.ResourcesMap()
	tf_export.DatasourcesMap = tf_provider.DataSourcesMap()
	tf_export.TfHclVersionvar = &tf_export.TfHclVersion12{}

	tf_export.ResourcesMap["oci_test_parent"] = testParentResource()
	tf_export.ResourcesMap["oci_test_child"] = testChildResource()
	tf_export.ResourcesMap["oci_test_error_child"] = testChildResourceWithError()
	tf_export.ResourcesMap["oci_test_404_error_child"] = testChildResourceWith404Error()

	tf_export.DatasourcesMap["oci_test_parents"] = testParentsDatasource()
	tf_export.DatasourcesMap["oci_test_children"] = testChildrenDatasource()
	tf_export.DatasourcesMap["oci_test_error_parents"] = testParentsDatasourceWithError()

	tf_export.DatasourcesMap["oci_test_panic_children"] = testParentsDatasourceWithPanic()

	tf_export.TenancyResourceGraphs["tenancy_testing"] = tenancyTestingResourceGraph
	tf_export.CompartmentResourceGraphs["compartment_testing"] = compartmentTestingResourceGraph

	initTestResources()
}

func cleanupResourceDiscoveryTests() {
	delete(tf_export.ResourcesMap, "oci_test_parent")
	delete(tf_export.ResourcesMap, "oci_test_child")
	delete(tf_export.ResourcesMap, "oci_test_error_child")
	delete(tf_export.DatasourcesMap, "oci_test_parents")
	delete(tf_export.DatasourcesMap, "oci_test_children")
	delete(tf_export.DatasourcesMap, "oci_test_error_children")
	delete(tf_export.DatasourcesMap, "oci_test_panic_children")
	delete(tf_export.TenancyResourceGraphs, "tenancy_testing")
	delete(tf_export.CompartmentResourceGraphs, "compartment_testing")
}

func initTestResources() {
	numParentResources := 4
	if parentResources == nil || len(parentResources) != numParentResources {
		parentResources = make(map[string]map[string]interface{}, numParentResources)
		for i := 0; i < numParentResources; i++ {
			parentResources[getTestResourceId("parent", i)] = generateTestResourceFromSchema(i, tf_export.ResourcesMap["oci_test_parent"].Schema)
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
				childResource := generateTestResourceFromSchema(i, tf_export.ResourcesMap["oci_test_child"].Schema)
				childResource["parent_id"] = parentId

				childrenResources[getTestResourceId("child", childCount)] = childResource
				childCount++
			}
		}
	}
}

func getRootCompartmentResource() *tf_export.OCIResource {
	return &tf_export.OCIResource{
		CompartmentId: resourceDiscoveryTestCompartmentOcid,
		TerraformResource: tf_export.TerraformResource{
			Id:             resourceDiscoveryTestCompartmentOcid,
			TerraformClass: "oci_identity_compartment",
			TerraformName:  "export",
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

	tfHclVersions := []tf_export.TfHclVersion{&tf_export.TfHclVersion11{}, &tf_export.TfHclVersion12{}}
	for _, tfVersion := range tfHclVersions {
		tf_export.TfHclVersionvar = tfVersion
		args := &tf_export.ExportCommandArgs{
			CompartmentId: &compartmentId,
			Services:      []string{"compartment_testing", "tenancy_testing"},
			OutputDir:     &outputDir,
			GenerateState: false,
			TFVersion:     &tf_export.TfHclVersionvar,
			Parallelism:   1,
		}
		getProviderEnvSettingWithDefaultVar = func(varName string, defaultValue string) string {
			return defaultValue
		}
		getEnvSettingWithBlankDefaultVar = func(varName string) string {
			return resourceDiscoveryTestTenancyOcid
		}
		getExportConfigVar = func(d *schema.ResourceData) (interface{}, error) {
			return getTestClients(), nil
		}
		exportConfigProvider = acctest.MockConfigurationProvider{}
		if err, _ = RunExportCommand(args); err != nil {
			t.Logf("(TF version %s) export command failed due to err: %v", tf_export.TfHclVersionvar.ToString(), err)
			t.Fail()
		}

		if _, err = os.Stat(fmt.Sprintf("%s%stenancy_testing.tf", outputDir, string(os.PathSeparator))); !os.IsNotExist(err) {
			t.Logf("(TF version %s) tenancy_testing.tf file generated even though it wasn't expected", tf_export.TfHclVersionvar.ToString())
			t.Fail()
		}

		if _, err = os.Stat(fmt.Sprintf("%s%scompartment_testing.tf", outputDir, string(os.PathSeparator))); os.IsNotExist(err) {
			t.Logf("(TF version %s) no compartment_testing.tf file generated", tf_export.TfHclVersionvar.ToString())
			t.Fail()
		}

		if _, err = os.Stat(fmt.Sprintf("%s%sterraform.tfstate", outputDir, string(os.PathSeparator))); !os.IsNotExist(err) {
			t.Logf("(TF version %s) found terraform.tfstate even though it wasn't expected", tf_export.TfHclVersionvar.ToString())
		}
	}

	os.RemoveAll(outputDir)
}

// Basic test to ensure that RunExportCommand generates TF artifacts when running with filter
// issue-routing-tag: terraform/default
func TestUnitRunExportCommandFilterResourceType_basic(t *testing.T) {
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

	tfHclVersions := []tf_export.TfHclVersion{&tf_export.TfHclVersion11{}, &tf_export.TfHclVersion12{}}
	for _, tfVersion := range tfHclVersions {
		tf_export.TfHclVersionvar = tfVersion
		args := &tf_export.ExportCommandArgs{
			CompartmentId: &compartmentId,
			Services:      []string{"compartment_testing", "tenancy_testing"},
			OutputDir:     &outputDir,
			GenerateState: false,
			TFVersion:     &tf_export.TfHclVersionvar,
			Parallelism:   1,
			Filters: []tf_export.ResourceFilter{&tf_export.ResourceTypeFilter{
				ResourceType:         map[string]bool{"oci_test_error_child": true},
				ResourceTypeOperator: tf_export.EXCLUDE,
			}},
		}
		getProviderEnvSettingWithDefaultVar = func(varName string, defaultValue string) string {
			return defaultValue
		}
		getEnvSettingWithBlankDefaultVar = func(varName string) string {
			return resourceDiscoveryTestTenancyOcid
		}
		getExportConfigVar = func(d *schema.ResourceData) (interface{}, error) {
			return getTestClients(), nil
		}
		exportConfigProvider = acctest.MockConfigurationProvider{}
		if err, _ = RunExportCommand(args); err != nil {
			t.Logf("(TF version %s) export command failed due to err: %v", tf_export.TfHclVersionvar.ToString(), err)
			t.Fail()
		}

		if _, err = os.Stat(fmt.Sprintf("%s%stenancy_testing.tf", outputDir, string(os.PathSeparator))); !os.IsNotExist(err) {
			t.Logf("(TF version %s) tenancy_testing.tf file generated even though it wasn't expected", tf_export.TfHclVersionvar.ToString())
			t.Fail()
		}

		if _, err = os.Stat(fmt.Sprintf("%s%scompartment_testing.tf", outputDir, string(os.PathSeparator))); os.IsNotExist(err) {
			t.Logf("(TF version %s) no compartment_testing.tf file generated", tf_export.TfHclVersionvar.ToString())
			t.Fail()
		}

		if _, err = os.Stat(fmt.Sprintf("%s%sterraform.tfstate", outputDir, string(os.PathSeparator))); !os.IsNotExist(err) {
			t.Logf("(TF version %s) found terraform.tfstate even though it wasn't expected", tf_export.TfHclVersionvar.ToString())
		}
	}

	os.RemoveAll(outputDir)
}

// issue-routing-tag: terraform/default
func TestUnitRunExportCommand_Parallel(t *testing.T) {
	initResourceDiscoveryTests()
	// add more services to compartment graphs
	tf_export.CompartmentResourceGraphs["compartment_testing_1"] = compartmentTestingResourceGraph
	tf_export.CompartmentResourceGraphs["compartment_testing_2"] = compartmentTestingResourceGraph
	tf_export.CompartmentResourceGraphs["compartment_testing_3"] = compartmentTestingResourceGraph
	tf_export.CompartmentResourceGraphs["compartment_testing_4"] = compartmentTestingResourceGraph
	tf_export.CompartmentResourceGraphs["compartment_testing_5"] = compartmentTestingResourceGraph
	tf_export.CompartmentResourceGraphs["compartment_testing_6"] = compartmentTestingResourceGraph
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

	tfHclVersions := []tf_export.TfHclVersion{&tf_export.TfHclVersion11{}, &tf_export.TfHclVersion12{}}
	for _, tfVersion := range tfHclVersions {
		tf_export.TfHclVersionvar = tfVersion
		args := &tf_export.ExportCommandArgs{
			CompartmentId: &compartmentId,
			Services:      []string{"compartment_testing", "compartment_testing_1", "compartment_testing_2", "compartment_testing_3", "compartment_testing_4", "compartment_testing_5", "compartment_testing_6", "tenancy_testing"},
			OutputDir:     &outputDir,
			GenerateState: false,
			TFVersion:     &tf_export.TfHclVersionvar,
			Parallelism:   4,
		}

		getProviderEnvSettingWithDefaultVar = func(varName string, defaultValue string) string {
			return defaultValue
		}
		getEnvSettingWithBlankDefaultVar = func(varName string) string {
			return resourceDiscoveryTestTenancyOcid
		}
		getExportConfigVar = func(d *schema.ResourceData) (interface{}, error) {
			return getTestClients(), nil
		}
		exportConfigProvider = acctest.MockConfigurationProvider{}
		if err, _ = RunExportCommand(args); err != nil {
			t.Logf("(TF version %s) export command failed due to err: %v", tf_export.TfHclVersionvar.ToString(), err)
			t.Fail()
		}

		if _, err = os.Stat(fmt.Sprintf("%s%stenancy_testing.tf", outputDir, string(os.PathSeparator))); !os.IsNotExist(err) {
			t.Logf("(TF version %s) tenancy_testing.tf file generated even though it wasn't expected", tf_export.TfHclVersionvar.ToString())
			t.Fail()
		}

		if _, err = os.Stat(fmt.Sprintf("%s%scompartment_testing.tf", outputDir, string(os.PathSeparator))); os.IsNotExist(err) {
			t.Logf("(TF version %s) no compartment_testing.tf file generated", tf_export.TfHclVersionvar.ToString())
			t.Fail()
		}

		if _, err = os.Stat(fmt.Sprintf("%s%sterraform.tfstate", outputDir, string(os.PathSeparator))); !os.IsNotExist(err) {
			t.Logf("(TF version %s) found terraform.tfstate even though it wasn't expected", tf_export.TfHclVersionvar.ToString())
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

	tfHclVersions := []tf_export.TfHclVersion{&tf_export.TfHclVersion11{}, &tf_export.TfHclVersion12{}}
	for _, tfVersion := range tfHclVersions {
		tf_export.TfHclVersionvar = tfVersion
		args := &tf_export.ExportCommandArgs{
			CompartmentId: &compartmentId,
			Services:      []string{"compartment_testing", "compartment_testing_1", "compartment_testing_2", "compartment_testing_3", "compartment_testing_4", "compartment_testing_5", "compartment_testing_6", "tenancy_testing"},
			OutputDir:     &outputDir,
			GenerateState: false,
			TFVersion:     &tf_export.TfHclVersionvar,
			Parallelism:   -1,
		}
		getProviderEnvSettingWithDefaultVar = func(varName string, defaultValue string) string {
			return defaultValue
		}
		getEnvSettingWithBlankDefaultVar = func(varName string) string {
			return resourceDiscoveryTestTenancyOcid
		}
		getExportConfigVar = func(d *schema.ResourceData) (interface{}, error) {
			return getTestClients(), nil
		}
		exportConfigProvider = acctest.MockConfigurationProvider{}
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
	tf_export.TfHclVersionvar = &tf_export.TfHclVersion12{}
	args := &tf_export.ExportCommandArgs{
		CompartmentId: &compartmentId,
		Services:      []string{"compartment_testing", "tenancy_testing"},
		OutputDir:     &nonexistentOutputDir,
		GenerateState: false,
		TFVersion:     &tf_export.TfHclVersionvar,
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

	tf_export.TfHclVersionvar = &tf_export.TfHclVersion12{}

	// nil args will cause panic and if panic is not handled test will fail else it will pass
	_, _ = RunExportCommand(nil)

	os.RemoveAll(outputDir)
}

// Test exit status in case of partial success
// issue-routing-tag: terraform/default
func TestUnitRunExportCommand_exitStatusForPartialSuccess(t *testing.T) {
	initResourceDiscoveryTests()
	// Replace commonexport.CompartmentResourceGraphs with the one having resource that has error in read
	// Status returned should be StatusPartialSuccess
	tf_export.CompartmentResourceGraphs["compartment_testing"] = compartmentTestingResourceGraphWithFaultyParentResource

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

	tf_export.TfHclVersionvar = &tf_export.TfHclVersion12{}
	args := &tf_export.ExportCommandArgs{
		CompartmentId: &compartmentId,
		Services:      []string{"compartment_testing", "core"},
		OutputDir:     &outputDir,
		GenerateState: false,
		TFVersion:     &tf_export.TfHclVersionvar,
		Parallelism:   1,
	}
	getProviderEnvSettingWithDefaultVar = func(varName string, defaultValue string) string {
		return defaultValue
	}
	getEnvSettingWithBlankDefaultVar = func(varName string) string {
		return resourceDiscoveryTestTenancyOcid
	}
	getExportConfigVar = func(d *schema.ResourceData) (interface{}, error) {
		return getTestClients(), nil
	}
	exportConfigProvider = acctest.MockConfigurationProvider{}
	err, status := RunExportCommand(args)
	if err != nil && status == StatusFail {
		t.Logf("(TF version %s) export command failed due to err: %v", tf_export.TfHclVersionvar.ToString(), err)
		t.Fail()
	}

	if err == nil && status == StatusPartialSuccess {
		if _, err = os.Stat(fmt.Sprintf("%s%score.tf", outputDir, string(os.PathSeparator))); os.IsNotExist(err) {
			t.Logf("(TF version %s) no core.tf file generated", tf_export.TfHclVersionvar.ToString())
			t.Fail()
		}
		if _, err = os.Stat(fmt.Sprintf("%s%scompartment_testing.tf", outputDir, string(os.PathSeparator))); !os.IsNotExist(err) {
			t.Logf("(TF version %s) no compartment_testing.tf file generated", tf_export.TfHclVersionvar.ToString())
			t.Fail()
		}

		if _, err = os.Stat(fmt.Sprintf("%s%sterraform.tfstate", outputDir, string(os.PathSeparator))); os.IsNotExist(err) {
			t.Logf("(TF version %s) found terraform.tfstate even though it wasn't expected", tf_export.TfHclVersionvar.ToString())
		}
	}

	os.RemoveAll(outputDir)
}

// Test exit status in case of partial success
// issue-routing-tag: terraform/default
func TestUnitRunExportCommand_errorSuggestionForPartialSuccess(t *testing.T) {
	type testFormat struct {
		name        string
		errorLength int
		ctx         *tf_export.ResourceDiscoveryContext
	}
	errors := []*tf_export.ResourceDiscoveryError{}
	parentResource := "oci_test_parent:ocid1.parent.abcdefghiklmnop.0"

	resourceDiscoveryContext := &tf_export.ResourceDiscoveryContext{
		ResourceHintsLookup: map[string]*tf_export.TerraformResourceHints{"oci_test_parent": exportParentDefinition},
		ExpectedResourceIds: map[string]bool{"oci_test_parent:ocid1.parent.abcdefghiklmnop.0": false},
		ErrorList: tf_export.ErrorList{
			Errors: append(errors, &tf_export.ResourceDiscoveryError{
				ResourceType:   "load_balancer",
				ParentResource: "export",
				Error:          fmt.Errorf("Error Message: [ERROR] Error while discovering below resources:\n" + parentResource),
				ResourceGraph:  nil,
			}),
		},
		TargetSpecificResources: false,
	}
	tests := []testFormat{
		{
			name:        "Test error message for partially not discovered resources",
			ctx:         resourceDiscoveryContext,
			errorLength: 1,
		},
	}
	for _, test := range tests {
		t.Logf("Running %s", test.name)
		if error, _ := getListOfNotDiscoveredResources(test.ctx); error != nil {
			expectedErrorMessage := error.Error()
			fmt.Println("\n", expectedErrorMessage)
			if !strings.Contains(expectedErrorMessage, parentResource) {
				t.Errorf("Output error - %s which is not equal to expected error - %s", expectedErrorMessage, test.ctx.ErrorList.Errors[0].Error.Error())
			}

		}
	}
}

// Test that resources can be found using a resource dependency graph
// issue-routing-tag: terraform/default
func TestUnitFindResources_basic(t *testing.T) {
	initResourceDiscoveryTests()
	defer cleanupResourceDiscoveryTests()
	rootResource := getRootCompartmentResource()

	ctx := &tf_export.ResourceDiscoveryContext{
		ErrorList: tf_export.ErrorList{},
	}
	results, err := findResources(ctx, rootResource, compartmentTestingResourceGraph, true)
	if err != nil {
		t.Logf("got error from findResources: %v", err)
		t.Fail()
	}

	if len(results) != len(childrenResources)+len(parentResources) {
		t.Logf("got %d results but expected %d results", len(results), len(childrenResources)+len(parentResources))
		t.Fail()
	}

	for _, foundResource := range results {
		if foundResource.TerraformClass == "oci_test_child" {
			if _, resourceRefreshAttributeExists := foundResource.SourceAttributes["a_nested"]; !resourceRefreshAttributeExists {
				t.Logf("child resource is missing an expected attribute that should have been filled by a resource refresh")
				t.Fail()
			}

			expectedTfNamePrefix := fmt.Sprintf("%s_%s", foundResource.Parent.TerraformName, exportChildDefinition.ResourceAbbreviation)
			if !strings.HasPrefix(foundResource.TerraformName, expectedTfNamePrefix) {
				t.Logf("child resource should have a name with prefix '%s' but name is '%s' instead", expectedTfNamePrefix, foundResource.TerraformName)
				t.Fail()
			}
		}
	}
}

func TestUnitFindResourcesInActiveLifeCycle_basic(t *testing.T) {
	initResourceDiscoveryTests()
	defer cleanupResourceDiscoveryTests()
	rootResource := getRootCompartmentResource()

	ctx := &tf_export.ResourceDiscoveryContext{
		ErrorList: tf_export.ErrorList{},
	}
	os.Setenv(globalvar.DiscoverAllStatesEnv, "1")

	results, err := findResources(ctx, rootResource, compartmentTestingResourceInactiveLifeCycleGraphGraph, true)
	if err != nil {
		t.Logf("got error from findResources: %v", err)
		t.Fail()
	}

	foundInactiveResource := false
	for _, foundResource := range results {
		if foundResource.TerraformClass == "oci_test_child_inactive" {
			foundInactiveResource = true
			break
		}
	}
	if foundInactiveResource == false {
		t.Logf("Inactive Resources not found")
		t.Fail()
	}
}

func TestUnitFindResourcesInActiveLifeCycleWithGlobalVariableNotSet_basic(t *testing.T) {
	initResourceDiscoveryTests()
	defer cleanupResourceDiscoveryTests()
	rootResource := getRootCompartmentResource()

	ctx := &tf_export.ResourceDiscoveryContext{
		ErrorList: tf_export.ErrorList{},
	}

	os.Unsetenv(globalvar.DiscoverAllStatesEnv)

	results, err := findResources(ctx, rootResource, compartmentTestingResourceInactiveLifeCycleGraphGraph, true)
	if err != nil {
		t.Logf("got error from findResources: %v", err)
		t.Fail()
	}

	for _, foundResource := range results {
		if foundResource.TerraformClass == "oci_test_child_inactive" {
			t.Logf("Inactive resource found even when export variable TF_DISCOVER_ALL_STATES is not set to 1: %v", err)
			t.Fail()
		}
	}
}

// Test that resources can be found using a resource dependency graph
// issue-routing-tag: terraform/default
func TestUnitFindResources_filter(t *testing.T) {
	initResourceDiscoveryTests()
	defer cleanupResourceDiscoveryTests()
	rootResource := getRootCompartmentResource()

	ctx := &tf_export.ResourceDiscoveryContext{
		ErrorList: tf_export.ErrorList{},
		ExportCommandArgs: &tf_export.ExportCommandArgs{
			Filters: []tf_export.ResourceFilter{&tf_export.ResourceTypeFilter{
				ResourceType:         map[string]bool{"oci_test_child": true},
				ResourceTypeOperator: tf_export.EXCLUDE,
			}},
		},
	}
	results, err := findResources(ctx, rootResource, compartmentTestingResourceGraph, false)
	if err != nil {
		t.Logf("got error from findResources: %v", err)
		t.Fail()
	}

	if len(results) == len(childrenResources)+len(parentResources) {
		t.Logf("got %d results but expected %d results", len(results), 4)
		t.Fail()
	}

	for _, foundResource := range results {
		if foundResource.TerraformClass == "oci_test_child" {
			t.Logf("resource of type oci_test_child should not be discovered as its set to be omitted by the filter")
			t.Fail()

		}
	}
}

// Test that resource with 404 Not found error do not show up in results
// issue-routing-tag: terraform/default
func TestUnitFindResources_404Error(t *testing.T) {
	initResourceDiscoveryTests()
	// Replace commonexport.CompartmentResourceGraphs with the one having resource that has 404 error in read
	// Resource with 404 error should be skipped

	defer cleanupResourceDiscoveryTests()
	rootResource := getRootCompartmentResource()

	ctx := &tf_export.ResourceDiscoveryContext{
		ErrorList: tf_export.ErrorList{},
	}
	results, err := findResources(ctx, rootResource, compartmentTestingResourceGraphWith404ErrorResource, true)
	if err != nil {
		t.Logf("got error from findResources: %v", err)
		t.Fail()
	}

	for _, resource := range results {
		if resource.SourceAttributes["id"] == "" {
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
	// Replace commonexport.CompartmentResourceGraphs with the one having resource that will panic

	defer cleanupResourceDiscoveryTests()
	rootResource := getRootCompartmentResource()

	ctx := &tf_export.ResourceDiscoveryContext{
		ErrorList: tf_export.ErrorList{},
	}
	results, err := findResources(ctx, rootResource, compartmentTestingResourceGraphWithPanicResource, true)
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

	ctx := &tf_export.ResourceDiscoveryContext{
		ErrorList: tf_export.ErrorList{},
	}
	_, err := findResources(ctx, rootResource, compartmentTestingResourceGraphWithFaultyChildResource, true)
	if err != nil {
		t.Logf("got error from findResources: %v", err)
		t.Fail()
	}
	if len(ctx.ErrorList.Errors) == 0 {
		t.Logf("expected errors for failed resources in ResourceDiscoveryContext errorList but found none")
		t.Fail()
	}
}

// Test that only targeted ocid resources are exportable
// issue-routing-tag: terraform/default
func TestUnitFindResources_restrictedOcids(t *testing.T) {
	initResourceDiscoveryTests()
	defer cleanupResourceDiscoveryTests()
	rootResource := getRootCompartmentResource()

	// Parent resources are defined as AlwaysExportable. So even if it's not specified in the ocids, it should be exported.
	restrictedOcidTests := []map[string]interface{}{
		{
			"ocids":                map[string]bool{getTestResourceId("parent", 0): false, getTestResourceId("child", 0): false},
			"numExpectedResources": len(parentResources) + 2,
		},
		{
			"ocids":                map[string]bool{getTestResourceId("parent", 0): false, getTestResourceId("child", 3): false},
			"numExpectedResources": len(parentResources) + 3,
		},
		{
			"ocids":                map[string]bool{getTestResourceId("parent", 0): false, getTestResourceId("child", 0): false, "nonexistentID": false},
			"numExpectedResources": len(parentResources) + 2,
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

		ctx := &tf_export.ResourceDiscoveryContext{
			ExpectedResourceIds: restrictedOcids,
			ErrorList:           tf_export.ErrorList{},
		}

		results, err := findResources(ctx, rootResource, compartmentTestingResourceGraph, true)
		if err != nil {
			t.Logf("got error from findResources: %v", err)
			t.Fail()
		}

		exportResourceCount := 0
		for _, resource := range results {
			if !resource.OmitFromExport {
				exportResourceCount++
			}
		}

		if exportResourceCount != testCase["numExpectedResources"].(int) {
			t.Logf("expected %d resources to be exported, but got %d", testCase["numExpectedResources"].(int), exportResourceCount)
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
	exportChildDefinition.FindResourcesOverrideFn = func(*tf_export.ResourceDiscoveryContext, *tf_export.TerraformResourceAssociation, *tf_export.OCIResource, *tf_export.TerraformResourceGraph) ([]*tf_export.OCIResource, error) {
		return []*tf_export.OCIResource{}, nil
	}
	defer func() { exportChildDefinition.FindResourcesOverrideFn = nil }()

	ctx := &tf_export.ResourceDiscoveryContext{
		ErrorList: tf_export.ErrorList{},
	}
	results, err := findResources(ctx, rootResource, compartmentTestingResourceGraph, true)
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
		if foundResource.TerraformClass == "oci_test_child" {
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
	exportChildDefinition.ProcessDiscoveredResourcesFn = func(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
		for _, resource := range resources {
			resource.SourceAttributes["added_by_process_function"] = true
		}
		return resources, nil
	}
	defer func() { exportChildDefinition.ProcessDiscoveredResourcesFn = nil }()

	ctx := &tf_export.ResourceDiscoveryContext{
		ErrorList: tf_export.ErrorList{},
	}

	results, err := findResources(ctx, rootResource, compartmentTestingResourceGraph, true)
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
		if foundResource.TerraformClass == "oci_test_child" {
			if _, ok := foundResource.SourceAttributes["added_by_process_function"]; !ok {
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
		result, err := tf_export.GenerateTerraformNameFromResource(test.resource, test.schema)
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

	ctx := &tf_export.ResourceDiscoveryContext{
		ErrorList: tf_export.ErrorList{},
	}
	results, err := findResources(ctx, rootResource, compartmentTestingResourceGraph, true)
	if err != nil {
		t.Logf("got error from findResources: %v", err)
		t.Fail()
	}

	targetResourceOcid := getTestResourceId("child", len(childrenResources)-1)
	testStringBuilder := &strings.Builder{}
	var targetResource *tf_export.OCIResource
	for _, resource := range results {
		if resource.Id == targetResourceOcid {
			targetResource = resource
			break
		}
	}

	if err := targetResource.GetHCLString(testStringBuilder, nil); err != nil {
		t.Logf("got error '%v' when trying to get HCL string", err)
		t.Fail()
	}
	resultHcl := testStringBuilder.String()

	expectedHclResult := `resource oci_test_child export_string3_child_1 {
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
func TestUnitGetHCLStringFromMap(t *testing.T) {
	initResourceDiscoveryTests()
	defer cleanupResourceDiscoveryTests()
	rootResource := getRootCompartmentResource()

	ctx := &tf_export.ResourceDiscoveryContext{}
	results, err := findResources(ctx, rootResource, compartmentTestingResourceGraph, true)
	if err != nil {
		t.Logf("got error from findResources: %v", err)
		t.Fail()
	}

	targetResourceOcid := getTestResourceId("child", len(childrenResources)-1)
	var targetResource *tf_export.OCIResource
	for _, resource := range results {
		if resource.Id == targetResourceOcid {
			targetResource = resource
			break
		}
	}

	// Test the syntax version generated
	tf_export.TfHclVersionvar = &tf_export.TfHclVersion11{}
	interpolationMap := map[string]string{targetResource.Parent.Id: targetResource.Parent.GetHclReferenceIdString()}
	err = tf_export.GetHCLStringFromMap(&strings.Builder{}, targetResource.Parent.SourceAttributes, testParentResource(), interpolationMap, rootResource, "")
	assert.NoError(t, err)
}

// Test that HCL can be generated when optional or required fields are missing
// issue-routing-tag: terraform/default
func TestUnitGetHCLString_missingFields(t *testing.T) {
	initResourceDiscoveryTests()
	defer cleanupResourceDiscoveryTests()
	rootResource := getRootCompartmentResource()

	ctx := &tf_export.ResourceDiscoveryContext{
		ErrorList: tf_export.ErrorList{},
	}
	results, err := findResources(ctx, rootResource, compartmentTestingResourceGraph, true)
	if err != nil {
		t.Logf("got error from findResources: %v", err)
		t.Fail()
	}

	targetResourceOcid := getTestResourceId("child", len(childrenResources)-1)
	testStringBuilder := &strings.Builder{}
	var targetResource *tf_export.OCIResource
	for _, resource := range results {
		if resource.Id == targetResourceOcid {
			targetResource = resource
			break
		}
	}

	delete(targetResource.SourceAttributes, "compartment_id")
	delete(targetResource.SourceAttributes, "a_string")
	targetResource.SourceAttributes["a_map"] = nil
	if err := targetResource.GetHCLString(testStringBuilder, nil); err != nil {
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

	ctx := &tf_export.ResourceDiscoveryContext{
		ErrorList: tf_export.ErrorList{},
	}
	results, err := findResources(ctx, rootResource, compartmentTestingResourceGraph, true)
	if err != nil {
		t.Logf("got error from findResources: %v", err)
		t.Fail()
	}

	targetResourceOcid := getTestResourceId("child", len(childrenResources)-1)
	testStringBuilder := &strings.Builder{}
	var targetResource *tf_export.OCIResource
	for _, resource := range results {
		if resource.Id == targetResourceOcid {
			targetResource = resource
			break
		}
	}

	// Test that ocids can be replaced with parent ID references
	interpolationMap := map[string]string{targetResource.Parent.Id: targetResource.Parent.GetHclReferenceIdString()}
	if err := targetResource.GetHCLString(testStringBuilder, interpolationMap); err != nil {
		t.Logf("got error '%v' when trying to get HCL string", err)
		t.Fail()
	}
	resultHcl := testStringBuilder.String()

	if !strings.Contains(resultHcl, targetResource.Parent.GetHclReferenceIdString()) || strings.Contains(resultHcl, targetResource.Parent.Id) {
		t.Logf("expected hcl to replace parent ocid '%s' with '%s', but it wasn't", targetResource.Parent.Id, targetResource.Parent.GetHclReferenceIdString())
		t.Fail()
	}

	// Test that self-referencing IDs are ignored and do not show up in result hcl
	interpolationMap = map[string]string{targetResource.Parent.Id: targetResource.GetHclReferenceIdString()}
	if err := targetResource.GetHCLString(testStringBuilder, interpolationMap); err != nil {
		t.Logf("got error '%v' when trying to get HCL string", err)
		t.Fail()
	}
	resultHcl = testStringBuilder.String()

	if strings.Contains(resultHcl, targetResource.GetHclReferenceIdString()) || !strings.Contains(resultHcl, targetResource.Parent.Id) {
		t.Logf("expected hcl to avoid cyclical reference '%s' but found one", targetResource.GetHclReferenceIdString())
		t.Fail()
	}
}

// issue-routing-tag: terraform/default
func TestUnitGetHCLString_tfSyntaxVersion(t *testing.T) {
	initResourceDiscoveryTests()
	defer cleanupResourceDiscoveryTests()
	rootResource := getRootCompartmentResource()

	ctx := &tf_export.ResourceDiscoveryContext{}
	results, err := findResources(ctx, rootResource, compartmentTestingResourceGraph, true)
	if err != nil {
		t.Logf("got error from findResources: %v", err)
		t.Fail()
	}

	targetResourceOcid := getTestResourceId("child", len(childrenResources)-1)
	var targetResource *tf_export.OCIResource
	for _, resource := range results {
		if resource.Id == targetResourceOcid {
			targetResource = resource
			break
		}
	}

	// Test the syntax version generated
	tf_export.TfHclVersionvar = &tf_export.TfHclVersion11{}
	interpolationMap := map[string]string{targetResource.Parent.Id: targetResource.Parent.GetHclReferenceIdString()}
	r, _ := regexp.Compile("\\$\\{.*}")
	if !r.MatchString(interpolationMap[targetResource.Parent.Id]) {
		t.Logf("incorrect syntax generated for version %v", tf_export.TfHclVersionvar.ToString())
		t.Fail()
	}

	tf_export.TfHclVersionvar = &tf_export.TfHclVersion12{}
	interpolationMap = map[string]string{targetResource.Parent.Id: targetResource.Parent.GetHclReferenceIdString()}
	r, _ = regexp.Compile("[^${}]")
	if !r.MatchString(interpolationMap[targetResource.Parent.Id]) {
		t.Logf("incorrect syntax generated for version %v", tf_export.TfHclVersionvar.ToString())
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
func TestUnitGetExportConfigWithFakeProviderClient(t *testing.T) {

	tfProviderGetSdkConfigProvider = func(d *schema.ResourceData, clients *tf_client.OracleClients) (oci_common.ConfigurationProvider, error) {
		return acctest.MockConfigurationProvider{}, nil
	}
	sdkConfigProviderTenancyOCIDVar = func(sdkConfigProvider oci_common.ConfigurationProvider) (string, error) {
		return "fake ocid", nil
	}
	tfProviderBuildConfigureClientFn = func(configProvider oci_common.ConfigurationProvider, httpClient *http.Client) (tf_client.ConfigureClient, error) {
		fakeConfigureClientFn := func(client *oci_common.BaseClient) error {
			return nil
		}
		return fakeConfigureClientFn, nil
	}
	tests := []struct {
		name      string
		mock      func()
		wantError bool
	}{
		{
			name: "with sucessful client",
			mock: func() {
				createSDKClientsVar = func(clients *tf_client.OracleClients, configProvider oci_common.ConfigurationProvider, configureClient tf_client.ConfigureClient) (err error) {
					return nil
				}
			},
			wantError: false,
		},
		{
			name: " client with error",
			mock: func() {
				createSDKClientsVar = func(clients *tf_client.OracleClients, configProvider oci_common.ConfigurationProvider, configureClient tf_client.ConfigureClient) (err error) {
					return fmt.Errorf("expected error from createClient")
				}
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			_, err := getExportConfig(nil)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err, "")
			}
		})
	}
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
   	for serviceName, _ := range commonexport.CompartmentResourceGraphs {
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
	tf_export.CompartmentResourceGraphs["compartment_testing"] = compartmentTestingResourceGraph
	tf_export.CompartmentResourceGraphs["compartment_testing_2"] = compartmentTestingResourceGraph
	tf_export.TenancyResourceGraphs["tenancy_testing"] = tenancyTestingResourceGraph
	tf_export.TenancyResourceGraphs["tenancy_testing_2"] = tenancyTestingResourceGraph

	defer func() {
		delete(tf_export.CompartmentResourceGraphs, "compartment_testing")
		delete(tf_export.CompartmentResourceGraphs, "compartment_testing_2")
		delete(tf_export.CompartmentResourceGraphs, "tenancy_testing")
		delete(tf_export.CompartmentResourceGraphs, "tenancy_testing_2")
	}()

	tf_export.CompartmentScopeServices = []string{"compartment_testing", "compartment_testing_2"}
	tf_export.TenancyScopeServices = []string{"tenancy_testing", "tenancy_testing_2"}
	tenancyOcid := resourceDiscoveryTestTenancyOcid
	compartmentId := resourceDiscoveryTestCompartmentOcid
	type fields struct {
		FinalizedServices []string
	}
	type args struct {
		ctx *tf_export.ResourceDiscoveryContext
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
				ctx: &tf_export.ResourceDiscoveryContext{
					TenancyOcid: tenancyOcid,
					ExportCommandArgs: &tf_export.ExportCommandArgs{
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
				ctx: &tf_export.ResourceDiscoveryContext{
					TenancyOcid: tenancyOcid,
					ExportCommandArgs: &tf_export.ExportCommandArgs{
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
				ctx: &tf_export.ResourceDiscoveryContext{
					TenancyOcid: tenancyOcid,
					ExportCommandArgs: &tf_export.ExportCommandArgs{
						CompartmentId:   &tenancyOcid,
						ExcludeServices: []string{"compartment_testing_2"},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.args.ctx.ExportCommandArgs.FinalizeServices(tt.args.ctx)
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

	ctx := &tf_export.ResourceDiscoveryContext{
		ErrorList: tf_export.ErrorList{},
	}
	_, err = findResources(ctx, rootResource, compartmentTestingResourceGraph, true)
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
func TestUnitRunListExportableServicesCommand(t *testing.T) {

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
func TestUnit_deleteInvalidReferences(t *testing.T) {
	discoveredResources := []*tf_export.OCIResource{
		{
			CompartmentId: resourceDiscoveryTestCompartmentOcid,
			TerraformResource: tf_export.TerraformResource{
				Id:             "ocid1.a.b.c",
				TerraformClass: "oci_resource_type1",
				TerraformName:  "type1_res1",
			},
		},
		{
			// resource with import failure
			CompartmentId: resourceDiscoveryTestCompartmentOcid,
			TerraformResource: tf_export.TerraformResource{
				Id:             "ocid1.d.e.f",
				TerraformClass: "oci_resource_type2",
				TerraformName:  "type2_res1",
			},
			IsErrorResource: true,
		},
		{
			CompartmentId: resourceDiscoveryTestCompartmentOcid,
			TerraformResource: tf_export.TerraformResource{
				Id:             "ocid1.g.h.i",
				TerraformClass: "oci_resource_type2",
				TerraformName:  "type2_res2",
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
	args := &tf_export.ExportCommandArgs{
		OutputDir: &outputDir,
	}
	tf_export.TfHclVersionvar = &tf_export.TfHclVersion12{}
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
	defer os.RemoveAll(outputDir)
}
func TestUnitCreateTerraformStruct(t *testing.T) {
	outputDir, err := os.Getwd()
	osStatvar = func(name string) (os.FileInfo, error) {
		return nil, nil
	}
	isDirVar = func(file os.FileInfo) bool {
		return false
	}
	outputDir = fmt.Sprintf("%s%sdiscoveryTest-%d", outputDir, string(os.PathSeparator), time.Now().Nanosecond())
	if err = os.Mkdir(outputDir, os.ModePerm); err != nil {
		t.Logf("unable to mkdir %s. err: %v", outputDir, err)
		t.Fail()
	}
	tests := []struct {
		name      string
		mock      func()
		wantError bool
		errorMsg  string
	}{
		{
			name: "without terraform bin path",
			mock: func() {
				getEnvSettingWithBlankDefaultVar = func(varName string) string {
					return ""
				}
				tfInstallFindVar = func(ctx context.Context, opts ...tfinstall.ExecPathFinder) (string, error) {
					return "", fmt.Errorf("error")
				}
			},
			wantError: true,
			errorMsg:  "error  expected for unset variable",
		},
		{
			name: "invalid path",
			mock: func() {
				getEnvSettingWithBlankDefaultVar = func(varName string) string {
					return "invalidPath"
				}
			},
			wantError: true,
			errorMsg:  "error  expected for invalid path",
		},
		{
			name: "bin as directory",
			mock: func() {
				getEnvSettingWithBlankDefaultVar = func(varName string) string {
					return "./"
				}
			},
			wantError: true,
			errorMsg:  "error  expected for directory",
		},
		{
			name: "with valid terraform bin path with nil version",
			mock: func() {
				getEnvSettingWithBlankDefaultVar = func(varName string) string {
					return ""
				}
				tfVersionVar = func(tf *tfexec.Terraform, backgroundCtx context.Context) (*version.Version, map[string]*version.Version, error) {
					return nil, nil, fmt.Errorf("mock")
				}
			},
			wantError: true,
			errorMsg:  "error  expected for nil version",
		},
		{
			name: "with valid terraform bin path with  version 1.2.2",
			mock: func() {
				getEnvSettingWithBlankDefaultVar = func(varName string) string {
					return "terraform_valid_bin"
				}
				tfVersionVar = func(tf *tfexec.Terraform, backgroundCtx context.Context) (*version.Version, map[string]*version.Version, error) {
					dummyVersion, _ := version.NewVersion("1.2.2")
					return dummyVersion, nil, nil
				}
				osStatvar = func(name string) (os.FileInfo, error) {
					return nil, nil
				}
				isDirVar = func(file os.FileInfo) bool {
					return false
				}
			},
			wantError: false,
			errorMsg:  "error  not expected for 1.2.2 version",
		},
	}
	args := &tf_export.ExportCommandArgs{
		OutputDir: &outputDir,
	}
	tf_export.TfHclVersionvar = &tf_export.TfHclVersion12{}

	for _, tt := range tests {
		t.Run(tt.name, func(test *testing.T) {
			tt.mock()
			_, _, err := createTerraformStruct(args)
			if tt.wantError {
				assert.Error(test, err, tt.errorMsg)
			} else {
				assert.NoError(test, err, tt.errorMsg)
			}
		})
	}
	defer os.RemoveAll(outputDir)
}

func TestUnitPrintResourceGraphResources(t *testing.T) {
	resourceGraphs := map[string]tf_export.TerraformResourceGraph{
		"tenancyTestingResourceGraph": tenancyTestingResourceGraph,
	}
	err := printResourceGraphResources(resourceGraphs, "testing")
	assert.NoError(t, err, "error not expected for tenancyTestingResourceGraph")
}

func TestUnitRunListExportableResourcesCommand(t *testing.T) {

	err := RunListExportableResourcesCommand()
	assert.NoError(t, err, "error not expected")
}

func TestUnitGenerateStateParallel(t *testing.T) {

	//err = readEnvironmentVars(d)

	ctx := getTestCtx()
	steps := make([]resourceDiscoveryStep, 1)
	steps[0] = &resourceDiscoveryWithTargetIds{
		resourceDiscoveryBaseStep: resourceDiscoveryBaseStep{
			ctx:                 ctx,
			name:                "resources",
			discoveredResources: []*tf_export.OCIResource{},
			omittedResources:    []*tf_export.OCIResource{},
		},
	}
	t.Logf(fmt.Sprintf("%v", steps))
	err := generateStateParallel(ctx, steps)
	assert.NoError(t, err, "")
}

func TestUnitGenerateStateParallelWhenTfInitFails(t *testing.T) {

	ctx := getTestCtx()
	nSteps := 80 // number of steps
	steps := make([]resourceDiscoveryStep, nSteps)

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

	for i := 0; i < nSteps; i++ {
		discoveredResources := []*tf_export.OCIResource{}
		discoveredResources = append(discoveredResources, &tf_export.OCIResource{
			CompartmentId: compartmentId,
			TerraformResource: tf_export.TerraformResource{
				Id:             "ocid1.a.b.c",
				TerraformClass: "oci_resource_type1",
				TerraformName:  "type1_res1",
			},
			Parent: &tf_export.OCIResource{
				TerraformResource: tf_export.TerraformResource{TerraformName: "tf"},
			},
		})

		steps[i] = &resourceDiscoveryWithTargetIds{
			resourceDiscoveryBaseStep: resourceDiscoveryBaseStep{
				ctx:                 ctx,
				name:                "resources" + fmt.Sprint(i),
				discoveredResources: discoveredResources,
				omittedResources:    []*tf_export.OCIResource{},
			},
		}
	}

	type args struct {
		steps []resourceDiscoveryStep
		ctx   *tf_export.ResourceDiscoveryContext
	}
	t_args := args{
		steps: steps,
		ctx:   ctx,
	}
	tests := []struct {
		name      string
		args      args
		mock      func()
		wantError bool
	}{
		{
			name: "If Import failed ,should Return error",
			args: t_args,
			mock: func() {
				t_args.ctx.TerraformProviderBinaryPath = "tf"
				t_args.ctx.OutputDir = &outputDir
				ctxTerraformImportVar = func(ctx *tf_export.ResourceDiscoveryContext, ctxBackground context.Context, address, id string, importArgs ...tfexec.ImportOption) error {
					return nil
				}
				terraformInitMockVar = func(r *resourceDiscoveryBaseStep, backgroundCtx context.Context, initArgs []tfexec.InitOption) error {
					return errors.New("Init failed")
				}
				sem = make(chan struct{}, 4) // Parallelism=4
				tf_export.ResourcesMap = mockResourcesMap()
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := generateStateParallel(ctx, steps)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err, "")
			}
		})
	}
}

func TestUnitGenerateState(test *testing.T) {
	defer func() {
		outputDir := getOutputDir()
		os.RemoveAll(outputDir)
	}()
	ctx := getTestCtx()
	steps := make([]resourceDiscoveryStep, 1)
	steps[0] = &resourceDiscoveryWithTargetIds{
		resourceDiscoveryBaseStep: resourceDiscoveryBaseStep{
			ctx:                 ctx,
			name:                "resources",
			discoveredResources: []*tf_export.OCIResource{},
			omittedResources:    []*tf_export.OCIResource{},
		},
	}
	type args struct {
		steps []resourceDiscoveryStep
		ctx   *tf_export.ResourceDiscoveryContext
	}
	t_args := args{
		steps: steps,
		ctx:   ctx,
	}
	tests := []struct {
		name      string
		args      args
		mock      func()
		wantError bool
	}{{
		name: "Run with Default Steps",
		args: t_args,
		mock: func() {
			terraformInitVar = func(ctx *tf_export.ResourceDiscoveryContext, backgroundCtx context.Context, initArgs []tfexec.InitOption) error {
				return nil
			}
			//err := generateState(ctx, steps)
		},
		wantError: false,
	},
		{
			name: "Run with terraformProviderBinaryPath value int context",
			args: t_args,
			mock: func() {
				t_args.ctx.TerraformProviderBinaryPath = "tf"
				terraformInitVar = func(ctx *tf_export.ResourceDiscoveryContext, backgroundCtx context.Context, initArgs []tfexec.InitOption) error {
					return nil
				}
				//err := generateState(ctx, steps)
			},
			wantError: false,
		},
		{
			name: "If Init failed ,should Return error",
			args: t_args,
			mock: func() {
				t_args.ctx.TerraformProviderBinaryPath = "tf"
				terraformInitVar = func(ctx *tf_export.ResourceDiscoveryContext, backgroundCtx context.Context, initArgs []tfexec.InitOption) error {
					return errors.New("init failed")
				}
				//err := generateState(ctx, steps)
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		test.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := generateState(ctx, steps)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err, "")
			}
		})
	}
}

func TestUnitGetOciResource(t *testing.T) {
	childResource := testChildResource()
	d := childResource.TestResourceData()
	resource, err := tf_export.GetOciResource(d, childResource.Schema, "dummy", exportChildDefinition, "dummy resource id")
	assert.NoError(t, err)
	assert.NotNil(t, resource, "should return a dummy resource")
	assert.Equal(t, "dummy", resource.CompartmentId)
	assert.Equal(t, "dummy resource id", resource.Id)
}
func TestUnitConvertDatasourceItemToMap(t *testing.T) {
	parentResource := testParentResource()
	childResource := getTypeSetResourceSchema()
	parentResource.Schema["a_nested_set"] = childResource
	d := parentResource.TestResourceData()
	dummyData := []interface{}{}
	dummyData = append(dummyData, map[string]string{
		"nested_string": "hello",
	})
	d.Set("a_nested_set", schema.NewSet(func(interface{}) int {
		return 1
	}, dummyData))
	result, err := tf_export.ConvertDatasourceItemToMap(d, "", parentResource.Schema)
	assert.NoError(t, err)
	assert.NotNil(t, result, "should return a map")
}
func TestUnitConvertResourceDataToMap(t *testing.T) {
	parentResource := testParentResource()
	d := parentResource.TestResourceData()
	result := tf_export.ConvertResourceDataToMap(parentResource.Schema, d)
	assert.NotNil(t, result, "should return a map")
}
func TestUnitResolveCompartmentId(t *testing.T) {
	client := getTestClients()
	compartmentName := "dummy_commpartment"
	id := "1"
	exportConfigProvider = acctest.MockConfigurationProvider{}
	identityClientListCompartmentsVar = func(clients *tf_client.OracleClients, req oci_identity.ListCompartmentsRequest) (oci_identity.ListCompartmentsResponse, error) {
		opcRequestId := "dummy_opc"
		return oci_identity.ListCompartmentsResponse{
			RawResponse:  &http.Response{},
			OpcRequestId: &opcRequestId,
			Items: []oci_identity.Compartment{{
				Id:   &id,
				Name: &compartmentName,
			},
			},
		}, nil
	}
	compartmentId, err := resolveCompartmentId(client, &compartmentName)
	assert.NoError(t, err)
	assert.NotNil(t, compartmentId)
	assert.Equal(t, id, *compartmentId)
}

func TestUnitGetTenancyOcidFromCompartment(t *testing.T) {
	compartmentId := "dummy_compartment_id"
	identityClientGetCompartmentVar = func(clients *tf_client.OracleClients, getCompartmentRequest oci_identity.GetCompartmentRequest) (oci_identity.GetCompartmentResponse, error) {
		return oci_identity.GetCompartmentResponse{
			Compartment: oci_identity.Compartment{
				//CompartmentId: &compartmentId,
				Id: &compartmentId,
			},
		}, nil
	}
	resp, err := getTenancyOcidFromCompartment(nil, compartmentId)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, compartmentId, resp)
}

func TestUnitImportResource(t *testing.T) {
	compartmentId := "dummy_compartment_id"
	outputDir := "tmp/"

	resource := &tf_export.OCIResource{
		CompartmentId: compartmentId,
		TerraformResource: tf_export.TerraformResource{
			Id:             "ocid1.a.b.c",
			TerraformClass: "oci_resource_type1",
			TerraformName:  "type1_res1",
		},
		Parent: &tf_export.OCIResource{
			TerraformResource: tf_export.TerraformResource{TerraformName: "tf"},
		},
	}
	ctx := getTestCtx()
	tfexecConfigVar = func(outputDir string) *tfexec.ConfigOption {
		return nil
	}
	tfexecStateVar = func(tmpStateOutputFile string) *tfexec.StateOption {
		return nil
	}

	//with Unknown Resource
	tf_export.ResourcesMap = make(map[string]*schema.Resource)
	importResource(ctx, resource, outputDir)

	//without importer
	tf_export.ResourcesMap = mockResourcesMap()
	tf_export.ResourcesMap["oci_resource_type1"].Importer = nil
	importResource(ctx, resource, outputDir)

	// Without Import Error
	ctxTerraformImportVar = func(ctx *tf_export.ResourceDiscoveryContext, ctxBackground context.Context, address, id string, importArgs ...tfexec.ImportOption) error {
		return nil
	}
	tf_export.ResourcesMap = mockResourcesMap()
	importResource(ctx, resource, outputDir)
	assert.Equal(t, 0, len(ctx.ErrorList.Errors))

	//With Import Error
	ctxTerraformImportVar = func(ctx *tf_export.ResourceDiscoveryContext, ctxBackground context.Context, address, id string, importArgs ...tfexec.ImportOption) error {
		return errors.New("dummy error to cover code")
	}
	tf_export.ResourcesMap = mockResourcesMap()
	importResource(ctx, resource, outputDir)
	assert.NotNil(t, ctx.ErrorList)
	assert.Equal(t, 1, len(ctx.ErrorList.Errors))
}

func mockResourcesMap() map[string]*schema.Resource {
	r := &schema.Resource{
		Schema: tf_provider.SchemaMap(),
		Importer: &schema.ResourceImporter{State: func(data *schema.ResourceData, v interface{}) ([]*schema.ResourceData, error) {
			var resourceDataSlice []*schema.ResourceData
			resourceDataSlice = append(resourceDataSlice, data)
			return resourceDataSlice, nil
		}},
	}
	dummyOCIResource := make(map[string]*schema.Resource)
	dummyOCIResource["oci_resource_type1"] = r
	return dummyOCIResource
}

func TestUnitGetDiscoverResourceSteps(t *testing.T) {

	//without targetResource
	t.Run("GetDiscoverResourceSteps without targetResource", func(t *testing.T) {
		ctx := getTestCtx()
		ctx.TargetSpecificResources = false
		result, err := getDiscoverResourceSteps(ctx)
		assert.NoError(t, err)
		assert.Equal(t, 0, len(result))
	})

	// With targetResource
	t.Run("GetDiscoverResourceSteps With targetResource", func(t *testing.T) {
		ctx := getTestCtx()
		ctx.TargetSpecificResources = true
		result, err := getDiscoverResourceSteps(ctx)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(result))
	})

}
func TestUnitGetDiscoverResourceWithGraphSteps(t *testing.T) {

	//without targetResource
	initResourceDiscoveryTests()
	t.Run("GetDiscoverResourceWithGraphSteps without targetResource", func(t *testing.T) {
		ctx := getTestCtx()
		ctx.TargetSpecificResources = false
		ctx.Services = []string{"identity"}
		result, err := getDiscoverResourceWithGraphSteps(ctx)
		assert.NoError(t, err)
		assert.Equal(t, 0, len(result))
	})

	// With targetResource
	t.Run("GetDiscoverResourceWithGraphSteps With targetResource", func(t *testing.T) {
		ctx := getTestCtx()
		ctx.TargetSpecificResources = true
		ctx.Services = []string{"budget"}
		result, err := getDiscoverResourceWithGraphSteps(ctx)
		assert.NoError(t, err)
		assert.Equal(t, 0, len(result))
	})
	// With targetResource
	t.Run("GetDiscoverResourceWithGraphSteps With tanaceyocid = compartment ocid", func(t *testing.T) {
		ctx := getTestCtx()
		ctx.TargetSpecificResources = true
		ctx.CompartmentId = &ctx.TenancyOcid
		ctx.Services = []string{"budget"}
		result, err := getDiscoverResourceWithGraphSteps(ctx)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(result))
	})

}
func TestUnitDeleteInvalidReferences(t *testing.T) {
	t.Run("Zero Error resource", func(t *testing.T) {

		referenceMap := make(map[string]string)
		referenceMap["1"] = "ocid1.3.b.c"
		referenceMap["2"] = "ocid2.a.b.c"
		referenceMap["3"] = "ocid3.d.b.c"

		discoveredResources := make([]*tf_export.OCIResource, 0, 3)
		resource := tf_export.OCIResource{
			IsErrorResource:   false,
			TerraformResource: tf_export.TerraformResource{Id: "1"},
		}
		discoveredResources = append(discoveredResources, &resource)
		deleteInvalidReferences(referenceMap, discoveredResources)
		assert.Equal(t, 3, len(referenceMap), "resources shouldn't be removed")
	})

	t.Run("One Error resource ", func(t *testing.T) {
		referenceMap := make(map[string]string)
		referenceMap["1"] = "ocid1.3.b.c"
		referenceMap["2"] = "ocid2.a.b.c"
		referenceMap["3"] = "ocid3.d.b.c"

		discoveredResources := make([]*tf_export.OCIResource, 0, 3)
		resource := tf_export.OCIResource{
			IsErrorResource:   true,
			TerraformResource: tf_export.TerraformResource{Id: "2"},
		}
		discoveredResources = append(discoveredResources, &resource)
		deleteInvalidReferences(referenceMap, discoveredResources)
		assert.Equal(t, 2, len(referenceMap), "Error resources should be removed")
		_, ok := referenceMap["2"]
		assert.False(t, ok, "Error resource should not be into referencemap")
	})
}
func TestUnitHasFreeformTag(t *testing.T) {
	resource := tf_export.OCIResource{
		SourceAttributes: map[string]interface{}{
			"freeform_tags": map[string]interface{}{
				"myPresentTag": "present",
			},
		},
	}
	assert.True(t, resource.HasFreeformTag("myPresentTag"), "should return True for present tag")
	assert.False(t, resource.HasFreeformTag("myNotPresentTag"), "should return False for not present tag")
}
func TestUnitHasDefinedTag(t *testing.T) {
	resource := tf_export.OCIResource{
		SourceAttributes: map[string]interface{}{
			"defined_tags": map[string]interface{}{
				"myDefinedTag": "YES",
			},
		},
	}
	assert.True(t, resource.HasDefinedTag("myDefinedTag", "YES"), "should return True for defined tag")
	assert.False(t, resource.HasDefinedTag("myDefinedTag", "NO"), "should return False for Not Defined tag")
}

func TestUnitParseDeliveryPolicy(t *testing.T) {
	policy := make(map[string]interface{})
	policy["backoff_retry_policy"] = []interface{}{
		map[string]interface{}{
			"max_retry_duration": "2h",
			"policy_type":        "NA",
		},
	}
	assert.NotNil(t, tf_export.ParseDeliveryPolicy(policy))
}

/*
   func TestUnitDeleteInvalidReferencesWithReferenceResource(t *testing.T) {
   	t.Run("One Error resource with another reference resource", func(t *testing.T) {
   		referenceMap := make(map[string]string)
   		referenceMap["1"] = "ocid1.3.b.c"
   		referenceMap["2"] = "ocid2.a.b.c"
   		referenceMap["3"] = "ocid3.1.b.c"

   		discoveredResources := make([]*commonexport.OCIResource, 0, 3)
   		resource := commonexport.OCIResource{
   			IsErrorResource:   true,
   			TerraformResource: commonexport.TerraformResource{Id: "1", TerraformName: "1"},
   		}
   		discoveredResources = append(discoveredResources, &resource)
   		deleteInvalidReferences(referenceMap, discoveredResources)
   		assert.Equal(t, 1, len(referenceMap), "Error resources should be removed")
   		_, ok := referenceMap["3"]
   		assert.False(t, ok, "Error resource should not be into referencemap")
   	})
   }

*/
