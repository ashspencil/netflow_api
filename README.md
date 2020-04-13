# Before running

## Replace Form, in auth.yaml
```
  .dockerconfigjson: <base64-code>
```

Find base64 code by
```
cat ~/.docker/config.json | base64 -w0
```

## Replace parameter, in models/model.go
```
Community             = <password>
CISCO_NMG_2010        = <ip>
CISCO_NMG_3F_N        = <ip>
CISCO_NMG_4F_S        = <ip>
CISCO_NMG_4F_N        = <ip>

```

## Quickly run
```
dep ensure
```
```
go run cmd/main.go
```

## Kubernetes Usage

replace <external-ip> and

```
kubectl create -f deploy.yaml
```
