# LeetCode Solutions

A personal collection of LeetCode solutions, primarily in Go, with Python and C available when useful.

## Layout

Keep each problem in its own directory under the language you used:

```text
go/0001-two-sum/solution.go
python/0001-two-sum/solution.py
c/0001-two-sum/solution.c
```

Use the zero-padded LeetCode problem number and a short, lowercase, hyphenated title. A problem directory may contain its own tests or supporting files. Different languages stay separate; there is no required shared API or build system across them.

## Go

Each problem directory is an independent Go package. After adding at least one Go solution, run all Go tests from the repository root:

```sh
go test ./...
```

Use the package name `solution` (or a short valid name) in each problem directory. Keep any local test harness and helper types alongside that problem.

## Python and C

Use standard-library-only Python unless a problem needs otherwise. Keep C solutions and any problem-specific headers or test harnesses together in the problem directory. Compile C solutions with the compiler and harness appropriate to that problem.

## Key Learnings

### Hash Map

- 'frequency array' as the hash key (go/0049-group-anagrams)
-  
