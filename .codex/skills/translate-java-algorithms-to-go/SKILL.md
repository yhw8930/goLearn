---
name: translate-java-algorithms-to-go
description: Translate Java algorithm lesson code from /Users/yhw/IdeaProjects/algorithmbasic/src into Go under /Users/yhw/go/src/goLearn/algorithm/lession. Use when the user asks to translate, port, migrate, or supplement Java classXX algorithm lessons into Go classXX folders, especially with Chinese explanations, algorithm ideas, and time/space complexity notes.
---

# Translate Java Algorithms To Go

## Overview

Translate Java algorithm course classes into idiomatic Go files for this repo while preserving the lesson structure. Add concise Chinese explanations and time/space complexity comments for each algorithm.

## Project Paths

- Java source root: `/Users/yhw/IdeaProjects/algorithmbasic/src`
- Go target root: `/Users/yhw/go/src/goLearn/algorithm/lession`
- Work from repo root: `/Users/yhw/go/src/goLearn`

For a request like “把 Java class18 翻译到 Go class18”, read Java files from:

```text
/Users/yhw/IdeaProjects/algorithmbasic/src/class18
```

and write Go files into:

```text
/Users/yhw/go/src/goLearn/algorithm/lession/class18
```

## Workflow

1. Inspect the Java class folder with `rg --files` or `find`.
2. Inspect nearby Go lesson folders to match local style before writing code.
3. Create the destination `classXX` folder if missing.
4. Translate each Java file into a corresponding Go file.
5. Keep shared structures in focused files when multiple algorithms need them, for example `graph.go`.
6. Add Chinese comments explaining the core idea and complexity.
7. Run `gofmt` on changed Go files.
8. Run a scoped compile check:

```bash
env GOCACHE=/private/tmp/go-build-cache go test ./algorithm/lession/classXX
```

Use `/private/tmp/go-build-cache` because the default Go build cache may be outside the sandbox.

## Go Style

- Use `package main`, matching the existing lesson folders.
- Prefer exported function names for algorithm entry points, such as `Dijkstra1`, `Hanoi2`, `Permutation3`.
- When translating or updating a single Java lesson file, preserve the Java `main` example as a Go `func main()` so the user can directly run that file with `go run path/to/File.go`.
- Add a concise Chinese problem statement comment immediately above `func main()`, describing the input, constraints, and required output.
- When translating a whole lesson folder, be aware that multiple `main` functions in the same `package main` prevent package-level `go test`; still keep per-file runnable examples when the Java source has `main`, and validate individual files when needed.
- Avoid generic package-level helper names such as `dp`, `process`, `f`, or `g` in lesson folders with multiple files. Prefix them with the problem name, for example `knapsackDP`, `knapsackProcess`, or `convertToLetterStringDP`, so files in the same package do not collide.
- Use Go slices as stacks/queues unless a standard library heap is clearly useful.
- Use `container/heap` for priority queues.
- Use `map[*Node]struct{}` as a set.
- Keep names close to the Java source when that helps course comparison, but use Go casing.
- Avoid broad refactors of unrelated lesson folders.

## Comment Requirements

For each main algorithm function, include:

- What problem it solves.
- The key idea in one or two Chinese sentences.
- Important preconditions, such as “边权非负” or “有向无环图”.
- Time complexity.
- Space complexity.

For `func main()` examples, include a Chinese problem statement comment above `main` rather than only describing the demo data.

Preferred format:

```go
// Dijkstra1 求 from 到所有可达点的最短距离，要求边权非负。
// 朴素版本每轮从 distanceMap 中线性找一个未确定的最小距离点。
//
// 时间复杂度：O(V^2+E)。
// 空间复杂度：O(V)。
func Dijkstra1(from *Node) map[*Node]int {
```

## Translation Guidance

- Java `ArrayList<T>` -> Go `[]T`.
- Java `HashMap<K,V>` -> Go `map[K]V`.
- Java `HashSet<T>` -> Go `map[T]struct{}`.
- Java `Stack<T>` -> Go slice with tail as stack top.
- Java `Queue<T>` -> Go slice queue for lessons; use index or pop front depending on simplicity.
- Java `PriorityQueue` -> Go `container/heap`.
- Java inner classes can become small Go structs in the same file.
- If Java duplicates graph structures across classes, duplicate them in the matching Go class folder unless the repo already has a shared package.

## Verification

After edits, run:

```bash
gofmt -w /Users/yhw/go/src/goLearn/algorithm/lession/classXX
env GOCACHE=/private/tmp/go-build-cache go test ./algorithm/lession/classXX
```

Report whether validation passed. If unrelated dirty files exist in git status, mention them but do not modify or revert them.
