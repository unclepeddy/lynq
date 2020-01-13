# lynq

## Mission

Enable everyone to declaratively organize, edit and create content

## Summary

Simply going about our everyday lives, we create troves of information (photos, videos and audio recorded via our phone, content posted and viewed on social media, messages sent and received, songs listened to, websites visited, etc.), but most people are unable to programmatically reason about and manipulate this data. Although these pieces of data are represented using well-defined schemas and exposed by services that own them, most people have to rely on the service provider to do any post-processing, organization and creation of derivative artifacts based on this information (edited photos, time capsules, personalized playlists, etc.). lynq envisions a world where ordinary people are empowered to use content they interact with on these platforms on their own terms to organize or create whatever derivative artifacts they want, subject to their own creativity. In a world where our information is immensely valuable, so is this ability to easily manipulate, augment and build on top of that information. lynq allows multiple personas to, at multiple levels of abstraction, implement machinery that will intelligently manipulate and build on top of one’s information.

## Design

### Personas: chefs, gardeners, and farmers

**Chefs**: Most lynq end-users are chefs, who use natural language to build `recipes` from already-made `ingredients`. Once `recipes` are written, they are compiled into lynq’s DSL and if valid, they can be activated to continuously operate over the user’s information until turned off.

**Gardeners**: Users who build ingredients by adding functionality to a lynq provider are called Gardeners. Gardeners use lynq’s DSL to write reusable components that can understand, reason about and manipulate `Things`.  

**Farmers**: The group of software developers that implement the actual machinery that interacts with external services, called `Providers`. Farmers define the lynq-internal data model for a Provider and use already existing or new API clients to implement the `Thing` API spec for every resource of the service that is to be represented. 

Everyone can be a chef since the interface is natural language. most chefs with some reasonable programming experience should be able to transition to be gardeners; farmers should be able to bring their existing work into the ecosystem. 

## Concepts

### Provider

A provider is a component that abstracts away an external service. It provides a few functionalities for the service it represents:
* Defines a lynq-internal, versioned schema (a `Thing`) for each user-visible artifact exposed by that service (photos, messages, etc.)
* Expose some set of standardized APIs for interacting with those objects to the rest of lynq ecosystem

### Thing

Each `Thing` is a type with a canonical schema that is owned by a single provider, representing a resource owned by a third-party digital service. A Provider that provides a particular Thing exposes a standard set of APIs to the rest of lynq ecosystem and is resonponsible for implementing the machinery that can talk to an external service to actualize actions performed using those APIs. This means that although most Things are only ever accessed by a single provider, it is possible for provider A to ask B to manipulate or reason about things that are provided by provider B: for example the Calendar provider can create Events (Things provided by Google Calendar Provider) based on Photos (Things provided by Google Photos Provider) you have taken at certain locations.

### Ingredient

Ingredients are composable transformations defined over Things. An ingredient is simply a function that parameterizes the creation or mutation of a Thing. An ingredient takes as input zero or more Things, zero or more configuration parameters and yields one or more Things and thus can be thought of as a source of Things. 

Ingredients can be used to import Things into the ecosystem, create Things based on other things, filter a stream of incoming Things, rank groups of Things, and much more.
Ingredients are defined and implemented by Gardeners using Things exposed by one or more Providers via the lynq DSL and used by Chefs to create Recipes.

### Recipe

Recipes can in a way be thought of as the main abstraction in lynq: the programs written by the end-user (chefs). These recipes are written by composing ingredients provided by installed Providers. Some example recipes:
* Apply certain types of edits to my photos after I take them
* Add certain photos to an album based on some conditions (who is in them, where or when they were taken, or any other set of conditions)
* Tag, restructure and summarize text (on notes, Google documents, conversations created by me)

When recipes are turned on, they create derivative `Things` on some recurring basis until turned off.
