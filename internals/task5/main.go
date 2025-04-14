package task5

// Interface 2D
type Hitung2D interface {
	Luas() float64
	Keliling() float64
}

// Interface 3D
type Hitung3D interface {
	Volume() float64
}

// Interface gabungan
type Hitung interface {
	Hitung2D
	Hitung3D
}

// Struct Persegi Panjang
type PersegiPanjang struct {
	Panjang float64
	Lebar   float64
}

// Implementasi Hitung2D
func (p *PersegiPanjang) Luas() float64 {
	return p.Panjang * p.Lebar
}

func (p *PersegiPanjang) Keliling() float64 {
	return 2 * (p.Panjang + p.Lebar)
}

// Struct Balok
type Kubus struct {
	Panjang float64
	Lebar   float64
	Tinggi  float64
}

// Implementasi Hitung (Hitung2D + Hitung3D)
func (b *Kubus) Luas() float64 {
	return b.Panjang * b.Lebar
}

func (b *Kubus) Keliling() float64 {
	return 2 * (b.Panjang + b.Lebar)
}

func (b *Kubus) Volume() float64 {
	return b.Panjang * b.Lebar * b.Tinggi
}
