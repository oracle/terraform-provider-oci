// // Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// // Licensed under the Mozilla Public License v2.0
package database_migration

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database_migration "github.com/oracle/oci-go-sdk/v65/databasemigration"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseMigrationConnectionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseMigrationConnection,
		Read:     readDatabaseMigrationConnection,
		Update:   updateDatabaseMigrationConnection,
		Delete:   deleteDatabaseMigrationConnection,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"connection_type": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"MYSQL",
					"ORACLE",
				}, true),
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"key_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"technology_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vault_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"additional_attributes": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"connection_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"database_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"database_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"db_system_id": {
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
			"host": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"replication_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},
			"replication_username": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_protocol": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_host": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_sudo_location": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_user": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_ca": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_cert": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_crl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wallet": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"ingress_ips": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"ingress_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_endpoint_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secret_id": {
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
	connection := *s.Res
	return *connection.GetId()
}

func (s *DatabaseMigrationConnectionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database_migration.ConnectionLifecycleStateCreating),
	}
}

func (s *DatabaseMigrationConnectionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database_migration.ConnectionLifecycleStateActive),
	}
}

func (s *DatabaseMigrationConnectionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database_migration.ConnectionLifecycleStateDeleting),
	}
}

func (s *DatabaseMigrationConnectionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database_migration.ConnectionLifecycleStateDeleted),
	}
}

func (s *DatabaseMigrationConnectionResourceCrud) Create() error {
	request := oci_database_migration.CreateConnectionRequest{}
	err := s.populateTopLevelPolymorphicCreateConnectionRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database_migration")

	response, err := s.Client.CreateConnection(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.GetId()
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

		// Only stop if the time Finished is set
		/*if workRequestResponse, ok := response.Response.(oci_database_migration.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}*/
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
	err := s.populateTopLevelPolymorphicUpdateConnectionRequest(&request)
	if err != nil {
		return err
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
	switch v := (*s.Res).(type) {
	case oci_database_migration.MysqlConnection:
		s.D.Set("connection_type", "MYSQL")

		additionalAttributes := []interface{}{}
		for _, item := range v.AdditionalAttributes {
			additionalAttributes = append(additionalAttributes, NameValuePairToMap(item))
		}
		s.D.Set("additional_attributes", additionalAttributes)

		if v.DatabaseName != nil {
			s.D.Set("database_name", *v.DatabaseName)
		}

		if v.DbSystemId != nil {
			s.D.Set("db_system_id", *v.DbSystemId)
		}

		if v.Host != nil {
			s.D.Set("host", *v.Host)
		}

		if v.Port != nil {
			s.D.Set("port", *v.Port)
		}

		s.D.Set("security_protocol", v.SecurityProtocol)

		s.D.Set("ssl_mode", v.SslMode)

		s.D.Set("technology_type", v.TechnologyType)

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

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		//s.D.Set("nsg_ids", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds))
		s.D.Set("nsg_ids", nsgIds)

		if v.Password != nil {
			s.D.Set("password", *v.Password)
		}

		if v.PrivateEndpointId != nil {
			s.D.Set("private_endpoint_id", *v.PrivateEndpointId)
		}

		if v.ReplicationPassword != nil {
			s.D.Set("replication_password", *v.ReplicationPassword)
		}

		if v.ReplicationUsername != nil {
			s.D.Set("replication_username", *v.ReplicationUsername)
		}

		if v.SecretId != nil {
			s.D.Set("secret_id", *v.SecretId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_database_migration.OracleConnection:
		s.D.Set("connection_type", "ORACLE")

		if v.ConnectionString != nil {
			s.D.Set("connection_string", *v.ConnectionString)
		}

		if v.DatabaseId != nil {
			s.D.Set("database_id", *v.DatabaseId)
		}

		if v.SshHost != nil {
			s.D.Set("ssh_host", *v.SshHost)
		}

		if v.SshKey != nil {
			s.D.Set("ssh_key", *v.SshKey)
		}

		if v.SshSudoLocation != nil {
			s.D.Set("ssh_sudo_location", *v.SshSudoLocation)
		}

		if v.SshUser != nil {
			s.D.Set("ssh_user", *v.SshUser)
		}

		s.D.Set("technology_type", v.TechnologyType)

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

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		//s.D.Set("nsg_ids", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds))
		s.D.Set("nsg_ids", v.NsgIds)

		if v.Password != nil {
			s.D.Set("password", *v.Password)
		}

		if v.PrivateEndpointId != nil {
			s.D.Set("private_endpoint_id", *v.PrivateEndpointId)
		}

		if v.ReplicationPassword != nil {
			s.D.Set("replication_password", *v.ReplicationPassword)
		}

		if v.ReplicationUsername != nil {
			s.D.Set("replication_username", *v.ReplicationUsername)
		}

		if v.SecretId != nil {
			s.D.Set("secret_id", *v.SecretId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	default:
		log.Printf("[WARN] Received 'connection_type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func ConnectionSummaryToMap(obj oci_database_migration.ConnectionSummary, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_database_migration.MysqlConnectionSummary:
		result["connection_type"] = "MYSQL"

		additionalAttributes := []interface{}{}
		for _, item := range v.AdditionalAttributes {
			additionalAttributes = append(additionalAttributes, NameValuePairToMap(item))
		}
		result["additional_attributes"] = additionalAttributes

		if v.DatabaseName != nil {
			result["database_name"] = string(*v.DatabaseName)
		}

		if v.DbSystemId != nil {
			result["db_system_id"] = string(*v.DbSystemId)
		}

		if v.Host != nil {
			result["host"] = string(*v.Host)
		}

		if v.Port != nil {
			result["port"] = int(*v.Port)
		}

		result["security_protocol"] = string(v.SecurityProtocol)

		result["ssl_mode"] = string(v.SslMode)

		result["technology_type"] = string(v.TechnologyType)
	case oci_database_migration.OracleConnectionSummary:
		result["connection_type"] = "ORACLE"

		if v.ConnectionString != nil {
			result["connection_string"] = string(*v.ConnectionString)
		}

		if v.DatabaseId != nil {
			result["database_id"] = string(*v.DatabaseId)
		}

		result["technology_type"] = string(v.TechnologyType)
	default:
		log.Printf("[WARN] Received 'connection_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func IngressIpDetailsToMap(obj oci_database_migration.IngressIpDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IngressIp != nil {
		result["ingress_ip"] = string(*obj.IngressIp)
	}

	return result
}

func (s *DatabaseMigrationConnectionResourceCrud) mapToNameValuePair(fieldKeyFormat string) (oci_database_migration.NameValuePair, error) {
	result := oci_database_migration.NameValuePair{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func NameValuePairToMap(obj oci_database_migration.NameValuePair) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *DatabaseMigrationConnectionResourceCrud) populateTopLevelPolymorphicCreateConnectionRequest(request *oci_database_migration.CreateConnectionRequest) error {
	//discriminator
	connectionTypeRaw, ok := s.D.GetOkExists("connection_type")
	var connectionType string
	if ok {
		connectionType = connectionTypeRaw.(string)
	} else {
		connectionType = "" // default value
	}
	switch strings.ToLower(connectionType) {
	case strings.ToLower("MYSQL"):
		details := oci_database_migration.CreateMysqlConnectionDetails{}
		if additionalAttributes, ok := s.D.GetOkExists("additional_attributes"); ok {
			interfaces := additionalAttributes.([]interface{})
			tmp := make([]oci_database_migration.NameValuePair, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "additional_attributes", stateDataIndex)
				converted, err := s.mapToNameValuePair(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("additional_attributes") {
				details.AdditionalAttributes = tmp
			}
		}
		if databaseName, ok := s.D.GetOkExists("database_name"); ok {
			tmp := databaseName.(string)
			details.DatabaseName = &tmp
		}
		if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
			tmp := dbSystemId.(string)
			details.DbSystemId = &tmp
		}
		if host, ok := s.D.GetOkExists("host"); ok {
			tmp := host.(string)
			details.Host = &tmp
		}
		if port, ok := s.D.GetOkExists("port"); ok {
			tmp := port.(int)
			details.Port = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_database_migration.MysqlConnectionSecurityProtocolEnum(securityProtocol.(string))
		}
		if sslCa, ok := s.D.GetOkExists("ssl_ca"); ok {
			tmp := sslCa.(string)
			details.SslCa = &tmp
		}
		if sslCert, ok := s.D.GetOkExists("ssl_cert"); ok {
			tmp := sslCert.(string)
			details.SslCert = &tmp
		}
		if sslCrl, ok := s.D.GetOkExists("ssl_crl"); ok {
			tmp := sslCrl.(string)
			details.SslCrl = &tmp
		}
		if sslKey, ok := s.D.GetOkExists("ssl_key"); ok {
			tmp := sslKey.(string)
			details.SslKey = &tmp
		}
		if sslMode, ok := s.D.GetOkExists("ssl_mode"); ok {
			details.SslMode = oci_database_migration.MysqlConnectionSslModeEnum(sslMode.(string))
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_database_migration.MysqlConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
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
				details.NsgIds = tmp
			}
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if replicationPassword, ok := s.D.GetOkExists("replication_password"); ok {
			tmp := replicationPassword.(string)
			details.ReplicationPassword = &tmp
		}
		if replicationUsername, ok := s.D.GetOkExists("replication_username"); ok {
			tmp := replicationUsername.(string)
			details.ReplicationUsername = &tmp
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_database_migration.MysqlConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("ORACLE"):
		details := oci_database_migration.CreateOracleConnectionDetails{}
		if connectionString, ok := s.D.GetOkExists("connection_string"); ok {
			tmp := connectionString.(string)
			details.ConnectionString = &tmp
		}
		if databaseId, ok := s.D.GetOkExists("database_id"); ok {
			tmp := databaseId.(string)
			details.DatabaseId = &tmp
		}
		if sshHost, ok := s.D.GetOkExists("ssh_host"); ok {
			tmp := sshHost.(string)
			details.SshHost = &tmp
		}
		if sshKey, ok := s.D.GetOkExists("ssh_key"); ok {
			tmp := sshKey.(string)
			details.SshKey = &tmp
		}
		if sshSudoLocation, ok := s.D.GetOkExists("ssh_sudo_location"); ok {
			tmp := sshSudoLocation.(string)
			details.SshSudoLocation = &tmp
		}
		if sshUser, ok := s.D.GetOkExists("ssh_user"); ok {
			tmp := sshUser.(string)
			details.SshUser = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_database_migration.OracleConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if wallet, ok := s.D.GetOkExists("wallet"); ok {
			tmp := wallet.(string)
			details.Wallet = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
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
				details.NsgIds = tmp
			}
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if replicationPassword, ok := s.D.GetOkExists("replication_password"); ok {
			tmp := replicationPassword.(string)
			details.ReplicationPassword = &tmp
		}
		if replicationUsername, ok := s.D.GetOkExists("replication_username"); ok {
			tmp := replicationUsername.(string)
			details.ReplicationUsername = &tmp
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_database_migration.OracleConnectionTechnologyTypeEnum(technologyType.(string)) //MysqlConnectionTechnologyTypeEnum
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	default:
		return fmt.Errorf("unknown connection_type '%v' was specified", connectionType)
	}
	return nil
}

func (s *DatabaseMigrationConnectionResourceCrud) populateTopLevelPolymorphicUpdateConnectionRequest(request *oci_database_migration.UpdateConnectionRequest) error {
	//discriminator
	connectionTypeRaw, ok := s.D.GetOkExists("connection_type")
	var connectionType string
	if ok {
		connectionType = connectionTypeRaw.(string)
	} else {
		connectionType = "" // default value
	}
	switch strings.ToLower(connectionType) {
	case strings.ToLower("MYSQL"):
		details := oci_database_migration.UpdateMysqlConnectionDetails{}
		if additionalAttributes, ok := s.D.GetOkExists("additional_attributes"); ok {
			interfaces := additionalAttributes.([]interface{})
			tmp := make([]oci_database_migration.NameValuePair, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "additional_attributes", stateDataIndex)
				converted, err := s.mapToNameValuePair(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("additional_attributes") {
				details.AdditionalAttributes = tmp
			}
		}
		if databaseName, ok := s.D.GetOkExists("database_name"); ok {
			tmp := databaseName.(string)
			details.DatabaseName = &tmp
		}
		if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
			tmp := dbSystemId.(string)
			details.DbSystemId = &tmp
		}
		if host, ok := s.D.GetOkExists("host"); ok {
			tmp := host.(string)
			details.Host = &tmp
		}
		if port, ok := s.D.GetOkExists("port"); ok {
			tmp := port.(int)
			details.Port = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_database_migration.MysqlConnectionSecurityProtocolEnum(securityProtocol.(string))
		}
		if sslCa, ok := s.D.GetOkExists("ssl_ca"); ok {
			tmp := sslCa.(string)
			details.SslCa = &tmp
		}
		if sslCert, ok := s.D.GetOkExists("ssl_cert"); ok {
			tmp := sslCert.(string)
			details.SslCert = &tmp
		}
		if sslCrl, ok := s.D.GetOkExists("ssl_crl"); ok {
			tmp := sslCrl.(string)
			details.SslCrl = &tmp
		}
		if sslKey, ok := s.D.GetOkExists("ssl_key"); ok {
			tmp := sslKey.(string)
			details.SslKey = &tmp
		}
		if sslMode, ok := s.D.GetOkExists("ssl_mode"); ok {
			details.SslMode = oci_database_migration.MysqlConnectionSslModeEnum(sslMode.(string))
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
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
				details.NsgIds = tmp
			}
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if replicationPassword, ok := s.D.GetOkExists("replication_password"); ok {
			tmp := replicationPassword.(string)
			details.ReplicationPassword = &tmp
		}
		if replicationUsername, ok := s.D.GetOkExists("replication_username"); ok {
			tmp := replicationUsername.(string)
			details.ReplicationUsername = &tmp
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("ORACLE"):
		details := oci_database_migration.UpdateOracleConnectionDetails{}
		if connectionString, ok := s.D.GetOkExists("connection_string"); ok {
			tmp := connectionString.(string)
			details.ConnectionString = &tmp
		}
		if databaseId, ok := s.D.GetOkExists("database_id"); ok {
			tmp := databaseId.(string)
			details.DatabaseId = &tmp
		}
		if sshHost, ok := s.D.GetOkExists("ssh_host"); ok {
			tmp := sshHost.(string)
			details.SshHost = &tmp
		}
		if sshKey, ok := s.D.GetOkExists("ssh_key"); ok {
			tmp := sshKey.(string)
			details.SshKey = &tmp
		}
		if sshSudoLocation, ok := s.D.GetOkExists("ssh_sudo_location"); ok {
			tmp := sshSudoLocation.(string)
			details.SshSudoLocation = &tmp
		}
		if sshUser, ok := s.D.GetOkExists("ssh_user"); ok {
			tmp := sshUser.(string)
			details.SshUser = &tmp
		}
		if wallet, ok := s.D.GetOkExists("wallet"); ok {
			tmp := wallet.(string)
			details.Wallet = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
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
				details.NsgIds = tmp
			}
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if replicationPassword, ok := s.D.GetOkExists("replication_password"); ok {
			tmp := replicationPassword.(string)
			details.ReplicationPassword = &tmp
		}
		if replicationUsername, ok := s.D.GetOkExists("replication_username"); ok {
			tmp := replicationUsername.(string)
			details.ReplicationUsername = &tmp
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	default:
		return fmt.Errorf("unknown connection_type '%v' was specified", connectionType)
	}
	return nil
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
