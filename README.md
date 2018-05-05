# kube-openapi-generator

**This is now maintained as library https://github.com/appscode/kutil/tree/release-7.0/openapi**

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

### Known Issues:
- https://github.com/kubernetes-client/gen/issues/59#issuecomment-381395227
