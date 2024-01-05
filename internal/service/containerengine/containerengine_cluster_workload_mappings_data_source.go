// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package containerengine

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/v65/containerengine"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ContainerengineClusterWorkloadMappingsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readContainerengineClusterWorkloadMappings,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"workload_mappings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(ContainerengineClusterWorkloadMappingResource()),
			},
		},
	}
}

func readContainerengineClusterWorkloadMappings(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineClusterWorkloadMappingsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerEngineClient()

	return tfresource.ReadResource(sync)
}

type ContainerengineClusterWorkloadMappingsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.ListWorkloadMappingsResponse
}

func (s *ContainerengineClusterWorkloadMappingsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerengineClusterWorkloadMappingsDataSourceCrud) Get() error {
	request := oci_containerengine.ListWorkloadMappingsRequest{}

	if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
		tmp := clusterId.(string)
		request.ClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "containerengine")

	response, err := s.Client.ListWorkloadMappings(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListWorkloadMappings(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ContainerengineClusterWorkloadMappingsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ContainerengineClusterWorkloadMappingsDataSource-", ContainerengineClusterWorkloadMappingsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		clusterWorkloadMapping := map[string]interface{}{
			"cluster_id": *r.ClusterId,
		}

		if r.DefinedTags != nil {
			clusterWorkloadMapping["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		clusterWorkloadMapping["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			clusterWorkloadMapping["id"] = *r.Id
		}

		if r.MappedCompartmentId != nil {
			clusterWorkloadMapping["mapped_compartment_id"] = *r.MappedCompartmentId
		}

		if r.MappedTenancyId != nil {
			clusterWorkloadMapping["mapped_tenancy_id"] = *r.MappedTenancyId
		}

		if r.Namespace != nil {
			clusterWorkloadMapping["namespace"] = *r.Namespace
		}

		clusterWorkloadMapping["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			clusterWorkloadMapping["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, clusterWorkloadMapping)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, ContainerengineClusterWorkloadMappingsDataSource().Schema["workload_mappings"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("workload_mappings", resources); err != nil {
		return err
	}

	return nil
}
