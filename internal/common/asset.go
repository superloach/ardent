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

	AnimWidth, AnimHeight uint16
	AnimationMap          map[string]Animation
}

const invalidAssetType = "Invalid asset type: %X"

func NewAsset() *Asset {
	return &Asset{
		AtlasMap:     make(map[string]AtlasRegion, 0),
		AnimationMap: make(map[string]Animation, 0),
	}
}

func (a *Asset) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)

	buf.WriteString("ardent")
	buf.WriteByte(0)
	buf.WriteByte(byte(a.Type))

	switch a.Type {
	case AssetTypeImage:
	case AssetTypeAtlas:
		buf.WriteByte(byte(len(a.AtlasMap)))

		for k, v := range a.AtlasMap {
			buf.WriteString(k)
			buf.WriteByte(0)

			data := make([]byte, 8)
			binary.LittleEndian.PutUint16(data[:2], v.X)
			binary.LittleEndian.PutUint16(data[2:4], v.Y)
			binary.LittleEndian.PutUint16(data[4:6], v.W)
			binary.LittleEndian.PutUint16(data[6:8], v.H)

			buf.Write(data)
		}

	case AssetTypeAnimation:
		buf.WriteByte(byte(len(a.AnimationMap)))

		data := make([]byte, 4)
		binary.LittleEndian.PutUint16(data[:2], a.AnimWidth)
		binary.LittleEndian.PutUint16(data[2:4], a.AnimHeight)

		buf.Write(data)

		for k, v := range a.AnimationMap {
			buf.WriteString(k)
			buf.WriteByte(0)

			data = make([]byte, 6)
			binary.LittleEndian.PutUint16(data[:2], v.Fps)
			binary.LittleEndian.PutUint16(data[2:4], v.Start)
			binary.LittleEndian.PutUint16(data[4:6], v.End)

			buf.Write(data)
		}
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

	magic, err := buf.ReadString(0)
	if err != nil {
		return err
	}

	if magic[:len(magic)-1] != "ardent" {
		return fmt.Errorf("Invalid filetype")
	}

	t, err := buf.ReadByte()
	if err != nil {
		return err
	}

	a.Type = AssetType(t)
	switch a.Type {
	case AssetTypeImage:
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

			a.AtlasMap[k[:len(k)-1]] = AtlasRegion{
				X: binary.LittleEndian.Uint16(regData[:2]),
				Y: binary.LittleEndian.Uint16(regData[2:4]),
				W: binary.LittleEndian.Uint16(regData[4:6]),
				H: binary.LittleEndian.Uint16(regData[6:8]),
			}
		}

	case AssetTypeAnimation:
		count, err := buf.ReadByte()
		if err != nil {
			return err
		}

		animSize := make([]byte, 4)
		if n, err := buf.Read(animSize); n != len(animSize) {
			if err != nil {
				return err
			}
			return fmt.Errorf("Expected %d bytes, got %d", len(animSize), n)
		}

		a.AnimWidth = binary.LittleEndian.Uint16(animSize[:2])
		a.AnimHeight = binary.LittleEndian.Uint16(animSize[2:4])

		for i := 0; i < int(count); i++ {
			k, err := buf.ReadString(0)
			if err != nil {
				return err
			}

			animData := make([]byte, 6)
			if n, err := buf.Read(animData); n != len(animData) {
				if err != nil {
					return err
				}
				return fmt.Errorf("Expected %d bytes, go %d", len(animData), n)
			}

			a.AnimationMap[k[:len(k)-1]] = Animation{
				Fps:   binary.LittleEndian.Uint16(animData[:2]),
				Start: binary.LittleEndian.Uint16(animData[2:4]),
				End:   binary.LittleEndian.Uint16(animData[4:6]),
			}
		}

	case AssetTypeSound:
	default:
		panic(fmt.Sprintf(invalidAssetType, t))
	}

	if a.Type == AssetTypeSound {
		// TODO
		return nil
	}

	a.Img, err = png.Decode(buf)

	return err
}
