# TDD

Basic procedure:

1. Write a failing test (red/yellow)
2. Change the implementation to make it pass
3. Refactor
4. Go to Step 1 with the next requirement

## Detroid (or "Classic" or "bottom-up") vs. London (or "Mockist" or "top-down")

The difference between Detroid and London concerns the direction of development.

### Detroid

In the Detroid model, the end result is assembled from individual sub-steps. Each individual step is developed and tested in a unit (e.g. a function). The final result is then the sum of the individual parts and is verified at the very end with an integration test.

This procedure is suitable for tasks where the requirements lead step by step to the final result. This is almost always the case with the first small katas.

### London

In the London model, the end result of a unit is directly given. In practice, this is very often the case with interfaces (REST, gRPC, ...).

The first step is to write a test that asks for the final result. Then you go down from top to bottom to the individual sub-requirements and implement them.

Since the partial requirements are not available for the integration test at the beginning, dummies are used. Depending on the technology used, these can be empty objects, mocks or stubs.