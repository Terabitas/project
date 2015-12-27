# `project` - automate initial project setup [![Circle CI](https://circleci.com/gh/nildev/project/tree/master.svg?style=svg&circle-token=5f985f466e03e2f89df778c37ff6c6ab0b370a09)](https://circleci.com/gh/nildev/project/tree/master)

## Why

Whenever I need to setup new *Go* project I find myself copy/pasting from previous one. To solve that I created `project` tool
that allows me to prepare different project *template* repositories and generate new project based on it.

# How to use

## Install

Either download binary from [here](https://github.com/nildev/project/releases) or `go get`
```
go get https://github.com/nildev/project
```

## Project templates

Using Golang template syntax prepare your project template, for example take a look at [this one](https://github.com/nildev/prj-tpl-cli-app).
Then create `project.json` file with actual values for each variable you have defined in your project template.

For example, content of such json file could be:

```
{
  "GitRepoFullPath": "https://bitbucket.org/nildev/selis",
  "OrgPath":"bitbucket.org/nildev",
  "BinaryName" : "selisd",
  "Org": "nildev",
  "ServiceName": "selis",
  "ServicePort" : "1080",
  "ServiceVersion" : "0.1.0"
}
```

## Generate new project

- replace $PATH_TO_NEW_PROJECT with path to directory where newly generated project should be created
- replace $PATH_TO_CONFIG_JSON with path to json file with values for variables defined in template 
- replace $PATH_TO_TEMPLATE_REPO to project template git repo for example `git@github.com:nildev/prj-tpl-cli-app.git`

Run:
```
project setup --destDir=$PATH_TO_NEW_PROJECT --configFile=$PATH_TO_CONFIG_JSON --templateRepo="$PATH_TO_TEMPLATE_REPO"
for example:
project setup --destDir=$GOPATH/src/github.com/nildev/newapp --configFile=./newapp.config.json --templateRepo="git@github.com:nildev/prj-tpl-cli-app.git"
```

## Template in private repository

If template is in private repository make sure that required ssh key is added to ssh agent.

# Available templates

Here is a list of available templates.
If you have created one please do a pull request with link to that repo.

 * https://github.com/nildev/prj-tpl-cli-app (`git@github.com:nildev/prj-tpl-cli-app.git`)
 

# Building from source

1) Get repo

```
go get bitbucket.org/nildev/tools
```

2) Restore deps

```
godep restore 
```

3) Run `build`

```
./build
```

## Project Details

### Release Notes

See the [releases tab](https://github.com/nildev/project/releases) for more information on each release.

### Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details on submitting patches and contacting developers via IRC and mailing lists.

### License

project is released under the Apache 2.0 license. See the [LICENSE](LICENSE) file for details.

Specific components of project use code derivative from software distributed under other licenses; in those cases the appropriate licenses are stipulated alongside the code.