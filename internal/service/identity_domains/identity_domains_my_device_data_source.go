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

func IdentityDomainsMyDeviceDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularIdentityDomainsMyDevice,
		Schema: map[string]*schema.Schema{
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
			"idcs_endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},
			"my_device_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_type_schema_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"additional_attributes": {
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
			"app_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"authentication_factors": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"public_key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"authentication_method": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"base_public_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"country_code": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"delete_in_progress": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"device_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"device_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"expires_on": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"external_id": {
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
			"is_acc_rec_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_compliant": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"last_sync_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_validated_time": {
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
			"non_compliances": {
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
						"name": {
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
			"ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"package_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"phone_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"platform": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"push_notification_target": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
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
			"reason": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"schemas": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"seed": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"seed_dek_id": {
				Type:     schema.TypeString,
				Computed: true,
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
			"third_party_factor": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"third_party_factor_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"third_party_vendor_name": {
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
			"user": {
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
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularIdentityDomainsMyDevice(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsMyDeviceDataSourceCrud{}
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

type IdentityDomainsMyDeviceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity_domains.IdentityDomainsClient
	Res    *oci_identity_domains.GetMyDeviceResponse
}

func (s *IdentityDomainsMyDeviceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityDomainsMyDeviceDataSourceCrud) Get() error {
	request := oci_identity_domains.GetMyDeviceRequest{}

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

	if myDeviceId, ok := s.D.GetOkExists("my_device_id"); ok {
		tmp := myDeviceId.(string)
		request.MyDeviceId = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity_domains")

	response, err := s.Client.GetMyDevice(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IdentityDomainsMyDeviceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	additionalAttributes := []interface{}{}
	for _, item := range s.Res.AdditionalAttributes {
		additionalAttributes = append(additionalAttributes, MyDeviceAdditionalAttributesToMap(item))
	}
	s.D.Set("additional_attributes", additionalAttributes)

	if s.Res.AppVersion != nil {
		s.D.Set("app_version", *s.Res.AppVersion)
	}

	authenticationFactors := []interface{}{}
	for _, item := range s.Res.AuthenticationFactors {
		authenticationFactors = append(authenticationFactors, MyDeviceAuthenticationFactorsToMap(item))
	}
	s.D.Set("authentication_factors", authenticationFactors)

	if s.Res.AuthenticationMethod != nil {
		s.D.Set("authentication_method", *s.Res.AuthenticationMethod)
	}

	if s.Res.BasePublicKey != nil {
		s.D.Set("base_public_key", *s.Res.BasePublicKey)
	}

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	if s.Res.CountryCode != nil {
		s.D.Set("country_code", *s.Res.CountryCode)
	}

	if s.Res.DeleteInProgress != nil {
		s.D.Set("delete_in_progress", *s.Res.DeleteInProgress)
	}

	if s.Res.DeviceType != nil {
		s.D.Set("device_type", *s.Res.DeviceType)
	}

	if s.Res.DeviceUUID != nil {
		s.D.Set("device_uuid", *s.Res.DeviceUUID)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DomainOcid != nil {
		s.D.Set("domain_ocid", *s.Res.DomainOcid)
	}

	if s.Res.ExpiresOn != nil {
		s.D.Set("expires_on", *s.Res.ExpiresOn)
	}

	if s.Res.ExternalId != nil {
		s.D.Set("external_id", *s.Res.ExternalId)
	}

	if s.Res.IdcsCreatedBy != nil {
		s.D.Set("idcs_created_by", []interface{}{idcsCreatedByToMap(s.Res.IdcsCreatedBy)})
	} else {
		s.D.Set("idcs_created_by", nil)
	}

	if s.Res.IdcsLastModifiedBy != nil {
		s.D.Set("idcs_last_modified_by", []interface{}{idcsLastModifiedByToMap(s.Res.IdcsLastModifiedBy)})
	} else {
		s.D.Set("idcs_last_modified_by", nil)
	}

	if s.Res.IdcsLastUpgradedInRelease != nil {
		s.D.Set("idcs_last_upgraded_in_release", *s.Res.IdcsLastUpgradedInRelease)
	}

	s.D.Set("idcs_prevented_operations", s.Res.IdcsPreventedOperations)

	if s.Res.IsAccRecEnabled != nil {
		s.D.Set("is_acc_rec_enabled", *s.Res.IsAccRecEnabled)
	}

	if s.Res.IsCompliant != nil {
		s.D.Set("is_compliant", *s.Res.IsCompliant)
	}

	if s.Res.LastSyncTime != nil {
		s.D.Set("last_sync_time", *s.Res.LastSyncTime)
	}

	if s.Res.LastValidatedTime != nil {
		s.D.Set("last_validated_time", *s.Res.LastValidatedTime)
	}

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	nonCompliances := []interface{}{}
	for _, item := range s.Res.NonCompliances {
		nonCompliances = append(nonCompliances, MyDeviceNonCompliancesToMap(item))
	}
	s.D.Set("non_compliances", nonCompliances)

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	if s.Res.PackageId != nil {
		s.D.Set("package_id", *s.Res.PackageId)
	}

	if s.Res.PhoneNumber != nil {
		s.D.Set("phone_number", *s.Res.PhoneNumber)
	}

	s.D.Set("platform", s.Res.Platform)

	if s.Res.PushNotificationTarget != nil {
		s.D.Set("push_notification_target", []interface{}{MyDevicePushNotificationTargetToMap(s.Res.PushNotificationTarget)})
	} else {
		s.D.Set("push_notification_target", nil)
	}

	if s.Res.Reason != nil {
		s.D.Set("reason", *s.Res.Reason)
	}

	s.D.Set("schemas", s.Res.Schemas)

	if s.Res.Seed != nil {
		s.D.Set("seed", *s.Res.Seed)
	}

	if s.Res.SeedDekId != nil {
		s.D.Set("seed_dek_id", *s.Res.SeedDekId)
	}

	s.D.Set("status", s.Res.Status)

	tags := []interface{}{}
	for _, item := range s.Res.Tags {
		tags = append(tags, tagsToMap(item))
	}
	s.D.Set("tags", tags)

	if s.Res.TenancyOcid != nil {
		s.D.Set("tenancy_ocid", *s.Res.TenancyOcid)
	}

	if s.Res.ThirdPartyFactor != nil {
		s.D.Set("third_party_factor", []interface{}{MyDeviceThirdPartyFactorToMap(s.Res.ThirdPartyFactor)})
	} else {
		s.D.Set("third_party_factor", nil)
	}

	if s.Res.User != nil {
		s.D.Set("user", []interface{}{MyDeviceUserToMap(s.Res.User)})
	} else {
		s.D.Set("user", nil)
	}

	return nil
}
