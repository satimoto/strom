// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"fmt"
	"time"
)

type CallAction string

const (
	CallActionAuthorize                     CallAction = "Authorize"
	CallActionBootNotification              CallAction = "BootNotification"
	CallActionCancelReservation             CallAction = "CancelReservation"
	CallActionChangeAvailability            CallAction = "ChangeAvailability"
	CallActionChangeConfiguration           CallAction = "ChangeConfiguration"
	CallActionClearCache                    CallAction = "ClearCache"
	CallActionClearChargingProfile          CallAction = "ClearChargingProfile"
	CallActionDataTransfer                  CallAction = "DataTransfer"
	CallActionDiagnosticsStatusNotification CallAction = "DiagnosticsStatusNotification"
	CallActionFirmwareStatusNotification    CallAction = "FirmwareStatusNotification"
	CallActionGetCompositeSchedule          CallAction = "GetCompositeSchedule"
	CallActionGetConfiguration              CallAction = "GetConfiguration"
	CallActionGetDiagnostics                CallAction = "GetDiagnostics"
	CallActionGetLocalListVersion           CallAction = "GetLocalListVersion"
	CallActionHeartbeat                     CallAction = "Heartbeat"
	CallActionMeterValues                   CallAction = "MeterValues"
	CallActionRemoteStartTransaction        CallAction = "RemoteStartTransaction"
	CallActionRemoteStopTransaction         CallAction = "RemoteStopTransaction"
	CallActionReserveNow                    CallAction = "ReserveNow"
	CallActionReset                         CallAction = "Reset"
	CallActionSendLocalList                 CallAction = "SendLocalList"
	CallActionSetChargingProfile            CallAction = "SetChargingProfile"
	CallActionStartTransaction              CallAction = "StartTransaction"
	CallActionStatusNotification            CallAction = "StatusNotification"
	CallActionStopTransaction               CallAction = "StopTransaction"
	CallActionTriggerMessage                CallAction = "TriggerMessage"
	CallActionUnlockConnector               CallAction = "UnlockConnector"
	CallActionUpdateFirmware                CallAction = "UpdateFirmware"
)

func (e *CallAction) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CallAction(s)
	case string:
		*e = CallAction(s)
	default:
		return fmt.Errorf("unsupported scan type for CallAction: %T", src)
	}
	return nil
}

type ChargePointErrorCode string

const (
	ChargePointErrorCodeConnectorLockFailure ChargePointErrorCode = "ConnectorLockFailure"
	ChargePointErrorCodeEVCommunicationError ChargePointErrorCode = "EVCommunicationError"
	ChargePointErrorCodeGroundFailure        ChargePointErrorCode = "GroundFailure"
	ChargePointErrorCodeHighTemperature      ChargePointErrorCode = "HighTemperature"
	ChargePointErrorCodeInternalError        ChargePointErrorCode = "InternalError"
	ChargePointErrorCodeLocalListConflict    ChargePointErrorCode = "LocalListConflict"
	ChargePointErrorCodeNoError              ChargePointErrorCode = "NoError"
	ChargePointErrorCodeOtherError           ChargePointErrorCode = "OtherError"
	ChargePointErrorCodeOverCurrentFailure   ChargePointErrorCode = "OverCurrentFailure"
	ChargePointErrorCodeOverVoltage          ChargePointErrorCode = "OverVoltage"
	ChargePointErrorCodePowerMeterFailure    ChargePointErrorCode = "PowerMeterFailure"
	ChargePointErrorCodePowerSwitchFailure   ChargePointErrorCode = "PowerSwitchFailure"
	ChargePointErrorCodeReaderFailure        ChargePointErrorCode = "ReaderFailure"
	ChargePointErrorCodeResetFailure         ChargePointErrorCode = "ResetFailure"
	ChargePointErrorCodeUnderVoltage         ChargePointErrorCode = "UnderVoltage"
	ChargePointErrorCodeWeakSignal           ChargePointErrorCode = "WeakSignal"
)

func (e *ChargePointErrorCode) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ChargePointErrorCode(s)
	case string:
		*e = ChargePointErrorCode(s)
	default:
		return fmt.Errorf("unsupported scan type for ChargePointErrorCode: %T", src)
	}
	return nil
}

type ChargePointStatus string

const (
	ChargePointStatusAvailable     ChargePointStatus = "Available"
	ChargePointStatusPreparing     ChargePointStatus = "Preparing"
	ChargePointStatusCharging      ChargePointStatus = "Charging"
	ChargePointStatusSuspendedEVSE ChargePointStatus = "SuspendedEVSE"
	ChargePointStatusSuspendedEV   ChargePointStatus = "SuspendedEV"
	ChargePointStatusFinishing     ChargePointStatus = "Finishing"
	ChargePointStatusReserved      ChargePointStatus = "Reserved"
	ChargePointStatusUnavailable   ChargePointStatus = "Unavailable"
	ChargePointStatusFaulted       ChargePointStatus = "Faulted"
)

func (e *ChargePointStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ChargePointStatus(s)
	case string:
		*e = ChargePointStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for ChargePointStatus: %T", src)
	}
	return nil
}

type ReservationStatus string

const (
	ReservationStatusAccepted    ReservationStatus = "Accepted"
	ReservationStatusCancelled   ReservationStatus = "Cancelled"
	ReservationStatusCompleted   ReservationStatus = "Completed"
	ReservationStatusFaulted     ReservationStatus = "Faulted"
	ReservationStatusOccupied    ReservationStatus = "Occupied"
	ReservationStatusRejected    ReservationStatus = "Rejected"
	ReservationStatusUnavailable ReservationStatus = "Unavailable"
)

func (e *ReservationStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ReservationStatus(s)
	case string:
		*e = ReservationStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for ReservationStatus: %T", src)
	}
	return nil
}

type Call struct {
	ID            int64      `db:"id" json:"id"`
	ChargePointID int64      `db:"charge_point_id" json:"chargePointID"`
	ReqID         string     `db:"req_id" json:"reqID"`
	Action        CallAction `db:"action" json:"action"`
	CreatedAt     time.Time  `db:"created_at" json:"createdAt"`
}

type ChargePoint struct {
	ID                int64          `db:"id" json:"id"`
	Identity          string         `db:"identity" json:"identity"`
	Model             string         `db:"model" json:"model"`
	Vendor            string         `db:"vendor" json:"vendor"`
	SerialNumber      sql.NullString `db:"serial_number" json:"serialNumber"`
	FirmwareVerion    sql.NullString `db:"firmware_verion" json:"firmwareVerion"`
	ModemIccid        sql.NullString `db:"modem_iccid" json:"modemIccid"`
	ModemImsi         sql.NullString `db:"modem_imsi" json:"modemImsi"`
	MeterSerialNumber sql.NullString `db:"meter_serial_number" json:"meterSerialNumber"`
	MeterType         sql.NullString `db:"meter_type" json:"meterType"`
	CreatedAt         time.Time      `db:"created_at" json:"createdAt"`
	UpdatedAt         time.Time      `db:"updated_at" json:"updatedAt"`
}

type Connector struct {
	ID              int64                `db:"id" json:"id"`
	ConnectorID     int32                `db:"connector_id" json:"connectorID"`
	ChargePointID   int64                `db:"charge_point_id" json:"chargePointID"`
	ErrorCode       ChargePointErrorCode `db:"error_code" json:"errorCode"`
	Status          ChargePointStatus    `db:"status" json:"status"`
	Info            sql.NullString       `db:"info" json:"info"`
	VendorID        sql.NullString       `db:"vendor_id" json:"vendorID"`
	VendorErrorCode sql.NullString       `db:"vendor_error_code" json:"vendorErrorCode"`
	CreatedAt       time.Time            `db:"created_at" json:"createdAt"`
	UpdatedAt       time.Time            `db:"updated_at" json:"updatedAt"`
}

type Reservation struct {
	ID            int64             `db:"id" json:"id"`
	ConnectorID   int32             `db:"connector_id" json:"connectorID"`
	ChargePointID int64             `db:"charge_point_id" json:"chargePointID"`
	ReqID         string            `db:"req_id" json:"reqID"`
	ExpiryDate    time.Time         `db:"expiry_date" json:"expiryDate"`
	Status        ReservationStatus `db:"status" json:"status"`
	IDTag         string            `db:"id_tag" json:"idTag"`
	ParentIDTag   sql.NullString    `db:"parent_id_tag" json:"parentIDTag"`
	CreatedAt     time.Time         `db:"created_at" json:"createdAt"`
	UpdatedAt     time.Time         `db:"updated_at" json:"updatedAt"`
}
