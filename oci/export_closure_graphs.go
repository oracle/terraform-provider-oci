package oci

import (
	oci_core "github.com/oracle/oci-go-sdk/v25/core"
)

/*
 * resourceClosureGraph specify the related resources for a given resource type
 */
var resourceClosureGraph = TerraformResourceGraph{

	/*
		INSTANCES
	*/
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
	/*
		LOAD BALANCERS
	*/
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
		{
			TerraformResourceHints: exportLoadBalancerBackendSetHints,
			datasourceQueryParams: map[string]string{
				"load_balancer_id": "id",
			},
		},
		{
			TerraformResourceHints: exportLoadBalancerCertificateHints,
			datasourceQueryParams: map[string]string{
				"load_balancer_id": "id",
			},
		},
	},
}

// Separate hints for closure as we want to discover volumes related to an instance only
// and list data source for these give volumes for an AD or a volume group only
var exportCoreVolumeClosureHints = &TerraformResourceHints{
	resourceClass:        "oci_core_volume",
	datasourceClass:      "oci_core_volume",
	resourceAbbreviation: "volume",
	discoverableLifecycleStates: []string{
		string(oci_core.VolumeLifecycleStateAvailable),
	},
}
