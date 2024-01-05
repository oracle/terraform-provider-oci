// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	LogAnalyticsNamespaceStorageEncryptionKeyInfoSingularDataSourceRepresentation = map[string]interface{}{
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	}

	LogAnalyticsNamespaceStorageEncryptionKeyInfoResourceConfig = ""
	LogAnalyticsNamespaceStorageEncryptionKeyInfoDependencies   = acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsNamespaceStorageEncryptionKeyInfoResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsNamespaceStorageEncryptionKeyInfoResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	singularDatasourceName := "data.oci_log_analytics_namespace_storage_encryption_key_info.test_namespace_storage_encryption_key_info"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace_storage_encryption_key_info", "test_namespace_storage_encryption_key_info", acctest.Required, acctest.Create, LogAnalyticsNamespaceStorageEncryptionKeyInfoSingularDataSourceRepresentation) +
				LogAnalyticsNamespaceStorageEncryptionKeyInfoDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "2"),
			),
		},
	})
}
