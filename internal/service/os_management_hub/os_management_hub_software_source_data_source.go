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

func OsManagementHubSoftwareSourceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["software_source_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OsManagementHubSoftwareSourceResource(), fieldMap, readSingularOsManagementHubSoftwareSource)
}

func readSingularOsManagementHubSoftwareSource(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubSoftwareSourceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SoftwareSourceClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubSoftwareSourceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.SoftwareSourceClient
	Res    *oci_os_management_hub.GetSoftwareSourceResponse
}

func (s *OsManagementHubSoftwareSourceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubSoftwareSourceDataSourceCrud) Get() error {
	request := oci_os_management_hub.GetSoftwareSourceRequest{}

	if softwareSourceId, ok := s.D.GetOkExists("software_source_id"); ok {
		tmp := softwareSourceId.(string)
		request.SoftwareSourceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.GetSoftwareSource(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OsManagementHubSoftwareSourceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.GetId())
	switch v := (s.Res.SoftwareSource).(type) {
	case oci_os_management_hub.CustomSoftwareSource:
		s.D.Set("software_source_type", "CUSTOM")

		if v.CustomSoftwareSourceFilter != nil {
			s.D.Set("custom_software_source_filter", []interface{}{CustomSoftwareSourceFilterToMap(v.CustomSoftwareSourceFilter)})
		} else {
			s.D.Set("custom_software_source_filter", nil)
		}

		if v.IsAutoResolveDependencies != nil {
			s.D.Set("is_auto_resolve_dependencies", *v.IsAutoResolveDependencies)
		}

		if v.IsAutomaticallyUpdated != nil {
			s.D.Set("is_automatically_updated", *v.IsAutomaticallyUpdated)
		}

		if v.IsCreatedFromPackageList != nil {
			s.D.Set("is_created_from_package_list", *v.IsCreatedFromPackageList)
		}

		if v.IsLatestContentOnly != nil {
			s.D.Set("is_latest_content_only", *v.IsLatestContentOnly)
		}

		s.D.Set("packages", v.Packages)

		vendorSoftwareSources := []interface{}{}
		for _, item := range v.VendorSoftwareSources {
			vendorSoftwareSources = append(vendorSoftwareSources, IdToMap(&item))
		}
		s.D.Set("vendor_software_sources", vendorSoftwareSources)

		s.D.Set("arch_type", v.ArchType)

		s.D.Set("availability", v.Availability)

		s.D.Set("availability_at_oci", v.AvailabilityAtOci)

		s.D.Set("checksum_type", v.ChecksumType)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.GpgKeyFingerprint != nil {
			s.D.Set("gpg_key_fingerprint", *v.GpgKeyFingerprint)
		}

		if v.GpgKeyId != nil {
			s.D.Set("gpg_key_id", *v.GpgKeyId)
		}

		if v.GpgKeyUrl != nil {
			s.D.Set("gpg_key_url", *v.GpgKeyUrl)
		}

		s.D.Set("os_family", v.OsFamily)

		if v.PackageCount != nil {
			s.D.Set("package_count", strconv.FormatInt(*v.PackageCount, 10))
		}

		if v.RepoId != nil {
			s.D.Set("repo_id", *v.RepoId)
		}

		if v.Size != nil {
			s.D.Set("size", *v.Size)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.Url != nil {
			s.D.Set("url", *v.Url)
		}
	case oci_os_management_hub.VendorSoftwareSource:
		s.D.Set("software_source_type", "VENDOR")

		if v.IsMandatoryForAutonomousLinux != nil {
			s.D.Set("is_mandatory_for_autonomous_linux", *v.IsMandatoryForAutonomousLinux)
		}

		if v.OriginSoftwareSourceId != nil {
			s.D.Set("origin_software_source_id", *v.OriginSoftwareSourceId)
		}

		s.D.Set("vendor_name", v.VendorName)

		s.D.Set("arch_type", v.ArchType)

		s.D.Set("availability", v.Availability)

		s.D.Set("availability_at_oci", v.AvailabilityAtOci)

		s.D.Set("checksum_type", v.ChecksumType)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.GpgKeyFingerprint != nil {
			s.D.Set("gpg_key_fingerprint", *v.GpgKeyFingerprint)
		}

		if v.GpgKeyId != nil {
			s.D.Set("gpg_key_id", *v.GpgKeyId)
		}

		if v.GpgKeyUrl != nil {
			s.D.Set("gpg_key_url", *v.GpgKeyUrl)
		}

		if v.IsMandatoryForAutonomousLinux != nil {
			s.D.Set("is_mandatory_for_autonomous_linux", *v.IsMandatoryForAutonomousLinux)
		}

		s.D.Set("os_family", v.OsFamily)

		if v.PackageCount != nil {
			s.D.Set("package_count", strconv.FormatInt(*v.PackageCount, 10))
		}

		if v.RepoId != nil {
			s.D.Set("repo_id", *v.RepoId)
		}

		if v.Size != nil {
			s.D.Set("size", *v.Size)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.Url != nil {
			s.D.Set("url", *v.Url)
		}

		s.D.Set("vendor_name", v.VendorName)
	case oci_os_management_hub.VersionedCustomSoftwareSource:
		s.D.Set("software_source_type", "VERSIONED")

		if v.CustomSoftwareSourceFilter != nil {
			s.D.Set("custom_software_source_filter", []interface{}{CustomSoftwareSourceFilterToMap(v.CustomSoftwareSourceFilter)})
		} else {
			s.D.Set("custom_software_source_filter", nil)
		}

		if v.IsAutoResolveDependencies != nil {
			s.D.Set("is_auto_resolve_dependencies", *v.IsAutoResolveDependencies)
		}

		if v.IsCreatedFromPackageList != nil {
			s.D.Set("is_created_from_package_list", *v.IsCreatedFromPackageList)
		}

		if v.IsLatestContentOnly != nil {
			s.D.Set("is_latest_content_only", *v.IsLatestContentOnly)
		}

		s.D.Set("packages", v.Packages)

		if v.SoftwareSourceVersion != nil {
			s.D.Set("software_source_version", *v.SoftwareSourceVersion)
		}

		vendorSoftwareSources := []interface{}{}
		for _, item := range v.VendorSoftwareSources {
			vendorSoftwareSources = append(vendorSoftwareSources, IdToMap(&item))
		}
		s.D.Set("vendor_software_sources", vendorSoftwareSources)

		s.D.Set("arch_type", v.ArchType)

		s.D.Set("availability", v.Availability)

		s.D.Set("availability_at_oci", v.AvailabilityAtOci)

		s.D.Set("checksum_type", v.ChecksumType)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.GpgKeyFingerprint != nil {
			s.D.Set("gpg_key_fingerprint", *v.GpgKeyFingerprint)
		}

		if v.GpgKeyId != nil {
			s.D.Set("gpg_key_id", *v.GpgKeyId)
		}

		if v.GpgKeyUrl != nil {
			s.D.Set("gpg_key_url", *v.GpgKeyUrl)
		}

		s.D.Set("os_family", v.OsFamily)

		if v.PackageCount != nil {
			s.D.Set("package_count", strconv.FormatInt(*v.PackageCount, 10))
		}

		if v.RepoId != nil {
			s.D.Set("repo_id", *v.RepoId)
		}

		if v.Size != nil {
			s.D.Set("size", *v.Size)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.Url != nil {
			s.D.Set("url", *v.Url)
		}
	default:
		log.Printf("[WARN] Received 'software_source_type' of unknown type %v", s.Res.SoftwareSource)
		return nil
	}

	return nil
}
