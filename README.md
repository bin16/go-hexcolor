### func Parse(s string) color.NRGBA

Hex color string will be silently parsed as `color.NRGBA`

``` go
// color.NRGBA{ 0xab, 0xc1, 0x23, 0xff }
hexcolor.Parse("#AbC123") 
hexcolor.Parse("AbC123") 

// color.NRGBA{ 0xaa, 0xbb, 0xcc, 0xff }
hexcolor.Parse("#aBc")
hexcolor.Parse("abC") 

// color.NRGBA{ 0xab, 0xc1, 0x23, 0xff }
hexcolor.Parse("#abc123FF")
hexcolor.Parse("Abc123FF") 

// invalid format will return hexcolor.Default, which is magenta
// color.NRGBA{ 0xff, 0x00, 0xff, 0xff }
hexcolor.Parse("invalid format")
```

*Note: No errors will be returned, whether the input is valid or not. Invalid inputs will return the default color (Default `color.NRGBA{0x00, 0x00, 0x00, 0x00}`), but you should never expect this.*

### func SafeParse(s string) color.NRGBA

`Parse()` is fast, but for some invalid colors, `Parse()` may return unexpected result.

`SafeParse()` will check if input is valid first.

``` go
hexcolor.Parse("#12345X")     // -> not sure
hexcolor.SafeParse("#12345X") // -> Default color.NRGBA{}
```
