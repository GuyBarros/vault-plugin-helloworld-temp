#!/bin/bash

set -ex

# Setup
TMPDIR=~/devhelloworldvault
mkdir -p $TMPDIR
export VAULT_ADDR=http://localhost:8200
export VAULT_TOKEN=devroot

# Cleanup old instance
if kill `ps axuw |grep 'vault server -dev' |awk '{print $2}'`; then
    sleep 1
fi
rm -f $TMPDIR/vault.log

cat >$TMPDIR/config.hcl <<EOF
plugin_directory = "$PWD/bin"
EOF

# Start server
vault server -dev -dev-root-token-id=$VAULT_TOKEN -config=$TMPDIR/config.hcl -log-level=trace 2> $TMPDIR/vault.log &
vault write sys/plugins/catalog/hello-world \
    sha_256="$(shasum -a 256 $PWD/bin/vault-plugin-hello-world| cut -d " " -f1)" \
    command="vault-plugin-hello-world"
vault secrets enable --plugin-name='hello-world' --path="hello" plugin
