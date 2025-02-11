FROM golang:1.23.2

WORKDIR /app

COPY go.mod go.sum ./
COPY . .

ENV NODE_VERSION=20.0.0
RUN apt install -y curl;
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.0/install.sh | bash
ENV NVM_DIR=/root/.nvm

RUN . "$NVM_DIR/nvm.sh" && nvm install ${NODE_VERSION};
RUN . "$NVM_DIR/nvm.sh" && nvm use v${NODE_VERSION};
RUN . "$NVM_DIR/nvm.sh" && nvm alias default v${NODE_VERSION};
ENV PATH="/root/.nvm/versions/node/v${NODE_VERSION}/bin/:${PATH}"

RUN npm install -g pnpm;
RUN pnpm install;

RUN pnpm tailwind:build;
RUN go mod download;
RUN go mod tidy;
RUN go build -o /app/main

FROM golang:1.23.2

WORKDIR /app

COPY go.mod go.sum ./
COPY . .

ENV NODE_VERSION=20.0.0
RUN apt install -y curl;
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.0/install.sh | bash
ENV NVM_DIR=/root/.nvm

RUN . "$NVM_DIR/nvm.sh" && nvm install ${NODE_VERSION};
RUN . "$NVM_DIR/nvm.sh" && nvm use v${NODE_VERSION};
RUN . "$NVM_DIR/nvm.sh" && nvm alias default v${NODE_VERSION};
ENV PATH="/root/.nvm/versions/node/v${NODE_VERSION}/bin/:${PATH}"

RUN npm install -g pnpm;
RUN pnpm install;

RUN pnpm tailwind:build;
RUN go mod download;
RUN go mod tidy;
RUN go build -o /app/main
