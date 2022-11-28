# colors

Minimalistic terminal colors for Go that work together - inner color doesn't reset surrounding ones.

![Example output](/docs/assets/example_output.png)

## Usage

```go
text := colors.New(colors.Codes["red"], colors.Codes["bold"]).Format("Red and Bold")
```
