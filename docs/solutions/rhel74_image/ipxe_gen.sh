#!/bin/bash
# ipxe_gen.sh - Generate ipxe.sh script for use in building 
# IPXE boot server specific to RHEL 7.4 image creation
# 
# Steven B. Nelson - 27 Oct 2017

# Script takes input from terraform and generates all scripts needed
# for building ipxe server.  Resulting script is then base64encoded via
# terraform and run at instance provisioning time.

# Set input for stdin to capture JSON from terraform
INPUT_JSON=`cat -`

# Set location of the template to use, plus the location of the 
# source build files to put into the ipxe boot environment
IPXE_BUILD_TEMPLATE="./ipxe.sh.template"
IPXE_KS="ks.cfg"
IPXE_FWCFG="direct.xml"
IPXE_CLOUDINIT="cloud.cfg"
IPXE_SOURCE_DIR="./sources"

# Capture all the JSON info passed by the terraform into local variables
SSH_PUBLIC_KEY=`echo ${INPUT_JSON} | jq -r '.ssh_public_key'`
OCI_API_PRIVATE_KEY=`echo ${INPUT_JSON} | jq -r '.private_key_path'`
OCI_API_TENANCY=`echo ${INPUT_JSON} | jq -r '.tenancy_ocid'`
OCI_API_USER=`echo ${INPUT_JSON} | jq -r '.user_ocid'`
OCI_API_REGION=`echo ${INPUT_JSON} | jq -r '.region'`
OCI_PRKEY_PW=`echo ${INPUT_JSON} | jq -r '.private_key_password'`
OCI_BUCKET=`echo ${INPUT_JSON} | jq -r '.bucket'`
OCI_ISO_NAME=`echo ${INPUT_JSON} | jq -r '.iso_name'`
OCI_OS_SHORT_NAME=`echo ${INPUT_JSON} | jq -r '.os_short_name'`
RHEL_UNAME=`echo ${INPUT_JSON} | jq -r '.rhel_user'`
RHEL_PW=`echo ${INPUT_JSON} | jq -r '.rhel_pw'`
ZEROS_OCID=`echo ${INPUT_JSON} | jq -r '.zeros_ocid'`

# If we get a pass phrase for the private key, use it, otherwise dont
if [ -z "${OCI_PRKEY_PW}" ]
then
	OCI_API_FINGERPRINT=`openssl rsa -pubout -outform DER -in ${OCI_API_PRIVATE_KEY} 2> /dev/null | openssl md5 -c`
else
	OCI_API_FINGERPRINT=`openssl rsa -pubout -outform DER -passin pass:${OCI_PRKEY_PW} -in ${OCI_API_PRIVATE_KEY} 2> /dev/null | openssl md5 -c`
fi

# Create the head of the script and initialize our first function.  All functions are 
# simply encapsulations of uuencoded shar files.  This design pattern repeats
# for each of the files needed during build - cloud.cfg, direct.xml (firewalld), 
# ks.cfg (kickstart), and private key (OCI CLI)
cat > ./ipxe.sh <<-EOF
#!/bin/bash
function cloud {
EOF

cloud=${IPXE_SOURCE_DIR}"/"${IPXE_CLOUDINIT}
uuencode ${cloud} ${IPXE_CLOUDINIT} > ./temp.uue
shar temp.uue | sed '/^#/d' | grep -v exit >> ./ipxe.sh

cat >> ./ipxe.sh <<-EOF
uudecode ./temp.uue
rm temp.uue
}
function firewallcfg {
EOF

fw=${IPXE_SOURCE_DIR}"/"${IPXE_FWCFG}
uuencode ${fw} ${IPXE_FWCFG} > ./temp.uue
shar temp.uue | sed '/^#/d' | grep -v exit >> ./ipxe.sh

cat >> ./ipxe.sh <<-EOF
uudecode ./temp.uue
rm temp.uue
}
function ks {
EOF

ks=${IPXE_SOURCE_DIR}"/"${IPXE_KS}
uuencode ${ks} ${IPXE_KS} > ./temp.uue
shar temp.uue | sed '/^#/d' | grep -v exit >> ./ipxe.sh

cat >> ./ipxe.sh <<-EOF
uudecode ./temp.uue
rm temp.uue
}
function privkey {
EOF

uuencode ${OCI_API_PRIVATE_KEY} oci_api_key.pem > ./temp.uue
shar temp.uue | sed '/^#/d' | grep -v exit >> ./ipxe.sh

cat >> ./ipxe.sh <<-EOF
uudecode ./temp.uue
rm temp.uue
}
EOF

# Add the template file to the shell script being built
cat ${IPXE_BUILD_TEMPLATE} >> ./ipxe.sh

# Replace all the tags in the shell script with actual values
sed -i.bak 's|<PUBLIC_KEY>|\"'"${SSH_PUBLIC_KEY}"'\"|g
s|<TENANCY>|\"'"${OCI_API_TENANCY}"'\"|g
s|<USER>|\"'"${OCI_API_USER}"'\"|g
s|<FINGERPRINT>|\"'"${OCI_API_FINGERPRINT}"'\"|g
s|<REGION>|\"'"${OCI_API_REGION}"'\"|g
s|<PASSPHRASE>|\"'"${OCI_PRKEY_PW}"'\"|g
s|<BUCKET>|\"'"${OCI_BUCKET}"'\"|g
s|<ISO_NAME>|\"'"${OCI_ISO_NAME}"'\"|g
s|<OS_NAME>|'"${OCI_OS_SHORT_NAME}"'|g
s|<RHEL_UNAME>|'"${RHEL_UNAME}"'|g
s|<RHEL_PASS>|'"${RHEL_PW}"'|g 
s|<ZEROS_OCID>|\"'"${ZEROS_OCID}"'\"|g' ./ipxe.sh

chmod 755 ./ipxe.sh
rm ./temp.uue

jq -n --arg shell "./ipxe.sh" '{ "shell":$shell }'