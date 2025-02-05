// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}

// InputFileLocation represents TL type `inputFileLocation#dfdaabe1`.
// DEPRECATED location of a photo
//
// See https://core.telegram.org/constructor/inputFileLocation for reference.
type InputFileLocation struct {
	// Server volume
	VolumeID int64
	// File identifier
	LocalID int
	// Check sum to access the file
	Secret int64
	// File reference¹
	//
	// Links:
	//  1) https://core.telegram.org/api/file_reference
	FileReference []byte
}

// InputFileLocationTypeID is TL type id of InputFileLocation.
const InputFileLocationTypeID = 0xdfdaabe1

// String implements fmt.Stringer.
func (i *InputFileLocation) String() string {
	if i == nil {
		return "InputFileLocation(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputFileLocation")
	sb.WriteString("{\n")
	sb.WriteString("\tVolumeID: ")
	sb.WriteString(fmt.Sprint(i.VolumeID))
	sb.WriteString(",\n")
	sb.WriteString("\tLocalID: ")
	sb.WriteString(fmt.Sprint(i.LocalID))
	sb.WriteString(",\n")
	sb.WriteString("\tSecret: ")
	sb.WriteString(fmt.Sprint(i.Secret))
	sb.WriteString(",\n")
	sb.WriteString("\tFileReference: ")
	sb.WriteString(fmt.Sprint(i.FileReference))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *InputFileLocation) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputFileLocation#dfdaabe1 as nil")
	}
	b.PutID(InputFileLocationTypeID)
	b.PutLong(i.VolumeID)
	b.PutInt(i.LocalID)
	b.PutLong(i.Secret)
	b.PutBytes(i.FileReference)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputFileLocation) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputFileLocation#dfdaabe1 to nil")
	}
	if err := b.ConsumeID(InputFileLocationTypeID); err != nil {
		return fmt.Errorf("unable to decode inputFileLocation#dfdaabe1: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputFileLocation#dfdaabe1: field volume_id: %w", err)
		}
		i.VolumeID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputFileLocation#dfdaabe1: field local_id: %w", err)
		}
		i.LocalID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputFileLocation#dfdaabe1: field secret: %w", err)
		}
		i.Secret = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode inputFileLocation#dfdaabe1: field file_reference: %w", err)
		}
		i.FileReference = value
	}
	return nil
}

// construct implements constructor of InputFileLocationClass.
func (i InputFileLocation) construct() InputFileLocationClass { return &i }

// Ensuring interfaces in compile-time for InputFileLocation.
var (
	_ bin.Encoder = &InputFileLocation{}
	_ bin.Decoder = &InputFileLocation{}

	_ InputFileLocationClass = &InputFileLocation{}
)

// InputEncryptedFileLocation represents TL type `inputEncryptedFileLocation#f5235d55`.
// Location of encrypted secret chat file.
//
// See https://core.telegram.org/constructor/inputEncryptedFileLocation for reference.
type InputEncryptedFileLocation struct {
	// File ID, id parameter value from encryptedFile¹
	//
	// Links:
	//  1) https://core.telegram.org/constructor/encryptedFile
	ID int64
	// Checksum, access_hash parameter value from encryptedFile¹
	//
	// Links:
	//  1) https://core.telegram.org/constructor/encryptedFile
	AccessHash int64
}

// InputEncryptedFileLocationTypeID is TL type id of InputEncryptedFileLocation.
const InputEncryptedFileLocationTypeID = 0xf5235d55

// String implements fmt.Stringer.
func (i *InputEncryptedFileLocation) String() string {
	if i == nil {
		return "InputEncryptedFileLocation(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputEncryptedFileLocation")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(i.ID))
	sb.WriteString(",\n")
	sb.WriteString("\tAccessHash: ")
	sb.WriteString(fmt.Sprint(i.AccessHash))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *InputEncryptedFileLocation) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputEncryptedFileLocation#f5235d55 as nil")
	}
	b.PutID(InputEncryptedFileLocationTypeID)
	b.PutLong(i.ID)
	b.PutLong(i.AccessHash)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputEncryptedFileLocation) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputEncryptedFileLocation#f5235d55 to nil")
	}
	if err := b.ConsumeID(InputEncryptedFileLocationTypeID); err != nil {
		return fmt.Errorf("unable to decode inputEncryptedFileLocation#f5235d55: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedFileLocation#f5235d55: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedFileLocation#f5235d55: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// construct implements constructor of InputFileLocationClass.
func (i InputEncryptedFileLocation) construct() InputFileLocationClass { return &i }

// Ensuring interfaces in compile-time for InputEncryptedFileLocation.
var (
	_ bin.Encoder = &InputEncryptedFileLocation{}
	_ bin.Decoder = &InputEncryptedFileLocation{}

	_ InputFileLocationClass = &InputEncryptedFileLocation{}
)

// InputDocumentFileLocation represents TL type `inputDocumentFileLocation#bad07584`.
// Document location (video, voice, audio, basically every type except photo)
//
// See https://core.telegram.org/constructor/inputDocumentFileLocation for reference.
type InputDocumentFileLocation struct {
	// Document ID
	ID int64
	// access_hash parameter from the document¹ constructor
	//
	// Links:
	//  1) https://core.telegram.org/constructor/document
	AccessHash int64
	// File reference¹
	//
	// Links:
	//  1) https://core.telegram.org/api/file_reference
	FileReference []byte
	// Thumbnail size to download the thumbnail
	ThumbSize string
}

// InputDocumentFileLocationTypeID is TL type id of InputDocumentFileLocation.
const InputDocumentFileLocationTypeID = 0xbad07584

// String implements fmt.Stringer.
func (i *InputDocumentFileLocation) String() string {
	if i == nil {
		return "InputDocumentFileLocation(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputDocumentFileLocation")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(i.ID))
	sb.WriteString(",\n")
	sb.WriteString("\tAccessHash: ")
	sb.WriteString(fmt.Sprint(i.AccessHash))
	sb.WriteString(",\n")
	sb.WriteString("\tFileReference: ")
	sb.WriteString(fmt.Sprint(i.FileReference))
	sb.WriteString(",\n")
	sb.WriteString("\tThumbSize: ")
	sb.WriteString(fmt.Sprint(i.ThumbSize))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *InputDocumentFileLocation) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputDocumentFileLocation#bad07584 as nil")
	}
	b.PutID(InputDocumentFileLocationTypeID)
	b.PutLong(i.ID)
	b.PutLong(i.AccessHash)
	b.PutBytes(i.FileReference)
	b.PutString(i.ThumbSize)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputDocumentFileLocation) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputDocumentFileLocation#bad07584 to nil")
	}
	if err := b.ConsumeID(InputDocumentFileLocationTypeID); err != nil {
		return fmt.Errorf("unable to decode inputDocumentFileLocation#bad07584: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputDocumentFileLocation#bad07584: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputDocumentFileLocation#bad07584: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode inputDocumentFileLocation#bad07584: field file_reference: %w", err)
		}
		i.FileReference = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputDocumentFileLocation#bad07584: field thumb_size: %w", err)
		}
		i.ThumbSize = value
	}
	return nil
}

// construct implements constructor of InputFileLocationClass.
func (i InputDocumentFileLocation) construct() InputFileLocationClass { return &i }

// Ensuring interfaces in compile-time for InputDocumentFileLocation.
var (
	_ bin.Encoder = &InputDocumentFileLocation{}
	_ bin.Decoder = &InputDocumentFileLocation{}

	_ InputFileLocationClass = &InputDocumentFileLocation{}
)

// InputSecureFileLocation represents TL type `inputSecureFileLocation#cbc7ee28`.
// Location of encrypted telegram passport¹ file.
//
// Links:
//  1) https://core.telegram.org/passport
//
// See https://core.telegram.org/constructor/inputSecureFileLocation for reference.
type InputSecureFileLocation struct {
	// File ID, id parameter value from secureFile¹
	//
	// Links:
	//  1) https://core.telegram.org/constructor/secureFile
	ID int64
	// Checksum, access_hash parameter value from secureFile¹
	//
	// Links:
	//  1) https://core.telegram.org/constructor/secureFile
	AccessHash int64
}

// InputSecureFileLocationTypeID is TL type id of InputSecureFileLocation.
const InputSecureFileLocationTypeID = 0xcbc7ee28

// String implements fmt.Stringer.
func (i *InputSecureFileLocation) String() string {
	if i == nil {
		return "InputSecureFileLocation(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputSecureFileLocation")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(i.ID))
	sb.WriteString(",\n")
	sb.WriteString("\tAccessHash: ")
	sb.WriteString(fmt.Sprint(i.AccessHash))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *InputSecureFileLocation) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputSecureFileLocation#cbc7ee28 as nil")
	}
	b.PutID(InputSecureFileLocationTypeID)
	b.PutLong(i.ID)
	b.PutLong(i.AccessHash)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputSecureFileLocation) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputSecureFileLocation#cbc7ee28 to nil")
	}
	if err := b.ConsumeID(InputSecureFileLocationTypeID); err != nil {
		return fmt.Errorf("unable to decode inputSecureFileLocation#cbc7ee28: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputSecureFileLocation#cbc7ee28: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputSecureFileLocation#cbc7ee28: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// construct implements constructor of InputFileLocationClass.
func (i InputSecureFileLocation) construct() InputFileLocationClass { return &i }

// Ensuring interfaces in compile-time for InputSecureFileLocation.
var (
	_ bin.Encoder = &InputSecureFileLocation{}
	_ bin.Decoder = &InputSecureFileLocation{}

	_ InputFileLocationClass = &InputSecureFileLocation{}
)

// InputTakeoutFileLocation represents TL type `inputTakeoutFileLocation#29be5899`.
// Empty constructor for takeout
//
// See https://core.telegram.org/constructor/inputTakeoutFileLocation for reference.
type InputTakeoutFileLocation struct {
}

// InputTakeoutFileLocationTypeID is TL type id of InputTakeoutFileLocation.
const InputTakeoutFileLocationTypeID = 0x29be5899

// String implements fmt.Stringer.
func (i *InputTakeoutFileLocation) String() string {
	if i == nil {
		return "InputTakeoutFileLocation(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputTakeoutFileLocation")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *InputTakeoutFileLocation) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputTakeoutFileLocation#29be5899 as nil")
	}
	b.PutID(InputTakeoutFileLocationTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputTakeoutFileLocation) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputTakeoutFileLocation#29be5899 to nil")
	}
	if err := b.ConsumeID(InputTakeoutFileLocationTypeID); err != nil {
		return fmt.Errorf("unable to decode inputTakeoutFileLocation#29be5899: %w", err)
	}
	return nil
}

// construct implements constructor of InputFileLocationClass.
func (i InputTakeoutFileLocation) construct() InputFileLocationClass { return &i }

// Ensuring interfaces in compile-time for InputTakeoutFileLocation.
var (
	_ bin.Encoder = &InputTakeoutFileLocation{}
	_ bin.Decoder = &InputTakeoutFileLocation{}

	_ InputFileLocationClass = &InputTakeoutFileLocation{}
)

// InputPhotoFileLocation represents TL type `inputPhotoFileLocation#40181ffe`.
// Use this object to download a photo with upload.getFile¹ method
//
// Links:
//  1) https://core.telegram.org/method/upload.getFile
//
// See https://core.telegram.org/constructor/inputPhotoFileLocation for reference.
type InputPhotoFileLocation struct {
	// Photo ID, obtained from the photo¹ object
	//
	// Links:
	//  1) https://core.telegram.org/constructor/photo
	ID int64
	// Photo's access hash, obtained from the photo¹ object
	//
	// Links:
	//  1) https://core.telegram.org/constructor/photo
	AccessHash int64
	// File reference¹
	//
	// Links:
	//  1) https://core.telegram.org/api/file_reference
	FileReference []byte
	// The PhotoSize¹ to download: must be set to the type field of the desired PhotoSize object of the photo²
	//
	// Links:
	//  1) https://core.telegram.org/type/PhotoSize
	//  2) https://core.telegram.org/constructor/photo
	ThumbSize string
}

// InputPhotoFileLocationTypeID is TL type id of InputPhotoFileLocation.
const InputPhotoFileLocationTypeID = 0x40181ffe

// String implements fmt.Stringer.
func (i *InputPhotoFileLocation) String() string {
	if i == nil {
		return "InputPhotoFileLocation(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputPhotoFileLocation")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(i.ID))
	sb.WriteString(",\n")
	sb.WriteString("\tAccessHash: ")
	sb.WriteString(fmt.Sprint(i.AccessHash))
	sb.WriteString(",\n")
	sb.WriteString("\tFileReference: ")
	sb.WriteString(fmt.Sprint(i.FileReference))
	sb.WriteString(",\n")
	sb.WriteString("\tThumbSize: ")
	sb.WriteString(fmt.Sprint(i.ThumbSize))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *InputPhotoFileLocation) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPhotoFileLocation#40181ffe as nil")
	}
	b.PutID(InputPhotoFileLocationTypeID)
	b.PutLong(i.ID)
	b.PutLong(i.AccessHash)
	b.PutBytes(i.FileReference)
	b.PutString(i.ThumbSize)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPhotoFileLocation) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPhotoFileLocation#40181ffe to nil")
	}
	if err := b.ConsumeID(InputPhotoFileLocationTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPhotoFileLocation#40181ffe: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhotoFileLocation#40181ffe: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhotoFileLocation#40181ffe: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhotoFileLocation#40181ffe: field file_reference: %w", err)
		}
		i.FileReference = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhotoFileLocation#40181ffe: field thumb_size: %w", err)
		}
		i.ThumbSize = value
	}
	return nil
}

// construct implements constructor of InputFileLocationClass.
func (i InputPhotoFileLocation) construct() InputFileLocationClass { return &i }

// Ensuring interfaces in compile-time for InputPhotoFileLocation.
var (
	_ bin.Encoder = &InputPhotoFileLocation{}
	_ bin.Decoder = &InputPhotoFileLocation{}

	_ InputFileLocationClass = &InputPhotoFileLocation{}
)

// InputPhotoLegacyFileLocation represents TL type `inputPhotoLegacyFileLocation#d83466f3`.
// Legacy photo file location
//
// See https://core.telegram.org/constructor/inputPhotoLegacyFileLocation for reference.
type InputPhotoLegacyFileLocation struct {
	// Photo ID
	ID int64
	// Access hash
	AccessHash int64
	// File reference
	FileReference []byte
	// Volume ID
	VolumeID int64
	// Local ID
	LocalID int
	// Secret
	Secret int64
}

// InputPhotoLegacyFileLocationTypeID is TL type id of InputPhotoLegacyFileLocation.
const InputPhotoLegacyFileLocationTypeID = 0xd83466f3

// String implements fmt.Stringer.
func (i *InputPhotoLegacyFileLocation) String() string {
	if i == nil {
		return "InputPhotoLegacyFileLocation(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputPhotoLegacyFileLocation")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(i.ID))
	sb.WriteString(",\n")
	sb.WriteString("\tAccessHash: ")
	sb.WriteString(fmt.Sprint(i.AccessHash))
	sb.WriteString(",\n")
	sb.WriteString("\tFileReference: ")
	sb.WriteString(fmt.Sprint(i.FileReference))
	sb.WriteString(",\n")
	sb.WriteString("\tVolumeID: ")
	sb.WriteString(fmt.Sprint(i.VolumeID))
	sb.WriteString(",\n")
	sb.WriteString("\tLocalID: ")
	sb.WriteString(fmt.Sprint(i.LocalID))
	sb.WriteString(",\n")
	sb.WriteString("\tSecret: ")
	sb.WriteString(fmt.Sprint(i.Secret))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *InputPhotoLegacyFileLocation) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPhotoLegacyFileLocation#d83466f3 as nil")
	}
	b.PutID(InputPhotoLegacyFileLocationTypeID)
	b.PutLong(i.ID)
	b.PutLong(i.AccessHash)
	b.PutBytes(i.FileReference)
	b.PutLong(i.VolumeID)
	b.PutInt(i.LocalID)
	b.PutLong(i.Secret)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPhotoLegacyFileLocation) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPhotoLegacyFileLocation#d83466f3 to nil")
	}
	if err := b.ConsumeID(InputPhotoLegacyFileLocationTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPhotoLegacyFileLocation#d83466f3: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhotoLegacyFileLocation#d83466f3: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhotoLegacyFileLocation#d83466f3: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhotoLegacyFileLocation#d83466f3: field file_reference: %w", err)
		}
		i.FileReference = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhotoLegacyFileLocation#d83466f3: field volume_id: %w", err)
		}
		i.VolumeID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhotoLegacyFileLocation#d83466f3: field local_id: %w", err)
		}
		i.LocalID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhotoLegacyFileLocation#d83466f3: field secret: %w", err)
		}
		i.Secret = value
	}
	return nil
}

// construct implements constructor of InputFileLocationClass.
func (i InputPhotoLegacyFileLocation) construct() InputFileLocationClass { return &i }

// Ensuring interfaces in compile-time for InputPhotoLegacyFileLocation.
var (
	_ bin.Encoder = &InputPhotoLegacyFileLocation{}
	_ bin.Decoder = &InputPhotoLegacyFileLocation{}

	_ InputFileLocationClass = &InputPhotoLegacyFileLocation{}
)

// InputPeerPhotoFileLocation represents TL type `inputPeerPhotoFileLocation#27d69997`.
// Location of profile photo of channel/group/supergroup/user
//
// See https://core.telegram.org/constructor/inputPeerPhotoFileLocation for reference.
type InputPeerPhotoFileLocation struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to download the high-quality version of the picture
	Big bool
	// The peer whose profile picture should be downloaded
	Peer InputPeerClass
	// Volume ID from FileLocation¹ met in the profile photo container.
	//
	// Links:
	//  1) https://core.telegram.org/type/FileLocation
	VolumeID int64
	// Local ID from FileLocation¹ met in the profile photo container.
	//
	// Links:
	//  1) https://core.telegram.org/type/FileLocation
	LocalID int
}

// InputPeerPhotoFileLocationTypeID is TL type id of InputPeerPhotoFileLocation.
const InputPeerPhotoFileLocationTypeID = 0x27d69997

// String implements fmt.Stringer.
func (i *InputPeerPhotoFileLocation) String() string {
	if i == nil {
		return "InputPeerPhotoFileLocation(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputPeerPhotoFileLocation")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(i.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(fmt.Sprint(i.Peer))
	sb.WriteString(",\n")
	sb.WriteString("\tVolumeID: ")
	sb.WriteString(fmt.Sprint(i.VolumeID))
	sb.WriteString(",\n")
	sb.WriteString("\tLocalID: ")
	sb.WriteString(fmt.Sprint(i.LocalID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *InputPeerPhotoFileLocation) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPeerPhotoFileLocation#27d69997 as nil")
	}
	b.PutID(InputPeerPhotoFileLocationTypeID)
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputPeerPhotoFileLocation#27d69997: field flags: %w", err)
	}
	if i.Peer == nil {
		return fmt.Errorf("unable to encode inputPeerPhotoFileLocation#27d69997: field peer is nil")
	}
	if err := i.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputPeerPhotoFileLocation#27d69997: field peer: %w", err)
	}
	b.PutLong(i.VolumeID)
	b.PutInt(i.LocalID)
	return nil
}

// SetBig sets value of Big conditional field.
func (i *InputPeerPhotoFileLocation) SetBig(value bool) {
	if value {
		i.Flags.Set(0)
		i.Big = true
	} else {
		i.Flags.Unset(0)
		i.Big = false
	}
}

// Decode implements bin.Decoder.
func (i *InputPeerPhotoFileLocation) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPeerPhotoFileLocation#27d69997 to nil")
	}
	if err := b.ConsumeID(InputPeerPhotoFileLocationTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPeerPhotoFileLocation#27d69997: %w", err)
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputPeerPhotoFileLocation#27d69997: field flags: %w", err)
		}
	}
	i.Big = i.Flags.Has(0)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerPhotoFileLocation#27d69997: field peer: %w", err)
		}
		i.Peer = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerPhotoFileLocation#27d69997: field volume_id: %w", err)
		}
		i.VolumeID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerPhotoFileLocation#27d69997: field local_id: %w", err)
		}
		i.LocalID = value
	}
	return nil
}

// construct implements constructor of InputFileLocationClass.
func (i InputPeerPhotoFileLocation) construct() InputFileLocationClass { return &i }

// Ensuring interfaces in compile-time for InputPeerPhotoFileLocation.
var (
	_ bin.Encoder = &InputPeerPhotoFileLocation{}
	_ bin.Decoder = &InputPeerPhotoFileLocation{}

	_ InputFileLocationClass = &InputPeerPhotoFileLocation{}
)

// InputStickerSetThumb represents TL type `inputStickerSetThumb#dbaeae9`.
// Location of stickerset thumbnail (see files¹)
//
// Links:
//  1) https://core.telegram.org/api/files
//
// See https://core.telegram.org/constructor/inputStickerSetThumb for reference.
type InputStickerSetThumb struct {
	// Sticker set
	Stickerset InputStickerSetClass
	// Volume ID
	VolumeID int64
	// Local ID
	LocalID int
}

// InputStickerSetThumbTypeID is TL type id of InputStickerSetThumb.
const InputStickerSetThumbTypeID = 0xdbaeae9

// String implements fmt.Stringer.
func (i *InputStickerSetThumb) String() string {
	if i == nil {
		return "InputStickerSetThumb(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputStickerSetThumb")
	sb.WriteString("{\n")
	sb.WriteString("\tStickerset: ")
	sb.WriteString(fmt.Sprint(i.Stickerset))
	sb.WriteString(",\n")
	sb.WriteString("\tVolumeID: ")
	sb.WriteString(fmt.Sprint(i.VolumeID))
	sb.WriteString(",\n")
	sb.WriteString("\tLocalID: ")
	sb.WriteString(fmt.Sprint(i.LocalID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *InputStickerSetThumb) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStickerSetThumb#dbaeae9 as nil")
	}
	b.PutID(InputStickerSetThumbTypeID)
	if i.Stickerset == nil {
		return fmt.Errorf("unable to encode inputStickerSetThumb#dbaeae9: field stickerset is nil")
	}
	if err := i.Stickerset.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputStickerSetThumb#dbaeae9: field stickerset: %w", err)
	}
	b.PutLong(i.VolumeID)
	b.PutInt(i.LocalID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputStickerSetThumb) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStickerSetThumb#dbaeae9 to nil")
	}
	if err := b.ConsumeID(InputStickerSetThumbTypeID); err != nil {
		return fmt.Errorf("unable to decode inputStickerSetThumb#dbaeae9: %w", err)
	}
	{
		value, err := DecodeInputStickerSet(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputStickerSetThumb#dbaeae9: field stickerset: %w", err)
		}
		i.Stickerset = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputStickerSetThumb#dbaeae9: field volume_id: %w", err)
		}
		i.VolumeID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputStickerSetThumb#dbaeae9: field local_id: %w", err)
		}
		i.LocalID = value
	}
	return nil
}

// construct implements constructor of InputFileLocationClass.
func (i InputStickerSetThumb) construct() InputFileLocationClass { return &i }

// Ensuring interfaces in compile-time for InputStickerSetThumb.
var (
	_ bin.Encoder = &InputStickerSetThumb{}
	_ bin.Decoder = &InputStickerSetThumb{}

	_ InputFileLocationClass = &InputStickerSetThumb{}
)

// InputFileLocationClass represents InputFileLocation generic type.
//
// See https://core.telegram.org/type/InputFileLocation for reference.
//
// Example:
//  g, err := DecodeInputFileLocation(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *InputFileLocation: // inputFileLocation#dfdaabe1
//  case *InputEncryptedFileLocation: // inputEncryptedFileLocation#f5235d55
//  case *InputDocumentFileLocation: // inputDocumentFileLocation#bad07584
//  case *InputSecureFileLocation: // inputSecureFileLocation#cbc7ee28
//  case *InputTakeoutFileLocation: // inputTakeoutFileLocation#29be5899
//  case *InputPhotoFileLocation: // inputPhotoFileLocation#40181ffe
//  case *InputPhotoLegacyFileLocation: // inputPhotoLegacyFileLocation#d83466f3
//  case *InputPeerPhotoFileLocation: // inputPeerPhotoFileLocation#27d69997
//  case *InputStickerSetThumb: // inputStickerSetThumb#dbaeae9
//  default: panic(v)
//  }
type InputFileLocationClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputFileLocationClass
	fmt.Stringer
}

// DecodeInputFileLocation implements binary de-serialization for InputFileLocationClass.
func DecodeInputFileLocation(buf *bin.Buffer) (InputFileLocationClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputFileLocationTypeID:
		// Decoding inputFileLocation#dfdaabe1.
		v := InputFileLocation{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputFileLocationClass: %w", err)
		}
		return &v, nil
	case InputEncryptedFileLocationTypeID:
		// Decoding inputEncryptedFileLocation#f5235d55.
		v := InputEncryptedFileLocation{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputFileLocationClass: %w", err)
		}
		return &v, nil
	case InputDocumentFileLocationTypeID:
		// Decoding inputDocumentFileLocation#bad07584.
		v := InputDocumentFileLocation{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputFileLocationClass: %w", err)
		}
		return &v, nil
	case InputSecureFileLocationTypeID:
		// Decoding inputSecureFileLocation#cbc7ee28.
		v := InputSecureFileLocation{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputFileLocationClass: %w", err)
		}
		return &v, nil
	case InputTakeoutFileLocationTypeID:
		// Decoding inputTakeoutFileLocation#29be5899.
		v := InputTakeoutFileLocation{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputFileLocationClass: %w", err)
		}
		return &v, nil
	case InputPhotoFileLocationTypeID:
		// Decoding inputPhotoFileLocation#40181ffe.
		v := InputPhotoFileLocation{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputFileLocationClass: %w", err)
		}
		return &v, nil
	case InputPhotoLegacyFileLocationTypeID:
		// Decoding inputPhotoLegacyFileLocation#d83466f3.
		v := InputPhotoLegacyFileLocation{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputFileLocationClass: %w", err)
		}
		return &v, nil
	case InputPeerPhotoFileLocationTypeID:
		// Decoding inputPeerPhotoFileLocation#27d69997.
		v := InputPeerPhotoFileLocation{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputFileLocationClass: %w", err)
		}
		return &v, nil
	case InputStickerSetThumbTypeID:
		// Decoding inputStickerSetThumb#dbaeae9.
		v := InputStickerSetThumb{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputFileLocationClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputFileLocationClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputFileLocation boxes the InputFileLocationClass providing a helper.
type InputFileLocationBox struct {
	InputFileLocation InputFileLocationClass
}

// Decode implements bin.Decoder for InputFileLocationBox.
func (b *InputFileLocationBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputFileLocationBox to nil")
	}
	v, err := DecodeInputFileLocation(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputFileLocation = v
	return nil
}

// Encode implements bin.Encode for InputFileLocationBox.
func (b *InputFileLocationBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputFileLocation == nil {
		return fmt.Errorf("unable to encode InputFileLocationClass as nil")
	}
	return b.InputFileLocation.Encode(buf)
}
