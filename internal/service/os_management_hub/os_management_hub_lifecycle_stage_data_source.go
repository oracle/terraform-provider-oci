// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubLifecycleStageDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularOsManagementHubLifecycleStage,
		Schema: map[string]*schema.Schema{
			"lifecycle_stage_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"arch_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"lifecycle_environment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"managed_instance_ids": {
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
			"os_family": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"rank": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"software_source_id": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
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
						"software_source_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vendor_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularOsManagementHubLifecycleStage(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubLifecycleStageDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LifecycleEnvironmentClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubLifecycleStageDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.LifecycleEnvironmentClient
	Res    *oci_os_management_hub.GetLifecycleStageResponse
}

func (s *OsManagementHubLifecycleStageDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubLifecycleStageDataSourceCrud) Get() error {
	request := oci_os_management_hub.GetLifecycleStageRequest{}

	if lifecycleStageId, ok := s.D.GetOkExists("lifecycle_stage_id"); ok {
		tmp := lifecycleStageId.(string)
		request.LifecycleStageId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.GetLifecycleStage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsManagementHubLifecycleStageDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("arch_type", s.Res.ArchType)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleEnvironmentId != nil {
		s.D.Set("lifecycle_environment_id", *s.Res.LifecycleEnvironmentId)
	}

	managedInstanceIds := []interface{}{}
	for _, item := range s.Res.ManagedInstanceIds {
		managedInstanceIds = append(managedInstanceIds, ManagedInstanceDetailsToMap(item))
	}
	s.D.Set("managed_instance_ids", managedInstanceIds)

	s.D.Set("os_family", s.Res.OsFamily)

	if s.Res.Rank != nil {
		s.D.Set("rank", *s.Res.Rank)
	}

	if s.Res.SoftwareSourceId != nil {
		s.D.Set("software_source_id", []interface{}{SoftwareSourceDetailsToMap(*s.Res.SoftwareSourceId)})
	} else {
		s.D.Set("software_source_id", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeModified != nil {
		s.D.Set("time_modified", s.Res.TimeModified.String())
	}

	s.D.Set("vendor_name", s.Res.VendorName)

	return nil
}
