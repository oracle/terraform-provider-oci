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

func OsManagementHubLifecycleStagesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsManagementHubLifecycleStages,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"arch_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"display_name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"lifecycle_stage_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"location": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"location_not_equal_to": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"os_family": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"software_source_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"lifecycle_stage_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

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
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_environment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_environment_display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"location": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"managed_instances": {
										Type:     schema.TypeInt,
										Computed: true,
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
												"is_mandatory_for_autonomous_linux": {
													Type:     schema.TypeBool,
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
							},
						},
					},
				},
			},
		},
	}
}

func readOsManagementHubLifecycleStages(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubLifecycleStagesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LifecycleEnvironmentClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubLifecycleStagesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.LifecycleEnvironmentClient
	Res    *oci_os_management_hub.ListLifecycleStagesResponse
}

func (s *OsManagementHubLifecycleStagesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubLifecycleStagesDataSourceCrud) Get() error {
	request := oci_os_management_hub.ListLifecycleStagesRequest{}

	if archType, ok := s.D.GetOkExists("arch_type"); ok {
		request.ArchType = oci_os_management_hub.ListLifecycleStagesArchTypeEnum(archType.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		interfaces := displayName.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("display_name") {
			request.DisplayName = tmp
		}
	}

	if displayNameContains, ok := s.D.GetOkExists("display_name_contains"); ok {
		tmp := displayNameContains.(string)
		request.DisplayNameContains = &tmp
	}

	if lifecycleStageId, ok := s.D.GetOkExists("id"); ok {
		tmp := lifecycleStageId.(string)
		request.LifecycleStageId = &tmp
	}

	if location, ok := s.D.GetOkExists("location"); ok {
		interfaces := location.([]interface{})
		tmp := make([]oci_os_management_hub.ManagedInstanceLocationEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_os_management_hub.ManagedInstanceLocationEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("location_not_equal_to") {
			request.Location = tmp
		}
	}

	if locationNotEqualTo, ok := s.D.GetOkExists("location_not_equal_to"); ok {
		interfaces := locationNotEqualTo.([]interface{})
		tmp := make([]oci_os_management_hub.ManagedInstanceLocationEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_os_management_hub.ManagedInstanceLocationEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("location_not_equal_to") {
			request.LocationNotEqualTo = tmp
		}
	}

	if osFamily, ok := s.D.GetOkExists("os_family"); ok {
		request.OsFamily = oci_os_management_hub.ListLifecycleStagesOsFamilyEnum(osFamily.(string))
	}

	if softwareSourceId, ok := s.D.GetOkExists("software_source_id"); ok {
		tmp := softwareSourceId.(string)
		request.SoftwareSourceId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_os_management_hub.LifecycleStageLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.ListLifecycleStages(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListLifecycleStages(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsManagementHubLifecycleStagesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubLifecycleStagesDataSource-", OsManagementHubLifecycleStagesDataSource(), s.D))
	resources := []map[string]interface{}{}
	lifecycleStage := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, LifecycleStageSummaryToMap(item))
	}
	lifecycleStage["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OsManagementHubLifecycleStagesDataSource().Schema["lifecycle_stage_collection"].Elem.(*schema.Resource).Schema)
		lifecycleStage["items"] = items
	}

	resources = append(resources, lifecycleStage)
	if err := s.D.Set("lifecycle_stage_collection", resources); err != nil {
		return err
	}

	return nil
}

func LifecycleStageSummaryToMap(obj oci_os_management_hub.LifecycleStageSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["arch_type"] = string(obj.ArchType)

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleEnvironmentDisplayName != nil {
		result["lifecycle_environment_display_name"] = string(*obj.LifecycleEnvironmentDisplayName)
	}

	if obj.LifecycleEnvironmentId != nil {
		result["lifecycle_environment_id"] = string(*obj.LifecycleEnvironmentId)
	}

	result["location"] = string(obj.Location)

	if obj.ManagedInstances != nil {
		result["managed_instances"] = int(*obj.ManagedInstances)
		//result["managed_instance_ids"] = []string{}
	}

	result["os_family"] = string(obj.OsFamily)

	if obj.Rank != nil {
		result["rank"] = int(*obj.Rank)
	}

	if obj.SoftwareSourceId != nil {
		result["software_source_id"] = []interface{}{SoftwareSourceDetailsToMap(*obj.SoftwareSourceId)}
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeModified != nil {
		result["time_modified"] = obj.TimeModified.String()
	}

	result["vendor_name"] = string(obj.VendorName)

	return result
}

func ManagedInstanceDetailsToMap(obj oci_os_management_hub.ManagedInstanceDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}
