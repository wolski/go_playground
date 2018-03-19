package main

import (
	"fmt"

	"github.com/mcgrew/gomzlib"
)

func main() {
	rawdata := mzlib.RawData{}
	rawdata.ReadMzXml("C:\\Users\\wewol\\Downloads\\cleavableICAT_ms2x2_2.mzXML")
	fmt.Println(rawdata.MinMz(), rawdata.MaxMz(), rawdata.ScanCount, len(rawdata.Scans))
	i := rawdata.Instrument
	fmt.Println(i.Accuracy, "\n",
		i.Detector, "\n",
		i.Ionization)
	fmt.Print("Hello World")
}
