// +build !cl12

package opencl

func (k *Kernel) ArgName(index int) (string, error) {
	return "", ErrUnsupported
}
