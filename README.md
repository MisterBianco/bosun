# Bosun

![logo](assets/linkedin_banner_image_2.png)

Bosun, or boatswain, is a utility to build docker images using a yaml config file
The yaml file is quite simple at the moment. There are only a couple values required.

- name
- tags

A bosun.yml file might look like the following:

```yaml
name: app_config
tags:
  - "1.0.1-alpine"
```

The tag values MUST follow [semver](https://semver.org/) formatting.

## Install

```bash
curl -sSL https://github.com/MisterBianco/bosun/releases/download/v0.1.0/bosun_0.1.0_Darwin_x86_64.tar.gz | tar -xz
```

## Commands

### build

The command that given a path that contains a bosun.yml to build a docker image

### version

A command to list the tool version

## Questions

Can I use this in CI?

A: Yes, that is actually the main goal here.