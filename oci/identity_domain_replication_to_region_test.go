// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DomainReplicationToRegionRequiredOnlyResource = DomainReplicationToRegionResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_identity_domain_replication_to_region", "test_domain_replication_to_region", Required, Create, domainReplicationToRegionRepresentation)

	domainReplicationToRegionRepresentation = map[string]interface{}{
		"domain_id":      Representation{RepType: Required, Create: `${oci_identity_domain.test_domain.id}`},
		"replica_region": Representation{RepType: Required, Create: `us-sanjose-1`},
	}

	DomainReplicationToRegionResourceDependencies = GenerateResourceFromRepresentationMap("oci_identity_domain", "test_domain", Required, Create, domainRepresentation)
)

// issue-routing-tag: identity/default
func TestIdentityDomainReplicationToRegionResource_basic(t *testing.T) {
	t.Skip("Skip this test because henosis tenancy is needed")
	httpreplay.SetScenario("TestIdentityDomainReplicationToRegionResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_identity_domain_replication_to_region.test_domain_replication_to_region"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+DomainReplicationToRegionResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_identity_domain_replication_to_region", "test_domain_replication_to_region", Optional, Create, domainReplicationToRegionRepresentation), "identity", "domainReplicationToRegion", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DomainReplicationToRegionResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_identity_domain_replication_to_region", "test_domain_replication_to_region", Required, Create, domainReplicationToRegionRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "domain_id"),
			),
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr +
				GenerateResourceFromRepresentationMap("oci_identity_domain", "test_domain", Required, Create,
					RepresentationCopyWithNewProperties(domainRepresentation, map[string]interface{}{
						"state": Representation{RepType: Required, Create: `inactive`},
					})),
		},
	})

}
