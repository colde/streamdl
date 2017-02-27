package main

type SmoothStreamingMedia struct {
	MinorVersion		int	`xml:"MinorVersion,attr"`
	Duration		int	`xml:"Duration,attr"`
	MajorVersion		int	`xml:"MajorVersion,attr"`
	ProtectionHeader	ProtectionHeader	`xml:"Protection>ProtectionHeader"`
	StreamIndex		[]StreamIndex	`xml:"StreamIndex"`
}

type ProtectionHeader struct {
	SystemID		string	`xml:"SystemID,attr"`
	Text			string	`xml:",chardata"`
}
type StreamIndex struct {
	Language		string	`xml:"Language,attr"`
	Chunks			int	`xml:"Chunks,attr"`
	Url			string	`xml:"Url,attr"`
	QualityLevels		int	`xml:"QualityLevels,attr"`
	Type			string	`xml:"Type,attr"`
	Name			string	`xml:"Name,attr"`
	Subtype			string	`xml:"Subtype,attr"`
	QualityLevel		[]QualityLevel	`xml:"QualityLevel"`
	Fragment		[]Fragment	`xml:"c"`
}
type QualityLevel struct {
	Bitrate			int	`xml:"Bitrate,attr"`
	MaxWidth		int	`xml:"MaxWidth,attr"`
	FourCC			string	`xml:"FourCC,attr"`
	CodecPrivateData	string	`xml:"CodecPrivateData,attr"`
	MaxHeight		int	`xml:"MaxHeight,attr"`
	Index			int	`xml:"Index,attr"`
	PacketSize		int	`xml:"PacketSize,attr"`
	Channels		int	`xml:"Channels,attr"`
	AudioTag		string	`xml:"AudioTag,attr"`
	BitsPerSample		int	`xml:"BitsPerSample,attr"`
	SamplingRate		int	`xml:"SamplingRate,attr"`
}
type Fragment struct {
	Duration		int	`xml:"d,attr"`
	Timestamp		int	`xml:"t,attr"`
}