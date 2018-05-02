// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	CertificateRequiredOnlyResource = CertificateResourceDependencies + `
resource "oci_load_balancer_certificate" "test_certificate" {
	#Required
	certificate_name = "${var.certificate_certificate_name}"
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
}
`

	CertificateResourceConfig = CertificateResourceDependencies + `
resource "oci_load_balancer_certificate" "test_certificate" {
	#Required
	certificate_name = "${var.certificate_certificate_name}"
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"

	#Optional
	ca_certificate = "${var.certificate_ca_certificate}"
	passphrase = "${var.certificate_passphrase}"
	private_key = "${var.certificate_private_key}"
	public_certificate = "${var.certificate_public_certificate}"
}
`
	CertificatePropertyVariables = `
variable "certificate_ca_certificate" { default = "-----BEGIN CERTIFICATE-----\nMIIBNzCB4gIJAKtwJkxUgNpzMA0GCSqGSIb3DQEBCwUAMCMxITAfBgNVBAoTGElu\ndGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0xNzA0MTIyMTU3NTZaFw0xODA0MTIy\nMTU3NTZaMCMxITAfBgNVBAoTGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDBcMA0G\nCSqGSIb3DQEBAQUAA0sAMEgCQQDlM8lz3BFJA6zBlsF63k9ajPVq3Q1WQoHQ3j35\n08DRKIfwqfV+CxL63W3dZrwL4TrjqorP5CQ36+I6OWALH2zVAgMBAAEwDQYJKoZI\nhvcNAQELBQADQQCEjHVQJoiiVpIIvDWF+4YDRReVuwzrvq2xduWw7CIsDWlYuGZT\nQKVY6tnTy2XpoUk0fqUvMB/M2HGQ1WqZGHs6\n-----END CERTIFICATE-----" }
variable "certificate_certificate_name" { default = "example_certificate_bundle" }
variable "certificate_passphrase" { default = "Mysecretunlockingcode42!1!" }
variable "certificate_private_key" { default = "-----BEGIN RSA PRIVATE KEY-----\nMIIBOgIBAAJBAOUzyXPcEUkDrMGWwXreT1qM9WrdDVZCgdDePfnTwNEoh/Cp9X4L\nEvrdbd1mvAvhOuOqis/kJDfr4jo5YAsfbNUCAwEAAQJAJz8k4bfvJceBT2zXGIj0\noZa9d1z+qaSdwfwsNJkzzRyGkj/j8yv5FV7KNdSfsBbStlcuxUm4i9o5LXhIA+iQ\ngQIhAPzStAN8+Rz3dWKTjRWuCfy+Pwcmyjl3pkMPSiXzgSJlAiEA6BUZWHP0b542\nu8AizBT3b3xKr1AH2nkIx9OHq7F/QbECIHzqqpDypa8/QVuUZegpVrvvT/r7mn1s\nddS6cDtyJgLVAiEA1Z5OFQeuL2sekBRbMyP9WOW7zMBKakLL3TqL/3JCYxECIAkG\nl96uo1MjK/66X5zQXBG7F2DN2CbcYEz0r3c3vvfq\n-----END RSA PRIVATE KEY-----" }
variable "certificate_public_certificate" { default = "-----BEGIN CERTIFICATE-----\nMIIBNzCB4gIJAKtwJkxUgNpzMA0GCSqGSIb3DQEBCwUAMCMxITAfBgNVBAoTGElu\ndGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0xNzA0MTIyMTU3NTZaFw0xODA0MTIy\nMTU3NTZaMCMxITAfBgNVBAoTGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDBcMA0G\nCSqGSIb3DQEBAQUAA0sAMEgCQQDlM8lz3BFJA6zBlsF63k9ajPVq3Q1WQoHQ3j35\n08DRKIfwqfV+CxL63W3dZrwL4TrjqorP5CQ36+I6OWALH2zVAgMBAAEwDQYJKoZI\nhvcNAQELBQADQQCEjHVQJoiiVpIIvDWF+4YDRReVuwzrvq2xduWw7CIsDWlYuGZT\nQKVY6tnTy2XpoUk0fqUvMB/M2HGQ1WqZGHs6\n-----END CERTIFICATE-----" }

`
	CertificateResourceDependencies = LoadBalancerPropertyVariables + LoadBalancerResourceConfig
)

func TestLoadBalancerCertificateResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_load_balancer_certificate.test_certificate"
	datasourceName := "data.oci_load_balancer_certificates.test_certificates"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config + CertificatePropertyVariables + compartmentIdVariableStr + CertificateRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "certificate_name", "example_certificate_bundle"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + CertificateResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + CertificatePropertyVariables + compartmentIdVariableStr + CertificateResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ca_certificate", "-----BEGIN CERTIFICATE-----\nMIIBNzCB4gIJAKtwJkxUgNpzMA0GCSqGSIb3DQEBCwUAMCMxITAfBgNVBAoTGElu\ndGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0xNzA0MTIyMTU3NTZaFw0xODA0MTIy\nMTU3NTZaMCMxITAfBgNVBAoTGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDBcMA0G\nCSqGSIb3DQEBAQUAA0sAMEgCQQDlM8lz3BFJA6zBlsF63k9ajPVq3Q1WQoHQ3j35\n08DRKIfwqfV+CxL63W3dZrwL4TrjqorP5CQ36+I6OWALH2zVAgMBAAEwDQYJKoZI\nhvcNAQELBQADQQCEjHVQJoiiVpIIvDWF+4YDRReVuwzrvq2xduWw7CIsDWlYuGZT\nQKVY6tnTy2XpoUk0fqUvMB/M2HGQ1WqZGHs6\n-----END CERTIFICATE-----"),
					resource.TestCheckResourceAttr(resourceName, "certificate_name", "example_certificate_bundle"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "passphrase", "Mysecretunlockingcode42!1!"),
					resource.TestCheckResourceAttr(resourceName, "private_key", "-----BEGIN RSA PRIVATE KEY-----\nMIIBOgIBAAJBAOUzyXPcEUkDrMGWwXreT1qM9WrdDVZCgdDePfnTwNEoh/Cp9X4L\nEvrdbd1mvAvhOuOqis/kJDfr4jo5YAsfbNUCAwEAAQJAJz8k4bfvJceBT2zXGIj0\noZa9d1z+qaSdwfwsNJkzzRyGkj/j8yv5FV7KNdSfsBbStlcuxUm4i9o5LXhIA+iQ\ngQIhAPzStAN8+Rz3dWKTjRWuCfy+Pwcmyjl3pkMPSiXzgSJlAiEA6BUZWHP0b542\nu8AizBT3b3xKr1AH2nkIx9OHq7F/QbECIHzqqpDypa8/QVuUZegpVrvvT/r7mn1s\nddS6cDtyJgLVAiEA1Z5OFQeuL2sekBRbMyP9WOW7zMBKakLL3TqL/3JCYxECIAkG\nl96uo1MjK/66X5zQXBG7F2DN2CbcYEz0r3c3vvfq\n-----END RSA PRIVATE KEY-----"),
					resource.TestCheckResourceAttr(resourceName, "public_certificate", "-----BEGIN CERTIFICATE-----\nMIIBNzCB4gIJAKtwJkxUgNpzMA0GCSqGSIb3DQEBCwUAMCMxITAfBgNVBAoTGElu\ndGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0xNzA0MTIyMTU3NTZaFw0xODA0MTIy\nMTU3NTZaMCMxITAfBgNVBAoTGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDBcMA0G\nCSqGSIb3DQEBAQUAA0sAMEgCQQDlM8lz3BFJA6zBlsF63k9ajPVq3Q1WQoHQ3j35\n08DRKIfwqfV+CxL63W3dZrwL4TrjqorP5CQ36+I6OWALH2zVAgMBAAEwDQYJKoZI\nhvcNAQELBQADQQCEjHVQJoiiVpIIvDWF+4YDRReVuwzrvq2xduWw7CIsDWlYuGZT\nQKVY6tnTy2XpoUk0fqUvMB/M2HGQ1WqZGHs6\n-----END CERTIFICATE-----"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to Force New parameters.
			{
				Config: config + `
variable "certificate_ca_certificate" { default = "-----BEGIN CERTIFICATE-----\nMIIE6zCCA9OgAwIBAgIJAPfyzvb6nDnyMA0GCSqGSIb3DQEBBQUAMIGpMQswCQYD\nVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHU2VhdHRrZTEP\nMA0GA1UEChMGT3JhY2xlMQwwCgYDVQQLEwNPQ0kxJzAlBgNVBAMTHmNlcnRpZmlj\nYXRlR2VuZXJhdGVkVGVzdFVwZGF0ZTErMCkGCSqGSIb3DQEJARYcb3JjaGVzdHJh\ndGlvbl9ncnBAb3JhY2xlLmNvbTAeFw0xODAyMjYyMjEyNDdaFw0xOTAyMjYyMjEy\nNDdaMIGpMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UE\nBxMHU2VhdHRrZTEPMA0GA1UEChMGT3JhY2xlMQwwCgYDVQQLEwNPQ0kxJzAlBgNV\nBAMTHmNlcnRpZmljYXRlR2VuZXJhdGVkVGVzdFVwZGF0ZTErMCkGCSqGSIb3DQEJ\nARYcb3JjaGVzdHJhdGlvbl9ncnBAb3JhY2xlLmNvbTCCASIwDQYJKoZIhvcNAQEB\nBQADggEPADCCAQoCggEBAMSpOzv8Y2O2O9ChqZkR7lkxtI/NrDWO2y8ogWiqUTb4\nVuXyNflrO31lCZOMKJSxqVUiPJSmT40YhFGoJYC8GXbB9cxZX5JLExWDwasn7Yeb\nEhaPBr85Qo0LunmOa6DsN47xxzpPJI/9+3xemvn4GJhFmraIS81fF261GbMoCCyg\ng2IyeBWN4HJxJjSH0o0CRokJelDj1LD0wLhswKWceJ9MPNXTwVnbrPbMWTF//XwT\nNrIz68TUKYQchAFY0EBNLTdN3GXM1heHEJT5of2DSNeQMqPmLwTaKK02IAIN3K67\nTMMtFsoYmjQu6WWkjb5ctU9SxR5WeDY3HWzJ50p2OxMCAwEAAaOCARIwggEOMB0G\nA1UdDgQWBBS0m/24zvCsED+neDIWwubApJTRMDCB3gYDVR0jBIHWMIHTgBS0m/24\nzvCsED+neDIWwubApJTRMKGBr6SBrDCBqTELMAkGA1UEBhMCVVMxEzARBgNVBAgT\nCldhc2hpbmd0b24xEDAOBgNVBAcTB1NlYXR0a2UxDzANBgNVBAoTBk9yYWNsZTEM\nMAoGA1UECxMDT0NJMScwJQYDVQQDEx5jZXJ0aWZpY2F0ZUdlbmVyYXRlZFRlc3RV\ncGRhdGUxKzApBgkqhkiG9w0BCQEWHG9yY2hlc3RyYXRpb25fZ3JwQG9yYWNsZS5j\nb22CCQD38s72+pw58jAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBBQUAA4IBAQCM\nK2G+XvJG2H6N7e3nr4I3f6FuCEo/6LL3leEiYeWlz+Nqm/z7PmLZOseTsTZRTmv7\ncIvoboQnQnUeV1kShW/smr3EKkiuVDz8nngxHkYneUKlSQqyr5J3JrUb53VGyvqi\nh9MWzYYfFvbtQdiz0VelY520D0YJmsvFCAY+ilY4Y7rG5sdqZhhmD1FIJwdi3gJl\nLyRL9XCB3ZfS1Q9marFsAbNQOTh5mKfkZs0aUo3zU1Dhh+usodp3azu5Nergb95+\nGzT4US7kxywB95qBSH7oZX0cBLv/4y3dfkZtIleHG/luAbpEiWHOL0GPaXtJuc0X\nGUpcphW5CCBfWVB5HrS/\n-----END CERTIFICATE-----" }
variable "certificate_certificate_name" { default = "certificateName2" }
variable "certificate_passphrase" { default = "someP@ssphrase" }
variable "certificate_private_key" { default = "-----BEGIN RSA PRIVATE KEY-----\nMIIEpQIBAAKCAQEAxKk7O/xjY7Y70KGpmRHuWTG0j82sNY7bLyiBaKpRNvhW5fI1\n+Ws7fWUJk4wolLGpVSI8lKZPjRiEUaglgLwZdsH1zFlfkksTFYPBqyfth5sSFo8G\nvzlCjQu6eY5roOw3jvHHOk8kj/37fF6a+fgYmEWatohLzV8XbrUZsygILKCDYjJ4\nFY3gcnEmNIfSjQJGiQl6UOPUsPTAuGzApZx4n0w81dPBWdus9sxZMX/9fBM2sjPr\nxNQphByEAVjQQE0tN03cZczWF4cQlPmh/YNI15Ayo+YvBNoorTYgAg3crrtMwy0W\nyhiaNC7pZaSNvly1T1LFHlZ4NjcdbMnnSnY7EwIDAQABAoIBAQCZttEW1PJ2qKTe\nMM4YV9yeGOzSDeeRwURu+tETlzR9q+R4PTaU6o0IOSYgDshVWPxLD8ScR9YcKdKa\n5LvTgLLCkV5nSGAcP9P92AtTlZTijoG34jla1A3Boy4w/vH/SUMv1qlrWRrx9YpR\nLr0hrxrCQn1oOOZO8wmkvacF0r946OoHcYPcyf7oX2X6L7hPoQBmMCNpQBekPaq1\nt9SlIZBKxyj/Cmjx1/Vh7kce4v5DZ2I3FZeeACH/+nXLHx4QfbiAovbBCJv8bi0R\nkh7tvS4LwrQVjF0jgFOqkZKBoGl2Kk42KCj45YwiIxek1cqD8MmoJImjfK5HME49\n+XFtDYaxAoGBAO4bMpfC3Pjqg1FFr3ppL2n+/3XbZSocjeGr+yu+/wVEUqD8hI94\nIACYmtHoFz4OgwvkRAJMcSnOXiSc8WB2Zer0Ck4NhT5aIKNPBVhEzVrW+ztMWEPt\nS0x05zXJqrG4rgedKLLviLfzbLn/o66azcyBQigXQQmyFS8YuGdNzgq5AoGBANNw\nsQIcKIriZfiq6Odab3mRDEsWaKCsQO9fq5tDTwlqVxv4ZatAUUEPlqs36DrSHF7i\n/OMcQ1B4Tmxl36/hHPPhFaGqWQd0gUBSTBpD8S/UlvIsW7RUVgbzGdIKg+LrHkt/\ngve0ztFjRUMiWWlh+zLS+0UYcZKoveM+DkEo/t4rAoGBAJCMflhXeiK66+Go9nDP\n7nLg1WFNQcsg7plz+NWD6/ncknUdF7SpbnQuML8BsSqBUNklSIxEr+Z7W/fKN1ci\nSZkrch4UYzGJMYsy68G8cxaUsIw9OtBx/GZ8CelDdpbR0QTaSdznQg66fuUj5tCX\nNhzS08DW7Srfq7Cdx0UfnNgxAoGAN7e8jWfKLz8Vy/+NvFmSgqB8ctUG4UBDJFE7\nsYB9NWf2aIJ+mDAsuU5YT+o9ojJ4E3ERXu+1tWDemaYG2hwiOwoXXzC6oUJMRRzq\nvQkmZ4pH9K1HxS9sNAyfwz/OHWAD7bU+V/Qng/n66IQGt8SAI5aBbGXSl+krkNRr\ngTRCZV0CgYEAwQlPqRY88/v9reFHcE5ylGiaV2An4R/83VGic2ELwC7W01ZLPc0u\nGDCTWDM1d4f9otU81g8QTKjes5cvoFIcPs+PiEFKzjTIs/B/rkZF9fmPrZlm07bb\n8O9dnl1VE1Odns9NOHagf+SEmr8okSrYJ30PY/gdowx294/fYpHV94U=\n-----END RSA PRIVATE KEY-----" }
variable "certificate_public_certificate" { default = "-----BEGIN CERTIFICATE-----\nMIIE6zCCA9OgAwIBAgIJAPfyzvb6nDnyMA0GCSqGSIb3DQEBBQUAMIGpMQswCQYD\nVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHU2VhdHRrZTEP\nMA0GA1UEChMGT3JhY2xlMQwwCgYDVQQLEwNPQ0kxJzAlBgNVBAMTHmNlcnRpZmlj\nYXRlR2VuZXJhdGVkVGVzdFVwZGF0ZTErMCkGCSqGSIb3DQEJARYcb3JjaGVzdHJh\ndGlvbl9ncnBAb3JhY2xlLmNvbTAeFw0xODAyMjYyMjEyNDdaFw0xOTAyMjYyMjEy\nNDdaMIGpMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UE\nBxMHU2VhdHRrZTEPMA0GA1UEChMGT3JhY2xlMQwwCgYDVQQLEwNPQ0kxJzAlBgNV\nBAMTHmNlcnRpZmljYXRlR2VuZXJhdGVkVGVzdFVwZGF0ZTErMCkGCSqGSIb3DQEJ\nARYcb3JjaGVzdHJhdGlvbl9ncnBAb3JhY2xlLmNvbTCCASIwDQYJKoZIhvcNAQEB\nBQADggEPADCCAQoCggEBAMSpOzv8Y2O2O9ChqZkR7lkxtI/NrDWO2y8ogWiqUTb4\nVuXyNflrO31lCZOMKJSxqVUiPJSmT40YhFGoJYC8GXbB9cxZX5JLExWDwasn7Yeb\nEhaPBr85Qo0LunmOa6DsN47xxzpPJI/9+3xemvn4GJhFmraIS81fF261GbMoCCyg\ng2IyeBWN4HJxJjSH0o0CRokJelDj1LD0wLhswKWceJ9MPNXTwVnbrPbMWTF//XwT\nNrIz68TUKYQchAFY0EBNLTdN3GXM1heHEJT5of2DSNeQMqPmLwTaKK02IAIN3K67\nTMMtFsoYmjQu6WWkjb5ctU9SxR5WeDY3HWzJ50p2OxMCAwEAAaOCARIwggEOMB0G\nA1UdDgQWBBS0m/24zvCsED+neDIWwubApJTRMDCB3gYDVR0jBIHWMIHTgBS0m/24\nzvCsED+neDIWwubApJTRMKGBr6SBrDCBqTELMAkGA1UEBhMCVVMxEzARBgNVBAgT\nCldhc2hpbmd0b24xEDAOBgNVBAcTB1NlYXR0a2UxDzANBgNVBAoTBk9yYWNsZTEM\nMAoGA1UECxMDT0NJMScwJQYDVQQDEx5jZXJ0aWZpY2F0ZUdlbmVyYXRlZFRlc3RV\ncGRhdGUxKzApBgkqhkiG9w0BCQEWHG9yY2hlc3RyYXRpb25fZ3JwQG9yYWNsZS5j\nb22CCQD38s72+pw58jAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBBQUAA4IBAQCM\nK2G+XvJG2H6N7e3nr4I3f6FuCEo/6LL3leEiYeWlz+Nqm/z7PmLZOseTsTZRTmv7\ncIvoboQnQnUeV1kShW/smr3EKkiuVDz8nngxHkYneUKlSQqyr5J3JrUb53VGyvqi\nh9MWzYYfFvbtQdiz0VelY520D0YJmsvFCAY+ilY4Y7rG5sdqZhhmD1FIJwdi3gJl\nLyRL9XCB3ZfS1Q9marFsAbNQOTh5mKfkZs0aUo3zU1Dhh+usodp3azu5Nergb95+\nGzT4US7kxywB95qBSH7oZX0cBLv/4y3dfkZtIleHG/luAbpEiWHOL0GPaXtJuc0X\nGUpcphW5CCBfWVB5HrS/\n-----END CERTIFICATE-----" }

                ` + compartmentIdVariableStr2 + CertificateResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ca_certificate", "-----BEGIN CERTIFICATE-----\nMIIE6zCCA9OgAwIBAgIJAPfyzvb6nDnyMA0GCSqGSIb3DQEBBQUAMIGpMQswCQYD\nVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHU2VhdHRrZTEP\nMA0GA1UEChMGT3JhY2xlMQwwCgYDVQQLEwNPQ0kxJzAlBgNVBAMTHmNlcnRpZmlj\nYXRlR2VuZXJhdGVkVGVzdFVwZGF0ZTErMCkGCSqGSIb3DQEJARYcb3JjaGVzdHJh\ndGlvbl9ncnBAb3JhY2xlLmNvbTAeFw0xODAyMjYyMjEyNDdaFw0xOTAyMjYyMjEy\nNDdaMIGpMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UE\nBxMHU2VhdHRrZTEPMA0GA1UEChMGT3JhY2xlMQwwCgYDVQQLEwNPQ0kxJzAlBgNV\nBAMTHmNlcnRpZmljYXRlR2VuZXJhdGVkVGVzdFVwZGF0ZTErMCkGCSqGSIb3DQEJ\nARYcb3JjaGVzdHJhdGlvbl9ncnBAb3JhY2xlLmNvbTCCASIwDQYJKoZIhvcNAQEB\nBQADggEPADCCAQoCggEBAMSpOzv8Y2O2O9ChqZkR7lkxtI/NrDWO2y8ogWiqUTb4\nVuXyNflrO31lCZOMKJSxqVUiPJSmT40YhFGoJYC8GXbB9cxZX5JLExWDwasn7Yeb\nEhaPBr85Qo0LunmOa6DsN47xxzpPJI/9+3xemvn4GJhFmraIS81fF261GbMoCCyg\ng2IyeBWN4HJxJjSH0o0CRokJelDj1LD0wLhswKWceJ9MPNXTwVnbrPbMWTF//XwT\nNrIz68TUKYQchAFY0EBNLTdN3GXM1heHEJT5of2DSNeQMqPmLwTaKK02IAIN3K67\nTMMtFsoYmjQu6WWkjb5ctU9SxR5WeDY3HWzJ50p2OxMCAwEAAaOCARIwggEOMB0G\nA1UdDgQWBBS0m/24zvCsED+neDIWwubApJTRMDCB3gYDVR0jBIHWMIHTgBS0m/24\nzvCsED+neDIWwubApJTRMKGBr6SBrDCBqTELMAkGA1UEBhMCVVMxEzARBgNVBAgT\nCldhc2hpbmd0b24xEDAOBgNVBAcTB1NlYXR0a2UxDzANBgNVBAoTBk9yYWNsZTEM\nMAoGA1UECxMDT0NJMScwJQYDVQQDEx5jZXJ0aWZpY2F0ZUdlbmVyYXRlZFRlc3RV\ncGRhdGUxKzApBgkqhkiG9w0BCQEWHG9yY2hlc3RyYXRpb25fZ3JwQG9yYWNsZS5j\nb22CCQD38s72+pw58jAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBBQUAA4IBAQCM\nK2G+XvJG2H6N7e3nr4I3f6FuCEo/6LL3leEiYeWlz+Nqm/z7PmLZOseTsTZRTmv7\ncIvoboQnQnUeV1kShW/smr3EKkiuVDz8nngxHkYneUKlSQqyr5J3JrUb53VGyvqi\nh9MWzYYfFvbtQdiz0VelY520D0YJmsvFCAY+ilY4Y7rG5sdqZhhmD1FIJwdi3gJl\nLyRL9XCB3ZfS1Q9marFsAbNQOTh5mKfkZs0aUo3zU1Dhh+usodp3azu5Nergb95+\nGzT4US7kxywB95qBSH7oZX0cBLv/4y3dfkZtIleHG/luAbpEiWHOL0GPaXtJuc0X\nGUpcphW5CCBfWVB5HrS/\n-----END CERTIFICATE-----"),
					resource.TestCheckResourceAttr(resourceName, "certificate_name", "certificateName2"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "passphrase", "someP@ssphrase"),
					resource.TestCheckResourceAttr(resourceName, "private_key", "-----BEGIN RSA PRIVATE KEY-----\nMIIEpQIBAAKCAQEAxKk7O/xjY7Y70KGpmRHuWTG0j82sNY7bLyiBaKpRNvhW5fI1\n+Ws7fWUJk4wolLGpVSI8lKZPjRiEUaglgLwZdsH1zFlfkksTFYPBqyfth5sSFo8G\nvzlCjQu6eY5roOw3jvHHOk8kj/37fF6a+fgYmEWatohLzV8XbrUZsygILKCDYjJ4\nFY3gcnEmNIfSjQJGiQl6UOPUsPTAuGzApZx4n0w81dPBWdus9sxZMX/9fBM2sjPr\nxNQphByEAVjQQE0tN03cZczWF4cQlPmh/YNI15Ayo+YvBNoorTYgAg3crrtMwy0W\nyhiaNC7pZaSNvly1T1LFHlZ4NjcdbMnnSnY7EwIDAQABAoIBAQCZttEW1PJ2qKTe\nMM4YV9yeGOzSDeeRwURu+tETlzR9q+R4PTaU6o0IOSYgDshVWPxLD8ScR9YcKdKa\n5LvTgLLCkV5nSGAcP9P92AtTlZTijoG34jla1A3Boy4w/vH/SUMv1qlrWRrx9YpR\nLr0hrxrCQn1oOOZO8wmkvacF0r946OoHcYPcyf7oX2X6L7hPoQBmMCNpQBekPaq1\nt9SlIZBKxyj/Cmjx1/Vh7kce4v5DZ2I3FZeeACH/+nXLHx4QfbiAovbBCJv8bi0R\nkh7tvS4LwrQVjF0jgFOqkZKBoGl2Kk42KCj45YwiIxek1cqD8MmoJImjfK5HME49\n+XFtDYaxAoGBAO4bMpfC3Pjqg1FFr3ppL2n+/3XbZSocjeGr+yu+/wVEUqD8hI94\nIACYmtHoFz4OgwvkRAJMcSnOXiSc8WB2Zer0Ck4NhT5aIKNPBVhEzVrW+ztMWEPt\nS0x05zXJqrG4rgedKLLviLfzbLn/o66azcyBQigXQQmyFS8YuGdNzgq5AoGBANNw\nsQIcKIriZfiq6Odab3mRDEsWaKCsQO9fq5tDTwlqVxv4ZatAUUEPlqs36DrSHF7i\n/OMcQ1B4Tmxl36/hHPPhFaGqWQd0gUBSTBpD8S/UlvIsW7RUVgbzGdIKg+LrHkt/\ngve0ztFjRUMiWWlh+zLS+0UYcZKoveM+DkEo/t4rAoGBAJCMflhXeiK66+Go9nDP\n7nLg1WFNQcsg7plz+NWD6/ncknUdF7SpbnQuML8BsSqBUNklSIxEr+Z7W/fKN1ci\nSZkrch4UYzGJMYsy68G8cxaUsIw9OtBx/GZ8CelDdpbR0QTaSdznQg66fuUj5tCX\nNhzS08DW7Srfq7Cdx0UfnNgxAoGAN7e8jWfKLz8Vy/+NvFmSgqB8ctUG4UBDJFE7\nsYB9NWf2aIJ+mDAsuU5YT+o9ojJ4E3ERXu+1tWDemaYG2hwiOwoXXzC6oUJMRRzq\nvQkmZ4pH9K1HxS9sNAyfwz/OHWAD7bU+V/Qng/n66IQGt8SAI5aBbGXSl+krkNRr\ngTRCZV0CgYEAwQlPqRY88/v9reFHcE5ylGiaV2An4R/83VGic2ELwC7W01ZLPc0u\nGDCTWDM1d4f9otU81g8QTKjes5cvoFIcPs+PiEFKzjTIs/B/rkZF9fmPrZlm07bb\n8O9dnl1VE1Odns9NOHagf+SEmr8okSrYJ30PY/gdowx294/fYpHV94U=\n-----END RSA PRIVATE KEY-----"),
					resource.TestCheckResourceAttr(resourceName, "public_certificate", "-----BEGIN CERTIFICATE-----\nMIIE6zCCA9OgAwIBAgIJAPfyzvb6nDnyMA0GCSqGSIb3DQEBBQUAMIGpMQswCQYD\nVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHU2VhdHRrZTEP\nMA0GA1UEChMGT3JhY2xlMQwwCgYDVQQLEwNPQ0kxJzAlBgNVBAMTHmNlcnRpZmlj\nYXRlR2VuZXJhdGVkVGVzdFVwZGF0ZTErMCkGCSqGSIb3DQEJARYcb3JjaGVzdHJh\ndGlvbl9ncnBAb3JhY2xlLmNvbTAeFw0xODAyMjYyMjEyNDdaFw0xOTAyMjYyMjEy\nNDdaMIGpMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UE\nBxMHU2VhdHRrZTEPMA0GA1UEChMGT3JhY2xlMQwwCgYDVQQLEwNPQ0kxJzAlBgNV\nBAMTHmNlcnRpZmljYXRlR2VuZXJhdGVkVGVzdFVwZGF0ZTErMCkGCSqGSIb3DQEJ\nARYcb3JjaGVzdHJhdGlvbl9ncnBAb3JhY2xlLmNvbTCCASIwDQYJKoZIhvcNAQEB\nBQADggEPADCCAQoCggEBAMSpOzv8Y2O2O9ChqZkR7lkxtI/NrDWO2y8ogWiqUTb4\nVuXyNflrO31lCZOMKJSxqVUiPJSmT40YhFGoJYC8GXbB9cxZX5JLExWDwasn7Yeb\nEhaPBr85Qo0LunmOa6DsN47xxzpPJI/9+3xemvn4GJhFmraIS81fF261GbMoCCyg\ng2IyeBWN4HJxJjSH0o0CRokJelDj1LD0wLhswKWceJ9MPNXTwVnbrPbMWTF//XwT\nNrIz68TUKYQchAFY0EBNLTdN3GXM1heHEJT5of2DSNeQMqPmLwTaKK02IAIN3K67\nTMMtFsoYmjQu6WWkjb5ctU9SxR5WeDY3HWzJ50p2OxMCAwEAAaOCARIwggEOMB0G\nA1UdDgQWBBS0m/24zvCsED+neDIWwubApJTRMDCB3gYDVR0jBIHWMIHTgBS0m/24\nzvCsED+neDIWwubApJTRMKGBr6SBrDCBqTELMAkGA1UEBhMCVVMxEzARBgNVBAgT\nCldhc2hpbmd0b24xEDAOBgNVBAcTB1NlYXR0a2UxDzANBgNVBAoTBk9yYWNsZTEM\nMAoGA1UECxMDT0NJMScwJQYDVQQDEx5jZXJ0aWZpY2F0ZUdlbmVyYXRlZFRlc3RV\ncGRhdGUxKzApBgkqhkiG9w0BCQEWHG9yY2hlc3RyYXRpb25fZ3JwQG9yYWNsZS5j\nb22CCQD38s72+pw58jAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBBQUAA4IBAQCM\nK2G+XvJG2H6N7e3nr4I3f6FuCEo/6LL3leEiYeWlz+Nqm/z7PmLZOseTsTZRTmv7\ncIvoboQnQnUeV1kShW/smr3EKkiuVDz8nngxHkYneUKlSQqyr5J3JrUb53VGyvqi\nh9MWzYYfFvbtQdiz0VelY520D0YJmsvFCAY+ilY4Y7rG5sdqZhhmD1FIJwdi3gJl\nLyRL9XCB3ZfS1Q9marFsAbNQOTh5mKfkZs0aUo3zU1Dhh+usodp3azu5Nergb95+\nGzT4US7kxywB95qBSH7oZX0cBLv/4y3dfkZtIleHG/luAbpEiWHOL0GPaXtJuc0X\nGUpcphW5CCBfWVB5HrS/\n-----END CERTIFICATE-----"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated but it wasn't.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config + `
variable "certificate_ca_certificate" { default = "-----BEGIN CERTIFICATE-----\nMIIE6zCCA9OgAwIBAgIJAPfyzvb6nDnyMA0GCSqGSIb3DQEBBQUAMIGpMQswCQYD\nVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHU2VhdHRrZTEP\nMA0GA1UEChMGT3JhY2xlMQwwCgYDVQQLEwNPQ0kxJzAlBgNVBAMTHmNlcnRpZmlj\nYXRlR2VuZXJhdGVkVGVzdFVwZGF0ZTErMCkGCSqGSIb3DQEJARYcb3JjaGVzdHJh\ndGlvbl9ncnBAb3JhY2xlLmNvbTAeFw0xODAyMjYyMjEyNDdaFw0xOTAyMjYyMjEy\nNDdaMIGpMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UE\nBxMHU2VhdHRrZTEPMA0GA1UEChMGT3JhY2xlMQwwCgYDVQQLEwNPQ0kxJzAlBgNV\nBAMTHmNlcnRpZmljYXRlR2VuZXJhdGVkVGVzdFVwZGF0ZTErMCkGCSqGSIb3DQEJ\nARYcb3JjaGVzdHJhdGlvbl9ncnBAb3JhY2xlLmNvbTCCASIwDQYJKoZIhvcNAQEB\nBQADggEPADCCAQoCggEBAMSpOzv8Y2O2O9ChqZkR7lkxtI/NrDWO2y8ogWiqUTb4\nVuXyNflrO31lCZOMKJSxqVUiPJSmT40YhFGoJYC8GXbB9cxZX5JLExWDwasn7Yeb\nEhaPBr85Qo0LunmOa6DsN47xxzpPJI/9+3xemvn4GJhFmraIS81fF261GbMoCCyg\ng2IyeBWN4HJxJjSH0o0CRokJelDj1LD0wLhswKWceJ9MPNXTwVnbrPbMWTF//XwT\nNrIz68TUKYQchAFY0EBNLTdN3GXM1heHEJT5of2DSNeQMqPmLwTaKK02IAIN3K67\nTMMtFsoYmjQu6WWkjb5ctU9SxR5WeDY3HWzJ50p2OxMCAwEAAaOCARIwggEOMB0G\nA1UdDgQWBBS0m/24zvCsED+neDIWwubApJTRMDCB3gYDVR0jBIHWMIHTgBS0m/24\nzvCsED+neDIWwubApJTRMKGBr6SBrDCBqTELMAkGA1UEBhMCVVMxEzARBgNVBAgT\nCldhc2hpbmd0b24xEDAOBgNVBAcTB1NlYXR0a2UxDzANBgNVBAoTBk9yYWNsZTEM\nMAoGA1UECxMDT0NJMScwJQYDVQQDEx5jZXJ0aWZpY2F0ZUdlbmVyYXRlZFRlc3RV\ncGRhdGUxKzApBgkqhkiG9w0BCQEWHG9yY2hlc3RyYXRpb25fZ3JwQG9yYWNsZS5j\nb22CCQD38s72+pw58jAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBBQUAA4IBAQCM\nK2G+XvJG2H6N7e3nr4I3f6FuCEo/6LL3leEiYeWlz+Nqm/z7PmLZOseTsTZRTmv7\ncIvoboQnQnUeV1kShW/smr3EKkiuVDz8nngxHkYneUKlSQqyr5J3JrUb53VGyvqi\nh9MWzYYfFvbtQdiz0VelY520D0YJmsvFCAY+ilY4Y7rG5sdqZhhmD1FIJwdi3gJl\nLyRL9XCB3ZfS1Q9marFsAbNQOTh5mKfkZs0aUo3zU1Dhh+usodp3azu5Nergb95+\nGzT4US7kxywB95qBSH7oZX0cBLv/4y3dfkZtIleHG/luAbpEiWHOL0GPaXtJuc0X\nGUpcphW5CCBfWVB5HrS/\n-----END CERTIFICATE-----" }
variable "certificate_certificate_name" { default = "certificateName2" }
variable "certificate_passphrase" { default = "someP@ssphrase" }
variable "certificate_private_key" { default = "-----BEGIN RSA PRIVATE KEY-----\nMIIEpQIBAAKCAQEAxKk7O/xjY7Y70KGpmRHuWTG0j82sNY7bLyiBaKpRNvhW5fI1\n+Ws7fWUJk4wolLGpVSI8lKZPjRiEUaglgLwZdsH1zFlfkksTFYPBqyfth5sSFo8G\nvzlCjQu6eY5roOw3jvHHOk8kj/37fF6a+fgYmEWatohLzV8XbrUZsygILKCDYjJ4\nFY3gcnEmNIfSjQJGiQl6UOPUsPTAuGzApZx4n0w81dPBWdus9sxZMX/9fBM2sjPr\nxNQphByEAVjQQE0tN03cZczWF4cQlPmh/YNI15Ayo+YvBNoorTYgAg3crrtMwy0W\nyhiaNC7pZaSNvly1T1LFHlZ4NjcdbMnnSnY7EwIDAQABAoIBAQCZttEW1PJ2qKTe\nMM4YV9yeGOzSDeeRwURu+tETlzR9q+R4PTaU6o0IOSYgDshVWPxLD8ScR9YcKdKa\n5LvTgLLCkV5nSGAcP9P92AtTlZTijoG34jla1A3Boy4w/vH/SUMv1qlrWRrx9YpR\nLr0hrxrCQn1oOOZO8wmkvacF0r946OoHcYPcyf7oX2X6L7hPoQBmMCNpQBekPaq1\nt9SlIZBKxyj/Cmjx1/Vh7kce4v5DZ2I3FZeeACH/+nXLHx4QfbiAovbBCJv8bi0R\nkh7tvS4LwrQVjF0jgFOqkZKBoGl2Kk42KCj45YwiIxek1cqD8MmoJImjfK5HME49\n+XFtDYaxAoGBAO4bMpfC3Pjqg1FFr3ppL2n+/3XbZSocjeGr+yu+/wVEUqD8hI94\nIACYmtHoFz4OgwvkRAJMcSnOXiSc8WB2Zer0Ck4NhT5aIKNPBVhEzVrW+ztMWEPt\nS0x05zXJqrG4rgedKLLviLfzbLn/o66azcyBQigXQQmyFS8YuGdNzgq5AoGBANNw\nsQIcKIriZfiq6Odab3mRDEsWaKCsQO9fq5tDTwlqVxv4ZatAUUEPlqs36DrSHF7i\n/OMcQ1B4Tmxl36/hHPPhFaGqWQd0gUBSTBpD8S/UlvIsW7RUVgbzGdIKg+LrHkt/\ngve0ztFjRUMiWWlh+zLS+0UYcZKoveM+DkEo/t4rAoGBAJCMflhXeiK66+Go9nDP\n7nLg1WFNQcsg7plz+NWD6/ncknUdF7SpbnQuML8BsSqBUNklSIxEr+Z7W/fKN1ci\nSZkrch4UYzGJMYsy68G8cxaUsIw9OtBx/GZ8CelDdpbR0QTaSdznQg66fuUj5tCX\nNhzS08DW7Srfq7Cdx0UfnNgxAoGAN7e8jWfKLz8Vy/+NvFmSgqB8ctUG4UBDJFE7\nsYB9NWf2aIJ+mDAsuU5YT+o9ojJ4E3ERXu+1tWDemaYG2hwiOwoXXzC6oUJMRRzq\nvQkmZ4pH9K1HxS9sNAyfwz/OHWAD7bU+V/Qng/n66IQGt8SAI5aBbGXSl+krkNRr\ngTRCZV0CgYEAwQlPqRY88/v9reFHcE5ylGiaV2An4R/83VGic2ELwC7W01ZLPc0u\nGDCTWDM1d4f9otU81g8QTKjes5cvoFIcPs+PiEFKzjTIs/B/rkZF9fmPrZlm07bb\n8O9dnl1VE1Odns9NOHagf+SEmr8okSrYJ30PY/gdowx294/fYpHV94U=\n-----END RSA PRIVATE KEY-----" }
variable "certificate_public_certificate" { default = "-----BEGIN CERTIFICATE-----\nMIIE6zCCA9OgAwIBAgIJAPfyzvb6nDnyMA0GCSqGSIb3DQEBBQUAMIGpMQswCQYD\nVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHU2VhdHRrZTEP\nMA0GA1UEChMGT3JhY2xlMQwwCgYDVQQLEwNPQ0kxJzAlBgNVBAMTHmNlcnRpZmlj\nYXRlR2VuZXJhdGVkVGVzdFVwZGF0ZTErMCkGCSqGSIb3DQEJARYcb3JjaGVzdHJh\ndGlvbl9ncnBAb3JhY2xlLmNvbTAeFw0xODAyMjYyMjEyNDdaFw0xOTAyMjYyMjEy\nNDdaMIGpMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UE\nBxMHU2VhdHRrZTEPMA0GA1UEChMGT3JhY2xlMQwwCgYDVQQLEwNPQ0kxJzAlBgNV\nBAMTHmNlcnRpZmljYXRlR2VuZXJhdGVkVGVzdFVwZGF0ZTErMCkGCSqGSIb3DQEJ\nARYcb3JjaGVzdHJhdGlvbl9ncnBAb3JhY2xlLmNvbTCCASIwDQYJKoZIhvcNAQEB\nBQADggEPADCCAQoCggEBAMSpOzv8Y2O2O9ChqZkR7lkxtI/NrDWO2y8ogWiqUTb4\nVuXyNflrO31lCZOMKJSxqVUiPJSmT40YhFGoJYC8GXbB9cxZX5JLExWDwasn7Yeb\nEhaPBr85Qo0LunmOa6DsN47xxzpPJI/9+3xemvn4GJhFmraIS81fF261GbMoCCyg\ng2IyeBWN4HJxJjSH0o0CRokJelDj1LD0wLhswKWceJ9MPNXTwVnbrPbMWTF//XwT\nNrIz68TUKYQchAFY0EBNLTdN3GXM1heHEJT5of2DSNeQMqPmLwTaKK02IAIN3K67\nTMMtFsoYmjQu6WWkjb5ctU9SxR5WeDY3HWzJ50p2OxMCAwEAAaOCARIwggEOMB0G\nA1UdDgQWBBS0m/24zvCsED+neDIWwubApJTRMDCB3gYDVR0jBIHWMIHTgBS0m/24\nzvCsED+neDIWwubApJTRMKGBr6SBrDCBqTELMAkGA1UEBhMCVVMxEzARBgNVBAgT\nCldhc2hpbmd0b24xEDAOBgNVBAcTB1NlYXR0a2UxDzANBgNVBAoTBk9yYWNsZTEM\nMAoGA1UECxMDT0NJMScwJQYDVQQDEx5jZXJ0aWZpY2F0ZUdlbmVyYXRlZFRlc3RV\ncGRhdGUxKzApBgkqhkiG9w0BCQEWHG9yY2hlc3RyYXRpb25fZ3JwQG9yYWNsZS5j\nb22CCQD38s72+pw58jAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBBQUAA4IBAQCM\nK2G+XvJG2H6N7e3nr4I3f6FuCEo/6LL3leEiYeWlz+Nqm/z7PmLZOseTsTZRTmv7\ncIvoboQnQnUeV1kShW/smr3EKkiuVDz8nngxHkYneUKlSQqyr5J3JrUb53VGyvqi\nh9MWzYYfFvbtQdiz0VelY520D0YJmsvFCAY+ilY4Y7rG5sdqZhhmD1FIJwdi3gJl\nLyRL9XCB3ZfS1Q9marFsAbNQOTh5mKfkZs0aUo3zU1Dhh+usodp3azu5Nergb95+\nGzT4US7kxywB95qBSH7oZX0cBLv/4y3dfkZtIleHG/luAbpEiWHOL0GPaXtJuc0X\nGUpcphW5CCBfWVB5HrS/\n-----END CERTIFICATE-----" }

data "oci_load_balancer_certificates" "test_certificates" {
	#Required
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"

    filter {
    	name = "certificate_name"
    	values = ["${oci_load_balancer_certificate.test_certificate.certificate_name}"]
    }
}
                ` + compartmentIdVariableStr2 + CertificateResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "load_balancer_id"),

					resource.TestCheckResourceAttr(datasourceName, "certificates.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "certificates.0.ca_certificate", "-----BEGIN CERTIFICATE-----\nMIIE6zCCA9OgAwIBAgIJAPfyzvb6nDnyMA0GCSqGSIb3DQEBBQUAMIGpMQswCQYD\nVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHU2VhdHRrZTEP\nMA0GA1UEChMGT3JhY2xlMQwwCgYDVQQLEwNPQ0kxJzAlBgNVBAMTHmNlcnRpZmlj\nYXRlR2VuZXJhdGVkVGVzdFVwZGF0ZTErMCkGCSqGSIb3DQEJARYcb3JjaGVzdHJh\ndGlvbl9ncnBAb3JhY2xlLmNvbTAeFw0xODAyMjYyMjEyNDdaFw0xOTAyMjYyMjEy\nNDdaMIGpMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UE\nBxMHU2VhdHRrZTEPMA0GA1UEChMGT3JhY2xlMQwwCgYDVQQLEwNPQ0kxJzAlBgNV\nBAMTHmNlcnRpZmljYXRlR2VuZXJhdGVkVGVzdFVwZGF0ZTErMCkGCSqGSIb3DQEJ\nARYcb3JjaGVzdHJhdGlvbl9ncnBAb3JhY2xlLmNvbTCCASIwDQYJKoZIhvcNAQEB\nBQADggEPADCCAQoCggEBAMSpOzv8Y2O2O9ChqZkR7lkxtI/NrDWO2y8ogWiqUTb4\nVuXyNflrO31lCZOMKJSxqVUiPJSmT40YhFGoJYC8GXbB9cxZX5JLExWDwasn7Yeb\nEhaPBr85Qo0LunmOa6DsN47xxzpPJI/9+3xemvn4GJhFmraIS81fF261GbMoCCyg\ng2IyeBWN4HJxJjSH0o0CRokJelDj1LD0wLhswKWceJ9MPNXTwVnbrPbMWTF//XwT\nNrIz68TUKYQchAFY0EBNLTdN3GXM1heHEJT5of2DSNeQMqPmLwTaKK02IAIN3K67\nTMMtFsoYmjQu6WWkjb5ctU9SxR5WeDY3HWzJ50p2OxMCAwEAAaOCARIwggEOMB0G\nA1UdDgQWBBS0m/24zvCsED+neDIWwubApJTRMDCB3gYDVR0jBIHWMIHTgBS0m/24\nzvCsED+neDIWwubApJTRMKGBr6SBrDCBqTELMAkGA1UEBhMCVVMxEzARBgNVBAgT\nCldhc2hpbmd0b24xEDAOBgNVBAcTB1NlYXR0a2UxDzANBgNVBAoTBk9yYWNsZTEM\nMAoGA1UECxMDT0NJMScwJQYDVQQDEx5jZXJ0aWZpY2F0ZUdlbmVyYXRlZFRlc3RV\ncGRhdGUxKzApBgkqhkiG9w0BCQEWHG9yY2hlc3RyYXRpb25fZ3JwQG9yYWNsZS5j\nb22CCQD38s72+pw58jAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBBQUAA4IBAQCM\nK2G+XvJG2H6N7e3nr4I3f6FuCEo/6LL3leEiYeWlz+Nqm/z7PmLZOseTsTZRTmv7\ncIvoboQnQnUeV1kShW/smr3EKkiuVDz8nngxHkYneUKlSQqyr5J3JrUb53VGyvqi\nh9MWzYYfFvbtQdiz0VelY520D0YJmsvFCAY+ilY4Y7rG5sdqZhhmD1FIJwdi3gJl\nLyRL9XCB3ZfS1Q9marFsAbNQOTh5mKfkZs0aUo3zU1Dhh+usodp3azu5Nergb95+\nGzT4US7kxywB95qBSH7oZX0cBLv/4y3dfkZtIleHG/luAbpEiWHOL0GPaXtJuc0X\nGUpcphW5CCBfWVB5HrS/\n-----END CERTIFICATE-----"),
					resource.TestCheckResourceAttr(datasourceName, "certificates.0.certificate_name", "certificateName2"),
					//resource.TestCheckResourceAttrSet(datasourceName, "certificates.0.load_balancer_id"), @ CODEGEN this field is not returned by the service
					resource.TestCheckResourceAttr(datasourceName, "certificates.0.public_certificate", "-----BEGIN CERTIFICATE-----\nMIIE6zCCA9OgAwIBAgIJAPfyzvb6nDnyMA0GCSqGSIb3DQEBBQUAMIGpMQswCQYD\nVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHU2VhdHRrZTEP\nMA0GA1UEChMGT3JhY2xlMQwwCgYDVQQLEwNPQ0kxJzAlBgNVBAMTHmNlcnRpZmlj\nYXRlR2VuZXJhdGVkVGVzdFVwZGF0ZTErMCkGCSqGSIb3DQEJARYcb3JjaGVzdHJh\ndGlvbl9ncnBAb3JhY2xlLmNvbTAeFw0xODAyMjYyMjEyNDdaFw0xOTAyMjYyMjEy\nNDdaMIGpMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UE\nBxMHU2VhdHRrZTEPMA0GA1UEChMGT3JhY2xlMQwwCgYDVQQLEwNPQ0kxJzAlBgNV\nBAMTHmNlcnRpZmljYXRlR2VuZXJhdGVkVGVzdFVwZGF0ZTErMCkGCSqGSIb3DQEJ\nARYcb3JjaGVzdHJhdGlvbl9ncnBAb3JhY2xlLmNvbTCCASIwDQYJKoZIhvcNAQEB\nBQADggEPADCCAQoCggEBAMSpOzv8Y2O2O9ChqZkR7lkxtI/NrDWO2y8ogWiqUTb4\nVuXyNflrO31lCZOMKJSxqVUiPJSmT40YhFGoJYC8GXbB9cxZX5JLExWDwasn7Yeb\nEhaPBr85Qo0LunmOa6DsN47xxzpPJI/9+3xemvn4GJhFmraIS81fF261GbMoCCyg\ng2IyeBWN4HJxJjSH0o0CRokJelDj1LD0wLhswKWceJ9MPNXTwVnbrPbMWTF//XwT\nNrIz68TUKYQchAFY0EBNLTdN3GXM1heHEJT5of2DSNeQMqPmLwTaKK02IAIN3K67\nTMMtFsoYmjQu6WWkjb5ctU9SxR5WeDY3HWzJ50p2OxMCAwEAAaOCARIwggEOMB0G\nA1UdDgQWBBS0m/24zvCsED+neDIWwubApJTRMDCB3gYDVR0jBIHWMIHTgBS0m/24\nzvCsED+neDIWwubApJTRMKGBr6SBrDCBqTELMAkGA1UEBhMCVVMxEzARBgNVBAgT\nCldhc2hpbmd0b24xEDAOBgNVBAcTB1NlYXR0a2UxDzANBgNVBAoTBk9yYWNsZTEM\nMAoGA1UECxMDT0NJMScwJQYDVQQDEx5jZXJ0aWZpY2F0ZUdlbmVyYXRlZFRlc3RV\ncGRhdGUxKzApBgkqhkiG9w0BCQEWHG9yY2hlc3RyYXRpb25fZ3JwQG9yYWNsZS5j\nb22CCQD38s72+pw58jAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBBQUAA4IBAQCM\nK2G+XvJG2H6N7e3nr4I3f6FuCEo/6LL3leEiYeWlz+Nqm/z7PmLZOseTsTZRTmv7\ncIvoboQnQnUeV1kShW/smr3EKkiuVDz8nngxHkYneUKlSQqyr5J3JrUb53VGyvqi\nh9MWzYYfFvbtQdiz0VelY520D0YJmsvFCAY+ilY4Y7rG5sdqZhhmD1FIJwdi3gJl\nLyRL9XCB3ZfS1Q9marFsAbNQOTh5mKfkZs0aUo3zU1Dhh+usodp3azu5Nergb95+\nGzT4US7kxywB95qBSH7oZX0cBLv/4y3dfkZtIleHG/luAbpEiWHOL0GPaXtJuc0X\nGUpcphW5CCBfWVB5HrS/\n-----END CERTIFICATE-----"),
				),
			},
		},
	})
}

func TestLoadBalancerCertificateResource_forcenew(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_load_balancer_certificate.test_certificate"

	var lbId string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + CertificatePropertyVariables + compartmentIdVariableStr + CertificateResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ca_certificate", "-----BEGIN CERTIFICATE-----\nMIIBNzCB4gIJAKtwJkxUgNpzMA0GCSqGSIb3DQEBCwUAMCMxITAfBgNVBAoTGElu\ndGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0xNzA0MTIyMTU3NTZaFw0xODA0MTIy\nMTU3NTZaMCMxITAfBgNVBAoTGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDBcMA0G\nCSqGSIb3DQEBAQUAA0sAMEgCQQDlM8lz3BFJA6zBlsF63k9ajPVq3Q1WQoHQ3j35\n08DRKIfwqfV+CxL63W3dZrwL4TrjqorP5CQ36+I6OWALH2zVAgMBAAEwDQYJKoZI\nhvcNAQELBQADQQCEjHVQJoiiVpIIvDWF+4YDRReVuwzrvq2xduWw7CIsDWlYuGZT\nQKVY6tnTy2XpoUk0fqUvMB/M2HGQ1WqZGHs6\n-----END CERTIFICATE-----"),
					resource.TestCheckResourceAttr(resourceName, "certificate_name", "example_certificate_bundle"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "passphrase", "Mysecretunlockingcode42!1!"),
					resource.TestCheckResourceAttr(resourceName, "private_key", "-----BEGIN RSA PRIVATE KEY-----\nMIIBOgIBAAJBAOUzyXPcEUkDrMGWwXreT1qM9WrdDVZCgdDePfnTwNEoh/Cp9X4L\nEvrdbd1mvAvhOuOqis/kJDfr4jo5YAsfbNUCAwEAAQJAJz8k4bfvJceBT2zXGIj0\noZa9d1z+qaSdwfwsNJkzzRyGkj/j8yv5FV7KNdSfsBbStlcuxUm4i9o5LXhIA+iQ\ngQIhAPzStAN8+Rz3dWKTjRWuCfy+Pwcmyjl3pkMPSiXzgSJlAiEA6BUZWHP0b542\nu8AizBT3b3xKr1AH2nkIx9OHq7F/QbECIHzqqpDypa8/QVuUZegpVrvvT/r7mn1s\nddS6cDtyJgLVAiEA1Z5OFQeuL2sekBRbMyP9WOW7zMBKakLL3TqL/3JCYxECIAkG\nl96uo1MjK/66X5zQXBG7F2DN2CbcYEz0r3c3vvfq\n-----END RSA PRIVATE KEY-----"),
					resource.TestCheckResourceAttr(resourceName, "public_certificate", "-----BEGIN CERTIFICATE-----\nMIIBNzCB4gIJAKtwJkxUgNpzMA0GCSqGSIb3DQEBCwUAMCMxITAfBgNVBAoTGElu\ndGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0xNzA0MTIyMTU3NTZaFw0xODA0MTIy\nMTU3NTZaMCMxITAfBgNVBAoTGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDBcMA0G\nCSqGSIb3DQEBAQUAA0sAMEgCQQDlM8lz3BFJA6zBlsF63k9ajPVq3Q1WQoHQ3j35\n08DRKIfwqfV+CxL63W3dZrwL4TrjqorP5CQ36+I6OWALH2zVAgMBAAEwDQYJKoZI\nhvcNAQELBQADQQCEjHVQJoiiVpIIvDWF+4YDRReVuwzrvq2xduWw7CIsDWlYuGZT\nQKVY6tnTy2XpoUk0fqUvMB/M2HGQ1WqZGHs6\n-----END CERTIFICATE-----"),

					func(s *terraform.State) (err error) {
						lbId, err = fromInstanceState(s, resourceName, "load_balancer_id")
						return err
					},
				),
			},
			// force new tests, test that changing a parameter would result in creation of a new resource.

			// @CODEGEN passphrase, private_key, ca_certificate, and public certificate all need to change at the same time for it to be valid.
			// Test for certificate name change
			{
				Config: config + `
variable "certificate_ca_certificate" { default = "-----BEGIN CERTIFICATE-----\nMIIBNzCB4gIJAKtwJkxUgNpzMA0GCSqGSIb3DQEBCwUAMCMxITAfBgNVBAoTGElu\ndGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0xNzA0MTIyMTU3NTZaFw0xODA0MTIy\nMTU3NTZaMCMxITAfBgNVBAoTGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDBcMA0G\nCSqGSIb3DQEBAQUAA0sAMEgCQQDlM8lz3BFJA6zBlsF63k9ajPVq3Q1WQoHQ3j35\n08DRKIfwqfV+CxL63W3dZrwL4TrjqorP5CQ36+I6OWALH2zVAgMBAAEwDQYJKoZI\nhvcNAQELBQADQQCEjHVQJoiiVpIIvDWF+4YDRReVuwzrvq2xduWw7CIsDWlYuGZT\nQKVY6tnTy2XpoUk0fqUvMB/M2HGQ1WqZGHs6\n-----END CERTIFICATE-----" }
variable "certificate_certificate_name" { default = "certificateName2" }
variable "certificate_passphrase" { default = "Mysecretunlockingcode42!1!" }
variable "certificate_private_key" { default = "-----BEGIN RSA PRIVATE KEY-----\nMIIBOgIBAAJBAOUzyXPcEUkDrMGWwXreT1qM9WrdDVZCgdDePfnTwNEoh/Cp9X4L\nEvrdbd1mvAvhOuOqis/kJDfr4jo5YAsfbNUCAwEAAQJAJz8k4bfvJceBT2zXGIj0\noZa9d1z+qaSdwfwsNJkzzRyGkj/j8yv5FV7KNdSfsBbStlcuxUm4i9o5LXhIA+iQ\ngQIhAPzStAN8+Rz3dWKTjRWuCfy+Pwcmyjl3pkMPSiXzgSJlAiEA6BUZWHP0b542\nu8AizBT3b3xKr1AH2nkIx9OHq7F/QbECIHzqqpDypa8/QVuUZegpVrvvT/r7mn1s\nddS6cDtyJgLVAiEA1Z5OFQeuL2sekBRbMyP9WOW7zMBKakLL3TqL/3JCYxECIAkG\nl96uo1MjK/66X5zQXBG7F2DN2CbcYEz0r3c3vvfq\n-----END RSA PRIVATE KEY-----" }
variable "certificate_public_certificate" { default = "-----BEGIN CERTIFICATE-----\nMIIBNzCB4gIJAKtwJkxUgNpzMA0GCSqGSIb3DQEBCwUAMCMxITAfBgNVBAoTGElu\ndGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0xNzA0MTIyMTU3NTZaFw0xODA0MTIy\nMTU3NTZaMCMxITAfBgNVBAoTGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDBcMA0G\nCSqGSIb3DQEBAQUAA0sAMEgCQQDlM8lz3BFJA6zBlsF63k9ajPVq3Q1WQoHQ3j35\n08DRKIfwqfV+CxL63W3dZrwL4TrjqorP5CQ36+I6OWALH2zVAgMBAAEwDQYJKoZI\nhvcNAQELBQADQQCEjHVQJoiiVpIIvDWF+4YDRReVuwzrvq2xduWw7CIsDWlYuGZT\nQKVY6tnTy2XpoUk0fqUvMB/M2HGQ1WqZGHs6\n-----END CERTIFICATE-----" }
				` + compartmentIdVariableStr + CertificateResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ca_certificate", "-----BEGIN CERTIFICATE-----\nMIIBNzCB4gIJAKtwJkxUgNpzMA0GCSqGSIb3DQEBCwUAMCMxITAfBgNVBAoTGElu\ndGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0xNzA0MTIyMTU3NTZaFw0xODA0MTIy\nMTU3NTZaMCMxITAfBgNVBAoTGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDBcMA0G\nCSqGSIb3DQEBAQUAA0sAMEgCQQDlM8lz3BFJA6zBlsF63k9ajPVq3Q1WQoHQ3j35\n08DRKIfwqfV+CxL63W3dZrwL4TrjqorP5CQ36+I6OWALH2zVAgMBAAEwDQYJKoZI\nhvcNAQELBQADQQCEjHVQJoiiVpIIvDWF+4YDRReVuwzrvq2xduWw7CIsDWlYuGZT\nQKVY6tnTy2XpoUk0fqUvMB/M2HGQ1WqZGHs6\n-----END CERTIFICATE-----"),
					resource.TestCheckResourceAttr(resourceName, "certificate_name", "certificateName2"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "passphrase", "Mysecretunlockingcode42!1!"),
					resource.TestCheckResourceAttr(resourceName, "private_key", "-----BEGIN RSA PRIVATE KEY-----\nMIIBOgIBAAJBAOUzyXPcEUkDrMGWwXreT1qM9WrdDVZCgdDePfnTwNEoh/Cp9X4L\nEvrdbd1mvAvhOuOqis/kJDfr4jo5YAsfbNUCAwEAAQJAJz8k4bfvJceBT2zXGIj0\noZa9d1z+qaSdwfwsNJkzzRyGkj/j8yv5FV7KNdSfsBbStlcuxUm4i9o5LXhIA+iQ\ngQIhAPzStAN8+Rz3dWKTjRWuCfy+Pwcmyjl3pkMPSiXzgSJlAiEA6BUZWHP0b542\nu8AizBT3b3xKr1AH2nkIx9OHq7F/QbECIHzqqpDypa8/QVuUZegpVrvvT/r7mn1s\nddS6cDtyJgLVAiEA1Z5OFQeuL2sekBRbMyP9WOW7zMBKakLL3TqL/3JCYxECIAkG\nl96uo1MjK/66X5zQXBG7F2DN2CbcYEz0r3c3vvfq\n-----END RSA PRIVATE KEY-----"),
					resource.TestCheckResourceAttr(resourceName, "public_certificate", "-----BEGIN CERTIFICATE-----\nMIIBNzCB4gIJAKtwJkxUgNpzMA0GCSqGSIb3DQEBCwUAMCMxITAfBgNVBAoTGElu\ndGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0xNzA0MTIyMTU3NTZaFw0xODA0MTIy\nMTU3NTZaMCMxITAfBgNVBAoTGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDBcMA0G\nCSqGSIb3DQEBAQUAA0sAMEgCQQDlM8lz3BFJA6zBlsF63k9ajPVq3Q1WQoHQ3j35\n08DRKIfwqfV+CxL63W3dZrwL4TrjqorP5CQ36+I6OWALH2zVAgMBAAEwDQYJKoZI\nhvcNAQELBQADQQCEjHVQJoiiVpIIvDWF+4YDRReVuwzrvq2xduWw7CIsDWlYuGZT\nQKVY6tnTy2XpoUk0fqUvMB/M2HGQ1WqZGHs6\n-----END CERTIFICATE-----"),
				),
			},

			//test loadBalancerId change via compartmentId Change
			{
				Config: config + `
variable "certificate_ca_certificate" { default = "-----BEGIN CERTIFICATE-----\nMIIBNzCB4gIJAKtwJkxUgNpzMA0GCSqGSIb3DQEBCwUAMCMxITAfBgNVBAoTGElu\ndGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0xNzA0MTIyMTU3NTZaFw0xODA0MTIy\nMTU3NTZaMCMxITAfBgNVBAoTGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDBcMA0G\nCSqGSIb3DQEBAQUAA0sAMEgCQQDlM8lz3BFJA6zBlsF63k9ajPVq3Q1WQoHQ3j35\n08DRKIfwqfV+CxL63W3dZrwL4TrjqorP5CQ36+I6OWALH2zVAgMBAAEwDQYJKoZI\nhvcNAQELBQADQQCEjHVQJoiiVpIIvDWF+4YDRReVuwzrvq2xduWw7CIsDWlYuGZT\nQKVY6tnTy2XpoUk0fqUvMB/M2HGQ1WqZGHs6\n-----END CERTIFICATE-----" }
variable "certificate_certificate_name" { default = "certificateName2" }
variable "certificate_passphrase" { default = "Mysecretunlockingcode42!1!" }
variable "certificate_private_key" { default = "-----BEGIN RSA PRIVATE KEY-----\nMIIBOgIBAAJBAOUzyXPcEUkDrMGWwXreT1qM9WrdDVZCgdDePfnTwNEoh/Cp9X4L\nEvrdbd1mvAvhOuOqis/kJDfr4jo5YAsfbNUCAwEAAQJAJz8k4bfvJceBT2zXGIj0\noZa9d1z+qaSdwfwsNJkzzRyGkj/j8yv5FV7KNdSfsBbStlcuxUm4i9o5LXhIA+iQ\ngQIhAPzStAN8+Rz3dWKTjRWuCfy+Pwcmyjl3pkMPSiXzgSJlAiEA6BUZWHP0b542\nu8AizBT3b3xKr1AH2nkIx9OHq7F/QbECIHzqqpDypa8/QVuUZegpVrvvT/r7mn1s\nddS6cDtyJgLVAiEA1Z5OFQeuL2sekBRbMyP9WOW7zMBKakLL3TqL/3JCYxECIAkG\nl96uo1MjK/66X5zQXBG7F2DN2CbcYEz0r3c3vvfq\n-----END RSA PRIVATE KEY-----" }
variable "certificate_public_certificate" { default = "-----BEGIN CERTIFICATE-----\nMIIBNzCB4gIJAKtwJkxUgNpzMA0GCSqGSIb3DQEBCwUAMCMxITAfBgNVBAoTGElu\ndGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0xNzA0MTIyMTU3NTZaFw0xODA0MTIy\nMTU3NTZaMCMxITAfBgNVBAoTGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDBcMA0G\nCSqGSIb3DQEBAQUAA0sAMEgCQQDlM8lz3BFJA6zBlsF63k9ajPVq3Q1WQoHQ3j35\n08DRKIfwqfV+CxL63W3dZrwL4TrjqorP5CQ36+I6OWALH2zVAgMBAAEwDQYJKoZI\nhvcNAQELBQADQQCEjHVQJoiiVpIIvDWF+4YDRReVuwzrvq2xduWw7CIsDWlYuGZT\nQKVY6tnTy2XpoUk0fqUvMB/M2HGQ1WqZGHs6\n-----END CERTIFICATE-----" }
				` + compartmentIdVariableStr2 + CertificateResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ca_certificate", "-----BEGIN CERTIFICATE-----\nMIIBNzCB4gIJAKtwJkxUgNpzMA0GCSqGSIb3DQEBCwUAMCMxITAfBgNVBAoTGElu\ndGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0xNzA0MTIyMTU3NTZaFw0xODA0MTIy\nMTU3NTZaMCMxITAfBgNVBAoTGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDBcMA0G\nCSqGSIb3DQEBAQUAA0sAMEgCQQDlM8lz3BFJA6zBlsF63k9ajPVq3Q1WQoHQ3j35\n08DRKIfwqfV+CxL63W3dZrwL4TrjqorP5CQ36+I6OWALH2zVAgMBAAEwDQYJKoZI\nhvcNAQELBQADQQCEjHVQJoiiVpIIvDWF+4YDRReVuwzrvq2xduWw7CIsDWlYuGZT\nQKVY6tnTy2XpoUk0fqUvMB/M2HGQ1WqZGHs6\n-----END CERTIFICATE-----"),
					resource.TestCheckResourceAttr(resourceName, "certificate_name", "certificateName2"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "passphrase", "Mysecretunlockingcode42!1!"),
					resource.TestCheckResourceAttr(resourceName, "private_key", "-----BEGIN RSA PRIVATE KEY-----\nMIIBOgIBAAJBAOUzyXPcEUkDrMGWwXreT1qM9WrdDVZCgdDePfnTwNEoh/Cp9X4L\nEvrdbd1mvAvhOuOqis/kJDfr4jo5YAsfbNUCAwEAAQJAJz8k4bfvJceBT2zXGIj0\noZa9d1z+qaSdwfwsNJkzzRyGkj/j8yv5FV7KNdSfsBbStlcuxUm4i9o5LXhIA+iQ\ngQIhAPzStAN8+Rz3dWKTjRWuCfy+Pwcmyjl3pkMPSiXzgSJlAiEA6BUZWHP0b542\nu8AizBT3b3xKr1AH2nkIx9OHq7F/QbECIHzqqpDypa8/QVuUZegpVrvvT/r7mn1s\nddS6cDtyJgLVAiEA1Z5OFQeuL2sekBRbMyP9WOW7zMBKakLL3TqL/3JCYxECIAkG\nl96uo1MjK/66X5zQXBG7F2DN2CbcYEz0r3c3vvfq\n-----END RSA PRIVATE KEY-----"),
					resource.TestCheckResourceAttr(resourceName, "public_certificate", "-----BEGIN CERTIFICATE-----\nMIIBNzCB4gIJAKtwJkxUgNpzMA0GCSqGSIb3DQEBCwUAMCMxITAfBgNVBAoTGElu\ndGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0xNzA0MTIyMTU3NTZaFw0xODA0MTIy\nMTU3NTZaMCMxITAfBgNVBAoTGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDBcMA0G\nCSqGSIb3DQEBAQUAA0sAMEgCQQDlM8lz3BFJA6zBlsF63k9ajPVq3Q1WQoHQ3j35\n08DRKIfwqfV+CxL63W3dZrwL4TrjqorP5CQ36+I6OWALH2zVAgMBAAEwDQYJKoZI\nhvcNAQELBQADQQCEjHVQJoiiVpIIvDWF+4YDRReVuwzrvq2xduWw7CIsDWlYuGZT\nQKVY6tnTy2XpoUk0fqUvMB/M2HGQ1WqZGHs6\n-----END CERTIFICATE-----"),
					func(s *terraform.State) (err error) {
						lbId2, err := fromInstanceState(s, resourceName, "load_balancer_id")
						if lbId == lbId2 {
							return fmt.Errorf("Resource was expected to be recreated with the new loadBalancerId but it wasn't.")
						}
						return err
					},
				),
			},

			//test for certificate info change
			{
				Config: config + `
variable "certificate_ca_certificate" { default = "-----BEGIN CERTIFICATE-----\nMIIE6zCCA9OgAwIBAgIJAPfyzvb6nDnyMA0GCSqGSIb3DQEBBQUAMIGpMQswCQYD\nVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHU2VhdHRrZTEP\nMA0GA1UEChMGT3JhY2xlMQwwCgYDVQQLEwNPQ0kxJzAlBgNVBAMTHmNlcnRpZmlj\nYXRlR2VuZXJhdGVkVGVzdFVwZGF0ZTErMCkGCSqGSIb3DQEJARYcb3JjaGVzdHJh\ndGlvbl9ncnBAb3JhY2xlLmNvbTAeFw0xODAyMjYyMjEyNDdaFw0xOTAyMjYyMjEy\nNDdaMIGpMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UE\nBxMHU2VhdHRrZTEPMA0GA1UEChMGT3JhY2xlMQwwCgYDVQQLEwNPQ0kxJzAlBgNV\nBAMTHmNlcnRpZmljYXRlR2VuZXJhdGVkVGVzdFVwZGF0ZTErMCkGCSqGSIb3DQEJ\nARYcb3JjaGVzdHJhdGlvbl9ncnBAb3JhY2xlLmNvbTCCASIwDQYJKoZIhvcNAQEB\nBQADggEPADCCAQoCggEBAMSpOzv8Y2O2O9ChqZkR7lkxtI/NrDWO2y8ogWiqUTb4\nVuXyNflrO31lCZOMKJSxqVUiPJSmT40YhFGoJYC8GXbB9cxZX5JLExWDwasn7Yeb\nEhaPBr85Qo0LunmOa6DsN47xxzpPJI/9+3xemvn4GJhFmraIS81fF261GbMoCCyg\ng2IyeBWN4HJxJjSH0o0CRokJelDj1LD0wLhswKWceJ9MPNXTwVnbrPbMWTF//XwT\nNrIz68TUKYQchAFY0EBNLTdN3GXM1heHEJT5of2DSNeQMqPmLwTaKK02IAIN3K67\nTMMtFsoYmjQu6WWkjb5ctU9SxR5WeDY3HWzJ50p2OxMCAwEAAaOCARIwggEOMB0G\nA1UdDgQWBBS0m/24zvCsED+neDIWwubApJTRMDCB3gYDVR0jBIHWMIHTgBS0m/24\nzvCsED+neDIWwubApJTRMKGBr6SBrDCBqTELMAkGA1UEBhMCVVMxEzARBgNVBAgT\nCldhc2hpbmd0b24xEDAOBgNVBAcTB1NlYXR0a2UxDzANBgNVBAoTBk9yYWNsZTEM\nMAoGA1UECxMDT0NJMScwJQYDVQQDEx5jZXJ0aWZpY2F0ZUdlbmVyYXRlZFRlc3RV\ncGRhdGUxKzApBgkqhkiG9w0BCQEWHG9yY2hlc3RyYXRpb25fZ3JwQG9yYWNsZS5j\nb22CCQD38s72+pw58jAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBBQUAA4IBAQCM\nK2G+XvJG2H6N7e3nr4I3f6FuCEo/6LL3leEiYeWlz+Nqm/z7PmLZOseTsTZRTmv7\ncIvoboQnQnUeV1kShW/smr3EKkiuVDz8nngxHkYneUKlSQqyr5J3JrUb53VGyvqi\nh9MWzYYfFvbtQdiz0VelY520D0YJmsvFCAY+ilY4Y7rG5sdqZhhmD1FIJwdi3gJl\nLyRL9XCB3ZfS1Q9marFsAbNQOTh5mKfkZs0aUo3zU1Dhh+usodp3azu5Nergb95+\nGzT4US7kxywB95qBSH7oZX0cBLv/4y3dfkZtIleHG/luAbpEiWHOL0GPaXtJuc0X\nGUpcphW5CCBfWVB5HrS/\n-----END CERTIFICATE-----" }
variable "certificate_certificate_name" { default = "certificateName2" }
variable "certificate_passphrase" { default = "someP@ssphrase" }
variable "certificate_private_key" { default = "-----BEGIN RSA PRIVATE KEY-----\nMIIEpQIBAAKCAQEAxKk7O/xjY7Y70KGpmRHuWTG0j82sNY7bLyiBaKpRNvhW5fI1\n+Ws7fWUJk4wolLGpVSI8lKZPjRiEUaglgLwZdsH1zFlfkksTFYPBqyfth5sSFo8G\nvzlCjQu6eY5roOw3jvHHOk8kj/37fF6a+fgYmEWatohLzV8XbrUZsygILKCDYjJ4\nFY3gcnEmNIfSjQJGiQl6UOPUsPTAuGzApZx4n0w81dPBWdus9sxZMX/9fBM2sjPr\nxNQphByEAVjQQE0tN03cZczWF4cQlPmh/YNI15Ayo+YvBNoorTYgAg3crrtMwy0W\nyhiaNC7pZaSNvly1T1LFHlZ4NjcdbMnnSnY7EwIDAQABAoIBAQCZttEW1PJ2qKTe\nMM4YV9yeGOzSDeeRwURu+tETlzR9q+R4PTaU6o0IOSYgDshVWPxLD8ScR9YcKdKa\n5LvTgLLCkV5nSGAcP9P92AtTlZTijoG34jla1A3Boy4w/vH/SUMv1qlrWRrx9YpR\nLr0hrxrCQn1oOOZO8wmkvacF0r946OoHcYPcyf7oX2X6L7hPoQBmMCNpQBekPaq1\nt9SlIZBKxyj/Cmjx1/Vh7kce4v5DZ2I3FZeeACH/+nXLHx4QfbiAovbBCJv8bi0R\nkh7tvS4LwrQVjF0jgFOqkZKBoGl2Kk42KCj45YwiIxek1cqD8MmoJImjfK5HME49\n+XFtDYaxAoGBAO4bMpfC3Pjqg1FFr3ppL2n+/3XbZSocjeGr+yu+/wVEUqD8hI94\nIACYmtHoFz4OgwvkRAJMcSnOXiSc8WB2Zer0Ck4NhT5aIKNPBVhEzVrW+ztMWEPt\nS0x05zXJqrG4rgedKLLviLfzbLn/o66azcyBQigXQQmyFS8YuGdNzgq5AoGBANNw\nsQIcKIriZfiq6Odab3mRDEsWaKCsQO9fq5tDTwlqVxv4ZatAUUEPlqs36DrSHF7i\n/OMcQ1B4Tmxl36/hHPPhFaGqWQd0gUBSTBpD8S/UlvIsW7RUVgbzGdIKg+LrHkt/\ngve0ztFjRUMiWWlh+zLS+0UYcZKoveM+DkEo/t4rAoGBAJCMflhXeiK66+Go9nDP\n7nLg1WFNQcsg7plz+NWD6/ncknUdF7SpbnQuML8BsSqBUNklSIxEr+Z7W/fKN1ci\nSZkrch4UYzGJMYsy68G8cxaUsIw9OtBx/GZ8CelDdpbR0QTaSdznQg66fuUj5tCX\nNhzS08DW7Srfq7Cdx0UfnNgxAoGAN7e8jWfKLz8Vy/+NvFmSgqB8ctUG4UBDJFE7\nsYB9NWf2aIJ+mDAsuU5YT+o9ojJ4E3ERXu+1tWDemaYG2hwiOwoXXzC6oUJMRRzq\nvQkmZ4pH9K1HxS9sNAyfwz/OHWAD7bU+V/Qng/n66IQGt8SAI5aBbGXSl+krkNRr\ngTRCZV0CgYEAwQlPqRY88/v9reFHcE5ylGiaV2An4R/83VGic2ELwC7W01ZLPc0u\nGDCTWDM1d4f9otU81g8QTKjes5cvoFIcPs+PiEFKzjTIs/B/rkZF9fmPrZlm07bb\n8O9dnl1VE1Odns9NOHagf+SEmr8okSrYJ30PY/gdowx294/fYpHV94U=\n-----END RSA PRIVATE KEY-----" }
variable "certificate_public_certificate" { default = "-----BEGIN CERTIFICATE-----\nMIIE6zCCA9OgAwIBAgIJAPfyzvb6nDnyMA0GCSqGSIb3DQEBBQUAMIGpMQswCQYD\nVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHU2VhdHRrZTEP\nMA0GA1UEChMGT3JhY2xlMQwwCgYDVQQLEwNPQ0kxJzAlBgNVBAMTHmNlcnRpZmlj\nYXRlR2VuZXJhdGVkVGVzdFVwZGF0ZTErMCkGCSqGSIb3DQEJARYcb3JjaGVzdHJh\ndGlvbl9ncnBAb3JhY2xlLmNvbTAeFw0xODAyMjYyMjEyNDdaFw0xOTAyMjYyMjEy\nNDdaMIGpMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UE\nBxMHU2VhdHRrZTEPMA0GA1UEChMGT3JhY2xlMQwwCgYDVQQLEwNPQ0kxJzAlBgNV\nBAMTHmNlcnRpZmljYXRlR2VuZXJhdGVkVGVzdFVwZGF0ZTErMCkGCSqGSIb3DQEJ\nARYcb3JjaGVzdHJhdGlvbl9ncnBAb3JhY2xlLmNvbTCCASIwDQYJKoZIhvcNAQEB\nBQADggEPADCCAQoCggEBAMSpOzv8Y2O2O9ChqZkR7lkxtI/NrDWO2y8ogWiqUTb4\nVuXyNflrO31lCZOMKJSxqVUiPJSmT40YhFGoJYC8GXbB9cxZX5JLExWDwasn7Yeb\nEhaPBr85Qo0LunmOa6DsN47xxzpPJI/9+3xemvn4GJhFmraIS81fF261GbMoCCyg\ng2IyeBWN4HJxJjSH0o0CRokJelDj1LD0wLhswKWceJ9MPNXTwVnbrPbMWTF//XwT\nNrIz68TUKYQchAFY0EBNLTdN3GXM1heHEJT5of2DSNeQMqPmLwTaKK02IAIN3K67\nTMMtFsoYmjQu6WWkjb5ctU9SxR5WeDY3HWzJ50p2OxMCAwEAAaOCARIwggEOMB0G\nA1UdDgQWBBS0m/24zvCsED+neDIWwubApJTRMDCB3gYDVR0jBIHWMIHTgBS0m/24\nzvCsED+neDIWwubApJTRMKGBr6SBrDCBqTELMAkGA1UEBhMCVVMxEzARBgNVBAgT\nCldhc2hpbmd0b24xEDAOBgNVBAcTB1NlYXR0a2UxDzANBgNVBAoTBk9yYWNsZTEM\nMAoGA1UECxMDT0NJMScwJQYDVQQDEx5jZXJ0aWZpY2F0ZUdlbmVyYXRlZFRlc3RV\ncGRhdGUxKzApBgkqhkiG9w0BCQEWHG9yY2hlc3RyYXRpb25fZ3JwQG9yYWNsZS5j\nb22CCQD38s72+pw58jAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBBQUAA4IBAQCM\nK2G+XvJG2H6N7e3nr4I3f6FuCEo/6LL3leEiYeWlz+Nqm/z7PmLZOseTsTZRTmv7\ncIvoboQnQnUeV1kShW/smr3EKkiuVDz8nngxHkYneUKlSQqyr5J3JrUb53VGyvqi\nh9MWzYYfFvbtQdiz0VelY520D0YJmsvFCAY+ilY4Y7rG5sdqZhhmD1FIJwdi3gJl\nLyRL9XCB3ZfS1Q9marFsAbNQOTh5mKfkZs0aUo3zU1Dhh+usodp3azu5Nergb95+\nGzT4US7kxywB95qBSH7oZX0cBLv/4y3dfkZtIleHG/luAbpEiWHOL0GPaXtJuc0X\nGUpcphW5CCBfWVB5HrS/\n-----END CERTIFICATE-----" }
				` + compartmentIdVariableStr2 + CertificateResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ca_certificate", "-----BEGIN CERTIFICATE-----\nMIIE6zCCA9OgAwIBAgIJAPfyzvb6nDnyMA0GCSqGSIb3DQEBBQUAMIGpMQswCQYD\nVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHU2VhdHRrZTEP\nMA0GA1UEChMGT3JhY2xlMQwwCgYDVQQLEwNPQ0kxJzAlBgNVBAMTHmNlcnRpZmlj\nYXRlR2VuZXJhdGVkVGVzdFVwZGF0ZTErMCkGCSqGSIb3DQEJARYcb3JjaGVzdHJh\ndGlvbl9ncnBAb3JhY2xlLmNvbTAeFw0xODAyMjYyMjEyNDdaFw0xOTAyMjYyMjEy\nNDdaMIGpMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UE\nBxMHU2VhdHRrZTEPMA0GA1UEChMGT3JhY2xlMQwwCgYDVQQLEwNPQ0kxJzAlBgNV\nBAMTHmNlcnRpZmljYXRlR2VuZXJhdGVkVGVzdFVwZGF0ZTErMCkGCSqGSIb3DQEJ\nARYcb3JjaGVzdHJhdGlvbl9ncnBAb3JhY2xlLmNvbTCCASIwDQYJKoZIhvcNAQEB\nBQADggEPADCCAQoCggEBAMSpOzv8Y2O2O9ChqZkR7lkxtI/NrDWO2y8ogWiqUTb4\nVuXyNflrO31lCZOMKJSxqVUiPJSmT40YhFGoJYC8GXbB9cxZX5JLExWDwasn7Yeb\nEhaPBr85Qo0LunmOa6DsN47xxzpPJI/9+3xemvn4GJhFmraIS81fF261GbMoCCyg\ng2IyeBWN4HJxJjSH0o0CRokJelDj1LD0wLhswKWceJ9MPNXTwVnbrPbMWTF//XwT\nNrIz68TUKYQchAFY0EBNLTdN3GXM1heHEJT5of2DSNeQMqPmLwTaKK02IAIN3K67\nTMMtFsoYmjQu6WWkjb5ctU9SxR5WeDY3HWzJ50p2OxMCAwEAAaOCARIwggEOMB0G\nA1UdDgQWBBS0m/24zvCsED+neDIWwubApJTRMDCB3gYDVR0jBIHWMIHTgBS0m/24\nzvCsED+neDIWwubApJTRMKGBr6SBrDCBqTELMAkGA1UEBhMCVVMxEzARBgNVBAgT\nCldhc2hpbmd0b24xEDAOBgNVBAcTB1NlYXR0a2UxDzANBgNVBAoTBk9yYWNsZTEM\nMAoGA1UECxMDT0NJMScwJQYDVQQDEx5jZXJ0aWZpY2F0ZUdlbmVyYXRlZFRlc3RV\ncGRhdGUxKzApBgkqhkiG9w0BCQEWHG9yY2hlc3RyYXRpb25fZ3JwQG9yYWNsZS5j\nb22CCQD38s72+pw58jAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBBQUAA4IBAQCM\nK2G+XvJG2H6N7e3nr4I3f6FuCEo/6LL3leEiYeWlz+Nqm/z7PmLZOseTsTZRTmv7\ncIvoboQnQnUeV1kShW/smr3EKkiuVDz8nngxHkYneUKlSQqyr5J3JrUb53VGyvqi\nh9MWzYYfFvbtQdiz0VelY520D0YJmsvFCAY+ilY4Y7rG5sdqZhhmD1FIJwdi3gJl\nLyRL9XCB3ZfS1Q9marFsAbNQOTh5mKfkZs0aUo3zU1Dhh+usodp3azu5Nergb95+\nGzT4US7kxywB95qBSH7oZX0cBLv/4y3dfkZtIleHG/luAbpEiWHOL0GPaXtJuc0X\nGUpcphW5CCBfWVB5HrS/\n-----END CERTIFICATE-----"),
					resource.TestCheckResourceAttr(resourceName, "certificate_name", "certificateName2"),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "passphrase", "someP@ssphrase"),
					resource.TestCheckResourceAttr(resourceName, "private_key", "-----BEGIN RSA PRIVATE KEY-----\nMIIEpQIBAAKCAQEAxKk7O/xjY7Y70KGpmRHuWTG0j82sNY7bLyiBaKpRNvhW5fI1\n+Ws7fWUJk4wolLGpVSI8lKZPjRiEUaglgLwZdsH1zFlfkksTFYPBqyfth5sSFo8G\nvzlCjQu6eY5roOw3jvHHOk8kj/37fF6a+fgYmEWatohLzV8XbrUZsygILKCDYjJ4\nFY3gcnEmNIfSjQJGiQl6UOPUsPTAuGzApZx4n0w81dPBWdus9sxZMX/9fBM2sjPr\nxNQphByEAVjQQE0tN03cZczWF4cQlPmh/YNI15Ayo+YvBNoorTYgAg3crrtMwy0W\nyhiaNC7pZaSNvly1T1LFHlZ4NjcdbMnnSnY7EwIDAQABAoIBAQCZttEW1PJ2qKTe\nMM4YV9yeGOzSDeeRwURu+tETlzR9q+R4PTaU6o0IOSYgDshVWPxLD8ScR9YcKdKa\n5LvTgLLCkV5nSGAcP9P92AtTlZTijoG34jla1A3Boy4w/vH/SUMv1qlrWRrx9YpR\nLr0hrxrCQn1oOOZO8wmkvacF0r946OoHcYPcyf7oX2X6L7hPoQBmMCNpQBekPaq1\nt9SlIZBKxyj/Cmjx1/Vh7kce4v5DZ2I3FZeeACH/+nXLHx4QfbiAovbBCJv8bi0R\nkh7tvS4LwrQVjF0jgFOqkZKBoGl2Kk42KCj45YwiIxek1cqD8MmoJImjfK5HME49\n+XFtDYaxAoGBAO4bMpfC3Pjqg1FFr3ppL2n+/3XbZSocjeGr+yu+/wVEUqD8hI94\nIACYmtHoFz4OgwvkRAJMcSnOXiSc8WB2Zer0Ck4NhT5aIKNPBVhEzVrW+ztMWEPt\nS0x05zXJqrG4rgedKLLviLfzbLn/o66azcyBQigXQQmyFS8YuGdNzgq5AoGBANNw\nsQIcKIriZfiq6Odab3mRDEsWaKCsQO9fq5tDTwlqVxv4ZatAUUEPlqs36DrSHF7i\n/OMcQ1B4Tmxl36/hHPPhFaGqWQd0gUBSTBpD8S/UlvIsW7RUVgbzGdIKg+LrHkt/\ngve0ztFjRUMiWWlh+zLS+0UYcZKoveM+DkEo/t4rAoGBAJCMflhXeiK66+Go9nDP\n7nLg1WFNQcsg7plz+NWD6/ncknUdF7SpbnQuML8BsSqBUNklSIxEr+Z7W/fKN1ci\nSZkrch4UYzGJMYsy68G8cxaUsIw9OtBx/GZ8CelDdpbR0QTaSdznQg66fuUj5tCX\nNhzS08DW7Srfq7Cdx0UfnNgxAoGAN7e8jWfKLz8Vy/+NvFmSgqB8ctUG4UBDJFE7\nsYB9NWf2aIJ+mDAsuU5YT+o9ojJ4E3ERXu+1tWDemaYG2hwiOwoXXzC6oUJMRRzq\nvQkmZ4pH9K1HxS9sNAyfwz/OHWAD7bU+V/Qng/n66IQGt8SAI5aBbGXSl+krkNRr\ngTRCZV0CgYEAwQlPqRY88/v9reFHcE5ylGiaV2An4R/83VGic2ELwC7W01ZLPc0u\nGDCTWDM1d4f9otU81g8QTKjes5cvoFIcPs+PiEFKzjTIs/B/rkZF9fmPrZlm07bb\n8O9dnl1VE1Odns9NOHagf+SEmr8okSrYJ30PY/gdowx294/fYpHV94U=\n-----END RSA PRIVATE KEY-----"),
					resource.TestCheckResourceAttr(resourceName, "public_certificate", "-----BEGIN CERTIFICATE-----\nMIIE6zCCA9OgAwIBAgIJAPfyzvb6nDnyMA0GCSqGSIb3DQEBBQUAMIGpMQswCQYD\nVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHU2VhdHRrZTEP\nMA0GA1UEChMGT3JhY2xlMQwwCgYDVQQLEwNPQ0kxJzAlBgNVBAMTHmNlcnRpZmlj\nYXRlR2VuZXJhdGVkVGVzdFVwZGF0ZTErMCkGCSqGSIb3DQEJARYcb3JjaGVzdHJh\ndGlvbl9ncnBAb3JhY2xlLmNvbTAeFw0xODAyMjYyMjEyNDdaFw0xOTAyMjYyMjEy\nNDdaMIGpMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UE\nBxMHU2VhdHRrZTEPMA0GA1UEChMGT3JhY2xlMQwwCgYDVQQLEwNPQ0kxJzAlBgNV\nBAMTHmNlcnRpZmljYXRlR2VuZXJhdGVkVGVzdFVwZGF0ZTErMCkGCSqGSIb3DQEJ\nARYcb3JjaGVzdHJhdGlvbl9ncnBAb3JhY2xlLmNvbTCCASIwDQYJKoZIhvcNAQEB\nBQADggEPADCCAQoCggEBAMSpOzv8Y2O2O9ChqZkR7lkxtI/NrDWO2y8ogWiqUTb4\nVuXyNflrO31lCZOMKJSxqVUiPJSmT40YhFGoJYC8GXbB9cxZX5JLExWDwasn7Yeb\nEhaPBr85Qo0LunmOa6DsN47xxzpPJI/9+3xemvn4GJhFmraIS81fF261GbMoCCyg\ng2IyeBWN4HJxJjSH0o0CRokJelDj1LD0wLhswKWceJ9MPNXTwVnbrPbMWTF//XwT\nNrIz68TUKYQchAFY0EBNLTdN3GXM1heHEJT5of2DSNeQMqPmLwTaKK02IAIN3K67\nTMMtFsoYmjQu6WWkjb5ctU9SxR5WeDY3HWzJ50p2OxMCAwEAAaOCARIwggEOMB0G\nA1UdDgQWBBS0m/24zvCsED+neDIWwubApJTRMDCB3gYDVR0jBIHWMIHTgBS0m/24\nzvCsED+neDIWwubApJTRMKGBr6SBrDCBqTELMAkGA1UEBhMCVVMxEzARBgNVBAgT\nCldhc2hpbmd0b24xEDAOBgNVBAcTB1NlYXR0a2UxDzANBgNVBAoTBk9yYWNsZTEM\nMAoGA1UECxMDT0NJMScwJQYDVQQDEx5jZXJ0aWZpY2F0ZUdlbmVyYXRlZFRlc3RV\ncGRhdGUxKzApBgkqhkiG9w0BCQEWHG9yY2hlc3RyYXRpb25fZ3JwQG9yYWNsZS5j\nb22CCQD38s72+pw58jAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBBQUAA4IBAQCM\nK2G+XvJG2H6N7e3nr4I3f6FuCEo/6LL3leEiYeWlz+Nqm/z7PmLZOseTsTZRTmv7\ncIvoboQnQnUeV1kShW/smr3EKkiuVDz8nngxHkYneUKlSQqyr5J3JrUb53VGyvqi\nh9MWzYYfFvbtQdiz0VelY520D0YJmsvFCAY+ilY4Y7rG5sdqZhhmD1FIJwdi3gJl\nLyRL9XCB3ZfS1Q9marFsAbNQOTh5mKfkZs0aUo3zU1Dhh+usodp3azu5Nergb95+\nGzT4US7kxywB95qBSH7oZX0cBLv/4y3dfkZtIleHG/luAbpEiWHOL0GPaXtJuc0X\nGUpcphW5CCBfWVB5HrS/\n-----END CERTIFICATE-----"),
				),
			},
		},
	})
}
