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

func DatascienceModelGroupModelsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatascienceModelGroupModels,
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
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"model_group_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"model_group_models": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"category": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"created_by": {
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
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_model_by_reference": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"model_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"project_id": {
							Type:     schema.TypeString,
							Computed: true,
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
						"time_updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDatascienceModelGroupModels(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelGroupModelsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

type DatascienceModelGroupModelsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.ListModelGroupModelsResponse
}

func (s *DatascienceModelGroupModelsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceModelGroupModelsDataSourceCrud) Get() error {
	request := oci_datascience.ListModelGroupModelsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if createdBy, ok := s.D.GetOkExists("created_by"); ok {
		tmp := createdBy.(string)
		request.CreatedBy = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if modelGroupId, ok := s.D.GetOkExists("model_group_id"); ok {
		tmp := modelGroupId.(string)
		request.ModelGroupId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_datascience.ListModelGroupModelsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.ListModelGroupModels(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListModelGroupModels(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatascienceModelGroupModelsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatascienceModelGroupModelsDataSource-", DatascienceModelGroupModelsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		modelGroupModel := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		modelGroupModel["category"] = r.Category

		if r.CreatedBy != nil {
			modelGroupModel["created_by"] = *r.CreatedBy
		}

		if r.DefinedTags != nil {
			modelGroupModel["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			modelGroupModel["display_name"] = *r.DisplayName
		}

		modelGroupModel["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			modelGroupModel["id"] = *r.Id
		}

		if r.IsModelByReference != nil {
			modelGroupModel["is_model_by_reference"] = *r.IsModelByReference
		}

		if r.LifecycleDetails != nil {
			modelGroupModel["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.ModelId != nil {
			modelGroupModel["model_id"] = *r.ModelId
		}

		if r.ProjectId != nil {
			modelGroupModel["project_id"] = *r.ProjectId
		}

		modelGroupModel["state"] = r.LifecycleState

		if r.SystemTags != nil {
			modelGroupModel["system_tags"] = tfresource.SystemTagsToMap(r.SystemTags)
		}

		if r.TimeCreated != nil {
			modelGroupModel["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			modelGroupModel["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, modelGroupModel)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatascienceModelGroupModelsDataSource().Schema["model_group_models"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("model_group_models", resources); err != nil {
		return err
	}

	return nil
}
