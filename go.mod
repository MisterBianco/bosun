module github.com/misterbianco/bosun

go 1.14

require gopkg.in/yaml.v2 v2.2.8

require (
	docker.io/go-docker v1.0.0
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/blang/semver v3.5.1+incompatible
	github.com/containerd/containerd v1.3.3 // indirect
	github.com/containerd/continuity v0.0.0-20200228182428-0f16d7a0959c // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v1.13.1
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/jhoonb/archivex v0.0.0-20180718040744-0488e4ce1681
	github.com/mitchellh/go-homedir v1.1.0
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/opencontainers/runc v0.1.1 // indirect
	github.com/sirupsen/logrus v1.4.1
	github.com/spf13/cobra v0.0.6
	github.com/spf13/viper v1.4.0
	golang.org/dl v0.0.0-20200302224518-306f3096cb2f // indirect
	golang.org/x/net v0.0.0-20190522155817-f3200d17e092
)

replace github.com/docker/docker => github.com/docker/engine v1.4.2-0.20200309214505-aa6a9891b09c
