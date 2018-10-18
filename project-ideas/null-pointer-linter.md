# Proposal: Identify null pointer exception via linter

Author(s): James Pudney

## Abstract

Customise a linter to identify potential null pointer exceptions

## Background

Linters are already used to identify code smells such as shadowing variables. Null pointer exceptions can be annoying, especially if the code is not in a hot/happy path and you stumble into the issue in production, though it might suggest that there's lacking tests.

## Proposal

Extend a linter to add support for null pointer exceptions.

## Rationale

Try to root out problems before they happen.

## Implementation (if applicable)

…
