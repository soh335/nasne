package nasne

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/soh335/nasne/xsrs"
)

var getRecordScheduleLisTmpl = `<?xml version="1.0" encoding="utf-8"?>
<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/">
  <s:Body>
    <u:X_GetRecordScheduleList xmlns:u="urn:schemas-xsrs-org:service:X_ScheduledRecording:2">
      <SearchCriteria />
      <Filter />
      <StartingIndex />
      <RequestedCount />
      <SortCriteria />
    </u:X_GetRecordScheduleList>
  </s:Body>
</s:Envelope>
`

type Root struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    Body
}

type Body struct {
	XMLName                       xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	GetRecordScheduleListResponse GetRecordScheduleListResponse
}

type GetRecordScheduleListResponse struct {
	XMLName        xml.Name `xml:"urn:schemas-xsrs-org:service:X_ScheduledRecording:2 X_GetRecordScheduleListResponse"`
	Result         string   `xml:"Result"`
	NumberReturned int      `xml:"NumberReturned"`
	TotalMatches   int      `xml:"TotalMatches"`
	UpdateID       int      `xml:"UpdateID"`
}

func GetRecordScheduleList(addr string) (*xsrs.Root, error) {

	req, err := http.NewRequest("POST", "/XSRS", strings.NewReader(getRecordScheduleLisTmpl))

	if err != nil {
		return nil, err
	}

	conn, err := net.Dial("tcp", addr)

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	req.Header.Set("Content-Type", `text/xml; charset="utf-8"`)

	var b bytes.Buffer
	if err := req.Write(&b); err != nil {
		return nil, err
	}

	if _, err := conn.Write(b.Bytes()); err != nil {
		return nil, err
	}

	res, err := http.ReadResponse(bufio.NewReader(conn), req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("http error: %s", res.Status)
	}

	defer res.Body.Close()

	var r Root
	if err := xml.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	var xsrsRoot xsrs.Root
	if err := xml.NewDecoder(strings.NewReader(r.Body.GetRecordScheduleListResponse.Result)).Decode(&xsrsRoot); err != nil {
		return nil, err
	}

	return &xsrsRoot, nil

}
