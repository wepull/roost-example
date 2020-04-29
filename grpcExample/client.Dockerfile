# Copyright 2019 American Express Travel Related Services Company, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
# in compliance with the License. You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software distributed under the License
# is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
# or implied. See the License for the specific language governing permissions and limitations under
# the License.

FROM golang:1.12-alpine as builder
ENV PROJECT github.com/roost-io/roost-example/grpc-k8s-health-check
WORKDIR /go/src/$PROJECT
COPY . .
WORKDIR /go/src/$PROJECT/client-grpc/
RUN go build -gcflags='-N -l' -o /app
# Make the dockerfile more optimized by using multistage dockerbuild which we copy the binary from the BUILD_STAGE container to the final container.
# The second FROM instruction starts a new build stage with the alpine image as its base.
FROM alpine:3.9
COPY --from=builder /app /app
# The ENTRYPOINT of an image specifies what executable to run when the container starts.
ENTRYPOINT ["/app"]