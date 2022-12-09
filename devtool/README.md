<!--
title: CLI plugin
weight: 5020
pre: "<i class=\"fa fa-terminal\" aria-hidden=\"true\"></i> "
-->

# FLOGO CLI PLUGIN

## Installation
### Prerequisites
Flogo ClI version v0.9.0 or above

### Install the Plugin
To install this plugin 

```
flogo plugin install github.com/project-flogo/cli-plugins/devtool
```

## Usage
You can create a generate a sample Trigger/Activity/Action using the plugin:

```
flogo dev gen-activity myActivity
```
or 

```
flogo dev gen-trigger
```

You can also use this plugin to create descriptor JSON from the metadata.
In same location as your metadata.go enter the following command
```
flogo dev sync-metadata
```
