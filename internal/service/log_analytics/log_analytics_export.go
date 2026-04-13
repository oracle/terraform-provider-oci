package log_analytics

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"
	oci_objectstorage "github.com/oracle/oci-go-sdk/v65/objectstorage"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportLogAnalyticsLogAnalyticsObjectCollectionRuleHints.GetIdFn = getLogAnalyticsLogAnalyticsObjectCollectionRuleId
	exportLogAnalyticsLogAnalyticsLogGroupHints.GetIdFn = getLogAnalyticsLogAnalyticsLogGroupId
	exportLogAnalyticsLogAnalyticsPreferencesManagementHints.GetIdFn = getLogAnalyticsLogAnalyticsPreferencesManagementId
	exportLogAnalyticsNamespaceScheduledTaskHints.GetIdFn = getLogAnalyticsNamespaceScheduledTaskId
	exportLogAnalyticsNamespaceIngestTimeRuleHints.GetIdFn = getLogAnalyticsNamespaceIngestTimeRuleId
	exportLogAnalyticsNamespaceStorageArchivalConfigHints.GetIdFn = getLogAnalyticsNamespaceStorageArchivalConfigId
	exportLogAnalyticsLogAnalyticsObjectCollectionRuleHints.FindResourcesOverrideFn = findLogAnalyticsObjectCollectionRules
	exportLogAnalyticsLogAnalyticsResourceCategoriesManagementHints.FindResourcesOverrideFn = findLogAnalyticsLogAnalyticsResourceCategoriesManagement
	exportLogAnalyticsNamespaceLookupHints.FindResourcesOverrideFn = findLogAnalyticsNamespaceLookups
	exportLogAnalyticsNamespaceAssociationHints.FindResourcesOverrideFn = findLogAnalyticsNamespaceAssociations
	exportLogAnalyticsLogAnalyticsObjectCollectionRuleHints.ProcessDiscoveredResourcesFn = processLogAnalyticsObjectCollectionRules
	tf_export.RegisterCompartmentGraphs("log_analytics", logAnalyticsResourceGraph)
	tf_export.RegisterTenancyGraphs("log_analytics_tenancy", logAnalyticsTenancyResourceGraph)
}

var logAnalyticsTerraformNameSanitizer = regexp.MustCompile(`[^a-zA-Z0-9\-_]+`)

func readLogAnalyticsResourceForExport(resource *schema.Resource, d *schema.ResourceData, clients interface{}) error {
	if resource.ReadContext != nil {
		if diags := resource.ReadContext(context.Background(), d, clients); diags.HasError() {
			return fmt.Errorf("%s", strings.Join(tf_export.ParseDiagToError(diags), " | "))
		}
		return nil
	}

	if resource.Read != nil {
		return resource.Read(d, clients)
	}

	return fmt.Errorf("[ERROR] resource has neither ReadContext nor Read implemented")
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
			resource.TerraformName = tf_export.CheckDuplicateResourceName(resource.TerraformName)
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

	ingestTimeRuleId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find ingestTimeRuleId for LogAnalytics NamespaceIngestTimeRule")
	}
	namespace, ok := resource.SourceAttributes["namespace"].(string)
	if (!ok || namespace == "") && resource.Parent != nil {
		namespace, ok = resource.Parent.SourceAttributes["namespace"].(string)
	}
	if !ok || namespace == "" {
		return "", fmt.Errorf("[ERROR] unable to find namespace for LogAnalytics NamespaceIngestTimeRule")
	}
	return GetNamespaceIngestTimeRuleCompositeId(ingestTimeRuleId, namespace), nil
}

func findLogAnalyticsNamespaceLookups(ctx *tf_export.ResourceDiscoveryContext, tfMeta *tf_export.TerraformResourceAssociation, parent *tf_export.OCIResource, resourceGraph *tf_export.TerraformResourceGraph) ([]*tf_export.OCIResource, error) {
	results := []*tf_export.OCIResource{}

	namespace, ok := parent.SourceAttributes["namespace"].(string)
	if !ok || namespace == "" {
		return results, fmt.Errorf("[ERROR] unable to find namespace for LogAnalytics NamespaceLookup")
	}

	request := oci_log_analytics.ListLookupsRequest{}
	request.NamespaceName = &namespace
	request.CompartmentId = ctx.CompartmentId
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "log_analytics")

	response, err := ctx.Clients.LogAnalyticsClient().ListLookups(context.Background(), request)
	if err != nil {
		return results, err
	}

	request.Page = response.OpcNextPage

	for request.Page != nil {
		listResponse, err := ctx.Clients.LogAnalyticsClient().ListLookups(context.Background(), request)
		if err != nil {
			return results, err
		}

		response.Items = append(response.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	for _, namespaceLookup := range response.Items {
		if namespaceLookup.Name == nil || *namespaceLookup.Name == "" {
			continue
		}

		namespaceLookupResource := tf_export.ResourcesMap[tfMeta.ResourceClass]
		d := namespaceLookupResource.TestResourceData()
		d.SetId(GetNamespaceLookupCompositeId(*namespaceLookup.Name, namespace))

		if err := readLogAnalyticsResourceForExport(namespaceLookupResource, d, ctx.Clients); err != nil {
			rdError := &tf_export.ResourceDiscoveryError{ResourceType: tfMeta.ResourceClass, ParentResource: parent.TerraformName, Error: err, ResourceGraph: resourceGraph}
			ctx.AddErrorToList(rdError)
			continue
		}

		resource := &tf_export.OCIResource{
			CompartmentId:    *ctx.CompartmentId,
			SourceAttributes: tf_export.ConvertResourceDataToMap(namespaceLookupResource.Schema, d),
			RawResource:      namespaceLookup,
			TerraformResource: tf_export.TerraformResource{
				Id:             d.Id(),
				TerraformClass: tfMeta.ResourceClass,
			},
			GetHclStringFn: tf_export.GetHclStringFromGenericMap,
			Parent:         parent,
		}

		if resource.TerraformName, err = tf_export.GenerateTerraformNameFromResource(resource.SourceAttributes, namespaceLookupResource.Schema); err != nil {
			resource.TerraformName = fmt.Sprintf("%s_%s", parent.TerraformName, *namespaceLookup.Name)
			resource.TerraformName = tf_export.CheckDuplicateResourceName(resource.TerraformName)
		}

		results = append(results, resource)
	}

	return results, nil
}

func getLogAnalyticsNamespaceAssociationStringAttribute(sourceAttributes map[string]interface{}, keys ...string) string {
	for _, key := range keys {
		if value, ok := sourceAttributes[key].(string); ok && value != "" {
			return value
		}
	}

	return ""
}

func listLogAnalyticsEntitySourceAssociationsForExport(ctx *tf_export.ResourceDiscoveryContext, namespace string, compartmentId string) ([]oci_log_analytics.LogAnalyticsAssociation, error) {
	request := oci_log_analytics.ListEntitySourceAssociationsRequest{}
	request.NamespaceName = &namespace
	request.CompartmentId = &compartmentId
	request.LifeCycleState = oci_log_analytics.ListEntitySourceAssociationsLifeCycleStateAll
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "log_analytics")

	response, err := ctx.Clients.LogAnalyticsClient().ListEntitySourceAssociations(context.Background(), request)
	if err != nil {
		return nil, err
	}

	request.Page = response.OpcNextPage
	for request.Page != nil {
		listResponse, err := ctx.Clients.LogAnalyticsClient().ListEntitySourceAssociations(context.Background(), request)
		if err != nil {
			return nil, err
		}

		response.Items = append(response.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return response.Items, nil
}

func getLogAnalyticsNamespaceAssociationSourceAttributes(namespace string, compartmentId string, namespaceAssociation oci_log_analytics.LogAnalyticsAssociation) map[string]interface{} {
	sourceAttributes := map[string]interface{}{
		"namespace":      namespace,
		"compartment_id": compartmentId,
	}

	if namespaceAssociation.EntityId != nil && *namespaceAssociation.EntityId != "" {
		sourceAttributes["entity_id"] = *namespaceAssociation.EntityId
	}
	if namespaceAssociation.LogGroupId != nil && *namespaceAssociation.LogGroupId != "" {
		sourceAttributes["log_group_id"] = *namespaceAssociation.LogGroupId
	}
	if namespaceAssociation.SourceName != nil && *namespaceAssociation.SourceName != "" {
		sourceAttributes["source_name"] = *namespaceAssociation.SourceName
	}
	if namespaceAssociation.AgentId != nil && *namespaceAssociation.AgentId != "" {
		sourceAttributes["agent_id"] = *namespaceAssociation.AgentId
	}
	if namespaceAssociation.AgentEntityName != nil && *namespaceAssociation.AgentEntityName != "" {
		sourceAttributes["agent_entity_name"] = *namespaceAssociation.AgentEntityName
	}
	if namespaceAssociation.AssociationProperties != nil {
		associationProperties := make([]interface{}, 0, len(namespaceAssociation.AssociationProperties))
		for _, associationProperty := range namespaceAssociation.AssociationProperties {
			associationProperties = append(associationProperties, AssociationPropertyToMap(associationProperty))
		}
		sourceAttributes["association_properties"] = associationProperties
	}
	if namespaceAssociation.EntityName != nil && *namespaceAssociation.EntityName != "" {
		sourceAttributes["entity_name"] = *namespaceAssociation.EntityName
	}
	if namespaceAssociation.EntityTypeDisplayName != nil && *namespaceAssociation.EntityTypeDisplayName != "" {
		sourceAttributes["entity_type_display_name"] = *namespaceAssociation.EntityTypeDisplayName
	}
	if namespaceAssociation.EntityTypeName != nil && *namespaceAssociation.EntityTypeName != "" {
		sourceAttributes["entity_type_name"] = *namespaceAssociation.EntityTypeName
	}
	if namespaceAssociation.FailureMessage != nil && *namespaceAssociation.FailureMessage != "" {
		sourceAttributes["failure_message"] = *namespaceAssociation.FailureMessage
	}
	if namespaceAssociation.Host != nil && *namespaceAssociation.Host != "" {
		sourceAttributes["host"] = *namespaceAssociation.Host
	}
	if namespaceAssociation.LifeCycleState != "" {
		sourceAttributes["lifecycle_state"] = string(namespaceAssociation.LifeCycleState)
	}
	if namespaceAssociation.LogGroupCompartment != nil && *namespaceAssociation.LogGroupCompartment != "" {
		sourceAttributes["log_group_compartment"] = *namespaceAssociation.LogGroupCompartment
	}
	if namespaceAssociation.LogGroupName != nil && *namespaceAssociation.LogGroupName != "" {
		sourceAttributes["log_group_name"] = *namespaceAssociation.LogGroupName
	}
	if namespaceAssociation.RetryCount != nil {
		sourceAttributes["retry_count"] = int(*namespaceAssociation.RetryCount)
	}
	if namespaceAssociation.SourceDisplayName != nil && *namespaceAssociation.SourceDisplayName != "" {
		sourceAttributes["source_display_name"] = *namespaceAssociation.SourceDisplayName
	}
	if namespaceAssociation.SourceTypeName != nil && *namespaceAssociation.SourceTypeName != "" {
		sourceAttributes["source_type_name"] = *namespaceAssociation.SourceTypeName
	}
	if namespaceAssociation.TimeLastAttempted != nil {
		sourceAttributes["time_last_attempted"] = namespaceAssociation.TimeLastAttempted.String()
	}

	return sourceAttributes
}

func findLogAnalyticsNamespaceAssociations(ctx *tf_export.ResourceDiscoveryContext, tfMeta *tf_export.TerraformResourceAssociation, parent *tf_export.OCIResource, resourceGraph *tf_export.TerraformResourceGraph) ([]*tf_export.OCIResource, error) {
	results := []*tf_export.OCIResource{}

	namespace, ok := parent.SourceAttributes["namespace"].(string)
	if !ok || namespace == "" {
		return results, fmt.Errorf("[ERROR] unable to find namespace for LogAnalytics NamespaceAssociation")
	}

	compartmentId := parent.CompartmentId
	if ctx.CompartmentId != nil && *ctx.CompartmentId != "" {
		compartmentId = *ctx.CompartmentId
	}
	if compartmentId == "" {
		return results, fmt.Errorf("[ERROR] unable to find compartmentId for LogAnalytics NamespaceAssociation")
	}

	namespaceAssociations, err := listLogAnalyticsEntitySourceAssociationsForExport(ctx, namespace, compartmentId)
	if err != nil {
		return results, err
	}

	namespaceAssociationResource := tf_export.ResourcesMap[tfMeta.ResourceClass]
	discoveredAssociationIds := map[string]struct{}{}

	for _, namespaceAssociation := range namespaceAssociations {
		if namespaceAssociation.EntityId == nil || *namespaceAssociation.EntityId == "" {
			continue
		}
		if namespaceAssociation.SourceName == nil || *namespaceAssociation.SourceName == "" {
			continue
		}

		sourceName := *namespaceAssociation.SourceName
		entityId := *namespaceAssociation.EntityId
		compositeId := GetNamespaceAssociationCompositeId(namespace, compartmentId, entityId, sourceName)
		if _, exists := discoveredAssociationIds[compositeId]; exists {
			continue
		}
		discoveredAssociationIds[compositeId] = struct{}{}

		resource := &tf_export.OCIResource{
			CompartmentId:    compartmentId,
			SourceAttributes: getLogAnalyticsNamespaceAssociationSourceAttributes(namespace, compartmentId, namespaceAssociation),
			RawResource:      namespaceAssociation,
			TerraformResource: tf_export.TerraformResource{
				Id:             compositeId,
				ImportId:       compositeId,
				TerraformClass: tfMeta.ResourceClass,
			},
			GetHclStringFn: tf_export.GetHclStringFromGenericMap,
			Parent:         parent,
		}

		if resource.TerraformName, err = tf_export.GenerateTerraformNameFromResource(resource.SourceAttributes, namespaceAssociationResource.Schema); err != nil {
			entityName := getLogAnalyticsNamespaceAssociationStringAttribute(resource.SourceAttributes, "entity_name")
			resource.TerraformName = getLogAnalyticsNamespaceAssociationTerraformName(parent.TerraformName, sourceName, entityName, entityId)
		}

		results = append(results, resource)
	}

	return results, nil
}

func getLogAnalyticsNamespaceAssociationTerraformName(parentTerraformName string, sourceName string, entityName string, entityId string) string {
	associationName := entityName
	if associationName == "" {
		associationName = entityId
	}

	normalizedName := logAnalyticsTerraformNameSanitizer.ReplaceAllString(fmt.Sprintf("%s_%s", sourceName, associationName), "-")
	return tf_export.CheckDuplicateResourceName(fmt.Sprintf("%s_%s", parentTerraformName, normalizedName))
}

func getLogAnalyticsNamespaceStorageArchivalConfigId(resource *tf_export.OCIResource) (string, error) {

	namespace, ok := resource.SourceAttributes["namespace"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find namespace for LogAnalytics NamespaceStorageArchivalConfig")
	}
	return GetNamespaceStorageArchivalConfigCompositeId(namespace), nil
}

func getLogAnalyticsLogAnalyticsLogGroupId(resource *tf_export.OCIResource) (string, error) {
	logAnalyticsLogGroupId, ok := resource.SourceAttributes["id"].(string)
	if !ok || logAnalyticsLogGroupId == "" {
		return "", fmt.Errorf("[ERROR] unable to find logAnalyticsLogGroupId for LogAnalytics LogAnalyticsLogGroup")
	}

	namespace := ""
	if ns, ok := resource.SourceAttributes["namespace"].(string); ok && ns != "" {
		namespace = ns
	}
	if namespace == "" && resource.Parent != nil {
		if parentNamespace, ok := resource.Parent.SourceAttributes["namespace"].(string); ok && parentNamespace != "" {
			namespace = parentNamespace
		}
	}
	if namespace == "" {
		return "", fmt.Errorf("[ERROR] unable to find namespace for LogAnalytics LogAnalyticsLogGroup")
	}

	resource.SourceAttributes["namespace"] = namespace
	compositeId := getLogAnalyticsLogGroupCompositeId(logAnalyticsLogGroupId, namespace)
	resource.ImportId = compositeId
	return compositeId, nil
}

func getLogAnalyticsLogAnalyticsPreferencesManagementId(resource *tf_export.OCIResource) (string, error) {
	namespace := ""
	if ns, ok := resource.SourceAttributes["namespace"].(string); ok && ns != "" {
		namespace = ns
	}
	if namespace == "" && resource.Parent != nil {
		if parentNamespace, ok := resource.Parent.SourceAttributes["namespace"].(string); ok && parentNamespace != "" {
			namespace = parentNamespace
		}
	}
	if namespace == "" {
		return "", fmt.Errorf("[ERROR] unable to find namespace for LogAnalytics LogAnalyticsPreferencesManagement")
	}

	resource.SourceAttributes["namespace"] = namespace
	compositeId := getLogAnalyticsPreferencesManagementId(namespace)
	resource.ImportId = compositeId
	return compositeId, nil
}

func getLogAnalyticsLogAnalyticsResourceCategoriesManagementId(resource *tf_export.OCIResource) (string, error) {
	namespace := ""
	if ns, ok := resource.SourceAttributes["namespace"].(string); ok && ns != "" {
		namespace = ns
	}
	if namespace == "" && resource.Parent != nil {
		if parentNamespace, ok := resource.Parent.SourceAttributes["namespace"].(string); ok && parentNamespace != "" {
			namespace = parentNamespace
		}
	}
	if namespace == "" {
		return "", fmt.Errorf("[ERROR] unable to find namespace for LogAnalytics LogAnalyticsResourceCategoriesManagement")
	}

	resourceId, ok := resource.SourceAttributes["resource_id"].(string)
	if !ok || resourceId == "" {
		return "", fmt.Errorf("[ERROR] unable to find resourceId for LogAnalytics LogAnalyticsResourceCategoriesManagement")
	}

	resourceType, ok := resource.SourceAttributes["resource_type"].(string)
	if !ok || resourceType == "" {
		return "", fmt.Errorf("[ERROR] unable to find resourceType for LogAnalytics LogAnalyticsResourceCategoriesManagement")
	}

	resource.SourceAttributes["namespace"] = namespace
	compositeId := getLogAnalyticsResourceCategoriesManagementId(namespace, resourceId, resourceType)
	resource.ImportId = compositeId
	return compositeId, nil
}

func findLogAnalyticsLogAnalyticsResourceCategoriesManagement(ctx *tf_export.ResourceDiscoveryContext, tfMeta *tf_export.TerraformResourceAssociation, parent *tf_export.OCIResource, resourceGraph *tf_export.TerraformResourceGraph) ([]*tf_export.OCIResource, error) {
	results := []*tf_export.OCIResource{}

	namespace, ok := parent.SourceAttributes["namespace"].(string)
	if !ok || namespace == "" {
		return results, fmt.Errorf("[ERROR] unable to find namespace for LogAnalytics LogAnalyticsResourceCategoriesManagement")
	}

	request := oci_log_analytics.ListResourceCategoriesRequest{}
	request.NamespaceName = &namespace
	request.CompartmentId = ctx.CompartmentId
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "log_analytics")

	response, err := ctx.Clients.LogAnalyticsClient().ListResourceCategories(context.Background(), request)
	if err != nil {
		return results, err
	}

	request.Page = response.OpcNextPage
	for request.Page != nil {
		listResponse, err := ctx.Clients.LogAnalyticsClient().ListResourceCategories(context.Background(), request)
		if err != nil {
			return results, err
		}

		response.Items = append(response.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	groupedResources := map[string]*tf_export.OCIResource{}
	groupedCategoryNames := map[string]map[string]struct{}{}
	orderedResourceIds := make([]string, 0)

	for _, item := range response.Items {
		if item.ResourceId == nil || *item.ResourceId == "" || item.ResourceType == nil || *item.ResourceType == "" || item.CategoryName == nil || *item.CategoryName == "" {
			continue
		}

		sourceAttributes := LogAnalyticsResourceCategoryToMap(item)
		sourceAttributes["namespace"] = namespace

		resource := &tf_export.OCIResource{
			CompartmentId:    *ctx.CompartmentId,
			SourceAttributes: sourceAttributes,
			RawResource:      item,
			TerraformResource: tf_export.TerraformResource{
				TerraformClass:    tfMeta.ResourceClass,
				TerraformTypeInfo: tfMeta.TerraformResourceHints,
			},
			GetHclStringFn: tf_export.GetHclStringFromGenericMap,
			Parent:         parent,
		}

		compositeId, err := getLogAnalyticsLogAnalyticsResourceCategoriesManagementId(resource)
		if err != nil {
			return results, err
		}

		existingResource, exists := groupedResources[compositeId]
		if !exists {
			resource.Id = compositeId
			resource.ImportId = compositeId
			resource.SourceAttributes["resource_categories"] = []interface{}{*item.CategoryName}

			groupedResources[compositeId] = resource
			groupedCategoryNames[compositeId] = map[string]struct{}{
				*item.CategoryName: {},
			}
			orderedResourceIds = append(orderedResourceIds, compositeId)
			continue
		}

		if _, categoryExists := groupedCategoryNames[compositeId][*item.CategoryName]; categoryExists {
			continue
		}

		existingCategories := existingResource.SourceAttributes["resource_categories"].([]interface{})
		existingResource.SourceAttributes["resource_categories"] = append(existingCategories, *item.CategoryName)
		groupedCategoryNames[compositeId][*item.CategoryName] = struct{}{}
	}

	sort.Strings(orderedResourceIds)
	for _, compositeId := range orderedResourceIds {
		resource := groupedResources[compositeId]
		categories := resource.SourceAttributes["resource_categories"].([]interface{})
		sortedCategories := make([]string, len(categories))
		for i, category := range categories {
			sortedCategories[i] = category.(string)
		}
		sort.Strings(sortedCategories)

		resourceCategories := make([]interface{}, len(sortedCategories))
		for i, category := range sortedCategories {
			resourceCategories[i] = category
		}
		resource.SourceAttributes["resource_categories"] = resourceCategories

		resourceId := resource.SourceAttributes["resource_id"].(string)
		resourceType := resource.SourceAttributes["resource_type"].(string)
		resource.TerraformName = getLogAnalyticsResourceCategoriesManagementTerraformName(parent.TerraformName, resourceType, resourceId)
		results = append(results, resource)
	}

	return results, nil
}

func getLogAnalyticsResourceCategoriesManagementTerraformName(parentTerraformName string, resourceType string, resourceId string) string {
	normalizedName := logAnalyticsTerraformNameSanitizer.ReplaceAllString(fmt.Sprintf("%s_%s", resourceType, resourceId), "-")
	return tf_export.CheckDuplicateResourceName(fmt.Sprintf("%s_%s", parentTerraformName, normalizedName))
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

var exportLogAnalyticsLogAnalyticsLogGroupHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_log_analytics_log_analytics_log_group",
	DatasourceClass:        "oci_log_analytics_log_analytics_log_groups",
	DatasourceItemsAttr:    "log_analytics_log_group_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "log_analytics_log_group",
	RequireResourceRefresh: true,
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
	DatasourceClass:      "oci_log_analytics_log_analytics_preference",
	ResourceAbbreviation: "log_analytics_preferences_management",
}

var exportLogAnalyticsLogAnalyticsUnprocessedDataBucketManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_log_analytics_log_analytics_unprocessed_data_bucket_management",
	ResourceAbbreviation: "log_analytics_unprocessed_data_bucket_management",
}

var exportLogAnalyticsLogAnalyticsResourceCategoriesManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_log_analytics_log_analytics_resource_categories_management",
	DatasourceClass:      "oci_log_analytics_log_analytics_resource_categories_list",
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

var exportLogAnalyticsNamespaceLookupHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_log_analytics_namespace_lookup",
	DatasourceClass:      "oci_log_analytics_namespace_lookup",
	ResourceAbbreviation: "namespace_lookup",
}

var exportLogAnalyticsNamespaceLookupsUpdateDataManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_log_analytics_namespace_lookups_update_data_management",
	ResourceAbbreviation: "namespace_lookups_update_data_management",
}

var exportLogAnalyticsNamespaceLookupsAppendDataManagementHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_log_analytics_namespace_lookups_append_data_management",
	ResourceAbbreviation: "namespace_lookups_append_data_management",
}

var exportLogAnalyticsNamespaceStorageArchivalConfigHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_log_analytics_namespace_storage_archival_config",
	DatasourceClass:      "oci_log_analytics_namespace_storage_archival_config",
	ResourceAbbreviation: "namespace_storage_archival_config",
}

var exportLogAnalyticsNamespaceAssociationHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_log_analytics_namespace_association",
	ResourceAbbreviation: "namespace_association",
}

var exportLogAnalyticsObjectStorageNamespaceHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_objectstorage_namespace",
	DatasourceClass:      "oci_objectstorage_namespace",
	ResourceAbbreviation: "objectstorage_namespace",
	IsDataSource:         true,
	ProcessDiscoveredResourcesFn: func(_ *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
		for _, resource := range resources {
			resource.OmitFromExport = true
		}
		return resources, nil
	},
}

var logAnalyticsResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportLogAnalyticsObjectStorageNamespaceHints},
		{TerraformResourceHints: exportLogAnalyticsLogAnalyticsObjectCollectionRuleHints},
	},
	"oci_objectstorage_namespace": {
		{
			TerraformResourceHints: exportLogAnalyticsLogAnalyticsLogGroupHints,
			DatasourceQueryParams: map[string]string{
				"namespace": "namespace",
			},
		},
		{
			TerraformResourceHints: exportLogAnalyticsNamespaceIngestTimeRuleHints,
			DatasourceQueryParams: map[string]string{
				"namespace": "namespace",
			},
		},
		{
			TerraformResourceHints: exportLogAnalyticsNamespaceLookupHints,
			DatasourceQueryParams: map[string]string{
				"namespace": "namespace",
			},
		},
		{
			TerraformResourceHints: exportLogAnalyticsNamespaceAssociationHints,
		},
	},
}

var logAnalyticsTenancyResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_tenancy": {
		{TerraformResourceHints: exportLogAnalyticsObjectStorageNamespaceHints},
	},
	"oci_objectstorage_namespace": {
		{
			TerraformResourceHints: exportLogAnalyticsLogAnalyticsPreferencesManagementHints,
			DatasourceQueryParams: map[string]string{
				"namespace": "namespace",
			},
		},
		{
			TerraformResourceHints: exportLogAnalyticsLogAnalyticsResourceCategoriesManagementHints,
			DatasourceQueryParams: map[string]string{
				"namespace": "namespace",
			},
		},
	},
}
