// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	publicationPackageSingularDataSourceRepresentation = map[string]interface{}{
		"package_version": acctest.Representation{RepType: acctest.Required, Create: `packageVersion`},
		"publication_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_marketplace_publication.test_publication.id}`},
	}

	publicationPackageDataSourceRepresentation = map[string]interface{}{
		"publication_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_marketplace_publication.test_publication.id}`},
		"package_type":    acctest.Representation{RepType: acctest.Optional, Create: `packageType`},
		"package_version": acctest.Representation{RepType: acctest.Optional, Create: `packageVersion`},
	}

	PublicationPackageResourceConfig = PublicationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_marketplace_publication", "test_publication", acctest.Optional, acctest.Create, publicationRepresentation)
)

// issue-routing-tag: marketplace/default
func TestMarketplacePublicationPackageResource_basic(t *testing.T) {
	t.Skip("Skip this test till Marketplace automates background processes and reduces the turnaround time.")
	httpreplay.SetScenario("TestMarketplacePublicationPackageResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_marketplace_publication_packages.test_publication_packages"
	singularDatasourceName := "data.oci_marketplace_publication_package.test_publication_package"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_marketplace_publication_packages", "test_publication_packages", acctest.Required, acctest.Create, publicationPackageDataSourceRepresentation) +
				compartmentIdVariableStr + PublicationPackageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "publication_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "publication_packages.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "publication_packages.0.listing_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "publication_packages.0.package_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "publication_packages.0.resource_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "publication_packages.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_marketplace_publication_package", "test_publication_package", acctest.Required, acctest.Create, publicationPackageSingularDataSourceRepresentation) +
				compartmentIdVariableStr + PublicationPackageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "package_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "publication_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "image_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "listing_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "operating_system.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "package_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "version"),
			),
		},
	})
}
