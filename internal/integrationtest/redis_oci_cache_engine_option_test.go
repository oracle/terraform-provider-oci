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
	RedisOciCacheEngineOptionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	RedisOciCacheEngineOptionResourceConfig = ""
)

// issue-routing-tag: redis/default
func TestRedisOciCacheEngineOptionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestRedisOciCacheEngineOptionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_redis_oci_cache_engine_options.test_oci_cache_engine_options"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_redis_oci_cache_engine_options", "test_oci_cache_engine_options", acctest.Required, acctest.Create, RedisOciCacheEngineOptionDataSourceRepresentation) +
				compartmentIdVariableStr + RedisOciCacheEngineOptionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "oci_cache_engine_options_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "oci_cache_engine_options_collection.0.items.#", "1"),
			),
		},
	})
}
