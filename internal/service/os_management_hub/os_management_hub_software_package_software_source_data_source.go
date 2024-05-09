// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubSoftwarePackageSoftwareSourceDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsManagementHubSoftwarePackageSoftwareSource,
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
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"os_family": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"software_package_name": {
				Type:     schema.TypeString,
				Required: true,
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
			"software_source_collection": {
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
									"is_mandatory_for_autonomous_linux": {
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
									"software_source_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"software_source_version": {
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
									"url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"vendor_name": {
										Type:     schema.TypeString,
										Computed: true,
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
								},
							},
						},
					},
				},
			},
		},
	}
}

func readOsManagementHubSoftwarePackageSoftwareSource(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubSoftwarePackageSoftwareSourceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SoftwareSourceClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubSoftwarePackageSoftwareSourceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.SoftwareSourceClient
	Res    *oci_os_management_hub.ListSoftwarePackageSoftwareSourcesResponse
}

func (s *OsManagementHubSoftwarePackageSoftwareSourceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubSoftwarePackageSoftwareSourceDataSourceCrud) Get() error {
	request := oci_os_management_hub.ListSoftwarePackageSoftwareSourcesRequest{}

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
		if len(tmp) != 0 || s.D.HasChange("availability_at_oci") {
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

	if softwarePackageName, ok := s.D.GetOkExists("software_package_name"); ok {
		tmp := softwarePackageName.(string)
		request.SoftwarePackageName = &tmp
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.ListSoftwarePackageSoftwareSources(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSoftwarePackageSoftwareSources(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsManagementHubSoftwarePackageSoftwareSourceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubSoftwarePackageSoftwareSourceDataSource-", OsManagementHubSoftwarePackageSoftwareSourceDataSource(), s.D))
	resources := []map[string]interface{}{}
	softwarePackageSoftwareSource := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SoftwareSourceSummaryToMap(item))
	}
	softwarePackageSoftwareSource["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OsManagementHubSoftwarePackageSoftwareSourceDataSource().Schema["software_source_collection"].Elem.(*schema.Resource).Schema)
		softwarePackageSoftwareSource["items"] = items
	}

	resources = append(resources, softwarePackageSoftwareSource)
	if err := s.D.Set("software_source_collection", resources); err != nil {
		return err
	}

	return nil
}

func SoftwareSourceSummaryToMap(obj oci_os_management_hub.SoftwareSourceSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_os_management_hub.CustomSoftwareSourceSummary:
		result["software_source_type"] = "CUSTOM"

		vendorSoftwareSources := []interface{}{}
		for _, item := range v.VendorSoftwareSources {
			vendorSoftwareSources = append(vendorSoftwareSources, IdToMap(&item))
		}
		result["vendor_software_sources"] = vendorSoftwareSources

		result["arch_type"] = string(v.ArchType)

		result["availability"] = string(v.Availability)

		result["availability_at_oci"] = string(v.AvailabilityAtOci)

		if v.CompartmentId != nil {
			result["compartment_id"] = string(*v.CompartmentId)
		}

		if v.DefinedTags != nil {
			result["defined_tags"] = tfresource.DefinedTagsToMap(v.DefinedTags)
		}

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}

		if v.DisplayName != nil {
			result["display_name"] = string(*v.DisplayName)
		}

		result["freeform_tags"] = v.FreeformTags

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}

		result["os_family"] = string(v.OsFamily)

		if v.PackageCount != nil {
			result["package_count"] = strconv.FormatInt(*v.PackageCount, 10)
		}

		if v.RepoId != nil {
			result["repo_id"] = string(*v.RepoId)
		}

		if v.Size != nil {
			result["size"] = float64(*v.Size)
		}

		result["state"] = string(v.LifecycleState)

		if v.SystemTags != nil {
			result["system_tags"] = tfresource.SystemTagsToMap(v.SystemTags)
		}

		if v.TimeCreated != nil {
			result["time_created"] = v.TimeCreated.String()
		}

		if v.TimeUpdated != nil {
			result["time_updated"] = v.TimeUpdated.String()
		}

		if v.Url != nil {
			result["url"] = string(*v.Url)
		}
	case oci_os_management_hub.VendorSoftwareSourceSummary:
		result["software_source_type"] = "VENDOR"

		if v.IsMandatoryForAutonomousLinux != nil {
			result["is_mandatory_for_autonomous_linux"] = bool(*v.IsMandatoryForAutonomousLinux)
		}

		result["vendor_name"] = string(v.VendorName)

		result["arch_type"] = string(v.ArchType)

		result["availability"] = string(v.Availability)

		result["availability_at_oci"] = string(v.AvailabilityAtOci)

		if v.CompartmentId != nil {
			result["compartment_id"] = string(*v.CompartmentId)
		}

		if v.DefinedTags != nil {
			result["defined_tags"] = tfresource.DefinedTagsToMap(v.DefinedTags)
		}

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}

		if v.DisplayName != nil {
			result["display_name"] = string(*v.DisplayName)
		}

		result["freeform_tags"] = v.FreeformTags

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}

		result["os_family"] = string(v.OsFamily)

		if v.PackageCount != nil {
			result["package_count"] = strconv.FormatInt(*v.PackageCount, 10)
		}

		if v.RepoId != nil {
			result["repo_id"] = string(*v.RepoId)
		}

		if v.Size != nil {
			result["size"] = float64(*v.Size)
		}

		result["state"] = string(v.LifecycleState)

		if v.SystemTags != nil {
			result["system_tags"] = tfresource.SystemTagsToMap(v.SystemTags)
		}

		if v.TimeCreated != nil {
			result["time_created"] = v.TimeCreated.String()
		}

		if v.TimeUpdated != nil {
			result["time_updated"] = v.TimeUpdated.String()
		}

		if v.Url != nil {
			result["url"] = string(*v.Url)
		}
	case oci_os_management_hub.VersionedCustomSoftwareSourceSummary:
		result["software_source_type"] = "VERSIONED"

		if v.SoftwareSourceVersion != nil {
			result["software_source_version"] = string(*v.SoftwareSourceVersion)
		}

		vendorSoftwareSources := []interface{}{}
		for _, item := range v.VendorSoftwareSources {
			vendorSoftwareSources = append(vendorSoftwareSources, IdToMap(&item))
		}
		result["vendor_software_sources"] = vendorSoftwareSources

		result["arch_type"] = string(v.ArchType)

		result["availability"] = string(v.Availability)

		result["availability_at_oci"] = string(v.AvailabilityAtOci)

		if v.CompartmentId != nil {
			result["compartment_id"] = string(*v.CompartmentId)
		}

		if v.DefinedTags != nil {
			result["defined_tags"] = tfresource.DefinedTagsToMap(v.DefinedTags)
		}

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}

		if v.DisplayName != nil {
			result["display_name"] = string(*v.DisplayName)
		}

		result["freeform_tags"] = v.FreeformTags

		if v.Id != nil {
			result["id"] = string(*v.Id)
		}

		result["os_family"] = string(v.OsFamily)

		if v.PackageCount != nil {
			result["package_count"] = strconv.FormatInt(*v.PackageCount, 10)
		}

		if v.RepoId != nil {
			result["repo_id"] = string(*v.RepoId)
		}

		if v.Size != nil {
			result["size"] = float64(*v.Size)
		}

		result["state"] = string(v.LifecycleState)

		if v.SystemTags != nil {
			result["system_tags"] = tfresource.SystemTagsToMap(v.SystemTags)
		}

		if v.TimeCreated != nil {
			result["time_created"] = v.TimeCreated.String()
		}

		if v.TimeUpdated != nil {
			result["time_updated"] = v.TimeUpdated.String()
		}

		if v.Url != nil {
			result["url"] = string(*v.Url)
		}
	default:
		log.Printf("[WARN] Received 'software_source_type' of unknown type %v", obj)
		return nil
	}

	return result
}
