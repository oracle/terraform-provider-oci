package network_load_balancer

import (
	"context"
	"fmt"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_network_load_balancer "github.com/oracle/oci-go-sdk/v65/networkloadbalancer"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportNetworkLoadBalancerBackendSetHints.GetIdFn = getNetworkLoadBalancerBackendSetId
	exportNetworkLoadBalancerBackendHints.GetIdFn = getNetworkLoadBalancerBackendId
	exportNetworkLoadBalancerListenerHints.GetIdFn = getNetworkLoadBalancerListenerId
	exportNetworkLoadBalancerBackendHints.ProcessDiscoveredResourcesFn = processNetworkLoadBalancerBackends
	exportNetworkLoadBalancerBackendSetHints.ProcessDiscoveredResourcesFn = processNetworkLoadBalancerBackendSets
	exportNetworkLoadBalancerListenerHints.FindResourcesOverrideFn = findNetworkLoadBalancerListeners
	exportNetworkLoadBalancerListenerHints.ProcessDiscoveredResourcesFn = processNetworkLoadBalancerListeners
	tf_export.RegisterCompartmentGraphs("network_load_balancer", networkLoadBalancerResourceGraph)
	tf_export.RegisterRelatedResourcesGraph("oci_network_load_balancer_backend_set", relatedNetworkLoadBalancerBackendSet)
}

// Custom overrides for generating composite IDs within the resource discovery framework
func processNetworkLoadBalancerBackends(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, backend := range resources {
		if backend.Parent == nil {
			continue
		}

		backend.Id = GetNlbBackendCompositeId(backend.SourceAttributes["name"].(string), backend.Parent.SourceAttributes["name"].(string), backend.Parent.SourceAttributes["network_load_balancer_id"].(string))
		backend.SourceAttributes["network_load_balancer_id"] = backend.Parent.SourceAttributes["network_load_balancer_id"].(string)

		// Don't use references to parent resources if they will be omitted from final result
		if !backend.Parent.OmitFromExport {
			backend.SourceAttributes["backend_set_name"] = tf_export.InterpolationString{
				ResourceReference: backend.Parent.GetTerraformReference(),
				Interpolation:     tf_export.TfHclVersionvar.GetDoubleExpHclString(backend.Parent.GetTerraformReference(), "name"),
				Value:             backend.Parent.SourceAttributes["name"].(string),
			}
		} else {
			backend.SourceAttributes["backend_set_name"] = backend.Parent.SourceAttributes["name"].(string)
		}
	}

	return resources, nil
}

func processNetworkLoadBalancerBackendSets(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, backendSet := range resources {
		if backendSet.Parent == nil {
			continue
		}

		backendSetName := backendSet.SourceAttributes["name"].(string)
		backendSet.Id = GetNlbBackendSetCompositeId(backendSetName, backendSet.Parent.Id)
		backendSet.SourceAttributes["network_load_balancer_id"] = backendSet.Parent.Id
	}

	return resources, nil
}

func findNetworkLoadBalancerListeners(ctx *tf_export.ResourceDiscoveryContext, tfMeta *tf_export.TerraformResourceAssociation, parent *tf_export.OCIResource, resourceGraph *tf_export.TerraformResourceGraph) ([]*tf_export.OCIResource, error) {
	networkLoadBalancerId := parent.SourceAttributes["network_load_balancer_id"].(string)
	backendSetName := parent.SourceAttributes["name"].(string)

	request := oci_network_load_balancer.GetNetworkLoadBalancerRequest{}
	request.NetworkLoadBalancerId = &networkLoadBalancerId
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_load_balancer")

	response, err := ctx.Clients.NetworkLoadBalancerClient().GetNetworkLoadBalancer(context.Background(), request)
	if err != nil {
		return nil, err
	}

	listenerResource := tf_export.ResourcesMap[tfMeta.ResourceClass]

	results := []*tf_export.OCIResource{}
	for listenerName, listener := range response.NetworkLoadBalancer.Listeners {
		if *listener.DefaultBackendSetName != backendSetName {
			continue
		}

		d := listenerResource.TestResourceData()
		d.SetId(GetNlbListenerCompositeId(listenerName, networkLoadBalancerId))

		// This calls into the listener resource's Read fn which has the unfortunate implementation of
		// calling GetNetworkLoadBalancer and looping through the listeners to find the expected one. So this entire method
		// may require O(n^^2) time. However, the benefits of having Read populate the ResourceData struct is better than duplicating it here.
		if err := listenerResource.Read(d, ctx.Clients); err != nil {
			// add error to the errorList and continue discovering rest of the resources
			rdError := &tf_export.ResourceDiscoveryError{ResourceType: tfMeta.ResourceClass, ParentResource: parent.TerraformName, Error: err, ResourceGraph: resourceGraph}
			ctx.AddErrorToList(rdError)
			continue
		}

		resource := &tf_export.OCIResource{
			CompartmentId:    parent.CompartmentId,
			SourceAttributes: tf_export.ConvertResourceDataToMap(listenerResource.Schema, d),
			RawResource:      listener,
			TerraformResource: tf_export.TerraformResource{
				Id:             d.Id(),
				TerraformClass: tfMeta.ResourceClass,
				TerraformName:  fmt.Sprintf("%s_%s", parent.Parent.TerraformName, listenerName),
			},
			GetHclStringFn: tf_export.GetHclStringFromGenericMap,
			Parent:         parent,
		}

		if !parent.OmitFromExport {
			resource.SourceAttributes["default_backend_set_name"] = tf_export.InterpolationString{
				ResourceReference: parent.GetTerraformReference(),
				Interpolation:     tf_export.TfHclVersionvar.GetDoubleExpHclString(parent.GetTerraformReference(), "name"),
				Value:             parent.SourceAttributes["name"].(string),
			}
		} else {
			resource.SourceAttributes["default_backend_set_name"] = parent.SourceAttributes["name"].(string)
		}
		resource.TerraformName = tf_export.GetValidUniqueTerraformName(resource.TerraformName)
		results = append(results, resource)
	}

	return results, nil
}

func processNetworkLoadBalancerListeners(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, listener := range resources {
		if listener.Parent == nil {
			continue
		}

		listenerName := listener.SourceAttributes["name"].(string)
		listener.Id = GetNlbListenerCompositeId(listenerName, listener.Parent.SourceAttributes["network_load_balancer_id"].(string))
		listener.SourceAttributes["network_load_balancer_id"] = listener.Parent.SourceAttributes["network_load_balancer_id"].(string)

		// Don't use references to parent resources if they will be omitted from final result
		if !listener.Parent.OmitFromExport {
			listener.SourceAttributes["default_backend_set_name"] = tf_export.InterpolationString{
				ResourceReference: listener.Parent.GetTerraformReference(),
				Interpolation:     tf_export.TfHclVersionvar.GetDoubleExpHclString(listener.Parent.GetTerraformReference(), "name"),
				Value:             listener.Parent.SourceAttributes["name"].(string),
			}
		} else {
			listener.SourceAttributes["default_backend_set_name"] = listener.Parent.SourceAttributes["name"].(string)
		}
	}

	return resources, nil
}

func getNetworkLoadBalancerBackendSetId(resource *tf_export.OCIResource) (string, error) {

	backendSetName, ok := resource.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find backendSetName for NetworkLoadBalancer BackendSet")
	}
	networkLoadBalancerId := resource.Parent.Id
	return GetNlbBackendSetCompositeId(backendSetName, networkLoadBalancerId), nil
}

func getNetworkLoadBalancerBackendId(resource *tf_export.OCIResource) (string, error) {

	backendName, ok := resource.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find backendName for NetworkLoadBalancer Backend")
	}
	backendsetName, ok := resource.Parent.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find backendSetName for NetworkLoadBalancer Backend")
	}
	networkLoadBalancerId := resource.Parent.Parent.Id
	return GetNlbBackendCompositeId(backendName, backendsetName, networkLoadBalancerId), nil
}

func getNetworkLoadBalancerListenerId(resource *tf_export.OCIResource) (string, error) {

	listenerName, ok := resource.SourceAttributes["name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find listenerName for NetworkLoadBalancer Listener")
	}
	networkLoadBalancerId := resource.Parent.Parent.Id
	return GetNlbListenerCompositeId(listenerName, networkLoadBalancerId), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportNetworkLoadBalancerNetworkLoadBalancerHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_network_load_balancer_network_load_balancer",
	DatasourceClass:        "oci_network_load_balancer_network_load_balancers",
	DatasourceItemsAttr:    "network_load_balancer_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "network_load_balancer",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_network_load_balancer.LifecycleStateActive),
	},
}

var exportNetworkLoadBalancerBackendSetHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_network_load_balancer_backend_set",
	DatasourceClass:        "oci_network_load_balancer_backend_sets",
	DatasourceItemsAttr:    "backend_set_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "backend_set",
	RequireResourceRefresh: true,
}

var exportNetworkLoadBalancerBackendHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_network_load_balancer_backend",
	DatasourceClass:        "oci_network_load_balancer_backends",
	DatasourceItemsAttr:    "backend_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "backend",
	RequireResourceRefresh: true,
}

var exportNetworkLoadBalancerListenerHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_network_load_balancer_listener",
	DatasourceClass:        "oci_network_load_balancer_listeners",
	DatasourceItemsAttr:    "listener_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "listener",
	RequireResourceRefresh: true,
}

var networkLoadBalancerResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportNetworkLoadBalancerNetworkLoadBalancerHints},
	},
	"oci_network_load_balancer_network_load_balancer": {
		{
			TerraformResourceHints: exportNetworkLoadBalancerBackendSetHints,
			DatasourceQueryParams: map[string]string{
				"network_load_balancer_id": "id",
			},
		},
	},
	"oci_network_load_balancer_backend_set": {
		{
			TerraformResourceHints: exportNetworkLoadBalancerBackendHints,
			DatasourceQueryParams: map[string]string{
				"backend_set_name":         "name",
				"network_load_balancer_id": "network_load_balancer_id",
			},
		},
		{
			TerraformResourceHints: exportNetworkLoadBalancerListenerHints,
			DatasourceQueryParams: map[string]string{
				"network_load_balancer_id": "network_load_balancer_id",
			},
		},
	},
}

var relatedNetworkLoadBalancerBackendSet = []tf_export.TerraformResourceAssociation{
	{
		TerraformResourceHints: exportNetworkLoadBalancerBackendHints,
		DatasourceQueryParams: map[string]string{
			"backend_set_name":         "name",
			"network_load_balancer_id": "network_load_balancer_id",
		},
	},
	{TerraformResourceHints: exportNetworkLoadBalancerListenerHints},
}
