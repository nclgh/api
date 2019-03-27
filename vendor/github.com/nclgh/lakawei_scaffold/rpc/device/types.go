package device

import (
	"time"
	"github.com/nclgh/lakawei_scaffold/rpc/common"
)

// timeFilter
type TimeFilter struct {
	Field string
	Start time.Time
	End   time.Time
}

// Manufacturer
type Manufacturer struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type AddManufacturerRequest struct {
	Name string
}

type AddManufacturerResponse struct {
	Code common.RspCode
	Msg  string
}

type GetManufacturerByIdRequest struct {
	Ids []int64
}

type GetManufacturerByIdResponse struct {
	Manufacturers map[int64]*Manufacturer
	Code          common.RspCode
	Msg           string
}

type QueryManufacturerRequest struct {
	Manufacturer *Manufacturer
	Page         int64
	PageSize     int64
}

type QueryManufacturerResponse struct {
	Manufacturers map[int64]*Manufacturer
	TotalCount    int64
	Code          common.RspCode
	Msg           string
}

// Device
type Device struct {
	Id               int64     `json:"id"`
	Code             string    `json:"code"`
	Name             string    `json:"name"`
	Model            string    `json:"model"`
	Brand            string    `json:"brand"`
	TagCode          string    `json:"tag_code"`
	DepartmentId     int64     `json:"department_id"`
	ManufacturerId   int64     `json:"manufacturer_id"`
	ManufacturerDate time.Time `json:"manufacturer_date"`
	RentStatus       int64     `json:"rent_status"`
	Description      string    `json:"description"`
}

type AddDeviceRequest struct {
	Device Device
}

type AddDeviceResponse struct {
	Code common.RspCode
	Msg  string
}

type DeleteDeviceRequest struct {
	Id int64
}

type DeleteDeviceResponse struct {
	Code common.RspCode
	Msg  string
}

type GetDeviceByIdRequest struct {
	Ids []int64
}

type GetDeviceByIdResponse struct {
	Devices map[int64]*Device
	Code    common.RspCode
	Msg     string
}

type QueryDeviceRequest struct {
	Device     *Device
	TimeFilter []*TimeFilter
	Page       int64
	PageSize   int64
}

type QueryDeviceResponse struct {
	Devices    map[int64]*Device
	TotalCount int64
	Code       common.RspCode
	Msg        string
}

// Achievement
type Achievement struct {
	Id                     int64     `json:"id"`
	DeviceId               int64     `json:"device_id"`
	MemberId               int64     `json:"member_id"`
	DepartmentId           int64     `json:"department_id"`
	AchievementDate        time.Time `json:"achievement_date"`
	AchievementDescription string    `json:"achievement_description"`
	AchievementRemark      string    `json:"achievement_remark"`
	PatentDescription      string    `json:"patent_description"`
	PaperDescription       string    `json:"paper_description"`
	CompetitionDescription string    `json:"competition_description"`
}

type AddAchievementRequest struct {
	Achievement Achievement
}

type AddAchievementResponse struct {
	Code common.RspCode
	Msg  string
}

type DeleteAchievementRequest struct {
	Id int64
}

type DeleteAchievementResponse struct {
	Code common.RspCode
	Msg  string
}

type GetAchievementByIdRequest struct {
	Ids []int64
}

type GetAchievementByIdResponse struct {
	Achievements map[int64]*Achievement
	Code         common.RspCode
	Msg          string
}

type QueryAchievementRequest struct {
	Achievement *Achievement
	TimeFilter  []*TimeFilter
	Page        int64
	PageSize    int64
}

type QueryAchievementResponse struct {
	Achievements map[int64]*Achievement
	TotalCount   int64
	Code         common.RspCode
	Msg          string
}
