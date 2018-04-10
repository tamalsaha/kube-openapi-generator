# kube-openapi-generator

## Generate swagger.json

```console
$ go install -v
$ kube-openapi-generator > swagger.json
```

## Install swagger-codegen

From https://github.com/swagger-api/swagger-codegen#getting-started :

```console

cd hack
git clone https://github.com/swagger-api/swagger-codegen
cd swagger-codegen
mvn clean package
cd ..
```

## Generate Java client

```console
java -jar ./hack/swagger-codegen/modules/swagger-codegen-cli/target/swagger-codegen-cli.jar generate \
  -i swagger.json \
  -l java \
  -o java-client
```

Known Issues:
https://github.com/kubernetes-client/gen/issues/59
