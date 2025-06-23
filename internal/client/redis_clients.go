// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_redis "github.com/oracle/oci-go-sdk/v65/redis"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_redis.OciCacheConfigSetClient", &OracleClient{InitClientFn: initRedisOciCacheConfigSetClient})
	RegisterOracleClient("oci_redis.OciCacheDefaultConfigSetClient", &OracleClient{InitClientFn: initRedisOciCacheDefaultConfigSetClient})
	RegisterOracleClient("oci_redis.OciCacheUserClient", &OracleClient{InitClientFn: initRedisOciCacheUserClient})
	RegisterOracleClient("oci_redis.RedisClusterClient", &OracleClient{InitClientFn: initRedisRedisClusterClient})
	RegisterOracleClient("oci_redis.RedisIdentityClient", &OracleClient{InitClientFn: initRedisRedisIdentityClient})
}

func initRedisOciCacheConfigSetClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_redis.NewOciCacheConfigSetClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) OciCacheConfigSetClient() *oci_redis.OciCacheConfigSetClient {
	return m.GetClient("oci_redis.OciCacheConfigSetClient").(*oci_redis.OciCacheConfigSetClient)
}

func initRedisOciCacheDefaultConfigSetClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_redis.NewOciCacheDefaultConfigSetClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) OciCacheDefaultConfigSetClient() *oci_redis.OciCacheDefaultConfigSetClient {
	return m.GetClient("oci_redis.OciCacheDefaultConfigSetClient").(*oci_redis.OciCacheDefaultConfigSetClient)
}

func initRedisOciCacheUserClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_redis.NewOciCacheUserClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) OciCacheUserClient() *oci_redis.OciCacheUserClient {
	return m.GetClient("oci_redis.OciCacheUserClient").(*oci_redis.OciCacheUserClient)
}

func initRedisRedisClusterClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_redis.NewRedisClusterClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) RedisClusterClient() *oci_redis.RedisClusterClient {
	return m.GetClient("oci_redis.RedisClusterClient").(*oci_redis.RedisClusterClient)
}

func initRedisRedisIdentityClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_redis.NewRedisIdentityClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) RedisIdentityClient() *oci_redis.RedisIdentityClient {
	return m.GetClient("oci_redis.RedisIdentityClient").(*oci_redis.RedisIdentityClient)
}
