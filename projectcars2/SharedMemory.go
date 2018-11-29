package projectcars2

// *** Types ***

// SharedMemoryVersion Header version number to test against
const SharedMemoryVersion = 9

// StringLengthMax Maximum allowed length of string
const StringLengthMax = 64

// StoredParticipantsMax Maximum number of general participant information allowed to be stored in memory-mapped file
const StoredParticipantsMax = 64

// TyreCompoundNameLengthMax Maximum length of a tyre compound name
const TyreCompoundNameLengthMax = 40

// Tyres data struct
type Tyres int

const (
	// TyreFrontLeft front left tyre designation
	TyreFrontLeft Tyres = iota
	// TyreFrontRight front right tyre designation
	TyreFrontRight
	// TyreRearLeft rear left tyre designation
	TyreRearLeft
	// TyreRearRight rear right tyre designation
	TyreRearRight
	// TyreMax struct maximum value
	TyreMax
)

// Vector data struct
type Vector int

const (
	// VecX X vector
	VecX Vector = iota
	// VecY Y vector
	VecY
	// VecZ Z vector
	VecZ
	// VecMax struct maximum value
	VecMax
)

// GameState (Type#1) (to be used with 'mGameState')
type GameState int

const (
	// GameExited current game state
	GameExited GameState = iota
	// GameFrontEnd current game state
	GameFrontEnd
	// GameIngamePlaying current game state
	GameIngamePlaying
	// GameIngamePaused current game state
	GameIngamePaused
	// GameIngameInmenuTimeTicking current game state
	GameIngameInmenuTimeTicking
	// GameIngameRestarting current game state
	GameIngameRestarting
	// GameIngameReplay current game state
	GameIngameReplay
	// GameFrontEndReplay current game state
	GameFrontEndReplay
	// GameMax struct maximum value
	GameMax
)

// SessionState (Type#2) (to be used with 'mSessionState')
type SessionState int

const (
	// SessionInvalid state
	SessionInvalid SessionState = iota
	// SessionPractice state
	SessionPractice
	// SessionTest state
	SessionTest
	// SessionQualify state
	SessionQualify
	// SessionFormationLap state
	SessionFormationLap
	// SessionRace state
	SessionRace
	// SessionTimeAttack state
	SessionTimeAttack
	// SessionMax struct maximum value
	SessionMax
)

// RaceState (Type#3) (to be used with 'mRaceState' and 'mRaceStates')
type RaceState int

const (
	// RacestateInvalid race state
	RacestateInvalid RaceState = iota
	// RacestateNotStarted race state
	RacestateNotStarted
	// RacestateRacing race state
	RacestateRacing
	// RacestateFinished race state
	RacestateFinished
	// RacestateDisqualified race state
	RacestateDisqualified
	// RacestateRetired race state
	RacestateRetired
	// RacestateDnf race state
	RacestateDnf
	//RacestateMax struct maximum value
	RacestateMax
)

// FlagColours (Type#5) (to be used with 'mHighestFlagColour')
type FlagColours int

const (
	// FlagColourNone Not used for actual flags, only for some query functions
	FlagColourNone FlagColours = iota
	// FlagColourGreen End of danger zone, or race started
	FlagColourGreen
	// FlagColourBlue Faster car wants to overtake the participant
	FlagColourBlue
	// FlagColourWhiteSlowCar Slow car in area
	FlagColourWhiteSlowCar
	// FlagColourWhiteFinalLap Final Lap
	FlagColourWhiteFinalLap
	// FlagColourRed Huge collisions where one or more cars become wrecked and block the track
	FlagColourRed
	// FlagColourYellow Danger on the racing surface itself
	FlagColourYellow
	// FlagColourDoubleYellow Danger that wholly or partly blocks the racing surface
	FlagColourDoubleYellow
	// FlagColourBlackAndWhite Unsportsmanlike conduct
	FlagColourBlackAndWhite
	// FlagColourBlackOrangeCircle Mechanical Failure
	FlagColourBlackOrangeCircle
	// FlagColourBlack Participant disqualified
	FlagColourBlack
	// FlagColourChequered Chequered flag
	FlagColourChequered
	// FlagColourMax struct maximum value
	FlagColourMax
)

// FlagReason (Type#6) (to be used with 'mHighestFlagReason')
type FlagReason int

const (
	// FlagReasonNone flag reason
	FlagReasonNone FlagReason = iota
	// FlagReasonSoloCrash flag reason
	FlagReasonSoloCrash
	// FlagReasonVehicleCrash flag reason
	FlagReasonVehicleCrash
	// FlagReasonVehicleObstruction flag reason
	FlagReasonVehicleObstruction
	// FlagReasonMax struct maximum value
	FlagReasonMax
)

// PitMode (Type#7) (to be used with 'mPitMode')
type PitMode int

const (
	// PitModeNone current pit mode
	PitModeNone PitMode = iota
	// PitModeDrivingIntoPits current pit mode
	PitModeDrivingIntoPits
	// PitModeInPit current pit mode
	PitModeInPit
	// PitModeDrivingOutOfPits current pit mode
	PitModeDrivingOutOfPits
	// PitModeInGarage current pit mode
	PitModeInGarage
	// PitModeDrivingOutOfGarage current pit mode
	PitModeDrivingOutOfGarage
	// PitModeMax struct maximum value current pit mode
	PitModeMax
)

// PitStopSchedule (Type#8) (to be used with 'mPitSchedule')
type PitStopSchedule int

const (
	// PitScheduleNone Nothing scheduled
	PitScheduleNone PitStopSchedule = iota
	// PitSchedulePlayerRequested used for standard pit sequence - requested by player
	PitSchedulePlayerRequested
	// PitScheduleEngineerRequested used for standard pit sequence - requested by engineer
	PitScheduleEngineerRequested
	// PitScheduleDamageRequested used for standard pit sequence - requested by engineer for damage
	PitScheduleDamageRequested
	// PitScheduleMandatory used for standard pit sequence - requested by engineer from career enforced lap number
	PitScheduleMandatory
	// PitScheduleDriveThrough used for drive-through penalty
	PitScheduleDriveThrough
	// PitScheduleStopGo used for stop-go penalty
	PitScheduleStopGo
	// PitSchedulePitspotOccupied used for drive-through when pitspot is occupied
	PitSchedulePitspotOccupied
	// PitScheduleMax for mem safety
	PitScheduleMax
)

// CarFlags (Type#9) (to be used with 'mCarFlags')
type CarFlags int

const (
	// CarHeadlight Car status flags
	CarHeadlight CarFlags = (1 << iota)
	// CarEngineActive Car status flags
	CarEngineActive CarFlags = (1 << iota)
	// CarEngineWarning Car status flags
	CarEngineWarning CarFlags = (1 << iota)
	// CarSpeedLimiter Car status flags
	CarSpeedLimiter CarFlags = (1 << iota)
	// CarAbs Car status flags
	CarAbs CarFlags = (1 << iota)
	// CarHandbrake Car status flags
	CarHandbrake CarFlags = (1 << iota)
)

// TyreFlags (Type#10) (to be used with 'mTyreFlags')
type TyreFlags int

const (
	// TyreAttached Tyre status flags
	TyreAttached TyreFlags = (1 << iota)
	// TyreInflated  Tyre status flags
	TyreInflated TyreFlags = (1 << iota)
	// TyreIsOnGround  Tyre status flags
	TyreIsOnGround TyreFlags = (1 << iota)
)

// Terrain (Type#11) Materials (to be used with 'mTerrain')
type Terrain int

const (
	// TerrainRoad terrain status
	TerrainRoad Terrain = iota
	// TerrainLowGripRoad terrain status
	TerrainLowGripRoad
	// TerrainBumpyRoad1 terrain status
	TerrainBumpyRoad1
	// TerrainBumpyRoad2 terrain status
	TerrainBumpyRoad2
	// TerrainBumpyRoad3 terrain status
	TerrainBumpyRoad3
	// TerrainMarbles terrain status
	TerrainMarbles
	// TerrainGrassyBerms terrain status
	TerrainGrassyBerms
	// TerrainGrass terrain status
	TerrainGrass
	// TerrainGravel terrain status
	TerrainGravel
	// TerrainBumpyGravel terrain status
	TerrainBumpyGravel
	// TerrainRumbleStrips terrain status
	TerrainRumbleStrips
	// TerrainDrains terrain status
	TerrainDrains
	// TerrainTyrewalls terrain status
	TerrainTyrewalls
	// TerrainCementwalls terrain status
	TerrainCementwalls
	// TerrainGuardrails terrain status
	TerrainGuardrails
	// TerrainSand terrain status
	TerrainSand
	// TerrainBumpySand terrain status
	TerrainBumpySand
	// TerrainDirt terrain status
	TerrainDirt
	// TerrainBumpyDirt terrain status
	TerrainBumpyDirt
	// TerrainDirtRoad terrain status
	TerrainDirtRoad
	// TerrainBumpyDirtRoad terrain status
	TerrainBumpyDirtRoad
	// TerrainPavement terrain status
	TerrainPavement
	// TerrainDirtBank terrain status
	TerrainDirtBank
	// TerrainWood terrain status
	TerrainWood
	// TerrainDryVerge terrain status
	TerrainDryVerge
	// TerrainExitRumbleStrips terrain status
	TerrainExitRumbleStrips
	// TerrainGrasscrete terrain status
	TerrainGrasscrete
	// TerrainLongGrass terrain status
	TerrainLongGrass
	// TerrainSlopeGrass terrain status
	TerrainSlopeGrass
	// TerrainCobbles terrain status
	TerrainCobbles
	// TerrainSandRoad terrain status
	TerrainSandRoad
	// TerrainBakedClay terrain status
	TerrainBakedClay
	// TerrainAstroturf terrain status
	TerrainAstroturf
	// TerrainSnowhalf terrain status
	TerrainSnowhalf
	// TerrainSnowfull terrain status
	TerrainSnowfull
	// TerrainDamagedRoad1 terrain status
	TerrainDamagedRoad1
	// TerrainTrainTrackRoad terrain status
	TerrainTrainTrackRoad
	// TerrainBumpycobbles terrain status
	TerrainBumpycobbles
	// TerrainAriesOnly terrain status
	TerrainAriesOnly
	// TerrainOrionOnly terrain status
	TerrainOrionOnly
	// TerrainB1rumbles terrain status
	TerrainB1rumbles
	// TerrainB2rumbles terrain status
	TerrainB2rumbles
	// TerrainRoughSandMedium terrain status
	TerrainRoughSandMedium
	// TerrainRoughSandHeavy terrain status
	TerrainRoughSandHeavy
	// TerrainSnowwalls terrain status
	TerrainSnowwalls
	// TerrainIceRoad terrain status
	TerrainIceRoad
	// TerrainRunoffRoad terrain status
	TerrainRunoffRoad
	// TerrainIllegalStrip terrain status
	TerrainIllegalStrip
	// TerrainPaintConcrete terrain status
	TerrainPaintConcrete
	// TerrainPaintConcreteIllegal terrain status
	TerrainPaintConcreteIllegal
	// TerrainRallyTarmac terrain status
	TerrainRallyTarmac

	// TerrainMax struct max value
	TerrainMax
)

// CrashState Damage (Type#12) (to be used with 'mCrashState')
type CrashState int

const (
	// CrashDamageNone crash state
	CrashDamageNone CrashState = iota
	// CrashDamageOfftrack crash state
	CrashDamageOfftrack
	// CrashDamageLargeProp crash state
	CrashDamageLargeProp
	// CrashDamageSpinning crash state
	CrashDamageSpinning
	// CrashDamageRolling crash state
	CrashDamageRolling
	//CrashMax struct max value
	CrashMax
)

// ParticipantInfo struct (Type#13) (to be used with 'mParticipantInfo')
type ParticipantInfo struct {
	mIsActive           bool
	mName               [StringLengthMax]byte // [ string ]
	mWorldPosition      [VecMax]float32       // [ UNITS = World Space  X  Y  Z ]
	mCurrentLapDistance float32               // [ UNITS = Metres ]   [ RANGE = 0.0f->... ]    [ UNSET = 0.0f ]
	mRacePosition       uint16                // [ RANGE = 1->... ]   [ UNSET = 0 ]
	mLapsCompleted      uint16                // [ RANGE = 0->... ]   [ UNSET = 0 ]
	mCurrentLap         uint16                // [ RANGE = 0->... ]   [ UNSET = 0 ]
	mCurrentSector      int16                 // [ RANGE = 0->... ]   [ UNSET = -1 ]
}

// SharedMemory is our struct containing Project Cars 2's telemetry data
type SharedMemory struct {
	// Version Number
	mVersion            uint16 // [ RANGE = 0->... ]
	mBuildVersionNumber uint16 // [ RANGE = 0->... ]   [ UNSET = 0 ]

	// Game States
	mGameState    uint16 // [ enum (Type#1) Game state ]
	mSessionState uint16 // [ enum (Type#2) Session state ]
	mRaceState    uint16 // [ enum (Type#3) Race State ]

	// Participant Info
	mViewedParticipantIndex int16                                  // [ RANGE = 0->StoredParticipantsMax ]   [ UNSET = -1 ]
	mNumParticipants        int16                                  // [ RANGE = 0->StoredParticipantsMax ]   [ UNSET = -1 ]
	mParticipantInfo        [StoredParticipantsMax]ParticipantInfo // [ struct (Type#13) ParticipantInfo struct ]

	// Unfiltered Input
	mUnfilteredThrottle float32 // [ RANGE = 0.0f->1.0f ]
	mUnfilteredBrake    float32 // [ RANGE = 0.0f->1.0f ]
	mUnfilteredSteering float32 // [ RANGE = -1.0f->1.0f ]
	mUnfilteredClutch   float32 // [ RANGE = 0.0f->1.0f ]

	// Vehicle information
	mCarName      [StringLengthMax]byte // [ string ]
	mCarClassName [StringLengthMax]byte // [ string ]

	// Event information
	mLapsInEvent    uint16                // [ RANGE = 0->... ]   [ UNSET = 0 ]
	mTrackLocation  [StringLengthMax]byte // [ string ] - untranslated shortened English name
	mTrackVariation [StringLengthMax]byte // [ string ]- untranslated shortened English variation description
	mTrackLength    float32               // [ UNITS = Metres ]   [ RANGE = 0.0f->... ]    [ UNSET = 0.0f ]

	// Timings
	mNumSectors                 int16   // [ RANGE = 0->... ]   [ UNSET = -1 ]
	mLapInvalidated             bool    // [ UNITS = boolean ]   [ RANGE = false->true ]   [ UNSET = false ]
	mBestLapTime                float32 // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mLastLapTime                float32 // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = 0.0f ]
	mCurrentTime                float32 // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = 0.0f ]
	mSplitTimeAhead             float32 // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mSplitTimeBehind            float32 // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mSplitTime                  float32 // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = 0.0f ]
	mEventTimeRemaining         float32 // [ UNITS = milli-seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mPersonalFastestLapTime     float32 // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mWorldFastestLapTime        float32 // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mCurrentSector1Time         float32 // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mCurrentSector2Time         float32 // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mCurrentSector3Time         float32 // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mFastestSector1Time         float32 // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mFastestSector2Time         float32 // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mFastestSector3Time         float32 // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mPersonalFastestSector1Time float32 // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mPersonalFastestSector2Time float32 // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mPersonalFastestSector3Time float32 // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mWorldFastestSector1Time    float32 // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mWorldFastestSector2Time    float32 // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mWorldFastestSector3Time    float32 // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]

	// Flags
	mHighestFlagColour uint16 // [ enum (Type#5) Flag Colour ]
	mHighestFlagReason uint16 // [ enum (Type#6) Flag Reason ]

	// Pit Info
	mPitMode     uint16 // [ enum (Type#7) Pit Mode ]
	mPitSchedule uint16 // [ enum (Type#8) Pit Stop Schedule ]

	// Car State
	mCarFlags                       uint16  // [ enum (Type#9) Car Flags ]
	mOilTempCelsius                 float32 // [ UNITS = Celsius ]   [ UNSET = 0.0f ]
	mOilPressureKPa                 float32 // [ UNITS = Kilopascal ]   [ RANGE = 0.0f->... ]   [ UNSET = 0.0f ]
	mWaterTempCelsius               float32 // [ UNITS = Celsius ]   [ UNSET = 0.0f ]
	mWaterPressureKPa               float32 // [ UNITS = Kilopascal ]   [ RANGE = 0.0f->... ]   [ UNSET = 0.0f ]
	mFuelPressureKPa                float32 // [ UNITS = Kilopascal ]   [ RANGE = 0.0f->... ]   [ UNSET = 0.0f ]
	mFuelLevel                      float32 // [ RANGE = 0.0f->1.0f ]
	mFuelCapacity                   float32 // [ UNITS = Liters ]   [ RANGE = 0.0f->1.0f ]   [ UNSET = 0.0f ]
	mSpeed                          float32 // [ UNITS = Metres per-second ]   [ RANGE = 0.0f->... ]
	mRpm                            float32 // [ UNITS = Revolutions per minute ]   [ RANGE = 0.0f->... ]   [ UNSET = 0.0f ]
	mMaxRPM                         float32 // [ UNITS = Revolutions per minute ]   [ RANGE = 0.0f->... ]   [ UNSET = 0.0f ]
	mBrake                          float32 // [ RANGE = 0.0f->1.0f ]
	mThrottle                       float32 // [ RANGE = 0.0f->1.0f ]
	mClutch                         float32 // [ RANGE = 0.0f->1.0f ]
	mSteering                       float32 // [ RANGE = -1.0f->1.0f ]
	mGear                           int16   // [ RANGE = -1 (Reverse)  0 (Neutral)  1 (Gear 1)  2 (Gear 2)  etc... ]   [ UNSET = 0 (Neutral) ]
	mNumGears                       int16   // [ RANGE = 0->... ]   [ UNSET = -1 ]
	mOdometerKM                     float32 // [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mAntiLockActive                 bool    // [ UNITS = boolean ]   [ RANGE = false->true ]   [ UNSET = false ]
	mLastOpponentCollisionIndex     int16   // [ RANGE = 0->StoredParticipantsMax ]   [ UNSET = -1 ]
	mLastOpponentCollisionMagnitude float32 // [ RANGE = 0.0f->... ]
	mBoostActive                    bool    // [ UNITS = boolean ]   [ RANGE = false->true ]   [ UNSET = false ]
	mBoostAmount                    float32 // [ RANGE = 0.0f->100.0f ]

	// Motion & Device Related
	mOrientation       [VecMax]float32 // [ UNITS = Euler Angles ]
	mLocalVelocity     [VecMax]float32 // [ UNITS = Metres per-second ]
	mWorldVelocity     [VecMax]float32 // [ UNITS = Metres per-second ]
	mAngularVelocity   [VecMax]float32 // [ UNITS = Radians per-second ]
	mLocalAcceleration [VecMax]float32 // [ UNITS = Metres per-second ]
	mWorldAcceleration [VecMax]float32 // [ UNITS = Metres per-second ]
	mExtentsCentre     [VecMax]float32 // [ UNITS = Local Space  X  Y  Z ]

	// Wheels / Tyres
	mTyreFlags             [TyreMax]uint16  // [ enum (Type#10) Tyre Flags ]
	mTerrain               [TyreMax]uint16  // [ enum (Type#11) Terrain Materials ]
	mTyreY                 [TyreMax]float32 // [ UNITS = Local Space  Y ]
	mTyreRPS               [TyreMax]float32 // [ UNITS = Revolutions per second ]
	fTyreSlipSpeed         [TyreMax]float32 // OBSOLETE, kept for backward compatibility only
	mTyreTemp              [TyreMax]float32 // [ UNITS = Celsius ]   [ UNSET = 0.0f ]
	fTyreGrip              [TyreMax]float32 // OBSOLETE, kept for backward compatibility only
	mTyreHeightAboveGround [TyreMax]float32 // [ UNITS = Local Space  Y ]
	fTyreLateralStiffness  [TyreMax]float32 // OBSOLETE, kept for backward compatibility only
	mTyreWear              [TyreMax]float32 // [ RANGE = 0.0f->1.0f ]
	mBrakeDamage           [TyreMax]float32 // [ RANGE = 0.0f->1.0f ]
	mSuspensionDamage      [TyreMax]float32 // [ RANGE = 0.0f->1.0f ]
	mBrakeTempCelsius      [TyreMax]float32 // [ UNITS = Celsius ]
	mTyreTreadTemp         [TyreMax]float32 // [ UNITS = Kelvin ]
	mTyreLayerTemp         [TyreMax]float32 // [ UNITS = Kelvin ]
	mTyreCarcassTemp       [TyreMax]float32 // [ UNITS = Kelvin ]
	mTyreRimTemp           [TyreMax]float32 // [ UNITS = Kelvin ]
	mTyreInternalAirTemp   [TyreMax]float32 // [ UNITS = Kelvin ]

	// Car Damage
	mCrashState   uint16  // [ enum (Type#12) Crash Damage State ]
	mAeroDamage   float32 // [ RANGE = 0.0f->1.0f ]
	mEngineDamage float32 // [ RANGE = 0.0f->1.0f ]

	// Weather
	mAmbientTemperature float32 // [ UNITS = Celsius ]   [ UNSET = 25.0f ]
	mTrackTemperature   float32 // [ UNITS = Celsius ]   [ UNSET = 30.0f ]
	mRainDensity        float32 // [ UNITS = How much rain will fall ]   [ RANGE = 0.0f->1.0f ]
	mWindSpeed          float32 // [ RANGE = 0.0f->100.0f ]   [ UNSET = 2.0f ]
	mWindDirectionX     float32 // [ UNITS = Normalised Vector X ]
	mWindDirectionY     float32 // [ UNITS = Normalised Vector Y ]
	mCloudBrightness    float32 // [ RANGE = 0.0f->... ]

	//PCars2 additions start, version 8
	// Sequence Number to help slightly with data integrity reads
	// OG volatile unsigned int mSequenceNumber
	mSequenceNumber uint16 // 0 at the start, incremented at start and end of writing, so odd when Shared Memory is being filled, even when the memory is not being touched

	//Additional car variables
	mWheelLocalPositionY [TyreMax]float32 // [ UNITS = Local Space  Y ]
	mSuspensionTravel    [TyreMax]float32 // [ UNITS = meters ] [ RANGE 0.f =>... ]  [ UNSET =  0.0f ]
	mSuspensionVelocity  [TyreMax]float32 // [ UNITS = Rate of change of pushrod deflection ] [ RANGE 0.f =>... ]  [ UNSET =  0.0f ]
	mAirPressure         [TyreMax]float32 // [ UNITS = PSI ]  [ RANGE 0.f =>... ]  [ UNSET =  0.0f ]
	mEngineSpeed         float32          // [ UNITS = Rad/s ] [UNSET = 0.f ]
	mEngineTorque        float32          // [ UNITS = Newton Meters] [UNSET = 0.f ] [ RANGE = 0.0f->... ]
	mWings               [2]float32       // [ RANGE = 0.0f->1.0f ] [UNSET = 0.f ]
	mHandBrake           float32          // [ RANGE = 0.0f->1.0f ] [UNSET = 0.f ]

	// additional race variables
	mCurrentSector1Times [StoredParticipantsMax]float32               // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mCurrentSector2Times [StoredParticipantsMax]float32               // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mCurrentSector3Times [StoredParticipantsMax]float32               // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mFastestSector1Times [StoredParticipantsMax]float32               // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mFastestSector2Times [StoredParticipantsMax]float32               // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mFastestSector3Times [StoredParticipantsMax]float32               // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mFastestLapTimes     [StoredParticipantsMax]float32               // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mLastLapTimes        [StoredParticipantsMax]float32               // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mLapsInvalidated     [StoredParticipantsMax]bool                  // [ UNITS = boolean for all participants ]   [ RANGE = false->true ]   [ UNSET = false ]
	mRaceStates          [StoredParticipantsMax]uint16                // [ enum (Type#3) Race State ]
	mPitModes            [StoredParticipantsMax]uint16                // [ enum (Type#7)  Pit Mode ]
	mOrientations        [StoredParticipantsMax][VecMax]float32       // [ UNITS = Euler Angles ]
	mSpeeds              [StoredParticipantsMax]float32               // [ UNITS = Metres per-second ]   [ RANGE = 0.0f->... ]
	mCarNames            [StoredParticipantsMax][StringLengthMax]byte // [ string ]
	mCarClassNames       [StoredParticipantsMax][StringLengthMax]byte // [ string ]

	// additional race variables
	mEnforcedPitStopLap       int16                                    // [ UNITS = in which lap there will be a mandatory pitstop] [ RANGE = 0.0f->... ] [ UNSET = -1 ]
	mTranslatedTrackLocation  [StringLengthMax]byte                    // [ string ]
	mTranslatedTrackVariation [StringLengthMax]byte                    // [ string ]
	mBrakeBias                float32                                  // [ RANGE = 0.0f->1.0f... ]   [ UNSET = -1.0f ]
	mTurboBoostPressure       float32                                  //	 RANGE = 0.0f->1.0f... ]   [ UNSET = -1.0f ]
	mTyreCompound             [TyreMax][TyreCompoundNameLengthMax]byte // [ strings  ]
	mPitSchedules             [StoredParticipantsMax]uint16            // [ enum (Type#7)  Pit Mode ]
	mHighestFlagColours       [StoredParticipantsMax]uint16            // [ enum (Type#5) Flag Colour ]
	mHighestFlagReasons       [StoredParticipantsMax]uint16            // [ enum (Type#6) Flag Reason ]
	mNationalities            [StoredParticipantsMax]uint16            // [ nationality table , SP AND UNSET = 0 ] See nationalities.txt file for details
	mSnowDensity              float32                                  // [ UNITS = How much snow will fall ]   [ RANGE = 0.0f->1.0f ], this will be non zero only in Snow season, in other seasons whatever is falling from the sky is reported as rain

}
