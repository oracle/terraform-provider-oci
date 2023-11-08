// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity_domains

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IdentityDomainsMyDevicesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityDomainsMyDevices,
		Schema: map[string]*schema.Schema{
			"idcs_endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"my_device_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"my_device_filter": {
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
			"my_devices": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(IdentityDomainsMyDeviceDataSource()),
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

func readIdentityDomainsMyDevices(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsMyDevicesDataSourceCrud{}
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

type IdentityDomainsMyDevicesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity_domains.IdentityDomainsClient
	Res    *oci_identity_domains.ListMyDevicesResponse
}

func (s *IdentityDomainsMyDevicesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityDomainsMyDevicesDataSourceCrud) Get() error {
	request := oci_identity_domains.ListMyDevicesRequest{}

	if myDeviceCount, ok := s.D.GetOkExists("my_device_count"); ok {
		tmp := myDeviceCount.(int)
		request.Count = &tmp
	}

	if myDeviceFilter, ok := s.D.GetOkExists("my_device_filter"); ok {
		tmp := myDeviceFilter.(string)
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity_domains")

	response, err := s.Client.ListMyDevices(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	// IDCS pagination
	startIndex := *response.StartIndex
	for startIndex+*response.ItemsPerPage <= *response.TotalResults {
		startIndex += *response.ItemsPerPage
		request.StartIndex = &startIndex
		listResponse, err := s.Client.ListMyDevices(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Resources = append(s.Res.Resources, listResponse.Resources...)
	}

	return nil
}

func (s *IdentityDomainsMyDevicesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentityDomainsMyDevicesDataSource-", IdentityDomainsMyDevicesDataSource(), s.D))

	resources := []interface{}{}
	for _, item := range s.Res.Resources {
		resources = append(resources, MyDeviceToMap(item))
	}
	s.D.Set("my_devices", resources)

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

func MyDeviceToMap(obj oci_identity_domains.MyDevice) map[string]interface{} {
	result := map[string]interface{}{}

	additionalAttributes := []interface{}{}
	for _, item := range obj.AdditionalAttributes {
		additionalAttributes = append(additionalAttributes, MyDeviceAdditionalAttributesToMap(item))
	}
	result["additional_attributes"] = additionalAttributes

	if obj.AppVersion != nil {
		result["app_version"] = string(*obj.AppVersion)
	}

	authenticationFactors := []interface{}{}
	for _, item := range obj.AuthenticationFactors {
		authenticationFactors = append(authenticationFactors, MyDeviceAuthenticationFactorsToMap(item))
	}
	result["authentication_factors"] = authenticationFactors

	if obj.AuthenticationMethod != nil {
		result["authentication_method"] = string(*obj.AuthenticationMethod)
	}

	if obj.BasePublicKey != nil {
		result["base_public_key"] = string(*obj.BasePublicKey)
	}

	if obj.CompartmentOcid != nil {
		result["compartment_ocid"] = string(*obj.CompartmentOcid)
	}

	if obj.CountryCode != nil {
		result["country_code"] = string(*obj.CountryCode)
	}

	if obj.DeleteInProgress != nil {
		result["delete_in_progress"] = bool(*obj.DeleteInProgress)
	}

	if obj.DeviceType != nil {
		result["device_type"] = string(*obj.DeviceType)
	}

	if obj.DeviceUUID != nil {
		result["device_uuid"] = string(*obj.DeviceUUID)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.DomainOcid != nil {
		result["domain_ocid"] = string(*obj.DomainOcid)
	}

	if obj.ExpiresOn != nil {
		result["expires_on"] = int(*obj.ExpiresOn)
	}

	if obj.ExternalId != nil {
		result["external_id"] = string(*obj.ExternalId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IdcsCreatedBy != nil {
		result["idcs_created_by"] = []interface{}{idcsCreatedByToMap(obj.IdcsCreatedBy)}
	}

	if obj.IdcsLastModifiedBy != nil {
		result["idcs_last_modified_by"] = []interface{}{idcsLastModifiedByToMap(obj.IdcsLastModifiedBy)}
	}

	if obj.IdcsLastUpgradedInRelease != nil {
		result["idcs_last_upgraded_in_release"] = string(*obj.IdcsLastUpgradedInRelease)
	}

	result["idcs_prevented_operations"] = obj.IdcsPreventedOperations

	if obj.IsAccRecEnabled != nil {
		result["is_acc_rec_enabled"] = bool(*obj.IsAccRecEnabled)
	}

	if obj.IsCompliant != nil {
		result["is_compliant"] = bool(*obj.IsCompliant)
	}

	if obj.LastSyncTime != nil {
		result["last_sync_time"] = string(*obj.LastSyncTime)
	}

	if obj.LastValidatedTime != nil {
		result["last_validated_time"] = string(*obj.LastValidatedTime)
	}

	if obj.Meta != nil {
		result["meta"] = []interface{}{metaToMap(obj.Meta)}
	}

	nonCompliances := []interface{}{}
	for _, item := range obj.NonCompliances {
		nonCompliances = append(nonCompliances, MyDeviceNonCompliancesToMap(item))
	}
	result["non_compliances"] = nonCompliances

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.PackageId != nil {
		result["package_id"] = string(*obj.PackageId)
	}

	if obj.PhoneNumber != nil {
		result["phone_number"] = string(*obj.PhoneNumber)
	}

	result["platform"] = string(obj.Platform)

	if obj.PushNotificationTarget != nil {
		result["push_notification_target"] = []interface{}{MyDevicePushNotificationTargetToMap(obj.PushNotificationTarget)}
	}

	if obj.Reason != nil {
		result["reason"] = string(*obj.Reason)
	}

	result["schemas"] = obj.Schemas

	if obj.Seed != nil {
		result["seed"] = string(*obj.Seed)
	}

	if obj.SeedDekId != nil {
		result["seed_dek_id"] = string(*obj.SeedDekId)
	}

	result["status"] = string(obj.Status)

	tags := []interface{}{}
	for _, item := range obj.Tags {
		tags = append(tags, tagsToMap(item))
	}
	result["tags"] = tags

	if obj.TenancyOcid != nil {
		result["tenancy_ocid"] = string(*obj.TenancyOcid)
	}

	if obj.ThirdPartyFactor != nil {
		result["third_party_factor"] = []interface{}{MyDeviceThirdPartyFactorToMap(obj.ThirdPartyFactor)}
	}

	if obj.User != nil {
		result["user"] = []interface{}{MyDeviceUserToMap(obj.User)}
	}

	return result
}

func MyDeviceAdditionalAttributesToMap(obj oci_identity_domains.MyDeviceAdditionalAttributes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func MyDeviceAuthenticationFactorsToMap(obj oci_identity_domains.MyDeviceAuthenticationFactors) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.PublicKey != nil {
		result["public_key"] = string(*obj.PublicKey)
	}

	result["status"] = string(obj.Status)

	result["type"] = string(obj.Type)

	return result
}

func MyDeviceNonCompliancesToMap(obj oci_identity_domains.MyDeviceNonCompliances) map[string]interface{} {
	result := map[string]interface{}{}

	result["action"] = string(obj.Action)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func MyDevicePushNotificationTargetToMap(obj *oci_identity_domains.MyDevicePushNotificationTarget) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func MyDeviceThirdPartyFactorToMap(obj *oci_identity_domains.MyDeviceThirdPartyFactor) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.ThirdPartyFactorType != nil {
		result["third_party_factor_type"] = string(*obj.ThirdPartyFactorType)
	}

	if obj.ThirdPartyVendorName != nil {
		result["third_party_vendor_name"] = string(*obj.ThirdPartyVendorName)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func MyDeviceUserToMap(obj *oci_identity_domains.MyDeviceUser) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}
