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
		"action":      Representation{repType: Required, create: `ARCHIVE`, update: `ABORT`},
		"is_enabled":  Representation{repType: Required, create: `false`, update: `true`},
		"name":        Representation{repType: Required, create: `sampleRule`, update: `name2`},
		"time_amount": Representation{repType: Required, create: `10`, update: `11`},
		"time_unit":   Representation{repType: Required, create: `DAYS`, update: `YEARS`},
		"target":      Representation{repType: Optional, update: `multipart-uploads`},
	}
)

func TestObjectStorageObjectLifecyclePolicyResource_validations(t *testing.T) {
	httpreplay.SetScenario("TestObjectStorageObjectLifecyclePolicyResource_validations")
	defer httpreplay.SaveScenario()
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_objectstorage_object_lifecycle_policy.test_object_lifecycle_policy"

	objectLifecyclePolicyRulesObjectNameFilterDifferentOrderRepresentation := map[string]interface{}{
		"inclusion_prefixes": Representation{repType: Optional, create: []string{bucketName, bucketName2}, update: []string{bucketName, bucketName2, bucketName3}},
		"inclusion_patterns": Representation{repType: Optional, create: []string{`inclusionPattern1`, `inclusionPattern2`}, update: []string{`inclusionPattern1`, `inclusionPattern2`, `inclusionPattern3`}},
		"exclusion_patterns": Representation{repType: Optional, create: []string{`exclusionPattern1`, `exclusionPattern2`}, update: []string{`exclusionPattern1`, `exclusionPattern2`, `exclusionPattern3`}},
	}

	objectLifecyclePolicyRulesObjectNameFilterOneValueIncludeRepresentation := map[string]interface{}{
		"inclusion_patterns": Representation{repType: Optional, create: []string{`inclusionPattern1`}},
	}

	objectLifecyclePolicyRulesObjectNameFilterOneValueExcludeRepresentation := map[string]interface{}{
		"exclusion_patterns": Representation{repType: Optional, create: []string{`inclusionPattern1`}},
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckObjectStorageObjectLifecyclePolicyDestroy,
		Steps: []resource.TestStep{
			// verify baseline create
			{
				Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Create, objectLifecyclePolicyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "bucket", bucketName),
					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),

					func(s *terraform.State) (err error) {
						_, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// change order of inclusion prefixes
			{
				Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Create,
						getUpdatedRepresentationCopy("rules.object_name_filter.inclusion_prefixes", Representation{repType: Optional, create: []string{bucketName2, bucketName}}, objectLifecyclePolicyRepresentation)),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			// Remove inclusion prefixes to see plan has changed
			{
				Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Create,
						getUpdatedRepresentationCopy("rules.object_name_filter.inclusion_prefixes", Representation{repType: Optional, create: []string{}}, objectLifecyclePolicyRepresentation)),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			// Change the value for the inclusion prefixes to see plan has changed
			{
				Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Create,
						getUpdatedRepresentationCopy("rules.object_name_filter.inclusion_prefixes", Representation{repType: Optional, create: []string{bucketName, bucketName2 + "_test"}}, objectLifecyclePolicyRepresentation)),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			// change order of inclusion patterns
			{
				Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Create,
						getUpdatedRepresentationCopy("rules.object_name_filter.inclusion_patterns", Representation{repType: Optional, create: []string{`inclusionPattern2`, `inclusionPattern1`}}, objectLifecyclePolicyRepresentation)),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			// Remove inclusion patterns to see plan has changed
			{
				Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Create,
						getUpdatedRepresentationCopy("rules.object_name_filter.inclusion_patterns", Representation{repType: Optional, create: []string{}}, objectLifecyclePolicyRepresentation)),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			// change order of exclusion patterns
			{
				Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Create,
						getUpdatedRepresentationCopy("rules.object_name_filter.exclusion_patterns", Representation{repType: Optional, create: []string{`exclusionPattern2`, `exclusionPattern1`}}, objectLifecyclePolicyRepresentation)),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			// Remove exclusion patterns to see plan has changed
			{
				Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Create,
						getUpdatedRepresentationCopy("rules.object_name_filter.exclusion_patterns", Representation{repType: Optional, create: []string{}}, objectLifecyclePolicyRepresentation)),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			// change order of object_name_filter
			{
				Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Create,
						getUpdatedRepresentationCopy("rules.object_name_filter", RepresentationGroup{Optional, objectLifecyclePolicyRulesObjectNameFilterDifferentOrderRepresentation}, objectLifecyclePolicyRepresentation)),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},

			// update the object_name_filter properties with the only one inclusion_patterns value
			{
				Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Create,
						getUpdatedRepresentationCopy("rules.object_name_filter", RepresentationGroup{Optional, objectLifecyclePolicyRulesObjectNameFilterOneValueIncludeRepresentation}, objectLifecyclePolicyRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "bucket", bucketName),
					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),

					func(s *terraform.State) (err error) {
						_, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// change to the same value for the exclusion_patterns
			{
				Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Create,
						getUpdatedRepresentationCopy("rules.object_name_filter", RepresentationGroup{Optional, objectLifecyclePolicyRulesObjectNameFilterOneValueExcludeRepresentation}, objectLifecyclePolicyRepresentation)),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestObjectStorageObjectLifecyclePolicyResource_MultiPartUploadsRule(t *testing.T) {
	httpreplay.SetScenario("TestObjectStorageObjectLifecyclePolicyResource_MultiPartUploadsRule")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_objectstorage_object_lifecycle_policy.test_object_lifecycle_policy"

	singularDatasourceName := "data.oci_objectstorage_object_lifecycle_policy.test_object_lifecycle_policy"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckObjectStorageObjectLifecyclePolicyDestroy,
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + ObjectLifecyclePolicyResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Create, objectLifecyclePolicyRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
					generateResourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Optional, Update,
						representationCopyWithNewProperties(objectLifecyclePolicyRepresentation, map[string]interface{}{
							"rules": RepresentationGroup{Optional, objectLifecyclePolicyRulesRepresentation_ForMultiPartUploads},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_objectstorage_object_lifecycle_policy", "test_object_lifecycle_policy", Required, Create, objectLifecyclePolicySingularDataSourceRepresentation) +
					compartmentIdVariableStr + ObjectLifecyclePolicyResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
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
		},
	})
}
