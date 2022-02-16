// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v58/core"
)

func CoreClusterNetworkInstancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreClusterNetworkInstances,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"cluster_network_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"instances": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"availability_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"fault_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_configuration_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"load_balancer_backends": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"backend_health_status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"backend_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"backend_set_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"load_balancer_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									// internal for work request access
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"shape": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readCoreClusterNetworkInstances(d *schema.ResourceData, m interface{}) error {
	sync := &CoreClusterNetworkInstancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeManagementClient()

	return tfresource.ReadResource(sync)
}

type CoreClusterNetworkInstancesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeManagementClient
	Res    *oci_core.ListClusterNetworkInstancesResponse
}

func (s *CoreClusterNetworkInstancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreClusterNetworkInstancesDataSourceCrud) Get() error {
	request := oci_core.ListClusterNetworkInstancesRequest{}

	if clusterNetworkId, ok := s.D.GetOkExists("cluster_network_id"); ok {
		tmp := clusterNetworkId.(string)
		request.ClusterNetworkId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListClusterNetworkInstances(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListClusterNetworkInstances(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreClusterNetworkInstancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreClusterNetworkInstancesDataSource-", CoreClusterNetworkInstancesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		clusterNetworkInstance := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AvailabilityDomain != nil {
			clusterNetworkInstance["availability_domain"] = *r.AvailabilityDomain
		}

		if r.DisplayName != nil {
			clusterNetworkInstance["display_name"] = *r.DisplayName
		}

		if r.FaultDomain != nil {
			clusterNetworkInstance["fault_domain"] = *r.FaultDomain
		}

		if r.Id != nil {
			clusterNetworkInstance["id"] = *r.Id
		}

		if r.InstanceConfigurationId != nil {
			clusterNetworkInstance["instance_configuration_id"] = *r.InstanceConfigurationId
		}

		loadBalancerBackends := []interface{}{}
		for _, item := range r.LoadBalancerBackends {
			loadBalancerBackends = append(loadBalancerBackends, InstancePoolInstanceLoadBalancerBackendToMap(item))
		}
		clusterNetworkInstance["load_balancer_backends"] = loadBalancerBackends

		if r.Region != nil {
			clusterNetworkInstance["region"] = *r.Region
		}

		if r.Shape != nil {
			clusterNetworkInstance["shape"] = *r.Shape
		}

		if r.State != nil {
			clusterNetworkInstance["state"] = *r.State
		}

		if r.TimeCreated != nil {
			clusterNetworkInstance["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, clusterNetworkInstance)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreClusterNetworkInstancesDataSource().Schema["instances"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("instances", resources); err != nil {
		return err
	}

	return nil
}
