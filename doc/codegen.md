# Code Generation

All functions will be compiled into equivalent machine code.

Functions with multiple parallel execution branches are broken into multiple pieces called `thunks`. Each thunk corresponds to a single function and a stack frame/context when executed. When the thunk ends it outputs its value into the graph execution manager whichs feeds all connected inputs with it.

## Thunk Spliting

If a function has multiple Execution Thread IDs defined in it's nodes, it should be splitted in several thunks.
