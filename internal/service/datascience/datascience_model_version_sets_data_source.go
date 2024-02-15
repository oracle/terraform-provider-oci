// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatascienceModelVersionSetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatascienceModelVersionSets,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"created_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
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
			"model_version_sets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatascienceModelVersionSetResource()),
			},
		},
	}
}

func readDatascienceModelVersionSets(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelVersionSetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

type DatascienceModelVersionSetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.ListModelVersionSetsResponse
}

func (s *DatascienceModelVersionSetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceModelVersionSetsDataSourceCrud) Get() error {
	request := oci_datascience.ListModelVersionSetsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if createdBy, ok := s.D.GetOkExists("created_by"); ok {
		tmp := createdBy.(string)
		request.CreatedBy = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_datascience.ListModelVersionSetsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.ListModelVersionSets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListModelVersionSets(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatascienceModelVersionSetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatascienceModelVersionSetsDataSource-", DatascienceModelVersionSetsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		modelVersionSet := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CreatedBy != nil {
			modelVersionSet["created_by"] = *r.CreatedBy
		}

		if r.DefinedTags != nil {
			modelVersionSet["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		modelVersionSet["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			modelVersionSet["id"] = *r.Id
		}

		if r.Name != nil {
			modelVersionSet["name"] = *r.Name
		}

		if r.ProjectId != nil {
			modelVersionSet["project_id"] = *r.ProjectId
		}

		modelVersionSet["state"] = r.LifecycleState

		if r.SystemTags != nil {
			modelVersionSet["system_tags"] = tfresource.SystemTagsToMap(r.SystemTags)
		}

		if r.TimeCreated != nil {
			modelVersionSet["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			modelVersionSet["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, modelVersionSet)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatascienceModelVersionSetsDataSource().Schema["model_version_sets"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("model_version_sets", resources); err != nil {
		return err
	}

	return nil
}
