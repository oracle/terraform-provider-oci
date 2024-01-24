// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/v65/containerengine"
)

func ContainerengineClustersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readContainerengineClusters,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"clusters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(ContainerengineClusterResource()),
			},
		},
	}
}

func readContainerengineClusters(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClustersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.ReadResource(sync)
}

type ContainerengineClustersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.ListClustersResponse
}

func (s *ContainerengineClustersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerengineClustersDataSourceCrud) Get() error {
	request := oci_containerengine.ListClustersRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if states, ok := s.D.GetOkExists("state"); ok {
		var enumStates []oci_containerengine.ClusterLifecycleStateEnum
		for _, r := range states.([]interface{}) {
			enumStates = append(enumStates, oci_containerengine.ClusterLifecycleStateEnum(r.(string)))
		}
		request.LifecycleState = enumStates
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "containerengine")

	response, err := s.Client.ListClusters(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListClusters(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ContainerengineClustersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ContainerengineClustersDataSource-", ContainerengineClustersDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		cluster := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		cluster["available_kubernetes_upgrades"] = r.AvailableKubernetesUpgrades

		clusterPodNetworkOptions := []interface{}{}
		for _, item := range r.ClusterPodNetworkOptions {
			clusterPodNetworkOptions = append(clusterPodNetworkOptions, ClusterPodNetworkOptionDetailsToMap(item))
		}
		cluster["cluster_pod_network_options"] = clusterPodNetworkOptions

		if r.DefinedTags != nil {
			cluster["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.EndpointConfig != nil {
			cluster["endpoint_config"] = []interface{}{ClusterEndpointConfigToMap(r.EndpointConfig, true)}
		} else {
			cluster["endpoint_config"] = nil
		}

		if r.Endpoints != nil {
			cluster["endpoints"] = []interface{}{ClusterEndpointsToMap(r.Endpoints)}
		} else {
			cluster["endpoints"] = nil
		}

		cluster["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			cluster["id"] = *r.Id
		}

		if r.ImagePolicyConfig != nil {
			cluster["image_policy_config"] = []interface{}{ImagePolicyConfigToMap(r.ImagePolicyConfig)}
		} else {
			cluster["image_policy_config"] = nil
		}

		if r.KubernetesVersion != nil {
			cluster["kubernetes_version"] = *r.KubernetesVersion
		}

		if r.LifecycleDetails != nil {
			cluster["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.Metadata != nil {
			cluster["metadata"] = []interface{}{ClusterMetadataToMap(r.Metadata)}
		} else {
			cluster["metadata"] = nil
		}

		if r.Name != nil {
			cluster["name"] = *r.Name
		}

		if r.Options != nil {
			cluster["options"] = []interface{}{ClusterCreateOptionsToMap(r.Options)}
		} else {
			cluster["options"] = nil
		}

		cluster["state"] = r.LifecycleState

		cluster["type"] = r.Type

		if r.VcnId != nil {
			cluster["vcn_id"] = *r.VcnId
		}

		resources = append(resources, cluster)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, ContainerengineClustersDataSource().Schema["clusters"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("clusters", resources); err != nil {
		return err
	}

	return nil
}
