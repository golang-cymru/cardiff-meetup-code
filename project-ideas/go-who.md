# Proposal: go-who - github org user management

Author(s): J. Gregory

## Abstract

A Go implementation of a tool similar to [gu-who](https://github.com/guardian/gu-who) but in a more understandable language and more easily configured package

## Background

As organisations grow, managing memebers in their github org becomes difficult so a tool to easily manage it is desirable.
The `gu-who` tool provides this in a fairly opinionated manner though is not overly friendly to extend or alter it's rulesets.

## Proposal

Create a new tool (in Go), using the github graphql api, borrowing from the good points of `gu-who` but enabling easier maintenance and configurable rules.

## Rationale

The existing tool is written in Scala and as such is fairly inpenetratable to anyone not expert in that lanaguage. It's rulesets are also highly opinionated and not easy to configure.

## Implementation (if applicable)

Use the github graphql api
