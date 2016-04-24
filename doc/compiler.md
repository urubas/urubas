# Compiler

The compiler should execute the following tasks:

1. Verify the integrity of the DAG
2. Construct the runtime representation of all types
3. Generate LLVM IR for each function thunk(see also: code generation)
4. Build a runtime representation of all exported(external linkage and parallel execution) functions
