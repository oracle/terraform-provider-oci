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

func ContainerInstancesContainerInstancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readContainerInstancesContainerInstances,
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
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"container_instance_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(ContainerInstancesContainerInstanceResource()),
						},
					},
				},
			},
		},
	}
}

func readContainerInstancesContainerInstances(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerInstancesContainerInstancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerInstanceClient()

	return tfresource.ReadResource(sync)
}

type ContainerInstancesContainerInstancesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_container_instances.ContainerInstanceClient
	Res    *oci_container_instances.ListContainerInstancesResponse
}

func (s *ContainerInstancesContainerInstancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerInstancesContainerInstancesDataSourceCrud) Get() error {
	request := oci_container_instances.ListContainerInstancesRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_container_instances.ContainerInstanceLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "containerinstance")

	response, err := s.Client.ListContainerInstances(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListContainerInstances(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ContainerInstancesContainerInstancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ContainerInstancesContainerInstancesDataSource-", ContainerInstancesContainerInstancesDataSource(), s.D))
	resources := []map[string]interface{}{}
	containerInstance := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ContainerInstanceSummaryToMap(item))
	}
	containerInstance["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ContainerInstancesContainerInstancesDataSource().Schema["container_instance_collection"].Elem.(*schema.Resource).Schema)
		containerInstance["items"] = items
	}

	resources = append(resources, containerInstance)
	if err := s.D.Set("container_instance_collection", resources); err != nil {
		return err
	}

	return nil
}
