# imerge

`imerge` is a small Go package for merging potentially 
large amounts of possibly overlapping integer intervals.

The intervals are held in a binary tree.
Merges occur upon each insert if possible. 
The output is deferred until explicitly asked for,
eliminating the need for unnecessary sorts in case
of updates.

This provides an advantage compared to slurping the
entire data set at once, if a large amount of overlapping intervals
can be reasonably expected.

# Example usage: 

```
// Create a new instance
n, err := NewNode(11, 15)

// Add an interval (will be merged if overlaps found)
err = n.Merge(42, 56)

// Output the intervals as a slice of slices
intervals := n.Intervals()
```

See also `cmd/merge/main.go`

# Building Command from Source

The package includes a demo command `merge`. To build it,
a Go compiler is required. The package was tested against 
`v1.13`, but any modern Go version should work.

```
git clone https://github.com/tvendelin/imerge.git
cd imerge
go build -o merge cmd/merge/main.go
```

Run

```
./merge
```

and follow the on-screen instructions.

