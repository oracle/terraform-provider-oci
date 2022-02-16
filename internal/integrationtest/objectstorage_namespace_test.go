// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	oci_object_storage "github.com/oracle/oci-go-sdk/v58/objectstorage"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	namespaceSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}

	NamespaceResourceConfig = ""
)

// issue-routing-tag: object_storage/default
func TestObjectStorageNamespaceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestObjectStorageNamespaceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_objectstorage_namespace.test_namespace"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Optional, acctest.Create, namespaceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + NamespaceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
			),
		},
	})
}

func getNamespaces(compartment string) ([]string, error) {
	var resourceIds []string
	compartmentId := compartment
	objectStorageClient := acctest.GetTestClients(&schema.ResourceData{}).ObjectStorageClient()

	getNamespacesRequest := oci_object_storage.GetNamespaceRequest{}
	getNamespacesResponse, err := objectStorageClient.GetNamespace(context.Background(), getNamespacesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Bucket NameSpace list for compartment id : %s , %s \n", compartmentId, err)
	}

	resourceIds = append(resourceIds, *getNamespacesResponse.Value)
	return resourceIds, nil
}
