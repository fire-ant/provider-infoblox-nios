# Provider Template

`provider-infoblox-nios` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/upbound/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Template API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/fire-ant/provider-infoblox-nios):
```
up ctp provider install fire-ant/provider-infoblox-nios:v0.1.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-infoblox-nios
spec:
  package: fire-ant/provider-infoblox-nios:v0.1.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/fire-ant/provider-infoblox-nios).

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/fire-ant/provider-infoblox-nios/issues).


Testing Notes:

1: get a curl endpoint
```
kubectl get allocation testipallocation -o jsonpath='{.status.atProvider.ref}'
```

2: get a curl pod to GET from the endpoint
```
kubectl run --image=curlimages/curl:latest -i --tty mycurlpod -- sh
```

3: run a curl test using the Status.AtProvider.Ref from the resource your
```
curl -k1 -u admin:infoblox -X GET https://nios-test.infoblox-system.svc.
cluster.local/wapi/v2.9/record:host/ZG5zLnJlY29yZDpob3N0JE5vbmU=:ip-12-34-56
-78.us-west-2.compute.internal/false
```

REST Doc: https://www.infoblox.com/wp-content/uploads/infoblox-deployment-infoblox-rest-api.pdf