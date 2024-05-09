// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package golden_gate

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
	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GoldenGateConnectionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createGoldenGateConnection,
		Read:     readGoldenGateConnection,
		Update:   updateGoldenGateConnection,
		Delete:   deleteGoldenGateConnection,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"connection_type": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"AMAZON_KINESIS",
					"AMAZON_REDSHIFT",
					"AMAZON_S3",
					"AZURE_DATA_LAKE_STORAGE",
					"AZURE_SYNAPSE_ANALYTICS",
					"DB2",
					"ELASTICSEARCH",
					"GENERIC",
					"GOLDENGATE",
					"GOOGLE_BIGQUERY",
					"GOOGLE_CLOUD_STORAGE",
					"HDFS",
					"JAVA_MESSAGE_SERVICE",
					"KAFKA",
					"KAFKA_SCHEMA_REGISTRY",
					"MICROSOFT_SQLSERVER",
					"MONGODB",
					"MYSQL",
					"OCI_OBJECT_STORAGE",
					"ORACLE",
					"ORACLE_NOSQL",
					"POSTGRESQL",
					"REDIS",
					"SNOWFLAKE",
				}, true),
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"technology_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"access_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"account_key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"account_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
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
			"authentication_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authentication_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"azure_tenant_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bootstrap_servers": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"host": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"private_ip": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"client_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_secret": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connection_factory": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connection_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connection_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"consumer_properties": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"core_site_xml": {
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
			"deployment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"endpoint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fingerprint": {
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
			"jndi_connection_factory": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"jndi_initial_context_factory": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"jndi_provider_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"jndi_security_credentials": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"jndi_security_principal": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"key_store": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"key_store_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},
			"locks": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"message": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
						"related_resource_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"is_lock_override": {
				Type:     schema.TypeBool,
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
			"password": {
				Type:      schema.TypeString,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"private_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"private_key_file": {
				Type:      schema.TypeString,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},
			"private_key_passphrase": {
				Type:      schema.TypeString,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},
			"producer_properties": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"public_key_fingerprint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redis_cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"routing_method": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sas_token": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secret_access_key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_protocol": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servers": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_account_key_file": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"session_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"should_use_jndi": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"should_validate_server_certificate": {
				Type:     schema.TypeBool,
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
			"ssl_client_keystash": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_client_keystoredb": {
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
			"ssl_key_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},
			"ssl_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_server_certificate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stream_pool_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenancy_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trust_store": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trust_store_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},
			"url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vault_id": {
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

func createGoldenGateConnection(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.CreateResource(d, sync)
}

func readGoldenGateConnection(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.ReadResource(sync)
}

func updateGoldenGateConnection(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteGoldenGateConnection(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type GoldenGateConnectionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_golden_gate.GoldenGateClient
	Res                    *oci_golden_gate.Connection
	DisableNotFoundRetries bool
}

func (s *GoldenGateConnectionResourceCrud) ID() string {
	connection := *s.Res
	return *connection.GetId()
}

func (s *GoldenGateConnectionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_golden_gate.ConnectionLifecycleStateCreating),
	}
}

func (s *GoldenGateConnectionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_golden_gate.ConnectionLifecycleStateActive),
	}
}

func (s *GoldenGateConnectionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_golden_gate.ConnectionLifecycleStateDeleting),
	}
}

func (s *GoldenGateConnectionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_golden_gate.ConnectionLifecycleStateDeleted),
	}
}

func (s *GoldenGateConnectionResourceCrud) Create() error {
	request := oci_golden_gate.CreateConnectionRequest{}
	err := s.populateTopLevelPolymorphicCreateConnectionRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

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
	return s.getConnectionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate"), oci_golden_gate.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *GoldenGateConnectionResourceCrud) getConnectionFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_golden_gate.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	connectionId, err := connectionWaitForWorkRequest(workId, "goldengateconnection",
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
		if tfresource.ShouldRetry(response, false, "golden_gate", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_golden_gate.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func connectionWaitForWorkRequest(wId *string, entityType string, action oci_golden_gate.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_golden_gate.GoldenGateClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "golden_gate")
	retryPolicy.ShouldRetryOperation = connectionWorkRequestShouldRetryFunc(timeout)

	response := oci_golden_gate.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_golden_gate.OperationStatusInProgress),
			string(oci_golden_gate.OperationStatusAccepted),
			string(oci_golden_gate.OperationStatusCanceled),
		},
		Target: []string{
			string(oci_golden_gate.OperationStatusSucceeded),
			string(oci_golden_gate.OperationStatusFailed),
			string(oci_golden_gate.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_golden_gate.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_golden_gate.OperationStatusFailed || response.Status == oci_golden_gate.OperationStatusCanceled {
		return nil, getErrorFromGoldenGateConnectionWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromGoldenGateConnectionWorkRequest(client *oci_golden_gate.GoldenGateClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_golden_gate.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_golden_gate.ListWorkRequestErrorsRequest{
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

func (s *GoldenGateConnectionResourceCrud) Get() error {
	request := oci_golden_gate.GetConnectionRequest{}

	tmp := s.D.Id()
	request.ConnectionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.GetConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Connection
	return nil
}

func (s *GoldenGateConnectionResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_golden_gate.UpdateConnectionRequest{}
	err := s.populateTopLevelPolymorphicUpdateConnectionRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.UpdateConnection(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getConnectionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate"), oci_golden_gate.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *GoldenGateConnectionResourceCrud) Delete() error {
	request := oci_golden_gate.DeleteConnectionRequest{}

	tmp := s.D.Id()
	request.ConnectionId = &tmp

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		request.IsLockOverride = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.DeleteConnection(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := connectionWaitForWorkRequest(workId, "goldengateconnection",
		oci_golden_gate.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *GoldenGateConnectionResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_golden_gate.AmazonKinesisConnection:
		s.D.Set("connection_type", "AMAZON_KINESIS")

		if v.AccessKeyId != nil {
			s.D.Set("access_key_id", *v.AccessKeyId)
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

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.AmazonRedshiftConnection:
		s.D.Set("connection_type", "AMAZON_REDSHIFT")

		if v.ConnectionUrl != nil {
			s.D.Set("connection_url", *v.ConnectionUrl)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

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

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.AmazonS3Connection:
		s.D.Set("connection_type", "AMAZON_S3")

		if v.AccessKeyId != nil {
			s.D.Set("access_key_id", *v.AccessKeyId)
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

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.AzureDataLakeStorageConnection:
		s.D.Set("connection_type", "AZURE_DATA_LAKE_STORAGE")

		if v.AccountName != nil {
			s.D.Set("account_name", *v.AccountName)
		}

		s.D.Set("authentication_type", v.AuthenticationType)

		if v.AzureTenantId != nil {
			s.D.Set("azure_tenant_id", *v.AzureTenantId)
		}

		if v.ClientId != nil {
			s.D.Set("client_id", *v.ClientId)
		}

		if v.Endpoint != nil {
			s.D.Set("endpoint", *v.Endpoint)
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

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.AzureSynapseConnection:
		s.D.Set("connection_type", "AZURE_SYNAPSE_ANALYTICS")

		if v.ConnectionString != nil {
			s.D.Set("connection_string", *v.ConnectionString)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

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

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.Db2Connection:
		s.D.Set("connection_type", "DB2")

		additionalAttributes := []interface{}{}
		for _, item := range v.AdditionalAttributes {
			additionalAttributes = append(additionalAttributes, NameValuePairToMap(item))
		}
		s.D.Set("additional_attributes", additionalAttributes)

		if v.DatabaseName != nil {
			s.D.Set("database_name", *v.DatabaseName)
		}

		if v.Host != nil {
			s.D.Set("host", *v.Host)
		}

		if v.Port != nil {
			s.D.Set("port", *v.Port)
		}

		s.D.Set("security_protocol", v.SecurityProtocol)

		s.D.Set("technology_type", v.TechnologyType)

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

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
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.ElasticsearchConnection:
		s.D.Set("connection_type", "ELASTICSEARCH")

		s.D.Set("authentication_type", v.AuthenticationType)

		s.D.Set("security_protocol", v.SecurityProtocol)

		if v.Servers != nil {
			s.D.Set("servers", *v.Servers)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

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

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.GenericConnection:
		s.D.Set("connection_type", "GENERIC")

		if v.Host != nil {
			s.D.Set("host", *v.Host)
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

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.GoldenGateConnection:
		s.D.Set("connection_type", "GOLDENGATE")

		if v.DeploymentId != nil {
			s.D.Set("deployment_id", *v.DeploymentId)
		}

		if v.Host != nil {
			s.D.Set("host", *v.Host)
		}

		if v.Port != nil {
			s.D.Set("port", *v.Port)
		}

		if v.PrivateIp != nil {
			s.D.Set("private_ip", *v.PrivateIp)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

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

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.GoogleBigQueryConnection:
		s.D.Set("connection_type", "GOOGLE_BIGQUERY")

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

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.GoogleCloudStorageConnection:
		s.D.Set("connection_type", "GOOGLE_CLOUD_STORAGE")

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

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.HdfsConnection:
		s.D.Set("connection_type", "HDFS")

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

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.JavaMessageServiceConnection:
		s.D.Set("connection_type", "JAVA_MESSAGE_SERVICE")

		s.D.Set("authentication_type", v.AuthenticationType)

		if v.ConnectionFactory != nil {
			s.D.Set("connection_factory", *v.ConnectionFactory)
		}

		if v.ConnectionUrl != nil {
			s.D.Set("connection_url", *v.ConnectionUrl)
		}

		if v.JndiConnectionFactory != nil {
			s.D.Set("jndi_connection_factory", *v.JndiConnectionFactory)
		}

		if v.JndiInitialContextFactory != nil {
			s.D.Set("jndi_initial_context_factory", *v.JndiInitialContextFactory)
		}

		if v.JndiProviderUrl != nil {
			s.D.Set("jndi_provider_url", *v.JndiProviderUrl)
		}

		if v.JndiSecurityPrincipal != nil {
			s.D.Set("jndi_security_principal", *v.JndiSecurityPrincipal)
		}

		if v.PrivateIp != nil {
			s.D.Set("private_ip", *v.PrivateIp)
		}

		s.D.Set("security_protocol", v.SecurityProtocol)

		if v.ShouldUseJndi != nil {
			s.D.Set("should_use_jndi", *v.ShouldUseJndi)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

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

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.KafkaConnection:
		s.D.Set("connection_type", "KAFKA")

		bootstrapServers := []interface{}{}
		for _, item := range v.BootstrapServers {
			bootstrapServers = append(bootstrapServers, KafkaBootstrapServerToMap(item))
		}
		s.D.Set("bootstrap_servers", bootstrapServers)

		s.D.Set("security_protocol", v.SecurityProtocol)

		if v.StreamPoolId != nil {
			s.D.Set("stream_pool_id", *v.StreamPoolId)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

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

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.KafkaSchemaRegistryConnection:
		s.D.Set("connection_type", "KAFKA_SCHEMA_REGISTRY")

		s.D.Set("authentication_type", v.AuthenticationType)

		if v.PrivateIp != nil {
			s.D.Set("private_ip", *v.PrivateIp)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.Url != nil {
			s.D.Set("url", *v.Url)
		}

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

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

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.MicrosoftSqlserverConnection:
		s.D.Set("connection_type", "MICROSOFT_SQLSERVER")

		additionalAttributes := []interface{}{}
		for _, item := range v.AdditionalAttributes {
			additionalAttributes = append(additionalAttributes, NameValuePairToMap(item))
		}
		s.D.Set("additional_attributes", additionalAttributes)

		if v.DatabaseName != nil {
			s.D.Set("database_name", *v.DatabaseName)
		}

		if v.Host != nil {
			s.D.Set("host", *v.Host)
		}

		if v.Port != nil {
			s.D.Set("port", *v.Port)
		}

		if v.PrivateIp != nil {
			s.D.Set("private_ip", *v.PrivateIp)
		}

		s.D.Set("security_protocol", v.SecurityProtocol)

		if v.ShouldValidateServerCertificate != nil {
			s.D.Set("should_validate_server_certificate", *v.ShouldValidateServerCertificate)
		}

		if v.SslCa != nil {
			s.D.Set("ssl_ca", *v.SslCa)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

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

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.MongoDbConnection:
		s.D.Set("connection_type", "MONGODB")

		if v.ConnectionString != nil {
			s.D.Set("connection_string", *v.ConnectionString)
		}

		if v.DatabaseId != nil {
			s.D.Set("database_id", *v.DatabaseId)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

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

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.MysqlConnection:
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

		if v.PrivateIp != nil {
			s.D.Set("private_ip", *v.PrivateIp)
		}

		s.D.Set("security_protocol", v.SecurityProtocol)

		s.D.Set("ssl_mode", v.SslMode)

		s.D.Set("technology_type", v.TechnologyType)

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

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

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.OciObjectStorageConnection:
		s.D.Set("connection_type", "OCI_OBJECT_STORAGE")

		if v.Region != nil {
			s.D.Set("region", *v.Region)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.TenancyId != nil {
			s.D.Set("tenancy_id", *v.TenancyId)
		}

		if v.UserId != nil {
			s.D.Set("user_id", *v.UserId)
		}

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

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.OracleConnection:
		s.D.Set("connection_type", "ORACLE")

		s.D.Set("authentication_mode", v.AuthenticationMode)

		if v.ConnectionString != nil {
			s.D.Set("connection_string", *v.ConnectionString)
		}

		if v.DatabaseId != nil {
			s.D.Set("database_id", *v.DatabaseId)
		}

		if v.PrivateIp != nil {
			s.D.Set("private_ip", *v.PrivateIp)
		}

		s.D.Set("session_mode", v.SessionMode)

		s.D.Set("technology_type", v.TechnologyType)

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

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

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.OracleNosqlConnection:
		s.D.Set("connection_type", "ORACLE_NOSQL")

		if v.Region != nil {
			s.D.Set("region", *v.Region)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.TenancyId != nil {
			s.D.Set("tenancy_id", *v.TenancyId)
		}

		if v.UserId != nil {
			s.D.Set("user_id", *v.UserId)
		}

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

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.PostgresqlConnection:
		s.D.Set("connection_type", "POSTGRESQL")

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

		if v.PrivateIp != nil {
			s.D.Set("private_ip", *v.PrivateIp)
		}

		s.D.Set("security_protocol", v.SecurityProtocol)

		s.D.Set("ssl_mode", v.SslMode)

		s.D.Set("technology_type", v.TechnologyType)

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

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

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.RedisConnection:
		s.D.Set("connection_type", "REDIS")

		s.D.Set("authentication_type", v.AuthenticationType)

		if v.RedisClusterId != nil {
			s.D.Set("redis_cluster_id", *v.RedisClusterId)
		}

		s.D.Set("security_protocol", v.SecurityProtocol)

		if v.Servers != nil {
			s.D.Set("servers", *v.Servers)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

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

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.SnowflakeConnection:
		s.D.Set("connection_type", "SNOWFLAKE")

		s.D.Set("authentication_type", v.AuthenticationType)

		if v.ConnectionUrl != nil {
			s.D.Set("connection_url", *v.ConnectionUrl)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

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

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
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

func (s *GoldenGateConnectionResourceCrud) mapToAddResourceLockDetails(fieldKeyFormat string) (oci_golden_gate.AddResourceLockDetails, error) {
	result := oci_golden_gate.AddResourceLockDetails{}

	if message, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "message")); ok {
		tmp := message.(string)
		result.Message = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_golden_gate.AddResourceLockDetailsTypeEnum(type_.(string))
	}

	return result, nil
}

func ConnectionSummaryToMap(obj oci_golden_gate.ConnectionSummary, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	//set common fields
	result["id"] = *obj.GetId()

	if obj.GetDisplayName() != nil {
		result["display_name"] = *obj.GetDisplayName()
	}
	result["compartment_id"] = *obj.GetCompartmentId()
	result["state"] = string(obj.GetLifecycleState())

	if obj.GetTimeCreated() != nil {
		result["time_created"] = obj.GetTimeCreated().String()
	}
	if obj.GetTimeUpdated() != nil {
		result["time_updated"] = obj.GetTimeUpdated().String()
	}
	if obj.GetDescription() != nil {
		result["description"] = *obj.GetDescription()
	}
	result["freeform_tags"] = obj.GetFreeformTags()
	if obj.GetDefinedTags() != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.GetDefinedTags())
	}
	if obj.GetSystemTags() != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.GetSystemTags())
	}
	if obj.GetLifecycleDetails() != nil {
		result["lifecycle_details"] = *obj.GetLifecycleDetails()
	}

	if obj.GetVaultId() != nil {
		result["vault_id"] = *obj.GetVaultId()
	}
	if obj.GetSubnetId() != nil {
		result["subnet_id"] = *obj.GetSubnetId()
	}

	ingressIps := []interface{}{}
	for _, item := range obj.GetIngressIps() {
		ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
	}
	result["ingress_ips"] = ingressIps

	if obj.GetNsgIds() != nil {
		result["nsg_ids"] = obj.GetNsgIds()
	}
	result["routing_method"] = string(obj.GetRoutingMethod())

	// set type specific fields
	switch v := (obj).(type) {
	case oci_golden_gate.AmazonKinesisConnectionSummary:
		result["connection_type"] = "AMAZON_KINESIS"

		if v.AccessKeyId != nil {
			result["access_key_id"] = string(*v.AccessKeyId)
		}

		result["technology_type"] = string(v.TechnologyType)
	case oci_golden_gate.AmazonRedshiftConnectionSummary:
		result["connection_type"] = "AMAZON_REDSHIFT"

		if v.ConnectionUrl != nil {
			result["connection_url"] = string(*v.ConnectionUrl)
		}

		result["technology_type"] = string(v.TechnologyType)

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	case oci_golden_gate.AmazonS3ConnectionSummary:
		result["connection_type"] = "AMAZON_S3"

		if v.AccessKeyId != nil {
			result["access_key_id"] = string(*v.AccessKeyId)
		}

		result["technology_type"] = string(v.TechnologyType)
	case oci_golden_gate.AzureDataLakeStorageConnectionSummary:
		result["connection_type"] = "AZURE_DATA_LAKE_STORAGE"

		if v.AccountName != nil {
			result["account_name"] = string(*v.AccountName)
		}

		result["authentication_type"] = string(v.AuthenticationType)

		if v.AzureTenantId != nil {
			result["azure_tenant_id"] = string(*v.AzureTenantId)
		}

		if v.ClientId != nil {
			result["client_id"] = string(*v.ClientId)
		}

		if v.Endpoint != nil {
			result["endpoint"] = string(*v.Endpoint)
		}

		result["technology_type"] = string(v.TechnologyType)
	case oci_golden_gate.AzureSynapseConnectionSummary:
		result["connection_type"] = "AZURE_SYNAPSE_ANALYTICS"

		if v.ConnectionString != nil {
			result["connection_string"] = string(*v.ConnectionString)
		}

		result["technology_type"] = string(v.TechnologyType)

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	case oci_golden_gate.Db2ConnectionSummary:
		result["connection_type"] = "DB2"

		additionalAttributes := []interface{}{}
		for _, item := range v.AdditionalAttributes {
			additionalAttributes = append(additionalAttributes, NameValuePairToMap(item))
		}
		result["additional_attributes"] = additionalAttributes

		if v.DatabaseName != nil {
			result["database_name"] = string(*v.DatabaseName)
		}

		if v.Host != nil {
			result["host"] = string(*v.Host)
		}

		if v.Port != nil {
			result["port"] = int(*v.Port)
		}

		result["security_protocol"] = string(v.SecurityProtocol)

		result["technology_type"] = string(v.TechnologyType)

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	case oci_golden_gate.ElasticsearchConnectionSummary:
		result["connection_type"] = "ELASTICSEARCH"

		result["authentication_type"] = string(v.AuthenticationType)

		result["security_protocol"] = string(v.SecurityProtocol)

		if v.Servers != nil {
			result["servers"] = string(*v.Servers)
		}

		result["technology_type"] = string(v.TechnologyType)

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	case oci_golden_gate.GenericConnectionSummary:
		result["connection_type"] = "GENERIC"

		if v.Host != nil {
			result["host"] = string(*v.Host)
		}

		result["technology_type"] = string(v.TechnologyType)
	case oci_golden_gate.GoldenGateConnectionSummary:
		result["connection_type"] = "GOLDENGATE"

		if v.DeploymentId != nil {
			result["deployment_id"] = string(*v.DeploymentId)
		}

		if v.Host != nil {
			result["host"] = string(*v.Host)
		}

		if v.Port != nil {
			result["port"] = int(*v.Port)
		}

		if v.PrivateIp != nil {
			result["private_ip"] = string(*v.PrivateIp)
		}

		result["technology_type"] = string(v.TechnologyType)

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	case oci_golden_gate.GoogleBigQueryConnectionSummary:
		result["connection_type"] = "GOOGLE_BIGQUERY"

		result["technology_type"] = string(v.TechnologyType)
	case oci_golden_gate.GoogleCloudStorageConnectionSummary:
		result["connection_type"] = "GOOGLE_CLOUD_STORAGE"

		result["technology_type"] = string(v.TechnologyType)
	case oci_golden_gate.HdfsConnectionSummary:
		result["connection_type"] = "HDFS"

		result["technology_type"] = string(v.TechnologyType)
	case oci_golden_gate.JavaMessageServiceConnectionSummary:
		result["connection_type"] = "JAVA_MESSAGE_SERVICE"

		result["authentication_type"] = string(v.AuthenticationType)

		if v.ConnectionFactory != nil {
			result["connection_factory"] = string(*v.ConnectionFactory)
		}

		if v.ConnectionUrl != nil {
			result["connection_url"] = string(*v.ConnectionUrl)
		}

		if v.JndiConnectionFactory != nil {
			result["jndi_connection_factory"] = string(*v.JndiConnectionFactory)
		}

		if v.JndiInitialContextFactory != nil {
			result["jndi_initial_context_factory"] = string(*v.JndiInitialContextFactory)
		}

		if v.JndiProviderUrl != nil {
			result["jndi_provider_url"] = string(*v.JndiProviderUrl)
		}

		if v.JndiSecurityPrincipal != nil {
			result["jndi_security_principal"] = string(*v.JndiSecurityPrincipal)
		}

		if v.PrivateIp != nil {
			result["private_ip"] = string(*v.PrivateIp)
		}

		result["security_protocol"] = string(v.SecurityProtocol)

		if v.ShouldUseJndi != nil {
			result["should_use_jndi"] = bool(*v.ShouldUseJndi)
		}

		result["technology_type"] = string(v.TechnologyType)

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	case oci_golden_gate.KafkaConnectionSummary:
		result["connection_type"] = "KAFKA"

		bootstrapServers := []interface{}{}
		for _, item := range v.BootstrapServers {
			bootstrapServers = append(bootstrapServers, KafkaBootstrapServerToMap(item))
		}
		result["bootstrap_servers"] = bootstrapServers

		result["security_protocol"] = string(v.SecurityProtocol)

		if v.StreamPoolId != nil {
			result["stream_pool_id"] = string(*v.StreamPoolId)
		}

		result["technology_type"] = string(v.TechnologyType)

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	case oci_golden_gate.KafkaSchemaRegistryConnectionSummary:
		result["connection_type"] = "KAFKA_SCHEMA_REGISTRY"

		result["authentication_type"] = string(v.AuthenticationType)

		if v.PrivateIp != nil {
			result["private_ip"] = string(*v.PrivateIp)
		}

		result["technology_type"] = string(v.TechnologyType)

		if v.Url != nil {
			result["url"] = string(*v.Url)
		}

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	case oci_golden_gate.MicrosoftSqlserverConnectionSummary:
		result["connection_type"] = "MICROSOFT_SQLSERVER"

		additionalAttributes := []interface{}{}
		for _, item := range v.AdditionalAttributes {
			additionalAttributes = append(additionalAttributes, NameValuePairToMap(item))
		}
		result["additional_attributes"] = additionalAttributes

		if v.DatabaseName != nil {
			result["database_name"] = string(*v.DatabaseName)
		}

		if v.Host != nil {
			result["host"] = string(*v.Host)
		}

		if v.Port != nil {
			result["port"] = int(*v.Port)
		}

		if v.PrivateIp != nil {
			result["private_ip"] = string(*v.PrivateIp)
		}

		result["security_protocol"] = string(v.SecurityProtocol)

		if v.ShouldValidateServerCertificate != nil {
			result["should_validate_server_certificate"] = bool(*v.ShouldValidateServerCertificate)
		}

		if v.SslCa != nil {
			result["ssl_ca"] = string(*v.SslCa)
		}

		result["technology_type"] = string(v.TechnologyType)

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	case oci_golden_gate.MongoDbConnectionSummary:
		result["connection_type"] = "MONGODB"

		if v.ConnectionString != nil {
			result["connection_string"] = string(*v.ConnectionString)
		}

		if v.DatabaseId != nil {
			result["database_id"] = string(*v.DatabaseId)
		}

		result["technology_type"] = string(v.TechnologyType)

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	case oci_golden_gate.MysqlConnectionSummary:
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

		if v.PrivateIp != nil {
			result["private_ip"] = string(*v.PrivateIp)
		}

		result["security_protocol"] = string(v.SecurityProtocol)

		result["ssl_mode"] = string(v.SslMode)

		result["technology_type"] = string(v.TechnologyType)

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	case oci_golden_gate.OciObjectStorageConnectionSummary:
		result["connection_type"] = "OCI_OBJECT_STORAGE"

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}

		result["technology_type"] = string(v.TechnologyType)

		if v.TenancyId != nil {
			result["tenancy_id"] = string(*v.TenancyId)
		}

		if v.UserId != nil {
			result["user_id"] = string(*v.UserId)
		}
	case oci_golden_gate.OracleConnectionSummary:
		result["connection_type"] = "ORACLE"

		result["authentication_mode"] = string(v.AuthenticationMode)

		if v.ConnectionString != nil {
			result["connection_string"] = string(*v.ConnectionString)
		}

		if v.DatabaseId != nil {
			result["database_id"] = string(*v.DatabaseId)
		}

		if v.PrivateIp != nil {
			result["private_ip"] = string(*v.PrivateIp)
		}

		result["session_mode"] = string(v.SessionMode)

		result["technology_type"] = string(v.TechnologyType)

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	case oci_golden_gate.OracleNosqlConnectionSummary:
		result["connection_type"] = "ORACLE_NOSQL"

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}

		result["technology_type"] = string(v.TechnologyType)

		if v.TenancyId != nil {
			result["tenancy_id"] = string(*v.TenancyId)
		}

		if v.UserId != nil {
			result["user_id"] = string(*v.UserId)
		}
	case oci_golden_gate.PostgresqlConnectionSummary:
		result["connection_type"] = "POSTGRESQL"

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

		if v.PrivateIp != nil {
			result["private_ip"] = string(*v.PrivateIp)
		}

		result["security_protocol"] = string(v.SecurityProtocol)

		result["ssl_mode"] = string(v.SslMode)

		result["technology_type"] = string(v.TechnologyType)

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	case oci_golden_gate.RedisConnectionSummary:
		result["connection_type"] = "REDIS"

		result["authentication_type"] = string(v.AuthenticationType)

		if v.RedisClusterId != nil {
			result["redis_cluster_id"] = string(*v.RedisClusterId)
		}

		result["security_protocol"] = string(v.SecurityProtocol)

		if v.Servers != nil {
			result["servers"] = string(*v.Servers)
		}

		result["technology_type"] = string(v.TechnologyType)

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	case oci_golden_gate.SnowflakeConnectionSummary:
		result["connection_type"] = "SNOWFLAKE"

		result["authentication_type"] = string(v.AuthenticationType)

		if v.ConnectionUrl != nil {
			result["connection_url"] = string(*v.ConnectionUrl)
		}

		result["technology_type"] = string(v.TechnologyType)

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	default:
		log.Printf("[WARN] Received 'connection_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func IngressIpDetailsToMap(obj oci_golden_gate.IngressIpDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IngressIp != nil {
		result["ingress_ip"] = string(*obj.IngressIp)
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) mapToKafkaBootstrapServer(fieldKeyFormat string) (oci_golden_gate.KafkaBootstrapServer, error) {
	result := oci_golden_gate.KafkaBootstrapServer{}

	if host, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "host")); ok {
		tmp := host.(string)
		result.Host = &tmp
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(int)
		result.Port = &tmp
	}

	if privateIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_ip")); ok {
		tmp := privateIp.(string)
		result.PrivateIp = &tmp
	}

	return result, nil
}

func KafkaBootstrapServerToMap(obj oci_golden_gate.KafkaBootstrapServer) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Host != nil {
		result["host"] = string(*obj.Host)
	}

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	if obj.PrivateIp != nil {
		result["private_ip"] = string(*obj.PrivateIp)
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) mapToNameValuePair(fieldKeyFormat string) (oci_golden_gate.NameValuePair, error) {
	result := oci_golden_gate.NameValuePair{}

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

func NameValuePairToMap(obj oci_golden_gate.NameValuePair) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) populateTopLevelPolymorphicCreateConnectionRequest(request *oci_golden_gate.CreateConnectionRequest) error {
	//discriminator
	connectionTypeRaw, ok := s.D.GetOkExists("connection_type")
	var connectionType string
	if ok {
		connectionType = connectionTypeRaw.(string)
	} else {
		connectionType = "" // default value
	}
	switch strings.ToLower(connectionType) {
	case strings.ToLower("AMAZON_KINESIS"):
		details := oci_golden_gate.CreateAmazonKinesisConnectionDetails{}
		if accessKeyId, ok := s.D.GetOkExists("access_key_id"); ok {
			tmp := accessKeyId.(string)
			details.AccessKeyId = &tmp
		}
		if secretAccessKey, ok := s.D.GetOkExists("secret_access_key"); ok {
			tmp := secretAccessKey.(string)
			details.SecretAccessKey = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.AmazonKinesisConnectionTechnologyTypeEnum(technologyType.(string))
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
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("AMAZON_REDSHIFT"):
		details := oci_golden_gate.CreateAmazonRedshiftConnectionDetails{}
		if connectionUrl, ok := s.D.GetOkExists("connection_url"); ok {
			tmp := connectionUrl.(string)
			details.ConnectionUrl = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.AmazonRedshiftConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
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
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("AMAZON_S3"):
		details := oci_golden_gate.CreateAmazonS3ConnectionDetails{}
		if accessKeyId, ok := s.D.GetOkExists("access_key_id"); ok {
			tmp := accessKeyId.(string)
			details.AccessKeyId = &tmp
		}
		if secretAccessKey, ok := s.D.GetOkExists("secret_access_key"); ok {
			tmp := secretAccessKey.(string)
			details.SecretAccessKey = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.AmazonS3ConnectionTechnologyTypeEnum(technologyType.(string))
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
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("AZURE_DATA_LAKE_STORAGE"):
		details := oci_golden_gate.CreateAzureDataLakeStorageConnectionDetails{}
		if accountKey, ok := s.D.GetOkExists("account_key"); ok {
			tmp := accountKey.(string)
			details.AccountKey = &tmp
		}
		if accountName, ok := s.D.GetOkExists("account_name"); ok {
			tmp := accountName.(string)
			details.AccountName = &tmp
		}
		if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
			details.AuthenticationType = oci_golden_gate.AzureDataLakeStorageConnectionAuthenticationTypeEnum(authenticationType.(string))
		}
		if azureTenantId, ok := s.D.GetOkExists("azure_tenant_id"); ok {
			tmp := azureTenantId.(string)
			details.AzureTenantId = &tmp
		}
		if clientId, ok := s.D.GetOkExists("client_id"); ok {
			tmp := clientId.(string)
			details.ClientId = &tmp
		}
		if clientSecret, ok := s.D.GetOkExists("client_secret"); ok {
			tmp := clientSecret.(string)
			details.ClientSecret = &tmp
		}
		if endpoint, ok := s.D.GetOkExists("endpoint"); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		if sasToken, ok := s.D.GetOkExists("sas_token"); ok {
			tmp := sasToken.(string)
			details.SasToken = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.AzureDataLakeStorageConnectionTechnologyTypeEnum(technologyType.(string))
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
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("AZURE_SYNAPSE_ANALYTICS"):
		details := oci_golden_gate.CreateAzureSynapseConnectionDetails{}
		if connectionString, ok := s.D.GetOkExists("connection_string"); ok {
			tmp := connectionString.(string)
			details.ConnectionString = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.AzureSynapseConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
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
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("DB2"):
		details := oci_golden_gate.CreateDb2ConnectionDetails{}
		if additionalAttributes, ok := s.D.GetOkExists("additional_attributes"); ok {
			interfaces := additionalAttributes.([]interface{})
			tmp := make([]oci_golden_gate.NameValuePair, len(interfaces))
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
		if host, ok := s.D.GetOkExists("host"); ok {
			tmp := host.(string)
			details.Host = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if port, ok := s.D.GetOkExists("port"); ok {
			tmp := port.(int)
			details.Port = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.Db2ConnectionSecurityProtocolEnum(securityProtocol.(string))
		}
		if sslClientKeystash, ok := s.D.GetOkExists("ssl_client_keystash"); ok {
			tmp := sslClientKeystash.(string)
			details.SslClientKeystash = &tmp
		}
		if sslClientKeystoredb, ok := s.D.GetOkExists("ssl_client_keystoredb"); ok {
			tmp := sslClientKeystoredb.(string)
			details.SslClientKeystoredb = &tmp
		}
		if sslServerCertificate, ok := s.D.GetOkExists("ssl_server_certificate"); ok {
			tmp := sslServerCertificate.(string)
			details.SslServerCertificate = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.Db2ConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("ELASTICSEARCH"):
		details := oci_golden_gate.CreateElasticsearchConnectionDetails{}
		if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
			details.AuthenticationType = oci_golden_gate.ElasticsearchConnectionAuthenticationTypeEnum(authenticationType.(string))

		}
		if fingerprint, ok := s.D.GetOkExists("fingerprint"); ok {
			tmp := fingerprint.(string)
			details.Fingerprint = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.ElasticsearchConnectionSecurityProtocolEnum(securityProtocol.(string))

		}
		if servers, ok := s.D.GetOkExists("servers"); ok {
			tmp := servers.(string)
			details.Servers = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.ElasticsearchConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
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
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("GENERIC"):
		details := oci_golden_gate.CreateGenericConnectionDetails{}
		if host, ok := s.D.GetOkExists("host"); ok {
			tmp := host.(string)
			details.Host = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.GenericConnectionTechnologyTypeEnum(technologyType.(string))
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
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("GOLDENGATE"):
		details := oci_golden_gate.CreateGoldenGateConnectionDetails{}
		if deploymentId, ok := s.D.GetOkExists("deployment_id"); ok {
			tmp := deploymentId.(string)
			details.DeploymentId = &tmp
		}
		if host, ok := s.D.GetOkExists("host"); ok {
			tmp := host.(string)
			details.Host = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if port, ok := s.D.GetOkExists("port"); ok {
			tmp := port.(int)
			details.Port = &tmp
		}
		if privateIp, ok := s.D.GetOkExists("private_ip"); ok {
			tmp := privateIp.(string)
			details.PrivateIp = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.GoldenGateConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
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
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("GOOGLE_BIGQUERY"):
		details := oci_golden_gate.CreateGoogleBigQueryConnectionDetails{}
		if serviceAccountKeyFile, ok := s.D.GetOkExists("service_account_key_file"); ok {
			tmp := serviceAccountKeyFile.(string)
			details.ServiceAccountKeyFile = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.GoogleBigQueryConnectionTechnologyTypeEnum(technologyType.(string))
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
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("GOOGLE_CLOUD_STORAGE"):
		details := oci_golden_gate.CreateGoogleCloudStorageConnectionDetails{}
		if serviceAccountKeyFile, ok := s.D.GetOkExists("service_account_key_file"); ok {
			tmp := serviceAccountKeyFile.(string)
			details.ServiceAccountKeyFile = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.GoogleCloudStorageConnectionTechnologyTypeEnum(technologyType.(string))
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
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("HDFS"):
		details := oci_golden_gate.CreateHdfsConnectionDetails{}
		if coreSiteXml, ok := s.D.GetOkExists("core_site_xml"); ok {
			tmp := coreSiteXml.(string)
			details.CoreSiteXml = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.HdfsConnectionTechnologyTypeEnum(technologyType.(string))
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
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("JAVA_MESSAGE_SERVICE"):
		details := oci_golden_gate.CreateJavaMessageServiceConnectionDetails{}
		if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
			details.AuthenticationType = oci_golden_gate.JavaMessageServiceConnectionAuthenticationTypeEnum(authenticationType.(string))
		}
		if connectionFactory, ok := s.D.GetOkExists("connection_factory"); ok {
			tmp := connectionFactory.(string)
			details.ConnectionFactory = &tmp
		}
		if connectionUrl, ok := s.D.GetOkExists("connection_url"); ok {
			tmp := connectionUrl.(string)
			details.ConnectionUrl = &tmp
		}
		if jndiConnectionFactory, ok := s.D.GetOkExists("jndi_connection_factory"); ok {
			tmp := jndiConnectionFactory.(string)
			details.JndiConnectionFactory = &tmp
		}
		if jndiInitialContextFactory, ok := s.D.GetOkExists("jndi_initial_context_factory"); ok {
			tmp := jndiInitialContextFactory.(string)
			details.JndiInitialContextFactory = &tmp
		}
		if jndiProviderUrl, ok := s.D.GetOkExists("jndi_provider_url"); ok {
			tmp := jndiProviderUrl.(string)
			details.JndiProviderUrl = &tmp
		}
		if jndiSecurityCredentials, ok := s.D.GetOkExists("jndi_security_credentials"); ok {
			tmp := jndiSecurityCredentials.(string)
			details.JndiSecurityCredentials = &tmp
		}
		if jndiSecurityPrincipal, ok := s.D.GetOkExists("jndi_security_principal"); ok {
			tmp := jndiSecurityPrincipal.(string)
			details.JndiSecurityPrincipal = &tmp
		}
		if keyStore, ok := s.D.GetOkExists("key_store"); ok {
			tmp := keyStore.(string)
			details.KeyStore = &tmp
		}
		if keyStorePassword, ok := s.D.GetOkExists("key_store_password"); ok {
			tmp := keyStorePassword.(string)
			details.KeyStorePassword = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if privateIp, ok := s.D.GetOkExists("private_ip"); ok {
			tmp := privateIp.(string)
			details.PrivateIp = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.JavaMessageServiceConnectionSecurityProtocolEnum(securityProtocol.(string))

		}
		if shouldUseJndi, ok := s.D.GetOkExists("should_use_jndi"); ok {
			tmp := shouldUseJndi.(bool)
			details.ShouldUseJndi = &tmp
		}
		if sslKeyPassword, ok := s.D.GetOkExists("ssl_key_password"); ok {
			tmp := sslKeyPassword.(string)
			details.SslKeyPassword = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.JavaMessageServiceConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if trustStore, ok := s.D.GetOkExists("trust_store"); ok {
			tmp := trustStore.(string)
			details.TrustStore = &tmp
		}
		if trustStorePassword, ok := s.D.GetOkExists("trust_store_password"); ok {
			tmp := trustStorePassword.(string)
			details.TrustStorePassword = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
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
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("KAFKA"):
		details := oci_golden_gate.CreateKafkaConnectionDetails{}
		if bootstrapServers, ok := s.D.GetOkExists("bootstrap_servers"); ok {
			interfaces := bootstrapServers.([]interface{})
			tmp := make([]oci_golden_gate.KafkaBootstrapServer, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "bootstrap_servers", stateDataIndex)
				converted, err := s.mapToKafkaBootstrapServer(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("bootstrap_servers") {
				details.BootstrapServers = tmp
			}
		}
		if consumerProperties, ok := s.D.GetOkExists("consumer_properties"); ok {
			tmp := consumerProperties.(string)
			details.ConsumerProperties = &tmp
		}
		if keyStore, ok := s.D.GetOkExists("key_store"); ok {
			tmp := keyStore.(string)
			details.KeyStore = &tmp
		}
		if keyStorePassword, ok := s.D.GetOkExists("key_store_password"); ok {
			tmp := keyStorePassword.(string)
			details.KeyStorePassword = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if producerProperties, ok := s.D.GetOkExists("producer_properties"); ok {
			tmp := producerProperties.(string)
			details.ProducerProperties = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.KafkaConnectionSecurityProtocolEnum(securityProtocol.(string))
		}
		if sslKeyPassword, ok := s.D.GetOkExists("ssl_key_password"); ok {
			tmp := sslKeyPassword.(string)
			details.SslKeyPassword = &tmp
		}
		if streamPoolId, ok := s.D.GetOkExists("stream_pool_id"); ok {
			tmp := streamPoolId.(string)
			details.StreamPoolId = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.KafkaConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if trustStore, ok := s.D.GetOkExists("trust_store"); ok {
			tmp := trustStore.(string)
			details.TrustStore = &tmp
		}
		if trustStorePassword, ok := s.D.GetOkExists("trust_store_password"); ok {
			tmp := trustStorePassword.(string)
			details.TrustStorePassword = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
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
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("KAFKA_SCHEMA_REGISTRY"):
		details := oci_golden_gate.CreateKafkaSchemaRegistryConnectionDetails{}
		if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
			details.AuthenticationType = oci_golden_gate.KafkaSchemaRegistryConnectionAuthenticationTypeEnum(authenticationType.(string))
		}
		if keyStore, ok := s.D.GetOkExists("key_store"); ok {
			tmp := keyStore.(string)
			details.KeyStore = &tmp
		}
		if keyStorePassword, ok := s.D.GetOkExists("key_store_password"); ok {
			tmp := keyStorePassword.(string)
			details.KeyStorePassword = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if privateIp, ok := s.D.GetOkExists("private_ip"); ok {
			tmp := privateIp.(string)
			details.PrivateIp = &tmp
		}
		if sslKeyPassword, ok := s.D.GetOkExists("ssl_key_password"); ok {
			tmp := sslKeyPassword.(string)
			details.SslKeyPassword = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.KafkaSchemaRegistryConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if trustStore, ok := s.D.GetOkExists("trust_store"); ok {
			tmp := trustStore.(string)
			details.TrustStore = &tmp
		}
		if trustStorePassword, ok := s.D.GetOkExists("trust_store_password"); ok {
			tmp := trustStorePassword.(string)
			details.TrustStorePassword = &tmp
		}
		if url, ok := s.D.GetOkExists("url"); ok {
			tmp := url.(string)
			details.Url = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
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
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("MICROSOFT_SQLSERVER"):
		details := oci_golden_gate.CreateMicrosoftSqlserverConnectionDetails{}
		if additionalAttributes, ok := s.D.GetOkExists("additional_attributes"); ok {
			interfaces := additionalAttributes.([]interface{})
			tmp := make([]oci_golden_gate.NameValuePair, len(interfaces))
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
		if host, ok := s.D.GetOkExists("host"); ok {
			tmp := host.(string)
			details.Host = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if port, ok := s.D.GetOkExists("port"); ok {
			tmp := port.(int)
			details.Port = &tmp
		}
		if privateIp, ok := s.D.GetOkExists("private_ip"); ok {
			tmp := privateIp.(string)
			details.PrivateIp = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.MicrosoftSqlserverConnectionSecurityProtocolEnum(securityProtocol.(string))
		}
		if shouldValidateServerCertificate, ok := s.D.GetOkExists("should_validate_server_certificate"); ok {
			tmp := shouldValidateServerCertificate.(bool)
			details.ShouldValidateServerCertificate = &tmp
		}
		if sslCa, ok := s.D.GetOkExists("ssl_ca"); ok {
			tmp := sslCa.(string)
			details.SslCa = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.MicrosoftSqlserverConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
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
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("MONGODB"):
		details := oci_golden_gate.CreateMongoDbConnectionDetails{}
		if connectionString, ok := s.D.GetOkExists("connection_string"); ok {
			tmp := connectionString.(string)
			details.ConnectionString = &tmp
		}
		if databaseId, ok := s.D.GetOkExists("database_id"); ok {
			tmp := databaseId.(string)
			details.DatabaseId = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.MongoDbConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
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
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("MYSQL"):
		details := oci_golden_gate.CreateMysqlConnectionDetails{}
		if additionalAttributes, ok := s.D.GetOkExists("additional_attributes"); ok {
			interfaces := additionalAttributes.([]interface{})
			tmp := make([]oci_golden_gate.NameValuePair, len(interfaces))
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
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if port, ok := s.D.GetOkExists("port"); ok {
			tmp := port.(int)
			details.Port = &tmp
		}
		if privateIp, ok := s.D.GetOkExists("private_ip"); ok {
			tmp := privateIp.(string)
			details.PrivateIp = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.MysqlConnectionSecurityProtocolEnum(securityProtocol.(string))
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
			details.SslMode = oci_golden_gate.MysqlConnectionSslModeEnum(sslMode.(string))
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.MysqlConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
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
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("OCI_OBJECT_STORAGE"):
		details := oci_golden_gate.CreateOciObjectStorageConnectionDetails{}
		if privateKeyFile, ok := s.D.GetOkExists("private_key_file"); ok {
			tmp := privateKeyFile.(string)
			details.PrivateKeyFile = &tmp
		}
		if privateKeyPassphrase, ok := s.D.GetOkExists("private_key_passphrase"); ok {
			tmp := privateKeyPassphrase.(string)
			details.PrivateKeyPassphrase = &tmp
		}
		if publicKeyFingerprint, ok := s.D.GetOkExists("public_key_fingerprint"); ok {
			tmp := publicKeyFingerprint.(string)
			details.PublicKeyFingerprint = &tmp
		}
		if region, ok := s.D.GetOkExists("region"); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.OciObjectStorageConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if tenancyId, ok := s.D.GetOkExists("tenancy_id"); ok {
			tmp := tenancyId.(string)
			details.TenancyId = &tmp
		}
		if userId, ok := s.D.GetOkExists("user_id"); ok {
			tmp := userId.(string)
			details.UserId = &tmp
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
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("ORACLE"):
		details := oci_golden_gate.CreateOracleConnectionDetails{}
		if authenticationMode, ok := s.D.GetOkExists("authentication_mode"); ok {
			details.AuthenticationMode = oci_golden_gate.OracleConnectionAuthenticationModeEnum(authenticationMode.(string))
		}
		if connectionString, ok := s.D.GetOkExists("connection_string"); ok {
			tmp := connectionString.(string)
			details.ConnectionString = &tmp
		}
		if databaseId, ok := s.D.GetOkExists("database_id"); ok {
			tmp := databaseId.(string)
			details.DatabaseId = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if privateIp, ok := s.D.GetOkExists("private_ip"); ok {
			tmp := privateIp.(string)
			details.PrivateIp = &tmp
		}
		if sessionMode, ok := s.D.GetOkExists("session_mode"); ok {
			details.SessionMode = oci_golden_gate.OracleConnectionSessionModeEnum(sessionMode.(string))
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.OracleConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
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
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("ORACLE_NOSQL"):
		details := oci_golden_gate.CreateOracleNosqlConnectionDetails{}
		if privateKeyFile, ok := s.D.GetOkExists("private_key_file"); ok {
			tmp := privateKeyFile.(string)
			details.PrivateKeyFile = &tmp
		}
		if privateKeyPassphrase, ok := s.D.GetOkExists("private_key_passphrase"); ok {
			tmp := privateKeyPassphrase.(string)
			details.PrivateKeyPassphrase = &tmp
		}
		if publicKeyFingerprint, ok := s.D.GetOkExists("public_key_fingerprint"); ok {
			tmp := publicKeyFingerprint.(string)
			details.PublicKeyFingerprint = &tmp
		}
		if region, ok := s.D.GetOkExists("region"); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.OracleNosqlConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if tenancyId, ok := s.D.GetOkExists("tenancy_id"); ok {
			tmp := tenancyId.(string)
			details.TenancyId = &tmp
		}
		if userId, ok := s.D.GetOkExists("user_id"); ok {
			tmp := userId.(string)
			details.UserId = &tmp
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
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("POSTGRESQL"):
		details := oci_golden_gate.CreatePostgresqlConnectionDetails{}
		if additionalAttributes, ok := s.D.GetOkExists("additional_attributes"); ok {
			interfaces := additionalAttributes.([]interface{})
			tmp := make([]oci_golden_gate.NameValuePair, len(interfaces))
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
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if port, ok := s.D.GetOkExists("port"); ok {
			tmp := port.(int)
			details.Port = &tmp
		}
		if privateIp, ok := s.D.GetOkExists("private_ip"); ok {
			tmp := privateIp.(string)
			details.PrivateIp = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.PostgresqlConnectionSecurityProtocolEnum(securityProtocol.(string))
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
			details.SslMode = oci_golden_gate.PostgresqlConnectionSslModeEnum(sslMode.(string))
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.PostgresqlConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
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
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("REDIS"):
		details := oci_golden_gate.CreateRedisConnectionDetails{}
		if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
			details.AuthenticationType = oci_golden_gate.RedisConnectionAuthenticationTypeEnum(authenticationType.(string))
		}
		if keyStore, ok := s.D.GetOkExists("key_store"); ok {
			tmp := keyStore.(string)
			details.KeyStore = &tmp
		}
		if keyStorePassword, ok := s.D.GetOkExists("key_store_password"); ok {
			tmp := keyStorePassword.(string)
			details.KeyStorePassword = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if redisClusterId, ok := s.D.GetOkExists("redis_cluster_id"); ok {
			tmp := redisClusterId.(string)
			details.RedisClusterId = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.RedisConnectionSecurityProtocolEnum(securityProtocol.(string))

		}
		if servers, ok := s.D.GetOkExists("servers"); ok {
			tmp := servers.(string)
			details.Servers = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.RedisConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if trustStore, ok := s.D.GetOkExists("trust_store"); ok {
			tmp := trustStore.(string)
			details.TrustStore = &tmp
		}
		if trustStorePassword, ok := s.D.GetOkExists("trust_store_password"); ok {
			tmp := trustStorePassword.(string)
			details.TrustStorePassword = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
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
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("SNOWFLAKE"):
		details := oci_golden_gate.CreateSnowflakeConnectionDetails{}
		if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
			details.AuthenticationType = oci_golden_gate.SnowflakeConnectionAuthenticationTypeEnum(authenticationType.(string))
		}
		if connectionUrl, ok := s.D.GetOkExists("connection_url"); ok {
			tmp := connectionUrl.(string)
			details.ConnectionUrl = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if privateKeyFile, ok := s.D.GetOkExists("private_key_file"); ok {
			tmp := privateKeyFile.(string)
			details.PrivateKeyFile = &tmp
		}
		if privateKeyPassphrase, ok := s.D.GetOkExists("private_key_passphrase"); ok {
			tmp := privateKeyPassphrase.(string)
			details.PrivateKeyPassphrase = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.SnowflakeConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
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
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
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

func (s *GoldenGateConnectionResourceCrud) populateTopLevelPolymorphicUpdateConnectionRequest(request *oci_golden_gate.UpdateConnectionRequest) error {
	//discriminator
	connectionTypeRaw, ok := s.D.GetOkExists("connection_type")
	var connectionType string
	if ok {
		connectionType = connectionTypeRaw.(string)
	} else {
		connectionType = "" // default value
	}
	switch strings.ToLower(connectionType) {
	case strings.ToLower("AMAZON_KINESIS"):
		details := oci_golden_gate.UpdateAmazonKinesisConnectionDetails{}
		if accessKeyId, ok := s.D.GetOkExists("access_key_id"); ok {
			tmp := accessKeyId.(string)
			details.AccessKeyId = &tmp
		}
		if secretAccessKey, ok := s.D.GetOkExists("secret_access_key"); ok {
			tmp := secretAccessKey.(string)
			details.SecretAccessKey = &tmp
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
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("AMAZON_REDSHIFT"):
		details := oci_golden_gate.UpdateAmazonRedshiftConnectionDetails{}
		if connectionUrl, ok := s.D.GetOkExists("connection_url"); ok {
			tmp := connectionUrl.(string)
			details.ConnectionUrl = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
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
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("AMAZON_S3"):
		details := oci_golden_gate.UpdateAmazonS3ConnectionDetails{}
		if accessKeyId, ok := s.D.GetOkExists("access_key_id"); ok {
			tmp := accessKeyId.(string)
			details.AccessKeyId = &tmp
		}
		if secretAccessKey, ok := s.D.GetOkExists("secret_access_key"); ok {
			tmp := secretAccessKey.(string)
			details.SecretAccessKey = &tmp
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
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("AZURE_DATA_LAKE_STORAGE"):
		details := oci_golden_gate.UpdateAzureDataLakeStorageConnectionDetails{}
		if accountKey, ok := s.D.GetOkExists("account_key"); ok {
			tmp := accountKey.(string)
			details.AccountKey = &tmp
		}
		if accountName, ok := s.D.GetOkExists("account_name"); ok {
			tmp := accountName.(string)
			details.AccountName = &tmp
		}
		if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
			details.AuthenticationType = oci_golden_gate.AzureDataLakeStorageConnectionAuthenticationTypeEnum(authenticationType.(string))
		}
		if azureTenantId, ok := s.D.GetOkExists("azure_tenant_id"); ok {
			tmp := azureTenantId.(string)
			details.AzureTenantId = &tmp
		}
		if clientId, ok := s.D.GetOkExists("client_id"); ok {
			tmp := clientId.(string)
			details.ClientId = &tmp
		}
		if clientSecret, ok := s.D.GetOkExists("client_secret"); ok {
			tmp := clientSecret.(string)
			details.ClientSecret = &tmp
		}
		if endpoint, ok := s.D.GetOkExists("endpoint"); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		if sasToken, ok := s.D.GetOkExists("sas_token"); ok {
			tmp := sasToken.(string)
			details.SasToken = &tmp
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
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("AZURE_SYNAPSE_ANALYTICS"):
		details := oci_golden_gate.UpdateAzureSynapseConnectionDetails{}
		if connectionString, ok := s.D.GetOkExists("connection_string"); ok {
			tmp := connectionString.(string)
			details.ConnectionString = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
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
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("DB2"):
		details := oci_golden_gate.UpdateDb2ConnectionDetails{}
		if additionalAttributes, ok := s.D.GetOkExists("additional_attributes"); ok {
			interfaces := additionalAttributes.([]interface{})
			tmp := make([]oci_golden_gate.NameValuePair, len(interfaces))
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
		if host, ok := s.D.GetOkExists("host"); ok {
			tmp := host.(string)
			details.Host = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if port, ok := s.D.GetOkExists("port"); ok {
			tmp := port.(int)
			details.Port = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.Db2ConnectionSecurityProtocolEnum(securityProtocol.(string))
		}
		if sslClientKeystash, ok := s.D.GetOkExists("ssl_client_keystash"); ok {
			tmp := sslClientKeystash.(string)
			details.SslClientKeystash = &tmp
		}
		if sslClientKeystoredb, ok := s.D.GetOkExists("ssl_client_keystoredb"); ok {
			tmp := sslClientKeystoredb.(string)
			details.SslClientKeystoredb = &tmp
		}
		if sslServerCertificate, ok := s.D.GetOkExists("ssl_server_certificate"); ok {
			tmp := sslServerCertificate.(string)
			details.SslServerCertificate = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("ELASTICSEARCH"):
		details := oci_golden_gate.UpdateElasticsearchConnectionDetails{}
		if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
			details.AuthenticationType = oci_golden_gate.ElasticsearchConnectionAuthenticationTypeEnum(authenticationType.(string))
		}
		if fingerprint, ok := s.D.GetOkExists("fingerprint"); ok {
			tmp := fingerprint.(string)
			details.Fingerprint = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.ElasticsearchConnectionSecurityProtocolEnum(securityProtocol.(string))

		}
		if servers, ok := s.D.GetOkExists("servers"); ok {
			tmp := servers.(string)
			details.Servers = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
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
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("GENERIC"):
		details := oci_golden_gate.UpdateGenericConnectionDetails{}
		if host, ok := s.D.GetOkExists("host"); ok {
			tmp := host.(string)
			details.Host = &tmp
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
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("GOLDENGATE"):
		details := oci_golden_gate.UpdateGoldenGateConnectionDetails{}
		if deploymentId, ok := s.D.GetOkExists("deployment_id"); ok {
			tmp := deploymentId.(string)
			details.DeploymentId = &tmp
		}
		if host, ok := s.D.GetOkExists("host"); ok {
			tmp := host.(string)
			details.Host = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if port, ok := s.D.GetOkExists("port"); ok {
			tmp := port.(int)
			details.Port = &tmp
		}
		if privateIp, ok := s.D.GetOkExists("private_ip"); ok {
			tmp := privateIp.(string)
			details.PrivateIp = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
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
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("GOOGLE_BIGQUERY"):
		details := oci_golden_gate.UpdateGoogleBigQueryConnectionDetails{}
		if serviceAccountKeyFile, ok := s.D.GetOkExists("service_account_key_file"); ok {
			tmp := serviceAccountKeyFile.(string)
			details.ServiceAccountKeyFile = &tmp
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
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("GOOGLE_CLOUD_STORAGE"):
		details := oci_golden_gate.UpdateGoogleCloudStorageConnectionDetails{}
		if serviceAccountKeyFile, ok := s.D.GetOkExists("service_account_key_file"); ok {
			tmp := serviceAccountKeyFile.(string)
			details.ServiceAccountKeyFile = &tmp
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
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("HDFS"):
		details := oci_golden_gate.UpdateHdfsConnectionDetails{}
		if coreSiteXml, ok := s.D.GetOkExists("core_site_xml"); ok {
			tmp := coreSiteXml.(string)
			details.CoreSiteXml = &tmp
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
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("JAVA_MESSAGE_SERVICE"):
		details := oci_golden_gate.UpdateJavaMessageServiceConnectionDetails{}
		if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
			details.AuthenticationType = oci_golden_gate.JavaMessageServiceConnectionAuthenticationTypeEnum(authenticationType.(string))
		}
		if connectionFactory, ok := s.D.GetOkExists("connection_factory"); ok {
			tmp := connectionFactory.(string)
			details.ConnectionFactory = &tmp
		}
		if connectionUrl, ok := s.D.GetOkExists("connection_url"); ok {
			tmp := connectionUrl.(string)
			details.ConnectionUrl = &tmp
		}
		if jndiConnectionFactory, ok := s.D.GetOkExists("jndi_connection_factory"); ok {
			tmp := jndiConnectionFactory.(string)
			details.JndiConnectionFactory = &tmp
		}
		if jndiInitialContextFactory, ok := s.D.GetOkExists("jndi_initial_context_factory"); ok {
			tmp := jndiInitialContextFactory.(string)
			details.JndiInitialContextFactory = &tmp
		}
		if jndiProviderUrl, ok := s.D.GetOkExists("jndi_provider_url"); ok {
			tmp := jndiProviderUrl.(string)
			details.JndiProviderUrl = &tmp
		}
		if jndiSecurityCredentials, ok := s.D.GetOkExists("jndi_security_credentials"); ok {
			tmp := jndiSecurityCredentials.(string)
			details.JndiSecurityCredentials = &tmp
		}
		if jndiSecurityPrincipal, ok := s.D.GetOkExists("jndi_security_principal"); ok {
			tmp := jndiSecurityPrincipal.(string)
			details.JndiSecurityPrincipal = &tmp
		}
		if keyStore, ok := s.D.GetOkExists("key_store"); ok {
			tmp := keyStore.(string)
			details.KeyStore = &tmp
		}
		if keyStorePassword, ok := s.D.GetOkExists("key_store_password"); ok {
			tmp := keyStorePassword.(string)
			details.KeyStorePassword = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if privateIp, ok := s.D.GetOkExists("private_ip"); ok {
			tmp := privateIp.(string)
			details.PrivateIp = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.JavaMessageServiceConnectionSecurityProtocolEnum(securityProtocol.(string))
		}
		if shouldUseJndi, ok := s.D.GetOkExists("should_use_jndi"); ok {
			tmp := shouldUseJndi.(bool)
			details.ShouldUseJndi = &tmp
		}
		if sslKeyPassword, ok := s.D.GetOkExists("ssl_key_password"); ok {
			tmp := sslKeyPassword.(string)
			details.SslKeyPassword = &tmp
		}
		if trustStore, ok := s.D.GetOkExists("trust_store"); ok {
			tmp := trustStore.(string)
			details.TrustStore = &tmp
		}
		if trustStorePassword, ok := s.D.GetOkExists("trust_store_password"); ok {
			tmp := trustStorePassword.(string)
			details.TrustStorePassword = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
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
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("KAFKA"):
		details := oci_golden_gate.UpdateKafkaConnectionDetails{}
		if bootstrapServers, ok := s.D.GetOkExists("bootstrap_servers"); ok {
			interfaces := bootstrapServers.([]interface{})
			tmp := make([]oci_golden_gate.KafkaBootstrapServer, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "bootstrap_servers", stateDataIndex)
				converted, err := s.mapToKafkaBootstrapServer(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("bootstrap_servers") {
				details.BootstrapServers = tmp
			}
		}
		if consumerProperties, ok := s.D.GetOkExists("consumer_properties"); ok {
			tmp := consumerProperties.(string)
			details.ConsumerProperties = &tmp
		}
		if keyStore, ok := s.D.GetOkExists("key_store"); ok {
			tmp := keyStore.(string)
			details.KeyStore = &tmp
		}
		if keyStorePassword, ok := s.D.GetOkExists("key_store_password"); ok {
			tmp := keyStorePassword.(string)
			details.KeyStorePassword = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if producerProperties, ok := s.D.GetOkExists("producer_properties"); ok {
			tmp := producerProperties.(string)
			details.ProducerProperties = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.KafkaConnectionSecurityProtocolEnum(securityProtocol.(string))
		}
		if sslKeyPassword, ok := s.D.GetOkExists("ssl_key_password"); ok {
			tmp := sslKeyPassword.(string)
			details.SslKeyPassword = &tmp
		}
		if streamPoolId, ok := s.D.GetOkExists("stream_pool_id"); ok {
			tmp := streamPoolId.(string)
			details.StreamPoolId = &tmp
		}
		if trustStore, ok := s.D.GetOkExists("trust_store"); ok {
			tmp := trustStore.(string)
			details.TrustStore = &tmp
		}
		if trustStorePassword, ok := s.D.GetOkExists("trust_store_password"); ok {
			tmp := trustStorePassword.(string)
			details.TrustStorePassword = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
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
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("KAFKA_SCHEMA_REGISTRY"):
		details := oci_golden_gate.UpdateKafkaSchemaRegistryConnectionDetails{}
		if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
			details.AuthenticationType = oci_golden_gate.KafkaSchemaRegistryConnectionAuthenticationTypeEnum(authenticationType.(string))
		}
		if keyStore, ok := s.D.GetOkExists("key_store"); ok {
			tmp := keyStore.(string)
			details.KeyStore = &tmp
		}
		if keyStorePassword, ok := s.D.GetOkExists("key_store_password"); ok {
			tmp := keyStorePassword.(string)
			details.KeyStorePassword = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if privateIp, ok := s.D.GetOkExists("private_ip"); ok {
			tmp := privateIp.(string)
			details.PrivateIp = &tmp
		}
		if sslKeyPassword, ok := s.D.GetOkExists("ssl_key_password"); ok {
			tmp := sslKeyPassword.(string)
			details.SslKeyPassword = &tmp
		}
		if trustStore, ok := s.D.GetOkExists("trust_store"); ok {
			tmp := trustStore.(string)
			details.TrustStore = &tmp
		}
		if trustStorePassword, ok := s.D.GetOkExists("trust_store_password"); ok {
			tmp := trustStorePassword.(string)
			details.TrustStorePassword = &tmp
		}
		if url, ok := s.D.GetOkExists("url"); ok {
			tmp := url.(string)
			details.Url = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
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
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("MICROSOFT_SQLSERVER"):
		details := oci_golden_gate.UpdateMicrosoftSqlserverConnectionDetails{}
		if additionalAttributes, ok := s.D.GetOkExists("additional_attributes"); ok {
			interfaces := additionalAttributes.([]interface{})
			tmp := make([]oci_golden_gate.NameValuePair, len(interfaces))
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
		if host, ok := s.D.GetOkExists("host"); ok {
			tmp := host.(string)
			details.Host = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if port, ok := s.D.GetOkExists("port"); ok {
			tmp := port.(int)
			details.Port = &tmp
		}
		if privateIp, ok := s.D.GetOkExists("private_ip"); ok {
			tmp := privateIp.(string)
			details.PrivateIp = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.MicrosoftSqlserverConnectionSecurityProtocolEnum(securityProtocol.(string))
		}
		if shouldValidateServerCertificate, ok := s.D.GetOkExists("should_validate_server_certificate"); ok {
			tmp := shouldValidateServerCertificate.(bool)
			details.ShouldValidateServerCertificate = &tmp
		}
		if sslCa, ok := s.D.GetOkExists("ssl_ca"); ok {
			tmp := sslCa.(string)
			details.SslCa = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
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
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("MONGODB"):
		details := oci_golden_gate.UpdateMongoDbConnectionDetails{}
		if connectionString, ok := s.D.GetOkExists("connection_string"); ok {
			tmp := connectionString.(string)
			details.ConnectionString = &tmp
		}
		if databaseId, ok := s.D.GetOkExists("database_id"); ok {
			tmp := databaseId.(string)
			details.DatabaseId = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
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
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("MYSQL"):
		details := oci_golden_gate.UpdateMysqlConnectionDetails{}
		if additionalAttributes, ok := s.D.GetOkExists("additional_attributes"); ok {
			interfaces := additionalAttributes.([]interface{})
			tmp := make([]oci_golden_gate.NameValuePair, len(interfaces))
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
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if port, ok := s.D.GetOkExists("port"); ok {
			tmp := port.(int)
			details.Port = &tmp
		}
		if privateIp, ok := s.D.GetOkExists("private_ip"); ok {
			tmp := privateIp.(string)
			details.PrivateIp = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.MysqlConnectionSecurityProtocolEnum(securityProtocol.(string))
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
			details.SslMode = oci_golden_gate.MysqlConnectionSslModeEnum(sslMode.(string))
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
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
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("OCI_OBJECT_STORAGE"):
		details := oci_golden_gate.UpdateOciObjectStorageConnectionDetails{}
		if privateKeyFile, ok := s.D.GetOkExists("private_key_file"); ok {
			tmp := privateKeyFile.(string)
			details.PrivateKeyFile = &tmp
		}
		if privateKeyPassphrase, ok := s.D.GetOkExists("private_key_passphrase"); ok {
			tmp := privateKeyPassphrase.(string)
			details.PrivateKeyPassphrase = &tmp
		}
		if publicKeyFingerprint, ok := s.D.GetOkExists("public_key_fingerprint"); ok {
			tmp := publicKeyFingerprint.(string)
			details.PublicKeyFingerprint = &tmp
		}
		if region, ok := s.D.GetOkExists("region"); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if tenancyId, ok := s.D.GetOkExists("tenancy_id"); ok {
			tmp := tenancyId.(string)
			details.TenancyId = &tmp
		}
		if userId, ok := s.D.GetOkExists("user_id"); ok {
			tmp := userId.(string)
			details.UserId = &tmp
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
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("ORACLE"):
		details := oci_golden_gate.UpdateOracleConnectionDetails{}
		if authenticationMode, ok := s.D.GetOkExists("authentication_mode"); ok {
			details.AuthenticationMode = oci_golden_gate.OracleConnectionAuthenticationModeEnum(authenticationMode.(string))
		}
		if connectionString, ok := s.D.GetOkExists("connection_string"); ok {
			tmp := connectionString.(string)
			details.ConnectionString = &tmp
		}
		if databaseId, ok := s.D.GetOkExists("database_id"); ok {
			tmp := databaseId.(string)
			details.DatabaseId = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if privateIp, ok := s.D.GetOkExists("private_ip"); ok {
			tmp := privateIp.(string)
			details.PrivateIp = &tmp
		}
		if sessionMode, ok := s.D.GetOkExists("session_mode"); ok {
			details.SessionMode = oci_golden_gate.OracleConnectionSessionModeEnum(sessionMode.(string))
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
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
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("ORACLE_NOSQL"):
		details := oci_golden_gate.UpdateOracleNosqlConnectionDetails{}
		if privateKeyFile, ok := s.D.GetOkExists("private_key_file"); ok {
			tmp := privateKeyFile.(string)
			details.PrivateKeyFile = &tmp
		}
		if privateKeyPassphrase, ok := s.D.GetOkExists("private_key_passphrase"); ok {
			tmp := privateKeyPassphrase.(string)
			details.PrivateKeyPassphrase = &tmp
		}
		if publicKeyFingerprint, ok := s.D.GetOkExists("public_key_fingerprint"); ok {
			tmp := publicKeyFingerprint.(string)
			details.PublicKeyFingerprint = &tmp
		}
		if region, ok := s.D.GetOkExists("region"); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if tenancyId, ok := s.D.GetOkExists("tenancy_id"); ok {
			tmp := tenancyId.(string)
			details.TenancyId = &tmp
		}
		if userId, ok := s.D.GetOkExists("user_id"); ok {
			tmp := userId.(string)
			details.UserId = &tmp
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
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("POSTGRESQL"):
		details := oci_golden_gate.UpdatePostgresqlConnectionDetails{}
		if additionalAttributes, ok := s.D.GetOkExists("additional_attributes"); ok {
			interfaces := additionalAttributes.([]interface{})
			tmp := make([]oci_golden_gate.NameValuePair, len(interfaces))
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
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if port, ok := s.D.GetOkExists("port"); ok {
			tmp := port.(int)
			details.Port = &tmp
		}
		if privateIp, ok := s.D.GetOkExists("private_ip"); ok {
			tmp := privateIp.(string)
			details.PrivateIp = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.PostgresqlConnectionSecurityProtocolEnum(securityProtocol.(string))
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
			details.SslMode = oci_golden_gate.PostgresqlConnectionSslModeEnum(sslMode.(string))
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
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
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("REDIS"):
		details := oci_golden_gate.UpdateRedisConnectionDetails{}
		if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
			details.AuthenticationType = oci_golden_gate.RedisConnectionAuthenticationTypeEnum(authenticationType.(string))
		}
		if keyStore, ok := s.D.GetOkExists("key_store"); ok {
			tmp := keyStore.(string)
			details.KeyStore = &tmp
		}
		if keyStorePassword, ok := s.D.GetOkExists("key_store_password"); ok {
			tmp := keyStorePassword.(string)
			details.KeyStorePassword = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if redisClusterId, ok := s.D.GetOkExists("redis_cluster_id"); ok {
			tmp := redisClusterId.(string)
			details.RedisClusterId = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.RedisConnectionSecurityProtocolEnum(securityProtocol.(string))
		}
		if servers, ok := s.D.GetOkExists("servers"); ok {
			tmp := servers.(string)
			details.Servers = &tmp
		}
		if trustStore, ok := s.D.GetOkExists("trust_store"); ok {
			tmp := trustStore.(string)
			details.TrustStore = &tmp
		}
		if trustStorePassword, ok := s.D.GetOkExists("trust_store_password"); ok {
			tmp := trustStorePassword.(string)
			details.TrustStorePassword = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
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
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("SNOWFLAKE"):
		details := oci_golden_gate.UpdateSnowflakeConnectionDetails{}
		if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
			details.AuthenticationType = oci_golden_gate.SnowflakeConnectionAuthenticationTypeEnum(authenticationType.(string))
		}
		if connectionUrl, ok := s.D.GetOkExists("connection_url"); ok {
			tmp := connectionUrl.(string)
			details.ConnectionUrl = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if privateKeyFile, ok := s.D.GetOkExists("private_key_file"); ok {
			tmp := privateKeyFile.(string)
			details.PrivateKeyFile = &tmp
		}
		if privateKeyPassphrase, ok := s.D.GetOkExists("private_key_passphrase"); ok {
			tmp := privateKeyPassphrase.(string)
			details.PrivateKeyPassphrase = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
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
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
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
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
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

func (s *GoldenGateConnectionResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_golden_gate.ChangeConnectionCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ConnectionId = &idTmp

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		changeCompartmentRequest.IsLockOverride = &tmp
	}

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.ChangeConnectionCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getConnectionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate"), oci_golden_gate.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
