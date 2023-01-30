package log_analytics

import (
	"context"
	"fmt"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"
	oci_objectstorage "github.com/oracle/oci-go-sdk/v65/objectstorage"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportLogAnalyticsLogAnalyticsObjectCollectionRuleHints.GetIdFn = getLogAnalyticsLogAnalyticsObjectCollectionRuleId
	exportLogAnalyticsNamespaceScheduledTaskHints.GetIdFn = getLogAnalyticsNamespaceScheduledTaskId
	exportLogAnalyticsNamespaceIngestTimeRuleHints.GetIdFn = getLogAnalyticsNamespaceIngestTimeRuleId
	exportLogAnalyticsLogAnalyticsObjectCollectionRuleHints.FindResourcesOverrideFn = findLogAnalyticsObjectCollectionRules
	exportLogAnalyticsLogAnalyticsObjectCollectionRuleHints.ProcessDiscoveredResourcesFn = processLogAnalyticsObjectCollectionRules
	tf_export.RegisterCompartmentGraphs("log_analytics", logAnalyticsResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework
func processLogAnalyticsObjectCollectionRules(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, resource := range resources {
		namespace := resource.SourceAttributes["namespace"].(string)
		logAnalyticsObjectCollectionRuleId := resource.Id
		resource.ImportId = GetLogAnalyticsObjectCollectionRuleCompositeId(logAnalyticsObjectCollectionRuleId, namespace)
	}

	return resources, nil
}

func findLogAnalyticsObjectCollectionRules(ctx *tf_export.ResourceDiscoveryContext, tfMeta *tf_export.TerraformResourceAssociation, parent *tf_export.OCIResource, resourceGraph *tf_export.TerraformResourceGraph) ([]*tf_export.OCIResource, error) {
	// List on LogAnalyticsObjectCollectionRules requires namespaceName path parameter.
	// Getting namespace from ObjectStorage.GetNamespace API before calling ListLogAnalyticsObjectCollectionRules API.
	results := []*tf_export.OCIResource{}

	namespaceRequest := oci_objectstorage.GetNamespaceRequest{}
	namespaceResponse, err := ctx.Clients.ObjectStorageClient().GetNamespace(context.Background(), namespaceRequest)
	if err != nil {
		return results, err
	}
	namespace := namespaceResponse.Value
	request := oci_log_analytics.ListLogAnalyticsObjectCollectionRulesRequest{}

	request.NamespaceName = namespace
	request.CompartmentId = ctx.CompartmentId
	request.LifecycleState = oci_log_analytics.ListLogAnalyticsObjectCollectionRulesLifecycleStateActive

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "log_analytics")

	response, err := ctx.Clients.LogAnalyticsClient().ListLogAnalyticsObjectCollectionRules(context.Background(), request)
	if err != nil {
		return results, err
	}

	request.Page = response.OpcNextPage

	for request.Page != nil {
		listResponse, err := ctx.Clients.LogAnalyticsClient().ListLogAnalyticsObjectCollectionRules(context.Background(), request)
		if err != nil {
			return results, err
		}

		response.Items = append(response.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	for _, logAnalyticsObjectCollectionRule := range response.Items {
		logAnalyticsObjectCollectionRuleResource := tf_export.ResourcesMap[tfMeta.ResourceClass]

		d := logAnalyticsObjectCollectionRuleResource.TestResourceData()
		d.SetId(GetLogAnalyticsObjectCollectionRuleCompositeId(*logAnalyticsObjectCollectionRule.Id, *namespace))

		if err := logAnalyticsObjectCollectionRuleResource.Read(d, ctx.Clients); err != nil {
			rdError := &tf_export.ResourceDiscoveryError{ResourceType: tfMeta.ResourceClass, ParentResource: parent.TerraformName, Error: err, ResourceGraph: resourceGraph}
			ctx.AddErrorToList(rdError)
			continue
		}

		resource := &tf_export.OCIResource{
			CompartmentId:    *ctx.CompartmentId,
			SourceAttributes: tf_export.ConvertResourceDataToMap(logAnalyticsObjectCollectionRuleResource.Schema, d),
			RawResource:      logAnalyticsObjectCollectionRule,
			TerraformResource: tf_export.TerraformResource{
				Id:             d.Id(),
				TerraformClass: tfMeta.ResourceClass,
			},
			GetHclStringFn: tf_export.GetHclStringFromGenericMap,
			Parent:         parent,
		}

		if resource.TerraformName, err = tf_export.GenerateTerraformNameFromResource(resource.SourceAttributes, logAnalyticsObjectCollectionRuleResource.Schema); err != nil {
			resource.TerraformName = fmt.Sprintf("%s_%s", parent.Parent.TerraformName, *logAnalyticsObjectCollectionRule.Name)
		}

		results = append(results, resource)
	}

	return results, nil

}

func getLogAnalyticsLogAnalyticsObjectCollectionRuleId(resource *tf_export.OCIResource) (string, error) {

	logAnalyticsObjectCollectionRuleId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find logAnalyticsObjectCollectionRuleId for LogAnalytics LogAnalyticsObjectCollectionRule")
	}
	namespace, ok := resource.SourceAttributes["namespace"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find namespace for LogAnalytics LogAnalyticsObjectCollectionRule")
	}
	return GetLogAnalyticsObjectCollectionRuleCompositeId(logAnalyticsObjectCollectionRuleId, namespace), nil
}

func getLogAnalyticsNamespaceScheduledTaskId(resource *tf_export.OCIResource) (string, error) {

	namespace, ok := resource.SourceAttributes["namespace"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find namespace for LogAnalytics NamespaceScheduledTask")
	}
	scheduledTaskId, ok := resource.SourceAttributes["scheduled_task_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find scheduledTaskId for LogAnalytics NamespaceScheduledTask")
	}
	return GetNamespaceScheduledTaskCompositeId(namespace, scheduledTaskId), nil
}

func getLogAnalyticsNamespaceIngestTimeRuleId(resource *tf_export.OCIResource) (string, error) {

	ingestTimeRuleId, ok := resource.SourceAttributes["ingest_time_rule_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find ingestTimeRuleId for LogAnalytics NamespaceIngestTimeRule")
	}
	namespace, ok := resource.SourceAttributes["namespace"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find namespace for LogAnalytics NamespaceIngestTimeRule")
	}
	return GetNamespaceIngestTimeRuleCompositeId(ingestTimeRuleId, namespace), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportLogAnalyticsLogAnalyticsObjectCollectionRuleHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_log_analytics_log_analytics_object_collection_rule",
	DatasourceClass:        "oci_log_analytics_log_analytics_object_collection_rules",
	DatasourceItemsAttr:    "log_analytics_object_collection_rule_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "log_analytics_object_collection_rule",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_log_analytics.ObjectCollectionRuleLifecycleStatesActive),
	},
}

var exportLogAnalyticsLogAnalyticsImportCustomContentHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_log_analytics_log_analytics_import_custom_content",
	ResourceAbbreviation: "log_analytics_import_custom_content",
}

var exportLogAnalyticsNamespaceScheduledTaskHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_log_analytics_namespace_scheduled_task",
	DatasourceClass:        "oci_log_analytics_namespace_scheduled_tasks",
	DatasourceItemsAttr:    "scheduled_task_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "namespace_scheduled_task",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_log_analytics.ScheduledTaskLifecycleStateActive),
	},
}

var exportLogAnalyticsLogAnalyticsPreferencesManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_log_analytics_log_analytics_preferences_management",
	ResourceAbbreviation: "log_analytics_preferences_management",
}

var exportLogAnalyticsLogAnalyticsUnprocessedDataBucketManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_log_analytics_log_analytics_unprocessed_data_bucket_management",
	ResourceAbbreviation: "log_analytics_unprocessed_data_bucket_management",
}

var exportLogAnalyticsLogAnalyticsResourceCategoriesManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_log_analytics_log_analytics_resource_categories_management",
	ResourceAbbreviation: "log_analytics_resource_categories_management",
}

var exportLogAnalyticsNamespaceIngestTimeRuleHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_log_analytics_namespace_ingest_time_rule",
	DatasourceClass:        "oci_log_analytics_namespace_ingest_time_rules",
	DatasourceItemsAttr:    "ingest_time_rule_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "namespace_ingest_time_rule",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_log_analytics.ConfigLifecycleStateActive),
	},
}

var logAnalyticsResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportLogAnalyticsLogAnalyticsObjectCollectionRuleHints},
	},
}
