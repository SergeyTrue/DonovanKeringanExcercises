# Exercise 3.8
Rendering fractals at high zoom levels demands great arithmetic precision.
Implement the same fractal using four different representations
of numbers: `complex64`, `complex128`, `big.Float`, and `big.Rat`.
(The latter two types are found in the `math/big` package.
`Float` uses arbitrary but bounded-precision floating-point;
`Rat` uses unbounded-precision rational numbers.)
How do they compare in performance and memory usage?
At what zoom levels do rendering artifacts become visible?

## Usage

You can run the program in two modes:

1. **Image generation and saving to files.**
2. **Running a web server for dynamic image generation.**

### Image Generation

By default, the program will generate four images of different Mandelbrot fractal versions and save them in the current directory.

```bash
go run fractal.go
```
## Running the Web Server

To run the web server, use the following command:

```bash
go run fractal.go server
```
The server will be started on port 8000. You can specify fractal parameters via the URL.

Example Request
```
http://localhost:8000/?scale=11.05&fn=128&centerX=-.140009&centerY=0.887
```
Parameters:
centerX (float64) — the center along the X-axis.
centerY (float64) — the center along the Y-axis.
scale (float64) — the scale of the image.
fn (string) — the method of forming the fractal:
64 — using complex numbers with floating-point (float32).
128 — using complex numbers with floating-point (float64).
bf — using big numbers big.Float.
br — using rational numbers big.Rat.


# Mandelbrot 64-bit
![Mandelbrot 64-bit](3BasicDataTypes/Excercise_3_8_FractalBenchmarking/mandelbrot64.png)

# Mandelbrot 128-bit
![Mandelbrot 128-bit](3BasicDataTypes/Excercise_3_8_FractalBenchmarking/mandelbrot128.png)

# Mandelbrot BigFloat
![Mandelbrot BigFloat](3BasicDataTypes/Excercise_3_8_FractalBenchmarking/mandelbrotBigFloat.png)