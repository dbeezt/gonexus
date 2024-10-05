# gonexus [![DepShield Badge](https://depshield.sonatype.org/badges/sonatype-nexus-community/gonexus/depshield.svg)](https://depshield.github.io) [![CircleCI](https://circleci.com/gh/sonatype-nexus-community/gonexus.svg?style=svg)](https://circleci.com/gh/sonatype-nexus-community/gonexus)

Provides a go client library for connecting to, and interacting with, [Sonatype](//www.sonatype.com) Nexus applications such as Nexus Repository Manager and Nexus IQ Server.

## Organization of this library

The library is broken into two packages. One for each application.

### nexusrm [![GoDoc](http://godoc.org/github.com/sonatype-nexus-community/gonexus/rm?status.png)](http://godoc.org/github.com/sonatype-nexus-community/gonexus/rm) [![nexusrm coverage](https://gocover.io/_badge/github.com/sonatype-nexus-community/gonexus/rm?0 "nexusrm coverage")](http://gocover.io/github.com/sonatype-nexus-community/gonexus/rm)

Create a connection to an instance of Nexus Repository Manager

```go
// import "github.com/sonatype-nexus-community/gonexus/rm"
rm, err := nexusrm.New("http://localhost:8081", "username", "password")
if err != nil {
    panic(err)
}
```

#### Supported RM Endpoints

| Endpoint                                                                                  |         Status         | Min RM Version |
| ----------------------------------------------------------------------------------------- | :--------------------: | :------------: |
| [Assets](https://help.sonatype.com/en/assets-api.html)                                    |      :full_moon:       |                |
| [Blob Store](https://help.sonatype.com/en/blob-store-api.html)                            |       :new_moon:       |      3.19      |
| [Cleanup Policies](https://help.sonatype.com/en/cleanup-policies-api.html)                |       :new_moon:       |      3.70      |
| [Components](https://help.sonatype.com/en/components-api.html)                            | :waning_gibbous_moon:  |                |
| [Content Selectors](https://help.sonatype.com/en/content-selectors.html)                  |       :new_moon:       |      3.19      |
| [Email](https://help.sonatype.com/en/email-api.html)                                      |       :new_moon:       |      3.19      |
| [HTTP Configuration](https://help.sonatype.com/en/http-configuration-api.html)            |       :new_moon:       |                |
| [IQ Server](?)                                                                            |       :new_moon:       |      3.19      |
| [Licensing](https://help.sonatype.com/en/licensing-api.html)                              |       :new_moon:       |      3.19      |
| [Lifecycle](https://help.sonatype.com/en/lifecycle-api.html)                              |       :new_moon:       |                |
| [Maintenance](https://help.sonatype.com/en/maintenance-api.html) _pro_                    | :waning_crescent_moon: |                |
| [Nodes](https://help.sonatype.com/en/nodes-api.html) _pro_                                |       :new_moon:       |                |
| [Read-Only](https://help.sonatype.com/en/read-only-api.html)                              |      :full_moon:       |                |
| [Repositories](https://help.sonatype.com/en/repositories-api.html)                        |      :full_moon:       |                |
| [Repository Firewall](https://help.sonatype.com/en/sonatype-repository-firewall-api.html) |       :new_moon:       |                |
| [Routing Rules](https://help.sonatype.com/en/routing-rules.html)                          |       :new_moon:       |      3.17      |
| [Search](https://help.sonatype.com/en/search-api.html)                                    | :waning_gibbous_moon:  |                |
| [Script](https://help.sonatype.com/en/script-api.html)                                    |      :full_moon:       |                |
| [Security Management](https://help.sonatype.com/en/security-management-api.html)          |       :new_moon:       |      3.19      |
| [Staging](https://help.sonatype.com/repomanager3/staging) _pro_                           | :waning_gibbous_moon:  |                |
| [Status](https://help.sonatype.com/en/status-api.html)                                    |      :full_moon:       |                |
| [Support](https://help.sonatype.com/en/support-api.html)                                  |      :full_moon:       |                |
| [Tagging](https://help.sonatype.com/en/tagging.html) _pro_                                | :waning_gibbous_moon:  |                |
| [Tasks](https://help.sonatype.com/en/tasks-api.html)                                      |       :new_moon:       |                |

#### Supported Provisioning API

| API        |        Status         |
| ---------- | :-------------------: |
| Core       |      :new_moon:       |
| Security   |      :new_moon:       |
| Blob Store | :waning_gibbous_moon: |
| Repository | :waning_gibbous_moon: |

_Legend_: :full_moon: complete :new_moon: untouched :waning_crescent_moon::last_quarter_moon::waning_gibbous_moon: partial support

### nexusiq [![GoDoc](http://godoc.org/github.com/sonatype-nexus-community/gonexus/iq?status.png)](http://godoc.org/github.com/sonatype-nexus-community/gonexus/iq) [![nexusiq coverage](https://gocover.io/_badge/github.com/sonatype-nexus-community/gonexus/iq?0 "nexusiq coverage")](http://gocover.io/github.com/sonatype-nexus-community/gonexus/iq)

Create a connection to an instance of Nexus IQ Server

```go
// import "github.com/sonatype-nexus-community/gonexus/iq"
iq, err := nexusiq.New("http://localhost:8070", "username", "password")
if err != nil {
    panic(err)
}

```

#### Supported IQ Endpoints

| Endpoint                                                                                                                            |   Status    | Min IQ Version |
| ----------------------------------------------------------------------------------------------------------------------------------- | :---------: | :------------: |
| [Advanced Search](https://help.sonatype.com/en/advanced-search-rest-api.html)                                                       | :new_moon:  |                |
| [Applicable Waivers](https://help.sonatype.com/en/applicable-waivers-rest-api.html)                                                 | :new_moon:  |                |
| [Application Categories](https://help.sonatype.com/en/application-categories-rest-api.html)                                         | :new_moon:  |                |
| [Application](https://help.sonatype.com/en/application-rest-api.html)                                                               | :full_moon: |                |
| [Atlassian Crowd](https://help.sonatype.com/en/atlassian-crowd-rest-api.html)                                                       | :new_moon:  |                |
| [Audit Log](https://help.sonatype.com/en/audit-log-rest-api.html)                                                                   | :new_moon:  |      r175      |
| [Authorization Configuration](https://help.sonatype.com/en/authorization-configuration-rest-api.html)                               | :full_moon: |      r70       |
| [Call Flow Analysis](https://help.sonatype.com/en/call-flow-analysis-rest-api.html)                                                 | :new_moon:  |      r176      |
| [Component Claim](https://help.sonatype.com/en/component-claim-rest-api.html)                                                       | :new_moon:  |                |
| [Component Details](https://help.sonatype.com/en/component-details-rest-api.html)                                                   | :full_moon: |                |
| [Component Evaluation](https://help.sonatype.com/en/component-evaluation-rest-api.html)                                             | :full_moon: |                |
| [Component Labels](https://help.sonatype.com/en/component-labels-rest-api.html)                                                     | :full_moon: |                |
| [Component Remediation](https://help.sonatype.com/en/component-remediation-rest-api.html)                                           | :full_moon: |      r64       |
| [Component Search](https://help.sonatype.com/en/component-search-rest-api.html)                                                     | :full_moon: |                |
| [Component in Quarantine](https://help.sonatype.com/en/components-in-quarantine-rest-api.html)                                      | :new_moon:  |                |
| [Component Versions](https://help.sonatype.com/en/component-versions-rest-api.html)                                                 | :full_moon: |                |
| [Component Waivers](https://help.sonatype.com/en/component-waivers-rest-api.html)                                                   | :new_moon:  |      r76       |
| [Configuration](https://help.sonatype.com/en/configuration-rest-api.html)                                                           | :new_moon:  |      r65       |
| [Cross-Stage Policy Violation](https://help.sonatype.com/en/cross-stage-policy-violation-rest-api.html)                             | :new_moon:  |                |
| [CycloneDx](https://help.sonatype.com/en/cyclonedx-rest-api.html)                                                                   | :new_moon:  |                |
| [Data Retention Policy](https://help.sonatype.com/en/data-retention-policy-rest-api.html)                                           | :full_moon: |                |
| [Feature Configuration](https://help.sonatype.com/en/feature-configuration-rest-api.html)                                           | :new_moon:  |      r149      |
| [Firewall](https://help.sonatype.com/en/firewall-rest-api.html)                                                                     | :new_moon:  |      r171      |
| [HTTP Proxy Server Configuration](https://help.sonatype.com/en/http-proxy-server-configuration-rest-api.html)                       | :new_moon:  |                |
| [Jira Configuration](https://help.sonatype.com/en/jira-configuration-rest-api.html)                                                 | :new_moon:  |      r139      |
| [License Legal](https://help.sonatype.com/en/license-legal-rest-api.html)                                                           | :new_moon:  |                |
| [Mail](https://help.sonatype.com/en/mail-rest-api.html)                                                                             | :new_moon:  |                |
| [Manifest Evaluation](https://help.sonatype.com/en/manifest-evaluation-rest-api.html)                                               | :new_moon:  |                |
| [Organizations](https://help.sonatype.com/en/organizations-rest-api.html)                                                           | :full_moon: |                |
| [Policy Violation](https://help.sonatype.com/en/policy-violation-rest-api.html)                                                     | :full_moon: |                |
| [Policy Waiver](https://help.sonatype.com/en/policy-waiver-rest-api.html)                                                           | :new_moon:  |      r71       |
| [Product License](https://help.sonatype.com/en/product-license-rest-api.html)                                                       | :new_moon:  |                |
| [Promote Scan](https://help.sonatype.com/en/promote-scan-rest-api.html)                                                             | :new_moon:  |                |
| [Release from Quarantine](https://help.sonatype.com/en/release-from-quarantine-rest-api.html)                                       | :new_moon:  |                |
| [Report](https://help.sonatype.com/en/report-rest-api.html)                                                                         | :new_moon:  |                |
| [Repository Firewall Evaluation](https://help.sonatype.com/en/firewall-evaluation-api.html)                                         | :new_moon:  |                |
| [Repository Results View](https://help.sonatype.com/en/repository-results-view-rest-api.html)                                       | :new_moon:  |                |
| [Reverse Proxy Authentication Configuration](https://help.sonatype.com/en/reverse-proxy-authentication-configuration-rest-api.html) | :new_moon:  |                |
| [Role](https://help.sonatype.com/en/role-rest-api.html)                                                                             | :full_moon: |      r70       |
| [SAML](https://help.sonatype.com/en/saml-rest-api.html)                                                                             | :new_moon:  |      r74       |
| [Security Vulnerability Override](https://help.sonatype.com/en/security-vulnerability-override-api.html)                            | :new_moon:  |                |
| [Source Control](https://help.sonatype.com/en/source-control-rest-api.html)                                                         | :full_moon: |                |
| [Source Control Configuration](https://help.sonatype.com/en/source-control-configuration-rest-api.html)                             | :new_moon:  |      r140      |
| [Source Control Evaluation](https://help.sonatype.com/en/source-control-evaluation-rest-api.html)                                   | :new_moon:  |                |
| [SPDX](https://help.sonatype.com/en/spdx-rest-api.html)                                                                             | :new_moon:  |                |
| [Similar Waivers](https://help.sonatype.com/en/similar-waivers-rest-api.html)                                                       | :new_moon:  |                |
| [Stale Waivers](https://help.sonatype.com/en/stale-waivers-rest-api.html)                                                           | :new_moon:  |                |
| [Success Metrics Data](https://help.sonatype.com/en/success-metrics-data-rest-api.html)                                             | :full_moon: |                |
| [Third-Party Analysis](https://help.sonatype.com/en/third-party-scan-rest-api.html)                                                 | :new_moon:  |                |
| [Transitive Waivers](https://help.sonatype.com/en/transitive-waivers-rest-api.html)                                                 | :new_moon:  |                |
| [User](https://help.sonatype.com/en/user-rest-api.html)                                                                             | :full_moon: |      r70       |
| [User Token](https://help.sonatype.com/en/user-token-rest-api.html)                                                                 | :new_moon:  |      r76       |
| [Vulnerability Analysis Details](https://help.sonatype.com/en/vulnerability-analysis-details-rest-api.html)                         | :new_moon:  |      r168      |
| [Vulnerability Custom Attributes](https://help.sonatype.com/en/vulnerability-custom-attributes-rest-api.html)                       | :new_moon:  |      r161      |
| [Vulnerability Details](https://help.sonatype.com/iqserver/automating/rest-apis/vulnerability-details-rest-api---v2)                | :new_moon:  |      r75       |
| [Vulnerability Group](https://help.sonatype.com/en/vulnerability-group-rest-api.html)                                               | :new_moon:  |      162       |
| [Webhooks](https://help.sonatype.com/en/lifecycle-webhooks-events-types.html)                                                       | :new_moon:  |                |

_Legend_: :full_moon: complete :new_moon: untouched :waning_crescent_moon::last_quarter_moon::waning_gibbous_moon: partial support

##### iqwebhooks [![GoDoc](http://godoc.org/github.com/sonatype-nexus-community/gonexus/iq/iqwebhooks?status.png)](http://godoc.org/github.com/sonatype-nexus-community/gonexus/iq/iqwebhooks) [![nexusiq webhooks coverage](https://gocover.io/_badge/github.com/sonatype-nexus-community/gonexus/iq/iqwebhooks/?0 "nexusiq webhooks coverage")](http://gocover.io/github.com/sonatype-nexus-community/gonexus/iq/iqwebhooks)

The `iq/iqwebhooks` subpackage provides structs for all of the event types along with helper functions.

Most notably it provides a function called `Listen` which is an `http.HandlerFunc` that can be used as an endpoint handler for a server functioning as a webhook listener.
The handler will place any webhook event it finds in a channel to be consumed at will.

An example of using the handler to listen for Application Evaluation events:

```go
// import "github.com/sonatype-nexus-community/gonexus/iq/webhooks"
appEvals, _ := iqwebhooks.ApplicationEvaluationEvents()

go func() {
    for _ = range appEvals:
        log.Println("Received Application Evaluation event")
    }
}()

http.HandleFunc("/ingest", iqwebhooks.Listen)
```

See the [documentation](https://godoc.org/github.com/sonatype-nexus-community/gonexus/iq/iqwebhooks#example-Listen) for a full example showing other event types.

## The Fine Print

It is worth noting that this is **NOT SUPPORTED** by [Sonatype](//www.sonatype.com), and is a contribution of [@HokieGeek](https://github.com/HokieGeek)
plus us to the open source community (read: you!)

Remember:

- Use this contribution at the risk tolerance that you have
- Do **NOT** file Sonatype support tickets related to this
- **DO** file issues here on GitHub, so that the community can pitch in
