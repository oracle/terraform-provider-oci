// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package license_manager

import (
	"context"
	"net/url"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_license_manager "github.com/oracle/oci-go-sdk/v65/licensemanager"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LicenseManagerConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createLicenseManagerConfiguration,
		Read:     readLicenseManagerConfiguration,
		Update:   updateLicenseManagerConfiguration,
		Delete:   deleteLicenseManagerConfiguration,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"email_ids": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional

			// Computed
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createLicenseManagerConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &LicenseManagerConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LicenseManagerClient()

	return tfresource.CreateResource(d, sync)
}

func readLicenseManagerConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &LicenseManagerConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LicenseManagerClient()

	return tfresource.ReadResource(sync)
}

func updateLicenseManagerConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &LicenseManagerConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LicenseManagerClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteLicenseManagerConfiguration(d *schema.ResourceData, m interface{}) error {
	return nil
}

type LicenseManagerConfigurationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_license_manager.LicenseManagerClient
	Res                    *oci_license_manager.Configuration
	DisableNotFoundRetries bool
}

func (s *LicenseManagerConfigurationResourceCrud) ID() string {
	return GetConfigurationId(*s.Res.CompartmentId)
}

func (s *LicenseManagerConfigurationResourceCrud) Create() error {
	request := oci_license_manager.UpdateConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if emailIds, ok := s.D.GetOkExists("email_ids"); ok {
		interfaces := emailIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("email_ids") {
			request.EmailIds = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "license_manager")

	response, err := s.Client.UpdateConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Configuration
	return nil
}

func (s *LicenseManagerConfigurationResourceCrud) Get() error {
	request := oci_license_manager.GetConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	compartmentId := s.D.Id()
	request.CompartmentId = &compartmentId

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "license_manager")

	response, err := s.Client.GetConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Configuration
	return nil
}

func (s *LicenseManagerConfigurationResourceCrud) Update() error {
	request := oci_license_manager.UpdateConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if emailIds, ok := s.D.GetOkExists("email_ids"); ok {
		interfaces := emailIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("email_ids") {
			request.EmailIds = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "license_manager")

	response, err := s.Client.UpdateConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Configuration
	return nil
}

func (s *LicenseManagerConfigurationResourceCrud) SetData() error {

	compartmentId := s.D.Id()
	s.D.Set("compartment_id", &compartmentId)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("email_ids", s.Res.EmailIds)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func GetConfigurationId(compartmentId string) string {
	compartmentId = url.PathEscape(compartmentId)
	return compartmentId
}
