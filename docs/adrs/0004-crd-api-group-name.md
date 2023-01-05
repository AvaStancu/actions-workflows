# ADR 0004: Technical detail about actions-runner-controller repository transfer
**Date**: 2022-11-04

**Status**: In Progress

# Context

As part of [[Epic] ARC Private Beta: Repository Migration & Open Sourcing Process](https://github.com/github/c2c-actions/issues/5640), we have decided to transfer the current [actions-runner-controller repository](https://github.com/actions-runner-controller/actions-runner-controller) into the [Actions org](https://github.com/actions).

**Goals:**
- A clear signal that GitHub will start taking over ARC and provide support.
- Since we are going to deprecate the existing auto-scale mode in ARC at some point, we want to have a clear separation between the legacy mode (not supported) and the new mode (supported).
- Avoid disrupting customers as much as we can, existing ARC customers will not notice any difference after the repository transfer, they can keep upgrading to the newer version of ARC and keep using the legacy mode. 

**Challenges**
- The original creator's name (`summerwind`) is all over the place, including some critical parts of ARC:
    - The k8s customer resource API's full name is `actions.summerwind.dev/v1alpha1/RunnerDeployment`, renaming it to `actions.github.com` is a breaking change and will force the customer to rebuild their entire k8s cluster. 
    - All docker images around ARC (controller + default runner) is published to [dockerhub/summerwind](https://hub.docker.com/u/summerwind)
- The helm chart for ARC is currently hosted on [GitHub pages](https://actions-runner-controller.github.io/actions-runner-controller) for https://github.com/actions-runner-controller/actions-runner-controller, moving the repository means we will break customers who install ARC via the helm chart


# Decisions

## APIs group names for k8s custom resources, `actions.summerwind` or `actions.github`

- We will not rename any existing ARC resources API name after moving the repository under Actions org. (keep `summerwind` for old stuff)
- For any new resource API we are going to add, those will be named properly under GitHub, ex: `actions.github.com/v1alpha1/AutoScalingRunnerSet`

Benefits:
- A clear separation from existing ARC:
    - Easy for the support engineer to triage income tickets and figure out whether we need to support the use case from the customer
- We won't break existing customers when they upgrade to a newer version of ARC after the repository transfer

Based on the spike done by `@nikola-jokic` https://github.com/actions/dev-arc/pull/8, we have confidence that we can host multiple resources with different API names under the same repository, and the published ARC controller can handle both resources properly.

Other considered approaches can be found in the [google doc](https://docs.google.com/document/d/1pBSiuBCdx2y7RYnMmD-p2nx5EA77bD2pmXjnb_F0Ois) as well.

## ARC Docker images

We will not start using the GitHub container registry for hosting ARC images (controller + runner images) right after the repository transfer.

But over time, we will start using GHCR for hosting those images along with our deprecation story.

## Helm chart

We will recreate the https://github.com/actions-runner-controller/actions-runner-controller repository after the repository transfer.

The recreated repository will only contain the helm chart assets which keep powering the https://actions-runner-controller.github.io/actions-runner-controller for customers to install ARC via Helm.

Long term, we will switch to hosting the helm chart on GHCR (OCI) instead of using GitHub Pages.

This will require a one-time change to our customers by running
`helm repo remove actions-runner-controller` and `helm repo add actions-runner-controller oci://ghcr.io/actions`
