// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubManagementStationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOsManagementHubManagementStation,
		Read:     readOsManagementHubManagementStation,
		Update:   updateOsManagementHubManagementStation,
		Delete:   deleteOsManagementHubManagementStation,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mirror": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"directory": {
							Type:     schema.TypeString,
							Required: true,
						},
						"port": {
							Type:     schema.TypeString,
							Required: true,
						},
						"sslport": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"is_sslverify_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"sslcert": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"proxy": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"is_enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},

						// Optional
						"forward": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"hosts": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"port": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_auto_config_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"refresh_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			// Computed
			"health": {
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
						"state": {
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
			"managed_instance_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mirror_capacity": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"mirror_package_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"mirror_size": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mirror_storage_available_size": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mirror_storage_size": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mirror_sync_status": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"failed": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"queued": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"synced": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"syncing": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"unsynced": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"mirror_unique_package_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"overall_percentage": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"overall_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_management_stations": {
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
			"profile_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scheduled_job_id": {
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
			"total_mirrors": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func createOsManagementHubManagementStation(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubManagementStationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementStationClient()

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if _, ok := sync.D.GetOkExists("refresh_trigger"); ok {
		err := sync.RefreshManagementStationConfig()
		if err != nil {
			return err
		}
	}
	return nil

}

func readOsManagementHubManagementStation(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubManagementStationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementStationClient()

	return tfresource.ReadResource(sync)
}

func updateOsManagementHubManagementStation(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubManagementStationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementStationClient()

	if _, ok := sync.D.GetOkExists("refresh_trigger"); ok && sync.D.HasChange("refresh_trigger") {
		oldRaw, newRaw := sync.D.GetChange("refresh_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.RefreshManagementStationConfig()

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("refresh_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	return nil
}

func deleteOsManagementHubManagementStation(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubManagementStationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementStationClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type OsManagementHubManagementStationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_os_management_hub.ManagementStationClient
	Res                    *oci_os_management_hub.ManagementStation
	DisableNotFoundRetries bool
}

func (s *OsManagementHubManagementStationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OsManagementHubManagementStationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_os_management_hub.ManagementStationLifecycleStateCreating),
	}
}

func (s *OsManagementHubManagementStationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_os_management_hub.ManagementStationLifecycleStateActive),
	}
}

func (s *OsManagementHubManagementStationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_os_management_hub.ManagementStationLifecycleStateDeleting),
	}
}

func (s *OsManagementHubManagementStationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_os_management_hub.ManagementStationLifecycleStateDeleted),
	}
}

func (s *OsManagementHubManagementStationResourceCrud) Create() error {
	request := oci_os_management_hub.CreateManagementStationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if hostname, ok := s.D.GetOkExists("hostname"); ok {
		tmp := hostname.(string)
		request.Hostname = &tmp
	}

	if isAutoConfigEnabled, ok := s.D.GetOkExists("is_auto_config_enabled"); ok {
		tmp := isAutoConfigEnabled.(bool)
		request.IsAutoConfigEnabled = &tmp
	}

	if mirror, ok := s.D.GetOkExists("mirror"); ok {
		if tmpList := mirror.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "mirror", 0)
			tmp, err := s.mapToCreateMirrorConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Mirror = &tmp
		}
	}

	if proxy, ok := s.D.GetOkExists("proxy"); ok {
		if tmpList := proxy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "proxy", 0)
			tmp, err := s.mapToCreateProxyConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Proxy = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	response, err := s.Client.CreateManagementStation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagementStation
	return nil
}

func (s *OsManagementHubManagementStationResourceCrud) Get() error {
	request := oci_os_management_hub.GetManagementStationRequest{}

	tmp := s.D.Id()
	request.ManagementStationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	response, err := s.Client.GetManagementStation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagementStation
	return nil
}

func (s *OsManagementHubManagementStationResourceCrud) Update() error {

	if _, ok := s.D.GetOkExists("compartmentId"); ok && s.D.HasChange("compartmentId") {
		err := s.ChangeManagementStationCompartment()
		if err != nil {
			return err
		}
	}
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_os_management_hub.UpdateManagementStationRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if hostname, ok := s.D.GetOkExists("hostname"); ok {
		tmp := hostname.(string)
		request.Hostname = &tmp
	}

	if isAutoConfigEnabled, ok := s.D.GetOkExists("is_auto_config_enabled"); ok {
		tmp := isAutoConfigEnabled.(bool)
		request.IsAutoConfigEnabled = &tmp
	}

	tmp := s.D.Id()
	request.ManagementStationId = &tmp

	if mirror, ok := s.D.GetOkExists("mirror"); ok {
		if tmpList := mirror.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "mirror", 0)
			tmp, err := s.mapToUpdateMirrorConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Mirror = &tmp
		}
	}

	if proxy, ok := s.D.GetOkExists("proxy"); ok {
		if tmpList := proxy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "proxy", 0)
			tmp, err := s.mapToUpdateProxyConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Proxy = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	response, err := s.Client.UpdateManagementStation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagementStation
	return nil
}

func (s *OsManagementHubManagementStationResourceCrud) Delete() error {
	request := oci_os_management_hub.DeleteManagementStationRequest{}

	tmp := s.D.Id()
	request.ManagementStationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err := s.Client.DeleteManagementStation(context.Background(), request)
	return err
}

func (s *OsManagementHubManagementStationResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Health != nil {
		s.D.Set("health", []interface{}{StationHealthToMap(s.Res.Health)})
	} else {
		s.D.Set("health", nil)
	}

	if s.Res.Hostname != nil {
		s.D.Set("hostname", *s.Res.Hostname)
	}

	if s.Res.IsAutoConfigEnabled != nil {
		s.D.Set("is_auto_config_enabled", *s.Res.IsAutoConfigEnabled)
	}

	s.D.Set("location", s.Res.Location)

	if s.Res.ManagedInstanceId != nil {
		s.D.Set("managed_instance_id", *s.Res.ManagedInstanceId)
	}

	if s.Res.Mirror != nil {
		s.D.Set("mirror", []interface{}{MirrorConfigurationToMap(s.Res.Mirror)})
	} else {
		s.D.Set("mirror", nil)
	}

	if s.Res.MirrorCapacity != nil {
		s.D.Set("mirror_capacity", *s.Res.MirrorCapacity)
	}

	if s.Res.MirrorPackageCount != nil {
		s.D.Set("mirror_package_count", *s.Res.MirrorPackageCount)
	}

	if s.Res.MirrorSize != nil {
		s.D.Set("mirror_size", strconv.FormatInt(*s.Res.MirrorSize, 10))
	}

	if s.Res.MirrorStorageAvailableSize != nil {
		s.D.Set("mirror_storage_available_size", strconv.FormatInt(*s.Res.MirrorStorageAvailableSize, 10))
	}

	if s.Res.MirrorStorageSize != nil {
		s.D.Set("mirror_storage_size", strconv.FormatInt(*s.Res.MirrorStorageSize, 10))
	}

	if s.Res.MirrorSyncStatus != nil {
		s.D.Set("mirror_sync_status", []interface{}{MirrorSyncStatusToMap(s.Res.MirrorSyncStatus)})
	} else {
		s.D.Set("mirror_sync_status", nil)
	}

	if s.Res.MirrorUniquePackageCount != nil {
		s.D.Set("mirror_unique_package_count", *s.Res.MirrorUniquePackageCount)
	}

	if s.Res.OverallPercentage != nil {
		s.D.Set("overall_percentage", *s.Res.OverallPercentage)
	}

	s.D.Set("overall_state", s.Res.OverallState)

	peerManagementStations := []interface{}{}
	for _, item := range s.Res.PeerManagementStations {
		peerManagementStations = append(peerManagementStations, PeerManagementStationToMap(item))
	}
	s.D.Set("peer_management_stations", peerManagementStations)

	if s.Res.ProfileId != nil {
		s.D.Set("profile_id", *s.Res.ProfileId)
	}

	if s.Res.Proxy != nil {
		s.D.Set("proxy", []interface{}{ProxyConfigurationToMap(s.Res.Proxy)})
	} else {
		s.D.Set("proxy", nil)
	}

	if s.Res.ScheduledJobId != nil {
		s.D.Set("scheduled_job_id", *s.Res.ScheduledJobId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TotalMirrors != nil {
		s.D.Set("total_mirrors", *s.Res.TotalMirrors)
	}

	return nil
}

func (s *OsManagementHubManagementStationResourceCrud) RefreshManagementStationConfig() error {
	request := oci_os_management_hub.RefreshManagementStationConfigRequest{}

	idTmp := s.D.Id()
	request.ManagementStationId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err := s.Client.RefreshManagementStationConfig(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("refresh_trigger")
	s.D.Set("refresh_trigger", val)

	return s.Get()
}

func (s *OsManagementHubManagementStationResourceCrud) ChangeManagementStationCompartment() error {
	request := oci_os_management_hub.ChangeManagementStationCompartmentRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	idTmp := s.D.Id()
	request.ManagementStationId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err := s.Client.ChangeManagementStationCompartment(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *OsManagementHubManagementStationResourceCrud) mapToCreateMirrorConfigurationDetails(fieldKeyFormat string) (oci_os_management_hub.CreateMirrorConfigurationDetails, error) {
	result := oci_os_management_hub.CreateMirrorConfigurationDetails{}

	if directory, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "directory")); ok {
		tmp := directory.(string)
		result.Directory = &tmp
	}

	if isSslverifyEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_sslverify_enabled")); ok {
		tmp := isSslverifyEnabled.(bool)
		result.IsSslverifyEnabled = &tmp
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(string)
		result.Port = &tmp
	}

	if sslcert, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sslcert")); ok {
		tmp := sslcert.(string)
		result.Sslcert = &tmp
	}

	if sslport, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sslport")); ok {
		tmp := sslport.(string)
		result.Sslport = &tmp
	}

	return result, nil
}

func (s *OsManagementHubManagementStationResourceCrud) mapToUpdateMirrorConfigurationDetails(fieldKeyFormat string) (oci_os_management_hub.UpdateMirrorConfigurationDetails, error) {
	result := oci_os_management_hub.UpdateMirrorConfigurationDetails{}

	if directory, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "directory")); ok {
		tmp := directory.(string)
		result.Directory = &tmp
	}

	if isSslverifyEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_sslverify_enabled")); ok {
		tmp := isSslverifyEnabled.(bool)
		result.IsSslverifyEnabled = &tmp
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(string)
		result.Port = &tmp
	}

	if sslcert, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sslcert")); ok {
		tmp := sslcert.(string)
		result.Sslcert = &tmp
	}

	if sslport, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sslport")); ok {
		tmp := sslport.(string)
		result.Sslport = &tmp
	}

	return result, nil
}

func MirrorConfigurationToMap(obj *oci_os_management_hub.MirrorConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Directory != nil {
		result["directory"] = string(*obj.Directory)
	}

	if obj.IsSslverifyEnabled != nil {
		result["is_sslverify_enabled"] = bool(*obj.IsSslverifyEnabled)
	}

	if obj.Port != nil {
		result["port"] = string(*obj.Port)
	}

	if obj.Sslcert != nil {
		result["sslcert"] = string(*obj.Sslcert)
	}

	if obj.Sslport != nil {
		result["sslport"] = string(*obj.Sslport)
	}

	return result
}

func (s *OsManagementHubManagementStationResourceCrud) mapToCreateProxyConfigurationDetails(fieldKeyFormat string) (oci_os_management_hub.CreateProxyConfigurationDetails, error) {
	result := oci_os_management_hub.CreateProxyConfigurationDetails{}

	if forward, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "forward")); ok {
		tmp := forward.(string)
		result.Forward = &tmp
	}

	if hosts, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hosts")); ok {
		interfaces := hosts.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "hosts")) {
			result.Hosts = tmp
		}
	}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(string)
		result.Port = &tmp
	}

	return result, nil
}

func (s *OsManagementHubManagementStationResourceCrud) mapToUpdateProxyConfigurationDetails(fieldKeyFormat string) (oci_os_management_hub.UpdateProxyConfigurationDetails, error) {
	result := oci_os_management_hub.UpdateProxyConfigurationDetails{}

	if forward, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "forward")); ok {
		tmp := forward.(string)
		result.Forward = &tmp
	}

	if hosts, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hosts")); ok {
		interfaces := hosts.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "hosts")) {
			result.Hosts = tmp
		}
	}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(string)
		result.Port = &tmp
	}

	return result, nil
}

func ProxyConfigurationToMap(obj *oci_os_management_hub.ProxyConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Forward != nil {
		result["forward"] = string(*obj.Forward)
	}

	result["hosts"] = obj.Hosts

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	if obj.Port != nil {
		result["port"] = string(*obj.Port)
	}

	return result
}

func ManagementStationSummaryToMap(obj oci_os_management_hub.ManagementStationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	result["health_state"] = string(obj.HealthState)

	if obj.Hostname != nil {
		result["hostname"] = string(*obj.Hostname)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["location"] = string(obj.Location)

	if obj.ManagedInstanceId != nil {
		result["managed_instance_id"] = string(*obj.ManagedInstanceId)
	}

	if obj.MirrorCapacity != nil {
		result["mirror_capacity"] = int(*obj.MirrorCapacity)
	}

	if obj.OverallPercentage != nil {
		result["overall_percentage"] = int(*obj.OverallPercentage)
	}

	result["overall_state"] = string(obj.OverallState)

	if obj.ProfileId != nil {
		result["profile_id"] = string(*obj.ProfileId)
	}

	if obj.ScheduledJobId != nil {
		result["scheduled_job_id"] = string(*obj.ScheduledJobId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	//if obj.TimeNextExecution != nil {
	//	result["time_next_execution"] = obj.TimeNextExecution.String()
	//}

	return result
}

func MirrorSyncStatusToMap(obj *oci_os_management_hub.MirrorSyncStatus) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Failed != nil {
		result["failed"] = int(*obj.Failed)
	}

	if obj.Queued != nil {
		result["queued"] = int(*obj.Queued)
	}

	if obj.Synced != nil {
		result["synced"] = int(*obj.Synced)
	}

	if obj.Syncing != nil {
		result["syncing"] = int(*obj.Syncing)
	}

	if obj.Unsynced != nil {
		result["unsynced"] = int(*obj.Unsynced)
	}

	return result
}

func PeerManagementStationToMap(obj oci_os_management_hub.PeerManagementStation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}

func StationHealthToMap(obj *oci_os_management_hub.StationHealth) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["state"] = string(obj.State)

	return result
}

func (s *OsManagementHubManagementStationResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_os_management_hub.ChangeManagementStationCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ManagementStationId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err := s.Client.ChangeManagementStationCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
