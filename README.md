# go-papermc

A papermc download api client make by GoLang

## install

Install by Go:

```
go get -u github.com/zhixuan2333/go-papermc
```

## Usage

```go
import "github.com/zhixuan2333/go-papermc"
```

```go
client := gopapermc.NewClient(nil,nil)

// Gets a list of all available projects.
projects, err := client.ListProjects()

// Gets infomastion about a project.
project, err := client.GetProject("paper")

// Gets infomation about a project version.
version, err := client.GetVersion("paper", "1.19")

// Gets all available builds for a project's version.
builds, err := client.ListVersionBuilds("paper", "1.19")

// Gets infomation related to a specific build.
build, err := client.GetBuild("paper", "1.19", "1")

// Downloads the given file from a build'd data.
file, err := client.DownloadFile("paper", "1.19", "1", "minecraft_server.1.19.jar")

// Gets infomation about a project's version group.
versionGroup, err := client.GetVersionGroup("paper", "1.18")

// Gets all available builds for a project's version group.
builds, err := client.ListVersionGroupBuilds("paper", "1.18")

```

## License

Licensed under the [MIT](./LICENSE).
