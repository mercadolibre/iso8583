package field

import (
	"reflect"

	"github.com/mercadolibre/iso8583/encoding"
	"github.com/mercadolibre/iso8583/padding"
	"github.com/mercadolibre/iso8583/prefix"
	"github.com/mercadolibre/iso8583/sort"
)

// TagSpec is used to define the format of field tags (sometimes defined as field IDs).
// This is most commonly used by composite field types such as TLVs, subfields
// and subelements. TagSpecs need not be defined for primitive field types
// such as numerics, strings or for composite field types that don't consist
// of tags in the message payload.
type TagSpec struct {
	// Length is defined for subfields and subelements whose tag
	// lengths are fixed and can be defined statically.
	// This field should not be populated in conjunction with the TLV Tag
	// encoder as their lengths are determined dynamically.
	Length int
	// Enc defines the encoder used to marshal and unmarshal
	// the tag.
	Enc encoding.Encoder
	// Pad sets the padding direction and type of the tag.
	// This is most commonly used for composite field types
	// whose tags hold leading 0s e.g. '003' would be unpadded to '3'.
	Pad padding.Padder
	// Sort defines the order in which Tags defined within the subfields
	// spec must be packed. This ordering may also be used for unpacking
	// if Spec.Tag.Enc == nil.
	Sort sort.StringSlice
	// SkipUnknownTLVTags is a flag which indicates whether TLV tags that are not found in
	// the spec should be skipped or throw an error.
	// This flag is only meant to be used in Composite fields with TLV encoding.
	SkipUnknownTLVTags bool
}

// Spec defines the structure of a field.
type Spec struct {
	// Length defines the maximum length of field (bytes, characters or
	// digits), for both fixed and variable lengths.
	Length int
	// Tag sets the tag specification. Only applicable to composite field
	// types.
	Tag *TagSpec
	// Description of what data the field holds.
	Description string
	// Enc defines the encoder used to marshal and unmarshal the field.
	// Only applicable to primitive field types e.g. numerics, strings,
	// binary etc
	Enc encoding.Encoder
	// Pref defines the prefixer of the field used to encode and decode the
	// length of the field.
	Pref prefix.Prefixer
	// Pad sets the padding direction and type of the field.
	Pad padding.Padder
	// Subfields defines the subfields held within the field. Only
	// applicable to composite field types.
	Subfields map[string]Field
	// BitmapOptions could be used on a bitmap field, and is optional.
	// Contains options that modifies the behaviour of the bitmap.
	BitmapOptions *BitmapOptions
}

func NewSpec(length int, desc string, enc encoding.Encoder, pref prefix.Prefixer) *Spec {
	return &Spec{
		Length:      length,
		Description: desc,
		Enc:         enc,
		Pref:        pref,
	}
}

// CreateSubfield creates a new instance of a field based on the input
// provided.
func CreateSubfield(specField Field) Field {
	fieldType := reflect.TypeOf(specField).Elem()

	// create new field and convert it to Field interface
	fl := reflect.New(fieldType).Interface().(Field)
	fl.SetSpec(specField.Spec())

	if composite, ok := fl.(CompositeWithSubfields); ok {
		composite.ConstructSubfields()
	}

	return fl
}

func CreateSubfields(s *Spec) map[string]Field {
	subfields := map[string]Field{}

	for k, specField := range s.Subfields {
		subfields[k] = CreateSubfield(specField)
	}

	return subfields
}
