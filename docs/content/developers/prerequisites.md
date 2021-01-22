# Development Prerequisites
This document provides an overview of the tools and packages needed to work on and to build onos-config.
Developers are expected to have these tools installed on the machine where the project is built.

## Git tools
Git needs to be installed as the project is versioned on GitHub.

Set your username: 
```bash
git config --global user.name "FIRST_NAME LAST_NAME"
```
Set your email address: 
```bash
git config --global user.email "MY_NAME@example.com"
```

An SSH key is necessary to work with Github through the SSH protocol.

If you do not already have on in youe local `~/.ssh` folder then generate an SSH key with:
```bash
ssh-keygen -t ecdsa
```
> It is easiest to leave the passphrase blank.

Once created the public part of this file can be viewed with:
```bash
cat ~/.ssh/id_ecdsa.pub
```

The contents of this should be copied over to Github as a new key in **SSH and GPG keys**

https://github.com/settings/keys

### Dealing with private Go repos in Github
For private Go repos in Github, the default access protocol is HTTPS. It is necessary
to add the following text to your `~/.gitconfig` file to change this to SSH:
```
[url "ssh://git@github.com/"]
insteadOf = https://github.com/
```

See https://golang.org/doc/faq#git_https for more details


## Go Tools
Since the project is authored mainly in the Go programming language, the project requires [Go tools] 
in order to build and execute the code.

## Go Linters
[golangci-lint] is required to validate that the Go source code complies with the established style 
guidelines. To install the tool, use this command:
```bash
curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin latest
```
> The directory ~/go/bin needs to be present in your $PATH environment variable
>for this to work after being installed 

## Docker
[Docker] is required to build the project Docker images and also to compile `*.proto` files into Go source files.

## Local kubernetes environment
Some form of local kubernetes development environment is also needed.
The core team uses [Kind], but there are other options such as [Minikube] and [MicroK8s].
> Some docker containers may need access to privileged rights e.g. `onos-config:debug`,
> `onos-topo:debug` and `opennetworking/mn-stratum` and so may not be suited to
> Kubernetes environments that cannot grant these rights

## Python 3
Python 3 needs to be installed to run the license checking tool in many on the Makefiles.
> The version provided by your OS will usually be sufficient

Verify it is installed with
```bash
python3 --version
```  

## IDE
Some form of an integrated development environment is also recommended.
The core team uses the [GoLand IDE] from JetBrains, but there are many other options. 
Microsoft's [Visual Studio Code] is one such option and is available as a free download.

Note that when using [GoLand IDE] you should enable integration with Go modules in `Preferences -> Go -> Go Modules`.

## License
The project requires that all Go source files are properly annotated using the Apache 2.0 License.
Since this requirement is enforced by the CI process, it is strongly recommended that developers
setup their IDE to include the [license text](https://github.com/onosproject/onos-config/blob/master/build/licensing/boilerplate.go.txt)
automatically.

[GoLand IDE can be easily setup to do this](license_goland.md) and other IDEs will have a similar mechanism.


[Go tools]: https://golang.org/doc/install
[golangci-lint]: https://github.com/golangci/golangci-lint
[Docker]: https://docs.docker.com/install/
[Kind]: https://github.com/kubernetes-sigs/kind
[Minikube]: https://kubernetes.io/docs/tasks/tools/install-minikube/
[MicroK8s]: https://microk8s.io/
[GoLand IDE]: https://www.jetbrains.com/go/
[Visual Studio Code]: https://code.visualstudio.com
