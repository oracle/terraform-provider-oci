package network_firewall

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

func GetNetworkFirewallPolicySubResourceCompositeId(resourceName string, networkFirewallPolicyId string, resourceType string) string {
	resourceName = url.PathEscape(resourceName)
	networkFirewallPolicyId = url.PathEscape(networkFirewallPolicyId)
	compositeId := "networkFirewallPolicies/" + networkFirewallPolicyId + "/" + resourceType + "/" + resourceName
	return compositeId
}

func parseNetworkFirewallPolicySubResourceCompositeId(compositeId string, resourceType string) (resourceName string, networkFirewallPolicyId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("networkFirewallPolicies/.*/"+resourceType+"/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	networkFirewallPolicyId, _ = url.PathUnescape(parts[1])
	resourceName, _ = url.PathUnescape(parts[3])

	return
}
