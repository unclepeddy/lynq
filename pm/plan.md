# Project Plan

lynq is developed using a iterative methodology, with each iteration (phase) lasting about 1 month.

During each iteration, our goals are to 
* Learn from other technologies in adjacent spaces that interest us
* Have fun building infrastructure and when valuable, make it re-usable
* Develop end-to-end working software to validate some hypotheses

---

## Phase 1

The overall aim in this phase is to use Terraform to sync state from one application to another on an ongoing basis.

O: Build re-usable infrastructure
- KR: Build 3 service clients (Spotify, Seatgeek, Google Calendar)

O: Learn from other technologies
- KR: Build 2 Terraform providers
- KR: Build automation on top of Terraform

O: Programmatically interact with third-party resources
- KR: Create Google Calendar events based on Spotify listening history and Seatgeek data

Lessons learned:
* Terraform does a number of things really well
  * For end-users, regardless of what providers you use, your tf configuration file looks the same
  * Composability and sharing: variables, modules, private and public repositories - it has the pieces needed for a healthy ecosystem
  * For developers, it provides nice abstractions (eg. schema library) that make writing providers extremely easy
* Terraform has a number of shortcomings we should avoid
  * A few magical things (eg. associating resource types with a provider by taking the first word of the resource type name, separated by underscores)
  * It remembers a view of the world and tries to match it with a newer state using resource identifiers; when immutable, this works fine; but when the IDs change, or the order in which they appear changes across syncs, it has a tough time reconciling the changes.
  * Methods of procuring and using user and application credentials differ widely across services

## Phase 2

The overall aim in this phase is to use Kubernetes-native constructs to create resources based on other resources.

O: Build re-usable infrastructure
- KR: Build 1 service client (Google Photos)

O: Learn from other technologies
- KR: Define CRDs for photos and albums in Google Photos
- KR: Build a Kubernetes operator (kubebuilder) using best practices 

O:  Programmatically interact with third-party resources
- KR: Continuously add photos to an album based on a criteria defined on the metadata of photos added to users' library
