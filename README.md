Hex color string will be silently parsed as `color.Color`

``` go
// color.RGBA{ 0xab, 0xc1, 0x23, 0xff }
hexcolor.Parse("#AbC123") 
hexcolor.Parse("AbC123") 

// color.RGBA{ 0xaa, 0xbb, 0xcc, 0xff }
hexcolor.Parse("#aBc")
hexcolor.Parse("abC") 

// color.RGBA{ 0xab, 0xc1, 0x23, 0xff }
hexcolor.Parse("#abc123FF")
hexcolor.Parse("Abc123FF") 

// invalid format will return hexcolor.Default, which is magenta
// color.RGBA{ 0xff, 0x00, 0xff, 0xff }
hexcolor.Parse("invalid format")
```

*Note: No errors will be returned, whether the input is valid or not. Invalid inputs will return the default color (Magenta `color.RGBA{0xff, 0x00, 0xff, 0xff}`), but you should never expect this.*