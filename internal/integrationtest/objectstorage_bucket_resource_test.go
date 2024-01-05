// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	retentionRuleRepresentation1 = map[string]interface{}{
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `sampleRetentionRule1`, Update: `sampleRetentionRule1`},
		"duration":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: retentionRuleDurationRepresentation},
		"time_rule_locked": acctest.Representation{RepType: acctest.Optional, Create: `2120-05-04T17:23:46Z`, Update: `2120-05-06T17:23:46Z`},
	}

	retentionRuleRepresentation2 = map[string]interface{}{
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `sampleRetentionRule2`, Update: `sampleRetentionRule2`},
		"duration":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: retentionRuleDurationRepresentation},
		"time_rule_locked": acctest.Representation{RepType: acctest.Optional, Create: `2120-05-04T17:23:46Z`, Update: `2120-05-06T17:23:46Z`},
	}

	retentionRuleRepresentation3 = map[string]interface{}{
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `sampleRetentionRule3`, Update: `sampleRetentionRule3`},
		"duration":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: retentionRuleDurationRepresentation},
		"time_rule_locked": acctest.Representation{RepType: acctest.Optional, Create: `2120-05-04T17:23:46Z`, Update: `2120-05-06T17:23:46Z`},
	}

	retentionRuleRepresentation4 = map[string]interface{}{
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `sampleRetentionRule4`, Update: `sampleRetentionRule4`},
		"duration":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: retentionRuleDurationRepresentation},
	}

	retentionRuleDurationRepresentation = map[string]interface{}{
		"time_amount": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `15`},
		"time_unit":   acctest.Representation{RepType: acctest.Required, Create: `DAYS`, Update: `DAYS`},
	}

	bucketRepresentationWithoutUpdateToForceNewFields = acctest.RepresentationCopyWithNewProperties(
		acctest.RepresentationCopyWithRemovedProperties(ObjectStorageBucketRepresentation, []string{"name", "namespace", "storage_tier", "versioning"}),
		map[string]interface{}{
			"name":         acctest.Representation{RepType: acctest.Required, Create: testBucketName},
			"namespace":    acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
			"storage_tier": acctest.Representation{RepType: acctest.Optional, Create: `Standard`},
		})

	bucketRepresentationForRetentionRules = acctest.RepresentationCopyWithNewProperties(bucketRepresentationWithoutUpdateToForceNewFields,
		map[string]interface{}{
			"retention_rules": []acctest.RepresentationGroup{
				{RepType: acctest.Optional, Group: retentionRuleRepresentation1},
				{RepType: acctest.Optional, Group: retentionRuleRepresentation2},
				{RepType: acctest.Optional, Group: retentionRuleRepresentation3},
			},
		})

	bucketRepresentationForRetentionRulesWithoutLock = acctest.RepresentationCopyWithNewProperties(bucketRepresentationWithoutUpdateToForceNewFields,
		map[string]interface{}{
			"retention_rules": []acctest.RepresentationGroup{
				{RepType: acctest.Optional, Group: retentionRuleRepresentation1},
				{RepType: acctest.Optional, Group: retentionRuleRepresentation2},
				{RepType: acctest.Optional, Group: retentionRuleRepresentation3},
				{RepType: acctest.Optional, Group: retentionRuleRepresentation4},
			},
		})

	bucketRepresentationForRetentionRulesReordered = acctest.RepresentationCopyWithNewProperties(bucketRepresentationWithoutUpdateToForceNewFields,
		map[string]interface{}{
			"retention_rules": []acctest.RepresentationGroup{
				{RepType: acctest.Optional, Group: retentionRuleRepresentation1},
				{RepType: acctest.Optional, Group: retentionRuleRepresentation3},
				{RepType: acctest.Optional, Group: retentionRuleRepresentation2},
			},
		})

	bucketRepresentationForRetentionRulesDelete = acctest.RepresentationCopyWithNewProperties(bucketRepresentationWithoutUpdateToForceNewFields,
		map[string]interface{}{
			"retention_rules": []acctest.RepresentationGroup{
				{RepType: acctest.Optional, Group: retentionRuleRepresentation1},
			},
		})

	bucketSingularDataSourceRetentionRulesRepresentation = map[string]interface{}{
		"name":      acctest.Representation{RepType: acctest.Required, Create: testBucketName},
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	}
)

// issue-routing-tag: object_storage/default
func TestResourceBucket_retentionRules(t *testing.T) {
	httpreplay.SetScenario("TestObjectStorageBucketResource_retentionRules")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_objectstorage_bucket.test_bucket"
	singularDatasourceName := "data.oci_objectstorage_bucket.test_bucket"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckObjectStorageBucketDestroy,
		Steps: []resource.TestStep{
			//verify Create with optionals
			{
				Config: config + compartmentIdVariableStr + ObjectStorageBucketResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Optional, acctest.Create, bucketRepresentationForRetentionRules),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "access_type", "NoPublicAccess"),
					resource.TestCheckResourceAttrSet(resourceName, "bucket_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
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
					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule1",
							"duration.#":             "1",
							"duration.0.time_amount": "10",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-04T17:23:46Z",
						},
						[]string{}),

					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule2",
							"duration.#":             "1",
							"duration.0.time_amount": "10",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-04T17:23:46Z",
						},
						[]string{}),

					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule3",
							"duration.#":             "1",
							"duration.0.time_amount": "10",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-04T17:23:46Z",
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
			// verify updates to updatable parameters of retention rules
			// changing display Name forces deletion of old rule and creation of the new rule
			{
				Config: config + compartmentIdVariableStr + ObjectStorageBucketResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Optional, acctest.Update, bucketRepresentationForRetentionRules),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "access_type", "ObjectRead"),
					resource.TestCheckResourceAttrSet(resourceName, "bucket_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
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

					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule1",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule2",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule3",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				Config: config + compartmentIdVariableStr + ObjectStorageBucketResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Optional, acctest.Update, bucketRepresentationForRetentionRulesReordered),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "access_type", "ObjectRead"),
					resource.TestCheckResourceAttrSet(resourceName, "bucket_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
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

					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule1",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule2",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule3",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				Config: config + compartmentIdVariableStr + ObjectStorageBucketResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Optional, acctest.Update, bucketRepresentationForRetentionRulesDelete),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "access_type", "ObjectRead"),
					resource.TestCheckResourceAttrSet(resourceName, "bucket_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
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

					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule1",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				Config: config + compartmentIdVariableStr + ObjectStorageBucketResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Optional, acctest.Update, bucketRepresentationForRetentionRules),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "access_type", "ObjectRead"),
					resource.TestCheckResourceAttrSet(resourceName, "bucket_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
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

					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule1",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule2",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule3",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				Config: config + compartmentIdVariableStr + ObjectStorageBucketResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Optional, acctest.Update, bucketRepresentationForRetentionRulesWithoutLock),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "access_type", "ObjectRead"),
					resource.TestCheckResourceAttrSet(resourceName, "bucket_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
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

					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule1",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule2",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule3",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule4",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
						},
						[]string{}),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				Config: config + compartmentIdVariableStr + ObjectStorageBucketResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Optional, acctest.Update, bucketRepresentationForRetentionRules),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "access_type", "ObjectRead"),
					resource.TestCheckResourceAttrSet(resourceName, "bucket_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
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

					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule1",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule2",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "retention_rules",
						map[string]string{
							"display_name":           "sampleRetentionRule3",
							"duration.#":             "1",
							"duration.0.time_amount": "15",
							"duration.0.time_unit":   "DAYS",
							"time_rule_locked":       "2120-05-06T17:23:46Z",
						},
						[]string{}),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				Config: config + compartmentIdVariableStr + ObjectStorageBucketResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Optional, acctest.Update, bucketRepresentationForRetentionRulesWithoutLock) +
					acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, bucketSingularDataSourceRetentionRulesRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(singularDatasourceName, "name", testBucketName),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
					resource.TestCheckResourceAttr(singularDatasourceName, "access_type", "ObjectRead"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "approximate_count"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "approximate_size"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "bucket_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
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
			//verify resource import
			{
				Config:                  config + ObjectStorageBucketRequiredOnlyResource,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}
