package slide

import (
	"github.com/jackvalmadre/go-cv/rimg64"
	"github.com/jackvalmadre/go-fftw/fftw"

	"image"
	"math/cmplx"
)

// Returns the number of positions such that the template g lies entirely inside the image f.
func outputSize(f, g image.Point) image.Point {
	var h image.Point
	h.X = max(f.X-g.X+1, 0)
	h.Y = max(f.Y-g.Y+1, 0)
	return h
}

// Computes approximate expense of each approach, return true if Fourier is better.
func useFourier(f, g image.Point) bool {
	h := outputSize(f, g)
	// One dot product per output pixel.
	naive := h.X * h.Y * g.X * g.Y
	// Two forward transforms and an inverse transform.
	fourier := 3 * f.X * f.Y * logb(f.X*f.Y)
	return fourier < naive
}

// Computes correlation of template g with image f.
//
// Takes inner product of g with f at all positions such that it lies entirely within f.
//
// Automatically selects between naive and Fourier-domain convolution.
func Corr(f, g *rimg64.Image) *rimg64.Image {
	size := outputSize(f.Size(), g.Size())
	// Return empty image if that's the result.
	if size.Eq(image.ZP) {
		return nil
	}
	// Switch implementation based on image size.
	if useFourier(f.Size(), g.Size()) {
		return corrFFT(f, g)
	}
	return corrNaive(f, g)
}

func corrNaive(f, g *rimg64.Image) *rimg64.Image {
	size := outputSize(f.Size(), g.Size())
	// Return empty image if that's the result.
	if size.Eq(image.ZP) {
		return nil
	}

	h := rimg64.New(size.X, size.Y)
	for i := 0; i < size.X; i++ {
		for j := 0; j < size.Y; j++ {
			var total float64
			for u := 0; u < g.Width; u++ {
				for v := 0; v < g.Height; v++ {
					total += f.At(i+u, j+v) * g.At(u, v)
				}
			}
			h.Set(i, j, total)
		}
	}
	return h
}

// Computes correlation of template g with image f.
func corrFFT(f, g *rimg64.Image) *rimg64.Image {
	size := outputSize(f.Size(), g.Size())
	// Return empty image if that's the result.
	if size.Eq(image.ZP) {
		return nil
	}

	x := fftw.NewArray2(f.Width, f.Height)
	y := fftw.NewArray2(f.Width, f.Height)
	// Copy into FFT arrays.
	for u := 0; u < f.Width; u++ {
		for v := 0; v < f.Height; v++ {
			x.Set(u, v, complex(f.At(u, v), 0))
		}
	}
	for u := 0; u < g.Width; u++ {
		for v := 0; v < g.Height; v++ {
			y.Set(u, v, complex(g.At(u, v), 0))
		}
	}

	// Take forward transforms.
	x = fftw.FFT2(x)
	y = fftw.FFT2(y)
	// Multiply in Fourier domain.
	for u := 0; u < f.Width; u++ {
		for v := 0; v < f.Height; v++ {
			x.Set(u, v, x.At(u, v)*cmplx.Conj(y.At(u, v)))
		}
	}
	// Take inverse transform.
	x = fftw.IFFT2(x)

	// Extract desired region.
	h := rimg64.New(size.X, size.Y)
	// Scale such that convolution theorem holds.
	n := float64(f.Width) * float64(f.Height)
	for u := 0; u < size.X; u++ {
		for v := 0; v < size.Y; v++ {
			h.Set(u, v, real(x.At(u, v))/n)
		}
	}
	return h
}