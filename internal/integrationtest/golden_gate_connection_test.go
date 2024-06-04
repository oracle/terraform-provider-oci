// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
package integrationtest

/*
Note:
	For each connection test you must define a descriptor in connectionTestDescriptors array.
     One connectionTestDescriptor must define:
     - connectionType
     - technologyType
     - representation - which contains only the fields which are connectionType & technologyType specific.
   	   Note, that common connection attributes are represented in CommonConnectionRepresentation
     - excludesFieldsFromDataCheck - list of fields which are not presented as part of the data resource
	   Note: usually these are the sensitive fields

For each connectionTestDescriptor the following tests will be executed:
  0. create connection resource and validate the populated attributes value - also check that resource id is saved properly
  1. extend test 1. with singular datasource and validates
	- if singular datasource could load connection based on its id
    - all of it's populated attributes (except the ones which are excluded from the field list explicitly in excludesFieldsFromDataCheck)
  2. still keeps resource created by step 1., but updates updatable values, validates
    - that changing compartment can successfully move connection to a new compartment
    - that update connection is successful and connection is not recreated
  3. extend test with datasource, and validate that listConnection find actual connection based on its values
  4. verifies resource import
 (5. deletes connections)

Following environment variables are required:
 - TF_VAR_compartment_id
 - TF_VAR_compartment_id_for_move
 - TF_VAR_kms_key_id
 - TF_VAR_vault_id
 - TF_VAR_subnet_id
 - TF_VAR_oracle_wallet - for oracle connection creation
 - TF_VAR_password - password used for create connection, due to sec central issues, we must use environment variables instead of hardcoded passwords
 - TF_VAR_new_password - new password used for update connection, due to sec central issues, we must use environment variables instead of hardcoded passwords
"
*/
import (
	"context"
	b64 "encoding/base64"
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

type ConnectionTestDescriptor struct {
	connectionType              oci_golden_gate.ConnectionTypeEnum
	technologyType              oci_golden_gate.TechnologyTypeEnum
	representation              map[string]interface{}
	excludedFieldsFromDataCheck []string
}

var (
	// TODO this representation is used only by deployment_test, so if both are modified, we should move it there
	GoldenGateConnectionRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"connection_type": acctest.Representation{RepType: acctest.Required, Create: `GOLDENGATE`},
		"display_name":    acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"technology_type": acctest.Representation{RepType: acctest.Required, Create: `GOLDENGATE`},
		"description":     acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"bar-key": "value"}},
		"key_id":          acctest.Representation{RepType: acctest.Optional, Create: `${var.kms_key_id}`},
		"private_ip":      acctest.Representation{RepType: acctest.Optional, Create: `10.0.1.78`},
		"subnet_id":       acctest.Representation{RepType: acctest.Optional, Create: `${var.subnet_id}`},
		"vault_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.vault_id}`},
	}

	EnabledConnectionTests = []oci_golden_gate.ConnectionTypeEnum{
		oci_golden_gate.ConnectionTypeAmazonS3,
		oci_golden_gate.ConnectionTypeAmazonKinesis,
		oci_golden_gate.ConnectionTypeAmazonRedshift,
		oci_golden_gate.ConnectionTypeAzureDataLakeStorage,
		oci_golden_gate.ConnectionTypeAzureSynapseAnalytics,
		//		oci_golden_gate.ConnectionTypeCassandra, //TODO 2023. 09. 19. lmadaras: to be tested
		oci_golden_gate.ConnectionTypeDb2,
		oci_golden_gate.ConnectionTypeElasticsearch,
		oci_golden_gate.ConnectionTypeGoldengate,
		oci_golden_gate.ConnectionTypeGeneric,
		oci_golden_gate.ConnectionTypeGoogleBigquery,
		oci_golden_gate.ConnectionTypeGoogleCloudStorage,
		oci_golden_gate.ConnectionTypeHdfs,
		oci_golden_gate.ConnectionTypeJavaMessageService,
		oci_golden_gate.ConnectionTypeKafka,
		oci_golden_gate.ConnectionTypeKafkaSchemaRegistry,
		oci_golden_gate.ConnectionTypeMicrosoftSqlserver,
		oci_golden_gate.ConnectionTypeMongodb,
		oci_golden_gate.ConnectionTypeMysql,
		oci_golden_gate.ConnectionTypeOciObjectStorage,
		oci_golden_gate.ConnectionTypeOracle,
		oci_golden_gate.ConnectionTypeOracleNosql,
		oci_golden_gate.ConnectionTypePostgresql,
		oci_golden_gate.ConnectionTypeRedis,
		oci_golden_gate.ConnectionTypeSnowflake,
	}

	CommonConnectionRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`, Update: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `TF-connection-test`, Update: `TF-connection-test-updated`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `new-description`},
		"key_id":         acctest.Representation{RepType: acctest.Optional, Create: `${var.kms_key_id}`},
		"subnet_id":      acctest.Representation{RepType: acctest.Optional, Update: `${var.subnet_id}`},
		"vault_id":       acctest.Representation{RepType: acctest.Optional, Create: `${var.vault_id}`},
		"routing_method": acctest.Representation{RepType: acctest.Optional, Create: `SHARED_DEPLOYMENT_ENDPOINT`, Update: `DEDICATED_ENDPOINT`},
	}

	connectionLocksRepresentation = map[string]interface{}{
		"type":    acctest.Representation{RepType: acctest.Required, Create: `FULL`},
		"message": acctest.Representation{RepType: acctest.Optional, Create: `message`},
	}

	ConnectionTestDescriptors = []ConnectionTestDescriptor{
		// AmazonS3, create a locked resource, only for this type. Resource locking is generic it applies the same way for the other types
		{connectionType: oci_golden_gate.ConnectionTypeAmazonS3, technologyType: oci_golden_gate.TechnologyTypeAmazonS3,
			representation: map[string]interface{}{
				// Override compartment to test move compartment too.
				"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`, Update: `${var.compartment_id_for_move}`},
				"access_key_id":     acctest.Representation{RepType: acctest.Required, Create: `AKIAIOSFODNN7EXAMPLE`, Update: `AKIAIOSFODNN7UPDATED`},
				"secret_access_key": acctest.Representation{RepType: acctest.Required, Create: `mysecret`},
				"locks":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: connectionLocksRepresentation},
				"is_lock_override":  acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
				"lifecycle":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsAndLocks},
			},
		},

		// Amazon Kinesis
		{connectionType: oci_golden_gate.ConnectionTypeAmazonKinesis, technologyType: oci_golden_gate.TechnologyTypeAmazonKinesis,
			representation: map[string]interface{}{
				"access_key_id":     acctest.Representation{RepType: acctest.Required, Create: `AKIAIOSFODNN7EXAMPLE`, Update: `AKIAIOSFODNN7UPDATED`},
				"secret_access_key": acctest.Representation{RepType: acctest.Required, Create: `mysecret`},
			},
		},

		// Amazon Redshift
		{connectionType: oci_golden_gate.ConnectionTypeAmazonRedshift, technologyType: oci_golden_gate.TechnologyTypeAmazonRedshift,
			representation: map[string]interface{}{
				"connection_url": acctest.Representation{RepType: acctest.Required,
					Create: `jdbc:redshift://aws-redshift-instance.aaaaaaaaaaaa.us-east-2.redshift.amazonaws.com:5439/mydb1`,
					Update: `jdbc:redshift://aws-redshift-instance.aaaaaaaaaaaa.us-east-2.redshift.amazonaws.com:5438/mydb2`},
				"username":          acctest.Representation{RepType: acctest.Required, Create: `user`, Update: `updatedUser`},
				"secret_access_key": acctest.Representation{RepType: acctest.Required, Create: `mysecret`},
				"password":          acctest.Representation{RepType: acctest.Required, Create: `${var.password}`, Update: `${var.new_password}`},
			},
		},

		// Azure DataLake
		{connectionType: oci_golden_gate.ConnectionTypeAzureDataLakeStorage, technologyType: oci_golden_gate.TechnologyTypeAzureDataLakeStorage,
			representation: map[string]interface{}{
				"authentication_type": acctest.Representation{RepType: acctest.Required, Create: string(oci_golden_gate.AzureDataLakeStorageConnectionAuthenticationTypeAzureActiveDirectory)},
				"account_name":        acctest.Representation{RepType: acctest.Required, Create: `myAccount`, Update: `updatedAccount`},
				"endpoint":            acctest.Representation{RepType: acctest.Required, Create: `https://whatever.com`, Update: `https://exactly.com`},
				"azure_tenant_id":     acctest.Representation{RepType: acctest.Required, Create: `14593954-d337-4a61-a364-9f758c64f97f`},
				"client_id":           acctest.Representation{RepType: acctest.Required, Create: `06ecaabf-8b80-4ec8-a0ec-20cbf463703d`},
				"client_secret":       acctest.Representation{RepType: acctest.Required, Create: `dO29Q~F5-VwnA.lZdd11xFF_t5NAXCaGwDl9NbT1`},
			},
		},

		// Azure Synapse
		{connectionType: oci_golden_gate.ConnectionTypeAzureSynapseAnalytics, technologyType: oci_golden_gate.TechnologyTypeAzureSynapseAnalytics,
			representation: map[string]interface{}{
				"connection_string": acctest.Representation{RepType: acctest.Required,
					Create: `jdbc:sqlserver://ws1.sql.azuresynapse.net:1433;database=db1;encrypt=true;trustServerCertificate=false;hostNameInCertificate=*.sql.azuresynapse.net;loginTimeout=300;'`,
					Update: `jdbc:sqlserver://ws1.sql.azuresynapse.net:1433;database=db2;encrypt=true;trustServerCertificate=false;hostNameInCertificate=*.sql.azuresynapse.net;loginTimeout=300;'`},
				"username": acctest.Representation{RepType: acctest.Required, Create: `user`, Update: `updatedUser`},
				"password": acctest.Representation{RepType: acctest.Required, Create: `${var.password}`, Update: `${var.new_password}`},
			},
		},

		//// Casandra
		//{connectionType: oci_golden_gate.ConnectionTypeCassandra, technologyType: oci_golden_gate.TechnologyTypeApacheCassandra,
		//	representation: map[string]interface{}{
		//		"contact_points": acctest.Representation{RepType: acctest.Required,
		//			Create: `server1.example.com,server2.example.com`,
		//			Update: `server3.example.com:9043,server4.example.com:9043`},
		//		"username": acctest.Representation{RepType: acctest.Required, Create: `user`, Update: `updatedUser`},
		//		"password": acctest.Representation{RepType: acctest.Required, Create: `${var.password}`, Update: `${var.new_password}`},
		//		"config_file": acctest.Representation{RepType: acctest.Required,
		//			Create: b64.StdEncoding.EncodeToString([]byte(
		//				"cluster_name: 'MyCluster1'\n" +
		//					"initial_token: null\n" +
		//					"listen_address: localhost\n" +
		//					"rpc_address: localhost\nseed_provider:\n" +
		//					"  - class_name: org.apache.cassandra.locator.SimpleSeedProvider\n" +
		//					"    parameters:\n" +
		//					"      - seeds: \"127.0.0.1\"\n" +
		//					"data_file_directories:\n" +
		//					"  - /var/lib/cassandra/data\n" +
		//					"commitlog_directory: /var/lib/cassandra/commitlog\n" +
		//					"saved_caches_directory: /var/lib/cassandra/saved_caches\n" +
		//					"disk_failure_policy: stop\n" +
		//					"memtable_allocation_type: heap_buffers\nmemtable_flush_writers: 0\n")),
		//			Update: b64.StdEncoding.EncodeToString([]byte(
		//				"cluster_name: 'MyCluster2'\n" +
		//					"initial_token: null\n" +
		//					"listen_address: localhost\n" +
		//					"rpc_address: localhost\nseed_provider:\n" +
		//					"  - class_name: org.apache.cassandra.locator.SimpleSeedProvider\n" +
		//					"    parameters:\n" +
		//					"      - seeds: \"127.0.0.1\"\n" +
		//					"data_file_directories:\n" +
		//					"  - /var/lib/cassandra/data\n" +
		//					"commitlog_directory: /var/lib/cassandra/commitlog\n" +
		//					"saved_caches_directory: /var/lib/cassandra/saved_caches\n" +
		//					"disk_failure_policy: stop\n" +
		//					"memtable_allocation_type: heap_buffers\nmemtable_flush_writers: 0\n"))},
		//	},
		//},
		//
		// DB2
		{connectionType: oci_golden_gate.ConnectionTypeDb2, technologyType: oci_golden_gate.TechnologyTypeDb2Zos,
			representation: map[string]interface{}{
				"database_name":     acctest.Representation{RepType: acctest.Required, Create: `database1`, Update: `database2`},
				"host":              acctest.Representation{RepType: acctest.Required, Create: `whatever1.fqdn.com`, Update: `whatever2.fqdn.com`},
				"port":              acctest.Representation{RepType: acctest.Required, Create: `10000`, Update: `10001`},
				"username":          acctest.Representation{RepType: acctest.Required, Create: `admin`, Update: `new_admin`},
				"password":          acctest.Representation{RepType: acctest.Required, Create: `${var.password}`, Update: `${var.new_password}`},
				"security_protocol": acctest.Representation{RepType: acctest.Required, Create: string(oci_golden_gate.Db2ConnectionSecurityProtocolPlain)},
			},
		},

		// Elasticsearch
		{connectionType: oci_golden_gate.ConnectionTypeElasticsearch, technologyType: oci_golden_gate.TechnologyTypeElasticsearch,
			representation: map[string]interface{}{
				"servers": acctest.Representation{RepType: acctest.Required,
					Create: `server1.example.com:9200,server2.example.com:9200`,
					Update: `server3.example.com:9201,server4.example.com:9202`},
				"security_protocol": acctest.Representation{RepType: acctest.Required, Create: string(oci_golden_gate.ElasticsearchConnectionSecurityProtocolPlain)},
				"authentication_type": acctest.Representation{RepType: acctest.Required,
					Create: string(oci_golden_gate.ElasticsearchConnectionAuthenticationTypeNone),
					Update: string(oci_golden_gate.ElasticsearchConnectionAuthenticationTypeBasic)},
				"username": acctest.Representation{RepType: acctest.Optional, Update: `new_admin`},
				"password": acctest.Representation{RepType: acctest.Optional, Update: `${var.new_password}`},
			},
		},

		// Generic
		{connectionType: oci_golden_gate.ConnectionTypeGeneric, technologyType: oci_golden_gate.TechnologyTypeGeneric,
			representation: map[string]interface{}{
				"host": acctest.Representation{RepType: acctest.Required,
					Create: `server1.example.com:1111,`,
					Update: `server2.example.com:2222,server3.example.com:3333`},
			},
		},

		// Google BigQuery
		{connectionType: oci_golden_gate.ConnectionTypeGoogleBigquery, technologyType: oci_golden_gate.TechnologyTypeGoogleBigquery,
			representation: map[string]interface{}{
				"service_account_key_file": acctest.Representation{RepType: acctest.Required,
					Create: b64.StdEncoding.EncodeToString([]byte(
						"{\n  \"type\": \"service_account\",\n" +
							"  \"project_id\": \"your-project-id\",\n" +
							"  \"private_key_id\": \"your-private-key-id\",\n" +
							"  \"private_key\": \"-----BEGIN PRIVATE KEY-----\\nYour_Private_Key_Here\\n-----END PRIVATE KEY-----\",\n" +
							"  \"client_email\": \"your-service-account-email@your-project-id.iam.gserviceaccount.com\",\n" +
							"  \"client_id\": \"your-client-id1\",\n" +
							"  \"auth_uri\": \"https://accounts.google.com/o/oauth2/auth\",\n" +
							"  \"token_uri\": \"https://accounts.google.com/o/oauth2/token\",\n" +
							"  \"auth_provider_x509_cert_url\": \"https://www.googleapis.com/oauth2/v1/certs\",\n" +
							"  \"client_x509_cert_url\": \"https://www.googleapis.com/robot/v1/metadata/x509/your-service-account-email%40your-project-id.iam.gserviceaccount.com\"\n}\n")),
					Update: b64.StdEncoding.EncodeToString([]byte(
						"{\n  \"type\": \"service_account\",\n" +
							"  \"project_id\": \"your-project-id2\",\n" +
							"  \"private_key_id\": \"your-private-key-id2\",\n" +
							"  \"private_key\": \"-----BEGIN PRIVATE KEY-----\\nYour_Private_Key2_Here\\n-----END PRIVATE KEY-----\",\n" +
							"  \"client_email\": \"your-service-account-email@your-project-id.iam.gserviceaccount.com\",\n" +
							"  \"client_id\": \"your-client-id2\",\n" +
							"  \"auth_uri\": \"https://accounts.google.com/o/oauth2/auth\",\n" +
							"  \"token_uri\": \"https://accounts.google.com/o/oauth2/token\",\n" +
							"  \"auth_provider_x509_cert_url\": \"https://www.googleapis.com/oauth2/v1/certs\",\n" +
							"  \"client_x509_cert_url\": \"https://www.googleapis.com/robot/v1/metadata/x509/your-service-account-email%40your-project-id.iam.gserviceaccount.com\"\n}\n"))},
			},
		},

		// Google Cloud Storage
		{connectionType: oci_golden_gate.ConnectionTypeGoogleCloudStorage, technologyType: oci_golden_gate.TechnologyTypeGoogleCloudStorage,
			representation: map[string]interface{}{
				"service_account_key_file": acctest.Representation{RepType: acctest.Required,
					Create: b64.StdEncoding.EncodeToString([]byte(
						"{\n  \"type\": \"service_account\",\n" +
							"  \"project_id\": \"your-project-id\",\n" +
							"  \"private_key_id\": \"your-private-key-id\",\n" +
							"  \"private_key\": \"-----BEGIN PRIVATE KEY-----\\nYour_Private_Key_Here\\n-----END PRIVATE KEY-----\",\n" +
							"  \"client_email\": \"your-service-account-email@your-project-id.iam.gserviceaccount.com\",\n" +
							"  \"client_id\": \"your-client-id1\",\n" +
							"  \"auth_uri\": \"https://accounts.google.com/o/oauth2/auth\",\n" +
							"  \"token_uri\": \"https://accounts.google.com/o/oauth2/token\",\n" +
							"  \"auth_provider_x509_cert_url\": \"https://www.googleapis.com/oauth2/v1/certs\",\n" +
							"  \"client_x509_cert_url\": \"https://www.googleapis.com/robot/v1/metadata/x509/your-service-account-email%40your-project-id.iam.gserviceaccount.com\"\n}\n")),
					Update: b64.StdEncoding.EncodeToString([]byte(
						"{\n  \"type\": \"service_account\",\n" +
							"  \"project_id\": \"your-project-id\",\n" +
							"  \"private_key_id\": \"your-private-key-id\",\n" +
							"  \"private_key\": \"-----BEGIN PRIVATE KEY-----\\nYour_Private_Key_Here\\n-----END PRIVATE KEY-----\",\n" +
							"  \"client_email\": \"your-service-account-email@your-project-id.iam.gserviceaccount.com\",\n" +
							"  \"client_id\": \"your-client-id2\",\n" +
							"  \"auth_uri\": \"https://accounts.google.com/o/oauth2/auth\",\n" +
							"  \"token_uri\": \"https://accounts.google.com/o/oauth2/token\",\n" +
							"  \"auth_provider_x509_cert_url\": \"https://www.googleapis.com/oauth2/v1/certs\",\n" +
							"  \"client_x509_cert_url\": \"https://www.googleapis.com/robot/v1/metadata/x509/your-service-account-email%40your-project-id.iam.gserviceaccount.com\"\n}\n"))},
			},
		},

		// Goldengate
		{connectionType: oci_golden_gate.ConnectionTypeGoldengate, technologyType: oci_golden_gate.TechnologyTypeGoldengate,
			representation: map[string]interface{}{
				"host":       acctest.Representation{RepType: acctest.Required, Create: `goldengate.oci.oraclecloud.com`, Update: `goldengate2.oci.oraclecloud.com`},
				"port":       acctest.Representation{RepType: acctest.Required, Create: `9090`, Update: `9091`},
				"password":   acctest.Representation{RepType: acctest.Required, Create: `${var.password}`, Update: `${var.new_password}`},
				"username":   acctest.Representation{RepType: acctest.Required, Create: `user`, Update: `updatedUser`},
				"private_ip": acctest.Representation{RepType: acctest.Required, Create: `10.0.0.1`, Update: `10.0.0.2`},
			},
		},

		// HDFS
		{connectionType: oci_golden_gate.ConnectionTypeHdfs, technologyType: oci_golden_gate.TechnologyTypeHdfs, representation: map[string]interface{}{
			"core_site_xml": acctest.Representation{RepType: acctest.Required,
				Create: b64.StdEncoding.EncodeToString([]byte(
					"<?xml version=\"1.0\"?><?xml-stylesheet type=\"text/xsl\" href=\"configuration.xsl\"?><configuration>" +
						"<property><name>fs.defaultFS</name><value>hdfs://foo.bar.com:8020</value><description>DefaultDescription</description><final>true</final></property></configuration>")),
				Update: b64.StdEncoding.EncodeToString([]byte(
					"<?xml version=\"1.0\"?><?xml-stylesheet type=\"text/xsl\" href=\"configuration.xsl\"?><configuration>" +
						"<property><name>fs.defaultFS</name><value>hdfs://foo.bar.com:8021</value><description>UpdatedDescription</description><final>true</final></property></configuration>"))},
		},
		},

		// JMS
		{connectionType: oci_golden_gate.ConnectionTypeJavaMessageService, technologyType: oci_golden_gate.TechnologyTypeOracleWeblogicJms, representation: map[string]interface{}{
			"should_use_jndi":    acctest.Representation{RepType: acctest.Required, Create: `false`},
			"connection_url":     acctest.Representation{RepType: acctest.Required, Create: `mq://foo.bar.com:7676`, Update: `mq://foo.bar.com:7677`},
			"connection_factory": acctest.Representation{RepType: acctest.Required, Create: `com.stc.jmsjca.core.JConnectionFactoryXA`, Update: `mq://foo.bar.com:7677`},
			"password":           acctest.Representation{RepType: acctest.Required, Create: `${var.password}`, Update: `${var.new_password}`},
			"username":           acctest.Representation{RepType: acctest.Required, Create: `user`, Update: `updatedUser`},
			"private_ip":         acctest.Representation{RepType: acctest.Required, Create: `10.0.0.1`, Update: `10.0.0.2`},
		},
		},

		// Kafka
		{connectionType: oci_golden_gate.ConnectionTypeKafka, technologyType: oci_golden_gate.TechnologyTypeApacheKafka,
			representation: map[string]interface{}{
				"security_protocol": acctest.Representation{RepType: acctest.Required, Create: string(oci_golden_gate.KafkaConnectionSecurityProtocolSaslSsl)},
				"username":          acctest.Representation{RepType: acctest.Required, Create: `username`, Update: `newUsername`},
				"password":          acctest.Representation{RepType: acctest.Required, Create: `${var.password}`, Update: `${var.new_password}`},
				"bootstrap_servers": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
					"host":       acctest.Representation{RepType: acctest.Required, Create: `whatever.fqdn.oraclecloud.com`},
					"port":       acctest.Representation{RepType: acctest.Required, Create: `9093`},
					"private_ip": acctest.Representation{RepType: acctest.Required, Create: `10.0.0.1`},
				}},
			},
			excludedFieldsFromDataCheck: []string{"bootstrap_servers"},
		},

		// Kafka Schema Registry
		{connectionType: oci_golden_gate.ConnectionTypeKafkaSchemaRegistry, technologyType: oci_golden_gate.TechnologyTypeConfluentSchemaRegistry,
			representation: map[string]interface{}{
				"authentication_type": acctest.Representation{RepType: acctest.Required, Create: string(oci_golden_gate.KafkaSchemaRegistryConnectionAuthenticationTypeBasic)},
				"username":            acctest.Representation{RepType: acctest.Required, Create: `username`, Update: `newUsername`},
				"password":            acctest.Representation{RepType: acctest.Required, Create: `${var.password}`, Update: `${var.new_password}`},
				"url":                 acctest.Representation{RepType: acctest.Required, Create: `https://10.1.1.1:9091`, Update: `https://10.1.1.2:9091`},
			},
		},

		// Microsoft SQLServer
		{connectionType: oci_golden_gate.ConnectionTypeMicrosoftSqlserver, technologyType: oci_golden_gate.TechnologyTypeMicrosoftSqlserver,
			representation: map[string]interface{}{
				"security_protocol": acctest.Representation{RepType: acctest.Required, Create: string(oci_golden_gate.MicrosoftSqlserverConnectionSecurityProtocolPlain)},
				"database_name":     acctest.Representation{RepType: acctest.Required, Create: `database`, Update: `database2`},
				"host":              acctest.Representation{RepType: acctest.Required, Create: `whatever.fqdn.com`, Update: `whatever.fqdn.com`},
				"port":              acctest.Representation{RepType: acctest.Required, Create: `10000`, Update: `10001`},
				"username":          acctest.Representation{RepType: acctest.Required, Create: `username`, Update: `newUsername`},
				"password":          acctest.Representation{RepType: acctest.Required, Create: `${var.password}`, Update: `${var.new_password}`},
				"private_ip":        acctest.Representation{RepType: acctest.Required, Create: `10.0.0.1`, Update: `10.0.0.2`},
			},
		},

		// MongoDb
		{connectionType: oci_golden_gate.ConnectionTypeMongodb, technologyType: oci_golden_gate.TechnologyTypeMongodb,
			representation: map[string]interface{}{
				"connection_string": acctest.Representation{RepType: acctest.Required, Create: `mongodb://10.0.0.1:9000`,
					Update: `mongodb://10.0.0.1:9001`},
				"username": acctest.Representation{RepType: acctest.Required, Create: `username`, Update: `newUsername`},
				"password": acctest.Representation{RepType: acctest.Required, Create: `${var.password}`, Update: `${var.new_password}`},
			},
		},

		// MYSQL
		{connectionType: oci_golden_gate.ConnectionTypeMysql, technologyType: oci_golden_gate.TechnologyTypeOciMysql,
			representation: map[string]interface{}{
				"username":          acctest.Representation{RepType: acctest.Required, Create: `username`, Update: `newUsername`},
				"password":          acctest.Representation{RepType: acctest.Required, Create: `${var.password}`, Update: `${var.new_password}`},
				"database_name":     acctest.Representation{RepType: acctest.Required, Create: `database`, Update: `anotherdatabase`},
				"security_protocol": acctest.Representation{RepType: acctest.Required, Create: string(oci_golden_gate.MysqlConnectionSecurityProtocolPlain)},
				"private_ip":        acctest.Representation{RepType: acctest.Required, Create: `10.0.0.1`, Update: `10.0.0.2`},
				"host":              acctest.Representation{RepType: acctest.Required, Create: `validfqdn.com`, Update: `updatedfqdn.com`},
				"port":              acctest.Representation{RepType: acctest.Required, Create: `10000`, Update: `10001`},
			},
		},

		// OCI ObjectStorage
		{connectionType: oci_golden_gate.ConnectionTypeOciObjectStorage, technologyType: oci_golden_gate.TechnologyTypeOciObjectStorage,
			representation: map[string]interface{}{
				"tenancy_id": acctest.Representation{RepType: acctest.Required, Create: `ocid1.tenancy.oc1..fakeaaaak44klio2gjjzh2jknouk2xeh5w5fzpbybtljoanddubxipbtfake`,
					Update: `ocid1.tenancy.oc2..fakeaaaak44klio2gjjzh2jknouk2xeh5w5fzpbybtljoanddubxipbtfake`},
				"region": acctest.Representation{RepType: acctest.Required, Create: `us-phoenix-1`, Update: `us-ashburn-1`},
				"user_id": acctest.Representation{RepType: acctest.Required, Create: `ocid1.user.oc1..fakeaaaatswfukd4gymkjhngu3yp7galhoqzax6mi4ypgdt44ggbjaz2fake`,
					Update: `ocid1.user.oc2..fakeaaaatswfukd4gymkjhngu3yp7galhoqzax6mi4ypgdt44ggbjaz2fake`},
				"private_key_file":       acctest.Representation{RepType: acctest.Required, Create: `my-private-key-file`, Update: `new-private-key-file`},
				"private_key_passphrase": acctest.Representation{RepType: acctest.Required, Create: `${var.password}`, Update: `${var.new_password}`},
				"public_key_fingerprint": acctest.Representation{RepType: acctest.Required, Create: `myfingerprint`, Update: `newfingerprint`},
			},
		},

		// Oracle
		{connectionType: oci_golden_gate.ConnectionTypeOracle, technologyType: oci_golden_gate.TechnologyTypeAmazonRdsOracle,
			representation: map[string]interface{}{
				"username":          acctest.Representation{RepType: acctest.Required, Create: `username`, Update: `newUsername`},
				"password":          acctest.Representation{RepType: acctest.Required, Create: `${var.password}`, Update: `${var.new_password}`},
				"session_mode":      acctest.Representation{RepType: acctest.Required, Create: string(oci_golden_gate.OracleConnectionSessionModeDirect)},
				"connection_string": acctest.Representation{RepType: acctest.Required, Create: `alert-henry-IMUny-dev7-ggs.sub05140125230.integrationvcn.oraclevcn.com:1521/DB0609_phx2hg.sub05140125230.integrationvcn.oraclevcn.com`},
				"wallet":            acctest.Representation{RepType: acctest.Required, Create: `${var.oracle_wallet}`},
			},
		},

		// OracleNoSql
		{connectionType: oci_golden_gate.ConnectionTypeOracleNosql, technologyType: oci_golden_gate.TechnologyTypeOracleNosql,
			representation: map[string]interface{}{
				"private_key_file":       acctest.Representation{RepType: acctest.Required, Create: `my-private-key-file`, Update: `new-private-key-file`},
				"public_key_fingerprint": acctest.Representation{RepType: acctest.Required, Create: `myfingerprint`, Update: `newfingerprint`},
				"user_id": acctest.Representation{RepType: acctest.Required, Create: `ocid1.user.oc1..fakeaaaatswfukd4gymkjhngu3yp7galhoqzax6mi4ypgdt44ggbjaz2fake`,
					Update: `ocid1.user.oc2..fakeaaaatswfukd4gymkjhngu3yp7galhoqzax6mi4ypgdt44ggbjaz2fake`},
				"tenancy_id": acctest.Representation{RepType: acctest.Required, Create: `ocid1.tenancy.oc1..fakeaaaak44klio2gjjzh2jknouk2xeh5w5fzpbybtljoanddubxipbtfake`,
					Update: `ocid1.tenancy.oc2..fakeaaaak44klio2gjjzh2jknouk2xeh5w5fzpbybtljoanddubxipbtfake`},
				"region": acctest.Representation{RepType: acctest.Required, Create: `us-phoenix-1`, Update: `us-ashburn-1`},
			},
		},

		// Postgresql
		{connectionType: oci_golden_gate.ConnectionTypePostgresql, technologyType: oci_golden_gate.TechnologyTypePostgresqlServer,
			representation: map[string]interface{}{
				"database_name":     acctest.Representation{RepType: acctest.Required, Create: `database`, Update: `database2`},
				"host":              acctest.Representation{RepType: acctest.Required, Create: `whatever.fqdn.com`, Update: `whatever.fqdn.com`},
				"port":              acctest.Representation{RepType: acctest.Required, Create: `10000`, Update: `10001`},
				"username":          acctest.Representation{RepType: acctest.Required, Create: `admin`, Update: `new_admin`},
				"password":          acctest.Representation{RepType: acctest.Required, Create: `${var.password}`, Update: `${var.new_password}`},
				"security_protocol": acctest.Representation{RepType: acctest.Required, Create: string(oci_golden_gate.PostgresqlConnectionSecurityProtocolPlain)},
				"private_ip":        acctest.Representation{RepType: acctest.Required, Create: `10.0.0.1`, Update: `10.0.0.2`},
			},
		},

		// Redis
		{connectionType: oci_golden_gate.ConnectionTypeRedis, technologyType: oci_golden_gate.TechnologyTypeRedis,
			representation: map[string]interface{}{
				"servers": acctest.Representation{RepType: acctest.Required,
					Create: `server1.example.com:6379,server2.example.com:6379`,
					Update: `server3.example.com:6372,server4.example.com:6378`},
				"security_protocol": acctest.Representation{RepType: acctest.Required, Create: string(oci_golden_gate.RedisConnectionSecurityProtocolPlain)},
				"authentication_type": acctest.Representation{RepType: acctest.Required,
					Create: string(oci_golden_gate.RedisConnectionAuthenticationTypeNone),
					Update: string(oci_golden_gate.RedisConnectionAuthenticationTypeBasic)},
				"username": acctest.Representation{RepType: acctest.Optional, Update: `new_admin`},
				"password": acctest.Representation{RepType: acctest.Optional, Update: `${var.new_password}`},
			},
		},

		// Snowflake
		{connectionType: oci_golden_gate.ConnectionTypeSnowflake, technologyType: oci_golden_gate.TechnologyTypeSnowflake,
			representation: map[string]interface{}{
				"connection_url": acctest.Representation{RepType: acctest.Required, Create: `jdbc:snowflake://myaccount.snowflakecomputing.com/?warehouse=dawarehous&db=database`,
					Update: `jdbc:snowflake://myaccount.snowflakecomputing.com/?warehouse=dawarehous2&db=database2`},
				"authentication_type": acctest.Representation{RepType: acctest.Required, Create: string(oci_golden_gate.SnowflakeConnectionAuthenticationTypeBasic)},
				"username":            acctest.Representation{RepType: acctest.Required, Create: `admin`, Update: `new_admin`},
				"password":            acctest.Representation{RepType: acctest.Required, Create: `${var.password}`, Update: `${var.new_password}`},
			},
		},
	}

	ExcludedFields = []string{
		"account_key",
		"client_secret",
		"consumer_properties",
		"key_store",
		"key_store_password",
		"password",
		"private_key_file",
		"private_key_passphrase",
		"producer_properties",
		"public_key_fingerprint",
		"sas_token",
		"ssl_ca",
		"ssl_cert",
		"ssl_crl",
		"ssl_key",
		"ssl_key_password",
		"trust_store",
		"trust_store_password",
		"wallet",
		"core_site_xml",
		"secret_access_key",
		"service_account_key_file",
		"is_lock_override",
	}
)

// issue-routing-tag: golden_gate/default
func TestGoldenGateConnectionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGoldenGateConnectionResource_basic")
	defer httpreplay.SaveScenario()

	const (
		COMPARTMENT_ID          = "compartment_id"
		COMPARTMENT_ID_FOR_MOVE = "compartment_id_for_move"
		KMS_KEY_ID              = "kms_key_id"
		SUBNET_ID               = "subnet_id"
		VAULT_ID                = "vault_id"
		CONNECTION_TYPE         = "connection_type"
		TECHNOLOGY_TYPE         = "technology_type"
		ORACLE_WALLET           = "oracle_wallet"
		PASSWORD                = "password"
		NEW_PASSWORD            = "new_password"
	)

	config := acctest.ProviderTestConfig() +
		makeVariableStr(COMPARTMENT_ID, t) +
		makeVariableStr(COMPARTMENT_ID_FOR_MOVE, t) +
		makeVariableStr(KMS_KEY_ID, t) +
		makeVariableStr(SUBNET_ID, t) +
		makeVariableStr(VAULT_ID, t) +
		makeVariableStr(ORACLE_WALLET, t) +
		makeVariableStr(PASSWORD, t) +
		makeVariableStr(NEW_PASSWORD, t)

	var createResourcesConfig, dataSourceConfig, listDataSourceConfig, updateResourcesConfig string
	// CREATE CHECK FUNCTION MAPS
	resourceCheckFunctions := []resource.TestCheckFunc{}
	updatedResourceCheckFunctions := []resource.TestCheckFunc{}
	dataValidatorFunctions := []resource.TestCheckFunc{}
	dataSourcesValidatorFunctions := []resource.TestCheckFunc{}
	resourceNames := make([]string, 0, len(EnabledConnectionTests))
	for _, connectionTestDescriptor := range ConnectionTestDescriptors {
		if !containsConnection(EnabledConnectionTests, connectionTestDescriptor.connectionType) {
			log.Printf("Skip connectionType, because it's not enabled: %s", connectionTestDescriptor.connectionType)
			continue
		}
		// DEFINE RESOURCE NAMES
		connectionType := string(connectionTestDescriptor.connectionType)
		technologyType := string(connectionTestDescriptor.technologyType)
		resourceName := connectionType + "_" + technologyType
		checkResourceName := "oci_golden_gate_connection." + connectionType + "_" + technologyType
		resourceNames = append(resourceNames, checkResourceName)
		checkDataSourceName := "data.oci_golden_gate_connection." + connectionType + "_" + technologyType
		checkDataSourcesName := "data.oci_golden_gate_connections." + connectionType + "_" + technologyType

		// CREATE BASIC RESOURCE STRUCTURE
		connectionRepresentation := acctest.RepresentationCopyWithNewProperties(CommonConnectionRepresentation, map[string]interface{}{
			"connection_type": acctest.Representation{RepType: acctest.Required, Create: connectionType},
			"technology_type": acctest.Representation{RepType: acctest.Required, Create: technologyType},
		})
		connectionRepresentation = acctest.RepresentationCopyWithNewProperties(connectionRepresentation, connectionTestDescriptor.representation)
		// ADD connectionTypeCheck
		connectionTypeCheck := resource.TestCheckResourceAttr(checkResourceName, CONNECTION_TYPE, connectionType)
		resourceCheckFunctions = append(resourceCheckFunctions, connectionTypeCheck)
		updatedResourceCheckFunctions = append(updatedResourceCheckFunctions, connectionTypeCheck)
		dataValidatorFunctions = append(dataValidatorFunctions, resource.TestCheckResourceAttr(checkDataSourceName, CONNECTION_TYPE, connectionType))
		log.Printf("Check singular-data / resource: %s, property: %s, expected value: %s ", checkResourceName, CONNECTION_TYPE, connectionType)

		// ADD connectionTypeCheck
		technologyTypeCheck := resource.TestCheckResourceAttr(checkResourceName, TECHNOLOGY_TYPE, technologyType)
		resourceCheckFunctions = append(resourceCheckFunctions, technologyTypeCheck)
		updatedResourceCheckFunctions = append(updatedResourceCheckFunctions, technologyTypeCheck)
		dataValidatorFunctions = append(dataValidatorFunctions, resource.TestCheckResourceAttr(checkDataSourceName, TECHNOLOGY_TYPE, technologyType))
		log.Printf("Check singular-data / resource: %s, property: %s, expected value: %s ", checkResourceName, TECHNOLOGY_TYPE, technologyType)

		// ADD checks for all the attributes
		for propName, propertyRepresentation := range connectionRepresentation {
			if propName == CONNECTION_TYPE || propName == TECHNOLOGY_TYPE {
				continue
			}
			_, ok := propertyRepresentation.(acctest.Representation)
			if !ok {
				continue
			}

			if propertyRepresentation.(acctest.Representation).Create != nil {
				expectedPropertyValue := getPropertyValue(propertyRepresentation.(acctest.Representation).Create.(string))
				log.Printf("Check singular-data / resource: %s, property: %s, expected value: %s ", checkResourceName, propName, expectedPropertyValue)

				checkAttribute := resource.TestCheckResourceAttr(checkResourceName, propName, expectedPropertyValue)
				resourceCheckFunctions = append(resourceCheckFunctions, checkAttribute)
				if propertyRepresentation.(acctest.Representation).Update == nil {
					updatedResourceCheckFunctions = append(updatedResourceCheckFunctions, checkAttribute)
				}
				if !contains(connectionTestDescriptor.excludedFieldsFromDataCheck, propName) && !contains(ExcludedFields, propName) {
					dataValidatorFunctions = append(dataValidatorFunctions, resource.TestCheckResourceAttr(checkDataSourceName, propName, expectedPropertyValue))
				}
			}

			if propertyRepresentation.(acctest.Representation).Update != nil {
				expectedUpdatePropertyValue := getPropertyValue(propertyRepresentation.(acctest.Representation).Update.(string))

				log.Printf("Check update resource: %s, property: %s, expected value: %s ", checkResourceName, propName, expectedUpdatePropertyValue)
				updatedResourceCheckFunctions = append(updatedResourceCheckFunctions, resource.TestCheckResourceAttr(checkResourceName, propName, expectedUpdatePropertyValue))
			}
		}

		// ADD create validator function
		var resId, resId2 string
		resourceCheckFunctions = append(resourceCheckFunctions, func(s *terraform.State) (err error) {
			resId, err = acctest.FromInstanceState(s, checkResourceName, "id")
			return err
		})
		// Checks if the ingress IPs are empty after create - with routingMethod SHARED_DEPLOYMENT_ENDPOINT
		resourceCheckFunctions = append(resourceCheckFunctions, resource.TestCheckNoResourceAttr(checkResourceName, "ingressIps.#"))

		// ADD update validator function
		updatedResourceCheckFunctions = append(updatedResourceCheckFunctions, func(s *terraform.State) (err error) {
			resId2, err = acctest.FromInstanceState(s, checkResourceName, "id")
			if resId != resId2 {
				return fmt.Errorf("resource recreated when it was supposed to be updated")
			}
			return err
		})
		// Checks if the ingress IPs are populated after the uptade - with routingMethod DEDICATED_ENDPOINT
		updatedResourceCheckFunctions = append(updatedResourceCheckFunctions, resource.TestCheckResourceAttrSet(checkResourceName, "ingress_ips.#"))

		// DataSource representation
		resourceCompartmentLocation := connectionRepresentation["compartment_id"].(acctest.Representation).Update.(string)
		displayName := connectionRepresentation["display_name"].(acctest.Representation).Update.(string)
		dataSourceRepresentation := map[string]interface{}{
			"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: resourceCompartmentLocation},
			"connection_type": acctest.Representation{RepType: acctest.Optional, Create: []string{connectionType}},
			"display_name":    acctest.Representation{RepType: acctest.Optional, Create: displayName},
			"state":           acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
			"technology_type": acctest.Representation{RepType: acctest.Optional, Create: []string{technologyType}},
		}
		for propName, propertyRepresentation := range dataSourceRepresentation {
			if propName == CONNECTION_TYPE || propName == TECHNOLOGY_TYPE {
				continue
			}
			expectedPropertyValue := getPropertyValue(propertyRepresentation.(acctest.Representation).Create.(string))
			dataSourcesValidatorFunctions = append(dataSourcesValidatorFunctions, resource.TestCheckResourceAttr(checkDataSourcesName, propName, expectedPropertyValue))
		}
		dataSourcesValidatorFunctions = append(dataSourcesValidatorFunctions, resource.TestCheckResourceAttr(checkDataSourcesName, "connection_collection.#", "1"))
		dataSourcesValidatorFunctions = append(dataSourcesValidatorFunctions, resource.TestCheckResourceAttr(checkDataSourcesName, "connection_collection.0.items.#", "1"))

		createResourcesConfig += acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_connection", resourceName, acctest.Optional, acctest.Create, connectionRepresentation)
		dataSourceConfig += acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_connection", resourceName, acctest.Optional, acctest.Create, map[string]interface{}{
			"connection_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_golden_gate_connection.` + resourceName + `.id}`},
		})
		updateResourcesConfig += acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_connection", resourceName, acctest.Optional, acctest.Update, connectionRepresentation)
		listDataSourceConfig += acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_connection", resourceName, acctest.Optional, acctest.Update, connectionRepresentation) +
			acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_connections", resourceName, acctest.Optional, acctest.Create, dataSourceRepresentation)
	}
	steps := []resource.TestStep{
		// 0. resource test
		{
			Config: config + createResourcesConfig,
			Check:  acctest.ComposeAggregateTestCheckFuncArrayWrapper(resourceCheckFunctions),
		},
		// 1. singular datasource test
		{
			Config: config + createResourcesConfig + dataSourceConfig,
			Check:  acctest.ComposeAggregateTestCheckFuncArrayWrapper(dataValidatorFunctions),
		},
		// 2. update resource
		{
			Config: config + updateResourcesConfig,
			Check:  acctest.ComposeAggregateTestCheckFuncArrayWrapper(updatedResourceCheckFunctions),
		},
		// 3. datasource test
		{
			Config: config + listDataSourceConfig,
			Check:  acctest.ComposeAggregateTestCheckFuncArrayWrapper(dataSourcesValidatorFunctions),
		},
	}
	// It's not recommended to import multiple resources in the same test step
	// A separate step is executed for each connection's import
	// https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/import
	for _, resourceName := range resourceNames {
		steps = append(steps, resource.TestStep{
			Config:                  config + updateResourcesConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: ExcludedFields,
			ResourceName:            resourceName,
		})
	}
	// EXECUTE TESTS
	acctest.ResourceTest(t, testAccCheckGoldenGateConnectionDestroy, steps)

}

func testAccCheckGoldenGateConnectionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).GoldenGateClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_golden_gate_connection" {
			noResourceFound = false
			request := oci_golden_gate.GetConnectionRequest{}

			tmp := rs.Primary.ID
			request.ConnectionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "golden_gate")

			response, err := client.GetConnection(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_golden_gate.ConnectionLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.GetLifecycleState())
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("GoldenGateConnection") {
		resource.AddTestSweepers("GoldenGateConnection", &resource.Sweeper{
			Name:         "GoldenGateConnection",
			Dependencies: acctest.DependencyGraph["connection"],
			F:            sweepGoldenGateConnectionResource,
		})
	}
}

func sweepGoldenGateConnectionResource(compartment string) error {
	goldenGateClient := acctest.GetTestClients(&schema.ResourceData{}).GoldenGateClient()
	connectionIds, err := getGoldenGateConnectionIds(compartment)
	if err != nil {
		return err
	}
	for _, connectionId := range connectionIds {
		if ok := acctest.SweeperDefaultResourceId[connectionId]; !ok {
			deleteConnectionRequest := oci_golden_gate.DeleteConnectionRequest{}

			deleteConnectionRequest.ConnectionId = &connectionId

			var overrideLock = true
			deleteConnectionRequest.IsLockOverride = &overrideLock

			deleteConnectionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "golden_gate")
			_, error := goldenGateClient.DeleteConnection(context.Background(), deleteConnectionRequest)
			if error != nil {
				fmt.Printf("Error deleting Connection %s %s, It is possible that the resource is already deleted. Please verify manually \n", connectionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &connectionId, GoldenGateConnectionSweepWaitCondition, time.Duration(3*time.Minute),
				GoldenGateConnectionSweepResponseFetchOperation, "golden_gate", true)
		}
	}
	return nil
}

func getGoldenGateConnectionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ConnectionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	goldenGateClient := acctest.GetTestClients(&schema.ResourceData{}).GoldenGateClient()

	listConnectionsRequest := oci_golden_gate.ListConnectionsRequest{}
	listConnectionsRequest.CompartmentId = &compartmentId
	listConnectionsRequest.LifecycleState = oci_golden_gate.ConnectionLifecycleStateActive
	listConnectionsResponse, err := goldenGateClient.ListConnections(context.Background(), listConnectionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Connection list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, connection := range listConnectionsResponse.Items {
		id := *connection.GetId()
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ConnectionId", id)
	}
	return resourceIds, nil
}

func GoldenGateConnectionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if connectionResponse, ok := response.Response.(oci_golden_gate.GetConnectionResponse); ok {
		return connectionResponse.GetLifecycleState() != oci_golden_gate.ConnectionLifecycleStateDeleted
	}
	return false
}

func GoldenGateConnectionSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.GoldenGateClient().GetConnection(context.Background(), oci_golden_gate.GetConnectionRequest{
		ConnectionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

func makeVariableStrWithValue(envVar string, value string) string {
	return fmt.Sprintf("variable \"%s\" { default = \"%s\" }\n", envVar, value)
}

func getPropertyValue(propertyValue string) string {
	if propertyValue != "" && strings.HasPrefix(propertyValue, "${var.") {
		// it's a variable and its value must be loaded from
		propertyEnvVariable := strings.Replace(strings.Replace(propertyValue, "${var.", "", 1), "}", "", 1)
		return utils.GetEnvSettingWithBlankDefault(propertyEnvVariable)
	}
	return propertyValue
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func containsConnection(s []oci_golden_gate.ConnectionTypeEnum, connection oci_golden_gate.ConnectionTypeEnum) bool {
	for _, v := range s {
		if v == connection {
			return true
		}
	}

	return false
}
