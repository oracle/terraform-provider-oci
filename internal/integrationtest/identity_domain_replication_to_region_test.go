// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	IdentityDomainReplicationToRegionRequiredOnlyResource = IdentityDomainReplicationToRegionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domain_replication_to_region", "test_domain_replication_to_region", acctest.Required, acctest.Create, IdentityDomainReplicationToRegionRepresentation)

	IdentityDomainReplicationToRegionRepresentation = map[string]interface{}{
		"domain_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_identity_domain.test_domain.id}`},
		"replica_region": acctest.Representation{RepType: acctest.Required, Create: `us-sanjose-1`},
	}

	IdentityDomainReplicationToRegionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_identity_domain", "test_domain", acctest.Required, acctest.Create, IdentityDomainRepresentation)
)

// issue-routing-tag: identity/default
func TestIdentityDomainReplicationToRegionResource_basic(t *testing.T) {
	t.Skip("Skip this test because henosis tenancy is needed")
	httpreplay.SetScenario("TestIdentityDomainReplicationToRegionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domain_replication_to_region.test_domain_replication_to_region"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+IdentityDomainReplicationToRegionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_domain_replication_to_region", "test_domain_replication_to_region", acctest.Optional, acctest.Create, IdentityDomainReplicationToRegionRepresentation), "identity", "domainReplicationToRegion", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + IdentityDomainReplicationToRegionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domain_replication_to_region", "test_domain_replication_to_region", acctest.Required, acctest.Create, IdentityDomainReplicationToRegionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "domain_id"),
			),
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_domain", "test_domain", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(IdentityDomainRepresentation, map[string]interface{}{
						"state": acctest.Representation{RepType: acctest.Required, Create: `inactive`},
					})),
		},
	})

}
