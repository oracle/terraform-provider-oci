// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubDynamicSetManagedInstancesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readOsManagementHubDynamicSetManagedInstancesWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_set_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"managed_instance_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"agent_version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"architecture": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"autonomous_settings": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"is_data_collection_authorized": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"scheduled_job_id": {
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
									"is_managed_by_autonomous_linux": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_management_station": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_reboot_required": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"lifecycle_environment": {
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
									"lifecycle_stage": {
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
									"location": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"managed_instance_group": {
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
									"notification_topic_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"os_family": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"tenancy_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_last_boot": {
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
				},
			},
		},
	}
}

func readOsManagementHubDynamicSetManagedInstancesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &OsManagementHubDynamicSetManagedInstancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DynamicSetClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type OsManagementHubDynamicSetManagedInstancesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.DynamicSetClient
	Res    *oci_os_management_hub.ListManagedInstancesInDynamicSetResponse
}

func (s *OsManagementHubDynamicSetManagedInstancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubDynamicSetManagedInstancesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_os_management_hub.ListManagedInstancesInDynamicSetRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if displayNameContains, ok := s.D.GetOkExists("display_name_contains"); ok {
		tmp := displayNameContains.(string)
		request.DisplayNameContains = &tmp
	}

	if dynamicSetId, ok := s.D.GetOkExists("dynamic_set_id"); ok {
		tmp := dynamicSetId.(string)
		request.DynamicSetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.ListManagedInstancesInDynamicSet(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListManagedInstancesInDynamicSet(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsManagementHubDynamicSetManagedInstancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubDynamicSetManagedInstancesDataSource-", OsManagementHubDynamicSetManagedInstancesDataSource(), s.D))
	resources := []map[string]interface{}{}
	dynamicSetManagedInstance := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ManagedInstanceSummaryToMap(item))
	}
	dynamicSetManagedInstance["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OsManagementHubDynamicSetManagedInstancesDataSource().Schema["managed_instance_collection"].Elem.(*schema.Resource).Schema)
		dynamicSetManagedInstance["items"] = items
	}

	resources = append(resources, dynamicSetManagedInstance)
	if err := s.D.Set("managed_instance_collection", resources); err != nil {
		return err
	}

	return nil
}
