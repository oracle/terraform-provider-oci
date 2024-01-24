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
	ServiceCatalogServiceCatalogPrivateApplicationPackageSingularDataSourceRepresentation = map[string]interface{}{
		"private_application_package_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_service_catalog_private_application_packages.test_private_application_packages.private_application_package_collection.0.items.0.id}`},
	}

	ServiceCatalogServiceCatalogPrivateApplicationPackageDataSourceRepresentation = map[string]interface{}{
		"private_application_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_service_catalog_private_application.test_private_application.id}`},
		"display_name":                   acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"package_type":                   acctest.Representation{RepType: acctest.Optional, Create: []string{`packageType`}},
		"private_application_package_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_service_catalog_private_application_package.test_private_application_package.id}`},
	}

	ServiceCatalogPrivateApplicationPackageResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_service_catalog_private_application_packages", "test_private_application_packages", acctest.Required, acctest.Create, ServiceCatalogServiceCatalogPrivateApplicationPackageDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_service_catalog_private_application", "test_private_application", acctest.Required, acctest.Create, ServiceCatalogPrivateApplicationRepresentation)
)

// issue-routing-tag: service_catalog/default
func TestServiceCatalogPrivateApplicationPackageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestServiceCatalogPrivateApplicationPackageResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_service_catalog_private_application_packages.test_private_application_packages"
	singularDatasourceName := "data.oci_service_catalog_private_application_package.test_private_application_package"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + ServiceCatalogPrivateApplicationPackageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "private_application_package_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "private_application_package_collection.0.items.0.display_name", "version"),
				resource.TestCheckResourceAttrSet(datasourceName, "private_application_package_collection.0.items.0.package_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "private_application_package_collection.0.items.0.private_application_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "private_application_package_collection.0.items.0.id"),

				resource.TestCheckResourceAttrSet(datasourceName, "private_application_package_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_service_catalog_private_application_package", "test_private_application_package", acctest.Required, acctest.Create, ServiceCatalogServiceCatalogPrivateApplicationPackageSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ServiceCatalogPrivateApplicationPackageResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_application_package_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "content_url"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "mime_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "package_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "version"),
			),
		},
	})
}
