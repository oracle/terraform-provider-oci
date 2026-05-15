// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

type importedCertificateDetails struct {
	serialNumber string
	notBefore    string
	notAfter     string
	dnsNames     []string
}

func mustParseImportedCertificateDetails(certificatePem string) importedCertificateDetails {
	block, _ := pem.Decode([]byte(certificatePem))
	if block == nil {
		panic("failed to decode imported certificate PEM")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		panic(fmt.Sprintf("failed to parse imported certificate PEM: %v", err))
	}

	return importedCertificateDetails{
		serialNumber: cert.SerialNumber.String(),
		notBefore:    cert.NotBefore.UTC().String(),
		notAfter:     cert.NotAfter.UTC().String(),
		dnsNames:     cert.DNSNames,
	}
}

func importedPemStateValue(pem string) string {
	return pem + "\n"
}

var (
	requiredCertName         = "test-required-cert-" + utils.RandomString(10, utils.CharsetWithoutDigits)
	certNameForOptionalTests = "test-optional-cert-" + utils.RandomString(10, utils.CharsetWithoutDigits)
	importedCertName         = "test-imported-cert-" + utils.RandomString(10, utils.CharsetWithoutDigits)
	certificateTestUserName  = fmt.Sprintf(
		"certmgmt-%s@example.com",
		utils.RandomStringOrHttpReplayValue(10, utils.CharsetLowerCaseWithoutDigits, "basicuser"),
	)

	certNotBeforeCreate = time.Now().UTC().AddDate(0, 0, 1).Truncate(time.Millisecond).Format(time.RFC3339Nano)
	certNotBeforeUpdate = time.Now().UTC().AddDate(0, 0, 7).Truncate(time.Millisecond).Format(time.RFC3339Nano)
	certNotAfterCreate  = time.Now().UTC().AddDate(1, 0, 1).Truncate(time.Millisecond).Format(time.RFC3339Nano)
	certNotAfterUpdate  = time.Now().UTC().AddDate(1, 0, 7).Truncate(time.Millisecond).Format(time.RFC3339Nano)

	csrPem                   = "-----BEGIN CERTIFICATE REQUEST-----\nMIICzzCCAbcCAQAwgYkxCzAJBgNVBAYTAnVzMRMwEQYDVQQIDApXYXNoaW5ndG9u\nMREwDwYDVQQHDAhCZWxsZXZ1ZTEMMAoGA1UECgwDT0NJMQwwCgYDVQQLDANTRUMx\nGDAWBgNVBAMMD3d3dy50ZXN0Y3NyLmNvbTEcMBoGCSqGSIb3DQEJARYNdGVzdEB0\nZXN0LmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALrlpiuehIMz\nre0pa+G8WjPaltQOKsA5LZWFUmnkRvy9zGtVKXmcZ70m8+SIIW0svwhDENtvl6wA\nPfQX8G+9XelIzyUv62EG9PwcSqW7rLpq8VygH/6fbYLU/LG8wuIjVq8yU8AQSvcJ\nBSNEiPMxN7qmXviKXHyY8A3ramgY+2wIfqNPRfrM/udPGd8xls/MviqCXLl3rT5c\nLDwCOi+rqOgWwt9wkOzAVNQWtwCaGaed1j98T3QQtYjeSVeV/HaGUquK9nare0dF\n49SEgRuCpVSeD/PHN5lz1YE99bioEiqfdATw05xZ08wTezgscCKfuaMMAXDF2Q0N\nqMNaynCFuuECAwEAAaAAMA0GCSqGSIb3DQEBCwUAA4IBAQBXsLUdghNHVHM54DRJ\nwNJdICec9HX5OuwHP1C/QpwwcEeO4lIKTxww4WWwReJVIoUewgGaxYLh+izp40bP\no+wFR0NDC9eBaRtnBZiPYn8bVA2PBl9z0VS2+AsQHR9hlKH0G3iU+C0O2wVSJpQW\nAIU3c/2CNg9GCTWmJE+Jd8dTd21WDlARBKw7GXwToHQL7J3vY+2/S2e6hdbxh9aG\n7ZEOYssclMtVZxKPdGp9l3JbwBrk/9hv3kxFZvbtxjnn3ah5Sas+YwIewn2VPSwY\nfoewMRgyYZKUFmw/T3YH2ony9ouT8mgPMJys72iVXO+Ri31VdL7AVQHMA0uQEpKM\n+Pn4\n-----END CERTIFICATE REQUEST-----"
	importedCertificatePemV1 = `-----BEGIN CERTIFICATE-----
MIID1DCCArygAwIBAgIUUSNkzHmCRJWIZeCu/+RNSolSWgYwDQYJKoZIhvcNAQEL
BQAwSTEkMCIGA1UEAwwbQ29kZXggSW1wb3J0ZWQgQ2VydCBSb290IENBMRQwEgYD
VQQKDAtPcmFjbGUgVGVzdDELMAkGA1UEBhMCVVMwHhcNMjYwNDI5MjAxNjU4WhcN
MjgwODAxMjAxNjU4WjBBMRwwGgYDVQQDDBNpbXBvcnRlZC1jZXJ0LmxvY2FsMRQw
EgYDVQQKDAtPcmFjbGUgVGVzdDELMAkGA1UEBhMCVVMwggEiMA0GCSqGSIb3DQEB
AQUAA4IBDwAwggEKAoIBAQCuc7mFa9Zo29aQBg3FwPQXofPDaucPgdqngvYJXJFY
wgTPidgZUgiRi1npOY8PCDm7dn5B/vQuLK/GMrI3LqyOvzYNo36ilDHkt1zit8MP
Amxw2LBUg+6WOHTOjD2U7GogTBTXu8NhKhgdlehx6u6Nd+HvrPEaD9rb2jOZUcBZ
mi1zmctZTd+4kZqWJx42ktHzRLrWaRltu7gvr/cIjQ879tKQOnCncxFbLJyFUtGB
bTO1/2kL83l69GQGfb7qdJVy0r6MxlOYyNnI+kbclyagyhbTyPWUnn8Wy2ztvBYs
CoRc165Al8n4prGFGvtCapgGRRpOIMatHUT4JL+6q6fHAgMBAAGjgbswgbgwOQYD
VR0RBDIwMIITaW1wb3J0ZWQtY2VydC5sb2NhbIIZaW1wb3J0ZWQtY2VydC5leGFt
cGxlLmNvbTAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB/wQEAwIFoDAdBgNVHSUEFjAU
BggrBgEFBQcDAQYIKwYBBQUHAwIwHQYDVR0OBBYEFLBPB1H6lhA/vh9muRT88YBZ
jk6qMB8GA1UdIwQYMBaAFNs69dZIK4qugHpj93QD4Gv3p/yaMA0GCSqGSIb3DQEB
CwUAA4IBAQDGidVB+XQDB1dvjJ1v+kxKFjmqHH//9MAq+EYVrPXpNnUk03OfPuPv
oj8ybEwIU8lMX+acDFMnJaDGMzCOce2bCgmvGSNboUGpRE1jzYhPozQK1MjHt111
QUe7nOG/A+/qmdnYjNg1ZTRjaxoufGYJQSdFowhAGxyi3rtXVXeCVgsYPKljURI1
cKmcp1YPIz9oaFZdGTNCUvxY113Y/GHsArROfI48fZTzxUIvZ+3zCa3LFqQeRLe5
OmBV10uveTHS5e6GeHol1yEtpxFEhs0w8JtjxdAjKaBrnB88s4rvsHIyXmkvW9L3
zriS1Nol4DlLdoH9POvN3y4PpZ6asH4y
-----END CERTIFICATE-----`
	importedPrivateKeyPemV1 = `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCuc7mFa9Zo29aQ
Bg3FwPQXofPDaucPgdqngvYJXJFYwgTPidgZUgiRi1npOY8PCDm7dn5B/vQuLK/G
MrI3LqyOvzYNo36ilDHkt1zit8MPAmxw2LBUg+6WOHTOjD2U7GogTBTXu8NhKhgd
lehx6u6Nd+HvrPEaD9rb2jOZUcBZmi1zmctZTd+4kZqWJx42ktHzRLrWaRltu7gv
r/cIjQ879tKQOnCncxFbLJyFUtGBbTO1/2kL83l69GQGfb7qdJVy0r6MxlOYyNnI
+kbclyagyhbTyPWUnn8Wy2ztvBYsCoRc165Al8n4prGFGvtCapgGRRpOIMatHUT4
JL+6q6fHAgMBAAECggEABiQWvMaeBpTuiaVc6MAxF3/g99kUhyw4CTuH865pg2lZ
CAE/NSz5Fj6EjLw09/9/IMU8DpCuAc4a5ew079JZyrCac4J0275Bv15hC64SVL3u
GU7XbjIyvOEB4592L37fAYIMHcgDSbFUpf3378KoYRx+PJSm7Bbwylp1BImnLsA6
npTSC/gfJzPtpgKlPuADcn2eHlZkqxF1GEdzNbee69sEf7UoY54y5Z7CaaObRB55
jB4qmSHD5Zk3Le7h7oXFKp39BFRPiKJlnBPmZ10Ca2pdPH4wXDpNj7qm5gG/qwWN
SeaaXvhUbJ+Xtj4LJ04CrRLjXYL5z2NSVnZuaLiCgQKBgQD0OF/X8mWj1/OzIz3H
/Van334ef0VMDMUcW4T+vGdL/4lnYHY85ZNleh/f0P9XnypcjwQwStNPWwiq+83w
iJC0Gmsf71eUyDPIJI2rOp3COdVeWo6+NfDp2iNoXYksneGbL4pETn1fuiCFdgO+
MLPqAvvx243nR0+FZBkMVclnQQKBgQC23drIXsm0WeTwcaiA/7m+KWg4soUCJc3d
qUY2pGKAYZunBzDT6DPy6vbDGKNFib0NPPwHrI52UYKpfWta0jixecndm9e8KF1W
Kr+OuzqNvt2RNQv1oaDYXkFqPx8V20HYCbGykLUrcTvl8kQ/P/fuJGb1lJpFjrxb
Zj37ndGVBwKBgQDfWxe9BoISwRSvB1IfsAEq82oDNl8IWL2FW+Zcm67VgNYH5977
fvO5YWH5bsm5N2ak1F/+UtDDyPw6dKU6kYvWTDshL+Knvhott9Chg7B4zC/qZWQb
bnXxSoILl1A3sAV/ypDoO+jKPg40QWTW8u/429XUsvorRbriwlNuRp82gQKBgBlp
POTJHzhboY+0C5lZb+M3986hKBCDVlcuFZ60Oyfh6HJOcn+P7eAcvIuh2dZo7Idx
aQhb9zQD4adR0WA+hDhNVzrfVBxGc6nsqM/ZvqpB0qJB/9ZzTlmix25PNS+hV0YB
GttYdoSB6LPsdYVgi0BheJk5vDJcQlplG8niRWRzAoGAIvTSa/DG8XR4EsEUFqtX
POFZD+rtbo3billYf5JenBtAuhUw0Ub1Xs5A+z+zpMbEjlv26qFH+5JUhc+bFxit
tCbAF7RXkhU1c5MTy+R4dAq8ZDfy44eZbNnSycfxhHku8zP7a3rwfoqxQOqGHsML
ERgQfPPismBqinuoMVUCHLU=
-----END PRIVATE KEY-----`
	importedCertificatePemV2 = `-----BEGIN CERTIFICATE-----
MIID1DCCArygAwIBAgIUUSNkzHmCRJWIZeCu/+RNSolSWgcwDQYJKoZIhvcNAQEL
BQAwSTEkMCIGA1UEAwwbQ29kZXggSW1wb3J0ZWQgQ2VydCBSb290IENBMRQwEgYD
VQQKDAtPcmFjbGUgVGVzdDELMAkGA1UEBhMCVVMwHhcNMjYwNDI5MjAyMjI3WhcN
MjgxMDE1MjAyMjI3WjBBMRwwGgYDVQQDDBNpbXBvcnRlZC1jZXJ0LmxvY2FsMRQw
EgYDVQQKDAtPcmFjbGUgVGVzdDELMAkGA1UEBhMCVVMwggEiMA0GCSqGSIb3DQEB
AQUAA4IBDwAwggEKAoIBAQCc9HIrs95AGbUYZiwjR8d2nX86dS+BlEjkjptemkj8
vifT4U9dgaG91124DRPMSU60WwyJYg9zmGj9UF95nBfStwnmG/op/9EcgDZ9k6U6
5FAgiZfvFChtWGSPFDYj6YISMX9mB0+gxIT5O58g1FmZ+HvaJM7CFIz2R309bj2Q
pVyTXwdzD6M9TnroG0G4LUSMPHtytDVmbbYDVXvB8HVs/Vsp8krlQZbAXz6Ht9ZA
uJBjgu4tQNnTnSIt9LLl9yqEMuNgyekkQLeDhuu5e3toZxqXD/LnNNt3ch++eNYJ
Ic2ieUydjlzLkl6eYILYg6QAO6Fy5KOI/brjiRlYpOCdAgMBAAGjgbswgbgwOQYD
VR0RBDIwMIITaW1wb3J0ZWQtY2VydC5sb2NhbIIZaW1wb3J0ZWQtY2VydC5leGFt
cGxlLmNvbTAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB/wQEAwIFoDAdBgNVHSUEFjAU
BggrBgEFBQcDAQYIKwYBBQUHAwIwHQYDVR0OBBYEFMoeegvvEOgKnx53URED9u/v
ArU5MB8GA1UdIwQYMBaAFNs69dZIK4qugHpj93QD4Gv3p/yaMA0GCSqGSIb3DQEB
CwUAA4IBAQBNR95nC9EirgItcxaVraI5Os2LdOdW6WDm+RQqyo5geMfuc1ne11Wm
EFUzn8JWg1L2Wv0M8QJ5Y/gW6IZ8AkOZh7QjhQmIodFIPhJ1Z3CnTj7d1KdCLEDr
Od3fQbe6KFpKAgXfikqvnMxU+anV0BAMGPWrdpdEGXseEQCqnCz9PcOfdS7ryrva
HCzkW4X5x7Dnw3VFH3dUX+tXh103mJZF94wnAjx3krOKuZbLUMEfIbrWJR2tXJ5C
yfbwV/g1X8QQW/E2qDXHvoeSJnrKkl8jDNxVyJqFms90S8oNk2GtX7BtnSLH/gu9
196E02L2sH+XjUqyIOMdIK37VxlQXcr3
-----END CERTIFICATE-----`
	importedPrivateKeyPemV2 = `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCc9HIrs95AGbUY
ZiwjR8d2nX86dS+BlEjkjptemkj8vifT4U9dgaG91124DRPMSU60WwyJYg9zmGj9
UF95nBfStwnmG/op/9EcgDZ9k6U65FAgiZfvFChtWGSPFDYj6YISMX9mB0+gxIT5
O58g1FmZ+HvaJM7CFIz2R309bj2QpVyTXwdzD6M9TnroG0G4LUSMPHtytDVmbbYD
VXvB8HVs/Vsp8krlQZbAXz6Ht9ZAuJBjgu4tQNnTnSIt9LLl9yqEMuNgyekkQLeD
huu5e3toZxqXD/LnNNt3ch++eNYJIc2ieUydjlzLkl6eYILYg6QAO6Fy5KOI/brj
iRlYpOCdAgMBAAECggEAOqYZF32BugAhLHnZWmGTBh66EPUDCG+JgWZMZmBOj6ra
pRDXIEC1WIPmcETqtfZao4g2OgvbbD1yVd23n3CVclaVK1ge9Zyy3eyG8AvAvS3d
RxtDG8IzjRBvmwtZY7f4s6WtTknDFnb4nU0LHDBdPuycExqX6+vT9J8oDFFqTeb2
uQDS2JHKtK1hg3fspRoxYs6jcEgankYlqcdiuo9YatippEgJKWPyzyRFG7TJMjb2
6C5e9/dPSuyV6KqzKh1bX/4+CBhg7GYhtdc7f70L58lF0uO1jimaXj4jNiXYaFA3
LMJm3o031hX2xifeJA79NpQmWKJrSAnUzR62gD8RvQKBgQDPchG38txBFkRfWro1
/eBh3zYF0O7MHTwAEyGv59ZgX62yTbCpLpq1P1w0Vj1zIxUfsXIWZC+R6GiRievT
u5Nd/9coQdCb68+GYiT2aldQtm/feETs+j3iN1Ktc/csnY0w1nZC/WAp5VnX5vXp
pnShAE5LYZDdQpqziamqlGD5uwKBgQDBsQYaCV39XR2MqZ7OD61bd9OEJHe91M4x
QCbJYACQp3sppnRmio6gUwkpviIjZo/D37l64U6BYVKZo4QcopdXQSfZXtIRna92
mxW8qTcCDZRyo1cOT5KU2uJZtiMM3UdSyQFxK3mIze219JL6VPOeyO0wzXbmVOKI
qdvh5LUdhwKBgQCX+fzLI+yC7ppGtYRIe+dALLQ6TJg2B5DmiR2b7qBa2J27TDmJ
qLAE6xdU7H12dZRLNMfdaaaEIZ1CSqXv57MOIZnyT01Sfc76p7pmoP+5DP8m28i7
SJW/vkIe+x9xmzdOOUBI0bH636ETkIxGkgFN1piFl0P3JY2KzOb1JmFpcwKBgQCy
E+2GNqKBdQDb+H8W+Pbrp6mCTn0AvQyt30tJzcf3NuRXkMBoHm+vG2jLRhHg9kSr
U6AWR9uDvFTTCLFIGnyGLlNS2IjOxWky5LcBc5otewrGs0Iw2BvuW6fLz1c+Vpti
k4eN4uikTKCGr8wJ99zcWeMpAV6VuelkuLZ+Y2XZNQKBgDxTeE4hH0GBtTjzsbwL
Gy552fkLFB4MNdBGD3HYZNkEa0K771cDAkOprqMWlIbHuHMrol8ySry7Jqini/pe
r1zh3TW/lG5Bi3YsuO3xQ+/QZvFWOJ9KEzYHZ48fCPvsGG0lIpjXzGEwWb6Fd/14
C25C8uyG0BVJmH3aXUfuKfNI
-----END PRIVATE KEY-----`
	importedCertChainPem = `-----BEGIN CERTIFICATE-----
MIIDgzCCAmugAwIBAgIUPkRQkm0NaovEUyojIoqkYWHtYQswDQYJKoZIhvcNAQEL
BQAwSTEkMCIGA1UEAwwbQ29kZXggSW1wb3J0ZWQgQ2VydCBSb290IENBMRQwEgYD
VQQKDAtPcmFjbGUgVGVzdDELMAkGA1UEBhMCVVMwHhcNMjYwNDI5MjAxNjU4WhcN
MzYwNDI2MjAxNjU4WjBJMSQwIgYDVQQDDBtDb2RleCBJbXBvcnRlZCBDZXJ0IFJv
b3QgQ0ExFDASBgNVBAoMC09yYWNsZSBUZXN0MQswCQYDVQQGEwJVUzCCASIwDQYJ
KoZIhvcNAQEBBQADggEPADCCAQoCggEBAM5yng/WTdPusublw9tqzpfqN5D6AAZR
XelergeDwsh3xRul0mqwE26S9SwmRZO0GnJJ4Gmf/Wnoj/2+LGgGs4Z5/qjgg6FG
4Uy6QzDTapbGldqzoKGEFXvJ4AQgwpotuosp7Xtm9OG0LMU8V+X/OePBEfixQaln
n0fXifVZDHQDQZOTF0XJAnDnu8oe8q9U35DHRjwoghAcjk4lN9E1025FFEVvrFTR
jK0H6Y5W/ZJNIMUk9XPRpatLQphLrQAKRF3mH0xMHxJttUPrCQSRYyBj9BXK4aIv
5Z3+AQP40mceptckAi91vyCzEfISJPr3KZltpIbRYc2kEPi55pZSRWkCAwEAAaNj
MGEwHQYDVR0OBBYEFNs69dZIK4qugHpj93QD4Gv3p/yaMB8GA1UdIwQYMBaAFNs6
9dZIK4qugHpj93QD4Gv3p/yaMA8GA1UdEwEB/wQFMAMBAf8wDgYDVR0PAQH/BAQD
AgGGMA0GCSqGSIb3DQEBCwUAA4IBAQCBuL130bVAtzwKbS5IAqM7NlqrJazVOOhi
nNQDvY1C8CrRqZJIdMyTNURIiILXrdGlAjX/7xD3NLR5Ea5Ka6QGN8tVsTrtLdYw
32rGrsUQ38FbEj8dD4N5wPDhrqsYw/lDn2yka2VjTIDBnwVB7r/g9TOyElPIoCyS
oV1HruBPRqeF78GvUxdIJky15JM6LSAVR+7nqbcMZOEFyxMRdCXztyY5jrQmZJug
TFgSRG/Q92we69f7VsZ36+b+qztokiJkoCdfLE6irdWdVlu7KeGK0vRnenXWE/Jx
	L/T/I7wj7Ol+eF5VGVybNyEEZL1TZjIWzbg8rtO8RcJqcdt6vRpM
	-----END CERTIFICATE-----`

	importedCertificateDetailsV1 = mustParseImportedCertificateDetails(importedCertificatePemV1)
	importedCertificateDetailsV2 = mustParseImportedCertificateDetails(importedCertificatePemV2)

	CertificatesManagementCertificateRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate", "test_certificate", acctest.Required, acctest.Create, certificatesManagementCertificateRepresentation)

	CertificatesManagementCertificateResourceConfig = CertificatesManagementCertificateResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate", "test_certificate", acctest.Optional, acctest.Update, certificatesManagementCertificateRepresentation)

	CertificatesManagementcertificateSingularDataSourceRepresentation = map[string]interface{}{
		"certificate_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_certificates_management_certificate.test_certificate.id}`},
	}

	certificatesManagementCertificateDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: certNameForOptionalTests},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: certificatesManagementCertificateDataSourceFilterRepresentation}}
	certificatesManagementCertificateDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_certificates_management_certificate.test_certificate.id}`}},
	}

	certificatesManagementCertificateRepresentation = map[string]interface{}{
		"certificate_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: certificateCertificateConfigRepresentationInternal},
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":               acctest.Representation{RepType: acctest.Required, Create: certNameForOptionalTests},
		"certificate_rules":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificateCertificateRulesRepresentation},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":        acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	certificateCertificateConfigRepresentation = map[string]interface{}{
		"config_type":                     acctest.Representation{RepType: acctest.Required, Create: `ISSUED_BY_INTERNAL_CA`},
		"cert_chain_pem":                  acctest.Representation{RepType: acctest.Optional, Create: `certChainPem`, Update: `certChainPem2`},
		"certificate_pem":                 acctest.Representation{RepType: acctest.Optional, Create: `certificatePem`, Update: `certificatePem2`},
		"certificate_profile_type":        acctest.Representation{RepType: acctest.Optional, Create: `TLS_SERVER_OR_CLIENT`},
		"csr_pem":                         acctest.Representation{RepType: acctest.Optional, Create: `csrPem`, Update: `csrPem2`},
		"issuer_certificate_authority_id": acctest.Representation{RepType: acctest.Optional, Create: issuerCaId},
		"key_algorithm":                   acctest.Representation{RepType: acctest.Optional, Create: `RSA2048`},
		"private_key_pem":                 acctest.Representation{RepType: acctest.Optional, Create: `privateKeyPem`, Update: `privateKeyPem2`},
		"private_key_pem_passphrase":      acctest.Representation{RepType: acctest.Optional, Create: `privateKeyPemPassphrase`, Update: `privateKeyPemPassphrase2`},
		"signature_algorithm":             acctest.Representation{RepType: acctest.Optional, Create: `SHA256_WITH_RSA`},
		"subject":                         acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificatesManagementCertificateSubjectRepresentation},
		"subject_alternative_names":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificateCertificateConfigSubjectAlternativeNamesRepresentation},
		"validity":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificateCertificateConfigValidityRepresentation},
		"version_name":                    acctest.Representation{RepType: acctest.Optional, Create: `versionName`, Update: `versionName2`},
		"lifecycle":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: certificatesManagementIgnoreChangesRepresentation},
	}

	// Internal managed cert config
	certificateCertificateConfigRepresentationInternal = map[string]interface{}{
		"config_type":                     acctest.Representation{RepType: acctest.Required, Create: `ISSUED_BY_INTERNAL_CA`, Update: `ISSUED_BY_INTERNAL_CA`},
		"certificate_profile_type":        acctest.Representation{RepType: acctest.Required, Create: `TLS_SERVER_OR_CLIENT`},
		"issuer_certificate_authority_id": acctest.Representation{RepType: acctest.Required, Create: issuerCaId},
		"key_algorithm":                   acctest.Representation{RepType: acctest.Optional, Create: `RSA2048`},
		"signature_algorithm":             acctest.Representation{RepType: acctest.Optional, Create: `SHA256_WITH_RSA`},
		"subject":                         acctest.RepresentationGroup{RepType: acctest.Required, Group: certificatesManagementCertificateSubjectRepresentation},
		"subject_alternative_names":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificateCertificateConfigSubjectAlternativeNamesRepresentation},
		"validity":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificateCertificateConfigValidityRepresentation},
		"version_name":                    acctest.Representation{RepType: acctest.Optional, Create: `versionName`, Update: `versionName2`},
	}

	certificateCertificateConfigRepresentationCSR = map[string]interface{}{
		"config_type":                     acctest.Representation{RepType: acctest.Required, Create: `MANAGED_EXTERNALLY_ISSUED_BY_INTERNAL_CA`, Update: `MANAGED_EXTERNALLY_ISSUED_BY_INTERNAL_CA`},
		"csr_pem":                         acctest.Representation{RepType: acctest.Required, Create: `-----BEGIN CERTIFICATE REQUEST-----\nMIICzzCCAbcCAQAwgYkxCzAJBgNVBAYTAnVzMRMwEQYDVQQIDApXYXNoaW5ndG9u\nMREwDwYDVQQHDAhCZWxsZXZ1ZTEMMAoGA1UECgwDT0NJMQwwCgYDVQQLDANTRUMx\nGDAWBgNVBAMMD3d3dy50ZXN0Y3NyLmNvbTEcMBoGCSqGSIb3DQEJARYNdGVzdEB0\nZXN0LmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALrlpiuehIMz\nre0pa+G8WjPaltQOKsA5LZWFUmnkRvy9zGtVKXmcZ70m8+SIIW0svwhDENtvl6wA\nPfQX8G+9XelIzyUv62EG9PwcSqW7rLpq8VygH/6fbYLU/LG8wuIjVq8yU8AQSvcJ\nBSNEiPMxN7qmXviKXHyY8A3ramgY+2wIfqNPRfrM/udPGd8xls/MviqCXLl3rT5c\nLDwCOi+rqOgWwt9wkOzAVNQWtwCaGaed1j98T3QQtYjeSVeV/HaGUquK9nare0dF\n49SEgRuCpVSeD/PHN5lz1YE99bioEiqfdATw05xZ08wTezgscCKfuaMMAXDF2Q0N\nqMNaynCFuuECAwEAAaAAMA0GCSqGSIb3DQEBCwUAA4IBAQBXsLUdghNHVHM54DRJ\nwNJdICec9HX5OuwHP1C/QpwwcEeO4lIKTxww4WWwReJVIoUewgGaxYLh+izp40bP\no+wFR0NDC9eBaRtnBZiPYn8bVA2PBl9z0VS2+AsQHR9hlKH0G3iU+C0O2wVSJpQW\nAIU3c/2CNg9GCTWmJE+Jd8dTd21WDlARBKw7GXwToHQL7J3vY+2/S2e6hdbxh9aG\n7ZEOYssclMtVZxKPdGp9l3JbwBrk/9hv3kxFZvbtxjnn3ah5Sas+YwIewn2VPSwY\nfoewMRgyYZKUFmw/T3YH2ony9ouT8mgPMJys72iVXO+Ri31VdL7AVQHMA0uQEpKM\n+Pn4\n-----END CERTIFICATE REQUEST-----`, Update: `csrPem2`},
		"issuer_certificate_authority_id": acctest.Representation{RepType: acctest.Required, Create: issuerCaId},
		"validity":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: certificateCertificateConfigValidityRepresentation},
		"version_name":                    acctest.Representation{RepType: acctest.Optional, Create: `versionName`, Update: `versionName2`},
	}

	certificateCertificateRulesRepresentation = map[string]interface{}{
		"advance_renewal_period": acctest.Representation{RepType: acctest.Required, Create: `P30D`, Update: `P45D`},
		"renewal_interval":       acctest.Representation{RepType: acctest.Required, Create: `P365D`, Update: `P100D`},
		"rule_type":              acctest.Representation{RepType: acctest.Required, Create: `CERTIFICATE_RENEWAL_RULE`},
	}

	certificateCertificateConfigSubjectAlternativeNamesRepresentation = map[string]interface{}{
		"type":  acctest.Representation{RepType: acctest.Optional, Create: `DNS`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: `www.oracle.com`},
	}
	certificateCertificateConfigValidityRepresentation = map[string]interface{}{
		"time_of_validity_not_after":  acctest.Representation{RepType: acctest.Required, Create: certNotAfterCreate, Update: certNotAfterUpdate},
		"time_of_validity_not_before": acctest.Representation{RepType: acctest.Optional, Create: certNotBeforeCreate, Update: certNotBeforeUpdate},
	}

	CertificatesManagementCertificateResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_user", "test_user", acctest.Required, acctest.Create, IdentityUserRepresentation)
)

func TestCertificatesManagementCertificateResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCertificatesManagementCertificateResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_certificates_management_certificate.test_certificate"
	datasourceName := "data.oci_certificates_management_certificates.test_certificates"
	singularDatasourceName := "data.oci_certificates_management_certificate.test_certificate"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CertificatesManagementCertificateResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate", "test_certificate", acctest.Optional, acctest.Create, certificatesManagementCertificateRepresentation), "certificatesmanagement", "certificate", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		Steps: []resource.TestStep{

			// verify create on CSR cert
			{
				Config: config + compartmentIdVariableStr +
					acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate", "test_certificate", acctest.Required, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(certificatesManagementCertificateRepresentation, map[string]interface{}{
							"certificate_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: certificateCertificateConfigRepresentationCSR},
							"name":               acctest.Representation{RepType: acctest.Required, Create: requiredCertName},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "certificate_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.config_type", "MANAGED_EXTERNALLY_ISSUED_BY_INTERNAL_CA"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.csr_pem", csrPem),
					resource.TestCheckResourceAttrSet(resourceName, "certificate_config.0.issuer_certificate_authority_id"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "name", requiredCertName),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + CertificatesManagementCertificateResourceDependencies,
			},

			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + CertificatesManagementCertificateResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate", "test_certificate", acctest.Optional, acctest.Create, certificatesManagementCertificateRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "certificate_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.certificate_profile_type", "TLS_SERVER_OR_CLIENT"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.config_type", "ISSUED_BY_INTERNAL_CA"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate_config.0.issuer_certificate_authority_id"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.key_algorithm", "RSA2048"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.signature_algorithm", "SHA256_WITH_RSA"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.common_name", "www.example.com"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.country", "US"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.distinguished_name_qualifier", "distinguishedNameQualifier"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.domain_component", "domainComponent"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.generation_qualifier", "JR"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.given_name", "Sir"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.initials", "HAM"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.locality_name", "Seattle"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.organization", "OCI"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.organizational_unit", "SecurityProducts"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.pseudonym", "pseudonym"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.serial_number", "serialNumber"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.state_or_province_name", "Washington"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.street", "123 Main Street"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.surname", "Last"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.title", "Lord"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate_config.0.subject.0.user_id"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject_alternative_names.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject_alternative_names.0.type", "DNS"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject_alternative_names.0.value", "www.oracle.com"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.validity.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.validity.0.time_of_validity_not_after", certNotAfterCreate),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.validity.0.time_of_validity_not_before", certNotBeforeCreate),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.version_name", "versionName"),
					resource.TestCheckResourceAttr(resourceName, "certificate_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_rules.0.advance_renewal_period", "P30D"),
					resource.TestCheckResourceAttr(resourceName, "certificate_rules.0.renewal_interval", "P365D"),
					resource.TestCheckResourceAttr(resourceName, "certificate_rules.0.rule_type", "CERTIFICATE_RENEWAL_RULE"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "config_type"),
					resource.TestCheckResourceAttrSet(resourceName, "current_version.0.certificate_id"),
					resource.TestCheckResourceAttrSet(resourceName, "current_version.0.issuer_ca_version_number"),
					resource.TestCheckResourceAttrSet(resourceName, "current_version.0.serial_number"),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.stages.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.subject_alternative_names.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.subject_alternative_names.0.type", "DNS"),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.subject_alternative_names.0.value", "www.oracle.com"),
					resource.TestCheckResourceAttrSet(resourceName, "current_version.0.time_created"),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.validity.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "current_version.0.validity.0.time_of_validity_not_after"),
					resource.TestCheckResourceAttrSet(resourceName, "current_version.0.validity.0.time_of_validity_not_before"),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.version_name", "versionName"),
					resource.TestCheckResourceAttrSet(resourceName, "current_version.0.version_number"),
					resource.TestCheckResourceAttrSet(resourceName, "issuer_certificate_authority_id"),
					resource.TestCheckResourceAttr(resourceName, "key_algorithm", "RSA2048"),
					resource.TestCheckResourceAttr(resourceName, "signature_algorithm", "SHA256_WITH_RSA"),
					resource.TestCheckResourceAttr(resourceName, "subject.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "subject.0.common_name", "www.example.com"),
					resource.TestCheckResourceAttrSet(resourceName, "subject.0.country"),
					resource.TestCheckResourceAttr(resourceName, "subject.0.distinguished_name_qualifier", "distinguishedNameQualifier"),
					resource.TestCheckResourceAttr(resourceName, "subject.0.domain_component", "domainComponent"),
					resource.TestCheckResourceAttr(resourceName, "subject.0.generation_qualifier", "JR"),
					resource.TestCheckResourceAttr(resourceName, "subject.0.given_name", "Sir"),
					resource.TestCheckResourceAttr(resourceName, "subject.0.initials", "HAM"),
					resource.TestCheckResourceAttr(resourceName, "subject.0.locality_name", "Seattle"),
					resource.TestCheckResourceAttr(resourceName, "subject.0.organization", "OCI"),
					resource.TestCheckResourceAttr(resourceName, "subject.0.organizational_unit", "SecurityProducts"),
					resource.TestCheckResourceAttr(resourceName, "subject.0.pseudonym", "pseudonym"),
					resource.TestCheckResourceAttr(resourceName, "subject.0.serial_number", "serialNumber"),
					resource.TestCheckResourceAttr(resourceName, "subject.0.state_or_province_name", "Washington"),
					resource.TestCheckResourceAttr(resourceName, "subject.0.street", "123 Main Street"),
					resource.TestCheckResourceAttr(resourceName, "subject.0.surname", "Last"),
					resource.TestCheckResourceAttr(resourceName, "subject.0.title", "Lord"),
					resource.TestCheckResourceAttrSet(resourceName, "subject.0.user_id"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", certNameForOptionalTests),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CertificatesManagementCertificateResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate", "test_certificate", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(certificatesManagementCertificateRepresentation, map[string]interface{}{
							"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "certificate_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.certificate_profile_type", "TLS_SERVER_OR_CLIENT"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.config_type", "ISSUED_BY_INTERNAL_CA"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate_config.0.issuer_certificate_authority_id"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.key_algorithm", "RSA2048"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.signature_algorithm", "SHA256_WITH_RSA"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.common_name", "www.example.com"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.country", "US"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.distinguished_name_qualifier", "distinguishedNameQualifier"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.domain_component", "domainComponent"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.generation_qualifier", "JR"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.given_name", "Sir"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.initials", "HAM"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.locality_name", "Seattle"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.organization", "OCI"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.organizational_unit", "SecurityProducts"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.pseudonym", "pseudonym"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.serial_number", "serialNumber"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.state_or_province_name", "Washington"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.street", "123 Main Street"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.surname", "Last"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.title", "Lord"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate_config.0.subject.0.user_id"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject_alternative_names.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject_alternative_names.0.type", "DNS"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject_alternative_names.0.value", "www.oracle.com"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.validity.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.validity.0.time_of_validity_not_after", certNotAfterCreate),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.validity.0.time_of_validity_not_before", certNotBeforeCreate),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.version_name", "versionName"),
					resource.TestCheckResourceAttr(resourceName, "certificate_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_rules.0.advance_renewal_period", "P30D"),
					resource.TestCheckResourceAttr(resourceName, "certificate_rules.0.renewal_interval", "P365D"),
					resource.TestCheckResourceAttr(resourceName, "certificate_rules.0.rule_type", "CERTIFICATE_RENEWAL_RULE"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttrSet(resourceName, "config_type"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", certNameForOptionalTests),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + CertificatesManagementCertificateResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate", "test_certificate", acctest.Optional, acctest.Update, certificatesManagementCertificateRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "certificate_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.certificate_profile_type", "TLS_SERVER_OR_CLIENT"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.config_type", "ISSUED_BY_INTERNAL_CA"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate_config.0.issuer_certificate_authority_id"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.key_algorithm", "RSA2048"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.signature_algorithm", "SHA256_WITH_RSA"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.common_name", "www.example.com"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.country", "US"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.distinguished_name_qualifier", "distinguishedNameQualifier"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.domain_component", "domainComponent"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.generation_qualifier", "JR"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.given_name", "Sir"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.initials", "HAM"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.locality_name", "Seattle"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.organization", "OCI"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.organizational_unit", "SecurityProducts"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.pseudonym", "pseudonym"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.serial_number", "serialNumber"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.state_or_province_name", "Washington"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.street", "123 Main Street"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.surname", "Last"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject.0.title", "Lord"),
					resource.TestCheckResourceAttrSet(resourceName, "certificate_config.0.subject.0.user_id"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject_alternative_names.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject_alternative_names.0.type", "DNS"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.subject_alternative_names.0.value", "www.oracle.com"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.validity.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.validity.0.time_of_validity_not_after", certNotAfterUpdate),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.validity.0.time_of_validity_not_before", certNotBeforeUpdate),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.version_name", "versionName2"),
					resource.TestCheckResourceAttr(resourceName, "certificate_rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_rules.0.advance_renewal_period", "P45D"),
					resource.TestCheckResourceAttr(resourceName, "certificate_rules.0.renewal_interval", "P100D"),
					resource.TestCheckResourceAttr(resourceName, "certificate_rules.0.rule_type", "CERTIFICATE_RENEWAL_RULE"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "config_type"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", certNameForOptionalTests),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},

			// verify datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_certificates_management_certificates", "test_certificates", acctest.Optional, acctest.Create, certificatesManagementCertificateDataSourceRepresentation) +
					compartmentIdVariableStr + CertificatesManagementCertificateResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_certificates_management_certificate", "test_certificate", acctest.Optional, acctest.Update, certificatesManagementCertificateRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "name", certNameForOptionalTests),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "certificate_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "certificate_collection.0.items.#", "1"),
				),
			},

			// verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_certificates_management_certificate", "test_certificate", acctest.Required, acctest.Create, CertificatesManagementcertificateSingularDataSourceRepresentation) +
					compartmentIdVariableStr + CertificatesManagementCertificateResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "certificate_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "certificate_profile_type"),
					resource.TestCheckResourceAttr(singularDatasourceName, "certificate_rules.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "certificate_rules.0.advance_renewal_period", "P45D"),
					resource.TestCheckResourceAttr(singularDatasourceName, "certificate_rules.0.renewal_interval", "P100D"),
					resource.TestCheckResourceAttr(singularDatasourceName, "certificate_rules.0.rule_type", "CERTIFICATE_RENEWAL_RULE"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "config_type"),
					resource.TestCheckResourceAttr(singularDatasourceName, "current_version.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "key_algorithm"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", certNameForOptionalTests),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "signature_algorithm"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttr(singularDatasourceName, "subject.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// verify resource import
			{
				Config:            config + CertificatesManagementCertificateRequiredOnlyResource,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"certificate_config",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func TestCertificatesManagementCertificateResource_imported(t *testing.T) {
	httpreplay.SetScenario("TestCertificatesManagementCertificateResource_imported")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")

	resourceName := "oci_certificates_management_certificate.imported_test_certificate"

	var resId string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			{
				Config: config + importedCertificateTestConfig(compartmentId, importedCertificatePemV1, importedPrivateKeyPemV1, "importedVersion1", "imported certificate create", "create"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "certificate_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.cert_chain_pem", importedPemStateValue(importedCertChainPem)),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.certificate_pem", importedPemStateValue(importedCertificatePemV1)),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.config_type", "IMPORTED"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.version_name", "importedVersion1"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "description", "imported certificate create"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.scenario", "create"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "key_algorithm", "RSA2048"),
					resource.TestCheckResourceAttr(resourceName, "name", importedCertName),
					resource.TestCheckResourceAttr(resourceName, "signature_algorithm", "SHA256_WITH_RSA"),
					resource.TestCheckResourceAttrSet(resourceName, "current_version.0.certificate_id"),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.serial_number", importedCertificateDetailsV1.serialNumber),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.stages.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.subject_alternative_names.#", strconv.Itoa(len(importedCertificateDetailsV1.dnsNames))),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.subject_alternative_names.0.type", "DNS"),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.subject_alternative_names.0.value", importedCertificateDetailsV1.dnsNames[0]),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.subject_alternative_names.1.type", "DNS"),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.subject_alternative_names.1.value", importedCertificateDetailsV1.dnsNames[1]),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.validity.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.validity.0.time_of_validity_not_after", importedCertificateDetailsV1.notAfter),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.validity.0.time_of_validity_not_before", importedCertificateDetailsV1.notBefore),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.version_name", "importedVersion1"),
					resource.TestCheckResourceAttrSet(resourceName, "current_version.0.version_number"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			{
				Config: config + importedCertificateTestConfig(compartmentId, importedCertificatePemV2, importedPrivateKeyPemV2, "importedVersion2", "imported certificate update", "update"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "certificate_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.cert_chain_pem", importedPemStateValue(importedCertChainPem)),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.certificate_pem", importedPemStateValue(importedCertificatePemV2)),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.config_type", "IMPORTED"),
					resource.TestCheckResourceAttr(resourceName, "certificate_config.0.version_name", "importedVersion2"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "description", "imported certificate update"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.scenario", "update"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "key_algorithm", "RSA2048"),
					resource.TestCheckResourceAttr(resourceName, "name", importedCertName),
					resource.TestCheckResourceAttr(resourceName, "signature_algorithm", "SHA256_WITH_RSA"),
					resource.TestCheckResourceAttrSet(resourceName, "current_version.0.certificate_id"),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.serial_number", importedCertificateDetailsV2.serialNumber),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.stages.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.subject_alternative_names.#", strconv.Itoa(len(importedCertificateDetailsV2.dnsNames))),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.subject_alternative_names.0.type", "DNS"),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.subject_alternative_names.0.value", importedCertificateDetailsV2.dnsNames[0]),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.subject_alternative_names.1.type", "DNS"),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.subject_alternative_names.1.value", importedCertificateDetailsV2.dnsNames[1]),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.validity.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.validity.0.time_of_validity_not_after", importedCertificateDetailsV2.notAfter),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.validity.0.time_of_validity_not_before", importedCertificateDetailsV2.notBefore),
					resource.TestCheckResourceAttr(resourceName, "current_version.0.version_name", "importedVersion2"),
					resource.TestCheckResourceAttrSet(resourceName, "current_version.0.version_number"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					func(s *terraform.State) (err error) {
						resId2, err := acctest.FromInstanceState(s, resourceName, "id")
						if err != nil {
							return err
						}
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return nil
					},
				),
			},
			{
				Config:            config + importedCertificateTestConfig(compartmentId, importedCertificatePemV2, importedPrivateKeyPemV2, "importedVersion2", "imported certificate update", "update"),
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"certificate_config",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func importedCertificateTestConfig(compartmentId, certificatePem, privateKeyPem, versionName, description, scenario string) string {
	return fmt.Sprintf(`
variable "compartment_id" {
	default = "%s"
}

resource "oci_certificates_management_certificate" "imported_test_certificate" {
	compartment_id = var.compartment_id
	name           = "%s"
	description    = "%s"

	certificate_config {
		config_type = "IMPORTED"
		certificate_pem = <<-EOT
%s
EOT
		private_key_pem = <<-EOT
%s
EOT
		cert_chain_pem = <<-EOT
%s
EOT
		version_name = "%s"
	}

	freeform_tags = {
		scenario = "%s"
	}
}
`, compartmentId, importedCertName, description, certificatePem, privateKeyPem, importedCertChainPem, versionName, scenario)
}
