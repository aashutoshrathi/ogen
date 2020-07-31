// Code generated by fastssz. DO NOT EDIT.
package bls

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the Multipub object
func (m *Multipub) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(m)
}

// MarshalSSZTo ssz marshals the Multipub object to a target array
func (m *Multipub) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(12)

	// Offset (0) 'PublicKeys'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(m.PublicKeys) * 48

	// Field (1) 'NumNeeded'
	dst = ssz.MarshalUint64(dst, m.NumNeeded)

	// Field (0) 'PublicKeys'
	if len(m.PublicKeys) > 32 {
		err = ssz.ErrListTooBig
		return
	}
	for ii := 0; ii < len(m.PublicKeys); ii++ {
		dst = append(dst, m.PublicKeys[ii][:]...)
	}

	return
}

// UnmarshalSSZ ssz unmarshals the Multipub object
func (m *Multipub) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 12 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'PublicKeys'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	// Field (1) 'NumNeeded'
	m.NumNeeded = ssz.UnmarshallUint64(buf[4:12])

	// Field (0) 'PublicKeys'
	{
		buf = tail[o0:]
		num, err := ssz.DivideInt2(len(buf), 48, 32)
		if err != nil {
			return err
		}
		m.PublicKeys = make([][48]byte, num)
		for ii := 0; ii < num; ii++ {
			copy(m.PublicKeys[ii][:], buf[ii*48:(ii+1)*48])
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Multipub object
func (m *Multipub) SizeSSZ() (size int) {
	size = 12

	// Field (0) 'PublicKeys'
	size += len(m.PublicKeys) * 48

	return
}

// HashTreeRoot ssz hashes the Multipub object
func (m *Multipub) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(m)
}

// HashTreeRootWith ssz hashes the Multipub object with a hasher
func (m *Multipub) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'PublicKeys'
	{
		if len(m.PublicKeys) > 32 {
			err = ssz.ErrListTooBig
			return
		}
		subIndx := hh.Index()
		for _, i := range m.PublicKeys {
			hh.Append(i[:])
		}
		numItems := uint64(len(m.PublicKeys))
		hh.MerkleizeWithMixin(subIndx, numItems, ssz.CalculateLimit(32, numItems, 32))
	}

	// Field (1) 'NumNeeded'
	hh.PutUint64(m.NumNeeded)

	hh.Merkleize(indx)
	return
}

// MarshalSSZ ssz marshals the Multisig object
func (m *Multisig) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(m)
}

// MarshalSSZTo ssz marshals the Multisig object to a target array
func (m *Multisig) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(12)

	// Offset (0) 'PublicKey'
	dst = ssz.WriteOffset(dst, offset)
	if m.PublicKey == nil {
		m.PublicKey = new(Multipub)
	}
	offset += m.PublicKey.SizeSSZ()

	// Offset (1) 'Signatures'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(m.Signatures) * 96

	// Offset (2) 'KeysSigned'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(m.KeysSigned)

	// Field (0) 'PublicKey'
	if dst, err = m.PublicKey.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'Signatures'
	if len(m.Signatures) > 32 {
		err = ssz.ErrListTooBig
		return
	}
	for ii := 0; ii < len(m.Signatures); ii++ {
		dst = append(dst, m.Signatures[ii][:]...)
	}

	// Field (2) 'KeysSigned'
	if len(m.KeysSigned) > 33 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, m.KeysSigned...)

	return
}

// UnmarshalSSZ ssz unmarshals the Multisig object
func (m *Multisig) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 12 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o1, o2 uint64

	// Offset (0) 'PublicKey'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	// Offset (1) 'Signatures'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return ssz.ErrOffset
	}

	// Offset (2) 'KeysSigned'
	if o2 = ssz.ReadOffset(buf[8:12]); o2 > size || o1 > o2 {
		return ssz.ErrOffset
	}

	// Field (0) 'PublicKey'
	{
		buf = tail[o0:o1]
		if m.PublicKey == nil {
			m.PublicKey = new(Multipub)
		}
		if err = m.PublicKey.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (1) 'Signatures'
	{
		buf = tail[o1:o2]
		num, err := ssz.DivideInt2(len(buf), 96, 32)
		if err != nil {
			return err
		}
		m.Signatures = make([][96]byte, num)
		for ii := 0; ii < num; ii++ {
			copy(m.Signatures[ii][:], buf[ii*96:(ii+1)*96])
		}
	}

	// Field (2) 'KeysSigned'
	{
		buf = tail[o2:]
		if err = ssz.ValidateBitlist(buf, 33); err != nil {
			return err
		}
		m.KeysSigned = append(m.KeysSigned, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Multisig object
func (m *Multisig) SizeSSZ() (size int) {
	size = 12

	// Field (0) 'PublicKey'
	if m.PublicKey == nil {
		m.PublicKey = new(Multipub)
	}
	size += m.PublicKey.SizeSSZ()

	// Field (1) 'Signatures'
	size += len(m.Signatures) * 96

	// Field (2) 'KeysSigned'
	size += len(m.KeysSigned)

	return
}

// HashTreeRoot ssz hashes the Multisig object
func (m *Multisig) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(m)
}

// HashTreeRootWith ssz hashes the Multisig object with a hasher
func (m *Multisig) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'PublicKey'
	if err = m.PublicKey.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'Signatures'
	{
		if len(m.Signatures) > 32 {
			err = ssz.ErrListTooBig
			return
		}
		subIndx := hh.Index()
		for _, i := range m.Signatures {
			hh.Append(i[:])
		}
		numItems := uint64(len(m.Signatures))
		hh.MerkleizeWithMixin(subIndx, numItems, ssz.CalculateLimit(32, numItems, 32))
	}

	// Field (2) 'KeysSigned'
	hh.PutBitlist(m.KeysSigned, 33)

	hh.Merkleize(indx)
	return
}
