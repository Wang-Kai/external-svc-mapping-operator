# external-svc-mapping-operator

## How to use kubebuilder ?

```
# 初始化脚手架
kubebuilder init  --domain sf.ucloud.cn

# add API
kubebuilder create api --group mapping --version v1 --kind ExternalSvc
```

## How to run it ?

```console
# install CRD
make install

# run program
make run ENABLE_WEBHOOKS=false

# create ExternalSvc resource
kubectl create -f config/samples/mapping_v1_externalsvc.yaml
```