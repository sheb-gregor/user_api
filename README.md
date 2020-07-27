# Go Web-Service Scaffold

This is a project example build with Lancer-Kit tool set.

#### Quick start

1. Clone this repo:

```shell script
git clone https://github.com/sheb-gregor/user_api
cd user_api
```

2. Prepare a local configuration:

```shell script
## here is secrets and other env vars
cp ./env/tmpl.env ./env/local.env

## here is configuration details
cp ./env/tmpl.config.yaml ./env/local.config.yaml
```

3. Build docker image:

```shell script
make build_docker image=user_api config=local
```

4. Start all:

```shell script
docker-compose up -d
```

## Development 

- Get `forge` — a tool for code generation:

```shell script
go get -u github.com/lancer-kit/forge
```



