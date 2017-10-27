#!/usr/local/bin/python2.7

import os
import ConfigParser
import argparse

envList = {
    "tenancy": "tenancy_ocid",
    "fingerprint": "fingerprint",
    "key_file": "private_key_path",
    "user": "user_ocid",
    "pass_phrase": "private_key_password",
    "public_key": "ssh_public_key",
    "region": "region",
    "profile": "profile"
}

argparser = argparse.ArgumentParser()
argparser.add_argument('--config_path', required=False,
                       default='~/.oraclebmc/config', dest="config_path")
argparser.add_argument('--profile', required=False,
                       default='DEFAULT', dest="profile")
argparser.add_argument('--dest_file', required=False,
                       default='./tfenv.sh', dest='dest_file')
argparser.add_argument('--public_key', required=False,
                       default="~/.ssh/id_rsa.pub", dest='public_key')
args = vars(argparser.parse_args())

outputEnv = {}
outputFile = os.open(args["dest_file"], os.O_RDWR |
                     os.O_TRUNC | os.O_CREAT, 0755)

if len(args["config_path"]) <= 0:
    obmcsConfig = "~/.oraclebmc/config"
else:
    obmcsConfig = args["config_path"]

obmcsConfig = os.path.expanduser(obmcsConfig)

parser = ConfigParser.ConfigParser()
parser.read(obmcsConfig)

if (args["profile"] in parser.sections()):
    for options in parser.options(args["profile"]):
        outputEnv[envList[options]] = parser.get(
            args["profile"], options)
elif args["profile"] == "DEFAULT":
    for options in parser.defaults().keys():
        outputEnv[envList[options]] = parser.defaults()[options]

if "~" in outputEnv["private_key_path"]:
    outputEnv["private_key_path"] = \
        os.path.expanduser(outputEnv["private_key_path"])

publicKeyFile = os.path.expanduser(args['public_key'])
if (os.path.exists(publicKeyFile)):
    pkFile = os.open(publicKeyFile, os.O_RDONLY)
    outputEnv["ssh_public_key"] = os.read(pkFile, 512).rstrip()
    os.close(pkFile)

outputEnv["profile"] = args["profile"]

os.write(outputFile, "#!/bin/bash\n")
for key in envList.keys():
    if envList[key] not in outputEnv:
        outputEnv[envList[key]] = ""
    os.write(outputFile, "export TF_VAR_{0}=\"{1}\"\n".format(
        envList[key], outputEnv[envList[key]]))

os.close(outputFile)
