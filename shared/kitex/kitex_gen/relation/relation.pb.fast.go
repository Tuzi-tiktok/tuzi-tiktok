// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package relation

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
	auth "tuzi-tiktok/kitex/kitex_gen/auth"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *RelationRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RelationRequest[number], err)
}

func (x *RelationRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Token, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RelationRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.ToUserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *RelationRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.ActionType, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *RelationResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RelationResponse[number], err)
}

func (x *RelationResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *RelationResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.StatusMsg = &tmp
	return offset, err
}

func (x *RelationFollowListRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RelationFollowListRequest[number], err)
}

func (x *RelationFollowListRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *RelationFollowListRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Token, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RelationFollowListResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RelationFollowListResponse[number], err)
}

func (x *RelationFollowListResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *RelationFollowListResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.StatusMsg = &tmp
	return offset, err
}

func (x *RelationFollowListResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v auth.User
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.UserList = append(x.UserList, &v)
	return offset, nil
}

func (x *RelationFollowerListRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RelationFollowerListRequest[number], err)
}

func (x *RelationFollowerListRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *RelationFollowerListRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Token, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RelationFollowerListResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RelationFollowerListResponse[number], err)
}

func (x *RelationFollowerListResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *RelationFollowerListResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.StatusMsg = &tmp
	return offset, err
}

func (x *RelationFollowerListResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v auth.User
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.UserList = append(x.UserList, &v)
	return offset, nil
}

func (x *RelationFriendListRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RelationFriendListRequest[number], err)
}

func (x *RelationFriendListRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *RelationFriendListRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Token, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RelationFriendListResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RelationFriendListResponse[number], err)
}

func (x *RelationFriendListResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *RelationFriendListResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.StatusMsg = &tmp
	return offset, err
}

func (x *RelationFriendListResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v FriendUser
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.UserList = append(x.UserList, &v)
	return offset, nil
}

func (x *FriendUser) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 7:
		offset, err = x.fastReadField7(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 8:
		offset, err = x.fastReadField8(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 9:
		offset, err = x.fastReadField9(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 10:
		offset, err = x.fastReadField10(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 11:
		offset, err = x.fastReadField11(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 12:
		offset, err = x.fastReadField12(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 13:
		offset, err = x.fastReadField13(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_FriendUser[number], err)
}

func (x *FriendUser) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *FriendUser) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Name, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *FriendUser) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadInt64(buf, _type)
	x.FollowCount = &tmp
	return offset, err
}

func (x *FriendUser) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadInt64(buf, _type)
	x.FollowerCount = &tmp
	return offset, err
}

func (x *FriendUser) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.IsFollow, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *FriendUser) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Avatar = &tmp
	return offset, err
}

func (x *FriendUser) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.BackgroundImage = &tmp
	return offset, err
}

func (x *FriendUser) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Signature = &tmp
	return offset, err
}

func (x *FriendUser) fastReadField9(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadInt64(buf, _type)
	x.TotalFavorited = &tmp
	return offset, err
}

func (x *FriendUser) fastReadField10(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadInt64(buf, _type)
	x.WorkCount = &tmp
	return offset, err
}

func (x *FriendUser) fastReadField11(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadInt64(buf, _type)
	x.FavoriteCount = &tmp
	return offset, err
}

func (x *FriendUser) fastReadField12(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Message = &tmp
	return offset, err
}

func (x *FriendUser) fastReadField13(buf []byte, _type int8) (offset int, err error) {
	x.MsgType, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *RelationRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *RelationRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Token == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetToken())
	return offset
}

func (x *RelationRequest) fastWriteField2(buf []byte) (offset int) {
	if x.ToUserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetToUserId())
	return offset
}

func (x *RelationRequest) fastWriteField3(buf []byte) (offset int) {
	if x.ActionType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, x.GetActionType())
	return offset
}

func (x *RelationResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *RelationResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetStatusCode())
	return offset
}

func (x *RelationResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetStatusMsg())
	return offset
}

func (x *RelationFollowListRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *RelationFollowListRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *RelationFollowListRequest) fastWriteField2(buf []byte) (offset int) {
	if x.Token == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetToken())
	return offset
}

func (x *RelationFollowListResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *RelationFollowListResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetStatusCode())
	return offset
}

func (x *RelationFollowListResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetStatusMsg())
	return offset
}

func (x *RelationFollowListResponse) fastWriteField3(buf []byte) (offset int) {
	if x.UserList == nil {
		return offset
	}
	for i := range x.GetUserList() {
		offset += fastpb.WriteMessage(buf[offset:], 3, x.GetUserList()[i])
	}
	return offset
}

func (x *RelationFollowerListRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *RelationFollowerListRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *RelationFollowerListRequest) fastWriteField2(buf []byte) (offset int) {
	if x.Token == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetToken())
	return offset
}

func (x *RelationFollowerListResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *RelationFollowerListResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetStatusCode())
	return offset
}

func (x *RelationFollowerListResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetStatusMsg())
	return offset
}

func (x *RelationFollowerListResponse) fastWriteField3(buf []byte) (offset int) {
	if x.UserList == nil {
		return offset
	}
	for i := range x.GetUserList() {
		offset += fastpb.WriteMessage(buf[offset:], 3, x.GetUserList()[i])
	}
	return offset
}

func (x *RelationFriendListRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *RelationFriendListRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *RelationFriendListRequest) fastWriteField2(buf []byte) (offset int) {
	if x.Token == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetToken())
	return offset
}

func (x *RelationFriendListResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *RelationFriendListResponse) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetStatusCode())
	return offset
}

func (x *RelationFriendListResponse) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMsg == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetStatusMsg())
	return offset
}

func (x *RelationFriendListResponse) fastWriteField3(buf []byte) (offset int) {
	if x.UserList == nil {
		return offset
	}
	for i := range x.GetUserList() {
		offset += fastpb.WriteMessage(buf[offset:], 3, x.GetUserList()[i])
	}
	return offset
}

func (x *FriendUser) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	offset += x.fastWriteField8(buf[offset:])
	offset += x.fastWriteField9(buf[offset:])
	offset += x.fastWriteField10(buf[offset:])
	offset += x.fastWriteField11(buf[offset:])
	offset += x.fastWriteField12(buf[offset:])
	offset += x.fastWriteField13(buf[offset:])
	return offset
}

func (x *FriendUser) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetId())
	return offset
}

func (x *FriendUser) fastWriteField2(buf []byte) (offset int) {
	if x.Name == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetName())
	return offset
}

func (x *FriendUser) fastWriteField3(buf []byte) (offset int) {
	if x.FollowCount == nil {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetFollowCount())
	return offset
}

func (x *FriendUser) fastWriteField4(buf []byte) (offset int) {
	if x.FollowerCount == nil {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 4, x.GetFollowerCount())
	return offset
}

func (x *FriendUser) fastWriteField5(buf []byte) (offset int) {
	if !x.IsFollow {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 5, x.GetIsFollow())
	return offset
}

func (x *FriendUser) fastWriteField6(buf []byte) (offset int) {
	if x.Avatar == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetAvatar())
	return offset
}

func (x *FriendUser) fastWriteField7(buf []byte) (offset int) {
	if x.BackgroundImage == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 7, x.GetBackgroundImage())
	return offset
}

func (x *FriendUser) fastWriteField8(buf []byte) (offset int) {
	if x.Signature == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 8, x.GetSignature())
	return offset
}

func (x *FriendUser) fastWriteField9(buf []byte) (offset int) {
	if x.TotalFavorited == nil {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 9, x.GetTotalFavorited())
	return offset
}

func (x *FriendUser) fastWriteField10(buf []byte) (offset int) {
	if x.WorkCount == nil {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 10, x.GetWorkCount())
	return offset
}

func (x *FriendUser) fastWriteField11(buf []byte) (offset int) {
	if x.FavoriteCount == nil {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 11, x.GetFavoriteCount())
	return offset
}

func (x *FriendUser) fastWriteField12(buf []byte) (offset int) {
	if x.Message == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 12, x.GetMessage())
	return offset
}

func (x *FriendUser) fastWriteField13(buf []byte) (offset int) {
	if x.MsgType == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 13, x.GetMsgType())
	return offset
}

func (x *RelationRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *RelationRequest) sizeField1() (n int) {
	if x.Token == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetToken())
	return n
}

func (x *RelationRequest) sizeField2() (n int) {
	if x.ToUserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetToUserId())
	return n
}

func (x *RelationRequest) sizeField3() (n int) {
	if x.ActionType == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, x.GetActionType())
	return n
}

func (x *RelationResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *RelationResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetStatusCode())
	return n
}

func (x *RelationResponse) sizeField2() (n int) {
	if x.StatusMsg == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetStatusMsg())
	return n
}

func (x *RelationFollowListRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *RelationFollowListRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetUserId())
	return n
}

func (x *RelationFollowListRequest) sizeField2() (n int) {
	if x.Token == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetToken())
	return n
}

func (x *RelationFollowListResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *RelationFollowListResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetStatusCode())
	return n
}

func (x *RelationFollowListResponse) sizeField2() (n int) {
	if x.StatusMsg == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetStatusMsg())
	return n
}

func (x *RelationFollowListResponse) sizeField3() (n int) {
	if x.UserList == nil {
		return n
	}
	for i := range x.GetUserList() {
		n += fastpb.SizeMessage(3, x.GetUserList()[i])
	}
	return n
}

func (x *RelationFollowerListRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *RelationFollowerListRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetUserId())
	return n
}

func (x *RelationFollowerListRequest) sizeField2() (n int) {
	if x.Token == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetToken())
	return n
}

func (x *RelationFollowerListResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *RelationFollowerListResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetStatusCode())
	return n
}

func (x *RelationFollowerListResponse) sizeField2() (n int) {
	if x.StatusMsg == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetStatusMsg())
	return n
}

func (x *RelationFollowerListResponse) sizeField3() (n int) {
	if x.UserList == nil {
		return n
	}
	for i := range x.GetUserList() {
		n += fastpb.SizeMessage(3, x.GetUserList()[i])
	}
	return n
}

func (x *RelationFriendListRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *RelationFriendListRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetUserId())
	return n
}

func (x *RelationFriendListRequest) sizeField2() (n int) {
	if x.Token == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetToken())
	return n
}

func (x *RelationFriendListResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *RelationFriendListResponse) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetStatusCode())
	return n
}

func (x *RelationFriendListResponse) sizeField2() (n int) {
	if x.StatusMsg == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetStatusMsg())
	return n
}

func (x *RelationFriendListResponse) sizeField3() (n int) {
	if x.UserList == nil {
		return n
	}
	for i := range x.GetUserList() {
		n += fastpb.SizeMessage(3, x.GetUserList()[i])
	}
	return n
}

func (x *FriendUser) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	n += x.sizeField7()
	n += x.sizeField8()
	n += x.sizeField9()
	n += x.sizeField10()
	n += x.sizeField11()
	n += x.sizeField12()
	n += x.sizeField13()
	return n
}

func (x *FriendUser) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetId())
	return n
}

func (x *FriendUser) sizeField2() (n int) {
	if x.Name == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetName())
	return n
}

func (x *FriendUser) sizeField3() (n int) {
	if x.FollowCount == nil {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetFollowCount())
	return n
}

func (x *FriendUser) sizeField4() (n int) {
	if x.FollowerCount == nil {
		return n
	}
	n += fastpb.SizeInt64(4, x.GetFollowerCount())
	return n
}

func (x *FriendUser) sizeField5() (n int) {
	if !x.IsFollow {
		return n
	}
	n += fastpb.SizeBool(5, x.GetIsFollow())
	return n
}

func (x *FriendUser) sizeField6() (n int) {
	if x.Avatar == nil {
		return n
	}
	n += fastpb.SizeString(6, x.GetAvatar())
	return n
}

func (x *FriendUser) sizeField7() (n int) {
	if x.BackgroundImage == nil {
		return n
	}
	n += fastpb.SizeString(7, x.GetBackgroundImage())
	return n
}

func (x *FriendUser) sizeField8() (n int) {
	if x.Signature == nil {
		return n
	}
	n += fastpb.SizeString(8, x.GetSignature())
	return n
}

func (x *FriendUser) sizeField9() (n int) {
	if x.TotalFavorited == nil {
		return n
	}
	n += fastpb.SizeInt64(9, x.GetTotalFavorited())
	return n
}

func (x *FriendUser) sizeField10() (n int) {
	if x.WorkCount == nil {
		return n
	}
	n += fastpb.SizeInt64(10, x.GetWorkCount())
	return n
}

func (x *FriendUser) sizeField11() (n int) {
	if x.FavoriteCount == nil {
		return n
	}
	n += fastpb.SizeInt64(11, x.GetFavoriteCount())
	return n
}

func (x *FriendUser) sizeField12() (n int) {
	if x.Message == nil {
		return n
	}
	n += fastpb.SizeString(12, x.GetMessage())
	return n
}

func (x *FriendUser) sizeField13() (n int) {
	if x.MsgType == 0 {
		return n
	}
	n += fastpb.SizeInt64(13, x.GetMsgType())
	return n
}

var fieldIDToName_RelationRequest = map[int32]string{
	1: "Token",
	2: "ToUserId",
	3: "ActionType",
}

var fieldIDToName_RelationResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
}

var fieldIDToName_RelationFollowListRequest = map[int32]string{
	1: "UserId",
	2: "Token",
}

var fieldIDToName_RelationFollowListResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "UserList",
}

var fieldIDToName_RelationFollowerListRequest = map[int32]string{
	1: "UserId",
	2: "Token",
}

var fieldIDToName_RelationFollowerListResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "UserList",
}

var fieldIDToName_RelationFriendListRequest = map[int32]string{
	1: "UserId",
	2: "Token",
}

var fieldIDToName_RelationFriendListResponse = map[int32]string{
	1: "StatusCode",
	2: "StatusMsg",
	3: "UserList",
}

var fieldIDToName_FriendUser = map[int32]string{
	1:  "Id",
	2:  "Name",
	3:  "FollowCount",
	4:  "FollowerCount",
	5:  "IsFollow",
	6:  "Avatar",
	7:  "BackgroundImage",
	8:  "Signature",
	9:  "TotalFavorited",
	10: "WorkCount",
	11: "FavoriteCount",
	12: "Message",
	13: "MsgType",
}

var _ = auth.File_user_proto
