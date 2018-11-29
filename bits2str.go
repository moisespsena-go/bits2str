package bits2str

const (
	Bit  Bits = 1
	Byte      = 8 * Bit
	// https://en.wikipedia.org/wiki/Orders_of_magnitude_(data)
	KB = 1000 * Byte
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
	PB = 1000 * TB
	EB = 1000 * PB
)

// Bits represents a quantity of Bits, bytes, kilobytes or megabytes. Bits are
// parsed and formatted using the IEEE / SI standards, which use multiples of
// 1000 to represent kilobytes and megabytes (instead of multiples of 1024). For
// more information see https://en.wikipedia.org/wiki/Megabyte#Definitions.
type Bits int64

// Bytes returns the size as a floating point number of bytes.
func (b Bits) Bytes() float64 {
	bytes := b / Byte
	bits := b % Byte
	return float64(bytes) + float64(bits)/8
}

// Kilobytes returns the size as a floating point number of kilobytes.
func (b Bits) Kilobytes() float64 {
	bytes := b / KB
	bits := b % KB
	return float64(bytes) + float64(bits)/(8*1000)
}

// Megabytes returns the size as a floating point number of megabytes.
func (b Bits) Megabytes() float64 {
	bytes := b / MB
	bits := b % MB
	return float64(bytes) + float64(bits)/(8*1000*1000)
}

// Gigabytes returns the size as a floating point number of gigabytes.
func (b Bits) Gigabytes() float64 {
	bytes := b / GB
	bits := b % GB
	return float64(bytes) + float64(bits)/(8*1000*1000*1000)
}

// String returns a string representation of b using the largest unit that has a
// positive number before the decimal. At most three decimal places of precision
// are printed.
func (b Bits) String() string {
	if b == 0 {
		return "0"
	}
	// Largest value is "-123.150EB"
	var buf [10]byte
	w := len(buf) - 1
	u := uint64(b)
	neg := b < 0
	if neg {
		u = -u
	}
	if u < uint64(Byte) {
		w -= 2
		copy(buf[w:], "Bit")
		w = fmtInt(buf[:w], u)
	} else {
		switch {
		case u < uint64(KB):
			w -= 0
			buf[w] = 'B'
			u = (u * 1e3 / 8)
		case u < uint64(MB):
			w -= 1
			copy(buf[w:], "kB")
			u /= 8
		case u < uint64(GB):
			w -= 1
			copy(buf[w:], "MB")
			u /= 8 * 1e3
		case u < uint64(TB):
			w -= 1
			copy(buf[w:], "GB")
			u /= 8 * 1e6
		case u < uint64(PB):
			w -= 1
			copy(buf[w:], "TB")
			u /= 8 * 1e9
		case u < uint64(EB):
			w -= 1
			copy(buf[w:], "PB")
			u /= 8 * 1e12
		case u >= uint64(EB):
			w -= 1
			copy(buf[w:], "EB")
			u /= 8 * 1e15
		}
		w, u = fmtFrac(buf[:w], u, 3)
		w = fmtInt(buf[:w], u)
	}
	if neg {
		w--
		buf[w] = '-'
	}
	return string(buf[w:])
}
