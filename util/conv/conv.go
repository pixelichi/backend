package conv

func MbUint8ToBytes(mb uint8) int64{
  	return int64(mb) <<20
}