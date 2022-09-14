package objectstorage

import (
	"fmt"
	"strings"
	"time"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportObjectStorageBucketHints.GetIdFn = getObjectStorageBucketId
	exportObjectStorageObjectLifecyclePolicyHints.GetIdFn = getObjectStorageObjectLifecyclePolicyId
	exportObjectStorageObjectHints.GetIdFn = getObjectStorageObjectId
	exportObjectStoragePreauthenticatedRequestHints.GetIdFn = getObjectStoragePreauthenticatedRequestId
	exportObjectStorageReplicationPolicyHints.GetIdFn = getObjectStorageReplicationPolicyId
	exportObjectStorageNamespaceHints.IsDataSource = true
	exportObjectStorageNamespaceHints.ProcessDiscoveredResourcesFn = processObjectStorageNamespace
	exportObjectStorageNamespaceHints.GetHCLStringOverrideFn = getObjectStorageNamespaceHCLDatasource
	exportObjectStorageNamespaceHints.AlwaysExportable = true
	exportObjectStorageObjectHints.RequireResourceRefresh = true
	exportObjectStoragePreauthenticatedRequestHints.ProcessDiscoveredResourcesFn = processObjectStoragePreauthenticatedRequest
	exportObjectStorageReplicationPolicyHints.ProcessDiscoveredResourcesFn = processObjectStorageReplicationPolicy
	tf_export.RegisterCompartmentGraphs("object_storage", objectStorageResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework
func processObjectStorageNamespace(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, ns := range resources {
		namespaceName, ok := ns.SourceAttributes["namespace"].(string)
		if !ok || namespaceName == "" {
			return resources, fmt.Errorf("[ERROR] object storage namespace data source has no name\n")
		}
		tf_export.RefMapLock.Lock()
		tf_export.ReferenceMap[namespaceName] = tf_export.TfHclVersionvar.GetDataSourceHclString(ns.GetTerraformReference(), "namespace")
		tf_export.RefMapLock.Unlock()
	}

	return resources, nil
}

func getObjectStorageNamespaceHCLDatasource(builder *strings.Builder, ociRes *tf_export.OCIResource, varMap map[string]string) error {
	builder.WriteString(fmt.Sprintf("data %s %s {\n", ociRes.TerraformClass, ociRes.TerraformName))
	builder.WriteString(fmt.Sprintf("compartment_id = %v\n", varMap[ociRes.CompartmentId]))
	builder.WriteString("}\n")
	return nil
}

func processObjectStoragePreauthenticatedRequest(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, resource := range resources {
		if resource.Parent == nil {
			continue
		}
		resource.SourceAttributes["bucket"] = resource.Parent.SourceAttributes["name"].(string)
		resource.SourceAttributes["namespace"] = resource.Parent.SourceAttributes["namespace"].(string)

		// Check if time is already in RFC3339Nano format
		timeExpires, err := time.Parse(time.RFC3339Nano, resource.SourceAttributes["time_expires"].(string))
		if err != nil {
			// parse time using format in time.String()
			timeExpires, err = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", resource.SourceAttributes["time_expires"].(string))
			if err != nil {
				return resources, err
			}
			// Format to RFC3339Nano
			resource.SourceAttributes["time_expires"] = timeExpires.Format(time.RFC3339Nano)
		}

	}
	return resources, nil
}

func processObjectStorageReplicationPolicy(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, resource := range resources {
		if resource.Parent == nil {
			continue
		}
		resource.SourceAttributes["bucket"] = resource.Parent.SourceAttributes["name"].(string)
		resource.SourceAttributes["namespace"] = resource.Parent.SourceAttributes["namespace"].(string)
	}
	return resources, nil
}

func getObjectStorageBucketId(resource *tf_export.OCIResource) (string, error) {

	bucket, ok := resource.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find bucket for ObjectStorage Bucket")
	}
	namespace, ok := resource.SourceAttributes["namespace"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find namespace for ObjectStorage Bucket")
	}
	return GetBucketCompositeId(bucket, namespace), nil
}

func getObjectStorageObjectLifecyclePolicyId(resource *tf_export.OCIResource) (string, error) {

	bucket, ok := resource.SourceAttributes["bucket"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find bucket for ObjectStorage ObjectLifecyclePolicy")
	}
	namespace, ok := resource.SourceAttributes["namespace"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find namespace for ObjectStorage ObjectLifecyclePolicy")
	}
	return GetObjectLifecyclePolicyCompositeId(bucket, namespace), nil
}

func getObjectStorageObjectId(resource *tf_export.OCIResource) (string, error) {

	bucket, ok := resource.Parent.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find bucket for ObjectStorage Object")
	}
	namespace, ok := resource.Parent.SourceAttributes["namespace"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find namespace for ObjectStorage Object")
	}
	object, ok := resource.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find object for ObjectStorage Object")
	}
	return GetObjectCompositeId(bucket, namespace, object), nil
}

func getObjectStoragePreauthenticatedRequestId(resource *tf_export.OCIResource) (string, error) {

	bucket, ok := resource.Parent.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find bucket for ObjectStorage PreauthenticatedRequest")
	}
	namespace, ok := resource.Parent.SourceAttributes["namespace"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find namespace for ObjectStorage PreauthenticatedRequest")
	}
	parId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find parId for ObjectStorage PreauthenticatedRequest")
	}
	return GetPreauthenticatedRequestCompositeId(bucket, namespace, parId), nil
}

func getObjectStorageReplicationPolicyId(resource *tf_export.OCIResource) (string, error) {

	bucket, ok := resource.Parent.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find bucket for ObjectStorage ReplicationPolicy")
	}
	namespace, ok := resource.Parent.SourceAttributes["namespace"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find namespace for ObjectStorage ReplicationPolicy")
	}
	replicationId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find replicationId for ObjectStorage ReplicationPolicy")
	}
	return GetReplicationPolicyCompositeId(bucket, namespace, replicationId), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportObjectStorageBucketHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_objectstorage_bucket",
	DatasourceClass:        "oci_objectstorage_bucket_summaries",
	DatasourceItemsAttr:    "bucket_summaries",
	ResourceAbbreviation:   "bucket",
	RequireResourceRefresh: true,
}

var exportObjectStorageObjectLifecyclePolicyHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_objectstorage_object_lifecycle_policy",
	DatasourceClass:      "oci_objectstorage_object_lifecycle_policy",
	ResourceAbbreviation: "object_lifecycle_policy",
}

var exportObjectStorageNamespaceHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_objectstorage_namespace",
	DatasourceClass:      "oci_objectstorage_namespace",
	ResourceAbbreviation: "namespace",
}

var exportObjectStorageObjectHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_objectstorage_object",
	DatasourceClass:      "oci_objectstorage_objects",
	DatasourceItemsAttr:  "objects",
	ResourceAbbreviation: "object",
}

var exportObjectStoragePreauthenticatedRequestHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_objectstorage_preauthrequest",
	DatasourceClass:      "oci_objectstorage_preauthrequests",
	DatasourceItemsAttr:  "preauthenticated_requests",
	ResourceAbbreviation: "preauthenticated_request",
}

var exportObjectStorageReplicationPolicyHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_objectstorage_replication_policy",
	DatasourceClass:      "oci_objectstorage_replication_policies",
	DatasourceItemsAttr:  "replication_policies",
	ResourceAbbreviation: "replication_policy",
}

var objectStorageResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportObjectStorageNamespaceHints},
	},
	"oci_objectstorage_bucket": {
		{
			TerraformResourceHints: exportObjectStorageObjectHints,
			DatasourceQueryParams: map[string]string{
				"bucket":    "name",
				"namespace": "namespace",
			},
		},
		{
			TerraformResourceHints: exportObjectStorageObjectLifecyclePolicyHints,
			DatasourceQueryParams: map[string]string{
				"namespace": "namespace",
				"bucket":    "name",
			},
		},
		{
			TerraformResourceHints: exportObjectStoragePreauthenticatedRequestHints,
			DatasourceQueryParams: map[string]string{
				"namespace": "namespace",
				"bucket":    "name",
			},
		},
		{
			TerraformResourceHints: exportObjectStorageReplicationPolicyHints,
			DatasourceQueryParams: map[string]string{
				"namespace": "namespace",
				"bucket":    "name",
			},
		},
	},
	"oci_objectstorage_namespace": {
		{
			TerraformResourceHints: exportObjectStorageBucketHints,
			DatasourceQueryParams: map[string]string{
				"namespace": "namespace",
			},
		},
	},
}
