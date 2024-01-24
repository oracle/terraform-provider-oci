// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package container_instances

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_container_instances "github.com/oracle/oci-go-sdk/v65/containerinstances"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ContainerInstancesContainerInstanceShapesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readContainerInstancesContainerInstanceShapes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"container_instance_shape_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"memory_options": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"default_per_ocpu_in_gbs": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"max_in_gbs": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"max_per_ocpu_in_gbs": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"min_in_gbs": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"min_per_ocpu_in_gbs": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
											},
										},
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"networking_bandwidth_options": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"default_per_ocpu_in_gbps": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"max_in_gbps": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"min_in_gbps": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
											},
										},
									},
									"ocpu_options": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"max": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"min": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
											},
										},
									},
									"processor_description": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readContainerInstancesContainerInstanceShapes(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerInstancesContainerInstanceShapesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerInstanceClient()

	return tfresource.ReadResource(sync)
}

type ContainerInstancesContainerInstanceShapesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_container_instances.ContainerInstanceClient
	Res    *oci_container_instances.ListContainerInstanceShapesResponse
}

func (s *ContainerInstancesContainerInstanceShapesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerInstancesContainerInstanceShapesDataSourceCrud) Get() error {
	request := oci_container_instances.ListContainerInstanceShapesRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "containerinstance")

	response, err := s.Client.ListContainerInstanceShapes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListContainerInstanceShapes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ContainerInstancesContainerInstanceShapesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ContainerInstancesContainerInstanceShapesDataSource-", ContainerInstancesContainerInstanceShapesDataSource(), s.D))
	resources := []map[string]interface{}{}
	containerInstanceShape := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ContainerInstanceShapeSummaryToMap(item))
	}
	containerInstanceShape["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ContainerInstancesContainerInstanceShapesDataSource().Schema["container_instance_shape_collection"].Elem.(*schema.Resource).Schema)
		containerInstanceShape["items"] = items
	}

	resources = append(resources, containerInstanceShape)
	if err := s.D.Set("container_instance_shape_collection", resources); err != nil {
		return err
	}

	return nil
}
