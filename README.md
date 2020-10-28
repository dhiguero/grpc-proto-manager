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

## TODO

- Add tests

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
