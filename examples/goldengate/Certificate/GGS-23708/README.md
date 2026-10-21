# GGS-23708 reproduction

This reproduces the certificate create/read failure reported with OCI Terraform provider v9.3.0. It uses an existing GoldenGate deployment so the only Terraform-managed object is the deployment certificate.

## Prerequisites

- A valid OCI session-token profile. The example defaults to `GGSTEST`, which is the profile used by the successful deployment GET; override it with `TF_VAR_oci_profile` if needed. Session tokens expire after one hour, so refresh the session when needed.
- An existing GoldenGate deployment OCID in `us-sanjose-1`, with an identity allowed to manage certificates for that deployment.
- OpenSSL, Terraform, and the OCI CLI.
- A PEM-encoded X.509 certificate. You can use a certificate you already have, or make a temporary self-signed certificate for the repro:

```sh
cd examples/goldengate/Certificate/GGS-23708
umask 077
openssl req -x509 -newkey rsa:2048 -nodes -sha256 -days 365 \
  -keyout certificate-private-key.pem \
  -out certificate.pem \
  -subj "/CN=GGS-23708-repro"
rm certificate-private-key.pem
```

`main.tf` base64-encodes `certificate.pem` for the provider.

## Steps

1. Set the existing deployment OCID and choose a certificate key that does not already exist in that deployment:

   ```sh
   export TF_VAR_deployment_id='ocid1.goldengatedeployment.oc1..replace-with-your-deployment-ocid'
   export TF_VAR_certificate_key='GGS23708Repro'
   ```

   Before running Terraform, verify that this same profile can see the deployment and list its certificates:

   ```sh
   oci --config-file "$HOME/.oci/config" \
     --profile GGSTEST --auth security_token --region us-sanjose-1 \
     --endpoint https://goldengate-dev16.us-sanjose-1.oci.oraclecloud.com/20200407 \
     goldengate deployment get \
     --deployment-id "$TF_VAR_deployment_id"

   oci --config-file "$HOME/.oci/config" \
     --profile GGSTEST --auth security_token --region us-sanjose-1 \
     --endpoint https://goldengate-dev16.us-sanjose-1.oci.oraclecloud.com/20200407 \
     goldengate certificate list \
     --deployment-id "$TF_VAR_deployment_id" --all
   ```

   Both commands must succeed before the reproduction. If `deployment get` fails, confirm the OCID and that this profile belongs to a tenancy with access to it. If only certificate listing fails, have an OCI administrator check the deployment's **Actions → View recommended policies**; OCI returns the same `NotAuthorizedOrNotFound` response for missing access and missing resources.

2. Initialize and apply with the pinned affected provider version:

   ```sh
   terraform init
   terraform apply
   ```

3. The certificate create request should succeed, then Terraform should fail on its follow-up `GetCertificate`. The error's `Request Target` should show the certificate key in both path positions:

   ```text
   /20200407/deployments/GGS23708Repro/certificates/GGS23708Repro
   ```

4. Confirm the certificate exists despite Terraform's read error:

   ```sh
   oci goldengate certificate list \
     --deployment-id "$TF_VAR_deployment_id" \
     --all
   ```

   The list should include `GGS23708Repro` in `ACTIVE` state.

5. Run `terraform apply` a second time to observe the duplicate-key create failure, because the first apply created the certificate remotely before its read failed.

## Cleanup

Delete the remotely created certificate after collecting the result:

```sh
oci goldengate certificate delete \
  --certificate-key "$TF_VAR_certificate_key" \
  --deployment-id "$TF_VAR_deployment_id" \
  --force
```

Wait for deletion to finish and verify it no longer appears in the certificate list. The provider resource may be tainted in the local Terraform state; this repro directory uses local state, so discard its state after cleanup or remove the tainted resource from state before reusing the directory.

```sh
terraform state rm oci_golden_gate_deployment_certificate.repro
```
