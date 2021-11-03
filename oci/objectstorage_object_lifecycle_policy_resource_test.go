package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	objectLifecyclePolicyRulesRepresentation_ForMultiPartUploads = map[string]interface{}{
		"action":      Representation{RepType: Required, Create: `ARCHIVE`, Update: `ABORT`},
		"is_enabled":  Representation{RepType: Required, Create: `false`, Update: `true`},
		"name":        Representation{RepType: Required, Create: `sampleRule`, Update: `name2`},
		"time_amount": Representation{RepType: Required, Create: `10`, Update: `11`},
		"time_unit":   Representation{RepType: Required, Create: `DAYS`, Update: `YEARS`},
		"target":      Representation{RepType: Optional, Update: `multipart-uploads`},
	}
)

// issue-routing-tag: object_storage/default
func TestResourceObjectLifecyclePolicy_validations(t *testing.T) {
	httpreplay.SetScenario("TestObjectStorageObjectLifecyclePolicyResource_validations")
	defer httpreplay.SaveScenario()
	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_objectstorage_object_lifecycle_policy.test_object_lifecycle_policy"

	objectLifecyclePolicyRulesObjectNameFilterDifferentOrderRepresentation := map[string]interface{}{
		"inclusion_prefixes": Representation{RepType: Optional, Create: []string{bucketName, bucketName2}, Update: []string{bucketName, bucketName2, bucketName3}},
		"inclusion_patterns": Representation{RepType: Optional, Create: []string{`inclusionPattern1`, `inclusionPattern2`}, Update: []string{`inclusionPattern1`, `inclusionPattern2`, `inclusionPattern3`}},
		"exclusion_patterns": Representation{RepType: Optional, Create: []string{`exclusionPattern1`, `exclusionPattern2`}, Update: []string{`exclusionPattern1`, `exclusionPattern2`, `exclusionPattern3`}},
	}

	objectLifecyclePolicyRulesObjectNameFilterOneValueIncludeRepresentation := map[string]interface{}{
		"inclusion_patterns": Representation{RepType: Optional, Create: []string{`inclusionPattern1`}},
	}

	objectLifecyclePolicyRulesObjectNameFilterOneValueExcludeRepresentation := map[string]interface{}{
		"exclusion_patterns": Representation{RepType: Optional, Create: []string{`inclusionPattern1`}},
	}

	ResourceTest(t, testAccCheckObjectStorageObjectLifecyclePolicyDestroy, []resource.TestStep{
		// verify baseline Create
		{
			Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Create, objectLifecyclePolicyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bucket", bucketName),
				resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),

				func(s *terraform.State) (err error) {
					_, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// change order of inclusion prefixes
		{
			Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Create,
					GetUpdatedRepresentationCopy("rules.object_name_filter.inclusion_prefixes", Representation{RepType: Optional, Create: []string{bucketName2, bucketName}}, objectLifecyclePolicyRepresentation)),
			PlanOnly:           true,
			ExpectNonEmptyPlan: false,
		},
		// Remove inclusion prefixes to see plan has changed
		{
			Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Create,
					GetUpdatedRepresentationCopy("rules.object_name_filter.inclusion_prefixes", Representation{RepType: Optional, Create: []string{}}, objectLifecyclePolicyRepresentation)),
			PlanOnly:           true,
			ExpectNonEmptyPlan: true,
		},
		// Change the value for the inclusion prefixes to see plan has changed
		{
			Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Create,
					GetUpdatedRepresentationCopy("rules.object_name_filter.inclusion_prefixes", Representation{RepType: Optional, Create: []string{bucketName, bucketName2 + "_test"}}, objectLifecyclePolicyRepresentation)),
			PlanOnly:           true,
			ExpectNonEmptyPlan: true,
		},
		// change order of inclusion patterns
		{
			Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Create,
					GetUpdatedRepresentationCopy("rules.object_name_filter.inclusion_patterns", Representation{RepType: Optional, Create: []string{`inclusionPattern2`, `inclusionPattern1`}}, objectLifecyclePolicyRepresentation)),
			PlanOnly:           true,
			ExpectNonEmptyPlan: false,
		},
		// Remove inclusion patterns to see plan has changed
		{
			Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Create,
					GetUpdatedRepresentationCopy("rules.object_name_filter.inclusion_patterns", Representation{RepType: Optional, Create: []string{}}, objectLifecyclePolicyRepresentation)),
			PlanOnly:           true,
			ExpectNonEmptyPlan: true,
		},
		// change order of exclusion patterns
		{
			Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Create,
					GetUpdatedRepresentationCopy("rules.object_name_filter.exclusion_patterns", Representation{RepType: Optional, Create: []string{`exclusionPattern2`, `exclusionPattern1`}}, objectLifecyclePolicyRepresentation)),
			PlanOnly:           true,
			ExpectNonEmptyPlan: false,
		},
		// Remove exclusion patterns to see plan has changed
		{
			Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Create,
					GetUpdatedRepresentationCopy("rules.object_name_filter.exclusion_patterns", Representation{RepType: Optional, Create: []string{}}, objectLifecyclePolicyRepresentation)),
			PlanOnly:           true,
			ExpectNonEmptyPlan: true,
		},
		// change order of object_name_filter
		{
			Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Create,
					GetUpdatedRepresentationCopy("rules.object_name_filter", RepresentationGroup{Optional, objectLifecyclePolicyRulesObjectNameFilterDifferentOrderRepresentation}, objectLifecyclePolicyRepresentation)),
			PlanOnly:           true,
			ExpectNonEmptyPlan: false,
		},

		// Update the object_name_filter properties with the only one inclusion_patterns value
		{
			Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Create,
					GetUpdatedRepresentationCopy("rules.object_name_filter", RepresentationGroup{Optional, objectLifecyclePolicyRulesObjectNameFilterOneValueIncludeRepresentation}, objectLifecyclePolicyRepresentation)),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bucket", bucketName),
				resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),

				func(s *terraform.State) (err error) {
					_, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// change to the same value for the exclusion_patterns
		{
			Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Create,
					GetUpdatedRepresentationCopy("rules.object_name_filter", RepresentationGroup{Optional, objectLifecyclePolicyRulesObjectNameFilterOneValueExcludeRepresentation}, objectLifecyclePolicyRepresentation)),
			PlanOnly:           true,
			ExpectNonEmptyPlan: true,
		},
	})
}

// issue-routing-tag: object_storage/default
func TestResourceObjectLifecyclePolicy_MultiPartUploadsRule(t *testing.T) {
	httpreplay.SetScenario("TestObjectStorageObjectLifecyclePolicyResource_MultiPartUploadsRule")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_objectstorage_object_lifecycle_policy.test_object_lifecycle_policy"

	singularDatasourceName := "data.oci_objectstorage_object_lifecycle_policy.test_object_lifecycle_policy"

	var resId, resId2 string

	ResourceTest(t, testAccCheckObjectStorageObjectLifecyclePolicyDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Create, objectLifecyclePolicyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bucket", bucketName),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
				CheckResourceSetContainsElementWithProperties(resourceName, "rules", map[string]string{
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
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates the rule for multipart-uploads abort feature
		{
			Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Update,
					RepresentationCopyWithNewProperties(objectLifecyclePolicyRepresentation, map[string]interface{}{
						"rules": RepresentationGroup{Optional, objectLifecyclePolicyRulesRepresentation_ForMultiPartUploads},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bucket", bucketName),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
				CheckResourceSetContainsElementWithProperties(resourceName, "rules", map[string]string{
					"action":      "ABORT",
					"is_enabled":  "true",
					"name":        "name2",
					"target":      "multipart-uploads",
					"time_amount": "11",
					"time_unit":   "YEARS",
				},
					[]string{}),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Required, Create, objectLifecyclePolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + ObjectLifecyclePolicyResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "bucket", bucketName),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.#", "1"),
				CheckResourceSetContainsElementWithProperties(singularDatasourceName, "rules", map[string]string{},
					[]string{}),

				resource.TestCheckResourceAttr(singularDatasourceName, "rules.#", "1"),
				CheckResourceSetContainsElementWithProperties(singularDatasourceName, "rules", map[string]string{
					"action":      "ABORT",
					"is_enabled":  "true",
					"name":        "name2",
					"target":      "multipart-uploads",
					"time_amount": "11",
					"time_unit":   "YEARS",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
	})
}
