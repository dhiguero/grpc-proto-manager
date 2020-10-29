# grpc-proto-manager

Generating gRPC protos can be a critical tasks for development. However, automating this task is usually done in a non-reusable, script-based and pipelined way for CI integration. This project aims to provide an easy-to-use tool to generate a collection of protos assuming a simple structure:

```
<my_base_path>/myservice -> github.com/my_user/grpc-myservice-<target_language>
```


## Layout structure

The tool relies in a simple directory structure to create the protos.

1. Independently of whether proto definitions are stored on a mono-repo or are spread accross different repositories, the tool expects a directory with the name of high-level entity that is associated with the protos. For example, if a microservice is named `login`, or a high-level entity in your system is `user` the tool assumes you are storing protos definitions (e.g., entities.proto & services.proto) in a directory named `login` or `user` respectively.

2. Target repos must exist on your account. This may be addressed for now, but for now, your administrator should create a repo named `grpc-<high_level_entity>-<target_language>` that will store the generated code.

3. To specify the target languages use a file named `.protolang` inside each directory. Be aware that this has been tested for now for Golang, other languages may not work :).

## Can I see a working example?

Yes! For a working example, check the following repositories:

* [gpm-example](https://github.com/dhiguero/gpm-example): Simple monorepo with the layout structure expected by GPM. It contains two protobuf services: ping and agenda.
* [grpc-agenda-go](https://github.com/dhiguero/grpc-agenda-go): Repository containing the golang generated code for the agenda service.
* [grpc-ping-go](https://github.com/dhiguero/grpc-ping-go): Repository containing the golang generated code for the ping service.

```
$ ./bin/darwin/gpm generate /<full_path>/gpm-example
9:03PM INF configuration loaded path=/<full_path>/gpm-example/.gpm.yaml
9:03PM DBG Launching GPM
9:03PM INF app config commit=0465ed57eeb8a3e890c33b1c23ac8467ad1be365 version=v0.0.1
9:03PM INF Paths Project=/<full_path>/gpm-example Temp=/tmp/gpm
9:03PM INF Providers Repository=github generator=docker
9:03PM INF Defaults Language=go
9:03PM INF generated code repository URL=dhiguero
9:03PM INF processing proto directory path=/<full_path>/gpm-example/agenda
9:03PM INF publishing new version newVersion=v0.1.0 repo=grpc-agenda-go
9:03PM INF processing proto directory path=/<full_path>/gpm-example/ping
9:03PM INF publishing new version newVersion=v0.1.0 repo=grpc-ping-go
```


## Roadmap/TODO

- Add tests
- Improve parametrization
- Brew installation
- Integration with GitHub actions

## Inspiration

Organizing and managing a protobuf repository is something that may developers and companies face, some of the resources that inspired this tool are:

- https://medium.com/namely-labs/how-we-build-grpc-services-at-namely-52a3ae9e7c35
- https://medium.com/building-ibotta/building-a-scaleable-protocol-buffers-grpc-artifact-pipeline-5265c5118c9d

## License

Copyright 2020 Daniel Higuero.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
