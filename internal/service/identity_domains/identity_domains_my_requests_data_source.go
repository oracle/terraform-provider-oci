// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity_domains

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IdentityDomainsMyRequestsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityDomainsMyRequests,
		Schema: map[string]*schema.Schema{
			"idcs_endpoint": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"my_request_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"my_request_filter": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"attribute_sets": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"attributes": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"authorization": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_type_schema_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_index": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sort_order": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"my_requests": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"action": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"approval_details": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"approval_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"approver_display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"approver_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"justification": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"order": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"status": {
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
						"compartment_ocid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"delete_in_progress": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"domain_ocid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"expires": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idcs_created_by": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"display": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ocid": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"idcs_last_modified_by": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"display": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ocid": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"idcs_last_upgraded_in_release": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"idcs_prevented_operations": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"justification": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"meta": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"last_modified": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"location": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_type": {
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
						"ocid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"requesting": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"display": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"requestor": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"display": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"schemas": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tags": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"tenancy_ocid": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"items_per_page": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"schemas": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"total_results": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func readIdentityDomainsMyRequests(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsMyRequestsDataSourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpoint(d)
	if err != nil {
		return err
	}
	client, err := m.(*client.OracleClients).IdentityDomainsClientWithEndpoint(idcsEndpoint)
	if err != nil {
		return err
	}
	sync.Client = client

	return tfresource.ReadResource(sync)
}

type IdentityDomainsMyRequestsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity_domains.IdentityDomainsClient
	Res    *oci_identity_domains.ListMyRequestsResponse
}

func (s *IdentityDomainsMyRequestsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityDomainsMyRequestsDataSourceCrud) Get() error {
	request := oci_identity_domains.ListMyRequestsRequest{}

	if myRequestCount, ok := s.D.GetOkExists("my_request_count"); ok {
		tmp := myRequestCount.(int)
		request.Count = &tmp
	}

	if myRequestFilter, ok := s.D.GetOkExists("my_request_filter"); ok {
		tmp := myRequestFilter.(string)
		request.Filter = &tmp
	}

	if attributeSets, ok := s.D.GetOkExists("attribute_sets"); ok {
		interfaces := attributeSets.([]interface{})
		tmp := make([]oci_identity_domains.AttributeSetsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_identity_domains.AttributeSetsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("attribute_sets") {
			request.AttributeSets = tmp
		}
	}

	if attributes, ok := s.D.GetOkExists("attributes"); ok {
		tmp := attributes.(string)
		request.Attributes = &tmp
	}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	if startIndex, ok := s.D.GetOkExists("start_index"); ok {
		tmp := startIndex.(int)
		request.StartIndex = &tmp
	}

	if sortOrder, ok := s.D.GetOkExists("sort_order"); ok {
		tmp := oci_identity_domains.ListMyRequestsSortOrderEnum(sortOrder.(string))
		request.SortOrder = tmp
	}

	if sortBy, ok := s.D.GetOkExists("sort_by"); ok {
		tmp := sortBy.(string)
		request.SortBy = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity_domains")

	response, err := s.Client.ListMyRequests(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	// IDCS pagination
	startIndex := *response.StartIndex
	for startIndex+*response.ItemsPerPage <= *response.TotalResults {
		startIndex += *response.ItemsPerPage
		request.StartIndex = &startIndex
		listResponse, err := s.Client.ListMyRequests(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Resources = append(s.Res.Resources, listResponse.Resources...)
	}

	return nil
}

func (s *IdentityDomainsMyRequestsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentityDomainsMyRequestsDataSource-", IdentityDomainsMyRequestsDataSource(), s.D))

	resources := []interface{}{}
	for _, item := range s.Res.Resources {
		resources = append(resources, MyRequestToMap(item))
	}
	s.D.Set("my_requests", resources)

	if s.Res.ItemsPerPage != nil {
		s.D.Set("items_per_page", *s.Res.ItemsPerPage)
	}

	s.D.Set("schemas", s.Res.Schemas)

	if s.Res.StartIndex != nil {
		s.D.Set("start_index", *s.Res.StartIndex)
	}

	if s.Res.TotalResults != nil {
		s.D.Set("total_results", *s.Res.TotalResults)
	}

	return nil
}
