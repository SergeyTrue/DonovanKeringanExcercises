# Exercise 3.8
Rendering fractals at high zoom levels demands great arithmetic precision.
Implement the same fractal using four different representations
of numbers: `complex64`, `complex128`, `big.Float`, and `big.Rat`.
(The latter two types are found in the `math/big` package.
`Float` uses arbitrary but bounded-precision floating-point;
`Rat` uses unbounded-precision rational numbers.)
How do they compare in performance and memory usage?
At what zoom levels do rendering artifacts become visible?
## Mandelbrot 64-bit
![Mandelbrot 64-bit](3BasicDataTypes/Excercise_3_8_FractalBenchmarking/mandelbrot64.png)

## Mandelbrot 128-bit
![Mandelbrot 128-bit](3BasicDataTypes/Excercise_3_8_FractalBenchmarking/mandelbrot128.png)

## Mandelbrot BigFloat
![Mandelbrot BigFloat](3BasicDataTypes/Excercise_3_8_FractalBenchmarking/mandelbrotBigFloat.png)