// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package capacity_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_capacity_management "github.com/oracle/oci-go-sdk/v65/capacitymanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CapacityManagementOccmDemandSignalCatalogResourcesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCapacityManagementOccmDemandSignalCatalogResources,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"demand_signal_namespace": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"occm_demand_signal_catalog_resource_collection": {
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
									"availability_domain": {
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
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"region": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_properties": {
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
															"is_editable": {
																Type:     schema.TypeBool,
																Computed: true,
															},
															"property_max_value": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"property_min_value": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"property_name": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"property_options": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional

																		// Computed
																		"option_key": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"option_value": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"property_unit": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"property_value": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"resource_property_constraints": {
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
															"constraint_name": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"constraint_value": {
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
									"target_compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
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
				},
			},
		},
	}
}

func readCapacityManagementOccmDemandSignalCatalogResources(d *schema.ResourceData, m interface{}) error {
	sync := &CapacityManagementOccmDemandSignalCatalogResourcesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DemandSignalClient()

	return tfresource.ReadResource(sync)
}

type CapacityManagementOccmDemandSignalCatalogResourcesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_capacity_management.DemandSignalClient
	Res    *oci_capacity_management.ListOccmDemandSignalCatalogResourcesResponse
}

func (s *CapacityManagementOccmDemandSignalCatalogResourcesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CapacityManagementOccmDemandSignalCatalogResourcesDataSourceCrud) Get() error {
	request := oci_capacity_management.ListOccmDemandSignalCatalogResourcesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if demandSignalNamespace, ok := s.D.GetOkExists("demand_signal_namespace"); ok {
		request.DemandSignalNamespace = oci_capacity_management.ListOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum(demandSignalNamespace.(string))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "capacity_management")

	response, err := s.Client.ListOccmDemandSignalCatalogResources(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOccmDemandSignalCatalogResources(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CapacityManagementOccmDemandSignalCatalogResourcesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CapacityManagementOccmDemandSignalCatalogResourcesDataSource-", CapacityManagementOccmDemandSignalCatalogResourcesDataSource(), s.D))
	resources := []map[string]interface{}{}
	occmDemandSignalCatalogResource := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OccmDemandSignalCatalogResourceSummaryToMap(item))
	}
	occmDemandSignalCatalogResource["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CapacityManagementOccmDemandSignalCatalogResourcesDataSource().Schema["occm_demand_signal_catalog_resource_collection"].Elem.(*schema.Resource).Schema)
		occmDemandSignalCatalogResource["items"] = items
	}

	resources = append(resources, occmDemandSignalCatalogResource)
	if err := s.D.Set("occm_demand_signal_catalog_resource_collection", resources); err != nil {
		return err
	}

	return nil
}

func OccmDemandSignalCatalogResourceSummaryToMap(obj oci_capacity_management.OccmDemandSignalCatalogResourceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["namespace"] = string(obj.Namespace)

	if obj.Region != nil {
		result["region"] = string(*obj.Region)
	}

	if obj.ResourceProperties != nil {
		result["resource_properties"] = []interface{}{OccmDemandSignalResourcePropertiesCollectionToMap(obj.ResourceProperties)}
	}

	if obj.ResourcePropertyConstraints != nil {
		result["resource_property_constraints"] = []interface{}{OccmDemandSignalResourcePropertyConstraintsCollectionToMap(obj.ResourcePropertyConstraints)}
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TargetCompartmentId != nil {
		result["target_compartment_id"] = string(*obj.TargetCompartmentId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
