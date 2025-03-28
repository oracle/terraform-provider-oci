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

func OsManagementHubSoftwareSourcesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsManagementHubSoftwareSources,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"arch_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"availability": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"availability_anywhere": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"availability_at_oci": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
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
			"display_name_not_equal_to": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"is_mandatory_for_autonomous_linux": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_mirror_sync_allowed": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"os_family": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"software_source_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"software_source_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"state": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"vendor_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"software_source_collection": {
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
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"software_source_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"software_source_sub_type": {
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
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"software_source_version": {
										Type:     schema.TypeString,
										Computed: true,
										ForceNew: true,
									},
									"vendor_software_sources": {
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
									"arch_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"availability": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"availability_at_oci": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_mandatory_for_autonomous_linux": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_mirror_sync_allowed": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"os_family": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"package_count": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"repo_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"size": {
										Type:     schema.TypeFloat,
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
									"url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"vendor_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
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

func readOsManagementHubSoftwareSources(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubSoftwareSourcesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SoftwareSourceClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubSoftwareSourcesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.SoftwareSourceClient
	Res    *oci_os_management_hub.ListSoftwareSourcesResponse
}

func (s *OsManagementHubSoftwareSourcesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubSoftwareSourcesDataSourceCrud) Get() error {
	request := oci_os_management_hub.ListSoftwareSourcesRequest{}

	if archType, ok := s.D.GetOkExists("arch_type"); ok {
		interfaces := archType.([]interface{})
		tmp := make([]oci_os_management_hub.ArchTypeEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_os_management_hub.ArchTypeEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("arch_type") {
			request.ArchType = tmp
		}
	}

	if availability, ok := s.D.GetOkExists("availability"); ok {
		interfaces := availability.([]interface{})
		tmp := make([]oci_os_management_hub.AvailabilityEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_os_management_hub.AvailabilityEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("availability") {
			request.Availability = tmp
		}
	}

	if availabilityAnywhere, ok := s.D.GetOkExists("availability_anywhere"); ok {
		interfaces := availabilityAnywhere.([]interface{})
		tmp := make([]oci_os_management_hub.AvailabilityEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_os_management_hub.AvailabilityEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("availability_anywhere") {
			request.AvailabilityAnywhere = tmp
		}
	}

	if availabilityAtOci, ok := s.D.GetOkExists("availability_at_oci"); ok {
		interfaces := availabilityAtOci.([]interface{})
		tmp := make([]oci_os_management_hub.AvailabilityEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_os_management_hub.AvailabilityEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("availability") {
			request.AvailabilityAtOci = tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if displayNameContains, ok := s.D.GetOkExists("display_name_contains"); ok {
		tmp := displayNameContains.(string)
		request.DisplayNameContains = &tmp
	}

	if displayNameNotEqualTo, ok := s.D.GetOkExists("display_name_not_equal_to"); ok {
		interfaces := displayNameNotEqualTo.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("display_name_not_equal_to") {
			request.DisplayNameNotEqualTo = tmp
		}
	}

	if isMandatoryForAutonomousLinux, ok := s.D.GetOkExists("is_mandatory_for_autonomous_linux"); ok {
		tmp := isMandatoryForAutonomousLinux.(bool)
		request.IsMandatoryForAutonomousLinux = &tmp
	}

	if isMirrorSyncAllowed, ok := s.D.GetOkExists("is_mirror_sync_allowed"); ok {
		tmp := isMirrorSyncAllowed.(bool)
		request.IsMirrorSyncAllowed = &tmp
	}

	if osFamily, ok := s.D.GetOkExists("os_family"); ok {
		interfaces := osFamily.([]interface{})
		tmp := make([]oci_os_management_hub.OsFamilyEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_os_management_hub.OsFamilyEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("os_family") {
			request.OsFamily = tmp
		}
	}

	if softwareSourceId, ok := s.D.GetOkExists("id"); ok {
		tmp := softwareSourceId.(string)
		request.SoftwareSourceId = &tmp
	}

	if softwareSourceType, ok := s.D.GetOkExists("software_source_type"); ok {
		interfaces := softwareSourceType.([]interface{})
		tmp := make([]oci_os_management_hub.SoftwareSourceTypeEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_os_management_hub.SoftwareSourceTypeEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("software_source_type") {
			request.SoftwareSourceType = tmp
		}
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		interfaces := state.([]interface{})
		tmp := make([]oci_os_management_hub.SoftwareSourceLifecycleStateEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_os_management_hub.SoftwareSourceLifecycleStateEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("state") {
			request.LifecycleState = tmp
		}
	}

	if vendorName, ok := s.D.GetOkExists("vendor_name"); ok {
		request.VendorName = oci_os_management_hub.ListSoftwareSourcesVendorNameEnum(vendorName.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.ListSoftwareSources(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSoftwareSources(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsManagementHubSoftwareSourcesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubSoftwareSourcesDataSource-", OsManagementHubSoftwareSourcesDataSource(), s.D))
	resources := []map[string]interface{}{}
	softwareSource := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SoftwareSourceSummaryToMap(item))
	}
	softwareSource["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OsManagementHubSoftwareSourcesDataSource().Schema["software_source_collection"].Elem.(*schema.Resource).Schema)
		softwareSource["items"] = items
	}

	resources = append(resources, softwareSource)
	if err := s.D.Set("software_source_collection", resources); err != nil {
		return err
	}

	return nil
}
