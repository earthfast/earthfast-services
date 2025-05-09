# this gets built by docker-compose.base.yml
FROM earthfast-contracts:latest

# Install the earthfast-cli build dependencies (for node-gyp, node-usb, etc)
RUN apt-get update && apt-get install -y libusb-1.0 libudev-dev libsecret-1-dev

# Set branch name for earthfast-cli
ARG CLI_BRANCH_NAME
ENV CLI_BRANCH_NAME=${CLI_BRANCH_NAME:-main}

# Install earthfast-cli
# The ADD command acts as a Docker build cache buster whenever HEAD is updated
ADD https://api.github.com/repos/earthfast/earthfast-cli/git/refs/heads/${CLI_BRANCH_NAME} version.json

RUN git clone --depth=1 -b ${CLI_BRANCH_NAME} https://github.com/earthfast/earthfast-cli /earthfast-cli
WORKDIR /earthfast-cli
RUN npm i
RUN npm run build
RUN npm link

WORKDIR /earthfast-contracts
ENTRYPOINT ["npx", "hardhat", "node", "--no-deploy"]
