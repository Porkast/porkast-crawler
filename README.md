# Pokast Crawler

## Overview

Pokast Crawler is a web scraping tool designed to parse popular podcast websites, extract podcast information from RSS links, and keep the RSS sources up to date.

## Features

- **Web Scraping:** Extracts podcast information from popular podcast websites.
- **RSS Parsing:** Parses RSS feeds to retrieve relevant podcast details.
- **Data Storage:** Stores podcast information in a [database/file format].
- **Automated Updates:** Periodically updates the RSS sources to ensure the latest information.


## GoFrame 

This project base on [GoFrame](https://github.com/gogf/gf), Project Makefile Commands: 
- `make cli`: Install or Update to the latest GoFrame CLI tool.
- `make dao`: Generate go files for `Entity/DAO/DO` according to the configuration file from `hack` folder.
- `make service`: Parse `logic` folder to generate interface go files into `service` folder.
- `make image TAG=xxx`: Run `docker build` to build image according `manifest/docker`.
- `make image.push TAG=xxx`: Run `docker build` and `docker push` to build and push image according `manifest/docker`.
- `make deploy TAG=xxx`: Run `kustomize build` to build and deploy deployment to kubernetes server group according `manifest/deploy`.
