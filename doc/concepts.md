# UVM

## Execution Graph

Node: 

* ID
* Type
* Named Directed Edges
* Flags
  - Pure?
  - Lazy/Eager

Node Instance:

* Running?
* State (Deferred/Accepted/Rejected)
* Input

Block:

* Name
* Function

## Execution

1. Wait for all inputs to arrive if the block is lazy
2. Run block
3. Signal output edges with result

