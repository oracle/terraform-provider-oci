// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

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
	retentionRuleRepresentation1 = map[string]interface{}{
		"display_name":     Representation{repType: Optional, create: `sampleRetentionRule1`, update: `sampleRetentionRule1`},
		"duration":         RepresentationGroup{Optional, retentionRuleDurationRepresentation},
		"time_rule_locked": Representation{repType: Optional, create: `2120-05-04T17:23:46Z`, update: `2120-05-06T17:23:46Z`},
	}

	retentionRuleRepresentation2 = map[string]interface{}{
		"display_name":     Representation{repType: Optional, create: `sampleRetentionRule2`, update: `sampleRetentionRule2`},
		"duration":         RepresentationGroup{Optional, retentionRuleDurationRepresentation},
		"time_rule_locked": Representation{repType: Optional, create: `2120-05-04T17:23:46Z`, update: `2120-05-06T17:23:46Z`},
	}

	retentionRuleRepresentation3 = map[string]interface{}{
		"display_name":     Representation{repType: Optional, create: `sampleRetentionRule3`, update: `sampleRetentionRule3`},
		"duration":         RepresentationGroup{Optional, retentionRuleDurationRepresentation},
		"time_rule_locked": Representation{repType: Optional, create: `2120-05-04T17:23:46Z`, update: `2120-05-06T17:23:46Z`},
	}

	retentionRuleRepresentation4 = map[string]interface{}{
		"display_name": Representation{repType: Optional, create: `sampleRetentionRule4`, update: `sampleRetentionRule4`},
		"duration":     RepresentationGroup{Optional, retentionRuleDurationRepresentation},
	}

	retentionRuleDurationRepresentation = map[string]interface{}{
		"time_amount": Representation{repType: Required, create: `10`, update: `15`},
		"time_unit":   Representation{repType: Required, create: `DAYS`, update: `DAYS`},
	}

	bucketRepresentationWithoutUpdateToForceNewFields = representationCopyWithNewProperties(
		representationCopyWithRemovedProperties(bucketRepresentation, []string{"name", "namespace", "storage_tier", "versioning"}),
		map[string]interface{}{
			"name":         Representation{repType: Required, create: testBucketName},
			"namespace":    Representation{repType: Required, create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
			"storage_tier": Representation{repType: Optional, create: `Standard`},
		})

	bucketRepresentationForRetentionRules = representationCopyWithNewProperties(bucketRepresentationWithoutUpdateToForceNewFields,
		map[string]interface{}{
			"retention_rules": []RepresentationGroup{
				{Optional, retentionRuleRepresentation1},
				{Optional, retentionRuleRepresentation2},
				{Optional, retentionRuleRepresentation3},
			},
		})

	bucketRepresentationForRetentionRulesWithoutLock = representationCopyWithNewProperties(bucketRepresentationWithoutUpdateToForceNewFields,
		map[string]interface{}{
			"retention_rules": []RepresentationGroup{
				{Optional, retentionRuleRepresentation1},
				{Optional, retentionRuleRepresentation2},
				{Optional, retentionRuleRepresentation3},
				{Optional, retentionRuleRepresentation4},
			},
		})

	bucketRepresentationForRetentionRulesReordered = representationCopyWithNewProperties(bucketRepresentationWithoutUpdateToForceNewFields,
		map[string]interface{}{
			"retention_rules": []RepresentationGroup{
				{Optional, retentionRuleRepresentation1},
				{Optional, retentionRuleRepresentation3},
				{Optional, retentionRuleRepresentation2},
			},
		})

	bucketRepresentationForRetentionRulesDelete = representationCopyWithNewProperties(bucketRepresentationWithoutUpdateToForceNewFields,
		map[string]interface{}{
			"retention_rules": []RepresentationGroup{
				{Optional, retentionRuleRepresentation1},
			},
		})

	bucketSingularDataSourceRetentionRulesRepresentation = map[string]interface{}{
		"name":      Representation{repType: Required, create: testBucketName},
		"namespace": Representation{repType: Required, create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	}
)

func TestObjectStorageBucketResource_retentionRules(t *testing.T) {
	httpreplay.SetScenario("TestObjectStorageBucketResource_retentionRules")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_objectstorage_bucket.test_bucket"
	singularDatasourceName := "data.oci_objectstorage_bucket.test_bucket"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckObjectStorageBucketDestroy,
		Steps: []resource.TestStep{
			//verify create with optionals
			{
				Config: config + compartmentIdVariableStr + BucketResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Optional, Create, bucketRepresentationForRetentionRules),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "access_type", "NoPublicAccess"),
					resource.TestCheckResourceAttrSet(resourceName, "bucket_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "etag"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "name", testBucketName),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object_events_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "storage_tier", "Standard"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					resource.TestCheckResourceAttr(resourceName, "retention_rules.#", "3"),
					CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule1",
							"duration.#":             "1",
							"duration.0.time_amount": "10",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-04T17:23:46Z",
						},
						[]string{}),

					CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule2",
							"duration.#":             "1",
							"duration.0.time_amount": "10",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-04T17:23:46Z",
						},
						[]string{}),

					CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule3",
							"duration.#":             "1",
							"duration.0.time_amount": "10",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-04T17:23:46Z",
						},
						[]string{}),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},
			// verify updates to updatable parameters of retention rules
			// changing display Name forces deletion of old rule and creation of the new rule
			{
				Config: config + compartmentIdVariableStr + BucketResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Optional, Update, bucketRepresentationForRetentionRules),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "access_type", "ObjectRead"),
					resource.TestCheckResourceAttrSet(resourceName, "bucket_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "etag"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					//resource.TestCheckResourceAttr(resourceName, "name", testBucketName2),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object_events_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "storage_tier", "Standard"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					resource.TestCheckResourceAttr(resourceName, "retention_rules.#", "3"),

					CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule1",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule2",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule3",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						// The id changes when the name changes
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify that no change on reordering
			{
				Config: config + compartmentIdVariableStr + BucketResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Optional, Update, bucketRepresentationForRetentionRulesReordered),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "access_type", "ObjectRead"),
					resource.TestCheckResourceAttrSet(resourceName, "bucket_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "etag"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "name", testBucketName),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object_events_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "storage_tier", "Standard"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					resource.TestCheckResourceAttr(resourceName, "retention_rules.#", "3"),

					CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule1",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule2",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule3",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						// The id changes when the name changes
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify retention rules delete
			{
				Config: config + compartmentIdVariableStr + BucketResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Optional, Update, bucketRepresentationForRetentionRulesDelete),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "access_type", "ObjectRead"),
					resource.TestCheckResourceAttrSet(resourceName, "bucket_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "etag"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "name", testBucketName),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object_events_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "storage_tier", "Standard"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					resource.TestCheckResourceAttr(resourceName, "retention_rules.#", "1"),

					CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule1",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						// The id changes when the name changes
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify retention rules add new rules
			{
				Config: config + compartmentIdVariableStr + BucketResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Optional, Update, bucketRepresentationForRetentionRules),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "access_type", "ObjectRead"),
					resource.TestCheckResourceAttrSet(resourceName, "bucket_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "etag"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "name", testBucketName),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object_events_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "storage_tier", "Standard"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					resource.TestCheckResourceAttr(resourceName, "retention_rules.#", "3"),

					CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule1",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule2",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule3",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						// The id changes when the name changes
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify adding a new retention rule without timeRuleLocked
			{
				Config: config + compartmentIdVariableStr + BucketResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Optional, Update, bucketRepresentationForRetentionRulesWithoutLock),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "access_type", "ObjectRead"),
					resource.TestCheckResourceAttrSet(resourceName, "bucket_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "etag"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "name", testBucketName),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object_events_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "storage_tier", "Standard"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					resource.TestCheckResourceAttr(resourceName, "retention_rules.#", "4"),

					CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule1",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule2",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule3",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule4",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
						},
						[]string{}),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						// The id changes when the name changes
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify deleting a old retention rule without timeRuleLocked
			{
				Config: config + compartmentIdVariableStr + BucketResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Optional, Update, bucketRepresentationForRetentionRules),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "access_type", "ObjectRead"),
					resource.TestCheckResourceAttrSet(resourceName, "bucket_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "etag"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
					resource.TestCheckResourceAttr(resourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "name", testBucketName),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object_events_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "storage_tier", "Standard"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					resource.TestCheckResourceAttr(resourceName, "retention_rules.#", "3"),

					CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule1",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule2",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule3",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						// The id changes when the name changes
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			//verify singular datasource
			{
				Config: config + compartmentIdVariableStr + BucketResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Optional, Update, bucketRepresentationForRetentionRulesWithoutLock) +
					generateDataSourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Required, Create, bucketSingularDataSourceRetentionRulesRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(singularDatasourceName, "name", testBucketName),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
					resource.TestCheckResourceAttr(singularDatasourceName, "access_type", "ObjectRead"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "approximate_count"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "approximate_size"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "bucket_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "etag"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "metadata.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", testBucketName),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
					// This is difficult to test because TF is eager in creating the datasource and gives stale results.
					// If a depends_on is added, we get an error like "After applying this step and refreshing, the plan was not empty:"
					//resource.TestCheckResourceAttrSet(singularDatasourceName, "object_lifecycle_policy_etag"),
					resource.TestCheckResourceAttr(singularDatasourceName, "object_events_enabled", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "storage_tier", "Standard"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "retention_rules.#"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "retention_rules.0.display_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "retention_rules.0.duration.#"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "retention_rules.0.duration.0.time_amount"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "retention_rules.0.duration.0.time_unit"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "retention_rules.0.retention_rule_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "retention_rules.0.time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "retention_rules.0.time_modified"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "retention_rules.0.time_rule_locked"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "retention_rules.1.display_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "retention_rules.1.duration.#"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "retention_rules.1.duration.0.time_amount"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "retention_rules.1.duration.0.time_unit"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "retention_rules.1.retention_rule_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "retention_rules.1.time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "retention_rules.1.time_modified"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "retention_rules.1.time_rule_locked"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "retention_rules.2.display_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "retention_rules.2.duration.#"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "retention_rules.2.duration.0.time_amount"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "retention_rules.2.duration.0.time_unit"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "retention_rules.2.retention_rule_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "retention_rules.2.time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "retention_rules.2.time_modified"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "retention_rules.2.time_rule_locked"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + BucketResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Optional, Update, bucketRepresentationForRetentionRules),
			},
			//verify resource import
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
