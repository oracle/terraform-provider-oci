// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/hashicorp/terraform/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/database"
)

func DatabaseExadataInfrastructureResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createDatabaseExadataInfrastructure,
		Read:     readDatabaseExadataInfrastructure,
		Update:   updateDatabaseExadataInfrastructure,
		Delete:   deleteDatabaseExadataInfrastructure,
		Schema: map[string]*schema.Schema{
			// Required
			"admin_network_cidr": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cloud_control_plane_server1": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cloud_control_plane_server2": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"corporate_proxy": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dns_server": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"gateway": {
				Type:     schema.TypeString,
				Required: true,
			},
			"infini_band_network_cidr": {
				Type:     schema.TypeString,
				Required: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ntp_server": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"shape": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"time_zone": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"activation_file": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"cpus_enabled": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"data_storage_size_in_tbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseExadataInfrastructure(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExadataInfrastructureResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return CreateResource(d, sync)
}

func readDatabaseExadataInfrastructure(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExadataInfrastructureResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return ReadResource(sync)
}

func updateDatabaseExadataInfrastructure(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExadataInfrastructureResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return UpdateResource(d, sync)
}

func deleteDatabaseExadataInfrastructure(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExadataInfrastructureResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type DatabaseExadataInfrastructureResourceCrud struct {
	BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.ExadataInfrastructure
	DisableNotFoundRetries bool
}

func (s *DatabaseExadataInfrastructureResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseExadataInfrastructureResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.ExadataInfrastructureLifecycleStateCreating),
		string(oci_database.ExadataInfrastructureLifecycleStateActivating),
	}
}

func (s *DatabaseExadataInfrastructureResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.ExadataInfrastructureLifecycleStateRequiresActivation),
		string(oci_database.ExadataInfrastructureLifecycleStateActive),
	}
}

func (s *DatabaseExadataInfrastructureResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.ExadataInfrastructureLifecycleStateDeleting),
		"TERMINATING",
	}
}

func (s *DatabaseExadataInfrastructureResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.ExadataInfrastructureLifecycleStateDeleted),
		"TERMINATED",
	}
}

func (s *DatabaseExadataInfrastructureResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_database.ExadataInfrastructureLifecycleStateActivating),
		string(oci_database.ExadataInfrastructureLifecycleStateUpdating),
	}
}

func (s *DatabaseExadataInfrastructureResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_database.ExadataInfrastructureLifecycleStateRequiresActivation),
		string(oci_database.ExadataInfrastructureLifecycleStateActive),
		string(oci_database.ExadataInfrastructureLifecycleStateActivationFailed),
		string(oci_database.ExadataInfrastructureLifecycleStateOffline),
	}
}

func (s *DatabaseExadataInfrastructureResourceCrud) Create() error {
	request := oci_database.CreateExadataInfrastructureRequest{}

	if adminNetworkCIDR, ok := s.D.GetOkExists("admin_network_cidr"); ok {
		tmp := adminNetworkCIDR.(string)
		request.AdminNetworkCIDR = &tmp
	}

	if cloudControlPlaneServer1, ok := s.D.GetOkExists("cloud_control_plane_server1"); ok {
		tmp := cloudControlPlaneServer1.(string)
		request.CloudControlPlaneServer1 = &tmp
	}

	if cloudControlPlaneServer2, ok := s.D.GetOkExists("cloud_control_plane_server2"); ok {
		tmp := cloudControlPlaneServer2.(string)
		request.CloudControlPlaneServer2 = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if corporateProxy, ok := s.D.GetOkExists("corporate_proxy"); ok {
		tmp := corporateProxy.(string)
		request.CorporateProxy = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if dnsServer, ok := s.D.GetOkExists("dns_server"); ok {
		request.DnsServer = []string{}
		interfaces := dnsServer.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.DnsServer = tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if gateway, ok := s.D.GetOkExists("gateway"); ok {
		tmp := gateway.(string)
		request.Gateway = &tmp
	}

	if infiniBandNetworkCIDR, ok := s.D.GetOkExists("infini_band_network_cidr"); ok {
		tmp := infiniBandNetworkCIDR.(string)
		request.InfiniBandNetworkCIDR = &tmp
	}

	if netmask, ok := s.D.GetOkExists("netmask"); ok {
		tmp := netmask.(string)
		request.Netmask = &tmp
	}

	if ntpServer, ok := s.D.GetOkExists("ntp_server"); ok {
		request.NtpServer = []string{}
		interfaces := ntpServer.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.NtpServer = tmp
	}

	if shape, ok := s.D.GetOkExists("shape"); ok {
		tmp := shape.(string)
		request.Shape = &tmp
	}

	if timeZone, ok := s.D.GetOkExists("time_zone"); ok {
		tmp := timeZone.(string)
		request.TimeZone = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateExadataInfrastructure(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExadataInfrastructure

	//wait for Exadata Infrastructure to not be in creating state
	if waitErr := waitForCreatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	if activationFile, ok := s.D.GetOkExists("activation_file"); ok {
		response, err := s.activateExadataInfrastructure(activationFile.(string), s.D.Id())
		if err != nil {
			return err
		}
		s.Res = &response.ExadataInfrastructure
	}
	return nil
}

func (s *DatabaseExadataInfrastructureResourceCrud) Get() error {
	request := oci_database.GetExadataInfrastructureRequest{}

	tmp := s.D.Id()
	request.ExadataInfrastructureId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetExadataInfrastructure(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExadataInfrastructure
	return nil
}

func (s *DatabaseExadataInfrastructureResourceCrud) Update() error {

	if s.D.Get("state").(string) == string(oci_database.ExadataInfrastructureLifecycleStateActive) {
		return fmt.Errorf("update not allowed on activated exadata infrastructure")
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

	request := oci_database.UpdateExadataInfrastructureRequest{}

	if adminNetworkCIDR, ok := s.D.GetOkExists("admin_network_cidr"); ok && s.D.HasChange("admin_network_cidr") {
		tmp := adminNetworkCIDR.(string)
		request.AdminNetworkCIDR = &tmp
	}

	if cloudControlPlaneServer1, ok := s.D.GetOkExists("cloud_control_plane_server1"); ok && s.D.HasChange("cloud_control_plane_server1") {
		tmp := cloudControlPlaneServer1.(string)
		request.CloudControlPlaneServer1 = &tmp
	}

	if cloudControlPlaneServer2, ok := s.D.GetOkExists("cloud_control_plane_server2"); ok && s.D.HasChange("cloud_control_plane_server2") {
		tmp := cloudControlPlaneServer2.(string)
		request.CloudControlPlaneServer2 = &tmp
	}

	if corporateProxy, ok := s.D.GetOkExists("corporate_proxy"); ok && s.D.HasChange("corporate_proxy") {
		tmp := corporateProxy.(string)
		request.CorporateProxy = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if dnsServer, ok := s.D.GetOkExists("dns_server"); ok && s.D.HasChange("dns_server") {
		request.DnsServer = []string{}
		interfaces := dnsServer.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.DnsServer = tmp
	}

	tmp := s.D.Id()
	request.ExadataInfrastructureId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if gateway, ok := s.D.GetOkExists("gateway"); ok && s.D.HasChange("gateway") {
		tmp := gateway.(string)
		request.Gateway = &tmp
	}

	if infiniBandNetworkCIDR, ok := s.D.GetOkExists("infini_band_network_cidr"); ok && s.D.HasChange("infini_band_network_cidr") {
		tmp := infiniBandNetworkCIDR.(string)
		request.InfiniBandNetworkCIDR = &tmp
	}

	if netmask, ok := s.D.GetOkExists("netmask"); ok && s.D.HasChange("netmask") {
		tmp := netmask.(string)
		request.Netmask = &tmp
	}

	if ntpServer, ok := s.D.GetOkExists("ntp_server"); ok && s.D.HasChange("ntp_server") {
		request.NtpServer = []string{}
		interfaces := ntpServer.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.NtpServer = tmp
	}

	if timeZone, ok := s.D.GetOkExists("time_zone"); ok && s.D.HasChange("time_zone") {
		tmp := timeZone.(string)
		request.TimeZone = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateExadataInfrastructure(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ExadataInfrastructure

	//wait for Exadata Infrastructure to not be in updating state after the update
	if waitErr := waitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	if activationFile, ok := s.D.GetOkExists("activation_file"); ok {
		response, err := s.activateExadataInfrastructure(activationFile.(string), s.D.Id())
		if err != nil {
			return err
		}
		s.Res = &response.ExadataInfrastructure
	}
	return nil
}

func (s *DatabaseExadataInfrastructureResourceCrud) Delete() error {
	request := oci_database.DeleteExadataInfrastructureRequest{}

	tmp := s.D.Id()
	request.ExadataInfrastructureId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.DeleteExadataInfrastructure(context.Background(), request)
	return err
}

func (s *DatabaseExadataInfrastructureResourceCrud) SetData() error {
	if s.Res.AdminNetworkCIDR != nil {
		s.D.Set("admin_network_cidr", *s.Res.AdminNetworkCIDR)
	}

	if s.Res.CloudControlPlaneServer1 != nil {
		s.D.Set("cloud_control_plane_server1", *s.Res.CloudControlPlaneServer1)
	}

	if s.Res.CloudControlPlaneServer2 != nil {
		s.D.Set("cloud_control_plane_server2", *s.Res.CloudControlPlaneServer2)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CorporateProxy != nil {
		s.D.Set("corporate_proxy", *s.Res.CorporateProxy)
	}

	if s.Res.CpusEnabled != nil {
		s.D.Set("cpus_enabled", *s.Res.CpusEnabled)
	}

	if s.Res.DataStorageSizeInTBs != nil {
		s.D.Set("data_storage_size_in_tbs", *s.Res.DataStorageSizeInTBs)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("dns_server", s.Res.DnsServer)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Gateway != nil {
		s.D.Set("gateway", *s.Res.Gateway)
	}

	if s.Res.InfiniBandNetworkCIDR != nil {
		s.D.Set("infini_band_network_cidr", *s.Res.InfiniBandNetworkCIDR)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Netmask != nil {
		s.D.Set("netmask", *s.Res.Netmask)
	}

	s.D.Set("ntp_server", s.Res.NtpServer)

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeZone != nil {
		s.D.Set("time_zone", *s.Res.TimeZone)
	}

	return nil
}

func (s *DatabaseExadataInfrastructureResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database.ChangeExadataInfrastructureCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ExadataInfrastructureId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.ChangeExadataInfrastructureCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}

func (s *DatabaseExadataInfrastructureResourceCrud) activateExadataInfrastructure(activationFile string, exadataInfrastructureId string) (*oci_database.ActivateExadataInfrastructureResponse, error) {
	request := oci_database.ActivateExadataInfrastructureRequest{}

	activationKeyFile, err := ioutil.ReadFile(activationFile)
	if err != nil {
		return nil, fmt.Errorf("unable to open activation key file: %s", err)
	}
	request.ActivationFile = activationKeyFile

	request.ExadataInfrastructureId = &exadataInfrastructureId

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.ActivateExadataInfrastructure(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
