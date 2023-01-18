package service_mesh

import (
	"bytes"
	"fmt"

	oci_service_mesh "github.com/oracle/oci-go-sdk/v65/servicemesh"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

func AccessLoggingConfigurationToMap(obj *oci_service_mesh.AccessLoggingConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	return result
}

func routeRulesHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if destinations, ok := m["destinations"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", destinations))
	}
	if isGrpc, ok := m["is_grpc"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", isGrpc))
	}
	if path, ok := m["path"]; ok && path != "" {
		buf.WriteString(fmt.Sprintf("%v-", path))
	}
	if pathType, ok := m["path_type"]; ok && pathType != "" {
		buf.WriteString(fmt.Sprintf("%v-", pathType))
	}
	if requestTimeoutInMs, ok := m["request_timeout_in_ms"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", requestTimeoutInMs))
	}
	if type_, ok := m["type"]; ok && type_ != "" {
		buf.WriteString(fmt.Sprintf("%v-", type_))
	}
	return utils.GetStringHashcode(buf.String())
}
