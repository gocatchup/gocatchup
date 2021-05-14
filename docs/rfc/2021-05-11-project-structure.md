# Project Structure

## Introduction

Decide on the structure of the application and the repository. Where the
code is going to be hosted, structured, and built.

### Background

Every time a project starts it doesn't usually start with the technical
details on how the project structure is going to look like and the
design philosophy that it follows which usually results in either a
messy code base or unspoken best practices. One could argue the
designing something so early might also lead us to a project that is
structure in a way that didn't have the full understanding of the
problem which is a valid argument.

### Goals

Have a good understanding of how to structure the project and what tools
are we going to use to build `gocatchup`.

## Solutions

### Code hosting and CI

Solution: **GitHub**

Currently, I work at GitLab, but I'm choosing GitHub and its surrounding
tools to host this code base and for CI/CD to get a better understanding
of the tools outside of GitLab, to prevent me from getting pigeonholed
on what tools I use to deploy code, the plan is to more up to date with
the industry.

### Repsitory

Soltuion: **Monorepo**

The repository is going to be a mono repository where every service is
inside this inside of this repository. The reason for this is to have
everything in 1 place, all services, build tools, and frontend code. The
only exception to this rule is going to be deployments. If possible
infrastructure and deployments are going to be under a different
repository for security reasons. Infrastructure needs to be separate
because it might need a different permissions model and a different set of
tools.

#### Repository structure

Solution: **Packes as layers**

There a lot of ways to struct a Go application. We will try the
[Packages as layers](https://www.gobeyond.dev/packages-as-layers/)
an approach which is also explained in detail in [Standard Package
Layout](https://www.gobeyond.dev/standard-package-layout/). The root of
the repository will hold all the business logic and domain models. I'm
still not certain about this structure because it seems a bit flat and
going against the monorepo approach. However, we'd need to experiment with
this to be sure if it's a bad idea or not.

### Deployments

No infrastructure was chosen yet but running it should be fairly cheap
and ideally, no servers should be managed. Using something like [Cloud
Run](https://cloud.google.com/run/) or [fly.io](https://fly.io/) is
ideal.

Deployment shouldn't take long and should be automated. The mean time of
merge to deployment should be less than 15 minutes.
