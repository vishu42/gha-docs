# Setup Node.js environment

Setup a Node.js environment by adding problem matchers and optionally downloading and adding it to the PATH.
## Overview
```
first line
second line
```


## Inputs

| Input | Description | Default | Required |
| -------- | -------- | -------- | ------- |
| always-auth | Set always-auth in npmrc. | false | false |
| architecture | Target architecture for Node to use. Examples: x86, x64. Will use system architecture by default. |  | false |
| cache | Used to specify a package manager for caching in the default directory. Supported values: npm, yarn, pnpm. |  | false |
| cache-dependency-path | Used to specify the path to a dependency file: package-lock.json, yarn.lock, etc. Supports wildcards or a list of file names for caching multiple dependencies. |  | false |
| check-latest | Set this option if you want the action to check for the latest available version that satisfies the version spec. | false | false |
| node-version | Version Spec of the version to use. Examples: 12.x, 10.15.1, &gt;=10.15.0. |  | false |
| node-version-file | File containing the version Spec of the version to use.  Examples: .nvmrc, .node-version, .tool-versions. |  | false |
| registry-url | Optional registry to set up for auth. Will set the registry in a project level .npmrc and .yarnrc file, and set up auth to read in from env.NODE_AUTH_TOKEN. |  | false |
| scope | Optional scope for authenticating against scoped registries. Will fall back to the repository owner when using the GitHub Packages registry (https://npm.pkg.github.com/). |  | false |
| token | Used to pull node distributions from node-versions. Since there&#39;s a default, this is typically not supplied by the user. When running this action on github.com, the default value is sufficient. When running on GHES, you can pass a personal access token for github.com if you are experiencing rate limiting. | ${{ github.server_url == &#39;https://github.com&#39; &amp;&amp; github.token || &#39;&#39; }} | false |

## Outputs

| Output | Description |
| --- | --- |
| cache-hit | A boolean value to indicate if a cache was hit. |
| node-version | The installed node version. |
