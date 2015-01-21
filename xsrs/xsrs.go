package xsrs

import (
	"encoding/xml"
)

type Root struct {
	XMLName xml.Name `xml:"urn:schemas-xsrs-org:metadata-1-0/x_srs/ xsrs"`
	Items   []Item   `xml:"item"`
}

type Item struct {
	XMLName                xml.Name `xml:"item"`
	Id                     string   `xml:"id,attr"`
	Title                  string   `xml:"title"`
	ScheduledStartDateTime string   `xml:"scheduledStartDateTime"`
	ScheduledDuration      string   `xml:"scheduledDuration"`
	ScheduledConditionID   string   `xml:"scheduledConditionID"`
	ScheduledChannelID     ScheduledChannelID
	DesiredMatchingID      DesiredMatchingID
	DesiredQualityMode     string `xml:"desiredQualityMode"`
	GenreID                GenreID
	ConflictID             string `xml:"conflictID"`
	MediaRemainAlertID     string `xml:"mediaRemainAlertID"`
	ReservationCreatorID   string `xml:"reservationCreatorID"`
	RecordingFlag          string `xml:"recordingFlag"`
	RecordDestinationID    string `xml:"recordDestinationID"`
	RecordSize             string `xml:"recordSize"`
	PortableRecordFile     PortableRecordFile
}

type ScheduledChannelID struct {
	XMLName          xml.Name `xml:"scheduledChannelID"`
	BroadcastingType int      `xml:"broadcastingType,attr"`
	ChannelType      int      `xml:"channelType,attr"`
	Value            string   `xml:",chardata"`
}

type DesiredMatchingID struct {
	XMLName xml.Name `xml:"desiredMatchingID"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}

type GenreID struct {
	XMLName xml.Name `xml:"genreID"`
	Type    string   `xml:"type,attr"`
	Value   int      `xml:",chardata"`
}

type PortableRecordFile struct {
	XMLName      xml.Name `xml:"portableRecordFile"`
	Target       string   `xml:"target,attr"`
	TransferPath string   `xml:"transferPath,attr"`
	Value        int      `xml:",chardata"`
}
