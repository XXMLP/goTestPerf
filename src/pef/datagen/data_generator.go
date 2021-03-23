package datagen

import (
	"errors"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/icrowley/fake"
	"github.com/satori/go.uuid"
)

var (
	fieldGenerator = map[string]func() interface{}{
		"ip": randomIPv4,
		"email": func() interface{} {
			return fake.EmailAddress()
		},
		"phone_number": randomPhoneNumber,
		"user_id": func() interface{} {
			return fake.Word()
		},
		"const_id": RandomConstID,
		"hardid":   RandomConstID,
	}

	constIDChars     = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	constIDCharsLen  = len(constIDChars)
	constIDTypes     = []string{"1", "2", "3"}
	constIDTypesLen  = len(constIDTypes)
	numbers          = []byte("0123456789")
	phonePrefixes    = []string{"130", "131", "132", "155", "156", "186", "133", "153", "189", "180", "137", "138", "139", "157", "158", "150", "188", "183", "182"}
	phonePrefixesLen = len(phonePrefixes)

	firstIPByte = byte(rand.Intn(254) + 1)
	lastIPByte  = byte(rand.Intn(256))

	dynamicFieldsMap = map[string]int{ // type: max
		"string": 32,
		"int":    100,
		"float":  100,
		"num":    10,
		"double": 100,
	}
)

func init() {
	if err := fake.SetLang("en"); err != nil {
		log.Fatalln(err)
	}

	rand.Seed(time.Now().UnixNano())
}

// GetGenerator is used to return a generator of specified field
func GetGenerator(field string) (string, func() interface{}, error) {
	if strings.ContainsRune(field, ':') { // name:type
		pairs := strings.SplitN(field, ":", 2)
		name, handlerType := pairs[0], pairs[1]
		handler, err := getDynamicGenerator(handlerType)
		return name, handler, err
	}

	handler, exists := fieldGenerator[field]
	if !exists {
		return "", nil, errors.New(field + " is not supported!")
	}
	return field, handler, nil
}

func getDynamicGenerator(handlerType string) (func() interface{}, error) {
	var dataType string
	var max int
	var handler func() interface{}

	if strings.ContainsRune(handlerType, '_') { // string_32, int_64
		pairs := strings.SplitN(handlerType, "_", 2)
		dataType = pairs[0]
		i, err := strconv.Atoi(pairs[1])
		if err != nil || i <= 0 {
			return nil, errors.New("invalid max size " + pairs[1])
		}
		max = i
	} else {
		dataType = handlerType
		if i, exists := dynamicFieldsMap[dataType]; exists {
			max = i
		}
	}

	switch dataType {
	case "string":
		handler = func() interface{} {
			return string(randomBytes(10))
		}
	case "int":
		handler = func() interface{} {
			return rand.Intn(max)
		}
	case "num":
		handler = func() interface{} {
			result := make([]byte, max)
			for i := 0; i < max; i++ {
				result[i] = numbers[rand.Intn(10)]
			}
				return string(result)
			}
	case "double":
		handler = func() interface{} {
			return rand.Float64() * float64(max)
		}
	case "float":
		handler = func() interface{} {
			return rand.Float32() * float32(max)
		}
	case "bool":
		handler = func() interface{} {
			return rand.Intn(2) != 0
		}
	case "uuid":
		handler = func() interface{} {
			return randomUUID()
		}
	default:
		return nil, errors.New("unsupported data type " + dataType)
	}

	return handler, nil
}

// randomUUID return a UUID v4 string
func randomUUID() string {
	return uuid.Must(uuid.NewV4()).String()
}

//randomBytes return a specified length of []byte
func randomBytes(size int) []byte {
	if size <= 0 {
		return []byte{}
	}

	result := make([]byte, size)

	for i := 0; i < size; i++ {
		result[i] = constIDChars[rand.Intn(constIDCharsLen)]
	}

	return result
}

func randomBoolean() interface{} {
	return rand.Intn(2) != 0
}

func randomStrBoolean() string {
	return strconv.FormatBool(rand.Intn(2) != 0)
}

func RandomConstID() interface{} {
	randomID := make([]byte, 31)
	for i := 0; i < 31; i++ {
		randomID[i] = constIDChars[rand.Intn(constIDCharsLen)]
	}

	return strings.Join(
		[]string{strconv.FormatInt(time.Now().Unix(), 16),
			string(randomID),
			constIDTypes[rand.Intn(constIDTypesLen)]},
		"")
}

func randomPhoneNumber() interface{} {
	result := make([]byte, 8)
	for i := 0; i < 8; i++ {
		result[i] = numbers[rand.Intn(10)]
	}

	return phonePrefixes[rand.Intn(phonePrefixesLen)] + string(result)
}

func randomIPv4() interface{} {
	ip := make([]byte, 4)
	ip[0] = byte(rand.Intn(254) + 1)
	for i := 1; i < 4; i++ {
		ip[i] = byte(rand.Intn(256))
	}

	return net.IP(ip).To4().String()
}
