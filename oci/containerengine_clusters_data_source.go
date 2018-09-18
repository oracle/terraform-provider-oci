// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/containerengine"
)

func ClustersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readClusters,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
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
				Elem:     GetDataSourceItemSchema(ClusterResource()),
			},
		},
	}
}

func readClusters(d *schema.ResourceData, m interface{}) error {
	sync := &ClustersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).containerEngineClient

	return ReadResource(sync)
}

type ClustersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.ListClustersResponse
}

func (s *ClustersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ClustersDataSourceCrud) Get() error {
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
		var enumStates []oci_containerengine.ListClustersLifecycleStateEnum
		for _, r := range states.([]string) {
			enumStates = append(enumStates, oci_containerengine.ListClustersLifecycleStateEnum(r))
		}
		request.LifecycleState = enumStates
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "containerengine")

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

func (s *ClustersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		cluster := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		cluster["available_kubernetes_upgrades"] = r.AvailableKubernetesUpgrades

		if r.Endpoints != nil {
			cluster["endpoints"] = []interface{}{ClusterEndpointsToMap(r.Endpoints)}
		} else {
			cluster["endpoints"] = nil
		}

		if r.Id != nil {
			cluster["id"] = *r.Id
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

		if r.VcnId != nil {
			cluster["vcn_id"] = *r.VcnId
		}

		resources = append(resources, cluster)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, ClustersDataSource().Schema["clusters"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("clusters", resources); err != nil {
		return err
	}

	return nil
}
