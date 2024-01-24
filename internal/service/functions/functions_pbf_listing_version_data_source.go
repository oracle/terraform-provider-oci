// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package functions

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_functions "github.com/oracle/oci-go-sdk/v65/functions"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FunctionsPbfListingVersionDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularFunctionsPbfListingVersion,
		Schema: map[string]*schema.Schema{
			"pbf_listing_version_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"change_summary": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"config": {
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
						"is_optional": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pbf_listing_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"requirements": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"min_memory_required_in_mbs": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"policies": {
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
									"policy": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"triggers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularFunctionsPbfListingVersion(d *schema.ResourceData, m interface{}) error {
	sync := &FunctionsPbfListingVersionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FunctionsManagementClient()

	return tfresource.ReadResource(sync)
}

type FunctionsPbfListingVersionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_functions.FunctionsManagementClient
	Res    *oci_functions.GetPbfListingVersionResponse
}

func (s *FunctionsPbfListingVersionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FunctionsPbfListingVersionDataSourceCrud) Get() error {
	request := oci_functions.GetPbfListingVersionRequest{}

	if pbfListingVersionId, ok := s.D.GetOkExists("pbf_listing_version_id"); ok {
		tmp := pbfListingVersionId.(string)
		request.PbfListingVersionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "functions")

	response, err := s.Client.GetPbfListingVersion(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FunctionsPbfListingVersionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ChangeSummary != nil {
		s.D.Set("change_summary", *s.Res.ChangeSummary)
	}

	config := []interface{}{}
	for _, item := range s.Res.Config {
		config = append(config, ConfigDetailsToMap(item))
	}
	s.D.Set("config", config)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.PbfListingId != nil {
		s.D.Set("pbf_listing_id", *s.Res.PbfListingId)
	}

	if s.Res.Requirements != nil {
		s.D.Set("requirements", []interface{}{RequirementDetailsToMap(s.Res.Requirements)})
	} else {
		s.D.Set("requirements", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	triggers := []interface{}{}
	for _, item := range s.Res.Triggers {
		triggers = append(triggers, TriggerToMap(item))
	}
	s.D.Set("triggers", triggers)

	return nil
}
