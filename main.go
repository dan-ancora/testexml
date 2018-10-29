package main

import (
	"encoding/xml"
	"fmt"
	//"io/ioutil"
	//"net/http"
	//"strings"
)

//Cities List strcut for decoding XML
type XmlEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    XmlBody  `xml:"Body"`
}

//type body
type XmlBody struct {
	XMLName  xml.Name           `xml:"Body"`
	Fault    XmlFault           `xml:"Fault`
	CityList ListCitiesResponse `xml:"ListCitiesResponse`
}

//type fault
type XmlFault struct {
	XMLName        xml.Name `xml:"Fault"`
	XMLFaultCode   xml.Name `xml:"faultcode"`
	XMLFaultString string   `xml:"faultstring"`
	XMLFaultDetail xml.Name `xml:"detail"`
}

type ListCitiesResponse struct {
	XMLName  xml.Name `xml:"ListCitiesResponse"`
	RespBody CityResp `xml:"cityResp"`
}
type CityResp struct {
	XMLName xml.Name  `xml:"cityResp"`
	Cities  []XMLCity `xml:"body"`
}

type XMLCity struct {
	XMLName         xml.Name `xml:"body"`
	XMLCityCode     string   `xml:"CityCode"`
	XMLCityName     string   `xml:"CityName"`
	XMLCityCountry  string   `xml:"Country"`
	XMLDisplayGroup string   `xml:"DisplayGroup"`
}

func main() {

	xresponse := []byte(`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
	<soapenv:Body>
	<!--
		<soapenv:Fault>
			<faultcode>soapenv:Server</faultcode>
			<faultstring> Error - Login Invalid</faultstring>
			<detail/>
		</soapenv:Fault>
	-->
		<xsd:ListCitiesResponse xmlns:xsd="http://www.wso2.org/php/xsd" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
         <xsd:cityResp>
            <xsd:body>
               <xsd:CityCode>ABB</xsd:CityCode>
               <xsd:CityName>*Abbotsford, BC, CANADA</xsd:CityName>
               <xsd:Country>CANADA</xsd:Country>
               <xsd:DisplayGroup xsi:nil="true"/>
            </xsd:body>
           <xsd:body>
               <xsd:CityCode>ZXR</xsd:CityCode>
               <xsd:CityName>Carmel Valley</xsd:CityName>
               <xsd:Country>USA</xsd:Country>
               <xsd:DisplayGroup xsi:nil="true"/>
            </xsd:body>
         </xsd:cityResp>
      </xsd:ListCitiesResponse>
	</soapenv:Body>
</soapenv:Envelope>`)

	var envelope XmlEnvelope
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	err := xml.Unmarshal(xresponse, &envelope)
	if err != nil {
		fmt.Println(err)
	}
	if envelope.Body.Fault.XMLFaultString != "" {
		fmt.Println(envelope.Body.Fault.XMLFaultString)

	} else {
		for i, city := range envelope.Body.CityList.RespBody.Cities {
			fmt.Println(i, city.XMLCityCode, city.XMLCityCountry, city.XMLCityName, city.XMLDisplayGroup)
		}

	}

}
