package shop

import (
	"v19go/internal/database"
	"v19go/internal/utils/bytestream"
)

type OfferItem struct {
	ID         int32
	Multiplier int32
	DataRef    [2]int32
	SkinID     int32
}

type Offer struct {
	Name        string
	Cost        int32
	Timer       int32
	ShopDisplay int32
	Type        int32
	Claim       bool
	Includes    []OfferItem
}

type Shop struct {
	offers []*Offer
}

func NewShop() *Shop {
	return &Shop{
		offers: []*Offer{
			{
				Name:        " ",
				Cost:        0,
				Timer:       86400,
				ShopDisplay: 1,
				Type:        0,
				Claim:       false,
				Includes: []OfferItem{
					{ID: 1, Multiplier: 100, DataRef: [2]int32{0, 0}, SkinID: 0},
				},
			},
			{
				Name:        " ",
				Cost:        8,
				Timer:       86400,
				ShopDisplay: 1,
				Type:        1,
				Claim:       false,
				Includes: []OfferItem{
					{ID: 8, Multiplier: 37, DataRef: [2]int32{16, 8}, SkinID: 0},
				},
			},
		},
	}
}

func (s *Shop) GetAll() []*Offer {
	return s.offers
}

func (s *Shop) Encode(stream *bytestream.ByteStream, account *database.Account) {
	stream.WriteVInt(int32(len(s.offers)))

	for _, offer := range s.offers {
		stream.WriteVInt(int32(len(offer.Includes)))
		for _, item := range offer.Includes {
			stream.WriteVInt(item.ID)
			stream.WriteVInt(item.Multiplier)
			stream.WriteDataReference(item.DataRef[0], item.DataRef[1])
			stream.WriteVInt(item.SkinID)
		}

		stream.WriteVInt(offer.Type)
		stream.WriteVInt(offer.Cost)
		stream.WriteVInt(offer.Timer)
		stream.WriteVInt(0) // Offer View
		stream.WriteVInt(100)
		stream.WriteBoolean(offer.Claim)
		stream.WriteBoolean(false)
		stream.WriteVInt(offer.ShopDisplay)
		stream.WriteVInt(0) // OldCost
		stream.WriteInt(0)
		stream.WriteString(offer.Name)
		stream.WriteBoolean(false)
	}
}