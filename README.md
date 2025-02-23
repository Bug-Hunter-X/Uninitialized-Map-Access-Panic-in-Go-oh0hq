# Uninitialized Map Access Panic in Go

This repository demonstrates a common error in Go: panicking due to accessing an uninitialized map.  The `bug.go` file shows the problematic code, while `bugSolution.go` provides a corrected version.

## Problem

Attempting to assign a value to a map that has not been initialized will result in a runtime panic.  This is because Go maps are not automatically initialized; you must explicitly create them before use.

## Solution

The solution is to initialize the map using the `make()` function before accessing its elements.  The corrected code in `bugSolution.go` demonstrates this crucial step.

## How to run

1. Clone the repository: `git clone [repository_url]`
2. Navigate to the repository directory: `cd [repository_name]`
3. Run the buggy code (expect a panic): `go run bug.go`
4. Run the corrected code: `go run bugSolution.go`