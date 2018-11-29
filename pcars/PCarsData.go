package pcars

/*
 #cgo CFLAGS: -g -Wall
 #include <SharedMemory.h>
 #include <ReadData.h>
*/
import (
	"C"
)

// CarState is our car status
type CarState struct {
	mOilTempCelsius                 float64 // [ UNITS = Celsius ]   [ UNSET = 0.0f ]
	mOilPressureKPa                 float64 // [ UNITS = Kilopascal ]   [ RANGE = 0.0f->... ]   [ UNSET = 0.0f ]
	mWaterTempCelsius               float64 // [ UNITS = Celsius ]   [ UNSET = 0.0f ]
	mWaterPressureKPa               float64 // [ UNITS = Kilopascal ]   [ RANGE = 0.0f->... ]   [ UNSET = 0.0f ]
	mFuelPressureKPa                float64 // [ UNITS = Kilopascal ]   [ RANGE = 0.0f->... ]   [ UNSET = 0.0f ]
	mFuelLevel                      float64 // [ RANGE = 0.0f->1.0f ]
	mFuelCapacity                   float64 // [ UNITS = Liters ]   [ RANGE = 0.0f->1.0f ]   [ UNSET = 0.0f ]
	mSpeed                          float64 // [ UNITS = Metres per-second ]   [ RANGE = 0.0f->... ]
	mRpm                            float64 // [ UNITS = Revolutions per minute ]   [ RANGE = 0.0f->... ]   [ UNSET = 0.0f ]
	mMaxRPM                         float64 // [ UNITS = Revolutions per minute ]   [ RANGE = 0.0f->... ]   [ UNSET = 0.0f ]
	mBrake                          float64 // [ RANGE = 0.0f->1.0f ]
	mThrottle                       float64 // [ RANGE = 0.0f->1.0f ]
	mClutch                         float64 // [ RANGE = 0.0f->1.0f ]
	mSteering                       float64 // [ RANGE = -1.0f->1.0f ]
	mGear                           int64   // [ RANGE = -1 (Reverse)  0 (Neutral)  1 (Gear 1)  2 (Gear 2)  etc... ]   [ UNSET = 0 (Neutral) ]
	mNumGears                       int64   // [ RANGE = 0->... ]   [ UNSET = -1 ]
	mOdometerKM                     float64 // [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mAntiLockActive                 bool    // [ UNITS = boolean ]   [ RANGE = false->true ]   [ UNSET = false ]
	mLastOpponentCollisionIndex     int64   // [ RANGE = 0->STORED_PARTICIPANTS_MAX ]   [ UNSET = -1 ]
	mLastOpponentCollisionMagnitude float64 // [ RANGE = 0.0f->... ]
	mBoostActive                    bool    // [ UNITS = boolean ]   [ RANGE = false->true ]   [ UNSET = false ]
	mBoostAmount                    float64 // [ RANGE = 0.0f->100.0f ]
}

// GetCarState returns RPM from PCars2 shared memory
func GetCarState(memAddr string) CarState {
	localCpy := C.getSharedMemory(memAddr)

	return CarState{}
}
