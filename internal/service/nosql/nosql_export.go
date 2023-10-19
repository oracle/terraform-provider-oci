package nosql

import (
	"fmt"

	oci_nosql "github.com/oracle/oci-go-sdk/v65/nosql"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportNosqlIndexHints.GetIdFn = getNosqlIndexId
	exportNosqlIndexHints.ProcessDiscoveredResourcesFn = processNosqlIndex
	tf_export.RegisterCompartmentGraphs("nosql", nosqlResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func processNosqlIndex(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, index := range resources {
		if index.Parent == nil {
			continue
		}
		index.SourceAttributes["table_name_or_id"] = index.Parent.Id
	}
	return resources, nil
}

func getNosqlIndexId(resource *tf_export.OCIResource) (string, error) {
	indexName, ok := resource.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find indexName for Nosql Index")
	}
	tableNameOrId := resource.Parent.Id
	return GetIndexCompositeId(indexName, tableNameOrId), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportNosqlTableHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_nosql_table",
	DatasourceClass:        "oci_nosql_tables",
	DatasourceItemsAttr:    "table_collection",
	ResourceAbbreviation:   "table",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_nosql.TableLifecycleStateActive),
	},
}

var exportNosqlIndexHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_nosql_index",
	DatasourceClass:        "oci_nosql_indexes",
	DatasourceItemsAttr:    "index_collection",
	ResourceAbbreviation:   "index",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_nosql.IndexLifecycleStateActive),
	},
}

var exportNosqlTableReplicaHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_nosql_table_replica",
	ResourceAbbreviation: "table_replica",
}

var nosqlResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportNosqlTableHints},
	},
	"oci_nosql_table": {
		{
			TerraformResourceHints: exportNosqlIndexHints,
			DatasourceQueryParams: map[string]string{
				"table_name_or_id": "id",
			},
		},
	},
}
