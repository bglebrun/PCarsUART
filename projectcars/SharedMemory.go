package projectcars

// *** Types ***

// SHARED_MEMORY_VERSION Header version number to test against
const SHARED_MEMORY_VERSION = 9

// STRING_LENGTH_MAX Maximum allowed length of string
const STRING_LENGTH_MAX = 64

// STORED_PARTICIPANTS_MAX Maximum number of general participant information allowed to be stored in memory-mapped file
const STORED_PARTICIPANTS_MAX = 64

// TYRE_COMPOUND_NAME_LENGTH_MAX Maximum length of a tyre compound name
const TYRE_COMPOUND_NAME_LENGTH_MAX = 40

// Tyres data struct
type Tyres int

const (
	TYRE_FRONT_LEFT Tyres = iota
	TYRE_FRONT_RIGHT
	TYRE_REAR_LEFT
	TYRE_REAR_RIGHT
	//--------------
	TYRE_MAX
)

// Vector data struct
type Vector int

const (
	VEC_X Vector = iota
	VEC_Y
	VEC_Z
	//-------------
	VEC_MAX
)

// GameState (Type#1) (to be used with 'mGameState')
type GameState int

const (
	GAME_EXITED GameState = iota
	GAME_FRONT_END
	GAME_INGAME_PLAYING
	GAME_INGAME_PAUSED
	GAME_INGAME_INMENU_TIME_TICKING
	GAME_INGAME_RESTARTING
	GAME_INGAME_REPLAY
	GAME_FRONT_END_REPLAY
	//-------------
	GAME_MAX
)

// SessionState (Type#2) (to be used with 'mSessionState')
type SessionState int

const (
	SESSION_INVALID SessionState = iota
	SESSION_PRACTICE
	SESSION_TEST
	SESSION_QUALIFY
	SESSION_FORMATION_LAP
	SESSION_RACE
	SESSION_TIME_ATTACK
	//-------------
	SESSION_MAX
)

// RaceState (Type#3) (to be used with 'mRaceState' and 'mRaceStates')
type RaceState int

const (
	RACESTATE_INVALID RaceState = iota
	RACESTATE_NOT_STARTED
	RACESTATE_RACING
	RACESTATE_FINISHED
	RACESTATE_DISQUALIFIED
	RACESTATE_RETIRED
	RACESTATE_DNF
	//-------------
	RACESTATE_MAX
)

// FlagColours (Type#5) (to be used with 'mHighestFlagColour')
type FlagColours int

const (
	// FLAG_COLOUR_NONE Not used for actual flags, only for some query functions
	FLAG_COLOUR_NONE FlagColours = iota
	// FLAG_COLOUR_GREEN End of danger zone, or race started
	FLAG_COLOUR_GREEN
	// FLAG_COLOUR_BLUE Faster car wants to overtake the participant
	FLAG_COLOUR_BLUE
	// FLAG_COLOUR_WHITE_SLOW_CAR Slow car in area
	FLAG_COLOUR_WHITE_SLOW_CAR
	// FLAG_COLOUR_WHITE_FINAL_LAP Final Lap
	FLAG_COLOUR_WHITE_FINAL_LAP
	// FLAG_COLOUR_RED Huge collisions where one or more cars become wrecked and block the track
	FLAG_COLOUR_RED
	// FLAG_COLOUR_YELLOW Danger on the racing surface itself
	FLAG_COLOUR_YELLOW
	// FLAG_COLOUR_DOUBLE_YELLOW Danger that wholly or partly blocks the racing surface
	FLAG_COLOUR_DOUBLE_YELLOW
	// FLAG_COLOUR_BLACK_AND_WHITE Unsportsmanlike conduct
	FLAG_COLOUR_BLACK_AND_WHITE
	// FLAG_COLOUR_BLACK_ORANGE_CIRCLE Mechanical Failure
	FLAG_COLOUR_BLACK_ORANGE_CIRCLE
	// FLAG_COLOUR_BLACK Participant disqualified
	FLAG_COLOUR_BLACK
	// FLAG_COLOUR_CHEQUERED Chequered flag
	FLAG_COLOUR_CHEQUERED
	// FLAG_COLOUR_MAX mem safety
	FLAG_COLOUR_MAX
)

// (Type#6) Flag Reason (to be used with 'mHighestFlagReason')
type FlagReason int

const (
	FLAG_REASON_NONE FlagReason = iota
	FLAG_REASON_SOLO_CRASH
	FLAG_REASON_VEHICLE_CRASH
	FLAG_REASON_VEHICLE_OBSTRUCTION
	//-------------
	FLAG_REASON_MA
)

// PitMode (Type#7) (to be used with 'mPitMode')
type PitMode int

const (
	PIT_MODE_NONE PitMode = iota
	PIT_MODE_DRIVING_INTO_PITS
	PIT_MODE_IN_PIT
	PIT_MODE_DRIVING_OUT_OF_PITS
	PIT_MODE_IN_GARAGE
	PIT_MODE_DRIVING_OUT_OF_GARAGE
	//-------------
	PIT_MODE_MAX
)

// PitStopSchedule (Type#8) (to be used with 'mPitSchedule')
type PitStopSchedule int

const (
	// PIT_SCHEDULE_NONE Nothing scheduled
	PIT_SCHEDULE_NONE PitStopSchedule = iota
	// PIT_SCHEDULE_PLAYER_REQUESTED Used for standard pit sequence - requested by player
	PIT_SCHEDULE_PLAYER_REQUESTED
	// PIT_SCHEDULE_ENGINEER_REQUESTED Used for standard pit sequence - requested by engineer
	PIT_SCHEDULE_ENGINEER_REQUESTED
	// PIT_SCHEDULE_DAMAGE_REQUESTED Used for standard pit sequence - requested by engineer for damage
	PIT_SCHEDULE_DAMAGE_REQUESTED
	// PIT_SCHEDULE_MANDATORY Used for standard pit sequence - requested by engineer from career enforced lap number
	PIT_SCHEDULE_MANDATORY
	// PIT_SCHEDULE_DRIVE_THROUGH Used for drive-through penalty
	PIT_SCHEDULE_DRIVE_THROUGH
	// PIT_SCHEDULE_STOP_GO Used for stop-go penalty
	PIT_SCHEDULE_STOP_GO
	// PIT_SCHEDULE_PITSPOT_OCCUPIED Used for drive-through when pitspot is occupied
	PIT_SCHEDULE_PITSPOT_OCCUPIED
	// PIT_SCHEDULE_MAX for mem safety
	PIT_SCHEDULE_MAX
)

// CarFlags (Type#9) (to be used with 'mCarFlags')
type CarFlags int

const (
	CAR_HEADLIGHT      CarFlags = (1 << iota)
	CAR_ENGINE_ACTIVE  CarFlags = (1 << iota)
	CAR_ENGINE_WARNING CarFlags = (1 << iota)
	CAR_SPEED_LIMITER  CarFlags = (1 << iota)
	CAR_ABS            CarFlags = (1 << iota)
	CAR_HANDBRAKE      CarFlags = (1 << iota)
)

// TyreFlags (Type#10) (to be used with 'mTyreFlags')
type TyreFlags int

const (
	TYRE_ATTACHED     TyreFlags = (1 << iota)
	TYRE_INFLATED     TyreFlags = (1 << iota)
	TYRE_IS_ON_GROUND TyreFlags = (1 << iota)
)

// Terrain (Type#11) Materials (to be used with 'mTerrain')
type Terrain int

const (
	TERRAIN_ROAD Terrain = iota
	TERRAIN_LOW_GRIP_ROAD
	TERRAIN_BUMPY_ROAD1
	TERRAIN_BUMPY_ROAD2
	TERRAIN_BUMPY_ROAD3
	TERRAIN_MARBLES
	TERRAIN_GRASSY_BERMS
	TERRAIN_GRASS
	TERRAIN_GRAVEL
	TERRAIN_BUMPY_GRAVEL
	TERRAIN_RUMBLE_STRIPS
	TERRAIN_DRAINS
	TERRAIN_TYREWALLS
	TERRAIN_CEMENTWALLS
	TERRAIN_GUARDRAILS
	TERRAIN_SAND
	TERRAIN_BUMPY_SAND
	TERRAIN_DIRT
	TERRAIN_BUMPY_DIRT
	TERRAIN_DIRT_ROAD
	TERRAIN_BUMPY_DIRT_ROAD
	TERRAIN_PAVEMENT
	TERRAIN_DIRT_BANK
	TERRAIN_WOOD
	TERRAIN_DRY_VERGE
	TERRAIN_EXIT_RUMBLE_STRIPS
	TERRAIN_GRASSCRETE
	TERRAIN_LONG_GRASS
	TERRAIN_SLOPE_GRASS
	TERRAIN_COBBLES
	TERRAIN_SAND_ROAD
	TERRAIN_BAKED_CLAY
	TERRAIN_ASTROTURF
	TERRAIN_SNOWHALF
	TERRAIN_SNOWFULL
	TERRAIN_DAMAGED_ROAD1
	TERRAIN_TRAIN_TRACK_ROAD
	TERRAIN_BUMPYCOBBLES
	TERRAIN_ARIES_ONLY
	TERRAIN_ORION_ONLY
	TERRAIN_B1RUMBLES
	TERRAIN_B2RUMBLES
	TERRAIN_ROUGH_SAND_MEDIUM
	TERRAIN_ROUGH_SAND_HEAVY
	TERRAIN_SNOWWALLS
	TERRAIN_ICE_ROAD
	TERRAIN_RUNOFF_ROAD
	TERRAIN_ILLEGAL_STRIP
	TERRAIN_PAINT_CONCRETE
	TERRAIN_PAINT_CONCRETE_ILLEGAL
	TERRAIN_RALLY_TARMAC

	//-------------
	TERRAIN_MAX
)

// CrashState Damage (Type#12) (to be used with 'mCrashState')
type CrashState int

const (
	CRASH_DAMAGE_NONE CrashState = iota
	CRASH_DAMAGE_OFFTRACK
	CRASH_DAMAGE_LARGE_PROP
	CRASH_DAMAGE_SPINNING
	CRASH_DAMAGE_ROLLING
	//-------------
	CRASH_MAX
)

// ParticipantInfo struct (Type#13) (to be used with 'mParticipantInfo')
type ParticipantInfo struct {
	mIsActive           bool
	mName               [STRING_LENGTH_MAX]byte // [ string ]
	mWorldPosition      [VEC_MAX]float32        // [ UNITS = World Space  X  Y  Z ]
	mCurrentLapDistance float32                 // [ UNITS = Metres ]   [ RANGE = 0.0f->... ]    [ UNSET = 0.0f ]
	mRacePosition       uint16                  // [ RANGE = 1->... ]   [ UNSET = 0 ]
	mLapsCompleted      uint16                  // [ RANGE = 0->... ]   [ UNSET = 0 ]
	mCurrentLap         uint16                  // [ RANGE = 0->... ]   [ UNSET = 0 ]
	mCurrentSector      int16                   // [ RANGE = 0->... ]   [ UNSET = -1 ]
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
	mViewedParticipantIndex int16                                    // [ RANGE = 0->STORED_PARTICIPANTS_MAX ]   [ UNSET = -1 ]
	mNumParticipants        int16                                    // [ RANGE = 0->STORED_PARTICIPANTS_MAX ]   [ UNSET = -1 ]
	mParticipantInfo        [STORED_PARTICIPANTS_MAX]ParticipantInfo // [ struct (Type#13) ParticipantInfo struct ]

	// Unfiltered Input
	mUnfilteredThrottle float32 // [ RANGE = 0.0f->1.0f ]
	mUnfilteredBrake    float32 // [ RANGE = 0.0f->1.0f ]
	mUnfilteredSteering float32 // [ RANGE = -1.0f->1.0f ]
	mUnfilteredClutch   float32 // [ RANGE = 0.0f->1.0f ]

	// Vehicle information
	mCarName      [STRING_LENGTH_MAX]byte // [ string ]
	mCarClassName [STRING_LENGTH_MAX]byte // [ string ]

	// Event information
	mLapsInEvent    uint16                  // [ RANGE = 0->... ]   [ UNSET = 0 ]
	mTrackLocation  [STRING_LENGTH_MAX]byte // [ string ] - untranslated shortened English name
	mTrackVariation [STRING_LENGTH_MAX]byte // [ string ]- untranslated shortened English variation description
	mTrackLength    float32                 // [ UNITS = Metres ]   [ RANGE = 0.0f->... ]    [ UNSET = 0.0f ]

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
	mLastOpponentCollisionIndex     int16   // [ RANGE = 0->STORED_PARTICIPANTS_MAX ]   [ UNSET = -1 ]
	mLastOpponentCollisionMagnitude float32 // [ RANGE = 0.0f->... ]
	mBoostActive                    bool    // [ UNITS = boolean ]   [ RANGE = false->true ]   [ UNSET = false ]
	mBoostAmount                    float32 // [ RANGE = 0.0f->100.0f ]

	// Motion & Device Related
	mOrientation       [VEC_MAX]float32 // [ UNITS = Euler Angles ]
	mLocalVelocity     [VEC_MAX]float32 // [ UNITS = Metres per-second ]
	mWorldVelocity     [VEC_MAX]float32 // [ UNITS = Metres per-second ]
	mAngularVelocity   [VEC_MAX]float32 // [ UNITS = Radians per-second ]
	mLocalAcceleration [VEC_MAX]float32 // [ UNITS = Metres per-second ]
	mWorldAcceleration [VEC_MAX]float32 // [ UNITS = Metres per-second ]
	mExtentsCentre     [VEC_MAX]float32 // [ UNITS = Local Space  X  Y  Z ]

	// Wheels / Tyres
	mTyreFlags             [TYRE_MAX]uint16  // [ enum (Type#10) Tyre Flags ]
	mTerrain               [TYRE_MAX]uint16  // [ enum (Type#11) Terrain Materials ]
	mTyreY                 [TYRE_MAX]float32 // [ UNITS = Local Space  Y ]
	mTyreRPS               [TYRE_MAX]float32 // [ UNITS = Revolutions per second ]
	fTyreSlipSpeed         [TYRE_MAX]float32 // OBSOLETE, kept for backward compatibility only
	mTyreTemp              [TYRE_MAX]float32 // [ UNITS = Celsius ]   [ UNSET = 0.0f ]
	fTyreGrip              [TYRE_MAX]float32 // OBSOLETE, kept for backward compatibility only
	mTyreHeightAboveGround [TYRE_MAX]float32 // [ UNITS = Local Space  Y ]
	fTyreLateralStiffness  [TYRE_MAX]float32 // OBSOLETE, kept for backward compatibility only
	mTyreWear              [TYRE_MAX]float32 // [ RANGE = 0.0f->1.0f ]
	mBrakeDamage           [TYRE_MAX]float32 // [ RANGE = 0.0f->1.0f ]
	mSuspensionDamage      [TYRE_MAX]float32 // [ RANGE = 0.0f->1.0f ]
	mBrakeTempCelsius      [TYRE_MAX]float32 // [ UNITS = Celsius ]
	mTyreTreadTemp         [TYRE_MAX]float32 // [ UNITS = Kelvin ]
	mTyreLayerTemp         [TYRE_MAX]float32 // [ UNITS = Kelvin ]
	mTyreCarcassTemp       [TYRE_MAX]float32 // [ UNITS = Kelvin ]
	mTyreRimTemp           [TYRE_MAX]float32 // [ UNITS = Kelvin ]
	mTyreInternalAirTemp   [TYRE_MAX]float32 // [ UNITS = Kelvin ]

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
	mWheelLocalPositionY [TYRE_MAX]float32 // [ UNITS = Local Space  Y ]
	mSuspensionTravel    [TYRE_MAX]float32 // [ UNITS = meters ] [ RANGE 0.f =>... ]  [ UNSET =  0.0f ]
	mSuspensionVelocity  [TYRE_MAX]float32 // [ UNITS = Rate of change of pushrod deflection ] [ RANGE 0.f =>... ]  [ UNSET =  0.0f ]
	mAirPressure         [TYRE_MAX]float32 // [ UNITS = PSI ]  [ RANGE 0.f =>... ]  [ UNSET =  0.0f ]
	mEngineSpeed         float32           // [ UNITS = Rad/s ] [UNSET = 0.f ]
	mEngineTorque        float32           // [ UNITS = Newton Meters] [UNSET = 0.f ] [ RANGE = 0.0f->... ]
	mWings               [2]float32        // [ RANGE = 0.0f->1.0f ] [UNSET = 0.f ]
	mHandBrake           float32           // [ RANGE = 0.0f->1.0f ] [UNSET = 0.f ]

	// additional race variables
	mCurrentSector1Times [STORED_PARTICIPANTS_MAX]float32                 // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mCurrentSector2Times [STORED_PARTICIPANTS_MAX]float32                 // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mCurrentSector3Times [STORED_PARTICIPANTS_MAX]float32                 // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mFastestSector1Times [STORED_PARTICIPANTS_MAX]float32                 // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mFastestSector2Times [STORED_PARTICIPANTS_MAX]float32                 // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mFastestSector3Times [STORED_PARTICIPANTS_MAX]float32                 // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mFastestLapTimes     [STORED_PARTICIPANTS_MAX]float32                 // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mLastLapTimes        [STORED_PARTICIPANTS_MAX]float32                 // [ UNITS = seconds ]   [ RANGE = 0.0f->... ]   [ UNSET = -1.0f ]
	mLapsInvalidated     [STORED_PARTICIPANTS_MAX]bool                    // [ UNITS = boolean for all participants ]   [ RANGE = false->true ]   [ UNSET = false ]
	mRaceStates          [STORED_PARTICIPANTS_MAX]uint16                  // [ enum (Type#3) Race State ]
	mPitModes            [STORED_PARTICIPANTS_MAX]uint16                  // [ enum (Type#7)  Pit Mode ]
	mOrientations        [STORED_PARTICIPANTS_MAX][VEC_MAX]float32        // [ UNITS = Euler Angles ]
	mSpeeds              [STORED_PARTICIPANTS_MAX]float32                 // [ UNITS = Metres per-second ]   [ RANGE = 0.0f->... ]
	mCarNames            [STORED_PARTICIPANTS_MAX][STRING_LENGTH_MAX]byte // [ string ]
	mCarClassNames       [STORED_PARTICIPANTS_MAX][STRING_LENGTH_MAX]byte // [ string ]

	// additional race variables
	mEnforcedPitStopLap       int16                                         // [ UNITS = in which lap there will be a mandatory pitstop] [ RANGE = 0.0f->... ] [ UNSET = -1 ]
	mTranslatedTrackLocation  [STRING_LENGTH_MAX]byte                       // [ string ]
	mTranslatedTrackVariation [STRING_LENGTH_MAX]byte                       // [ string ]
	mBrakeBias                float32                                       // [ RANGE = 0.0f->1.0f... ]   [ UNSET = -1.0f ]
	mTurboBoostPressure       float32                                       //	 RANGE = 0.0f->1.0f... ]   [ UNSET = -1.0f ]
	mTyreCompound             [TYRE_MAX][TYRE_COMPOUND_NAME_LENGTH_MAX]byte // [ strings  ]
	mPitSchedules             [STORED_PARTICIPANTS_MAX]uint16               // [ enum (Type#7)  Pit Mode ]
	mHighestFlagColours       [STORED_PARTICIPANTS_MAX]uint16               // [ enum (Type#5) Flag Colour ]
	mHighestFlagReasons       [STORED_PARTICIPANTS_MAX]uint16               // [ enum (Type#6) Flag Reason ]
	mNationalities            [STORED_PARTICIPANTS_MAX]uint16               // [ nationality table , SP AND UNSET = 0 ] See nationalities.txt file for details
	mSnowDensity              float32                                       // [ UNITS = How much snow will fall ]   [ RANGE = 0.0f->1.0f ], this will be non zero only in Snow season, in other seasons whatever is falling from the sky is reported as rain

}
