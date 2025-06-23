// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	RedisOciCacheDefaultConfigSetSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"oci_cache_default_config_set_id": acctest.Representation{RepType: acctest.Required, Create: `${var.default_config_set_id}`},
	}

	RedisOciCacheDefaultConfigSetDataSourceRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `default-redis_7_0-v1`},
		"id":               acctest.Representation{RepType: acctest.Optional, Create: `${var.default_config_set_id}`},
		"software_version": acctest.Representation{RepType: acctest.Optional, Create: `REDIS_7_0`},
		"state":            acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":           acctest.RepresentationGroup{RepType: acctest.Required, Group: RedisOciCacheDefaultConfigSetDataSourceFilterRepresentation}}
	RedisOciCacheDefaultConfigSetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `items.id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${var.default_config_set_id}`}},
	}
	RedisOciCacheDefaultConfigSetResourceConfig = ""
)

// issue-routing-tag: redis/default
func TestRedisOciCacheDefaultConfigSetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestRedisOciCacheDefaultConfigSetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_redis_oci_cache_default_config_sets.test_oci_cache_default_config_sets"
	singularDatasourceName := "data.oci_redis_oci_cache_default_config_set.test_oci_cache_default_config_set"
	defaultConfigSetId := utils.GetEnvSettingWithBlankDefault("default_config_set_id")
	defaultConfigSetIdVariableStr := fmt.Sprintf("variable \"default_config_set_id\" { default = \"%s\" }\n", defaultConfigSetId)

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_redis_oci_cache_default_config_sets", "test_oci_cache_default_config_sets", acctest.Optional, acctest.Create, RedisOciCacheDefaultConfigSetDataSourceRepresentation) +
				compartmentIdVariableStr + defaultConfigSetIdVariableStr + RedisOciCacheDefaultConfigSetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "default-redis_7_0-v1"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "software_version", "REDIS_7_0"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttrSet(datasourceName, "oci_cache_default_config_set_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_redis_oci_cache_default_config_set", "test_oci_cache_default_config_set", acctest.Required, acctest.Create, RedisOciCacheDefaultConfigSetSingularDataSourceRepresentation) +
				compartmentIdVariableStr + defaultConfigSetIdVariableStr + RedisOciCacheDefaultConfigSetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "oci_cache_default_config_set_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "default_configuration_details.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "software_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
	})
}
