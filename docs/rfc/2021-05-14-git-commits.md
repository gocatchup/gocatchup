# Git Commits

## Introduction

Create a convention for the git commits so they always follow the same
structure.

## Background

Using a convention for commits will help generating changelogs, and
automation around it. It's also good to have consistent and clear
commits when you are searching through history and running `git blame`

## Solution

Use [conventional commits](https://www.conventionalcommits.org/). The
allowed types will be the [angular
types](https://github.com/angular/angular/blob/22b96b9/CONTRIBUTING.md#type).

For example

```shell
<type>(<scope>): <subject>
<BLANK LINE>
<body>
<BLANK LINE>
<footer>
```

```shell
fix(github): use latest github client

The version in our package.json gets copied to the one we publish, and
users need the latest of these.
```

### Enforcing the commits

We will be using [commitsar](https://commitsar.tech/docs/usage/github)
which will run on every pull request with GitHub action and will prevent
merging the pull request.
