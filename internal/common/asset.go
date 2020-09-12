package common

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image"
	"image/png"
)

type AssetType byte

const (
	AssetTypeImage AssetType = 1 << iota
	AssetTypeAtlas
	AssetTypeAnimation
	AssetTypeSound
)

type Asset struct {
	Type AssetType

	Img      image.Image
	AtlasMap map[string]AtlasRegion
}

const invalidAssetType = "Invalid asset type: %X"

func NewAsset() *Asset {
	return &Asset{
		AtlasMap: make(map[string]AtlasRegion, 0),
	}
}

func (a *Asset) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)

	buf.WriteByte(byte(a.Type))

	switch a.Type {
	case AssetTypeAtlas:
		buf.WriteByte(byte(len(a.AtlasMap)))

		for k, v := range a.AtlasMap {
			buf.WriteString(k)
			buf.WriteByte(0)

			data := make([]byte, 8)
			binary.LittleEndian.PutUint16(data, v.X)
			binary.LittleEndian.PutUint16(data, v.Y)
			binary.LittleEndian.PutUint16(data, v.W)
			binary.LittleEndian.PutUint16(data, v.H)

			buf.Write(data)
		}

	case AssetTypeAnimation:
		// TODO
	case AssetTypeSound:
		// TODO
	default:
		panic(fmt.Sprintf(invalidAssetType, byte(a.Type)))
	}

	if a.Type != AssetTypeSound {
		if err := png.Encode(buf, a.Img); err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

func (a *Asset) UnmarshalBinary(data []byte) error {
	buf := bytes.NewBuffer(data)

	t, err := buf.ReadByte()
	if err != nil {
		return err
	}

	switch AssetType(t) {
	case AssetTypeAtlas:
		count, err := buf.ReadByte()
		if err != nil {
			return err
		}

		for i := 0; i < int(count); i++ {
			k, err := buf.ReadString(0)
			if err != nil {
				return err
			}

			regData := make([]byte, 8)
			if n, err := buf.Read(regData); n != len(regData) {
				if err != nil {
					return err
				}
				return fmt.Errorf("Expected %d bytes, got %d", len(regData), n)
			}

			a.AtlasMap[k] = AtlasRegion{
				X: binary.LittleEndian.Uint16(regData[:2]),
				Y: binary.LittleEndian.Uint16(regData[2:4]),
				W: binary.LittleEndian.Uint16(regData[4:6]),
				H: binary.LittleEndian.Uint16(regData[6:8]),
			}
		}

	case AssetTypeAnimation:
	case AssetTypeSound:
	default:
		panic(fmt.Sprintf(invalidAssetType, t))
	}

	if AssetType(t) == AssetTypeSound {
		// TODO
		return nil
	}

	a.Img, err = png.Decode(buf)

	return err
}
