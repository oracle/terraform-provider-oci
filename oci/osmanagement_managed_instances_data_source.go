// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_osmanagement "github.com/oracle/oci-go-sdk/osmanagement"
)

func OsmanagementManagedInstancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsmanagementManagedInstances,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_instances": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"child_software_sources": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_boot": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_checkin": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"managed_instance_groups": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"os_kernel_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"os_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"os_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"parent_software_source": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"updates_available": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readOsmanagementManagedInstances(d *schema.ResourceData, m interface{}) error {
	sync := &OsmanagementManagedInstancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).osManagementClient

	return ReadResource(sync)
}

type OsmanagementManagedInstancesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_osmanagement.OsManagementClient
	Res    *oci_osmanagement.ListManagedInstancesResponse
}

func (s *OsmanagementManagedInstancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsmanagementManagedInstancesDataSourceCrud) Get() error {
	request := oci_osmanagement.ListManagedInstancesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "osmanagement")

	response, err := s.Client.ListManagedInstances(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListManagedInstances(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsmanagementManagedInstancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		managedInstance := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.Description != nil {
			managedInstance["description"] = *r.Description
		}

		if r.DisplayName != nil {
			managedInstance["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			managedInstance["id"] = *r.Id
		}

		if r.LastBoot != nil {
			managedInstance["last_boot"] = *r.LastBoot
		}

		if r.LastCheckin != nil {
			managedInstance["last_checkin"] = *r.LastCheckin
		}

		managedInstance["status"] = r.Status

		if r.UpdatesAvailable != nil {
			managedInstance["updates_available"] = *r.UpdatesAvailable
		}

		resources = append(resources, managedInstance)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, OsmanagementManagedInstancesDataSource().Schema["managed_instances"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("managed_instances", resources); err != nil {
		return err
	}

	return nil
}

func SoftwareSourceIdToMap(obj *oci_osmanagement.SoftwareSourceId) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}
