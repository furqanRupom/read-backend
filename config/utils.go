package config
import(
	"strconv"
)
func parseBool(value string)(*bool,error){
	val,error:= strconv.ParseBool(value)
	return &val,error
}
func parseInt(value string)(*int,error){
	val,error := strconv.ParseInt(value,10,64)
	i:= int(val)
	return &i,error
}