## Development Setup

### Requirements

* [Go 1.18](https://golang.org/dl/)
* [Docker](https://docker.com)

### Setup

1. Change into the `inexorable-node/dev` directory:

```sh
cd dev
```

2. Create `inexorable-node/dev/env/dev` by copying the template:

```sh
cp env/dev.tmpl env/dev
```

3. Install dnsmasq and setup local self-signed SSL certificate

The local system uses dnsmasq to route a wildcard subdomain locally to localhost (because /etc/hosts does not support wildcards). The domain is `armadalocal.test`.

An nginx container paired with dnsmasq routes all requests to content and domain nodes to the proper container and provides SSL termination with a locally provisioned, self-signed certificate. Here's a full example:

The system stars with all content nodes given `dev-content-n-1.armadalocal.test` urls registered on chain. Projects will be given <project-name>.armadalocal.test as they're created.

User creates a project hosted on `<project-id>.armadalocal.test`. The domain to project mapping gets pushed to a file which the domain node can download and update. When the user goes to the browser and loads this url, the domain node returns the service worker code with a `dev-content-n-1.armadalocal.test` content node. Since everything is https end to end it all works

Instructions to install dnsmasq, create ssl cert and install on your local system. Adapted from https://gist.github.com/ogrrd/5831371?permalink_comment_id=4790334#gistcomment-4790334

```sh
brew install dnsmasq
export LOCAL_CUSTOM_TLD="test"
sudo echo "address=/.${LOCAL_CUSTOM_TLD}/127.0.0.1" >> $(brew --prefix)/etc/dnsmasq.conf
cat <<EOF | sudo tee /etc/resolver/${LOCAL_CUSTOM_TLD}
nameserver 127.0.0.1
EOF
sudo brew services restart dnsmasq
sudo killall -HUP mDNSResponder
scutil --dns
```

Run this from the `earthfast-services/dev/ssl_key` folder

```sh
# `$(brew --prefix)/etc/dnsmasq.conf` should contain `address=/.test/127.0.0.1`
# `cat /etc/resolver/test` should contain `nameserver 127.0.0.1`

# generate ssl cert for local
openssl req -x509 -out armadalocal.crt -keyout armadalocal.key -newkey rsa:2048 -nodes -sha256 -days 3650 -subj '/CN=armadalocal.test' -extensions EXT -config <( \
printf "[dn]\nCN=armadalocal.test\n[req]\ndistinguished_name = dn\n[EXT]\nsubjectAltName=DNS:armadalocal.test,DNS:*.armadalocal.test\nkeyUsage=digitalSignature\nextendedKeyUsage=serverAuth")
```

To install this certificate on mac, open Keychain Access, go to System under System Keychains on the left. Go to the Certificates tab and import `armadalocal.crt`. Double click to edit toggle the "Trust" dropdown and select "Trust Always". This means your browser will not throw self signed certificate errors locally
/l

4. Start the dev environment:

```sh
# optional params to configure do script options
# export ENVIRONMENT=dev
# export NUM_CONTENT_NODES=5
./do start
```
