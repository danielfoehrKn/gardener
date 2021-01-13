# Automatic Deployment of Gardenlets

The gardenlet can automatically deploy itself into shoot clusters, and register this cluster as a seed cluster. 
These clusters are called shooted seeds. 
This procedure is the preferred way to add additional seed clusters, because shoot clusters already come with production-grade qualities that are also demanded for seed clusters.

## Prerequisites

The only prerequisite is to register an initial cluster as a seed cluster that has already a gardenlet deployed:

* This gardenlet was either deployed as part of a Gardener installation using a setup tool (for example, `gardener/garden-setup`) or
* the gardenlet was deployed manually 
  - for a step-by-step manual installation Guide see: [Deploy a Gardenlet Manually](deploy_gardenlet_manually.md))
  - for a Gardenlet deployment using an installation tool see: [Gardenlet landscaper component](../../.landscaper/gardenlet/README.md).

> The initial cluster can be the garden cluster itself.

## Self-Deployment of Gardenlets in Additional Shooted Seed Clusters

For a better scalability, you usually need more seed clusters that you can create as follows:

1. Use the initial cluster as the seed cluster for other shooted seed clusters. It hosts the control planes of the other seed clusters.
1. The gardenlet deployed in the initial cluster deploys itself automatically into the shooted seed clusters.  

The advantage of this approach is that there’s only one initial gardenlet installation required. Every other shooted seed cluster has a gardenlet deployed automatically.

## Related Links

[Create Shooted Seed Cluster](../usage/shooted_seed.md)

[garden-setup](http://github.com/gardener/garden-setup)

