package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"reflect"
	"strings"

	v1 "github.com/star-table/interface/golang/table/v1"

	jsoniter "github.com/json-iterator/go"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

// UnmarshalOptions is a configurable JSON format parser.
var UnmarshalOptions = protojson.UnmarshalOptions{
	DiscardUnknown: true,
}

// MarshalOptions is a configurable JSON format marshaller.
var MarshalOptions = protojson.MarshalOptions{
	EmitUnpopulated: true,
}

type ts struct {
	A int64 `json:"a"`
}

type tss struct {
	ts
	B string
}

var jsonIter = jsoniter.ConfigCompatibleWithStandardLibrary

type GetAppInfoByAppIdResp struct {
	Data *GetAppInfoByAppIdRespData `json:"data"`
}

type GetAppInfoByAppIdRespData struct {
	Id          int64  `json:"id"`
	PkgId       int64  `json:"pkgId"`
	OrgId       int64  `json:"orgId"`
	ExtendsId   string `json:"extendsId"`
	Type        int    `json:"type"`
	ProjectId   int64  `json:"projectId"`
	MirrorAppId int64  `json:"mirrorAppId"`
}

type TableColumnData struct {
	Name              string           `json:"name"`
	Label             string           `json:"label"`
	AliasTitle        string           `json:"aliasTitle"`
	Description       string           `json:"description"`
	IsSys             bool             `json:"isSys"`
	IsOrg             bool             `json:"isOrg"`
	Writable          bool             `json:"writable"`
	Editable          bool             `json:"editable"`
	Unique            bool             `json:"unique"`
	UniquePreHandler  string           `json:"uniquePreHandler"`
	SensitiveStrategy string           `json:"sensitiveStrategy"`
	SensitiveFlag     int32            `json:"sensitiveFlag"`
	Field             TableColumnField `json:"field"`
}

type TableColumnField struct {
	Type       string                 `json:"type,omitempty"`
	CustomType string                 `json:"customType"`
	DataType   string                 `json:"dataType,omitempty"`
	Props      map[string]interface{} `json:"props"`
}

type LcCtSelectFieldPropsSelectOptions struct {
	Id string `json:"id"`
}

func main() {
	gzipStr := "H4sIAAAAAAACA9WW3Y6aQBSA32WuSSOKduVOWnRlRSuwbsvGi5EZ3anDj8PgTwhJb/sKvd5X24u+RQcqC82mvbCJCXfD4ZtzvnNmojymwAt9HwYIqMDGfEgwRQPOGZAAP0UYqG0JQI+TMIiB+pie1wJGkMN38V92cLiieJzn5CuK5NMH8243cEfmRhPUnuBD8W6PD16P7j7bpoiu8zxFeE3RbG6R3sQV4ZCRDQk+imJATUEAfVEAvDz/+Pnt+2tBWQIRCyPM+AmoQUKpBBCOPUai365pJnqgNDxgpCPCzVC8zbOJthNIgcpZgiUQe1Cwa0hjnIkN6J8lO/WSKQijakIk70EEOt3FsO9sxZZzClksvZCGDKitTKrAu7mNurdJBXYqUK6DTqApqOdWYLsC23VwRXx79nVXgd0K7NTBL3NrpMNjBSoVqGRLqWzMKZpuZW8Gm+XZ3twJC3shQ/95IViRpIiLpW0YQ983zrRTnnt5SrU7I55ez2gPaYL/PI7rKS/cBb5taU1S9mB/u7InTVIezMeacT9qkvJk62rRCF2uXP5iXE+5vV3IY75u0pQN+cbSuk9NUt6w6Wli6E1SfqKdVmJML1cu/9Sup3zUuGEqD02a8o3uEPOhd7ly+UVwPeVPM/Z+cAyaNGWrrytTeH+5cvk5JZSX2fIXrZOKMVwLAAA="
	bts, _ := (&base64.Encoding{}).DecodeString(gzipStr)
	gzipReader, err := gzip.NewReader(bytes.NewReader(bts))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer gzipReader.Close()
	var buf bytes.Buffer
	_, err = io.Copy(&buf, gzipReader)

	fmt.Println(buf.String())

	//html := "<p>fsf<span style=\"font-size:14px\">fsfsf<span style=\"font-size:40px\">sfsfsfs<strong>gsfsdf<u>fsdfsdf</u></strong></span></span></p><div class=\"media-wrap image-wrap\"><img src=\"https://attachments.startable.cn/org_2578/feedback_resource/2022/10/10/d69f4d98cdba4011a6ae1089c6695af01665381939769.png\"/></div><p>sfsfsdfsdfsdfsdfsdff</p>"
	//html = strings.Replace(html, "</p>", "\n", -1)
	//regx, err := regexp.Compile("<[^>]*>")
	//if err != nil {
	//	return
	//}
	//fmt.Println(regx.ReplaceAllString(html, ""))
	//t, _ := time.Parse("2006-01-02 15:04:05", "1971-01-01 00:00:00")
	//fmt.Println(t.Unix())
	//f := excelize.NewFile()
	//// 创建一个工作表
	//f.SetColWidth("Sheet1", "A", "B", 30)
	//f.SetRowHeight("Sheet1", 1, 60)
	//f.SetRowHeight("Sheet1", 2, 55)
	////f.GetRowHeight("Sheet1", "A")
	//
	//for i := 1; i <= 3; i++ {
	//	formatStyle := fmt.Sprintf(`{
	//    		"x_offset": %v,
	//    		"y_offset": %v,
	//    		"hyperlink": "%s",
	//    		"hyperlink_type": "External",
	//    		"print_obj": false,
	//    		"lock_aspect_ratio": true,
	//    		"locked": false,
	//    		"positioning": "oneCell",
	//			"autofit": false
	//		 }`, 20+(i-1)*60, 10, "https://www.baidu.com")
	//	err := f.AddPicture("Sheet1", "B2", fmt.Sprintf("./test/static/%d.png", i), formatStyle)
	//	fmt.Println(err)
	//}
	//
	//f.SetColWidth("Sheet1", "A", "A", 9.6)
	//f.SetRowHeight("Sheet1", 1, 20)
	//// 根据指定路径保存文件
	//if err := f.SaveAs("Book1.xlsx"); err != nil {
	//	fmt.Println(err)
	//}
}

func getAttachmentRecycleJson(resourceIds []int64, columnIds []string, recycleFlag int) string {
	updateJsons := make([]string, 0, len(columnIds))
	defaultValue := `jsonb_set(%s, '{%s,%d,recycleFlag}', '%d', false)`
	for _, resourceId := range resourceIds {
		dataValue := defaultValue
		if len(columnIds) == 1 {
			dataValue = fmt.Sprintf(defaultValue, "data", columnIds[0], resourceId, recycleFlag)
		} else {
			for i, columnId := range columnIds {
				if i == len(columnIds)-1 {
					dataValue = fmt.Sprintf(dataValue, fmt.Sprintf(defaultValue, "data", columnId, resourceId, recycleFlag))
				} else if i == 0 {
					dataValue = fmt.Sprintf(defaultValue, "%s", columnId, resourceId, recycleFlag)
				} else {
					dataValue = fmt.Sprintf(defaultValue, dataValue, columnId, resourceId, recycleFlag)
				}
			}
		}

		updateJsons = append(updateJsons, dataValue)
	}

	updateJson := updateJsons[0]
	for i := 1; i < len(updateJsons); i++ {
		updateJson = strings.Replace(updateJson, "data", updateJsons[i], 1)
	}

	return updateJson
}

func getStatusName(id int64, issueStatusColumn *TableColumnData) (string, error) {
	groupSelect := v1.ColumnType_groupSelect.String()
	if issueStatusColumn != nil && issueStatusColumn.Field.Type == groupSelect && issueStatusColumn.Field.Props != nil &&
		issueStatusColumn.Field.Props[groupSelect] != nil {
		groupSelectInfo := issueStatusColumn.Field.Props[groupSelect]
		if m, ok := groupSelectInfo.(map[string]interface{}); ok {
			groupStruct, err := structpb.NewStruct(m)
			if err != nil {
				return "", err
			}
			if groupStruct.Fields["options"] != nil && groupStruct.Fields["options"].GetListValue() != nil {
				for _, value := range groupStruct.Fields["options"].GetListValue().Values {
					if value.GetStructValue() != nil && value.GetStructValue().Fields["id"] != nil {
						if id == int64(value.GetStructValue().Fields["id"].GetNumberValue()) {
							return value.GetStructValue().Fields["value"].GetStringValue(), nil
						}
					}
				}
			}
		}
	}

	return "", nil
}

func Unmarshal(data []byte, v interface{}) error {
	switch m := v.(type) {
	default:
		rv := reflect.ValueOf(v)
		for rv := rv; rv.Kind() == reflect.Ptr; {
			if rv.IsNil() {
				rv.Set(reflect.New(rv.Type().Elem()))
			}
			rv = rv.Elem()
		}
		if m, ok := reflect.Indirect(rv).Interface().(proto.Message); ok {
			return UnmarshalOptions.Unmarshal(data, m)
		}
		return jsonIter.Unmarshal(data, m)
	}
}
