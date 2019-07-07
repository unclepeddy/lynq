## Vision

## Motivation

## Tech Strategy

We'll do a lot of prototyping - in order to not waste a lot of effort and time, we'll pick what we'll be intentional about and what we don't mind scratching.
We suspect that the following capabilities will emerge as important aspects of the design of the system.

1. A language to declare an invariant over the state of resources owned by disparate digital services
2. A state storage interface for keeping invariants on resources owned by arbitrary services (CRD definitions)
3. A flexible interface for scheduling continual/periodic/one-shot execution of functions/containers
4. A mechanism for continually observing and reacting to the state of resources elsewhere on the internet
5. A daemon for hyper-user-task-specific speech-to-text models

## Roadmap

1. For 1 provider (Calendar), I can decleratively manage 1 type of resource
2. For 2 provider (Calendar, FB Messanger), I can decleratively manage any types of resources
3. I can run a script to do a 1-time sync of the birthdays of my last K LRU contacts on messanger to my Calendar
4. I can decleratively keep my last K LRU contacts on messanger synced to my Calendar (will be enforced until sync deleted)
5. I can declare a simple relation between state managed by my FB Messanger and Calendar instances.
