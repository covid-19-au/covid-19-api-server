package util

import (
    "github.com/golang/protobuf/ptypes"
    "github.com/golang/protobuf/ptypes/timestamp"
    "time"
)

func Unix2Timestamp(unixSeconds int64) *timestamp.Timestamp {
    t := time.Unix(unixSeconds, 0)

    if timeProto, err := ptypes.TimestampProto(t); err != nil {
        return nil
    } else {
        return timeProto
    }
}
