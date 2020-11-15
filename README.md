# HS PROTOS

## Summary

HS protos is main repository to store protobuf schema (.proto) that will be used through out services accross hungerstation ecosystem.

## Usage
### Setup Protobuf

- Install protobuf compiler (`protoc`): `$ brew install protobuf`
- Install proto file generator for Go: `$ go get -u github.com/golang/protobuf/protoc-gen-go`

### Create Proto file

Define your proto files (`.proto`) and put them in appropriate packages.

### Generate Code

Generate for both Ruby and Go languages:
```
$ protoc path/to/file.proto --go_out=. --ruby_out=.
```

### Usage
#### Add as a git-submodule

Include this repository as a git-submodule. From your project's root directory run following commands:

```
$ git submodule add git@github.com:HungerStation/hs_protos.git lib/protos
$ git submodule init
$ git submodule update
```

#### Update proto files

When this repository is updated, you need to pull the updates manually:

- Change directory: `$ cd lib/notification`
- Pull the latest proto files: `$ git checkout master && git pull`
- Back to project's root directory: `$ cd -`
- Commit changes: `$ git add lib/notification && git commit`
