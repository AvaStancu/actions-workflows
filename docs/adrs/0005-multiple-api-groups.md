# ADR 0005: Supporting multiple API groups

**Date**: 2022-10-28

**Status**: In Progress

## Context

[actions-runner-controller](https://github.com/actions-runner-controller/actions-runner-controller) currently uses `actions.summerwind.dev` as a group name for its custom resources. As we draft the migration plan of ARC’s repository to our [actions](https://github.com/actions) organization, the question of changing the group name was brought up.

`summerwind` is the handle of the original developer and maintainer of ARC. They have since handed over the maintenance of ARC to Yusuke (mumoshu). Given that GitHub is taking over the maintenance of this project moving forward, and will be adding new custom resources, there are arguments that can be made in favor of having these resources belong to: `actions.github.com` instead of `actions.summerwind.dev`.

## Decision

Following an evaluation detailed in [this document](https://docs.google.com/document/d/1pBSiuBCdx2y7RYnMmD-p2nx5EA77bD2pmXjnb_F0Ois/edit#) as well as a [spike](https://github.com/actions/dev-arc/pull/8) to test the feasibility of the proposed solution, we have decided to move forward with the following:

> Create a new group name `actions.github.com` and keep using `actions.summerwind.dev` for the existing resources.

The spike helped eliminate the technical uncertainty associated with managing 2 group names. The spike was also reviewed and approved by Yusuke (ARC’s current maintainer) without any mentionable issues or concerns.

### Pros

- Change is invisible to customers who don’t opt to use the new resources
- Existing group name will disappear anyway if we decide to deprecate the existing scaling modes & resources

### Cons

- The most technically complex of the 3 options and the least explored (we don’t know what we don’t know)
- More effort in documenting the change and creating a migration path for customers

### Other considerations & nuances

- Engineering team believes this to be a good compromise especially if we plan to deprecate the existing modes
