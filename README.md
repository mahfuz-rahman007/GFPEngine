
PROJECT 1 — Concurrent File Processing Engine (GFPEngine)
Build a CLI that recursively scans directories, analyzes files, calculates hashes, detects duplicates, and processes work concurrently. The purpose is to turn your existing Go syntax knowledge into real understanding.
Target Architecture
CLI → Scanner → Jobs Channel → Worker Pool → Results Channel → Aggregator → Console/JSON Report
Build Checklist

1. ☐ CLI foundation
   Create `goprocess <directory>` with worker count and output options. Validate input and return useful errors.
   Learn/understand: main, packages, modules, flag, os, filepath, error returns.
   ```Markdown
   goprocess ./data
   goprocess ./data --workers 8
   goprocess ./data --output report.json
   ```
2. ☐ Directory scanner
   Walk nested directories and create a FileInfo model containing path, size, extension and modification time.
   Learn/understand: structs, methods, slices, filesystem APIs
3. ☐ Processor abstraction
   Create a Processor interface and multiple implementations such as hash/text/metadata processors.
   Learn/understand: interfaces, composition, dependency injection, pointer/value receivers
4. ☐ Streaming hashing
   Hash files without loading entire files into memory.
   Learn/understand: io.Reader, io.Writer, buffers, file handles, crypto/sha256
5. ☐ Duplicate detection
   Group files by hash and produce duplicate reports.
   Learn/understand: maps, equality, aggregation
6. ☐ Concurrent worker pool
   Send file jobs through channels to multiple workers and collect results.
   Learn/understand: goroutines, channels, WaitGroup, buffered channels
7. ☐ Cancellation
   Cancel processing with context and handle Ctrl+C cleanly.
   Learn/understand: context, signal handling, lifecycle ownership
8. ☐ Race investigation
   Introduce shared state, run the race detector, then fix the design.
   Learn/understand: mutex, synchronization, race detector
9. ☐ Tests and benchmarks
   Add unit tests, table-driven tests, concurrency tests and benchmarks.

Final Architecture
                  CLI
                     │
                     ↓
                File Scanner
                     │
                     ↓
                Job Producer
                     │
                     ↓
               jobs channel
              /      |
             ↓       ↓       ↓
          Worker   Worker   Worker
             │       │       │
        └──┼──┘
                     ↓
              results channel
                     │
                     ↓
                 Aggregator
                     │
             ┌─┴───┐
             ↓               ↓
          Console           JSON
           Report           Report
