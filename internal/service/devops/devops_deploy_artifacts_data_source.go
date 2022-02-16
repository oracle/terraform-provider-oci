// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_devops "github.com/oracle/oci-go-sdk/v58/devops"
)

func DevopsDeployArtifactsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDevopsDeployArtifacts,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"deploy_artifact_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DevopsDeployArtifactResource()),
						},
					},
				},
			},
		},
	}
}

func readDevopsDeployArtifacts(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsDeployArtifactsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

type DevopsDeployArtifactsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.ListDeployArtifactsResponse
}

func (s *DevopsDeployArtifactsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsDeployArtifactsDataSourceCrud) Get() error {
	request := oci_devops.ListDeployArtifactsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_devops.DeployArtifactLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "devops")

	response, err := s.Client.ListDeployArtifacts(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDeployArtifacts(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DevopsDeployArtifactsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DevopsDeployArtifactsDataSource-", DevopsDeployArtifactsDataSource(), s.D))
	resources := []map[string]interface{}{}
	deployArtifact := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DeployArtifactSummaryToMap(item))
	}
	deployArtifact["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DevopsDeployArtifactsDataSource().Schema["deploy_artifact_collection"].Elem.(*schema.Resource).Schema)
		deployArtifact["items"] = items
	}

	resources = append(resources, deployArtifact)
	if err := s.D.Set("deploy_artifact_collection", resources); err != nil {
		return err
	}

	return nil
}
