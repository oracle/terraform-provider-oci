// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_management_agent "github.com/oracle/oci-go-sdk/v32/managementagent"
)

func init() {
	RegisterDatasource("oci_management_agent_management_agent_images", ManagementAgentManagementAgentImagesDataSource())
}

func ManagementAgentManagementAgentImagesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readManagementAgentManagementAgentImages,
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
				Type:     schema.TypeString,
				Optional: true,
			},
			"management_agent_images": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"checksum": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"object_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"platform_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"platform_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"size": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readManagementAgentManagementAgentImages(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentManagementAgentImagesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).managementAgentClient()

	return ReadResource(sync)
}

type ManagementAgentManagementAgentImagesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_management_agent.ManagementAgentClient
	Res    *oci_management_agent.ListManagementAgentImagesResponse
}

func (s *ManagementAgentManagementAgentImagesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ManagementAgentManagementAgentImagesDataSourceCrud) Get() error {
	request := oci_management_agent.ListManagementAgentImagesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_management_agent.ListManagementAgentImagesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "management_agent")

	response, err := s.Client.ListManagementAgentImages(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListManagementAgentImages(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ManagementAgentManagementAgentImagesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("ManagementAgentManagementAgentImagesDataSource-", ManagementAgentManagementAgentImagesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		managementAgentImage := map[string]interface{}{}

		if r.Checksum != nil {
			managementAgentImage["checksum"] = *r.Checksum
		}

		if r.Id != nil {
			managementAgentImage["id"] = *r.Id
		}

		if r.ObjectUrl != nil {
			managementAgentImage["object_url"] = *r.ObjectUrl
		}

		if r.PlatformName != nil {
			managementAgentImage["platform_name"] = *r.PlatformName
		}

		managementAgentImage["platform_type"] = r.PlatformType

		if r.Size != nil {
			managementAgentImage["size"] = *r.Size
		}

		managementAgentImage["state"] = r.LifecycleState

		if r.Version != nil {
			managementAgentImage["version"] = *r.Version
		}

		resources = append(resources, managementAgentImage)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, ManagementAgentManagementAgentImagesDataSource().Schema["management_agent_images"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("management_agent_images", resources); err != nil {
		return err
	}

	return nil
}
