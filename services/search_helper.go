package services

import (
	"bufio"
	"bytes"
	"io"
	"regexp"
	"strings"
	"time"

	"github.com/sant470/devetron/lib/s3"
	apptypes "github.com/sant470/devetron/types"
	"github.com/sant470/devetron/types/consts"
)

// const (
// 	_prefix = "app-logs-bucket"
// )

func dateTimeFormat(timestamp int64) (hour string, date string) {
	ut := time.Unix(timestamp, 0)
	fDate := ut.Format("2006-01-02 15")
	dateComp := strings.Split(fDate, " ")
	return dateComp[1], dateComp[0]
}

func (searchSvc *SearchService) searchRemoteFile(cli *s3.Client, pathSuffix string, query string, out chan<- apptypes.Match, cancelSig <-chan struct{}) {
	barr, err := cli.GetS3ObjectBuffer(consts.Bucket, pathSuffix)
	if err != nil {
		searchSvc.lgr.Printf("error getting file: %s, error: %s", pathSuffix, err.Error())
	}
	reader := bufio.NewReader(bytes.NewReader(barr))
	lineNo := 0
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			searchSvc.lgr.Printf("error reading lines from the file: %s, error: %s", pathSuffix, err.Error())
			continue
		}
		lineNo++
		match, err := regexp.MatchString(query, str)
		if err != nil {
			continue
		}
		if match {
			select {
			case <-cancelSig:
				return
			default:
				out <- apptypes.Match{
					Line:     lineNo,
					FilePath: pathSuffix,
					Text:     str,
				}
			}
		}
	}
}

// func (searchSvc *SearchService) searchDiskFile(pathSuffix string, query string, out chan<- apptypes.Match, cancelSig <-chan struct{}) {
// 	filePath := filepath.Join(_prefix, pathSuffix)
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		searchSvc.lgr.Printf("error opening file: %s, error: %s", filePath, err.Error())
// 		return
// 	}
// 	defer file.Close()
// 	reader := bufio.NewReader(file)
// 	lineNo := 0
// 	for {
// 		str, err := reader.ReadString('\n')
// 		if err != nil {
// 			if err == io.EOF {
// 				break
// 			}
// 			searchSvc.lgr.Printf("error reading lines from the file: %s, error: %s", filePath, err.Error())
// 			continue
// 		}
// 		lineNo++
// 		match, err := regexp.MatchString(query, str)
// 		if err != nil {
// 			continue
// 		}
// 		if match {
// 			select {
// 			case <-cancelSig:
// 				return
// 			default:
// 				out <- apptypes.Match{
// 					Line:     lineNo,
// 					FilePath: filePath,
// 					Text:     str,
// 				}
// 			}
// 		}
// 	}
// }
