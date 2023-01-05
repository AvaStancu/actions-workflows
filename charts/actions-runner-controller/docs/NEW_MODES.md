# Notes on setting up ARC via OCI

## Packaging new Helm Charts

Until the repository migration is complete, we will manually release our Charts via the 'Publish New Helm Chart to OCI' Action.

## Installing our new ARC modes

This credential setup step can be removed once our new packages are public.
Until then, we need a Docker secret in order to allow the k8s cluster to pull the container image.

```

GH_USER=<your username on GH>
GH_PAT=<a PAT for your account with read:packages scope>
AUTH=$(printf %s:%s ${GH_USER} ${GH_PAT} | base64)
echo "{\"auths\":{\"ghcr.io\":{\"auth\":\"${AUTH}\"}}}" | kubectl create secret docker-registry regcred \
  --from-file=.dockerconfigjson=/dev/stdin -n actions-runner-system
```

```
helm upgrade --install \                                                                
  --namespace actions-runner-system --create-namespace \
  actions-runner-controller \
  oci://ghcr.io/actions/actions-runner-controller-helm-chart-2/actions-runner-controller
```

## Uninstallation steps

You can also remove our controller, but the CRD **will** remain.
Info at https://helm.sh/docs/chart_best_practices/custom_resource_definitions/.

```
helm uninstall actions-runner-controller-2 --namespace actions-runner-system
```