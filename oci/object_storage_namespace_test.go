// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	oci_object_storage "github.com/oracle/oci-go-sdk/objectstorage"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	namespaceSingularDataSourceRepresentation = map[string]interface{}{}

	NamespaceResourceConfig = ""
)

func TestObjectStorageNamespaceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestObjectStorageNamespaceResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation) +
					NamespaceResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(),
			},
		},
	})
}

func getNamespaces(compartment string) ([]string, error) {
	var resourceIds []string
	compartmentId := compartment
	objectStorageClient := GetTestClients(&schema.ResourceData{}).objectStorageClient

	getNamespacesRequest := oci_object_storage.GetNamespaceRequest{}
	getNamespacesResponse, err := objectStorageClient.GetNamespace(context.Background(), getNamespacesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Bucket NameSpace list for compartment id : %s , %s \n", compartmentId, err)
	}

	resourceIds = append(resourceIds, *getNamespacesResponse.Value)
	return resourceIds, nil
}
