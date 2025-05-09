##
## Builder: install earthfast-cli
##
FROM node:16 AS builder

# Install the earthfast-cli build dependencies (for node-gyp, node-usb, etc)
RUN apt-get update && apt-get install -y libusb-1.0 libudev-dev libsecret-1-dev

# Install earthfast-cli at HEAD
# The ADD command acts as a Docker build cache buster whenever HEAD is updated
ADD https://api.github.com/repos/earthfast/earthfast-cli/git/refs/heads/main version.json
RUN git clone --depth=1 https://github.com/earthfast/earthfast-cli.git /earthfast-cli
WORKDIR /earthfast-cli
RUN npm i
RUN npm run build
RUN npm link

##
## hello-world
##
FROM builder AS hello-world

WORKDIR /build

COPY hello-world .

RUN earthfast bundle create hello-world.tgz public

##
## hello-world-react
##
FROM builder AS hello-world-react

WORKDIR /build

COPY hello-world-react .

RUN npm install
RUN npm run build
RUN earthfast bundle create hello-world-react.tgz build

##
## The final container is a static file server hosting every site,
## along with a manifest.json file containing each site's checksum.
##
FROM node:16-alpine

WORKDIR /www
RUN npm install http-server -g

COPY --from=hello-world /build/hello-world.tgz .
COPY --from=hello-world-react /build/hello-world-react.tgz .

WORKDIR /
COPY manifest-gen.js .
RUN node manifest-gen.js www > /tmp/manifest.json
RUN mv /tmp/manifest.json www

WORKDIR /www
ENTRYPOINT ["http-server", "-c-1"]
