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
make run ENABLE_WEBHOOKS=false
```