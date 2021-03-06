package slide

import (
	"fmt"

	"github.com/jvlmdr/go-cv/rimg64"
)

// Algo identifies an algorithm.
type Algo int

const (
	Auto Algo = iota
	Naive
	FFT
	BLAS
)

func CorrAlgo(f, g *rimg64.Image, algo Algo) (*rimg64.Image, error) {
	switch algo {
	case Auto:
		return CorrAuto(f, g)
	case Naive:
		return CorrNaive(f, g)
	case FFT:
		return CorrFFT(f, g)
	case BLAS:
		return CorrBLAS(f, g)
	default:
		panic(fmt.Sprintf("unsupported algorithm: %v", algo))
	}
}

func CorrStrideAlgo(f, g *rimg64.Image, stride int, algo Algo) (*rimg64.Image, error) {
	switch algo {
	case Naive:
		return CorrStrideNaive(f, g, stride)
	case FFT:
		return CorrStrideFFT(f, g, stride)
	case BLAS:
		return CorrStrideBLAS(f, g, stride)
	default:
		panic(fmt.Sprintf("unsupported algorithm: %v", algo))
	}
}

func CorrBankAlgo(f *rimg64.Image, g *Bank, algo Algo) (*rimg64.Multi, error) {
	switch algo {
	case Naive:
		return CorrBankNaive(f, g)
	case FFT:
		return CorrBankFFT(f, g)
	case BLAS:
		return CorrBankBLAS(f, g)
	default:
		panic(fmt.Sprintf("unsupported algorithm: %v", algo))
	}
}

func CorrBankStrideAlgo(f *rimg64.Image, g *Bank, stride int, algo Algo) (*rimg64.Multi, error) {
	switch algo {
	case Naive:
		return CorrBankStrideNaive(f, g, stride)
	case FFT:
		return CorrBankStrideFFT(f, g, stride)
	case BLAS:
		return CorrBankStrideBLAS(f, g, stride)
	default:
		panic(fmt.Sprintf("unsupported algorithm: %v", algo))
	}
}

func CorrMultiAlgo(f, g *rimg64.Multi, algo Algo) (*rimg64.Image, error) {
	switch algo {
	case Auto:
		return CorrMultiAuto(f, g)
	case Naive:
		return CorrMultiNaive(f, g)
	case FFT:
		return CorrMultiFFT(f, g)
	case BLAS:
		return CorrMultiBLAS(f, g)
	default:
		panic(fmt.Sprintf("unsupported algorithm: %v", algo))
	}
}

func CorrMultiStrideAlgo(f, g *rimg64.Multi, stride int, algo Algo) (*rimg64.Image, error) {
	switch algo {
	case Naive:
		return CorrMultiStrideNaive(f, g, stride)
	case FFT:
		return CorrMultiStrideFFT(f, g, stride)
	case BLAS:
		return CorrMultiStrideBLAS(f, g, stride)
	default:
		panic(fmt.Sprintf("unsupported algorithm: %v", algo))
	}
}

func CorrMultiBankAlgo(f *rimg64.Multi, g *MultiBank, algo Algo) (*rimg64.Multi, error) {
	switch algo {
	case Naive:
		return CorrMultiBankNaive(f, g)
	case FFT:
		return CorrMultiBankFFT(f, g)
	case BLAS:
		return CorrMultiBankBLAS(f, g)
	default:
		panic(fmt.Sprintf("unsupported algorithm: %v", algo))
	}
}

func CorrMultiBankStrideAlgo(f *rimg64.Multi, g *MultiBank, stride int, algo Algo) (*rimg64.Multi, error) {
	switch algo {
	case Naive:
		return CorrMultiBankStrideNaive(f, g, stride)
	case FFT:
		return CorrMultiBankStrideFFT(f, g, stride)
	case BLAS:
		return CorrMultiBankStrideBLAS(f, g, stride)
	default:
		panic(fmt.Sprintf("unsupported algorithm: %v", algo))
	}
}
