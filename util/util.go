package util

import (
    "github.com/kolo/xmlrpc"
    "time"
)

func ToStringSlice(is []interface{}) []string {
    ss := make([]string, len(is), len(is))

    for i, _ := range is {
        ss[i] = is[i].(string)
    }
    return ss
}

func ToString(i interface{}) string {
    if v, ok := i.(string); ok {
        return v
    }
    return ""
}

func ToTime(i interface{}) time.Time {
    if v, ok := i.(time.Time); ok {
        return v
    }
    var t time.Time
    return t
}

func ToInterfaceSlice(i interface{}) []interface{} {
    if v, ok := i.([]interface{}); ok {
        return v
    }
    var s []interface{}
    return s
}

func ToInt64(i interface{}) int64 {
    if v, ok := i.(int64); ok {
        return v
    }
    var n int64
    return n
}

func ToFloat64(i interface{}) float64 {
    if v, ok := i.(float64); ok {
        return v
    }
    var n float64
    return n
}

func ToXmlrpcStruct(i interface{}) xmlrpc.Struct {
    if v, ok := i.(xmlrpc.Struct); ok {
        return v
    }
    var s xmlrpc.Struct
    return s
}

func ToBool(i interface{}) bool {
    if v, ok := i.(bool); ok {
        return v
    }
    var b bool
    return b
}

func ToIntSlice(is []interface{}) []int {
    numbers := make([]int, len(is), len(is))

    for i, _ := range is {
        numbers[i] = int(is[i].(int64))
    }
    return numbers
}