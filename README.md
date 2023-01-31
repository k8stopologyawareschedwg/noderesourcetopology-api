# NodeResourceTopology API

## Purpose
This repository contains the CRD based API definition used for enabling NUMA aware Scheduling in Kubernetes. 
Please refer to [this](https://docs.google.com/document/d/12kj3fK8boNuPNq) document for more details.

This repo was created to enable experimentation and for development of Topology-aware scheduling components
that rely on this API like NFD/RTE (for exposing node resource information while taking topology into consideration)
and NodeResourceTopologyMatch scheduler plugin (for taking node resource topology into consideration while making
scheduling decisions). This repo allows the ability to experiement with new features but once the API reaches
a level of stability, the long term plan is to deprecate it and move entirely to the repo under Kubernetes staging:
https://github.com/kubernetes/noderesourcetopology-api.

## Community, discussion, contribution, and support

Learn how to engage with the Kubernetes community on the [community
page](http://kubernetes.io/community/).

You can reach the maintainers of this repository at:
- Slack: [#topology-aware-scheduling](https://kubernetes.slack.com/archives/C012XSGFZQE)

### Code of Conduct

Participation in the Kubernetes community is governed by the [Kubernetes
Code of Conduct](code-of-conduct.md).

### Contibution Guidelines

See [CONTRIBUTING.md](CONTRIBUTING.md) for more information. Please note that [kubernetes/noderesourcetopology-api](https://github.com/kubernetes/noderesourcetopology-api)
is a readonly mirror repository, all development is done at [kubernetes/kubernetes](https://github.com/kubernetes/kubernetes).