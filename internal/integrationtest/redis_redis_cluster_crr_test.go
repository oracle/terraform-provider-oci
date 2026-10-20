// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

// issue-routing-tag: redis/default
func TestRedisRedisClusterCRRResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestRedisRedisClusterCRRResource_basic")
	defer httpreplay.SaveScenario()

	inputs := redisCRRTestInputs{
		compartmentID:      utils.GetEnvSettingWithBlankDefault("compartment_ocid"),
		primaryClusterID:   utils.GetEnvSettingWithBlankDefault("redis_crr_primary_cluster_id"),
		subnetID:           utils.GetEnvSettingWithBlankDefault("redis_crr_subnet_id"),
		configSetID:        utils.GetEnvSettingWithBlankDefault("redis_crr_oci_cache_config_set_id"),
		nodeCount:          utils.GetEnvSettingWithBlankDefault("redis_crr_node_count"),
		nodeMemoryInGBs:    utils.GetEnvSettingWithBlankDefault("redis_crr_node_memory_in_gbs"),
		softwareVersion:    utils.GetEnvSettingWithBlankDefault("redis_crr_software_version"),
		clusterMode:        utils.GetEnvSettingWithDefault("redis_crr_cluster_mode", "NONSHARDED"),
		alternatePrimaryID: utils.GetEnvSettingWithBlankDefault("redis_crr_alternate_primary_cluster_id"),
	}
	if missing := inputs.missing(); missing != "" {
		t.Skipf("skipping Redis CRR acceptance test; set %s", missing)
	}
	if inputs.alternatePrimaryID == "" {
		inputs.alternatePrimaryID = differentPrimaryClusterID(inputs.primaryClusterID)
	}

	config := acctest.ProviderTestConfig()
	resourceName := "oci_redis_redis_cluster.test_cluster"
	updatedDisplayName := "tf-acc-crr-lifecycle-secondary"
	var clusterID string

	acctest.SaveConfigContent(config+inputs.terraformConfig(inputs.primaryClusterID), "redis", "redisClusterCrr", t)

	acctest.ResourceTest(t, testAccCheckRedisRedisClusterDestroy, []resource.TestStep{
		// Create a secondary cluster by supplying primary_cluster_id at creation.
		{
			Config: config + inputs.terraformConfig(inputs.primaryClusterID),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "primary_cluster_id", inputs.primaryClusterID),
				resource.TestCheckResourceAttr(resourceName, "cluster_role", "SECONDARY"),
				func(s *terraform.State) error {
					var err error
					clusterID, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// Remove primary_cluster_id to convert the secondary cluster to standalone.
		{
			Config: config + inputs.terraformConfig(""),
			Check:  assertRedisCRRClusterState(resourceName, &clusterID, "STANDALONE", false),
		},

		// Update display_name and re-add primary_cluster_id in one apply. Regular updates must
		// complete before the standalone cluster is converted back to a secondary.
		{
			Config: config + inputs.terraformConfigWithDisplayName(inputs.primaryClusterID, updatedDisplayName),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				assertRedisCRRClusterState(resourceName, &clusterID, "SECONDARY", true),
				resource.TestCheckResourceAttr(resourceName, "display_name", updatedDisplayName),
			),
		},

		// Terraform must reject changing a secondary directly to a different primary.
		{
			Config:      config + inputs.terraformConfigWithDisplayName(inputs.alternatePrimaryID, updatedDisplayName),
			ExpectError: regexp.MustCompile("changing primary_cluster_id from one primary cluster to another is not supported"),
		},
	})
}

func assertRedisCRRClusterState(resourceName string, expectedID *string, expectedRole string, expectPrimaryClusterID bool) resource.TestCheckFunc {
	checks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr(resourceName, "cluster_role", expectedRole),
		func(s *terraform.State) error {
			actualID, err := acctest.FromInstanceState(s, resourceName, "id")
			if err != nil {
				return err
			}
			if *expectedID != actualID {
				return fmt.Errorf("cluster was recreated instead of converted: before=%s after=%s", *expectedID, actualID)
			}
			return nil
		},
	}
	if expectPrimaryClusterID {
		checks = append(checks, resource.TestCheckResourceAttrSet(resourceName, "primary_cluster_id"))
	} else {
		// An omitted optional string is represented by Terraform as an empty value in state.
		checks = append(checks, resource.TestCheckResourceAttr(resourceName, "primary_cluster_id", ""))
	}
	return acctest.ComposeAggregateTestCheckFuncWrapper(checks...)
}

type redisCRRTestInputs struct {
	compartmentID      string
	primaryClusterID   string
	alternatePrimaryID string
	subnetID           string
	configSetID        string
	nodeCount          string
	nodeMemoryInGBs    string
	softwareVersion    string
	clusterMode        string
}

func (i redisCRRTestInputs) missing() string {
	for name, value := range map[string]string{
		"TF_VAR_redis_crr_primary_cluster_id":      i.primaryClusterID,
		"TF_VAR_redis_crr_subnet_id":               i.subnetID,
		"TF_VAR_redis_crr_oci_cache_config_set_id": i.configSetID,
		"TF_VAR_redis_crr_node_count":              i.nodeCount,
		"TF_VAR_redis_crr_node_memory_in_gbs":      i.nodeMemoryInGBs,
		"TF_VAR_redis_crr_software_version":        i.softwareVersion,
	} {
		if value == "" {
			return name
		}
	}
	return ""
}

func (i redisCRRTestInputs) terraformConfig(primaryClusterID string) string {
	return i.terraformConfigWithDisplayName(primaryClusterID, "tf-acc-crr-lifecycle")
}

func (i redisCRRTestInputs) terraformConfigWithDisplayName(primaryClusterID string, displayName string) string {
	primaryClusterConfig := ""
	if primaryClusterID != "" {
		primaryClusterConfig = fmt.Sprintf("\n  primary_cluster_id = %q", primaryClusterID)
	}

	return fmt.Sprintf(`
resource "oci_redis_redis_cluster" "test_cluster" {
  compartment_id          = %q
  display_name            = %q
  node_count              = %s
  node_memory_in_gbs      = %s
  software_version        = %q
  subnet_id               = %q
  cluster_mode            = %q
  oci_cache_config_set_id = %q%s
}
`,
		i.compartmentID, displayName, i.nodeCount, i.nodeMemoryInGBs, i.softwareVersion, i.subnetID, i.clusterMode, i.configSetID, primaryClusterConfig,
	)
}

func differentPrimaryClusterID(primaryClusterID string) string {
	if len(primaryClusterID) == 0 {
		return "different-primary-cluster-id"
	}
	if primaryClusterID[len(primaryClusterID)-1] == 'a' {
		return primaryClusterID[:len(primaryClusterID)-1] + "b"
	}
	return primaryClusterID[:len(primaryClusterID)-1] + "a"
}
