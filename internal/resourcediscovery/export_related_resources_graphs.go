package resourcediscovery

import oci_core "github.com/oracle/oci-go-sdk/v56/core"

/*
 * exportRelatedResourcesGraph specify the related resources for a given resource type
 */
var exportRelatedResourcesGraph = TerraformResourceGraph{
	//INSTANCES
	"oci_core_instance": {
		{
			TerraformResourceHints: exportCoreVolumeAttachmentHints,
			datasourceQueryParams: map[string]string{
				"availability_domain": "availability_domain",
				"instance_id":         "id",
			},
		},
		{
			TerraformResourceHints: exportCoreVnicAttachmentHints,
			datasourceQueryParams: map[string]string{
				"instance_id": "id",
			},
		},
	},
	"oci_core_volume_attachment": {
		{
			TerraformResourceHints: exportCoreVolumeClosureHints,
			datasourceQueryParams: map[string]string{
				"volume_id": "volume_id",
			},
		},
	},

	//LOAD BALANCERS

	"oci_load_balancer_backend_set": {
		{
			TerraformResourceHints: exportLoadBalancerBackendHints,
			datasourceQueryParams: map[string]string{
				"backendset_name":  "name",
				"load_balancer_id": "load_balancer_id",
			},
		},
		{TerraformResourceHints: exportLoadBalancerListenerHints},
	},
	"oci_load_balancer_load_balancer": {
		// certificates have to be discovered before listeners in order to populate
		// the references for certificate_name in listeners (dependency)
		// If moving to parallel execution in future, this dependency needs to be maintained
		{
			TerraformResourceHints: exportLoadBalancerCertificateHints,
			datasourceQueryParams: map[string]string{
				"load_balancer_id": "id",
			},
		},
		{
			TerraformResourceHints: exportLoadBalancerBackendSetHints,
			datasourceQueryParams: map[string]string{
				"load_balancer_id": "id",
			},
		},
	},

	//NETWORK LOAD BALANCERS

	"oci_network_load_balancer_backend_set": {
		{
			TerraformResourceHints: exportNetworkLoadBalancerBackendHints,
			datasourceQueryParams: map[string]string{
				"backend_set_name":         "name",
				"network_load_balancer_id": "network_load_balancer_id",
			},
		},
		{TerraformResourceHints: exportNetworkLoadBalancerListenerHints},
	},
}

// Separate hints for closure as we want to discover volumes related to an instance only
// and list data source for these give volumes for an AD or a volume Group only

var exportCoreVolumeClosureHints = &TerraformResourceHints{
	resourceClass:        "oci_core_volume",
	datasourceClass:      "oci_core_volume",
	resourceAbbreviation: "volume",
	discoverableLifecycleStates: []string{
		string(oci_core.VolumeLifecycleStateAvailable),
	},
}
