![cvescan](img/cvescan-readme-logo.png)

[![](https://img.shields.io/badge/Status-BETA-yellow)](CONTRIBUTING.md)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/cvescan/cvescan) 
[![Go Report Card](https://goreportcard.com/badge/github.com/cvescan/cvescan)](https://goreportcard.com/report/github.com/cvescan/cvescan) 
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/6409/badge)](https://bestpractices.coreinfrastructure.org/projects/6409)
[![codecov](https://codecov.io/gh/cvescan/cvescan/branch/main/graph/badge.svg?token=P9WBOBQTOB)](https://codecov.io/gh/cvescan/cvescan) 
[![SBOM](https://img.shields.io/badge/CyloneDX-SBOM-informational)](sbom/cvescan.cyclonedx.json)


```cvescan``` is an application that scans SBOMs for security vulnerabilities.

## Overview

So you've asked a vendor for an Software Bill of Materials (SBOM) for one of their closed source products, and they provided one to you in a JSON file... now what?

The first thing you're going to want to do is see if any of the components listed inside the SBOM have security vulnerabilities, and what kind of licenses these components have. This will help you identify what kind of risk you will be taking on by using the product. 

Finding security vulnerabilities and license information for components identified in a SBOM is exactly what ```cvescan``` is meant to do. ```cvescan``` can read any JSON or XML based [CycloneDX](https://cyclonedx.org) format, or a JSON [SPDX](https://spdx.dev) or [Syft](https://github.com/anchore/syft) formatted SBOM, and tell you pretty quickly if there are any vulnerabilities. 

### Open vs. Closed Source

Software can either be open or closed source. You can look at third party components you'll find in Github, or any public source repository as open source. Technically, the software you create internally at your own company is open source as well - it's not public, but your internal teams can see it. Closed source software can also be internal, but usually this is software that you purchase from external vendors. 

Companies can use SCA tools provided by vendors such as Github, Sonatype, Snyk, etc. to scan any kind of open source and provide vulnerability data - and even generate SBOMs in some cases. What they can't do (yet...) is scan closed source software that you don't have visibility into. This is where SBOMs and ```cvescan``` come into play. SBOMs provide the composition of software that you can't access, and ```cvescan``` determines if anything in the SBOM has vulnerabilities.

### Purpose

We created ```cvescan``` to scan the closed source SBOMs that are provided when you receive them from vendors. It can scan open source SBOMs too, and technically you could use ```cvescan``` as an open source SCA tool if you wanted to.


### Supported SBOM formats

There are quite a few SBOM formats available today. ```cvescan``` supports the following:

- [SPDX](https://spdx.dev)
- [CycloneDX](https://cyclonedx.org)
- [Syft](https://github.com/anchore/syft)

## Providers

![](img/providers/banner.png)

```cvescan``` supports multiple sources for vulnerability information. We call these *providers*. Currently, ```cvescan``` uses [OSV](doc/providers/osv.md) as the *default* provider, but you can also use the [Sonatype OSS Index](doc/providers/ossindex.md), or [Snyk](doc/providers/snyk.md).

At this time, please note that [OSV](doc/providers/osv.md) is free and does not require any credentials to use, [Sonatype OSS Index](doc/providers/ossindex.md) is free but requires you to register and obtain a token, and [Snyk](doc/providers/snyk.md) support requires a Snyk license.

In addition to data ```cvescan``` collects from Providers, it also [enriches](#data-enrichment) vulnerability data with extra information such as exploitation probabilities.

### Provider Support

Please note that *each provider supports different ecosystems*, so if you're not seeing any vulnerabilities in one, try another. An ecosystem is simply the package manager, or type of package. Examples include rpm, npm, gems, etc. It is important to understand that each provider may report different vulnerabilities. If in doubt, look at a few of them.

If ```cvescan``` does not find any vulnerabilities, it doesn't mean that there aren't any. All it means is that the provider being used didn't detect any, or it doesn't support the ecosystem. Some providers have vulnerabilities that come back with no Severity information. In this case, the Severity will be listed as "UNDEFINED"

### Provider Documentation

Provider documentation for ```cvescan``` can be found:

* [OSV](doc/providers/osv.md)
* [OSSINDEX](doc/providers/ossindex.md)
* [Snyk](doc/providers/snyk.md)

## Installation

### Mac

You can use [Homebrew](https://brew.sh) to install ```cvescan``` using the following:

``` bash
brew tap devops-kung-fu/homebrew-tap
brew install devops-kung-fu/homebrew-tap/cvescan
```

If you do not have Homebrew, you can still [download the latest release](https://github.com/cvescan/cvescan/releases) (ex: ```cvescan_0.4.1_darwin_all.tar.gz```), extract the files from the archive, and use the ```cvescan``` binary.  

If you wish, you can move the ```cvescan``` binary to your ```/usr/local/bin``` directory or anywhere on your path.

### Linux

To install ```cvescan```,  [download the latest release](https://github.com/cvescan/cvescan/releases) for your platform and install locally. For example, install ```cvescan``` on Ubuntu:

```bash
dpkg -i cvescan_0.4.1_linux_arm64.deb
```

## Using cvescan

You can scan either an entire folder of SBOMs or an individual SBOM with ```cvescan```.  ```cvescan``` doesn't care if you have multiple formats in a single folder. It'll sort everything out for you.

Note that the default output for ```cvescan``` is to STDOUT. Options to output in HTML or JSON are described later in this document.

### Single SBOM scan

``` bash
# Using OSV (the default provider) which does not require any credentials
cvescan scan cyclonedx.sbom.json

# Using a provider that requires credentials (ossindex)
cvescan scan --provider=xxx --username=xxx --token=xxx spdx-sbom.json
```
If the provider finds vulnerabilities you'll see an output similar to the following:

![](img/cvescan-example.png)

If the provider doesn't return any vulnerabilities you'll see something like the following:

![](img/cvescan-example-novulns.png)

### Entire folder scan

This is good for when you receive multiple SBOMs from a vendor for the same product. Or, maybe you want to find out what vulnerabilities you have in your entire organization. A folder scan will find all components, de-duplicate them, and then scan them for vulnerabilities.

```bash
# scan a folder of SBOMs (the following command will scan a folder in your current folder named "sboms")
cvescan scan --provider=xxx --username=xxx --token=xxx ./sboms
```

You'll see a similar result to what a Single SBOM scan will provide.

## Output Formats

```cvescan``` outputs data into three useful formats. By default, output is rendered to the command line. For enhanced reporting, you can output to HTML using the ```--output=html``` flag. To output to JSON, utilize the ```--output=json``` flag.

### HTML Output

If you would like a readable report generated with detailed vulnerability information, you can utilized the ```--output``` flag to save a report to an HTML file.

Example command:

``` bash
cvescan scan bad-bom.json --output=html
```

This will save a file in your current folder in the format "YYYY-MM-DD-HH-MM-SS-cvescan-results.html". If you open this file in a web browser, you'll see output like the following:

![](img/cvescan-html.png)

### JSON Output

```cvescan``` can output vulnerability data in JSON format using the ```--output``` flag. The default output is to STDOUT. There is a ton of more information in the JSON output than what gets displayed in the terminal. You'll be able to see a package description and what it's purpose is, what the vulnerability name is, a summary of the vulnerability, and more.

![](img/cvescan-json.png)

Example command:

``` bash
cvescan scan bad-bom.json --output=json > filename.json
```

## Ignoring Vulnerabilities

If needed, you can use the ```--ignore-file``` flag to load a list of CVEs to ignore in the vulnerability output. This list needs to be in a specific format where each CVE to ignore is entered on a separate line similar to the following:

```
CVE-2022-31163
CVE-2022-23520
```

There is an example ```cvescan.ignore``` file [here](./_TESTDATA_/ignore/cvescan.ignore)

To use the ```cvescan.ignore``` file, use the syntax as follows:

``` bash
cvescan --ignore-file=cvescan.ignore scan bom.json
```

## Data Enrichment

```cvescan``` has the ability to enrich vulnerability data it obtains from the [Providers](#providers). The first "enricher" we have implemented for is for [EPSS](https://www.first.org/epss/)

### Exploit Prediction Scoring System (EPSS)

[EPSS](https://www.first.org/epss/) stands for Exploit Prediction Scoring System and is framework that predicts the probability of a vulnerability being exploited. [EPSS](https://www.first.org/epss/) is often used to help in identifying high risk vulnerabilities to prioritize for remediation. 

[EPSS](https://www.first.org/epss/) uses a percentage for probability. So if you see 94, the score is that is trying to say that vulnerability has a 94% probability of exploitation.  And it stands to reason that a vulnerability with a score like 94, is something that deserves immediate attention, where a vulnerability with a score of like say 20 deserves to take a lower priority.

## Advanced stuff

If you wish, you can set two environment variables to store your credentials, and not have to type them on the command line. Check out the [Environment Variables](####Environment-Variables) information later in this README.

### Scanning SBOMs from STDIN

If you're using ```cvescan``` in your CI/CD pipelines, you can do an all in one command with Syft to generate and scan a SBOM for vulnerabilities. To do this, you can do something like the following command:

``` bash
# Make sure you include the - character at the end of the command. This triggers cvescan to read from STDIN
syft packages . -o cyclonedx-json | cvescan scan --provider ossindex --output json - 
```
This command creates a SBOM, pipes it into cvescan, and generates results in JSON format.

### Environment Variables

If you don't want to enter credentials all the time, you can add the following to your ```.bashrc``` or ```.bash_profile```

``` bash
export cvescan_PROVIDER_USERNAME={{your OSS Index user name}}
export cvescan_PROVIDER_TOKEN={{your OSS Index API Token}}
```

### Messing around

If you want to kick the tires on ```cvescan``` you'll find a selection of test SBOMs in the [test](sbom/test/) folder. 

## Notes

- It's pretty rare to see SBOMs with license information. Most of the time, the generators like Syft need a flag like ```--license```. If you need license info, make sure you ask for it with the SBOM.
- OSV. It's great, but the API is also wonky. They have a batch endpoint that would make it a ton quicker to get information back, but at the time of writing it doesn't work as expected. ```cvescan``` needs to send one PURL at a time to get vulnerabilities back, so in a big SBOM it will take some time. We'll keep an eye on that.

## Contributing

If you would like to contribute to the development of ```cvescan``` please refer to the [CONTRIBUTING.md](CONTRIBUTING.md) file in this repository. Please read the [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) file before contributing.

## Software Bill of Materials

```cvescan``` uses Syft to generate a Software Bill of Materials every time a developer commits code to this repository (as long as [Hookz](https://github.com/devops-kung-fu/hookz) is being used and is has been initialized in the working directory). More information for CycloneDX is available [here](https://cyclonedx.org).

The current CycloneDX SBOM for ```cvescan``` is available [here](./sbom/cvescan.cyclonedx.json).

## Sponsors

Thank you to the sponsors and supporters of ```cvescan```

![](img/sponsors/zero-logo.png)

## Credits

A big thank-you to our friends at [ZERO](https://zero.health) for the ```cvescan``` logo.

Thank you to [Sonatype](https://sonatype.com) for providing a wicked tool like the [Sonatype OSS Index](https://ossindex.sonatype.org).

Many thanks to our friends and fellow ```cvescan``` contributors at [Snyk](https://snyk.io) for creating a provider, and coding up processing a SBOM from STDIN. You guys rock.

EPSS description comes from the team at [Nucleus](https://nucleussec.com/blog/what-is-epss/). Thank you!