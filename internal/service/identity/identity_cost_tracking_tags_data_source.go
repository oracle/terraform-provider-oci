// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v58/identity"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func IdentityCostTrackingTagsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityCostTrackingTags,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"defined_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"description": {
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
						"is_cost_tracking": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_retired": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tag_namespace_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tag_namespace_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"validator": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"validator_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"values": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
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

func readIdentityCostTrackingTags(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityCostTrackingTagsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

type IdentityCostTrackingTagsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListCostTrackingTagsResponse
}

func (s *IdentityCostTrackingTagsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityCostTrackingTagsDataSourceCrud) Get() error {
	request := oci_identity.ListCostTrackingTagsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity")

	response, err := s.Client.ListCostTrackingTags(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCostTrackingTags(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *IdentityCostTrackingTagsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentityCostTrackingTagsDataSource-", IdentityCostTrackingTagsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		costTrackingTag := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			costTrackingTag["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			costTrackingTag["description"] = *r.Description
		}

		costTrackingTag["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			costTrackingTag["id"] = *r.Id
		}

		if r.IsCostTracking != nil {
			costTrackingTag["is_cost_tracking"] = *r.IsCostTracking
		}

		if r.IsRetired != nil {
			costTrackingTag["is_retired"] = *r.IsRetired
		}

		if r.Name != nil {
			costTrackingTag["name"] = *r.Name
		}

		costTrackingTag["state"] = r.LifecycleState

		if r.TagNamespaceId != nil {
			costTrackingTag["tag_namespace_id"] = *r.TagNamespaceId
		}

		if r.TagNamespaceName != nil {
			costTrackingTag["tag_namespace_name"] = *r.TagNamespaceName
		}

		if r.TimeCreated != nil {
			costTrackingTag["time_created"] = r.TimeCreated.String()
		}

		if r.Validator != nil {
			validatorArray := []interface{}{}
			if validatorMap := BaseTagDefinitionValidatorToMap(&r.Validator); validatorMap != nil {
				validatorArray = append(validatorArray, validatorMap)
			}
			costTrackingTag["validator"] = validatorArray
		} else {
			costTrackingTag["validator"] = nil
		}

		resources = append(resources, costTrackingTag)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, IdentityCostTrackingTagsDataSource().Schema["tags"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("tags", resources); err != nil {
		return err
	}

	return nil
}
