package ec2b

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
)

const (
	N          = 312
	M          = 156
	MATRIX_A   = 0xB5026F5AA96619E9
	UPPER_MASK = 0xFFFFFFFF80000000
	LOWER_MASK = 0x7FFFFFFF
)

type source struct {
	array [312]uint64 // state vector
	index uint64      // array index
}

func NewRand() *rand.Rand        { return rand.New(NewSource64()) }
func NewRand64() *rand.Rand      { return rand.New(NewSource64()) }
func NewSource() rand.Source     { return &source{index: N + 1} }
func NewSource64() rand.Source64 { return &source{index: N + 1} }

func (s *source) Seed(seed int64) {
	s.array[0] = uint64(seed)
	for s.index = 1; s.index < N; s.index++ {
		s.array[s.index] = (0x5851F42D4C957F2D * (s.array[s.index-1] ^ (s.array[s.index-1] >> 62)) + s.index)
	}
}

func (s *source) Int63() int64 {
	return int64(s.Uint64() & 0x7FFFFFFFFFFFFFFE)
}

func (s *source) Uint64() uint64 {
	var i int
	var x uint64
	magic := []uint64{0, MATRIX_A}
	if s.index >= N {
		if s.index == N+1 {
			s.Seed(int64(5489))
		}
		for i = 0; i < N-M; i++ {
			x = (s.array[i] & UPPER_MASK) | (s.array[i+1] & LOWER_MASK)
			s.array[i] = s.array[i+(M)] ^ (x >> 1) ^ magic[int(x&uint64(1))]
		}
		for ; i < N-1; i++ {
			x = (s.array[i] & UPPER_MASK) | (s.array[i+1] & LOWER_MASK)
			s.array[i] = s.array[i+(M-N)] ^ (x >> 1) ^ magic[int(x&uint64(1))]
		}
		x = (s.array[N-1] & UPPER_MASK) | (s.array[0] & LOWER_MASK)
		s.array[N-1] = s.array[M-1] ^ (x >> 1) ^ magic[int(x&uint64(1))]
		s.index = 0
	}
	x = s.array[s.index]
	s.index++
	x ^= (x >> 29) & 0x5555555555555555
	x ^= (x << 17) & 0x71D67FFFEDA60000
	x ^= (x << 37) & 0xFFF7EEE000000000
	x ^= (x >> 43)
	return x
}

type KeyBlock struct {
	seed uint64
	data [4096]byte
}

func NewKeyBlock(seed uint64) *KeyBlock {
	b := &KeyBlock{seed: seed}
	r := NewRand()
	r.Seed(int64(b.seed))
	r.Seed(int64(r.Uint64()))
	r.Uint64()
	for i := 0; i < 4096>>3; i++ {
		binary.BigEndian.PutUint64(b.data[i<<3:], r.Uint64())
	}
	return b
}

func (b *KeyBlock) Seed() uint64 {
	return b.seed
}

func (b *KeyBlock) Xor(data []byte) {
	for i := 0; i < len(data); i++ {
		data[i] ^= b.data[i%4096]
	}
}

type Ec2b struct {
	key  []byte
	data []byte
	seed uint64
	temp []byte
}

func LoadKey(b []byte) (*Ec2b, error) {
	if len(b) < 4+4+16+4+2048 {
		return nil, fmt.Errorf("invalid ec2b key")
	}
	if string(b[0:4]) != "Ec2b" {
		return nil, fmt.Errorf("invalid ec2b key")
	}
	keyLen := binary.LittleEndian.Uint32(b[4:])
	if keyLen != 16 {
		return nil, fmt.Errorf("invalid ec2b key")
	}
	dataLen := binary.LittleEndian.Uint32(b[24:])
	if dataLen != 2048 {
		return nil, fmt.Errorf("invalid ec2b key")
	}
	e := &Ec2b{
		key:  b[8:24],
		data: b[28 : 28+2048],
	}
	e.init()
	return e, nil
}

func NewEc2b() *Ec2b {
	e := &Ec2b{
		key:  make([]byte, 16),
		data: make([]byte, 2048),
	}
	rand.Read(e.key)
	rand.Read(e.data)
	e.init()
	return e
}

func (e *Ec2b) init() {
	k := make([]byte, 16)
	copy(k[:], e.key)
	keyScramble(k)
	e.SetSeed(getSeed(k, e.data))
}

func (e *Ec2b) Bytes() []byte {
	b := make([]byte, 4+4+16+4+2048)
	copy(b[0:4], []byte("Ec2b"))
	binary.LittleEndian.PutUint32(b[4:], 16)
	copy(b[8:], e.key)
	binary.LittleEndian.PutUint32(b[24:], 2048)
	copy(b[28:], e.data)
	return b
}

func (e *Ec2b) SetSeed(seed uint64) {
	e.seed = seed
	r := ec2b.NewRand64()
	r.Seed(int64(e.seed))
	e.temp = make([]byte, 4096)
	for i := 0; i < 4096>>3; i++ {
		binary.LittleEndian.PutUint64(e.temp[i<<3:], r.Uint64())
	}
}

func (e *Ec2b) Seed() uint64 {
	return e.seed
}

func (e *Ec2b) Key() []byte {
	b := make([]byte, 4+4+16+4+2048)
	copy(b[0:4], []byte("Ec2b"))
	binary.LittleEndian.PutUint32(b[4:], 16)
	copy(b[8:], e.key)
	binary.LittleEndian.PutUint32(b[24:], 2048)
	copy(b[28:], e.data)
	return b
}

func (e *Ec2b) Xor(data []byte) {
	for i := 0; i < len(data); i++ {
		data[i] ^= e.temp[i%4096]
	}
}

func keyScramble(key []byte) {
	var roundKeys [11][16]byte
	for r := 0; r < 11; r++ {
		for i := 0; i < 16; i++ {
			for j := 0; j < 16; j++ {
				idx := (r << 8) + (i << 4) + j
				roundKeys[r][i] ^= aesXorTable[1][idx] ^ aesXorTable[0][idx]
			}
		}
	}
	xorRoundKey(key, roundKeys[0][:])
	for r := 1; r < 10; r++ {
		subBytesInv(key)
		shiftRowsInv(key)
		mixColsInv(key)
		xorRoundKey(key, roundKeys[r][:])
	}
	subBytesInv(key)
	shiftRowsInv(key)
	xorRoundKey(key, roundKeys[10][:])
	for i := 0; i < 16; i++ {
		key[i] ^= keyXorTable[i]
	}
}

func xorRoundKey(key, roundKey []byte) {
	for i := 0; i < 16; i++ {
		key[i] ^= roundKey[i]
	}
}

func subBytes(key []byte) {
	for i := 0; i < 16; i++ {
		key[i] = lookupSbox[key[i]]
	}
}

func subBytesInv(key []byte) {
	for i := 0; i < 16; i++ {
		key[i] = lookupSboxInv[key[i]]
	}
}

func shiftRows(key []byte) {
	var temp [16]byte
	copy(temp[:], key[:])
	for i := 0; i < 16; i++ {
		key[i] = temp[shiftRowsTable[i]]
	}
}

func shiftRowsInv(key []byte) {
	var temp [16]byte
	copy(temp[:], key[:])
	for i := 0; i < 16; i++ {
		key[i] = temp[shiftRowsTableInv[i]]
	}
}

func mixColInv(key []byte) {
	a0, a1, a2, a3 := key[0], key[1], key[2], key[3]
	key[0] = lookupG14[a0] ^ lookupG9[a3] ^ lookupG13[a2] ^ lookupG11[a1]
	key[1] = lookupG14[a1] ^ lookupG9[a0] ^ lookupG13[a3] ^ lookupG11[a2]
	key[2] = lookupG14[a2] ^ lookupG9[a1] ^ lookupG13[a0] ^ lookupG11[a3]
	key[3] = lookupG14[a3] ^ lookupG9[a2] ^ lookupG13[a1] ^ lookupG11[a0]
}

func mixColsInv(key []byte) {
	mixColInv(key[0:])
	mixColInv(key[4:])
	mixColInv(key[8:])
	mixColInv(key[12:])
}

func getSeed(key, data []byte) uint64 {
	v := ^uint64(0xCEAC3B5A867837AC)
	v ^= binary.LittleEndian.Uint64(key[0:])
	v ^= binary.LittleEndian.Uint64(key[8:])
	for i := 0; i < len(data)>>3; i++ {
		v ^= binary.LittleEndian.Uint64(data[i<<3:])
	}
	return v
}
