// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_osmanagement "github.com/oracle/oci-go-sdk/osmanagement"
)

func init() {
	RegisterDatasource("oci_osmanagement_managed_instance", OsmanagementManagedInstanceDataSource())
}

func OsmanagementManagedInstanceDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularOsmanagementManagedInstance,
		Schema: map[string]*schema.Schema{
			"managed_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
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
	}
}

func readSingularOsmanagementManagedInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OsmanagementManagedInstanceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).osManagementClient()

	return ReadResource(sync)
}

type OsmanagementManagedInstanceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_osmanagement.OsManagementClient
	Res    *oci_osmanagement.GetManagedInstanceResponse
}

func (s *OsmanagementManagedInstanceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsmanagementManagedInstanceDataSourceCrud) Get() error {
	request := oci_osmanagement.GetManagedInstanceRequest{}

	if managedInstanceId, ok := s.D.GetOkExists("managed_instance_id"); ok {
		tmp := managedInstanceId.(string)
		request.ManagedInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "osmanagement")

	response, err := s.Client.GetManagedInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsmanagementManagedInstanceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	childSoftwareSources := []interface{}{}
	for _, item := range s.Res.ChildSoftwareSources {
		childSoftwareSources = append(childSoftwareSources, SoftwareSourceIdToMap(&item))
	}
	s.D.Set("child_software_sources", childSoftwareSources)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.LastBoot != nil {
		s.D.Set("last_boot", *s.Res.LastBoot)
	}

	if s.Res.LastCheckin != nil {
		s.D.Set("last_checkin", *s.Res.LastCheckin)
	}

	managedInstanceGroups := []interface{}{}
	for _, item := range s.Res.ManagedInstanceGroups {
		managedInstanceGroups = append(managedInstanceGroups, IdToMap(item))
	}
	s.D.Set("managed_instance_groups", managedInstanceGroups)

	if s.Res.OsKernelVersion != nil {
		s.D.Set("os_kernel_version", *s.Res.OsKernelVersion)
	}

	if s.Res.OsName != nil {
		s.D.Set("os_name", *s.Res.OsName)
	}

	if s.Res.OsVersion != nil {
		s.D.Set("os_version", *s.Res.OsVersion)
	}

	if s.Res.ParentSoftwareSource != nil {
		s.D.Set("parent_software_source", []interface{}{SoftwareSourceIdToMap(s.Res.ParentSoftwareSource)})
	} else {
		s.D.Set("parent_software_source", nil)
	}

	s.D.Set("status", s.Res.Status)

	if s.Res.UpdatesAvailable != nil {
		s.D.Set("updates_available", *s.Res.UpdatesAvailable)
	}

	return nil
}
