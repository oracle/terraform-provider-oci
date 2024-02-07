// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_migration

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database_migration "github.com/oracle/oci-go-sdk/v65/databasemigration"
)

func DatabaseMigrationConnectionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,

		Create: createDatabaseMigrationConnection,
		Read:   readDatabaseMigrationConnection,
		Update: updateDatabaseMigrationConnection,
		Delete: deleteDatabaseMigrationConnection,
		Schema: map[string]*schema.Schema{
			// Required
			"admin_credentials": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"password": {
							Type:      schema.TypeString,
							Required:  true,
							Sensitive: true,
						},
						"username": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"database_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vault_details": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"compartment_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"key_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"vault_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},

			// Optional
			"certificate_tdn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connect_descriptor": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"connect_string": {
							Type:          schema.TypeString,
							Optional:      true,
							Computed:      true,
							ConflictsWith: []string{"connect_descriptor.0.database_service_name", "connect_descriptor.0.host", "connect_descriptor.0.port"},
						},
						"database_service_name": {
							Type:          schema.TypeString,
							Optional:      true,
							Computed:      true,
							ConflictsWith: []string{"connect_descriptor.0.connect_string"},
						},
						"host": {
							Type:          schema.TypeString,
							Optional:      true,
							Computed:      true,
							ConflictsWith: []string{"connect_descriptor.0.connect_string"},
						},
						"port": {
							Type:          schema.TypeInt,
							Optional:      true,
							Computed:      true,
							ConflictsWith: []string{"connect_descriptor.0.connect_string"},
						},

						// Computed
					},
				},
			},
			"database_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
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
			"manual_database_sub_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"nsg_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"private_endpoint": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"compartment_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"vcn_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"replication_credentials": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"password": {
							Type:      schema.TypeString,
							Required:  true,
							Sensitive: true,
						},
						"username": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"ssh_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"host": {
							Type:     schema.TypeString,
							Required: true,
						},
						"sshkey": {
							Type:     schema.TypeString,
							Required: true,
						},
						"user": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"sudo_location": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"tls_keystore": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tls_wallet": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"credentials_secret_id": {
				Type:     schema.TypeString,
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
		},
	}
}

func createDatabaseMigrationConnection(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseMigrationConnection(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseMigrationConnection(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseMigrationConnection(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseMigrationConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseMigrationConnectionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database_migration.DatabaseMigrationClient
	Res                    *oci_database_migration.Connection
	DisableNotFoundRetries bool
}

func (s *DatabaseMigrationConnectionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseMigrationConnectionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database_migration.LifecycleStatesCreating),
	}
}

func (s *DatabaseMigrationConnectionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database_migration.LifecycleStatesActive),
	}
}

func (s *DatabaseMigrationConnectionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database_migration.LifecycleStatesDeleting),
	}
}

func (s *DatabaseMigrationConnectionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database_migration.LifecycleStatesDeleted),
	}
}

func (s *DatabaseMigrationConnectionResourceCrud) Create() error {
	request := oci_database_migration.CreateConnectionRequest{}

	if adminCredentials, ok := s.D.GetOkExists("admin_credentials"); ok {
		if tmpList := adminCredentials.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "admin_credentials", 0)
			tmp, err := s.mapToCreateAdminCredentials(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.AdminCredentials = &tmp
		}
	}

	if certificateTdn, ok := s.D.GetOkExists("certificate_tdn"); ok {
		tmp := certificateTdn.(string)
		request.CertificateTdn = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if connectDescriptor, ok := s.D.GetOkExists("connect_descriptor"); ok {
		if tmpList := connectDescriptor.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connect_descriptor", 0)
			tmp, err := s.mapToCreateConnectDescriptor(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ConnectDescriptor = &tmp
		}
	}

	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
	}

	if databaseType, ok := s.D.GetOkExists("database_type"); ok {
		request.DatabaseType = oci_database_migration.DatabaseConnectionTypesEnum(databaseType.(string))
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if manualDatabaseSubType, ok := s.D.GetOkExists("manual_database_sub_type"); ok {
		request.ManualDatabaseSubType = oci_database_migration.DatabaseManualConnectionSubTypesEnum(manualDatabaseSubType.(string))
	}

	if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
			request.NsgIds = tmp
		}
	}

	if privateEndpoint, ok := s.D.GetOkExists("private_endpoint"); ok {
		if tmpList := privateEndpoint.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "private_endpoint", 0)
			tmp, err := s.mapToCreatePrivateEndpoint(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.PrivateEndpoint = &tmp
		}
	}

	if replicationCredentials, ok := s.D.GetOkExists("replication_credentials"); ok {
		if tmpList := replicationCredentials.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "replication_credentials", 0)
			tmp, err := s.mapToCreateAdminCredentials(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ReplicationCredentials = &tmp
		}
	}

	if sshDetails, ok := s.D.GetOkExists("ssh_details"); ok {
		if tmpList := sshDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "ssh_details", 0)
			tmp, err := s.mapToCreateSshDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SshDetails = &tmp
		}
	}

	if tlsKeystore, ok := s.D.GetOkExists("tls_keystore"); ok {
		tmp := tlsKeystore.(string)
		request.TlsKeystore = &tmp
	}

	if tlsWallet, ok := s.D.GetOkExists("tls_wallet"); ok {
		tmp := tlsWallet.(string)
		request.TlsWallet = &tmp
	}

	if vaultDetails, ok := s.D.GetOkExists("vault_details"); ok {
		if tmpList := vaultDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "vault_details", 0)
			tmp, err := s.mapToCreateVaultDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.VaultDetails = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration")

	response, err := s.Client.CreateConnection(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getConnectionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration"), oci_database_migration.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DatabaseMigrationConnectionResourceCrud) getConnectionFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,

	actionTypeEnum oci_database_migration.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	connectionId, err := connectionWaitForWorkRequest(workId, "connection",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*connectionId)

	return s.Get()
}

func connectionWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "database_migration", startTime) {
			return true
		}

		return false
	}
}

func connectionWaitForWorkRequest(wId *string, entityType string, action oci_database_migration.WorkRequestResourceActionTypeEnum,

	timeout time.Duration, disableFoundRetries bool, client *oci_database_migration.DatabaseMigrationClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "database_migration")
	retryPolicy.ShouldRetryOperation = connectionWorkRequestShouldRetryFunc(timeout)

	response := oci_database_migration.GetWorkRequestResponse{}

	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_database_migration.OperationStatusInProgress),
			string(oci_database_migration.OperationStatusAccepted),
			string(oci_database_migration.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_database_migration.OperationStatusSucceeded),
			string(oci_database_migration.OperationStatusFailed),
			string(oci_database_migration.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_database_migration.GetWorkRequestRequest{
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.WorkRequest
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}
	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_database_migration.OperationStatusFailed || response.Status == oci_database_migration.OperationStatusCanceled {
		return nil, getErrorFromDatabaseMigrationConnectionWorkRequest(client, wId, retryPolicy, entityType, action)
	}
	return identifier, nil
}

func getErrorFromDatabaseMigrationConnectionWorkRequest(client *oci_database_migration.DatabaseMigrationClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_database_migration.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_database_migration.ListWorkRequestErrorsRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		})
	if err != nil {
		return err
	}

	allErrs := make([]string, 0)
	for _, wrkErr := range response.Items {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *DatabaseMigrationConnectionResourceCrud) Get() error {
	request := oci_database_migration.GetConnectionRequest{}

	tmp := s.D.Id()
	request.ConnectionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration")

	response, err := s.Client.GetConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Connection
	return nil
}

func (s *DatabaseMigrationConnectionResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database_migration.UpdateConnectionRequest{}

	if adminCredentials, ok := s.D.GetOkExists("admin_credentials"); ok && s.D.HasChange("admin_credentials") {
		if tmpList := adminCredentials.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "admin_credentials", 0)
			tmp, err := s.mapToUpdateAdminCredentials(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.AdminCredentials = &tmp
		}
	}

	if certificateTdn, ok := s.D.GetOkExists("certificate_tdn"); ok {
		tmp := certificateTdn.(string)
		request.CertificateTdn = &tmp
	}

	if connectDescriptor, ok := s.D.GetOkExists("connect_descriptor"); ok {
		if tmpList := connectDescriptor.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connect_descriptor", 0)
			tmp, err := s.mapToUpdateConnectDescriptor(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ConnectDescriptor = &tmp
		}
	}

	tmp := s.D.Id()
	request.ConnectionId = &tmp

	if databaseId, ok := s.D.GetOkExists("database_id"); ok && s.D.HasChange("database_id") {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok && s.D.HasChange("defined_tags") {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok && s.D.HasChange("freeform_tags") {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
			request.NsgIds = tmp
		}
	}

	if privateEndpoint, ok := s.D.GetOkExists("private_endpoint"); ok {
		if tmpList := privateEndpoint.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "private_endpoint", 0)
			tmp, err := s.mapToUpdatePrivateEndpoint(fieldKeyFormat)

			if err != nil {
				return err
			}
			request.PrivateEndpoint = &tmp
		}
	}

	if replicationCredentials, ok := s.D.GetOkExists("replication_credentials"); ok {
		if tmpList := replicationCredentials.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "replication_credentials", 0)
			tmp, err := s.mapToUpdateAdminCredentials(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ReplicationCredentials = &tmp
		}
	}

	if sshDetails, ok := s.D.GetOkExists("ssh_details"); ok {
		if tmpList := sshDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "ssh_details", 0)
			tmp, err := s.mapToUpdateSshDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SshDetails = &tmp
		}
	}

	if tlsKeystore, ok := s.D.GetOkExists("tls_keystore"); ok {
		tmp := tlsKeystore.(string)
		request.TlsKeystore = &tmp
	}

	if tlsWallet, ok := s.D.GetOkExists("tls_wallet"); ok {
		tmp := tlsWallet.(string)
		request.TlsWallet = &tmp
	}

	if vaultDetails, ok := s.D.GetOkExists("vault_details"); ok && s.D.HasChange("vault_details") {
		if tmpList := vaultDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "vault_details", 0)
			tmp, err := s.mapToUpdateVaultDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.VaultDetails = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration")

	response, err := s.Client.UpdateConnection(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getConnectionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration"), oci_database_migration.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DatabaseMigrationConnectionResourceCrud) Delete() error {
	request := oci_database_migration.DeleteConnectionRequest{}

	tmp := s.D.Id()
	request.ConnectionId = &tmp
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration")

	response, err := s.Client.DeleteConnection(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := connectionWaitForWorkRequest(workId, "connection",
		oci_database_migration.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DatabaseMigrationConnectionResourceCrud) SetData() error {
	if s.Res.AdminCredentials != nil {
		s.D.Set("admin_credentials", []interface{}{AdminCredentialsToMapPassword(s.Res.AdminCredentials, s.D)})

	} else {
		s.D.Set("admin_credentials", nil)
	}

	if s.Res.CertificateTdn != nil {
		s.D.Set("certificate_tdn", *s.Res.CertificateTdn)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectDescriptor != nil {
		s.D.Set("connect_descriptor", []interface{}{ConnectDescriptorToMap(s.Res.ConnectDescriptor)})
	} else {
		s.D.Set("connect_descriptor", nil)
	}

	if s.Res.CredentialsSecretId != nil {
		s.D.Set("credentials_secret_id", *s.Res.CredentialsSecretId)
	}

	if s.Res.DatabaseId != nil {
		s.D.Set("database_id", *s.Res.DatabaseId)
	}

	s.D.Set("database_type", s.Res.DatabaseType)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	nsgIds := []interface{}{}
	for _, item := range s.Res.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	s.D.Set("nsg_ids", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds))
	s.D.Set("manual_database_sub_type", s.Res.ManualDatabaseSubType)

	if s.Res.PrivateEndpoint != nil {
		s.D.Set("private_endpoint", []interface{}{PrivateEndpointDetailsToMap(s.Res.PrivateEndpoint)})
	} else {
		s.D.Set("private_endpoint", nil)
	}

	if s.Res.ReplicationCredentials != nil {
		s.D.Set("replication_credentials", []interface{}{AdminCredentialsToMapPassword2(s.Res.ReplicationCredentials, s.D)})
	} else {
		s.D.Set("replication_credentials", nil)
	}

	if s.Res.SshDetails != nil {
		s.D.Set("ssh_details", []interface{}{SshDetailsToMapPass(s.Res.SshDetails, s.D)})

	} else {
		s.D.Set("ssh_details", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.VaultDetails != nil {
		s.D.Set("vault_details", []interface{}{VaultDetailsToMap(s.Res.VaultDetails)})
	} else {
		s.D.Set("vault_details", nil)
	}
	return nil
}

func ConnectionSummaryToMap(obj oci_database_migration.ConnectionSummary) map[string]interface{} {

	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DatabaseId != nil {
		result["database_id"] = string(*obj.DatabaseId)
	}

	result["database_type"] = string(obj.DatabaseType)

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

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	nsgIds := []interface{}{}
	for _, item := range obj.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	result["manual_database_sub_type"] = string(obj.ManualDatabaseSubType)

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *DatabaseMigrationConnectionResourceCrud) mapToCreateAdminCredentials(fieldKeyFormat string) (oci_database_migration.CreateAdminCredentials, error) {
	result := oci_database_migration.CreateAdminCredentials{}

	if password, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password")); ok {
		tmp := password.(string)
		result.Password = &tmp
	}

	if username, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "username")); ok {
		tmp := username.(string)
		result.Username = &tmp
	}

	return result, nil
}

func (s *DatabaseMigrationConnectionResourceCrud) mapToUpdateAdminCredentials(fieldKeyFormat string) (oci_database_migration.UpdateAdminCredentials, error) {
	result := oci_database_migration.UpdateAdminCredentials{}

	if password, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password")); ok {
		tmp := password.(string)
		result.Password = &tmp
	}

	if username, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "username")); ok {
		tmp := username.(string)
		result.Username = &tmp
	}

	return result, nil
}

func AdminCredentialsToMap(obj *oci_database_migration.AdminCredentials) map[string]interface{} {
	result := map[string]interface{}{}
	if obj.Username != nil {
		result["username"] = string(*obj.Username)
	}

	return result
}

func AdminCredentialsToMapPassword(obj *oci_database_migration.AdminCredentials, resourceData *schema.ResourceData) map[string]interface{} {
	result := map[string]interface{}{}
	if adminCredentialsValue, ok := resourceData.GetOkExists("admin_credentials"); ok {
		if tmpList := adminCredentialsValue.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "admin_credentials", 0)
			if adminPassword, ok := resourceData.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password")); ok {
				tmp := adminPassword.(string)
				result["password"] = &tmp
			}
		}
	}

	if obj.Username != nil {
		result["username"] = string(*obj.Username)
	}
	return result
}

func AdminCredentialsToMapPassword2(obj *oci_database_migration.AdminCredentials, resourceData *schema.ResourceData) map[string]interface{} {
	result := map[string]interface{}{}
	if adminPass, ok := resourceData.GetOkExists("replication_credentials.0.password"); ok && adminPass != nil {
		result["password"] = adminPass.(string)
	}

	if obj.Username != nil {
		result["username"] = string(*obj.Username)
	}
	return result
}

func AdminCredentialsToMapPasswordRest(obj *oci_database_migration.GoldenGateHub, resourceData *schema.ResourceData) map[string]interface{} {
	result := map[string]interface{}{}

	if adminPasswordR, ok := resourceData.GetOkExists("golden_gate_details.0.hub.0.rest_admin_credentials.0.password"); ok && adminPasswordR != nil {
		result["password"] = adminPasswordR.(string)
	}

	if obj.RestAdminCredentials != nil {
		result["username"] = string(*obj.RestAdminCredentials.Username)
	}
	return result
}

func AdminCredentialsToMapPasswordContainer(obj *oci_database_migration.GoldenGateHub, resourceData *schema.ResourceData) map[string]interface{} {
	result := map[string]interface{}{}

	if adminPasswordC, ok := resourceData.GetOkExists("golden_gate_details.0.hub.0.source_container_db_admin_credentials.0.password"); ok && adminPasswordC != nil {
		result["password"] = adminPasswordC.(string)
	}

	if obj.SourceContainerDbAdminCredentials != nil {
		result["username"] = string(*obj.SourceContainerDbAdminCredentials.Username)
	}
	return result
}

func AdminCredentialsToMapPasswordSource(obj *oci_database_migration.GoldenGateHub, resourceData *schema.ResourceData) map[string]interface{} {
	result := map[string]interface{}{}

	if adminPasswordS, ok := resourceData.GetOkExists("golden_gate_details.0.hub.0.source_db_admin_credentials.0.password"); ok && adminPasswordS != nil {
		result["password"] = adminPasswordS.(string)
	}
	if obj.SourceDbAdminCredentials != nil {
		result["username"] = string(*obj.SourceDbAdminCredentials.Username)
	}
	return result
}

func AdminCredentialsToMapPasswordTarget(obj *oci_database_migration.GoldenGateHub, resourceData *schema.ResourceData) map[string]interface{} {
	result := map[string]interface{}{}

	if adminPasswordT, ok := resourceData.GetOkExists("golden_gate_details.0.hub.0.target_db_admin_credentials.0.password"); ok && adminPasswordT != nil {
		result["password"] = adminPasswordT.(string)
	}

	if obj.TargetDbAdminCredentials != nil {
		result["username"] = string(*obj.TargetDbAdminCredentials.Username)
	}

	return result
}

func (s *DatabaseMigrationConnectionResourceCrud) mapToCreateConnectDescriptor(fieldKeyFormat string) (oci_database_migration.CreateConnectDescriptor, error) {
	result := oci_database_migration.CreateConnectDescriptor{}

	if connectString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connect_string")); ok {
		tmp := connectString.(string)
		result.ConnectString = &tmp
	}

	if databaseServiceName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database_service_name")); ok && s.D.HasChange("database_service_name") {
		tmp := databaseServiceName.(string)
		result.DatabaseServiceName = &tmp
	}

	if host, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "host")); ok && s.D.HasChange("host") {
		tmp := host.(string)
		result.Host = &tmp
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok && s.D.HasChange("port") {
		tmp := port.(int)
		result.Port = &tmp
	}

	return result, nil
}

func (s *DatabaseMigrationConnectionResourceCrud) mapToUpdateConnectDescriptor(fieldKeyFormat string) (oci_database_migration.UpdateConnectDescriptor, error) {
	result := oci_database_migration.UpdateConnectDescriptor{}

	if connectString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connect_string")); ok {
		tmp := connectString.(string)
		result.ConnectString = &tmp
	}

	if databaseServiceName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "database_service_name")); ok && s.D.HasChange("database_service_name") {
		tmp := databaseServiceName.(string)
		result.DatabaseServiceName = &tmp
	}

	if host, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "host")); ok && s.D.HasChange("host") {
		tmp := host.(string)
		result.Host = &tmp
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok && s.D.HasChange("port") {
		tmp := port.(int)
		result.Port = &tmp
	}

	return result, nil
}

func ConnectDescriptorToMap(obj *oci_database_migration.ConnectDescriptor) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConnectString != nil {
		result["connect_string"] = string(*obj.ConnectString)
	}

	if obj.DatabaseServiceName != nil {
		result["database_service_name"] = *obj.DatabaseServiceName
	}

	if obj.Host != nil {
		result["host"] = *obj.Host
	}

	if obj.Port != nil {
		result["port"] = *obj.Port
	}

	return result
}

func (s *DatabaseMigrationConnectionResourceCrud) mapToCreatePrivateEndpoint(fieldKeyFormat string) (oci_database_migration.CreatePrivateEndpoint, error) {
	result := oci_database_migration.CreatePrivateEndpoint{}

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		result.CompartmentId = &tmp
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	if vcnId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vcn_id")); ok {
		tmp := vcnId.(string)
		result.VcnId = &tmp
	}

	return result, nil
}

func (s *DatabaseMigrationConnectionResourceCrud) mapToUpdatePrivateEndpoint(fieldKeyFormat string) (oci_database_migration.UpdatePrivateEndpoint, error) {
	result := oci_database_migration.UpdatePrivateEndpoint{}

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		result.CompartmentId = &tmp
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	if vcnId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vcn_id")); ok {
		tmp := vcnId.(string)
		result.VcnId = &tmp
	}

	return result, nil
}

func PrivateEndpointDetailsToMap(obj *oci_database_migration.PrivateEndpointDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	if obj.VcnId != nil {
		result["vcn_id"] = string(*obj.VcnId)
	}

	return result
}

func (s *DatabaseMigrationConnectionResourceCrud) mapToCreateSshDetails(fieldKeyFormat string) (oci_database_migration.CreateSshDetails, error) {
	result := oci_database_migration.CreateSshDetails{}

	if host, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "host")); ok {
		tmp := host.(string)
		result.Host = &tmp
	}

	if sshkey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sshkey")); ok {
		tmp := sshkey.(string)
		result.Sshkey = &tmp
	}

	if sudoLocation, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sudo_location")); ok {
		tmp := sudoLocation.(string)
		result.SudoLocation = &tmp
	}

	if user, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user")); ok {
		tmp := user.(string)
		result.User = &tmp
	}

	return result, nil
}

func (s *DatabaseMigrationConnectionResourceCrud) mapToUpdateSshDetails(fieldKeyFormat string) (oci_database_migration.UpdateSshDetails, error) {
	result := oci_database_migration.UpdateSshDetails{}

	if host, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "host")); ok {
		tmp := host.(string)
		result.Host = &tmp
	}

	if sshkey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sshkey")); ok {
		tmp := sshkey.(string)
		result.Sshkey = &tmp
	}

	if sudoLocation, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sudo_location")); ok {
		tmp := sudoLocation.(string)
		result.SudoLocation = &tmp
	}

	if user, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user")); ok {
		tmp := user.(string)
		result.User = &tmp
	}

	return result, nil
}

func SshDetailsToMap(obj *oci_database_migration.SshDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Host != nil {
		result["host"] = string(*obj.Host)
	}
	if obj.SudoLocation != nil {
		result["sudo_location"] = string(*obj.SudoLocation)
	}

	if obj.User != nil {
		result["user"] = string(*obj.User)
	}

	return result
}

func SshDetailsToMapPass(obj *oci_database_migration.SshDetails, resourceData *schema.ResourceData) map[string]interface{} {
	result := map[string]interface{}{}

	if sshDetailsValue, ok := resourceData.GetOkExists("ssh_details"); ok {
		if tmpList := sshDetailsValue.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "ssh_details", 0)
			if sshKey, ok := resourceData.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sshkey")); ok {
				tmp := sshKey.(string)
				result["sshkey"] = &tmp
			}
		}
	}
	if obj.Host != nil {
		result["host"] = string(*obj.Host)
	}

	if obj.SudoLocation != nil {
		result["sudo_location"] = string(*obj.SudoLocation)
	}

	if obj.User != nil {
		result["user"] = string(*obj.User)
	}
	return result
}

func (s *DatabaseMigrationConnectionResourceCrud) mapToCreateVaultDetails(fieldKeyFormat string) (oci_database_migration.CreateVaultDetails, error) {
	result := oci_database_migration.CreateVaultDetails{}

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		result.CompartmentId = &tmp
	}

	if keyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_id")); ok {
		tmp := keyId.(string)
		result.KeyId = &tmp
	}

	if vaultId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_id")); ok {
		tmp := vaultId.(string)
		result.VaultId = &tmp
	}

	return result, nil
}

func (s *DatabaseMigrationConnectionResourceCrud) mapToUpdateVaultDetails(fieldKeyFormat string) (oci_database_migration.UpdateVaultDetails, error) {
	result := oci_database_migration.UpdateVaultDetails{}

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		result.CompartmentId = &tmp
	}

	if keyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_id")); ok {
		tmp := keyId.(string)
		result.KeyId = &tmp
	}

	if vaultId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_id")); ok {
		tmp := vaultId.(string)
		result.VaultId = &tmp
	}

	return result, nil
}
func VaultDetailsToMap(obj *oci_database_migration.VaultDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.KeyId != nil {
		result["key_id"] = string(*obj.KeyId)
	}

	if obj.VaultId != nil {
		result["vault_id"] = string(*obj.VaultId)
	}

	return result
}

func (s *DatabaseMigrationConnectionResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database_migration.ChangeConnectionCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ConnectionId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration")

	_, err := s.Client.ChangeConnectionCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
