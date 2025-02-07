// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package golden_gate

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GoldenGateConnectionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["connection_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(GoldenGateConnectionResource(), fieldMap, readSingularGoldenGateConnection)
}

func readSingularGoldenGateConnection(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateConnectionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.ReadResource(sync)
}

type GoldenGateConnectionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_golden_gate.GoldenGateClient
	Res    *oci_golden_gate.GetConnectionResponse
}

func (s *GoldenGateConnectionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GoldenGateConnectionDataSourceCrud) Get() error {
	request := oci_golden_gate.GetConnectionRequest{}

	if connectionId, ok := s.D.GetOkExists("connection_id"); ok {
		tmp := connectionId.(string)
		request.ConnectionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "golden_gate")

	response, err := s.Client.GetConnection(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *GoldenGateConnectionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.GetId())
	switch v := (s.Res.Connection).(type) {
	case oci_golden_gate.AmazonKinesisConnection:
		s.D.Set("connection_type", "AMAZON_KINESIS")

		if v.AccessKeyId != nil {
			s.D.Set("access_key_id", *v.AccessKeyId)
		}

		if v.SecretAccessKeySecretId != nil {
			s.D.Set("secret_access_key_secret_id", *v.SecretAccessKeySecretId)
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

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

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

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
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

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

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

		if v.SecretAccessKeySecretId != nil {
			s.D.Set("secret_access_key_secret_id", *v.SecretAccessKeySecretId)
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

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

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

		if v.AccountKeySecretId != nil {
			s.D.Set("account_key_secret_id", *v.AccountKeySecretId)
		}

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

		if v.ClientSecretSecretId != nil {
			s.D.Set("client_secret_secret_id", *v.ClientSecretSecretId)
		}

		if v.Endpoint != nil {
			s.D.Set("endpoint", *v.Endpoint)
		}

		if v.SasTokenSecretId != nil {
			s.D.Set("sas_token_secret_id", *v.SasTokenSecretId)
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

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

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

		s.D.Set("technology_type", v.TechnologyType)

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

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
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

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

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
	case oci_golden_gate.DatabricksConnection:
		s.D.Set("connection_type", "DATABRICKS")

		s.D.Set("authentication_type", v.AuthenticationType)

		if v.ClientId != nil {
			s.D.Set("client_id", *v.ClientId)
		}

		if v.ClientSecretSecretId != nil {
			s.D.Set("client_secret_secret_id", *v.ClientSecretSecretId)
		}

		if v.ConnectionUrl != nil {
			s.D.Set("connection_url", *v.ConnectionUrl)
		}

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
		}

		if v.StorageCredentialName != nil {
			s.D.Set("storage_credential_name", *v.StorageCredentialName)
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

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

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

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
		}

		if v.Port != nil {
			s.D.Set("port", *v.Port)
		}

		s.D.Set("security_protocol", v.SecurityProtocol)

		if v.SslClientKeystashSecretId != nil {
			s.D.Set("ssl_client_keystash_secret_id", *v.SslClientKeystashSecretId)
		}

		if v.SslClientKeystoredbSecretId != nil {
			s.D.Set("ssl_client_keystoredb_secret_id", *v.SslClientKeystoredbSecretId)
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

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

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
	case oci_golden_gate.ElasticsearchConnection:
		s.D.Set("connection_type", "ELASTICSEARCH")

		s.D.Set("authentication_type", v.AuthenticationType)

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
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

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

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

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

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

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
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

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

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

		if v.ServiceAccountKeyFileSecretId != nil {
			s.D.Set("service_account_key_file_secret_id", *v.ServiceAccountKeyFileSecretId)
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

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

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

		if v.ServiceAccountKeyFileSecretId != nil {
			s.D.Set("service_account_key_file_secret_id", *v.ServiceAccountKeyFileSecretId)
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

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

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
	case oci_golden_gate.GooglePubSubConnection:
		s.D.Set("connection_type", "GOOGLE_PUBSUB")

		if v.ServiceAccountKeyFileSecretId != nil {
			s.D.Set("service_account_key_file_secret_id", *v.ServiceAccountKeyFileSecretId)
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

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

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

		s.D.Set("technology_type", v.TechnologyType)

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

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

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

		if v.JndiSecurityCredentialsSecretId != nil {
			s.D.Set("jndi_security_credentials_secret_id", *v.JndiSecurityCredentialsSecretId)
		}

		if v.JndiSecurityPrincipal != nil {
			s.D.Set("jndi_security_principal", *v.JndiSecurityPrincipal)
		}

		if v.KeyStorePasswordSecretId != nil {
			s.D.Set("key_store_password_secret_id", *v.KeyStorePasswordSecretId)
		}

		if v.KeyStoreSecretId != nil {
			s.D.Set("key_store_secret_id", *v.KeyStoreSecretId)
		}

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
		}

		if v.PrivateIp != nil {
			s.D.Set("private_ip", *v.PrivateIp)
		}

		s.D.Set("security_protocol", v.SecurityProtocol)

		if v.ShouldUseJndi != nil {
			s.D.Set("should_use_jndi", *v.ShouldUseJndi)
		}

		if v.SslKeyPasswordSecretId != nil {
			s.D.Set("ssl_key_password_secret_id", *v.SslKeyPasswordSecretId)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.TrustStorePasswordSecretId != nil {
			s.D.Set("trust_store_password_secret_id", *v.TrustStorePasswordSecretId)
		}

		if v.TrustStoreSecretId != nil {
			s.D.Set("trust_store_secret_id", *v.TrustStoreSecretId)
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

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

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

		if v.ConsumerProperties != nil {
			s.D.Set("consumer_properties", *v.ConsumerProperties)
		}

		if v.KeyStorePasswordSecretId != nil {
			s.D.Set("key_store_password_secret_id", *v.KeyStorePasswordSecretId)
		}

		if v.KeyStoreSecretId != nil {
			s.D.Set("key_store_secret_id", *v.KeyStoreSecretId)
		}

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
		}

		if v.ProducerProperties != nil {
			s.D.Set("producer_properties", *v.ProducerProperties)
		}

		s.D.Set("security_protocol", v.SecurityProtocol)

		if v.SslKeyPasswordSecretId != nil {
			s.D.Set("ssl_key_password_secret_id", *v.SslKeyPasswordSecretId)
		}

		if v.StreamPoolId != nil {
			s.D.Set("stream_pool_id", *v.StreamPoolId)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.TrustStorePasswordSecretId != nil {
			s.D.Set("trust_store_password_secret_id", *v.TrustStorePasswordSecretId)
		}

		if v.TrustStoreSecretId != nil {
			s.D.Set("trust_store_secret_id", *v.TrustStoreSecretId)
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

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

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

		if v.KeyStorePasswordSecretId != nil {
			s.D.Set("key_store_password_secret_id", *v.KeyStorePasswordSecretId)
		}

		if v.KeyStoreSecretId != nil {
			s.D.Set("key_store_secret_id", *v.KeyStoreSecretId)
		}

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
		}

		if v.PrivateIp != nil {
			s.D.Set("private_ip", *v.PrivateIp)
		}

		if v.SslKeyPasswordSecretId != nil {
			s.D.Set("ssl_key_password_secret_id", *v.SslKeyPasswordSecretId)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.TrustStorePasswordSecretId != nil {
			s.D.Set("trust_store_password_secret_id", *v.TrustStorePasswordSecretId)
		}

		if v.TrustStoreSecretId != nil {
			s.D.Set("trust_store_secret_id", *v.TrustStoreSecretId)
		}

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

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

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
	case oci_golden_gate.MicrosoftFabricConnection:
		s.D.Set("connection_type", "MICROSOFT_FABRIC")

		if v.ClientId != nil {
			s.D.Set("client_id", *v.ClientId)
		}

		if v.ClientSecretSecretId != nil {
			s.D.Set("client_secret_secret_id", *v.ClientSecretSecretId)
		}

		if v.Endpoint != nil {
			s.D.Set("endpoint", *v.Endpoint)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.TenantId != nil {
			s.D.Set("tenant_id", *v.TenantId)
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

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

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

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
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

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

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

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
		}

		s.D.Set("security_protocol", v.SecurityProtocol)

		s.D.Set("technology_type", v.TechnologyType)

		if v.TlsCertificateKeyFilePasswordSecretId != nil {
			s.D.Set("tls_certificate_key_file_password_secret_id", *v.TlsCertificateKeyFilePasswordSecretId)
		}

		if v.TlsCertificateKeyFileSecretId != nil {
			s.D.Set("tls_certificate_key_file_secret_id", *v.TlsCertificateKeyFileSecretId)
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

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

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

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
		}

		if v.Port != nil {
			s.D.Set("port", *v.Port)
		}

		if v.PrivateIp != nil {
			s.D.Set("private_ip", *v.PrivateIp)
		}

		s.D.Set("security_protocol", v.SecurityProtocol)

		if v.SslKeySecretId != nil {
			s.D.Set("ssl_key_secret_id", *v.SslKeySecretId)
		}

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

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

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

		if v.PrivateKeyFileSecretId != nil {
			s.D.Set("private_key_file_secret_id", *v.PrivateKeyFileSecretId)
		}

		if v.PrivateKeyPassphraseSecretId != nil {
			s.D.Set("private_key_passphrase_secret_id", *v.PrivateKeyPassphraseSecretId)
		}

		if v.Region != nil {
			s.D.Set("region", *v.Region)
		}

		if v.ShouldUseResourcePrincipal != nil {
			s.D.Set("should_use_resource_principal", *v.ShouldUseResourcePrincipal)
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

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

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

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
		}

		if v.PrivateIp != nil {
			s.D.Set("private_ip", *v.PrivateIp)
		}

		s.D.Set("session_mode", v.SessionMode)

		s.D.Set("technology_type", v.TechnologyType)

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

		if v.WalletSecretId != nil {
			s.D.Set("wallet_secret_id", *v.WalletSecretId)
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

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

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

		if v.PrivateKeyFileSecretId != nil {
			s.D.Set("private_key_file_secret_id", *v.PrivateKeyFileSecretId)
		}

		if v.PrivateKeyPassphraseSecretId != nil {
			s.D.Set("private_key_passphrase_secret_id", *v.PrivateKeyPassphraseSecretId)
		}

		if v.Region != nil {
			s.D.Set("region", *v.Region)
		}

		if v.ShouldUseResourcePrincipal != nil {
			s.D.Set("should_use_resource_principal", *v.ShouldUseResourcePrincipal)
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

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

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

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
		}

		if v.Port != nil {
			s.D.Set("port", *v.Port)
		}

		if v.PrivateIp != nil {
			s.D.Set("private_ip", *v.PrivateIp)
		}

		s.D.Set("security_protocol", v.SecurityProtocol)

		if v.SslKeySecretId != nil {
			s.D.Set("ssl_key_secret_id", *v.SslKeySecretId)
		}

		s.D.Set("ssl_mode", v.SslMode)

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

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

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

		s.D.Set("technology_type", v.TechnologyType)

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

		if v.KeyStorePasswordSecretId != nil {
			s.D.Set("key_store_password_secret_id", *v.KeyStorePasswordSecretId)
		}

		if v.KeyStoreSecretId != nil {
			s.D.Set("key_store_secret_id", *v.KeyStoreSecretId)
		}

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
		}

		if v.RedisClusterId != nil {
			s.D.Set("redis_cluster_id", *v.RedisClusterId)
		}

		s.D.Set("security_protocol", v.SecurityProtocol)

		if v.Servers != nil {
			s.D.Set("servers", *v.Servers)
		}

		if v.TrustStorePasswordSecretId != nil {
			s.D.Set("trust_store_password_secret_id", *v.TrustStorePasswordSecretId)
		}

		if v.TrustStoreSecretId != nil {
			s.D.Set("trust_store_secret_id", *v.TrustStoreSecretId)
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

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

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

		s.D.Set("technology_type", v.TechnologyType)

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

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
		}

		if v.PrivateKeyFileSecretId != nil {
			s.D.Set("private_key_file_secret_id", *v.PrivateKeyFileSecretId)
		}

		if v.PrivateKeyPassphraseSecretId != nil {
			s.D.Set("private_key_passphrase_secret_id", *v.PrivateKeyPassphraseSecretId)
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

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

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

		s.D.Set("technology_type", v.TechnologyType)

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
		log.Printf("[WARN] Received 'connection_type' of unknown type %v", s.Res.Connection)
		return nil
	}

	return nil
}
