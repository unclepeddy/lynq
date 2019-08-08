## Motivation

It's no secret that most of the world has gone through a process of digitization in the past 3+ decades. Computers have taken over our physical spaces.
It's natural to think of this as a consequence of the increase in availability and decrease in price of digital computers (Moore's Law), which is correct. But over the past 2+ decades, that growth has accelerated partly because of the internet, which has allowed us to build a global virtual world by connecting all these devices together.
In the 1990s, people wondered why the internet matters. The answer often given was that "it's a distributed, decentralised, permissionless network where anyone can build applications." While it took over a decade for that promise to materialize (in the early 90's most "applications" were the unimaginative email and file transfer, with the www comprising less than 5% of the total internet traffic) nevertheless it did, leading to the Google's, Youtube's and the Facebook's of the world. We now live in a world where anyone in their garage could build an application and use the internet as a platform to reach billions of people.

Today, people wonder why a technology we call "blockchain" matters - and the answer we hear often is that some implementation of that technology will be a distributed, decentralised, permissionless network where anyone can build applications. Most people don't take this seriously because for anyone that's experimented with even the most advanced blockchain implementations, they know that while in theory it sounds useful, it looks a bit like the internet in the early '90s. Understandably, it's hard to imagine Youtube, when all you have is FTP. Similarly, when all we have is barely-working smart contracts, it's hard to imagine the world that will be possible once the technology exists. Much like we needed to spend more than a decade building the internet network layer by layer ([literally](https://en.wikipedia.org/wiki/OSI_model)), invent and re-invent browser technologies and experiences, and create web-native and cloud-native programming paradigms, we will go through an evolution of abstraction layers on this new stack. The only difference is that now instead of software, we'll have [software 2.0](https://medium.com/@karpathy/software-2-0-a64152b37c35) to program it with. 

ML apps on the blockchain.

Most of people's lives is coming online.
Most of those companies are building software 1.0
This means most of the data is going to be owned by software 1.0 services.
As the technology for building the new type of software becomes more capable, we will start building some services on it.
There will be a technological divide and whole new ways of addressing it.
Same way today almost most industries are becoming coomputerized and already-dizitized ones are moving onto the cloud, we will have communities and applications (companies) needing to offer services on the new internet, and an opportunity for bridging two worlds. 

This is where lynq comes in.


## Vision

Lynq is a software ecosystem built for developers powered by developers. Lynq transpiles into itself, meaning source code written by users can inter-operate with those written by developers. Lynq programs are executed on lynq's cloud-based runtime environment.

The language for developers is an abstraction over swift, with an abstraction layer over it provided to the end users which probabilistically transpiles into lynq.

A special part of the language is the standard library it comes with. It offers developers (and sometimes end-users) the ability to easily create their own Lynq operators, which allow for interactions with new services.

The API server is implemented on top of Kubernetes and all the computation scheduling, state storage and reconciliation is done via its machinery. 

The end user writes and submits lynq progams to the collection of their programs, stored as a git repository.

The files kept in the repository declare desired state about the user's world that are actualized via a set of Lynq operators, which is an abstraction layer on top of kubernetes operators. The Lynq operator framework is optimized for ease of use, but allows inter-operability with [one of] Kubernetes operator framework for advanced use cases.

Lynq is a language for expressing intent about the state of your world. While it's possible to use it to query about and change the state of your bedroom light, you will most likely want to use it for managing more complex interaction of information in your digital world. Today, that's what emails you receive, what phone and text conversations you have, and what pictures you take. Soon, it will be transcript of everything you heard and saw during the day, including what you may have missed or forgotten. We will have the desire to asynchronously digest that information, and know certain things about it - Lynq allows us to do this.

## Tech Strategy

We'll do a lot of prototyping - in order to not waste a lot of effort and time, we'll pick what we'll be intentional about and what we don't mind scratching.
We suspect that the following capabilities will emerge as important aspects of the design of the system.

1. A language to declare an invariant over the state of resources owned by disparate digital services
2. A state storage interface for keeping invariants on resources owned by arbitrary services (CRD definitions)
3. A flexible interface for scheduling continual/periodic/one-shot execution of functions/containers
4. A mechanism for continually observing and reacting to the state of resources elsewhere on the internet

## Roadmap

1. For 1 provider (Calendar), I can decleratively manage 1 type of resource
2. For 2 providers (Calendar, FB Messanger), I can decleratively manage any types of resources
3. I can run a script to do a 1-time sync of the birthdays of my last K LRU contacts on messanger to my Calendar
4. I can decleratively keep my last K LRU contacts on messanger synced to my Calendar (will be enforced until sync deleted)
5. I can declare a simple relation between state managed by my FB Messanger and Calendar instances.
