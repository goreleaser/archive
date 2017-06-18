# archive

A simple Go archiving library.

Example:

```go
file, err := os.Create("file.zip")
if err != nil {
  // deal with the error
}
archive := archive.New(file)
defer archive.Close()
archive.Add("file.txt", "/path/to/file.txt")
```
