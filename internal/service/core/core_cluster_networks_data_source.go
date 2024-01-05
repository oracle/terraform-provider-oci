// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"
)

func CoreClusterNetworksDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreClusterNetworks,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cluster_networks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreClusterNetworkResource()),
			},
		},
	}
}

func readCoreClusterNetworks(d *schema.ResourceData, m interface{}) error {
	sync := &CoreClusterNetworksDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeManagementClient()

	return tfresource.ReadResource(sync)
}

type CoreClusterNetworksDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeManagementClient
	Res    *oci_core.ListClusterNetworksResponse
}

func (s *CoreClusterNetworksDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreClusterNetworksDataSourceCrud) Get() error {
	request := oci_core.ListClusterNetworksRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.ClusterNetworkSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListClusterNetworks(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListClusterNetworks(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreClusterNetworksDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreClusterNetworksDataSource-", CoreClusterNetworksDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		clusterNetwork := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			clusterNetwork["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			clusterNetwork["display_name"] = *r.DisplayName
		}

		clusterNetwork["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			clusterNetwork["id"] = *r.Id
		}

		instancePools := []interface{}{}
		for _, item := range r.InstancePools {
			instancePools = append(instancePools, InstancePoolSummaryToMap(item))
		}
		clusterNetwork["instance_pools"] = instancePools

		clusterNetwork["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			clusterNetwork["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			clusterNetwork["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, clusterNetwork)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreClusterNetworksDataSource().Schema["cluster_networks"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("cluster_networks", resources); err != nil {
		return err
	}

	return nil
}

func InstancePoolSummaryToMap(obj oci_core.InstancePoolSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.InstanceConfigurationId != nil {
		result["instance_configuration_id"] = string(*obj.InstanceConfigurationId)
	}

	if obj.Size != nil {
		result["size"] = int(*obj.Size)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}
