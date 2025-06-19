// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	RedisRedisClusterCreateIdentityTokenRequiredOnlyResource = RedisRedisClusterCreateIdentityTokenResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_redis_redis_cluster_create_identity_token", "test_redis_cluster_create_identity_token", acctest.Required, acctest.Create, RedisRedisClusterCreateIdentityTokenRepresentation)

	RedisRedisClusterCreateIdentityTokenRepresentation = map[string]interface{}{
		"public_key":       acctest.Representation{RepType: acctest.Required, Create: `LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUFycHJDUWl5RVpwWFlaY044c0RGSAp4QVlYZ0hXRFd1NTF2ejBFWDg5dzMzaDgwNllQbUVBcFFjYzRETjJFbGNvcWJ0K1djb3lZWGdMemVmWGswcnVICjI1TDR4alZPQXJWQ2FwbVRubXBUVFY4RW9zYVorUjdBbmFkZC9LYjVHa21NV2ZZbmwvT05hT3RPQ1ZBMVZzOWsKTXZoZTFJZEJZNWM2bTlHYWxEanlaL2hOamlONVZyRkh1K0puajhodzNOSk9XSEtMNW9JT3NIU3A4MFozWmZHRApvUDFHREs3SVpXNUJ5bEhXSzhPU1pzYjJ6cDQ5d3JGVjZEWVJlR1F1RHFPbzhpR0lXUzJFQk9MU2xvMk9OU3RJCktySUJ3eTN5aElXK2hPVXVheWorRnc5OEpmdjVhYnpNK2hPa3BXUktqeWhsUGY0b05FK2J2bm9xT1NFdTIwUnoKR1FJREFRQUIKLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0tCg`},
		"redis_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `rediscluster.luiw7q`},
		"redis_user":       acctest.Representation{RepType: acctest.Required, Create: `OCI_REDIS_OWNER`},
		"defined_tags":     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
	}

	RedisRedisClusterCreateIdentityTokenResourceDependencies = ""

	// RedisRedisClusterCreateIdentityTokenResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
	//
	//	acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
	//	DefinedTagsDependencies +
	//	acctest.GenerateResourceFromRepresentationMap("oci_redis_redis_cluster", "test_redis_cluster", acctest.Required, acctest.Create, RedisRedisClusterRepresentation)
)

// issue-routing-tag: redis/default
func TestRedisRedisClusterCreateIdentityTokenResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestRedisRedisClusterCreateIdentityTokenResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_redis_redis_cluster_create_identity_token.test_redis_cluster_create_identity_token"

	// var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+RedisRedisClusterCreateIdentityTokenResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_redis_redis_cluster_create_identity_token", "test_redis_cluster_create_identity_token", acctest.Optional, acctest.Create, RedisRedisClusterCreateIdentityTokenRepresentation), "redis", "redisClusterCreateIdentityToken", t)
	log.Printf("[DEBUG] *** 0010 RedisRedisClusterCreateIdentityTokenRepresentation ***: %v", RedisRedisClusterCreateIdentityTokenRepresentation)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + RedisRedisClusterCreateIdentityTokenResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_redis_cluster_create_identity_token", "test_redis_cluster_create_identity_token", acctest.Required, acctest.Create, RedisRedisClusterCreateIdentityTokenRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "public_key", "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUFycHJDUWl5RVpwWFlaY044c0RGSAp4QVlYZ0hXRFd1NTF2ejBFWDg5dzMzaDgwNllQbUVBcFFjYzRETjJFbGNvcWJ0K1djb3lZWGdMemVmWGswcnVICjI1TDR4alZPQXJWQ2FwbVRubXBUVFY4RW9zYVorUjdBbmFkZC9LYjVHa21NV2ZZbmwvT05hT3RPQ1ZBMVZzOWsKTXZoZTFJZEJZNWM2bTlHYWxEanlaL2hOamlONVZyRkh1K0puajhodzNOSk9XSEtMNW9JT3NIU3A4MFozWmZHRApvUDFHREs3SVpXNUJ5bEhXSzhPU1pzYjJ6cDQ5d3JGVjZEWVJlR1F1RHFPbzhpR0lXUzJFQk9MU2xvMk9OU3RJCktySUJ3eTN5aElXK2hPVXVheWorRnc5OEpmdjVhYnpNK2hPa3BXUktqeWhsUGY0b05FK2J2bm9xT1NFdTIwUnoKR1FJREFRQUIKLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0tCg"),
				resource.TestCheckResourceAttrSet(resourceName, "redis_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "redis_user", "OCI_REDIS_OWNER"),
			),
		},
	})
}
