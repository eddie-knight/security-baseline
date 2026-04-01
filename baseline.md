# Open Source Project Security Baseline - osps-baseline 


_Open Source Project Security Baseline_ is a Gemara ControlCatalog, conforming to specification version **v1.0.0-rc.1** by OSPS Baseline Authors (Human, `baseline-authors`).

**Warning**: This is a draft, and should only be used as a preview of the upcoming artifact publication.



### Description

The Open Source Project Security (OSPS) Baseline is a set of security criteria that projects should meet to demonstrate a strong security posture.


### Requirement Applicability Groups

The following groups are used to specify the circumstance within which an assessment requirement is mandated.

- **maturity-1** — Maturity Level 1: for any code or non-code project with any number of maintainers or users
- **maturity-2** — Maturity Level 2: for any code project that has at least 2 maintainers and a small number of consistent users
- **maturity-3** — Maturity Level 3: for any code project that has a large number of consistent users



_Summary: 41 control(s), 65 assessment requirement(s)._



## Table of contents

- [Access Control](#ac)
  - [OSPS-AC-01: Use MFA for Sensitive Actions
](#osps-ac-01-use-mfa-for-sensitive-actions)
  - [OSPS-AC-02: Restrict Collaborator Permissions
](#osps-ac-02-restrict-collaborator-permissions)
  - [OSPS-AC-03: Protect the Primary Branch from Accidental Modification
](#osps-ac-03-protect-the-primary-branch-from-accidental-modification)
  - [OSPS-AC-04: Enforce Least Privilege on CI/CD Pipelines
](#osps-ac-04-enforce-least-privilege-on-ci-cd-pipelines)
- [Build and Release](#br)
  - [OSPS-BR-01: Prevent Untrusted Input When Building & Releasing
](#osps-br-01-prevent-untrusted-input-when-building-releasing)
  - [OSPS-BR-02: Assign Unique Version Identifiers
](#osps-br-02-assign-unique-version-identifiers)
  - [OSPS-BR-03: Use Encrypted Channels for Development & Release Activity
](#osps-br-03-use-encrypted-channels-for-development-release-activity)
  - [OSPS-BR-04: Publish Change Log With Release
](#osps-br-04-publish-change-log-with-release)
  - [OSPS-BR-05: Use Standardized Dependency Management Tools
](#osps-br-05-use-standardized-dependency-management-tools)
  - [OSPS-BR-06: Include Signatures and Hashes With Release
](#osps-br-06-include-signatures-and-hashes-with-release)
  - [OSPS-BR-07: Secure Secrets and Credentials
](#osps-br-07-secure-secrets-and-credentials)
- [Documentation](#do)
  - [OSPS-DO-01: Publish User Guides for Basic Functionality
](#osps-do-01-publish-user-guides-for-basic-functionality)
  - [OSPS-DO-02: Provide Mechanisms for Reporting Defects
](#osps-do-02-provide-mechanisms-for-reporting-defects)
  - [OSPS-DO-03: Publish Provenance Verification Instructions
](#osps-do-03-publish-provenance-verification-instructions)
  - [OSPS-DO-04: Publish Support Scope and Duration
](#osps-do-04-publish-support-scope-and-duration)
  - [OSPS-DO-05: Document Security Update Scope and Duration
](#osps-do-05-document-security-update-scope-and-duration)
  - [OSPS-DO-06: Publish Dependency Management Policy
](#osps-do-06-publish-dependency-management-policy)
  - [OSPS-DO-07: Provide Instructions on How to Build From Source
](#osps-do-07-provide-instructions-on-how-to-build-from-source)
- [Governance](#gv)
  - [OSPS-GV-01: Publish Project Roles and Responsibilities
](#osps-gv-01-publish-project-roles-and-responsibilities)
  - [OSPS-GV-02: Provide Public Discussion Mechanisms
](#osps-gv-02-provide-public-discussion-mechanisms)
  - [OSPS-GV-03: Publish Contribution Guide
](#osps-gv-03-publish-contribution-guide)
  - [OSPS-GV-04: Require Formal Review of Permission Grants
](#osps-gv-04-require-formal-review-of-permission-grants)
- [Legal](#le)
  - [OSPS-LE-01: Require Code Contributors to Assert Right to Commit
](#osps-le-01-require-code-contributors-to-assert-right-to-commit)
  - [OSPS-LE-02: Ensure Project Licenses are Fully Open Source
](#osps-le-02-ensure-project-licenses-are-fully-open-source)
  - [OSPS-LE-03: Maintain and Release Licenses in a Well Known Location
](#osps-le-03-maintain-and-release-licenses-in-a-well-known-location)
- [Quality](#qa)
  - [OSPS-QA-01: Publish Source Code and Change History
](#osps-qa-01-publish-source-code-and-change-history)
  - [OSPS-QA-02: Publish Software Dependencies
](#osps-qa-02-publish-software-dependencies)
  - [OSPS-QA-03: Address Pass/Fail Checks Before Accepting Changes
](#osps-qa-03-address-pass-fail-checks-before-accepting-changes)
  - [OSPS-QA-04: Enforce Security Requirements on All Codebases
](#osps-qa-04-enforce-security-requirements-on-all-codebases)
  - [OSPS-QA-05: Prevent Executables in the Codebase
](#osps-qa-05-prevent-executables-in-the-codebase)
  - [OSPS-QA-06: Use Automated Testing in CI/CD Pipelines
](#osps-qa-06-use-automated-testing-in-ci-cd-pipelines)
  - [OSPS-QA-07: Require Merge Approvals
](#osps-qa-07-require-merge-approvals)
- [Security Assessment](#sa)
  - [OSPS-SA-01: Publish Design Descriptions of System Actors and Actions
](#osps-sa-01-publish-design-descriptions-of-system-actors-and-actions)
  - [OSPS-SA-02: Publish External Interface Descriptions
](#osps-sa-02-publish-external-interface-descriptions)
  - [OSPS-SA-03: Maintain a Project Security Assessment
](#osps-sa-03-maintain-a-project-security-assessment)
- [Vulnerability Management](#vm)
  - [OSPS-VM-01: Publish Coordinated Vulnerability Disclosure Policy
](#osps-vm-01-publish-coordinated-vulnerability-disclosure-policy)
  - [OSPS-VM-02: Publish Contacts and Process for Reporting Vulnerabilities.
](#osps-vm-02-publish-contacts-and-process-for-reporting-vulnerabilities)
  - [OSPS-VM-03: Maintain Private Vulnerability Reporting Process
](#osps-vm-03-maintain-private-vulnerability-reporting-process)
  - [OSPS-VM-04: Publish Discovered Vulnerabilities
](#osps-vm-04-publish-discovered-vulnerabilities)
  - [OSPS-VM-05: Publish and Enforce a Dependency Remediation Policy
](#osps-vm-05-publish-and-enforce-a-dependency-remediation-policy)
  - [OSPS-VM-06: Publish and Enforce an Application Security Testing Policy
](#osps-vm-06-publish-and-enforce-an-application-security-testing-policy)








## AC: Access Control

Access Control focuses on the mechanisms and
policies that control access to the project's version
control system and CI/CD pipelines. These controls help
ensure that only authorized users can access sensitive
data, modify repository settings, or execute build and
release processes.

### OSPS-AC-01: Use MFA for Sensitive Actions




**Objective**

Reduce the risk of account compromise or insider threats by requiring
multi-factor authentication for collaborators modifying the project
repository settings or accessing sensitive data.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: AC-4(21) · AC-17 · CM-5 · CM-6 · IA-2 · IA-5 · 1.2e · 1.2f

- **BPB**: CC-G-1

- **BSI-TR-03185-2**: GV.02

- **CRA**: 1.2d · 1.2e · 1.2f

- **CSF**: PR.A-02 · PR.A-05

- **OpenCRE**: 486-813 · 124-564 · 347-352 · 333-858 · 152-725 · 201-246

- **PCIDSS**: 2.2.1 · 8.2.1 · 8.3.1

- **PSSCRM**: G2.6 · P3.3 · E1.2 · E1.3 · E1.4 · E3.1

- **SAMM**: Operations -Environment Management -Configuration Hardening Lvl1

- **SSDF**: PO.3.2 · PS.1 · PS.2

- **UKSSCOP**: Claim 1.4.2 · Claim 2.1.5 · Claim 2.2.2



#### OSPS-AC-01.01



When a user attempts to read or modify a sensitive resource in the project's
authoritative repository, the system MUST require the user to complete
a multi-factor authentication process.


**Applicability:** maturity-1, maturity-2, maturity-3

**Recommendation**: Enforce multi-factor authentication for the project's version
control system, requiring collaborators to provide a second form of
authentication when accessing sensitive data or modifying repository
settings. Passkeys are acceptable for this control.


### OSPS-AC-02: Restrict Collaborator Permissions




**Objective**

Reduce the risk of unauthorized access to the project's repository by
limiting the permissions granted to new collaborators.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: AC-2 · AC-3 · AC-4(21) · AC-5 · AC-6 · CM-5 · CM-7

- **BSI-TR-03185-2**: GV.02

- **CRA**: 1.2f

- **CSF**: PR.AA-02 · PR.AA-05

- **OpenCRE**: 486-813 · 124-564 · 802-056 · 368-633 · 152-725

- **PCIDSS**: 2.2.1

- **PSSCRM**: P2.3 · E1.2 · E3.3

- **SSDF**: PO.2 · PO.3.2 · PS.1 · PS.2

- **UKSSCOP**: Claim 2.2.2



#### OSPS-AC-02.01



When a new collaborator is added, the version control system MUST
require manual permission assignment, or restrict the collaborator
permissions to the lowest available privileges by default.


**Applicability:** maturity-1, maturity-2, maturity-3

**Recommendation**: Most public version control systems are configured in this manner.
Ensure the project's version control system always assigns the lowest
available permissions to collaborators by default when added, granting
additional permissions only when necessary.


### OSPS-AC-03: Protect the Primary Branch from Accidental Modification




**Objective**

Reduce the risk of accidental changes or deletion of the primary branch
of the project's repository by preventing unintentional modification.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: AC-3 · AC-5 · CM-3 · CM-3(2) · CM-5

- **BSI-TR-03185-2**: GV.02 · QA.06

- **CRA**: 1.2f

- **CSF**: PR.A-02 · PR.A-05

- **OpenCRE**: 486-813 · 124-564 · 152-725

- **PCIDSS**: 2.2.1

- **PSSCRM**: P3.2 · P3.5 · E1.5 · E3.1

- **Scorecard**: Branch-Protection

- **SSDF**: PO.3.2 · PS.1 · PS.2

- **UKSSCOP**: Claim 1.1.4 · Claim 2.2.2



#### OSPS-AC-03.01



When a direct commit is attempted on the project's primary branch,
an enforcement mechanism MUST prevent the change from being applied.


**Applicability:** maturity-1, maturity-2, maturity-3

**Recommendation**: If the VCS is centralized, set branch protection on the primary branch
in the project's VCS. Alternatively, use a decentralized approach,
like the Linux kernel's, where changes are first proposed in another
repository, and merging changes into the primary repository requires a
specific separate act.

#### OSPS-AC-03.02



When an attempt is made to delete the project's primary branch,
the version control system MUST treat this as a sensitive activity
and require explicit confirmation of intent.


**Applicability:** maturity-1, maturity-2, maturity-3

**Recommendation**: Set branch protection on the primary branch in the project's version
control system to prevent deletion.


### OSPS-AC-04: Enforce Least Privilege on CI/CD Pipelines




**Objective**

Reduce the risk of unauthorized access to the project's build and release
processes by limiting the permissions granted to steps within the CI/CD
pipelines.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: AC-3(8) · AC-4 · AC-4(6) · AC-6 · AC-20 · AC-20(1) · CM-5 · CM-7

- **CRA**: 1.2d · 1.2e · 1.2f

- **CSF**: PR.AA-02 · PR.AA-05

- **OpenCRE**: 486-813 · 124-564 · 347-507 · 263-284 · 123-124

- **PCIDSS**: 2.2.1 · 8.2.1

- **PSSCRM**: P3.2

- **SAMM**: Operations -Environment Management -Configuration Hardening Lvl1

- **SLSA**: Producer - Choose an appropriate build platform · Build platform - Isolation strength - Isolated

- **SSDF**: PO.2 · PO.3.2 · PS.1 · PS.2

- **UKSSCOP**: Claim 2.1.1 · Claim 2.1.3 · Claim 2.2.2



#### OSPS-AC-04.01



When a CI/CD task is executed with no permissions specified, the
CI/CD system MUST default the task's permissions to the lowest
permissions granted in the pipeline.


**Applicability:** maturity-2, maturity-3

**Recommendation**: Configure the project's settings to assign the lowest available
permissions to new pipelines by default, granting additional
permissions only when necessary for specific tasks.

#### OSPS-AC-04.02



When a job is assigned permissions in a CI/CD pipeline, the source
code or configuration MUST only assign the minimum privileges
necessary for the corresponding activity.


**Applicability:** maturity-3

**Recommendation**: Configure the project's CI/CD pipelines to assign the lowest available
permissions to users and services by default, elevating permissions
only when necessary for specific tasks. In some version control
systems, this may be possible at the organizational or repository
level. If not, set permissions at the top level of the pipeline.




## BR: Build and Release

Build and Release focuses on the processes and
tools used to compile, package, and distribute the
project's software. These controls help ensure that the
project's build and release pipelines are secure,
consistent, and reliable, reducing the risk of
vulnerabilities or errors in the software distribution
process.

### OSPS-BR-01: Prevent Untrusted Input When Building & Releasing




**Objective**

Reduce the risk of code injection or other security vulnerabilities in the
project's build and release pipelines by preventing untrusted input from
accessing privileged resources.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: AC-3 · AC-4 · AC-4(21) · CM-5 · CM-7 · SI-7

- **CRA**: 1.2f

- **CSF**: PR.AA-02

- **OpenCRE**: 486-813 · 124-564 · 357-352

- **PCIDSS**: 2.2.1 · 6.4.1

- **PSSCRM**: P2.3 · P3.2 · P3.5 · E2.4 · E2.5 · D2.2

- **Scorecard**: Dangerous-Workflow

- **SLSA**: Choose an appropriate build platform

- **SSDF**: PO.3.2 · PO.5.2 · PS.1 · PS.2

- **UKSSCOP**: Claim 2.1.2 · Claim 2.2.2



#### OSPS-BR-01.01



When a CI/CD pipeline operates on untrusted metadata, those
parameters MUST be sanitized and validated prior to use in the
pipeline.


**Applicability:** maturity-1, maturity-2, maturity-3

**Recommendation**: CI/CD pipelines should sanitize (quote, escape or exit on expected
values) all metadata inputs which correspond to untrusted sources.
This includes data such as branch names, commit messages, tags, pull request titles,
and author information.

#### OSPS-BR-01.02



Retired in https://github.com/ossf/security-baseline/pull/443


**Applicability:** retired

#### OSPS-BR-01.03



When a CI/CD pipeline operates on untrusted code snapshots, it MUST
prevent access to privileged CI/CD credentials and assets.


**Applicability:** maturity-1, maturity-2, maturity-3

**Recommendation**: CI/CD pipelines should isolate untrusted code snapshots from
privileged credentials and assets. In particular, projects should be
careful to ensure that workflows which build or execute code prior
to review by a collaborator do not have access to CI/CD credentials.

#### OSPS-BR-01.04



CI/CD pipelines which accept trusted collaborator input MUST sanitize and
validate that input prior to use in the pipeline.


**Applicability:** maturity-3

**Recommendation**: CI/CD pipelines should sanitize (quote, escape or exit on expected
values) all collaborator inputs on explicit workflow executions.
While collaborators are generally trusted, manual inputs to a
workflow cannot be reviewed and could be abused by an account
takeover or insider threat.


### OSPS-BR-02: Assign Unique Version Identifiers




**Objective**

Ensure that each software asset produced by the project is uniquely
identified, enabling users to track changes and updates to the project
over time.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: IA-4 · SA-15 · SI-7 · SR-4

- **BPB**: CC-B-5 · CC-B-6 · CC-B-7

- **BSI-TR-03185-2**: BR.02

- **CRA**: 1.2f

- **OpenCRE**: 486-813 · 124-564

- **PCIDSS**: 6.4.3

- **PSSCRM**: G1.4 · E1.2 · E2.1 · E2.6

- **SLSA**: Follow a consistent build process · Provenance generation- Exists, Authentic

- **SSDF**: PO.3.2 · PS.1 · PS.2 · PS.3

- **UKSSCOP**: Claim 1.1.4 · Claim 3.1.1 · Claim 3.4.2



#### OSPS-BR-02.01



When an official release is created, that release MUST be assigned a
unique version identifier.


**Applicability:** maturity-2, maturity-3

**Recommendation**: Assign a unique version identifier to each release produced by the
project, following a consistent naming convention or numbering scheme.
Examples include SemVer, CalVer, or git commit id.

#### OSPS-BR-02.02



When an official release is created, all assets within that release
MUST be clearly associated with the release identifier or another
unique identifier for the asset.


**Applicability:** maturity-3

**Recommendation**: Assign a unique version identifier to each software asset produced by
the project, following a consistent naming convention or numbering
scheme. Examples include SemVer, CalVer, or git commit id.


### OSPS-BR-03: Use Encrypted Channels for Development & Release Activity




**Objective**

Protect the confidentiality and integrity of project source code during
development, reducing the risk of eavesdropping or data tampering.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: AC-4 · AC-4(21)

- **BPB**: B-B-11

- **BSI-TR-03185-2**: BR.03

- **CRA**: 1.2d · 1.2e · 1.2f · 1.2i · 1.2j · 1.2k

- **OpenCRE**: 483-813 · 124-564 · 263-184

- **PCIDSS**: 2.2.1 · 2.2.7 · 4.2.1 · 4.2.2 · 6.4.1 · 8.3.2

- **PSSCRM**: E1.1 · E2.2 · E2.4 · E2.5

- **SLSA**: Choose an appropriate build platform

- **SSDF**: PO.3.2 · PO.5.2 · PS.1 · PS.2

- **UKSSCOP**: Claim 3.1.2



#### OSPS-BR-03.01



When the project lists a URI as an official project channel, that URI
MUST be exclusively delivered using encrypted channels.


**Applicability:** maturity-1, maturity-2, maturity-3

**Recommendation**: Configure the project's websites and version control systems to use
encrypted channels such as SSH or HTTPS for data transmission.
Ensure all tools and domains referenced in project documentation can
only be accessed via encrypted channels.

#### OSPS-BR-03.02



When the project lists a URI as an official distribution channel,
that channel MUST be protected from adversary-in-the-middle
attacks using cryptographically authenticated channels.


**Applicability:** maturity-1, maturity-2, maturity-3

**Recommendation**: Artifacts distributed by the project should be distributed through
channels which ensure integrity and authenticity. Use of HTTPS for
downloads, signed releases, or distribution through trusted package
managers are all acceptable methods to protect against
adversary-in-the-middle attacks.


### OSPS-BR-04: Publish Change Log With Release




**Objective**

Provide transparency and accountability for changes made to the project's
software releases in such a way that users can understand the
modifications and improvements included in each release.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: AU-2 · AU-6 · AU-10 · CM-5 · CM-6 · MA-1 · MA-8 · SI-4 · SI-5

- **BPB**: CC-B-8 · CC-B-9 · Q-B-7 · A-B-1 · A-S-1

- **BSI-TR-03185-2**: BR.04

- **CRA**: 1.2d · 1.2f · 1.2h · 1.2j · 1.2l · 2.5

- **OpenCRE**: 483-813 · 068-486 · 124-564 · 757-271 · 347-352 · 263-184 · 208-355 · 745-356 · 732-148

- **PCIDSS**: 6.2.1 · 6.4.1 · 6.5.1 · 6.5.2 · 10.2.2

- **PSSCRM**: G1.4 · E2.1 · E2.4 · E2.5 · E3.1 · E3.6

- **SLSA**: Choose an appropriate build platform · Follow a consistent build process · Build platform - Isolation strength - isolated

- **SSDF**: PS.1 · PS.2 · PS.3 · PW.1.2

- **UKSSCOP**: Claim 1.1.4 · Claim 2.2.3 · Claim 3.1.1



#### OSPS-BR-04.01



When an official release is created, that release MUST contain
a descriptive log of functional and security
modifications.


**Applicability:** maturity-2, maturity-3

**Recommendation**: Ensure that all releases include a descriptive change log. It is
recommended to ensure that the change log is human-readable and
includes details beyond commit messages, such as descriptions of the
security impact or relevance to different use cases. To ensure
machine readability, place the content under a markdown header
such as "## Changelog".


### OSPS-BR-05: Use Standardized Dependency Management Tools




**Objective**

Ensure that the project's build and release pipelines use standardized tools
and processes to manage dependencies, reducing the risk of compatibility
issues or security vulnerabilities in the software.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: AC-4 · CM-2 · CM-7(4) · CM-7(5) · RA-5 · SA-15 · SR-3

- **BPB**: Q-B-2

- **CRA**: 1.2b · 1.2d · 1.2f · 1.2h · 1.2j · 2.1 · 2.2 · 2.3

- **OpenCRE**: 486-813 · 124-564 · 347-352 · 715-334

- **PCIDSS**: 6.4.3

- **PSSCRM**: P3.1 · P3.5 · E2.2 · E2.3 · E2.4 · E2.5

- **SAMM**: Implementation -Secure Build -Build Process Lvl2

- **SLSA**: Isolation strength - isolated

- **SSDF**: PO.3.2 · PS.1 · PS.2

- **UKSSCOP**: Claim 1.2.1 · Claim 1.2.5



#### OSPS-BR-05.01



When a build and release pipeline ingests dependencies, it MUST
use standardized tooling where available.


**Applicability:** maturity-2, maturity-3

**Recommendation**: Use a common tooling for your ecosystem, such as package managers or
dependency management tools to ingest dependencies at build time. This
may include using a dependency file, lock file, or manifest to specify
the required dependencies, which are then pulled in by the build
system.


### OSPS-BR-06: Include Signatures and Hashes With Release




**Objective**

Ensure released software assets can be verified by users to ensure
integrity of each asset when it is used.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: AU-10 · MP-1 · SA-15 · SI-7 · SI-7(14)

- **BSI-TR-03185-2**: BR.03

- **PCIDSS**: 2.2.1 · 2.2.7 · 3.5.1 · 4.2.1 · 4.2.2 · 6.4.1 · 8.3.2

- **PSSCRM**: P1.2 · P3.2 · P3.3 · E2.1 · E2.2 · E2.6

- **SAMM**: Implementation -Secure Deployment -Deployment Process Lvl3

- **Scorecard**: Signed-Releases

- **SLSA**: Distribute provenance - Exists

- **SSDF**: PO.5.2 · PS.2 · PS.2.1 · PW.6.2

- **UKSSCOP**: Claim 1.2.2 · Claim 3.1.1



#### OSPS-BR-06.01



When an official release is created, that release MUST be signed or
accounted for in a signed manifest including each asset's
cryptographic hashes.


**Applicability:** maturity-2, maturity-3

**Recommendation**: Sign all released software assets at build time with a cryptographic
signature or attestations, such as GPG or PGP signature, Sigstore
signatures, SLSA provenance, or SLSA VSAs. Include the cryptographic
hashes of each asset in a signed manifest or metadata file.


### OSPS-BR-07: Secure Secrets and Credentials




**Objective**

Ensure that data which can lead to security vulnerabilities or
supply chain compromise is not disclosed, compromised, or misused.


#### Guidelines

This control aids in the application of the following guidelines:

- **BPB**: S-B-5

- **SSDF**: PO.1.1 · P0.3.1 · P0.4.2 · PO.5.1 · PW.1.2 · PW.1.3 · PW.5.1

- **UKSSCOP**: Claim 1.4.3 · Claim 1.4.5



#### OSPS-BR-07.01



The project MUST prevent the unintentional storage of unencrypted sensitive data, such as secrets and credentials, in the version control system.


**Applicability:** maturity-1

**Recommendation**: Configure .gitignore or equivalent to exclude files that may contain sensitive information. Use pre-commit hooks and automated scanning tools to detect and prevent the inclusion of sensitive data in commits.

#### OSPS-BR-07.02



The project MUST define a policy for managing secrets and credentials used by the project. The policy should include guidelines for storing, accessing, and rotating secrets and credentials.


**Applicability:** maturity-3

**Recommendation**: Document how secrets and credentials are managed and used within the project. This should include details on how secrets are stored (e.g., using a secrets management tool), how access is controlled, and how secrets are rotated or updated. Ensure that sensitive information is not hard-coded in the source code or stored in version control systems.




## DO: Documentation

Documentation focuses on the information
provided to users, contributors, and maintainers
of the project. These controls help ensure that
the project's documentation is comprehensive,
accurate, and up-to-date, enabling users to
understand the project's features and functionality, maintenance, support,
security and release practices.

### OSPS-DO-01: Publish User Guides for Basic Functionality




**Objective**

Ensure that users have a clear and comprehensive understanding of the
project's current features in order to prevent damage from misuse or
misconfiguration.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: CM-2 · PL-2 · PL-8 · SA-15

- **BPB**: B-B-1 · B-B-9 · B-S-7 · B-S-9

- **BSI-TR-03185-2**: BR.01

- **CRA**: 1.2b · 1.2j · 1.2k

- **CSF**: GV.OC-04 · GV.OC-05

- **ISO-18974**: 4.1.4

- **OpenCRE**: 036-275

- **PCIDSS**: 2.1.1 · 2.2.1 · 3.1.1 · 4.1.1 · 5.1.1 · 6.1.1 · 6.2.1 · 7.1.1 · 8.1.1 · 11.1.1 · 12.10.5

- **PSSCRM**: G5.1 · E3.5

- **SSDF**: PW.1.2

- **UKSSCOP**: 4.1



#### OSPS-DO-01.01



When the project has made a release, the project documentation MUST
include user guides for all basic functionality.


**Applicability:** maturity-1, maturity-2, maturity-3

**Recommendation**: Create user guides or documentation for all basic functionality of the
project, explaining how to install, configure, and use the project's
features. If there are any known dangerous or destructive actions
available, include highly-visible warnings.


### OSPS-DO-02: Provide Mechanisms for Reporting Defects




**Objective**

Enable users and contributors to report defects or issues with the
released software assets, facilitating communication and collaboration on
defect fixes and improvements.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: IR-6 · SI-4 · SI-5

- **BPB**: B-B-3 · R-B-1+ · R-B-1 · R-B-2 · R-S-2

- **BSI-TR-03185-2**: QA.03

- **CRA**: 1.2c · 1.2l · 2.1 · 2.2 · 2.5 · 2.6

- **CSF**: RS.MA-02 · GV.RM-05

- **ISO-18974**: 4.2.1

- **PCIDSS**: 6.3.2 · 6.3.3 · 6.5.1 · 6.5.2 · 12.10.2

- **SAMM**: Implementation -Defect Management -Defect Tracking Lvl1 · Implementation -Defect Management -Defect Tracking Lvl2

- **SSDF**: PW.1.2 · RV.1.1 · RV.2.1 · RV.1.2

- **UKSSCOP**: 1.1 · 1.3



#### OSPS-DO-02.01



When the project has made a release, the project documentation MUST
include a guide for reporting defects.


**Applicability:** maturity-1, maturity-2, maturity-3

**Recommendation**: It is recommended that projects use their VCS default issue tracker.
If an external source is used, ensure that the project documentation
and contributing guide clearly and visibly explain how to use the
reporting system. It is recommended that project documentation also
sets expectations for how defects will be triaged and resolved.


### OSPS-DO-03: Publish Provenance Verification Instructions




**Objective**

Enable users to verify the authenticity and integrity of the project's
released software assets, reducing the risk of using tampered or
unauthorized versions of the software.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: CM-2 · IR-1 · MP-1 · SA-15 · SI-7 · SI-7(14)

- **BPB**: CC-B-8

- **BSI-TR-03185-2**: BR.01 · BR.03

- **CRA**: 1.2d

- **OpenCRE**: 171-222

- **PCIDSS**: 3.1.1 · 3.5.1 · 4.1.1 · 5.1.1 · 6.1.1 · 6.2.1 · 7.1.1 · 8.1.1 · 11.1.1

- **PSSCRM**: G1.3 · G2.5 · P1.2 · P3.1 · P3.2 · P3.3 · E2.6

- **SSDF**: PO.4.2 · PS.2 · PS.2.1 · PS.3.1 · RV.1.3

- **UKSSCOP**: 3.1



#### OSPS-DO-03.01



When the project has made a release, the project documentation MUST
contain instructions to verify the integrity and authenticity of the
release assets.


**Applicability:** maturity-3

**Recommendation**: Instructions in the project should contain information about the
technology used, the commands to run, and the expected output.
When possible, avoid storing this documentation in the same location
as the build and release pipeline to avoid a single breach
compromising both the software and the documentation for verifying the
integrity of the software.

#### OSPS-DO-03.02



When the project has made a release, the project documentation MUST
contain instructions to verify the expected identity of the person or
process authoring the software release.


**Applicability:** maturity-3

**Recommendation**: The expected identity may be in the form of key IDs used to sign,
issuer and identity from a sigstore certificate, or other similar
forms.
When possible, avoid storing this documentation in the same location
as the build and release pipeline to avoid a single breach
compromising both the software and the documentation for verifying the
integrity of the software.


### OSPS-DO-04: Publish Support Scope and Duration




**Objective**

Provide users with clear expectations regarding the project's support
lifecycle in such a way that enables downstream consumers to ensure the
continued functionality and security of their systems.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: PL-1 · PL-2 · SI-4

- **BPB**: R-B-3

- **BSI-TR-03185-2**: DE.01

- **ISO-18974**: 4.1 · 4.3.1

- **PCIDSS**: 2.1.1 · 3.1.1 · 3.2.1 · 4.1.1 · 5.1.1 · 6.1.1 · 6.3.3 · 7.1.1 · 8.1.1 · 11.1.1

- **PSSCRM**: E1.6

- **SAMM**: Operations -Operational Management -System Decommissioning -Legacy Management Lvl1

- **SSDF**: PO.4.2 · PS.3.1 · RV.1.3

- **UKSSCOP**: 4.1 · 4.2 · Claim 4.1.1 · Claim 4.1.2 · Claim 4.2.1



#### OSPS-DO-04.01



When the project has made a release, the project documentation MUST
include a descriptive statement about the scope and duration of
support for each release.


**Applicability:** maturity-3

**Recommendation**: In order to communicate the scope and duration of support for the
project's released software assets, the project should have a
SUPPORT.md file, a "Support" section in SECURITY.md, or
other documentation explaining the support lifecycle,
including the expected duration of support for each release, the
types of support provided (e.g., bug fixes, security updates), and
any relevant policies or procedures for obtaining support.


### OSPS-DO-05: Document Security Update Scope and Duration




**Objective**

Communicate when the project maintainers will no longer fix defects or
security vulnerabilities.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: PL-1 · PL-2 · SI-4 · SI-5

- **BSI-TR-03185-2**: DE.01

- **CRA**: 1.2c · 2.6

- **ISO-18974**: 4.1.1 · 4.3.1

- **OpenCRE**: 673-475 · 053-751

- **PCIDSS**: 3.1.1 · 3.2.1 · 4.1.1 · 5.1.1 · 6.1.1 · 6.3.2 · 7.1.1 · 8.1.1 · 11.1.1

- **PSSCRM**: E1.6

- **SAMM**: Operations -Operational Management -System Decommissioning -Legacy Management Lvl1 · Operations -Operational Management -System Decommissioning -Legacy Management Lvl2

- **UKSSCOP**: 3.5 · 4.1 · Claim 4.1.1 · Claim 4.2.1



#### OSPS-DO-05.01



When the project has made a release, the project documentation MUST
provide a descriptive statement when releases or versions will no
longer receive security updates.


**Applicability:** maturity-3

**Recommendation**: In order to communicate the scope and duration of support for security
fixes, the project should have a SUPPORT.md or other documentation
explaining the project's policy for security updates.


### OSPS-DO-06: Publish Dependency Management Policy




**Objective**

Provide information about how the project selects, obtains, and tracks
third-party components that are required for the software to function.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: CA-7 · CM-7(5) · CM-8 · PM-30 · RA-3(1) · SA-11 · SI-4 · SR-3 · SR-5 · SR-6 · SR-7

- **BPB**: A-S-1

- **BSI-TR-03185-2**: QA.01

- **CRA**: 2.1

- **OpenCRE**: 613-286 · 053-751

- **PCIDSS**: 2.1.1 · 3.1.1 · 4.1.1 · 5.1.1 · 6.1.1 · 6.3.2 · 6.4.3 · 7.1.1 · 8.1.1 · 11.1.1 · 12.5.2

- **PSSCRM**: G1.4 · G2.4 · P3.1 · P3.2 · P3.4

- **SAMM**: Design -Security Requirements -Supplier Security Lvl2

- **Scorecard**: Pinned-Dependencies

- **UKSSCOP**: 1.2 · 3.3 · Claim 1.2.1 · Claim 1.2.2



#### OSPS-DO-06.01



When the project has made a release, the project documentation MUST
include a description of how the project selects, obtains, and tracks
its dependencies.


**Applicability:** maturity-2, maturity-3

**Recommendation**: It is recommended to publish this information alongside the project's
technical & design documentation on a publicly viewable resource such
as the source code repository, project website, or other channel.  


### OSPS-DO-07: Provide Instructions on How to Build From Source




**Objective**

Ensure that users have a clear and comprehensive instructions on how
to build the software from source code.




#### OSPS-DO-07.01



The project documentation MUST
include instructions on how to build the software, including required
libraries, frameworks, SDKs, and dependencies.


**Applicability:** maturity-2, maturity-3

**Recommendation**: It is recommended to publish this information alongside the project's
contributor documentation, such as in `CONTRIBUTING.md` or other
developer task documentation. This may also be documented using
`Makefile` targets or other automation scripts.




## GV: Governance

Governance focuses on the policies and
procedures that guide the project's decision-making
and community interactions. These controls help ensure
that the project is well positioned to respond to
both threats and opportunities.

### OSPS-GV-01: Publish Project Roles and Responsibilities




**Objective**

Document project roles and responsibilities to help project participants,
potential contributors, and downstream consumers understand who is working
on the project and what areas of authority they may have.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: AC-2 · AC-3 · IA-2 · PL-1 · PL-4 · PM-30

- **BPB**: B-S-3 · B-S-4

- **OpenCRE**: 013-021

- **PCIDSS**: 2.1.2 · 3.1.1 · 3.1.2 · 4.1.1 · 4.1.2 · 5.1.1 · 5.1.2 · 6.1.1 · 6.1.2 · 6.5.4 · 7.1.1 · 7.1.2 · 8.1.1 · 8.1.2 · 11.1.1 · 11.1.2 · 12.1.3 · 12.5.2

- **PSSCRM**: G2.3 · E3.1 · E3.3

- **UKSSCOP**: Claim 2.1.1



#### OSPS-GV-01.01



While active, the project documentation MUST include a list of
project members with access to sensitive resources.


**Applicability:** maturity-2, maturity-3

**Recommendation**: Document project participants and their roles through such artifacts
as members.md, governance.md, maintainers.md, or similar file within
the source code repository of the project.
This may be as simple as including names or account handles in a list
of maintainers, or more complex depending on the project's governance.

#### OSPS-GV-01.02



While active, the project documentation MUST include descriptions of
the roles and responsibilities for members of the project.


**Applicability:** maturity-2, maturity-3

**Recommendation**: Document project participants and their roles through such artifacts
as members.md, governance.md, maintainers.md, or similar file within
the source code repository of the project.


### OSPS-GV-02: Provide Public Discussion Mechanisms




**Objective**

Encourage open communication and collaboration within the project
community by enabling users to provide feedback and discuss proposed
changes or usage challenges.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: AC-21 · AU-6 · PL-1

- **BPB**: B-B-3 · B-B-12

- **BSI-TR-03185-2**: QA.03

- **CRA**: 1.2l · 2.3 · 2.4 · 2.6

- **PCIDSS**: 12.5.2

- **SSDF**: PS.3 · PW.1.2



#### OSPS-GV-02.01



While active, the project MUST have one or more mechanisms for public
discussions about proposed changes and usage obstacles.


**Applicability:** maturity-1, maturity-2, maturity-3

**Recommendation**: Establish one or more mechanisms for public discussions within the
project, such as mailing lists, instant messaging, or issue trackers,
to facilitate open communication and feedback.


### OSPS-GV-03: Publish Contribution Guide




**Objective**

Provide guidance on how to participate in the project, outlining the steps
required to submit changes or enhancements to the project's codebase.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: AC-3 · AC-20 · PL-1

- **BPB**: B-B-4 · B-S-3 · B-B-4+ · R-B-1 · Q-G-2

- **BSI-TR-03185-2**: GV.01 · QA.04 · QA.05

- **CRA**: 1.2l · 2.4

- **ISO-18974**: 4.1.2

- **PCIDSS**: 2.1.1 · 6.5.4 · 8.2.1 · 12.5.2

- **PSSCRM**: G2.4 · P2.2

- **SSDF**: PW.1.2

- **UKSSCOP**: Claim 2.1.1



#### OSPS-GV-03.01



While active, the project documentation MUST include an explanation
of the contribution process.


**Applicability:** maturity-1, maturity-2, maturity-3

**Recommendation**: Create a CONTRIBUTING.md or CONTRIBUTING/ directory to outline the
contribution process including the steps for submitting changes, and
engaging with the project maintainers.

#### OSPS-GV-03.02



While active, the project documentation MUST include a guide for code
contributors that includes requirements for acceptable contributions.


**Applicability:** maturity-2, maturity-3

**Recommendation**: Extend the CONTRIBUTING.md or CONTRIBUTING/ contents in the project
documentation to outline the requirements for acceptable
contributions, including coding standards, testing requirements, and
submission guidelines for code contributors. It is recommended that
this guide is the source of truth for both contributors and approvers.


### OSPS-GV-04: Require Formal Review of Permission Grants




**Objective**

Ensure that code contributors are vetted and reviewed before being granted
elevated permissions to sensitive resources within the project, reducing
the risk of unauthorized access or misuse.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: AC-2 · AC-3 · AC-4(21) · AC-5 · AC-6 · AC-20 · CM-7 · IR-4(6) · PM-30 · SI-4

- **BPB**: B-B-5 · B-S-3 · B-B-4+ · Q-G-2

- **CRA**: 1.2d · 1.2l · 2.1 · 2.2 · 2.5 · 2.6

- **CSF**: PR.AA-02 · PR.AA-05

- **ISO-18974**: 4.1.2

- **OpenCRE**: 123-124 · 152-725

- **PCIDSS**: 2.1.1 · 6.5.4 · 8.2.1 · 8.2.2

- **PSSCRM**: E3.1 · E3.3

- **SSDF**: PO.2 · PO.3.2



#### OSPS-GV-04.01



While active, the project documentation MUST have a policy that code
collaborators are reviewed prior to granting escalated permissions to
sensitive resources.


**Applicability:** maturity-3

**Recommendation**: Publish an enforceable policy in the project documentation that
requires code collaborators to be reviewed and approved before being
granted escalated permissions to sensitive resources, such as merge
approval or access to secrets. It is recommended that vetting includes
establishing a justifiable lineage of identity such as confirming the
contributor's association with a known trusted organization.




## LE: Legal

Legal focuses on the policies and
procedures that govern the project's licensing
and intellectual property. These controls help
ensure that the project's source code is
distributed under a recognized and legally
enforceable open source software license,
reducing the risk of intellectual property
disputes or licensing violations.

### OSPS-LE-01: Require Code Contributors to Assert Right to Commit




**Objective**

Ensure that code contributors are aware of and acknowledge their legal
responsibility for the contributions they make to the project, reducing
the risk of intellectual property disputes against the project.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: PL-4

- **BPB**: B-S-1

- **CRA**: 1.2b · 1.2f

- **PCIDSS**: 12.8.5

- **PSSCRM**: E3.1

- **SSDF**: PO.3.2 · PS.1 · PW.1.2 · PW.2.1



#### OSPS-LE-01.01



While active, the version control system MUST require all code
contributors to assert that they are legally authorized to make the
associated contributions on every commit.


**Applicability:** maturity-2, maturity-3

**Recommendation**: Include a DCO in the project's repository, requiring code
contributors to assert that they are legally authorized to commit the
associated contributions on every commit. Use a status check to ensure
the assertion is made. A CLA also satisfies this requirement.
Some version control systems, such as GitHub, may include this in the
platform terms of service.

It is understood that projects with a lengthy history prior to
adopting OSPS Baseline may not be able to retroactively enforce this
requirement.


### OSPS-LE-02: Ensure Project Licenses are Fully Open Source




**Objective**

Ensure that the project's source code is distributed under a recognized
and legally enforceable open source software license, providing clarity on
how the code can be used and shared by others.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: PL-4

- **BPB**: B-B-6 · B-B-7

- **BSI-TR-03185-2**: LE.01

- **CRA**: 1.2b

- **CSF**: GV.OC-03

- **PCIDSS**: 3.2.1

- **PSSCRM**: G1.2

- **Scorecard**: License

- **SSDF**: PO.3.2



#### OSPS-LE-02.01



While active, the license for the source code MUST meet the OSI Open
Source Definition or the FSF Free Software Definition.


**Applicability:** maturity-1, maturity-2, maturity-3

**Recommendation**: Add a LICENSE file to the project's repo with a license that is an
approved license by the Open Source Initiative (OSI), or a free
license as approved by the Free Software Foundation (FSF). Examples of
such licenses include the MIT, BSD 2-clause, BSD 3-clause revised,
Apache 2.0, Lesser GNU General Public License (LGPL), and the GNU
General Public License (GPL). Releasing to the public domain meets
this control if there are no other encumbrances such as patents.

#### OSPS-LE-02.02



While active, the license for the released software assets MUST meet
the OSI Open Source Definition or the FSF Free Software Definition.


**Applicability:** maturity-1, maturity-2, maturity-3

**Recommendation**: If a different license is included with released software assets,
ensure it is an approved license by the Open Source Initiative (OSI),
or a free license as approved by the Free Software Foundation (FSF).
Examples of such licenses include the MIT, BSD 2-clause, BSD 3-clause
revised, Apache 2.0, Lesser GNU General Public License (LGPL), and the
GNU General Public License (GPL). Note that the license for the
released software assets may be different than the source code.


### OSPS-LE-03: Maintain and Release Licenses in a Well Known Location




**Objective**

Ensure that the project's source code and released software assets are
distributed with the appropriate license terms, making it clear to users
and contributors how each can be used and shared.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: PL-4

- **BPB**: B-B-8

- **BSI-TR-03185-2**: LE.02

- **CRA**: 1.2b

- **PCIDSS**: 3.2.1

- **PSSCRM**: G1.2

- **Scorecard**: License

- **SSDF**: PO.3.2



#### OSPS-LE-03.01



While active, the license for the source code MUST be maintained in
the corresponding repository's LICENSE file, COPYING file, or
LICENSE/ directory.


**Applicability:** maturity-1, maturity-2, maturity-3

**Recommendation**: Include the project's source code license in the project's LICENSE
file, COPYING file, or LICENSE/ directory to provide visibility and
clarity on the licensing terms. The filename MAY have an extension.
If the project has multiple repositories, ensure that each repository
includes the license file.

#### OSPS-LE-03.02



While active, the license for the released software assets MUST be
included in the released source code, or in a LICENSE file, COPYING
file, or LICENSE/ directory alongside the corresponding release
assets.


**Applicability:** maturity-1, maturity-2, maturity-3

**Recommendation**: Include the project's released software assets license in the released
source code, or in a LICENSE file, COPYING file, or LICENSE/ directory
alongside the corresponding release assets to provide visibility and
clarity on the licensing terms. The filename MAY have an extension.
If the project has multiple repositories, ensure that each repository
includes the license file.




## QA: Quality

Quality focuses on the processes and
practices used to ensure the quality and
reliability of the project's source code and
software assets. These controls help ensure
that the project's source code is well
maintained, secure, and reliable, reducing the
risk of defects or vulnerabilities in the
software.

### OSPS-QA-01: Publish Source Code and Change History




**Objective**

Enable users to access and review the project's source code and history,
promoting transparency and collaboration within the project community.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: RA-5 · SA-11 · SA-15

- **BPB**: CC-B-1 · CC-B-2 · CC-B-3 · R-B-5

- **BSI-TR-03185-2**: QA.02

- **CRA**: 1.2b · 1.2f · 1.2j

- **CSF**: ID.AM-02 · ID.RA-01 · ID.RA-08

- **ISO-18974**: 4.1.4

- **OpenCRE**: 486-813 · 124-564 · 757-271

- **PCIDSS**: 2.1.1 · 6.2.1 · 6.5.1 · 6.5.2

- **PSSCRM**: P3.5 · E2.2

- **SAMM**: Implementation -Secure Build -Build Process Lvl1

- **SLSA**: Build platform - isolation strength - Isolated

- **SSDF**: PS.1 · PS.2 · PS.3 · PW.1.2 · PW.2.1

- **UKSSCOP**: Claim 2.2.3



#### OSPS-QA-01.01



While active, the project's source code repository MUST be publicly
readable at a static URL.


**Applicability:** maturity-1, maturity-2, maturity-3

**Recommendation**: Use a common VCS such as GitHub, GitLab, or Bitbucket. Ensure the
repository is publicly readable. Avoid duplication or mirroring of
repositories unless highly visible documentation clarifies the primary
source. Avoid frequent changes to the repository that would impact the
repository URL. Ensure the repository is public.

#### OSPS-QA-01.02



The version control system MUST contain a publicly readable record of
all changes made, who made the changes, and when the changes were
made.


**Applicability:** maturity-1, maturity-2, maturity-3

**Recommendation**: Use a common VCS such as GitHub, GitLab, or Bitbucket to maintain a
publicly readable commit history. Avoid squashing or rewriting commits
in a way that would obscure the author of any commits.


### OSPS-QA-02: Publish Software Dependencies




**Objective**

Provide transparency and accountability for the project's dependencies
while enabling users and contributors to understand the software's direct
dependencies.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: CA-7 · CM-2 · CM-8 · PL-8 · RA-3(1) · RA-5 · SA-11 · SA-15 · SR-3 · SR-4

- **BPB**: Q-S-8 · Q-S-9

- **BSI-TR-03185-2**: QA.01

- **CRA**: 2.1 · 2.2 · 2.3

- **CSF**: ID.AM.01 · ID.AM-02

- **ISO-18974**: 4.1.5 · 4.3.1

- **OpenCRE**: 486-813 · 124-564 · 673-475 · 863-521 · 613-286

- **PCIDSS**: 6.3.2 · 6.4.3 · 12.5.1

- **PSSCRM**: G1.4 · G1.5 · G2.5 · P3.1 · P3.2 · P5.1 · P5.2 · E2.1 · E2.2

- **SAMM**: Implementation -Secure Build -Software Dependencies Lvl1

- **SSDF**: PO.3.3 · PS.1 · PS.2 · PS.3.2 · PW.4

- **UKSSCOP**: Claim 1.2.1 · Claim 1.2.2 · Claim 3.1.1



#### OSPS-QA-02.01



When the package management system supports it, the source code
repository MUST contain a dependency list that accounts for the direct
language dependencies.


**Applicability:** maturity-1, maturity-2, maturity-3

**Recommendation**: This may take the form of a package manager or language dependency file
that enumerates all direct dependencies such as package.json, Gemfile,
or go.mod.

#### OSPS-QA-02.02



When the project has made a release, all compiled released software
assets MUST be delivered with a software bill of materials.


**Applicability:** maturity-3

**Recommendation**: It is recommended to auto-generate SBOMs at build time using a tool
that has been vetted for accuracy. This enables users to ingest this
data in a standardized approach alongside other projects in their
environment.


### OSPS-QA-03: Address Pass/Fail Checks Before Accepting Changes




**Objective**

Ensure that the project's approvers do not become accustomed to tolerating
failing status checks, even if arbitrary, because it increases the risk of
overlooking security vulnerabilities or defects identified by automated
checks.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: AU-6 · CM-3 · CM-6 · PL-8 · SA-11 · SA-15 · SR-3

- **BSI-TR-03185-2**: QA.04

- **CRA**: 1.2f · 1.2k

- **CSF**: ID.IM-02

- **ISO-18974**: 4.1.5

- **OpenCRE**: 263-184 · 253-452

- **PCIDSS**: 6.3.1 · 6.3.2 · 6.5.2

- **PSSCRM**: G2.2 · G5.3 · G5.4 · P3.5 · P4.1 · P4.2

- **SAMM**: Implementation -Secure Build -Build Process Lvl3 · Implementation -Secure Build -Software Dependencies Lvl3 · Verification -Requirements Testing -Control Verification Lvl1 · Verification -Requirements Testing -Control Verification Lvl2 · Verification -Requirements Testing -Control Verification Lvl3

- **SSDF**: PO.4.1 · PS.1 · PS.2 · RV.1.2

- **UKSSCOP**: Claim 1.3.2 · Claim 1.3.3



#### OSPS-QA-03.01



When a commit is made to the primary branch, any automated status
checks for commits MUST pass or be manually bypassed.


**Applicability:** maturity-2, maturity-3

**Recommendation**: Configure the project's version control system to require that all
automated status checks pass or require manual acknowledgement before a
commit can be merged into the primary branch. It is recommended that
any optional status checks are NOT configured as a pass or fail
requirement that approvers may be tempted to bypass.


### OSPS-QA-04: Enforce Security Requirements on All Codebases




**Objective**

Ensure that all codebases produced by the project are well
documented and held to the same security standard.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: PL-8 · SA-15

- **CRA**: 1.2b · 1.2f

- **OpenCRE**: 486-813 · 124-564

- **PCIDSS**: 6.4.2

- **PSSCRM**: G2.2 · G5.4

- **Scorecard**: Binary-Artifacts

- **SLSA**: Build platform - isolation strength - Isolated

- **SSDF**: PO.3.2 · PO.4.1 · PS.1 · PS.2 · RV.1.2



#### OSPS-QA-04.01



Projects with multiple repositories MUST document a list of codebases
that are part of the project.


**Applicability:** maturity-1, maturity-2, maturity-3

**Recommendation**: Document any additional subproject code repositories produced by the
project and compiled into a release. This documentation should include
the status and intent of the respective codebase.

#### OSPS-QA-04.02



When the project has made a release comprising multiple source code
repositories, all subprojects MUST enforce security requirements that
are as strict or stricter than the primary codebase.


**Applicability:** maturity-3

**Recommendation**: Any additional subproject code repositories produced by the project
and compiled into a release must enforce security requirements as
applicable to the status and intent of the respective codebase.
In addition to following the corresponding OSPS Baseline requirements,
this may include requiring a security review, ensuring that it is
free of vulnerabilities, and ensuring that it is free of known
security issues.


### OSPS-QA-05: Prevent Executables in the Codebase




**Objective**

Reduce the risk of including generated executable artifacts in the
project's version control system, ensuring that only source code and
necessary files are stored in the repository.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: PL-8 · SA-15 · SR-3

- **BSI-TR-03185-2**: BR.05

- **CRA**: 1.2b

- **OpenCRE**: 486-813 · 124-564

- **PCIDSS**: 6.4.3

- **SSDF**: PS.1 · PS.2



#### OSPS-QA-05.01



While active, the version control system MUST NOT contain generated
executable artifacts.


**Applicability:** maturity-1, maturity-2, maturity-3

**Recommendation**: Remove generated executable artifacts in the project's version control
system. It is recommended that any scenario where a generated
executable artifact appears critical to a process such as testing, it
should be instead be generated at build time or stored separately and
fetched during a specific well-documented pipeline step.

#### OSPS-QA-05.02



While active, the version control system MUST NOT contain unreviewable
binary artifacts.


**Applicability:** maturity-1, maturity-2, maturity-3

**Recommendation**: Do not add any unreviewable binary artifacts to the project's version
control system. This includes executable application binaries, library
files, and similar artifacts. It does not include assets such as
graphical images, sound or music files, and similar content typically
stored in a binary format.


### OSPS-QA-06: Use Automated Testing in CI/CD Pipelines




**Objective**

Ensure that the project uses at least one automated test suite for the
source code repository and clearly documents when and how tests are run.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: SA-11 · SA-15 · SR-3

- **BPB**: Q-B-4 · Q-B-8 · Q-B-9 · Q-B-10 · Q-S-2

- **BSI-TR-03185-2**: QA.04 · QA.05

- **CRA**: 2.3

- **CSF**: ID.AM-02

- **ISO-18974**: 4.1.5

- **OpenCRE**: 207-435 · 088-377

- **PCIDSS**: 6.2.3 · 6.3.1 · 6.3.2 · 6.4.2

- **PSSCRM**: P4.1 · P4.2 · P4.3 · P4.4 · E2.4 · E2.5

- **SAMM**: Verification-Requirements -Testing -Control Verification Lvl1 · Verification-Requirements -Testing -Control Verification Lvl2 · Verification-Requirements -Testing -Control Verification Lvl3 · Verification -Security Testing -Scalable Baseline Lvl3

- **Scorecard**: CI-Tests

- **SSDF**: PW.8.2



#### OSPS-QA-06.01



Prior to a commit being accepted, the project's CI/CD pipelines MUST
run at least one automated test suite to ensure the changes meet
expectations.


**Applicability:** maturity-2, maturity-3

**Recommendation**: Automated tests should be run prior to every merge into the primary
branch. The test suite should be run in a CI/CD pipeline and the
results should be visible to all contributors. The test suite should
be run in a consistent environment and should be run in a way that
allows contributors to run the tests locally.
Examples of test suites include unit tests, integration tests, and
end-to-end tests.

#### OSPS-QA-06.02



While active, project's documentation MUST clearly document when and
how tests are run.


**Applicability:** maturity-3

**Recommendation**: Add a section to the contributing documentation that explains how to
run the tests locally and how to run the tests in the CI/CD pipeline.
The documentation should explain what the tests are testing and how to
interpret the results.

#### OSPS-QA-06.03



While active, the project's documentation MUST include a policy that
all major changes to the software produced by the project should add
or update tests of the functionality in an automated test suite.


**Applicability:** maturity-3

**Recommendation**: Add a section to the contributing documentation that explains the
policy for adding or updating tests. The policy should explain what
constitutes a major change and what tests should be added or updated.


### OSPS-QA-07: Require Merge Approvals




**Objective**

Ensure that the project's version control system requires at least one
non-author human approval of changes before merging into the release or
primary branch.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: AC-5 · AU-6 · PL-8 · SA-15 · SR-3

- **BPB**: B-G-3

- **BSI-TR-03185-2**: QA.06

- **PCIDSS**: 6.2.3.1 · 6.4.2 · 6.5.4

- **PSSCRM**: G2.4 · P3.3 · P3.5

- **Scorecard**: Code-Review



#### OSPS-QA-07.01



When a commit is made to the primary branch, the project's version
control system MUST require at least one non-author human approval of the
changes before merging.


**Applicability:** maturity-3

**Recommendation**: Configure the project's version control system to require at least one
non-author human approval of changes before merging into the release or
primary branch. This can be achieved by requiring a pull request to be
reviewed and approved by at least one other collaborator before it can
be merged.




## SA: Security Assessment

Security Assessment encourages practices that
help ensure that the project is well positioned
to identify and address security vulnerabilities
and threats in the software.

### OSPS-SA-01: Publish Design Descriptions of System Actors and Actions




**Objective**

Provide an overview of the project's design and architecture, illustrating
the interactions and components of the system to help contributors and
security reviewers understand the internal logic of the released software
assets.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: CM-2 · PL-8 · RA-3 · SA-15

- **BPB**: B-B-1 · B-S-7 · B-S-8

- **CRA**: 1.2a · 1.2b

- **CSF**: ID.AM-02

- **OpenCRE**: 155-155 · 326-704 · 068-102 · 036-275 · 162-655

- **PCIDSS**: 2.2.1 · 2.2.3 · 2.2.4 · 2.2.5 · 2.2.6 · 3.1.1 · 4.1.1 · 5.1.1 · 6.1.1 · 6.2.1 · 7.1.1 · 8.1.1 · 11.1.1 · 12.3.1 · 12.5.3

- **PSSCRM**: G5.1 · P1.1 · E3.4 · E3.7

- **SAMM**: Operations -Operational Management -Data Protection Lvl2

- **SSDF**: PO.1 · PO.2 · PO.3.2

- **UKSSCOP**: Claim 1.1.5



#### OSPS-SA-01.01



When the project has made a release, the project documentation MUST
include design documentation demonstrating all actions and actors
within the system.


**Applicability:** maturity-2, maturity-3

**Recommendation**: Include designs in the project documentation that explains the actions
and actors. Actors include any subsystem or entity that can influence
another segment in the system.
Ensure this is updated for new features or breaking changes.


### OSPS-SA-02: Publish External Interface Descriptions




**Objective**

Provide users and developers with an understanding of how to interact with
the project's software and integrate it with other systems, enabling them
to use the software effectively.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: CM-2 · PL-2 · PL-8 · RA-3 · SA-15

- **BPB**: B-B-10 · B-S-7

- **CRA**: 1.2a · 1.2b

- **CSF**: GV.OC-05 · ID.AM-01

- **ISO-18974**: 4.1.4

- **OpenCRE**: 155-155 · 068-102 · 072-713 · 820-878

- **PCIDSS**: 2.2.1 · 2.2.3 · 2.2.4 · 2.2.5 · 2.2.6 · 6.2.1 · 12.3.1 · 12.8.1

- **PSSCRM**: E3.4 · E3.7

- **SSDF**: PW.1.2

- **UKSSCOP**: Claim 1.1.5



#### OSPS-SA-02.01



When the project has made a release, the project documentation MUST
include descriptions of all external software interfaces of the
released software assets.


**Applicability:** maturity-2, maturity-3

**Recommendation**: Document all software interfaces (APIs) of the released software
assets, explaining how users can interact with the software and what
data is expected or produced.
Ensure this is updated for new features or breaking changes.


### OSPS-SA-03: Maintain a Project Security Assessment




**Objective**

Provide project maintainers an understanding of how the software can be
misused or broken, enabling them to plan mitigations accordingly.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: CA-2 · CA-2(3) · PM-30 · RA-3 · SA-11 · SA-15 · SA-15(3) · SA-15(8) · SI-3 · SR-3 · SR-3(3) · SR-6 · SR-7

- **BPB**: B-S-8 · S-G-1

- **CRA**: 1.1 · 1.2j · 1.2k · 2.2

- **CSF**: ID.RA-01 · ID.RA-04 · ID.RA-05 · DE.AE-07

- **ISO-18974**: 4.1.5

- **OpenCRE**: 068-102 · 154-031 · 888-770

- **PCIDSS**: 2.2.4 · 2.2.5 · 2.2.6 · 6.2.1 · 6.2.3.1 · 6.3.2 · 6.4.2 · 11.3.1 · 12.3.1

- **PSSCRM**: G4.3 · G5.2 · P2.1

- **SAMM**: Governance -Create and Promote Lvl1 · Design -Threat Assessment -Application Risk Profile Lvl1 · Design -Threat Assessment -Threat Modeling Lvl1 · Verification -Architecture Assessment -Architecture Mitigation Lvl2

- **SSDF**: PO.5.1 · PW.1.1

- **UKSSCOP**: Claim 1.4.1



#### OSPS-SA-03.01



When the project has made a release, the project MUST perform a
security assessment to understand the most likely and impactful
potential security problems that could occur within the software.


**Applicability:** maturity-2, maturity-3

**Recommendation**: Performing a security assessment informs both project members as well
as downstream consumers that the project understands what problems
could arise within the software. Understanding what threats could be
realized helps the project manage and address risk. This information
is useful to downstream consumers to demonstrate the security acumen
and practices of the project.
Ensure this is updated for new features or breaking changes.

#### OSPS-SA-03.02



When the project has made a release, the project MUST perform a threat
modeling and attack surface analysis to understand and protect against
attacks on critical code paths, functions, and interactions within the
system.


**Applicability:** maturity-3

**Recommendation**: Threat modeling is an activity where the project looks at the
codebase, associated processes and infrastructure, interfaces, key
components and "thinks like a hacker" and brainstorms how the system
be be broken or compromised. Each identified threat is listed out so
the project can then think about how to proactively avoid or close off
any gaps/vulnerabilities that could arise.
Ensure this is updated for new features or breaking changes.




## VM: Vulnerability Management

Vulnerability Management focuses on the
processes and practices used to identify and
address security vulnerabilities in the project's
software dependencies. These controls help ensure
that the project is well positioned to respond to
security threats and vulnerabilities in the software.

### OSPS-VM-01: Publish Coordinated Vulnerability Disclosure Policy




**Objective**

Establish a process for reporting and addressing vulnerabilities in the
project, ensuring that security issues are handled promptly and
transparently.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: IR-1 · IR-4 · IR-6 · IR-7(1) · IR-8 · SI-2

- **BPB**: R-B-6 · R-B-8 · R-S-2 · S-B-14 · S-B-15

- **BSI-TR-03185-2**: VM.01

- **CRA**: 2.1 · 2.2 · 2.3 · 2.6 · 2.7 · 2.8

- **CSF**: GV.PO-01 · GV.PO-02 · ID.RA-01 · ID.RA-08

- **ISO-18974**: 4.1.5 · 4.2.1 · 4.3.2

- **OpenCRE**: 887-750

- **PCIDSS**: 2.1.1 · 3.1.1 · 4.1.1 · 5.1.1 · 6.1.1 · 6.3.1 · 6.3.2 · 7.1.1 · 8.1.1 · 11.1.1 · 11.2.1 · 12.1.1 · 12.1.3

- **PSSCRM**: D1.1 · D1.2 · D1.3 · D1.5

- **SAMM**: Governance -Create and Promote Lvl2 · Governance -Policy & Compliance -Policy & Standards Lvl1 · Implementation -Defect Management -Defect Tracking Lvl1 · Implementation -Defect Management -Defect Tracking Lvl2 · Implementation -Defect Management -Defect Tracking Lvl3 · Operations -Incident Management -Incident Response Lvl1 · Operations -Incident Management -Incident Response Lvl2 · Operations -Incident Management -Incident Response Lvl3

- **Scorecard**: Security-Policy

- **SSDF**: RV.1.3

- **UKSSCOP**: Claim 3.4.1 · Claim 3.5.1 · Claim 4.1.2



#### OSPS-VM-01.01



While active, the project documentation MUST
include a policy for coordinated vulnerability disclosure (CVD), with a clear
timeframe for response.


**Applicability:** maturity-2, maturity-3

**Recommendation**: Create a SECURITY.md file at the root of the directory, outlining the
project's policy for coordinated vulnerability disclosure. Include a
method for reporting vulnerabilities. Set expectations for how the
project will respond and address reported issues.


### OSPS-VM-02: Publish Contacts and Process for Reporting Vulnerabilities.




**Objective**

Enable external parties, such as researchers, to easily understand who they
should contact in the event that a vulnerability is found, and what process
they should follow.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: IR-1 · IR-4 · IR-6 · IR-8

- **BPB**: B-S-8

- **BSI-TR-03185-2**: VM.01

- **CRA**: 2.5

- **CSF**: GV.PO-01 · GV.PO-02 · ID.RA-01

- **ISO-18974**: 4.1.1 · 4.1.3 · 4.1.5 · 4.2.2

- **OpenCRE**: 464-513

- **PCIDSS**: 6.3.3 · 12.1.1 · 12.10.2

- **SAMM**: Governance -Policy&Compliance -Policy&Standards Lvl2

- **Scorecard**: Security-Policy

- **SSDF**: RV.1.3

- **UKSSCOP**: 3.2



#### OSPS-VM-02.01



While active, the project documentation MUST contain
security contacts.


**Applicability:** maturity-1

**Recommendation**: Create a security.md (or similarly-named) file that contains security
contacts for the project.


### OSPS-VM-03: Maintain Private Vulnerability Reporting Process




**Objective**

Provide a means for reporting privately to the security contacts within
the project so that security vulnerabilities are not be shared with the
public until the project has been provided time to analyze and prepare
remediations to protect users of the project.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: IR-6

- **BPB**: R-B-7

- **BSI-TR-03185-2**: VM.01

- **CRA**: 2.5 · 2.6

- **OpenCRE**: 308-514

- **PCIDSS**: 6.3.1 · 6.3.3 · 12.10.2

- **SAMM**: Operations -Incident Management -Incident Response Lvl3

- **UKSSCOP**: 3.2



#### OSPS-VM-03.01



While active, the project documentation MUST
provide a means for private vulnerability reporting directly to
the security contacts within the project.


**Applicability:** maturity-2, maturity-3

**Recommendation**: Provide a means for security researchers to report vulnerabilities
privately to the project. This may be a dedicated email address, a
web form, VCS specialized tools, email addresses for security
contacts, or other methods.


### OSPS-VM-04: Publish Discovered Vulnerabilities




**Objective**

Ensure that project end users have a well known mechanism to understand
vulnerabilities found within the project.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: CA-7 · CM-3 · CM-8 · IR-5 · SI-2 · SI-4 · SI-5

- **BSI-TR-03185-2**: VM.02

- **CRA**: 1.2a · 1.2b · 2.1 · 2.4 · 2.6

- **CSF**: ID.RA-01

- **ISO-18974**: 4.1.5

- **PCIDSS**: 6.2.3 · 6.3.1 · 6.3.2 · 6.3.3 · 11.3.1

- **PSSCRM**: G2.2 · D1.1

- **SSDF**: PO.4.1 · RV.2.1 · RV.2.2

- **UKSSCOP**: 3.4 · 3.5 · 4.3



#### OSPS-VM-04.01



While active, the project documentation MUST
publicly publish data about discovered vulnerabilities.


**Applicability:** maturity-2, maturity-3

**Recommendation**: Provide information about known vulnerabilities in a predictable
public channel, such as a CVE entry, blog post, or other medium.
To the degree possible, this information should include affected
version(s), how a consumer can determine if they are vulnerable, and
instructions for mitigation or remediation.

#### OSPS-VM-04.02



While active, any vulnerabilities in the
software components not affecting the project MUST be accounted for
in a VEX document, augmenting the vulnerability report with
non-exploitability details.


**Applicability:** maturity-3

**Recommendation**: Establish a VEX feed communicating the exploitability status of
known vulnerabilities, including assessment details or any
mitigations in place preventing vulnerable code from being
executed.


### OSPS-VM-05: Publish and Enforce a Dependency Remediation Policy




**Objective**

Identify and address defects and security weaknesses in the project's
imported code early in the development process, reducing the risk of
shipping insecure software.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: CA-7 · RA-5 · SA-11 · SI-2 · SI-3

- **BPB**: B-S-8 · Q-B-12 · Q-S-9 · S-B-14 · S-B-15 · A-B-1 · A-B-3 · A-B-8 · A-S-1

- **CRA**: 1.2a · 1.2b · 1.2c · 2.1 · 2.2 · 2.3 · 2.4

- **CSF**: GV.RM-05 · GV.RM-06 · GV.PO-01 · GV.PO-02 · ID.RA-01 · ID.RA-08 · ID.IM-02

- **ISO-18974**: 4.1.5 · 4.2.1 · 4.2.2 · 4.3.2

- **OpenCRE**: 155-155 · 124-564 · 757-271 · 464-513 · 611-158 · 207-435 · 088-377

- **PCIDSS**: 6.2.3 · 6.3.1 · 6.3.2 · 6.4.1 · 6.4.2

- **PSSCRM**: G5.4 · P4.1 · P4.2 · P4.3 · P4.4 · P4.5

- **SAMM**: Implementation -Secure Build-Build Process Lvl3 · Implementation -Software Dependencies Lvl3 · Verification -Security Testing -Scalable Baseline Lvl1 · Verification -Security Testing -Scalable Baseline Lvl3

- **Scorecard**: Security-Policy · Vulnerabilities

- **SSDF**: PO.4 · PW.1.2 · PW.8.1 · RV.1.2 · RV.1.3 · RV.2.1 · RV.2.2

- **UKSSCOP**: 1.2 · 3.3



#### OSPS-VM-05.01



While active, the project documentation MUST include a policy that
defines a threshold for remediation of SCA findings related to
vulnerabilities and licenses.


**Applicability:** maturity-3

**Recommendation**: Document a policy in the project that defines a threshold for
remediation of SCA findings related to vulnerabilities and licenses.
Include the process for identifying, prioritizing, and remediating
these findings.

#### OSPS-VM-05.02



While active, the project documentation MUST include a policy to
address SCA violations prior to any release.


**Applicability:** maturity-3

**Recommendation**: Document a policy in the project to address applicable Software
Composition Analysis results before any release, and add status checks
that verify compliance with that policy prior to release.

#### OSPS-VM-05.03



While active, all changes to the project's codebase MUST be
automatically evaluated against a documented policy for malicious
dependencies and known vulnerabilities in dependencies, then blocked
in the event of violations, except when declared and suppressed as
non-exploitable.


**Applicability:** maturity-3

**Recommendation**: Create a status check in the project's version control system that
runs a Software Composition Analysis tool on all changes
to the codebase. Require that the status check passes before changes
can be merged.


### OSPS-VM-06: Publish and Enforce an Application Security Testing Policy




**Objective**

Identify and address defects and security weaknesses in the project's
codebase early in the development process, reducing the risk of shipping
insecure software.


#### Guidelines

This control aids in the application of the following guidelines:

- **800-161**: CA-7 · RA-5 · SA-11 · SI-2 · SI-3

- **BPB**: B-S-8 · Q-B-12 · Q-S-9 · S-B-14 · S-B-15 · A-B-1 · A-B-3 · A-B-8 · A-S-1

- **BSI-TR-03185-2**: QA.05

- **CRA**: 1.2a · 1.2b · 1.2c · 2.1 · 2.2 · 2.3 · 2.4

- **CSF**: GV.RM-05 · GV.RM-06 · GV.PO-01 · GV.PO-02 · ID.RA-01 · ID.RA-08 · ID.IM-02

- **ISO-18974**: 4.1.5 · 4.2.1 · 4.2.2 · 4.3.2

- **OpenCRE**: 155-155 · 124-564 · 757-271 · 464-513 · 611-158 · 207-435 · 088-377

- **PCIDSS**: 6.2.3 · 6.3.1 · 6.3.2 · 6.4.1 · 6.4.2 · 6.5.2

- **PSSCRM**: G5.4 · P4.1 · P4.2 · P4.3 · P4.4 · P4.5

- **SAMM**: Implementation -Secure Build-Build Process Lvl3 · Implementation -Software Dependencies Lvl3 · Verification -Security Testing -Scalable Baseline Lvl1 · Verification -Security Testing -Scalable Baseline Lvl3

- **Scorecard**: Security-Policy · Vulnerabilities · SAST

- **SSDF**: PO.4 · PW.1.2 · PW.8.1 · RV.1.2 · RV.1.3 · RV.2.1 · RV 2.2

- **UKSSCOP**: 1.3 · 1.4



#### OSPS-VM-06.01



While active, the project documentation MUST include a policy that
defines a threshold for remediation of SAST findings.


**Applicability:** maturity-3

**Recommendation**: Document a policy in the project that defines a threshold for
remediation of Static Application Security Testing (SAST) findings.
Include the process for identifying, prioritizing, and remediating
these findings.

#### OSPS-VM-06.02



While active, all changes to the project's codebase MUST be
automatically evaluated against a documented policy for security
weaknesses and blocked in the event of violations except when declared
and suppressed as non-exploitable.


**Applicability:** maturity-3

**Recommendation**: Create a status check in the project's version control system that
runs a Static Application Security Testing (SAST) tool on all changes
to the codebase. Require that the status check passes before changes
can be merged.




### Mapping References

- **800-161** — NIST Special Publication 800-161 - Cybersecurity Supply Chain Risk Management Practices for Systems and Organizations (vr1-upd1) — [link](https://nvlpubs.nist.gov/nistpubs/SpecialPublications/NIST.SP.800-161r1-upd1.pdf)

  This publication provides guidance to organizations on identifying, assessing, and mitigating cybersecurity risks throughout the supply chain at all levels of their organizations. The publication integrates cybersecurity supply chain risk management (C-SCRM) into risk management activities by applying a multilevel, C-SCRM-specific approach, including guidance on the development of C-SCRM strategy implementation plans, C-SCRM policies, C-SCRM plans, and risk assessments for products and services.
- **BPB** — OpenSSF Best Practices Badge (v2024) — [link](https://github.com/coreinfrastructure/best-practices-badge/blob/main/criteria/criteria.yml)

  The Open Source Security Foundation (OpenSSF) Best Practices Badge is a way for Free/Libre and Open Source Software (FLOSS) projects to show that they follow best practices. Projects can voluntarily self-certify, at no cost, by using this web application to explain how they follow each best practice. The OpenSSF Best Practices Badge is inspired by the many badges available to projects on GitHub. Consumers of the badge can quickly assess which FLOSS projects are following best practices and, as a result, are more likely to produce higher-quality secure software.
- **BSI-TR-03185-2** — BSI TR-03185-2 Secure Software Lifecycle for Open Source Software (vv1.1.0) — [link](https://www.bsi.bund.de/SharedDocs/Downloads/EN/BSI/Publications/TechGuidelines/TR03185/BSI-TR-03185-2.pdf?__blob=publicationFile&v=5)

  Secure software is essential for the use of IT products in governments, businesses and societies.
- **CRA** — Cyber Resilience Act (v20.11.2024) — [link](https://eur-lex.europa.eu/legal-content/EN/TXT/HTML/?uri=OJ:L_202402847#tit_1)

  Regulation (EU) 2024/2847 of the European Parliament and of the Council of 23 October 2024 on horizontal cybersecurity requirements for products with digital elements and amending Regulations (EU) No 168/2013 and (EU) 2019/1020 and Directive (EU) 2020/1828 (Cyber Resilience Act) (Text with EEA relevance)
- **CSAG** — CISA Software Acquisition Guide (v2024-08-01) — [link](https://www.cisa.gov/resources-tools/resources/software-acquisition-guide-government-enterprise-consumers-software-assurance-cyber-supply-chain)

  The Software Acquisition Guide for Government Enterprise Consumers: Software Assurance in the Cyber-Supply Chain Risk Management (C-SCRM) Lifecycle product was developed in response to the core challenges of software assurance and cybersecurity transparency in the acquisition process, focusing primarily on software lifecycle activities.
- **CSbDP** — CISA Secure by Design Pledge (v2024-05-08) — [link](https://www.cisa.gov/sites/default/files/2024-05/CISA%20Secure%20by%20Design%20Pledge_508c.pdf)

  A voluntary pledge focused on seven goals to work towards, in addition to context and example approaches to achieve the goal and demonstrate measurable progress within enterprise software products and services.
- **CSF** — NIST Cybersecurity Framework (v2.0) — [link](https://nvlpubs.nist.gov/nistpubs/CSWP/NIST.CSWP.29.pdf)

  The NIST Cybersecurity Framework (CSF) 2.0 provides guidance to industry, government agencies, and other organizations to manage cybersecurity risks. It offers a taxonomy of high level cybersecurity outcomes that can be used by any organization — regardless of its size, sector, or maturity — to better understand, assess, prioritize, and communicate its cybersecurity efforts. The CSF does not prescribe how outcomes should be achieved. Rather, it links to online resources that provide additional guidance on practices and controls that could be used to achieve those outcomes.
- **DORA** — EU Digital Operational Resilience Act (DORA) (v2022-12-14) — [link](https://eur-lex.europa.eu/legal-content/EN/TXT/PDF/?uri=CELEX:32022R2554&from=FR)

  On digital operational resilience for the financial sector and amending Regulations (EC) No 1060/2009, (EU) No 648/2012, (EU) No 600/2014, (EU) No 909/2014 and (EU) 2016/1011.
- **MAF** — MITRE ATT&CK Framework (vv18) — [link](https://attack.mitre.org/)

  A globally-accessible knowledge base of adversary tactics and techniques based on real-world observations.
- **NIS2** — EU Network and Information Security Directive 2 (v2024-10-17) — [link](https://eur-lex.europa.eu/legal-content/EN/TXT/HTML/?uri=OJ:L_202402690#tit_1)

  Laying down rules for the application of Directive (EU) 2022/2555 as regards technical and methodological requirements of cybersecurity risk-management measures and further specification of the cases in which an incident is considered to be significant with regard to DNS service providers, TLD name registries, cloud computing service providers, data centre service providers, content delivery network providers, managed service providers, managed security service providers, providers of online market places, of online search engines and of social networking services platforms, and trust service providers.
- **PCIDSS** — Payment Card Industry Data Security Standard (v4.0.1) — [link](https://docs-prv.pcisecuritystandards.org/PCI%20DSS/Standard/PCI-DSS-v4_0_1.pdf)

  PCI Security Standards are technical and operational requirements set by the PCI Security Standards Council (PCI SSC) to protect cardholder data. The standards apply to all entities that store, process or transmit cardholder data – with requirements for software developers and manufacturers of applications and devices used in those transactions. The Council is responsible for managing the security standards, while compliance with the PCI set of standards is enforced by the founding members of the Council: American Express, Discover Financial Services, JCB, MasterCard and Visa Inc. The PCI Data Security Standard (PCI DSS) applies to all entities that store, process, and/or transmit cardholder data. It covers technical and operational system components included in or connected to cardholder data. If you accept or process payment cards, PCI DSS applies to you.
- **PSSCRM** — Proactive Software Supply Chain Risk Management Framework (v1.0) — [link](https://arxiv.org/pdf/2404.12300)

  The Proactive-Software Supply Chain Risk Management (P-SSCRM) Framework is designed to help you understand and plan a secure software supply chain risk management initiative. P-SSCRM was created through a process of understanding and analyzing real-world data from nine industry-leading software supply chain risk management initiatives as well as through the analysis and unification of ten government and industry documents, frameworks, and standards. Although individual methodologies and standards differ, many initiatives and standards share common ground. P-SSCRM describes this common ground and presents a model for understanding, quantifying, and developing a secure software supply chain risk management program and determining where your organization’s existing efforts stand when contrasted with other real-world software supply chain risk management initiatives.
- **SAMM** — OWASP Software Assurance Maturity Model (v2.0) — [link](https://owaspsamm.org/model/)

  The mission of OWASP Software Assurance Maturity Model (SAMM) is to be the prime maturity model for software assurance that provides an effective and measurable way for all types of organizations to analyze and improve their software security posture. OWASP SAMM supports the complete software lifecycle, including development and acquisition, and is technology and process agnostic. It is intentionally built to be evolutive and risk-driven in nature.
- **Scorecard** — OpenSSF Scorecard (vv5.2.1) — [link](https://scorecard.dev/)
- **SLSA** — Supply Chain Levels for Software Artifacts (v1.0) — [link](https://github.com/slsa-framework/slsa)

  SLSA (pronounced "salsa") is a security framework from source to service, giving anyone working with software a common language for increasing levels of software security and supply chain integrity. It’s how you get from safe enough to being as resilient as possible, at any link in the chain.
- **SSDF** — Secure Software Development Framework (v1.1) — [link](https://csrc.nist.gov/pubs/sp/800/218/final)

  The Secure Software Development Framework (SSDF) is a set of fundamental, sound, and secure software development practices based on established secure software development practice documents from organizations such as BSA, OWASP, and SAFECode. Few software development life cycle (SDLC) models explicitly address software security in detail, so practices like those in the SSDF need to be added to and integrated with each SDLC implementation. Following the SSDF practices should help software producers reduce the number of vulnerabilities in released software, reduce the potential impact of the exploitation of undetected or unaddressed vulnerabilities, and address the root causes of vulnerabilities to prevent recurrences. Also, because the SSDF provides a common language for describing secure software development practices, software producers and acquirers can use it to foster their communications for procurement processes and other management activities.
- **UKSSCOP** — United Kingdom National Cyber Security Centre Software Security Code of Practice (v2025-05-07) — [link](https://www.ncsc.gov.uk/guidance/software-security-code-of-practice-assurance-principles-claims)

  The Software Code of Practice has been created by DSIT and the National Cyber Security Centre (NCSC), the UK’s technical authority for cyber security, and is co-sealed by the Canadian Centre for Cyber Security (CCCS). The Code reflects the government’s ongoing focus on codifying minimum standards for technology providers to reduce cyber risk. It is aimed at professionals who are responsible for overseeing the development of ‘commodity’ software, including technical, compliance, and risk experts. For those organisations that require a higher level of assurance in the resilience of their connected products and technology, consider using the NCSC’s Cyber Resilience Testing scheme.
- **USCTM** — US Cyber Trust Mark (v2023-07-18) — [link](https://www.fcc.gov/CyberTrustMark)

  A voluntary cybersecurity labeling program for wireless consumer IoT products.


