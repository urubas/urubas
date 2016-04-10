# UVM

## Execution DAG

### Nodes

* ID
* Type
* Named Outputs
  - Directed Edges
* Flags
  - Lazy/Eager

### Node Instance

* Running?
* Inputs
* Outputs

### Input

* State
* Value Channel
* State Channel

### Block

* Name
* Handler
* Pure?
* Routing?

## Definitions

### Function
A DAG which takes may take inputs and returns one or more outputs.

### Routing block
Blocks that may accept or reject his edges separately.

### Branch level
The unique identifier for a execution path based on every branching that occurs in the code.

### Top-level node
Nodes that doesn't have any input from his branch level.
Literals are excluded from this category.

## Execution

1. Wait for all inputs to arrive if the block is lazy
2. Run block
3. Signal output edges with result

## Branching
A routing block may fork the execution context into one or more branches. Those branches can be accepted or rejected separately.

When branching, each execution path should be forked, creating new instances of each node above in the graph. Those intances should only differ in their inputs as folow:

* If the input comes from a parent branching level, use the same instance
* Otherwise, it should 
